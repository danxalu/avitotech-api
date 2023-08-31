package inter

type CreateUserAns struct {
	UserID string `json:"user_id"`
}

type CreateUserAsk struct {
	UserName string `json:"user_name"`
}

type ChangeSegmentAns struct {
	SegmentId int `json:"segments_id"`
	Slug       string `json:"slug"`
}

type ChangeSegmentAsk struct {
	Slug string `json:"slug"`
}


type ChangeUserSegmentAns struct {
	UserID   int `json:"user_id"`
	Segment []string `json:"segment"`
}

type ChangeUserSegmentAsk struct {
	UserID  int `json:"user_id"`
	AddSlug []int `json:"add_slug"`
	DeleteSlug []int `json:"delete_slug"`
}


type CheckSegmentAns struct {
	Segment []string `json:"segment"`
}

type CheckSegmentAsk struct {
	UserID int `json:"user_id"`
}
