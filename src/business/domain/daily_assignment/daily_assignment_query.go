package dailyassignment

const(
	insertDailyAssignment=`
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

	updateDailyAssignment=`
		UPDATE
			daily_assignments
	`
)