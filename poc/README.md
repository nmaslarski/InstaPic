### POC for real time updates of the feed

_WARNING: This is proof of concept only. Not to be used in any way in production. The code is not properly tested, with no error handling and is thread unsafe, with multiple race conditions and leaking resources_


For the sake of simplicity:
- Instead of having multiple services (Main service and Feed service), I've represented them as two go routines in the same server. Communicating trough a channel acting as a queue.
- Our client for the POC is a cli, which can easily be started in multiple terminals and for the sake of simplicity it's uploading strings, and not pictures

Running the server:
```
make build-server
make start-server
```
Running the cli:
```
make build-cli
make start-cli
```

Starting the CLI in multiple terminals we can observe that if you push a message trough one of them, all other users will receive the event
