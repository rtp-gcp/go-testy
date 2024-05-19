
package main

import "github.com/google/generative-ai-go/genai"
import "google.golang.org/api/option"

import (
    "fmt"
    "log"
    "context"
    "os"
//    "encoding/json"   // json.MarshalIndent()
)

func main() {
    fmt.Println("hello from go")
    ctx := context.Background()
    // Access your API key as an environment variable (see "Set up your API key" above)
    client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GOOGLE_API_KEY")))
    if err != nil {
      log.Fatal(err)
    }
    defer client.Close()

    // For text-only input, use the gemini-pro model
    model := client.GenerativeModel("gemini-pro")
    resp, err := model.GenerateContent(ctx, genai.Text("Write a story about a magic backpack."))
    if err != nil {
      log.Fatal(err)
    }

    // Sometimes, this returns a result:
    // 2024/03/24 09:49:24 blocked: candidate: FinishReasonSafety

    // https://pkg.go.dev/github.com/google/generative-ai-go/genai#GenerateContentResponse

    //it := genai.GenerateContentResponseIterator(resp)
    //it := genai.GenerateContentResponseIterator(resp)
    //it := client.GenerateContentResponseIterator(resp)
    //it := genai.client.GenerateContentResponseIterator(resp)

    //for it.HasNext() {
    //    res := it.Next()
    //    fmt.Println(res)
    //}

    // https://eli.thegreenplace.net/2023/using-gemini-models-from-go/
    //bs, _ := json.MarshalIndent(resp, "", "    ")
    //fmt.Println(string(bs))

    // Grab the specific text result
    //fmt.Println(resp.Candidates)
    fmt.Println(resp.Candidates[0].Content.Parts)

}






