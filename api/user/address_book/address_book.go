/*
 * @Author: weihua hu
 * @Date: 2024-12-26 20:26:18
 * @LastEditTime: 2024-12-28 00:01:15
 * @LastEditors: weihua hu
 * @Description:
 */
package address_book

import (
	"net/http"
	"sky-take-out-go/common/result"
	"sky-take-out-go/middlewares"
	"sky-take-out-go/models/entity"
	service "sky-take-out-go/service/user/address_book"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Create(ctx *gin.Context) {
	var addressBook entity.AddressBook
	if err := ctx.ShouldBindJSON(&addressBook); err != nil {
		ctx.JSON(http.StatusOK, result.Error(err.Error()))
		return
	}
	addressBook.UserId = middlewares.GetCurrentUserID(ctx)
	if err := service.Create(addressBook); err != nil {
		ctx.JSON(http.StatusOK, result.Error(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, result.Success())
}

func Update(ctx *gin.Context) {
	var addressBook entity.AddressBook
	if err := ctx.ShouldBindJSON(&addressBook); err != nil {
		ctx.JSON(http.StatusOK, result.Error(err.Error()))
		return
	}
	addressBook.UserId = middlewares.GetCurrentUserID(ctx)

	if err := service.Update(addressBook); err != nil {
		ctx.JSON(http.StatusOK, result.Error(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, result.Success())
}

func Delete(ctx *gin.Context) {
	idParam := ctx.Query("id")
	id, _ := strconv.ParseInt(idParam, 10, 64)

	if err := service.DeleteById(id); err != nil {
		ctx.JSON(http.StatusOK, result.Error(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, result.Success())

}

func SearchDefault(ctx *gin.Context) {

}

func GetById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, _ := strconv.ParseInt(idParam, 10, 64)
	addressBook, err := service.GetById(id)
	if err != nil {
		ctx.JSON(http.StatusOK, result.Error(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, result.Success(addressBook))

}

func SetDefault(ctx *gin.Context) {
	type SetDefaultDTO struct {
		ID int64 `json:"id"`
	}
	var dto SetDefaultDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusOK, result.Error(err.Error()))
		return
	}
	userId := middlewares.GetCurrentUserID(ctx)

	if err := service.SetDefault(dto.ID, userId); err != nil {
		ctx.JSON(http.StatusOK, result.Error(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, result.Success())

}

func GetDefault(ctx *gin.Context) {
	userId := middlewares.GetCurrentUserID(ctx)
	addressBook, err := service.GetDefault(userId)
	if err != nil {
		ctx.JSON(http.StatusOK, result.Error(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, result.Success(addressBook))
}

func List(ctx *gin.Context) {
	userId := middlewares.GetCurrentUserID(ctx)
	addressBooks, err := service.List(userId)
	if err != nil {
		ctx.JSON(http.StatusOK, result.Error(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, result.Success(addressBooks))
}
