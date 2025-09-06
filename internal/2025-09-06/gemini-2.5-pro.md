# google/gemini-2.5-pro

return the full Go script that:
1. Reads an MPEG-DASH MPD XML file path from the CLI: `go run main.go <mpd_file_path>`
2. Uses `http://test.test/test.mpd` as the initial MPD URL for resolving all relative BaseURLs
3. Outputs a JSON object mapping each `Representation@id` to a list of fully resolved segment URLs

<https://aistudio.google.com/prompts/new_chat?model=gemini-2.5-pro>

~~~sh
MODEL_ID="gemini-2.5-pro"
GENERATE_CONTENT_API="streamGenerateContent"

cat << EOF > request.json
{
    "contents": [
      {
        "role": "user",
        "parts": [
          {
            "text": "INSERT_INPUT_HERE"
          },
        ]
      },
    ],
    "generationConfig": {
      "thinkingConfig": {
        "thinkingBudget": -1,
      },
    },
    "tools": [
      {
        "googleSearch": {
        }
      },
    ],
}
EOF

curl \
-X POST \
-H "Content-Type: application/json" \
"https://generativelanguage.googleapis.com/v1beta/models/${MODEL_ID}:${GENERATE_CONTENT_API}?key=${GEMINI_API_KEY}" -d '@request.json'
~~~
