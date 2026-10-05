package main

import (
	"fmt"
	"log"
	"time"

	"github.com/mxschmitt/playwright-go"
)

func main() {
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}
	defer pw.Stop()

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(true),
	})
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}
	defer browser.Close()

	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}

	page.OnConsole(func(msg playwright.ConsoleMessage) {
		fmt.Printf("[CONSOLE %s] %s\n", msg.Type(), msg.Text())
	})
	page.OnPageError(func(err error) {
		fmt.Printf("[PAGE ERROR] %v\n", err)
	})

	fmt.Println("Visiting http://localhost:8080...")
	_, err = page.Goto("http://localhost:8080", playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateDomcontentloaded,
	})
	if err != nil {
		log.Fatalf("goto error: %v", err)
	}

	time.Sleep(2 * time.Second)

	tbody, _ := page.QuerySelector("#job-tbody")
	fmt.Printf("#job-tbody exists in DOM: %v\n", tbody != nil)

	// Now try clicking submit on the form without filling or with filling
	btn, _ := page.QuerySelector("button[type='submit']")
	if btn != nil {
		fmt.Println("Clicking Start Scraping...")
		btn.Click()
	}

	time.Sleep(2 * time.Second)
	tbody2, _ := page.QuerySelector("#job-tbody")
	fmt.Printf("After submit, #job-tbody exists: %v\n", tbody2 != nil)
}
