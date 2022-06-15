package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	// run concurrent goroutines to test transaction
	n := 6
	amount := int64(10)

	//Collect results of transfers
	errors := make(chan error)
	// results := make(chan TransferTxResult)

	for i := 0; i < n; i++ {
		fromAccountID := account1.ID
		toAccountID := account2.ID

		if i%2 == 1 {
			fromAccountID = account2.ID
			toAccountID = account1.ID
		}
		go func() {
			_, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: fromAccountID,
				ToAccountID:   toAccountID,
				Amount:        amount,
			})

			errors <- err
			// results <- result
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errors
		require.NoError(t, err)

		// require.NotEmpty(t, r)

		// transfer := r.Transfer
		// require.NotEmpty(t, transfer)
		// require.Equal(t, account1.ID, transfer.FromAccountID)
		// require.Equal(t, account2.ID, transfer.ToAccountID)
		// require.Equal(t, amount, transfer.Amount)
		// require.NotZero(t, transfer.ID)
		// require.NotZero(t, transfer.CreatedAt)

		// _, err = store.GetTransfer(context.Background(), transfer.ID)
		// require.NoError(t, err)

		// fromEntry := r.FromEntry
		// require.NotEmpty(t, fromEntry)
		// require.Equal(t, account1.ID, fromEntry.AccountID)
		// require.Equal(t, -amount, fromEntry.Amount)
		// require.NotZero(t, fromEntry.ID)
		// require.NotZero(t, fromEntry.CreatedAt)

		// _, err = store.GetEntry(context.Background(), fromEntry.ID)
		// require.NoError(t, err)

		// toEntry := r.ToEntry
		// require.NotEmpty(t, toEntry)
		// require.Equal(t, account2.ID, toEntry.AccountID)
		// require.Equal(t, amount, toEntry.Amount)
		// require.NotZero(t, toEntry.ID)
		// require.NotZero(t, toEntry.CreatedAt)

		// _, err = store.GetEntry(context.Background(), toEntry.ID)
		// require.NoError(t, err)

		// fromAccount := r.FromAccount
		// require.NotEmpty(t, fromAccount)
		// require.Equal(t, account1.ID, fromAccount.ID)

		// toAccount := r.ToAccount
		// require.NotEmpty(t, toAccount)
		// require.Equal(t, account2.ID, toAccount.ID)

		// diff1 := account1.Balance - fromAccount.Balance
		// diff2 := toAccount.Balance - account2.Balance
		// require.Equal(t, diff1, diff2)
		// require.True(t, diff1 > 0)
		// require.True(t, diff1%amount == 0)

		// k := int(diff1 / amount)
		// require.True(t, k >= 1 && k <= n)
		// require.NotContains(t, existed, k)
		// existed[k] = true
	}

	updateAccount1, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	updateAccount2, err := testQueries.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)

	require.Equal(t, account1.Balance, updateAccount1.Balance)
	require.Equal(t, account2.Balance, updateAccount2.Balance)
}
