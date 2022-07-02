# 1. Use Google Identity for Authentication
Date: 2022-07-01

## Status

Pending

## Context

We need a way to manage users - Create, Authenticate, manage sessions and store credentials.
Options are to grow in house identity providing mechanism or use an external service.

## Decision

We will use external IDP as it's more secure, fast and easy to work with.

## Consequences
Pros:
- Secure way of authenticating out of the box
Cons:
- We need to redirect unauthenticated users to Google Identity Service
