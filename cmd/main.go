package main

import (
	"fmt"
	"log"
	"time"

	"github.com/fcjr/frothy"
)

func main() {
	uri, err := frothy.DecodeFromScreen()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Raw URI: %s\n", uri)

	secret, err := frothy.ParseOTPSecretFromURI(uri)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Secret: %s\n", secret.Secret)

	totp, err := frothy.NewTOTP(secret.Secret)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Code: %s\n", totp.Code)
	fmt.Printf("ExpiresAt: %s", totp.ExpiresAt.Format(time.RFC850))
}
