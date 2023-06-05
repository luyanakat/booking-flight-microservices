package helper

import "mock-project/graphql/graph/model"

func CheckNil(input *string) string {
	if input == nil {
		return ""
	}
	return *input
}

func CheckNilUser(input *model.User) int {
	if input == nil {
		return 0
	}
	return 1
}
