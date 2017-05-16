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

  <link href="style.css" type="text/css" rel="stylesheet">
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
