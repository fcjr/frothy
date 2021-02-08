package frothy

import (
	"runtime"

	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
)

func (app *App) RunUI() {
	runtime.LockOSThread()

	a := cocoa.NSApp_WithDidLaunch(func(n objc.Object) {
		obj := cocoa.NSStatusBar_System().StatusItemWithLength(cocoa.NSVariableStatusItemLength)
		obj.Retain()

		// set icon
		data := core.NSData_WithBytes(logoBytes, uint64(len(logoBytes)))
		image := cocoa.NSImage_InitWithData(data)
		image.SetSize(core.Size(16.0, 16.0))
		image.SetTemplate(true)
		obj.Button().SetImage(image)

		// Setup Items
		itemAddQR := cocoa.NSMenuItem_New()
		itemAddQR.SetTitle("Add Via QR")
		itemAddQR.SetAction(objc.Sel("addQRClicked:"))
		cocoa.DefaultDelegateClass.AddMethod("addQRClicked:", func(_ objc.Object) {
			app.AddQR()
		})

		itemAddCode := cocoa.NSMenuItem_New()
		itemAddCode.SetTitle("Add Via Code")
		itemAddCode.SetAction(objc.Sel("addCodeClicked:"))
		cocoa.DefaultDelegateClass.AddMethod("addCodeClicked:", func(_ objc.Object) {
			app.AddCode()
		})

		// TODO add Codes

		itemQuit := cocoa.NSMenuItem_New()
		itemQuit.SetTitle("Quit")
		itemQuit.SetAction(objc.Sel("terminate:"))

		// Create Menu
		menu := cocoa.NSMenu_New()
		menu.AddItem(itemAddQR)
		menu.AddItem(itemAddCode)
		menu.AddItem(cocoa.NSMenuItem_Separator())
		menu.AddItem(itemQuit)
		obj.SetMenu(menu)

	})
	a.SetActivationPolicy(cocoa.NSApplicationActivationPolicyAccessory)
	a.Run()
}
