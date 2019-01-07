package api

import "encoding/json"

type Schedule struct {
	Days  []string `json:"days",from:"days"`
	Times []string `json:"times",from:"times"`
}
type Profile struct {
	Avatar            string                      `json:"avatar,omitempty"`
	FormattedUsername string                      `json:"formatted_username,omitempty"`
	CreatedAt         int64                       `json:"created_at,omitempty"`
	Default           bool                        `json:"default,omitempty"`
	Id                string                      `json:"id,omitempty"`
	Schedules         []Schedule                  `json:"schedules,omitempty"`
	Service           string                      `json:"service"`
	ServiceId         string                      `json:"service_id"`
	ServiceUsername   string                      `json:"service_username"`
	Statistics        map[string]*json.RawMessage `json:"statistics,omitempty"`
	TeamMembers       []string                    `json:"team_members,omitempty"`
	Timezone          string                      `json:"timezone,omitempty"`
	UserId            string                      `json:"user_id,omitempty"`
}

type SuccessResponse struct {
	Success bool `json:"success"`
}
type ProfileService struct {
	client Client
}

func (s *ProfileService) ListProfiles() ([]Profile, error) {
	req, err := s.client.newRequest("GET", "/profiles.json", nil)
	if err != nil {
		return nil, err
	}
	var users []Profile
	_, err = s.client.do(req, &users)
	return users, err
}

func (s *ProfileService) GetProfile(id string) (Profile, error) {
	req, err := s.client.newRequest("GET", "/profiles/"+id+".json", nil)
	if err != nil {
		return Profile{}, err
	}
	var user Profile
	_, err = s.client.do(req, &user)
	return user, err
}

func (s *ProfileService) GetProfileSchedules(id string) ([]Schedule, error) {
	req, err := s.client.newRequest("GET", "/profiles/"+id+"/schedules.json", nil)
	if err != nil {
		return []Schedule{}, err
	}
	var schedules []Schedule
	_, err = s.client.do(req, &schedules)
	return schedules, err
}

func (s *ProfileService) UpdateProfileSchedule(id string, schedule *Schedule) (bool, error) {
	req, err := s.client.newRequest("POST", "/profiles/"+id+"/schedules/update.json", schedule)
	if err != nil {
		return false, err
	}
	var res SuccessResponse
	_, err = s.client.do(req, &res)
	return res.Success, err
}
