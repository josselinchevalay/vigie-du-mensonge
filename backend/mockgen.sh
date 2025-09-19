#!/bin/bash

# Usage:
# ./mockgen.sh <destination> <package_name> <source>

DESTINATION="$1"
PACKAGE_NAME="$2"
SOURCE="$3"

if [ -z "$DESTINATION" ] || [ -z "$PACKAGE_NAME" ] || [ -z "$SOURCE" ]; then
  echo "Usage: $0 <destination> <package_name> <source>"
  exit 1
fi

go run github.com/golang/mock/mockgen --build_flags=--mod=vendor \
  -destination "$DESTINATION" \
  -package "$PACKAGE_NAME" \
  -source "$SOURCE"