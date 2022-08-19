package main

import (
	"log"
	"os/exec"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func runCommand(commandFlags []string) {
	cmd := exec.Command("shutdown", commandFlags...)

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	} else {
		log.Print("success")
	}
}

func main() {
	app := app.New()
	window := app.NewWindow("GoPowerManager")

	// Create buttons
	shutdownButton := widget.NewButton("Shutdown", func() {
		var flags = []string{"/s", "/f"}
		runCommand(flags)
	})
	rebootButton := widget.NewButton("Reboot", func() {
		var flags = []string{"/r", "/f"}
		runCommand(flags)
	})
	hibernateButton := widget.NewButton("Hibernate", func() {
		var flags = []string{"/h"}
		runCommand(flags)
	})
	quitButton := widget.NewButton("Quit", func() {
		app.Quit()
	})

	// Resize them all to a consistent size
	shutdownButton.Resize(fyne.NewSize(250, 80))
	rebootButton.Resize(fyne.NewSize(250, 80))
	hibernateButton.Resize(fyne.NewSize(250, 80))
	quitButton.Resize(fyne.NewSize(250, 80))

	// Setup the GUI layout
	window.SetContent(widget.NewVBox(
		widget.NewHBox(shutdownButton, rebootButton),
		widget.NewHBox(hibernateButton, quitButton),
	))

	// Show the interface
	window.ShowAndRun()
}
