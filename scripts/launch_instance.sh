#!/bin/bash

# Format Disk
sudo mkfs.ext4 /dev/xvdb

# Mount Disk
echo "creating folder for data"
mkdir -p /home/cynxhost/node

echo "mounting disk"
mount /dev/xvdb /home/cynxhost/node

echo "disk mounted"
cd /home/cynxhost/node

# Fetch metadata
TOKEN=$(curl -X PUT "http://169.254.169.254/latest/api/token" -H "X-aws-ec2-metadata-token-ttl-seconds: 21600")
AWS_INSTANCE_ID=$(curl -H "X-aws-ec2-metadata-token: $TOKEN" http://169.254.169.254/latest/meta-data/instance-id)
PUBLIC_IP=$(curl -H "X-aws-ec2-metadata-token: $TOKEN" http://169.254.169.254/latest/meta-data/public-ipv4)


# Fetch EBS Volume ID
VOLUME_ID=$(aws ec2 describe-volumes \
--filters "Name=attachment.instance-id,Values=$AWS_INSTANCE_ID" \
--query "Volumes[1].VolumeId" \
--output text)

# Send to backend
RESPONSE=$(curl -X POST {{LAUNCH_SUCCESS_CALLBACK_URL}} \
-H "Content-Type: application/json" \
-d '{
"aws_instance_id": "'"$AWS_INSTANCE_ID"'",
"public_ip": "'"$PUBLIC_IP"'",
"ebs_volume_id": "'"$VOLUME_ID"'"
"type": "{{LAUNCH_SUCCESS_TYPE}}"
}')

SCRIPT=$(echo "$RESPONSE" | jq -r '.data.script')
PERSISTENT_NODE_ID=$(echo "$RESPONSE" | jq -r '.data.persistent_node_id')

if [ "$SCRIPT" != "null" ] && [ -n "$SCRIPT" ]; then
echo "Received script: $SCRIPT"
# Execute the received script
eval "$SCRIPT"

# Send success response
echo "Sending success response"
curl -X POST {{SETUP_SUCCESS_CALLBACK_URL}} \
-H "Content-Type: application/json" \
-d '{
  "persistent_node_id": "'"$PERSISTENT_NODE_ID"'",
  "type": "{{SETUP_SUCCESS_TYPE}}"
}'

else
echo "No script found in response."
fi