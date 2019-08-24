package word

import (
	"fmt"

	"github.com/jinzhu/gorm"
	//We are using sqlite as database.
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//Word contains a single word
type Word struct {
	gorm.Model  `json:"-"`
	Word        string `json:"word"`
	Category    int    `json:"category,omitempty"`
	PicturePath string `json:"picturePath,omitempty"`
}

//Init initializes and opens the database.
func Init(dbname string) error {
	var err error
	db, err = gorm.Open("sqlite3", dbname)
	if err != nil {
		return fmt.Errorf("failed to open database words.db (%s)", err.Error())
	}
	db.AutoMigrate(&Word{})
	return nil
}

//Exit closes the database.
func Exit() {
	db.Close()
}

//Add adds a single word to the word database.
func Add(word Word) error {
	if err := db.Create(&word).Error; err != nil {
		return fmt.Errorf("failed to create word (%s)", err.Error())
	}
	return nil
}

//Find finds all words in database.
func Find() []Word {
	var words []Word
	db.Find(&words)
	return words
}

var db *gorm.DB
