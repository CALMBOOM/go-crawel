package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := "https://movie.douban.com/top250"
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	//由于豆瓣具有反爬机制 需要模拟成浏览器
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36 Edg/85.0.564.51")
	resp, err := client.Do(req)
	//url := "https://movie.douban.com/top250"
	dom, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		log.Fatalln(err)
		return
	}
	//fmt.Println(resp.Status)
	//fmt.Println(dom.Html())
	//	dom.Find("#content > div > div.article > ol > li").Each(func(i int, selection *goquery.Selection) {
	//	fmt.Println(selection.Text())
	//	})
	dom.Find(".paginator>a").Each(func(i int, selection *goquery.Selection) {
		//fmt.Println(selection.Text())
		sonurl, _ := selection.Attr("href")
		fmt.Println(sonurl)
	})
}

func getpages() {

}
