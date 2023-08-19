package constants

const (
	ErrorStatus        = "Error"
	SuccessStatus      = "Success"
	ErrorMessageSucces = ""
	SuccessInsertData  = "Success insert data"
	SuccessUpdateDate  = "Success update data"
	SuccessGetData     = "success retrieve data"
	SuccessLogin   = "success login"
	SuccessRfreshToken   = "success refresh token"
	ErrorIdNotValid    = "ID not valid"
	ErrorDataNotFound  = "Data not found"
	ErrorFailedLogin = "Email Or Password Wrong"
)

func ErrorEmailRegistered(email string) string {
	return "Email " + email + " sudah terdaftar"
}
