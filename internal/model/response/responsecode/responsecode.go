package responsecode

type ResponseCode string

const (

	// Expected Error
	CodeSuccess             ResponseCode = "SU"
	CodeValidationError     ResponseCode = "VE"
	CodeAuthenticationError ResponseCode = "AU"
	CodeNotAllowed          ResponseCode = "NA"
	CodeNotFound            ResponseCode = "NF"
	CodeInvalidCredentials  ResponseCode = "IC"
	CodeForbidden           ResponseCode = "FB"
	CodeFailJSON            ResponseCode = "FJ"

	// Internal
	CodeJwtError        ResponseCode = "JWTERR"
	CodeInternalError   ResponseCode = "IE"
	CodeAWSError        ResponseCode = "AWSERR"
	CodeEC2Error        ResponseCode = "EC2ERR"
	CodeECRError        ResponseCode = "ECRERR"
	CodeS3Error         ResponseCode = "S3ERR"
	CodeRCONError       ResponseCode = "RCONERR"
	CodePorkbunError    ResponseCode = "PBERR"
	CodeCloudflareError ResponseCode = "CFERR"

	// DB Error
	CodeTblServerTemplateError         ResponseCode = "TBLSTE"
	CodeTblServerTemplateCategoryError ResponseCode = "TBLSTC"
	CodeTblUserError                   ResponseCode = "TBLUSR"
	CodeTblInstanceError               ResponseCode = "TBLINT"
	CodeTblInstanceTypeError           ResponseCode = "TBLITT"
	CodeTblPersistentNodeError         ResponseCode = "TBLPND"
	CodeTblScriptError                 ResponseCode = "TBLSCP"
	CodeTblStorageError                ResponseCode = "TBLSTO"
)

var ResponseCodeNames = map[ResponseCode]string{
	CodeSuccess:             "Success",
	CodeValidationError:     "Validation Error",
	CodeAuthenticationError: "Authentication Error",
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
