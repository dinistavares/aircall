package aircall

import (
	"os"
	"reflect"
	"testing"
)

var (
	testAppID     = os.Getenv("AIRCALL_APP_ID")
	testAppSecret = os.Getenv("AIRCALL_APP_SECRET")
	testClient    = NewClient(testAppID, testAppSecret)

	expectedCompanyResponse = CompanyResponse{
		Company{
			Name:         "TestCompany",
			UsersCount:   1,
			NumbersCount: 1,
		},
	}

	expectedUsersResponse = UsersResponse{
		Meta: ResponseMeta{
			Count:            1,
			Total:            1,
			CurrentPage:      1,
			PerPage:          20,
			NextPageLink:     "",
			PreviousPageLink: "",
		},
		Users: []User{
			{
				ID:         147262,
				DirectLink: "https://api.aircall.io/v1/users/147262",
				Name:       "Test User",
				Email:      "leo.hetsch+1@gmail.com",
				Available:  true,
			},
		},
	}
)

func TestApiPing(t *testing.T) {
	res, err := testClient.Ping()
	if err != nil {
		t.Fatalf("Cannot call /ping endpoint [err: %v]", err)
	}
	if res.Ping != "pong" {
		t.Error("Expected pong response")
	}
}

func TestCompany(t *testing.T) {
	res, err := testClient.Company()
	if err != nil {
		t.Fatalf("Cannot call /company endpoint [err: %v]", err)
	}

	if !reflect.DeepEqual(res, expectedCompanyResponse) {
		t.Error("Invalid Company response")
	}
}
