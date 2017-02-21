package main

import "fmt"

func ExampleParseFile() {
  yml := parseFile("content_test.yml")

  fmt.Println(yml["subject"])
  fmt.Println(yml["object"])
  fmt.Println(yml["action"])
  fmt.Println(yml["adverbial"])

  // Output:
  // [Donald Trump]
  // [Migrants]
  // [will ban %{object}]
  // [tomorrow]
}
