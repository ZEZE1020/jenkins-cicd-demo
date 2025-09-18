#!/bin/bash

# Simple deployment script for Docker Compose

set -e

echo "Starting deployment with Docker Compose..."

# Navigate to the deployment directory
cd "$(dirname "$0")/../deployments"

# Pull the latest image
echo "Pulling latest Docker image..."
docker-compose pull

# Stop existing containers
echo "Stopping existing containers..."
docker-compose down

# Start new containers
echo "Starting new containers..."
docker-compose up -d

# Wait for health check
echo "Waiting for health check..."
sleep 15

# Check if the application is healthy
if curl -f http://localhost:3000/health > /dev/null 2>&1; then
    echo "Deployment successful! Application is healthy."
    echo "Application available at: http://localhost:3000"
    echo "Dashboard available at: http://localhost:3000/dashboard"
    echo "Health check: http://localhost:3000/health"
    echo "Metrics: http://localhost:3000/metrics"
else
    echo "Deployment failed! Application health check failed."
    echo "Checking logs..."
    docker-compose logs app
    exit 1
fi