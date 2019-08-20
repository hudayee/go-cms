package role

import (
	"cms/helper"
	"cms/models"
	"github.com/kataras/iris"
	"log"
)

func GetById(c iris.Context) {
	//params {id:int}
	id, _ := c.Params().GetInt("id")

	role, err := models.RoleDao.Get(id)
	if err != nil {
		c.JSON(helper.ToWeb(500, err.Error(), nil))
		log.Println(err)
		return
	}

	c.JSON(helper.ToWeb(200, "", role))
}

func GetAll(c iris.Context) {
	var (
		id     = c.URLParamIntDefault("id", 0)
		name   = c.URLParam("name")
		limit  = c.URLParamIntDefault("limit", 10)
		offset = c.URLParamIntDefault("offset", 0)
	)
	query := models.RoleQuery{
		Id:     id,
		Name:   name,
		Limit:  limit,
		Offset: offset,
	}

	roles, count, err := models.RoleDao.GetAll(query)
	if err != nil {
		c.JSON(helper.ToWeb(500, err.Error(), nil))
		log.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["list"] = roles
	data["count"] = count
	c.JSON(helper.ToWeb(200, "", data))
}

func Create(c iris.Context) {
	var role = models.Role{}
	if err := c.ReadJSON(&role); err != nil {
		c.JSON(helper.ToWeb(400, err.Error(), nil))
		log.Println(err)
		return
	}
	if err := models.RoleDao.Create(role); err != nil {
		c.JSON(helper.ToWeb(500, err.Error(), nil))
		log.Println(err)
		return
	}
	c.JSON(helper.ToWeb(200, "", nil))
}

func Update(c iris.Context) {
	//params {id:int}
	id, _ := c.Params().GetInt("id")
	var role = models.Role{}
	if err := c.ReadJSON(&role); err != nil {
		c.JSON(helper.ToWeb(400, err.Error(), nil))
		log.Println(err)
		return
	}
	role.Id = id
	num, err := models.RoleDao.Update(role)
	if err != nil {
		c.JSON(helper.ToWeb(500, err.Error(), nil))
		log.Println(err)
		return
	} else if num == 0 {
		c.JSON(helper.ToWeb(400, "not found", nil))
	}
	c.JSON(helper.ToWeb(200, "", nil))
}
