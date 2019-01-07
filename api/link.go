package api

import "net/url"

type LinkShares struct {
	Shares int64 `json:"shares"`
}

type LinkService struct {
	client Client
}

func (s *LinkService) GetLinkShares(linkUrl url.URL) (LinkShares, error) {

	parameters := url.Values{}
	parameters.Add("url", linkUrl.String())

	req, err := s.client.newRequest("GET", "/links/shares.json?"+parameters.Encode(), nil)
	if err != nil {
		return LinkShares{}, err
	}
	var res LinkShares
	_, err = s.client.do(req, &res)
	return res, err
}
