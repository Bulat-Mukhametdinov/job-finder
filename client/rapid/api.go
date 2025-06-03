package rapid

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Job struct {
	JobID         string  `json:"job_id"`
	JobTitle      string  `json:"job_title"`
	EmployerName  string  `json:"employer_name"`
	MinSalary     float64 `json:"job_min_salary"`
	MaxSalary     float64 `json:"job_max_salary"`
	JobApplyLink  string  `json:"job_apply_link"`
	JobIsRemote   bool    `json:"job_is_remote"`
	JobCity       string  `json:"job_city"`
	JobState      string  `json:"job_state"`
	JobPostTime   string  `json:"job_posted_at"`
	EmployerLogo  string  `json:"employer_logo"`
	JobHighlights struct {
		Qualifications   []string `json:"Qualifications"`
		Benefits         []string `json:"Benefits"`
		Responsibilities []string `json:"Responsibilities"`
	} `json:"job_highlights"`
}

type JobResponse struct {
	Data []Job `json:"data"`
}

func GetJob(query string, page string, num_pages string, country string, language string, date_posted string,
	work_from_home string, job_requirements string, exclude_job_publishers string, fields string) []Job {

	url, err := buildSearchURL(query, page, num_pages, country, language, date_posted,
		work_from_home, job_requirements, exclude_job_publishers, fields)

	if err != nil {
		log.Fatal("Error creating url:", err)
	}

	rapidapiKey := os.Getenv("RAPID_API_KEY")
	if rapidapiKey == "" {
		log.Fatal("RAPID_API_KEY not fond in .env")
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Error in requesting:", err)
	}

	req.Header.Add("x-rapidapi-key", rapidapiKey)
	req.Header.Add("x-rapidapi-host", "jsearch.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("Ошибка при выполнении запроса:", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Ошибка чтения тела ответа:", err)
	}
	var resp JobResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		log.Fatal("Ошибка при разборе JSON:", err)
	}

	return resp.Data

	// for _, job := range resp.Data {
	// 	fmt.Println("Title:", job.JobTitle)
	// 	fmt.Println("City:", job.JobCity)
	// 	if job.MinSalary != nil && job.MaxSalary != nil {
	// 		fmt.Printf("Salary: %.2f - %.2f\n", *job.MinSalary, *job.MaxSalary)
	// 	}
	// 	fmt.Println("Remote:", job.JobIsRemote)
	// 	fmt.Println("Apply link:", job.JobApplyLink)
	// 	fmt.Println("Qualifications:")
	// 	for i, qual := range job.JobHighlights.Qualifications {
	// 		fmt.Printf("  %d. %s\n", i+1, qual)
	// 	}

	// 	fmt.Println("Benefits:")
	// 	for i, benefit := range job.JobHighlights.Benefits {
	// 		fmt.Printf("  %d. %s\n", i+1, benefit)
	// 	}

	// 	fmt.Println("Responsibilities:")
	// 	for i, resp := range job.JobHighlights.Responsibilities {
	// 		fmt.Printf("  %d. %s\n", i+1, resp)
	// 	}
	// 	fmt.Println()
	// }
}

func buildSearchURL(
	query string,
	page string,
	numPages string,
	country string,
	language string,
	datePosted string,
	workFromHome string,
	jobRequirements string,
	excludeJobPublishers string,
	fields string,
) (string, error) {
	baseURL := "https://jsearch.p.rapidapi.com/search"

	params := url.Values{}

	if strings.TrimSpace(query) != "" {
		params.Add("query", query)
	}

	if page != "" && page != "1" {
		params.Add("page", page)
	}
	if numPages != "" && numPages != "1" {
		params.Add("num_pages", numPages)
	}
	if country != "" && strings.ToLower(country) != "us" {
		params.Add("country", country)
	}
	if language != "" {
		params.Add("language", language)
	}
	if datePosted != "" && strings.ToLower(datePosted) != "all" {
		params.Add("date_posted", datePosted)
	}
	if workFromHome != "" && strings.ToLower(workFromHome) != "false" {
		params.Add("work_from_home", workFromHome)
	}
	if jobRequirements != "" {
		params.Add("job_requirements", jobRequirements)
	}
	if excludeJobPublishers != "" {
		params.Add("exclude_job_publishers", excludeJobPublishers)
	}
	if fields != "" {
		params.Add("fields", fields)
	}

	fullURL := baseURL
	if len(params) > 0 {
		fullURL += "?" + params.Encode()
	}

	return fullURL, nil
}

func ParseJobs(jsonData []byte) ([]Job, error) {
	var jobResp JobResponse
	err := json.Unmarshal(jsonData, &jobResp)
	if err != nil {
		return nil, err
	}
	return jobResp.Data, nil
}
