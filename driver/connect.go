package driver

import (
	"database/sql"
	"fmt" // package used to read the .env file
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // mysql golang driver
)

const (
	user     = "root"
	password = "mymysql"
	dbname   = "article"
	host     = ""
	// port     = 3306
)

// Generate Database Function
func GenerateDatabase() {

	db, err := sql.Open("mysql", "root:mymysql@tcp(:"+os.Getenv("PORT")+")/")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + dbname)
	if err != nil {
		panic(err)
	}

	mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s:"+os.Getenv("PORT")+")/%s",
		user, password, host, dbname)

	db, err = sql.Open("mysql", mysqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS posts (
		id INT NOT NULL AUTO_INCREMENT NOT NULL PRIMARY KEY,
		title VARCHAR(200) NOT NULL,
		content text NOT NULL,
		category VARCHAR(100) NOT NULL,
		created_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_date TIMESTAMP NULL DEFAULT NULL,
		status ENUM('Publish','Draft','Thrash')
		)`)
	if err != nil {
		panic(err)
	}
}

//ConnectDB function
func ConnectDB() (*sql.DB, error) {
	mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s:"+os.Getenv("PORT")+")/%s",
		user, password, host, dbname)

	db, err := sql.Open("mysql", mysqlInfo)
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}

	fmt.Println("Successfully connected!")
	return db, nil
}
