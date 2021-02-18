package desktop

import (
	"fmt"
	"runtime"
	"time"

	"github.com/d-tsuji/clipboard"
	"github.com/fcjr/frothy"
	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
)

func (app *App) RunUI() error {
	runtime.LockOSThread()

	// Don't quit app when all windows are closed
	cocoa.DefaultDelegateClass.AddMethod("applicationShouldTerminateAfterLastWindowClosed:", func(notification objc.Object) bool {
		return false
	})

	a := cocoa.NSApp_WithDidLaunch(func(n objc.Object) {
		obj := cocoa.NSStatusBar_System().StatusItemWithLength(cocoa.NSVariableStatusItemLength)
		obj.Retain()

		// set icon
		data := core.NSData_WithBytes(frothy.MacLogoBytes, uint64(len(frothy.MacLogoBytes)))
		image := cocoa.NSImage_InitWithData(data)
		image.SetSize(core.Size(16.0, 16.0))
		image.SetTemplate(true)
		obj.Button().SetImage(image)

		menu := cocoa.NSMenu_New()

		// Setup Items
		itemAddQR := cocoa.NSMenuItem_New()
		itemAddQR.SetTitle("Add Via QR")
		itemAddQR.SetAction(objc.Sel("addQRClicked:"))
		cocoa.DefaultDelegateClass.AddMethod("addQRClicked:", func(_ objc.Object) {
			app.AddQR()
		})
		menu.AddItem(itemAddQR)

		itemAddCode := cocoa.NSMenuItem_New()
		itemAddCode.SetTitle("Add Via Code")
		itemAddCode.SetAction(objc.Sel("addCodeClicked:"))
		cocoa.DefaultDelegateClass.AddMethod("addCodeClicked:", func(_ objc.Object) {
			app.AddCode()
		})
		menu.AddItem(itemAddCode)

		// add title
		menu.AddItem(cocoa.NSMenuItem_Separator())

		codesTitle := cocoa.NSMenuItem_New()
		codesTitle.SetTitle("2FA Codes") // TODO make bold?
		codesTitle.SetEnabled(false)
		menu.AddItem(codesTitle)

		menu.AddItem(cocoa.NSMenuItem_Separator())

		itemQuit := cocoa.NSMenuItem_New()
		itemQuit.SetTitle("Quit")
		itemQuit.SetAction(objc.Sel("terminate:"))
		menu.AddItem(itemQuit)

		// set Menu
		obj.SetMenu(menu)

		go func() {
			secretItems := make(map[string]cocoa.NSMenuItem)

			for {
				app.secretLock.RLock()
				secrets := app.secrets
				app.secretLock.RUnlock()
				if len(secrets) > 0 {

					for _, secret := range secrets {
						secretsItem, itemExisted := secretItems[secret.UID]
						if !itemExisted {
							secretsItem = cocoa.NSMenuItem_New()
							secretItems[secret.UID] = secretsItem
						}

						totp, err := frothy.NewTOTP(secret.Secret)
						var title string
						if err != nil {
							title = fmt.Sprintf("%s: ERROR", secret.Name)
						} else {
							title = fmt.Sprintf("%s: %s (%.0fs)", secret.Name, totp.Code, time.Until(totp.ExpiresAt).Seconds())
						}

						if itemExisted {
							core.Dispatch(func() {
								secretsItem.SetTitle(title)
							})
						} else {
							methodName := fmt.Sprintf("copyToCliboard_%s:", secret.UID)
							cocoa.DefaultDelegateClass.AddMethod(methodName, getClipboardFunc(secret.Secret))
							secretsItem.SetAction(objc.Sel(methodName))
							menu.Send("insertItem:atIndex:", secretsItem, 4)
						}
					}
				}
			}
		}()
	})
	a.SetActivationPolicy(cocoa.NSApplicationActivationPolicyAccessory)
	a.Run()
	return nil
}

// TODO this is garbage clean upp the clipboard functionality
func getClipboardFunc(secret string) func(_ objc.Object) {
	return func(_ objc.Object) {
		totp, err := frothy.NewTOTP(secret)
		if err == nil {
			_ = clipboard.Set(totp.Code)
		}
	}
}
