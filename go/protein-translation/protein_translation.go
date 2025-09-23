package protein

import (
	"errors"
)

var codonMap = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine", "UUC": "Phenylalanine",
	"UUA": "Leucine", "UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine", "UCA": "Serine", "UCG": "Serine",
	"UAU": "Tyrosine", "UAC": "Tyrosine",
	"UGU": "Cysteine", "UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP", "UAG": "STOP", "UGA": "STOP",
}

var (
	ErrStop        = errors.New("stop codon")
	ErrInvalidBase = errors.New("invalid codon")
)

const codonLength = 3

func FromRNA(rna string) ([]string, error) {
	codons := make([]string, 0, len(rna)/codonLength)
	for i := 0; i < len(rna); i += codonLength {
		if i+codonLength > len(rna) {
			break
		}
		codon := rna[i : i+codonLength]
		aminoAcid, err := FromCodon(codon)
		if err != nil {
			if errors.Is(err, ErrStop) {
				break
			}
			return nil, err
		}
		codons = append(codons, aminoAcid)
	}
	return codons, nil

}

func FromCodon(codon string) (string, error) {
	aminoAcid, ok := codonMap[codon]
	if !ok {
		return "", ErrInvalidBase
	}
	if aminoAcid == "STOP" {
		return "", ErrStop
	}
	return aminoAcid, nil
}
