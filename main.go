package main

import (
	"fmt"
	"sync"
	"test/parse"
	"time"
)

var url string
var movie []parse.DoubanMovie

func main() {

	url = "https://movie.douban.com/top250"
	pages := []string{}
	pages = parse.Getpages(url)
	start := time.Now()
	var wg sync.WaitGroup
	for _, j := range pages {
		wg.Add(1)
		go func() {
			movie = append(movie, (parse.Getmovies(j))...)
			wg.Done()
		}()
		//fmt.Println(len(movie))
		//fmt.Println(j)
	}

	//end := time.Now().UnixNano()
	wg.Wait()
	fmt.Println(time.Since(start))
	//fmt.Println(pages[1])
	//printresult()
	fmt.Println(len(movie))
}

func printresult() {
	for _, j := range movie {
		fmt.Println("1:", j.Title)
		fmt.Println("2:", j.Subtitle)
		fmt.Println("3:", j.Other)
		fmt.Println("4:", j.Desc)
		fmt.Println("5:", j.Year)
		fmt.Println("6:", j.Area)
		fmt.Println("7:", j.Tag)
		fmt.Println("8:", j.Star)
		fmt.Println("9:", j.Comment)
		fmt.Println("10:", j.Quote)

	}
}
