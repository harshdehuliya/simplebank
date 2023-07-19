package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTxDeadlock(t *testing.T) {
	store := NewStore(testDB)

	account1 := CreateRandomAccount(t)
	account2 := CreateRandomAccount(t)

	n := 10
	amount := int64(10)

	errs := make(chan error)

	// run n concurrent transfer transaction
	for i := 0; i < n; i++ {
		fromAccountID := account1.ID
		ToAccountID := account2.ID

		if i%2 == 1 {
			fromAccountID = account2.ID
			ToAccountID = account1.ID
		}
		go func() {
			ctx := context.Background()
			_, err := store.TransfersTx(ctx, TransferTxParams{
				FromAccountID: fromAccountID,
				ToAccountId:   ToAccountID,
				Amount:        amount,
			})

			errs <- err
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)
	}

	//check final updated balance
	updateAccount1, err := testQuery.GetAccount(context.Background(), account1.ID)
	require.NotEmpty(t, updateAccount1)
	require.NoError(t, err)

	updateAccount2, err := testQuery.GetAccount(context.Background(), account2.ID)
	require.NotEmpty(t, updateAccount2)
	require.NoError(t, err)

	fmt.Println(">>>before Tx: ", updateAccount1.Balance, updateAccount1.Balance)

	require.Equal(t, account1.Balance, updateAccount1.Balance)
	require.Equal(t, account2.Balance, updateAccount2.Balance)
}
