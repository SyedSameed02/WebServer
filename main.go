package main

import ( 
	"log"
	"fmt"
	"net/http"	
 )

func helloHandler(w http.ResponseWriter , r *http.Request){

	if r.URL.Path != "/hello"{
	http.Error(w,"404 not found",http.StatusNotFound)
	return
	}
	if r.Method != "GET" {
		 http.Error(w,"method not supported",http.StatusNotFound)
		 return
	}
	fmt.Fprintf(w,"hello!")
}

func formHandler(w http.ResponseWriter, r *http.Request){

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w,"parseForm() err: %v",err)
		return
	}
	fmt.Fprintf(w,"Post request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w,"name = %s\n",name)
	fmt.Fprintf(w,"address = %s\n",address)
}



func main () {

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/",fileServer)

	http.HandleFunc("/form",formHandler)
	http.HandleFunc("/hello",helloHandler)

	fmt.Printf("Starting server at port 8080\n")

	if err := http.ListenAndServe(":8080",nil); err != nil {
		log.Fatal(err)
	}
}