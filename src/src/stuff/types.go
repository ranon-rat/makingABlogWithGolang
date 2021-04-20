package stuff

type Document struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Mineatura string `json:"mineatura"`
	Body      string `json:"bodyOfDocument"`
}
type Publications struct {
	Size         int        `json:"Size"`
	Publications []Document `json:"Publications"`
	Cantidad     int        `json:"Cantidad"`
}
