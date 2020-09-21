package models

type Article struct {
	Model

	TagID int `json:"tag_id" gorm:"index"`
	Tag   Tag `json:"tag"`

	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CreatedBy     string `json:"created_by"`
	ModifiedBy    string `json:"modified_by"`
	State         int    `json:"state"`
	CoverImageUrl string `json:"cover_image_url"`
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
func EditArticle(id int, data interface{}) error {

	if err := db.Model(&Article{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

/**
添加文章
 */
func AddArticle(data map[string]interface{}) error {
	article := Article{
		TagID: data["tag_id"].(int),
		Title: data["title"].(string),
		Desc: data["desc"].(string),
		Content: data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State: data["state"].(int),
		CoverImageUrl: data["cover_image_url"].(string),
	}

	if err := db.Create(&article).Error; err != nil {
		return  err
	}

	return nil
}

/**
删除文章
 */
func DeleteArticle(id int) error {

	if err := db.Where("id = ?", id).Delete(Article{}).Error; err != nil {
		return err
	}

	return nil
}

/**
删除所有文章
 */
func CleanAllArticle() error {
	if err := db.Unscoped().Where("delete_no != ?", 0).Delete(&Article{}).Error; err != nil {
		return err
	}

	return nil
}

