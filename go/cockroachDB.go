package swagger

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func getConnection() *sql.DB {

	roachConnection, err := sql.Open("postgres", "postgresql://root@localhost:26257/ptruoradb?sslmode=disable")

	if err != nil {
		panic("Could not establish connection to CockroachDB")
	}

	return roachConnection
}

/*func createRecord(post Post) error {
	qryString := fmt.Sprintf(
		"INSERT INTO posts (name, category, author, created_at, updated_at) VALUES ('%s', '%s', '%s', NOW(), NOW())",
		post.Name, post.Category, post.Author)

	_, err := getConnection().Exec(qryString)
	return err
}

func readRecords() []Post {
	var result []Post
	rows, err := roachConnection.Query("select * FROM posts;")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		var post Post
		if err := rows.Scan(&post.ID, &post.Name, &post.Author, &post.Category, &post.CreatedAt, &post.UpdatedAt); err != nil {
			log.Fatal(err)
		}
		result = append(result, post)
	}

	return result
}

func updateRecord(fieldName string, newVal string, condFieldName string, condFieldValue int64) error {
	qryString := fmt.Sprintf("UPDATE posts set %s='%s' WHERE %s=%d", fieldName, newVal, condFieldName, condFieldValue)

	fmt.Printf("Query: %v\n", qryString)
	_, err := getConnection().Exec(qryString)

	return err
}

func deleteRecord(id int64) error {
	qryString := fmt.Sprintf("DELETE from posts WHERE id = %d", id)
	fmt.Printf("Query: %v\n", qryString)

	_, err := getConnection().Exec(qryString)
	return err
}
*/
