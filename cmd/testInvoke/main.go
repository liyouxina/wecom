package main

import (
	"github.com/liyouxina/wecom/data/dbmysql"
	"log"
)

func main() {
	dbmysql.CreateToken(&dbmysql.Token{
		CorpId: "sad",
	})
	list, err := dbmysql.GetTokenList(dbmysql.Token{})
	if err != nil {
		log.Printf(err.Error())
	} else {
		log.Printf("%v", list)
	}
}

func getToken() {
	//resp := wecom.GetToken(wecom.CorpId, wecom.Secret)
	//bytes, _ := json.Marshal(resp)
	//fmt.Println(string(bytes))
}
