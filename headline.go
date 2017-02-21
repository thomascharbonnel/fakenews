package main

import (
  "log"
  "gopkg.in/yaml.v2"
  "io/ioutil"
  "math/rand"
  "time"
  "regexp"
)

type Content map[string][]string

func parseFile(location string) Content {
  var content Content

  // Get all the content
  raw, err := ioutil.ReadFile(location)
  if err != nil {
    panic(err)
  }

  // Parse content
  yaml.Unmarshal(raw, &content)

  return content
}

func generateHeadline(content Content) string {
  object_reg, err := regexp.Compile("%{object}")
	if err != nil {
	  log.Fatal(err)
	}

  s := rand.NewSource(time.Now().UnixNano())
  r := rand.New(s)

  subject := content["subject"][r.Intn(len(content["subject"]))]
  object := content["object"][r.Intn(len(content["object"]))]
  action := content["action"][r.Intn(len(content["action"]))]
  adverbial := content["adverbial"][r.Intn(len(content["adverbial"]))]

  full_action := object_reg.ReplaceAllString(action, object)

  return subject + " " + full_action + " " + adverbial
}
