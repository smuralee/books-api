package books

type FixedResponse struct {
	Status        string `json:"Status"`
	RemoteAddress string `json:"RemoteAddr"`
	Hostname      string `json:"Hostname"`
}
