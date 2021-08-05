package main

import (
	"github.com/ant0ine/go-json-rest/rest"
)

func newAPI() *rest.Api {
	api := rest.NewApi()

	api.Use([]rest.Middleware{
		&rest.AccessLogApacheMiddleware{},
		&rest.TimerMiddleware{},
		&rest.RecorderMiddleware{},
		&rest.PoweredByMiddleware{},
		&rest.RecoverMiddleware{
			EnableResponseStackTrace: true,
		},
		&rest.JsonIndentMiddleware{},
		&rest.ContentTypeCheckerMiddleware{},
		&rest.CorsMiddleware{
			AccessControlAllowCredentials: true,
			AccessControlMaxAge:           3600,
			AllowedMethods:                []string{"GET"},
			OriginValidator: func(origin string, request *rest.Request) bool {
				return true
			},
		},
	}...)

	return api
}

func addRoutes(api *rest.Api) (*rest.Api, error) {
	router, err := rest.MakeRouter(
		rest.Get("/twitter/recent/search/:handle", getTwitterRecentSearch),
	)
	if err != nil {
		return api, err
	}

	api.SetApp(router)

	return api, nil
}
