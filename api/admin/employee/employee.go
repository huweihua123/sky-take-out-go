/*
 * @Author: weihua hu
 * @Date: 2024-12-15 15:06:45
 * @LastEditTime: 2024-12-28 00:01:51
 * @LastEditors: weihua hu
 * @Description:
 */
package employee

import (
	"net/http"
	"sky-take-out-go/common/result"
	"sky-take-out-go/middlewares"
	"sky-take-out-go/models"
	"sky-take-out-go/models/dto"
	"sky-take-out-go/models/vo"
	service "sky-take-out-go/service/admin/employee"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {

	var employeeLoginDTO dto.EmployeeLoginDTO

	// 绑定并校验 JSON 数据
	if err := ctx.ShouldBindJSON(&employeeLoginDTO); err != nil {
		ctx.JSON(http.StatusOK, result.Error(err.Error()))
		return
	}
	// 调用service层的Login方法
	employee, err := service.Login(employeeLoginDTO)
	if err != nil {
		ctx.JSON(http.StatusOK, result.Error(err.Error()))
		return
	}
	// 生成JWT令牌
	j := middlewares.NewAdminJWT()
	claims := models.CustomClaims{
		ID:       employee.ID,
		NickName: employee.Username,
		Role:     "admin",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(3 * time.Hour).Unix(),
			Issuer:    "sky-take-out-go",
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		ctx.JSON(http.StatusOK, result.Error("生成令牌失败"))
		return
	}

	employeeLoginVO := vo.EmployeeLoginVO{
		ID:       employee.ID,
		Username: employee.Username,
		Name:     employee.Name,
		Token:    token,
	}

	// 返回成功响应
	ctx.JSON(http.StatusOK, result.Success(employeeLoginVO))

}

func Save(ctx *gin.Context) {
	var employeeDTO dto.EmployeeDTO
	if err := ctx.ShouldBindJSON(&employeeDTO); err != nil {
		ctx.JSON(http.StatusOK, result.Error(err.Error()))
		return
	}

	err := service.Save(employeeDTO)

	if err != nil {
		ctx.JSON(http.StatusOK, result.Error(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, result.Success())
}

func StartOrStop(ctx *gin.Context) {
	statusParam := ctx.Param("status")
	status, _ := strconv.Atoi(statusParam)
	idParam := ctx.Query("id")
	id, _ := strconv.ParseInt(idParam, 10, 64)

	err := service.StartOrStop(status, id)
	if err != nil {
		ctx.JSON(http.StatusOK, result.Error(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, result.Success())
}

func Update(ctx *gin.Context) {
	var employeeDTO dto.EmployeeDTO
	if err := ctx.ShouldBindJSON(&employeeDTO); err != nil {
		ctx.JSON(http.StatusOK, result.Error(err.Error()))
		return
	}
	err := service.Update(employeeDTO)
	if err != nil {
		ctx.JSON(http.StatusOK, result.Error(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, result.Success())
}

/**
 * @Author: weihua hu
 * @description: 通过id获取员工信息
 * @return {*}
 */
func GetById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, _ := strconv.ParseInt(idParam, 10, 64)
	employee, err := service.GetById(id)
	if err != nil {
		ctx.JSON(http.StatusOK, result.Error(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, result.Success(employee))
}

func Page(ctx *gin.Context) {
	var pageDTO dto.EmployeePageQueryDTO
	if err := ctx.ShouldBindQuery(&pageDTO); err != nil {
		ctx.JSON(http.StatusOK, result.Error(err.Error()))
		return
	}
	employees, err := service.PageQuery(pageDTO)
	if err != nil {
		ctx.JSON(http.StatusOK, result.Error(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, result.Success(employees))
}
