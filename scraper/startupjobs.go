package scraper

import (
	"Gojob-scraper/models"
	"fmt"

	"github.com/gocolly/colly"
)


func ScrapeStartupJobs() []models.Job {
	jobs := []models.Job{}
	c := colly.NewCollector()

	c.OnHTML("div.job-listing", func(e *colly.HTMLElement) {
			title := e.ChildText("a.job-title")
			company := e.ChildText("div.company")
			location := e.ChildText("div.location")
			url := e.ChildAttr("a.job-title", "href")

			job := models.Job{
					Title:      title,
					Company:    company,
					Location:   location,
					URL:        "https://startup.jobs" + url,
					DatePosted: "",
					Source:     "Startup Jobs",
			}

			jobs = append(jobs, job)
			fmt.Printf("Found job: %+v\n", job)
	})

	c.Visit("https://startup.jobs/")

	return jobs
}
