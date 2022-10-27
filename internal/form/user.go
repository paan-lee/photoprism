package form

import (
	"github.com/urfave/cli"

	"github.com/photoprism/photoprism/pkg/clean"
)

// User represents a user account form.
type User struct {
	UserName    string       `json:"Name,omitempty" yaml:"Name,omitempty"`
	UserEmail   string       `json:"Email,omitempty" yaml:"Email,omitempty"`
	DisplayName string       `json:"DisplayName,omitempty" yaml:"DisplayName,omitempty"`
	UserRole    string       `json:"Role,omitempty" yaml:"Role,omitempty"`
	SuperAdmin  bool         `json:"SuperAdmin,omitempty" yaml:"SuperAdmin,omitempty"`
	CanLogin    bool         `json:"CanLogin,omitempty" yaml:"CanLogin,omitempty"`
	WebDAV      bool         `json:"WebDAV,omitempty" yaml:"WebDAV,omitempty"`
	UserAttr    string       `json:"Attr,omitempty" yaml:"Attr,omitempty"`
	BasePath    string       `json:"BasePath,omitempty" yaml:"BasePath,omitempty"`
	UploadPath  string       `json:"UploadPath,omitempty" yaml:"UploadPath,omitempty"`
	Password    string       `json:"Password,omitempty" yaml:"Password,omitempty"`
	UserDetails *UserDetails `json:"Details,omitempty"`
}

// NewUserFromCli creates a new form with values from a CLI context.
func NewUserFromCli(ctx *cli.Context) User {
	return User{
		UserName:    clean.Username(ctx.Args().First()),
		UserEmail:   clean.Email(ctx.String("email")),
		DisplayName: clean.Name(ctx.String("name")),
		UserRole:    clean.Role(ctx.String("role")),
		SuperAdmin:  ctx.Bool("superadmin"),
		CanLogin:    !ctx.Bool("no-login"),
		WebDAV:      ctx.Bool("webdav"),
		UserAttr:    clean.Attr(ctx.String("attr")),
		BasePath:    clean.UserPath(ctx.String("base-path")),
		UploadPath:  clean.UserPath(ctx.String("upload-path")),
		Password:    clean.Password(ctx.String("password")),
	}
}

// Name returns the sanitized username in lowercase.
func (f *User) Name() string {
	return clean.Username(f.UserName)
}

// Email returns the sanitized email in lowercase.
func (f *User) Email() string {
	return clean.Email(f.UserEmail)
}

// Role returns the sanitized user role string.
func (f *User) Role() string {
	return clean.Role(f.UserRole)
}

// Attr returns the sanitized user account attributes.
func (f *User) Attr() string {
	return clean.Attr(f.UserAttr)
}
