package router

import (
	"fmt"
	"strings"
	"time"
	// "html/template"
	"net/http"
	"main/internal/rod"
)

func (router *Router) ScrapeJobDescription(w http.ResponseWriter, r *http.Request){
	fmt.Println(r.Method, " request on /scrape/jobdesc")

	r.ParseForm()
	jobPostURL := r.FormValue("posting-url")
	fmt.Println("jobPostURL: ", jobPostURL)

	browser := rod.GetBrowser()

  page := browser.MustPage(jobPostUrl).MustWaitLoad().MustWaitIdle(2 * time.Second)

    // Wait for the full job description container
    el := page.MustElement("#jobDescriptionText")
    desc := strings.TrimSpace(el.MustText())

    // Just return as plain string
    w.Header().Set("Content-Type", "text/plain; charset=utf-8")
    _, _ = w.Write([]byte(desc))
}

// helper to normalize Indeed links
func normalizeIndeedURL(u string) string {
    parsed, err := url.Parse(u)
    if err != nil {
        return u
    }
    q := parsed.Query()
    if vjk := q.Get("vjk"); vjk != "" {
        return "https://www.indeed.com/viewjob?jk=" + vjk
    }
    if jk := q.Get("jk"); jk != "" {
        return "https://www.indeed.com/viewjob?jk=" + jk
    }
    return u
}


