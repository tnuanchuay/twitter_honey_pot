package database

import "fmt"

var clearDatabase = fmt.Sprintf(`
	
	DROP DATABASE IF EXISTS %s;
	
	`, DbName)

var executionSeries = []string {
	createDatabase,
	honeyDb,
	catchDb,
}

var createDatabase = fmt.Sprintf(`
	
	CREATE DATABASE IF NOT EXISTS %s;
	
	`, DbName)

var honeyDb = fmt.Sprintf(`

	CREATE TABLE IF NOT EXISTS %s.honey (
		id 					INT 			PRIMARY KEY NOT NULL 	AUTO_INCREMENT,
		redirect_to 		VARCHAR(1000) 				NOT NULL,
		user_id 			VARCHAR(100) 				NOT NULL,
		create_date 		TIMESTAMP 					NOT NULL,
		url 				VARCHAR(200)
	);
	
	`, DbName)

var catchDb = fmt.Sprintf(`
	
	CREATE TABLE IF NOT EXISTS %s.catch (
		id 					INT 			PRIMARY KEY	NOT NULL	AUTO_INCREMENT,
		honey_id 			INT 						NOT NULL,
		hit_time 			TIMESTAMP 					NOT NULL,
		ip 					VARCHAR(200),
		referer_url 		VARCHAR(300),
		x_forwarded_for 	VARCHAR(100),
		city 				VARCHAR(100),
		country_name 		VARCHAR(100),
		country_code 		VARCHAR(10),
		continent_name 		VARCHAR(100),
		latitude 			DOUBLE,
		longitude 			DOUBLE,
		asn_id 				VARCHAR(25),
		asn_name 			VARCHAR(200),
		asn_domain 			VARCHAR(200),
		asn_route 			VARCHAR(200),
		asn_type 			VARCHAR(25),
		is_tor 				BOOLEAN,
		is_proxy 			BOOLEAN,
		is_anonymous		BOOLEAN,
		is_known_attacker	BOOLEAN,
		is_known_abuser		BOOLEAN,
		is_threat			BOOLEAN,
		is_bogon			BOOLEAN
	)
	
	`, DbName)