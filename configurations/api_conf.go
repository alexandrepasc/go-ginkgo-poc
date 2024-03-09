package configurations

var Endpoints = endpoints{
	Meowfacts: meowfacts{
		BaseURL: "https://meowfacts.herokuapp.com/",
	},
	Catfacts: catfacts{
		BaseURL: "https://cat-fact.herokuapp.com/facts",
	},
}

type endpoints struct {
	Meowfacts meowfacts
	Catfacts  catfacts
}

type meowfacts struct {
	BaseURL string
}

type catfacts struct {
	BaseURL string
}
