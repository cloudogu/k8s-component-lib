# Dogu format

The Component-CR can be used to install Cloudogu components in a Kubernetes cluster with the Component Operator.

All fields of a Component CR are described below.

## Example

```yaml
apiVersion: k8s.cloudogu.com/v1
kind: Component
metadata:
  name: k8s-service-discovery
spec:
  name: k8s-service-discovery
  namespace: k8s
  version: 3.0.0
  deployNamespace: ecosystem
  mappedValues:
    mainLogLevel: debug
  valuesYamlOverwrite: |
    loadBalancerService:
      externalTrafficPolicy: Cluster
```

- `.metadata.name`: The component name of the Kubernetes resource. This must be identical to `.spec.name`.
- `.spec.name`: The component name as it appears in the Helm registry. This must be identical to `.metadata.name`.
- `.spec.namespace`: The component namespace in the helm registry.
    - Using different component namespaces, different versions could be deployed (e.g. for debugging purposes).
    - This is _not_ the cluster namespace.
- `.spec.version`: The version of the component in the helm registry.
- `.spec.deployNamespace`: (optional) The k8s-namespace, where all resources of the component should be deployed. If this is empty the namespace of the component-operator will be used.
- `.spec.mappedValues`: (optional) Helm values used to override configurations from the Helm `values.yaml` file. These values are mapped according to the configuration defined in the `component-values-metadata.yaml` file.
- `.spec.valuesYamlOverwrite`: (optional) Helm-Values to overwrite configurations of the default values.yaml file. Should be written as a [multiline-yaml](https://yaml-multiline.info/) string for readability.

> [!WARNING]
> `.spec.mappedValues` and `.spec.valuesYamlOverwrite` should not be used at the same time.  
> If both values are configured, `mappedValues` will take precedence.

> [!WARNING]
> `.spec.mappedValues` and `.spec.valuesYamlOverwrite` must not overwrite list entries.
> Due to the structure of YAML, it is not possible to set individual elements within a list.
> Only the entire list can ever be overwritten.
