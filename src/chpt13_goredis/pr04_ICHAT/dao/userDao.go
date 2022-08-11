package dao

import (
	"encoding/json"
	"fmt"

	"github.com/gomodule/redigo/redis"
	"github.com/gomodule/redigo/redis/src/chpt13_goredis/pr04_ICHAT/model"
)

// 获取 Id
func GetId(conn redis.Conn) (int, error) {

	id, errDo := redis.Int(conn.Do("hget", "userdb", 0))
	if errDo != nil {
		if errDo == redis.ErrNil {
			fmt.Println("[错误]:Id获取错误.")
			errDo = model.ErrorUserdbNil
		} else {
			fmt.Println(errDo)
		}
	}
	return id, errDo
}

// 重置 Id
func ResetId(conn redis.Conn, id int) (err error) {

	_, err = conn.Do("Hset", "userdb", "0", id)
	if err != nil {
		fmt.Println("[错误]:Id重置错误.")
		fmt.Println(err)
	}
	return
}

// 获取 user
func GetUserByName(conn redis.Conn, name string) (user *model.User, err error) {

	defer conn.Close()
	id, errGet := GetId(conn)
	if id == 0 || errGet != nil {
		err = model.ErrorUserdbNil
		return
	}

	users, errDo := redis.StringMap(conn.Do("HGETALL", "userdb"))
	if errDo != nil {
		if errDo == redis.ErrNil {
			err = model.ErrorUserdbNil
		}
		return
	}

	user = &model.User{}
	for key, val := range users {
		if key != "0" {
			errJson := json.Unmarshal([]byte(val), user)
			if errJson != nil {
				fmt.Println("[错误]:用户数据解析错误.")
				fmt.Println(errJson)
				err = errJson
				return
			}

			if user.Name == name {
				return
			}
		}
	}
	user = nil
	return
}

// 判断 user 是否存在
func UserIsExist(conn redis.Conn, name string) (isExist bool, err error) {
	defer conn.Close()

	users, errDo := redis.StringMap(conn.Do("HGETALL", "userdb"))
	if errDo != nil {
		if errDo == redis.ErrNil {
			err = model.ErrorUserdbNil
		}
		return
	}
	var user = model.User{}
	for key, val := range users {
		if key != "0" {
			errJson := json.Unmarshal([]byte(val), &user)
			if errJson != nil {
				fmt.Println("[错误]:用户数据解析错误.")
				fmt.Println(errJson)
				err = errJson
				continue
			}

			if user.Name == name {
				isExist = true
				return
			}
		}
	}
	return
}

// 写存 user
func SetUser(conn redis.Conn, user model.User) (err error) {

	id, err := GetId(conn)
	if err != nil {
		fmt.Println("[错误]:获取初始Id错误.")
		return
	}

	id++
	user.Id = id

	data, errJson := json.Marshal(user)
	if errJson != nil {
		fmt.Println("[错误]:用户数据序列化错误.")
		fmt.Println(errJson)
		err = errJson
		return
	}

	_, err = conn.Do("Hset", "userdb", id, data)
	if err != nil {
		fmt.Println("[错误]:用户写存错误.")
		return
	}
	err = ResetId(conn, id)

	return
}
