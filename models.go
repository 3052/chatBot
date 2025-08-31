package chatBot

import "errors"

type models []*model

type model struct {
   err  error
   slug string
   url  string
}

var (
   errLegacy = errors.New("legacy")
   errPreview = errors.New("preview")
   errUnverified = errors.New("unverified")
)

var all_models = models{
   {
      slug: "ai21/jamba-large-1.7",
      url:  "studio.ai21.com",
   },
   {
      slug: "ai21/jamba-mini-1.7",
      url:  "studio.ai21.com",
   },
   //////////////////////////////////////////////////////////////////////////////
   {
      slug: "anthropic/claude-3.5-haiku-20241022",
      err: errLegacy,
   },
   {
      slug: "anthropic/claude-3.5-sonnet",
      err: errLegacy,
   },
   {
      slug: "anthropic/claude-3.5-sonnet-20240620",
      err: errLegacy,
   },
   {
      slug: "anthropic/claude-3.7-sonnet",
      err: errLegacy,
   },
   {
      slug: "anthropic/claude-opus-4",
      err: errLegacy,
   },
   {
      slug: "anthropic/claude-opus-4.1",
      url:  "claude.ai",
   },
   {
      slug: "anthropic/claude-sonnet-4",
      url:  "claude.ai",
   },
   //////////////////////////////////////////////////////////////////////////////
   {
      slug: "bytedance/ui-tars-1.5-7b",
      url: "openrouter.ai/chat?models=bytedance/ui-tars-1.5-7b",
   },
   //////////////////////////////////////////////////////////////////////////////
   {
      slug: "deepseek/deepseek-chat-v3.1",
      url: "chat.deepseek.com",
   },
   {
      slug: "deepseek/deepseek-r1-0528",
      url: "deepinfra.com/deepseek-ai/DeepSeek-R1-0528",
   },
   {
      slug: "deepseek/deepseek-r1",
      url: "deepinfra.com/deepseek-ai/DeepSeek-R1",
   },
   {
      slug: "deepseek/deepseek-prover-v2",
      url: "deepinfra.com/deepseek-ai/DeepSeek-Prover-V2-671B",
      err: errors.New(`due to low usage this model has been replaced by
      deepseek-ai/DeepSeek-V3-0324`),
   },
   {
      slug: "deepseek/deepseek-chat-v3-0324",
      url: "deepinfra.com/deepseek-ai/DeepSeek-V3-0324",
   },
   {
      slug: "deepseek/deepseek-r1-distill-llama-70b",
      url: "deepinfra.com/deepseek-ai/DeepSeek-R1-Distill-Llama-70B",
   },
   {
      slug: "deepseek/deepseek-r1-distill-qwen-32b",
      url: "deepinfra.com/deepseek-ai/DeepSeek-R1-Distill-Qwen-32B",
   },
   //////////////////////////////////////////////////////////////////////////////
   {
      slug: "arcee-ai/spotlight",
      url:  "api.together.ai/playground/arcee_ai/arcee-spotlight",
   },
   {
      slug: "arcee-ai/maestro-reasoning",
      url:  "api.together.ai/playground/arcee-ai/maestro-reasoning",
   },
   {
      slug: "arcee-ai/virtuoso-large",
      url:  "api.together.ai/playground/arcee-ai/virtuoso-large",
   },
   //////////////////////////////////////////////////////////////////////////////
   {
      slug: "google/gemini-2.0-flash-001",
      err: errLegacy,
   },
   {
      slug: "google/gemini-2.0-flash-lite-001",
      err: errLegacy,
   },
   {
      slug: "google/gemini-pro-1.5",
      err: errLegacy,
   },
   {
      slug: "google/gemini-flash-1.5",
      err: errLegacy,
   },
   {
      slug: "google/gemini-flash-1.5-8b",
      err: errLegacy,
   },
   {
      slug: "google/gemini-2.5-pro-preview",
      err: errPreview,
   },
   {
      slug: "google/gemini-2.5-flash-lite-preview-06-17",
      err: errPreview,
   },
   {
      slug: "google/gemini-2.5-pro-preview-05-06",
      err: errPreview,
   },
   {
      slug: "google/gemini-2.5-pro",
      url: "gemini.google.com",
   },
   {
      slug: "google/gemini-2.5-flash",
      url: "gemini.google.com",
   },
   {
      slug: "google/gemini-2.5-flash-lite",
      url: "aistudio.google.com/prompts/new_chat?model=gemini-2.5-flash-lite",
   },
   //////////////////////////////////////////////////////////////////////////////
   {
      slug: "inception/mercury",
      url: "chat.inceptionlabs.ai",
   },
   {
      slug: "inception/mercury-coder",
      url: "chat.inceptionlabs.ai",
   },
   //////////////////////////////////////////////////////////////////////////////
   {
      slug: "meta-llama/llama-4-scout",
      url: "deepinfra.com/meta-llama/Llama-4-Scout-17B-16E-Instruct",
   },
   {
      slug: "meta-llama/llama-4-maverick",
      url: "deepinfra.com/meta-llama/Llama-4-Maverick-17B-128E-Instruct-FP8",
   },
   {
      slug: "meta-llama/llama-guard-4-12b",
      url: "deepinfra.com/meta-llama/Llama-Guard-4-12B",
   },
   //////////////////////////////////////////////////////////////////////////////
   //////////////////////////////////////////////////////////////////////////////
   {
      slug: "microsoft/mai-ds-r1",
      err:  errUnverified,
   },
   //////////////////////////////////////////////////////////////////////////////
   {
      slug: "minimax/minimax-m1",
      err:  errUnverified,
   },
   //////////////////////////////////////////////////////////////////////////////
   {
      slug: "mistralai/codestral-2508",
      err:  errUnverified,
   },
   {
      slug: "mistralai/devstral-medium",
      err:  errUnverified,
   },
   {
      slug: "mistralai/devstral-small",
      err:  errUnverified,
   },
   {
      slug: "mistralai/devstral-small-2505",
      err:  errUnverified,
   },
   {
      slug: "mistralai/mistral-medium-3",
      err:  errUnverified,
   },
   {
      slug: "mistralai/mistral-medium-3.1",
      err:  errUnverified,
   },
   {
      slug: "mistralai/mistral-small-3.1-24b-instruct",
      err:  errUnverified,
   },
   {
      slug: "mistralai/mistral-small-3.2-24b-instruct",
      err:  errUnverified,
   },
   //////////////////////////////////////////////////////////////////////////////
   {
      slug: "moonshotai/kimi-dev-72b",
      err:  errUnverified,
   },
   {
      slug: "moonshotai/kimi-vl-a3b-thinking",
      err:  errUnverified,
   },
   //////////////////////////////////////////////////////////////////////////////
   {
      slug: "nousresearch/hermes-4-405b",
      err:  errUnverified,
   },
   {
      slug: "nousresearch/hermes-4-70b",
      err:  errUnverified,
   },
   //////////////////////////////////////////////////////////////////////////////
   {
      slug: "nvidia/llama-3.1-nemotron-ultra-253b-v1",
      err:  errUnverified,
   },
   {
      slug: "nvidia/llama-3.3-nemotron-super-49b-v1",
      err:  errUnverified,
   },
   //////////////////////////////////////////////////////////////////////////////
   {
      slug: "openai/codex-mini",
      err:  errUnverified,
   },
   {
      slug: "openai/gpt-4.1",
      err:  errUnverified,
   },
   {
      slug: "openai/gpt-4.1-mini",
      err:  errUnverified,
   },
   {
      slug: "openai/gpt-4.1-nano",
      err:  errUnverified,
   },
   {
      slug: "openai/gpt-4o",
      err:  errUnverified,
   },
   {
      slug: "openai/gpt-4o-2024-05-13",
      err:  errUnverified,
   },
   {
      slug: "openai/gpt-4o-2024-08-06",
      err:  errUnverified,
   },
   {
      slug: "openai/gpt-4o-2024-11-20",
      err:  errUnverified,
   },
   {
      slug: "openai/gpt-4o-audio-preview",
      err:  errUnverified,
   },
   {
      slug: "openai/gpt-4o-mini",
      err:  errUnverified,
   },
   {
      slug: "openai/gpt-4o-mini-2024-07-18",
      err:  errUnverified,
   },
   {
      slug: "openai/gpt-4o-mini-search-preview",
      err:  errUnverified,
   },
   {
      slug: "openai/gpt-4o-search-preview",
      err:  errUnverified,
   },
   {
      slug: "openai/gpt-5",
      url:  "platform.openai.com/docs/models/gpt-5",
   },
   {
      slug: "openai/gpt-5-chat",
      url:  "chatgpt.com?model=gpt-5",
   },
   {
      slug: "openai/gpt-5-mini",
      url:  "platform.openai.com/docs/models/gpt-5-mini",
   },
   {
      slug: "openai/gpt-5-nano",
      url:  "platform.openai.com/docs/models/gpt-5-nano",
   },
   {
      slug: "openai/gpt-oss-120b",
      url:  "gpt-oss.com",
   },
   {
      slug: "openai/gpt-oss-20b",
      url:  "gpt-oss.com",
   },
   {
      slug: "openai/o3",
      err:  errUnverified,
   },
   {
      slug: "openai/o4-mini",
      err:  errUnverified,
   },
   {
      slug: "openai/o4-mini-high",
      err:  errUnverified,
   },
   //////////////////////////////////////////////////////////////////////////////
   {
      slug: "perplexity/r1-1776",
      err:  errUnverified,
   },
   {
      slug: "perplexity/sonar-deep-research",
      err:  errUnverified,
   },
   {
      slug: "perplexity/sonar-pro",
      err:  errUnverified,
   },
   {
      slug: "perplexity/sonar-reasoning-pro",
      err:  errUnverified,
   },
   //////////////////////////////////////////////////////////////////////////////
   {
      slug: "qwen/qwen3-235b-a22b-2507",
      err:  errUnverified,
   },
   {
      slug: "qwen/qwen3-235b-a22b-thinking-2507",
      err:  errUnverified,
   },
   {
      slug: "qwen/qwen3-30b-a3b-instruct-2507",
      err:  errUnverified,
   },
   {
      slug: "qwen/qwen3-30b-a3b-thinking-2507",
      err:  errUnverified,
   },
   {
      slug: "qwen/qwen3-8b",
      err:  errUnverified,
   },
   {
      slug: "qwen/qwen3-coder",
   },
   {
      slug: "qwen/qwen3-coder-30b-a3b-instruct",
      err:  errUnverified,
   },
   {
      slug: "qwen/qwq-32b",
      err:  errUnverified,
   },
   //////////////////////////////////////////////////////////////////////////////
   {
      slug: "switchpoint/router",
      err:  errUnverified,
   },
   //////////////////////////////////////////////////////////////////////////////
   {
      slug: "tngtech/deepseek-r1t-chimera",
      err:  errUnverified,
   },
   //////////////////////////////////////////////////////////////////////////////
   {
      slug: "x-ai/grok-3",
      err:  errUnverified,
   },
   {
      slug: "x-ai/grok-3-beta",
      err:  errUnverified,
   },
   {
      slug: "x-ai/grok-3-mini",
      err:  errUnverified,
   },
   {
      slug: "x-ai/grok-3-mini-beta",
      err:  errUnverified,
   },
   {
      slug: "x-ai/grok-4",
      err:  errUnverified,
   },
   {
      slug: "x-ai/grok-code-fast-1",
      err:  errUnverified,
   },
   //////////////////////////////////////////////////////////////////////////////
   {
      slug: "z-ai/glm-4-32b",
      err:  errUnverified,
   },
   {
      slug: "z-ai/glm-4.5",
      err:  errUnverified,
   },
   {
      slug: "z-ai/glm-4.5-air",
      err:  errUnverified,
   },
}
