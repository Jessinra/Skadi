package enums

type OrderState = string

const (
	StateCreated   = "CREATED"
	StateAccepted  = "ACCEPTED"
	StatePurchased = "PURCHASED"
	StateOnTheWay  = "ON_THE_WAY"
	StateDelivered = "DELIVERED"
	StateRevived   = "REVIVED"
	StateCompleted = "COMPLETED"
)
