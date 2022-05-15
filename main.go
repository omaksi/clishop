package main

import (
	"fmt"
	"log"
	"os"

	"ondrejmaksi.com/db2project/ui"
)

func main() {

	setUpFileLogging()

	app := ui.Init()

	if err := app.Run(); err != nil {
		fmt.Printf("Error running application: %s\n", err)
		panic(err)
	}

}

func setUpFileLogging() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)
	log.Println("Hello world!")
}
