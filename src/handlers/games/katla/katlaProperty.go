package katla

type katlaData struct {
	Props katlaProps `json:"props"`
}

type katlaProps struct {
	PageProps katlaPageProps `json:"pageProps"`
}

type katlaPageProps struct {
	Hashed string `json:"hashed"`
}
