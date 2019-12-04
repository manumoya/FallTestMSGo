package main

import (
  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
  "github.com/labstack/echo/engine/standard"
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
  "log"
	"os"
  "fmt"

  "FallTestMSGo/fallfin/store"
  "FallTestMSGo/fallfin/api"
)

// Iniciar Log
func iniLog() *os.File{
  f, err := os.OpenFile("log.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
  if err != nil {
    //log.Fatalf("error opening file: %v", err)
  }
  defer f.Close()
  return f
}

func main(){
  var fLog *os.File
  fLog = iniLog()
  log.SetOutput(fLog)
  log.Println("Prueba LOG")

  /* SETUP */
  e := echo.New()
  e.Use(middleware.CORS())

  /* crear BD */
  var db *sql.DB
  db = store.CreateBD()

  /* Crear tabla*/
  store.CreateTable(db)

  /* Route => api*/
  e.GET("/", api.Home)
  e.POST("/beers", api.AddBeers)
  e.GET("/beers/:beerID", api.SearchBeerByIdGET)
  e.GET("/beers", api.SearchBeers)

  // Server
  fmt.Printf("Running... 8080")
  e.Run(standard.New(":8080"))

}
