package frothy

import (
	"fmt"
	"sync"

	"github.com/sqweek/dialog"
)

type App struct {
	secretLock  sync.RWMutex
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
				dialog.Message("This secret already exists under the name %s.", existingName).Error()
			}(existingSecret.Name)
			return
		}
	}

	app.secretLock.Lock()
	defer app.secretLock.Unlock()
	app.secrets = append(app.secrets, newSecret)
	if err := app.secretStore.SetSecrets(app.secrets); err != nil {
		go func(name string) {
			dialog.Message("Failed to store 2FA Key for %s to keychain.", name).Error()
		}(newSecret.Name)
		return
	}
}

func (app *App) AddCode() {
	fmt.Println("AddCode")
}
