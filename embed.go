package voo_su

import "embed"

//go:embed migrations
var migration embed.FS

func Migration() embed.FS {
	return migration
}
