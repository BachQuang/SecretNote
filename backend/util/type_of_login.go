package util

const (
	GOOGLE = "GOOGLE"
)

func IsSupportedLogin(typeOfLogin string) bool {
	switch typeOfLogin {
	case GOOGLE:
		return true
	}
	return false
}
