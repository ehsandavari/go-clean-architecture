package middlewares

import (
	"github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func I18n() gin.HandlerFunc {
	path, err := os.Getwd()
	if err != nil {
		log.Fatalln("error in os.Getwd in I18n middleware", err)
	}
	return i18n.Localize(i18n.WithBundle(&i18n.BundleCfg{
		RootPath:         path + "/presentation/api/localize",
		AcceptLanguage:   []language.Tag{language.English, language.Persian},
		DefaultLanguage:  language.Persian,
		UnmarshalFunc:    yaml.Unmarshal,
		FormatBundleFile: "yaml",
	}))
}
