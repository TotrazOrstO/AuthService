package user

type (
	TokenPair struct {
		AccessToken  string `json:"access_token" bson:"access_token"`
		RefreshToken string `json:"refresh_token" bson:"refresh_token"`
	}

	RefreshToken struct {
		ID           string `bson:"_id"`
		RefreshToken string `bson:"refresh_token"`
	}
)
