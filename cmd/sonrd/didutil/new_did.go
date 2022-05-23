package didutil

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonr-io/sonr/pkg/did"
	"github.com/sonr-io/sonr/pkg/did/ssi"
	"github.com/spf13/viper"
)

var (
	CONFIG_PATH = viper.GetString("LOCAL_CONFIG_DIR")
)

func surveyNewDid() error {
	// Webauthn confirmation
	webauthn := false
	prompt := &survey.Confirm{
		Message: "Would you like to generate a Webauthn Credential in your browser?",
	}
	survey.AskOne(prompt, &webauthn)

	if webauthn {
		// Open browser and serve index.html
	} else {
		pubKey := ""
		// Present public key types
		selectOpt := &survey.Select{
			Message: "What type of key would you like to generate",
			Options: []string{
				"Ed25519VerificationKey2018",
				"EcdsaSecp256k1VerificationKey2019",
				"RsaVerificationKey2018",
			},
			Default: "Ed25519VerificationKey2018",
		}
		// perform the questions
		survey.AskOne(selectOpt, &pubKey, nil)

		// Parse Account Addr
		addr := genAccAddressDid()
		didRoot, err := did.ParseDID(addr)
		if err != nil {
			return err
		}

		var vm *did.VerificationMethod

		// Parse DID Controller
		didController, err := did.ParseDID(fmt.Sprintf("%s#test", didRoot))
		if err != nil {
			return err
		}

		// Switch on pubKey
		switch pubKey {
		case "EcdsaSecp256k1VerificationKey2019":
			// Generate EcdsaSecp256k1VerificationKey2019
			priv := secp256k1.GenPrivKey()
			pub := priv.PubKey()
			vm, _ = did.NewVerificationMethod(*didRoot, ssi.ECDSASECP256K1VerificationKey2019, *didController, pub)
			// did.ParseDID(input string)
		case "RsaVerificationKey2018":
			// Generate RsaVerificationKey2018
			priv, err := rsa.GenerateKey(rand.Reader, 2048)
			if err != nil {
				return err
			}
			pub := priv.PublicKey
			vm, _ = did.NewVerificationMethod(*didRoot, ssi.RSAVerificationKey2018, *didController, pub)
			// did.ParseDID(input string)
		default:
			pub := ed25519.GenPrivKey().PubKey()
			vm, _ = did.NewVerificationMethod(*didRoot, ssi.ED25519VerificationKey2018, *didController, pub)
			// did.ParseDID(input string)
		}

		ctx, err := ssi.ParseURI("https://www.w3.org/ns/did/v1")
		if err != nil {
			return err
		}

		// Document properties
		doc := &did.DocumentImpl{
			Context: []ssi.URI{*ctx},
			ID:      *didRoot,
		}
		doc.AddController(*didController)
		doc.AddAuthenticationMethod(vm)
		buf, err := doc.MarshalJSON()
		if err != nil {
			return err
		}

		saveFile := false
		prompt := &survey.Confirm{
			Message: "Would you like to save the JSON output to disk?",
		}
		survey.AskOne(prompt, &saveFile)

		if saveFile {
			path := filepath.Join(CONFIG_PATH, "testutil", "sample", "did.json")
			ioutil.WriteFile(path, buf, 0644)
			fmt.Printf("Saved DID JSON Document to: %s\n", path)
		} else {
			fmt.Println(string(buf))
		}
	}
	return nil
}

// AccAddress returns a sample account address
func genAccAddressDid() string {
	pk := ed25519.GenPrivKey().PubKey()
	addr := pk.Address()
	addrStr := sdk.AccAddress(addr).String()
	addrStr = strings.TrimLeft(addrStr, "snr")
	addrStr = strings.TrimRight(addrStr, "cosmos")

	d, err := did.New(addrStr)
	if err != nil {
		fmt.Printf("Error: failed to create did: %s", err.Error())
		panic(err) // this should never happen
	}
	return d.String()
}
