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
