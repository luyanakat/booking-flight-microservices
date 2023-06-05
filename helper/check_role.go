package helper

import "mock-project/graphql/graph/model"

func CheckAdmin(user *model.User) bool {
	if user.AccessID == 3 {
		return true
	}
	return false
}
