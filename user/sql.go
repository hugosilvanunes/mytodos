package user

const (
	INSERT_INTO_USERS_SQL   = "INSERT INTO users(email, password, name) VALUES (:email, :password, :name)"
	FIND_BY_EMAIL_USERS_SQL = "SELECT * FROM users WHERE email=$1"
)
