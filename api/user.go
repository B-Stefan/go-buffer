package api

type User struct {
	Id         string `json:"id"`
	ActivityAt int64  `json:"activity_at"`
	CreatedAt  int64  `json:"created_at"`
	Plan       string `json:"plan"`
	Timezone   string `json:"timezone"`
}

type DeauthorizeResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type UserService struct {
	client Client
}

func (s *UserService) GetUser(id string) (User, error) {
	req, err := s.client.newRequest("GET", "/user.json", nil)
	if err != nil {
		return User{}, err
	}
	var user User
	_, err = s.client.do(req, &user)
	return user, err
}

func (s *UserService) DeauthorizeUser(id string) (bool, error) {
	req, err := s.client.newRequest("POST", "/user/deauthorize.json", nil)
	if err != nil {
		return false, err
	}
	var res DeauthorizeResponse
	_, err = s.client.do(req, &res)
	return res.Success, err
}
