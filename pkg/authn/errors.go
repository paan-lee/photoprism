package authn

import (
	"errors"
	"fmt"

	"github.com/photoprism/photoprism/pkg/txt"
)

// Generic error messages for authentication and authorization:
var (
	ErrUnauthorized           = errors.New("unauthorized")
	ErrAccountAlreadyExists   = errors.New("account already exists")
	ErrAccountNotFound        = errors.New("account not found")
	ErrAccountDisabled        = errors.New("account disabled")
	ErrInvalidCredentials     = errors.New("invalid credentials")
	ErrInvalidShareToken      = errors.New("invalid share token")
	ErrInsufficientScope      = errors.New("insufficient scope")
	ErrDisabledInPublicMode   = errors.New("disabled in public mode")
	ErrAuthenticationDisabled = errors.New("authentication disabled")
)

// OAuth2-related error messages:
var (
	ErrInvalidClientID     = errors.New("invalid client id")
	ErrInvalidClientSecret = errors.New("invalid client secret")
)

// Username-related error messages:
var (
	ErrUsernameRequired = errors.New("username required")
	ErrInvalidUsername  = errors.New("invalid username")
)

// Passcode-related error messages:
var (
	ErrPasscodeRequired         = errors.New("passcode required")
	ErrPasscodeNotSetUp         = errors.New("passcode required, but not set up")
	ErrPasscodeNotVerified      = errors.New("passcode not verified")
	ErrPasscodeAlreadyActivated = errors.New("passcode already activated")
	ErrPasscodeNotSupported     = errors.New("passcode not supported")
	ErrInvalidPasscode          = errors.New("invalid passcode")
	ErrInvalidPasscodeFormat    = errors.New("invalid passcode format")
	ErrInvalidPasscodeKey       = errors.New("invalid passcode key")
	ErrInvalidPasscodeType      = errors.New("invalid passcode type")
)

// Password-related error messages:
var (
	ErrInvalidPassword     = errors.New("invalid password")
	ErrPasswordRequired    = errors.New("password required")
	ErrPasswordTooShort    = errors.New("password is too short")
	ErrPasswordTooLong     = errors.New(fmt.Sprintf("password must have less than %d characters", txt.ClipPassword))
	ErrPasswordsDoNotMatch = errors.New("passwords do not match")
)

// WebDAV-related error messages:
var (
	ErrWebDAVAccessDisabled     = errors.New("webdav access is disabled")
	ErrFailedToCreateUploadPath = errors.New("failed to create upload path")
	ErrBasicAuthDoesNotMatch    = errors.New("basic auth username does not match")
)
