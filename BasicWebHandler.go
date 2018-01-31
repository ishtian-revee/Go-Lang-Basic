package main

import(
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request){

	//fmt.Fprint(w, "Index page !")

	fmt.Fprintf(w, "<h1>Index page</h1>")
	fmt.Fprintf(w, "<h4>Let's GO !!!</h4>")
	fmt.Fprintf(w, "<p>Go is awesome!</p>")
}

func about_handler(w http.ResponseWriter, r *http.Request){

	//fmt.Fprintf(w, "About page !")

	fmt.Fprintf(w, `
<h1>About page</h1>
<h4>Let's GO !!!</h4>
<p>Go is awesome!</p>`)
}

func main(){

	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about/", about_handler)
	http.ListenAndServe(":8000", nil)
}
