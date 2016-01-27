package goutyl

import (
	"net/http"
	"strconv"
	"strings"
)

const (
	allowedOrigin    = "Access-Control-Allow-Origin"
	allowedMethods   = "Access-Control-Allow-Methods"
	allowCredentials = "Access-Control-Allow-Credentials"
	maxAge           = "Access-Control-Max-Age"
	allowedHeaders   = "Access-Control-Allow-Headers"
	defaultMaxAge    = 86400
)

var (
	defaultAllowedOrigins = []string{"*"}
	defaultAllowedHeaders = []string{"X-Requested-With", "X-HTTP-Method-Override", "Content-Type", "Accept", "Authorization"}
	defaultAllowedMethods = []string{"POST", "GET", "PUT", "DELETE", "OPTIONS"}
)

type CORSOptions struct {
	AllowedOrigins   []string
	AllowedMethods   []string
	AllowedHeaders   []string
	AllowCredentials bool
	HandleOptions    bool
	OptionsHandler   func()
	MaxAge           int
}

type CORS struct {
	optionsMap map[string]string
	options    CORSOptions
}

func (c CORS) Handle(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for i, v := range c.optionsMap {
			w.Header().Set(i, v)
		}
		if c.options.HandleOptions && r.Method == "OPTIONS" {
			apiResponse := &ApiResponse{Message: "All good."}
			apiResponse.Status(http.StatusOK).Json(w)
		}
	})
}

func (c *CORS) AddHeader(header string) *CORS {
	var headerList []string
	if c.optionsMap[allowedHeaders] != "" {
		headerList = strings.Split(c.optionsMap[allowedHeaders], ",")
		headerList = append(headerList, header)
	} else {
		headerList = []string{header}
	}
	c.optionsMap[allowedHeaders] = strings.Join(headerList, ",")
	return c
}

func (c *CORS) New(opts CORSOptions) *CORS {
	c.options = opts

	if !c.options.HandleOptions && c.options.OptionsHandler == nil {
		c.options.HandleOptions = true
	}
	if len(opts.AllowedOrigins) <= 0 {
		opts.AllowedOrigins = []string{"*"}
	}
	if len(opts.AllowedMethods) <= 0 {
		opts.AllowedMethods = defaultAllowedMethods
	}
	if len(opts.AllowedHeaders) <= 0 {
		opts.AllowedHeaders = defaultAllowedHeaders
	}
	if opts.MaxAge == 0 {
		opts.MaxAge = defaultMaxAge
	}
	c.optionsMap[allowedOrigin] = strings.Join(opts.AllowedOrigins[:], ",")
	c.optionsMap[allowedMethods] = strings.Join(opts.AllowedMethods[:], ",")
	c.optionsMap[allowCredentials] = strconv.FormatBool(opts.AllowCredentials)
	c.optionsMap[maxAge] = strconv.Itoa(opts.MaxAge)
	if c.optionsMap[allowedHeaders] != nil {
		headerList := strings.Split(c.optionsMap[allowedHeaders], ",")
		headerList = append(opts.AllowedHeaders, headerList)
		c.optionsMap[allowedHeaders] = strings.Join(headerList[:], ",")
	} else {
		c.optionsMap[allowedHeaders] = strings.Join(opts.AllowedHeaders[:], ",")
	}
	return c
}

func (c *CORS) Default() *CORS {
	return c.New(CORSOptions{
		AllowedOrigins:   defaultAllowedOrigins,
		AllowedHeaders:   defaultAllowedHeaders,
		AllowedMethods:   defaultAllowedMethods,
		MaxAge:           defaultMaxAge,
		AllowCredentials: false,
	})
}
