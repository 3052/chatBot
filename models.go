package chatBot

import (
   "errors"
   "fmt"
)

func (m *model) String() string {
   var b []byte
   b = fmt.Appendln(b, "slug =", m.slug)
   b = fmt.Append(b, "url = ", m.url)
   if m.err != nil {
      b = fmt.Append(b, "\nerr = ", m.err)
   }
   return string(b)
}

type model struct {
   slug string
   url  string
   err  error
}

type models []*model

var (
   errLegacy  = errors.New("legacy")
   errPreview = errors.New("preview")
   errTrain   = errors.New("paid endpoints that may train on inputs")
)

var all_models = models{
   {
      slug: "ai21/jamba-large-1.7",
      url:  "https://studio.ai21.com",
   },
   {
      slug: "ai21/jamba-mini-1.7",
      url:  "https://studio.ai21.com",
   },
   //////////////////////////////////////////////////////////////////////////////
   {
      slug: "anthropic/claude-3.5-haiku-20241022",
      err:  errLegacy,
   },
   {
      slug: "anthropic/claude-3.5-sonnet",
      err:  errLegacy,
   },
   {
      slug: "anthropic/claude-3.5-sonnet-20240620",
      err:  errLegacy,
   },
   {
      slug: "anthropic/claude-3.7-sonnet",
      err:  errLegacy,
   },
   {
      slug: "anthropic/claude-opus-4",
      err:  errLegacy,
   },
   {
      slug: "anthropic/claude-opus-4.1",
      url:  "https://claude.ai",
   },
   {
      slug: "anthropic/claude-sonnet-4",
      url:  "https://claude.ai",
   },
   //////////////////////////////////////////////////////////////////////////////
   {
      slug: "bytedance/ui-tars-1.5-7b",
      url:  "https://openrouter.ai/chat?models=bytedance/ui-tars-1.5-7b",
   },
   //////////////////////////////////////////////////////////////////////////////
   {
      slug: "deepseek/deepseek-chat-v3.1",
      url:  "https://chat.deepseek.com",
   },
   {
      slug: "deepseek/deepseek-r1-0528",
      url:  "https://deepinfra.com/deepseek-ai/DeepSeek-R1-0528",
   },
   {
      slug: "deepseek/deepseek-r1",
      url:  "https://deepinfra.com/deepseek-ai/DeepSeek-R1",
   },
   {
      slug: "deepseek/deepseek-prover-v2",
      url:  "https://deepinfra.com/deepseek-ai/DeepSeek-Prover-V2-671B",
      err: errors.New(`due to low usage this model has been replaced by
      deepseek-ai/DeepSeek-V3-0324`),
   },
   {
      slug: "deepseek/deepseek-chat-v3-0324",
      url:  "https://deepinfra.com/deepseek-ai/DeepSeek-V3-0324",
   },
   {
      slug: "deepseek/deepseek-r1-distill-llama-70b",
      url:  "https://deepinfra.com/deepseek-ai/DeepSeek-R1-Distill-Llama-70B",
   },
   {
      slug: "deepseek/deepseek-r1-distill-qwen-32b",
      url:  "https://deepinfra.com/deepseek-ai/DeepSeek-R1-Distill-Qwen-32B",
   },
   //////////////////////////////////////////////////////////////////////////////
   {
      slug: "arcee-ai/spotlight",
      url:  "https://api.together.ai/playground/arcee_ai/arcee-spotlight",
   },
   {
      slug: "arcee-ai/maestro-reasoning",
      url:  "https://api.together.ai/playground/arcee-ai/maestro-reasoning",
   },
   {
      slug: "arcee-ai/virtuoso-large",
      url:  "https://api.together.ai/playground/arcee-ai/virtuoso-large",
   },
   //////////////////////////////////////////////////////////////////////////////
   {
      slug: "google/gemini-2.0-flash-001",
      err:  errLegacy,
   },
   {
      slug: "google/gemini-2.0-flash-lite-001",
      err:  errLegacy,
   },
   {
      slug: "google/gemini-pro-1.5",
      err:  errLegacy,
   },
   {
      slug: "google/gemini-flash-1.5",
      err:  errLegacy,
   },
   {
      slug: "google/gemini-flash-1.5-8b",
      err:  errLegacy,
   },
   {
      slug: "google/gemini-2.5-pro-preview",
      err:  errPreview,
   },
   {
      slug: "google/gemini-2.5-flash-lite-preview-06-17",
      err:  errPreview,
   },
   {
      slug: "google/gemini-2.5-pro-preview-05-06",
      err:  errPreview,
   },
   {
      slug: "google/gemini-2.5-pro",
      url:  "https://gemini.google.com",
   },
   {
      slug: "google/gemini-2.5-flash",
      url:  "https://gemini.google.com",
   },
   {
      slug: "google/gemini-2.5-flash-lite",
      url:  "https://aistudio.google.com/prompts/new_chat?model=gemini-2.5-flash-lite",
   },
   //////////////////////////////////////////////////////////////////////////////
   {
      slug: "inception/mercury",
      url:  "https://chat.inceptionlabs.ai",
   },
   {
      slug: "inception/mercury-coder",
      url:  "https://chat.inceptionlabs.ai",
   },
   //////////////////////////////////////////////////////////////////////////////
   {
      slug: "meta-llama/llama-4-scout",
      url:  "https://deepinfra.com/meta-llama/Llama-4-Scout-17B-16E-Instruct",
   },
   {
      slug: "meta-llama/llama-4-maverick",
      url:  "https://deepinfra.com/meta-llama/Llama-4-Maverick-17B-128E-Instruct-FP8",
   },
   {
      slug: "meta-llama/llama-guard-4-12b",
      url:  "https://deepinfra.com/meta-llama/Llama-Guard-4-12B",
   },
   //////////////////////////////////////////////////////////////////////////////
   {
      slug: "microsoft/mai-ds-r1",
      err:  errTrain,
   },
   //////////////////////////////////////////////////////////////////////////////
   {
      slug: "minimax/minimax-m1",
      url:  "https://chat.minimax.io",
   },
   //////////////////////////////////////////////////////////////////////////////
   {
      slug: "mistralai/mistral-medium-3.1",
      err:  errLegacy,
   },
   {
      slug: "mistralai/mistral-medium-3",
      err:  errLegacy,
   },
   {
      slug: "mistralai/devstral-small",
      err:  errLegacy,
   },
   {
      slug: "mistralai/mistral-small-3.1-24b-instruct",
      err:  errLegacy,
   },
   {
      slug: "mistralai/codestral-2508",
      url:  "https://console.mistral.ai/build/playground",
   },
   {
      slug: "mistralai/devstral-medium",
      url:  "https://console.mistral.ai/build/playground",
   },
   {
      slug: "mistralai/devstral-small-2505",
      url:  "https://console.mistral.ai/build/playground",
   },
   {
      slug: "mistralai/mistral-small-3.2-24b-instruct",
      url:  "https://deepinfra.com/mistralai/Mistral-Small-3.2-24B-Instruct-2506",
   },
   //////////////////////////////////////////////////////////////////////////////
   {
      slug: "moonshotai/kimi-dev-72b",
      url:  "https://cloud.siliconflow.com/playground/chat",
   },
   {
      slug: "moonshotai/kimi-vl-a3b-thinking",
      err:  errTrain,
   },
   //////////////////////////////////////////////////////////////////////////////
   {
      slug: "nousresearch/hermes-4-70b",
      url:  "https://chat.nousresearch.com",
   },
   {
      slug: "nousresearch/hermes-4-405b",
      url:  "https://chat.nousresearch.com",
   },
   //////////////////////////////////////////////////////////////////////////////
   {
      slug: "nvidia/llama-3.3-nemotron-super-49b-v1",
      url:  "https://build.nvidia.com/nvidia/llama-3_3-nemotron-super-49b-v1",
   },
   {
      slug: "nvidia/llama-3.1-nemotron-ultra-253b-v1",
      url:  "https://build.nvidia.com/nvidia/llama-3_1-nemotron-ultra-253b-v1",
   },
   //////////////////////////////////////////////////////////////////////////////
   {
      slug: "openai/gpt-4.1",
      err:  errLegacy,
   },
   {
      slug: "openai/gpt-4.1-mini",
      err:  errLegacy,
   },
   {
      slug: "openai/gpt-4.1-nano",
      err:  errLegacy,
   },
   {
      slug: "openai/gpt-4o",
      err:  errLegacy,
   },
   {
      slug: "openai/gpt-4o-2024-05-13",
      err:  errLegacy,
   },
   {
      slug: "openai/gpt-4o-2024-08-06",
      err:  errLegacy,
   },
   {
      slug: "openai/gpt-4o-2024-11-20",
      err:  errLegacy,
   },
   {
      slug: "openai/gpt-4o-audio-preview",
      err:  errLegacy,
   },
   {
      slug: "openai/gpt-4o-mini",
      err:  errLegacy,
   },
   {
      slug: "openai/gpt-4o-mini-2024-07-18",
      err:  errLegacy,
   },
   {
      slug: "openai/gpt-4o-mini-search-preview",
      err:  errLegacy,
   },
   {
      slug: "openai/gpt-4o-search-preview",
      err:  errLegacy,
   },
   {
      slug: "openai/o3",
      err:  errLegacy,
   },
   {
      slug: "openai/o4-mini",
      err:  errLegacy,
   },
   {
      slug: "openai/o4-mini-high",
      err:  errLegacy,
   },
   {
      slug: "openai/codex-mini",
      err:  errLegacy,
   },
   {
      slug: "openai/gpt-oss-120b",
      url:  "https://gpt-oss.com",
   },
   {
      slug: "openai/gpt-oss-20b",
      url:  "https://gpt-oss.com",
   },
   {
      slug: "openai/gpt-5-chat",
      url:  "https://chatgpt.com?model=gpt-5",
   },
   {
      slug: "openai/gpt-5-mini",
      url:  "https://platform.openai.com/docs/models/gpt-5-mini",
   },
   {
      slug: "openai/gpt-5-nano",
      url:  "https://platform.openai.com/docs/models/gpt-5-nano",
   },
   {
      slug: "openai/gpt-5",
      url:  "https://platform.openai.com/docs/models/gpt-5",
   },
   //////////////////////////////////////////////////////////////////////////////
   {
      slug: "perplexity/r1-1776",
      err:  errors.New(`provider returned error: {"error":{"code":400}}`),
   },
   {
      slug: "perplexity/sonar-deep-research",
      err:  errLegacy,
   },
   {
      slug: "perplexity/sonar-pro",
      url:  "https://playground.perplexity.ai",
   },
   {
      slug: "perplexity/sonar-reasoning-pro",
      url:  "https://playground.perplexity.ai",
   },
   //////////////////////////////////////////////////////////////////////////////
   {
      slug: "qwen/qwen3-8b",
      url:  "https://openrouter.ai/chat?models=qwen/qwen3-8b",
   },
   {
      slug: "qwen/qwen3-30b-a3b-instruct-2507",
      url:  "https://chat.qwen.ai",
   },
   {
      slug: "qwen/qwen3-30b-a3b-thinking-2507",
      url:  "https://chat.qwen.ai",
   },
   {
      slug: "qwen/qwen3-235b-a22b-2507",
      url:  "https://chat.qwen.ai",
   },
   {
      slug: "qwen/qwen3-235b-a22b-thinking-2507",
      url:  "https://chat.qwen.ai",
   },
   {
      slug: "qwen/qwen3-coder",
      url:  "https://chat.qwen.ai",
   },
   {
      slug: "qwen/qwen3-coder-30b-a3b-instruct",
      url:  "https://openrouter.ai/chat?models=qwen/qwen3-coder-30b-a3b-instruct",
   },
   {
      slug: "qwen/qwq-32b",
      url:  "https://openrouter.ai/chat?models=qwen/qwen3-32b",
   },
   //////////////////////////////////////////////////////////////////////////////
   {
      slug: "switchpoint/router",
      url:  "https://switchpoint.dev/chat",
   },
   //////////////////////////////////////////////////////////////////////////////
   {
      slug: "tngtech/deepseek-r1t-chimera",
      err:  errTrain,
   },
   //////////////////////////////////////////////////////////////////////////////
   {
      slug: "x-ai/grok-3",
      url:  "https://grok.com",
   },
   {
      slug: "x-ai/grok-3-beta",
      err:  errPreview,
   },
   {
      slug: "x-ai/grok-3-mini",
      url:  "https://openrouter.ai/chat?models=x-ai/grok-3-mini",
   },
   {
      slug: "x-ai/grok-3-mini-beta",
      err:  errPreview,
   },
   {
      slug: "x-ai/grok-4",
      url:  "https://grok.com",
   },
   {
      slug: "x-ai/grok-code-fast-1",
      url:  "https://openrouter.ai/chat?models=x-ai/grok-code-fast-1",
   },
   //////////////////////////////////////////////////////////////////////////////
   {
      slug: "z-ai/glm-4.5",
      url:  "https://chat.z.ai",
   },
   {
      slug: "z-ai/glm-4.5-air",
      url:  "https://chat.z.ai",
   },
   {
      slug: "z-ai/glm-4-32b",
      url:  "https://chat.z.ai",
   },
}
