package your_integration

import "encoding/json"

func (n *NeuroIntegration) registerActions() error {
	actions := []ActionDefinition{
		// THESE 3 ACTIONS ARE BROKEN ASF, DON'T UNCOMMENT
		ActionSchemas[string(CmdMouseMove)],
	}

	data := map[string]interface{}{
		"actions": actions,
	}
	dataBytes, _ := json.Marshal(data)

	return n.sendMessage(NeuroMessage{
		Command: "actions/register",
		Data:    dataBytes,
	})
}

func (n *NeuroIntegration) unregisterActions() error {
	data := map[string]interface{}{
		"actions": []string{
			string(
		},
	}
	dataBytes, _ := json.Marshal(data)

	return n.sendMessage(NeuroMessage{
		Command: "actions/unregister",
		Data:    dataBytes,
	})
}
