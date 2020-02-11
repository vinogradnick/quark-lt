package quark_lt

import "github.com/quark_lt/cmd/quark_node"

type QuarkScheduler struct {
	QuarkNodes map[int]quark_node.QuarkNode
}
