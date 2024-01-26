package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/iden3/go-iden3-core/v2/w3c"
	"github.com/iden3/go-service-template/pkg/credential"
	"github.com/iden3/go-service-template/pkg/logger"
	"github.com/iden3/go-service-template/pkg/router/http/middleware"
	"github.com/iden3/go-service-template/pkg/services/issuer"
	"github.com/iden3/iden3comm/v2/packers"
	"github.com/iden3/iden3comm/v2/protocol"
)

type IssuerHandlers struct {
	issuerService *issuer.IssuerService
	host          string
}

func NewIssuerHandlers(issuerService *issuer.IssuerService, host string) IssuerHandlers {
	return IssuerHandlers{
		issuerService: issuerService,
		host:          host,
	}
}

func (h *IssuerHandlers) CreateClaim(w http.ResponseWriter, r *http.Request) {
	credentialReq := &credential.CredentialRequest{}
	if err := json.NewDecoder(r.Body).Decode(credentialReq); err != nil {
		logger.WithContext(r.Context()).WithError(err).
			Error("error unmarshalizing request")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	issuerDID := r.Context().Value(middleware.DIDContextKey{}).(*w3c.DID)
	recordID, err := h.issuerService.IssueCredential(r.Context(), issuerDID, credentialReq)
	if err != nil {
		logger.WithContext(r.Context()).WithError(err).
			Error("error issuing credential")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err = json.NewEncoder(w).Encode(map[string]string{"id": recordID}); err != nil {
		logger.WithContext(r.Context()).WithError(err).
			Error("error marshalizing response")
	}
}

func (h *IssuerHandlers) IsRevokedClaim(w http.ResponseWriter, r *http.Request) {
	nonce := chi.URLParam(r, "nonce")
	n, err := strconv.ParseInt(nonce, 10, 64)
	if err != nil {
		logger.WithContext(r.Context()).WithError(err).
			Error("error parsing nonce")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	issuerDID := r.Context().Value(middleware.DIDContextKey{}).(*w3c.DID)
	isRevoked, err := h.issuerService.IsRevokedCredential(r.Context(), issuerDID, uint64(n))
	if err != nil {
		logger.WithContext(r.Context()).WithError(err).
			Error("error checking if credential is revoked")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"mtp":{"existence": %t}}`, isRevoked)
}

func (h *IssuerHandlers) RevokeClaim(w http.ResponseWriter, r *http.Request) {
	nonce := chi.URLParam(r, "nonce")
	n, err := strconv.ParseInt(nonce, 10, 64)
	if err != nil {
		logger.WithContext(r.Context()).WithError(err).
			Error("error parsing nonce")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	issuerDID := r.Context().Value(middleware.DIDContextKey{}).(*w3c.DID)
	if err := h.issuerService.RevokeCredential(r.Context(), issuerDID, uint64(n)); err != nil {
		logger.WithContext(r.Context()).WithError(err).
			Error("error revoking credential")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
}

func (h *IssuerHandlers) GetUserVCs(w http.ResponseWriter, r *http.Request) {
	subject := r.URL.Query().Get("subject")
	schemaType := r.URL.Query().Get("schemaType")
	issuerDID := r.Context().Value(middleware.DIDContextKey{}).(*w3c.DID)
	vcs, err := h.issuerService.GetUserCredentials(r.Context(), issuerDID, subject, schemaType)
	if err != nil {
		logger.WithContext(r.Context()).WithError(err).
			Error("error getting user credentials")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(vcs); err != nil {
		logger.WithContext(r.Context()).WithError(err).
			Error("error marshalizing response")
	}
}

func (h *IssuerHandlers) GetUserVCByID(w http.ResponseWriter, r *http.Request) {
	claimID := chi.URLParam(r, "claimId")
	issuerDID := r.Context().Value(middleware.DIDContextKey{}).(*w3c.DID)
	vc, err := h.issuerService.GetCredentialByID(r.Context(), issuerDID, claimID)
	if err != nil {
		logger.WithContext(r.Context()).WithError(err).
			Error("error getting user credential by id")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/ld+json")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(vc); err != nil {
		logger.WithContext(r.Context()).WithError(err).
			Error("error marshalizing response")
	}
}

func (h *IssuerHandlers) GetOffer(w http.ResponseWriter, r *http.Request) {
	claimID := r.URL.Query().Get("claimId")
	if claimID == "" {
		logger.WithContext(r.Context()).Error("claimId query param is required")
		http.Error(w, "claimId query param is required", http.StatusBadRequest)
		return
	}
	issuerDID := r.Context().Value(middleware.DIDContextKey{}).(*w3c.DID)
	subject := r.URL.Query().Get("subject")
	if subject == "" {
		logger.WithContext(r.Context()).Error("subject query param is required")
		http.Error(w, "subject query param is required", http.StatusBadRequest)
		return
	}

	offerMessage := protocol.CredentialsOfferMessage{
		ID:       uuid.New().String(),
		ThreadID: uuid.New().String(),
		Typ:      packers.MediaTypePlainMessage,
		Type:     protocol.CredentialOfferMessageType,
		Body: protocol.CredentialsOfferMessageBody{
			URL: fmt.Sprintf("%s/api/v1/agent", strings.Trim(h.host, "/")),
			Credentials: []protocol.CredentialOffer{
				{
					ID:          claimID,
					Description: "BalanceCredential",
				},
			},
		},
		From: issuerDID.String(),
		To:   subject,
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(offerMessage); err != nil {
		logger.WithContext(r.Context()).WithError(err).
			Error("error marshalizing response")
	}
}

func (h *IssuerHandlers) GetIssuersList(w http.ResponseWriter, r *http.Request) {
	issuers := h.issuerService.GetIssuersList(r.Context())
	w.Header().Set("Content-Type", "application/ld+json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(issuers); err != nil {
		logger.WithContext(r.Context()).WithError(err).
			Error("error marshalizing response")
	}
}
