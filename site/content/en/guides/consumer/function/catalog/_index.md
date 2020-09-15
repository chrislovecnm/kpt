---
title: "Functions Catalog"
linkTitle: "Functions Catalog"
weight: 6
type: docs
description: >
    Catalog of Config Functions.
---

<!---
DO NOT EDIT. Generated by: "cd catalog; npm run gen-docs"
-->

This repository documents a catalog of functions implementing the
[Configuration Functions Specification][spec].

Run functions either imperatively or declaratively by following the
[Functions User Guide].

Implement configuration functions using any toolchain such as the
[Typescript SDK][ts sdk] or [Golang Libraries][go libs].

## Sources

See [definition of source functions][source].

| Image                                         | Args              | Description                                                       | Example                                                                                                        | Source                                                                                                                  | Toolchain                                         |
| --------------------------------------------- | ----------------- | ----------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------- |
| gcr.io/kpt-dev/kpt                            | fn source         | Reads a directory of Kubernetes configuration recursively.        |                                                                                                                | [Source](https://github.com/kubernetes-sigs/kustomize/blob/master/cmd/config/internal/commands/source.go)               |                                                   |
| gcr.io/kpt-functions/read-yaml                |                   | [Demo] Reads a directory of Kubernetes configuration recursively. |                                                                                                                | [Source](https://github.com/GoogleContainerTools/kpt-functions-sdk/blob/master/ts/demo-functions/src/read_yaml.ts)      | [Typescript SDK](../../../producer/functions/ts/) |
| gcr.io/cloud-builders/kubectl                 | get [...] -o yaml | Get one or many resources from a Kubernetes cluster.              |                                                                                                                | [Source](https://github.com/GoogleCloudPlatform/cloud-builders/blob/master/kubectl/Dockerfile)                          |                                                   |
| gcr.io/kustomize-functions/create-application |                   | Add an Application CR to a group of resources.                    |                                                                                                                | [Source](https://github.com/kubernetes-sigs/kustomize/blob/master/functions/examples/application-cr/image/main.go)      | [Go Library](../../../producer/functions/golang/) |
| gcr.io/kpt-functions/kustomize-build          |                   | Run kustomize to build configuration                              | [Example](https://github.com/GoogleContainerTools/kpt-functions-catalog/tree/master/examples/kustomize-build/) | [Source](https://github.com/GoogleContainerTools/kpt-functions-catalog/blob/master/functions/ts/src/kustomize_build.ts) | [Typescript SDK](../../../producer/functions/ts/) |

## Sinks

See [definition of sink functions][sink].

| Image                           | Args    | Description                                                                                                                       | Example | Source                                                                                                              | Toolchain                                         |
| ------------------------------- | ------- | --------------------------------------------------------------------------------------------------------------------------------- | ------- | ------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------- |
| gcr.io/kpt-dev/kpt              | fn sink | Writes a directory of Kubernetes configuration. It maintains the original directory structure as read by source functions.        |         | [Source](https://github.com/kubernetes-sigs/kustomize/blob/master/cmd/config/internal/commands/sink.go)             |                                                   |
| gcr.io/kpt-functions/write-yaml |         | [Demo] Writes a directory of Kubernetes configuration. It maintains the original directory structure as read by source functions. |         | [Source](https://github.com/GoogleContainerTools/kpt-functions-sdk/blob/master/ts/demo-functions/src/write_yaml.ts) | [Typescript SDK](../../../producer/functions/ts/) |

## Validators

| Image                                                | Args | Description                                                                                                                             | Example                                                                                                         | Source                                                                                                                          | Toolchain                                         |
| ---------------------------------------------------- | ---- | --------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------- |
| gcr.io/kpt-functions/istioctl-analyze                |      | Istioctl analyze is a diagnostic tool that can detect potential issues with Istio configuration and output errors to the results field. | [Example](https://github.com/GoogleContainerTools/kpt-functions-catalog/tree/master/examples/istioctl-analyze/) | [Source](https://github.com/GoogleContainerTools/kpt-functions-catalog/blob/master/functions/ts/src/istioctl_analyze.ts)        | [Typescript SDK](../../../producer/functions/ts/) |
| gcr.io/kpt-functions/gatekeeper-validate             |      | Enforces OPA constraints on input objects. The constraints are also passed as part of the input to the function.                        |                                                                                                                 | [Source](https://github.com/GoogleContainerTools/kpt-functions-sdk/blob/master/go/pkg/functions/gatekeeper/validate.go)         |                                                   |
| gcr.io/kpt-functions/validate-rolebinding            |      | [Demo] Enforces a blacklist of subjects in RoleBinding objects.                                                                         |                                                                                                                 | [Source](https://github.com/GoogleContainerTools/kpt-functions-sdk/blob/master/ts/demo-functions/src/validate_rolebinding.ts)   | [Typescript SDK](../../../producer/functions/ts/) |
| gcr.io/kpt-functions/kubeval                         |      | Validates configuration using kubeval.                                                                                                  |                                                                                                                 | [Source](https://github.com/GoogleContainerTools/kpt-functions-catalog/blob/master/functions/ts/src/kubeval.ts)                 | [Typescript SDK](../../../producer/functions/ts/) |
| gcr.io/kustomize-functions/example-validator-kubeval |      | [Demo] Validates that all containers have cpu and memory reservations set.                                                              |                                                                                                                 | [Source](https://github.com/kubernetes-sigs/kustomize/blob/master/functions/examples/validator-kubeval/image/main.go)           | [Go Library](../../../producer/functions/golang/) |
| gcr.io/kustomize-functions/example-validator         |      | Validates Kubernetes configuration files using schemas from the Kubernetes OpenAPI spec.                                                |                                                                                                                 | [Source](https://github.com/kubernetes-sigs/kustomize/blob/master/functions/examples/validator-resource-requests/image/main.go) | [Go Library](../../../producer/functions/golang/) |
| gcr.io/kpt-functions/suggest-psp                     |      | [Demo] Lints PodSecurityPolicy by suggesting 'spec.allowPrivilegeEscalation' field be set to 'false'.                                   |                                                                                                                 | [Source](https://github.com/GoogleContainerTools/kpt-functions-sdk/blob/master/ts/demo-functions/src/suggest_psp.ts)            | [Typescript SDK](../../../producer/functions/ts/) |

## Generators

| Image                                          | Args | Description                                                                                          | Example                                                                                                      | Source                                                                                                                                           | Toolchain                                         |
| ---------------------------------------------- | ---- | ---------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------- |
| gcr.io/kpt-functions/expand-team-cr            |      | [Demo] Reads custom resources of type Team and generates multiple Namespace and RoleBinding objects. |                                                                                                              | [Source](https://github.com/GoogleContainerTools/kpt-functions-sdk/blob/master/ts/demo-functions/src/expand_team_cr.ts)                          | [Typescript SDK](../../../producer/functions/ts/) |
| gcr.io/kpt-functions/helm-template             |      | Render chart templates locally using helm template.                                                  | [Example](https://github.com/GoogleContainerTools/kpt-functions-catalog/tree/master/examples/helm-template/) | [Source](https://github.com/GoogleContainerTools/kpt-functions-catalog/blob/master/functions/ts/src/helm_template.ts)                            | [Typescript SDK](../../../producer/functions/ts/) |
| gcr.io/kustomize-functions/example-nginx       |      | Generate configuration from go templates using the functionConfig as the template input.             |                                                                                                              | [Source](https://github.com/kubernetes-sigs/kustomize/blob/master/functions/examples/template-go-nginx/image/main.go)                            | [Go Library](../../../producer/functions/golang/) |
| gcr.io/kustomize-functions/example-cockroachdb |      | Generate configuration from heredoc template using the functionConfig as the template input.         |                                                                                                              | [Source](https://github.com/kubernetes-sigs/kustomize/blob/master/functions/examples/template-heredoc-cockroachdb/image/cockroachdb-template.sh) |                                                   |

## Transformers

| Image                                                        | Args | Description                                                                                  | Example                                                                                                      | Source                                                                                                                     | Toolchain                                                 |
| ------------------------------------------------------------ | ---- | -------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------ | -------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------- |
| gcr.io/kpt-functions/mutate-psp                              |      | [Demo] Mutates PodSecurityPolicy objects by setting spec.allowPrivilegeEscalation to false.  |                                                                                                              | [Source](https://github.com/GoogleContainerTools/kpt-functions-sdk/blob/master/ts/demo-functions/src/mutate_psp.ts)        | [Typescript SDK](../../../producer/functions/ts/)         |
| N/A                                                          |      | Sets the namespace field of all configs passed in.                                           |                                                                                                              | [Source](https://github.com/GoogleContainerTools/kpt-functions-catalog/blob/master/functions/starlark/set_namespace.star)  | [Starlark Runtime](../../../producer/functions/starlark/) |
| gcr.io/kpt-functions/set-namespace                           |      | Sets the namespace field of all configs passed in.                                           |                                                                                                              | [Source](https://github.com/GoogleContainerTools/kpt-functions-catalog/blob/master/functions/go/set-namespace/main.go)     | [Go Library](../../../producer/functions/golang/)         |
| gcr.io/kubeflow-images-public/kustomize-fns/remove-namespace |      | Removes the namespace field of cluster-scoped configs.                                       |                                                                                                              | [Source](https://github.com/kubeflow/kfctl/tree/master/kustomize-fns/remove-namespace/)                                    | [Go Library](../../../producer/functions/golang/)         |
| gcr.io/kpt-functions/label-namespace                         |      | [Demo] Adds a label to all Namespaces.                                                       |                                                                                                              | [Source](https://github.com/GoogleContainerTools/kpt-functions-sdk/blob/master/ts/hello-world/src/label_namespace.ts)      | [Typescript SDK](../../../producer/functions/ts/)         |
| gcr.io/kpt-functions/helm-template                           |      | Render chart templates locally using helm template.                                          | [Example](https://github.com/GoogleContainerTools/kpt-functions-catalog/tree/master/examples/helm-template/) | [Source](https://github.com/GoogleContainerTools/kpt-functions-catalog/blob/master/functions/ts/src/helm_template.ts)      | [Typescript SDK](../../../producer/functions/ts/)         |
| gcr.io/kustomize-functions/example-tshirt                    |      | Sets cpu and memory reservations on all containers for Resources annotated with tshirt size. |                                                                                                              | [Source](https://github.com/kubernetes-sigs/kustomize/blob/master/functions/examples/injection-tshirt-sizes/image/main.go) | [Go Library](../../../producer/functions/golang/)         |
| gcr.io/kpt-functions/annotate-config                         |      | [Demo] Adds an annotation to all configs.                                                    |                                                                                                              | [Source](https://github.com/GoogleContainerTools/kpt-functions-sdk/blob/master/ts/demo-functions/src/annotate_config.ts)   | [Typescript SDK](../../../producer/functions/ts/)         |

## Miscellaneous

| Image                      | Args | Description                                | Example | Source                                                                                                         | Toolchain                                         |
| -------------------------- | ---- | ------------------------------------------ | ------- | -------------------------------------------------------------------------------------------------------------- | ------------------------------------------------- |
| gcr.io/kpt-functions/no-op |      | [Demo] No Op function for testing purposes |         | [Source](https://github.com/GoogleContainerTools/kpt-functions-sdk/blob/master/ts/demo-functions/src/no_op.ts) | [Typescript SDK](../../../producer/functions/ts/) |

## Next Steps

- Learn more ways of using the kpt fn command from the [reference] doc.
- Get a quickstart on writing functions from the [function producer docs].

[spec]: https://github.com/kubernetes-sigs/kustomize/blob/master/cmd/config/docs/api-conventions/functions-spec.md
[Functions User Guide]: ../
[ts sdk]: ../../../producer/functions/ts/
[go libs]: ../../../producer/functions/golang/
[source]: ../../../../concepts/functions/#source-function
[sink]: ../../../../concepts/functions/#sink-function
[reference]: ../../../../reference/fn/run/
[function producer docs]: ../../../producer/functions/