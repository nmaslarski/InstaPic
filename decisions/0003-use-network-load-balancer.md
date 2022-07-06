# 3. Use Network Load Balancer
Date: 2022-07-04

## Status

Pending

## Context

We're planning on having distributed and scalable architecture, meaning we might have multiple instances of our BE services. The client need to be able to connect to any one of them without knowing the specifics - where it is located, which once are available, how many are there etc...
We need a way to handle and distribute incoming traffic.

## Decision

Multiple decisions were investigated:
 - Reverse Proxy (f.e. nginx). It can provide all the required features, but has to be managed by the team, it's HA and Security is another dependancy to the dev team.
 - DNS Routing. Does not provide a way to handle autoscaling services. Have to register and propagate/delete records every time the instances are being scaled.
 DNS registration is slow.
 - Network Load Balancer: Provides the required features. Lightweight, easy to set up, secure by default. Can handle redirects for unauthenticated users.


## Consequences
Pros:
- Handles autoscaling BE apps
- Redirects unauthenticated users to the IDP services
- Secured against DDOS and other attacks

Cons:
- Single zone (EU/US/AP), will need to move to another load balancing solution for multiple regions
