package main

import "fmt"

func ExampleGenerateHTML() {
  fmt.Println(generateHTML("lol", "pony"))

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
  //   <style>
  //     #frame {
  //       height: 50%;
  //       width: 50%;
  //       margin: 0 auto;
  //     }
  //   </style>
  // </head>
  // <body>
  //   <div id="frame">lol</div>
  // </body>
  // </html>
}
