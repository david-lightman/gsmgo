# gsmgo
Client to read GSM secrets

Simple go-based client to read the **latest** Google Secrets Manager secret-version from the associated Google project. Perfect for use in GCE / containers that have the proper service-account with appropriate role to read a given GSM secret-version.  

## setup

Set/export an environment variable with the id of the GSM secret-version to read.

```shell
export SECRET='foo'  # e.g. projects/<project>/secrets/foo/versions/latest
./gsmgo              # outputs the secret-version
```
