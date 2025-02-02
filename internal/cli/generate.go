package cli

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	webPush "voo.su/pkg/push/web_push"
)

type GenerateProvider struct {
}

func Generate(ctx *cli.Context, app *GenerateProvider) error {
	privateKey, publicKey, err := webPush.GenerateVAPIDKeys()
	if err != nil {
		log.Println(err)
	}

	fmt.Println(privateKey)
	fmt.Println(publicKey)

	return nil
}
