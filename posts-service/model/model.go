package model

type Post struct {
	Id       int     `json:"id"`
	UserId   int     `json:"userId"`
	Created  string  `json:"created"`
	Tags     string  `json:"tags"`
	Content  string  `json:"content"`
	AtchType string  `json:"atchType"`
	AtchId   int     `json:"atchId"`
	AtchUrl  string  `json:"atchUrl"`
	Reaction []int32 `json:"reaction"`
	CmtCount int     `json:"cmtCount"`
}

type Comment struct {
	Id       int        `json:"id"`
	UserId   int        `json:"userId"`
	PostId   int        `json:"postId"`
	ParentId int        `json:"parentId"`
	Content  string     `json:"content"`
	Created  string     `json:"created"`
	Children []*Comment `json:"children,omitempty"`
}

type Reaction struct {
	UserId int    `json:"userId"`
	PostId int    `json:"postId"`
	T      string `json:"type"`
}

type Album struct {
	Id      int    `json:"id"`
	UserId  int    `json:"userId"`
	Descr   string `json:"descr"`
	Created string `json:"created"`
}

type Photo struct {
	Id      int    `json:"id"`
	UserId  int    `json:"userId"`
	AlbumId int    `json:"albumId"`
	Url     string `json:"url"`
	Created string `json:"created"`
}
