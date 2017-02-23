package main

import (
	"log"
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

func moveIT(from, to string) error {
	err := os.Rename(from, to)
	return err
}

func path() string {
	t := time.Now()
	path := "files/" + t.Format("2006-01-02") + "/"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0755)
		t = t.Add(-24 * time.Hour)
		moveIT("files/"+t.Format("2006-01-02"), "files/0ld/"+t.Format("2006-01-02"))
	}
	return path
}

func save2Path(r requestFile, wg *sync.WaitGroup) {
	s, err := getObject(r)
	if err != nil {
		log.Println(err)
	} else {
		if checkExist(s, r) == true {
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
		//===========================
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
		wg.Add(1)
		go save2Path(o, &wg)
		wg.Wait()
	}
}
