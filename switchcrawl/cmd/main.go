package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/tungtanium/go-course23/switchcrawl/pkg/repo/csvread"
)

// Entry URL components for
// https://shopee.vn/search?category=11035954&keyword=ktt%20strawberry&subcategory=11036000&thirdCategory=11036011
const UrlPrefix string = "https://shopee.vn/search?category=11035954&keyword="
const UrlSuffix string = "&subcategory=11036000&thirdCategory=11036011"

type ProductUrl struct {
	ItemId, ShopId string
}

// Generate URL to search for switch from proper categories
func GetEntryUrl(keyword string) string {
	return fmt.Sprintf("%s%s%s", UrlPrefix, keyword, UrlSuffix)
}

// Scrape local
func main() {
	pwd, _ := os.Getwd()
	csvDataPath := filepath.Join(pwd, "assets", "switchdata.csv")
	csvread.CsvReader(csvDataPath)
}
