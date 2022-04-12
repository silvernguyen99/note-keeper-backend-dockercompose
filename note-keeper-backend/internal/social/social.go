package social

type SocialRequest struct {
	SocialToken string `json:""`
}

type Social struct {
	SocialAccount SocialAccount
}

type SocialInfo struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Birthday string `json:"birthday"`
	Email    string `json:"email"`
	Picture  string `json:"picture"`
}

type SocialAccount interface {
	// Concrete types should implement this method.
	GetInfo(accessToken string) (*SocialInfo, error)
}

func NewSocialInstant(scAcc SocialAccount) *Social {
	return &Social{
		SocialAccount: scAcc,
	}
}

func (s *Social) PerformGetInfo(accessToken string) (*SocialInfo, error) {
	return s.SocialAccount.GetInfo(accessToken)
}
