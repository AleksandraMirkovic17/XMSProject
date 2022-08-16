package domain

type JobOffer struct {
	JobId          string
	PublisherId    string
	RequiredSkills []string
	DatePosted     string
	DateValid      string
	CompanyName    string
	Position       string
	JobDescription string
}

type UserJobNode struct {
	UserID   string
	Username string
	Skills   []string
}
