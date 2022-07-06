# 4. Split the backend in main and feed services
Date: 2022-07-06

## Status

Pending

## Context

InstaPic will allow for few operations:
 - Upload pictures
 - View Feed

Where uploading will be less frequent. And viewing the feed will be frequent operation - every user who has opened the app will be viewing the feed with live updates.

## Decision

As viewing the feed will be frequently executed by many users and will be data intense we want it to be separately scalable from the main service featuring not so frequently executed operations


## Consequences
Pros:
- Having separate Feed Service will allow us to scale it separately, have its own cache even deploy it multi-regionally if needed.
- We'll have a clear separation of concerns.

Cons:  
- We'll need to update it with the newly uploaded pictures
