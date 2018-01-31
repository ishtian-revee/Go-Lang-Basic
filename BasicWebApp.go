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

type Location struct {
	Loc string `xml:"loc"`
}

type SitemapIndex struct{
	Locations []Location `xml:"sitemap"`	// it's a slice
}

// value receiver
func (l Location) String() string{

	return fmt.Sprintf(l.Loc)
}

func main(){

	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	//bytes := post
	var s SitemapIndex
	xml.Unmarshal(bytes, &s)

	fmt.Println(s.Locations)
}
