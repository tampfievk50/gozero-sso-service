package types

type AccountExtend struct {
	FirstName                  string `json:"firstName,omitempty,optional"`
	LastName                   string `json:"lastName,omitempty,optional"`
	Gender                     int8   `json:"gender,omitempty,optional"`
	Dob                        string `json:"dob,omitempty,optional"`
	Address                    string `json:"address,omitempty,optional"`
	IdentityCardNumber         string `json:"identityCardNumber,omitempty,optional"`
	IdentityCardIssuedDate     string `json:"identityCardIssuedDate,omitempty,optional"`
	IdentityCardIssuedBy       string `json:"identityCardIssuedBy,omitempty,optional"`
	IdentityCardFrontViewPhoto string `json:"identityCardFrontViewPhoto,omitempty,optional"`
	IdentityCardBackViewPhoto  string `json:"identityCardBackViewPhoto,omitempty,optional"`
}
type AccountCreateRequest struct {
	Email    string        `json:"email"`
	Password string        `json:"password"`
	Phone    string        `json:"phone,omitempty,optional"`
	Extend   AccountExtend `json:"extend,omitempty,optional"`
	RoleIds  []int64       `json:"roleIds,omitempty,optional"`
}

type AccountUpdateRequest struct {
	Id          int64         `path:"id"`
	NewPassword string        `json:"newPassword,omitempty,optional"`
	Phone       string        `json:"phone,omitempty,optional"`
	Extend      AccountExtend `json:"extend,omitempty,optional"`
	RoleIds     []int64       `json:"roleIds,omitempty,optional"`
	GroupIds    []int64       `json:"groupIds,omitempty,optional"`
}

type AccountUpdateStateRequest struct {
	Id     int64 `path:"id"`
	Active bool  `json:"active"`
}

type AuthForgetPasswordRequest struct {
	Email   string `json:"email"`
	Token   string `json:"token,optional"`
	CiId    string `json:"ciId,optional"`
	CiToken string `json:"ciToken,optional"`
}

type AuthForgetPasswordResponse struct {
	Ttl int32 `json:"ttl"`
}

type AuthIsAllowRequest struct {
	Token  string `json:"token"`
	Method string `json:"method"`
	Path   string `json:"path"`
}

type AuthLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token,optional"`
	Otp      string `json:"otp,optional"`
	CiId     string `json:"ciId,optional"`
	CiToken  string `json:"ciToken,optional"`
}

type AuthLoginWithSecretRequest struct {
	Secret     string `json:"secret,optional"`
	AccessCode string `json:"accessCode,optional"`
	Otp        string `json:"otp,optional"`
}

type AuthResendOtpWithSecretRequest struct {
	Secret string `json:"secret"`
}

type AuthResetPasswordRequest struct {
	Email    string `json:"email"`
	Otp      string `json:"otp"`
	Password string `json:"password"`
}
