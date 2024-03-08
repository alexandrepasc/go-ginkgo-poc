package configurations

var Endpoints = endpoints{
	BaseURL:  "https://meowfacts.herokuapp.com/",
}

type endpoints struct {
	BaseURL   string
}
