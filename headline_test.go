package main

import "fmt"

func ExampleParseFile() {
  content := parseFile("content_test.yml")

  fmt.Println(content["subject"])
  fmt.Println(content["object"])
  fmt.Println(content["action"])
  fmt.Println(content["adverbial"])

  // Output:
  // [Donald Trump]
  // [Migrants]
  // [will ban %{object}]
  // [tomorrow]
}

func ExampleGenerateHeadline() {
  content := parseFile("content_test.yml")
  headline := generateHeadline(content)

  fmt.Println(headline)

  // Output: Donald Trump will ban Migrants tomorrow
}
