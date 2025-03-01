package response

const (
	ErrorCodeSuccess  = 20000
	ErrorParamInvalid = 40000

	ErrorInvalidToken = 30001
)

var msg = map[int]string{
	ErrorCodeSuccess:  "Success",
	ErrorParamInvalid: "Email is invalid",
}
