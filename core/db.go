package core

import (
	"encoding/json"
	"os"
)

const dbPath = "/var/lib/bitfer/db.json"

func loadDB() map[string]string {
	data, err := os.ReadFile(dbPath)
	if err != nil {
		return map[string]string{}
	}

	var db map[string]string
	json.Unmarshal(data, &db)
	return db
}

func saveDB(db map[string]string) {
	data, _ := json.MarshalIndent(db, "", "  ")
	os.WriteFile(dbPath, data, 0644)
}
