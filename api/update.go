package api

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
	Page  int   `form:"page",json:"page"`
	Count int   `form:"count",json:"count"`
	Since int   `form:"since",json:"since"`
	Utc   int64 `form:"utc",json:"utc"`
}
type SendUpdateOptions struct {
	Page   int    `form:"page",json:"page"`
	Count  int    `form:"count",json:"count"`
	Since  int    `form:"since",json:"since"`
	Utc    int64  `form:"utc",json:"utc"`
	Filter string `form:"filter",json:"filter"`
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

	parameters, err := getValuesWithoutEmpty(options)

	if err != nil {
		return CountUpdateResponse{}, err
	}

	req, err := s.client.newRequest("GET", "/profiles/"+profileId+"/updates/pending.json?"+parameters.Encode(), nil)
	if err != nil {
		return CountUpdateResponse{}, err
	}
	var res CountUpdateResponse
	_, err = s.client.do(req, &res)
	return res, err
}

func (s *UpdateService) GetSendUpdates(profileId string, options SendUpdateOptions) (CountUpdateResponse, error) {

	parameters, err := getValuesWithoutEmpty(options)

	if err != nil {
		return CountUpdateResponse{}, err
	}

	req, err := s.client.newRequest("GET", "/profiles/"+profileId+"/updates/send.json?"+parameters.Encode(), nil)
	if err != nil {
		return CountUpdateResponse{}, err
	}
	var res CountUpdateResponse
	_, err = s.client.do(req, &res)
	return res, err
}
