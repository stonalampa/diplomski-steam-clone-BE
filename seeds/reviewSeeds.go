package seeds

import repo "main/repository"

var reviews = []repo.Review{
	{
		ID:      generateId("63bee09d0d258a28b131251f"),
		UserId:  generateId("63bd6bd756fa7318db852016"),
		GameId:  generateId("63bed828e88e68df55def4e4"),
		Comment: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus molestie arcu nisi, at facilisis ligula rutrum id. Nullam nec metus sodales lectus vestibulum scelerisque sit amet non massa.",
	},
	{
		ID:      generateId("63bee0b2fb35bceee5aec758"),
		UserId:  generateId("63bd6d40fb221e7af31f30dd"),
		GameId:  generateId("63bed83dc0378d4abb338e72"),
		Comment: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus molestie arcu nisi, at facilisis ligula rutrum id. Nullam nec metus sodales lectus vestibulum scelerisque sit amet non massa.",
	},
	{
		ID:      generateId("63bee0afff5998b77251d340"),
		UserId:  generateId("63bd6ce339dd5484c8d7b7a9"),
		GameId:  generateId("63bed81ec1bd88d167f66a8f"),
		Comment: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus molestie arcu nisi, at facilisis ligula rutrum id. Nullam nec metus sodales lectus vestibulum scelerisque sit amet non massa.",
	},
	{
		ID:      generateId("63bee0ab8796ecb9f52e5dc4"),
		UserId:  generateId("63bd6d29bfdab9586c5bf162"),
		GameId:  generateId("63bed843597757fc6bb4953e"),
		Comment: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus molestie arcu nisi, at facilisis ligula rutrum id. Nullam nec metus sodales lectus vestibulum scelerisque sit amet non massa.",
	},
	{
		ID:      generateId("63bee0a87cfd99e13547fafc"),
		UserId:  generateId("63bd6d1fab372bf383a2dbc7"),
		GameId:  generateId("63bed848240a8702cbdf09d4"),
		Comment: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus molestie arcu nisi, at facilisis ligula rutrum id. Nullam nec metus sodales lectus vestibulum scelerisque sit amet non massa.",
	},
}
