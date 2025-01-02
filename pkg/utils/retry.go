// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package utils

import "time"

func Retry(num int, sleep time.Duration, fn func() error) error {
	var err error
	for i := 0; i < num; i++ {
		if err = fn(); err == nil {
			return nil
		}
		time.Sleep(sleep)
	}

	return err
}
