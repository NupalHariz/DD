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
)
