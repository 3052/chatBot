package main

import (
    "encoding/json"
    "flag"
    "fmt"
    "io/ioutil"
    "log"
    "net/url"
    "strings"
    "encoding/xml"
)

// MPD struct
type MPD struct {
    XMLName         xml.Name `xml:"MPD"` 
    BaseURLs        []string `xml:"BaseURL"`
    Periods         []Period `xml:"Period"`
    MPDFilePath     string
    InitialMPD      string
    ResolvedURLs    map[string][]string
}

// Period struct
type Period struct {
    AdaptationSets []AdaptationSet `xml:"AdaptationSet"`
}

// AdaptationSet struct
type AdaptationSet struct {
    Representations []Representation `xml:"Representation"`
}

// Representation struct
type Representation struct {
    ID       string `xml:"id,attr"`
    BaseURL  string `xml:"BaseURL"`
    SegmentBase *SegmentBase `xml:"SegmentBase"`
    SegmentList *SegmentList `xml:"SegmentList"`
    SegmentTemplate *SegmentTemplate `xml:"SegmentTemplate"`
}

// SegmentBase struct
type SegmentBase struct {
    InitializationInitialization string `xml:"Initialization@initialization"`
}

// Segment struct within SegmentList
type Segment struct {
    URL string `xml:"URL"`
}

// SegmentList struct
type SegmentList struct {
    Segments []Segment `xml:"Segment"`
}

// SegmentTemplate struct
type SegmentTemplate struct {
    Media string `xml:"media,attr"`
    StartNumber int `xml:"startNumber,attr"`
    EndNumber int `xml:"endNumber,attr"`
}

func main() {
    flag.Parse()

    mpdFilePath := flag.Arg(0)
    if mpdFilePath == "" {
        log.Fatalf("Please provide an MPD file path.")
    }

    initialMPD := `http://test.test/test.mpd`

    // Read MPD file
    mpdXML, err := ioutil.ReadFile(mpdFilePath)
    if err != nil {
        log.Fatalf("Error reading MPD file %s: %v", mpdFilePath, err)
    }

    var mpd MPD
    err = xml.Unmarshal(mpdXML, &mpd)
    if err != nil {
        log.Fatalf("Error unmarshaling MPD XML: %v", err)
    }

    mpd.MPDFilePath = mpdFilePath
    mpd.InitialMPD = initialMPD

    // Resolve all BaseURLs
    baseURLsMap := map[string]string{}
    for idx, baseURL := range mpd.BaseURLs {
        if !strings.HasPrefix(baseURL, "http") {
            baseURL = resolveURL(initialMPD, baseURL)
        }
        baseURLsMap[fmt.Sprintf("BaseURL[%d]", idx)] = baseURL
    }

    mpd.ResolvedURLs = map[string][]string{}

    // Process each Representation
    for _, period := range mpd.Periods {
        for _, adaptationSet := range period.AdaptationSets {
            for _, representation := range adaptationSet.Representations {
                representationID := representation.ID
                if representationID == "" {
                    log.Fatalf("Representation ID is missing.")
                }

                // Resolve BaseURL for Representation
                representationBaseURL := representation.BaseURL
                if !strings.HasPrefix(representationBaseURL, "http") {
                    representationBaseURL = resolveURL(initialMPD, representationBaseURL)
                }

                initializationURL := resolveURL(initialMPD, representation.SegmentBase.InitializationInitialization)

                segmentURLs := []string{}

                // get segment list or template URLs
                if representation.SegmentList != nil {
                    for _, segment := range representation.SegmentList.Segments {
                        segmentURL := resolveSegmentURL(initialMPD, segment)
                        segmentURLs = append(segmentURLs, segmentURL)
                    }
                }

                // get segment template URLs, if any
                if representation.SegmentTemplate != nil {
                    segmentTemplateURLs := generateSegmentTemplateURLs(initialMPD, representation.SegmentTemplate)
                    segmentURLs = append(segmentURLs, segmentTemplateURLs...)
                }

                mpd.ResolvedURLs[representationID] = append(segmentURLs, initializationURL)
            }
        }
    }

    // Convert to JSON
    output, err := json.MarshalIndent(mpd.ResolvedURLs, "", "  ")
    if err != nil {
        log.Fatalf("Error marshaling to JSON: %v", err)
    }

    // Output JSON
    fmt.Println(string(output))
}

func resolveURL(base string, relative string) string {
    if strings.HasPrefix(relative, "http") {
        return relative
    }

    baseURL, err := url.Parse(base)
    if err != nil {
        log.Fatalf("Error parsing base URL %s: %v", base, err)
    }

    relativeURL, err := url.Parse(relative)
    if err != nil {
        log.Fatalf("Error parsing relative URL %s: %v", relative, err)
    }

    return baseURL.ResolveReference(relativeURL).String()
}

func resolveSegmentURL(initialMPD string, segment Segment) string {
    if !strings.HasPrefix(segment.URL, "http") {
        return resolveURL(initialMPD, segment.URL)
    }
    return segment.URL
}

func generateSegmentTemplateURLs(initialMPD string, segmentTemplate *SegmentTemplate) []string {
    segmentURLs := []string{}

    // iterate to get URLs for all segments
    for i := segmentTemplate.StartNumber; i <= segmentTemplate.EndNumber; i++ {
        segmentURL := resolveURL(initialMPD, fmt.Sprintf(segmentTemplate.Media, i))
        segmentURLs = append(segmentURLs, segmentURL)
    }

    return segmentURLs
}
