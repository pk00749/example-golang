package one

import (
	'fmt'
	'strconv'
)

type Good interface {
	settleAccount() int
	orderInfo() string
}

type Phone struct {
	name string
	quantity int
	price int
}

func (phone Phone) settleAccount() int{
	return phone.price * phone.quantity
}

func (phone Phone) orderInfo() string {
	
}