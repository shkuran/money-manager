package transaction

type Transaction struct {
	ID    int64   `json:"id"`
	Value float64 `json:"value"`
}