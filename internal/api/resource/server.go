package resource

type ServerResource struct {
	SupportedBillingTypes     []string
	SupportedOperatingSystems []string
	SupportedPlans            []string
	SupportedSites            []string
	SupportedRAIDLevels       []string
}

func NewServerResource() *ServerResource {
	return &ServerResource{
		SupportedBillingTypes: []string{
			"SKIP",
			"hourly",
			"monthly",
			"yearly",
		},
		SupportedOperatingSystems: []string{
			"ipxe",
			"windows_server_2019_std_v1",
			"ubuntu_22_04_x64_lts",
			"debian_11",
			"rockylinux_8",
			"debian_10",
			"rhel8",
			"centos_7_4_x64",
			"centos_8_x64",
			"ubuntu_20_04_x64_lts",
			"debian_12",
			"ubuntu22_ml_in_a_box",
			"windows2022",
		},
		SupportedPlans: []string{
			"c2-large-x86",
			"c2-medium-x86",
			"c2-small-x86",
			"c3-large-x86",
			"c3-medium-x86",
			"c3-small-x86",
			"c3-xlarge-x86",
			"g3-large-x86",
			"g3-medium-x86",
			"g3-small-x86",
			"g3-xlarge-x86",
			"m3-large-x86",
			"s2-small-x86",
			"s3-large-x86",
		},
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
		SupportedRAIDLevels: []string{
			"SKIP",
			"raid-0",
			"raid-1",
		},
	}
}
