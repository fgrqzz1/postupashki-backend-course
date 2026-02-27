package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"io"
	"net/http"
	"os"
	"time"
)

type result struct {
	URL      string
	Response *http.Response
	Error    error
}

func fetch(url string, results chan<- result, client *http.Client) {
	resp, err := client.Get(url)
	if err != nil {
		results <- result{URL: url, Error: err}
		return
	}

	results <- result{url, resp, err}
}

func main() {
	var timeout int
	var help bool

	pflag.IntVarP(&timeout, "timeout", "t", 15, "Таймаут в секундах")
	pflag.BoolVarP(&help, "help", "h", false, "Справка по команде")

	pflag.Parse()

	if help {
		pflag.Usage()
		os.Exit(0)
	}

	urls := pflag.Args()
	if len(urls) == 0 {
		//fmt.Println("Ошибка: не указаны URL")
		os.Exit(1)
	}

	client := &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}

	results := make(chan result, len(urls))

	for _, url := range urls {
		go fetch(url, results, client)
	}

	errorsCount := 0
	for {
		select {
		case r := <-results:
			if r.Error != nil {
				errorsCount++
				if errorsCount == len(urls) {
					//fmt.Fprintf(os.Stderr, "Все запросы завершились с ошибкой\n")
					os.Exit(1)
				}
				continue
			}

			fmt.Println(r.Response.Status)
			for name, headers := range r.Response.Header {
				for _, value := range headers {
					fmt.Printf("%s: %s\n", name, value)
				}
			}
			fmt.Println()
			io.Copy(os.Stdout, r.Response.Body)
			r.Response.Body.Close()

			os.Exit(0)

		case <-time.After(time.Duration(timeout) * time.Second):
			fmt.Fprintf(os.Stderr, "Глобальный таймаут\n")
			os.Exit(228)
		}
	}

}
