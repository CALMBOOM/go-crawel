package parse

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type DoubanMovie struct {
	Title    string
	Subtitle string
	Other    string
	Desc     string
	Year     string
	Area     string
	Tag      string
	Star     string
	Comment  string
	Quote    string
}

func Getmovies(pages string) []DoubanMovie {
	var moviedata []DoubanMovie
	client := &http.Client{}
	req, _ := http.NewRequest("GET", pages, nil)
	//由于豆瓣具有反爬机制 需要模拟成浏览器
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36 Edg/85.0.564.51")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("no signal")
	}
	dom, _ := goquery.NewDocumentFromResponse(resp)
	dom.Find("#content > div > div.article > ol > li").Each(func(i int, selection *goquery.Selection) {
		//fmt.Println(selection.Text())
		//var t string
		var tmp DoubanMovie
		tmp.Title = selection.Find(".hd>a>span").Eq(0).Text()
		tmp.Subtitle = strings.TrimLeft(selection.Find(".hd>a>span").Eq(1).Text(), " / ")
		tmp.Other = strings.TrimLeft(selection.Find(".hd>a>span").Eq(2).Text(), " / ")
		s := strings.TrimSpace(selection.Find(".bd>p").Eq(0).Text())
		s_tmp := strings.Split(s, "\n")
		movieDesc := strings.Split(s_tmp[1], "/")
		tmp.Desc = s_tmp[0]
		//	fmt.Println(selection.Find(".star").Text())
		//fmt.Println(s_tmp)

		tmp.Year = strings.TrimSpace(movieDesc[0])
		tmp.Area = strings.TrimSpace(movieDesc[1])
		tmp.Star = strings.TrimSpace(selection.Find(".bd .star .rating_num").Text())
		tmp.Tag = strings.TrimSpace(movieDesc[2])
		tmp.Quote = strings.TrimSpace(selection.Find(".bd>.quote").Text())
		tmp.Comment = strings.TrimSpace(selection.Find(".bd .star span").Eq(3).Text())
		moviedata = append(moviedata, tmp)

	})
	return moviedata
}

//获取全部分页
func Getpages(url string) []string {
	ss := []string{}
	ss = append(ss, url)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	//由于豆瓣具有反爬机制 需要模拟成浏览器
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36 Edg/85.0.564.51")
	resp, _ := client.Do(req)
	dom, _ := goquery.NewDocumentFromResponse(resp)

	dom.Find(".paginator>a").Each(func(i int, selection *goquery.Selection) {
		sonurl, _ := selection.Attr("href")
		ss = append(ss, url+sonurl)

	})
	return ss
}
