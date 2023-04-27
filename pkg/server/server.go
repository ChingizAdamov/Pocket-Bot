package server

import (
	"net/http"
	"strconv"

	"github.com/ChingizAdamov/pocket_bot/pkg/repository"
	"github.com/zhashkevych/go-pocket-sdk"
)

type AuthServer struct {
	server *http.Server
	pocketClient *pocket.Client
	tokenRepository repository.TokenRepositore
	redirectURL string
}

func NewAuthServer(pocketClient *pocket.Client, tokenRepository repository.TokenRepositore, redirectURL string) *AuthServer {
	return &AuthServer{pocketClient: pocketClient, tokenRepository: tokenRepository,redirectURL: redirectURL}
}

func (s *AuthServer) Start() error {
	s.server = &http.Server{
		Handler: s,
		Addr: ":80",
	}

	return s.server.ListenAndServe()
}

func (s *AuthServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	chatIDParam := r.URL.Query().Get("chat_id")
	if chatIDParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	chatID, err := strconv.ParseInt(chatIDParam, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	reqToken, err := s.tokenRepository.Get(chatID, repository.RequestTokens)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	authResp, err := s.pocketClient.Authorize(r.Context(), reqToken)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = s.tokenRepository.Save(chatID, authResp.AccessToken, repository.AccessTokens)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("location", s.redirectURL)
	w.WriteHeader(http.StatusMovedPermanently)
}



