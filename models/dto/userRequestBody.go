package dto

type UserCreateRequest struct {
	Username  string `validate:"required,min=1,max=100" json:"username"`
	Firstname string `validate:"required,min=1,max=100" json:"firstname"`
	Lastname  string `validate:"required,min=1,max=100" json:"lastname"`
	Email     string `validate:"required,min=1,max=100" json:"email"`
	Password  string `validate:"required,min=8" json:"password"`
}

type UserEditRequest struct {
	Firstname string `validate:"required,min=1,max=100" json:"firstname"`
	Lastname  string `validate:"required,min=1,max=100" json:"lastname"`
}

type UserChangePasswordRequest struct {
	Oldpassword string `validate:"required,min=8" json:oldpassword`
	Newpassword string `validate:"required,min=8" json:newpassword`
}

type UserCreateOauthRequest struct {
	Email      string `validate:"required,min=1,max=100" json:"email"`
	FamilyName string `validate:"required,min=1,max=100" json:"family_name"`
	GivenName  string `validate:"required,min=1,max=100" json:"given_name"`
	Sub        string `validate:"required,min=1,max=100" json:"sub"`
}
