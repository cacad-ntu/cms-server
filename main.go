package main

import (
	"log"
	"flag"
	"github.com/gorilla/mux"
	"encoding/json"
	"strings"
	"strconv"
	"github.com/rs/cors"
	"./storage"
	"mime/multipart"
    "net/http"
	"os"
	"fmt"
	"bytes"
	"math/rand"
	"time"
	"io"
	"path/filepath"
	"./models"
)

func random(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}

func Upload(url, file string) (err error) {
    // Prepare a form that you will submit to that URL.
    var b bytes.Buffer
    w := multipart.NewWriter(&b)
	// Add your image file
	
    f, err := os.Open(file)
    if err != nil {
        return 
    }
    defer f.Close()
    fw, err := w.CreateFormFile("image", file)
    if err != nil {
        return 
    }
    if _, err = io.Copy(fw, f); err != nil {
        return
    }
    // Add the other fields
    if fw, err = w.CreateFormField("name"); err != nil {
        return
    }
    if _, err = fw.Write([]byte("new event here")); err != nil {
        return
	}
	fw, _ = w.CreateFormField("location");
	fw.Write([]byte("NTU"))
	
	fw, _ = w.CreateFormField("eventType");
	fw.Write([]byte("free form"))

	fw, _ = w.CreateFormField("crowd");
	fw.Write([]byte("123"))
    // Don't forget to close the multipart writer.
    // If you don't close it, your request will be missing the terminating boundary.
    w.Close()

    // Now that you have a form, you can submit it to your handler.
    req, err := http.NewRequest("POST", url, &b)
    if err != nil {
        return 
    }
    // Don't forget to set the content type, this will contain the boundary.
    req.Header.Set("Content-Type", w.FormDataContentType())

    // Submit the request
    client := &http.Client{}
    res, err := client.Do(req)
    if err != nil {
        return 
    }

    // Check the response
    if res.StatusCode != http.StatusOK {
        err = fmt.Errorf("bad status: %s", res.Status)
    }
    return
}


func studyPlaceHandler(rw http.ResponseWriter, req *http.Request) {
	db, _ := storage.NewDB("db.sqlite3")
	studyPlaces, _:= db.ListAllStudyPlaces()
	for ind, _ := range studyPlaces {
		studyPlaces[ind].Images = []string{studyPlaces[ind].Image}
	}
	result, _ := json.Marshal(studyPlaces)
	
	rw.Header().Set("Content-Type", "application/json")
		
	rw.Write(result)
}

func locationHandler(rw http.ResponseWriter, req *http.Request) {
	db, _ := storage.NewDB("db.sqlite3")
	param := req.URL.Query()["id"][0];
	
	locationID, _ := strconv.Atoi(param)
	studyPlace, _ := db.GetStudyPlaceById(locationID)
	
	studyPlace.Images = []string{studyPlace.Image}
	studyPlace.Timestamp = make([]string, 0)
	studyPlace.HistoricalDensity = [][]int{make([]int, 0), make([]int, 0)}
	history, _ := db.GetLocationHistoriesById(locationID)
	
	for _, elem := range history {
		studyPlace.Timestamp = append(studyPlace.Timestamp, elem.Timestamp)
		studyPlace.HistoricalDensity[0] = append(studyPlace.HistoricalDensity[0], elem.Value)
	}

	studyPlace.Levels, _ = db.GetSubLocationsById(locationID)
	result, _ := json.Marshal(studyPlace)
	
	rw.Header().Set("Content-Type", "application/json")
		
	rw.Write(result)
}

func eventsHandler(rw http.ResponseWriter, req *http.Request) {
	db, _ := storage.NewDB("db.sqlite3")
	events, _:= db.ListAllEvents()
	for ind, _ := range events {
		events[ind].Images = []string{events[ind].Image}
	}
	result, _ := json.Marshal(events)
	
	rw.Header().Set("Content-Type", "application/json")
		
	rw.Write(result)
}

func eventHandler(rw http.ResponseWriter, req *http.Request) {
	db, _ := storage.NewDB("db.sqlite3")
	param := req.URL.Query()["id"][0];
	
	eventID, _ := strconv.Atoi(param)
	event, _ := db.GetEventById(eventID)
	
	event.Images = []string{event.Image}
	result, _ := json.Marshal(event)
	
	rw.Header().Set("Content-Type", "application/json")
		
	rw.Write(result)
}

func testHandler(rw http.ResponseWriter, req *http.Request) {
	// db, _ := storage.NewDB("db.sqlite3")
	// value, _ := db.GetLastEventId()
	// fmt.Println(value)
	// ev := models.Event{value + 1, "testName", "testLocation", "testType", 150, "sdkjflj", make([]string, 0)}
	// db.AddEvent(&ev)
	//Upload("http://localhost:3000/upload", "C:/Users/Dian Bakti/Documents/repos/cms-server/static/img/events/CACGRAD.png")
}

func uploadHandler(rw http.ResponseWriter, req *http.Request) {
	db, _ := storage.NewDB("db.sqlite3")
	file, handler, err := req.FormFile("image") // img is the key of the form-data
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	
	name := filepath.Base(handler.Filename)
	fmt.Println(name)
	
	newPath := "./static/img/events/" + name
	dbPath := "img/events/" + name
	fmt.Println(newPath)
	f, err := os.OpenFile(newPath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)	
	req.ParseForm()

	value, _ := db.GetLastEventId()
	//crowdValue, _ := strconv.Atoi(req.Form["crowd"][0])
	ev := models.Event{value + 1, req.Form["name"][0], req.Form["location"][0], req.Form["eventType"][0], random(120, 200), dbPath, make([]string, 0)}
	db.AddEvent(&ev)
	newAddr := strings.Split(req.RemoteAddr, ":")[0]
	newAddr = "http://" + newAddr + ":8100"
	http.Redirect(rw, req, newAddr + "/#/tab/event", 301)
}

func main() {
	var dir string
	storage.NewDB("db.sqlite3")
	flag.StringVar(&dir, "dir", "./static/", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()
	r := mux.NewRouter()
    r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))
	r.HandleFunc("/study-place", studyPlaceHandler)
	r.HandleFunc("/location", locationHandler)

	r.HandleFunc("/events", eventsHandler)
	r.HandleFunc("/event", eventHandler)
	
	r.HandleFunc("/upload", uploadHandler)
	r.HandleFunc("/test", testHandler)

    // Apply the CORS middleware to our top-level router, with the defaults.
    err := http.ListenAndServe(":3000", cors.Default().Handler(r))
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
	}
	
}