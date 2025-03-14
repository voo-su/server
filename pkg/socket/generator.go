package socket

import "github.com/bwmarrin/snowflake"

type IIdGenerator interface {
	IdGen() int64
}

type SnowflakeGenerator struct {
	Snowflake *snowflake.Node
}

var defaultIdGenerator IIdGenerator

func init() {
	node, err := snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}

	defaultIdGenerator = &SnowflakeGenerator{
		Snowflake: node,
	}
}

func (s *SnowflakeGenerator) IdGen() int64 {
	return s.Snowflake.Generate().Int64()
}
