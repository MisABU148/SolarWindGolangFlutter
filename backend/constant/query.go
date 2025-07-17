package constant

const CreateDeckQuery = `
	WITH unsorted_users AS (
		SELECT id, preferred_gym_time
		FROM users
		WHERE city_id = $1
		AND age BETWEEN ($2 - 10) AND ($2 + 10)
		AND LOWER(gender) = LOWER($4)
		AND LOWER(preferred_gender) = LOWER($3)
	),
	initial_user AS (
		SELECT preferred_gym_time FROM users WHERE id = $5
	),
	initial_user_sport AS (
		SELECT sport_id FROM user_sport WHERE user_id = $5
	),
	matched_users AS (
		SELECT u.id,
			u.preferred_gym_time,
			COUNT(us.sport_id) AS matched_sports_count,
			LENGTH(REPLACE(LPAD((u.preferred_gym_time & iu.preferred_gym_time)::bit(7)::text, 7, '0'), '0', '')) AS matched_days_count
		FROM unsorted_users u
		CROSS JOIN initial_user iu
		JOIN user_sport us ON u.id = us.user_id
		JOIN initial_user_sport ius ON us.sport_id = ius.sport_id
		GROUP BY u.id, u.preferred_gym_time, iu.preferred_gym_time
	)
	SELECT u.id, u.telegram_id, u.username, u.description, u.age, u.gender, u.preferred_gender,
		u.city_id, u.preferred_gym_time,
		mu.matched_sports_count, mu.matched_days_count
	FROM matched_users mu
	JOIN users u ON u.id = mu.id
	WHERE u.id != $5
	ORDER BY mu.matched_sports_count DESC, mu.matched_days_count DESC
	`
