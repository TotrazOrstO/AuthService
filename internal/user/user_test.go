package user

import (
	"MedodsProject/pkg/client/mongodb"
	"MedodsProject/pkg/config"
	"context"
	"testing"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func testService(t *testing.T) *Service {
	t.Helper()
	cfg := config.New()
	ctx := context.Background()

	mongo, err := mongodb.NewClient(ctx, cfg.MongoDB)
	require.NoError(t, err)

	return NewService(mongo)
}

func TestGenerateTokenPair(t *testing.T) {
	s := testService(t)
	ctx := context.Background()

	userID := "test_user_id_1"
	tp, err := s.GenerateTokenPair(ctx, userID)
	require.NoError(t, err)
	assert.NotEmpty(t, tp.AccessToken)
	assert.NotEmpty(t, tp.RefreshToken)

	rt, err := s.rtStorage.RefreshTokenByID(ctx, userID)
	require.NoError(t, err)

	err = bcrypt.CompareHashAndPassword([]byte(rt), []byte(tp.RefreshToken))
	assert.NoError(t, err)
}

func TestRefreshTokenPair(t *testing.T) {
	s := testService(t)
	ctx := context.Background()

	userID := "test_user_id_2"
	tp, err := s.GenerateTokenPair(ctx, userID)
	require.NoError(t, err)
	time.Sleep(1 * time.Second)
	newTp, err := s.RefreshTokenPair(ctx, userID, tp.RefreshToken)
	require.NoError(t, err)

	assert.NotEmpty(t, newTp.AccessToken)
	assert.NotEmpty(t, newTp.RefreshToken)

	// проверяем что новая пара токенов не равна старым
	assert.NotEqual(t, tp.AccessToken, newTp.AccessToken)
	assert.NotEqual(t, tp.RefreshToken, newTp.RefreshToken)
}
