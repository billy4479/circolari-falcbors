package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

type circolare struct {
	Title string `json:"Title"`
	PDFs  []pdf  `json:"PDFs"`
}

type pdf struct {
	Url   string `json:"URL"`
	Title string `json:"Title"`
}

const mainURL = "https://www.liceofalcbors.edu.it/archivio-circolari/"

func scrap() ([]circolare, error) {
	fmt.Print("Downloading root page...")
	root, err := goquery.NewDocument(mainURL)
	if err != nil {
		return nil, err
	}
	fmt.Println("\tDone.")

	var errors error
	var circolari []circolare

	root.Find("section > div.pf-content > div > div > h4 > a").
		Each(func(i int, s *goquery.Selection) {
			if errors != nil {
				return
			}

			href, _ := s.Attr("href")
			title := s.Text()

			circular, err := goquery.NewDocument(href)
			if err != nil {
				errors = err
				return
			}

			fmt.Printf("Downloading \"%s\"...", title)

			var pdfs []pdf
			circular.Find("section > div.pf-content > p > a").
				Each(func(i int, s *goquery.Selection) {
					href, _ := s.Attr("href")
					pdfs = append(pdfs, pdf{
						Title: s.Text(),
						Url:   href,
					})
				})
			fmt.Println("\tDone.")

			circolari = append(circolari, circolare{
				Title: title,
				PDFs:  pdfs,
			})
		})

	if errors != nil {
		return nil, err
	}

	return circolari, nil
}
