package main

import (
	"encoding/xml"
	"fmt"
	"time"
)

type Item struct {
	Title      string `xml:"title"`
	PubDate    string `xml:"pubDate"`
	Categories []string `xml:"category"`
}

type Channel struct {
	Items []Item `xml:"item"`
}

type RSS struct {
	Channel Channel `xml:"channel"`
}

func Parse(xmlStr string) (*RSS, error) {
	data := new(RSS)
	if err := xml.Unmarshal([]byte(xmlStr), data); err != nil {
		return nil, err
	}
	return data, nil
}

func (i *Item) Date() time.Time {
	date, _ := time.Parse("Mon, 2 Jan 2006 15:04:05 -0700", i.PubDate)
	return date
}

func (r *RSS) Latest() Item {
	var latest Item
	for i, item := range r.Channel.Items {
		if i == 0 {
			latest = item
			fmt.Printf("init: %v\n", latest.Title)
			continue
		}

		if item.Date().After(latest.Date()) {
			latest = item
			fmt.Printf("latest: %v\n", latest.Title)
		}
	}
	return latest
}