package main

import ("fmt"
"net/http"
"io/ioutil"
"strings"
"encoding/xml")

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

    for titles, content := range newsMap {
        fmt.Println(titles, content)
    }
}
