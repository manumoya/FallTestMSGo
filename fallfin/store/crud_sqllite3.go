package store

import (
  "database/sql"
  "os"
  //"fallfin/models"
  "FallTestMSGo/fallfin/models"
  //"log"
)

/* Open BD */
func OpenBD() *sql.DB{
  //var db *sql.DB
	//var err error
	db, err := sql.Open("sqlite3", "./foo.db")
  if err != nil {
		//log.Fatal(err)
	}
  return db
}

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
  //db.Exec("create table if not exists BeerItemTable (id integer not null primary key,name text, brewery text)")
  db.Exec("create table if not exists BeerItemTable (id integer ,name text, brewery text,country text, price real, currency text)")

}

/* Agregar BeerItem*/
func AddBeerItem(db *sql.DB, id int, name string, brewery string, country string, price float32, currency string) {
  tx, _ := db.Begin()
  stmt, _ := tx.Prepare("insert into BeerItemTable (id,name,brewery,country,price,currency) values (?,?,?,?,?,?)")
  _, err := stmt.Exec(id, name, brewery, country, price, currency)
  if err != nil {
		//log.Fatal(err)
	}
  tx.Commit()
}

/* obtener una cerveza  por ID */
func GetBeerItem(db *sql.DB, id2 int) models.BeerItem{
  var stmt *sql.Stmt
	var err error

  stmt, err = db.Prepare("select * from BeerItemTable where id = ?")
	if err != nil {
		//log.Fatal(err)
	}
	defer stmt.Close()

  var tempBeerItem models.BeerItem
  err = stmt.QueryRow(id2).Scan(&tempBeerItem.Id, &tempBeerItem.Name,  &tempBeerItem.Brewery,
                                &tempBeerItem.Country, &tempBeerItem.Price, &tempBeerItem.Country)

  if err != nil {
		//log.Fatal(err)
	}
	return tempBeerItem
}

/*
rows, err := db.Query("select * from BeerItemTable")
//checkError(err)
if err != nil {
  log.Fatal(err)
}

*/

func SearchAllBeer(db *sql.DB) models.BeerItemList {
  var beers models.BeerItemList

  rows, err := db.Query("select * from BeerItemTable")
	if err != nil {
		//log.Fatal(err)
	}
	defer rows.Close()
  for rows.Next() {
    var tempBeerItem models.BeerItem
    err = rows.Scan(&tempBeerItem.Id, &tempBeerItem.Name,  &tempBeerItem.Brewery,
                    &tempBeerItem.Country, &tempBeerItem.Price, &tempBeerItem.Currency)

    beers.Beers = append(beers.Beers, tempBeerItem)
  }

  return beers
}
