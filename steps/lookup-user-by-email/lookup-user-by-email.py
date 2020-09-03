#!/usr/bin/env python
# queries the slack api with the email address and gets back a user id

import os
from relay_sdk import Interface, Dynamic as D
from slack import WebClient
from slack.errors import SlackApiError

relay = Interface()
api_key = relay.get(D.connection.apiKey),

client = WebClient(api_key)

try:
    response = client.users_lookupByEmail(email='geoff.woodburn@puppet.com')
    
    relay.outputs.set("member_id", response['user']['id'])
except SlackApiError as e:
    # You will get a SlackApiError if "ok" is False
    assert e.response["ok"] is False
    assert e.response["error"]  # str like 'invalid_auth', 'channel_not_found'
    print(e.response)