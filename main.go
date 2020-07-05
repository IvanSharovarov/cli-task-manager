package main

import (
	"fmt"
	"github.com/IvanSharovarov/cli-task-manager/cmd"
	"github.com/IvanSharovarov/cli-task-manager/db"
	"github.com/mitchellh/go-homedir"
	"os"
	"path/filepath"
)

func main() {
	home, err := homedir.Dir()
	must(err)

	dbPath := filepath.Join(home, "tasks.db")
	err = db.Init(dbPath)
	must(err)

	must(cmd.RootCmd.Execute())
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}