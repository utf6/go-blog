package models

type Tag struct {
	Model

	Name string `json:"name"`
	CreatedBy string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State int `json:"state"`
}

//func (tag *Tag) BeforeCreate (scope *gorm.Scope) error {
//	scope.SetColumn("CreatedOn", time.Now().Unix())
//
//	return nil
//}
//
//func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
//	scope.SetColumn("ModifiedOn", time.Now().Unix())
//
//	return nil
//}

/**
获取标签
 */
func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag)  {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)

	return
}

/**
获取标签总数
 */
func GetTagTotal(maps interface{}) (count int)  {
	db.Model(&Tag{}).Where(maps).Count(&count)

	return
}

/**
判断标签名是否存在
 */
func ExistTagByName(name string) bool  {
	var tag Tag
	db.Select("id").Where("name=?", name).First(&tag)

	if tag.ID > 0 {
		return true
	}

	return false
}

/**
添加标签
 */
func AddTag(name string, state int, createdBy string) bool {
	db.Create(&Tag{
		Name: name,
		State: state,
		CreatedBy: createdBy,
	})

	return true
}

/**
通过id判断用户是否存在
 */
func ExistTagByID(id int) bool {
	var tag Tag
	db.Select("id").Where("id = ?", id).First(&tag)

	if tag.ID > 0 {
		return true
	}

	return false
}

/**
删除标签
 */
func DeleteTag(id int) bool {
	db.Where("id = ?", id).Delete(&Tag{})

	return true
}

/**
编辑标签
 */
func EditTag(id int, data interface{}) bool {
	db.Model(&Tag{}).Where("id = ?", id).Updates(data)

	return true
}

/**
清除所有tag
 */
func CleanAllTag() bool {
	db.Unscoped().Where("delete_on != ?", 0).Delete(&Tag{})

	return true
}