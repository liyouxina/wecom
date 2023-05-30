package dbmysql

import (
	"gorm.io/gorm"
	"log"
)

type Token struct {
	gorm.Model
	CorpId string `json:"corp_id" gorm:"unique"`
	Secret string `json:"secret"`
	Token  string `json:"token"`
}

func (*Token) TableName() string {
	return "token"
}

func CreateToken(token *Token) int64 {
	result := db.Create(token)
	return result.RowsAffected
}

func GetTokenList(token Token) (result []Token, err error) {
	defer func() {
		err := recover()
		log.Printf("GetTokenList error %v\n", err)
	}()
	if token.ID != 0 {
		db.Where("id = ?", token.ID)
	}
	if token.CorpId != "" {
		db.Where("corp_id like ?", likeWrap(token.CorpId))
	}
	result = []Token{}
	db.Find(&result)
	return result, nil
}
