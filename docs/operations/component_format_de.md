# Dogu format

Die Component-CR kann genutzt werden, um Cloudogu-Komponenten in einem Kubernetescluster mit dem Component-Operator zu installieren.

Folgend werden alle Felder einer Component-CR beschrieben.

## Beispiel

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

- `.metadata.name`: Der Komponentenname der Kubernetes-Resource. Dieser muss identisch mit `.spec.name` sein.
- `.spec.name`: Der Komponentenname, wie er in der Helm-Registry lautet. Dieser muss identisch mit `.metadata.name` sein.
- `.spec.namespace`: Der Namespace der Komponente in der Helm-Registry.
    - Mittels unterschiedlicher Komponenten-Namespaces können unterschiedliche Versionen ausgebracht werden (z. B. zu Debugging-Zwecken).
    - Es handelt sich hierbei _nicht_ um den Cluster-Namespace.
- `.spec.version`: Die Version der Komponente in der Helm-Registry.
- `.spec.deployNamespace`: (optional) Der k8s-Namespace, in dem alle Ressourcen der Komponente deployed werden sollen. Wenn dieser leer ist, wird der Namespace des Komponenten-Operators verwendet.
- `.spec.mappedValues`: (optional) Helm-Werte zum Überschreiben von Konfigurationen aus der Helm-Datei values.yaml. Diese Werte werden durch die Konfiguration in component-values-metadata.yaml gemappt.
- `.spec.valuesYamlOverwrite`: (optional) Helm-Werte zum Überschreiben von Konfigurationen aus der Helm-Datei values.yaml. Sollte aus Gründen der Lesbarkeit als [multiline-yaml](https://yaml-multiline.info/) geschrieben werden.

> [!WARNING]
> `.spec.mappedValues` und `.spec.valuesYamlOverwrite` sollten nicht gleichzeitig verwendet werden. Sind beide Werte konfiguriert, so bekommen die mappedValues den Vorzug.

> [!WARNING]
> `.spec.mappedValues` und `.spec.valuesYamlOverwrite` dürfen keine Listeneinträge überschreiben. Es ist durch die Struktur von Yaml nicht möglich einzelne Elemente innerhalb einer Liste zu setzen.
>  Es kann immer nur die gesamte Liste überschrieben werden.
