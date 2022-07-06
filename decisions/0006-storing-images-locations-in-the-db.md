# 6. Store the images locations in the  DB
Date: 2022-07-06

## Status

Pending

## Context

The images will be saved in GSC which is a file storage, however we need a easy way to find them and filter by id or user.


## Decision

The image locations will be saved in an SQL database table with the following columns:
1. Unique ID - Incrementing with every inserted picture
1. User ID - The ID of the user who uploaded the picture
1. Picture Location - URL to GCS where the picture is located

## Consequences
We'll be able to easily query for:
 - Pictures uploaded by specific user
 - Last X pictures, as the ID will determine their order
