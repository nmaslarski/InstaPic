# 7. Updating the feed service about new images
Date: 2022-07-06

## Status

Pending

## Context

The feed service will update the users feed realtime, hence it has to know when new image is uploaded. There can be multiple instances of the feed service and all of them need to know about the update. Subscription based update fits this scenario perfectly.

But also on starting a new Feed service or when a user requests his feed the service will need to fetch specific set of pictures (last 20-30 or the once uploaded by a given user)

## Decision

Pub/Sub and API calls.

Pub/Sub - Every new instance of the Feed service will create it's own subscription for the Topic where the new images are pushed. Cache them in memory and in centralised cache.

The main service will expose an API trough which other services (the Feed service) will be able to query images by user, or id.


## Consequences
Having a queue will easily allow other services to know about images, without the main service knowing about them.
