package respone

const (
	ErrCodeSucces       = 20001 //Succes
	ErrCodeParamInvalid = 20003 //Email Invalid
	ErrInvalidToken     = 30001 //Token invalid
)

var msg = map[int]string{
	ErrCodeSucces:       "success",
	ErrCodeParamInvalid: "Email is invalid",
	ErrInvalidToken:     "Token is invalid",
}
