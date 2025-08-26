package chatBot

import (
   "encoding/json"
   "fmt"
   "io"
   "net/http"
   "time"
)

func delete_model(m *model) bool {
   if m.ContextLength < 128000 {
      return true
   }
   if m.Endpoint == nil {
      return true
   }
   if m.Endpoint.ModelVariantSlug != m.Slug {
      return true
   }
   const day = 24 * time.Hour
   const month = 30 * day
   return time.Since(m.UpdatedAt) >= 5*month
}

type model struct {
   Author string
   ContextLength int `json:"context_length"`
   Endpoint *struct { // DELETE
      ModelVariantSlug string `json:"model_variant_slug"`
   }
   ShortName string `json:"short_name"`
   Slug string
   UpdatedAt time.Time `json:"updated_at"`
}

func (m *model) String() string {
   var b []byte
   b = fmt.Appendln(b, "author =", m.Author)
   b = fmt.Appendln(b, "context =", m.ContextLength)
   if m.Endpoint != nil {
      b = fmt.Appendln(b, "endpoint model =", m.Endpoint.ModelVariantSlug)
   }
   b = fmt.Appendln(b, "short =", m.ShortName)
   b = fmt.Appendln(b, "slug =", m.Slug)
   b = fmt.Append(b, "updated = ", m.UpdatedAt)
   return string(b)
}

func get_models() (byte_slice[models], error) {
   req, _ := http.NewRequest("", "https://openrouter.ai", nil)
   req.URL.Path = "/api/frontend/models/find"
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   return io.ReadAll(resp.Body)
}

type byte_slice[T any] []byte

type models []*model

func (m *models) unmarshal(data byte_slice[models]) error {
   var value struct {
      Data struct {
         Models []*model
      }
   }
   err := json.Unmarshal(data, &value)
   if err != nil {
      return err
   }
   *m = value.Data.Models
   return nil
}

var slugs = []string{
   "ai21/jamba-large-1.7",
   "ai21/jamba-mini-1.7",
   "anthropic/claude-3.5-haiku-20241022",
   "anthropic/claude-3.5-sonnet",
   "anthropic/claude-3.5-sonnet-20240620",
   "anthropic/claude-3.7-sonnet",
   "anthropic/claude-opus-4",
   "anthropic/claude-opus-4.1",
   "anthropic/claude-sonnet-4",
   "arcee-ai/maestro-reasoning",
   "arcee-ai/spotlight",
   "arcee-ai/virtuoso-large",
   "bytedance/ui-tars-1.5-7b",
   "deepseek/deepseek-chat-v3-0324",
   "deepseek/deepseek-chat-v3.1",
   "deepseek/deepseek-prover-v2",
   "deepseek/deepseek-r1",
   "deepseek/deepseek-r1-0528",
   "deepseek/deepseek-r1-distill-llama-70b",
   "deepseek/deepseek-r1-distill-qwen-1.5b",
   "deepseek/deepseek-r1-distill-qwen-32b",
   "deepseek/deepseek-v3-base",
   "deepseek/deepseek-v3.1-base",
   "google/gemini-2.0-flash-001",
   "google/gemini-2.0-flash-lite-001",
   "google/gemini-2.5-flash",
   "google/gemini-2.5-flash-lite",
   "google/gemini-2.5-flash-lite-preview-06-17",
   "google/gemini-2.5-pro",
   "google/gemini-2.5-pro-exp-03-25",
   "google/gemini-2.5-pro-preview",
   "google/gemini-2.5-pro-preview-05-06",
   "google/gemini-flash-1.5",
   "google/gemini-flash-1.5-8b",
   "google/gemini-pro-1.5",
   "inception/mercury",
   "inception/mercury-coder",
   "meta-llama/llama-4-maverick",
   "meta-llama/llama-4-scout",
   "meta-llama/llama-guard-4-12b",
   "microsoft/mai-ds-r1",
   "minimax/minimax-m1",
   "mistralai/codestral-2508",
   "mistralai/devstral-medium",
   "mistralai/devstral-small",
   "mistralai/devstral-small-2505",
   "mistralai/mistral-medium-3",
   "mistralai/mistral-medium-3.1",
   "mistralai/mistral-small-3.1-24b-instruct",
   "mistralai/mistral-small-3.2-24b-instruct",
   "moonshotai/kimi-vl-a3b-thinking",
   "nvidia/llama-3.1-nemotron-ultra-253b-v1",
   "nvidia/llama-3.3-nemotron-super-49b-v1",
   "openai/codex-mini",
   "openai/gpt-4.1",
   "openai/gpt-4.1-mini",
   "openai/gpt-4.1-nano",
   "openai/gpt-4o",
   "openai/gpt-4o-2024-05-13",
   "openai/gpt-4o-2024-08-06",
   "openai/gpt-4o-2024-11-20",
   "openai/gpt-4o-audio-preview",
   "openai/gpt-4o-mini",
   "openai/gpt-4o-mini-2024-07-18",
   "openai/gpt-4o-mini-search-preview",
   "openai/gpt-4o-search-preview",
   "openai/gpt-5",
   "openai/gpt-5-chat",
   "openai/gpt-5-mini",
   "openai/gpt-5-nano",
   "openai/gpt-oss-120b",
   "openai/gpt-oss-20b",
   "openai/o3",
   "openai/o3-pro",
   "openai/o4-mini",
   "openai/o4-mini-high",
   "perplexity/r1-1776",
   "perplexity/sonar-deep-research",
   "perplexity/sonar-pro",
   "perplexity/sonar-reasoning-pro",
   "qwen/qwen3-235b-a22b-2507",
   "qwen/qwen3-235b-a22b-thinking-2507",
   "qwen/qwen3-30b-a3b-instruct-2507",
   "qwen/qwen3-8b",
   "qwen/qwen3-coder",
   "qwen/qwq-32b",
   "switchpoint/router",
   "tngtech/deepseek-r1t-chimera",
   "x-ai/grok-3",
   "x-ai/grok-3-beta",
   "x-ai/grok-3-mini",
   "x-ai/grok-3-mini-beta",
   "x-ai/grok-4",
   "z-ai/glm-4-32b",
   "z-ai/glm-4.5",
   "z-ai/glm-4.5-air",
}
