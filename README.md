## Usage Example:

The v2 code includes a comprehensive example showing how to:
- Create action handlers with validation
- Register actions
- Use action windows for turn-based games
- Handle context messages
- Configure priorities and options

## To use this SDK:

```bash
go get github.com/gorilla/websocket
```

Then create your action handlers by implementing the `ActionHandler` interface, which cleanly separates validation (where you check parameters and return success/failure) from execution (where you actually perform the action).
