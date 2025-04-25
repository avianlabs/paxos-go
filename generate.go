package paxos

// Download the latest OpenAPI spec from the Paxos API
//go:generate curl -s https://developer.paxos.com/docs/paxos-v2.openapi.json -o paxos-v2.openapi.json

// Clean up old files
//go:generate rm -rf test/ api_*.go model_*.go .openapi-generator/FILES

// Generate new files
//go:generate openapi-generator generate -i paxos-v2.openapi.json -g go --package-name paxos --git-user-id avianlabs --git-repo-id paxos-go --additional-properties=disallowAdditionalPropertiesIfNotPresent=false -t ./templates

// Clean up generated files
//go:generate go mod tidy
//go:generate rm -rf docs/
//go:generate mv README.md paxos-README.md
//go:generate git ch README.md
