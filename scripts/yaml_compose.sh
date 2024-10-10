#!/bin/bash

FOLDER=${1:-.}
OUTPUT_FILE=${2:-output.yaml}

echo "Cleaning up spec/gen"
rm $FOLDER/gen/*.gen.yaml > /dev/null 2>&1

echo "Generating openapi spec: $FOLDER/gen/faro.gen.yaml"
cp "$FOLDER/entrypoint.yaml" $FOLDER/gen/faro.gen.yaml > /dev/null 2>&1
find "$FOLDER" -type f -name "*.yaml" | while read -r file; do
    YQ_OUTPUT=$(yq --yaml-output --slurp '.[0] * .[1]' $FOLDER/gen/faro.gen.yaml "$file" 2>&1)
    if [ $? -ne 0 ]; then
        echo "Error processing file: $file"
        echo "$YQ_OUTPUT"
        continue
    fi
    echo "$YQ_OUTPUT" > $FOLDER/gen/faro.gen.yaml
done
sed -i '' '1s/^/# Code auto-generated. DO NOT EDIT.\n/' $FOLDER/gen/faro.gen.yaml

echo "Composed openapi spec: $OUTPUT_FILE"
