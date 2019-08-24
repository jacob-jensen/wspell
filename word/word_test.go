package word

import (
	"os"
	"testing"
)

func TestAddWord(t *testing.T) {
	data := []Word{
		Word{Word: "hej"},
		Word{Word: "med"},
		Word{Word: "dig"},
		Word{Word: "du"},
		Word{Word: "er"},
		Word{Word: "sød"},
	}
	dbname := "words_test.db"
	if err := os.Remove(dbname); err != nil {
		if !os.IsNotExist(err) {
			t.Errorf("failed to remove file: %s", dbname)
			return
		}
	}
	if err := Init(dbname); err != nil {
		t.Errorf("failed to open database: %s", err.Error())
		return
	}
	defer Exit()
	for _, w := range data {
		if err := Add(w); err != nil {
			t.Errorf("failed to add word: %s", err.Error())
			return
		}
	}
	words := Find()
	if len(data) != len(words) {
		t.Errorf("wrong number of words found %d vs %d", len(data), len(words))
		return
	}
	for idx, w := range words {
		if w.Word != data[idx].Word {
			t.Errorf("words do not match in idx: %d (%s vs %s)", idx, w.Word, data[idx].Word)
			return
		}
	}
}
