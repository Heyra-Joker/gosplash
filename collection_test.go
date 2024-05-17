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
	t.Logf(response.requestURL)
}

func TestCollection_GetACollection(t *testing.T) {
	collectionGroup := Collection{Client: NewClient(accessKey, SetClientProxy("http://127.0.0.1:7890"))}
	res, response, err := collectionGroup.GetACollection("IzkerwNZhf8")
	if err != nil {
		t.Fatal(err)
	}
	data, _ := json.Marshal(res)
	t.Log(string(data))
	t.Log(response.OriginResponseBody)
	t.Logf(response.requestURL)
}

func TestCollection_GetACollectionPhotos(t *testing.T) {
	collectionGroup := Collection{Client: NewClient(accessKey, SetClientProxy("http://127.0.0.1:7890"))}
	res, response, err := collectionGroup.GetACollectionPhotos(GetACollectionPhotosReq{
		ID:          "tCTsj5uaetE",
		Page:        "1",
		PerPage:     "",
		Orientation: "",
	})
	if err != nil {
		t.Fatal(err)
	}
	data, _ := json.Marshal(res)
	t.Log(string(data))
	t.Log(response.OriginResponseBody)
	t.Logf(response.requestURL)
}

func TestCollection_RelatedCollections(t *testing.T) {
	collectionGroup := Collection{Client: NewClient(accessKey, SetClientProxy("http://127.0.0.1:7890"))}
	res, response, err := collectionGroup.RelatedCollections("tCTsj5uaetE")
	if err != nil {
		t.Fatal(err)
	}
	data, _ := json.Marshal(res)
	t.Log(string(data))
	t.Log(response.OriginResponseBody)
	t.Logf(response.requestURL)
}

func TestCollection_NewCollections(t *testing.T) {
	token := "aOicmLdr6SWlZ3EBIN3SDnpe4EXDRmTKJNdq4YtQ7ww"
	collectionGroup := Collection{Client: NewClient(accessKey, SetClientProxy("http://127.0.0.1:7890"))}
	res, response, err := collectionGroup.NewCollections(token, NewCollectionsReq{
		Title:       "test",
		Description: "this test collection",
		Private:     false,
	})
	if err != nil {
		t.Fatal(err)
	}
	data, _ := json.Marshal(res)
	t.Log(string(data))
	t.Log(response.OriginResponseBody)
	t.Logf(response.requestURL)
	// RSjGi3z9pfk
}

func TestCollection_UpdateCollection(t *testing.T) {
	token := "aOicmLdr6SWlZ3EBIN3SDnpe4EXDRmTKJNdq4YtQ7ww"
	collectionGroup := Collection{Client: NewClient(accessKey, SetClientProxy("http://127.0.0.1:7890"))}
	res, response, err := collectionGroup.UpdateCollection(token, UpdateCollectionReq{
		ID:          "BhMKvHpksrM",
		Title:       "test update",
		Description: "",
		Private:     false,
	})
	if err != nil {
		t.Fatal(err)
	}
	data, _ := json.Marshal(res)
	t.Log(string(data))
	t.Log(response.OriginResponseBody)
	t.Logf(response.requestURL)
	// RSjGi3z9pfk
}

func TestCollection_DeleteCollection(t *testing.T) {
	token := "aOicmLdr6SWlZ3EBIN3SDnpe4EXDRmTKJNdq4YtQ7ww"
	collectionGroup := Collection{Client: NewClient(accessKey, SetClientProxy("http://127.0.0.1:7890"))}
	response, err := collectionGroup.DeleteCollection(token, "bzMDxUazqgw")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response.OriginResponseBody)
	t.Logf(response.requestURL)
}

func TestCollection_AddPhoto2Collection(t *testing.T) {
	token := "aOicmLdr6SWlZ3EBIN3SDnpe4EXDRmTKJNdq4YtQ7ww"
	collectionGroup := Collection{Client: NewClient(accessKey, SetClientProxy("http://127.0.0.1:7890"))}
	res, response, err := collectionGroup.AddPhoto(token, AddPhotoReq{
		CollectionId: "BhMKvHpksrM",
		PhotoId:      "a-red-race-car-sitting-on-top-of-a-dirt-field-YFcEYRE9lvc",
	})
	if err != nil {
		t.Fatal(err)
	}
	data, _ := json.Marshal(res)
	t.Log(string(data))
	t.Log(response.OriginResponseBody)
	t.Logf(response.requestURL)
}

func TestCollection_RemovePhoto(t *testing.T) {
	token := "aOicmLdr6SWlZ3EBIN3SDnpe4EXDRmTKJNdq4YtQ7ww"
	collectionGroup := Collection{Client: NewClient(accessKey, SetClientProxy("http://127.0.0.1:7890"))}
	res, response, err := collectionGroup.RemovePhoto(token, RemovePhotoReq{
		CollectionId: "BhMKvHpksrM",
		PhotoId:      "a-red-race-car-sitting-on-top-of-a-dirt-field-YFcEYRE9lvc",
	})
	if err != nil {
		t.Fatal(err)
	}
	data, _ := json.Marshal(res)
	t.Log(string(data))
	t.Log(response.OriginResponseBody)
	t.Logf(response.requestURL)
}
