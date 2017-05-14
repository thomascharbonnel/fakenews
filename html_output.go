package main

import (
  "html/template"
  "bytes"
)

func generateHTML(headline string) string {
  const tpl = `
<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">

  <title>{{.Headline}} — Fake News Generator</title>
  <meta name="description" content="Fake News Generator">
  <meta name="author" content="Thomas Charbonnel">

  <style>
    body {
      background-color: black;
    }

    /* http://stackoverflow.com/a/20593342/3745914 */
    #frame {
      width: 100vw; 
      height: 56.25vw; /* height:width ratio = 9/16 = .5625  */
      max-height: 100vh;
      max-width: 177.78vh; /* 16/9 = 1.778 */
      margin: auto;
      position: absolute;
      top:0;bottom:0; /* vertical center */
      left:0;right:0; /* horizontal center */
    }

    img#background {
      position: absolute;
      width: 100%;
      height: 100%;
    }

    #headline {
      position: absolute;
      background-color: white;
      z-index: 100;
      bottom: 2%;
      width: 80%;
      padding: 1%;
      margin: 1%;
      font-size: 24pt;
      font-weight: bold;
    }

    img#logo {
      position: absolute;
      z-index: 100;
      background-color: white;
      bottom: 2%;
      right: 0%;
      width: 13%;
      padding: 1%;
      margin: 1%;
      height: 7%;
    }
  </style>
</head>
<body>
  <div id="frame">
    <img id="background" src="photo">
    <p id="headline">{{.Headline}}</p>
    <img id="logo" src="logo">
  </div>
</body>
</html>
`

  output := new(bytes.Buffer)

  t, err := template.New("index").Parse(tpl)
  check(err)

  data := struct {
    Headline string
  }{
    Headline: headline,
  }

  err = t.Execute(output, data)

  return output.String()
}
