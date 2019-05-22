package swagger

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func getConnection() *sql.DB {

	roachConnection, err := sql.Open("postgres", "postgresql://root@localhost:26257/ptruoradb?sslmode=disable")

	if err != nil {
		panic("Could not establish connection to CockroachDB")
	}

	return roachConnection
}

func getServers(idDomain string) []ServerParameters {

	var servidores []ServerParameters

	db2 := getConnection()

	rows2, err2 := db2.Query("select id_domain,address,ssl_grade,country,owner FROM tbl_server where id_domain = '" + idDomain + "';")

	if err2 == nil {

		for rows2.Next() {
			var idDomain string
			var address string
			var sslGrade string
			var country string
			var owner string

			if err3 := rows2.Scan(&idDomain, &address, &sslGrade, &country, &owner); err3 == nil {

				item := ServerParameters{address, sslGrade, country, owner}

				servidores = append(servidores, item)

			}

		}

	}

	defer rows2.Close()

	return servidores

}

func saveTrace(idDomain string) error {
	qryString := fmt.Sprintf(
		"INSERT INTO tbl_traces (id_domain, datetime) VALUES ('%s', NOW())",
		idDomain)

	_, err := getConnection().Exec(qryString)
	return err
}

func saveDomain(pD DomainParameters) error {

	qryString := fmt.Sprintf(
		"INSERT INTO tbl_domain (name,server_change,ssl_grade,previus_ssl_grade,logo,title,is_down) VALUES ('%s',%t,'%s','%s','%s',%t)",
		pD.Name, pD.ServerChange, pD.PreviusSslGrade, pD.Logo, pD.Title, pD.IsDown)

	for _, serve1 := range pD.Servers {
		saveServe(serve1, pD.Name)
	}

	_, err := getConnection().Exec(qryString)

	return err

}

func saveServe(serve1 ServerParameters, idDomain string) error {

	qryString := fmt.Sprintf(
		"INSERT INTO tbl_server (id_domain, address, ssl_grade, country, owner) VALUES ('%s','%s','%s','%s',%s)",
		idDomain, serve1.Address, serve1.SslGrade, serve1.Country, serve1.Owner)

	_, err := getConnection().Exec(qryString)

	return err

}

func getDomain(nameDomain string) DomainParameters {

	var dominio1 DomainParameters

	db := getConnection()

	rows, err := db.Query("select name,server_change,ssl_grade,previus_ssl_grade,logo,title,is_down FROM tbl_domain where name = '" + nameDomain + "';")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		var serverChange bool
		var sslGrade string
		var previusSsl_grade string
		var logo string
		var title string
		var isDown bool

		if err := rows.Scan(&name, &serverChange, &sslGrade, &previusSsl_grade, &logo, &title, &isDown); err != nil {
			log.Fatal(err)
		}

		servidores := getServers(name)

		dominio1 = DomainParameters{
			ServerChange:    serverChange,
			SslGrade:        sslGrade,
			PreviusSslGrade: previusSsl_grade,
			Logo:            logo,
			Title:           title,
			IsDown:          isDown,
			Name:            name,
			Servers:         servidores,
		}

	}

	return dominio1

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
