package museum

const (
	createMuseumQuery = `
		INSERT INTO museum (
			id_owner, inn, kpp, founder, museum_activity_in_charter, name, museum_legal_status,
			is_memorial_reserve_museum, is_historical_memorial_reserve, is_art_museum,
			is_museum_reserve, is_estate_museum, is_palace_park_ensemble,
			is_historical_architectural_reserve, annual_visitor_capacity,
			internal_visitors_count, external_visitors_count,
			is_valuable_cultural_heritage, valuable_museum_items_count
		)
		VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10,
			$11, $12, $13, $14, $15, $16, $17, $18, $19
		)
		RETURNING id`
	// TODO: исправить костыль (убрал вывод inn)
	getMuseumByIDQuery = `
		SELECT
			id, id_owner, kpp, founder, museum_activity_in_charter, name, museum_legal_status,
			is_memorial_reserve_museum, is_historical_memorial_reserve, is_art_museum,
			is_museum_reserve, is_estate_museum, is_palace_park_ensemble,
			is_historical_architectural_reserve, annual_visitor_capacity,
			internal_visitors_count, external_visitors_count,
			is_valuable_cultural_heritage, valuable_museum_items_count
		FROM museum
		WHERE id = $1`

	updateMuseumQuery = `
		UPDATE museum SET
			id_owner = :id_owner,
			inn = :inn,
			kpp = :kpp,
			founder = :founder,
			museum_activity_in_charter = :museum_activity_in_charter,
			name = :name,
			museum_legal_status = :museum_legal_status,
			is_memorial_reserve_museum = :is_memorial_reserve_museum,
			is_historical_memorial_reserve = :is_historical_memorial_reserve,
			is_art_museum = :is_art_museum,
			is_museum_reserve = :is_museum_reserve,
			is_estate_museum = :is_estate_museum,
			is_palace_park_ensemble = :is_palace_park_ensemble,
			is_historical_architectural_reserve = :is_historical_architectural_reserve,
			annual_visitor_capacity = :annual_visitor_capacity,
			internal_visitors_count = :internal_visitors_count,
			external_visitors_count = :external_visitors_count,
			is_valuable_cultural_heritage = :is_valuable_cultural_heritage,
			valuable_museum_items_count = :valuable_museum_items_count
		WHERE id = :id`

	deleteMuseumQuery = `
		DELETE FROM museum
		WHERE id = $1`

	findMuseumByINNQuery = `
		SELECT
			id, id_owner, kpp, founder, museum_activity_in_charter, name, museum_legal_status,
			is_memorial_reserve_museum, is_historical_memorial_reserve, is_art_museum,
			is_museum_reserve, is_estate_museum, is_palace_park_ensemble,
			is_historical_architectural_reserve, annual_visitor_capacity,
			internal_visitors_count, external_visitors_count,
			is_valuable_cultural_heritage, valuable_museum_items_count
		FROM museum
		WHERE LOWER(inn) = LOWER($1)`

	findMuseumsByOwnerQuery = `
		SELECT
			id, id_owner, kpp, founder, museum_activity_in_charter, name, museum_legal_status,
			is_memorial_reserve_museum, is_historical_memorial_reserve, is_art_museum,
			is_museum_reserve, is_estate_museum, is_palace_park_ensemble,
			is_historical_architectural_reserve, annual_visitor_capacity,
			internal_visitors_count, external_visitors_count,
			is_valuable_cultural_heritage, valuable_museum_items_count
		FROM museum
		WHERE id_owner = $1`

	listMuseumsQuery = `
		SELECT
			id, id_owner, kpp, founder, museum_activity_in_charter, name, museum_legal_status,
			is_memorial_reserve_museum, is_historical_memorial_reserve, is_art_museum,
			is_museum_reserve, is_estate_museum, is_palace_park_ensemble,
			is_historical_architectural_reserve, annual_visitor_capacity,
			internal_visitors_count, external_visitors_count,
			is_valuable_cultural_heritage, valuable_museum_items_count
		FROM museum
		WHERE ($1::text IS NULL OR name ILIKE '%' || $1 || '%')
			AND (
				$2::text IS NULL OR
				($2 = 'memorial_reserve' AND is_memorial_reserve_museum = true) OR
				($2 = 'historical_memorial_reserve' AND is_historical_memorial_reserve = true) OR
				($2 = 'art' AND is_art_museum = true) OR
				($2 = 'museum_reserve' AND is_museum_reserve = true) OR
				($2 = 'estate' AND is_estate_museum = true) OR
				($2 = 'palace_park_ensemble' AND is_palace_park_ensemble = true) OR
				($2 = 'historical_architectural_reserve' AND is_historical_architectural_reserve = true)
			)
		ORDER BY id
		LIMIT $3 OFFSET $4`

	countMuseumsQuery = `
		SELECT COUNT(*)
		FROM museum
		WHERE ($1::text IS NULL OR name ILIKE '%' || $1 || '%')
			AND (
				$2::text IS NULL OR
				($2 = 'memorial_reserve' AND is_memorial_reserve_museum = true) OR
				($2 = 'historical_memorial_reserve' AND is_historical_memorial_reserve = true) OR
				($2 = 'art' AND is_art_museum = true) OR
				($2 = 'museum_reserve' AND is_museum_reserve = true) OR
				($2 = 'estate' AND is_estate_museum = true) OR
				($2 = 'palace_park_ensemble' AND is_palace_park_ensemble = true) OR
				($2 = 'historical_architectural_reserve' AND is_historical_architectural_reserve = true)
			)`
)
