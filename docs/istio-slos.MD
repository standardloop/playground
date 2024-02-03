# Istio for SLOs with Many CRDS

## Installtion

```sh
# base
$ helm repo add istio https://istio-release.storage.googleapis.com/charts
$ helm repo update
$ kubectl create namespace istio-system
$ helm install istio-base istio/base -n istio-system
$ helm install istiod istio/istiod -n istio-system --wait

# gateway
$ kubectl create namespace istio-gateway
$ helm install istio-gateway istio/gateway -n istio-gateway --set service.type="ClusterIP"
```

## Ingress Resources
For this document, I am using Kind Cluster locally.

I have deployed a simple Go API in the default namespace.

I'm using nginx-ingress to act as a GCLB, so nginx will route to Istio Gateway.

You can route directly to Istio using HostPorts or MetalLB but I already had this setup and continued use with it

```yaml
---
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: istio-api-ingress
  namespace: istio-gateway
spec:
  ingressClassName: nginx
  rules:
    - host: api.local
      http:
        paths:
          - path: /
            pathType: Prefix
            backend:
              service:
                name: istio-gateway
                port:
                  name: http2

---
apiVersion: networking.istio.io/v1alpha3
kind: Gateway
metadata:
  name: api-gateway
  namespace: istio-gateway
spec:
  selector:
    istio: gateway
  servers:
    - hosts:
        - api.local
      port:
        name: http
        number: 80
        protocol: HTTP

---
apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: api-vs
  namespace: default
spec:
  gateways:
    - istio-gateway/api-gateway
  hosts:
    - api.local
  http:
    - route:
        - destination:
            host: api.default.svc.cluster.local
            port:
              number: 80
          weight: 100

```

## Telemetry Related Resources

```yaml
---
apiVersion: monitoring.coreos.com/v1
kind: ServiceMonitor
metadata:
  name: istio-component-monitor
  namespace: istio-system
  labels:
    monitoring: istio-components
    release: kube-prometheus-stack
spec:
  jobLabel: istio
  targetLabels: [app]
  selector:
    matchExpressions:
      - { key: istio, operator: In, values: [pilot] }
  namespaceSelector:
    any: true
  endpoints:
    - port: http-monitoring
      interval: 2s

---
apiVersion: install.istio.io/v1alpha1
kind: IstioOperator
metadata:
  name: istio-operator
  namespace: istio-system
spec:
  meshConfig:
    accessLogFile: /dev/stdout
    accessLogFormat: |
      startTime=%START_TIME% req="%REQ(:METHOD)% %REQ(X-ENVOY-ORIGINAL-PATH?:PATH)% %PROTOCOL%" responseCode=%RESPONSE_CODE% responseFlags=%RESPONSE_FLAGS% responseCodeDetails=%RESPONSE_CODE_DETAILS% connectionTerminationDetails=%CONNECTION_TERMINATION_DETAILS% upstreamTransportFailureReason="%UPSTREAM_TRANSPORT_FAILURE_REASON%" bytesReceived=%BYTES_RECEIVED% bytesSent=%BYTES_SENT% duration=%DURATION% upstreamServiceTime=%RESP(X-ENVOY-UPSTREAM-SERVICE-TIME)% xForwardedFor="%REQ(X-FORWARDED-FOR)%" userAgent="%REQ(USER-AGENT)%" requestId="%REQ(X-REQUEST-ID)%" requestAuthority="%REQ(:AUTHORITY)%" upstreamHost="%UPSTREAM_HOST%" upstreamCluster=%UPSTREAM_CLUSTER% upstreamLocalAddress=%UPSTREAM_LOCAL_ADDRESS% downstreamLocalAddress=%DOWNSTREAM_LOCAL_ADDRESS% downstreamRemoteAddress=%DOWNSTREAM_REMOTE_ADDRESS% requestedServerName=%REQUESTED_SERVER_NAME% routeName=%ROUTE_NAME% traceID=%REQ(x-b3-traceid)%
    defaultConfig:
      tracing:
        sampling: 100
      proxyStatsMatcher:
        inclusionRegexps:
          - ".*"
  values:
    telemetry:
      enabled: true
      v2:
        enabled: true

---
apiVersion: monitoring.coreos.com/v1
kind: PodMonitor
metadata:
  name: envoy-stats-monitor
  namespace: istio-system
  labels:
    monitoring: istio-proxies
    release: kube-prometheus-stack
spec:
  selector:
    matchExpressions:
      - { key: istio-prometheus-ignore, operator: DoesNotExist }
  namespaceSelector:
    any: true
  jobLabel: envoy-stats
  podMetricsEndpoints:
    - path: /stats/prometheus
      interval: 15s
      relabelings:
        - action: keep
          sourceLabels: [__meta_kubernetes_pod_container_name]
          regex: "istio-proxy"
        - action: keep
          sourceLabels:
            [__meta_kubernetes_pod_annotationpresent_prometheus_io_scrape]
        - action: replace
          regex: (\d+);(([A-Fa-f0-9]{1,4}::?){1,7}[A-Fa-f0-9]{1,4})
          replacement: "[$2]:$1"
          sourceLabels:
            - __meta_kubernetes_pod_annotation_prometheus_io_port
            - __meta_kubernetes_pod_ip
          targetLabel: __address__
        - action: replace
          regex: (\d+);((([0-9]+?)(\.|$)){4})
          replacement: $2:$1
          sourceLabels:
            - __meta_kubernetes_pod_annotation_prometheus_io_port
            - __meta_kubernetes_pod_ip
          targetLabel: __address__
        - action: labeldrop
          regex: "__meta_kubernetes_pod_label_(.+)"
        - sourceLabels: [__meta_kubernetes_namespace]
          action: replace
          targetLabel: namespace
        - sourceLabels: [__meta_kubernetes_pod_name]
          action: replace
          targetLabel: pod_name

---
apiVersion: extensions.istio.io/v1alpha1
kind: WasmPlugin
metadata:
  name: istio-attributegen-filter
  namespace: istio-system
spec:
  selector:
    matchLabels:
      app: api
  url: https://storage.googleapis.com/istio-build/proxy/attributegen-359dcd3a19f109c50e97517fe6b1e2676e870c4d.wasm
  imagePullPolicy: Always
  phase: AUTHN
  pluginConfig:
    attributes:
      - output_attribute: "istio_operationId"
        match:
          - value: "RandNum"
            condition: "request.url_path == '/api/v1/rand' && request.method == 'GET'"
          - value: "RandNumSleep"
            condition: "request.url_path == '/api/v1/rand-sleep' && request.method == 'GET'"

---
apiVersion: telemetry.istio.io/v1alpha1
kind: Telemetry
metadata:
  name: custom-tags
  namespace: istio-system
spec:
  metrics:
    - overrides:
        - match:
            metric: ALL_METRICS
            mode: CLIENT_AND_SERVER
          tagOverrides:
            request_operation:
              value: istio_operationId
      providers:
        - name: prometheus
  accessLogging:
    - providers:
        - name: envoy
```
