package delivery

import (
	"fmt"
	"strconv"
	"strings"
)

func (od *OrderDelivery) GetOrderReturnsDelivery(args []string) {
	if len(args) != 0 {
		fmt.Println("bad number of params")
		return
	}

	orderPages, err := od.service.GetOrderReturnsService()
	if err != nil {
		fmt.Printf("error in getting order returns: %s\n", err)
		return
	}

	pageNum := 1
	pagesNums := make([]string, len(orderPages))
	for i, _ := range orderPages {
		pagesNums[i] = strconv.Itoa(i + 1)
	}
	pagesChoiceStr := strings.Join(pagesNums, " ")

	for {
		fmt.Printf("page %d\n", pageNum)
		for _, ord := range orderPages[pageNum-1] {
			fmt.Printf("%+v\n", ord)
		}
		fmt.Printf("switch page or exit\navailable pages: %s\n", pagesChoiceStr)
		var ch string
		_, err = fmt.Scan(&ch)
		if err != nil {
			fmt.Printf("input error: %s\n", err)
			return
		}
		pageNum, err = strconv.Atoi(ch)
		if err != nil || pageNum > len(orderPages) {
			return
		}
	}

}
