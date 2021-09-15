# TPR: Terraform private registry

## Run

```shell
tpr --config /path/to/config.yaml serv
```

## Storage implementations

- [x] filesystem
- [ ] s3
- [ ] artifactory

## Roadmap

- Implementation works with artifactory as a storage
- Store gpg keys in Hashicorp vault
- Implementation works with s3 as a storage
- Implementation works modules api
- Write a lot of tests :)