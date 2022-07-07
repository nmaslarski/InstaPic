# 9. Streaming the general feed to the FE app
Date: 2022-07-07

## Status

Pending

## Context

The frontend application will need update the feed realtime.

## Decision

Suitable option for live updates of many clients are sockets.

When the client requests specific users feed we will fetch it from the main service and stream it. Upon events from the main service we will have to validate if the picture is uploaded by the user who's viewing his feed and sent if it is


## Consequences

The Feed service will be able to stream the pictures directly to all online clients, when it receives them from the main service, or loads them from the cache
