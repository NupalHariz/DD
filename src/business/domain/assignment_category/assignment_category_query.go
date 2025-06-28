package assignmentcategory

const(
	insertAssignmentCategory=`
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
)