package store

import "server/model"

// Users data
var Users = []model.User{
	{
		ID:   1,
		Name: "Winter",
		MenuScores: []model.MenuScore{
			{
				MenuName: "짜장면",
				Score:    0,
				Count:    0,
			},
			{
				MenuName: "짬뽕",
				Score:    0,
				Count:    0,
			},
			{
				MenuName: "탕수육",
				Score:    0,
				Count:    0,
			},
			{
				MenuName: "볶음밥",
				Score:    0,
				Count:    0,
			},
		},
	},
}
