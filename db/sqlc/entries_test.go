package db

import (
	"context"
	"testing"

	"github.com/techschool/simplebank/util"

	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T, account Account) Entry {
	arg := CreateEntriesParams{
		AccountID: util.RandomInt(1, account.ID),
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntries(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)
	return entry
}

func TestCreateEntries(t *testing.T) {
	account := createRandomAccount(t)
	createRandomEntry(t, account)
}
