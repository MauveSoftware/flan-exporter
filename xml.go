package main

type NmapRun struct {
	Hosts []HostResult `xml:"host"`
}

type HostResult struct {
	Ports  PortsResult `xml:"ports"`
	Status struct {
		State string `xml:"state,attr"`
	} `xml:"status"`
}

type PortsResult struct {
	Ports []PortResult `xml:"port"`
}

type PortResult struct {
	Protocol string `xml:"protocol,attr"`
	Number   string `xml:"portid,attr"`
	State    struct {
		State string `xml:"state,attr"`
	} `xml:"state"`
	Script ScriptResult `xml:"script"`
}

type ScriptResult struct {
	ID    string `xml:"id,attr"`
	Table struct {
		Tables []Table `xml:"table"`
	} `xml:"table"`
}

type Table struct {
	Elements []TableElement `xml:"elem"`
}

type TableElement struct {
	Key  string `xml:"key,attr"`
	Text string `xml:",chardata"`
}
