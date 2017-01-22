package main

import (
	"log"
	"os"
)

func fileExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func hashreturn(file []byte) string {
	return "xxxx"
}

// Fix this function
func checkExist(s []byte, r requestFile) bool {
	log.Println(path() + r.SubPath + "/" + r.Filename)
	if fileExist(path()+r.SubPath+"/"+r.Filename) == false {
		if hashreturn(s) == hashreturn(s) {
			return true
		}
		return false
	}
	log.Println("file " + r.Filename + " is exist")
	return false
}
