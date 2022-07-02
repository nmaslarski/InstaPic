# 1. Use Google Kubernetes Engine
Date: 2022-07-01

## Status

Pending

## Context

We need a way to deploy and manage our service.

Evaluated  ideas:
1. Deploying directly on a VM - Hard to maintain, we have to orchestrate the deployment and monitoring of the process, not necessarily cheaper
1. Using managed platform - Handles a lot of complexity for you f.e monitoring and restarting the process. Network communication between containers. Better utilization of resources etc.


## Decision

KGE provides good abstraction over the hypervisor layer with enough access to the underlying resources to be flexible when you need. Kubernetes has good utilization of resources. Provides Liveness, Readiness and Start probes, flexible auto-scaling depending on the user needs and more.

Using it as a service provided from Google will enable us to focus on the InstaPic application itself, instead of managing infrastructure, as well as access to other GCP services via internal network.

## Consequences
Pros:
- The team will focus on developing the app itself
- Outsourcing responsibilities around keeping the process alive, zero downtime deployment, auto-scailing etc. 

Cons:
- Tight coupling with the provider.
- Might end up more expensive than the alternatives .
