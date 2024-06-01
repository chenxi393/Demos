package main

import "errors"

type OrderState uint8

const (
	Created OrderState = 1 << iota
	Paid
	Delivering
	Received
	Done
	Cancelling
	Returning
	Closed
)

type order struct {
	state OrderState
}

func NewOrder() *order {
	return &order{
		state: Created,
	}
}

func (o *order) CanPay() bool {
	return o.state == Created
}

func (o *order) CanDeliver() bool {
	return o.state == Paid
}

func (o *order) CanCancel() bool {
	return o.state == Created || o.state == Paid
}

func (o *order) CanReceive() bool {
	return o.state == Delivering
}

func (o *order) PaymentService() bool {
	//调用远程接口完成实际支付

	return false
}

// 然后是关键操作的实现，比如支付
func (o *order) Pay() (bool, error) {
	if o.CanPay() {
		if ok := o.PaymentService(); ok {
			o.state = Paid
			return true, nil
		}

	} else {
		//抛出错误
		return false, errors.New("Pay Error")
	}

	return false, nil
}

func (o *order) Cancel() (bool, error) {
	if o.CanCancel() {
		// 取消订单
		o.state = Cancelling

		// 取消订单，申请审批和清理数据，如果顺利成功再关闭订单
		o.state = Closed
	} else {
		// 抛出错误
		return false, errors.New("Cancel Error")
	}

	return false, nil
}
