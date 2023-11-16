package aircall

const (
	CallStatusInitial  = "initial"
	CallStatusAnswered = "answered"
	CallStatusDone     = "done"

	CallDirectionInbound  = "inbound"
	CallDirectionOutbound = "outbound"
)

type Company struct {
	Name         string `json:"name"`
	UsersCount   int    `json:"users_count"`
	NumbersCount int    `json:"numbers_count"`
}

type User struct {
	ID         int      `json:"id"`
	DirectLink string   `json:"direct_link"`
	Name       string   `json:"name"`
	Email      string   `json:"email"`
	Available  bool     `json:"availaible"`
	Numbers    []Number `json:"numbers,omitempty"`
}

type Number struct {
	ID         int    `json:"id"`
	DirectLink string `json:"direct_link"`
	Name       string `json:"name"`
	Digits     string `json:"digits"`
	Country    string `json:"country"`
	TimeZone   string `json:"time_zone"`
	Open       bool   `json:"open"`
}

type Contact struct {
	ID           int           `json:"id"`
	DirectLink   string        `json:"direct_link"`
	FirstName    string        `json:"first_name"`
	LastName     string        `json:"last_name"`
	CompanyName  string        `json:"company_name"`
	Information  string        `json:"information"`
	PhoneNumbers []ContactInfo `json:"phone_numbers"`
	Emails       []ContactInfo `json:"emails"`
}

type ContactInfo struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type Call struct {
	ID         int      `json:"id"`
	DirectLink string   `json:"direct_link"`
	Status     string   `json:"status"`
	Direction  string   `json:"direction"`
	StartedAt  int      `json:"started_at"`
	AnsweredAt int      `json:"answered_at"`
	EndedAt    int      `json:"ended_at"`
	Duration   int      `json:"duration"`
	RawDigits  string   `json:"raw_digits"`
	Voicemail  string   `json:"voicemail"`
	Recording  string   `json:"recording"`
	Archived   bool     `json:"archived"`
	Number     Number   `json:"number"`
	User       User     `json:"user"`
	Contact    Contact  `json:"contact"`
	AssignedTo User     `json:"assigned_to"`
	Comments   []string `json:"comments"`
	Tags       []string `json:"tags"`
}

type Webhook struct {
	ID         int      `json:"id"`
	WebhookID  string   `json:"webhook_id"`
	DirectLink string   `json:"direct_link"`
	Status     string   `json:"status"`
	CreatedAt  string   `json:"created_at"`
	CustomName string   `json:"custom_name"`
	URL        string   `json:"url"`
	Active     bool     `json:"active"`
	Token      string   `json:"token"`
	Events     []string `json:"events"`
}
