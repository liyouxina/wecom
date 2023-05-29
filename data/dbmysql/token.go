package dbmysql

import "gorm.io/gorm"

type Token struct {
	gorm.Model
	CorpId string `json:"corp_id"`
	Secret string `json:"secret"`
	Token  string `json:"token"`
}

func CreateToken(token *Token) error {
	result := db.Create(token)
	return result.Error
}
