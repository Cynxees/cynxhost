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

	CodeTblServerTemplateError ResponseCode = "TBLSTE"
	CodeTblUserError           ResponseCode = "TBLUSR"
	CodeTblInstanceError       ResponseCode = "TBLINT"
	CodeTblInstanceTypeError   ResponseCode = "TBLITT"
	CodeTblPersistentNodeError ResponseCode = "TBLPND"
	CodeTblScriptError         ResponseCode = "TBLSCP"
	CodeTblStorageError        ResponseCode = "TBLSTO"
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
	CodeTblInstanceError:       "TblInstance Error",
	CodeTblInstanceTypeError:   "TblInstanceType Error",
	CodeTblPersistentNodeError: "TblPersistentNode Error",
	CodeTblScriptError:         "TblScript Error",
	CodeTblStorageError:        "TblStorage Error",
}
