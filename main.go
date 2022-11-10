package main

import (
	"fmt"
	"time"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("www.nike.com"),
	)
	c.OnXML("//*[@id='root']/div/div/div[1]/div/div[1]/div[2]/div/section[1]/div[2]/aside/div/div[2]", func(e *colly.XMLElement) {
		text := e.Text
		fmt.Println(text)
		for true {
			time.Sleep(3 * time.Second)
			if text[72:85] == "StockUpcoming" {
				fmt.Println("Not available yet, still looking...")
			} else {
				fmt.Println("InStock, buying...")
			}
		}
	})
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})
	c.Visit("https://www.nike.com/ca/launch/t/air-jordan-7-x-trophy-room-true-red-and-obsidian")
}
