package model

import "github.com/channel-io/go-lib/pkg/errors/apierr"

const (
	youAre   = "youAre"
	required = "required"
)

type roleType string

const (
	RoleTypeOwner             roleType = "owner"
	RoleTypeGeneralSettings   roleType = "generalSettings"
	OwnerErrMessage                    = "{unauthorized.access.owner}"
	GeneralSettingsErrMessage          = "{unauthorized.access}"
)

type UnauthorizedRoleError struct {
	apierr.HTTPErrorBuildable

	youAre   string
	required roleType
	message  string
}

func (u *UnauthorizedRoleError) HTTPStatusCode() int {
	return 403
}

func (u *UnauthorizedRoleError) ErrorName() string {
	return "unauthorizedRoleError"
}

func (u *UnauthorizedRoleError) Causes() []*apierr.Cause {
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

func (u *UnauthorizedRoleError) Error() string {
	return u.ErrorName()
}

func NewUnauthorizedRoleError(youAre string, required roleType, message string) *UnauthorizedRoleError {
	return &UnauthorizedRoleError{
		message:  message,
		youAre:   youAre,
		required: required,
	}
}
