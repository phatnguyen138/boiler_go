#!/bin/bash

# Check if a new name argument is provided
if [ -z "$1" ]; then
  echo "Error: Please provide a new name as an argument."
  exit 1
fi

# Export variables with the script argument and the original name
export NEW="$1"
export CUR="bolier_go" # Example: github.com/user/old-lame-name

# Update module name
go mod edit -module "$NEW"

# Update Go files using sed (alternative to perl)
find . -type f -name "*.go" -exec sed -i "s/$CUR/$NEW/g" {} \;

echo "Successfully updated module name and Go files."
