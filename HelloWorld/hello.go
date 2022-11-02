package HelloWorld

const englishHelloPrefix = "Hello, "
const chineseHelloPrefix = "你好, "
const defaultEnglishName = "World"
const defaultChineseName = "世界"
const chinese = "Chinese"

func Hello(name string, language string) string {
	if name == "" {
		switch language {
		case chinese:
			name = defaultChineseName
		default:
			name = defaultEnglishName
		}
	}
	return changeLanguage(language) + name
}

func changeLanguage(language string) (prefix string) {
	switch language {
	case chinese:
		prefix = chineseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
