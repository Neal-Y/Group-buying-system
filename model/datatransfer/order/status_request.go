package order

type StatusRequest struct {
	Note   string `json:"note"`
	Status string `json:"status"`
}
