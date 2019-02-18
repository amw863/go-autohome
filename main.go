package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
	// "net/url"
	// "crypto/md5"
	"strings"
	"sync"
)

type Item struct {
	url   string
	depth int
}

func main() {

	// Items := make(map[string]Item)

	req_url := "https://mall.autohome.com.cn/list/0-310100-0-0-0-0-0-0-0-1.html"
	// wg := &sync.WaitGroup{}
	i := 0
	links := []string{}
	for {
		links = append(links, req_url)
		i++
		fmt.Println(i)
		next_url, IsEixst := getLink(req_url)
		if !IsEixst || i > 3000 {
			break
		}

		req_url = fmt.Sprintf("https://mall.autohome.com.cn/%s", next_url)
		// md5_url := fmt.Sprintf("%x", md5.Sum([]byte(req_url)))

		// if _, ok := Items[md5_url]; ok {
		// 	fmt.Println(req_url, "is exists.")
		// 	continue
		// }

		// item := Item{
		// 	req_url,
		// 	1,
		// }

		// Items[md5_url] = item
	}

	ch := make(chan int, 10)
	for _, link := range links {
		ch <- 1
		go downloader(link, ch)
	}

	filename := "autohome.txt"
	if ioutil.WriteFile(filename, []byte(strings.Join(links, "\n")), 0664) == nil {
		fmt.Println("写入文件成功.", links)
	}

	// fmt.Println(Items)
}

func getLink(req_url string) (string, bool) {
	client := &http.Client{}

	request, err := http.NewRequest("GET", req_url, nil)
	if err != nil {
		panic(err)
	}

	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.96 Safari/537.36")
	resp, err := client.Do(request)
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	return doc.Find(".pager .pager-next").Attr("href")
}

func downloader(banner string, ch chan int) {
	fmt.Println("download from", url)
	<-ch
}
