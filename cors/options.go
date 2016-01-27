package cors

type Options struct {
	AllowedOrigins   []string
	AllowedMethods   []string
	AllowedHeaders   []string
	AllowCredentials bool
	HandleOptions    bool
	OptionsHandler   func()
	MaxAge           int
}
