// Package main provides various examples of Fyne API capabilities
package main

import (
	"fmt"
	"fyne.io/fyne/dialog"
	"gzv-gui/screens"
	"net/url"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

var globalApp = app.New()
var globalWindow = globalApp.NewWindow("gzv-gui")

func welcomeScreen(a fyne.App) fyne.CanvasObject {
	logo := canvas.NewImageFromResource(theme.FyneLogo())
	logo.SetMinSize(fyne.NewSize(128, 128))

	link, err := url.Parse("https://fyne.io/")
	if err != nil {
		fyne.LogError("Could not parse URL", err)
	}

	return widget.NewVBox(
		widget.NewLabelWithStyle("Welcome to the Fyne toolkit demo app", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		layout.NewSpacer(),
		widget.NewHBox(layout.NewSpacer(), logo, layout.NewSpacer()),
		widget.NewHyperlinkWithStyle("fyne.io", link, fyne.TextAlignCenter, fyne.TextStyle{}),
		layout.NewSpacer(),

		widget.NewGroup("Theme",
			fyne.NewContainerWithLayout(layout.NewGridLayout(2),
				widget.NewButton("Dark", func() {
					a.Settings().SetTheme(theme.DarkTheme())
				}),
				widget.NewButton("Light", func() {
					a.Settings().SetTheme(theme.LightTheme())
				}),
			),
		),
	)
}

func main() {
	globalWindow.SetMainMenu(fyne.NewMainMenu(fyne.NewMenu("File",
		fyne.NewMenuItem("New", func() { fmt.Println("Menu New") }),
	), fyne.NewMenu("Edit",
		fyne.NewMenuItem("Cut", func() { fmt.Println("Menu Cut") }),
		fyne.NewMenuItem("Copy", func() { fmt.Println("Menu Copy") }),
		fyne.NewMenuItem("Paste", func() { fmt.Println("Menu Paste") }),
	)))

	tabs := widget.NewTabContainer(
		widget.NewTabItemWithIcon("Welcome", theme.HomeIcon(), welcomeScreen(globalApp)),
		widget.NewTabItemWithIcon("TX", theme.HomeIcon(), TxScreen()),
		widget.NewTabItemWithIcon("Widgets", theme.ContentCopyIcon(), screens.WidgetScreen()),
		widget.NewTabItemWithIcon("Graphics", theme.DocumentCreateIcon(), screens.GraphicsScreen()),
		widget.NewTabItemWithIcon("Windows", theme.ViewFullScreenIcon(), screens.DialogScreen(globalWindow)),
		widget.NewTabItemWithIcon("Advanced", theme.SettingsIcon(), screens.AdvancedScreen(globalWindow)))
	tabs.SetTabLocation(widget.TabLocationLeading)
	globalWindow.SetContent(tabs)
	globalWindow.ShowAndRun()
}

func TxScreen() fyne.CanvasObject {
	toolbar := widget.NewToolbar()

	return fyne.NewContainerWithLayout(layout.NewBorderLayout(toolbar, nil, nil, nil),
		toolbar,
		widget.NewTabContainer(
			widget.NewTabItem("NormalTx", makeNormalTxTab()),
		),
	)
}

func makeNormalTxTab() fyne.Widget {
	to := widget.NewEntry()
	to.SetPlaceHolder("to address")
	value := widget.NewEntry()
	value.SetPlaceHolder("send ZVC value")
	nonce := widget.NewEntry()
	nonce.SetPlaceHolder("nonce")
	gasPrice := widget.NewEntry()
	gasPrice.SetPlaceHolder("gas price")
	gasLimit := widget.NewEntry()
	gasLimit.SetPlaceHolder("gas limit")

	form := &widget.Form{
		OnCancel: func() {
			fmt.Println("Cancelled")
		},
		OnSubmit: func() {
			message := fmt.Sprintf("confirm:\n  To: %s\n  Value: %s ZVC\n  Nonce: %s\n  GasPrice: %s\n  GasLimit: %s",
				to.Text, value.Text, nonce.Text, gasPrice.Text, gasLimit.Text)
			cnf := dialog.NewConfirm("Confirmation", message, sendCallBack, globalWindow)
			cnf.SetDismissText("Nah")
			cnf.SetConfirmText("Oh Yes!")
			cnf.Show()
		},
	}
	form.Append("To", to)
	form.Append("Value", value)
	form.Append("Nonce", nonce)
	form.Append("GasPrice", gasPrice)
	form.Append("GasLimit", gasLimit)
	return form
}

func sendCallBack(yes bool) {
	if yes {
		fmt.Println("confirmed")
	} else {
		fmt.Println("canceled")
	}
}
