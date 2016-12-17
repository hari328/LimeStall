package viewmodels

type Category struct {
	ImageUrl string
	Title string
	Desc string
	IsRight bool
}

type Categories struct {
	Title string
	Active string
	Categories []Category
}


func GetCategories() Categories {

	displayedCategories := make([]Category, 0)

	juiceCategory := Category{
		ImageUrl: "lemon.png",
		Title: "Juices and Mixes",
		Desc: "Explore our wide assortment of juices and mixes expected bytoday's lemonade stand clientelle. Now featuring a full line of organic juices that are guaranteed to be obtained from trees that have never been treated with pesticides or artificial",
		IsRight: false,
	}

	supplyCategory := Category {
		ImageUrl: "kiwi.png",
		Title: "Signs and Advertising",
		Desc: "From paper cups to bio-degradable plastic to straws and napkins, LSS is your source for the sundries that keep your standrunning smoothly.",
		IsRight: true,
	}

	advertiseCategory := Category {
		ImageUrl: "pineapple.png",
		Title: "Signs and Advertising",
		Desc: "Sure, you could just wait for people to find your stand along the side of the road, but if you want to take it to the next level, our premium line of advertising supplies.",
		IsRight: false,
	}

	displayedCategories = append(displayedCategories, juiceCategory, supplyCategory, advertiseCategory)

	return Categories{
		Title: "Lemon Stand Society - shop",
		Active: "shop",
		Categories: displayedCategories,
	}
}