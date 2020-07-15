package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type PageVariables struct {
	PageTitle string
}

func main() {
	res, _ := http.Get("https://api.ipify.org")
	ip, _ := ioutil.ReadAll(res.Body)
	os.Stdout.Write(ip)
	http.HandleFunc("/", HomePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HomePage(w http.ResponseWriter, r *http.Request) {

	Title := "Terramino"
	MyPageVariables := PageVariables{
		PageTitle: Title,
	}

	t, err := template.ParseFiles("index.html") //parse the html file index.html
	if err != nil {                             // if there is an error
		log.Print("template parsing error: ", err) // log it
	}

	err = t.Execute(w, MyPageVariables) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil {                     // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}
