package chatBot

type model struct {
   slug string
   url  string
   info string
   ok   bool
}

type models []*model

const (
   legacyModel  = "legacy model"
   mayTrain   = "paid endpoints that may train on inputs"
   previewModel = "preview model"
)

var all_models = models{
   {
      slug: "ai21/jamba-large-1.7",
      url:  "https://studio.ai21.com",
      ok:   true,
   },
   {
      slug: "ai21/jamba-mini-1.7",
      url:  "https://studio.ai21.com",
      ok:   true,
   },
   //////////////////////////////////////////////////////////////////////////////
   {
      slug: "anthropic/claude-3.5-haiku-20241022",
      info: legacyModel,
   },
   {
      slug: "anthropic/claude-3.5-sonnet",
      info: legacyModel,
   },
   {
      slug: "anthropic/claude-3.5-sonnet-20240620",
      info: legacyModel,
   },
   {
      slug: "anthropic/claude-3.7-sonnet",
      info: legacyModel,
   },
   {
      slug: "anthropic/claude-opus-4",
      info: legacyModel,
   },
   {
      slug: "anthropic/claude-opus-4.1",
      url:  "https://claude.ai",
      ok:   true,
   },
   {
      slug: "anthropic/claude-sonnet-4",
      url:  "https://claude.ai",
      ok:   true,
   },
   //////////////////////////////////////////////////////////////////////////////
   {
      slug: "arcee-ai/spotlight",
      url:  "https://api.together.ai/playground/arcee_ai/arcee-spotlight",
      ok:   true,
   },
   {
      slug: "arcee-ai/maestro-reasoning",
      url:  "https://api.together.ai/playground/arcee-ai/maestro-reasoning",
      ok:   true,
   },
   {
      slug: "arcee-ai/virtuoso-large",
      url:  "https://api.together.ai/playground/arcee-ai/virtuoso-large",
      ok:   true,
   },
   //////////////////////////////////////////////////////////////////////////////
   {
      slug: "bytedance/ui-tars-1.5-7b",
      url:  "https://openrouter.ai/chat?models=bytedance/ui-tars-1.5-7b",
      ok:   true,
   },
   //////////////////////////////////////////////////////////////////////////////
   {
      slug: "deepseek/deepseek-chat-v3.1",
      url:  "https://chat.deepseek.com",
      ok:   true,
   },
   {
      slug: "deepseek/deepseek-r1-0528",
      url:  "https://deepinfra.com/deepseek-ai/DeepSeek-R1-0528",
      ok:   true,
   },
   {
      slug: "deepseek/deepseek-r1",
      url:  "https://deepinfra.com/deepseek-ai/DeepSeek-R1",
      ok:   true,
   },
   {
      slug: "deepseek/deepseek-prover-v2",
      url:  "https://deepinfra.com/deepseek-ai/DeepSeek-Prover-V2-671B",
      info: "due to low usage this model has been replaced by\n      deepseek-ai/DeepSeek-V3-0324",
   },
   {
      slug: "deepseek/deepseek-chat-v3-0324",
      url:  "https://deepinfra.com/deepseek-ai/DeepSeek-V3-0324",
      ok:   true,
   },
   {
      slug: "deepseek/deepseek-r1-distill-llama-70b",
      url:  "https://deepinfra.com/deepseek-ai/DeepSeek-R1-Distill-Llama-70B",
      ok:   true,
   },
   {
      slug: "deepseek/deepseek-r1-distill-qwen-32b",
      url:  "https://deepinfra.com/deepseek-ai/DeepSeek-R1-Distill-Qwen-32B",
      ok:   true,
   },
   //////////////////////////////////////////////////////////////////////////////
   {
      slug: "google/gemini-2.0-flash-001",
      info: legacyModel,
   },
   {
      slug: "google/gemini-2.0-flash-lite-001",
      info: legacyModel,
   },
   {
      slug: "google/gemini-pro-1.5",
      info: legacyModel,
   },
   {
      slug: "google/gemini-flash-1.5",
      info: legacyModel,
   },
   {
      slug: "google/gemini-flash-1.5-8b",
      info: legacyModel,
   },
   {
      slug: "google/gemini-2.5-pro-preview",
      info: previewModel,
   },
   {
      slug: "google/gemini-2.5-flash-lite-preview-06-17",
      info: previewModel,
   },
   {
      slug: "google/gemini-2.5-pro-preview-05-06",
      info: previewModel,
   },
   {
      slug: "google/gemini-2.5-pro",
      url:  "https://gemini.google.com",
      ok:   true,
   },
   {
      slug: "google/gemini-2.5-flash",
      url:  "https://gemini.google.com",
      ok:   true,
   },
   {
      slug: "google/gemini-2.5-flash-lite",
      url:  "https://aistudio.google.com/prompts/new_chat?model=gemini-2.5-flash-lite",
      ok:   true,
   },
   //////////////////////////////////////////////////////////////////////////////
   {
      slug: "inception/mercury",
      url:  "https://chat.inceptionlabs.ai",
      ok:   true,
   },
   {
      slug: "inception/mercury-coder",
      url:  "https://chat.inceptionlabs.ai",
      ok:   true,
   },
   //////////////////////////////////////////////////////////////////////////////
   {
      slug: "meta-llama/llama-4-scout",
      url:  "https://deepinfra.com/meta-llama/Llama-4-Scout-17B-16E-Instruct",
      ok:   true,
   },
   {
      slug: "meta-llama/llama-4-maverick",
      url:  "https://deepinfra.com/meta-llama/Llama-4-Maverick-17B-128E-Instruct-FP8",
      ok:   true,
   },
   {
      slug: "meta-llama/llama-guard-4-12b",
      url:  "https://deepinfra.com/meta-llama/Llama-Guard-4-12B",
      ok:   true,
   },
   //////////////////////////////////////////////////////////////////////////////
   {
      slug: "microsoft/mai-ds-r1",
      info: mayTrain,
   },
   //////////////////////////////////////////////////////////////////////////////
   {
      slug: "minimax/minimax-m1",
      url:  "https://chat.minimax.io",
      ok:   true,
   },
   //////////////////////////////////////////////////////////////////////////////
   {
      slug: "mistralai/mistral-medium-3.1",
      info: legacyModel,
   },
   {
      slug: "mistralai/codestral-2508",
      url:  "https://console.mistral.ai/build/playground",
      ok:   true,
   },
   {
      slug: "mistralai/devstral-medium",
      url:  "https://console.mistral.ai/build/playground",
      ok:   true,
   },
   {
      slug: "mistralai/devstral-small-2505",
      url:  "https://console.mistral.ai/build/playground",
      ok:   true,
   },
   {
      slug: "mistralai/mistral-medium-3",
      info: legacyModel,
   },
   {
      slug: "mistralai/devstral-small",
      info: legacyModel,
   },
   {
      slug: "mistralai/mistral-small-3.1-24b-instruct",
      info: legacyModel,
   },
   {
      slug: "mistralai/mistral-small-3.2-24b-instruct",
      url:  "https://deepinfra.com/mistralai/Mistral-Small-3.2-24B-Instruct-2506",
      ok:   true,
   },
   //////////////////////////////////////////////////////////////////////////////
   {
      slug: "moonshotai/kimi-dev-72b",
      url:  "https://cloud.siliconflow.com/playground/chat",
      ok:   true,
   },
   {
      slug: "moonshotai/kimi-vl-a3b-thinking",
      info: mayTrain,
   },
   //////////////////////////////////////////////////////////////////////////////
   {
      slug: "nousresearch/hermes-4-70b",
      url:  "https://chat.nousresearch.com",
      ok:   true,
   },
   {
      slug: "nousresearch/hermes-4-405b",
      url:  "https://chat.nousresearch.com",
      ok:   true,
   },
   //////////////////////////////////////////////////////////////////////////////
   {
      slug: "nvidia/llama-3.3-nemotron-super-49b-v1",
      url:  "https://build.nvidia.com/nvidia/llama-3_3-nemotron-super-49b-v1",
      ok:   true,
   },
   {
      slug: "nvidia/llama-3.1-nemotron-ultra-253b-v1",
      url:  "https://build.nvidia.com/nvidia/llama-3_1-nemotron-ultra-253b-v1",
      ok:   true,
   },
   //////////////////////////////////////////////////////////////////////////////
   {
      slug: "openai/gpt-4.1",
      info: legacyModel,
   },
   {
      slug: "openai/gpt-4.1-mini",
      info: legacyModel,
   },
   {
      slug: "openai/gpt-4.1-nano",
      info: legacyModel,
   },
   {
      slug: "openai/gpt-4o",
      info: legacyModel,
   },
   {
      slug: "openai/gpt-4o-2024-05-13",
      info: legacyModel,
   },
   {
      slug: "openai/gpt-4o-2024-08-06",
      info: legacyModel,
   },
   {
      slug: "openai/gpt-4o-2024-11-20",
      info: legacyModel,
   },
   {
      slug: "openai/gpt-4o-audio-preview",
      info: legacyModel,
   },
   {
      slug: "openai/gpt-4o-mini",
      info: legacyModel,
   },
   {
      slug: "openai/gpt-4o-mini-2024-07-18",
      info: legacyModel,
   },
   {
      slug: "openai/gpt-4o-mini-search-preview",
      info: legacyModel,
   },
   {
      slug: "openai/gpt-4o-search-preview",
      info: legacyModel,
   },
   {
      slug: "openai/o3",
      info: legacyModel,
   },
   {
      slug: "openai/o4-mini",
      info: legacyModel,
   },
   {
      slug: "openai/o4-mini-high",
      info: legacyModel,
   },
   {
      slug: "openai/codex-mini",
      info: legacyModel,
   },
   {
      slug: "openai/gpt-oss-120b",
      url:  "https://gpt-oss.com",
      ok:   true,
   },
   {
      slug: "openai/gpt-oss-20b",
      url:  "https://gpt-oss.com",
      ok:   true,
   },
   {
      slug: "openai/gpt-5-chat",
      url:  "https://chatgpt.com?model=gpt-5",
      ok:   true,
   },
   {
      slug: "openai/gpt-5-mini",
      url:  "https://platform.openai.com/docs/models/gpt-5-mini",
      ok:   true,
   },
   {
      slug: "openai/gpt-5-nano",
      url:  "https://platform.openai.com/docs/models/gpt-5-nano",
      ok:   true,
   },
   {
      slug: "openai/gpt-5",
      url:  "https://platform.openai.com/docs/models/gpt-5",
      ok:   true,
   },
   //////////////////////////////////////////////////////////////////////////////
   {
      slug: "perplexity/r1-1776",
      info: `provider returned error: {"error":{"code":400}}`,
   },
   {
      slug: "perplexity/sonar-deep-research",
      info: legacyModel,
   },
   {
      slug: "perplexity/sonar-pro",
      url:  "https://playground.perplexity.ai",
      ok:   true,
   },
   {
      slug: "perplexity/sonar-reasoning-pro",
      url:  "https://playground.perplexity.ai",
      ok:   true,
   },
   //////////////////////////////////////////////////////////////////////////////
   {
      slug: "qwen/qwen3-8b",
      url:  "https://openrouter.ai/chat?models=qwen/qwen3-8b",
      ok:   true,
   },
   {
      slug: "qwen/qwen3-30b-a3b-instruct-2507",
      url:  "https://chat.qwen.ai",
      ok:   true,
   },
   {
      slug: "qwen/qwen3-30b-a3b-thinking-2507",
      url:  "https://chat.qwen.ai",
      ok:   true,
   },
   {
      slug: "qwen/qwen3-235b-a22b-2507",
      url:  "https://chat.qwen.ai",
      ok:   true,
   },
   {
      slug: "qwen/qwen3-235b-a22b-thinking-2507",
      url:  "https://chat.qwen.ai",
      ok:   true,
   },
   {
      slug: "qwen/qwen3-coder",
      url:  "https://chat.qwen.ai",
      ok:   true,
   },
   {
      slug: "qwen/qwen3-coder-30b-a3b-instruct",
      url:  "https://openrouter.ai/chat?models=qwen/qwen3-coder-30b-a3b-instruct",
      ok:   true,
   },
   {
      slug: "qwen/qwq-32b",
      url:  "https://openrouter.ai/chat?models=qwen/qwen3-32b",
      ok:   true,
   },
   //////////////////////////////////////////////////////////////////////////////
   {
      slug: "switchpoint/router",
      url:  "https://switchpoint.dev/chat",
      ok:   true,
   },
   //////////////////////////////////////////////////////////////////////////////
   {
      slug: "tngtech/deepseek-r1t-chimera",
      info: mayTrain,
   },
   //////////////////////////////////////////////////////////////////////////////
   {
      slug: "x-ai/grok-3",
      url:  "https://grok.com",
      ok:   true,
   },
   {
      slug: "x-ai/grok-3-beta",
      info: previewModel,
   },
   {
      slug: "x-ai/grok-3-mini",
      url:  "https://openrouter.ai/chat?models=x-ai/grok-3-mini",
      ok:   true,
   },
   {
      slug: "x-ai/grok-3-mini-beta",
      info: previewModel,
   },
   {
      slug: "x-ai/grok-4",
      url:  "https://grok.com",
      ok:   true,
   },
   {
      slug: "x-ai/grok-code-fast-1",
      url:  "https://openrouter.ai/chat?models=x-ai/grok-code-fast-1",
      ok:   true,
   },
   //////////////////////////////////////////////////////////////////////////////
   {
      slug: "z-ai/glm-4.5",
      url:  "https://chat.z.ai",
      ok:   true,
   },
   {
      slug: "z-ai/glm-4.5-air",
      url:  "https://chat.z.ai",
      ok:   true,
   },
   {
      slug: "z-ai/glm-4-32b",
      url:  "https://chat.z.ai",
      ok:   true,
   },
}
