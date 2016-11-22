package viewmodels

type Home struct {
	Title string
	Active string
}

func GetHome() Home {
	result := Home{
		Title: "Lemode Stand Supply",
		Active: "home",
	}

	return result
}
