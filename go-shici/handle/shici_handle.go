package handle

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"go-shici/gofish"
	"io"
	"strings"
)

var baseUrl = "https://so.gushiwen.cn"

type AuthorHandle struct {
}

func (h *AuthorHandle) Worker(body io.Reader, url string) {
	doc, err := goquery.NewDocumentFromReader(body)

	if err != nil {
		fmt.Println("doc.err.", err)
	}

	doc.Find(".sons").Find(".cont").Find("a").Each(func(i int, selection *goquery.Selection) {
		author := selection.Text()
		fmt.Printf("%d author=%s\n", i, author)
		link, _ := selection.Attr("href")
		fmt.Printf("%d link=%s\n", i, link)

		h := PomeHomeHandle{}
		fish := gofish.NewGoFish()
		request, err := gofish.NewRequest("GET", baseUrl+link, gofish.UserAgent, &h, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		fish.Request = request
		fish.Visit()
	})
}

type PomeHomeHandle struct {
}

func (h *PomeHomeHandle) Worker(body io.Reader, url string) {
	doc, err := goquery.NewDocumentFromReader(body)

	if err != nil {
		fmt.Println("doc.err.", err)
	}

	doc.Find(".sonspic").Find(".cont").Find("p").Find("a").Each(func(i int, selection *goquery.Selection) {
		link, _ := selection.Attr("href")

		if !strings.HasPrefix(link, "/shiwens") {
			return
		}

		fmt.Println("作品主页=", baseUrl+link)

		h := PomeInfoHandle{}
		fish := gofish.NewGoFish()
		request, err := gofish.NewRequest("GET", baseUrl+link, gofish.UserAgent, &h, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		fish.Request = request
		fish.Visit()
	})
}
