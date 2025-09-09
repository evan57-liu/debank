package idutils

import (
	"sync/atomic"

	"github.com/bwmarrin/snowflake"
)

const nodeNumber = 1

var nodes []*snowflake.Node
var count int32 = 0

func init() {
	for i := 0; i < nodeNumber; i++ {
		node, err := snowflake.NewNode(int64(i))
		if err != nil {
			panic(err)
		}
		nodes = append(nodes, node)
	}
}

func GenID() int64 {
	// 原子增加 count 并且限制在 nodeNumber 范围内
	index := atomic.AddInt32(&count, 1) % nodeNumber

	return nodes[index].Generate().Int64()
}
