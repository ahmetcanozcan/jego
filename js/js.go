package js

import _ "embed"

var (
	//go:embed export.js
	ExportJS string
)

type Object map[string]any
