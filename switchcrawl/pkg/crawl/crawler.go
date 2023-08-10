package crawl

import (
	"fmt"
	"log"

	"github.com/gocolly/colly/v2"
)

func Crawl(url string, cachepath string) {
	c := colly.NewCollector(
		colly.AllowURLRevisit(),
		// colly.AllowedDomains("www.shopee.vn", "shopee.vn"),
		colly.CacheDir(cachepath), // Cache dir
	)

	// extensions.RandomUserAgent(c)

	c.OnRequest(func(r *colly.Request) {
		currentURL := r.URL.String()
		log.Println("Visiting: ", currentURL)
	})

	c.OnHTML("title", func(h *colly.HTMLElement) {
		// rawLink := h.ChildAttr("a", "href")
		fmt.Println("Dit Me May")
	})

	c.OnResponse(func(r *colly.Response) {
		log.Println("Response Code: ", r.StatusCode)
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Error: ", err)
	})

	c.Visit(url)

}
