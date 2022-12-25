package transactiondto

import "dewetour/models"

type TransResponse struct {
	ID         int                       `json:"id"`
	CounterQty int                       `json:"counterqty" form:"counterqty"`
	Total      int                       `json:"total" form:"total"`
	Status     string                    `json:"status" form:"status" gorm:"type: varchar(100)"`
	Attachment string                    `json:"attachment" form:"attachment" gorm:"type: varchar(100)"`
	UserId     int                       `json:"userid" form:"userid"`
	User       models.UserTransactionRes `json:"user"`
	TripId     int                       `json:"tripid" from:"tripid"`
	Trip       models.TripResponse       `json:"trip"`
}
