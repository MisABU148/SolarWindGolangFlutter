package model

type LikeDTO struct {
	LikerID int64 `json:"liker_id"` // из заголовка
	LikedID int64 `json:"liked_id"` // из тела запроса
	IsLiker bool  `json:"is_liker"` // true = like, false = dislike
	IsLiked bool  `json:"is_liked"` // true = like, false = dislike
}

type LikesCompositePrimaryKey struct {
	LikerID int64
	LikedID int64
}

type LikesEntity struct {
	ID            LikesCompositePrimaryKey
	IsFirstLikes  *bool // nullable boolean
	IsSecondLikes *bool
}
