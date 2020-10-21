package repository

import (
	"log"
	"twitter_honey_pod/app/database"
	"twitter_honey_pod/app/model"
)

func CreateCatch(catch model.Catch) (*model.Catch, error) {
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
		catch.RefererUrl,
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

func GetCatch(catch *model.Catch) ([]model.Catch, error) {
	db, err := database.GetDatabase()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	params := make([]interface{}, 0)
	query := `
		SELECT
			id,
			honey_id, 
			hit_time, 
			ip, 
			referer_url, 
			x_forwarded_for, 
			city, 
			country_name, 
			country_code, 
			continent_name, 
			latitude, 
			longitude,
			asn_id,
			asn_name,
			asn_domain,
			asn_route,
			asn_type,
			is_tor,
			is_proxy,
			is_anonymous,
			is_known_attacker,
			is_known_abuser,
			is_threat,
			is_bogon
		FROM ` + database.DbName + `.catch
	`

	if catch != nil {
		query = query + " WHERE 1 = 1"

		if catch.HoneyId != 0 {
			query = query + " AND honey_id = ?"
			params = append(params, catch.HoneyId)
		}

		if catch.HitTime != "" {
			query = query + " AND hit_time = ?"
			params = append(params, catch.HitTime)
		}

		if catch.Ip != "" {
			query = query + " AND ip = ?"
			params = append(params, catch.Ip)
		}

		if catch.RefererUrl != "" {
			query = query + " AND referer_url = ?"
			params = append(params, catch.RefererUrl)
		}

		if catch.XForwardedFor != "" {
			query = query + " AND x_forwarded_for = ?"
			params = append(params, catch.XForwardedFor)
		}

		if catch.City != "" {
			query = query + " AND city = ?"
			params = append(params, catch.City)
		}

		if catch.CountryName != "" {
			query = query + " AND country_name = ?"
			params = append(params, catch.CountryName)
		}

		if catch.CountryCode != "" {
			query = query + " AND country_code = ?"
			params = append(params, catch.CountryCode)
		}

		if catch.ContinentName != "" {
			query = query + " AND continent_name = ?"
			params = append(params, catch.ContinentName)
		}

		if catch.Latitude != 0 {
			query = query + " AND latitude = ?"
			params = append(params, catch.Latitude)
		}

		if catch.Longitude != 0 {
			query = query + " AND longitude = ?"
			params = append(params, catch.Longitude)
		}

		if catch.AsnId != "" {
			query = query + " AND asn_id = ?"
			params = append(params, catch.AsnId)
		}

		if catch.AsnName != "" {
			query = query + " AND asn_name = ?"
			params = append(params, catch.AsnId)
		}

		if catch.AsnDomain != "" {
			query = query + " AND asn_domain = ?"
			params = append(params, catch.AsnId)
		}

		if catch.AsnRoute != "" {
			query = query + " AND asn_route = ?"
			params = append(params, catch.AsnId)
		}

		if catch.AsnType != "" {
			query = query + " AND asn_type = ?"
			params = append(params, catch.AsnId)
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

	result := make([]model.Catch, 0)
	for rows.Next() {
		var honeyPot int64
		var hitTime, ip, refererUrl, xForwardedFor, city, countryName, countryCode,
		continentName, asnId, asnName, asnDomain, asnRoute, asnType string
		var latitude, longitude float64
		var id int64
		var isTor, isProxy, isAnonymous, isKnownAttacker, isKnownAbuser, isThreat, isBogon bool

		err := rows.Scan(&id, &honeyPot, &hitTime, &ip, &refererUrl, &xForwardedFor, &city, &countryName,
			&countryCode, &continentName, &latitude, &longitude, &asnId, &asnName, &asnDomain, &asnRoute, &asnType,
			&isTor, &isProxy, &isAnonymous, &isKnownAttacker, &isKnownAbuser, &isThreat, &isBogon)
		if err != nil {
			log.Println(err)
			continue
		}

		row := model.Catch{
			Id:              id,
			HoneyId:         honeyPot,
			HitTime:         hitTime,
			Ip:              ip,
			RefererUrl:      refererUrl,
			XForwardedFor:   xForwardedFor,
			City:            city,
			CountryName:     countryName,
			CountryCode:     countryCode,
			ContinentName:   continentName,
			Latitude:        latitude,
			Longitude:       longitude,
			AsnId:           asnId,
			AsnName:         asnName,
			AsnDomain:       asnDomain,
			AsnRoute:        asnRoute,
			AsnType:         asnType,
			IsTor:           isTor,
			IsProxy:         isProxy,
			IsAnonymous:     isAnonymous,
			IsKnownAttacker: isKnownAttacker,
			IsKnownAbUser:   isKnownAbuser,
			IsThreat:        isThreat,
			IsBogon:         isBogon,
		}

		result = append(result, row)
	}

	return result, nil
}
