package api

import (
	"net/url"
	"strconv"
)

type Update struct {
	Id              string           `json:"id"`
	CreatedAt       int64            `json:"created_at"`
	Day             string           `json:"day"`
	DueAt           int64            `json:"due_at"`
	DueTime         string           `json:"due_time"`
	ProfileId       string           `json:"profile_id"`
	ProfileService  string           `json:"profile_service"`
	SentAt          int64            `json:"sent_at"`
	ServiceUpdateId string           `json:"service_update_id"`
	Statistics      map[string]int64 `json:"statistics"`
	Status          string           `json:"status"`
	Text            string           `json:"text"`
	TextFormatted   string           `json:"text_formatted"`
	UserId          string           `json:"user_id"`
	Via             string           `json:"via"`
}

type CountUpdateResponse struct {
	Count   int      `json:"count"`
	Updates []Update `json:"updates"`
}
type PendingUpdateOptions struct {
	Page  int   `json:"page"`
	Count int   `json:"count"`
	Since int   `json:"since"`
	Utc   int64 `json:"utc"`
}
type SendUpdateOptions struct {
	PendingUpdateOptions
	Filter string `json:"filter"`
}
type UpdateService struct {
	client Client
}

func (s *UpdateService) GetUpdate(id string) (Update, error) {

	req, err := s.client.newRequest("GET", "/updates/"+id+".json", nil)
	if err != nil {
		return Update{}, err
	}
	var res Update
	_, err = s.client.do(req, &res)
	return res, err
}

func (s *UpdateService) GetPendingUpdates(profileId string, options PendingUpdateOptions) (CountUpdateResponse, error) {

	parameters := s.getValues(options)

	req, err := s.client.newRequest("GET", "/profiles/"+profileId+"/updates/pending.json?"+parameters.Encode(), nil)
	if err != nil {
		return CountUpdateResponse{}, err
	}
	var res CountUpdateResponse
	_, err = s.client.do(req, &res)
	return res, err
}

func (s *UpdateService) GetSendUpdates(profileId string, options SendUpdateOptions) (CountUpdateResponse, error) {

	parameters := url.Values{}

	// @TODO: Improve solution by adding custom encoder for values.
	// See: https://github.com/go-playground/form#registering-custom-types

	if options.Count > 0 {
		parameters.Add("count", strconv.Itoa(options.Count))
	}
	if options.Page > 0 {
		parameters.Add("page", strconv.Itoa(options.Page))
	}
	if options.Utc > 0 {
		parameters.Add("utc", strconv.FormatInt(options.Utc, 10))
	}
	if options.Since > 0 {
		parameters.Add("since", strconv.Itoa(options.Since))
	}
	if options.Filter != "" {
		parameters.Add("filter", options.Filter)
	}

	req, err := s.client.newRequest("GET", "/profiles/"+profileId+"/updates/send.json?"+parameters.Encode(), nil)
	if err != nil {
		return CountUpdateResponse{}, err
	}
	var res CountUpdateResponse
	_, err = s.client.do(req, &res)
	return res, err
}

func (s *UpdateService) getValues(options PendingUpdateOptions) url.Values {
	parameters := url.Values{}

	// @TODO: Improve solution by adding custom encoder for values.
	// See: https://github.com/go-playground/form#registering-custom-types

	if options.Count > 0 {
		parameters.Add("count", strconv.Itoa(options.Count))
	}
	if options.Page > 0 {
		parameters.Add("page", strconv.Itoa(options.Page))
	}
	if options.Utc > 0 {
		parameters.Add("utc", strconv.FormatInt(options.Utc, 10))
	}
	if options.Since > 0 {
		parameters.Add("since", strconv.Itoa(options.Since))
	}
	return parameters
}
