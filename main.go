package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

type requestFile struct {
	URL      string
	Filename string
	SubPath  string
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
		log.Printf("%s", err)
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
}

func path() string {
	t := time.Now()
	path := "files/" + t.Format("2006-01-02") + "/"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0755)
	}
	return path
}

func fileExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func save2Path(r requestFile, wg *sync.WaitGroup) {
	s, err := getObject(r)
	if err != nil {
		log.Println(err)
	} else {
		switch r.SubPath {
		case "Aydar":
			os.Mkdir(path()+"Aydar/", 0755)
			saveObject(s, path()+"Aydar/"+r.Filename)
			log.Println("file " + r.Filename + " downlowaded (for Aydar)")
		case "Yuliya":
			os.Mkdir(path()+"Yuliya/", 0755)
			saveObject(s, path()+"Yuliya/"+r.Filename)
			log.Println("file " + r.Filename + " downlowaded (for Yuliya)")
		default:
			saveObject(s, path()+r.Filename)
			log.Println("file " + r.Filename + " downlowaded")
		}
	}
	defer wg.Done()
}

func main() {
	longurl := "https://cloud-api.yandex.net:443/v1/disk/public/resources/download?public_key=DhLa7f6nRVrD8AZj9EGmFkyE8goTvQr0vPDb6WsdgtQ%3D&path=%2Fhomework%2F"
	allFiles := []requestFile{
		{longurl + "vocabulary%2Fru-en.html",
			"ru-en.html", ""},
		{longurl + "vocabulary%2Fen-ru.html",
			"en-ru.html", ""},
		{longurl + "homework-analysis%2FYuliya.pdf",
			"Yuliya.pdf", ""},
		{longurl + "homework-analysis%2FAydar.pdf",
			"Aydar.pdf", ""},
		{longurl + "irregular-verbs%2Ffollow-and-click.html",
			"follow-and-click.html", ""},
		{longurl + "irregular-verbs%2Forder.pdf",
			"order.pdf", ""},
		{longurl + "irregular-verbs%2Fpractice-and-check.html",
			"practice-and-check.html", ""},
		{longurl + "irregular-verbs%2Fwords.mp3",
			"words.mp3", ""},
		{longurl + "exercises.jpg",
			"exercises.jpg", ""},
		{longurl + "pronunciation%2FAydar%2Fconfusable.pdf",
			"confusable.pdf", "Aydar"},
		{longurl + "pronunciation%2FAydar%2Ffollow-and-click.html",
			"follow-and-click.html", "Aydar"},
		{longurl + "pronunciation%2FAydar%2Fn-back.mp3",
			"n-back.mp3", "Aydar"},
		{longurl + "pronunciation%2FAydar%2Fpractice-and-check.html",
			"practice-and-check.html", "Aydar"},
		{longurl + "pronunciation%2FAydar%2Fpronunciation.pdf",
			"pronunciation.pdf", "Aydar"},
		{longurl + "pronunciation%2FAydar%2Fsounds.mp3",
			"sounds.mp3", "Aydar"},
		{longurl + "pronunciation%2FAydar%2Fwords.mp3",
			"words.mp3", "Aydar"},
		//===========================
		{longurl + "pronunciation%2FYuliya%2Fconfusable.pdf",
			"confusable.pdf", "Yuliya"},
		{longurl + "pronunciation%2FYuliya%2Ffollow-and-click.html",
			"follow-and-click.html", "Yuliya"},
		{longurl + "pronunciation%2FYuliya%2Fn-back.mp3",
			"n-back.mp3", "Yuliya"},
		{longurl + "pronunciation%2FYuliya%2Fpractice-and-check.html",
			"practice-and-check.html", "Yuliya"},
		{longurl + "pronunciation%2FYuliya%2Fpronunciation.pdf",
			"pronunciation.pdf", "Yuliya"},
		{longurl + "pronunciation%2FYuliya%2Fsounds.mp3",
			"sounds.mp3", "Yuliya"},
		{longurl + "pronunciation%2FYuliya%2Fwords.mp3",
			"words.mp3", "Yuliya"},
	}
	var wg sync.WaitGroup

	// TODO rewrite as multithread func
	for _, o := range allFiles {
		if fileExist(path()+o.SubPath+"/"+o.Filename) == false {
			wg.Add(1)
			go save2Path(o, &wg)
		} else {
			log.Println("file " + o.Filename + " is exist")
		}
		wg.Wait()
	}
}
