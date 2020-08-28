package gotest

import "errors"

func Division(a, b int64) (int64, error) {
	if b == 0 {
		return 0, errors.New("分母不能为0")
	}
	return a / b, nil
}
