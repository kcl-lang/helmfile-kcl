# Helmfile KCL Plugin

[![Go Report Card](https://goreportcard.com/badge/github.com/kcl-lang/helmfile-kcl)](https://goreportcard.com/report/github.com/kcl-lang/helmfile-kcl)
[![GoDoc](https://godoc.org/github.com/kcl-lang/helmfile-kcl?status.svg)](https://godoc.org/github.com/kcl-lang/helmfile-kcl)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/kcl-lang/helmfile-kcl/blob/main/LICENSE)

[KCL](https://github.com/KusionStack/kcl) is a constraint-based record & functional domain language. Full documents of KCL can be found [here](https://kcl-lang.io/).

You can use the `helmfile-kcl plugin` to

+ Edit the helm charts in a hook way to separate data and logic for the Kubernetes manifests management.
+ For multi-environment and multi-tenant scenarios, you can maintain these configurations gracefully rather than simply copy and paste.
+ Validate all KRM resources using the KCL schema.

## Install

### Prerequisites

+ [Kustomize](https://github.com/kubernetes-sigs/kustomize)
+ [Helm](https://github.com/helm/helm)
+ [Helmfile](https://github.com/helmfile/helmfile)

## Quick Start

```shell
cd examples/hello-world && helmfile apply
```

The content of `helmfile.yaml` looks like this:

```yaml
repositories:
- name: prometheus-community
  url: https://prometheus-community.github.io/helm-charts

releases:
- name: prom-norbac-ubuntu
  namespace: prometheus
  chart: prometheus-community/prometheus
  set:
  - name: rbac.create
    value: false
  transformers:
  # Use KRM KCL Plugin to mutate or validate Kubernetes manifests.
  - apiVersion: krm.kcl.dev/v1alpha1
    kind: KCLRun
    metadata:
      name: "set-annotation"
      annotations:
        config.kubernetes.io/function: |
          container:
            image: docker.io/kcllang/kustomize-kcl:v0.1.2
    spec:
      params:
        annotations:
          config.kubernetes.io/local-config: "true"
      source: oci://ghcr.io/kcl-lang/set-annotation
```

## Guides for Developing KCL

Here's what you can do in the KCL script:

+ Read resources from `option("resource_list")`. The `option("resource_list")` complies with the [KRM Functions Specification](https://kpt.dev/book/05-developing-functions/01-functions-specification). You can read the input resources from `option("resource_list")["items"]` and the `functionConfig` from `option("resource_list")["functionConfig"]`.
+ Return a KPM list for output resources.
+ Return an error using `assert {condition}, {error_message}`.
+ Read the environment variables. e.g. `option("PATH")` (Not yet implemented).
+ Read the OpenAPI schema. e.g. `option("open_api")["definitions"]["io.k8s.api.apps.v1.Deployment"]` (Not yet implemented).

Full documents of KCL can be found [here](https://kcl-lang.io/).

## Examples

See [here](https://kcl-lang.io/krm-kcl/tree/main/examples) for more examples.

## Thanks

+ [helmfile](https://github.com/helmfile/helmfile)
+ [helm-diff](https://github.com/databus23/helm-diff)
+ [helm-secrets](https://github.com/jkroepke/helm-secrets)
+ [helm-s3](https://github.com/hypnoglow/helm-s3)
+ [helm-git](https://github.com/aslafy-z/helm-git)
