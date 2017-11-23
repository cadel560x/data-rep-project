package eliza

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Reflect my own attempt to reflect 'you' object pronoun to 'I' subject pronoun based on prepositions found
func Reflect(input string) string {
	// Split the input on word boundaries.
	// boundaries := regexp.MustCompile(`\b`)
	boundaries := regexp.MustCompile(`(?=\S*['-])([a-zA-Z'-]+)`)
	tokens := boundaries.Split(input, -1)

	// Some key prepositions
	prepositions := []string{
		"to",
		"by",
		"under",
		"about",
		"on",
		"according",
		"over",
		"of",
		"without",
	}

	// List the reflections.
	reflections := [][]string{
		{`was`, `were`},
		{`I`, `you`},
		{`I'm`, `you are`},
		{`I'd`, `you would`},
		{`I've`, `you have`},
		{`I'll`, `you will`},
		{`my`, `your`},
		{`you're`, `I am`},
		{`were`, `was`},
		{`you've`, `I have`},
		{`you'll`, `I will`},
		{`your`, `my`},
		{`yours`, `mine`},
		// {`you`, `me`},
		{`me`, `you`},
	}

	// Loop through each token, reflecting it if there's a match.
	for i, token := range tokens {
		for _, reflection := range reflections {

			// Let's take 'you' separately because it is the same as subject pronoun as object pronoun
			if token == "you" {

				// Loop through the prepositions
				for j, preposition := range prepositions {
					// Compare the previous word, that is 'token[i-2]' to the 'preposition'. 'token[i-1]' is the space character.
					if tokens[i-2] == preposition {
						// If 'you' is an object pronoun to a preposition, the swap it for 'me'
						tokens[i] = "me"
						break
					}

					// The previous word was not a preposition, swapping for a subject pronoun
					if j == len(prepositions)-1 {
						tokens[i] = "I"
					}

				} // for j, prepostition
				// As for the rest of reflections, keep doing the normal substitution
			} else if matched, _ := regexp.MatchString(reflection[0], token); matched {
				tokens[i] = reflection[1]
				break
			} // if - else if

		} // for 'reflection'
	} // for 'i'

	// Put the tokens back together.
	return strings.Join(tokens, ``)

} // Reflect

// Eliza chatbot engine
// As seen in https://github.com/data-representation/eliza teaching aid

// Replacer data structure of regexs and responses
type Replacer struct {
	original     *regexp.Regexp
	replacements []string
}

// ReadReplacersFromFile reads regexs and responses from file
func ReadReplacersFromFile(path string) []Replacer {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var replacers []Replacer

	for scanner, readoriginal := bufio.NewScanner(file), false; scanner.Scan(); {
		switch line := scanner.Text(); {
		case strings.HasPrefix(line, "#"):
		case len(line) == 0:
			readoriginal = false
		case readoriginal == false:
			replacers = append(replacers, Replacer{original: regexp.MustCompile(line)})
			readoriginal = true
		default:
			replacers[len(replacers)-1].replacements = append(replacers[len(replacers)-1].replacements, line)
		}
	}
	return replacers

} // ReadReplacersFromFile

// Eliza data structure
type Eliza struct {
	responses     []Replacer
	substitutions []Replacer
}

// FromFiles seeds 'Eliza' data structure from files
func FromFiles(responsePath string, substitutionPath string) Eliza {
	eliza := Eliza{}

	eliza.responses = ReadReplacersFromFile(responsePath)
	eliza.substitutions = ReadReplacersFromFile(substitutionPath)

	return eliza

} // ElizaFromFiles

// RespondTo returns a response based on user input
func (me *Eliza) RespondTo(input string) string {
	for _, response := range me.responses {
		if matches := response.original.FindStringSubmatch(input); matches != nil {
			output := response.replacements[rand.Intn(len(response.replacements))]
			boundaries := regexp.MustCompile(`[\s,.?!]+`)

			for m, match := range matches[1:] {
				tokens := boundaries.Split(match, -1)
				for t, token := range tokens {
					for _, substitution := range me.substitutions {
						if substitution.original.MatchString(token) {
							tokens[t] = substitution.replacements[rand.Intn(len(substitution.replacements))]
							break
						}
					} // for substitution
				} // for token

				output = strings.Replace(output, "$"+strconv.Itoa(m+1), strings.Join(tokens, " "), -1)
			}

			return output

		} // if matches
	} // for response

	return "I don't know what to say."

} // RespondTo
