package localizer

type (
	TextKey string
)

type Localizer interface {
	SetLanguages(languages ...string)
	GetTextByKey(textKey TextKey) string
	GetTextByKeyWithVariables(textKey TextKey, variables map[string]any) string
	GetDefaultText() string
}
