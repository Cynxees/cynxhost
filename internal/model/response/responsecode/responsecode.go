package responsecode

type ResponseCode string

const (
	// Success Codes
	CodeSuccess ResponseCode = "SU"

	// Validation Error
	CodeValidationError ResponseCode = "VE"

	// Authentication Error
	CodeAuthenticationError ResponseCode = "AU"

	// Generic Server Error
	CodeServerError ResponseCode = "SE"

	// Internal Server Error
	CodeInternalError ResponseCode = "IE"
)

var ResponseCodeNames = map[ResponseCode]string{
	CodeSuccess:             "Success",
	CodeValidationError:     "Validation Error",
	CodeAuthenticationError: "Authentication Error",
	CodeServerError:         "Server Error",
	CodeInternalError:       "Internal Error",
}
