package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {

	blog_titles, err := getBlogTitles("https://www.wikipedia.org/")

	if err != nil {
		fmt.Println("err:=", err)
		return
	}

	fmt.Printf("Blog Titles: %v\n", blog_titles)
}

func getBlogTitles(url string) (titles string, err error) {

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("status code error := ", resp.StatusCode, resp.Status)
	}

	document, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}

	document.Find(".other-project-title").Each(func(i int, s *goquery.Selection) {
		titles += fmt.Sprintf("%s\n", s.Text())
	})

	return titles, nil
}
