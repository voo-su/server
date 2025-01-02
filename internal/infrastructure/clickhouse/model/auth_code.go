// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package model

type AuthCode struct {
	Email        string
	Code         string
	Token        string
	ErrorMessage string
}
