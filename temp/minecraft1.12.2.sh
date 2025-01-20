#!/bin/bash

# Variables
JAVA_SWITCH_SCRIPT="/home/cynxhost/scripts/switch_java_to_version.sh"
MC_SERVER_DIR="/home/cynxhost/node"
MC_VERSION="1.12.2"
MC_SERVER_JAR="/home/cynxhost/node/server.jar"
MC_SERVER_URL="{{.VARIABLE_URL}}"
EULA_FILE="$MC_SERVER_DIR/eula.txt"
PROPERTIES_FILE="$MC_SERVER_DIR/server.properties"

# Step 1: Switch to the appropriate Java version
echo "Switching Java version..."
if [ -f "$JAVA_SWITCH_SCRIPT" ]; then
    bash "$JAVA_SWITCH_SCRIPT"
else
    echo "Java switch script not found: $JAVA_SWITCH_SCRIPT"
    exit 1
fi

# Step 2: Create the Minecraft server directory if it doesn't exist
echo "Creating Minecraft server directory at $MC_SERVER_DIR..."
mkdir -p "$MC_SERVER_DIR"
cd "$MC_SERVER_DIR" || { echo "Failed to change directory to $MC_SERVER_DIR"; exit 1; }

# Step 3: Download the Minecraft server JAR file
echo "Downloading Minecraft server JAR..."
wget -O "$MC_SERVER_JAR" "$MC_SERVER_URL"

if [ $? -ne 0 ]; then
    echo "Failed to download Minecraft server JAR."
    exit 1
fi

# Step 4: Accept the EULA
echo "Accepting EULA..."
echo "eula=true" > "$EULA_FILE"

# Step 5: Create or update server.properties
echo "Setting up server.properties..."
if [ ! -f "$PROPERTIES_FILE" ]; then
    echo "server.properties not found. Creating a new one..."
    cat <<EOL > "$PROPERTIES_FILE"
#Minecraft server properties
online-mode=false
EOL
else
    echo "Updating server.properties..."
    if grep -q "online-mode" "$PROPERTIES_FILE"; then
        sed -i 's/online-mode=.*/online-mode=false/' "$PROPERTIES_FILE"
    else
        echo "online-mode=false" >> "$PROPERTIES_FILE"
    fi
fi

# Create log files
AGENT_DIR="/home/cynxhost/cynxhost-agent"
AGENT_OUTPUT_FILE="$AGENT_DIR/output.log"
SESSION_NAME="cynxhost-node"

mkdir $AGENT_DIR
touch $AGENT_OUTPUT_FILE

# Start the Minecraft server
echo "Starting Minecraft server..."
tmux new-session -d -s $SESSION_NAME "java -Xmx1024M -Xms1024M -jar '$MC_SERVER_JAR' nogui > '$AGENT_OUTPUT_FILE' 2>&1"