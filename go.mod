module my-sample-operator

go 1.15

require (
	k8s.io/api v0.0.0-20201114085527-4a626d306b98
	k8s.io/apimachinery v0.0.0-20201114085355-859536f6dc9b
	k8s.io/client-go v0.0.0-20201114085741-77eda6a9395b
	k8s.io/code-generator v0.0.0-20201114085218-0843e0f44fb3
	k8s.io/klog v1.0.0
	k8s.io/klog/v2 v2.4.0
)

// replace k8s.io/client-go v0.0.0-20190425172711-65184652c889 => k8s.io/client-go v0.0.0-20201114085741-77eda6a9395b
