package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/protocol/webauthncose"
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/crypto/mpc"
	"github.com/onsonr/sonr/pkg/common"
	"github.com/onsonr/sonr/pkg/common/cookie"
	"github.com/onsonr/sonr/pkg/common/response"
	"github.com/onsonr/sonr/pkg/common/session"
	"github.com/onsonr/sonr/pkg/gateway/handlers/views/register"
)

func HandleRegisterView(c echo.Context) error {
	return response.TemplEcho(c, register.ProfileFormView())
}

func HandleRegisterStart(c echo.Context) error {
	firstName := c.FormValue("first_name")
	lastName := c.FormValue("last_name")
	handle := c.FormValue("handle")
	cookie.Write(c, cookie.UserName, fmt.Sprintf("%s %s", firstName, lastName))
	cookie.Write(c, cookie.UserHandle, handle)
	ks, err := mpc.NewKeyset()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	req := getLinkCredentialRequest(c, ks.Address(), handle, ks.UserJSON())
	return response.TemplEcho(c, register.LinkCredentialView(req))
}

func HandleRegisterFinish(c echo.Context) error {
	// cred := c.FormValue("credential")
	return response.TemplEcho(c, register.LoadingVaultView())
}

// ╭───────────────────────────────────────────────────────────╮
// │                  Registration Components                  │
// ╰───────────────────────────────────────────────────────────╯

func getLinkCredentialRequest(c echo.Context, addr string, handle string, userKSJSON string) register.LinkCredentialRequest {
	cc := session.GetData(c)
	usr := buildUserEntity(addr, handle)
	blob := buildLargeBlob(userKSJSON)
	service := buildServiceEntity(c)
	return register.LinkCredentialRequest{
		Platform:        cc.Platform,
		Handle:          handle,
		DeviceModel:     cc.DeviceModel,
		Architecture:    cc.UserArchitecture,
		Address:         addr,
		RegisterOptions: buildRegisterOptions(usr, blob, service),
	}
}

func buildRegisterOptions(user protocol.UserEntity, blob common.LargeBlob, service protocol.RelyingPartyEntity) protocol.PublicKeyCredentialCreationOptions {
	return protocol.PublicKeyCredentialCreationOptions{
		Timeout:     10000,
		Attestation: protocol.PreferDirectAttestation,
		AuthenticatorSelection: protocol.AuthenticatorSelection{
			AuthenticatorAttachment: "platform",
			ResidentKey:             protocol.ResidentKeyRequirementPreferred,
			UserVerification:        "preferred",
		},
		RelyingParty: service,
		User:         user,
		Extensions: protocol.AuthenticationExtensions{
			"largeBlob": blob,
		},
		Parameters: []protocol.CredentialParameter{
			{
				Type:      "public-key",
				Algorithm: webauthncose.AlgES256,
			},
			{
				Type:      "public-key",
				Algorithm: webauthncose.AlgES256K,
			},
			{
				Type:      "public-key",
				Algorithm: webauthncose.AlgEdDSA,
			},
		},
	}
}

func buildLargeBlob(userKeyshareJSON string) common.LargeBlob {
	return common.LargeBlob{
		Support: "required",
		Write:   userKeyshareJSON,
	}
}

func buildUserEntity(userAddress string, userHandle string) protocol.UserEntity {
	return protocol.UserEntity{
		ID:          userAddress,
		DisplayName: userHandle,
		CredentialEntity: protocol.CredentialEntity{
			Name: userAddress,
		},
	}
}

func buildServiceEntity(c echo.Context) protocol.RelyingPartyEntity {
	return protocol.RelyingPartyEntity{
		CredentialEntity: protocol.CredentialEntity{
			Name: "Sonr.ID",
		},
		ID: c.Request().Host,
	}
}
