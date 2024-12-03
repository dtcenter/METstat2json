package listColumnDefs

import (
    "fmt"
    "sync"
)

/*
{
    "lineType": "VL1L2",
    "columns": [
        {
            "name": "TOTAL",
            "type": "int"
        },
        {
            "name": "UFBAR",
            "type": "float"
        },
        {
            "name": "VFBAR",
            "type": "float"
        },
        {
            "name": "UOBAR",
            "type": "float"
        },
        {
            "name": "VOBAR",
            "type": "float"
        },
        {
            "name": "UVFOBAR",
            "type": "float"
        },
        {
            "name": "UVFFBAR",
            "type": "float"
        },
        {
            "name": "UVOOBAR",
            "type": "float"
        },
        {
            "name": "F_SPEED_BAR",
            "type": "float"
        },
        {
            "name": "O_SPEED_BAR",
            "type": "float"
        }
    ]
}

{
    "lineType": "VAL1L2",
    "columns": [
        {
            "name": "TOTAL",
            "type": "int"
        },
        {
            "name": "UFABAR",
            "type": "float"
        },
        {
            "name": "VFABAR",
            "type": "float"
        },
        {
            "name": "UOABAR",
            "type": "float"
        },
        {
            "name": "VOABAR",
            "type": "float"
        },
        {
            "name": "UVFOABAR",
            "type": "float"
        },
        {
            "name": "UVFFBAR",
            "type": "float"
        },
        {
            "name": "UVOOABAR",
            "type": "float"
        },
        {
            "name": "FA_SPEED_BAR",
            "type": "float"
        },
        {
            "name": "OA_SPEED_BAR",
            "type": "float"
        }
    ]
}

*/


type ColumnDef struct {
	Name         string
	HeaderFields []string
	DataFields   []string
}

// column definitions - add to these all you like
// Each column definition has a name, a list of header fields and a list of data fields
// Each column definition is a struct that implements the Parser interface
// Each new ColumnDef needs to be added to the parserMap

var VL1L2 = ColumnDef{
	Name: "VL1L2",
	HeaderFields: []string{"VERSION","MODEL","DESC","LINE_TYPE","FCST_VALID_BEG","FCST_VALID_END","OBS_VALID_BEG","OBS_VALID_END","FCST_VAR","FCST_UNITS","FCST_LEV","FCST_THRESH","OBS_VAR","OBS_UNITS","OBS_LEV","OBTYPE","OBS_THRESH","INTERP_MTHD","INTERP_PNTS","VX_MASK",
	},
	DataFields: []string{"TOTAL", "UFBAR", "VFBAR", "UOBAR", "VOBAR", "UVFOBAR", "UVFFBAR", "UVOOBAR", "F_SPEED_BAR", "O_SPEED_BAR"},
}

var VAL1L2 = ColumnDef{
	Name: "VAL1L2",
	HeaderFields: []string{"VERSION","MODEL","DESC","LINE_TYPE","FCST_VALID_BEG","FCST_VALID_END","OBS_VALID_BEG","OBS_VALID_END","FCST_VAR","FCST_UNITS","FCST_LEV","FCST_THRESH","OBS_VAR","OBS_UNITS","OBS_LEV","OBTYPE","OBS_THRESH","INTERP_MTHD","INTERP_PNTS","VX_MASK",
	},
	DataFields: []string{"TOTAL", "UFABAR", "VFABAR", "UOABAR", "VOABAR", "UVFOABAR", "UVFFBAR", "UVOOABAR", "FA_SPEED_BAR", "OA_SPEED_BAR"},
}

var parserMap = map[string] ColumnDef{
    "VL1L2": VL1L2,
    "VAL1L2": VAL1L2,
}

var builderMap = map[string] ParserBuilder{}

// Don't change below this line

type Parser interface {
	Parse(string) string
}

type ParserBuilder interface {
	Columns(ColumnDef) ParserBuilder
	Build() Parser
}

type parserBuilder struct {
	columnDef ColumnDef
}

type parser struct {
	Columns ColumnDef
}

func (pb *parserBuilder) Columns(columns ColumnDef) ParserBuilder {
	pb.columnDef = columns
	return pb
}

func (pb *parserBuilder) Build() Parser {
	return &parser{
		Columns: pb.columnDef,
	}
}

var lock = &sync.Mutex{}

func (p *parser) Parse(line string) string {
	// this is where the parsing routine would go
	fmt.Printf("Parsing the string: %s from %s\n", line, p.Columns.Name)
	return "parsed"
}

func getParser(lineType string) ParserBuilder {
    if builderMap[lineType] == nil {
        builderMap[lineType] = &parserBuilder{}
    }
	return builderMap[lineType]
}

func ParseIt(f func(ColumnDef, string), lineType string, line string) string{
    // getParser is a singleton pattern - these are pointers to the specific parser objects
    lock.Lock()
    parser := getParser(lineType).Columns(parserMap[lineType]).Build().(*parser)
    lock.Unlock()
    f(parser.Columns, line)
    return "parsed"
}

