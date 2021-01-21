package main

const (
	cantidad int = 9
)

type document struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Mineatura string `json:"mineatura"`
	Body      string `json:"bodyOfDocument"`
}
type publications struct {
	Size         int        `json:"Size"`
	Publications []document `json:"Publications"`
	Cantidad     int        `json:"Cantidad"`
}
