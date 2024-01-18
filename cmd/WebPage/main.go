package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gocolly/colly"
)

func main() {

	var webpageContent string

	collector := colly.NewCollector(
		colly.AllowedDomains("vitalitysouth.com"),
	)

	collector.OnHTML("html", func(element *colly.HTMLElement) {
		webpageContent = element.Text
	})

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL.String())
	})

	err := collector.Visit("https://vitalitysouth.com/about-us/")
	if err != nil {
		fmt.Println("Error visiting the website:", err)
		return
	}

	writeHTML(webpageContent)
}

func writeHTML(content string) {
	err := ioutil.WriteFile("vs_content.html", []byte(content), 0644)

	if err != nil {
		log.Println("Unable to create HTML File")
		return
	}

	fmt.Println("Page has been successfully saved to vs_content.html")
}
