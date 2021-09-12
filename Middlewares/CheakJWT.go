package Middlewares

import (
	"KeTangPai/services/DC/JWT"
	"KeTangPai/services/DC/KetangpaiDB"
	"KeTangPai/services/DC/UserCenter"
	"KeTangPai/services/RBAC"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

//检查JWT是否被修改，如果一切顺利会把用户名，用户ID以及用户注册用的邮箱放入上下文

//读取
func CheakJWT(
	jwt JWT.JWTClient,
	uc UserCenter.UserCenterClient,
	k KetangpaiDB.KetangpaiDBClient,
	r RBAC.RBACClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")//找token
		if token == ""{//没有token
			c.JSON(http.StatusBadRequest,gin.H{
				"error":"token == nil",
			})
			c.Abort()
			return
		}
		//检查token
		ctx,_:=context.WithTimeout(context.Background(), serviceTimeLimit)
		u,err:=jwt.CheckToken(ctx,&JWT.Token{Content: token})//检查传入的token
		if err!=nil {//对方出问题
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":err.Error(),
			})
			c.Abort()
			return
		}
		//获取用户Email
		ctx,_=context.WithTimeout(context.Background(), serviceTimeLimit)
		user,err:=uc.GetUserInfo(ctx,&UserCenter.Id{I: u.Id})//获取用户的相关信息
		if err!=nil {
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":err.Error(),
			})
			c.Abort()
			return
		}

		//获取user在本产品下的信息
		ctx,_=context.WithTimeout(context.Background(), serviceTimeLimit)
		kus,err:=k.GetUserInfo(ctx,&KetangpaiDB.User{Id:user.Id})
		if err!=nil {
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":err.Error(),
			})
			c.Abort()
			return
		}

		if kus.Uid<1 {
			ctx,_=context.WithTimeout(context.Background(), serviceTimeLimit)
			kus,err=k.CreateUser(ctx,&KetangpaiDB.User{Id:user.Id})//没有记录就创造记录
		}

		if err!=nil {
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":err.Error(),
			})
			c.Abort()
			return
		}

		//获取用户角色
		ctx,_=context.WithTimeout(context.Background(), serviceTimeLimit)
		role,err:=r.GetRole(ctx,&RBAC.Uids{Uids: []uint32{kus.Uid}})
		if err!=nil {
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":err.Error(),
			})
			c.Abort()
			return
		}

		//获取班级名
		ctx,_=context.WithTimeout(context.Background(), serviceTimeLimit)
		classname,err:=k.GetClassName(ctx,&KetangpaiDB.Classid{Classid: kus.Classid})
		if err!=nil {
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":err.Error(),
			})
			c.Abort()
			return
		}

		c.Set("id",user.Id)
		c.Set("role",role.Roles[0])
		c.Set("name",kus.Name)//存入用户名
		c.Set("uid",kus.Uid)//存入UID
		c.Set("email",user.Email)//存入邮箱
		c.Set("classid",kus.Classid)//存入classID
		c.Set("classname",classname.Name)
	}
}