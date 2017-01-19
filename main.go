package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type requestFile struct {
	URL      string
	Filename string
}

type responseFile struct {
	Href      string `json:"href"`
	Method    string `json:"method"`
	Templated string `json:"templated"`
}

// GetAllContents getting json from remote server
func getAllContents(URL string) []byte {
	response, err := http.Get(URL)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
	}
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("%s", err)
		os.Exit(1)
	}
	return contents
}

func getObject(r requestFile) ([]byte, error) {
	response := getAllContents(r.URL)
	var secondResponse responseFile
	json.Unmarshal(response, &secondResponse)
	if secondResponse.Href == "" {
		return response, errors.New(r.Filename + " file is not found")
	}
	f := getAllContents(secondResponse.Href)
	return f, nil
}

func saveObject(b []byte, f string) {
	err := ioutil.WriteFile(f, b, 0644)
	if err != nil {
		log.Println(err)
	}
	log.Println("DONE")
}

func createDirIfNotExist() string {

	t := time.Now()
	path := t.Format("2006-01-02")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0755)
	}
	return path + "/"
}

func main() {
	longurl := "https://cloud-api.yandex.net:443/v1/disk/public/resources/download?public_key=DhLa7f6nRVrD8AZj9EGmFkyE8goTvQr0vPDb6WsdgtQ%3D&path=%2Fhomework%2F"
	allFiles := []requestFile{
		{longurl + "vocabulary%2Fru-en.html",
			"ru-en.html"},
		{longurl + "vocabulary%2Fen-ru.html",
			"en-ru.html"},
		{longurl + "homework-analysis%2FYuliya.pdf",
			"Yuliya.pdf"},
		{longurl + "homework-analysis%2FAydar.pdf",
			"Aydar.pdf"},
		{longurl + "irregular-verbs%2Ffollow-and-click.html",
			"follow-and-click.html"},
		{longurl + "irregular-verbs%2Forder.pdf",
			"order.pdf"},
		{longurl + "irregular-verbs%2Fpractice-and-check.html",
			"practice-and-check.html"},
		{longurl + "irregular-verbs%2Fwords.mp3",
			"words.mp3"},
	}

	for _, o := range allFiles {
		s, err := getObject(o)
		if err != nil {
			log.Println(err)
		} else {
			path2save := createDirIfNotExist()
			saveObject(s, path2save+o.Filename)
		}
	}
}
