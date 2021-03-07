package adding

type Candy struct {
	Id       string `json:"id",omitempty`
	Category string `json:"category"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
}
