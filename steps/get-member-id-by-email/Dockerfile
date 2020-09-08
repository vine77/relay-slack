FROM python:3.8-alpine as base

RUN apk add --update gcc musl-dev

RUN pip install wheel && pip wheel slackclient --wheel-dir=/wheels

FROM relaysh/core:latest-python
COPY --from=base /wheels /wheels
RUN pip install --no-index --find-links=/wheels slackclient
COPY "./lookup-user-by-email.py" "/entrypoint.py"
ENTRYPOINT []
CMD ["python3", "/entrypoint.py"]

LABEL "org.opencontainers.image.title"="Slack Lookup Member ID by Email"
LABEL "org.opencontainers.image.description"="This step gets a slack users member id from an email"
LABEL "com.puppet.nebula.sdk.version"="v1"