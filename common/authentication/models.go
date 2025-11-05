package authentication

const (
	UnauthorisedErrorMessage = "USER IS NOT AUTHORISED TO PERFORM THE ACTION. PLEASE USE PROPER AUTHENTICATION TOKEN"
)

type UserDetail struct {
	UserId  string
	EmailId string
}
