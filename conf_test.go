package goini

import (
	"fmt"
	"testing"
)

func TestExists(t *testing.T) {
	conf, err := SetConfig("./conf/conf.ini")
	if err != nil {
		t.Errorf("SetConfig err %v", err)
	}
	username := conf.GetValue("database", "username")
	fmt.Println(username) //root
	conf.DeleteValue("database", "username")
	username = conf.GetValue("database", "username")
	if len(username) == 0 {
		fmt.Println("username is not exists") //this stdout username is not exists
	}
	conf.SetValue("database", "username", "widuu")
	username = conf.GetValue("database", "username")
	fmt.Println(username) //widuu

	data := conf.ReadList()
	fmt.Println(data)
}

func TestNotExists(t *testing.T) {
	_, err := SetConfig("./conf/not_exists.ini")
	if err == nil {
		t.Errorf("configuring file expects not exist")
	}
}
