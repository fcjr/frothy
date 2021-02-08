package frothy

import (
	"fmt"

	"github.com/sqweek/dialog"
)

type App struct {
	secretStore *SecretStore
	secrets     []*OTPSecret
}

func NewApp() (*App, error) {
	ss, err := OpenSecretStore()
	if err != nil {
		return nil, err
	}
	secrets, err := ss.GetSecrets()
	if err != nil {
		return nil, err
	}
	for _, secret := range secrets {
		fmt.Println(secret.Name)
	}

	return &App{
		secretStore: ss,
		secrets:     secrets,
	}, nil
}

func (app *App) AddQR() {
	qrData, err := DecodeFromScreen()
	if err != nil {
		go dialog.Message("Couldn't get code from screen, please make sure Screen Recording is enabled for Frothy in system preferences and that the QR code is visible on your screen.").Error()
		return
	}

	newSecret, err := ParseOTPSecretFromURI(qrData)
	if err != nil {
		go dialog.Message("Found QR code on screen but was not a valid 2FA code.").Error()
		return
	}
	for _, existingSecret := range app.secrets {
		if newSecret.Secret == existingSecret.Secret {
			go func(existingName string) {
				dialog.Message("This secret already exists under the name %s", existingName).Error()
			}(existingSecret.Name)
			return
		}
	}

	app.secrets = append(app.secrets, newSecret)
	app.secretStore.SetSecrets(app.secrets)
}

func (app *App) AddCode() {
	fmt.Println("AddCode")
}
