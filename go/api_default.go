/*
 * ptruora
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0a0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var MyServerName = "http://127.0.0.1:5501"

func RestDomainResourceCLASSINSTANCEGetDomain(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", MyServerName)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	params := mux.Vars(r)
	nameDomain := params["name"]

	objetoApi := consumeSslLabsToDomainParameters(consumeSslLabs(nameDomain))
	objetoDB := getDomain(nameDomain)

	if objetoDB.Name != "" && objetoApi.Name != "" {
		saveDomain(objetoApi)
	}

	if objetoApi.Name != "" {
		_ = saveTrace(nameDomain)

		json.NewEncoder(w).Encode(objetoApi)
		w.WriteHeader(http.StatusOK)

	} else {
		var message = Msg{
			"The domain '" + nameDomain + "' does not exist", 404,
		}
		json.NewEncoder(w).Encode(message)
		w.WriteHeader(http.StatusNotFound)
	}

}

func RestDomainResourceCLASSINSTANCEGetTracert(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", MyServerName)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var listaResp []string

	db := getConnection()

	rows, err := db.Query("SELECT * FROM ptruoradb.tbl_traces ORDER BY datetime DESC LIMIT 20;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id_domain string
		var datetime string

		if err := rows.Scan(&id_domain, &datetime); err != nil {
			log.Fatal(err)
		}

		listaResp = append(listaResp, id_domain)

	}

	objres := TracertParameters{
		Items: listaResp,
	}

	json.NewEncoder(w).Encode(objres)
	w.WriteHeader(http.StatusOK)

}
