package crypto

import (
	"github.com/urfave/cli/v2"
)

var (
	KEYGEN  = "keygen"
	ENCRYPT = "encrypt"
	DECRYPT = "decrypt"
)

var CryptoCommands = []*cli.Command{
	{
		Name:    KEYGEN,
		Usage:   "Generate an ecies key pair",
		Aliases: []string{"g", "kg"},
		Action: func(c *cli.Context) error {
			return eciesAction(c, KEYGEN)
		},
	},
	{
		Name:    ENCRYPT,
		Aliases: []string{"e", "en"},
		Usage:   "Encrypt a plaintext to a ciphertext using ecies",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "pk",
				Usage: "ecies public key",
			},
		},
		Action: func(c *cli.Context) error {
			return eciesAction(c, ENCRYPT)
		},
	},
	{
		Name:    DECRYPT,
		Usage:   "Decrypt a ciphertext to a plaintext using ecies",
		Aliases: []string{"d", "de"},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "sk",
				Usage: "ecies private key",
			},
		},
		Action: func(c *cli.Context) error {
			return eciesAction(c, DECRYPT)
		},
	},
}
