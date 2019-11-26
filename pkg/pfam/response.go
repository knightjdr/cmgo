package pfam

// Regions contains sequence level domain and region information.
type Regions struct {
	Domains []Domain `json:"regions"`
	Motifs  []Motif  `json:"motifs"`
}

// Domain contains protein domain information.
type Domain struct {
	End      int      `json:"end"`
	Metadata Metadata `json:"metadata"`
	Name     string
	Start    int `json:"start"`
}

// Metadata for domain.
type Metadata struct {
	Identified string `json:"identifier"`
}

// Motif contains protein motif/region information.
type Motif struct {
	End   int    `json:"end"`
	Name  string `json:"type"`
	Start int    `json:"start"`
}

// Response for data received from Pfam.
type Response []Regions

// AddDomainNames attaches the metadata identfier to the Domain struct.
func (r *Response) AddDomainNames() {
	for i, domain := range (*r)[0].Domains {
		(*r)[0].Domains[i].Name = domain.Metadata.Identified
	}
}
