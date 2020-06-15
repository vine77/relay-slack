# slack-step-message-send

This [Slack](https://slack.com) step sends a message to a desired channel.

## Specification

| Setting | Data type | Description | Default | Required |
|---------|-----------|-------------|---------|----------|
| `connection` | Relay Connection   | Connection to Slack requiring api token | None | True |
| `channel` | string | Channel to send message to | None | True | 
| `username` | string | Username to send message as | None | True | 
| `message` | string | Message to send | None | True | 

For more information on creating a Slack token, check out ["Create and regenerate API tokens"](https://slack.com/help/articles/215770388-Create-and-regenerate-API-tokens).

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
