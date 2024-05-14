/*
 * Copyright (c) 2024. Lorem ipsum dolor sit amet, consectetur adipiscing elit.
 * Morbi non lorem porttitor neque feugiat blandit. Ut vitae ipsum eget quam lacinia accumsan.
 * Etiam sed turpis ac ipsum condimentum fringilla. Maecenas magna.
 * Proin dapibus sapien vel ante. Aliquam erat volutpat. Pellentesque sagittis ligula eget metus.
 * Vestibulum commodo. Ut rhoncus gravida arcu.
 */

package main

import (
	"encoding/json"
	"testing"
)

func TestPhoto_Photos(t *testing.T) {
	photoGroup := Photo{Client: NewClient(accessKey, SetClientProxy("http://127.0.0.1:7890"))}
	res, response, err := photoGroup.Photos(PhotosReq{
		Page:    "1",
		PerPage: "1",
		OrderBy: "",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response.requestURL)
	t.Log(response.OriginResponseBody)
	a, _ := json.Marshal(res)
	t.Logf(string(a))
}

func TestPhoto_GetPhoto(t *testing.T) {
	photoGroup := Photo{Client: NewClient(accessKey, SetClientProxy("http://127.0.0.1:7890"))}
	res, response, err := photoGroup.GetAPhoto("Nx2k4q0dcJY")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response.requestURL)
	//t.Log(response.OriginResponseBody)
	a, _ := json.Marshal(res)
	t.Logf(string(a))
}

func TestPhoto_Random(t *testing.T) {
	photoGroup := Photo{Client: NewClient(accessKey, SetClientProxy("http://127.0.0.1:7890"))}
	res, response, err := photoGroup.Random(PhotoRandomReq{})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response.requestURL)
	t.Log(response.OriginResponseBody)
	a, _ := json.Marshal(res)
	t.Logf(string(a))
}
