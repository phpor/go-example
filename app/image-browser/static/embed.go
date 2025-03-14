package static

import "embed"

//go:embed js/*.js
var JS embed.FS

//go:embed html/*.html
var HTML embed.FS
