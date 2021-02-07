package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// Item is Alfred's item struct.
type Item struct {
	Type     string `json:"type"`
	Icon     string `json:"icon"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Arg      string `json:"arg"`
}

// Menu is Alfred's menu struct.
type Menu struct {
	Items []Item `json:"items"`
}

func main() {
	client := &http.Client{}
	url := "https://ifconfig.me/"
	userAgent := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15) AppleWebKit/537 (KHTML, like Gecko) Chrome/88 Safari/537"
	acceptLanguage := "en-US;q=0.9,en;q=0.8"

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("accept-language", acceptLanguage)
	res, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	// byteArray, _ := ioutil.ReadAll(res.Body)
	// fmt.Println(string(byteArray))

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var ip string
	doc.Find("#ip_address").Each(func(i int, s *goquery.Selection) {
		ip = s.Text()
	})

	var item Item
	item.Icon = "./icon.png"
	item.Title = "Your IP address is " + ip
	item.Subtitle = "Copy to Clipboard"
	item.Arg = ip

	var menu Menu
	menu.Items = append(menu.Items, item)

	menuJSON, _ := json.Marshal(menu)
	fmt.Println(string(menuJSON))
}
