package domain

import (
	"github.com/friendsofgo/errors"

	app "github.com/channel-io/ch-app-store/internal/app/domain"
)

type AutoCompleteRequest struct {
	Command   Key
	ChannelID string
	app.Body[AutoCompleteArgs]
}

type AutoCompleteArgs []AutoCompleteArg
type AutoCompleteArg struct {
	Focused bool
	Name    string
	Value   any
}

func (args AutoCompleteArgs) validate() error {
	for _, arg := range args {
		if len(arg.Name) <= 0 || arg.Value == nil {
			return errors.New("name and value must not be empty")
		}
	}
	return nil
}

type Choices []Choice
type Choice struct {
	Name  string
	Value any
}

func (choices Choices) validate() error {
	for _, c := range choices {
		if c.Value == nil || len(c.Name) == 0 {
			return errors.New("name and value of choice must not be empty")
		}
	}
	return nil
}
