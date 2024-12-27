/*
 * @Author: weihua hu
 * @Date: 2024-12-22 21:17:36
 * @LastEditTime: 2024-12-25 13:43:44
 * @LastEditors: weihua hu
 * @Description:
 */
package setmeal

import (
	"net/http"
	result "sky-take-out-go/common/result"
	"sky-take-out-go/models/dto"
	service "sky-take-out-go/service/admin/setmeal"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func Create(ctx *gin.Context) {
	var setmealDTO dto.SetMealDTO
	if err := ctx.ShouldBindJSON(&setmealDTO); err != nil {
		ctx.JSON(http.StatusOK, result.Error(err.Error()))
		return
	}
	if err := service.Create(setmealDTO); err != nil {
		ctx.JSON(http.StatusOK, result.Error(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, result.Success())

}

func Update(ctx *gin.Context) {
	var setmealDTO dto.SetMealDTO
	if err := ctx.ShouldBindJSON(&setmealDTO); err != nil {
		ctx.JSON(http.StatusOK, result.Error(err.Error()))
		return
	}
	if err := service.Update(setmealDTO); err != nil {
		ctx.JSON(http.StatusOK, result.Error(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, result.Success())
}

func Delete(ctx *gin.Context) {
	idsParam := ctx.Query("ids")
	if idsParam == "" {
		ctx.JSON(http.StatusOK, result.Error("ids is required"))
		return
	}
	idStrs := strings.Split(idsParam, ",")
	ids := make([]int64, 0, len(idStrs))
	for _, idStr := range idStrs {
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusOK, result.Error(err.Error()))
			return
		}
		ids = append(ids, id)
	}
	if err := service.Delete(ids); err != nil {
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
	if err := service.StartOrStop(id, status); err != nil {
		ctx.JSON(http.StatusOK, result.Error(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, result.Success())
}

func GetById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, _ := strconv.ParseInt(idParam, 10, 64)
	setmealDTO, err := service.GetById(id)
	if err != nil {
		ctx.JSON(http.StatusOK, result.Error(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, result.Success(setmealDTO))
}

func Page(ctx *gin.Context) {
	var setmealPageQueryDTO dto.SetmealPageQueryDTO
	if err := ctx.ShouldBindQuery(&setmealPageQueryDTO); err != nil {
		ctx.JSON(http.StatusOK, result.Error(err.Error()))
		return
	}
	PageResult, err := service.PageQuery(setmealPageQueryDTO)
	if err != nil {
		ctx.JSON(http.StatusOK, result.Error(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, result.Success(PageResult))
}
