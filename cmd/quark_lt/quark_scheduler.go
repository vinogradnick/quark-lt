package quark_lt

import (
	"github.com/quark_lt/cmd/quark_node"
)

type QuarkScheduler struct {
	QuarkNodes map[int]*quark_node.QuarkNode
}

func NewQuarkScheduler() *QuarkScheduler {
	return &QuarkScheduler{QuarkNodes: map[int]*quark_node.QuarkNode{}}
}
func (q QuarkScheduler) AddNode(node quark_node.QuarkNode) {
	q.QuarkNodes[node.Id] = &node
}
func (q QuarkScheduler) CheckNode(id int) {

}
