package level

type LevelInput struct {
	IDLevel string `json:"id_level" binding:"required"`
	Level string `json:"level" binding:"required"`
}