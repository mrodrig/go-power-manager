package main

import (
	"log"
	"os/exec"

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

	window.SetContent(widget.NewVBox(
		widget.NewHBox(
			widget.NewButton("Shutdown", func() {
				var flags = []string{"/s", "/f"}
				runCommand(flags)
			}),
			widget.NewButton("Reboot", func() {
				var flags = []string{"/r", "/f"}
				runCommand(flags)
			}),
		),
		widget.NewHBox(
			widget.NewButton("Hibernate", func() {
				var flags = []string{"/h"}
				runCommand(flags)
			}),
			widget.NewButton("Quit", func() {
				app.Quit()
			}),
		),
	))

	window.ShowAndRun()
}
