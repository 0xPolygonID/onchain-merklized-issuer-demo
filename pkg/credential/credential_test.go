package credential

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/iden3/go-iden3-core/v2/w3c"
	"github.com/stretchr/testify/require"
)

func TestNewCredential(t *testing.T) {
	tests := []struct {
		name               string
		issuer             string
		credentialRequest  *CredentialRequest
		jsonSchemaMetadata JSONSchemaMetadata
	}{
		{
			name:   "Create credential success",
			issuer: "did:polygonid:polygon:mumbai:2qCU58EJgrEMZqnh5RaCdNre9pFtMwtYdDosZ61Zjx",
			credentialRequest: &CredentialRequest{
				CredentialSchema: "ipfs://QmbgBjetG5V6DecQXTRrJ7s239b4aLydpjgB5Q6tiyZyUi",
				Type:             "BalanceCredential",
				Expiration:       1893456000,
				CredentialSubject: map[string]any{
					"balance": 100,
					"id":      "did:polygonid:polygon:mumbai:2qNt1BRRNQHHyz9L8RXPW3WcrRhaMiWn9j6Pz5MQDS",
					"type":    "BalanceCredential",
				},
				RevNonce: uintToPtr(100),
			},
			jsonSchemaMetadata: JSONSchemaMetadata{
				Metadata: struct {
					URIs map[string]string `json:"uris"`
				}{
					URIs: map[string]string{
						"jsonLdContext": "ipfs://QmbbTKPTJy5zpS2aWBRps1duU8V3zC84jthpWDXE9mLHBX",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			issuerDID, err := w3c.ParseDID(tt.issuer)
			require.NoError(t, err)

			credential, err := NewCredential(issuerDID, tt.credentialRequest, tt.jsonSchemaMetadata)
			require.NoError(t, err)

			credentialBytes, err := json.Marshal(credential)
			require.NoError(t, err)

			fmt.Println(string(credentialBytes))
		})
	}
}

func uintToPtr(v uint64) *uint64 {
	return &v
}
