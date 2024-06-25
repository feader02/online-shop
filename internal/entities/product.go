package entities

type Product struct {
	Id          int     `db:"id" json:"id"`
	Name        string  `db:"name" json:"name"`
	Price       float64 `db:"price" json:"price"`
	Description string  `db:"description" json:"description"`
	Height      int     `db:"height" json:"height"`
	Width       int     `db:"width" json:"width"`
	Depth       int     `db:"depth" json:"depth"`
	Photo       string  `db:"photo" json:"photo"`
	Type        string  `db:"type" json:"type"`
}
