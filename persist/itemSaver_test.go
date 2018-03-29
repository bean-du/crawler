package persist

import (
	"testing"

	"encoding/json"

	"github.com/bean-du/crawler/engine"
	"github.com/bean-du/crawler/model"
	"github.com/olivere/elastic"
	"golang.org/x/net/context"
)

func TestSave(t *testing.T) {

	profile := engine.Item{
		Url:  "https://city.zhenai.com/u/1234567890",
		Type: "zhenai",
		Id:   "1234567890",
		Payload: model.Profile{
			Name:   "bean",
			Age:    30,
			Gander: "男",
			Height: 180,
			Weight: 65,
			House:  "已购房",
			Car:    "未购车",
			Income: "10000-20000元/月",
		},
	}

	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	const index = "dating_test"
	err = save(client, profile, index)
	if err != nil {
		panic(err)
	}

	resp, err := client.Get().Index(index).Type(profile.Type).Id(profile.Id).Do(context.Background())
	if err != nil {
		panic(err)
	}
	t.Logf("%s", resp.Source)

	var actual engine.Item

	json.Unmarshal(*resp.Source, &actual)

	actualProfile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile

	if actual != profile {
		t.Errorf("Error :Got: %v, profile : %v", actual, profile)
	}
}
