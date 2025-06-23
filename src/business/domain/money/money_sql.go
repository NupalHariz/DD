package money

const (
	insertMoney = `
		INSERT INTO moneys(
			user_id,
			amount,
			category_id,
			type
		) VALUES(
			:user_id,
			:amount,
			:category_id,
			:type 
		)
	`
)
