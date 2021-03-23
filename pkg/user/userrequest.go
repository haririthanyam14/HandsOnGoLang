package user

type UserRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type UserResp struct {
	UserID  int
	Message string
}
