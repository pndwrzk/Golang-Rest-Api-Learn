package constants

const (
	ErrorStatus        = "Error"
	SuccessStatus      = "Success"
	ErrorMessageSucces = ""
	SuccessInsertData  = "Success Insert Data"
	SuccessGetData     = "success retrieve data"
)

func ErrorEmailRegistered(email string) string {
	return "Email " + email + " sudah terdaftar"
}
