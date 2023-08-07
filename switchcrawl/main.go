package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gocolly/colly"
)

// Entry URL components for
// https://shopee.vn/search?category=11035954&keyword=ktt%20strawberry&subcategory=11036000&thirdCategory=11036011
const UrlPrefix string = "https://shopee.vn/search?category=11035954&keyword="
const UrlSuffix string = "&subcategory=11036000&thirdCategory=11036011"

type ProductUrl struct {
	ItemId, ShopId string
}

type Product struct {
	Name, PricePerUnit, ImgUrl, Type string
	// Type could be default to ""
}

// Generate URL to search for switch from proper categories
func GetEntryUrl(keyword string) string {
	return fmt.Sprintf("%s%s%s", UrlPrefix, keyword, UrlSuffix)
}

// Scrape local
func main() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}

	t := &http.Transport{}
	t.RegisterProtocol("file", http.NewFileTransport(http.Dir("/")))

	c := colly.NewCollector()
	c.WithTransport(t)

	pages := []string{}

	c.OnHTML("h1", func(e *colly.HTMLElement) {
		pages = append(pages, e.Text)
	})

	c.OnHTML("a", func(e *colly.HTMLElement) {
		c.Visit("file://" + dir + "/html" + e.Attr("href"))
	})

	fmt.Println("file://" + dir + "/html/index.html")
	c.Visit("file://" + dir + "/html/index.html")
	c.Wait()
	for i, p := range pages {
		fmt.Printf("%d : %s\n", i, p)
	}
}
