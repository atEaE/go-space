package component

import "github.com/yohamta/donburi"

var (
	PlayerTag = donburi.NewTag() // プレイヤーエンティティ識別用タグ
	EnemyTag  = donburi.NewTag() // 敵エンティティ識別用タグ
	BulletTag = donburi.NewTag() // 弾エンティティ識別用タグ
	GemTag    = donburi.NewTag() // 経験値ジェム識別用タグ
)
