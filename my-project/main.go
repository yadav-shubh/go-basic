package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Image struct {
	DownloadURL string `json:"download_url"`
}

func main() {

	URL := "https://picsum.photos/v2/list?page=1&limit=10"
	response, e := http.Get(URL)
	if e != nil {
		log.Fatal(e)
	}
	defer response.Body.Close()

	var images []Image
	if response.StatusCode == http.StatusOK {
		err := json.NewDecoder(response.Body).Decode(&images)
		if err != nil {
			log.Fatal("error = ", err)
		}
	}

	for i := 0; i < len(images); i++ {
		// fileName := "temp" + strconv.Itoa(counter)
		// counter++
		// err := downloadFile(URL, fileName)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		fmt.Println("URLs ", i, " : ", images[i].DownloadURL)
	}

}
