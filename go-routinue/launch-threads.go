package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup
var webSiteList1 []string
var mutex sync.Mutex

func main() {
	webSiteList := []string{
		"https://www.google.com",
		"https://www.github.com",
		"https://www.facebook.com",
		"https://www.linkedin.com",
	}

	for _, web := range webSiteList {
		go findStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()

	fmt.Println(webSiteList1)
}

func findStatusCode(web string) {
	response, err := http.Get(web)
	if err != nil {
		fmt.Println("Error while fetching the website")
	}

	// add url in  the list
	mutex.Lock()
	webSiteList1 = append(webSiteList1, web)
	mutex.Unlock()

	fmt.Printf("status code is %d for %s\n", response.StatusCode, web)
	defer wg.Done()
}
