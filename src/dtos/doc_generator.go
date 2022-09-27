package dtos

//go:generate -command OAPI go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen -generate

//go:generate OAPI types --output-config --old-config-style --output-config -o dtos.gen.go -package dtos .dtos/res/swagger.yml
