package scraper

import (
	"Gojob-scraper/models"
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func ScrapeRemoteRocketship() []models.Job {
	jobs := []models.Job{}
	c := colly.NewCollector()

	c.OnHTML("h3.text-lg.font-semibold.text-primary.mr-4 a", func(e *colly.HTMLElement) {
		title := e.Text
		company := e.ChildText("span.text-sm")
		url := e.Attr("href")
		location := extractLocationFromURL(url)

		datePosted := e.ChildText("span.text-xs.text-gray-500")

		job := models.Job{
			      Title:      title,
            Company:    company,
            Location:   location,
            URL:        "https://www.remoterocketship.com" + url,
            DatePosted: datePosted,
            Source:     "Remote Rocketship",
		}

		jobs = append(jobs, job)
		fmt.Printf("Found job: %+v\n", job)
	})

	c.Visit("https://www.remoterocketship.com/?page=1&sort=DateAdded&locations=Belgium")
	return jobs
}


func extractLocationFromURL(url string) string {

	parts := strings.Split(url, "/")

	for _, part := range parts {

		if strings.Contains(part, "europe") || strings.Contains(part, "remote") || strings.Contains(part, "worldwide") {
			return part
		}
	}


	return "Unknown"}