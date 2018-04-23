package postgres

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

const (
	user   = "postgres"
	dbname = "postgres"
)

func setupDB(user, dbname string) (db *sql.DB, err error) {
	dbinfo := fmt.Sprintf("user=%s dbname=%s sslmode=disable", user, dbname)
	db, err = sql.Open("postgres", dbinfo)
	return db, err
}
func TestDB(t *testing.T) {
	db, err := setupDB(user, dbname)
	assert.Nil(t, err)

	err = db.Ping()
	assert.Nil(t, err)
}

func TestEndorse1st(t *testing.T) {
	db, err := setupDB(user, dbname)

	res, err := getEndorse1st(db)
	assert.Nil(t, err)
	fmt.Println(len(res))
}

func TestEndorse2nd(t *testing.T) {
	db, err := setupDB(user, dbname)

	res, err := getEndorse2nd(db)
	assert.Nil(t, err)
	fmt.Println(len(res))
}

func TestEndorse3rd(t *testing.T) {
	db, err := setupDB(user, dbname)

	res, err := getEndorse3rd(db)
	assert.Nil(t, err)
	fmt.Println(len(res))
}

func TestEndorse4th(t *testing.T) {
	db, err := setupDB(user, dbname)

	res, err := getEndorse4th(db)
	assert.Nil(t, err)
	fmt.Println(len(res))
}

func TestEndorse5th(t *testing.T) {
	db, err := setupDB(user, dbname)

	res, err := getEndorse5th(db)
	assert.Nil(t, err)
	fmt.Println(len(res))
}

// func TestCreateDataTables(t *testing.T) {
// 	db, err := setupDB(user, dbname)
// 	err = db.Ping()
// 	assert.Nil(t, err)

// 	_, err = db.Query(`
// 		CREATE TABLE person (
// 			id INTEGER PRIMARY KEY,
// 			name VARCHAR(255),
// 			job VARCHAR(255),
// 			birthday VARCHAR(255)
// 		);
// 	`)
// 	assert.Nil(t, err)

// 	_, err = db.Query(`
// 		CREATE TABLE relation (
// 			id1 INTEGER REFERENCES person(id),
// 			id2 INTEGER REFERENCES person(id)
// 		);
// 	`)
// 	assert.Nil(t, err)
// }

// func TestInsertData(t *testing.T) {
// 	db, err := setupDB(user, dbname)
// 	err = db.Ping()
// 	assert.Nil(t, err)

// 	_, err = db.Query(`
// 		COPY person FROM PROGRAM 'tail -n +2 /import/social_network_nodes.csv'
// 		WITH (FORMAT csv);
// 	`)
// 	assert.Nil(t, err)

// 	_, err = db.Query(`
// 		COPY relation FROM PROGRAM 'tail -n +2 /import/social_network_edges.csv'
// 		WITH (FORMAT csv);
// 	`)
// 	assert.Nil(t, err)
// }
