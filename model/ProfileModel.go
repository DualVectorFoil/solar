package model

type Profile struct {
	ID string
	PhoneNum string
	UserName string
	Email string
	AvatarUrl string
	Pwd string
	Locale string
	Bio string
	Followers int
	Following int
	ArtworkCount int
	RegisterAt int64
	LastLoginAt int64
	Token string
}
