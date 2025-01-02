package structColumnDefs

import (
	"encoding/json"
	"fmt"
	"os"
	"parser/pkg/buildHeaderLineTypeUtilities"
	"parser/pkg/structColumnTypes"
	"strings"
)

/*
This package is used to define the structs that are used to parse the data for MET output files.
The entry point to this package is the ParseLine function that takes a header line, a data line, a fileType,
and a pointer to a map of data structures indexed by a string representation of a dataKey. The header line
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

func ParseLine(headerLine string, dataLine string, fileType string, doc map[string]interface{}) (map[string]interface{}, error) {
	_err := error(nil)
	if headerLine == "" || dataLine == "" {
		return doc, _err
	}
	// get the lineType
	lineType, headerData, dataData, dataKey, err := buildHeaderLineTypeUtilities.GetLineType(headerLine, dataLine, fileType)
	if err != nil {
		// cannot process this line so return the doc as is - it is probably a truncated line
		fmt.Println("Error getting line type: ", err)
		return doc, nil
	}
	// create a tmpHeaderData and remove the NA values fromm the headerData - we also have to do this in the GetDocFoId function
	tmpHeaderData := []string{}
	for _, h := range headerData {
		if h != "NA" && h != "" {
			tmpHeaderData = append(tmpHeaderData, h)
		}
	}
	// make the headerData NA values into "" so that the fillXXXX_Header functions make those values empty
	for i, h := range headerData {
		if h == "NA" {
			headerData[i] = ""
		}
	}

	if doc == nil {
		newDoc := make(map[string]interface{})
		doc = newDoc
	}
	id := strings.Join(tmpHeaderData, ":")
	_, exists := doc[id]
	fileLineType := fileType + "_" + lineType
	if !exists {
		metaData := structColumnTypes.VxMetadata{ID: id, Subset: "MET", Type: "DD", SubType: "MET"}
		doc[id], _err = structColumnTypes.GetDocForId(fileLineType, metaData, headerData, dataData, dataKey)
		if _err != nil {
			return doc, _err
		}
	} else {
		tempDoc := doc[id].(map[string]interface{})
		_err := structColumnTypes.AddDataElement(dataKey, fileLineType, dataData, &tempDoc)
		if _err != nil {
			return doc, _err
		}
	}
	return doc, _err
}

func WriteJsonToFile(doc map[string]interface{}, filename string) error {
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
	return os.WriteFile(filename, jsonBytes, 0644)
}
