package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	GroupName = "glue.solo.io"
	Version   = "v1"
)

var KnownCRDs = []crd{
	{
		Plural:  "routes",
		Group:   GroupName,
		Version: Version,
		Kind:    "Route",
	},
	{
		Plural:  "upstreams",
		Group:   GroupName,
		Version: Version,
		Kind:    "Upstream",
	},
	{
		Plural:  "virtualhosts",
		Group:   GroupName,
		Version: Version,
		Kind:    "VirtualHost",
	},
}

type crd struct {
	Plural  string
	Group   string
	Version string
	Kind    string
}

func (d crd) FullName() string {
	return d.Plural + "." + d.Group
}

// SchemeGroupVersion is group version used to register these objects
var SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: Version}

// Kind takes an unqualified kind and returns back a Group qualified GroupKind
func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

var (
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	AddToScheme   = SchemeBuilder.AddToScheme
)

// Adds the list of known types to Scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&Route{},
		&RouteList{},
		&Upstream{},
		&UpstreamList{},
		&VirtualHost{},
		&VirtualHostList{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
