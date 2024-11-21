package dto

type UserCreateRequest struct {
	Username  string `validate:"required,min=1,max=100" json:"username"`
	Firstname string `validate:"required,min=1,max=100" json:"first_name"`
	Lastname  string `validate:"required,min=1,max=100" json:"last_name"`
	Email     string `validate:"required,min=1,max=100" json:"email"`
	Password  string `validate:"required,min=8" json:"password"`
}

type UserEditRequest struct {
	Firstname string `validate:"required,min=1,max=100" json:"first_name"`
	Lastname  string `validate:"required,min=1,max=100" json:"last_name"`
}

type UserChangePasswordRequest struct {
	Oldpassword string `validate:"required,min=8" json:old_password`
	Newpassword string `validate:"required,min=8" json:new_password`
}

type UserCreateOauthRequest struct {
	Email      string `validate:"required,min=1,max=100" json:"email"`
	FamilyName string `validate:"required,min=1,max=100" json:"family_name"`
	GivenName  string `validate:"required,min=1,max=100" json:"given_name"`
	Sub        string `validate:"required,min=1,max=100" json:"sub"`
}
