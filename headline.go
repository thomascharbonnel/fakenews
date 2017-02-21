package main

import (
  "gopkg.in/yaml.v2"
  "io/ioutil"
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
