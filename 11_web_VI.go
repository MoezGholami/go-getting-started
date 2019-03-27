package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "strings"
    "encoding/xml"
    "html/template"
)

type SitemapIndex struct {
    Locations []string `xml:"sitemap>loc"`
}

type Feed struct {
    Titles []string `xml:"url>news>title"`
    Keywords []string `xml:"url>news>keywords"`
    Locations []string `xml:"url>loc"`
}

type NewsContent struct {
    Keywords string
    Location string
}

type NewsAggregationPageContent struct {
    Title string
    News map[string]NewsContent
}

func (nc *NewsAggregationPageContent) showHandler(w http.ResponseWriter, r *http.Request) {
    t, err := template.ParseFiles("11_web_template.html")
    fmt.Println("rendering error: ", err)
    t.Execute(w, *nc)
}

func main() {
    response, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
    bytes, _ := ioutil.ReadAll(response.Body)
    response.Body.Close()

    var feed Feed
    var s SitemapIndex
    var newsMap map[string]NewsContent = make(map[string]NewsContent)
    xml.Unmarshal(bytes, &s)

    for _, Location := range(s.Locations) {
        Location = strings.TrimSpace(Location)
        fmt.Println("fetching ", "\""+Location+"\"", " ...")
        response, _ := http.Get(Location)
        bytes, _ := ioutil.ReadAll(response.Body)
        response.Body.Close()
        xml.Unmarshal(bytes, &feed)

        for i, title := range feed.Titles {
            newsMap[title] = NewsContent{Keywords: feed.Keywords[i], Location: feed.Locations[i]}
        }
    }

    var pageContent NewsAggregationPageContent = NewsAggregationPageContent{Title: "All News", News: newsMap}
    http.HandleFunc("/", pageContent.showHandler)
    fmt.Println("going to serve at http://localhost:4000 ")
    http.ListenAndServe(":4000", nil)
}
