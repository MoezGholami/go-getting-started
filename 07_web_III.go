package main

import ("fmt"
"net/http"
"io/ioutil"
"strings"
"encoding/xml")

type SitemapIndex struct {
    Locations []string `xml:"sitemap>loc"`
}

type News struct {
    Titles []string `xml:"url>news>title"`
    Keywords []string `xml:"url>news>keywords"`
    Locations []string `xml:"url>loc"`
}

func main() {
    response, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
    bytes, _ := ioutil.ReadAll(response.Body)
    response.Body.Close()

    var s SitemapIndex
    xml.Unmarshal(bytes, &s)

    for _, Location := range(s.Locations) {
        Location = strings.TrimSpace(Location)
        fmt.Println("fetching ", "\""+Location+"\"", " ...")
        response, _ := http.Get(Location)
        bytes, _ := ioutil.ReadAll(response.Body)
        response.Body.Close()
        var n News
        xml.Unmarshal(bytes, &n)
    }
}
