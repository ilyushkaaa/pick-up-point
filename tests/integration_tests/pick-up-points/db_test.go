package pick_up_points

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	database "homework/pkg/database/postgres"
	"homework/tests/fixtures"
	"homework/tests/states"
)

func setUp(t *testing.T, db database.DBops, tableName string) {
	t.Helper()
	ctx := context.Background()

	err := truncateTable(ctx, db, tableName)
	assert.NoError(t, err)

	err = resetIDSequence(ctx, db, tableName)
	assert.NoError(t, err)
}

func truncateTable(ctx context.Context, db database.DBops, tableName string) error {
	_, err := db.Exec(ctx, fmt.Sprintf("TRUNCATE table %s", tableName))
	return err
}

func resetIDSequence(ctx context.Context, db database.DBops, tableName string) error {
	_, err := db.Exec(ctx, fmt.Sprintf("ALTER SEQUENCE %s RESTART WITH 5010", tableName+"_id_seq"))
	return err
}

func fillDataBase(t *testing.T, db database.DBops) {
	t.Helper()

	ctx := context.Background()
	pp := fixtures.PickUpPoint().Valid().V()
	_, err := db.Exec(ctx,
		`INSERT INTO pick_up_points (id, name, phone_number, region, city, street, house_num)
              VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`,
		pp.ID, pp.Name, pp.PhoneNumber, pp.Address.Region, pp.Address.City, pp.Address.Street, pp.Address.HouseNum)
	assert.NoError(t, err)

	pp = fixtures.PickUpPoint().Valid().ID(states.PPID2).Name(states.PPName2).V()
	_, err = db.Exec(ctx,
		`INSERT INTO pick_up_points (id, name, phone_number, region, city, street, house_num)
              VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`,
		pp.ID, pp.Name, pp.PhoneNumber, pp.Address.Region, pp.Address.City, pp.Address.Street, pp.Address.HouseNum)
	assert.NoError(t, err)
}
