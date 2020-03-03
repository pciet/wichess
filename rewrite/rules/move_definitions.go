package rules

type Move struct {
	From Address `json:"from"`
	To   Address `json:"to"`
}

type MoveSet struct {
	From  Address   `json:"from"`
	Moves []Address `json:"moves"`
}
