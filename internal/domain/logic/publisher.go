// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package logic

import "context"

type Publisher struct {
}

func (p *Publisher) Publish(ctx context.Context, topic string, message any) error {
	return nil
}
