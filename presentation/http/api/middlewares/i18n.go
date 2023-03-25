package middlewares

import (
	"github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
)

func I18n() gin.HandlerFunc {
	return i18n.Localize(i18n.WithBundle(&i18n.BundleCfg{
		RootPath:         "./presentation/http/api/localize",
		AcceptLanguage:   []language.Tag{language.English, language.Persian},
		DefaultLanguage:  language.English,
		UnmarshalFunc:    yaml.Unmarshal,
		FormatBundleFile: "yaml",
	}))
}
