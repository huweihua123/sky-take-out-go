/*
 * @Author: weihua hu
 * @Date: 2024-12-20 00:17:25
 * @LastEditTime: 2024-12-20 18:30:29
 * @LastEditors: weihua hu
 * @Description:
 */
package category

import (
	"net/http"
	"sky-take-out-go/common/result"
	"sky-take-out-go/models/dto"
	service "sky-take-out-go/service/category"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Create(ctx *gin.Context) {
	var categoryDTO dto.CategoryDTO
	if err := ctx.ShouldBindJSON(&categoryDTO); err != nil {
		ctx.JSON(http.StatusOK, result.Error(err.Error()))
		return
	}

	// 调用service层的Create方法
	if err := service.Create(categoryDTO); err != nil {
		ctx.JSON(http.StatusOK, result.Error(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, result.Success())

}

func Page(ctx *gin.Context) {
	var pageDTO dto.CategoryPageQueryDTO
	if err := ctx.ShouldBindQuery(&pageDTO); err != nil {
		ctx.JSON(http.StatusOK, result.Error(err.Error()))
		return
	}

	// 调用service层的Page方法
	pageVO, err := service.PageQuery(pageDTO)
	if err != nil {
		ctx.JSON(http.StatusOK, result.Error(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, result.Success(pageVO))
}

func StartOrStop(ctx *gin.Context) {
	statusParam := ctx.Param("status")
	idParam := ctx.Query("id")
	status, _ := strconv.Atoi(statusParam)
	id, _ := strconv.ParseInt(idParam, 10, 64)
	// 调用service层的StartOrStop方法
	if err := service.StartOrStop(status, id); err != nil {
		ctx.JSON(http.StatusOK, result.Error(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, result.Success())
}

func Update(ctx *gin.Context) {
	var categoryDTO dto.CategoryDTO
	if err := ctx.ShouldBindJSON(&categoryDTO); err != nil {
		ctx.JSON(http.StatusOK, result.Error(err.Error()))
		return
	}

	// 调用service层的Update方法
	if err := service.Update(categoryDTO); err != nil {
		ctx.JSON(http.StatusOK, result.Error(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, result.Success())
}

func Delete(ctx *gin.Context) {
	idParam := ctx.Query("id")
	id, _ := strconv.ParseInt(idParam, 10, 64)
	// 调用service层的Delete方法
	if err := service.Delete(id); err != nil {
		ctx.JSON(http.StatusOK, result.Error(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, result.Success())
}

func List(ctx *gin.Context) {
	typeParam := ctx.Query("type")
	var typet *int

	if typeParam != "" {
		val, err := strconv.Atoi(typeParam)
		if err != nil {
			ctx.JSON(http.StatusOK, result.Error("无效的type参数"))
			return
		}
		typet = &val
	}

	categorys, err := service.List(typet)
	if err != nil {
		ctx.JSON(http.StatusOK, result.Error(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, result.Success(categorys))
}
