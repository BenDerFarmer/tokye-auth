package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/nrednav/cuid2"
	"gorm.io/gorm"
)

type User struct {
	ID          string      `gorm:"primaryKey"`
	Email       string      `gorm:"unique"`
	Credentials Credentials `gorm:"type:JSON"`
}

type Credentials []webauthn.Credential

func (c Credentials) Value() (driver.Value, error) {
	return json.Marshal(c)
}

func (c *Credentials) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("Failed to unmarshal JSONB value: %v", value)
	}

	return json.Unmarshal(bytes, c)
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = cuid2.Generate()
	return
}

func (user *User) WebAuthnID() []byte {
	return []byte(user.ID)
}

func (user *User) WebAuthnName() string {
	return user.Email
}

func (user *User) WebAuthnDisplayName() string {
	return user.Email
}

func (user *User) WebAuthnIcon() string {
	return ""
}

func (user *User) WebAuthnCredentials() []webauthn.Credential {
	return user.Credentials
}

func (user *User) AddCredential(credential *webauthn.Credential) {
	user.Credentials = append(user.Credentials, *credential)
}

func (user *User) UpdateCredential(credential *webauthn.Credential) {
	for i, c := range user.Credentials {
		if c.Descriptor().CredentialID.String() == credential.Descriptor().CredentialID.String() {
			user.Credentials[i] = *credential
		}
	}
}
