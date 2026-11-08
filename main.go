package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

var browser_config_url = "https://httpbin.org/"
var find_user_agent_url = "https://httpbin.org/user-agent"
var user_agent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/154.0.0.0 Safari/537.36"

func getHtml(body string, link string) {
	doc, err := html.Parse(strings.NewReader(body))

	if err != nil {
		log.Fatal(err)
	}

	for n := range doc.Descendants() {
		if n.Type == html.ElementNode && n.DataAtom == atom.A {
			for _, a := range n.Attr {
				if a.Key == "href" {
					Sublink := a.Val
					if !strings.HasPrefix(Sublink, "#") &&
						!strings.HasPrefix(Sublink, "https") &&
						!strings.HasPrefix(Sublink, "/") {
						fmt.Println(Sublink)
						crawl(link + "/" + Sublink)
					}
					break
				}
			}
		}
	}
}
func crawl(link string) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("User-Agent", user_agent)
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	bodyByte, err := io.ReadAll(resp.Body)

	body := string(bodyByte)
	fmt.Println("===============================> ", link)
	fmt.Println(body)

	getHtml(body, link)
}

func main() {
	link := "https://wingram.org/"
	crawl(link)
}
