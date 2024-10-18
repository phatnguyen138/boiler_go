#!/bin/bash

# Ask for the new project name
read -p "Enter your new project name: " project_name

# Update the go.mod file
sed -i '' "s|module bolier_go|module $project_name|" go.mod

# Update all Go import paths, but exclude the current script
find . -type f -name "*.go" ! -path "./scripts/update_modules.sh" -exec sed -i '' "s|bolier_go|$project_name|g" {} +

echo "Project $project_name setup completed!"
