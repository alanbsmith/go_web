package main

import (
  "github.com/codegangsta/martini"
  "github.com/martini-contrib/render"
  "net/http"
  _ "github.com/lib/pq"
)

func main() {
  m:= martini.Classic()
  m.Use(render.Renderer(render.Options{
    Layout: "layout",
  }))

  m.Get("/", func(ren render.Render, r *http.Request) {
  
  ren.HTML(200,"landing_page", nil)
  })

  m.Get("/ideas", func(ren render.Render, r *http.Request) {
  
  ren.HTML(200,"ideas/index", nil)
  })

  m.Run()
}