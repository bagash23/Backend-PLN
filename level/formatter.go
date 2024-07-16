package level

type LevelFormatter struct {
	IDLevel string `json:"id_level"`
	Level string `json:"level"`
}

func FormatterLevel(level Level) LevelFormatter {
	formatter := LevelFormatter {
		IDLevel: level.IDLevel,
		Level: level.Level,
	}
	return formatter
}