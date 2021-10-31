package scrapper

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	id          string
	title       string
	companyName string
	location    string
	summary     string
}

// Scrape Indeed by a term
func Scrape(term string) {
	var baseURL string = "https://kr.indeed.com/jobs?q=" + term + "&limit=50"
	var jobs []extractedJob
	mainC := make(chan []extractedJob)
	totalPages := getPages(baseURL)

	for i := 0; i < totalPages; i++ {
		go getPage(baseURL, i, mainC)
	}

	for i := 0; i < totalPages; i++ {
		extractedJobs := <-mainC
		jobs = append(jobs, extractedJobs...)
	}

	writeJobs(jobs)
	fmt.Println("Done, extracted", len(jobs))
}

func getPage(url string, page int, mainC chan<- []extractedJob) {
	var jobs []extractedJob
	c := make(chan extractedJob)
	pageURL := url + "&start=" + strconv.Itoa(page*50)
	fmt.Println(pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkStatusCode(res)

	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	tapItems := doc.Find(".tapItem")
	tapItems.Each(func(i int, item *goquery.Selection) {
		go extractJob(item, c)
	})

	for i := 0; i < tapItems.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}
	mainC <- jobs
}

func extractJob(item *goquery.Selection, c chan<- extractedJob) {
	id, _ := item.Attr("data-jk")
	title := cleanString(item.Find(".jobTitle>span").Text())
	companyName := cleanString(item.Find(".companyName").Text())
	location := cleanString(item.Find(".companyLocation").Text())
	summary := cleanString(item.Find(".job-snippet").Text())
	c <- extractedJob{
		id:          id,
		title:       title,
		companyName: companyName,
		location:    location,
		summary:     summary}
}

func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

func getPages(url string) int {
	pages := 0
	res, err := http.Get(url)
	checkErr(err)
	checkStatusCode(res)

	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})

	return pages
}

func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"ID", "TITLE", "COMPANYNAME", "LOCATION", "SUMMARY"}
	wErr := w.Write(headers)
	checkErr(wErr)

	jobSlice := make(chan []string)

	for _, job := range jobs {
		go sliceJob(job, jobSlice)
	}

	for i := 0; i < len(jobs); i++ {
		jwErr := w.Write(<-jobSlice)
		checkErr(jwErr)
	}
}

func sliceJob(job extractedJob, jobSlice chan<- []string) {
	jobSlice <- []string{job.id, job.title, job.companyName, job.location, job.summary}
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkStatusCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status:", res.StatusCode)
	}
}
