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
  extraAnnotations:
    comcast.com/testAnnotation: test.annotation
  extraLabels:
    testLabel: testLabel
  podSpec:
    containers:
      - env:
          - name: MINIO_ENDPOINT
            value: "https://changme.com:8443"
          - name: ACCESS_KEY
            value: "CHANGEME"
          - name: SECRET_KEY
            value: "CHANGEME"
        image: kuberhealthy/minio-test:v1.0.0
        imagePullPolicy: IfNotPresent
        name: main
        resources:
          requests:
            cpu: 10m
            memory: 50Mi
```

#### How-to

Apply a `.yaml` file similar to the one shown above with `kubectl apply -f`
