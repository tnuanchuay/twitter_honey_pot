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
		_, err = db.Exec(`
	
	DROP DATABASE ` + DbName + `;
	
	`)
		if err != nil {
			return err
		}
	}

	_, err = db.Exec(`
	
	CREATE DATABASE IF NOT EXISTS ` + DbName + `;
	
	`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`

	CREATE TABLE IF NOT EXISTS ` + DbName + `.honey (
		id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
		redirect_to VARCHAR(1000) NOT NULL,
		user_id VARCHAR(100) NOT NULL,
		create_date TIMESTAMP NOT NULL,
		url VARCHAR(200)
	);
	
	`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
	
	CREATE TABLE IF NOT EXISTS ` + DbName + `.catch (
		id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
		honey_id INT NOT NULL,
		hit_time TIMESTAMP NOT NULL,
		ip VARCHAR(1024),
		referer_url VARCHAR(300),
		x_forwarded_for VARCHAR(100)
	)
	
	`)

	if err != nil {
		return err
	}

	return nil
}
