package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/extensions"
)

const URL string = "https://www.jd-sports.com.au/men/collection/nike-air-force-1/"
const UAgent string = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/116.0"
const ProductURLPrefix = "https://www.jd-sports.com.au"

type NikeAF1 struct {
	Name    string `json:"name"`
	Price   string `json:"price"`
	ImgUrl  string `json:"imgurl"`
	ProdUrl string `json:"produrl"`
}

func cleanPriceData(rawPrice string) string {
	return strings.Join(strings.Fields(rawPrice), " ")
}

// Trying scrape all Nike AirForce1 collection from JD-Sports.com.au
func CrawlAF1FromJDSportsAU() []NikeAF1 {
	var af1s []NikeAF1

	c := colly.NewCollector(
		colly.AllowURLRevisit(),
		colly.AllowedDomains("www.jd-sports.com.au", "jd-sports.com.au"),
		colly.CacheDir("./jdsportsau_cache"), // Cache dir
	)

	// Add Random User Agents
	extensions.RandomUserAgent(c)

	// Find and visit all links
	c.OnHTML("span.itemContainer", func(h *colly.HTMLElement) {
		cleanedPrice := cleanPriceData(h.ChildText("span.pri"))
		prodLink := fmt.Sprintf("%s%s", ProductURLPrefix, h.ChildAttr("a.itemImage", "href"))

		af1 := NikeAF1{
			Name:    h.ChildText("span.itemTitle"),
			Price:   cleanedPrice,
			ImgUrl:  h.ChildAttr("img.thumbnail", "data-src"),
			ProdUrl: prodLink,
		}
		af1s = append(af1s, af1)
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL.String())
	})

	c.OnResponse(func(r *colly.Response) {
		log.Println("Response Code: ", r.StatusCode)
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Error: ", err)
	})

	c.Visit(URL)

	return af1s
}

func main() {
	af1s := CrawlAF1FromJDSportsAU()
	data, err := json.Marshal(af1s)

	if err != nil {
		fmt.Println(err.Error())
	}
	os.WriteFile("af1.json", data, 0644)

}
