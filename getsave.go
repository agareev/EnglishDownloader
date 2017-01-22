package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// GetAllContents internal function for getObject
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

// --------------------------------------------

func saveObject(b []byte, f string) {
	err := ioutil.WriteFile(f, b, 0644)
	if err != nil {
		log.Println(err)
	}
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
