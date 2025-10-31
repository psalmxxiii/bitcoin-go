package elliptic_curve

import (
	"fmt"
)

type FiniteElement struct {
	order uint64
	num   uint64
}

func NewFiniteElement(order uint64, num uint64) *FiniteElement {
	if num >= order {
		errMsg := fmt.Sprintf("num %d is out of the range [0 -  %d]", num, order-1)
		panic(errMsg)
	}
	return &FiniteElement{
		order: order,
		num:   num,
	}
}

func (fe *FiniteElement) String() string {
	return fmt.Sprintf("FiniteElement:{order: %d, num: %d}\n", fe.order, fe.num)
}

func (fe *FiniteElement) EqualTo(other *FiniteElement) bool {
	return fe.num == other.num && fe.order == other.order
}

func (fe *FiniteElement) Add(other *FiniteElement) *FiniteElement {
	if fe.order != other.order {
		errMsg := fmt.Sprintf("cannot add finite elements with different orders: %d and %d", fe.order, other.order)
		panic(errMsg)
	}
	sum := (fe.num + other.num) % fe.order
	return NewFiniteElement(fe.order, sum)
}

func (fe *FiniteElement) Sub(other *FiniteElement) *FiniteElement {
	if fe.order != other.order {
		errMsg := fmt.Sprintf("cannot subtract finite elements with different orders: %d and %d", fe.order, other.order)
		panic(errMsg)
	}
	diff := (fe.num + fe.order - other.num) % fe.order
	return NewFiniteElement(fe.order, diff)
}

func (fe *FiniteElement) Neg() *FiniteElement {
	neg := (fe.order - fe.num) % fe.order
	return NewFiniteElement(fe.order, neg)
}
