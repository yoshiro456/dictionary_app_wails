package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	sqlConn := sqlite.Open("assets/database.db")
	db, err := gorm.Open(sqlConn, &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db
}

type EMDictionary struct {
	Id      uint   `json:"id"`
	Word    string `json:"word"`
	Content string `json:"content"`
}

func (d *EMDictionary) TableName() string {
	return "eng_mym"
}

func (d *EMDictionary) FindById(id uint) (dict *EMDictionary) {
	DB.First(&dict)
	return
}

func (d *EMDictionary) FindManyByWord(word string) (dicts *[]EMDictionary) {
	DB.Where("word LIKE ?", word+"%").Limit(50).Find(&dicts)
	return
}

type MEDictionary struct {
	Id      uint   `json:"id"`
	Word    string `json:"word"`
	Content string `json:"content"`
}

func (d *MEDictionary) TableName() string {
	return "mym_eng"
}

func (d *MEDictionary) FindById(id uint) (dict *MEDictionary) {
	DB.First(&dict)
	return
}

func (d *MEDictionary) FindManyByWord(word string) (dicts *[]MEDictionary) {
	DB.Find(dicts, "word = ?", word)
	return
}
