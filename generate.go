package paxos

// Check for required tools
//go:generate sh -c "command -v curl >/dev/null 2>&1 || { echo 'Error: curl is required but not installed' >&2; exit 1; }"
//go:generate sh -c "command -v jq >/dev/null 2>&1 || { echo 'Error: jq is required but not installed' >&2; exit 1; }"
//go:generate sh -c "command -v openapi-generator >/dev/null 2>&1 || { echo 'Error: openapi-generator is required but not installed' >&2; exit 1; }"

// Download the latest OpenAPI spec from the Paxos API
//go:generate sh -c "curl -s https://developer.paxos.com/docs/paxos-v2.openapi.json > paxos-v2.openapi.json.tmp"

// Patch the spec to add the XLAYER enum value to CryptoNetwork
//go:generate sh -c "jq '.components.schemas.CryptoNetwork.enum += [\"XLAYER\"]' paxos-v2.openapi.json.tmp > paxos-v2.openapi.json"

// Clean up temporary file
//go:generate rm -f paxos-v2.openapi.json.tmp

// Clean up old files
//go:generate rm -rf test/ api_*.go model_*.go .openapi-generator/FILES

// Generate new files
//go:generate openapi-generator generate -i paxos-v2.openapi.json -g go --package-name paxos --git-user-id avianlabs --git-repo-id paxos-go --additional-properties=disallowAdditionalPropertiesIfNotPresent=false -t ./templates

// Clean up generated files
//go:generate go mod tidy
//go:generate rm -rf docs/
//go:generate mv README.md paxos-README.md
//go:generate git ch README.md
