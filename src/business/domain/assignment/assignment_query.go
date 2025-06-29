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
)
