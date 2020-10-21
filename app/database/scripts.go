package database

import "fmt"

var CreateCatch = fmt.Sprintf(`

	INSERT INTO %s.catch(
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
	) 
	VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)

`, DbName)
