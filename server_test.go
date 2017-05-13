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
  //     body {
  //       background-color: black;
  //     }
  // 
  // 
  //     #frame {
  //       width: 100vw;
  //       height: 56.25vw;
  //       max-height: 100vh;
  //       max-width: 177.78vh;
  //       margin: auto;
  //       position: absolute;
  //       top:0;bottom:0;
  //       left:0;right:0;
  //     }
  // 
  //     img#background {
  //       position: absolute;
  //       width: 100%;
  //       height: 100%;
  //     }
  // 
  //     #headline {
  //       position: absolute;
  //       background-color: white;
  //       z-index: 100;
  //       bottom: 2%;
  //       width: 80%;
  //       padding: 1%;
  //       margin: 1%;
  //       font-size: 24pt;
  //       font-weight: bold;
  //     }
  // 
  //     img#logo {
  //       position: absolute;
  //       z-index: 100;
  //       background-color: white;
  //       bottom: 2%;
  //       right: 0%;
  //       width: 13%;
  //       padding: 1%;
  //       margin: 1%;
  //       height: 7%;
  //     }
  //   </style>
  // </head>
  // <body>
  //   <div id="frame">
  //     <img id="background" src="pony">
  //     <p id="headline">lol</p>
  //     <img id="logo" src="./trump_logo.jpg">
  //   </div>
  // </body>
  // </html>
}
