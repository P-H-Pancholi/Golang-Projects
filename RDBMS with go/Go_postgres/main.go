package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

type Actor struct {
	actor_id   int64
	first_name string
	last_name  string
}

func main() {

	// link to golang db drivers : https://github.com/golang/go/wiki/SQLDrivers

	connString := "user=postgres password=7411 host=localhost port=5432 dbname=sakila sslmode=disable"

	var err error

	db, err = sql.Open("postgres", connString)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	pingErr := db.Ping()

	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected !")

	actorId, err := addActor("Jennifer", "Lawrence")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("actor added : %v\n", actorId)

	actor, err := GetActor(actorId)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Actor : %v\n", actor)

	RowsAffected, err := updateActor("Jennifer", actorId)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Total actors Update : %v\n", RowsAffected)

	RowsAffected, err = deleteActor(actorId)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Total actors deleted : %v\n", RowsAffected)

	actor, err = GetActor(actorId)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Actor : %v\n", actor)

}

func addActor(firstName, lastName string) (int64, error) {

	var id int64
	err := db.QueryRow("INSERT INTO actor (first_name , last_name) VALUES ( $1 , $2) RETURNING actor_id", firstName, lastName).Scan(&id)
	// QueryRow returns atmost one row, in postgres we have returning & lastInsertId() is not support so we use QueryRow instead of exec()

	if err != nil {
		return 0, fmt.Errorf("addActor: %v", err)
	}
	return id, nil
}

func updateActor(firstName string, actor_id int64) (int64, error) {

	var id int64
	result, err := db.Exec("UPDATE actor SET first_name = $1 WHERE actor_id = $2 ", firstName, actor_id)
	//Exec does not returns row
	if err != nil {
		return 0, fmt.Errorf("UpdateActor: %v", err)
	}

	id, err = result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("UpdateActor: %v", err)
	}
	return id, nil
}

func GetActor(id int64) ([]Actor, error) {
	var actors []Actor
	result, err := db.Query("SELECT actor_id, first_name, last_name FROM actor WHERE actor_id = $1", id)
	//Executes query & returns resultSet

	if err != nil {
		return nil, fmt.Errorf("GetActor : %v", err)
	}
	defer result.Close()

	for result.Next() {
		var act Actor
		if err := result.Scan(&act.actor_id, &act.first_name, &act.last_name); err != nil { // use scan on result sets
			return nil, fmt.Errorf("GetActor : %v", err)
		}
		actors = append(actors, act)

		if err := result.Err(); err != nil {
			return nil, fmt.Errorf("GetActor : %v", err)
		}
	}

	return actors, nil

}

func deleteActor(id int64) (int64, error) {
	result, err := db.Exec("DELETE FROM actor WHERE actor_id = $1", id)

	if err != nil {
		return 0, fmt.Errorf("deleteActor : %v", err)
	}

	rowAffect, err := result.RowsAffected()

	if err != nil {
		return 0, fmt.Errorf("deleteActor : %v", err)
	}

	return rowAffect, nil

}
