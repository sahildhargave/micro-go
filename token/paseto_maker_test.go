package token

import (
	"errors"
	"testing"
	"time"

	"github.com/sahil/simplebank/util"
	"github.com/stretchr/testify/require"
)

var (
	ErrEmptyToken = errors.New("token is empty")
	ErrInvalidKey = errors.New("invalid key provided for generating token")
)

func TestPasetoMaker(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	username := util.RandomOwner()
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, payload, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err = maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.Equal(t, username, payload.Username)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiresAt, time.Second)
}

func TestExpiredPasetoToken(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	token, payload, err := maker.CreateToken(util.RandomOwner(), -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	payload, err = maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredToken.Error())
	require.Nil(t, payload)

}

func TestInvalidSignaturePasetoToken(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	// Creating a valid token
	token, payload, err := maker.CreateToken(util.RandomOwner(), time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	// Modifying the token to invalidate the signature
	invalidToken := "invalid" + token[7:] // Modify a part of the token
	payload, err = maker.VerifyToken(invalidToken)
	require.Error(t, err)

	// Check if the error message contains the expected substring
	require.Contains(t, err.Error(), "incorrect token header")
	require.Nil(t, payload)
}

func TestEmptyPasetoToken(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	// Attempting to verify an empty token
	payload, err := maker.VerifyToken("")
	require.Error(t, err)
	// Update the expected error message to match the actual message from the library
	require.Contains(t, err.Error(), "incorrect token header")
	require.Nil(t, payload)
}

func TestInvalidPasetoMakerKey(t *testing.T) {
	// Creating a Paseto maker with an invalid key length
	invalidKey := util.RandomString(16)
	_, err := NewPasetoMaker(invalidKey)
	require.Error(t, err)

	// Update the expected error message to match the actual message from the library
	require.Contains(t, err.Error(), "Invalid key size: must be exactly 32 characters")
}
