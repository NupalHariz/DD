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

	readMoney = `
		SELECT
			id,
			user_id,
			amount,
			category_id,
			type
		FROM
			moneys
	`

	updateMoney = `
		UPDATE
			moneys
	`
)
