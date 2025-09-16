package strand

func ToRNA(dna string) string {
	rna := ""
	for _, nucleotide := range dna {
		switch nucleotide {
		case 'G':
			rna += "C"
		case 'C':
			rna += "G"
		case 'T':
			rna += "A"
		case 'A':
			rna += "U"
		default:
			return ""
		}
	}
	return rna
}
