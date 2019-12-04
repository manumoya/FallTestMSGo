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
  "FallTestMSGo/fallfin/api"

  "fmt"

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

  /* Persistencia  */
  /* crear BD */
  var db *sql.DB
  db = dblite.CreateBD()

  /* Crear tabla*/
  dblite.CreateTable(db)

  /* Grabar una cervza*/
  //dblite.AddBeerItem(db, "Pilsen",  "Cristal") // added data to database
  //var beerItem = dblite.GetBeerItem(db, 1) // printing the user

  /* Fin Persistencia */

  /* home */
  e.GET("/", func(c echo.Context) error {
    return c.String(http.StatusOK, "Hello , World!\n")
  })

  /* Route => api*/

  e.POST("/beers", api.AddBeers)
  e.GET("/beers", api.SearchBeerByIdGET)

  // Server
  fmt.Printf("Running... 8080")
  e.Run(standard.New(":8080"))

}
