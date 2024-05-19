
package genaitesty

import "github.com/google/generative-ai-go/genai"
import "google.golang.org/api/option"

ctx := context.Background()
// Access your API key as an environment variable (see "Set up your API key" above)
client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("API_KEY")))
if err != nil {
  log.Fatal(err)
}
defer client.Close()

//model := client.GenerativeModel("MODEL_NAME")
model := client.GenerativeModel("gemini-pro")

