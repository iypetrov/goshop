package users

const (
	INVALID  string = "INVALID"
	NONE     string = "NONE"
	GOOGLE   string = "GOOGLE"
	FACEBOOK string = "FACEBOOK"
	GITHUB   string = "GITHUB"
)

func ConvertToAuthProvider(provider string) string {
	switch provider {
	case NONE:
		return NONE
	case GOOGLE:
		return GOOGLE
	case FACEBOOK:
		return FACEBOOK
	case GITHUB:
		return GITHUB
	default:
		return INVALID
	}
}

const (
	CLIENT string = "CLIENT"
	ADMIN  string = "ADMIN"
)
