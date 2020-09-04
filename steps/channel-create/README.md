# slack-step-channel-create

This [Slack](https://slack.com) step creates a channel if it doesn't exist.

## Specification

| Setting      | Data type        | Description                             | Default | Required |
|--------------|------------------|-----------------------------------------|---------|----------|
| `connection` | Relay Connection | Connection to Slack requiring api token | None    | True     |
| `channel`    | string           | Channel to create                       | None    | True     |
| `topic`      | string           | Topic to set, if any                    | None    | False    |
| `memberID`   | string           | User to invite, if any                  | None    | False    |

## Outputs

| Name        | Data type | Description                 |
|-------------|-----------|-----------------------------|
| `channelID` | String    | The id of the slack channel |

In order to connect to Slack you will need to **Create a new Slack app** at [https://api.slack.com/apps/](https://api.slack.com/apps?new_app=1), then:

* Give it a name like "Relay.sh"
* Select the workspace you want to target
* Navigate to **Add features and functionality**
  * Click **Permissions**
  * On the **OAuth and Permissions** page, go to **Scopes**
  * Under **Bot Token Scopes**, add `channels:manage`
* For some organizations your Slack app may require approval before it can be installed.
* Once approved and installed, you'll see a **Bot User OAuth Access Token** on the site.
* In Relay, go to the **Connections** tab and click **Add connection**. Choose **Slack** from the drop-down.
* Give the connection a name which you'll reference in your workflows and paste the token in.

## Example

```yaml
parameters:
  incidentNumber:
    description: "The incident number being addressed by this war room"

steps:
# ...
- name: create-channel
  image: relaysh/slack-step-channel-create
  spec:
    connection: !Connection { type: slack, name: my-slack-account}
    channel: !Fn.concat ["incident-", !Parameter incidentNumber]
    topic: !Fn.concat ["Addressing incident ", !Parameter incidentNumber]
    memberID: "UDHPDS88Z"
```
