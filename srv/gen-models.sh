#!/usr/bin/env sh

set -e

OUT_DIR=$(pwd)/internal/apimodel

# Clean all models before generating new ones
rm -rf "${OUT_DIR}"
mkdir -p "${OUT_DIR}"

TEMPLATE_API_PATH="$(pwd)/template/api.yaml"
API_PATH="$(pwd)/.api.yaml"
cp "${TEMPLATE_API_PATH}" "${API_PATH}"

# replace "{{.SrvUrl}}" with sed
sed -i '' 's/{{.SrvUrl}}/http:\/\/localhost:8080/g' "${API_PATH}"

echo "About to generate models from ${API_PATH}"
openapi-generator generate -i "${API_PATH}" \
  --output "${OUT_DIR}" \
  --generator-name go \
  --global-property models,modelDocs=false \
  --additional-properties="packageName=api"

# go fmt
echo "About to fmt"
go fmt ./...