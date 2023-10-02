package db

import (
	"context"
	"database/sql"
	"simplebank/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func CreateRandomEntry(t *testing.T) Entry {
	account1 := CreateRandomAccount(t)
	arg := CreateEntryParams{
		AccountID: account1.ID,
		Amount:    util.RandomMoney(),
	}
	entry, err := testQueries.CreateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, entry.AccountID, account1.ID)
	require.Equal(t, entry.Amount, arg.Amount)

	return entry
}

func TestGetEntry(t *testing.T) {
	entry1 := CreateRandomEntry(t)

	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.Equal(t, entry2.AccountID, entry1.AccountID)
	require.Equal(t, entry2.CreatedAt, entry1.CreatedAt)
}

func TestListEntries(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomEntry(t)
	}
	arg := ListEntriesParams{
		Limit:  5,
		Offset: 5,
	}
	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}

func TestUpdateEntry(t *testing.T) {
	entry1 := CreateRandomEntry(t)

	arg := UpdateEntryParams{
		ID:     entry1.ID,
		Amount: util.RandomMoney(),
	}

	entry2, err := testQueries.UpdateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, arg.Amount, entry2.Amount)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.CreatedAt, entry2.CreatedAt)
}

func TestDeleteEntry(t *testing.T) {
	entry1 := CreateRandomEntry(t)

	err := testQueries.DeleteEntry(context.Background(), entry1.ID)
	require.NoError(t, err)

	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, entry2)
}
