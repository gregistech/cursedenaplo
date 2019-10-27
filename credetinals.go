package main

import "os"

func GetCredetinals() (string, string, string) {
	return os.Getenv("CK_INST"), os.Getenv("CK_USERNAME"), os.Getenv("CK_PASSWORD")
}
