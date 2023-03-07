/*
Create time at 2023年3月7日0007下午 18:48:42
Create User at Administrator
*/

package main

import (
	"fmt"
	"log"
	"net/http"
)

func hellofunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method not supported", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "hello")
}
func formfunc(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parse error: %v", err)
		return
	}
	fmt.Fprintf(w, "Parse successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "name: %v, address: %v", name, address)
	return
}
func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", hellofunc)
	http.HandleFunc("/form", formfunc)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
