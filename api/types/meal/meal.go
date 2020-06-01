package meal

type Meal struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Section string  `json:"section"`
	Price   float32 `json:"price"`
}
