package schemas

type Veracode struct {
	ApiCredetial Auth
}

type Auth struct {
	ApiKeyID     string
	ApiKeySecret string
	ApiUrl       string
}