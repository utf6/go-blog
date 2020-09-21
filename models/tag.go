package models

import "github.com/jinzhu/gorm"

type Tag struct {
	Model

	Name string `json:"name"`
	CreatedBy string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State int `json:"state"`
}

/**
获取标签
 */
func GetTags(pageNum int, pageSize int, maps interface{}) ([]Tag, error)  {
	var (
		tags []Tag
		err error
	)

	err = db.Where(maps).Find(&tags).Error
	
	if pageSize > 0 && pageNum > 0 {
		err = db.Where(maps).Find(&tags).Offset(pageNum).Limit(pageSize).Error
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return tags, nil
}

/**
获取标签总数
 */
func GetTagTotal(maps interface{}) (int, error)  {
	var count int
	if err := db.Model(&Tag{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

/**
判断标签名是否存在
 */
func ExistTagByName(name string) (bool, error)  {
	var tag Tag

	err := db.Select("id").Where("name = ? AND deleted_on = ?", name, 0).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if tag.ID > 0 {
		return true, nil
	}

	return false, nil
}

/**
添加标签
 */
func AddTag(name string, state int, createdBy string) error {
	tag := Tag{
		Name: name,
		State: state,
		CreatedBy: createdBy,
	}

	if err := db.Create(&tag).Error; err != nil {
		return err
	}

	return nil
}

/**
通过id判断用户是否存在
 */
func ExistTagByID(id int, data interface{}) error {
	if err := db.Model(&Tag{}).Where("id = ? AND deleted_no = ?", id, 0).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

/**
删除标签
 */
func DeleteTag(id int) (bool, error) {
	if err := db.Where("id = ?", id).Delete(&Tag{}).Error; err != nil {
		return false,err
	}

	return true, nil
}

/**
编辑标签
 */
func EditTag(id int, data interface{}) (bool, error) {
	if err := db.Model(&Tag{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return false,err
	}

	return true, nil
}

/**
清除所有tag
 */
func CleanAllTag() (bool, error) {
	if err := db.Unscoped().Where("delete_on != ?", 0).Delete(&Tag{}).Error; err != nil {
		return false, err
	}

	return true, nil
}