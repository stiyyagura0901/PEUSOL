package helpers

import "github.com/PEUSOL/problems"

// GetFuncs ...
func GetFuncs() map[string]interface{} {
	funcs := map[string]interface{}{
		"One": problems.One,
		"Two": problems.Two,
		"Three": problems.Three,
		"Four": problems.Four,
		"Five": problems.Five,
		"Six": problems.Six,
		"Seven": problems.Seven,
		"Eight": problems.Eight,
		"Nine": problems.Nine,
		"Ten": problems.Ten,
		"Eleven": problems.Eleven,
		"Twelve": problems.Twelve,
		"OneHundredFortyFive": problems.OneHundredFortyFive,
	}

	return funcs
}
