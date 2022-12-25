package models

type TransactionModels struct {
	ID         int                `json:"id" gorm:"primaryKey"`
	CounterQty int                `json:"counterqty"`
	Total      int                `json:"total"`
	Status     string             `json:"status" gorm:"type: varchar(50)"`
	Attachment string             `json:"attachment" gorm:"type: varchar(50)"`
	TripId     int                `json:"tripid" form:"tripid"`
	Trip       TripResponse       `json:"trip"`
	UserId     int                `json:"userid" form:"userid"`
	User       UserTransactionRes `json:"user"`
}
