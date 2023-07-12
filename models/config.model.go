package models

// default values given in this struct
//
//	type ConfigDatabase struct {
//		Port     string `env:"PORT" env-default:"3306"`
//		Host     string `env:"HOST" env-default:"127.0.0.1"`
//		Name     string `env:"NAME" env-default:"Ali"`
//		User     string `env:"USER" env-default:"root"`
//		Password string `env:"PASSWORD"`
//		DbName   string `env:"DBNAME" env-default:"test"`
//	}
type ConfigDatabase struct {
	Port                string `env:"PORT"`
	Host                string `env:"HOST"`
	Name                string `env:"NAME"`
	User                string `env:"USER"`
	Password            string `env:"PASSWORD"`
	DbName              string `env:"DBNAME"`
	JWTSecretSigningKey string `env:"JWT_SIGN_KEY"`
}
