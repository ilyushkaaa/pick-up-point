package command

import (
	"fmt"
	"strings"

	"homework/Homework-1/internal/command/model"
	"homework/Homework-1/internal/order/delivery"
)

func PrintCommandsDescription(commands []model.Command) {
	for i := 0; i < len(commands); i++ {
		params := strings.Join(commands[i].Params, " ")
		strToPrint := fmt.Sprintf("%d. %s %s\n   %s", i+1, commands[i].Name, params, commands[i].Description)
		fmt.Println(strToPrint)
	}
}

func InitCommands(fileOrderDelivery *delivery.OrderDelivery) []model.Command {
	commands := make([]model.Command, 6)
	commands[0] = model.NewCommand("add_order",
		`command process getting order from courier and adding it in pick-up point. 
   It gets 3 required params: order ID, client ID and expiration date of order storage in format "yyyy-mm-dd"`,
		[]string{"<order_id>", "<client_id", "<expire_date>"}, fileOrderDelivery.AddOrderDelivery)
	commands[1] = model.NewCommand("delete_order",
		`command process returning order to courier due to its expiration of shelf life. 
   It gets 1 required param: order ID`,
		[]string{"<order_id>"}, fileOrderDelivery.DeleteOrderDelivery)
	commands[2] = model.NewCommand("issue_orders",
		`command process issuing orders to a client. 
   It gets one or more params: IDs of orders`,
		[]string{"<order_id>..."}, fileOrderDelivery.IssueOrderDelivery)
	commands[3] = model.NewCommand("get_orders",
		`command process getting list of client's orders'. 
   It gets 1 required param: user ID; and 2 optional params: the first one limits the list to N last orders
   and second one shows orders which are in pick-up point now`,
		[]string{"<client_id>", "[<limit>]", "[PP-only]"}, fileOrderDelivery.GetUserOrdersDelivery)
	commands[4] = model.NewCommand("return_order",
		`command process returning order from client to pick-up point. 
   It gets 2 required params: client ID and order ID`,
		[]string{"<client_id>", "<order_id>"}, fileOrderDelivery.ReturnOrderDelivery)
	commands[5] = model.NewCommand("get_order_returns",
		`command shows list of all order returns with pagination. 
   It does not have any params`,
		[]string{}, fileOrderDelivery.GetOrderReturnsDelivery)

	return commands
}
