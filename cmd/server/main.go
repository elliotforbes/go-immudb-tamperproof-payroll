package main

import (
	"log"

	"github.com/TutorialEdge/tamperproof-payroll/internal/immudb"
)

func main() {
	db, err := immudb.NewImmuDB()
	if err != nil {
		log.Fatal("failed to connect to immudb")
	}

	err = db.UpdateSalary("12345", 100000)
	if err != nil {
		log.Fatal("failed to update salary")
	}
}
