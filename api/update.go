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
	Page  int   `form:"page" json:"page"`
	Count int   `form:"count" json:"count"`
	Since int   `form:"since" json:"since"`
	Utc   int64 `form:"utc" json:"utc"`
}
type SendUpdateOptions struct {
	Page   int    `form:"page" json:"page"`
	Count  int    `form:"count" json:"count"`
	Since  int    `form:"since" json:"since"`
	Utc    int64  `form:"utc" json:"utc"`
	Filter string `form:"filter" json:"filter"`
}

type SuccessUpdateResponse struct {
	Success bool     `json:"success"`
	Updates []Update `json:"updates"`
}
type SuccessUpdateResponseWithCount struct {
	Success          bool     `json:"success"`
	BufferCount      int      `json:"buffer_count"`
	BufferPercentage float32  `json:"buffer_percentage"`
	Updates          []Update `json:"updates"`
}

type ReorderUpdatesOptions struct {
	Order  int  `form:"order" json:"order"`
	Offset int  `form:"offset" json:"offset"`
	Utc    bool `form:"utc" json:"utc"`
}

type ShuffleUpdatesOptions struct {
	Count int  `form:"count" json:"count"`
	Utc   bool `form:"utc" json:"utc"`
}

type CreateUpdateOptions struct {
	ProfileIds  []string      `form:"profile_ids" json:"profile_ids"`
	Text        string        `form:"text,omitempty" json:"text,omitempty"`
	Shorten     bool          `form:"shorten,omitempty" json:"shorten,omitempty"`
	Now         bool          `form:"now,omitempty" json:"now,omitempty"`
	Top         bool          `form:"top,omitempty" json:"top,omitempty"`
	Media       []interface{} `form:"media,omitempty" json:"media,omitempty"`
	Attachment  bool          `form:"attachment,omitempty" json:"attachment,omitempty"`
	ScheduledAt bool          `form:"scheduled_at,omitempty" json:"scheduled_at,omitempty"`
	Retweet     struct {
		TweetId string `form:"tweet_id" json:"tweet_id"`
		Comment string `form:"tweet_id,omitempty" json:"tweet_id,omitempty"`
	} `form:"retweet,omitempty" json:"retweet,omitempty"`
}

type UpdateUpdateOptions struct {
	Text        string        `form:"text" json:"text"`
	Media       []interface{} `form:"media" json:"media"`
	Now         bool          `form:"now" json:"now"`
	Utc         bool          `form:"utc" json:"utc"`
	ScheduledAt bool          `form:"scheduled_at" json:"scheduled_at"`
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

func (s *UpdateService) ReorderUpdate(profileId string, options ReorderUpdatesOptions) (SuccessUpdateResponse, error) {

	parameters, err := getValuesWithoutEmpty(options)

	if err != nil {
		return SuccessUpdateResponse{}, err
	}

	req, err := s.client.newRequest("POST", "/profiles/"+profileId+"/updates/reorder.json?"+parameters.Encode(), nil)

	if err != nil {
		return SuccessUpdateResponse{}, err
	}
	var res SuccessUpdateResponse
	_, err = s.client.do(req, &res)
	return res, err
}

func (s *UpdateService) ShuffleUpdate(profileId string, options ShuffleUpdatesOptions) (SuccessUpdateResponse, error) {

	parameters, err := getValuesWithoutEmpty(options)

	if err != nil {
		return SuccessUpdateResponse{}, err
	}

	req, err := s.client.newRequest("POST", "/profiles/"+profileId+"/updates/shuffle.json?"+parameters.Encode(), nil)

	if err != nil {
		return SuccessUpdateResponse{}, err
	}
	var res SuccessUpdateResponse
	_, err = s.client.do(req, &res)
	return res, err
}

func (s *UpdateService) CreateUpdate(newUpdate CreateUpdateOptions) (SuccessUpdateResponseWithCount, error) {

	req, err := s.client.newRequest("POST", "/updates/create.json", newUpdate)

	if err != nil {
		return SuccessUpdateResponseWithCount{}, err
	}
	var res SuccessUpdateResponseWithCount
	_, err = s.client.do(req, &res)
	return res, err
}

func (s *UpdateService) UpdateUpdate(newUpdate UpdateUpdateOptions) (SuccessUpdateResponseWithCount, error) {

	req, err := s.client.newRequest("POST", "/updates/update.json", newUpdate)

	if err != nil {
		return SuccessUpdateResponseWithCount{}, err
	}
	var res SuccessUpdateResponseWithCount
	_, err = s.client.do(req, &res)
	return res, err
}

func (s *UpdateService) ShareUpdate(updateId string) (bool, error) {

	req, err := s.client.newRequest("POST", "/updates/"+updateId+"/share.json", nil)

	if err != nil {
		return false, err
	}
	var res SuccessResponse
	_, err = s.client.do(req, &res)
	return res.Success, err
}

func (s *UpdateService) DestroyUpdate(updateId string) (bool, error) {

	req, err := s.client.newRequest("POST", "/updates/"+updateId+"/destroy.json", nil)

	if err != nil {
		return false, err
	}
	var res SuccessResponse
	_, err = s.client.do(req, &res)
	return res.Success, err
}
