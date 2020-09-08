package mysql

const (
	queryInsertUser       = `INSERT INTO users (username, password, name) VALUES (?, ?, ?)`
	querySelectUser       = `SELECT user_id, username, password, name FROM users WHERE username=?`
	querySelectDetailUser = `SELECT user_id, username, password, name FROM users WHERE user_id=?`
	QueryUpdateUser       = `
		UPDATE users
		SET
			username = ?,
			name = ?
		WHERE user_id = ?
	`
)
