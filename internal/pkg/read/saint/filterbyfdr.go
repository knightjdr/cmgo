package saint

// FilterByFDR removes preys not passing an FDR.
func (s *SAINT) FilterByFDR(fdr float64) {
	for i := len(*s) - 1; i >= 0; i-- {
		if (*s)[i].FDR > fdr {
			*s = append((*s)[:i], (*s)[i+1:]...)
		}
	}
}
