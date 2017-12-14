// Automatically generated by swagger-doc. DO NOT EDIT!

package v1

func (VirtualMachine) SwaggerDoc() map[string]string {
	return map[string]string{
		"":       "VirtualMachine is *the* VM Definition. It represents a virtual machine in the runtime environment of kubernetes.\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object",
		"spec":   "VM Spec contains the VM specification.",
		"status": "Status is the high level overview of how the VM is doing. It contains information available to controllers and users.",
	}
}

func (VirtualMachineList) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "VirtualMachineList is a list of VirtualMachines\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object",
	}
}

func (VirtualMachineSpec) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                              "VirtualMachineSpec is a description of a VirtualMachine.",
		"domain":                        "Specification of the desired behavior of the VirtualMachine on the host.",
		"nodeSelector":                  "NodeSelector is a selector which must be true for the vm to fit on a node.\nSelector which must match a node's labels for the vm to be scheduled on that node.\nMore info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/\n+optional",
		"affinity":                      "If affinity is specifies, obey all the affinity rules",
		"terminationGracePeriodSeconds": "Grace period observed after signalling a VM to stop after which the VM is force terminated.",
		"volumes":                       "List of volumes that can be mounted by disks belonging to the vm.",
		"interfaces":                    "List of interfaces/networks which can be attached to an interface inside the vm.",
	}
}

func (Affinity) SwaggerDoc() map[string]string {
	return map[string]string{
		"":             "Affinity groups all the affinity rules related to a VM",
		"nodeAffinity": "Node affinity support",
	}
}

func (VirtualMachineStatus) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                  "VirtualMachineStatus represents information about the status of a VM. Status may trail the actual\nstate of a system.",
		"nodeName":          "NodeName is the name where the VM is currently running.",
		"migrationNodeName": "MigrationNodeName is the node where the VM is live migrating to.",
		"conditions":        "Conditions are specific points in VM's pod runtime.",
		"phase":             "Phase is the status of the VM in kubernetes world. It is not the VM status, but partially correlates to it.",
		"graphics":          "Graphics represent the details of available graphical consoles.",
	}
}

func (VirtualMachineGraphics) SwaggerDoc() map[string]string {
	return map[string]string{}
}

func (VirtualMachineCondition) SwaggerDoc() map[string]string {
	return map[string]string{}
}

func (Spice) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object",
	}
}

func (SpiceInfo) SwaggerDoc() map[string]string {
	return map[string]string{}
}

func (Migration) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "A Migration is a job that moves a Virtual Machine from one node to another\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object",
	}
}

func (MigrationSpec) SwaggerDoc() map[string]string {
	return map[string]string{
		"":             "MigrationSpec is a description of a VM Migration\nFor example \"destinationNodeName\": \"testvm\" will migrate a VM called \"testvm\" in the namespace \"default\"",
		"selector":     "Criterias for selecting the VM to migrate.\nFor example\nselector:\n  name: testvm\nwill select the VM `testvm` for migration",
		"nodeSelector": "Criteria to use when selecting the destination for the migration\nfor example, to select by the hostname, specify `kubernetes.io/hostname: master`\nother possible choices include the hardware required to run the vm or\nor lableing of the nodes to indicate their roles in larger applications.\nexamples:\ndisktype: ssd,\nrandomGenerator: /dev/random,\nrandomGenerator: superfastdevice,\napp: mysql,\nlicensedForServiceX: true\nNote that these selectors are additions to the node selectors on the VM itself and they must not exist on the VM.\nIf they are conflicting with the VM, no migration will be started.",
	}
}

func (VMSelector) SwaggerDoc() map[string]string {
	return map[string]string{
		"name": "Name of the VM to migrate",
	}
}

func (MigrationStatus) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "MigrationStatus is the last reported status of a VM Migratrion. Status may trail the actual\nstate of a migration.",
	}
}

func (MigrationList) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "A list of Migrations\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object",
	}
}

func (MigrationHostInfo) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "Host specific data, used by the migration controller to fetch host specific migration information from the target host",
	}
}

func (VirtualMachineReplicaSet) SwaggerDoc() map[string]string {
	return map[string]string{
		"":       "VM is *the* VM Definition. It represents a virtual machine in the runtime environment of kubernetes.\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object",
		"spec":   "VM Spec contains the VM specification.",
		"status": "Status is the high level overview of how the VM is doing. It contains information available to controllers and users.",
	}
}

func (VirtualMachineReplicaSetList) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "VMList is a list of VMs\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object",
	}
}

func (VMReplicaSetSpec) SwaggerDoc() map[string]string {
	return map[string]string{
		"replicas": "Number of desired pods. This is a pointer to distinguish between explicit\nzero and not specified. Defaults to 1.\n+optional",
		"selector": "Label selector for pods. Existing ReplicaSets whose pods are\nselected by this will be the ones affected by this deployment.\n+optional",
		"template": "Template describes the pods that will be created.",
		"paused":   "Indicates that the replica set is paused.\n+optional",
	}
}

func (VMReplicaSetStatus) SwaggerDoc() map[string]string {
	return map[string]string{
		"replicas":      "Total number of non-terminated pods targeted by this deployment (their labels match the selector).\n+optional",
		"readyReplicas": "The number of ready replicas for this replica set.\n+optional",
	}
}

func (VMReplicaSetCondition) SwaggerDoc() map[string]string {
	return map[string]string{}
}

func (VMTemplateSpec) SwaggerDoc() map[string]string {
	return map[string]string{
		"spec": "VM Spec contains the VM specification.",
	}
}
