package userModel

import (
	"sauth/db"
	"sauth/util"
	"database/sql"
)

func Insert(model *User) (sql.Result, error) {
	sql := `INSERT INTO user
			(uuid, user_name, password, nick_name, sex, position, tel, mail, address, remark, create_user, create_time, update_user, update_time) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	res, err := db.Engine.Exec(sql,
		util.GetUuid(), model.UserName, model.Password, model.NickName, model.Sex, model.Position, model.Tel, model.Mail, model.Address, model.Remark, model.CreateUser, util.CurrentTime(), model.UpdateUser, util.CurrentTime())
	return res, err
}

func FindCountByUserNameAndPassword(userName string, password string) ([]map[string][]byte, error) {
	sql := `SELECT COUNT(*) AS count FROM user
			where user_name = ? and password = ?`
	res, err := db.Engine.Query(sql,
		userName, password)
	return res, err
}

func FindCountByUserName(userName string) ([]map[string][]byte, error) {
	sql := `SELECT COUNT(*) AS count FROM user
			where user_name = ?`
	res, err := db.Engine.Query(sql, userName)
	return res, err
}
