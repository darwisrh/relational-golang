package transactiondto

import "time"

type TransResponse struct {
	ID             int       `json:"id"`
	CounterQty     int       `json:"counterqty"`
	Total          int       `json:"total"`
	Status         string    `json:"status" gorm:"type: varchar(100)"`
	Attachment     string    `json:"attachment" gorm:"type: varchar(100)"`
	Accomodation   string    `json:"accomodation" gorm:"type: varchar(100)"`
	Trasnportation string    `json:"transportation" gorm:"type: varchar(100)"`
	Eat            string    `json:"eat" gorm:"type: varcahr(50)`
	Day            int       `json:"day"`
	Night          int       `json:"night"`
	DateTrip       time.Time `json:"-"`
	Price          int       `json:"price"`
	Quota          int       `json:"quota"`
	Description    string    `json:"eat" gorm:"type: varcahr(50)`
	Image          string    `json:"eat" gorm:"type: varcahr(50)`
}
