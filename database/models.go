package database

type Device struct {
	Name        string
	ProductLine string
	Varients    []string
}

type Firmware struct {
	Id         int
	ExternalId int
}
