package validator

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

//? A Custom struct validation
func ValidateDataRequest(dataReq interface{}) (string, string) {
	//* Validate data request
	validate := validator.New()
	err := validate.Struct(dataReq)

	var errMsg string
	var errMsgInd string

	if err != nil {
		if strings.Contains(err.Error(), "failed on the 'required' tag") {
			field := between(err.Error(), "Field validation for ", " failed on the 'required")
			errMsg = fmt.Sprintf("Data %s cannot be empty", field)
			errMsgInd = fmt.Sprintf("Data %s tidak boleh kosong", field)
		} else if strings.Contains(err.Error(), "failed on the 'email' tag") {
			errMsg = "Data Email not valid"
			errMsgInd = "Data Email tidak valid"
		} else if strings.Contains(err.Error(), "failed on the 'confirmPassword' tag") {
			errMsg = "Password not equal"
			errMsgInd = "Password tidak sama"
		} else if strings.Contains(err.Error(), "failed on the 'minPassword' tag") {
			errMsg = "Password cannot be less than 8 characters"
			errMsgInd = "Kata sandi tidak boleh kurang dari 8 karakter"
		} else {
			errMsg = err.Error()
			errMsgInd = err.Error()
		}
		return errMsg, errMsgInd
	}
	return "", ""
}

func between(value string, a string, b string) string {
	// Get substring between two strings.
	posFirst := strings.Index(value, a)
	if posFirst == -1 {
		return ""
	}
	posLast := strings.Index(value, b)
	if posLast == -1 {
		return ""
	}
	posFirstAdjusted := posFirst + len(a)
	if posFirstAdjusted >= posLast {
		return ""
	}
	return value[posFirstAdjusted:posLast]
}
