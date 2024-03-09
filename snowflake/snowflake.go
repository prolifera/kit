package snowflake

import (
	"github.com/bwmarrin/snowflake"
	"sync"
)

var snowFlakeNode *snowflake.Node
var lock = sync.Mutex{}

func GetSnowUid() int64 {
	if snowFlakeNode != nil {
		return snowFlakeNode.Generate().Int64()
	}
	lock.Lock()
	defer lock.Unlock()
	if snowFlakeNode == nil {
		snowFlakeNode, _ = snowflake.NewNode(1)
		return snowFlakeNode.Generate().Int64()
	} else {
		return snowFlakeNode.Generate().Int64()
	}
}
