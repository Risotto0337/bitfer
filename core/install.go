package core

import (
	"os"
	"os/exec"
)

func Install(src string, name string) error {
	target := "/opt/bitfer/" + name
	os.MkdirAll(target, 0755)

	cmd := exec.Command("bash", "-c", "cp -r "+src+"/* "+target)
	err := cmd.Run()
	if err != nil {
		return err
	}

	db := loadDB()
	db[name] = target
	saveDB(db)

	return nil
}
