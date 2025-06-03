package metLineTypeParser

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	metLineTypeDefinitions_v10_0 "github.com/NOAA-GSL/METstat2json/pkg/metLineTypeDefinitions_v10_0"
	metLineTypeDefinitions_v10_1 "github.com/NOAA-GSL/METstat2json/pkg/metLineTypeDefinitions_v10_1"
	metLineTypeDefinitions_v11_0 "github.com/NOAA-GSL/METstat2json/pkg/metLineTypeDefinitions_v11_0"
	metLineTypeDefinitions_v11_1 "github.com/NOAA-GSL/METstat2json/pkg/metLineTypeDefinitions_v11_1"
	metLineTypeDefinitions_v12_0 "github.com/NOAA-GSL/METstat2json/pkg/metLineTypeDefinitions_v12_0"

	buildHeaderLineTypeUtilities "github.com/NOAA-GSL/METstat2json/pkg/buildHeaderLineTypeUtilities"
)

/*
Supportoing versions FYI, these are all the >= 10.0.0 releases of MET that exist: 10.0 10.1 11.0 11.1 12.0
This package is used to parse the data for MET output files.
The entry point to this package is the ParseLine function that takes a data set name, header line, a data line, a fileType,
and a map of documents indexed by the document id. The document pointer can be an empty document. The data set name is a string that identifies the actual MET dataset.
For example a MET user may run the same thing multiple times and without a unique data set name for
each run the id fields of the JSON documents in the parsed output would be the same and the data would
overwrite itself in the database. So the data set name is required and must be unique. The header line
is the first line of the file and contains the header field names.  The data line is any subsequent line of the file
that contains the header and the data fields.
The fileType is a string that represents the type of file being parsed. The docPtr is a pointer to a map of
documents that are indexed by an id that is derived from the header fields minus the dataKey fields.

A dataKey is an array of header field values. For example most line types have a dataKey of {"Fcst_lead"}
which would have a string representation of the value of the "Fcst_lead" element in the header string.
The dataKey fields are disallowed from the header id and are not included in the headerData. These keys serve
the purpose of actually merging line data with the same dataKey values into a single document. The dataKey is used
to index the data section of the document, which is a map[string]interface{}, where the interface is a specific concrete
data type.

The ParseLine function uses the GetLineType function to determine the lineType of the data line, the headerData,
dataKey, and descIndex. headerData are the ordered data fields for the header section of the line, the dataKey is the actual
dataKey i.e the concatenated dataKey values, and the descIndex is the ordinal index of the desc field.
The descIndex is used to trim the desc field to 10 characters.

The parseLine function also uses the
GetId function to determine the id of the data line. The id is derived from the headerData minus the dataKey fields and
is returned in the form of a VxMetadata struct. The VxMetadata struct is then converted to a map[string]interface{}
so that it can be passed to the GetDocForId function without the GetDocForId function needing to know the VxMetadata struct type.

There are a couple of utility functions that are used to get the headerData without the NA values and to convert the VxMetadata struct.
A document pointer is required as a place to store the parsed data. If the document is nil, a new document is created.
The header line values (minus the dataKey fields) are used to derive the id, with date fields converted to epochs.
If the data section of of the document[id] is nil, a new data section is created. The data section is then populated
with the data fields from the data line. If the data section is not nil, the data fields are added to the existing data map.

The parameter getExternalDocForId is a function pointer that is used to get an external document for a given id. This function
is used to get a document from an external source, such as a database, that is indexed by the id. If the external document
is not nil, it is added to the document map. If the external document is nil, a new document is created for the id.
*/

const DOC_NOT_FOUND = "document not found"

func getParserVersion(dataLine string) (string, error) {
	metVersion := strings.ToLower(strings.Fields(dataLine)[0])
	metVersionParts := strings.Split(metVersion, ".")
	if len(metVersionParts) != 3 {
		return "", fmt.Errorf("invalid MET version format: %s", metVersion)
	}
	lineVersion := metVersionParts[0] + "_" + metVersionParts[1]
	return lineVersion, nil
}

func ParseLine(dataSetName string, headerLine string, dataLine string, docPtr *map[string]interface{}, fileName string, getExternalDocForId func(id string) (map[string]interface{}, error)) (map[string]interface{}, error) {
	// recover from unexpected errors
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered: %v for fileName %s\n", r, fileName)
		}
	}()

	if dataSetName == "" {
		return *docPtr, fmt.Errorf("dataSetName is empty")
	}
	if len(dataSetName) > 10 {
		return *docPtr, fmt.Errorf("dataSetName is too long - must be <= 10 characters")
	}
	// get line version e.g. V12.0.0 -> v12_0
	parserVersion, _err := getParserVersion(dataLine)
	if _err != nil {
		return *docPtr, fmt.Errorf("error getting parser version from line %s: %w", dataLine, _err)
	}
	if headerLine == "" {
		return *docPtr, fmt.Errorf("empty header line")
	}
	if dataLine == "" {
		return *docPtr, fmt.Errorf("empty data line")
	}
	// make sure we have the basename here
	fileName = filepath.Base(fileName)
	filePathParts := strings.Split(fileName, ".")
	fileType := strings.ToUpper(filePathParts[1])
	if fileType == "SWP" {
		// skip the swp files - might be editing a file and don't want to parse the .swp file
		return *docPtr, fmt.Errorf("skipping swp file")
	}
	if fileType == "DS_STORE" {
		// skip the .DS_Store files
		return *docPtr, fmt.Errorf("skipping .DS_Store file")
	}

	// get the lineType
	fileLineType, headerData, dataData, dataKey, descIndex, err := buildHeaderLineTypeUtilities.GetLineType(headerLine, dataLine, fileName, parserVersion)
	if err != nil {
		// cannot process this line so return the docPtr as is - it is probably a truncated line
		fmt.Println("Error getting line type: ", err)
		return *docPtr, err
	}
	// if there are any disallowed fields in this linetype then add the disallowed data to the dataData array - in order
	disallowedFields := buildHeaderLineTypeUtilities.DataKeyMap[fileLineType].HeaderDisallow
	if len(disallowedFields) > 0 {
		for _, disallowedField := range disallowedFields {
			disAllowedFieldValue, err := buildHeaderLineTypeUtilities.GetHeaderValue(strings.Fields(headerLine), strings.Fields(dataLine), disallowedField)
			// if there is an error getting the disallowed field, just append "" to the dataData array
			if err != nil {
				fmt.Println("Error getting disallowed field: ", err)
			}
			// if an err it appends ""
			dataData = append(dataData, disAllowedFieldValue)
		}
	}

	// get the tmpHeaderData without the NA values
	tmpHeaderData := getTmpHeaderSanNA(headerData, descIndex)
	if *docPtr == nil {
		newDoc := make(map[string]interface{})
		docPtr = &newDoc
	}
	// GetId will fill in the id field of the metaData struct with the constructed id
	// metadata doesn't change between versions, we just use the latest one. Same with DOC
	metaData, _err := buildHeaderLineTypeUtilities.GetId(dataSetName, tmpHeaderData, &buildHeaderLineTypeUtilities.VxMetadata{Subset: "MET", Type: "DD", SubType: "MET"})
	if _err != nil {
		return *docPtr, fmt.Errorf("error getting id from line %s: %w", dataLine, _err)
	}
	_, exists := (*docPtr)[metaData.ID]
	if !exists {
		// check to see if there is an existing external document for this id
		externalExistingDoc, err := (getExternalDocForId)(metaData.ID)
		if err != nil && !strings.HasPrefix(err.Error(), DOC_NOT_FOUND) {
			return *docPtr, err
		}
		// if there is an external document for this id, use it, we will add the data into it
		if externalExistingDoc != nil {
			(*docPtr)[metaData.ID] = externalExistingDoc
		} else {
			// have to create a new document for this id
			metaDataMap, _err := getMetaDataMap(metaData)
			if _err != nil {
				return *docPtr, _err
			}
			// create a new document for the new metaData.ID
			// This function will also fill in the headerData fields
			// indexed by dataKey value in the document.
			// The document needs to be of the correct version.
			switch parserVersion {
			case "v10_0":
				(*docPtr)[metaData.ID], _err = metLineTypeDefinitions_v10_0.GetDocForId(fileLineType, metaDataMap, headerData, dataData, dataKey)
			case "v10_1":
				(*docPtr)[metaData.ID], _err = metLineTypeDefinitions_v10_1.GetDocForId(fileLineType, metaDataMap, headerData, dataData, dataKey)
			case "v11_0":
				(*docPtr)[metaData.ID], _err = metLineTypeDefinitions_v11_0.GetDocForId(fileLineType, metaDataMap, headerData, dataData, dataKey)
			case "v11_1":
				(*docPtr)[metaData.ID], _err = metLineTypeDefinitions_v11_1.GetDocForId(fileLineType, metaDataMap, headerData, dataData, dataKey)
			case "v12_0":
				(*docPtr)[metaData.ID], _err = metLineTypeDefinitions_v12_0.GetDocForId(fileLineType, metaDataMap, headerData, dataData, dataKey)
			default:
				return *docPtr, fmt.Errorf("unsupported version %s", parserVersion)
			}
			if _err != nil || (*docPtr)[metaData.ID] == nil {
				return *docPtr, fmt.Errorf("error creating doc for file: %s error: %w", fileName, _err)
			}
			// add the dataSetName to the header - dataSetName is not part of the structure
			(*docPtr)[metaData.ID].(map[string]interface{})["dataSetName"] = dataSetName
			// return the new doc - the doc was created and the data was added to it
			return *docPtr, _err
		}
	} else {
		// we either had the doc already, got it externally, or created it
		// now we need to add the data to the document
		docMap := (*docPtr)[metaData.ID].(map[string]interface{})
		switch parserVersion {
		case "v10_0":
			// add the data to the document
			(*docPtr)[metaData.ID], _err = metLineTypeDefinitions_v10_0.AddDataElement(dataKey, fileLineType, dataData, &docMap)
		case "v10_1":
			// add the data to the document
			(*docPtr)[metaData.ID], _err = metLineTypeDefinitions_v10_1.AddDataElement(dataKey, fileLineType, dataData, &docMap)
		case "v11_0":
			// add the data to the document
			(*docPtr)[metaData.ID], _err = metLineTypeDefinitions_v11_0.AddDataElement(dataKey, fileLineType, dataData, &docMap)
		case "v11_1":
			// add the data to the document
			(*docPtr)[metaData.ID], _err = metLineTypeDefinitions_v11_1.AddDataElement(dataKey, fileLineType, dataData, &docMap)
		case "v12_0":
			// add the data to the document
			(*docPtr)[metaData.ID], _err = metLineTypeDefinitions_v12_0.AddDataElement(dataKey, fileLineType, dataData, &docMap)
		default:
			return *docPtr, fmt.Errorf("unsupported version %s", parserVersion)
		}
		if _err != nil {
			return *docPtr, fmt.Errorf("error getting doc for file: %s error: %w", fileName, _err)
		}
	}
	return *docPtr, _err
}

/*
convert the fields of the metaData to a map[string]interface{} so it can be added to the doc without needing the VxMetadata struct type definition
*/
func getMetaDataMap(metaData buildHeaderLineTypeUtilities.VxMetadata) (map[string]interface{}, error) {
	var metaDataMap map[string]interface{}
	jsonBytes, err := json.Marshal(metaData)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(jsonBytes, &metaDataMap)
	if err != nil {
		return nil, err
	}
	return metaDataMap, nil
}

/*
create a tmpHeaderData and remove the "" and the NA values fromm the headerData.
This also has to be done in the GetDocForId i.e. (fill_XXXX_Header) functions,
and trim the desc field data to 10 chars, if it isn't empty ("")
*/
func getTmpHeaderSanNA(headerData []string, descIndex int) []string {
	tmpHeaderData := []string{}
	for i, h := range headerData {
		if h != "NA" && h != "" {
			if i == descIndex {
				if len(h) > 10 {
					h = h[:10]
				}
			}
			tmpHeaderData = append(tmpHeaderData, h)
		}
	}
	// make the headerData NA values into "" so that the fillXXXX_Header functions make those values empty
	for i, h := range headerData {
		if h == "NA" {
			headerData[i] = ""
		}
	}
	return tmpHeaderData
}

func WriteJsonToCompressedFile(doc map[string]interface{}, filename string) error {
	// get the documents as a list
	// Defines the Slice capacity to match the Map elements count
	docList := make([]interface{}, 0, len(doc))

	for _, tx := range doc {
		docList = append(docList, tx)
	}
	// Marshal the document struct to JSON
	jsonBytes, err := json.Marshal(docList)
	if err != nil {
		return err
	}
	// Write the JSON bytes to a file
	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	_, err = w.Write(jsonBytes)
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	// Write the compressed data to a file
	err = os.WriteFile(filename, b.Bytes(), 0o644)
	if err != nil {
		fmt.Printf("Failed to write file: %v", err)
	}
	return nil
}
