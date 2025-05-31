# METlinetypes

This is a parser for MET output files.

## Approach

The approach is to use a [Go program](https://github.com/NOAA-GSL/METstat2json/blob/main/generator/generator.go) that uses several files from the MET repo (which is versioned by METplus) to generate another GO package that can then be imported to a GO program that can be used for parsing MET output files. The MET output file versions that are supported are v10.0.0 and later.

The file [column defs](https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/data/table_files/met_header_columns_V12.0.txt) as well as met_header_columns_v11.1.0, ...\_v11.0.0, ...\_v10.1.0, and ...\_v10.0.0 are used to get a list of required terms that need to be type defined for eacg respective version. Then several source code files are searched for type conversion statements for those terms. If the type of a term cannot be determined from source code an attempt is made to look up the term in the MET user guide.

These are the src files that are searched...

- [mode_line.cc](https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/src/libcode/vx_analysis_util/mode_line.cc)
- [stat_job.cc](https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/src/libcode/vx_analysis_util/stat_job.cc)
- [unstructured_grid.cc](https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/src/libcode/vx_grid/unstructured_grid.cc)
- [prob_info_base.cc](https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/src/libcode/vx_tc_util/prob_info_base.cc)
- [prob_rirw_info.cc](https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/src/libcode/vx_tc_util/prob_rirw_info.cc)
- [prob_rirw_pair_info.cc](https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/src/libcode/vx_tc_util/prob_rirw_pair_info.cc)
- [track_pair_info.cc](https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/src/libcode/vx_tc_util/track_pair_info.cc)
- [parse_stat_line.cc](https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/src/tools/core/stat_analysis/parse_stat_line.cc)
- [tc_stat_job.cc](https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/src/tools/tc_utils/tc_stat/tc_stat_job.cc)

... and these are the user guide files that are searched ...

- [ensemble-stat.rst](https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/docs/Users_Guide/ensemble-stat.rst)
- [grid-stat.rst](https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/docs/Users_Guide/grid-stat.rst)
- [gsi-tools.rst](https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/docs/Users_Guide/gsi-tools.rst)
- [mode-td.rst](https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/docs/Users_Guide/mode-td.rst)
- [mode.rst](https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/docs/Users_Guide/mode.rst)
- [point-stat.rst](https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/docs/Users_Guide/point-stat.rst)
- [stat-analysis.rst](https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/docs/Users_Guide/stat-analysis.rst)
- [tc-gen.rst](https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/docs/Users_Guide/tc-gen.rst)
- [tc-pairs.rst](https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/docs/Users_Guide/tc-pairs.rst)
- [wavelet-stat.rst](https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/docs/Users_Guide/wavelet-stat.rst)

... in order to derive column definitions and type information that is used to create the GO packages ...

- [metLineTypeDefinitions_v10_0](https://github.com/NOAA-GSL/METstat2json/tree/main/pkg/metLineTypeDefinitions_v10_0)
- [metLineTypeDefinitions_v10_1](https://github.com/NOAA-GSL/METstat2json/tree/main/pkg/metLineTypeDefinitions_v10_1)
- [metLineTypeDefinitions_v11_0](https://github.com/NOAA-GSL/METstat2json/tree/main/pkg/metLineTypeDefinitions_v11_0)
- [metLineTypeDefinitions_v11_0](https://github.com/NOAA-GSL/METstat2json/tree/main/pkg/metLineTypeDefinitions_v11_0)
- [metLineTypeDefinitions_v11_1](https://github.com/NOAA-GSL/METstat2json/tree/main/pkg/metLineTypeDefinitions_v11_1)
- [metLineTypeDefinitions_v12_0](https://github.com/NOAA-GSL/METstat2json/tree/main/pkg/metLineTypeDefinitions_v12_0)

These packages contain the struct definitions, fill functions for each MET line type, and parse routines necessary to convert MET output files into json documents for use in a GSL AVID Couchbase database, according to the AVID Couchbase data schema.

In addition to the "metLineTypeDefinitions_v..." packages there are several other local packages, "metLineTypeParser", "sample_parser", "generator", and "buildHeaderLineTypeUtilities".
The "sample_parser" package demonstrates how to use the [metLineTypeParser](https://github.com/NOAA-GSL/METstat2json/tree/main/pkg/metLineTypeParser) package. The metLineTypeParser is the only package that is required to parse MET output files.

### metLineTypeParser

This package is used to parse the data for MET output files.

The entry point to this package is the ParseLine function that takes a header line, a data line, a fileType, a dataSetName, and a pointer to a map of documents indexed by the document id. The document pointer can be an empty document. The header line is the first line of the file and contains the header field names. The data line is any subsequent line of the file that contains the header and the data fields.

The fileType is a string that represents the type of file being parsed. The dataSetName is a required user defined string that will identify the particular group of data that is being parsed. This allows the database to keep seperate copies of data that have the same Line type, the same header fields and the same vlaid or init times e.g. output from multiple MET runs of the same data set. The docPtr is a pointer to a map of documents that are indexed by an id that is derived from the header fields minus the dataKey fields and the additional dataSetName.

A dataKey is an array of header field values. For example most line types have a dataKey of {"Fcst_lead"} which would have a string representation of the value of the "Fcst_lead" element in the header string.

The dataKey fields are disallowed from the header id and are not included in the headerData. These keys serve the purpose of actually merging line data with the same dataKey values into a single document. The dataKey is used to index the data section of the document, which is a map[string]interface{}, where the interface is a specific concrete data type.

The ParseLine function uses the GetLineType function to determine the lineType of the data line, the headerData, dataKey, and descIndex. headerData are the ordered data fields for the header section of the line, the dataKey is the actual dataKey i.e the concatenated dataKey values, and the descIndex is the ordinal index of the desc field.

The descIndex is used to trim the desc field to 10 characters.

The parseLine function also uses the GetId function to determine the id of the data line. The id is derived from the headerData minus the dataKey fields and is returned in the form of a VxMetaDa ta struct. The VxMetaData struct is then converted to a map[string]interface{} so that it can be passed to the GetDocForId function without the GetDocForId function needing to know the VxMetaData struct type.

There are a couple of utility functions that are used to get the headerData without the NA values and to convert the VxMetaData struct.
A document pointer is required as a place to store the parsed data. If the document is nil, a new document is created.
The header line values (minus the dataKey fields) are used to derive the id, with date fields converted to epochs.
If the data section of of the document[id] is nil, a new data section is created. The data section is then populated with the data fields from the data line. If the data section is not nil, the data fields are added to the existing data map

### buildHeaderLineTypeUtilities

This package contains the utilities for building the header and line type for the data files.
The package exists to avoid some circular dependencies because both the parsing and the "generator" generation program depend on these utilities.

The data files are MET output files that contain a header section and a data section. The header section contains the header fields that are used to identify the document. The data section contains the data fields that are used to populate the document.

This package is separate from the metLineTypeDefinitions package because the metLineTypeDefinitions package is automatically generated from the generator.go program and there is a desire to avoid a circular dependency. This package defines a VxMetaData struct that is used to store the metadata for the mapped documents. The metadata is used to uniquely identify each document and is used to merge documents with the same metadata.

This package also defines the DataKeyMap that is used to determine the key data fields for a given line type. The key data fields are used to merge documents with the same header field values excluding the key data fields. The key data fields can be either from the header or the data section of a record line. In addition there is a headerDisallow field that may have a list of header section fields that must be disallowed from the header structure.

### header key data field

If the key data field is type header then the associated array of fields must be from the header section. The fields will be concatenated in the order that they are defined and separated by "\_". The fields will be excluded from the header and the resultant key string will be used to index the data map. Each data map entry will represent a different data record from the source file resulting in the merging of all the records that have the same header keys minus the key data fields.

### data key data field

If the key data field is type data then the associated array of fields must all be from the data section of the record. The fields will be concatenated in the order that they are defined and separated by an "\_". Those fields wont be excluded from the data section but the data section will be partitioned into a map that is indexed by the values from the OBECT_ID field.

When the dataKeyFields above is applied to this document the lines with identical headers will be merged into a single document (as well as the lines with matching headers from other processed files) but because of the dataKeyMap type data definition the data records with the same OBJECT_ID will be mapped into a data map in that single document and that map will be indexed by the OBJECT_ID field. There is also a headerDisallow field and any header fields that are in this list will not be included in the header structure. This allows line definitions that have INIT, VALID, and LEAD times to have multiple data entries e.g. keyed by LEAD.

### other utilities

Other utilities exist to convert the date fields to epochs, to get the line type of the data line, to get the key data fields for a given line type, and to find the data type of a given field in addition to some other utility functions.

## generator

The output of this program is a series of structs that can be used to define the header
and data types in the buildHeaderTypes.go file and some parsing routines that are aware of the
header and data types. The outout is the metLineTypeDefinitions.go file which is the sole content of the metLineTypeDefinitions package and should never be directly edited.

This is how to build the metLineTypeDefinitions package. It is suggested redirect the output (which is the go program) to a temporary file and then after looking at it copy it to its proper destination. You may have to create the ...pkg/metLineTypeDefinitions/ directory, but it probably comes with the repo clone.

```console
go run generator -version=v12.0 > /pkg/metLineTypeDefinitions_v12_0/metLineTypeDefinitions.go
```

## Usage

This package is used to parse the data for MET output files.

The entry point to this package is the ParseLine function that takes a header line, a data line, a fileType, and a map of documents indexed by the document id. The document pointer can be an empty document. The header line is the first line of the file and contains the header field names.  The data line is any subsequent line of the file that contains the header and the data fields. The fileType is a string that represents the type of file being parsed. The docPtr is a pointer to a map of documents that are indexed by an id that is derived from the header fields minus the dataKey fields.

A dataKey is an array of either header or data field values. For example most line types have a dataKey of {"Fcst_lead"} from the header section which would have a string representation of the value of the "Fcst_lead" element in the header string. The dataKey fields are disallowed from the header id and are not included in the headerData. These keys serve the purpose of actually merging line data with the same dataKey values into a single document. The dataKey is used to index the data section of the document, which is a map[string]interface{}, where the interface is a specific concrete data type.

A dataKey can also be a data field name. In that case the data section of a single document will be divided into a map based on that dataKey.

The ParseLine function uses the GetLineType function to determine the lineType of the data line, the headerData, dataKey, and descIndex. headerData are the ordered data fields for the header section of the line, the dataKey is the actual dataKey i.e the concatenated dataKey values, and the descIndex is the ordinal index of the desc field.

The descIndex is used to trim the desc field to 10 characters.

The parseLine function also uses the GetId function to determine the id of the data line. The id is derived from the headerData minus the dataKey fields and is returned in the form of a VxMetaData struct. The VxMetaData struct is then converted to a map[string]interface{} so that it can be passed to the GetDocForId function without the GetDocForId function needing to know the VxMetaData struct type.

There are a couple of utility functions that are used to get the headerData without the NA values and to convert the VxMetaData struct. A document pointer is required as a place to store the parsed data. If the document is nil, a new document is created. The header line values (minus the dataKey fields) are used to derive the id, with date fields converted to epochs. If the data section of of the document[id] is nil, a new data section is created. The data section is then populated with the data fields from the data line. If the data section is not nil, the data fields are added to the existing data map.

### Formatting the generated structColumnDefs

There is a scripts directory at the top of the repo with a format-lint.sh script that can be used to format and lint the go code in the repo, including the derived structColumnDefs package. It must be run from the root of the repo.

## Testing

There are unit tests in the buildHeaderLineTypeUtilities package in the file buildHeaderLineTypeUtilities_test.go. You can run these in vscode or on the command line.
In vscode there is usually a link above the test case function name that allows you to run or debug the test case. For the command line this is how I do it

There are also integration level tests in the structColumnDefs/definitions_test.go file. These tests use the testdata directory. That test directory has its own repo [testdata](https://github.com/NOAA-GSL/MET-parser-testdata). It has most (if not all) of the various output file formats from the 12.0 version of MET, and enough data to get a good feel for performance. As more kinds of data are added to subsequent releases of MET this testdata directory can be modified to include sample data for new ouputfile versions. These tests are good examples of how to use the parser.

```text
ranpierce-mac1:buildHeaderLineTypeUtilities randy.pierce$ cd /Users/randy.pierce/metlinetypes/pkg/buildHeaderLineTypeUtilities/
ranpierce-mac1:buildHeaderLineTypeUtilities randy.pierce$ go clean --testcache
ranpierce-mac1:buildHeaderLineTypeUtilities randy.pierce$ go test ./...
ok   parser/pkg/buildHeaderLineTypeUtilities 0.185s
ranpierce-mac1:buildHeaderLineTypeUtilities randy.pierce$ cd /Users/randy.pierce/metlinetypes/pkg/structColumnDefs
ranpierce-mac1:buildHeaderLineTypeUtilities randy.pierce$ go clean --testcache
ranpierce-mac1:structColumnDefs randy.pierce$ go test ./...
ok  parser/pkg/structColumnDefs 15.705s

```

or if I want more log output....

```text
ranpierce-mac1:buildHeaderLineTypeUtilities randy.pierce$ cd /Users/randy.pierce/metlinetypes/pkg/buildHeaderLineTypeUtilities/
ranpierce-mac1:buildHeaderLineTypeUtilities randy.pierce$ go clean --testcache
ranpierce-mac1:buildHeaderLineTypeUtilities randy.pierce$ go test -v ./...
=== RUN   TestGetLineType
=== RUN   TestGetLineType/VERSION_MODEL_DESC_FCST_LEAD_FCST_VALID_BEG__FCST_VALID_END__OBS_LEAD_OBS_VALID_BEG___OBS_VALID_END___FCST_VAR__FCST_UNITS_FCST_LEV_OBS_VAR___OBS_UNITS_OBS_LEV_OBTYPE_VX_MASK_INTERP_MTHD_INTERP_PNTS_FCST_THRESH_OBS_THRESH_COV_THRESH_ALPHA_LINE_TYPE
=== RUN   TestGetLineType/VERSION_MODEL_DESC_FCST_LEAD_FCST_VALID_BEG__FCST_VALID_END__OBS_LEAD_OBS_VALID_BEG___OBS_VALID_END___FCST_VAR__FCST_UNITS_FCST_LEV_OBS_VAR___OBS_UNITS_OBS_LEV_OBTYPE_VX_MASK_INTERP_MTHD_INTERP_PNTS_FCST_THRESH_OBS_THRESH_COV_THRESH_ALPHA_LINE_TYPE#01
=== RUN   TestGetLineType/VERSION_MODEL_DESC_FCST_LEAD_FCST_VALID_BEG__FCST_VALID_END__OBS_LEAD_OBS_VALID_BEG___OBS_VALID_END___FCST_VAR__FCST_UNITS_FCST_LEV_OBS_VAR___OBS_UNITS_OBS_LEV_OBTYPE_VX_MASK_INTERP_MTHD_INTERP_PNTS_FCST_THRESH_OBS_THRESH_COV_THRESH_ALPHA_LINE_TYPE#02
--- PASS: TestGetLineType (0.00s)
    --- PASS: TestGetLineType/VERSION_MODEL_DESC_FCST_LEAD_FCST_VALID_BEG__FCST_VALID_END__OBS_LEAD_OBS_VALID_BEG___OBS_VALID_END___FCST_VAR__FCST_UNITS_FCST_LEV_OBS_VAR___OBS_UNITS_OBS_LEV_OBTYPE_VX_MASK_INTERP_MTHD_INTERP_PNTS_FCST_THRESH_OBS_THRESH_COV_THRESH_ALPHA_LINE_TYPE (0.00s)
    --- PASS: TestGetLineType/VERSION_MODEL_DESC_FCST_LEAD_FCST_VALID_BEG__FCST_VALID_END__OBS_LEAD_OBS_VALID_BEG___OBS_VALID_END___FCST_VAR__FCST_UNITS_FCST_LEV_OBS_VAR___OBS_UNITS_OBS_LEV_OBTYPE_VX_MASK_INTERP_MTHD_INTERP_PNTS_FCST_THRESH_OBS_THRESH_COV_THRESH_ALPHA_LINE_TYPE#01 (0.00s)
    --- PASS: TestGetLineType/VERSION_MODEL_DESC_FCST_LEAD_FCST_VALID_BEG__FCST_VALID_END__OBS_LEAD_OBS_VALID_BEG___OBS_VALID_END___FCST_VAR__FCST_UNITS_FCST_LEV_OBS_VAR___OBS_UNITS_OBS_LEV_OBTYPE_VX_MASK_INTERP_MTHD_INTERP_PNTS_FCST_THRESH_OBS_THRESH_COV_THRESH_ALPHA_LINE_TYPE#02 (0.00s)
=== RUN   TestSplitColumnDefLine
=== RUN   TestSplitColumnDefLine/MODE
=== RUN   TestSplitColumnDefLine/MTD
=== RUN   TestSplitColumnDefLine/STAT
=== RUN   TestSplitColumnDefLine/STAT#01
--- PASS: TestSplitColumnDefLine (0.00s)
    --- PASS: TestSplitColumnDefLine/MODE (0.00s)
    --- PASS: TestSplitColumnDefLine/MTD (0.00s)
    --- PASS: TestSplitColumnDefLine/STAT (0.00s)
    --- PASS: TestSplitColumnDefLine/STAT#01 (0.00s)
=== RUN   TestGetKeyDataFieldsForLineType
=== RUN   TestGetKeyDataFieldsForLineType/MODE
=== RUN   TestGetKeyDataFieldsForLineType/MTD
=== RUN   TestGetKeyDataFieldsForLineType/STAT
=== RUN   TestGetKeyDataFieldsForLineType/unknown
--- PASS: TestGetKeyDataFieldsForLineType (0.00s)
    --- PASS: TestGetKeyDataFieldsForLineType/MODE (0.00s)
    --- PASS: TestGetKeyDataFieldsForLineType/MTD (0.00s)
    --- PASS: TestGetKeyDataFieldsForLineType/STAT (0.00s)
    --- PASS: TestGetKeyDataFieldsForLineType/unknown (0.00s)
=== RUN   TestFindType
=== RUN   TestFindType/item1
=== RUN   TestFindType/item2
=== RUN   TestFindType/item3
=== RUN   TestFindType/item4
--- PASS: TestFindType (0.00s)
    --- PASS: TestFindType/item1 (0.00s)
    --- PASS: TestFindType/item2 (0.00s)
    --- PASS: TestFindType/item3 (0.00s)
    --- PASS: TestFindType/item4 (0.00s)
PASS
ok   parser/pkg/buildHeaderLineTypeUtilities 0.184s
```

There are more comprehensive tests in the structColumnDefs pkg.

**NOTE: There is an error in one of the data files.** This error is expected because the file is truncated.

```text
Error getting line type:  UNPARSABLE_LINE: lineTypeIndex is greater than the length of the data line
Skipping line: V12.0.0 GFS   NA   1080000   20241101_180000 20241101_180000 000000   20241101_180000 20241101_180000 RH   because it isn't parsable from file grid_stat_GFS_TMP_vs_ANLYS_TMP_Z2_1080000L_20241101_180000V.stat
.../testdata/G2G_v12/20241102-00z 96
```

```text
ranpierce-mac1:structColumnDefs randy.pierce$ go test -v ./...
=== RUN   TestParseVAL1L2
--- PASS: TestParseVAL1L2 (0.00s)
=== RUN   TestParseRegressionSuite
--- PASS: TestParseRegressionSuite (3.29s)
=== RUN   TestParseG2G_v12_Suite
../../testdata/G2G_v12 672
.../testdata/G2G_v12/20241031-00z 96
.../testdata/G2G_v12/20241031-00z/grid_stat 1280
.../testdata/G2G_v12/20241031-06z 96
.../testdata/G2G_v12/20241031-06z/grid_stat 1248
.../testdata/G2G_v12/20241031-12z 96
.../testdata/G2G_v12/20241031-12z/grid_stat 1280
.../testdata/G2G_v12/20241031-18z 96
.../testdata/G2G_v12/20241031-18z/grid_stat 1248
.../testdata/G2G_v12/20241101-00z 96
.../testdata/G2G_v12/20241101-00z/grid_stat 1280
.../testdata/G2G_v12/20241101-06z 96
.../testdata/G2G_v12/20241101-06z/grid_stat 1248
.../testdata/G2G_v12/20241101-12z 96
.../testdata/G2G_v12/20241101-12z/grid_stat 1280
.../testdata/G2G_v12/20241101-18z 96
.../testdata/G2G_v12/20241101-18z/grid_stat 1248
Error getting line type:  UNPARSABLE_LINE: lineTypeIndex is greater than the length of the data line
Skipping line: V12.0.0 GFS   NA   1080000   20241101_180000 20241101_180000 000000   20241101_180000 20241101_180000 RH   because it isn't parsable from file grid_stat_GFS_TMP_vs_ANLYS_TMP_Z2_1080000L_20241101_180000V.stat
.../testdata/G2G_v12/20241102-00z 96
.../testdata/G2G_v12/20241102-00z/grid_stat 1280
.../testdata/G2G_v12/20241102-06z 96
.../testdata/G2G_v12/20241102-06z/grid_stat 1248
.../testdata/G2G_v12/20241102-12z 96
.../testdata/G2G_v12/20241102-12z/grid_stat 1280
.../testdata/G2G_v12/20241102-18z 96
.../testdata/G2G_v12/20241102-18z/grid_stat 1248
.../testdata/G2G_v12/20241103-00z 96
.../testdata/G2G_v12/20241103-00z/grid_stat 1280
.../testdata/G2G_v12/20241103-06z 96
.../testdata/G2G_v12/20241103-06z/grid_stat 1248
.../testdata/G2G_v12/20241103-12z 96
.../testdata/G2G_v12/20241103-12z/grid_stat 1280
.../testdata/G2G_v12/20241103-18z 96
.../testdata/G2G_v12/20241103-18z/grid_stat 1248
.../testdata/G2G_v12/20241104-00z 96
.../testdata/G2G_v12/20241104-00z/grid_stat 1280
.../testdata/G2G_v12/20241104-06z 96
.../testdata/G2G_v12/20241104-06z/grid_stat 1248
.../testdata/G2G_v12/20241104-18z 96
.../testdata/G2G_v12/20241104-18z/grid_stat 1120
--- PASS: TestParseG2G_v12_Suite (11.22s)
PASS
ok   parser/pkg/structColumnDefs 14.742s
```

## sample program

There is a sample program called regression.go in pkg/regression/regregression.go.
To build this program use something like...

```text
    go build -o /tmp/regression pkg/regression/regression.go
```

To run this program you must provide a path to a data directory...

```text
/tmp/regression -path=.../testdata/statfiles/
```

## build and install

See [Go's build and install tutorial](https://go.dev/doc/tutorial/compile-install).

There is a sample build script for the sample_parser in scripts/build.sh. This is an example of how it might be used.

```console
scripts/build.sh
bin/darwin/arm64/sample_parser -outdir /tmp -path .../testdata
gunzip /tmp/sample_output.json.gz
jq . /tmp/sample_output.json > /tmp/sample_output-pretty.json
vim /tmp/sample_output-pretty.json
```
