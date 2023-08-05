package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

const startUrl = "https://books.toscrape.com/"

type Book struct {
	Title string
	Price string
}

func closeFile(file *os.File) {
	err := file.Close()
	if err != nil {
		log.Fatal("Cannot close file", err)
	}
}

func createFile() *os.File {
	file, err := os.Create("books.csv")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	return file
}

func main() {
	file := createFile()
	defer closeFile(file)

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{"Title", "Price"}
	err := writer.Write(headers)
	if err != nil {
		log.Fatal("Cannot write to file", err)
	}

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL.String())
	})

	c.OnHTML(".next > a", func(e *colly.HTMLElement) {
		nextPage := e.Request.AbsoluteURL(e.Attr("href"))
		c.Visit(nextPage)
	})

	c.OnHTML(".product_pod", func(e *colly.HTMLElement) {
		book := Book{}
		book.Title = e.ChildAttr(".image_container img", "alt")
		book.Price = e.ChildText(".price_color")
		row := []string{book.Title, book.Price}
		writer.Write(row)
	})

	c.Visit(startUrl)
}
