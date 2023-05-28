package model

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

type Command struct {
	Cmd    string
	Task   string
	IsTask bool
}

func (c *Command) String() string {
	if c.IsTask {
		return fmt.Sprintf("task: %s", c.Task)
	} else {
		return fmt.Sprintf("cmd: %s", c.Cmd)
	}
}

func (c *Command) UnmarshalYAML(node *yaml.Node) error {
	switch node.Kind {
	case yaml.SequenceNode:
		return fmt.Errorf("unexpected node %v", node.Value)
	case yaml.MappingNode:
		nodeMap := make(map[string]string)
		err := node.Decode(&nodeMap)

		if err != nil {
			return err
		}

		taskName, ok := nodeMap["task"]
		if ok {
			c.IsTask = true
			c.Task = taskName
		}

	case yaml.ScalarNode:
		c.Cmd = node.Value
		c.IsTask = false
	case yaml.AliasNode:
		return fmt.Errorf("unexpected node %v", node.Value)
	}
	return nil
}
