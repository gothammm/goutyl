package cors

import (
	api "github.com/peek4y/goutyl/api"
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

type CORS struct {
	optionsMap map[string]string
	options    Options
}

func (c CORS) Handle(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for i, v := range c.optionsMap {
			w.Header().Set(i, v)
		}
		if c.options.HandleOptions && r.Method == "OPTIONS" {
			apiResponse := &api.ApiResponse{Message: "All good."}
			apiResponse.Status(http.StatusOK).Json(w)
		}
	})
}

func (c *CORS) AddHeader(header string) *CORS {
	var headerList []string
	if len(c.optionsMap) <= 0 {
		c.optionsMap = make(map[string]string)
	}
	if c.optionsMap[allowedHeaders] != "" {
		headerList = strings.Split(c.optionsMap[allowedHeaders], ",")
		headerList = append(headerList, header)
	} else {
		headerList = []string{header}
	}
	c.optionsMap[allowedHeaders] = strings.Join(headerList, ",")
	if c.options.AllowedHeaders != nil {
		c.options.AllowedHeaders = append(c.options.AllowedHeaders, headerList...)
	}
	return c
}

func (c *CORS) New(opts Options) *CORS {
	c.options = opts

	if len(c.optionsMap) <= 0 {
		c.optionsMap = make(map[string]string)
	}

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
	if c.optionsMap[allowedHeaders] != "" {
		headerList := strings.Split(c.optionsMap[allowedHeaders], ",")
		c.optionsMap[allowedHeaders] = strings.Join(append(opts.AllowedHeaders, headerList...)[:], ",")
	} else {
		c.optionsMap[allowedHeaders] = strings.Join(opts.AllowedHeaders[:], ",")
	}
	return c
}

func (c *CORS) Default() *CORS {
	return c.New(Options{
		AllowedOrigins:   defaultAllowedOrigins,
		AllowedHeaders:   defaultAllowedHeaders,
		AllowedMethods:   defaultAllowedMethods,
		MaxAge:           defaultMaxAge,
		AllowCredentials: false,
	})
}

func New(opts Options) *CORS {
	cors := &CORS{}
	return cors.New(opts)
}

func Default() *CORS {
	cors := &CORS{}
	return cors.Default()
}
