package resource

import "embed"

//go:embed "template"
var template embed.FS

func Template() embed.FS {
	return template
}
