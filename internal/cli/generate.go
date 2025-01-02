// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package cli

import (
	"fmt"
	"github.com/urfave/cli/v2"
	webPush "voo.su/pkg/push/web_push"
)

type GenerateProvider struct {
}

func Generate(ctx *cli.Context, app *GenerateProvider) error {
	privateKey, publicKey, err := webPush.GenerateVAPIDKeys()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(privateKey)
	fmt.Println(publicKey)

	return nil
}
