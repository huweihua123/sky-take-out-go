/*
 * @Author: weihua hu
 * @Date: 2024-12-20 00:27:00
 * @LastEditTime: 2024-12-28 00:01:46
 * @LastEditors: weihua hu
 * @Description:
 */
package dish

import (
	"net/http"
	"sky-take-out-go/common/result"
	"sky-take-out-go/models/dto"
	service "sky-take-out-go/service/admin/dish"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func Create(ctx *gin.Context) {
	var dishDTO dto.DishDTO
	if err := ctx.ShouldBindJSON(&dishDTO); err != nil {
		ctx.JSON(http.StatusOK, result.Error(err.Error()))
		return
	}
	err := service.Create(dishDTO)
	if err != nil {
		ctx.JSON(http.StatusOK, result.Error(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, result.Success())
}

func Update(ctx *gin.Context) {
	var dishDTO dto.DishDTO
	if err := ctx.ShouldBindJSON(&dishDTO); err != nil {
		ctx.JSON(http.StatusOK, result.Error(err.Error()))
		return
	}
	err := service.Update(dishDTO)
	if err != nil {
		ctx.JSON(http.StatusOK, result.Error(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, result.Success())
}

func StartOrStop(ctx *gin.Context) {
	statusParam := ctx.Param("status")
	idParam := ctx.Query("id")
	status, _ := strconv.Atoi(statusParam)
	id, _ := strconv.ParseInt(idParam, 10, 64)
	err := service.StartOrStop(status, id)
	if err != nil {
		ctx.JSON(http.StatusOK, result.Error(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, result.Success())
}

func GetById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, _ := strconv.ParseInt(idParam, 10, 64)
	dish, err := service.GetById(id)
	if err != nil {
		ctx.JSON(http.StatusOK, result.Error(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, result.Success(dish))
}

func Delete(ctx *gin.Context) {
	// 获取ids参数
	idsParam := ctx.Query("ids")
	if idsParam == "" {
		ctx.JSON(http.StatusOK, result.Error("菜品id不能为空"))
		return
	}

	// 将ids字符串转换为[]int64
	idStrs := strings.Split(idsParam, ",")
	ids := make([]int64, 0, len(idStrs))
	for _, idStr := range idStrs {
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusOK, result.Error("无效的菜品id"))
			return
		}
		ids = append(ids, id)
	}

	// 调用service层批量删除
	if err := service.BatchDelete(ids); err != nil {
		ctx.JSON(http.StatusOK, result.Error(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, result.Success())
}
