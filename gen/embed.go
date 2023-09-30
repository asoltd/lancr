package gen

import (
	"embed"
)

//go:embed openapiv2/*
var OpenAPI embed.FS
