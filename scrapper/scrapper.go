package scrapper

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"zsh-go-auto/completions"

	"github.com/PuerkitoBio/goquery"
)

func Run() *completions.AutoComplete {
	groups, commands, flags := parsePage("https://cloud.google.com/sdk/gcloud/reference/")

	// this should be recursive not just one deep it could use go routines im sure google can handle it
	// and there are alot of them :P
	populateGroups(groups)

	return &completions.AutoComplete{
		Name:        "gcloud",
		Description: "gcloud bla bla bla",
		Groups:      groups,
		Commands:    commands,
		Flags:       flags,
	}
}

func populateGroups(groups []*completions.Group) {
	sliceLength := len(groups)
	var wg sync.WaitGroup
	wg.Add(sliceLength)
	for _, g := range groups {
		go func(g *completions.Group) {
			defer wg.Done()
			sg, sc, sf := parsePage(g.URI)
			g.Groups = sg
			g.Commands = sc
			g.Flags = sf
			if len(sg) > 0 {
				populateGroups(sg)
			}
			if len(sc) > 0 {
				populateCommands(sc)
			}
		}(g)
	}

	wg.Wait()
}

func populateCommands(commands []*completions.Command) {
	sliceLength := len(commands)
	var wg sync.WaitGroup
	wg.Add(sliceLength)
	for _, c := range commands {
		go func(c *completions.Command) {
			defer wg.Done()
			_, _, sa := parsePage(c.URI)
			c.Flags = sa
		}(c)
	}

	wg.Wait()
}

func parsePage(uri string) (groups []*completions.Group, commands []*completions.Command, flags []*completions.Flag) {
	res, err := http.Get(uri)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find("#GROUPS dd dl").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		s.Children().Each(func(x int, c *goquery.Selection) {
			if c.Is("dt") {
				uri, _ := c.Find("a").Attr("href")
				groupName := strings.Replace(c.Text(), "-", "_", -1)
				fmt.Printf("group - %v \n", uri)
				h := c.Next()
				g := &completions.Group{
					Name:        groupName,
					Description: strings.Trim(h.Text(), "\n"),
					URI:         uri,
				}
				groups = append(groups, g)
				//fmt.Println(c.Text() + " - " + h.Text())

			}
		})

	})

	doc.Find("#COMMANDS dd dl").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		s.Children().Each(func(x int, c *goquery.Selection) {
			if c.Is("dt") {
				uri, _ := c.Find("a").Attr("href")
				commandName := strings.Replace(c.Text(), "-", "_", -1)
				fmt.Printf("command - %v \n", uri)
				h := c.Next()
				g := &completions.Command{
					Name:        commandName,
					Description: strings.Trim(h.Text(), "\n"),
					URI:         uri,
				}
				commands = append(commands, g)
				//fmt.Println(c.Text() + " - " + h.Text())
			}
		})

	})

	doc.Find("#FLAGS dd dl").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		s.Children().Each(func(x int, c *goquery.Selection) {
			if c.Is("dt") {
				commandName := strings.Replace(c.Text(), "-", "_", -1)
				h := c.Next()
				g := &completions.Flag{
					Name:        commandName,
					Description: strings.Trim(h.Text(), "\n"),
				}
				flags = append(flags, g)
				//fmt.Println(c.Text() + " - " + h.Text())
			}
		})

	})

	return groups, commands, flags
}
