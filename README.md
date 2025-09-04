# minio-test
khcheck to create a minio bucket, create an object, then delete both

This check requires three environment variables in the kube spec used to deploy it.  A minio API endpoint, and the access/secret (or user/pass).

**Example yaml**

```yaml
---
apiVersion: comcast.github.io/v1
kind: KuberhealthyCheck
metadata:
  name: minio-test
  namespace: kuberhealthy
spec:
  runInterval: 30s
  timeout: 2m
  podSpec:
    containers:
      - name: minio-test
        image: kuberhealthy/minio-test:v1.0.0
        imagePullPolicy: IfNotPresent
        env:
          - name: MINIO_ENDPOINT
            value: "https://changme.com:8443"
          - name: ACCESS_KEY
            value: "CHANGEME"
          - name: SECRET_KEY
            value: "CHANGEME"
        resources:
          requests:
            cpu: 10m
            memory: 50Mi
```

Kuberhealthy automatically adds several environment variables to each check pod:

- `KH_REPORTING_URL` – endpoint for reporting check status.
- `KH_CHECK_RUN_DEADLINE` – UNIX deadline for the current check run.
- `KH_RUN_UUID` – unique identifier for the check run, used when reporting.
- `KH_POD_NAMESPACE` – namespace where the check pod is running.

#### How-to

Apply a `.yaml` file similar to the one shown above with `kubectl apply -f`
