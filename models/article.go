package models

import (
	"github.com/jinzhu/gorm"

	"time"
)

type Article struct {
	Model

	TagID int `json:"tag_id" gorm:"index"`
	Tag Tag `json:"tag"`

	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
	CreatedBy string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State int `json:"state"`
}

/**
通过id查询文章
 */
func ExistArticleByID(id int) bool {

	var article Article
	db.Select("id").Where("id = ?", id).First(&article)

	if article.ID > 0 {
		return true
	}
	return false
}

/**
获取文章总数
 */
func GetArticleTotal(maps interface{}) (count int) {
	db.Model(&Article{}).Where(maps).Count(&count)

	return
}

/**
获取文章列表
 */
func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)

	return
}

/**
获取文章，关联标签
 */
func GetArticle(id int) (article Article) {
	db.Where("id = ?", id).First(&article)
	db.Model(&article).Related(&article.Tag)

	return
}

/**
编辑文章
 */
func EditArticle(id int, data interface{}) bool {
	db.Model(&Article{}).Where("id = ?", id).Updates(data)

	return true
}

/**
添加文章
 */
func AddArticle(data map[string]interface{}) bool {
	db.Create(&Article{
		TagID: data["tag_id"].(int),
		Title: data["title"].(string),
		Desc: data["desc"].(string),
		Content: data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State: data["state"].(int),
	})

	return true
}

/**
删除文章
 */
func DeleteArticle(id int) bool {
	db.Where("id = ?", id).Delete(Article{})

	return true
}

/**
创建之前填充 CreatedOn
 */
func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

/**
更新之前填充 ModifiedOn
 */
func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}

/**
删除所有文章
 */
func CleanAllArticle() bool {
	db.Unscoped().Where("delete_no != ?", 0).Delete(&Article{})

	return true
}

