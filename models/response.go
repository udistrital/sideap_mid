package models

type Response struct {
	Status string
	Res    interface{} `json:"res"`
}
