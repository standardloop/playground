# playground

[![Kubernetes](https://img.shields.io/badge/Kubernetes-326CE5?logo=kubernetes&logoColor=fff)](#)

---

https://github.com/standardloop/playground

Playground is a local environment for playing around with various kubernetes tools.

## Tools Used

### Setup
- [task (go-task)](https://github.com/go-task/task) for the overall automation.
- [rancher-desktop](https://github.com/rancher-sandbox/rancher-desktop/) or [colima](https://github.com/abiosoft/colima) for handling docker engine.
- [kind](https://github.com/kubernetes-sigs/kind) for provisioning a local kubernetes cluster.
    - [local registry](https://kind.sigs.k8s.io/docs/user/local-registry/) is enabled on the cluster.
- [cloud-provider-kind](https://github.com/kubernetes-sigs/cloud-provider-kind) for the ability to have access to `gatewayClassName: cloud-provider-kind`.
- [opentofu](https://github.com/opentofu/opentofu) for deploying infrastructure applications.
    - [kustomize provider](https://github.com/kbst/terraform-provider-kustomization) for deploying [kustomizations](https://github.com/kubernetes-sigs/kustomize).
    - [helm provider](https://github.com/hashicorp/terraform-provider-helm) for deploying [helm](https://github.com/helm/helm) charts (this repo mostly relies on `flux` `HelmRelease` but there are a few cases of just using raw helm charts).
    - [null provider](https://github.com/hashicorp/terraform-provider-null) mostly for using `local-exec`.
- [flux](https://github.com/fluxcd/flux2) deployed (minimally) to have access to `HelmRelease` and `HelmRepository` crds.

### Infrastructure Applications

- [metrics-server](https://github.com/kubernetes-sigs/metrics-server) for playing around with autoscaling.
- [argo-rollouts](https://github.com/argoproj/argo-rollouts) for playing around with deployment strategies.
- [headlamp](https://github.com/kubernetes-sigs/headlamp) as a kubernetes web UI.
- [istio](https://github.com/istio/istio) as the entryway into the cluster.
- [kiali](https://github.com/kiali/kiali) for some insights into istio.
- [prometheus](https://github.com/prometheus/prometheus) deployed via `kube-prometheus-stack` for metrics.
- [grafana](https://github.com/grafana/grafana) deploy via [grafana-operator](https://github.com/grafana/grafana-operator) for monitoring dashboards.
- [homer](https://github.com/bastienwirtz/homer) as the home page of the cluster.

### Test Applications
- [tilt](https://github.com/tilt-dev/tilt) for managing test applications in the cluster.

### Documentation and Standards
- [prek](https://github.com/j178/prek) for pre-commit.
- [yamllint](https://github.com/adrienverge/yamllint) for enforcing yaml rules.
- [helm-docs](https://github.com/norwoodj/helm-docs) for managing helm documentation.
- [helm-schema](https://github.com/dadav/helm-schema) for handling helm `values.schema.json` generation.
- [terraform-docs](https://github.com/terraform-docs/terraform-docs) for handling tofu documentation.

## Screenshot

![alt text](https://raw.githubusercontent.com/standardloop/playground/refs/heads/main/docs/screenshot.png)


## Spin Up

```sh
$ task
```

Note: `sudo` is needed to run `cloud-provider-kind`.

## Clean Up

```sh
$ task clean
```

## Updating `/etc/hosts`

You will need to update your `/etc/hosts` to have access to the applications on the local cluster.

### Find you IP

```sh
$ kubectl get gateway playground-local-gateway  -n istio-gateway
NAME                       CLASS                 ADDRESS      PROGRAMMED   AGE
playground-local-gateway   cloud-provider-kind   172.18.0.5   True         17m
```

### `/etc/hosts` example

```sh
$ cat /etc/hosts`
172.18.0.5      playground.local
172.18.0.5      grafana.playground.local
172.18.0.5      prometheus.playground.local
172.18.0.5      api.playground.local
172.18.0.5      www.playground.local
172.18.0.5      rollouts.playground.local
172.18.0.5      api.playground.local
172.18.0.5      headlamp.playground.local
172.18.0.5      homer.playground.local
172.18.0.5      home.playground.local
172.18.0.5      kiali.playground.local
```

## Mise WIP

```sh
$ mise install
...
$ mise uninstall --all
```
