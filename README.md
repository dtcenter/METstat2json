# METlinetypes

This is a parser for MET output files.

## Approach

The approach is to use two files from the MET repo (which is versioned)

```https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/src/tools/core/stat_analysis/parse_stat_line.cc```

and

```
https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/data/table_files/met_header_columns_V12.0.txt"
```

to derive column definitions and type information which is used to create a GO package "pkg/structColumnTypes" automatically that contains the struct definitions, fill functions, and parse routines necessary to convert MET output files into json documents for use in a GSL AVID Couchbase database, according to the AVID Couchbase data schema.

In addition to the GO pkg "structColumnTypes" there are three other local packages, "structColumnDefs", "buildHeaderLineTypes", and "buildHeaderLineTypeUtilities".

### structColumnDefs

This package is used to parse the data for MET output files.
This package is used to parse the data for MET output files.
The entry point to this package is the ParseLine function that takes a header line, a data line, a fileType,
and a map of documents indexed by the document id. The document pointer can be an empty document. The header line
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
getId function to determine the id of the data line. The id is derived from the headerData minus the dataKey fields and
is returned in the form of a VxMetaData struct. The VxMetaData struct is then converted to a map[string]interface{}
so that it can be passed to the GetDocForId function without the GetDocForId function needing to know the VxMetaData struct type.

There are a couple of utility functions that are used to get the headerData without the NA values and to convert the VxMetaData struct.
A document pointer is required as a place to store the parsed data. If the document is nil, a new document is created.
The header line values (minus the dataKey fields) are used to derive the id, with date fields converted to epochs.
If the data section of of the document[id] is nil, a new data section is created. The data section is then populated
with the data fields from the data line. If the data section is not nil, the data fields are added to the existing data map
*/

### buildHeaderLineTypeUtilities

This package contains the utilities for building the header and line type for the data files.
The package exists to avoid some circular dependencies because both the parsing and the buildHeaderLineTypes generation program depend on these utilities.

The data files are MET output files that contain a header section and a data section. The header section contains the header fields that are used to identify the document. The data section contains the data fields that are used to populate the document.

This package is separate from the structColumnTypes package because the structColumnTypes package is automatically
generated from the buildHeaderLineTypes.go program and there is a desire to avoid a circular dependency.
This package defines a VxMetaData struct that is used to store the metadata for the mapped documents.
The metadata is used to uniquely identify each document and is used to merge documents with the same metadata.

This package also defines the DataKeyMap that is used to determine the key data fields for a given line type.
The key data fields are used to merge documents with the same header field values excluding the key data fields.
Other utilities exist to convert the date fields to epochs, to get the line type of the data line, to get the key
data fields for a given line type, and to find the data type of a given field in addition to some other utility functions.

### buildHeaderLineTypes

The output of this program is a series of structs that can be used to define the header
and data types in the buildHeaderTypes.go file and some parsing routines that are aware of the
header and data types. The outout is the structColumnTypes.go file which is the sole content of the structColumnTypes package and should never be directly edited.

## Usage

This package is used to parse the data for MET output files.
The entry point to this package is the ParseLine function that takes a header line, a data line, a fileType,
and a map of documents indexed by the document id. The document pointer can be an empty document. The header line
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
getId function to determine the id of the data line. The id is derived from the headerData minus the dataKey fields and
is returned in the form of a VxMetaData struct. The VxMetaData struct is then converted to a map[string]interface{}
so that it can be passed to the GetDocForId function without the GetDocForId function needing to know the VxMetaData struct type.

There are a couple of utility functions that are used to get the headerData without the NA values and to convert the VxMetaData struct.
A document pointer is required as a place to store the parsed data. If the document is nil, a new document is created.
The header line values (minus the dataKey fields) are used to derive the id, with date fields converted to epochs.
If the data section of of the document[id] is nil, a new data section is created. The data section is then populated
with the data fields from the data line. If the data section is not nil, the data fields are added to the existing data map.
