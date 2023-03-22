package configs

type AuthConfig struct {
	AccessHmacSecretKey     string `envconfig:"ACCESS_HMAC_SECRET_KEY" mapstructure:"access_hmac_secret_key"`
	RefreshHmacSecretKey    string `envconfig:"REFRESH_HMAC_SECRET_KEY" mapstructure:"refresh_hmac_secret_key"`
	JwtExpiration           int    `envconfig:"JWT_EXPRIVATION" mapstructure:"jwt_exprivation" default:"30"`
	ResetPasswordExpiration int    `envconfig:"RESET_PASSWORD_EXPIRATION" mapstructure:"reset_password_expiration" default:"5"`
	VerifyEmailExpiration   int    `envconfig:"VERIFY_EMAIL_EXPIRATION" mapstructure:"verify_email_expiration" default:"5"`
}

func LoadAuthConfig() (*AuthConfig, error) {
	//TODO
	return nil, nil
}