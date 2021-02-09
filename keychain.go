package frothy

import (
	"os"
	"sync"

	"github.com/99designs/keyring"
	cbor "github.com/fxamacker/cbor/v2"
	"github.com/gen2brain/dlgs"
	"github.com/google/uuid"
)

var keyringConfigDefaults = keyring.Config{
	ServiceName:              "frothy",
	FileDir:                  "~/.frothy/secrets/",
	FilePasswordFunc:         fileKeyringPassphrasePrompt,
	LibSecretCollectionName:  "frothy",
	KWalletAppID:             "frothy",
	KWalletFolder:            "frothy",
	KeychainTrustApplication: true,
	WinCredPrefix:            "frothy",
}

type SecretStore struct {
	lock sync.RWMutex
	ring keyring.Keyring
}

func OpenSecretStore() (*SecretStore, error) {
	ring, err := keyring.Open(keyringConfigDefaults)
	if err != nil {
		return nil, err
	}
	return &SecretStore{
		ring: ring,
	}, nil
}

func (ss *SecretStore) GetSecrets() ([]*OTPSecret, error) {
	ss.lock.RLock()
	defer ss.lock.RUnlock()

	keys, err := ss.ring.Keys()
	if err != nil {
		return nil, err
	}

	var secrets []*OTPSecret
	for _, key := range keys {
		item, err := ss.ring.Get(key)
		if err != nil {
			return nil, err
		}
		var secret OTPSecret
		if err := cbor.Unmarshal(item.Data, &secret); err != nil {
			return nil, err
		}
		secrets = append(secrets, &secret)
	}
	return secrets, nil
}

func (ss *SecretStore) SetSecrets(secrets []*OTPSecret) error {
	ss.lock.Lock()
	defer ss.lock.Unlock()

	for _, secret := range secrets {
		if secret.UID == "" {
			secret.UID = uuid.New().String()
		}

		b, err := cbor.Marshal(&secret)
		if err != nil {
			return err
		}

		err = ss.ring.Set(keyring.Item{
			Key:  secret.UID,
			Data: b,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func fileKeyringPassphrasePrompt(prompt string) (string, error) {
	if password := os.Getenv("FROTHY_FILE_PASSPHRASE"); password != "" {
		return password, nil
	}

	// TODO replace dlgs
	password, _, err := dlgs.Password("Unlock Frothy", "Enter your Frothy Password:")
	if err != nil {
		return "", err
	}
	return password, nil
}
