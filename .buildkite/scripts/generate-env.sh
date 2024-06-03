#!/bin/bash

TEMPLATE_FILE="template.env"
ENV_FILE=".env"

# check if .env file exists
if [ ! -f "$ENV_FILE" ]; then
    echo "Error: .env file not found!"
    exit 1
fi

# Renave .env to template.env
mv "$ENV_FILE" "$TEMPLATE_FILE"

# Create or clear the output .env file
> "$ENV_FILE"

# Read the template .env file line by line
while IFS= read -r line || [[ -n "$line" ]]; do
    # Check if the line is not empty and not a comment
    if [[ -n "$line" && ! "$line" =~ ^# ]]; then
        # Extract the variable name from the line
        VAR_NAME=$(echo "$line" | cut -d '=' -f 1)
        # Check if the variable is set in the environment
        if [ -n "${!VAR_NAME}" ]; then
            # Write the variable and its value to the output .env file
            echo "$VAR_NAME=${!VAR_NAME}" >> "$ENV_FILE"
        else
            # Write the original line to the output .env file
            echo "$line" >> "$ENV_FILE"
        fi
    else
        # Write the original line (empty or comment) to the output .env file
        echo "$line" >> "$ENV_FILE"
    fi
done < "$TEMPLATE_FILE"

echo "==============================="
echo ".env file created successfully."
echo "==============================="
cat .env
echo "==============================="
