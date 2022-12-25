package transactiondto

type TransRequest struct {
	CounterQty int    `json:"counterqty" form:"counterqty"`
	Total      int    `json:"total" form:"total"`
	Status     string `json:"status:" form:"status" gorm:"type: varchar(100)"`
	Attachment string `json:"attachment:" form:"attachment" gorm:"type: varchar(100)"`
	TripId     int    `json:"tripid" form:"tripid"`
	UserId     int    `json:"userid" form:"userid"`
}
