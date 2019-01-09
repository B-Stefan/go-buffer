package api

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestDecodeUpdate(t *testing.T) {

	jsonStr := `{
  		"id": "4eb8565e0acb04bb82000004",
  		"created_at": 1320703582,
  		"day": "Monday 7th November",
  		"due_at": 1320742680,
  		"due_time": "10:09 pm",
  		"profile_id": "4eb854340acb04e870000010",
  		"profile_service": "twitter",
  		"sent_at": 1320744001,
  		"service_update_id": "133667319959392256",
  		"statistics": {
    		"reach": 2460,
    		"clicks": 56,
    		"retweets": 20,
    		"favorites": 1,
    		"mentions": 1
  		},
  		"status": "sent",
  		"text": "This is just the beginning, the very beginning...",
  		"text_formatted": "This is just the beginning, the very beginning...",
  		"user_id": "4eb9276e0acb04bb81000067",
  		"via": "chrome"
	}`

	model := new(Update)
	err := tryJsonEncoding(jsonStr, model)
	if err != nil {
		t.Error(err)
	}
}

func TestGetUpdate(t *testing.T) {
	mockUpdate := Update{
		Status: "sent",
	}

	clientMock := ClientMock{}
	clientMock.On("do", mock.Anything).Return(mockUpdate)
	clientMock.On("newRequest", "GET", mock.Anything, mock.Anything).Return()

	service := UpdateService{&clientMock}
	user, err := service.GetUpdate("uuid")

	assert.NoErrorf(t, err, "Should not throw error")
	assert.Equal(t, mockUpdate, user)
	clientMock.AssertExpectations(t)
}

func TestGetPendingUpdates(t *testing.T) {
	mockResponse := CountUpdateResponse{
		Count: 30,
	}

	clientMock := ClientMock{}
	clientMock.On("do", mock.Anything).Return(mockResponse)
	clientMock.On("newRequest", "GET", "/profiles/uuid/updates/pending.json?count=1", mock.Anything).Return()

	service := UpdateService{&clientMock}
	user, err := service.GetPendingUpdates("uuid", PendingUpdateOptions{
		Count: 1,
	})

	assert.NoErrorf(t, err, "Should not throw error")
	assert.Equal(t, mockResponse, user)
	clientMock.AssertExpectations(t)
}

func TestGetSendUpdates(t *testing.T) {
	mockResponse := CountUpdateResponse{
		Count: 303,
	}

	clientMock := ClientMock{}
	clientMock.On("do", mock.Anything).Return(mockResponse)
	clientMock.On("newRequest", "GET", "/profiles/uuid/updates/send.json?filter=default", mock.Anything).Return()

	service := UpdateService{&clientMock}
	user, err := service.GetSendUpdates("uuid", SendUpdateOptions{
		Filter: "default",
	})

	assert.NoErrorf(t, err, "Should not throw error")
	assert.Equal(t, mockResponse, user)
	clientMock.AssertExpectations(t)
}

func TestReorderUpdates(t *testing.T) {
	mockResponse := SuccessUpdateResponse{
		Success: true,
	}

	options := ReorderUpdatesOptions{
		Offset: 1,
		Order:  1,
	}

	clientMock := ClientMock{}
	clientMock.On("do", mock.Anything).Return(mockResponse)
	clientMock.On("newRequest", "POST", "/profiles/uuid/updates/reorder.json?offset=1&order=1&utc=false", mock.Anything).Return()

	service := UpdateService{&clientMock}
	res, err := service.ReorderUpdate("uuid", options)

	assert.NoErrorf(t, err, "Should not throw error")
	assert.Equal(t, mockResponse, res)
	clientMock.AssertExpectations(t)
}

func TestShuffleUpdates(t *testing.T) {
	mockResponse := SuccessUpdateResponse{
		Success: true,
	}

	options := ShuffleUpdatesOptions{
		Count: 1,
		Utc:   true,
	}

	clientMock := ClientMock{}
	clientMock.On("do", mock.Anything).Return(mockResponse)
	clientMock.On("newRequest", "POST", "/profiles/uuid/updates/shuffle.json?count=1&utc=true", mock.Anything).Return()

	service := UpdateService{&clientMock}
	res, err := service.ShuffleUpdate("uuid", options)

	assert.NoErrorf(t, err, "Should not throw error")
	assert.Equal(t, mockResponse, res)
	clientMock.AssertExpectations(t)
}

func TestCreateUpdate(t *testing.T) {
	mockResponse := SuccessUpdateResponseWithCount{
		Success:     true,
		BufferCount: 10,
	}

	newUpdate := CreateUpdateOptions{
		ProfileIds: []string{"uudid1", "uuid2"},
		Text:       "Hello from go-buffer",
	}
	clientMock := ClientMock{}
	clientMock.On("do", mock.Anything).Return(mockResponse)
	clientMock.On("newRequest", "POST", "/updates/create.json", newUpdate).Return()

	service := UpdateService{&clientMock}
	res, err := service.CreateUpdate(newUpdate)

	assert.NoErrorf(t, err, "Should not throw error")
	assert.Equal(t, mockResponse, res)
	clientMock.AssertExpectations(t)
}

func TestUpdateUpdate(t *testing.T) {
	mockResponse := SuccessUpdateResponseWithCount{
		Success:     true,
		BufferCount: 10,
	}

	newUpdate := UpdateUpdateOptions{
		Text: "Hello from go-buffer",
	}
	clientMock := ClientMock{}
	clientMock.On("do", mock.Anything).Return(mockResponse)
	clientMock.On("newRequest", "POST", "/updates/update.json", newUpdate).Return()

	service := UpdateService{&clientMock}
	res, err := service.UpdateUpdate(newUpdate)

	assert.NoErrorf(t, err, "Should not throw error")
	assert.Equal(t, mockResponse, res)
	clientMock.AssertExpectations(t)
}

func TestShareUpdate(t *testing.T) {
	mockResponse := SuccessResponse{
		Success: true,
	}

	clientMock := ClientMock{}
	clientMock.On("do", mock.Anything).Return(mockResponse)
	clientMock.On("newRequest", "POST", "/updates/uuid/share.json", nil).Return()

	service := UpdateService{&clientMock}
	res, err := service.ShareUpdate("uuid")

	assert.NoErrorf(t, err, "Should not throw error")
	assert.Equal(t, mockResponse.Success, res)
	clientMock.AssertExpectations(t)
}

func TestDestroyUpdate(t *testing.T) {
	mockResponse := SuccessResponse{
		Success: true,
	}

	clientMock := ClientMock{}
	clientMock.On("do", mock.Anything).Return(mockResponse)
	clientMock.On("newRequest", "POST", "/updates/uuid/destroy.json", nil).Return()

	service := UpdateService{&clientMock}
	res, err := service.DestroyUpdate("uuid")

	assert.NoErrorf(t, err, "Should not throw error")
	assert.Equal(t, mockResponse.Success, res)
	clientMock.AssertExpectations(t)
}

func TestMoveToTopUpdate(t *testing.T) {
	mockResponse := Update{
		Text: "Greetings from go-buffer",
	}

	clientMock := ClientMock{}
	clientMock.On("do", mock.Anything).Return(mockResponse)
	clientMock.On("newRequest", "POST", "/updates/uuid/move_to_top.json", nil).Return()

	service := UpdateService{&clientMock}
	res, err := service.MoveToTopUpdate("uuid")

	assert.NoErrorf(t, err, "Should not throw error")
	assert.Equal(t, mockResponse, res)
	clientMock.AssertExpectations(t)
}
