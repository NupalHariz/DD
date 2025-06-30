package assignmentcategory

const (
	insertAssignmentCategory = `
		INSERT INTO
			assignment_categories(
				user_id,
				name
			)
		VALUES(
			:user_id,
			:name
		)
	`

	readAssignmentCategory = `
		SELECT
			id,
			user_id,
			name
		FROM
			assignment_categories
	`
)
