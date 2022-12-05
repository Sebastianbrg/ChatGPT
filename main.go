package main
import (
	"fmt"
	"strings"
	"github.com/openai/gpt-3-go/gpt3"
)
func main() {
	// lets start to play. I have no idea where this will end. 20 minutes. Go!
	input := "Hello world!"
	output := convertToBertStyle(input)
	fmt.Println(output) // should print "[CLS] Hello[SPACE]world!"
	// Define the input text.
	text := "This is an example of tokenized text."

	// Tokenize the text.
tokens := tokenizeText(text)

// Print the tokens.
fmt.Println(tokens) // ["This", "is", "an", "example", "of", "tokenized", "text."]
	key:= os.GetEnv("GPT_KEY")
// Create a new GPT-3 client
client := gpt3.NewClient(key)

// Generate text using the pre-trained GPT-3 model
response, err := client.Generate(gpt3.GenerateOpts{
	Prompt: "Hello world!",
	Length: 100,
})
if err != nil {
	fmt.Println(err)
} else {
	fmt.Println(response.Data.Choices[0].Text)
}


}


func convertToBertStyle(input string) string {
	// Add [CLS] token to the beginning of the input string
	output := "[CLS] " + input

	// Replace all whitespace with the [SPACE] token
	output = strings.Replace(output, " ", "[SPACE]", -1)

	return output
}


// Define the tokenizeText function.
func tokenizeText(text string) []string {
	// Split the text into a slice of tokens.
	tokens := strings.Fields(text)
  
	// Return the slice of tokens.
	return tokens
  }

  
