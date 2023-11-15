package db

import (
	// "context"
	"context"
	"fmt"
	"testing"
	// "github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T){
	store := NewStore(testDB)
	
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	fmt.Println("Before: ", account1.Balance, account2.Balance)

	
	// n := 1
	amount := int64(10)
	result, err := store.TransferTx(context.Background(), TransferTxParams{
		FromAccountID: account1.ID,
		ToAccountID: account2.ID,
		Amount: amount,
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v", result)

	// errs := make(chan error)
	// results := make(chan TransferTxResult)
	// // run n concurrent transfer transaction
	// for i := 0; i < n; i++ {
	// 	go func(){
	// 		result, err := store.TransferTx(context.Background(), TransferTxParams{
	// 			FromAccountID: account1.ID,
	// 			ToAccountID: account2.ID,
	// 			Amount: amount,
	// 		})
	// 		errs <- err
	// 		results <- result
	// 	}()
	// }
	
	// // Check results
	// for i := 0; i < n; i++ {
	// 	err := <- errs
	// 	require.NoError(t, err)

	// 	// check transfer
	// 	result := <- results
	// 	transfer := result.Transfer
	// 	require.NotEmpty(t, transfer)
	// 	require.Equal(t, account1.ID, transfer.FromAccountID)
	// 	require.Equal(t, account2.ID, transfer.ToAccountID)
	// 	require.Equal(t, amount, transfer.Amount)
	// 	require.NotZero(t, transfer.ID)
	// 	require.NotZero(t, transfer.CreatedAt)

	// 	_, err = store.GetTransfer(context.Background(), transfer.ID)
	// 	require.NoError(t, err)

		// TODO: check accounts' balance
	// }
}
