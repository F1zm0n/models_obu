package models

type OBUData struct {
	OBUID int     `json:"obu_id"`
	Lat   float64 `json:"lat"`
	Long  float64 `json:"long"`
}
type Invoice struct {
	OBUID         int     `db:"obu_id" json:"obu_id"`
	TotalDistance float64 `db:"value"  json:"total_distance"`
	TotalAmount   float64 `            json:"total_amount"`
}

type Distance struct {
	Value float64 `json:"value"`
	OBUID int     `json:"obu_id"`
	Unix  int64   `json:"unix"`
}
