package main

import (
	"DBwork/resource"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	fmt.Printf("Potatoler Banking System starting...\n")
	fmt.Printf("Connecting to the database...\n")

	myApp := app.New()
	myApp.Settings().SetTheme(&resource.MyTheme{})
	myWindow := myApp.NewWindow("Banking System")
	myWindow.Resize(fyne.NewSize(200, 300))
	myWindow.SetContent(ConnectingView(myWindow))
	myWindow.SetContent(LoginView(myWindow))
	myWindow.ShowAndRun()
}
