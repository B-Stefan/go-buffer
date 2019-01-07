package api

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestDecodeUser(t *testing.T) {

	jsonStr := `{
		"activity_at":1343654640,
  		"created_at":1326189062,
  		"id":"4f0c0a06512f7ef214000000",
  		"plan":"free",
  		"timezone":"Asia\/Tel_Aviv"
	}`

	var user User
	err := json.Unmarshal([]byte(jsonStr), &user)
	byteOut, err := json.Marshal(&user)

	eq, err := isJSONEqual(byteOut, []byte(jsonStr))

	fmt.Println("a=c\t", eq, "with error", err)

	if err != nil {
		fmt.Printf("%+v\n", user)
		t.Fatal(err)
	}

	if !eq {
		fmt.Printf("%+v\n", user)
		fmt.Printf("%+v\n", byteOut)
		t.Error("Json documents are not equal")
	}
}

func TestGetUser(t *testing.T) {

	userMock := User{
		Id:   "Bob",
		Plan: "free",
	}
	clientMock := ClientMock{}
	clientMock.On("do", mock.Anything).Return(userMock)
	clientMock.On("newRequest", "GET", mock.Anything, mock.Anything).Return()

	service := UserService{&clientMock}
	user, err := service.GetUser("<uuid>")

	assert.NoErrorf(t, err, "Should not throw error")
	assert.Equal(t, userMock, user)
	clientMock.AssertExpectations(t)
}

func TestDeauthorizeUser(t *testing.T) {

	clientMock := ClientMock{}
	clientMock.On("do", mock.Anything).Return(DeauthorizeResponse{Success: true})
	clientMock.On("newRequest", "POST", mock.Anything, mock.Anything).Return()

	service := UserService{&clientMock}
	user, err := service.DeauthorizeUser("<uuid>")

	assert.NoErrorf(t, err, "Should not throw error")
	assert.Equal(t, true, user)
	clientMock.AssertExpectations(t)
}
