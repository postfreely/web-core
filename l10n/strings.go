package l10n

// Strings returns a translation set that will take any term and return its
// translation.
func Strings(lang string) map[string]string {
	switch lang {
	case "fr":
		return phrasesFR
	case "hu":
		return phrasesHU
	case "it":
		return phrasesIT
	case "ro":
		return phrasesRO
	default:
		return phrases
	}
}
