package impl

import v1 "github.com/dgodyna/bazel-playground/api/petstore/v1"

func GetUser() *v1.User {
	return &v1.User{
		Username: "test",
	}
}
