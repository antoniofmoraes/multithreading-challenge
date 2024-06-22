package internal

import "regexp"

func RemoveNonDigits(s string) string {
	return regexp.MustCompile(`[^0-9]`).ReplaceAllString(s, "")
	// regex := regexp.MustCompile(`[0-9]+`)
	// return strings.Join(regex.FindAllString(res.Cep, -1), "")
}
