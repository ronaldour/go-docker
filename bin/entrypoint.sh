#!/bin/ash

# Sync with Parameter Store
eval $(aws-env)
# Run Go app
echo "running"
/app/app