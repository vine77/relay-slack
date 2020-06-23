# slack-step-message-send

This [Slack](https://slack.com) step sends a message to a desired channel.

## Specification

| Setting | Data type | Description | Default | Required |
|---------|-----------|-------------|---------|----------|
| `connection` | Relay Connection   | Connection to Slack requiring api token | None | True |
| `channel` | string | Channel to send message to | None | True | 
| `username` | string | Username to send message as | None | True | 
| `message` | string | Message to send | None | True | 

In order to connect to Slack you will need to ["create a new Slack app"](https://api.slack.com/apps) as they have deprecated the creation of new legacy tokens. Once you have created the app you will need to apply the correct OAuth scopes. For the purpose of sending messages from Relay, and allowing custom user names, you will need to ensure that the Bot Token Scopes include the chat:write, chat:write:public, and chat:write:customize scopes. For some organizations your Slack app may require approval before it can be installed. Once approved and installed, you can copy the Bot User OAuth Access Token and supply it to the Slack connection form in Relay.

For more information on creating a Slack app, check out ["Start building Slack apps"](https://api.slack.com/start).

## Outputs
None

## Example  

```yaml
steps:
# ...
- name: notify-slack
  image: relaysh/slack-step-message-send
  spec:
    channel: nebula-workflows
    connection: !Connection { type: slack, name: my-slack-account}
    message: "hello Relay!"
    username: "Relay Workflows"
```
