package budget

const (
	insertBudget = `
	INSERT INTO budgets(
		user_id, 
		category_id, 
		amount, 
		time_period
	) VALUES(
		:user_id,
		:category_id,
		:amount,
		:type
	)
	`

	updateCurrentExpense = `
		UPDATE budgets
		SET current_expense = current_expense + :current_expense
		WHERE
		user_id = :user_id
		AND
		category_id = :category_id
	`
)
