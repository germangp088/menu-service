package meal

type Meal struct {
	Name    string  `json:"name"`
	Section string  `json:"section"`
	Price   float32 `json:"price"`
}
