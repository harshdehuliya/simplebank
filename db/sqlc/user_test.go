package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/techschool/simplebank/db/util"
)

func CreateRandomUser(t *testing.T) Users {

	hashedPassword, err := util.HashedPassword(util.RandomString(6))
	require.NoError(t, err)
	arg := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}

	user, err := testQuery.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)

	require.NotZero(t, user.CreatedAt)
	require.True(t, user.PasswordChangedAt.IsZero())

	return user
}
func TestCreateUser(t *testing.T) {
	CreateRandomUser(t)
}
func TestGetUser(t *testing.T) {
	user1 := CreateRandomUser(t)
	user2, err := testQuery.GetUser(context.Background(), user1.Username)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.Email, user2.Email)

	require.NotZero(t, user2.Username)
	require.NotZero(t, user2.CreatedAt)

	require.WithinDuration(t, user1.PasswordChangedAt, user2.PasswordChangedAt, time.Second)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

// func TestUpdateAccount(t *testing.T) {
// 	account1 := CreateRandomAccount(t)
// 	arg := UpdateAccountParams{
// 		ID:      account1.ID,
// 		Balance: util.RandomMoney(),
// 	}

// 	account2, err := testQuery.UpdateAccount(context.Background(), arg)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, account2)

// 	require.Equal(t, account1.Owner, account2.Owner)
// 	require.Equal(t, arg.Balance, account2.Balance)
// 	require.Equal(t, account1.Currency, account2.Currency)

// 	require.NotZero(t, account2.ID)
// 	require.NotZero(t, account2.CreatedAt)

// 	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
// }
// func TestDeleteAccoun(t *testing.T) {
// 	account1 := CreateRandomAccount(t)
// 	err := testQuery.DeleteAccount(context.Background(), account1.ID)
// 	require.NoError(t, err)
// 	account2, err := testQuery.GetAccount(context.Background(), account1.ID)
// 	require.Error(t, err)
// 	require.EqualError(t, err, sql.ErrNoRows.Error())
// 	require.Empty(t, account2)

// }

// func TestListAccounts(t *testing.T) {
// 	n := 10

// 	for i := 0; i < n; i++ {
// 		CreateRandomAccount(t)
// 	}

// 	arg := ListAccountsParams{
// 		Limit:  5,
// 		Offset: 5,
// 	}

// 	accounts, err := testQuery.ListAccounts(context.Background(), arg)
// 	require.NoError(t, err)
// 	require.Len(t, accounts, 5)

// 	for _, account := range accounts {
// 		require.NotEmpty(t, account)
// 	}
// }
