package api

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/url"
	"testing"
)

func TestDecodeLinkShare(t *testing.T) {

	jsonStr := `{
    	"shares": 47348
	}`

	shares := new(LinkShares)
	err := tryJsonEncoding(jsonStr, shares)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetLinkShare(t *testing.T) {

	linkUrl, err := url.Parse("http://my.cool.url.de/äöü@!_/test")
	linkUrlEncoded := "http%3A%2F%2Fmy.cool.url.de%2F%25C3%25A4%25C3%25B6%25C3%25BC%40%2521_%2Ftest"
	apiUrl := "/links/shares.json?url="

	linkMock := LinkShares{
		Shares: 2304,
	}

	clientMock := ClientMock{}
	clientMock.On("do", mock.Anything).Return(linkMock)
	clientMock.On("newRequest", "GET", apiUrl+linkUrlEncoded, mock.Anything).Return()

	service := LinkService{&clientMock}
	user, err := service.GetLinkShares(*linkUrl)

	assert.NoErrorf(t, err, "Should not throw error")
	assert.Equal(t, linkMock, user)
	clientMock.AssertExpectations(t)
}
