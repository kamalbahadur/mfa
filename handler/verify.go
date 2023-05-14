package handler

import (
	"github.com/pquerna/otp/totp"
	"mfa/persistence"
	"net/http"
)

type verify struct {
	repo persistence.Repository
}

func NewVerifyHandler(repo persistence.Repository) *verify {
	return &verify{repo: repo}
}

func (s *verify) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	userId := request.URL.Query().Get("userId")
	code := request.URL.Query().Get("code")
	if code == "" || userId == "" {
		notVerified(writer)
		return
	}

	secret, err := s.repo.GetSharedSecret(userId)
	if err != nil {
		notVerified(writer)
		return
	}

	verified := totp.Validate(code, secret)

	if !verified {
		notVerified(writer)
		return
	}

	writer.Write([]byte("Verified!"))
	writer.WriteHeader(http.StatusOK)
}

func notVerified(writer http.ResponseWriter) {
	writer.Write([]byte("Not Verified!"))
	writer.WriteHeader(http.StatusInternalServerError)
}
