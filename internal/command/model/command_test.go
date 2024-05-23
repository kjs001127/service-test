package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/null/v8"

	"github.com/channel-io/ch-app-store/internal/command/model"
)

func testCommand() *model.Command {
	return &model.Command{
		AppID:       "test12",
		Name:        "testCommand",
		Description: null.StringFrom("한글 설명입니다").Ptr(),
		Scope:       model.ScopeDesk,
		AlfMode:     model.AlfModeDisable,
		NameDescI18NMap: map[string]model.I18nMap{
			"ko": {
				Name:        "한글커맨드12",
				Description: "한글 설명입니다!!",
			},
		},
		ParamDefinitions: []*model.ParamDefinition{
			{
				Name:        "testParam",
				Type:        model.ParamTypeInt,
				Required:    false,
				Description: "한글 설명입니다",
				NameDescI18nMap: map[string]model.ParamDefI18ns{
					"ko": {
						Name:        "한글파라미터12",
						Description: "한글 파라미터 설명",
					},
				},
			},
		},
	}
}

func TestInvalidName(t *testing.T) {
	invalidNames := []string{
		"한글이름",
		"name12",
		"name with space",
		"tooMuchLongLongLongName",
	}

	for _, invalidName := range invalidNames {
		command := testCommand()
		command.Name = invalidName
		assert.Error(t, command.Validate())
	}
}

func TestInvalidI18nName(t *testing.T) {
	invalidI18nNames := []string{
		"name with space",
		"tooMuchLongLongLongName",
		"한글인데 띄어쓰기가 있음",
	}
	for _, invalidName := range invalidI18nNames {
		command := testCommand()
		command.NameDescI18NMap["ko"] = model.I18nMap{
			Name: invalidName,
		}
		assert.Error(t, command.Validate())
	}
}

func TestParamDefinitionName(t *testing.T) {
	invalidNames := []string{
		"한글이름",
		"name with space",
		"name12",
		"tooMuchLongLongLongName",
	}
	for _, invalidName := range invalidNames {
		command := testCommand()
		command.ParamDefinitions[0].Name = model.ParamName(invalidName)

		assert.Error(t, command.Validate())
	}
}

func TestParamDefinitionI18nName(t *testing.T) {
	invalidI18nNames := []string{
		"name with space",
		"tooMuchLongLongLongName",
		"한글인데 띄어쓰기가 있음",
	}
	for _, invalidName := range invalidI18nNames {
		command := testCommand()
		command.ParamDefinitions[0].NameDescI18nMap["ko"] = model.ParamDefI18ns{
			Name: invalidName,
		}
		assert.Error(t, command.Validate())
	}
}

func TestValidCommand(t *testing.T) {
	assert.NoError(t, testCommand().Validate())
}
