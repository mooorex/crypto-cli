package crypto

import (
	"errors"
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/mooorex/crypto-cli/util"
	"github.com/urfave/cli/v2"
)

//todo: check the number of arguments

func eciesAction(c *cli.Context, action string) error {
	arg := c.Args().First()
	var (
		res string
		err error
	)
	if action == KEYGEN {
		res, err = eciesKeyGenAction(c)
	} else if action == ENCRYPT {
		res, err = eciesEncryptAction(c)
	} else if action == DECRYPT {
		res, err = eciesDecryptAction(c, arg)
	} else {
		err = errors.New("Unknown action")
	}
	if err != nil {
		return err
	}

	fmt.Printf("ecies %s result:\n%s\n", action, res)
	return nil
}

func eciesKeyGenAction(c *cli.Context) (string, error) {
	pk, sk, err := util.KeyGen()
	return "Public Key: " + pk + "\nPrivate Key: " + sk, err
}

func eciesEncryptAction(c *cli.Context) (string, error) {
	pk, err := getPublicKey(c)
	if err != nil {
		return "", err
	}

	prompt := promptui.Prompt{
		Label: "The message to be encrypted",
		Mask:  '*',
	}
	msg, err := prompt.Run()
	if err != nil {
		return "", err
	}

	ct, err := util.Encypt(pk, msg)
	return "Ciphertext: " + ct, err
}

func eciesDecryptAction(c *cli.Context, ct string) (string, error) {
	sk, err := getPrivateKey(c)
	if err != nil {
		return "", err
	}

	if len(ct) == 0 {
		prompt := promptui.Prompt{
			Label: "The ciphertext to be decrypted",
		}
		ct, err = prompt.Run()
		if err != nil {
			return "", err
		}
	}

	msg, err := util.Decypt(sk, ct)
	return "Message: " + msg, err
}

func getPublicKey(c *cli.Context) (string, error) {
	var key string
	if c.String("pk") != "" {
		key = c.String("pk")
	} else {
		prompt := promptui.Prompt{
			Label: "Public Key",
		}
		result, err := prompt.Run()
		if err != nil {
			return "", err
		}
		key = result
	}
	//todo: check the length
	return key, nil
}

func getPrivateKey(c *cli.Context) (string, error) {
	var key string
	if c.String("sk") != "" {
		key = c.String("sk")
	} else {
		prompt := promptui.Prompt{
			Label: "Private Key",
			Mask:  '*',
		}
		result, err := prompt.Run()
		if err != nil {
			return "", err
		}
		key = result
	}
	//todo: check the length
	return key, nil
}
