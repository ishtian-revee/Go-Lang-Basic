package main

import(
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/xml"
)

//var post = []byte(`
//<sitemapindex xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
//	<sitemap>
//		<loc>
//		http://www.washingtonpost.com/news-politics-sitemap.xml
//		</loc>
//	</sitemap>
//	<sitemap>
//		<loc>
//		http://www.washingtonpost.com/news-blogs-politics-sitemap.xml
//		</loc>
//	</sitemap>
//	<sitemap>
//		<loc>
//		http://www.washingtonpost.com/news-opinions-sitemap.xml
//		</loc>
//	</sitemap>
//</sitemapindex>
//`)

//type Location struct {
//	Loc string `xml:"loc"`
//}
//
//type SitemapIndex struct{
//	Locations []Location `xml:"sitemap"`	// it's a slice
//}
//
//// value receiver
//func (l Location) String() string{
//
//	return fmt.Sprintf(l.Loc)
//}

// instead of writing above 3 pieces of codes we can write this struct
type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Titles []string `xml:"url>news>title"`
	Keywords []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct{
	Keyword string
	Location string
}

func main(){

	var s SitemapIndex
	var n News
	news_map := make(map[string]NewsMap)

	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	//resp.Body.Close()

	//bytes := post
	xml.Unmarshal(bytes, &s)

	//fmt.Println(s.Locations)

	// looping through each contents from the slice and printing it
	for _, Location := range s.Locations {

		resp, _ := http.Get(Location)
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &n)

		for idx, _ := range n.Titles{

			news_map[n.Titles[idx]] = NewsMap{n.Keywords[idx], n.Locations[idx]}
		}
	}

	for idx, data := range news_map{

		fmt.Println("\n\n\n", idx)
		fmt.Println(data.Keyword)
		fmt.Println("\n", data.Location)
	}
}
