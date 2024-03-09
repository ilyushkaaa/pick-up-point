package command

import (
	"fmt"
	"strings"

	"homework/internal/order/delivery"
)

const commandsNumber = 6

type Commands []Command

func (cs Commands) Call(commandName string, params []string) error {
	switch commandName {
	case "help":
		cs.PrintCommandsDescription()
		return nil
	default:
		for _, cmd := range cs {
			if cmd.Name == commandName {
				return cmd.FuncToCall(params)
			}
		}
		return fmt.Errorf("unknown command")
	}
}

func (cs Commands) PrintCommandsDescription() {
	for i := 0; i < len(cs); i++ {
		params := strings.Join(cs[i].Params, " ")
		fmt.Printf("%d. %s %s\n   %s\n", i+1, cs[i].Name, params, cs[i].Description)
	}
}

func InitCommands(orderDelivery *delivery.OrderDelivery) Commands {
	commands := make([]Command, 0, commandsNumber)
	commands = append(commands, New("add_order",
		`command process getting order from courier and adding it in pick-up point. 
   It gets 3 required params: order ID, client ID and expiration date of order storage in format "yyyy-mm-dd"`,
		[]string{"<order_id>", "<client_id", "<expire_date>"}, orderDelivery.AddOrderDelivery))
	commands = append(commands, New("delete_order",
		`command process returning order to courier due to its expiration of shelf life. 
   It gets 1 required param: order ID`,
		[]string{"<order_id>"}, orderDelivery.DeleteOrderDelivery))
	commands = append(commands, New("issue_orders",
		`command process issuing orders to a client. 
   It gets one or more params: IDs of orders`,
		[]string{"<order_id>..."}, orderDelivery.IssueOrderDelivery))
	commands = append(commands, New("get_orders",
		`command process getting list of client's orders'. 
   It gets 1 required param: user ID; and 2 optional params: the first one limits the list to N last orders
   and second one shows orders which are in pick-up point now`,
		[]string{"<client_id>", "[<limit>]", "[PP-only]"}, orderDelivery.GetUserOrdersDelivery))
	commands = append(commands, New("return_order",
		`command process returning order from client to pick-up point. 
   It gets 2 required params: client ID and order ID`,
		[]string{"<client_id>", "<order_id>"}, orderDelivery.ReturnOrderDelivery))
	commands = append(commands, New("get_order_returns",
		`command shows list of all order returns with pagination. 
   It has 1 optional param: page number. Default value: 1`,
		[]string{"[<page_num>]"}, orderDelivery.GetOrderReturnsDelivery))

	return commands
}
