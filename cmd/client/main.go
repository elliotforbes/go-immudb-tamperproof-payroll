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

	salary, err := db.GetVerifiedSalary("12345")
	if err != nil {
		log.Fatal("failed to update salary")
	}

	log.Printf("Engineer Salary: %d\n", salary)
}
