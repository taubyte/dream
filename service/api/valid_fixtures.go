package api

import (
	httpIface "github.com/taubyte/http"
)

func validFixtures() {
	serviceApi.GET(&httpIface.RouteDefinition{
		Path:    "/spec/fixtures",
		Handler: fixturesHandler,
	})
}

func fixturesHandler(ctx httpIface.Context) (interface{}, error) {
	return mv.ValidFixtures(), nil
}
