package structColumnDefs

import (
    "fmt"
    "sync"
    "reflect"
    "strconv"
    "strings"
)

type ColumnDef struct {
	Name string
    doc interface{}
}

// column definitions - add to these all you like
// Each column definition has a name, and a list of fields.
// Each column definition is a struct that implements the Parser interface
// Each new ColumnDef needs to be added to the parserMap
//VERSION MODEL DESC FCST_LEAD FCST_VALID_BEG  FCST_VALID_END  OBS_LEAD OBS_VALID_BEG   OBS_VALID_END   FCST_VAR  FCST_UNITS FCST_LEV OBS_VAR   OBS_UNITS OBS_LEV OBTYPE VX_MASK INTERP_MTHD INTERP_PNTS FCST_THRESH OBS_THRESH COV_THRESH ALPHA LINE_TYPE

type statHeaderL1 struct {
    ID string `json:"ID"`
    Version string `json:"VERSION"`
    Model string `json:"MODEL"`
    Desc string `json:"DESC"`
    FcstLev string `json:"FCST_LEV"`
    FcstValidBeg int `json:"FCST_VALID_BEG"`
    FcstValidEnd int `json:"FCST_VALID_END"`
    ObsLev string `json:"OBS_LEV"`
    ObsValidBeg int `json:"OBS_VALID_BEG"`
    ObsValidEnd int `json:"OBS_VALID_END"`
    FcstVar string `json:"FCST_VAR"`
    FcstUnits string `json:"FCST_UNITS"`
    FcstLvl string `json:"FCST_LVL"`

    InterpMthd string `json:"INTERP_MTHD"`
    InterpPnts int `json:"INTERP_PNTS"`
    LineType string `json:"LINE_TYPE"`
    ObsUnits string `json:"OBS_UNITS"`
    ObsVar string `json:"OBS_VAR"`
    Obtype string `json:"OBTYPE"`
    VxMask string `json:"VX_MASK"`
}

type SAL1L2doc struct {
    HeaderFields statHeaderL1
        Total int `json:"TOTAL"`
        DataFields map[string]struct {
            Fabar float64 `json:"FABAR"`
            Oabar float64 `json:"OABAR"`
            Foabar float64 `json:"FOABAR"`
            Ffabar float64 `json:"FFABAR"`
            Ooabar float64 `json:"OOABAR"`
            Mae float64 `json:"MAE"`
        }
}

type VL1L2Doc struct {
    HeaderFields statHeaderL1
    DataFields map[string]struct {
        Fabar float64 `json:"FABAR"`
        Ffabar float64 `json:"FFABAR"`
        Foabar float64 `json:"FOABAR"`
        Mae float64 `json:"MAE"`
        Oabar float64 `json:"OABAR"`
        Ooabar float64 `json:"OOABAR"`
    }
}

type VAL1L2Doc struct {
    HeaderFields statHeaderL1
    DataFields map[string]struct {
        Total int `json:"TOTAL"`
        UFABAR float64 `json:"UFABAR"`
        VFABAR float64 `json:"VFABAR"`
        UOABAR float64 `json:"UOABAR"`
        VOABAR float64 `json:"VOABAR"`
        UVFOABAR float64 `json:"UVFOABAR"`
        UVFFBAR float64 `json:"UVFFBAR"`
        UVOOABAR float64 `json:"UVOOABAR"`
        FA_SPEED_BAR float64 `json:"FA_SPEED_BAR"`
        OA_SPEED_BAR float64 `json:"OA_SPEED_BAR"`
    }
}

var parserMap = map[string] ColumnDef{
    "VL1L2": ColumnDef {
        Name: "VL1L2",
        doc: VL1L2Doc{},
    },
    "VAL1L2": ColumnDef {
        Name: "VAL1L2",
        doc: VAL1L2Doc{},
    },
}


// DON'T CHANGE BELOW THIS LINE

var builderMap = map[string] ParserBuilder{}

type Parser interface {
	Parse(string) (string, error)
}

type ParserBuilder interface {
	Columns(ColumnDef) ParserBuilder
	Build(string) Parser
}

type parserBuilder struct {
	columnDef ColumnDef
}

type parser struct {
	Columns ColumnDef
}

// aStruct is either a header or a data DataFields struct
func (pb *parser) fillStruct(aStruct interface{}, lineData []string) {
    dataVal := reflect.ValueOf(aStruct)
    for i := 0; i < dataVal.NumField(); i++ {
        switch dataVal.Field(i).Kind() {
            case reflect.Int: {
                d, _ := strconv.Atoi(lineData[i])
                dataVal.Field(i).SetInt(int64(d))
            }
            case reflect.Float64: {
                d, _ := strconv.ParseFloat(lineData[i], 64)
                dataVal.Field(i).SetFloat(d)
            }
            case reflect.String: {
                dataVal.Field(i).SetString(lineData[i])
            }
            case reflect.Bool: {
                d, _ := strconv.ParseBool(lineData[i])
                dataVal.Field(i).SetBool(d)
            }
            default: {
                fmt.Printf("Unknown type %s\n", dataVal.Field(i).Kind())
            }
        }
    }
}

func (pb *parserBuilder) Columns(columns ColumnDef) ParserBuilder {
	pb.columnDef = columns
	return pb
}

func (pb *parserBuilder) Build(line string) Parser {
	myPb := &parser{
		Columns: pb.columnDef,
	}
    // Build only happens to brand new builders so fill in the header fields
    // of the column definition struct
    headerFieldStruct := reflect.ValueOf(myPb.Columns.doc).FieldByName("HeaderFields")
    lineData := strings.Fields(line)
    myPb.fillStruct(headerFieldStruct, lineData)
    return myPb
}

var lock = &sync.Mutex{}

func (p *parser) Parse(line string) (string, error) {
	// this is where the parsing routine would go
	fmt.Printf("Parsing the string: %s from %s\n", line, p.Columns.Name)
    dataFieldStruct := reflect.ValueOf(p.Columns.doc).FieldByName("DataFields")
    lineData := strings.Fields(line)
    p.fillStruct(dataFieldStruct, lineData)
	return "parsed", nil
}

func getParser(lineType string) ParserBuilder {
    if builderMap[lineType] == nil {
        builderMap[lineType] = &parserBuilder{}
    }
	return builderMap[lineType]
}

func ParseIt(f func(ColumnDef, string) (interface{}, error), lineType string, line string) (interface{}, error){
    // singleton pattern - these are pointers to the parser objects.
    // getParsers() returns a pointer to a particular parser implementation mapped to the lineType.
    // the parser implementation is a struct that implements the Parser interface by having a Parse method
    // and a Columns field that is a struct that implements the ColumnDef interface.

    lock.Lock()
    parser := getParser(lineType).Columns(parserMap[lineType]).Build(line).(*parser)
    lock.Unlock()
    ret, _err := parser.Parse(line)
    return ret, _err
}

