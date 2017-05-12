package main

import (
  "html/template"
  "bytes"
)

func generateHTML(headline string, image_path string) string {
  const tpl = `
<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">

  <title>{{.Title}} — Fake News Generator</title>
  <meta name="description" content="Fake News Generator">
  <meta name="author" content="Thomas Charbonnel">

  <style>
    #frame {
      height: 50%;
      width: 50%;
      margin: 0 auto;
    }
  </style>
</head>
<body>
  <div id="frame">{{.Headline}}</div>
</body>
</html>
`

  output := new(bytes.Buffer)

	check := func(err error) {
    if err != nil {
      panic(err)
    }
  }

  t, err := template.New("index").Parse(tpl)
  check(err)

  data := struct {
    Title string
    Headline string
  }{
    Title: headline,
    Headline: headline,
  }

  err = t.Execute(output, data)

  return output.String()
}
