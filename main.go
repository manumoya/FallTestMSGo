package main

import (
  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
  "github.com/labstack/echo/engine/standard"
  "net/http"

  "database/sql"
  _ "github.com/mattn/go-sqlite3"

  "log"
	"os"

  // fallfin/dblite
  "FallTestMSGo/fallfin/dblite"

)



func main(){

  f, err := os.OpenFile("log.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
  if err != nil {
    log.Fatalf("error opening file: %v", err)
  }
  defer f.Close()

  log.SetOutput(f)
  log.Println("This is a test log entry")

  /* SETUP */
  e := echo.New()
  e.Use(middleware.CORS())

  /* ROUTES */
  /*
  e.POST("/api/articles", postArticle)
  e.POST("/api/comments", postComment)

  e.PUT("/api/comments/:comment_id", updateComment)

  e.GET("/api/articles", getArticles)
  e.GET("/api/articles/:article_id/comments", getCommentsByArticle)
  e.GET("/api/comments/:comment_id", getCommentById)

  e.DELETE("/api/comments/:comment_id", deleteComment)
  e.DELETE("/api/articles/:article_id", deleteArticle)
  */


  /* Persistencia  */

  // crear BD

  var db *sql.DB
  db = dblite.CreateBD()

  /*
  os.Remove("./foo.db")
	db, err := sql.Open("sqlite3", "./foo.db")
  if err != nil {
		log.Fatal(err)
	}
  */
	//defer db.Close()
  //db.Exec("create table if not exists testTable (id integer,username text, surname text,age Integer,university text)")
  dblite.CreateTable(db)

  dblite.AddBeerItem(db, "Pilsen",  "Cristal") // added data to database
  var beerItem = dblite.GetBeerItem(db, 1) // printing the user

 /* Fin Persistencia */

  e.GET("/", func(c echo.Context) error {
    return c.String(http.StatusOK, "Hello  " + beerItem.Name+ " , World!\n")
  })

  // Server
  e.Run(standard.New(":8080"))

}
