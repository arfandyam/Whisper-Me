package web

type UserCreateRequest struct {
	Username  string `validate:"required,min=1,max=100" json:"username"`
	Firstname string `validate:"required,min=1,max=100" json:"firstname"`
	Lastname  string `validate:"required,min=1,max=100" json:"lastname"`
	Email     string `validate:"required,min=1,max=100" json:"email"`
	Password  string `validate:"required,min=8,max=100" json:"password"`
}

// type UserChangePasswordRequest struct {
// 	Oldpassword string `validate:"required,min=8,max=100" json:oldpassword`
// 	Newpassword string `validate:"required,min=8,max=100" json:newpassword`
// }
