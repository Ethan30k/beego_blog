package models

import "time"

type TagPost struct {
	Id int
	Tagid int //标签id
	Postid int //文章id
	Poststatus int //文章状态
	Posttime time.Time //文章发表时间
}
