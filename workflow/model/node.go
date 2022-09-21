package model

import (
	"container/list"
	"errors"
	"strings"
)

// Node represents a specific logical unit of processing and routing
// in a workflow.
// 流程中的一个节点
type Node struct {
	Name         string `json:"name"`                //子流程名字
	Type         string `json:"type"`                //子流程类型
	NodeId       string `json:"node_id"`             //流程ID
	PrevId       string `json:"prev_id,omitempty"`   //上一个流程ID
	AssigneeId   string `json:"assignee_id"`         //审批人ID
	AssigneeName string `json:"assignee_name"`       //审批人名字
	ChildNode    *Node  `json:"childNode,omitempty"` //子节点
}

// NodeInfo 节点信息
type NodeInfo struct {
	Name         string `json:"name"`
	Type         string `json:"type"`
	NodeID       string `json:"node_id"`
	PrevID       string `json:"prev_id,omitempty"`
	AssigneeID   string `json:"assignee_id"`
	AssigneeName string `json:"assignee_name"`
}

// IfProcessConifgIsValid 检查流程配置是否有效
func IfProcessConifgIsValid(n *Node) error {
	// 节点名称是否有效
	//if len(node.NodeID) == 0 {
	//	return errors.New("节点的【nodeId】不能为空！！")
	//}
	// 检查类型是否有效
	if len(n.Type) == 0 {
		return errors.New("没获取到流程节点类型")
	}

	const allowType = "start|approver"
	if !strings.Contains(allowType, strings.ToLower(n.Type)) { // 是否为允许类型
		return errors.New("没获取到流程节点类型")
	}

	// 当前节点是否设置有审批人
	if n.AssigneeId == "" {
		return errors.New("流程节点存在没有审批人ID")
	}

	// 子节点是否存在
	if n.ChildNode != nil {
		return IfProcessConifgIsValid(n.ChildNode)
	}
	return nil
}

func (n *Node) Add2ExecutionList(list *list.List) {

	list.PushBack(NodeInfo{
		Name:         n.Name,
		Type:         n.Type,
		NodeID:       n.NodeId,
		PrevID:       n.PrevId,
		AssigneeID:   n.AssigneeId,
		AssigneeName: n.AssigneeName,
	})
}
