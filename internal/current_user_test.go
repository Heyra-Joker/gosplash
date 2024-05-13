package internal

import (
	"testing"
)

const accessKey = "44POH-LuAn-Knh-JrASbP2yGcDDKGDArtAAJhk9x_Lk"

func TestCurrentUser_Me(t *testing.T) {
	currentUserGroup := CurrentUser{Client: NewClient(accessKey, SetClientProxy("http://127.0.0.1:7890"))}
	reply, responseHeader, err := currentUserGroup.Me("ntMHq_GwOOdmkwVrk9oxch7nnaKZ29XrNSne56-dPjY")
	t.Logf("%#v", responseHeader)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", reply)
}

func TestCurrentUser_UpdateMe(t *testing.T) {
	currentUserGroup := CurrentUser{Client: NewClient(accessKey, SetClientProxy("http://127.0.0.1:7890"))}
	req := UpdateMeRequest{InstagramUsername: "Joker", FirstName: "thanks", Email: "77581236@gmail.com"}
	reply, _, err := currentUserGroup.UpdateMe("Kb7QKq_NAHxX5YiV713fKzWd3nFDMzKpQIs7hAYx47w", req)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", reply)
}
