# METlinetypes

This is a parser for MET output files.

## Approach

The approach is to use a [GO program](https://github.com/ian-noaa/metlinetypes/blob/82d936b3aed5128676fed135bfda65897deb2c5a/pkg/buildHeaderLineTypes/buildHeaderLineTypes.go) that uses several files from the MET repo (which is versioned) to create a GO package that can then be
used to create a GO package that can be used for parsing MET output files.
The file [column defs](https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/data/table_files/met_header_columns_V12.0.txt) is used to get a list of required terms that need to be type defined. Then several source code files are searched for type conversion statements for those terms. If the type of a term cannot be determined from source code an attempt is made to look up the term in the MET user guide.

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

... and these are the usr guide files that are searched ...

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

... to derive column definitions and type information which is used to create the GO package "pkg/structColumnTypes" that contains the struct definitions, fill functions, and parse routines necessary to convert MET output files into json documents for use in a GSL AVID Couchbase database, according to the AVID Couchbase data schema.

In addition to the GO pkg "structColumnTypes" there are three other local packages, "structColumnDefs", "buildHeaderLineTypes", and "buildHeaderLineTypeUtilities".

### structColumnDefs

This package is used to parse the data for MET output files.
The entry point to this package is the ParseLine function that takes a header line, a data line, a fileType, and a map of documents indexed by the document id. The document pointer can be an empty document. The header line is the first line of the file and contains the header field names.  The data line is any subsequent line of the file that contains the header and the data fields.
The fileType is a string that represents the type of file being parsed. The docPtr is a pointer to a map of documents that are indexed by an id that is derived from the header fields minus the dataKey fields.

A dataKey is an array of header field values. For example most line types have a dataKey of {"Fcst_lead"} which would have a string representation of the value of the "Fcst_lead" element in the header string.
The dataKey fields are disallowed from the header id and are not included in the headerData. These keys serve the purpose of actually merging line data with the same dataKey values into a single document. The dataKey is used to index the data section of the document, which is a map[string]interface{}, where the interface is a specific concrete data type.

The ParseLine function uses the GetLineType function to determine the lineType of the data line, the headerData, dataKey, and descIndex. headerData are the ordered data fields for the header section of the line, the dataKey is the actual dataKey i.e the concatenated dataKey values, and the descIndex is the ordinal index of the desc field.
The descIndex is used to trim the desc field to 10 characters.

The parseLine function also uses the getId function to determine the id of the data line. The id is derived from the headerData minus the dataKey fields and is returned in the form of a VxMetaDa ta struct. The VxMetaData struct is then converted to a map[string]interface{}
so that it can be passed to the GetDocForId function without the GetDocForId function needing to know the VxMetaData struct type.

There are a couple of utility functions that are used to get the headerData without the NA values and to convert the VxMetaData struct.
A document pointer is required as a place to store the parsed data. If the document is nil, a new document is created.
The header line values (minus the dataKey fields) are used to derive the id, with date fields converted to epochs.
If the data section of of the document[id] is nil, a new data section is created. The data section is then populated with the data fields from the data line. If the data section is not nil, the data fields are added to the existing data map


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
The key data fields can be either from the header or the data section of a record line.

### header key data field

If the key data field is type header then the associated array of fields must be from the header section. The fields will be concatenated in the order that they are defined and separated by "_". The fields will be excluded from the header and the resultant key string will be used to index the data map. Each data map entry will represent a different data record from the source file resulting in the merging of all the records that have the same header keys minus the key data fields.

### data key data field

If the key data field is type data then the associated array of fields must all be from the data section of the record. The fields will be concatenated in the order that they are defined and separated by an "_". Those fields wont be excluded from the data section but the data section will be partitioned into a map that is indexed by the values from the OBECT_ID field.

When the dataKeyFields above is applied to this document the lines with identical headers will be merged into a single document (as well as the lines with matching headers from other processed files) but because of the dataKeyMap type data definition the data records with the same OBJECT_ID will be mapped into a data map in that single document and that map will be indexed by the OBJECT_ID field.

### other utilities

Other utilities exist to convert the date fields to epochs, to get the line type of the data line, to get the key
data fields for a given line type, and to find the data type of a given field in addition to some other utility functions.

## buildHeaderLineTypes

The output of this program is a series of structs that can be used to define the header
and data types in the buildHeaderTypes.go file and some parsing routines that are aware of the
header and data types. The outout is the structColumnTypes.go file which is the sole content of the structColumnTypes package and should never be directly edited.

This is how to build the structColumnTypes package. It is suggested redirect the output (which is the go program)
to a temporary file and then after looking at it copy it to its proper destination.
You may have to create the ...pkg/structColumnTypes/ directory, but it probably comes with the repo clone.

```text
ranpierce-mac1:buildHeaderLineTypes randy.pierce$ cd /Users/randy.pierce/metlinetypes/pkg/buildHeaderLineTypes
ranpierce-mac1:buildHeaderLineTypes randy.pierce$ go run . > /tmp/types.go
ranpierce-mac1:buildHeaderLineTypes randy.pierce$ cp /tmp/types.go ../structColumnTypes/structColumnTypes.go
```

## Usage

This package is used to parse the data for MET output files.
The entry point to this package is the ParseLine function that takes a header line, a data line, a fileType,
and a map of documents indexed by the document id. The document pointer can be an empty document. The header line
is the first line of the file and contains the header field names.  The data line is any subsequent line of the file
that contains the header and the data fields.
The fileType is a string that represents the type of file being parsed. The docPtr is a pointer to a map of
documents that are indexed by an id that is derived from the header fields minus the dataKey fields.

A dataKey is an array of either header or data field values. For example most line types have a dataKey of {"Fcst_lead"}
from the header section which would have a string representation of the value of the "Fcst_lead" element in the header string.
The dataKey fields are disallowed from the header id and are not included in the headerData. These keys serve
the purpose of actually merging line data with the same dataKey values into a single document. The dataKey is used
to index the data section of the document, which is a map[string]interface{}, where the interface is a specific concrete
data type.

A dataKey can also be a data field name. In that case the data section of a single document will be divided into a map based on that dataKey.

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

### Formatting the generated structColumnDefs

There is a scripts directory at the top of the repo with a format-lint.sh script that can be used to format and lint the go code in the repo, including the derived structColumnDefs package. It must be run from the root of the repo.

## Testing

There are unit tests in the buildHeaderLineTypeUtilities package in the file buildHeaderLineTypeUtilities_test.go. You can run these in vscode or on the command line.
In vscode there is usually a link above the test case function name that allows you to run or debug the test case.
For the command line this is how I do it

```text
ranpierce-mac1:buildHeaderLineTypeUtilities randy.pierce$ cd /Users/randy.pierce/metlinetypes/pkg/buildHeaderLineTypeUtilities/
ranpierce-mac1:buildHeaderLineTypeUtilities randy.pierce$ go clean --testcache
ranpierce-mac1:buildHeaderLineTypeUtilities randy.pierce$ go test ./...
ok   parser/pkg/buildHeaderLineTypeUtilities 0.185s
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

**NOTE: There is an error in one of the data files.** This error is expected.

```text
Error getting line type:  UNPARSABLE_LINE: lineTypeIndex is greater than the length of the data line
Skipping line: V12.0.0 GFS   NA   1080000   20241101_180000 20241101_180000 000000   20241101_180000 20241101_180000 RH   because it isn't parsable from file grid_stat_GFS_TMP_vs_ANLYS_TMP_Z2_1080000L_20241101_180000V.stat
/Users/randy.pierce/metlinetypes/test_data/G2G_v12/20241102-00z 96
```

```text
cd /Users/randy.pierce/metlinetypes/pkg/structColumnDefs/
go clean --testcache
ranpierce-mac1:structColumnDefs randy.pierce$ go test -v ./...
=== RUN   TestParseVAL1L2
--- FAIL: TestParseVAL1L2 (0.00s)
panic: interface conversion: interface {} is nil, not map[string]interface {} [recovered]
 panic: interface conversion: interface {} is nil, not map[string]interface {}

goroutine 4 [running]:
testing.tRunner.func1.2({0x1047a0500, 0x14000073800})
 /opt/homebrew/Cellar/go/1.23.3/libexec/src/testing/testing.go:1632 +0x1bc
testing.tRunner.func1()
 /opt/homebrew/Cellar/go/1.23.3/libexec/src/testing/testing.go:1635 +0x334
panic({0x1047a0500?, 0x14000073800?})
 /opt/homebrew/Cellar/go/1.23.3/libexec/src/runtime/panic.go:785 +0x124
parser/pkg/structColumnDefs.TestParseVAL1L2(0x1400007a9c0)
 /Users/randy.pierce/metlinetypes/pkg/structColumnDefs/definitions_test.go:100 +0x754
testing.tRunner(0x1400007a9c0, 0x1047e20b8)
 /opt/homebrew/Cellar/go/1.23.3/libexec/src/testing/testing.go:1690 +0xe4
created by testing.(*T).Run in goroutine 1
 /opt/homebrew/Cellar/go/1.23.3/libexec/src/testing/testing.go:1743 +0x314
FAIL parser/pkg/structColumnDefs 0.208s
FAIL
ranpierce-mac1:structColumnDefs randy.pierce$ go test -v ./...
=== RUN   TestParseVAL1L2
--- PASS: TestParseVAL1L2 (0.00s)
=== RUN   TestParseRegressionSuite
--- PASS: TestParseRegressionSuite (3.29s)
=== RUN   TestParseG2G_v12_Suite
/Users/randy.pierce/metlinetypes/pkg/structColumnDefs/../../test_data/G2G_v12 672
/Users/randy.pierce/metlinetypes/test_data/G2G_v12/20241031-00z 96
/Users/randy.pierce/metlinetypes/test_data/G2G_v12/20241031-00z/grid_stat 1280
/Users/randy.pierce/metlinetypes/test_data/G2G_v12/20241031-06z 96
/Users/randy.pierce/metlinetypes/test_data/G2G_v12/20241031-06z/grid_stat 1248
/Users/randy.pierce/metlinetypes/test_data/G2G_v12/20241031-12z 96
/Users/randy.pierce/metlinetypes/test_data/G2G_v12/20241031-12z/grid_stat 1280
/Users/randy.pierce/metlinetypes/test_data/G2G_v12/20241031-18z 96
/Users/randy.pierce/metlinetypes/test_data/G2G_v12/20241031-18z/grid_stat 1248
/Users/randy.pierce/metlinetypes/test_data/G2G_v12/20241101-00z 96
/Users/randy.pierce/metlinetypes/test_data/G2G_v12/20241101-00z/grid_stat 1280
/Users/randy.pierce/metlinetypes/test_data/G2G_v12/20241101-06z 96
/Users/randy.pierce/metlinetypes/test_data/G2G_v12/20241101-06z/grid_stat 1248
/Users/randy.pierce/metlinetypes/test_data/G2G_v12/20241101-12z 96
/Users/randy.pierce/metlinetypes/test_data/G2G_v12/20241101-12z/grid_stat 1280
/Users/randy.pierce/metlinetypes/test_data/G2G_v12/20241101-18z 96
/Users/randy.pierce/metlinetypes/test_data/G2G_v12/20241101-18z/grid_stat 1248
Error getting line type:  UNPARSABLE_LINE: lineTypeIndex is greater than the length of the data line
Skipping line: V12.0.0 GFS   NA   1080000   20241101_180000 20241101_180000 000000   20241101_180000 20241101_180000 RH   because it isn't parsable from file grid_stat_GFS_TMP_vs_ANLYS_TMP_Z2_1080000L_20241101_180000V.stat
/Users/randy.pierce/metlinetypes/test_data/G2G_v12/20241102-00z 96
/Users/randy.pierce/metlinetypes/test_data/G2G_v12/20241102-00z/grid_stat 1280
/Users/randy.pierce/metlinetypes/test_data/G2G_v12/20241102-06z 96
/Users/randy.pierce/metlinetypes/test_data/G2G_v12/20241102-06z/grid_stat 1248
/Users/randy.pierce/metlinetypes/test_data/G2G_v12/20241102-12z 96
/Users/randy.pierce/metlinetypes/test_data/G2G_v12/20241102-12z/grid_stat 1280
/Users/randy.pierce/metlinetypes/test_data/G2G_v12/20241102-18z 96
/Users/randy.pierce/metlinetypes/test_data/G2G_v12/20241102-18z/grid_stat 1248
/Users/randy.pierce/metlinetypes/test_data/G2G_v12/20241103-00z 96
/Users/randy.pierce/metlinetypes/test_data/G2G_v12/20241103-00z/grid_stat 1280
/Users/randy.pierce/metlinetypes/test_data/G2G_v12/20241103-06z 96
/Users/randy.pierce/metlinetypes/test_data/G2G_v12/20241103-06z/grid_stat 1248
/Users/randy.pierce/metlinetypes/test_data/G2G_v12/20241103-12z 96
/Users/randy.pierce/metlinetypes/test_data/G2G_v12/20241103-12z/grid_stat 1280
/Users/randy.pierce/metlinetypes/test_data/G2G_v12/20241103-18z 96
/Users/randy.pierce/metlinetypes/test_data/G2G_v12/20241103-18z/grid_stat 1248
/Users/randy.pierce/metlinetypes/test_data/G2G_v12/20241104-00z 96
/Users/randy.pierce/metlinetypes/test_data/G2G_v12/20241104-00z/grid_stat 1280
/Users/randy.pierce/metlinetypes/test_data/G2G_v12/20241104-06z 96
/Users/randy.pierce/metlinetypes/test_data/G2G_v12/20241104-06z/grid_stat 1248
/Users/randy.pierce/metlinetypes/test_data/G2G_v12/20241104-18z 96
/Users/randy.pierce/metlinetypes/test_data/G2G_v12/20241104-18z/grid_stat 1120
--- PASS: TestParseG2G_v12_Suite (11.22s)
PASS
ok   parser/pkg/structColumnDefs 14.742s

```
