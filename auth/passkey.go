package auth

import (
	"github.com/ChaotenHG/auth-server/config"
	"github.com/go-webauthn/webauthn/webauthn"
)

var WebAuthn *webauthn.WebAuthn

func LoadPasskeyConfig(cfg *config.Config) {
	wconfig := &webauthn.Config{
		RPDisplayName: cfg.PassKey.DisplayName, // Display Name for your site
		RPID:          cfg.PassKey.RPID,        // Generally the FQDN for your site
		RPOrigins:     cfg.PassKey.Origins,     // The origin URLs allowed for WebAuthn requests
	}

	WebAuthn, _ = webauthn.New(wconfig)
}
