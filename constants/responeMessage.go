package constants

const (
	ErrorStatus        = "Error"
	SuccessStatus      = "Success"
	ErrorMessageSucces = ""
	SuccessInsertData  = "Success Insert Data"
	SuccessGetData     = "success retrieve data"
	ErrorIdNotValid    = "ID not Valid"
	ErrorDataNotFound  = "Data Not Found"
)

func ErrorEmailRegistered(email string) string {
	return "Email " + email + " sudah terdaftar"
}
