package structColumnDefs

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
)

type ColumnDef struct {
	Name string
	doc  interface{}
}

// column definitions - add to these all you like
// Each column definition has a name, and a list of fields.
// Each column definition is a struct that implements the Parser interface
// Each new ColumnDef needs to be added to the parserMap
// VERSION MODEL DESC FCST_LEAD FCST_VALID_BEG  FCST_VALID_END  OBS_LEAD OBS_VALID_BEG   OBS_VALID_END   FCST_VAR  FCST_UNITS FCST_LEV OBS_VAR   OBS_UNITS OBS_LEV OBTYPE VX_MASK INTERP_MTHD INTERP_PNTS FCST_THRESH OBS_THRESH COV_THRESH ALPHA LINE_TYPE
type statHeader struct {
	Version        string `json:"VERSION"`
	Model          string `json:"MODEL"`
	Desc           string `json:"DESC"`
	Fcst_lead      string `json:"FCST_LEAD"`
	Fcst_valid_beg string `json:"FCST_VALID_BEG"`
	Fcst_valid_end string `json:"FCST_VALID_END"`
	Obs_lead       string `json:"OBS_LEAD"`
	Obs_valid_beg  string `json:"OBS_VALID_BEG"`
	Obs_valid_end  string `json:"OBS_VALID_END"`
	Fcst_var       string `json:"FCST_VAR"`
	Fcst_units     string `json:"FCST_UNITS"`
	Fcst_lev       string `json:"FCST_LEV"`
	Obs_var        string `json:"OBS_VAR"`
	Obs_units      string `json:"OBS_UNITS"`
	Obs_lev        string `json:"OBS_LEV"`
	Obtype         string `json:"OBTYPE"`
	Vx_mask        string `json:"VX_MASK"`
	Interp_mthd    string `json:"INTERP_MTHD"`
	Interp_pnts    string `json:"INTERP_PNTS"`
	Fcst_thresh    string `json:"FCST_THRESH"`
	Obs_thresh     string `json:"OBS_THRESH"`
	Cov_thresh     string `json:"COV_THRESH"`
	Alpha          string `json:"ALPHA"`
	Line_type      string `json:"LINE_TYPE"`
}

type vxMetadata struct {
	ID string `json:"ID"`
	// etc...
}

type VL1L2Doc struct {
	vxMetadata
	statHeader
	Data map[string]struct {
		Fabar  float64 `json:"FABAR"`
		Ffabar float64 `json:"FFABAR"`
		Foabar float64 `json:"FOABAR"`
		Mae    float64 `json:"MAE"`
		Oabar  float64 `json:"OABAR"`
		Ooabar float64 `json:"OOABAR"`
	}
}

type VAL1L2Doc struct {
	vxMetadata
	statHeader
	Data map[string]struct {
		Total        int     `json:"TOTAL"`
		UFABAR       float64 `json:"UFABAR"`
		VFABAR       float64 `json:"VFABAR"`
		UOABAR       float64 `json:"UOABAR"`
		VOABAR       float64 `json:"VOABAR"`
		UVFOABAR     float64 `json:"UVFOABAR"`
		UVFFBAR      float64 `json:"UVFFBAR"`
		UVOOABAR     float64 `json:"UVOOABAR"`
		FA_SPEED_BAR float64 `json:"FA_SPEED_BAR"`
		OA_SPEED_BAR float64 `json:"OA_SPEED_BAR"`
	}
}

var parserMap = map[string]ColumnDef{
	"VL1L2": {
		Name: "VL1L2",
		doc:  VL1L2Doc{},
	},
	"VAL1L2": {
		Name: "VAL1L2",
		doc:  VAL1L2Doc{},
	},
}

// DON'T CHANGE BELOW THIS LINE

var builderMap = map[string]ParserBuilder{}

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
		case reflect.Int:
			{
				d, _ := strconv.Atoi(lineData[i])
				dataVal.Field(i).SetInt(int64(d))
			}
		case reflect.Float64:
			{
				d, _ := strconv.ParseFloat(lineData[i], 64)
				dataVal.Field(i).SetFloat(d)
			}
		case reflect.String:
			{
				dataVal.Field(i).SetString(lineData[i])
			}
		case reflect.Bool:
			{
				d, _ := strconv.ParseBool(lineData[i])
				dataVal.Field(i).SetBool(d)
			}
		default:
			{
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

func ParseIt(f func(ColumnDef, string) (interface{}, error), lineType string, line string) (interface{}, error) {
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
