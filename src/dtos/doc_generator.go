package dtos

//go:generate -command OAPI go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen -generate

// Generate a dtos.gen.go and print the contents for a config yaml file to stdout, which could be used for
// "OAPI --config <config.yaml> ./res/swagger.yml" later on instead of setting all command line parameters.
// See https://github.com/deepmap/oapi-codegen/releases for details
//go:generate OAPI types --output-config --old-config-style --output-config -o dtos.gen.go -package dtos ./res/swagger.yml
