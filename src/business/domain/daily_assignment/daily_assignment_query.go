package dailyassignment

const (
	insertDailyAssignment = `
		INSERT INTO
			daily_assignments(
				user_id,
				name	
			)
		VALUES(
			:user_id,
			:name
		)
	`

	updateDailyAssignment = `
		UPDATE
			daily_assignments
	`

	readDailyAssignment = `
		SELECT
			id,
			user_id,
			name,
			is_done
		FROM
			daily_assignments
	`

	updateDailyAssignmentToFalse = `
		UPDATE
			daily_assignments
		SET
			is_done = false
		WHERE
			is_done = true
	`
)
