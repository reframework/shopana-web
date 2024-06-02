#!/bin/bash

# Path to the template .env file
TEMPLATE_ENV_FILE="template.env"

# Path to the output .env file
OUTPUT_ENV_FILE=".env"

# Create or clear the output .env file
> "$OUTPUT_ENV_FILE"

# Read the template .env file line by line
while IFS= read -r line || [[ -n "$line" ]]; do
    # Check if the line is not empty and not a comment
    if [[ -n "$line" && ! "$line" =~ ^# ]]; then
        # Extract the variable name from the line
        VAR_NAME=$(echo "$line" | cut -d '=' -f 1)
        # Check if the variable is set in the environment
        if [ -n "${!VAR_NAME}" ]; then
            # Write the variable and its value to the output .env file
            echo "$VAR_NAME=${!VAR_NAME}" >> "$OUTPUT_ENV_FILE"
        else
            # Write the original line to the output .env file
            echo "$line" >> "$OUTPUT_ENV_FILE"
        fi
    else
        # Write the original line (empty or comment) to the output .env file
        echo "$line" >> "$OUTPUT_ENV_FILE"
    fi
done < "$TEMPLATE_ENV_FILE"

echo "==============================="
echo ".env file created successfully."
echo "==============================="
cat .env
echo "==============================="
