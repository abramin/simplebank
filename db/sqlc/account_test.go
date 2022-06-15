package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/techschool/simplebank/util"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomBalance(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	a := createRandomAccount(t)
	account, err := testQueries.GetAccount(context.Background(), a.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, a.ID, account.ID)
}

func TestUpdateAccount(t *testing.T) {
	a := createRandomAccount(t)
	arg := UpdateAccountParams{
		ID:      a.ID,
		Balance: util.RandomBalance(),
	}
	account, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, a.ID, account.ID)
	require.Equal(t, account.Balance, arg.Balance)
}

func TestDeleteAccount(t *testing.T) {
	a := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), a.ID)
	require.NoError(t, err)

	account, err := testQueries.GetAccount(context.Background(), a.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account)
}
