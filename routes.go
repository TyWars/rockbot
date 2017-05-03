package main

// Route is the model for the router setup
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc HandlerFunc
}

// Routes are the main setup for our Router
type Routes []Route

var routes = Routes{
	Route{"Healthcheck", "GET", "/v1/healthcheck", HealthcheckHandler},
	//=== USERS ===
	Route{"ListUsers", "GET", "/v1/users", ListUsersHandler},
	Route{"GetUser", "GET", "/v1/users/{uid:[a-z0-9]+}", GetUserHandler},
	Route{"CreateUser", "POST", "/v1/users", CreateUserHandler},
	//Route{"UpdateUser", "PUT", "/v1/users/{uid:[a-z0-9]+}", UpdateUserHandler},
	Route{"DeleteUser", "DELETE", "/v1/users/{uid:[a-z0-9]+}", DeleteUserHandler},

	//=== PASSPORTS === Not implemented yet, defaulting to unimplemented PassportsHandler
	/*
		Route{"GetUserPassport", "GET", "/users/{uid}/passports", PassportsHandler},
		Route{"GetPassport", "GET", "/passports/{pid:[0-9]+}", PassportsHandler},
		Route{"CreateUserPassport", "POST", "/users/{uid}/passports", PassportsHandler},
		Route{"UpdatePassport", "PUT", "/passports/{pid:[0-9]+}", PassportsHandler},
		Route{"DeletePassport", "DELETE", "/passports/{pid:[0-9]+}", PassportsHandler},
	*/
}
