package neuro_integration_sdk

import (
  	"encoding/json"
)

func (n *NeuroIntegration) sendMessage(msg NeuroMessage) error {
	msg.Game = n.gameName
	return n.ws.WriteJSON(msg)
}

func (n *NeuroIntegration) startup() error {
	return n.sendMessage(NeuroMessage{
		Command: "startup",
	})
}

func (n *NeuroIntegration) sendContext(message string, silent bool) error {
	data := map[string]interface{}{
		"message": message,
		"silent":  silent,
	}
	dataBytes, _ := json.Marshal(data)

	return n.sendMessage(NeuroMessage{
		Command: "context",
		Data:    dataBytes,
	})
}

func (n *NeuroIntegration) sendActionResult(actionID string, success bool, message string) error {
	data := map[string]interface{}{
		"id":      actionID,
		"success": success,
		"message": message,
	}
	dataBytes, _ := json.Marshal(data)

	return n.sendMessage(NeuroMessage{
		Command: "action/result",
		Data:    dataBytes,
	})
}

func (n *NeuroIntegration) sendShutdownReady() error {
	return n.sendMessage(NeuroMessage{
		Command: "shutdown/ready",
	})
}
