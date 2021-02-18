package desktop

import (
	"fmt"
	"sync"

	"github.com/fcjr/alert"
	"github.com/fcjr/frothy"
	"github.com/martinlindhe/inputbox"
)

type App struct {
	secretLock  sync.RWMutex
	secretStore *SecretStore
	secrets     []*frothy.OTPSecret
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

	secret, err := frothy.ParseOTPSecretFromURI(qrData)
	if err != nil {
		go alert.Error("Error", "Found QR code on screen but was not a valid 2FA code.")
		return
	}

	app.addSecret(secret)
}

func (app *App) AddCode() {
	code, ok := inputbox.InputBox("Add TOTP via Code", "Enter Code given by the website", "")
	if ok {
		if code == "" {
			go app.AddCode()
			return
		}
		fmt.Println(code)
	}
}

func (app *App) addSecret(secret *frothy.OTPSecret) {
	for _, existingSecret := range app.secrets {
		if secret.Secret == existingSecret.Secret {
			go func(existingName string) {
				go alert.Error("Error", fmt.Sprint("This secret already exists under the name %s.", existingName))
			}(existingSecret.Name)
			return
		}
	}

	app.secretLock.Lock()
	defer app.secretLock.Unlock()
	app.secrets = append(app.secrets, secret)
	if err := app.secretStore.SetSecrets(app.secrets); err != nil {
		go func(name string) {
			alert.Error("Error", fmt.Sprint("Failed to store 2FA Key for %s to keychain.", name))
		}(secret.Name)
		return
	}
}
