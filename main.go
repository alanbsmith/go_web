package main

import (
  "fmt"
  "github.com/codegangsta/martini"
  "github.com/martini-contrib/render"
  "database/sql"
  "net/http"
  _ "github.com/lib/pq"
)

func SetupDB() *sql.DB {
  db, err := sql.Open("postgres", "dbname=ideabox sslmode=disable")
  PanicIf(err)
  return db
}

func PanicIf(err error) {
  if err != nil {
    panic(err)
  }
}

func main() {
  m:= martini.Classic()
  m.Map(SetupDB())

  m.Use(render.Renderer(render.Options{
    Layout: "layout",
  }))

  m.Get("/", func(rw http.ResponseWriter, r *http.Request, db *sql.DB) {
    rows, err := db.Query("SELECT title, description FROM ideas")
    PanicIf(err)
    defer rows.Close()

    var title, description string
    for rows.Next() {
      err := rows.Scan(&title, &description)
      PanicIf(err)
      fmt.Fprintf(rw, "Title: %s\nDescription: %s\n\n",title, description)
    }
    // ren.HTML(200,"landing_page", nil)
  })

  m.Get("/ideas", func(ren render.Render, r *http.Request) {
  
  ren.HTML(200,"ideas/index", nil)
  })

  m.Run()
}