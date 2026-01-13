# Minio Test

khcheck to create a minio bucket, create an object, then delete both

## What it is
This repository builds the container image used by Kuberhealthy to run the minio-test check.

## Image
- `docker.io/kuberhealthy/minio-test`
- Tags: short git SHA for `main` pushes and `vX.Y.Z` for releases.

## Quick start
- Apply the example manifest: `kubectl apply -f minio-test.yaml`
- Edit the manifest to set any required inputs for your environment.

## Build locally
- `docker build -f ./Dockerfile -t kuberhealthy/minio-test:dev .`

## Contributing
Issues and PRs are welcome. Please keep changes focused and add a short README update when behavior changes.

## License
See `LICENSE`.
