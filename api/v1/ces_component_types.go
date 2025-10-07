package v1

import (
	"fmt"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.
const (
	// ComponentStatusNotInstalled represents a status for a component that is not installed
	ComponentStatusNotInstalled = ""
	// ComponentStatusInstalling represents a status for a component that is currently being installed
	ComponentStatusInstalling = "installing"
	// ComponentStatusUpgrading represents a status for a component that is currently being upgraded
	ComponentStatusUpgrading = "upgrading"
	// ComponentStatusDeleting represents a status for a component that is currently being deleted
	ComponentStatusDeleting = "deleting"
	// ComponentStatusInstalled represents a status for a component that was successfully installed
	ComponentStatusInstalled = "installed"
	// ComponentStatusTryToInstall represents a status for a component that is not installed but its install process is in requeue loop.
	ComponentStatusTryToInstall = "tryToInstall"
	// ComponentStatusTryToUpgrade represents a status for a component that is installed but its actual upgrade process is in requeue loop.
	// In this state the component can be healthy but the version in the spec is not installed.
	ComponentStatusTryToUpgrade = "tryToUpgrade"
	// ComponentStatusTryToDelete represents a status for a component that is installed but its delete process is in requeue loop.
	// In this state the component can be healthy.
	ComponentStatusTryToDelete = "tryToDelete"
)

const FinalizerName = "component-finalizer"

const (
	ComponentNameLabelKey    = "k8s.cloudogu.com/component.name"
	ComponentVersionLabelKey = "k8s.cloudogu.com/component.version"
)

// ComponentSpec defines the desired state of a component.
type ComponentSpec struct {
	// Namespace of the component (e.g. k8s)
	Namespace string `json:"namespace,omitempty"`
	// Name of the component (e.g. k8s-dogu-operator)
	Name string `json:"name,omitempty"`
	// Desired version of the component (e.g. 2.4.48-3)
	Version string `json:"version,omitempty"`
	// Mapped values from component cr to specific component attributes, usually log levels
	MappedValues map[string]string `json:"mappedValues,omitempty"`
	// DeployNamespace is the namespace where the helm chart should be deployed in.
	// This value is optional. If it is empty the operator deploys the helm chart in the namespace where the operator is deployed.
	DeployNamespace string `json:"deployNamespace,omitempty"`
	// ValuesYamlOverwrite is a multiline-yaml string that is applied alongside the original values.yaml-file of the component.
	// It can be used to overwrite specific configurations. Lists are overwritten, maps are merged.
	// +optional
	ValuesYamlOverwrite string `json:"valuesYamlOverwrite,omitempty"`
}

type HealthStatus string

const (
	PendingHealthStatus     HealthStatus = ""
	AvailableHealthStatus   HealthStatus = "available"
	UnavailableHealthStatus HealthStatus = "unavailable"
	UnknownHealthStatus     HealthStatus = "unknown"
	mappingMetadataFileName              = "component-values-metadata.yaml"
)

// ComponentStatus defines the observed state of a Component.
type ComponentStatus struct {
	// Status represents the state of the component in the ecosystem.
	Status string `json:"status"`
	// RequeueTimeNanos contains the time in nanoseconds to wait until the next requeue.
	RequeueTimeNanos time.Duration `json:"requeueTimeNanos,omitempty"`
	// Health describes the health status of the component.
	// A component becomes 'available' if its Status is 'installed',
	// and all its deployments, stateful sets, and daemon sets are available.
	Health HealthStatus `json:"health,omitempty"`
	// Installed version of the component (e.g. 2.4.48-3)
	InstalledVersion string `json:"installedVersion,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels=app=ces;app.kubernetes.io/name=k8s-component-operator;k8s.cloudogu.com/component.name=k8s-component-operator-crd
// +kubebuilder:resource:shortName="comp"
// +kubebuilder:printcolumn:name="Spec-Version",type="string",JSONPath=".spec.version",description="The desired version of the component"
// +kubebuilder:printcolumn:name="Installed Version",type="string",JSONPath=".status.installedVersion",description="The current version of the component"
// +kubebuilder:printcolumn:name="Health",type="string",JSONPath=".status.health",description="The current health state of the component"
// +kubebuilder:printcolumn:name="Status",type="string",JSONPath=".status.status",description="The current status of the component"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp",description="The age of the component"

// Component is the Schema for the ces component API
type Component struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComponentSpec   `json:"spec,omitempty"`
	Status ComponentStatus `json:"status,omitempty"`
}

type Mapping struct {
	Path    string            `yaml:"path"`
	Mapping map[string]string `yaml:"Mapping"`
}

type MetaValue struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Keys        []Mapping
}

type MetadataMapping struct {
	ApiVersion string               `yaml:"apiVersion"`
	Metavalues map[string]MetaValue `yaml:"metavalues"`
}

// String returns a string representation of this component.
func (c *Component) String() string {
	return fmt.Sprintf("%s/%s:%s", c.Spec.Namespace, c.Spec.Name, c.Spec.Version)
}

// +kubebuilder:object:root=true

// ComponentList contains a list of Component
type ComponentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Component `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Component{}, &ComponentList{})
}
