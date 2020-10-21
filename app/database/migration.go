package database

import (
	"log"
	"os"
)

func MigrateDb() error {
	db, err := GetDatabase()
	if err != nil {
		return err
	}

	if os.Getenv("CLEAR") != "" {
		log.Println("clearing database")
		_, err = db.Exec(clearDatabase)
		if err != nil {
			return err
		}
	}

	for _, v := range executionSeries{
		_, err = db.Exec(v)
		if err != nil {
			return err
		}
	}

	return nil
}
