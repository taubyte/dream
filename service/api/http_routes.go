package api

func setUpHttpRoutes() {
	corsHttp()

	statusHttp()
	lesMiesrablesHttp()
	fixtureHttp()
	idHttp()

	// Inject
	injectSimpleHttp()
	injectServiceHttp()
	injectUniverseHttp()

	// Kill
	killServiceHttp()
	killSimpleHttp()
	killNodeIdHttp()
	killUniverseHttp()

	// Get
	validClients()
	validServices()
	validFixtures()
}
