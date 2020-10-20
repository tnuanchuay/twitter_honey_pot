package repository

import (
	"twitter_honey_pod/app/database"
	"twitter_honey_pod/app/model"
)

func CreateCatch(catch model.Catch) (*model.Catch, error){
	db, err := database.GetDatabase()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	statement, err := db.Prepare(`INSERT INTO ` + database.DbName + `.catch(honey_id, hit_time, ip, referer_url) VALUES(?, ?, ?, ?)`)
	if err != nil {
		return nil, err
	}
	defer statement.Close()

	result, err := statement.Exec(catch.HoneyId, catch.HitTime, catch.Ip, catch.ReferralUrl)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	catch.Id = id

	return &catch, nil
}