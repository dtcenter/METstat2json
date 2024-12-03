package structColumnDefs

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

type LField = struct {
	name string
	fieldType string
}

type ColumnDef struct {
	Name string
	Columns []LField
}


// column definitions - add to these all you like
// Each column definition has a name, and a list of fields.
// Each column definition is a struct that implements the Parser interface
// Each new ColumnDef needs to be added to the parserMap

var VL1L2 = ColumnDef {
	Name: "VL1L2",
	Columns: []LField {
		{name: "TOTAL", fieldType: "int"},
		{name: "UFBAR", fieldType: "float"},
		{name: "VFBAR", fieldType: "float"},
		{name: "UOBAR", fieldType: "float"},
		{name: "VOBAR", fieldType: "float"},
		{name: "UVFOBAR", fieldType: "float"},
		{name: "UVFFBAR", fieldType: "float"},
		{name: "UVOOBAR", fieldType: "float"},
		{name: "F_SPEED_BAR", fieldType: "float"},
		{name: "O_SPEED_BAR", fieldType: "float"},
	},
}

var VAL1L2 = ColumnDef {
	Name: "VAL1L2",
	Columns: []LField {
		{name: "TOTAL", fieldType: "int"},
		{name: "UFABAR", fieldType: "float"},
		{name: "VFABAR", fieldType: "float"},
		{name: "UOABAR", fieldType: "float"},
		{name: "VOABAR", fieldType: "float"},
		{name: "UVFOABAR", fieldType: "float"},
		{name: "UVFFBAR", fieldType: "float"},
		{name: "UVOOABAR", fieldType: "float"},
		{name: "FA_SPEED_BAR", fieldType: "float"},
		{name: "OA_SPEED_BAR", fieldType: "float"},
	},
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
    // singleton pattern - these are pointers to the parser objects
    lock.Lock()
    parser := getParser(lineType).Columns(parserMap[lineType]).Build().(*parser)
    lock.Unlock()
    f(parser.Columns, line)
    return "parsed"
}

