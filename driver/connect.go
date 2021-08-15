package driver

import (
	"database/sql"
	"fmt" // package used to read the .env file
	"log"

	_ "github.com/go-sql-driver/mysql" // mysql golang driver
)

const (
	user     = "root"
	password = "mymysql"
	dbname   = "article"
	host     = "127.0.0.1"
	port     = 3306
)

// Generate Database Function
func GenerateDatabase() {

	db, err := sql.Open("mysql", "root:mymysql@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + dbname)
	if err != nil {
		panic(err)
	}

	mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?multiStatements=true",
		user, password, host, port, dbname)

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

func InsertExampleData() {
	db, err := ConnectDB()

	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = db.Exec(`
		INSERT INTO posts VALUES (1, 'Updated Nice Article Title', 'Published quick, brown fox jumps over a lazy dog. DJs flock by when MTV ax quiz prog. Junk MTV quiz graced by fox whelps. Bawds jog, flick quartz, vex nymphs. Waltz, bad nymph, for quick jigs vex! Fox nymph', 'category3', '2021-08-13 16:14:25', '2021-08-15 19:15:56', 'Publish');
		INSERT INTO posts VALUES (3, 'Some Very Nice Article Title', 'The quick, brown fox jumps over a lazy dog. DJs flock by when MTV ax quiz prog. Junk MTV quiz graced by fox whelps. Bawds jog, flick quartz, vex nymphs. Waltz, bad nymph, for quick jigs vex! Fox nymph', 'category2', '2021-08-13 16:14:55', '2021-08-15 08:06:31', 'Draft');
		INSERT INTO posts VALUES (4, 'Thrashed Nice Article Title', 'New quick, brown fox jumps over a lazy dog. DJs flock by when MTV ax quiz prog. Junk MTV quiz graced by fox whelps. Bawds jog, flick quartz, vex nymphs. Waltz, bad nymph, for quick jigs vex! Fox nymph', 'category1', '2021-08-14 11:52:16', '2021-08-14 11:52:50', 'Thrash');
		INSERT INTO posts VALUES (5, 'First Kafka Article Title', 'One morning, when Gregor Samsa woke from troubled dreams, he found himself transformed in his bed into a horrible vermin. He lay on his armour-like back, and if he lifted his head a little he could see', 'Kafka', '2021-08-15 16:30:01', '2021-08-15 22:51:22', 'Publish');
		INSERT INTO posts VALUES (6, 'Second Kafka Article Title', 'One morning, when Gregor Samsa woke from troubled dreams, he found himself transformed in his bed into a horrible vermin. He lay on his armour-like back, and if he lifted his head a little he could see his brow', 'Kafka2', '2021-08-15 16:30:49', '2021-08-15 22:25:41', 'Publish');
		INSERT INTO posts VALUES (7, 'Third Kafka Article Title', 'Third morning, when Gregor Samsa woke from troubled dreams, he found himself transformed in his bed into a horrible vermin. He lay on his armour-like back, and if he lifted his head a little he could see his brow', 'Kafka3', '2021-08-15 22:48:09', NULL, 'Publish');
		INSERT INTO posts VALUES (8, 'Fourth Kafka Article Title', 'Fourth morning, when Gregor Samsa woke from troubled dreams, he found himself transformed in his bed into a horrible vermin. He lay on his armour-like back, and if he lifted his head a little he could see his brow', 'Kafka4', '2021-08-15 22:48:23', '2021-08-15 22:49:07', 'Publish');
		INSERT INTO posts VALUES (9, 'Fifth Kafka Article Title', 'Fifth morning, when Gregor Samsa woke from troubled dreams, he found himself transformed in his bed into a horrible vermin. He lay on his armour-like back, and if he lifted his head a little he could see his brow', 'Kafka5', '2021-08-15 22:49:37', NULL, 'Publish');
		INSERT INTO posts VALUES (10, 'Sixth Kafka Article Title', 'Sixth morning, when Gregor Samsa woke from troubled dreams, he found himself transformed in his bed into a horrible vermin. He lay on his armour-like back, and if he lifted his head a little he could see his brow', 'Kafka6', '2021-08-15 22:49:46', NULL, 'Publish');
		INSERT INTO posts VALUES (11, 'Seventh Kafka Article Title', 'Seventh morning, when Gregor Samsa woke from troubled dreams, he found himself transformed in his bed into a horrible vermin. He lay on his armour-like back, and if he lifted his head a little he could see his brow', 'Kafka7', '2021-08-15 22:49:57', NULL, 'Publish');
		INSERT INTO posts VALUES (12, 'Eighth Kafka Article Title', 'Eighth morning, when Gregor Samsa woke from troubled dreams, he found himself transformed in his bed into a horrible vermin. He lay on his armour-like back, and if he lifted his head a little he could see his brow', 'Kafka8', '2021-08-15 22:50:29', NULL, 'Publish');
		INSERT INTO posts VALUES (13, 'Ninth Kafka Article Title', 'Ninth morning, when Gregor Samsa woke from troubled dreams, he found himself transformed in his bed into a horrible vermin. He lay on his armour-like back, and if he lifted his head a little he could see his brow', 'Kafka9', '2021-08-15 22:50:42', NULL, 'Publish');
		INSERT INTO posts VALUES (14, 'Tenth Kafka Article Title', 'Tenth morning, when Gregor Samsa woke from troubled dreams, he found himself transformed in his bed into a horrible vermin. He lay on his armour-like back, and if he lifted his head a little he could see his brow', 'Kafka10', '2021-08-15 22:50:54', NULL, 'Publish');
		INSERT INTO posts VALUES (15, 'Eleventh Kafka Article Title', 'Eleventh morning, when Gregor Samsa woke from troubled dreams, he found himself transformed in his bed into a horrible vermin. He lay on his armour-like back, and if he lifted his head a little he could see his brow', 'Kafka11', '2021-08-15 22:51:06', '2021-08-15 22:51:29', 'Draft');
	`)
	if err != nil {
		panic(err)
	}
}

//ConnectDB function
func ConnectDB() (*sql.DB, error) {
	mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?multiStatements=true",
		user, password, host, port, dbname)

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
