package main

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
)

func Test_GetFromDB(t *testing.T) {
	ctrl:=gomock.NewController(t)
	defer ctrl.Finish() // 断言DB.Get()方法是否被调用

	m := NewMockDB(ctrl)
	m.EXPECT().Get(gomock.Not("Tom")).Return(100, errors.New("not exist")).Times(2)

	//if v := GetFromDB(m, "Tom"); v != -1 {
	//	t.Fatal("expected -1, but got", v)
	//}

	GetFromDB(m, "ABC")
	GetFromDB(m, "DEF")

}
