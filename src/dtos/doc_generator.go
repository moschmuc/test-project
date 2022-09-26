package dtos

//go:generate -command OAPI go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen -generate

//go:generate OAPI types -o dtos.gen.go -package dtos ./res/swagger.yml
