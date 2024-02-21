package messages

type Message struct {
	Username string `json:"username" form:"username" query:"username"`
	Color    string `json:"color" form:"color" query:"color"`
	Message  string `json:"message" form:"message" query:"message"`
}
