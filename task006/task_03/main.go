package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

func loadPage(title string) (*Page, error) {
	filename := title + ".html"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	// send Not found in such case
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	if r.Method == "GET" {
		title := "form"
		p, _ := loadPage(title)
		fmt.Fprintf(w, "%s", p.Body)
	} else if r.Method == "POST" {
		// get new value for name and address
		name := r.PostFormValue("name")
		address := r.PostFormValue("address")
		// creating a cookie
		cookie := &http.Cookie{
			Name:   "token",
			Value:  name + ":" + address,
			MaxAge: 300,
		}
		// setting the cookie
		http.SetCookie(w, cookie)
		title := "form"
		c, _ := loadPage(title)
		fmt.Fprintf(w, "%s", c.Body)
	} else {
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func main() {
	// server port number
	const port = 8080

	fmt.Printf("Launching server on port: %d \n\n", port)

	// set handler for route '/'
	http.HandleFunc("/", handler)
	// start server without ending
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
