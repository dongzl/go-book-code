package handle

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"go-shici/db"
	"io"
	"strings"
)

var poemHome = "https://so.gushiwen.cn/shiwens/default.aspx?astr=李白&page=4"

func getUrl(url string, size int) []string {
	urls := make([]string, 0)
	urlTpl := strings.Replace(url, "page=4", "page=%d", 1)

	for i := 0; i < size; i++ {
		urls = append(urls, fmt.Sprintf(urlTpl, i))
		fmt.Println(fmt.Sprintf(urlTpl, i))
	}

	return urls
}

type PomeInfoHandle struct {

}

func (h *PomeInfoHandle) Worker(body io.Reader, url string) {
	doc, err := goquery.NewDocumentFromReader(body)

	if err != nil {
		fmt.Println("doc.err.", err)
	}

	doc.Find(".cont").Each(func(i int, selection *goquery.Selection) {
		author := ""
		dynsty := ""
		content := ""
		title := ""

		title = strings.TrimSpace(selection.Find("p").Find("a").Find("b").Text())
		author = strings.TrimSpace(selection.Find(".source").Find("a").Eq(0).Text())
		dynsty = strings.TrimSpace(selection.Find(".source").Find("a").Eq(1).Text())
		fmt.Printf("作者：%s，朝代：%s，标题：%s\n", author, dynsty, title)
		selection.Find(".contson").Each(func(i int, selection *goquery.Selection) {
			content = strings.TrimSpace(selection.Text())
			fmt.Printf("内容：%s\n", content)
		})

		if author != "" && dynsty != "" && content != "" && title != "" {
			p := db.Poem{}
			p.Author = author
			p.Title = title
			p.Content = content
			p.Dynasty = dynsty
			p.Save()
		}
	})
}
