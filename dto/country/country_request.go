package countrydto

type CountryRequest struct {
	Name string `json:"name" gorm:"type: varchar(50)"`
}
