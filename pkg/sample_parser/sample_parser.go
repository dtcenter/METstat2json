package main

import (
	"compress/gzip"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"github.com/NOAA-GSL/MET-parser/pkg/structColumnDefs"
	"github.com/NOAA-GSL/MET-parser/pkg/structColumnTypes"
	"path/filepath"
	"runtime"
	"strings"
)

// dummy function to satisfy the function signature of getExternalDocForId
func getMissingExternalDocForId(id string) (map[string]interface{}, error) {
	// fmt.Println("getExternalDocForId called with id:", id)
	// Put your own code here in this method but always return this exact error if the document is not found
	return nil, fmt.Errorf("%s: %s", structColumnTypes.DOC_NOT_FOUND, id)
}

func ReadJsonFromGzipFile(filename string) (map[string]interface{}, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	gz, err := gzip.NewReader(file)
	if err != nil {
		fmt.Print(err)
	}
	defer gz.Close()
	decoder := json.NewDecoder(gz)
	var result []interface{}
	err = decoder.Decode(&result)
	if err != nil {
		return nil, err
	}
	parsedDoc := make(map[string]interface{})
	for _, v := range result {
		elem := v.(map[string]interface{})
		parsedDoc[elem["id"].(string)] = elem
	}
	return parsedDoc, nil
}

func ParseRegressionSuite() {
	var doc map[string]interface{}
	var err error
	var test_data_directory string
	output_directory := "/tmp"
	Usage := func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.StringVar(&test_data_directory, "path", "", "Required - Path to the regression test data")
	flag.StringVar(&output_directory, "outdir", "", "Optional - Path to the output directory - defaults to /tmp")
	flag.Parse()
	if test_data_directory == "" {
		Usage()
		os.Exit(1)
	}
	if !strings.HasSuffix(output_directory, "/") {
		output_directory += "/"
	}
	// walk through the directory to read files using filepath.Walk
	err = filepath.Walk(test_data_directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() { // skip directories - we only want the files
			return nil
		}
		doc = parseFile(path, info, doc)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	// write output to json	gzipped file
	err = structColumnDefs.WriteJsonToCompressedFile(doc, output_directory+"sample_output.json.gz")
	if err != nil {
		log.Fatal(err)
	}
}

func parseFile(fPath string, fileInfos os.FileInfo, doc map[string]interface{}) map[string]interface{} {
	file, err := os.Open(fPath) // open the file
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fName := fileInfos.Name()
	if strings.HasSuffix(fName, ".swp") || strings.HasSuffix(fName, ".DS_Store") {
		// skip the swp files - might be editing a file and don't want to parse the .swp file
		// and skip any .DS_Store files
		return doc
	}
	rawData, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(rawData), "\n")
	if len(lines) == 0 || lines[0] == "" {
		log.Printf("empty file %s - skipping\n", fPath)
		return doc
	}
	headerLine := lines[0]
	if !strings.HasPrefix(headerLine, "VERSION") {
		log.Printf("missing VERSION at start of header line - bad header line? for file %s - skipping rest of file\n", fPath)
		return doc
	}
	for line := range lines {
		if line == 0 || lines[line] == "" {
			continue
		}
		dataLine := lines[line]
		doc, err = structColumnDefs.ParseLine(headerLine, dataLine, &doc, fName, getMissingExternalDocForId)
		if err != nil {
			log.Printf("Error parsing line: %s for file %s\n", err, fName)
		}
	}
	return doc
}

func main() {
	fmt.Println("environment:" + runtime.GOOS + "_" + runtime.GOARCH)
	ParseRegressionSuite()
}
