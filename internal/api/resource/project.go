package resource

type ProjectResource struct {
	SupportedEnvironments      []string
	SupportedProvisioningTypes []string
}

func NewProjectResource() *ProjectResource {
	return &ProjectResource{
		SupportedEnvironments: []string{
			"SKIP",
			"Development",
			"Staging",
			"Production",
		},
		SupportedProvisioningTypes: []string{
			"on_demand",
			"reserved",
		},
	}
}
