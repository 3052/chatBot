package chatBot

import (
   "encoding/json"
   "fmt"
   "io"
   "net/http"
   "time"
)

var canonical = map[string]bool{
   "anthropic/claude-opus-4.1": true,
   "anthropic/claude-sonnet-4": true,
   "anthropic/claude-opus-4": false,
   "anthropic/claude-3.5-sonnet": false,
   "anthropic/claude-3.7-sonnet": false,
   "anthropic/claude-3.5-sonnet-20240620": false,
   "anthropic/claude-3.5-haiku-20241022": false,
   "google/gemini-2.5-flash-lite-preview-06-17": false,
   "google/gemini-2.0-flash-001": false,
   "google/gemini-2.0-flash-lite-001": false,
   "google/gemini-flash-1.5": false,
   "google/gemini-flash-1.5-8b": false,
   "google/gemini-2.5-pro-exp-03-25": false,
   "google/gemini-2.5-pro-preview-05-06": false,
   "google/gemini-pro-1.5": false,
   "google/gemini-2.5-pro-preview": false,
   "google/gemini-2.5-flash": true,
   "google/gemini-2.5-pro": true,
   // aistudio.google.com?model=gemini-2.5-flash-lite
   "google/gemini-2.5-flash-lite": true,
   "openai/gpt-4o-2024-05-13": false,
   "openai/gpt-4o-2024-08-06": false,
   "openai/gpt-4o-2024-11-20": false,
   "openai/gpt-4o-audio-preview": false,
   "openai/gpt-4o-mini-2024-07-18": false,
   "openai/gpt-4o-mini-search-preview": false,
   "openai/gpt-4o-search-preview": false,
   
   "openai/gpt-4o": false,
   "openai/gpt-4o-mini": false,
   "openai/codex-mini": false,
   "openai/gpt-4.1": false,
   "openai/gpt-4.1-mini": false,
   "openai/gpt-4.1-nano": false,
   "openai/gpt-5": true,
   "openai/gpt-5-chat": false,
   "openai/gpt-5-mini": false,
   "openai/gpt-5-nano": false,
   "openai/gpt-oss-120b": false,
   "openai/gpt-oss-20b": false,
   "openai/o3": false,
   "openai/o3-pro": false,
   "openai/o4-mini": false,
   "openai/o4-mini-high": false,
   
   "ai21/jamba-large-1.7": false,
   "ai21/jamba-mini-1.7": false,
   "arcee-ai/maestro-reasoning": false,
   "arcee-ai/spotlight": false,
   "arcee-ai/virtuoso-large": false,
   "bytedance/ui-tars-1.5-7b": false,
   "deepseek/deepseek-chat-v3-0324": false,
   "deepseek/deepseek-chat-v3.1": false,
   "deepseek/deepseek-prover-v2": false,
   "deepseek/deepseek-r1": false,
   "deepseek/deepseek-r1-0528": false,
   "deepseek/deepseek-r1-distill-llama-70b": false,
   "deepseek/deepseek-r1-distill-qwen-1.5b": false,
   "deepseek/deepseek-r1-distill-qwen-32b": false,
   "deepseek/deepseek-v3-base": false,
   "deepseek/deepseek-v3.1-base": false,
   "inception/mercury": false,
   "inception/mercury-coder": false,
   "meta-llama/llama-4-maverick": false,
   "meta-llama/llama-4-scout": false,
   "meta-llama/llama-guard-4-12b": false,
   "microsoft/mai-ds-r1": false,
   "minimax/minimax-m1": false,
   "mistralai/codestral-2508": false,
   "mistralai/devstral-medium": false,
   "mistralai/devstral-small": false,
   "mistralai/devstral-small-2505": false,
   "mistralai/mistral-medium-3": false,
   "mistralai/mistral-medium-3.1": false,
   "mistralai/mistral-small-3.1-24b-instruct": false,
   "mistralai/mistral-small-3.2-24b-instruct": false,
   "moonshotai/kimi-vl-a3b-thinking": false,
   "nvidia/llama-3.1-nemotron-ultra-253b-v1": false,
   "nvidia/llama-3.3-nemotron-super-49b-v1": false,
   "perplexity/r1-1776": false,
   "perplexity/sonar-deep-research": false,
   "perplexity/sonar-pro": false,
   "perplexity/sonar-reasoning-pro": false,
   "qwen/qwen3-235b-a22b-2507": false,
   "qwen/qwen3-235b-a22b-thinking-2507": false,
   "qwen/qwen3-30b-a3b-instruct-2507": false,
   "qwen/qwen3-8b": false,
   "qwen/qwen3-coder": true,
   "qwen/qwq-32b": false,
   "switchpoint/router": false,
   "tngtech/deepseek-r1t-chimera": false,
   "x-ai/grok-3": false,
   "x-ai/grok-3-beta": false,
   "x-ai/grok-3-mini": false,
   "x-ai/grok-3-mini-beta": false,
   "x-ai/grok-4": false,
   "z-ai/glm-4-32b": false,
   "z-ai/glm-4.5": false,
   "z-ai/glm-4.5-air": false,
}

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
