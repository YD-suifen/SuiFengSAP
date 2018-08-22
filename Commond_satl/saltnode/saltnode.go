package saltnode

import (
	"SuiFengSAP/Commond_satl/http"
	"SuiFengSAP/Commond_satl/over"
	"SuiFengSAP/Commond_satl/token"
	"encoding/json"
	"fmt"
)

func RequestKeylist() []string {
	var request over.Request
	request.Client = "wheel"
	request.Tgt = "*"
	request.Fun = "key.list_all"

	requesjson, err := json.Marshal(request)

	if err != nil {
		fmt.Println(err)
	}

	tocken, err := token.Token()

	err, respone := http.Httprequest("POST", http.Url, requesjson, tocken)
	if err != nil {
		fmt.Println(err)
	}

	var Respon over.KeylistRespon

	_ = json.Unmarshal(respone, &Respon)

	var keylist []string

	for _, v := range Respon.Return {
		keylist = v.Data.Return.Minions
	}
	return keylist

}
