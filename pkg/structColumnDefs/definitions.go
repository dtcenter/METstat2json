package structColumnDefs

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"os"
	"parser/pkg/buildHeaderLineTypeUtilities"
	"parser/pkg/structColumnTypes"
)

/*
This package is used to parse the data for MET output files.
The entry point to this package is the ParseLine function that takes a header line, a data line, a fileType,
and a map of documents indexed by the document id. The document pointer can be an empty document. The header line
is the first line of the file and contains the header field names. The data field names are optional and therefore
not used. The data line is any subsequent line of the file that contains the header and the data fields.
The fileType is a string that represents the type of file being parsed. The doc is a pointer to a map of
documents that are indexed by an id that is derived from the header fields without the dataKey fields.

A dataKey is an array of header field values, for example most line types have a dataKey of {"Fcst_lead"}
which would have a string representation of the value of the "Fcst_lead" element in the header string.
The dataKey fields are disallowed from the header id and are not included in the headerData. These keys serve
the purpose of actually merging line data with the same dataKey values into a single document. The dataKey is used
to index the data section of the document, which is a map[string]interface{}, where the interface is a specific concrete
data type.

The ParseLine function uses the getLineType function to determine the lineType of the data line. It also uses the
getId function to determine the id of the data line. The headerData and dataData are then extracted from the header
and data lines respectively. The metaData must be provided as a structColumnTypes.VxMetadata struct.

A document pointer is required as a place to store the parsed data. If the document is nil, a new document is created.
The header line values (minus the dataKey fields) are used to derive the id, with date fields converted to epochs.
If the data section of of the document[id] is nil, a new data section is created. The data section is then populated
with the data fields from the data line. If the data section is not nil, the data fields are added to the existing data map.
*/

func ParseLine(headerLine string, dataLine string, fileType string, docPtr *map[string]interface{}) (map[string]interface{}, error) {
	_err := error(nil)
	if headerLine == "" || dataLine == "" {
		return *docPtr, _err
	}
	// get the lineType
	lineType, headerData, dataData, dataKey, descIndex, err := buildHeaderLineTypeUtilities.GetLineType(headerLine, dataLine, fileType)
	if err != nil {
		// cannot process this line so return the docPtr as is - it is probably a truncated line
		fmt.Println("Error getting line type: ", err)
		return *docPtr, err
	}
	tmpHeaderData := getTmpHeaderSanNA(headerData, descIndex)
	if *docPtr == nil {
		newDoc := make(map[string]interface{})
		docPtr = &newDoc
	}
	// GetId will fill in the id field of the metaData struct with the constructed id
	metaData := buildHeaderLineTypeUtilities.GetId(tmpHeaderData, &buildHeaderLineTypeUtilities.VxMetadata{Subset: "MET", Type: "DD", SubType: "MET"})
	fileLineType := fileType + "_" + lineType
	_, exists := (*docPtr)[metaData.ID]
	if !exists {
		// TODO: check if the id exists in the database (has already been processed)
		// if it does, then we need to add the data from the database to the existing document here
		// need to do a fetch from the database to get the data
		// and then add the data to the existing document here doc[id] = data_from_the_db

		metaDataMap, _err := getMetaDataMap(metaData)
		if _err != nil {
			return *docPtr, _err
		}
		// create a new document for the new metaData.ID
		(*docPtr)[metaData.ID], _err = structColumnTypes.GetDocForId(fileLineType, metaDataMap, headerData, dataData, dataKey)
		if _err != nil {
			return *docPtr, _err
		}
	} else {
		tempDoc := (*docPtr)[metaData.ID].(map[string]interface{})
		_err := structColumnTypes.AddDataElement(dataKey, fileLineType, dataData, &tempDoc)
		if _err != nil {
			return *docPtr, _err
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
create a tmpHeaderData and remove the NA values fromm the headerData - we also have to do this in the GetDocFoId function
trim the desc field data to 10 chars, if it isn't empty ("")
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

func WriteJsonToGzipFile(doc map[string]interface{}, filename string) error {
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
	err = os.WriteFile(filename, b.Bytes(), 0644)
	if err != nil {
		fmt.Printf("Failed to write file: %v", err)
	}
	return nil
}
