package core

import "os"

func Remove(name string) error {
	db := loadDB()

	path := db[name]
	if path == "" {
		return nil
	}

	os.RemoveAll(path)

	delete(db, name)
	saveDB(db)

	return nil
}
