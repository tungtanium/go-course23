package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/tungtanium/go-course23/switchcrawl/pkg/crawl"
)

// Shopee.vn Entry URL components for
// https://shopee.vn/search?category=11035954&keyword=ktt%20strawberry&subcategory=11036000&thirdCategory=11036011
const UrlPrefix string = "https://shopee.vn/search?category=11035954&keyword="
const UrlSuffix string = "&subcategory=11036000&thirdCategory=11036011"

// Lazada.vn entry URL
// https://www.lazada.vn/phu-kien-ban-phim/?q=milky+yellow+pro

type ProductUrl struct {
	ItemId, ShopId string
}

// Generate URL to search for switch from proper categories
func GetEntryUrl(keyword string) string {
	// URLify the keywords first
	keyword = strings.ReplaceAll(keyword, " ", "%20")

	return fmt.Sprintf("%s%s%s", UrlPrefix, keyword, UrlSuffix)
}

// Scrape local
func main() {
	// Read switch data from csv to fill into a database
	pwd, _ := os.Getwd()
	// csvDataPath := filepath.Join(pwd, "assets", "switchdata.csv")
	collyCacheDir := filepath.Join(pwd, "cache", "shopeevn_cache")

	// switches := csvread.CsvReader(csvDataPath)

	// url := GetEntryUrl(switches[0].Name)
	url := "https://www.lazada.vn/catalog/?spm=a2o4n.searchlist.search.d_go.113e5e01qbZbbJ&q=Ajazz+Diced+Fruit+Banana"
	crawl.Crawl(url, collyCacheDir)
	// for idx := 0; idx < 10; idx++ {
	// 	url := GetEntryUrl(switches[idx].Name)
	// 	crawl.Crawl(url, collyCacheDir)

	// }
}
