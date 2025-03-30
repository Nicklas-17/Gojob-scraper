package models

type Job struct {
	Title 			string `json:"Title"`
	Company 		string `json:"Company"`
	Location 		string `json:"Location"`
	URL 				string `json:"Url"`
	DatePosted 	string `json:"Date_Posted"`
	Source 			string `json:"Source"`
}