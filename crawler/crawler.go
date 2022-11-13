package crawler

import (
	"fmt"
	"log"
	"strings"

	"github.com/playwright-community/playwright-go"
)

func Crawler() {
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}

	headlessOption := false

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: &headlessOption,
	})

	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}

	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}

	if _, err = page.Goto("https://portal.tjpr.jus.br/jurisprudencia/publico/pesquisa.do?actionType=pesquisar"); err != nil {
		log.Fatalf("could not goto: %v", err)
	}

	entries, err := page.QuerySelectorAll(".juris-tabela-dados")
	if err != nil {
		log.Fatalf("could not get entries: %v", err)
	}

	for i, entry := range entries {
		titleElement, err := entry.QuerySelector("a")
		if err != nil {
			log.Fatalf("could not get title element: %v", err)
		}
		title, err := titleElement.TextContent()
		if err != nil {
			log.Fatalf("could not get text content: %v", err)
		}
		fmt.Printf("%d: %s\n", i+1, strings.TrimSpace(title))
	}

	if err = browser.Close(); err != nil {
		log.Fatalf("could not close browser: %v", err)
	}

	if err = pw.Stop(); err != nil {
		log.Fatalf("could not stop Playwright: %v", err)
	}
}
