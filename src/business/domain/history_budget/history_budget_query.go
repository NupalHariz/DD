package historybudget

const (
	insertHistoryBudget=`
		INSERT INTO history_budgets(
			user_id,
			budget_id,
			category_id,
			spent,
			planned,
			type,
			period_start,
			period_end	
		) VALUES (
			:user_id,
			:budget_id,
			:category_id,
			:spent,
			:planned,
			:type,
			:period_start,
			:period_end
		)
	`
)