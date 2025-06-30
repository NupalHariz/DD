package assignment

const (
	insertAssignment = `
		INSERT INTO
			assignments(
				user_id,
				category_id,
				name,
				deadline,
				status,
				priority
			)VALUES(
				:user_id,
				:category_id,
				:name,
				:deadline,
				:status,
				:priority
			)
	`

	updateAssignment = `
		UPDATE
			assignments
	`

	readAssignment = `
		SELECT
			id,
			user_id,
			category_id,
			name,
			deadline,
			status,
			priority
		FROM
			assignments
	`

	countAssignments = `
		SELECT
			COUNT(*)
		FROM
			assignments
	`
)
