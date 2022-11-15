// Package model models that use in this project
package model

// Vendor : struct for vendor
type User struct {
	ID           string `bson,json:"id"`
	Name         string `bson,json:"name"`
	Password     string `bson,json:"password"`
	Balance      uint   `bson,json:"balance"`
	RefreshToken string `bson,json:"refreshtoken"`
	UserType     string `bson,json:"usertype"`
}

//ПокупкаТовара()----//СписаниеСредствСБаланса()

// Config struct create config
type Config struct {
	CurrentDB     string `env:"CURRENT_DB" envDefault:"postgres"`
	PostgresDBURL string `env:"POSTGRES_DB_URL"`
	MongoDBURL    string `env:"MONGO_DB_URL"`
	JwtKey        []byte `env:"JWT-KEY" `
}
