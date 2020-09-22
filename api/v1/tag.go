package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"github.com/utf6/go-blog/models"
	"github.com/utf6/go-blog/pkg/e"
	"github.com/utf6/go-blog/pkg/setting"
	"github.com/utf6/go-blog/pkg/util"
	"net/http"
)

/**
获取文章多个标签
 */
func GetTags(c *gin.Context)  {

	name := c.Query("name")
	
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps[name] = name
	}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := e.SUCCESS
	data["lists"] = models.GetArticles(util.GetPage(c), setting.AppSetting.PageSize, maps)
	data["total"] = models.GetArticleTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : data,
	})
}

// @Summary 新增文章标签
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags [post]
func AddTag(c *gin.Context)  {
	name := c.PostForm("name")
	state := com.StrTo(c.PostForm("state")).MustInt()
	createdBy := c.PostForm("created_by")

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	code := e.INVALID_PARAMS
	if ! valid.HasErrors() {
		if models.ExistTagByName(name) {
			code = e.ERROR_EXIST_TAG
		} else {
			code = e.SUCCESS
			models.AddTag(name, state, createdBy)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : make(map[string]string),
	})
}


// @Summary 修改文章标签
// @Produce  json
// @Param name query int true "id"
// @Param state query int true "name"
// @Param created_by query string false "modified_by"
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags/:id [put]
func EditTag(c *gin.Context)  {
	id := com.StrTo(c.Param("id")).MustInt()
	name := c.PostForm("name")
	modifiedBy := c.PostForm("modified_by")

	valid := validation.Validation{}
	var state int = -1

	if arg := c.PostForm("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	valid.Required(id, "id").Message("ID不能为空")
	valid.Required(modifiedBy, "modified_by").Message("修改者不能为空")
	valid.MaxSize(modifiedBy, 25, "modified_by").Message("修改者最长为25字符")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Required(name, "name").Message("名称不能为空")

	//返回验证信息
	code := e.INVALID_PARAMS
	if valid.HasErrors() {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"code" : code,
			"msg" : valid.Error("参数错误"),
		})
	}

	code = e.SUCCESS
	//判断数据是否存在
	if models.ExistTagByID(id) {
		if !models.ExistTagByName(name) {
			data := make(map[string]interface{})
			data["modified_by"] = modifiedBy

			if name != "" {
				data["name"] = name
			}

			if state != -1 {
				data["state"] = state
			}
			//编辑数据
			models.EditTag(id, data)
		} else {
			code = e.ERROR_EXIST_TAG
		}
	} else {
		code = e.ERROR_NOT_EXIST_TAG
	}

	//返回状态
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : make(map[string]string),
	})
}

/**
删除文章标签
 */
func DeleteTag(c *gin.Context)  {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	//返回验证信息
	code := e.INVALID_PARAMS
	if valid.HasErrors() {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"code" : code,
			"msg" : valid.Errors,
		})
	}

	//删除操作
	code = e.ERROR_NOT_EXIST_TAG
	if models.ExistTagByID(id) {
		models.DeleteTag(id)
		code = e.SUCCESS
	}

	//返回信息
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : make(map[string]string),
	})
}