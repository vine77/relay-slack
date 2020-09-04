# slack-step-lookup-user-by-email

This [Slack](https://slack.com) step gets a users member id by providing an email address.

## Specification

| Setting | Data type | Description | Default | Required |
|---------|-----------|-------------|---------|----------|
| `connection` | Relay Connection   | Connection to Slack requiring api token | None | True |
| `email` | string | Channel to send message to | None | True | 

## Permissions 

This step requires that you have the following OAuth scopes enabled in your Slack App:

* users:read
* users:read.email

## Example  

```yaml
steps:
# ...
- name: slack-lookup-member-id-by-email
  image: relaysh/slack-step-lookup-user-by-email
  spec:
    connection: !Connection { type: slack, name: my-slack-account}
    email: "foo@example.com"
```
