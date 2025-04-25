# `paxos-go`

To generate, simply run `go generate`

To see full docs, see: [paxos-README.md](paxos-README.md).

> [!WARNING]
> Note that we have altered the ./templates directory to account
> for the fact the default generation will fail with this OpenAPI
> spec because there are enums with duplicate names. To solve
> this we've had to alter enum generation to suffix the datatype.
