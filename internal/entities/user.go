package entities

type UserLogin struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	JwtToken string `json:"jwt_token"`
}

type UserReg struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	Email    string `json:"email" db:"email"`
	JwtToken string `json:"jwt_token"`
}
