package gosplash

import (
	"encoding/json"
	"testing"
)

func TestCollection_Collections(t *testing.T) {
	collectionGroup := Collection{Client: NewClient(accessKey, SetClientProxy("http://127.0.0.1:7890"))}
	res, response, err := collectionGroup.Collections(CollectionReq{
		Page:    "1",
		PerPage: "1",
	})
	if err != nil {
		t.Fatal(err)
	}
	data, _ := json.Marshal(res)
	t.Log(string(data))
	t.Log(response.OriginResponseBody)
}
