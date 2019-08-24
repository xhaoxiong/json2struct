package service

import (
	"github.com/ChimeraCoder/gojson"
	"github.com/spf13/cast"
	"json2struct/repositories"
	"strings"
)

type TestService struct {
	repo *repositories.TestRepositories
}

func NewTestService() *TestService {
	return &TestService{repo: repositories.NewTestRepositories()}
}

func (t *TestService) ParseJsonStr(m map[string]interface{}, userAgent, Ip, Addr string) (string, bool) {
	var parse gojson.Parser
	value := cast.ToString(m["input_val"])
	if !t.CheckParams(m) {
		return "", false
	}

	if m["cate"] == "YAML" {
		parse = gojson.ParseYaml
	} else {
		parse = gojson.ParseJson
	}

	i := strings.NewReader(value)
	if bytes, err := gojson.Generate(i,
		parse,
		"gojson",
		"json2struct",
		[]string{"json"},
		cast.ToBool(m["contact"])); err != nil {
		return "", false
	} else {
		return string(bytes), t.repo.Create(userAgent, Ip, Addr, string(bytes))
	}
}

func (t *TestService) CheckParams(m map[string]interface{}) bool {
	return !(m["input_val"] == "" || m["cate"] == "")
}

func (t *TestService) GetRotateCount(count *int) {
	t.repo.GetRotateCount(count)
}
