# 8. Caching in the Feed service
Date: 2022-07-06

## Status

Pending

## Context

When showing the general feed we will be loading a batch of the latest images. Loading them every time will add huge load to the database, the file storage and the network. Scrolling down the feed we will have to load more batches, for the user

## Decision

The Feed service will keep in-memory the latest batch of images, at it's going to be returned to all the users. The next 10 batches will be cached in Redis in binary format, so when the user scrolls we'll be able to retrieve them faster.



## Consequences
Pros:
- Showing the general feed will be fast and efficient until the cached batches.

Cons:
- Showing the feed past the cached images will be slow. It will require API call to the main service and fetching the images from GCS.
- Showing the personal feed of a user will require calls to the main service, to retrieve his already uploaded images.
