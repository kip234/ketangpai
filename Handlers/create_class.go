package Handlers

import (
	"KeTangPai/services/DC/KetangpaiDB"
	"KeTangPai/services/RBAC"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)
//新建一个班级-一个老师只能有一个班级，该操作会自动解散原有班级
//学生列表里可能会出现老师自己
func Create_class(u KetangpaiDB.KetangpaiDBClient,r RBAC.RBACClient) gin.HandlerFunc {
	return func(c *gin.Context){
		uid,err:=getUint("uid",c)//获取当前用户ID
		if err!=nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
			return
		}
		class:=KetangpaiDB.Classdb{}
		err=c.ShouldBind(&class)//获取传回来的班级信息
		if err!=nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
			return
		}
		class.Teacher=uid//指明老师
		ctx,_:=context.WithTimeout(context.Background(),serviceTimeLimit)
		re,err:=u.CreateClass(ctx,&KetangpaiDB.Class{
			Classid:class.Classid,
			Teacher:class.Teacher,
			Name:class.Name,
			Students:class.Students,
		})//创建记录
		if err!=nil {
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":err.Error(),
			})
			return
		}

		users:=make([]uint32,len(class.Students)+1)
		roles:=make([]string,len(class.Students)+1)
		index:=0
		for _,i:=range class.Students{
			if i==class.Teacher {
				continue
			}
			users[index]=i
			roles[index]="/student"
			index++
		}
		users[index]=class.Teacher
		roles[index]="/teacher"

		ctx,_=context.WithTimeout(context.Background(),serviceTimeLimit)
		_,err=r.RefreshUserRole(ctx,&RBAC.UsersRoles{Roles: roles,Uids: users})
		if err!=nil {
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK,gin.H{
			"return":re,
		})
	}
}
