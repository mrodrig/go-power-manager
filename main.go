package main

import (
	"log"
	"os/exec"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
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
		var flags = []string{"/s", "/f", "/t", "0"}
		runCommand(flags)
	})
	rebootButton := widget.NewButton("Reboot", func() {
		var flags = []string{"/r", "/f", "/t", "0"}
		runCommand(flags)
	})
	hibernateButton := widget.NewButton("Hibernate", func() {
		var flags = []string{"/h"}
		runCommand(flags)
	})
	quitButton := widget.NewButton("Quit", func() {
		app.Quit()
	})

	// Setup the GUI layout
	window.SetContent(fyne.NewContainerWithLayout(layout.NewAdaptiveGridLayout(2), shutdownButton, rebootButton, hibernateButton, quitButton))
	window.Resize(fyne.NewSize(400, 200))

	// Show the interface
	window.ShowAndRun()
}
