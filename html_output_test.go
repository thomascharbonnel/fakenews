package main

import "fmt"

func ExampleGenerateHTML() {
  fmt.Println(generateHTML("lol"))

  // Output:
  // <!doctype html>
  // <html lang="en">
  // <head>
  //   <meta charset="utf-8">
  // 
  //   <title>lol — Fake News Generator</title>
  //   <meta name="description" content="Fake News Generator">
  //   <meta name="author" content="Thomas Charbonnel">
  //
  //   <link href="style.css" type="text/css" rel="stylesheet">
  // </head>
  // <body>
  //   <div id="frame">
  //     <img id="background" src="photo">
  //     <p id="headline">lol</p>
  //     <img id="logo" src="logo">
  //   </div>
  // </body>
  // </html>
}
