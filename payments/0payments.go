package payments

type Payment struct {
	Qty     int
	Amount  float64
	Success bool
	Reference, Currency,
	Status, Message string
}
