package nlp

import (
	"reflect"
	"strings"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/stretchr/testify/require"
)

type tokenizeCases struct {
	Text   string
	Tokens []string
}

/*
	 var tokenizeCasesTest = []tokenizeCases{
		{"What's on second?", []string{"what", "s", "on", "second"}},
		{"", nil},
	}
*/
func loadTokenizeCases(t *testing.T) []tokenizeCases {
	/*
		data, err := os.ReadFile("tokenize_cases.toml")
		require.NoError(t, err, "Read File")
	*/
	var testCases struct {
		Cases []tokenizeCases
	}

	// err = toml.Unmarshal(data, &testCases)
	_, err := toml.DecodeFile("tokenize_cases.toml", &testCases)
	require.NoError(t, err, "Unmarshal TOML")
	return testCases.Cases
}
func TestTokenizeTable(t *testing.T) {
	for _, tc := range loadTokenizeCases(t) {
		t.Run(tc.Text, func(t *testing.T) {
			tokens := Tokenize(tc.Text)
			require.Equal(t, tc.Tokens, tokens)
		})
	}
}

/*
	func TestTokenizeTable(t *testing.T) {
		for _, tc := range tokenizeCasesTest {
			t.Run(tc.Text, func(t *testing.T) {
				tokens := Tokenize(tc.Text)
				require.Equal(t, tc.Tokens, tokens)
			})
		}
	}
*/
func TestTokenize(t *testing.T) {
	text := "What's on second?"
	expected := []string{"what", "on", "second"}
	tokens := Tokenize(text)
	require.Equal(t, expected, tokens)
	// Before testify
	if !reflect.DeepEqual(expected, tokens) {
		t.Fatalf("expected %#v, got %#v", expected, tokens)
	}
}

func FuzzTokenize(f *testing.F) {
	f.Fuzz(func(t *testing.T, text string) {
		tokens := Tokenize(text)
		lText := strings.ToLower(text)
		for _, tok := range tokens {
			if !strings.Contains(lText, tok) {
				t.Fatal(tok)
			}
		}
	})
}
