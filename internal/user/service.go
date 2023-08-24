package user

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type RefreshTokenStorage interface {
	SaveRefreshToken(ctx context.Context, userId, refreshToken string) error
	RefreshTokenByID(ctx context.Context, userId string) (string, error)
}

type Service struct {
	rtStorage RefreshTokenStorage
}

func NewService(db *mongo.Database) *Service {
	s := &Service{
		rtStorage: NewRefreshTokenDB(db),
	}

	return s
}

func (s *Service) GenerateTokenPair(ctx context.Context, userID string) (TokenPair, error) {
	tokenPair, err := s.generateTokenPair(userID)
	if err != nil {
		return TokenPair{}, err
	}

	hashedRefreshToken, err := bcrypt.GenerateFromPassword([]byte(tokenPair.RefreshToken), bcrypt.DefaultCost)
	if err != nil {
		return TokenPair{}, fmt.Errorf("hashing refresh token: %w", err)
	}

	err = s.rtStorage.SaveRefreshToken(ctx, userID, string(hashedRefreshToken))
	if err != nil {
		return TokenPair{}, fmt.Errorf("save refresh token: %w", err)
	}

	return tokenPair, nil
}

func (s *Service) generateTokenPair(userID string) (TokenPair, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodRS512, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	})

	privateKey, err := os.ReadFile("medods.rsa")
	if err != nil {
		return TokenPair{}, fmt.Errorf("error reading private key file: %w", err)
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		return TokenPair{}, fmt.Errorf("error parsing RSA private key: %w", err)
	}

	accessTokenString, err := accessToken.SignedString(key)
	if err != nil {
		return TokenPair{}, fmt.Errorf("sign token: %w", err)
	}

	refreshBytes := make([]byte, 32)
	_, err = rand.Read(refreshBytes)
	if err != nil {
		return TokenPair{}, fmt.Errorf("rand read for refresh token: %w", err)
	}
	refreshToken := base64.URLEncoding.EncodeToString(refreshBytes)

	return TokenPair{
		AccessToken:  accessTokenString,
		RefreshToken: refreshToken,
	}, nil
}

func (s *Service) RefreshTokenPair(ctx context.Context, userID, refreshToken string) (TokenPair, error) {

	rt, err := s.rtStorage.RefreshTokenByID(ctx, userID)
	if err != nil {
		return TokenPair{}, fmt.Errorf("refresh token: %w", err)
	}
	// проверка рефреш токена

	err = bcrypt.CompareHashAndPassword([]byte(rt), []byte(refreshToken))
	if err != nil {
		return TokenPair{}, fmt.Errorf("compare hash and password: %w", err)
	}

	tokenPair, err := s.generateTokenPair(userID)
	if err != nil {
		return TokenPair{}, fmt.Errorf("generate token pair: %w", err)
	}

	return tokenPair, nil
}
