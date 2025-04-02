package templates

import (
	"embed"
	"io/fs"
)

//go:embed *
var FS embed.FS

//go:embed static/*
var static embed.FS
var StaticFS, _ = fs.Sub(static, "static")
