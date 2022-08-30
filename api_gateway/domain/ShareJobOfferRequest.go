package domain

type ShareJobRequest struct {
	ApiToken string
	Job      Job
}

type Job struct {
	JobId          string
	PublisherId    string
	RequiredSkills []string
	DatePosted     string
	DateValid      string
	CompanyName    string
	Position       string
	JobDescription string
}
