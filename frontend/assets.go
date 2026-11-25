package frontend

import "embed"

//go:embed index.html style.css
var Assets embed.FS
