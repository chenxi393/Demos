package main

import (
	"context"
	"fmt"

	"github.com/looplab/fsm"
)

func main() {
	order := fsm.NewFSM(
		"created",
		fsm.Events{
			{Name: "pay", Src: []string{"created"}, Dst: "paid"}, //订单创建 - 支付 - 发货
			{Name: "deliver", Src: []string{"paid"}, Dst: "delivering"},
			{Name: "receive", Src: []string{"delivering"}, Dst: "received"}, // 收货
			{Name: "confirm", Src: []string{"received"}, Dst: "done"},       // 确认收货

			{Name: "cancel", Src: []string{"received", "paid"}, Dst: "cancelling"},      // 取消
			{Name: "return", Src: []string{"delivering", "received"}, Dst: "returning"}, // 退款
			{Name: "close", Src: []string{"cancelling", "returning"}, Dst: "closed"},    // 关闭订单
		},

		// Callbacks are added as a map specified as Callbacks where the key is parsed
		// as the callback event as follows, and called in the same order:
		//
		// 1. before_<EVENT> - called before event named <EVENT>
		//
		// 2. before_event - called before all events
		//
		// 3. leave_<OLD_STATE> - called before leaving <OLD_STATE>
		//
		// 4. leave_state - called before leaving all states
		//
		// 5. enter_<NEW_STATE> - called after entering <NEW_STATE>
		//
		// 6. enter_state - called after entering all states
		//
		// 7. after_<EVENT> - called after event named <EVENT>
		//
		// 8. after_event - called after all events

		fsm.Callbacks{
			"before_pay": func(_ context.Context, e *fsm.Event) {
				fmt.Println("支付服务申请中……")
				// 发送 before_pay 服务
			},
			"paid": func(_ context.Context, e *fsm.Event) {
				fmt.Println("支付成功")
			},
			"after_deliver": func(_ context.Context, e *fsm.Event) {
				fmt.Println("已通知用户：商品配送中")
			},
			"cancel": func(ctx context.Context, e *fsm.Event) {
				fmt.Println("订单取消")
				e.Cancel()
			},
			"return": func(ctx context.Context, e *fsm.Event) {
				fmt.Println("订单返回")
				e.Cancel()
			},
			"close": func(ctx context.Context, e *fsm.Event) {
				fmt.Println("订单关闭")
				e.Cancel()
			},
		},
	)

	fmt.Println("current status: ", order.Current())

	test := 1
	err := order.Event(context.Background(), "pay", test)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("current status: ", order.Current())

	err = order.Event(context.Background(), "deliver")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("current status: ", order.Current())

	err = order.Event(context.Background(), "receive")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("current status: ", order.Current())

	err = order.Event(context.Background(), "confirm")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("current status: ", order.Current())
}
