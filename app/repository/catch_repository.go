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

	statement, err := db.Prepare(database.CreateCatch)
	if err != nil {
		return nil, err
	}
	defer statement.Close()

	result, err := statement.Exec(
		catch.HoneyId,
		catch.HitTime,
		catch.Ip,
		catch.ReferralUrl,
		catch.XForwardedFor,
		catch.City,
		catch.CountryName,
		catch.CountryCode,
		catch.ContinentName,
		catch.Latitude,
		catch.Longitude,
		catch.AsnId,
		catch.AsnName,
		catch.AsnDomain,
		catch.AsnRoute,
		catch.AsnType,
		catch.IsTor,
		catch.IsProxy,
		catch.IsAnonymous,
		catch.IsKnownAttacker,
		catch.IsKnownAbUser,
		catch.IsThreat,
		catch.IsBogon,
		)

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