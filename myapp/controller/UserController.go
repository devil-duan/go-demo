package controller

import (
	"myapp/model"

	"github.com/kataras/iris/v12"

	"fmt"
)

// 查询多条记录
func List(ctx iris.Context) {

	sqlStr := "select id, name from user"
	rows, err := model.DB.Query(sqlStr)

	users := []model.User{}

	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	// 关闭rows释放持有的数据库链接
	defer rows.Close()
	// 循环读取结果集中的数据
	for rows.Next() {
		var u model.User
		err := rows.Scan(&u.Id, &u.Name)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s\n", u.Id, u.Name)
		users = append(users, u)
	}

	ctx.JSON(users)
}

//单条新增
func Create(ctx iris.Context) {

	var user model.User

	ctx.ReadJSON(&user)
	fmt.Printf("user.Name:%s\n", user.Name)
	sqlStr := "insert into user (name) values (?)"
	ret, err := model.DB.Exec(sqlStr, user.Name)

	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	// 新插入数据的id
	theID, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)

	ctx.StatusCode(iris.StatusCreated)

}
//单条删除
func DelOne(ctx iris.Context) {
	id := ctx.Params().GetInt64Default("id", 3)
	
	sqlStr := "delete from user where id =?"
	ret,err := model.DB.Exec(sqlStr,id)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)

}

// 查询单条数据
func QueryOne(ctx iris.Context) {
	id := ctx.Params().GetInt64Default("id", 3)
	var u model.User
	fmt.Printf("id: %d\n", id)
	sqlStr := "select id, name from user where id=?"
	err := model.DB.QueryRow(sqlStr, id).Scan(&u.Id, &u.Name)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s\n", u.Id, u.Name)
	ctx.JSON(u)
}
