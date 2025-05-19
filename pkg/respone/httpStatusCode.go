package respone

const (
	ErrCodeSucces       = 20001 //Succes
	ErrCodeParamInvalid = 20003 //Email Invalid
)

var msg = map[int]string{
	ErrCodeSucces:       "success",
	ErrCodeParamInvalid: "Email is invalid",
}
