package domain

type UserConn struct {
	UserID   string
	Username string
	IsPublic bool
}

type UserRecommendation struct {
	UserID   string
	Username string
	IsPublic bool
	IsMutual bool
	Mutual   int
}
