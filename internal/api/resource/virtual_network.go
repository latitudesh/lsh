package resource

type VirtualNetworkResource struct {
	SupportedSites []string
}

func NewVirtualNetworkResource() *VirtualNetworkResource {
	return &VirtualNetworkResource{
		SupportedSites: []string{
			"ASH",
			"BGT",
			"BUE",
			"CHI",
			"DAL",
			"FRA",
			"LAX",
			"LON",
			"MEX",
			"MEX2",
			"MIA",
			"MIA2",
			"NYC",
			"SAN",
			"SAN2",
			"SAO",
			"SAO2",
			"SYD",
			"TYO",
			"TYO2",
		},
	}
}
