package cache_service

import (
	"github.com/utf6/go-blog/pkg/e"
	"strconv"
	"strings"
)

type Article struct {
	ID		int
	TagID	int
	State	int

	PageNum		int
	PageSize	int
}

func (a *Article) GetArticleKey() string {
	return e.CACHE_ARTICLE + "_" + strconv.Itoa(a.ID)
}

func (a *Article) GetArticlesKey() string {
	keys := []string{
		e.CACHE_ARTICLE,
		"LIST",
	}

	if a.ID > 0 {
		keys = append(keys, strconv.Itoa(a.ID))
	}

	if a.TagID > 0 {
		keys = append(keys, strconv.Itoa(a.ID))
	}

	if a.State > 0 {
		keys = append(keys, strconv.Itoa(a.ID))
	}

	if a.PageNum > 0 {
		keys = append(keys, strconv.Itoa(a.ID))
	}

	if a.PageSize > 0 {
		keys = append(keys, strconv.Itoa(a.ID))
	}

	return strings.Join(keys, "_")
}
