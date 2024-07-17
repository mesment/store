/*
 * @Author: leon 714765233@qq.com
 * @Date: 2024/7/17 18:56:00
 * @File: store.go
 * @Description:
 *
 * Copyright (c) 2024 by leon email:714765233@qq.com, All Rights Reserved.
 */
package store

import (
	"context"
	"github.com/gogf/gf/v2/container/gvar"
	"time"
)

type OssManager interface {
	// GetObject 从oss 获取文件
	GetObject(ctx context.Context, file string) (b []byte, err error)
}

type Cache interface {
	// Get 当`timeout > 0`且缓存命中时，设置/重置`key`的过期时间
	Get(ctx context.Context, key string, timeout ...time.Duration) (val *gvar.Var, err error)
	// Set 当`timeout > 0`时，设置/重置`key`的过期时间
	Set(ctx context.Context, key string, val any, timeout ...time.Duration) (err error)
}

type Store struct {
	cacheAdapter Cache      // 缓存适配器
	Oss          OssManager // oss下载管理器
}

func New(cache Cache, ossManager OssManager) *Store {
	s := &Store{
		cacheAdapter: cache,
		Oss:          ossManager,
	}
	return s
}

func (s *Store) GetContent(ctx context.Context, req *GetContentReq) (content string, err error) {
	var value *gvar.Var
	var b []byte
	value, err = s.cacheAdapter.Get(ctx, req.Key)
	if err != nil {
		return
	}
	if !value.IsNil() {
		content = value.String()
		return
	}
	b, err = s.Oss.GetObject(ctx, req.Key)
	if err != nil {
		return
	}
	content = string(b)
	return
}
