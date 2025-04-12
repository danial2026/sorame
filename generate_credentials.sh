#!/bin/bash

# Check if .env file exists, if not copy from template
if [ ! -f .env ]; then
    echo "Creating .env file from template..."
    cp .env.template .env
fi

# Check if Python is installed
if ! command -v python3 &> /dev/null; then
    echo "Python 3 is required but not installed. Please install Python 3 and try again."
    exit 1
fi

# Run the Python script
python3 generate_credentials.py

# Source the .env file
source .env

# Create docker network if it doesn't exist
docker network create $NETWORK_NAME