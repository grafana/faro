#!/bin/bash

FILE_PATH=${1:-api.gen.go}

# Patch type alias into embedded struct
sed -i '' 's/type Traces = ptrace.Traces/type Traces struct {\n\tptrace.Traces\n}/' "$FILE_PATH"
