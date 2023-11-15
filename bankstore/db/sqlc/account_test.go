/*
Pattern: <filename>_test.go
*/
package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"
	"github.com/vadshi/go3/bankstore/util"
)

const (
	dbSource = "postgresql://postgres:postgres@localhost:5435/bankstoredb?sslmode=disable"
)

var ctx = context.Background()

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := pgx.Connect(ctx, dbSource)
	if err != nil {
		log.Fatal("Can not connect to db", err)
	}
	defer conn.Close(ctx)

	testQueries = New(conn)

	os.Exit(m.Run())
}

func TestCreateAccount(t *testing.T) {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomAmount(),
		Currency: Currency(util.RandomCurrency()),
	}
	account, err := testQueries.CreateAccount(ctx, arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}

func TestGetAccount(t *testing.T){
	
}