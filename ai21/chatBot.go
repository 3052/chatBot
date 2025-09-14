package main

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/url"
  "os"
  "path"
  "strings"
)

type MPD struct {
  BaseURLs     []string
  Periods      []Period
  Representations map[string]Representation
}

type Period struct {
  AdaptationSets []AdaptationSet
}

type AdaptationSet struct {
  Representations []Representation
}

type Representation struct {
  ID string
  BaseURL string
  SegmentBase SegmentBase
}

type SegmentBase struct {
  Initialization Initialization
  RepresentationIndex RepresentationIndex
}

type Initialization struct {
  SourceURL string
}

type RepresentationIndex struct {
  SourceURL string
}

func main() {
  if len(os.Args) != 2 {
    fmt.Println("Usage: go run main.go <mpd_file_path>")
    return
  }

  mpdFilePath := os.Args[1]
  mpdFileContent, err := ioutil.ReadFile(mpdFilePath)
  if err != nil {
    fmt.Println("Error reading MPD file:", err)
    return
  }

  var mpd MPD
  if err := json.Unmarshal(mpdFileContent, &mpd); err != nil {
    fmt.Println("Error unmarshalling MPD file:", err)
    return
  }

  initialMPDURL, err := url.Parse("http://test.test/test.mpd")
  if err != nil {
    fmt.Println("Error parsing initial MPD URL:", err)
    return
  }

  resolvedRepresentationSegmentURLs := resolveRepresentationSegmentURLs(mpd, initialMPDURL)

  jsonOutput, err := json.Marshal(resolvedRepresentationSegmentURLs)
  if err != nil {
    fmt.Println("Error marshalling JSON output:", err)
    return
  }

  fmt.Println(string(jsonOutput))
}

func resolveRepresentationSegmentURLs(mpd MPD, initialMPDURL *url.URL) map[string][]string {
  resolvedURLs := make(map[string][]string)

  for _, period := range mpd.Periods {
    for _, adaptationSet := range period.AdaptationSets {
      for _, representation := range adaptationSet.Representations {
        representationURL := resolveURL(representation.BaseURL, initialMPDURL)

        if representation.SegmentBase.Initialization.SourceURL != "" {
          resolvedInitializationURL := resolveURL(representation.SegmentBase.Initialization.SourceURL, initialMPDURL)
          resolvedURLs[representation.ID] = append(resolvedURLs[representation.ID], resolvedInitializationURL)
        }

        if representation.SegmentBase.RepresentationIndex.SourceURL != "" {
          resolvedRepresentationIndexURL := resolveURL(representation.SegmentBase.RepresentationIndex.SourceURL, initialMPDURL)
          resolvedURLs[representation.ID] = append(resolvedURLs[representation.ID], resolvedRepresentationIndexURL)
        }
      }
    }
  }

  return resolvedURLs
}

func resolveURL(relativeURL string, baseURL *url.URL) string {
  relativeURLParsed, err := url.Parse(relativeURL)
  if err != nil {
    fmt.Println("Error parsing relative URL:", err)
    return ""
  }

  resolvedURL := baseURL.ResolveReference(relativeURLParsed)
  return resolvedURL.String()
}
