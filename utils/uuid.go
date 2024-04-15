package utils

import "github.com/bwmarrin/snowflake"

var sn *snowflake.Node

func Init() error {
	node, err := snowflake.NewNode(0)
	if err != nil {
		return err
	}
	sn = node
	return nil
}

func GenerateUUID() int64 {
	return sn.Generate().Int64()
}
