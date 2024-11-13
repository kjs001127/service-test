package model

import "github.com/channel-io/go-lib/pkg/errors/apierr"

const (
	youAre   = "youAre"
	required = "required"
	action   = "action"
	scope    = "scope"
)

type roleType string

const (
	RoleTypeOwner             roleType = "owner"
	RoleTypeGeneralSettings   roleType = "generalSettings"
	OwnerErrMessage                    = "{unauthorized.access.owner}"
	GeneralSettingsErrMessage          = "{unauthorized.access}"
)

type UnauthorizedRoleError interface {
	HTTPStatusCode() int
	ErrorName() string
	Causes() []*apierr.Cause
	Error() string
}

type OwnerRoleError struct {
	apierr.HTTPErrorBuildable

	youAre   string
	required roleType
	message  string
}

func NewOwnerRoleError(youAre string, required roleType, message string) UnauthorizedRoleError {
	return &OwnerRoleError{
		message:  message,
		youAre:   youAre,
		required: required,
	}
}

func (u *OwnerRoleError) HTTPStatusCode() int {
	return 403
}

func (u *OwnerRoleError) ErrorName() string {
	return "unauthorizedRoleError"
}

func (u *OwnerRoleError) Causes() []*apierr.Cause {
	return []*apierr.Cause{
		{
			Message: u.message,
			Detail: map[string]interface{}{
				youAre:   u.youAre,
				required: u.required,
			},
		},
	}
}

func (u *OwnerRoleError) Error() string {
	return u.ErrorName()
}

type GeneralSettingsRoleError struct {
	apierr.HTTPErrorBuildable

	action  roleType
	message string
	scope   string
}

func NewGeneralSettingsRoleError(action roleType, message string, scope string) UnauthorizedRoleError {
	return &GeneralSettingsRoleError{
		action:  action,
		message: message,
		scope:   scope,
	}
}

func (g *GeneralSettingsRoleError) HTTPStatusCode() int {
	return 403
}

func (g *GeneralSettingsRoleError) ErrorName() string {
	return "unauthorizedPermissionError"
}

func (g *GeneralSettingsRoleError) Causes() []*apierr.Cause {
	return []*apierr.Cause{
		{
			Message: g.message,
			Detail: map[string]interface{}{
				action: g.action,
				scope:  g.scope,
			},
		},
	}
}

func (g *GeneralSettingsRoleError) Error() string {
	return g.ErrorName()
}
