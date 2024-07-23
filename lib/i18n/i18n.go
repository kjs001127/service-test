package i18n

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"

	goI18n "github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

const LocalizerKey = "localizer"

var DefaultLanguage = language.English

var i18nKeyRegex = regexp.MustCompile(`^\{.*\}$`)

type I18n interface {
	GetLocalizer(string) *goI18n.Localizer
}

func NewI18nImpl() *I18nImpl {
	bundle := goI18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.LoadMessageFile("./i18n/en.json")
	bundle.LoadMessageFile("./i18n/ko.json")
	bundle.LoadMessageFile("./i18n/ja.json")

	return &I18nImpl{bundle: bundle}
}

type I18nImpl struct {
	bundle *goI18n.Bundle
}

func IsValid(lang string) bool {
	switch lang {
	case language.Korean.String(), language.English.String(), language.Japanese.String():
		return true
	default:
		return false
	}
}

func (s *I18nImpl) GetLocalizer(lang string) *goI18n.Localizer {
	if !IsValid(lang) {
		return goI18n.NewLocalizer(s.bundle, language.English.String())
	}
	return goI18n.NewLocalizer(s.bundle, lang)
}

func Translate(context context.Context, key string, args ...any) string {
	if !i18nKeyRegex.MatchString(key) {
		return key
	}

	rawLocalizer := context.Value(LocalizerKey)
	if rawLocalizer == nil {
		return key
	}

	localizer := rawLocalizer.(*goI18n.Localizer)

	message, err := localizer.Localize(&goI18n.LocalizeConfig{MessageID: key[1 : len(key)-1]})

	if err != nil {
		return key
	}
	return fmt.Sprintf(message, args...)
}
