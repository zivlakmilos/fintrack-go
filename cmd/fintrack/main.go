package main

import (
	"fmt"

	"github.com/zivlakmilos/fintrack-go/pkg/core"
	"github.com/zivlakmilos/fintrack-go/pkg/db"
)

func main() {
	config, err := core.LoadConfig()
	if err != nil {
		fmt.Printf("got error %v\n", err)
		return
	}

	dbPath, err := core.GetDBPath(config.Year)
	if err != nil {
		return
	}

	sql, err := db.Open(dbPath)
	if err != nil {
		fmt.Printf("got error %v", err)
		return
	}
	defer sql.Close()

	err = db.InitDB(sql)
	if err != nil {
		fmt.Printf("got error %v\n", err)
		return
	}

	fmt.Printf("%v\n", config)
}
