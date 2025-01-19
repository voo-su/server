package voo_su

import "embed"

//go:embed migrations
var migration embed.FS

//go:embed internal/locale
var locale embed.FS

func Migration() embed.FS {
	return migration
}

func Locale() embed.FS {
	return locale
}
