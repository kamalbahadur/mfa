package handler

import (
	"bytes"
	"github.com/pquerna/otp/totp"
	"image/png"
	"log"
	"mfa/persistence"
	"net/http"
)

type signup struct {
	repo persistence.Repository
}

func NewSignupHandler(repo persistence.Repository) *signup {
	return &signup{repo: repo}
}

func (s *signup) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	userId := request.URL.Query().Get("userId")
	if userId == "" {
		writer.Write([]byte("userId is required"))
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "myfakedomain.com",
		AccountName: "user@myfakedomain.com",
	})
	if err != nil {
		panic(err)
	}
	// Convert TOTP key into a PNG
	var buf bytes.Buffer
	img, err := key.Image(200, 200)
	if err != nil {
		panic(err)
	}
	png.Encode(&buf, img)

	err = s.repo.SaveSharedSecret(userId, key.Secret())
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	writer.Header().Set("Content-Type", "image/x-png")
	writer.Write(buf.Bytes())
}
