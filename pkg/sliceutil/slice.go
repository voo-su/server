// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package sliceutil

import (
	"fmt"
	"strconv"
	"strings"
)

type IInt interface {
	uint | uint8 | uint16 | uint32 | uint64 | int | int8 | int16 | int32 | int64
}

type IFloat interface {
	float32 | float64
}

func Include[T IInt | string](find T, arr []T) bool {
	for _, value := range arr {
		if value == find {
			return true
		}
	}

	return false
}

func Unique[T IInt | string](data []T) []T {
	list, hash := make([]T, 0), make(map[T]struct{})
	for _, value := range data {
		if _, ok := hash[value]; !ok {
			list = append(list, value)
			hash[value] = struct{}{}
		}
	}

	return list
}

func Sum[T IInt | IFloat](arr []T) T {
	var count T
	for _, v := range arr {
		count += v
	}

	return count
}

func ToMap[T any, K int | string](arr []T, fn func(T) K) map[K]T {
	var m = make(map[K]T)
	for _, t := range arr {
		m[fn(t)] = t
	}

	return m
}

func ParseIds(str string) []int {
	str = strings.TrimSpace(str)
	ids := make([]int, 0)
	if str == "" {
		return ids
	}
	for _, value := range strings.Split(str, ",") {
		if id, err := strconv.Atoi(value); err == nil {
			ids = append(ids, id)
		}
	}

	return ids
}

func ToIds[T IInt](items []T) string {
	tmp := make([]string, 0, len(items))
	for _, item := range items {
		tmp = append(tmp, fmt.Sprintf("%d", item))
	}

	return strings.Join(tmp, ",")
}
