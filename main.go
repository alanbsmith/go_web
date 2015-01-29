package main

import (
  "github.com/codegangsta/martini"
  "github.com/martini-contrib/render"
  "database/sql"
  "net/http"
  _ "github.com/lib/pq"
)

type Idea struct {
  Title       string
  Description string
}

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

  m.Get("/", func(ren render.Render, r *http.Request) {
    ren.HTML(200,"landing_page", nil)
  })

  m.Get("/ideas", func(ren render.Render, r *http.Request, db *sql.DB) {
    rows, err := db.Query("SELECT title, description FROM ideas")
    PanicIf(err)
    defer rows.Close()

    ideas := []Idea{}
    for rows.Next() {
      i := Idea{}
      err := rows.Scan(&i.Title, &i.Description)
      PanicIf(err)
      ideas = append(ideas, i)
    }
    ren.HTML(200,"ideas/index", ideas)
  })

  m.Run()
}