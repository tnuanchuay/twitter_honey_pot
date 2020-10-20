package repository

import (
	"log"
	"twitter_honey_pod/app/database"
	"twitter_honey_pod/app/model"
)

func CreateHoney(honey model.Honey) (*model.Honey, error) {
	db, err := database.GetDatabase()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	statement, err := db.Prepare(`INSERT INTO ` + database.DbName + `.honey(redirect_to, url, user_id, create_date) VALUES (?, ?, ?, ?)`)
	if err != nil {
		return nil, err
	}
	defer statement.Close()

	result, err := statement.Exec(honey.RedirectTo, honey.Url, honey.UserId, honey.CreateDate)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	honey.Id = id

	return &honey, nil
}

func GetHoney(honey *model.Honey) ([]model.Honey, error) {
	db, err := database.GetDatabase()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := `SELECT create_date, user_id, url, redirect_to, id FROM ` + database.DbName + `.honey`
	params := make([]interface{}, 0)
	if honey != nil {
		query = query + " WHERE"

		if honey.CreateDate != "" {
			query = query + " create_date = ?"
			params = append(params, honey.CreateDate)
		}

		if honey.UserId != "" {
			query = query + " user_id = ?"
			params = append(params, honey.UserId)
		}

		if honey.Url != "" {
			query = query + " url = ? "
			params = append(params, honey.Url)
		}

		if honey.RedirectTo != "" {
			query = query + " redirect_to = ? "
			params = append(params, honey.RedirectTo)
		}
	}

	statement, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer statement.Close()

	rows, err := statement.Query(params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	result := make([]model.Honey, 0)
	for rows.Next() {
		var createDate, userId, url, redirectTo string
		var id int64

		err = rows.Scan(&createDate, &userId, &url, &redirectTo, &id)
		if err != nil {
			log.Println(err)
			continue
		}

		row := model.Honey{
			RedirectTo: redirectTo,
			Url:        url,
			UserId:     userId,
			CreateDate: createDate,
			Id:         id,
		}

		result = append(result, row)
	}

	return result, nil
}
