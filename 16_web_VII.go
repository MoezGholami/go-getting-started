// this is improvement of web example VI (11_web_VI.go) using channels
// significant improvement (though downloading sucks in go)
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

func getNewsFeedOfANewsMapLocation(resultChannel chan *Feed, location string) {
    var feed Feed
    location = strings.TrimSpace(location)
    fmt.Println("fetching ", "\""+location+"\"", " ...")
    response, _ := http.Get(location)
    bytes, _ := ioutil.ReadAll(response.Body)
    response.Body.Close()
    xml.Unmarshal(bytes, &feed)
    resultChannel <- &feed
}

func main() {
    response, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
    bytes, _ := ioutil.ReadAll(response.Body)
    response.Body.Close()

    var s SitemapIndex
    var newsMap map[string]NewsContent = make(map[string]NewsContent)
    xml.Unmarshal(bytes, &s)
    feedChannel := make(chan *Feed)

    for _, location := range(s.Locations) { go getNewsFeedOfANewsMapLocation(feedChannel, location) }
    for range(s.Locations) {
        feed := <- feedChannel
        for i, title := range feed.Titles {
            newsMap[title] = NewsContent{Keywords: feed.Keywords[i], Location: feed.Locations[i]}
        }
    }

    var pageContent NewsAggregationPageContent = NewsAggregationPageContent{Title: "All News", News: newsMap}
    http.HandleFunc("/", pageContent.showHandler)
    fmt.Println("going to serve at http://localhost:4000 ")
    http.ListenAndServe(":4000", nil)
}
