package frothy

import (
	"bytes"
	"image/png"

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

	// Exit
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
