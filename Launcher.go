package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	Launcher "github.com/dianchi/GoMC/src"
)

func main() {
	var version string
	a := app.New()
	w := a.NewWindow("HelloMinecraft")

	Input1 := widget.NewEntry()
	Input1.TextStyle = fyne.TextStyle{Bold: true}
	Input1.SetPlaceHolder("Entry your name")

	input2 := widget.NewEntry()
	input2.TextStyle = fyne.TextStyle{Bold: true}
	input2.SetPlaceHolder("Entry your password")
	provinceSelect := widget.NewSelect(Launcher.Checker(".minecraft/versions/"), func(value string) {
		fmt.Println("version:", value)
		version = value
	})
	btn1 := widget.NewButton("StartGame", func() {
		GameStart(Input1.Text, input2.Text, version)
	})

	btn1.Resize(fyne.NewSize(100, 100))
	w.SetContent(container.NewVBox(btn1, Input1, input2, provinceSelect))
	w.Resize(fyne.NewSize(500, 500))
	w.SetPadded(false)
	w.ShowAndRun()
}

func GameStart(name string, password string, version string) {
	fmt.Print(time.Now(), "\n")
	Launcher.CmdGen(name, version, "1.19", "G:/Minecraft/JDK17/bin/java.exe", "", "windows", "x64", "1", true, "128m", "4096m")
}
