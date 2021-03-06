package main

import (
  "net/http"
  "fmt"
  "log"
  "io/ioutil"
  "path/filepath"
  "mime"
  "math/rand"
  "flag"
  "strconv"
)

func ServePicture(w http.ResponseWriter, r *http.Request, image_path string) {
  w.Header().Set("Content-Type", mime.TypeByExtension(filepath.Ext(image_path)))
  http.ServeFile(w, r, image_path)
}

func main() {
  files, err := ioutil.ReadDir("./pictures")
  check(err)

  port := flag.Int("p", 8080, "a port number")
  flag.Parse()

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    headline := generateHeadline(parseFile("./content.yml"))

    fmt.Fprintf(w, generateHTML(headline))
  })

  // Background photo of Trump
  http.HandleFunc("/photo", func(w http.ResponseWriter, r *http.Request) {
    image_idx := rand.Int() % len(files)
    image_path := "pictures/" + files[image_idx].Name()

    ServePicture(w, r, image_path)
  })

  http.HandleFunc("/logo", func(w http.ResponseWriter, r *http.Request) {
    ServePicture(w, r, "trump_logo.jpg")
  })

  http.HandleFunc("/style.css", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/css")
    http.ServeFile(w, r, "style.css")
  })

  log.Printf("Server will start on port %d.", *port)
  log.Fatal(http.ListenAndServe(":" + strconv.Itoa(*port), nil))
}
