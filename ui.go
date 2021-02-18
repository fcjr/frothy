package frothy

import (
	"fmt"
	"sync"

	"github.com/fcjr/alert"
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
		go alert.Error("Error", "Couldn't get code from screen, please make sure Screen Recording is enabled for Frothy in system preferences and that the QR code is visible on your screen.")
		return
	}

	newSecret, err := ParseOTPSecretFromURI(qrData)
	if err != nil {
		go alert.Error("Error", "Found QR code on screen but was not a valid 2FA code.")
		return
	}
	for _, existingSecret := range app.secrets {
		if newSecret.Secret == existingSecret.Secret {
			go func(existingName string) {
				go alert.Error("Error", fmt.Sprint("This secret already exists under the name %s.", existingName))
			}(existingSecret.Name)
			return
		}
	}

	app.secretLock.Lock()
	defer app.secretLock.Unlock()
	app.secrets = append(app.secrets, newSecret)
	if err := app.secretStore.SetSecrets(app.secrets); err != nil {
		go func(name string) {
			alert.Error("Error", fmt.Sprint("Failed to store 2FA Key for %s to keychain.", name))
		}(newSecret.Name)
		return
	}
}

func (app *App) AddCode() {
	fmt.Println("AddCode")
}
