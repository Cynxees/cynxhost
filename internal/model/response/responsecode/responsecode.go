package responsecode

type ResponseCode string

const (
	CodeSuccess             ResponseCode = "SU"
	CodeValidationError     ResponseCode = "VE"
	CodeAuthenticationError ResponseCode = "AU"
	CodeServerError         ResponseCode = "SE"
	CodeInternalError       ResponseCode = "IE"
	CodeNotAllowed          ResponseCode = "NA"
	CodeNotFound            ResponseCode = "NF"

	CodeTblServerTemplateError ResponseCode = "ST"
	CodeTblUserError           ResponseCode = "TU"
)

var ResponseCodeNames = map[ResponseCode]string{
	CodeSuccess:             "Success",
	CodeValidationError:     "Validation Error",
	CodeAuthenticationError: "Authentication Error",
	CodeServerError:         "Server Error",
	CodeInternalError:       "Internal Error",
	CodeNotAllowed:          "Not Allowed",
	CodeNotFound:            "Not Found",

	CodeTblServerTemplateError: "TblServerTemplate Error",
	CodeTblUserError:           "TblUser Error",
}
