# jamba-large-1.7

https://studio.ai21.com

## prompt 1, 19s

return the full Go script that:
1. Reads an MPEG-DASH MPD XML file path from the CLI: `go run main.go <mpd_file_path>`
2. Uses `http://test.test/test.mpd` as the initial MPD URL for resolving all relative BaseURLs
3. Outputs a JSON object mapping each `Representation@id` to a list of fully resolved segment URLs
