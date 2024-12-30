# Temporal 101 - Go

https://temporal.talentlms.com/learner/courseinfo/id:126

Code repository for the Temporal 101 with Go course exercises.

This departs from the upstream source repository and instead uses the latest versions of Temporal libraries and Temporal Server (at time of writing). It also uses a Docker Compose local environment instead of relying on the Temporal CLI.

## Temporal Cluster

There is a local Temporal Cluster ready to use in `./temporal-server`

```bash
docker compose up
```

## Reference

- [General Temporal Documentation](https://docs.temporal.io/)
- [Temporal Go Developer's Guide](https://docs.temporal.io/develop/go/)
- [Temporal Go SDK API Reference](https://pkg.go.dev/go.temporal.io/sdk)
- [Go Language Documentation](https://go.dev/doc/)
