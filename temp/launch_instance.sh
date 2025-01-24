#!/bin/bash

DISK_DEVICE="/dev/xvdb"
MOUNT_DIR="/home/cynxhost/node"

# Format Disk
echo "Formatting disk $DISK_DEVICE"
sudo mkfs.ext4 "$DISK_DEVICE"

if [ $? -ne 0 ]; then
    echo "Failed to format the disk. Trying with nvme1n1..."
    DISK_DEVICE="/dev/nvme1n1"

    echo "Formatting disk $DISK_DEVICE"
    sudo mkfs.ext4 "$DISK_DEVICE"

    if [ $? -ne 0 ]; then
        echo "Failed to format the disk."
        exit 1
    fi
fi

# Step 1: Create the folder for the data
echo "Creating folder for data at $MOUNT_DIR..."
mkdir -p "$MOUNT_DIR"

# Step 2: Mount the disk
echo "Mounting disk $DISK_DEVICE to $MOUNT_DIR..."
mount "$DISK_DEVICE" "$MOUNT_DIR"

if [ $? -ne 0 ]; then
    echo "Failed to mount the disk."
    exit 1
fi

echo "Disk mounted successfully."

# Step 3: Set permissions to public (read/write for everyone)
echo "Setting permissions to public..."
chmod -R 777 "$MOUNT_DIR"

if [ $? -ne 0 ]; then
    echo "Failed to set permissions."
    exit 1
fi

echo "Permissions set to public successfully."

# Step 4: Change directory to the mounted folder
cd "$MOUNT_DIR" || { echo "Failed to change directory to $MOUNT_DIR"; exit 1; }

# Confirmation
echo "Disk setup is complete. Current directory: $(pwd)"

# Fetch metadata
TOKEN=$(curl -X PUT "http://169.254.169.254/latest/api/token" -H "X-aws-ec2-metadata-token-ttl-seconds: 21600")
AWS_INSTANCE_ID=$(curl -H "X-aws-ec2-metadata-token: $TOKEN" http://169.254.169.254/latest/meta-data/instance-id)
PUBLIC_IP=$(curl -H "X-aws-ec2-metadata-token: $TOKEN" http://169.254.169.254/latest/meta-data/public-ipv4)
PRIVATE_IP=$(curl -H "X-aws-ec2-metadata-token: $TOKEN" http://169.254.169.254/latest/meta-data/local-ipv4)

# Write .env for cynxhost-agent
cd /home/cynxhost/cynxhost-agent || { echo "Failed to change directory to /home/cynxhost/cynxhost-agent"; exit 1; }
touch .env

echo "Writing to .env file..."
{
  echo "JWT_SECRET=\"{{.JWT_SECRET}}\""
  echo "APP_PUBLIC_IP=\"$PUBLIC_IP\""
  echo "APP_PRIVATE_IP=\"$PRIVATE_IP\""
  echo "CENTRAL_PRIVATE_IP=\"{{.CENTRAL_PRIVATE_IP}}\""
  echo "CENTRAL_PUBLIC_IP=\"{{.CENTRAL_PUBLIC_IP}}\""
  echo "CENTRAL_PORT=\"{{.CENTRAL_PORT}}\""
} > .env

# Fetching cynxhostagent from s3
echo "Fetching cynxhostagent from s3..."
aws s3 cp s3://cynxhost/cynxhostagent/cynxhostagent . --region ap-southeast-1
aws s3 cp s3://cynxhost/{{.CONFIG_PATH}} ./config.json --region ap-southeast-1

sudo chmod +x cynxhostagent

# Restarting cynxhost agent service
echo "Restarting cynxhost agent service..."
sudo systemctl restart cynxhost-agent.service

# Fetch EBS Volume ID
VOLUME_ID=$(aws ec2 describe-volumes \
--filters "Name=attachment.instance-id,Values=$AWS_INSTANCE_ID" \
--query "Volumes[1].VolumeId" \
--output text)

# Send to backend
RESPONSE=$(curl -X POST {{.LAUNCH_SUCCESS_CALLBACK_URL}} \
-H "Content-Type: application/json" \
-d '{
"aws_instance_id": "'"$AWS_INSTANCE_ID"'",
"public_ip": "'"$PUBLIC_IP"'",
"ebs_volume_id": "'"$VOLUME_ID"'",
"type": "{{.LAUNCH_SUCCESS_TYPE}}"
}')

SCRIPT=$(echo "$RESPONSE" | jq -r '.data.Script')
PERSISTENT_NODE_ID=$(echo "$RESPONSE" | jq '.data.PersistentNodeId | select(. != null) | tonumber')

cd $MOUNT_DIR

if [ "$SCRIPT" != "null" ] && [ -n "$SCRIPT" ]; then
  echo "Received base64 encoded script"

  echo "Pulling base image..."
  aws ecr get-login-password --region ap-southeast-1 | docker login --username AWS --password-stdin 242201306378.dkr.ecr.ap-southeast-1.amazonaws.com
  docker pull 242201306378.dkr.ecr.ap-southeast-1.amazonaws.com/cynxhost/node:base

  # Make docker container with dockerfile
  echo "Creating Dockerfile..."
  cat > Dockerfile << EOF
FROM 242201306378.dkr.ecr.ap-southeast-1.amazonaws.com/cynxhost/node:cynxhost-base-image
USER cynxhost
WORKDIR $MOUNT_DIR
COPY script.sh $MOUNT_DIR/script.sh
USER root
RUN chmod +x $MOUNT_DIR/script.sh
USER cynxhost
CMD ["sh", "-c", "$MOUNT_DIR/script.sh && tail -f /dev/null"]
EOF

  # Decode the base64 encoded script and save it to script.sh
  echo "Decoding and saving the script..."
  echo "$SCRIPT" | base64 --decode > $MOUNT_DIR/script.sh
  
  # Build Docker image
  echo "Building Docker image..."
  docker build -t cynxhost-container .

  # Run Docker container
  echo "Running Docker container..."
  docker run -d --name cynxhost-container cynxhost-container

  # Send success response
  echo "Sending success response"
  curl -X POST {{.SETUP_SUCCESS_CALLBACK_URL}} \
  -H "Content-Type: application/json" \
  -d '{
    "persistent_node_id": '$PERSISTENT_NODE_ID',
    "type": "{{.SETUP_SUCCESS_TYPE}}"
  }'

else
echo "No script found in response."
fi

# Setup DNS
echo "Setting up DNS..."
DOMAIN="{{.DOMAIN}}"
NGINX_CONFIG="/etc/nginx/sites-available/${DOMAIN}"

# Create an NGINX configuration
echo "server {
    # HTTP traffic (port 80)
    listen 80;
    server_name $DOMAIN;

    location / {
        proxy_pass http://127.0.0.1:3001;  # Replace with your HTTP service's port
        proxy_http_version 1.1;
        proxy_set_header Upgrade \$http_upgrade;
        proxy_set_header Connection 'upgrade';
        proxy_set_header Host \$host;
        proxy_cache_bypass \$http_upgrade;
    }
}"> $NGINX_CONFIG

# Enable the configuration and restart NGINX
sudo ln -s $NGINX_CONFIG /etc/nginx/sites-enabled/
sudo nginx -t
sudo systemctl restart nginx

# Open necessary firewall ports
echo "Opening necessary ports in the firewall..."
sudo ufw allow 80         # HTTP
sudo ufw allow 8000       # WebSocket
sudo ufw allow 25565      # Minecraft
sudo ufw reload

# Final message
echo "DNS setup and services configuration complete. Access your services at:"
echo "HTTP: http://$DOMAIN"
echo "WebSocket: ws://$DOMAIN:8000"
echo "Minecraft: $DOMAIN:25565"