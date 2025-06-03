package main

import (
	"fmt"
	"job-finder/internal/client/rapid"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	query := "G0Lang"
	page := "1"
	num_pages := "1"
	country := "us"
	language := ""
	date_posted := "all"
	work_from_home := "false"
	job_requirements := ""
	exclude_job_publishers := ""
	fields := ""

	resp := rapid.GetJob(query, page, num_pages, country, language, date_posted,
		work_from_home, job_requirements, exclude_job_publishers, fields)
	fmt.Println(resp)
}
