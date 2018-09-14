package main

import (
	"fmt"
	"os"
	"zsh-go-auto/completions"
	"zsh-go-auto/scrapper"

	"zsh-go-auto/generator"
	"zsh-go-auto/persister"
)

func main() {
	file, err := os.Create("./_gcloud")
	cacheFile := "./cache.tmp"
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	p := persister.NewDiskPersister()

	var gcloudDocs completions.AutoComplete
	err = p.Load(cacheFile, &gcloudDocs)
	if err != nil {
		gcloudDocs = *scrapper.Run()
		p.Save(cacheFile, gcloudDocs)
	}

	generator.Run(file, gcloudDocs)
}
