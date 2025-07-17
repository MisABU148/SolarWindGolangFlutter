package model

type Like struct {
	LikerID       int64 `db:"liker_id"`        // тот, кто лайкает
	LikedID       int64 `db:"liked_id"`        // тот, кого лайкают
	IsFirstLikes  bool  `db:"is_first_likes"`  // лайкнул первый
	IsSecondLikes *bool `db:"is_second_likes"` // лайкнул второй
}
