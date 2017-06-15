package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/mizoguche/go-hatebu/hatebu"
	"github.com/mizoguche/hugo-hatebu/feed"
)

func main() {
	os.Exit(run())
}

func run() int {
	args := os.Args[1:]
	filename := args[0]
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("failted to read %v", filename)
		return 1
	}

	xml := string(b)
	x, err := feed.Parse(xml)
	if err != nil {
		fmt.Printf("failted to parse xml: %v\n", err)
		return 1
	}

	latest := x.Latest()

	duration, _ := time.ParseDuration("-1h")
	threshold := time.Now().Add(duration)
	if latest.Date().Before(threshold) {
		fmt.Println("no new article")
		return 0
	}

	category := ""
	for _, c := range latest.Categories {
		category = fmt.Sprintf("%s[%s]", category, c)
	}

	comment := ""
	if os.Getenv("HATENA_COMMENT") != "" {
		comment = os.Getenv("HATENA_COMMENT")
	}

	req := hatebu.AddRequest{
		URL:         latest.Link,
		Comment:     category + comment,
		PostTwitter: true,
	}

	return hatebu.Add(req)
}
