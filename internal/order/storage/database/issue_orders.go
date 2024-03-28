package storage

import (
	"context"
	"fmt"
	"strings"
	"time"
)

func (s *OrderStoragePG) IssueOrders(ctx context.Context, orderIDs map[uint64]struct{}) error {
	issueTime := time.Now()
	args := make([]interface{}, 0, len(orderIDs)+2)
	args = append(args, true)
	args = append(args, issueTime)
	placeholders := make([]string, 0, len(orderIDs))
	for i := 0; i < len(orderIDs); i++ {
		placeholders = append(placeholders, fmt.Sprintf("$%d", i+3))
	}
	for id := range orderIDs {
		args = append(args, id)
	}
	query := fmt.Sprintf("UPDATE orders SET is_issued = $1, order_issue_date = $2 WHERE id IN (%s)", strings.Join(placeholders, ", "))
	_, err := s.db.Cluster.Exec(ctx, query, args...)
	return err
}
