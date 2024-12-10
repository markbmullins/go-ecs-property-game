package components

type Tenant struct {
	Happiness        float64
	RentDue          float64
	MonthsWithoutPay int
	MoveOutChance    float64
	DesiredRent      float64
}

type TenantList struct {
	Tenants []Tenant
}
