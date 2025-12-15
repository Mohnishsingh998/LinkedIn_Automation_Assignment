package main

import (
	"log"

	"LinkedIn_Automation_Assignment/internal/auth"
	"LinkedIn_Automation_Assignment/internal/browser"
)

func main() {
	br := browser.NewBrowser()
	page := br.MustPage()

	if auth.LoadCookies(page) {
		log.Println("Cookies loaded")
	}

	if auth.IsLoggedIn(page) {
		log.Println("Session is authenticated")
	} else {
		log.Println("Session is NOT authenticated")
	}

	select {}
}
