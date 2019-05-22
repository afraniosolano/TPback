package swagger

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func consumeSslLabs(idDomain string) SslLabs {

	var info SslLabs

	response, err := http.Get("https://api.ssllabs.com/api/v3/analyze?host=" + idDomain)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {

		defer response.Body.Close()

		data, _ := ioutil.ReadAll(response.Body)

		//data2 := string(data)
		//fmt.Println(data2)

		err2 := json.Unmarshal(data, &info)
		if err2 != nil {
			fmt.Printf("Error %s\n", err2)
		}
	}

	return info

}
