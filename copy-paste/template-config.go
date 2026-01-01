package neuro_integration_sdk

const (
	Action1        CommandType = "Action1"
  Action2        CommandType = "Action2"
  Action3        CommandType = "Action3"
)

var ActionSchemas = map[string]ActionDefinition{
	string(Action3): {
		Name:        string(Action3),
		Description: "PETPETPET A!",
		Schema:      map[string]interface{}{},
	},
	string(Action2): {
		Name:        string(Action2),
		Description: "we love variable a! we also love cookies!",
		Schema: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"a": map[string]interface{}{
					"type":        "string",
					"description": "say hi to A!",
				},
				"cookies": map[string]interface{}{
					"type":        "integer",
					"description": "Give A some cookies!",
					"default":     100000,
				},
			},
			"required": []string{"cookies"},
		},
	},
	string(Action1): {
		Name:        string(Action1),
		Description: "Actyyy",
		Schema:      map[string]interface{}{},
	},
}
