// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package utils

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
)

type ImageMeta struct {
	Width  int
	Height int
}

func ReadImageMeta(r io.Reader) *ImageMeta {
	c, _, _ := image.DecodeConfig(r)

	return &ImageMeta{c.Width, c.Height}
}
