package storage

import (
	"context"
	"fmt"
	"strings"
	"time"
)

func (s *OrderStoragePG) IssueOrders(ctx context.Context, orderIDs map[uint64]struct{}) error {
	issueTime := time.Now()
	args := make([]interface{}, 0, len(orderIDs)+1)
	args = append(args, issueTime)
	placeholders := make([]string, 0, len(orderIDs))
	idCount := 0
	for id := range orderIDs {
		placeholders = append(placeholders, fmt.Sprintf("$%d", idCount+2))
		args = append(args, id)
		idCount++
	}
	query := fmt.Sprintf("UPDATE orders SET order_issue_date = $2 WHERE id IN (%s)", strings.Join(placeholders, ", "))
	_, err := s.db.Exec(ctx, query, args...)
	return err
}
