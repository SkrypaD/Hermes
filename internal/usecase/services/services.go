package services

type CustomClaims struct {
	UserID int    `json:"user_id"`
	RoleID int    `json:"role_id"`
	Role   string `json:"role"`
	Login  string `json:"login"`
}

type TokenGeneratorService interface {
	// Attempts to generate token with provided claims.
	GenerateToken(claims CustomClaims) (string, error)
	// Verifies validity of provided token.
	VerifyToken(token string) (*CustomClaims, error)
}

type HasherService interface {
	// Attempts to hash provided password.
	HashPassword(password string) (string, error)
	// Checks if provided password is the same as original.
	VerifyPassword(hashedOriginal string, candiate string) error
}
