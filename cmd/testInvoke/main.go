package main

import "github.com/liyouxina/wecom/data/dbmysql"

func main() {
	dbmysql.CreateToken(&dbmysql.Token{
		CorpId: "sad",
	})
}

func getToken() {
	//resp := wecom.GetToken(wecom.CorpId, wecom.Secret)
	//bytes, _ := json.Marshal(resp)
	//fmt.Println(string(bytes))
}
