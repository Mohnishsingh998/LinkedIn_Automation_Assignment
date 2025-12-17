package main

import (
	"log"

	"LinkedIn_Automation_Assignment/internal/auth"
	"LinkedIn_Automation_Assignment/internal/browser"
)

func main() {
	br := browser.NewBrowser()
	page := br.MustPage()

	if err := auth.Authenticate(page); err != nil {
		log.Fatal(err)
	}

	log.Println("Authenticated successfully")
	select {}
}
