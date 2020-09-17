# slack-step-get-member-id-by-email

This [Slack](https://slack.com) step gets a users member id by providing an email address.

## Permissions 

This step requires that you have the following OAuth scopes enabled in your Slack App:

* users:read
* users:read.email

## Example  

```yaml
steps:
# ...
- name: slack-get-member-id-by-email
  image: relaysh/slack-step-get-member-id-by-email
  spec:
    connection: !Connection { type: slack, name: my-slack-account}
    email: "foo@example.com"
```
