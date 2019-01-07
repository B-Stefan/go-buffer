package api

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestDecodeProfile(t *testing.T) {

	profileJson := `{
	"avatar" : "http://a3.twimg.com/profile_images/1405180232.png",
		"created_at" :  1320703028,
		"default" : true,
		"formatted_username" : "@skinnyoteam",
		"id" : "4eb854340acb04e870000010",
		"schedules" : [{
		"days" : [
			"mon",
		"tue",
		"wed",
		"thu",
		"fri"
	],
		"times" : [
		"12:00",
		"17:00",
		"18:00"
	]
	}],
    "service" : "twitter",
    "service_id" : "164724445",
    "service_username" : "skinnyoteam",
    "statistics" : {
        "followers" : 246
    },
    "team_members" : [
        "4eb867340acb04e670000001"
    ],
    "timezone" : "Europe/London",
    "user_id" : "4eb854340acb04e870000010"
    }`

	var profile Profile
	err := json.Unmarshal([]byte(profileJson), &profile)
	byteOut, err := json.Marshal(&profile)

	eq, err := isJSONEqual(byteOut, []byte(profileJson))

	fmt.Println("a=c\t", eq, "with error", err)

	if err != nil {
		fmt.Printf("%+v\n", profile)
		t.Fatal(err)
	}

	if !eq {
		fmt.Printf("%+v\n", profile)
		fmt.Printf("%+v\n", byteOut)
		t.Error("Json documents are not equal")
	}
}

func TestListProfiles(t *testing.T) {

	profilesMock := []Profile{
		{FormattedUsername: "Yoda"},
		{FormattedUsername: "Luke"},
	}
	clientMock := ClientMock{}
	clientMock.On("do", mock.Anything).Return(profilesMock)
	clientMock.On("newRequest", "GET", mock.Anything, mock.Anything).Return()

	service := ProfileService{&clientMock}
	profiles, err := service.ListProfiles()

	assert.NoErrorf(t, err, "Should not throw error")
	assert.Equal(t, profilesMock, profiles)
	clientMock.AssertExpectations(t)
}

func TestGetProfile(t *testing.T) {

	profileMock := Profile{
		FormattedUsername: "Bob",
		ServiceUsername:   "facebook",
	}
	clientMock := ClientMock{}
	clientMock.On("do", mock.Anything).Return(profileMock)
	clientMock.On("newRequest", "GET", mock.Anything, mock.Anything).Return()

	service := ProfileService{&clientMock}
	profile, err := service.GetProfile("<uuid>")

	assert.NoErrorf(t, err, "Should not throw error")
	assert.Equal(t, profileMock, profile)
	clientMock.AssertExpectations(t)
}

func TestGetProfileSchedules(t *testing.T) {

	scheduleMock := []Schedule{
		{Days: []string{"Mon", "Thu"}},
		{Times: []string{"09:00", "15:00"}},
	}
	clientMock := ClientMock{}
	clientMock.On("do", mock.Anything).Return(scheduleMock)
	clientMock.On("newRequest", "GET", mock.Anything, mock.Anything).Return()

	service := ProfileService{&clientMock}
	schedules, err := service.GetProfileSchedules("<uuid>")

	assert.NoErrorf(t, err, "Should not throw error")
	assert.Equal(t, scheduleMock, schedules)
	clientMock.AssertExpectations(t)
}

func TestUpdateProfileSchedules(t *testing.T) {

	scheduleMock := Schedule{
		Days: []string{"Mon", "Thu"},
	}
	clientMock := new(ClientMock)
	clientMock.On("do", mock.Anything).Return(SuccessResponse{Success: true})
	clientMock.On("newRequest", "POST", mock.AnythingOfType("string"), mock.Anything).Return()

	service := ProfileService{clientMock}
	schedules, err := service.UpdateProfileSchedule("<uuid>", &scheduleMock)

	assert.NoErrorf(t, err, "Should not throw error")
	assert.Equal(t, true, schedules)
	clientMock.AssertExpectations(t)
}
