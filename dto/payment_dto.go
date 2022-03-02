package dto

type GeneratePayment struct {
	Order []OrderWaitList
}

type StatusPayment struct {
}

type InvoicePayment struct {
	Message string
	Result  []CreateOrder
}
