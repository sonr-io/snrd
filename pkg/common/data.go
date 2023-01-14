package common

import (
	"encoding/base64"
	"errors"
	fmt "fmt"
	"strconv"
	"strings"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/shengdoushi/base58"
)

// Extensions are discussed in §9. WebAuthn Extensions (https://www.w3.org/TR/webauthn/#extensions).

// For a list of commonly supported extensions, see §10. Defined Extensions
// (https://www.w3.org/TR/webauthn/#sctn-defined-extensions).

type AuthenticationExtensionsClientOutputs map[string]interface{}

const (
	ExtensionAppID        = "appid"
	ExtensionAppIDExclude = "appidExclude"
)

// It converts a `WebauthnCredential` to a `webauthn.Credential`
func ConvertToWebauthnCredential(credential *WebauthnCredential) webauthn.Credential {
	return webauthn.Credential{
		ID:        credential.Id,
		PublicKey: credential.PublicKey,
		Authenticator: webauthn.Authenticator{
			AAGUID:       credential.Authenticator.Aaguid,
			SignCount:    credential.Authenticator.SignCount,
			CloneWarning: credential.Authenticator.CloneWarning,
		},
	}
}

// It converts a Package Struct to a WebauthnCredential Struct
func ConvertFromWebauthnCredential(credential *webauthn.Credential) *WebauthnCredential {
	return &WebauthnCredential{
		Id:        credential.ID,
		PublicKey: credential.PublicKey,
		Authenticator: &WebauthnAuthenticator{
			Aaguid:       credential.Authenticator.AAGUID,
			SignCount:    credential.Authenticator.SignCount,
			CloneWarning: credential.Authenticator.CloneWarning,
		},
	}
}

// VerifyCounter
// Step 17 of §7.2. about verifying attestation. If the signature counter value authData.signCount
// is nonzero or the value stored in conjunction with credential’s id attribute is nonzero, then
// run the following sub-step:
//
//	If the signature counter value authData.signCount is
//
//	→ Greater than the signature counter value stored in conjunction with credential’s id attribute.
//	Update the stored signature counter value, associated with credential’s id attribute, to be the value of
//	authData.signCount.
//
//	→ Less than or equal to the signature counter value stored in conjunction with credential’s id attribute.
//	This is a signal that the authenticator may be cloned, see CloneWarning above for more information.
func (a *WebauthnAuthenticator) UpdateCounter(authDataCount uint32) {
	if authDataCount <= a.SignCount && (authDataCount != 0 || a.SignCount != 0) {
		a.CloneWarning = true
		return
	}
	a.SignCount = authDataCount
}

func ConvertBoolToString(v bool) string {
	if v {
		return "TRUE"
	} else {
		return "FALSE"
	}
}

func ConvertStringToBool(v string) bool {
	if v == "TRUE" {
		return true
	}
	return false
}

// ConvertStdCredential creates a common.WebauthnCredential from a webauthn.Credential from the go-webauthn package
func ConvertStdCredential(wa *webauthn.Credential) *WebauthnCredential {
	transportsStr := []string{}
	for _, t := range wa.Transport {
		transportsStr = append(transportsStr, string(t))
	}
	return &WebauthnCredential{
		Id:              wa.ID,
		PublicKey:       wa.PublicKey,
		AttestationType: wa.AttestationType,
		Transport:       transportsStr,
		Authenticator: &WebauthnAuthenticator{
			Aaguid:       wa.Authenticator.AAGUID,
			SignCount:    wa.Authenticator.SignCount,
			CloneWarning: wa.Authenticator.CloneWarning,
		},
	}
}

// ToStdCredential converts a common WebauthnCredential to one that can be used for the go-webauthn package
func (c *WebauthnCredential) ToStdCredential() *webauthn.Credential {
	transports := []protocol.AuthenticatorTransport{}
	for _, t := range c.Transport {
		transports = append(transports, protocol.AuthenticatorTransport(t))
	}
	return &webauthn.Credential{
		ID:              c.Id,
		PublicKey:       c.PublicKey,
		AttestationType: c.AttestationType,
		Transport:       transports,
		Authenticator: webauthn.Authenticator{
			AAGUID:       c.Authenticator.Aaguid,
			SignCount:    c.Authenticator.SignCount,
			CloneWarning: c.Authenticator.CloneWarning,
		},
	}
}

func (c *WebauthnCredential) Did() string {
	return fmt.Sprintf("did:webauth:%s", base58.Encode(c.Id, base58.BitcoinAlphabet))
}

func (c *WebauthnCredential) PublicKeyMultibase() string {
	return "z" + base64.StdEncoding.EncodeToString(c.PublicKey)
}

// ToMetadata converts a common WebauthnCredential into a map[string]string
func (c *WebauthnCredential) ToMetadata() map[string]string {
	return map[string]string{
		"credential_id":               base64.StdEncoding.EncodeToString(c.Id),
		"authenticator.aaguid":        base64.StdEncoding.EncodeToString(c.Authenticator.Aaguid),
		"authenticator.clone_warning": ConvertBoolToString(c.Authenticator.CloneWarning),
		"authenticator.sign_count":    strconv.FormatUint(uint64(c.Authenticator.SignCount), 10),
		"transport":                   strings.Join(c.Transport, ","),
		"attestion_type":              c.AttestationType,
	}
}

// NewWebAuthnCredential creates a new WebauthnCredential from a ParsedCredentialCreationData and contains all needed information about a WebAuthn credential for storage.
// This is then used to create a VerificationMethod for the DID Document.
func NewWebAuthnCredential(c *protocol.ParsedCredentialCreationData) *WebauthnCredential {
	transportsStr := []string{}
	for _, t := range c.Transports {
		transportsStr = append(transportsStr, string(t))
	}
	return &WebauthnCredential{
		Id:              c.Response.AttestationObject.AuthData.AttData.CredentialID,
		PublicKey:       c.Response.AttestationObject.AuthData.AttData.CredentialPublicKey,
		AttestationType: c.Response.AttestationObject.Format,
		Transport:       transportsStr,
		Authenticator: &WebauthnAuthenticator{
			Aaguid:    c.Response.AttestationObject.AuthData.AttData.AAGUID,
			SignCount: c.Response.AttestationObject.AuthData.Counter,
		},
	}
}

// Validate verifies that this WebauthnCredential is identical to the go-webauthn package credential
func (c *WebauthnCredential) Validate(pc *webauthn.Credential) error {
	if len(c.PublicKey) != len(pc.PublicKey) {
		return errors.New("Credential Public Keys do not match")
	}
	return nil
}

// DeriveConfig returns a new WalletShareConfig with the prefix changed
func (c *WalletShareConfig) DeriveConfig(prefix string, childIndex int) *WalletShareConfig {
	newConf := c
	newConf.Bech32Prefix = prefix
	newConf.ChildIndex = int32(childIndex)
	return newConf
}

// NewWalletShare
func NewWalletShare(data []byte) (*WalletShareConfig, error) {
	wsc := &WalletShareConfig{}
	err := wsc.Unmarshal(data)
	if err != nil {
		return nil, err
	}
	return wsc, nil
}
