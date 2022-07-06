# 5. Store the images in GCS
Date: 2022-07-06

## Status

Pending

## Context

We need a place to keep the users images, with easy access.
Evaluated options:
- Database - databases are not meant for file storages, and storing the image files there is not impossible, but will put a lot of load on the DB
- GCS - meant for file storage, easily accessible, same datacenter and network as the services


## Decision

Save the images in GCS, and saving the location in the database, along with the other required properties

## Consequences
Pros:
- easily accessible pictures
- no load on the database
- distributed storage
- sending image location to the other services, and not the image itself
