package main

import (
	"fmt"
	"log"
	"net/http"
)

// The flow of this project will be that we create the server, the server creates routes for us to traverse through, and through functions will we
// get to those locations
// Create two basic html functions within the folder: index.html and form.html. Afterwards, write in main

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil { // submitting a form will make a POST request that will parse the form
		fmt.Fprintf(w, "ParseForm() error: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful.\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name: John Doe %s\n", name)
	fmt.Fprintf(w, "Address: 123 Apple Road %s\n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) { // a send and request interaction implemented
	if r.URL.Path != "/hello" {
		http.Error(w, "404 error occurred", http.StatusNotFound) // if something other than hello is ran
		return
	}

	if r.Method != "GET" { // In our handlefunc, /hello causes an HTTP GET request. In the case something incorrect occurs, we implement a fix
		http.Error(w, "Method not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello!") // Check later to see if a normal print statement would or wouldn't work on this.

}

func main() {
	fileServer := http.FileServer(http.Dir("./static")) // instructs goloang to check the static directory provided
	http.Handle("/", fileServer)                        // The "/" handles the default html
	http.HandleFunc("/form", formHandler)               // handles the form html file
	http.HandleFunc("/hello", helloHandler)             // handles the index html

	// Start server with error handling provided
	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil { // ListenAndServe starts an HTTP server
		log.Fatal(err)
	}
}

// After running your file, travel to a browser and provide "localhost:8080/" or "localhost:8080/hello" & "localhost:8080/form"
// to see your end results. Ctrl-c within the terminal closes the port.
