/*
 * @Author: leon 714765233@qq.com
 * @Date: 2024/7/17 18:59:00
 * @File: types.go
 * @Description:
 *
 * Copyright (c) 2024 by leon email:714765233@qq.com, All Rights Reserved.
 */
package store

type GetContentReq struct {
	Key string `json:"key"`
}

type GetContentRes struct {
	Content string `json:"content"`
}
