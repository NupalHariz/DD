package category

const (
	insertCategory = `
		INSERT INTO CATEGORIES(
			user_id,
			name
		) VALUES(
			:user_id,
			:name
		)
	`

	readCategories = `
		SELECT
			id,
			user_id,
			name
		FROM
			categories
	`
)
