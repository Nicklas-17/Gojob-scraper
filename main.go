package main

import (
	"Gojob-scraper/scraper"
	"fmt"
)


func main() {
	rocketShipjobs := scraper.ScrapeRemoteRocketship()
	startupJobs := scraper.ScrapeStartupJobs()

	alljobs := append(rocketShipjobs, startupJobs...)

	fmt.Println(alljobs)
}