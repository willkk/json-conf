package jsonconf

import (
	"testing"
)

type MyConf struct {
	Timeout int `json:"timeout"`
	Ip string `json:"ip"`
	Port int `json:"port"`
	Redis string `json:"redis"`
	Mysql string `json:"mysql"`
	Log string `json:"log"`
}

func TestUnmarshal(t *testing.T) {
	myconf := &MyConf{}
	err := Unmarshal("./testapp.conf", myconf)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("Parsed json:%v", *myconf)
}
