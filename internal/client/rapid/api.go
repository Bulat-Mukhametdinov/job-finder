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
	IsFavourite   bool
	JobComment    string
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
	SalaryPeriod  string  `json:"job_salary_period"`
	JobHighlights struct {
		Qualifications   []string `json:"Qualifications"`
		Benefits         []string `json:"Benefits"`
		Responsibilities []string `json:"Responsibilities"`
	} `json:"job_highlights"`
}

type JobResponse struct {
	Data []Job `json:"data"`
}

type RapidAPI struct {
	baseURL string
	apiKey  string
}

func NewRapidAPI() *RapidAPI {
	rapidapiKey := os.Getenv("RAPID_API_KEY")
	if rapidapiKey == "" {
		log.Println("RAPID_API_KEY not found in .env")
		return nil
	}
	return &RapidAPI{"https://jsearch.p.rapidapi.com/", rapidapiKey}
}

func (r RapidAPI) GetJob(job_id string) (Job, error) {
	url := r.baseURL + "job-details?job_id=" + job_id

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error creating request:", err)
		return Job{}, err
	}

	req.Header.Add("x-rapidapi-key", r.apiKey)
	req.Header.Add("x-rapidapi-host", "jsearch.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error executing request:", err)
		return Job{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return Job{}, err
	}
	var resp JobResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		log.Println("Error parsing JSON:", err)
		return Job{}, err
	}

	if len(resp.Data) != 1 {
		log.Println("Error in retrieving valid answer from rapid api:", string(body))
		return Job{}, nil
	}
	return resp.Data[0], nil
}

func (r RapidAPI) GetJobs(query string, page string, numPages string, country string, language string, datePosted string,
	workFromHome string, jobRequirements string, excludeJobPublishers string, fields string) ([]Job, error) {

	url, err := r.buildSearchURL(query, page, numPages, country, language, datePosted,
		workFromHome, jobRequirements, excludeJobPublishers, fields)

	if err != nil {
		log.Println("Error creating URL:", err)
		return []Job{}, err
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error creating request:", err)
		return []Job{}, err
	}

	req.Header.Add("x-rapidapi-key", r.apiKey)
	req.Header.Add("x-rapidapi-host", "jsearch.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error executing request:", err)
		return []Job{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return []Job{}, err
	}
	var resp JobResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		log.Println("Error parsing JSON:", err)
		return []Job{}, err
	}

	if len(resp.Data) == 0 {
		log.Println("No jobs found:", string(body))
	}
	return resp.Data, nil
}

func (r RapidAPI) buildSearchURL(
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

	fullURL := r.baseURL + "search"
	if len(params) > 0 {
		fullURL += "?" + params.Encode()
	}

	return fullURL, nil
}
