package res

import (
	"net/http"
	"user/pkg/e"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status uint        `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Err    string      `json:"err"`
}

type DataList struct {
	Item  interface{} `json:"item"`
	Total uint64      `json:"total"`
}

type TokenData struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}

func Ok(ctx *gin.Context, msgCode uint, data interface{}) {
	ctx.JSON(http.StatusOK, ginH(msgCode, data))
}

func Unauthorized(ctx *gin.Context, msgCode uint) {
	ctx.JSON(http.StatusUnauthorized, ginH(msgCode, nil))
}

func InternalError(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, ginH(e.Error, nil))
}

func ForbiddenError(ctx *gin.Context, msgCode uint) {
	ctx.JSON(http.StatusForbidden, ginH(msgCode, nil))
}

func Error(ctx *gin.Context, msgCode uint) {
	ctx.JSON(http.StatusInternalServerError, ginH(msgCode, nil))
}

func ginH(msgCode uint, data interface{}) gin.H {
	return gin.H{
		"code": msgCode,
		"msg":  e.GetMsg(msgCode),
		"data": data,
	}
}
