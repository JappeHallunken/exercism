package ledger

import (
	"fmt"
	"sort"
	"strings"
)

type Entry struct {
	Date, Description string
	Change            int
}

func FormatLedger(currency, locale string, entries []Entry) (string, error) {
	if currency == "" {
		return "", fmt.Errorf("unsupported currency: empty")
	}
	symbols := map[string]string{"USD": "$", "EUR": "€"}
	symbol, ok := symbols[currency]

	if !ok {
		return "", fmt.Errorf("unsupported currency: %q", currency)
	}

	entriesCopy := make([]Entry, len(entries))
	copy(entriesCopy, entries)

	if len(entriesCopy) == 0 {
		entriesCopy = []Entry{}
	}

	sort.Slice(entriesCopy, func(i, j int) bool {
		if entriesCopy[i].Date != entriesCopy[j].Date {
			return entriesCopy[i].Date < entriesCopy[j].Date
		}
		if entriesCopy[i].Description != entriesCopy[j].Description {
			return entriesCopy[i].Description < entriesCopy[j].Description
		}
		return entriesCopy[i].Change < entriesCopy[j].Change
	})

	var header string
	switch locale {
	case "nl-NL":
		header = fmt.Sprintf("%-10s | %-25s | %s\n", "Datum", "Omschrijving", "Verandering")
	case "en-US":
		header = fmt.Sprintf("%-10s | %-25s | %s\n", "Date", "Description", "Change")
	default:
		return "", fmt.Errorf("unsupported locale: %s", locale)
	}

	var sb strings.Builder
	sb.WriteString(header)

	for _, e := range entriesCopy {
		line, err := formatEntry(e, symbol, locale)
		if err != nil {
			return "", err
		}
		sb.WriteString(line)
	}

	return sb.String(), nil
}


func formatEntry(e Entry, symbol, locale string) (string, error) {
	if len(e.Date) != 10 || e.Date[4] != '-' || e.Date[7] != '-' {
		return "", fmt.Errorf("invalid date format: %q", e.Date)
	}
	y, m, d := e.Date[:4], e.Date[5:7], e.Date[8:]

	var date string
	switch locale {
	case "nl-NL":
		date = fmt.Sprintf("%s-%s-%s", d, m, y)
	case "en-US":
		date = fmt.Sprintf("%s/%s/%s", m, d, y)
	default:
		return "", fmt.Errorf("unsupported locale: %s", locale)
	}

	desc := e.Description
	if len(desc) > 25 {
		desc = desc[:22] + "..."
	} else {
		desc += strings.Repeat(" ", 25-len(desc))
	}

	amount := formatAmount(e.Change, symbol, locale)

	return fmt.Sprintf("%-10s | %s | %13s\n", date, desc, amount), nil
}

func formatAmount(cents int, symbol, locale string) string {
	neg := cents < 0
	if neg {
		cents = -cents
	}
	val := fmt.Sprintf("%03d", cents)
	intPart, frac := val[:len(val)-2], val[len(val)-2:]

	var parts []string
	for len(intPart) > 3 {
		parts = append([]string{intPart[len(intPart)-3:]}, parts...)
		intPart = intPart[:len(intPart)-3]
	}
	if intPart != "" {
		parts = append([]string{intPart}, parts...)
	}

	switch locale {
	case "nl-NL":
		amount := fmt.Sprintf("%s %s,%s", symbol, strings.Join(parts, "."), frac)
		if neg {
			amount += "-"
		} else {
			amount += " "
		}
		return amount
	case "en-US":
		amount := fmt.Sprintf("%s%s.%s", symbol, strings.Join(parts, ","), frac)
		if neg {
			return fmt.Sprintf("(%s)", amount)
		}
		return amount + " "
	default:
		return ""
	}
}
