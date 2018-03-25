package main

import (
	"time"
	"fmt"
	"log"
	"github.com/namsral/flag"
	"gopkg.in/virgil.v5/sdk"
	"gopkg.in/virgilsecurity/virgil-crypto-go.v5"
)

// crypto vars
var (
	crypto      = virgil_crypto_go.NewVirgilCrypto()
	cardCrypto  = virgil_crypto_go.NewVirgilCardCrypto()
	tokenSigner = virgil_crypto_go.NewVirgilAccessTokenSigner()
)

// config vars
var (
	config             string
	privateKeyStr      string
	privateKeyPassword string
	appId              string
	appPubKeyId        string
	identity           string
	searchCard         string
	ttl                time.Duration
)

func main() {

	flag.StringVar(&config, "config", "", "Config file with variables - optional. Can parse both variables from config and CLI")
	flag.StringVar(&privateKeyStr, "privateKeyStr", "", "Private Api Key generated at dashboard.virgilsecurity.com. Required")
	flag.StringVar(&privateKeyPassword, "privateKeyPassword", "", "Private key password - null by default")
	flag.StringVar(&appId, "appId", "", "APP_ID in virgil dashboard. Required")
	flag.StringVar(&appPubKeyId, "appPubKeyId", "", "API_KEY_ID in virgil dashboard. Required")
	flag.StringVar(&identity, "identity", "", "Identity. Required")
	flag.DurationVar(&ttl, "ttl", time.Hour*24, "Time to live of token. Max 24h")
	flag.StringVar(&searchCard, "searchCard", "example-card", "card to search - just for verify")
	flag.Parse()

	privateKeyByte := []byte(privateKeyStr)

	// verify required variables are set
	requiredParams := [] string{
		"privateKeyStr",
		"appPubKeyId",
		"appId",
		"identity",
	}

	flagset := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) { flagset[f.Name] = true })
	for _, param := range requiredParams {
		if ! flagset[param] {
			flag.Usage()
			log.Fatalln(param, "required. Please specify all required variables")
		}
	}

	// import a private key
	privateKey, err := crypto.ImportPrivateKey(privateKeyByte, privateKeyPassword)
	if err != nil {
		//handle error
		fmt.Println("PrivateKeyErr:", err)
	}

	// setup JWT generator
	jwtGenerator := sdk.NewJwtGenerator(privateKey, appPubKeyId, tokenSigner, appId, ttl)

	token, err := jwtGenerator.GenerateToken(identity, nil)
	fmt.Println("#token:", token, "\n")
	if err != nil {
		log.Fatal("GenerateTokenErr:", err)
	}

	// curl cli example
	SearchCard := fmt.Sprintf("#SearchCard example:\n curl -w '@curl-format.txt' -X POST -H 'Authorization: Virgil %s' -d '{\"identity\": \"%s\"}' https://api.virgilsecurity.com/card/v5/actions/search\n\n", token, identity)
	fmt.Printf(SearchCard)
	GetCard := fmt.Sprintf("#getCard example:\n curl -w '@curl-format.txt' -X GET -H 'Authorization: Virgil %s' https://api.virgilsecurity.com/card/v5/%s\n", token, searchCard)
	fmt.Printf(GetCard)
}
