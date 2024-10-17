package asciiart

import (
	"os"
	"strings"
)

// check if the character is printable
func Checkchars(s string) bool {
	for _, c := range s {
		if c < 32 || c > 126 {
			return false
		}
	}
	return true
}

// mapping
func MapBanner(filename string) (map[rune][]string, error) {
	data, err := os.ReadFile("banners/" + filename + ".txt")
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(data), "\n")

	banner := make(map[rune][]string)
	charIndex := 32
	for i := 0; i < len(lines); i += 9 {
		if i+8 < len(lines) {
			banner[rune(charIndex)] = lines[i+1 : i+9]
			charIndex++
		}
	}
	return banner, nil
}

// check there is newline
func Checknewline(inpultsplit []string) bool {
	for _, line := range inpultsplit {
		if len(line) != 0 {
			return false
		}
	}
	return true
}

func CheckInput(input string) string {
	var output []byte
	for i := 0; i < len(input); i++ {
		char := input[i]
		// Allow only printable ASCII characters or line breaks
		if (char >= 32 && char <= 126) || char == '\r' || char == '\n' {
			output = append(output, char)
		}
	}
	// Convert the byte slice back to a string
	return string(output)
}

// drawing the result
func Draw(banner map[rune][]string, inpultsplit []string) string {
	var output string
	for idx, v := range inpultsplit {
		if Checknewline(inpultsplit) && idx != len(inpultsplit)-1 {
			output += "\n"
			continue
		} else if len(v) == 0 && !Checknewline(inpultsplit) {
			output += "\n"
		} else if len(v) != 0 && !Checknewline(inpultsplit) {
			for i := 0; i < 8; i++ {
				for j := 0; j < len(v); j++ {
					output += (banner[rune(v[j])][i])
				}
				output += "\n"
			}
		}
	}
	return output
}
