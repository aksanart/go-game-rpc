package entities

type Enemy struct {
	*Sprite
	IsFollowPlayer bool
	Damage         int
}
