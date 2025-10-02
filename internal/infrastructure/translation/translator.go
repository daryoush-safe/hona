package translation

import (
	"log"

	"github.com/go-playground/locales/en_US"
	"github.com/go-playground/locales/fa_IR"
	ut "github.com/go-playground/universal-translator"
)

var translationMap = make(map[string]map[string]string)

func GetTranslator(locale string) ut.Translator {
	universalTranslator := createUniversalTranslator()

	loadAndAddTranslations(universalTranslator)

	translator, found := universalTranslator.GetTranslator(locale)
	if !found {
		translator, _ = universalTranslator.GetTranslator("fa_IR")
	}

	return translator
}

func createUniversalTranslator() *ut.UniversalTranslator {
	en := en_US.New()
	fa := fa_IR.New()
	return ut.New(fa, en, fa)
}

func loadAndAddTranslations(universalTranslator *ut.UniversalTranslator) {
	addTranslations("fa_IR", Persian, universalTranslator)
	addTranslations("en_US", English, universalTranslator)
}

func addTranslations(key string, translations map[string]interface{}, universalTranslator *ut.UniversalTranslator) {
	translator, found := universalTranslator.GetTranslator(key)
	if !found {
		log.Println("translator not found")
		return
	}
	flattenedTranslations := loadTranslation(key, translations)

	for key, translation := range flattenedTranslations {
		translator.Add(key, translation, true)
	}
}

func loadTranslation(key string, translations map[string]interface{}) map[string]string {
	if translations, found := translationMap[key]; found {
		return translations
	}
	flattenedTranslations := make(map[string]string)
	flattenMap("", translations, flattenedTranslations)
	return flattenedTranslations
}

func flattenMap(prefix string, input map[string]interface{}, output map[string]string) {
	for key, value := range input {
		fullKey := key
		if prefix != "" {
			fullKey = prefix + "." + key
		}
		switch v := value.(type) {
		case map[string]interface{}:
			flattenMap(fullKey, v, output)
		case string:
			output[fullKey] = v
		}
	}
}
