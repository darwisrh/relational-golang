package transactiondto

type TransRequest struct {
	CounterQty int    `json:"counterqty"`
	Total      int    `json:"counterqty"`
	Status     string `json:"status:" gorm:"type: varchar(100)"`
	Attachment string `json:"status:" gorm:"type: varchar(100)"`
	TripId     int    `json:"tripid"`
}
