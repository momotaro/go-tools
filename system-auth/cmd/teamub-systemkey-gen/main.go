package main

import (
	"fmt"
	"log"
	"os"

	auth "github.com/momotaro/go-tools/system-auth"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()

	app.Action = func(c *cli.Context) error {
		name := os.Getenv("CLIENT_SYSTEM_NAME")
		secret := os.Getenv("SECRET")

		auth := auth.NewAuth(name, secret)

		token, err := auth.CreateToken()
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(token)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}
