package dblite

import (
  "database/sql"
  "os"
  //"fallfin/models"
  "FallTestMSGo/fallfin/models"
  //"log"
)

/* Crear BD */
func CreateBD() *sql.DB{
  //var db *sql.DB
	//var err error
  os.Remove("./foo.db")
	db, err := sql.Open("sqlite3", "./foo.db")
  if err != nil {
		//log.Fatal(err)
	}
  return db
}

/* Crear Table */
func CreateTable(db *sql.DB){
  db.Exec("create table if not exists BeerItemTable (id integer not null primary key,name text, brewery text)")
}

/* Agregar BeerItem*/
func AddBeerItem(db *sql.DB, name string, brewery string) {
  tx, _ := db.Begin()
  stmt, _ := tx.Prepare("insert into BeerItemTable (name,brewery) values (?,?)")
  _, err := stmt.Exec(name, brewery)
  if err != nil {
		//log.Fatal(err)
	}
  tx.Commit()
}

/* obtener una cerveza  por ID */
func GetBeerItem(db *sql.DB, id2 int) models.BeerItem{
  var stmt *sql.Stmt
	var err error

  /*
  rows, err := db.Query("select * from BeerItemTable")
  //checkError(err)
  if err != nil {
		log.Fatal(err)
	}

  for rows.Next() {
    var tempBeerItem models.BeerItem
    err = rows.Scan(&tempBeerItem.Id, &tempBeerItem.Name,  &tempBeerItem.Brewery)
    if err != nil {
  		log.Fatal(err)
  	}
    if tempBeerItem.Id== id2 {
      return tempBeerItem
    }
  }

  return models.BeerItem{}
  */

  stmt, err = db.Prepare("select * from BeerItemTable where id = ?")
	if err != nil {
		//log.Fatal(err)
	}
	defer stmt.Close()

  var tempBeerItem models.BeerItem
  err = stmt.QueryRow(id2).Scan(&tempBeerItem.Id, &tempBeerItem.Name,  &tempBeerItem.Brewery)
	if err != nil {
		//log.Fatal(err)
	}
	//log.Println("Name:" + tempBeerItem.Name)

  return tempBeerItem

}
