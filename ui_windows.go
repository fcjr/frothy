package frothy

import (
	"bytes"
	"fmt"
	"image/png"
	"log"
	"time"

	"github.com/d-tsuji/clipboard"
	"github.com/lxn/walk"
)

func (app *App) RunUI() error {
	// We need either a walk.MainWindow or a walk.Dialog for their message loop.
	// We will not make it visible though.
	mw, err := walk.NewMainWindow()
	if err != nil {
		return err
	}

	img, err := png.Decode(bytes.NewReader(windowsLogoBytes))
	if err != nil {
		return err
	}

	icon, err := walk.NewBitmapFromImage(img)
	if err != nil {
		return err
	}

	// Create the notify icon and make sure we clean it up on exit.
	ni, err := walk.NewNotifyIcon(mw)
	if err != nil {
		return err
	}
	defer ni.Dispose()

	// Set the icon and a tool tip text.
	if err := ni.SetIcon(icon); err != nil {
		return err
	}
	if err := ni.SetToolTip("Frothy: TOTP Client"); err != nil {
		return err
	}

	// ni.MouseDown().Attach(func(x, y int, button walk.MouseButton) {
	// 	if button != walk.LeftButton {
	// 		return
	// 	}
	// 	fmt.Println("clicked")
	// })

	// Add Context Menu Items

	// Add Via QR
	actionAddQR := walk.NewAction()
	if err := actionAddQR.SetText("Add Via QR"); err != nil {
		return err
	}
	actionAddQR.Triggered().Attach(func() { app.AddQR() })
	if err := ni.ContextMenu().Actions().Add(actionAddQR); err != nil {
		return err
	}

	// Add Via Code
	actionAddCode := walk.NewAction()
	if err := actionAddCode.SetText("Add Via Code"); err != nil {
		return err
	}
	actionAddCode.Triggered().Attach(func() { app.AddCode() })
	if err := ni.ContextMenu().Actions().Add(actionAddCode); err != nil {
		return err
	}

	if err := ni.ContextMenu().Actions().Add(walk.NewSeparatorAction()); err != nil {
		return err
	}

	// Codes
	actionCodesTitle := walk.NewAction()
	if err := actionCodesTitle.SetText("2FA Codes"); err != nil {
		return err
	}
	if err := actionCodesTitle.SetEnabled(false); err != nil {
		return err
	}
	if err := ni.ContextMenu().Actions().Add(actionCodesTitle); err != nil {
		return err
	}

	if err := ni.ContextMenu().Actions().Add(walk.NewSeparatorAction()); err != nil {
		return err
	}

	go func() {
		secretActions := make(map[string]*walk.Action)

		for {
			app.secretLock.RLock()
			secrets := app.secrets
			app.secretLock.RUnlock()
			if len(secrets) > 0 {

				for _, secret := range secrets {
					secretAction, itemExisted := secretActions[secret.UID]
					if !itemExisted {
						secretAction = walk.NewAction()
						secretActions[secret.UID] = secretAction
					}

					totp, err := NewTOTP(secret.Secret)
					var title string
					if err != nil {
						title = fmt.Sprintf("%s: ERROR", secret.Name)
					} else {
						title = fmt.Sprintf("%s: %s (%.0fs)", secret.Name, totp.Code, time.Until(totp.ExpiresAt).Seconds())
					}

					if itemExisted {
						secretAction.SetText(title)
					} else {
						secretAction.Triggered().Attach(getClipboardFunc(secret.Secret))
						if err := ni.ContextMenu().Actions().Insert(4, secretAction); err != nil {
							log.Println(err) // TODO handle?
						}
					}
				}
			}
		}
	}()

	// Exit
	if err := ni.ContextMenu().Actions().Add(walk.NewSeparatorAction()); err != nil {
		return err
	}
	exitAction := walk.NewAction()
	if err := exitAction.SetText("Exit"); err != nil {
		return err
	}
	exitAction.Triggered().Attach(func() { walk.App().Exit(0) })
	if err := ni.ContextMenu().Actions().Add(exitAction); err != nil {
		return err
	}

	// show icon
	if err := ni.SetVisible(true); err != nil {
		return err
	}

	mw.Run()
	return nil
}

func getClipboardFunc(secret string) func() {
	return func() {
		totp, err := NewTOTP(secret)
		if err == nil {
			_ = clipboard.Set(totp.Code)
		}
	}
}
