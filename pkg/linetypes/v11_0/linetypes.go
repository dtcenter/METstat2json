/*
The following data types were not found in the MET user guide files or the MET source code files.
For simplicity, the values of these data types will be treated as strings in the generated code.

	Undefined data types: [S12 S13 S21 S23 S31 S32]

To resolve this, consult the github.com/dtcenter/MET repo to determine if there is a more appropriate type,
and, if there is, add an override to the overRideDefinedMetDataTypes function in generator/generator.go.
*/
package v11_0

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

/*
THIS CODE IS AUTOMATICALLY GENERATED - DO NOT EDIT THIS CODE
To modify this code - modify the generator.go file and run the generator.go program
cd  <repo_root>
go run generator -version=v12.0 > pkg/linetypes/v12_0/linetypes.go
*/

func GetLeadFromInitValid(data []string, dataFieldIndex int) string {
	initTime, _ := time.Parse("20060102_150405", data[dataFieldIndex-1])
	validTime, _ := time.Parse("20060102_150405", data[dataFieldIndex+1])
	lead := validTime.Sub(initTime)
	return strconv.FormatInt(int64(lead.Hours()), 10)
}

func SetValueForField(doc *map[string]interface{}, fileType string, term string, i int, dataLen int, fields []string, fieldIndex int, dataType string) {
	if term == "INIT" && fileType != "TCST" {
		// do not assign any value to the INIT field for TCST files
		// The INIT field needs to be moved from the header to the data section.
		return
	}
	if term == "LEAD" && fileType == "TCST" {
		// This is a special case for TCST files.
		// if the TERM is LEAD and the lead is NA then we
		// have to get the lead from the init and valid fields.
		// INIT is the prior field and VALID is the next field from LEAD.
		// the init and valid fields are in the format YYYYMMDD_HHMMSS
		// the lead is the difference between the valid and the init
		// in hours
		(*doc)["LEAD"], _ = strconv.Atoi(GetLeadFromInitValid(fields, fieldIndex))
		return
	}
	if i <= dataLen && fields[fieldIndex] != "" && fields[fieldIndex] != "NA" {
		switch dataType {
		case "int":
			(*doc)[term], _ = strconv.Atoi(fields[fieldIndex])
			return
		case "float64":
			(*doc)[term], _ = strconv.ParseFloat(fields[fieldIndex], 64)
			return
		case "string":
			(*doc)[term] = fields[fieldIndex]
			return
		default:
			(*doc)[term] = fields[fieldIndex]
			return
		}
	}
}

// Header struct definitions
type MODE_CTS_header struct {
	VERSION    string  `json:"VERSION"`
	MODEL      string  `json:"MODEL"`
	N_VALID    int     `json:"N_VALID"`
	GRID_RES   float64 `json:"GRID_RES"`
	DESC       string  `json:"DESC"`
	FCST_VALID string  `json:"FCST_VALID"`
	FCST_ACCUM string  `json:"FCST_ACCUM"`
	OBS_LEAD   int     `json:"OBS_LEAD"`
	OBS_VALID  string  `json:"OBS_VALID"`
	OBS_ACCUM  string  `json:"OBS_ACCUM"`
	FCST_RAD   int     `json:"FCST_RAD"`
	FCST_THR   string  `json:"FCST_THR"`
	OBS_RAD    int     `json:"OBS_RAD"`
	OBS_THR    string  `json:"OBS_THR"`
	FCST_VAR   string  `json:"FCST_VAR"`
	FCST_UNITS string  `json:"FCST_UNITS"`
	FCST_LEV   string  `json:"FCST_LEV"`
	OBS_VAR    string  `json:"OBS_VAR"`
	OBS_UNITS  string  `json:"OBS_UNITS"`
	OBS_LEV    string  `json:"OBS_LEV"`
	OBTYPE     string  `json:"OBTYPE"`
	LINE_TYPE  string  `json:"LINE_TYPE"`
}

type MODE_OBJ_header struct {
	VERSION    string  `json:"VERSION"`
	MODEL      string  `json:"MODEL"`
	N_VALID    int     `json:"N_VALID"`
	GRID_RES   float64 `json:"GRID_RES"`
	DESC       string  `json:"DESC"`
	FCST_VALID string  `json:"FCST_VALID"`
	FCST_ACCUM string  `json:"FCST_ACCUM"`
	OBS_LEAD   int     `json:"OBS_LEAD"`
	OBS_VALID  string  `json:"OBS_VALID"`
	OBS_ACCUM  string  `json:"OBS_ACCUM"`
	FCST_RAD   int     `json:"FCST_RAD"`
	FCST_THR   string  `json:"FCST_THR"`
	OBS_RAD    int     `json:"OBS_RAD"`
	OBS_THR    string  `json:"OBS_THR"`
	FCST_VAR   string  `json:"FCST_VAR"`
	FCST_UNITS string  `json:"FCST_UNITS"`
	FCST_LEV   string  `json:"FCST_LEV"`
	OBS_VAR    string  `json:"OBS_VAR"`
	OBS_UNITS  string  `json:"OBS_UNITS"`
	OBS_LEV    string  `json:"OBS_LEV"`
	OBTYPE     string  `json:"OBTYPE"`
	LINE_TYPE  string  `json:"LINE_TYPE"`
}

type STAT_CNT_header struct {
	VERSION        string  `json:"VERSION"`
	MODEL          string  `json:"MODEL"`
	DESC           string  `json:"DESC"`
	FCST_VALID_BEG int     `json:"FCST_VALID_BEG"`
	FCST_VALID_END int     `json:"FCST_VALID_END"`
	OBS_LEAD       int     `json:"OBS_LEAD"`
	OBS_VALID_BEG  int     `json:"OBS_VALID_BEG"`
	OBS_VALID_END  int     `json:"OBS_VALID_END"`
	FCST_VAR       string  `json:"FCST_VAR"`
	FCST_UNITS     string  `json:"FCST_UNITS"`
	FCST_LEV       string  `json:"FCST_LEV"`
	OBS_VAR        string  `json:"OBS_VAR"`
	OBS_UNITS      string  `json:"OBS_UNITS"`
	OBS_LEV        string  `json:"OBS_LEV"`
	OBTYPE         string  `json:"OBTYPE"`
	VX_MASK        string  `json:"VX_MASK"`
	INTERP_MTHD    string  `json:"INTERP_MTHD"`
	INTERP_PNTS    int     `json:"INTERP_PNTS"`
	FCST_THRESH    string  `json:"FCST_THRESH"`
	OBS_THRESH     string  `json:"OBS_THRESH"`
	COV_THRESH     string  `json:"COV_THRESH"`
	ALPHA          float64 `json:"ALPHA"`
	LINE_TYPE      string  `json:"LINE_TYPE"`
}

type STAT_CTC_header struct {
	VERSION        string  `json:"VERSION"`
	MODEL          string  `json:"MODEL"`
	DESC           string  `json:"DESC"`
	FCST_VALID_BEG int     `json:"FCST_VALID_BEG"`
	FCST_VALID_END int     `json:"FCST_VALID_END"`
	OBS_LEAD       int     `json:"OBS_LEAD"`
	OBS_VALID_BEG  int     `json:"OBS_VALID_BEG"`
	OBS_VALID_END  int     `json:"OBS_VALID_END"`
	FCST_VAR       string  `json:"FCST_VAR"`
	FCST_UNITS     string  `json:"FCST_UNITS"`
	FCST_LEV       string  `json:"FCST_LEV"`
	OBS_VAR        string  `json:"OBS_VAR"`
	OBS_UNITS      string  `json:"OBS_UNITS"`
	OBS_LEV        string  `json:"OBS_LEV"`
	OBTYPE         string  `json:"OBTYPE"`
	VX_MASK        string  `json:"VX_MASK"`
	INTERP_MTHD    string  `json:"INTERP_MTHD"`
	INTERP_PNTS    int     `json:"INTERP_PNTS"`
	FCST_THRESH    string  `json:"FCST_THRESH"`
	OBS_THRESH     string  `json:"OBS_THRESH"`
	COV_THRESH     string  `json:"COV_THRESH"`
	ALPHA          float64 `json:"ALPHA"`
	LINE_TYPE      string  `json:"LINE_TYPE"`
}

type STAT_CTS_header struct {
	VERSION        string  `json:"VERSION"`
	MODEL          string  `json:"MODEL"`
	DESC           string  `json:"DESC"`
	FCST_VALID_BEG int     `json:"FCST_VALID_BEG"`
	FCST_VALID_END int     `json:"FCST_VALID_END"`
	OBS_LEAD       int     `json:"OBS_LEAD"`
	OBS_VALID_BEG  int     `json:"OBS_VALID_BEG"`
	OBS_VALID_END  int     `json:"OBS_VALID_END"`
	FCST_VAR       string  `json:"FCST_VAR"`
	FCST_UNITS     string  `json:"FCST_UNITS"`
	FCST_LEV       string  `json:"FCST_LEV"`
	OBS_VAR        string  `json:"OBS_VAR"`
	OBS_UNITS      string  `json:"OBS_UNITS"`
	OBS_LEV        string  `json:"OBS_LEV"`
	OBTYPE         string  `json:"OBTYPE"`
	VX_MASK        string  `json:"VX_MASK"`
	INTERP_MTHD    string  `json:"INTERP_MTHD"`
	INTERP_PNTS    int     `json:"INTERP_PNTS"`
	FCST_THRESH    string  `json:"FCST_THRESH"`
	OBS_THRESH     string  `json:"OBS_THRESH"`
	COV_THRESH     string  `json:"COV_THRESH"`
	ALPHA          float64 `json:"ALPHA"`
	LINE_TYPE      string  `json:"LINE_TYPE"`
}

type STAT_DMAP_header struct {
	VERSION        string  `json:"VERSION"`
	MODEL          string  `json:"MODEL"`
	DESC           string  `json:"DESC"`
	FCST_VALID_BEG int     `json:"FCST_VALID_BEG"`
	FCST_VALID_END int     `json:"FCST_VALID_END"`
	OBS_LEAD       int     `json:"OBS_LEAD"`
	OBS_VALID_BEG  int     `json:"OBS_VALID_BEG"`
	OBS_VALID_END  int     `json:"OBS_VALID_END"`
	FCST_VAR       string  `json:"FCST_VAR"`
	FCST_UNITS     string  `json:"FCST_UNITS"`
	FCST_LEV       string  `json:"FCST_LEV"`
	OBS_VAR        string  `json:"OBS_VAR"`
	OBS_UNITS      string  `json:"OBS_UNITS"`
	OBS_LEV        string  `json:"OBS_LEV"`
	OBTYPE         string  `json:"OBTYPE"`
	VX_MASK        string  `json:"VX_MASK"`
	INTERP_MTHD    string  `json:"INTERP_MTHD"`
	INTERP_PNTS    int     `json:"INTERP_PNTS"`
	FCST_THRESH    string  `json:"FCST_THRESH"`
	OBS_THRESH     string  `json:"OBS_THRESH"`
	COV_THRESH     string  `json:"COV_THRESH"`
	ALPHA          float64 `json:"ALPHA"`
	LINE_TYPE      string  `json:"LINE_TYPE"`
}

type STAT_ECLV_header struct {
	VERSION        string  `json:"VERSION"`
	MODEL          string  `json:"MODEL"`
	DESC           string  `json:"DESC"`
	FCST_VALID_BEG int     `json:"FCST_VALID_BEG"`
	FCST_VALID_END int     `json:"FCST_VALID_END"`
	OBS_LEAD       int     `json:"OBS_LEAD"`
	OBS_VALID_BEG  int     `json:"OBS_VALID_BEG"`
	OBS_VALID_END  int     `json:"OBS_VALID_END"`
	FCST_VAR       string  `json:"FCST_VAR"`
	FCST_UNITS     string  `json:"FCST_UNITS"`
	FCST_LEV       string  `json:"FCST_LEV"`
	OBS_VAR        string  `json:"OBS_VAR"`
	OBS_UNITS      string  `json:"OBS_UNITS"`
	OBS_LEV        string  `json:"OBS_LEV"`
	OBTYPE         string  `json:"OBTYPE"`
	VX_MASK        string  `json:"VX_MASK"`
	INTERP_MTHD    string  `json:"INTERP_MTHD"`
	INTERP_PNTS    int     `json:"INTERP_PNTS"`
	FCST_THRESH    string  `json:"FCST_THRESH"`
	OBS_THRESH     string  `json:"OBS_THRESH"`
	COV_THRESH     string  `json:"COV_THRESH"`
	ALPHA          float64 `json:"ALPHA"`
	LINE_TYPE      string  `json:"LINE_TYPE"`
}

type STAT_ECNT_header struct {
	VERSION        string  `json:"VERSION"`
	MODEL          string  `json:"MODEL"`
	DESC           string  `json:"DESC"`
	FCST_VALID_BEG int     `json:"FCST_VALID_BEG"`
	FCST_VALID_END int     `json:"FCST_VALID_END"`
	OBS_LEAD       int     `json:"OBS_LEAD"`
	OBS_VALID_BEG  int     `json:"OBS_VALID_BEG"`
	OBS_VALID_END  int     `json:"OBS_VALID_END"`
	FCST_VAR       string  `json:"FCST_VAR"`
	FCST_UNITS     string  `json:"FCST_UNITS"`
	FCST_LEV       string  `json:"FCST_LEV"`
	OBS_VAR        string  `json:"OBS_VAR"`
	OBS_UNITS      string  `json:"OBS_UNITS"`
	OBS_LEV        string  `json:"OBS_LEV"`
	OBTYPE         string  `json:"OBTYPE"`
	VX_MASK        string  `json:"VX_MASK"`
	INTERP_MTHD    string  `json:"INTERP_MTHD"`
	INTERP_PNTS    int     `json:"INTERP_PNTS"`
	FCST_THRESH    string  `json:"FCST_THRESH"`
	OBS_THRESH     string  `json:"OBS_THRESH"`
	COV_THRESH     string  `json:"COV_THRESH"`
	ALPHA          float64 `json:"ALPHA"`
	LINE_TYPE      string  `json:"LINE_TYPE"`
}

type STAT_FHO_header struct {
	VERSION        string  `json:"VERSION"`
	MODEL          string  `json:"MODEL"`
	DESC           string  `json:"DESC"`
	FCST_VALID_BEG int     `json:"FCST_VALID_BEG"`
	FCST_VALID_END int     `json:"FCST_VALID_END"`
	OBS_LEAD       int     `json:"OBS_LEAD"`
	OBS_VALID_BEG  int     `json:"OBS_VALID_BEG"`
	OBS_VALID_END  int     `json:"OBS_VALID_END"`
	FCST_VAR       string  `json:"FCST_VAR"`
	FCST_UNITS     string  `json:"FCST_UNITS"`
	FCST_LEV       string  `json:"FCST_LEV"`
	OBS_VAR        string  `json:"OBS_VAR"`
	OBS_UNITS      string  `json:"OBS_UNITS"`
	OBS_LEV        string  `json:"OBS_LEV"`
	OBTYPE         string  `json:"OBTYPE"`
	VX_MASK        string  `json:"VX_MASK"`
	INTERP_MTHD    string  `json:"INTERP_MTHD"`
	INTERP_PNTS    int     `json:"INTERP_PNTS"`
	FCST_THRESH    string  `json:"FCST_THRESH"`
	OBS_THRESH     string  `json:"OBS_THRESH"`
	COV_THRESH     string  `json:"COV_THRESH"`
	ALPHA          float64 `json:"ALPHA"`
	LINE_TYPE      string  `json:"LINE_TYPE"`
}

type STAT_GENMPR_header struct {
	VERSION        string  `json:"VERSION"`
	MODEL          string  `json:"MODEL"`
	DESC           string  `json:"DESC"`
	FCST_VALID_BEG int     `json:"FCST_VALID_BEG"`
	FCST_VALID_END int     `json:"FCST_VALID_END"`
	OBS_LEAD       int     `json:"OBS_LEAD"`
	OBS_VALID_BEG  int     `json:"OBS_VALID_BEG"`
	OBS_VALID_END  int     `json:"OBS_VALID_END"`
	FCST_VAR       string  `json:"FCST_VAR"`
	FCST_UNITS     string  `json:"FCST_UNITS"`
	FCST_LEV       string  `json:"FCST_LEV"`
	OBS_VAR        string  `json:"OBS_VAR"`
	OBS_UNITS      string  `json:"OBS_UNITS"`
	OBS_LEV        string  `json:"OBS_LEV"`
	OBTYPE         string  `json:"OBTYPE"`
	VX_MASK        string  `json:"VX_MASK"`
	INTERP_MTHD    string  `json:"INTERP_MTHD"`
	INTERP_PNTS    int     `json:"INTERP_PNTS"`
	FCST_THRESH    string  `json:"FCST_THRESH"`
	OBS_THRESH     string  `json:"OBS_THRESH"`
	COV_THRESH     string  `json:"COV_THRESH"`
	ALPHA          float64 `json:"ALPHA"`
	LINE_TYPE      string  `json:"LINE_TYPE"`
}

type STAT_GRAD_header struct {
	VERSION        string  `json:"VERSION"`
	MODEL          string  `json:"MODEL"`
	DESC           string  `json:"DESC"`
	FCST_VALID_BEG int     `json:"FCST_VALID_BEG"`
	FCST_VALID_END int     `json:"FCST_VALID_END"`
	OBS_LEAD       int     `json:"OBS_LEAD"`
	OBS_VALID_BEG  int     `json:"OBS_VALID_BEG"`
	OBS_VALID_END  int     `json:"OBS_VALID_END"`
	FCST_VAR       string  `json:"FCST_VAR"`
	FCST_UNITS     string  `json:"FCST_UNITS"`
	FCST_LEV       string  `json:"FCST_LEV"`
	OBS_VAR        string  `json:"OBS_VAR"`
	OBS_UNITS      string  `json:"OBS_UNITS"`
	OBS_LEV        string  `json:"OBS_LEV"`
	OBTYPE         string  `json:"OBTYPE"`
	VX_MASK        string  `json:"VX_MASK"`
	INTERP_MTHD    string  `json:"INTERP_MTHD"`
	INTERP_PNTS    int     `json:"INTERP_PNTS"`
	FCST_THRESH    string  `json:"FCST_THRESH"`
	OBS_THRESH     string  `json:"OBS_THRESH"`
	COV_THRESH     string  `json:"COV_THRESH"`
	ALPHA          float64 `json:"ALPHA"`
	LINE_TYPE      string  `json:"LINE_TYPE"`
}

type STAT_ISC_header struct {
	VERSION        string  `json:"VERSION"`
	MODEL          string  `json:"MODEL"`
	DESC           string  `json:"DESC"`
	FCST_VALID_BEG int     `json:"FCST_VALID_BEG"`
	FCST_VALID_END int     `json:"FCST_VALID_END"`
	OBS_LEAD       int     `json:"OBS_LEAD"`
	OBS_VALID_BEG  int     `json:"OBS_VALID_BEG"`
	OBS_VALID_END  int     `json:"OBS_VALID_END"`
	FCST_VAR       string  `json:"FCST_VAR"`
	FCST_UNITS     string  `json:"FCST_UNITS"`
	FCST_LEV       string  `json:"FCST_LEV"`
	OBS_VAR        string  `json:"OBS_VAR"`
	OBS_UNITS      string  `json:"OBS_UNITS"`
	OBS_LEV        string  `json:"OBS_LEV"`
	OBTYPE         string  `json:"OBTYPE"`
	VX_MASK        string  `json:"VX_MASK"`
	INTERP_MTHD    string  `json:"INTERP_MTHD"`
	INTERP_PNTS    int     `json:"INTERP_PNTS"`
	FCST_THRESH    string  `json:"FCST_THRESH"`
	OBS_THRESH     string  `json:"OBS_THRESH"`
	COV_THRESH     string  `json:"COV_THRESH"`
	ALPHA          float64 `json:"ALPHA"`
	LINE_TYPE      string  `json:"LINE_TYPE"`
}

type STAT_MCTC_header struct {
	VERSION        string  `json:"VERSION"`
	MODEL          string  `json:"MODEL"`
	DESC           string  `json:"DESC"`
	FCST_VALID_BEG int     `json:"FCST_VALID_BEG"`
	FCST_VALID_END int     `json:"FCST_VALID_END"`
	OBS_LEAD       int     `json:"OBS_LEAD"`
	OBS_VALID_BEG  int     `json:"OBS_VALID_BEG"`
	OBS_VALID_END  int     `json:"OBS_VALID_END"`
	FCST_VAR       string  `json:"FCST_VAR"`
	FCST_UNITS     string  `json:"FCST_UNITS"`
	FCST_LEV       string  `json:"FCST_LEV"`
	OBS_VAR        string  `json:"OBS_VAR"`
	OBS_UNITS      string  `json:"OBS_UNITS"`
	OBS_LEV        string  `json:"OBS_LEV"`
	OBTYPE         string  `json:"OBTYPE"`
	VX_MASK        string  `json:"VX_MASK"`
	INTERP_MTHD    string  `json:"INTERP_MTHD"`
	INTERP_PNTS    int     `json:"INTERP_PNTS"`
	FCST_THRESH    string  `json:"FCST_THRESH"`
	OBS_THRESH     string  `json:"OBS_THRESH"`
	COV_THRESH     string  `json:"COV_THRESH"`
	ALPHA          float64 `json:"ALPHA"`
	LINE_TYPE      string  `json:"LINE_TYPE"`
}

type STAT_MCTS_header struct {
	VERSION        string  `json:"VERSION"`
	MODEL          string  `json:"MODEL"`
	DESC           string  `json:"DESC"`
	FCST_VALID_BEG int     `json:"FCST_VALID_BEG"`
	FCST_VALID_END int     `json:"FCST_VALID_END"`
	OBS_LEAD       int     `json:"OBS_LEAD"`
	OBS_VALID_BEG  int     `json:"OBS_VALID_BEG"`
	OBS_VALID_END  int     `json:"OBS_VALID_END"`
	FCST_VAR       string  `json:"FCST_VAR"`
	FCST_UNITS     string  `json:"FCST_UNITS"`
	FCST_LEV       string  `json:"FCST_LEV"`
	OBS_VAR        string  `json:"OBS_VAR"`
	OBS_UNITS      string  `json:"OBS_UNITS"`
	OBS_LEV        string  `json:"OBS_LEV"`
	OBTYPE         string  `json:"OBTYPE"`
	VX_MASK        string  `json:"VX_MASK"`
	INTERP_MTHD    string  `json:"INTERP_MTHD"`
	INTERP_PNTS    int     `json:"INTERP_PNTS"`
	FCST_THRESH    string  `json:"FCST_THRESH"`
	OBS_THRESH     string  `json:"OBS_THRESH"`
	COV_THRESH     string  `json:"COV_THRESH"`
	ALPHA          float64 `json:"ALPHA"`
	LINE_TYPE      string  `json:"LINE_TYPE"`
}

type STAT_MPR_header struct {
	VERSION        string  `json:"VERSION"`
	MODEL          string  `json:"MODEL"`
	DESC           string  `json:"DESC"`
	FCST_VALID_BEG int     `json:"FCST_VALID_BEG"`
	FCST_VALID_END int     `json:"FCST_VALID_END"`
	OBS_LEAD       int     `json:"OBS_LEAD"`
	OBS_VALID_BEG  int     `json:"OBS_VALID_BEG"`
	OBS_VALID_END  int     `json:"OBS_VALID_END"`
	FCST_VAR       string  `json:"FCST_VAR"`
	FCST_UNITS     string  `json:"FCST_UNITS"`
	FCST_LEV       string  `json:"FCST_LEV"`
	OBS_VAR        string  `json:"OBS_VAR"`
	OBS_UNITS      string  `json:"OBS_UNITS"`
	OBS_LEV        string  `json:"OBS_LEV"`
	OBTYPE         string  `json:"OBTYPE"`
	VX_MASK        string  `json:"VX_MASK"`
	INTERP_MTHD    string  `json:"INTERP_MTHD"`
	INTERP_PNTS    int     `json:"INTERP_PNTS"`
	FCST_THRESH    string  `json:"FCST_THRESH"`
	OBS_THRESH     string  `json:"OBS_THRESH"`
	COV_THRESH     string  `json:"COV_THRESH"`
	ALPHA          float64 `json:"ALPHA"`
	LINE_TYPE      string  `json:"LINE_TYPE"`
}

type STAT_NBRCNT_header struct {
	VERSION        string  `json:"VERSION"`
	MODEL          string  `json:"MODEL"`
	DESC           string  `json:"DESC"`
	FCST_VALID_BEG int     `json:"FCST_VALID_BEG"`
	FCST_VALID_END int     `json:"FCST_VALID_END"`
	OBS_LEAD       int     `json:"OBS_LEAD"`
	OBS_VALID_BEG  int     `json:"OBS_VALID_BEG"`
	OBS_VALID_END  int     `json:"OBS_VALID_END"`
	FCST_VAR       string  `json:"FCST_VAR"`
	FCST_UNITS     string  `json:"FCST_UNITS"`
	FCST_LEV       string  `json:"FCST_LEV"`
	OBS_VAR        string  `json:"OBS_VAR"`
	OBS_UNITS      string  `json:"OBS_UNITS"`
	OBS_LEV        string  `json:"OBS_LEV"`
	OBTYPE         string  `json:"OBTYPE"`
	VX_MASK        string  `json:"VX_MASK"`
	INTERP_MTHD    string  `json:"INTERP_MTHD"`
	INTERP_PNTS    int     `json:"INTERP_PNTS"`
	FCST_THRESH    string  `json:"FCST_THRESH"`
	OBS_THRESH     string  `json:"OBS_THRESH"`
	COV_THRESH     string  `json:"COV_THRESH"`
	ALPHA          float64 `json:"ALPHA"`
	LINE_TYPE      string  `json:"LINE_TYPE"`
}

type STAT_NBRCTC_header struct {
	VERSION        string  `json:"VERSION"`
	MODEL          string  `json:"MODEL"`
	DESC           string  `json:"DESC"`
	FCST_VALID_BEG int     `json:"FCST_VALID_BEG"`
	FCST_VALID_END int     `json:"FCST_VALID_END"`
	OBS_LEAD       int     `json:"OBS_LEAD"`
	OBS_VALID_BEG  int     `json:"OBS_VALID_BEG"`
	OBS_VALID_END  int     `json:"OBS_VALID_END"`
	FCST_VAR       string  `json:"FCST_VAR"`
	FCST_UNITS     string  `json:"FCST_UNITS"`
	FCST_LEV       string  `json:"FCST_LEV"`
	OBS_VAR        string  `json:"OBS_VAR"`
	OBS_UNITS      string  `json:"OBS_UNITS"`
	OBS_LEV        string  `json:"OBS_LEV"`
	OBTYPE         string  `json:"OBTYPE"`
	VX_MASK        string  `json:"VX_MASK"`
	INTERP_MTHD    string  `json:"INTERP_MTHD"`
	INTERP_PNTS    int     `json:"INTERP_PNTS"`
	FCST_THRESH    string  `json:"FCST_THRESH"`
	OBS_THRESH     string  `json:"OBS_THRESH"`
	COV_THRESH     string  `json:"COV_THRESH"`
	ALPHA          float64 `json:"ALPHA"`
	LINE_TYPE      string  `json:"LINE_TYPE"`
}

type STAT_NBRCTS_header struct {
	VERSION        string  `json:"VERSION"`
	MODEL          string  `json:"MODEL"`
	DESC           string  `json:"DESC"`
	FCST_VALID_BEG int     `json:"FCST_VALID_BEG"`
	FCST_VALID_END int     `json:"FCST_VALID_END"`
	OBS_LEAD       int     `json:"OBS_LEAD"`
	OBS_VALID_BEG  int     `json:"OBS_VALID_BEG"`
	OBS_VALID_END  int     `json:"OBS_VALID_END"`
	FCST_VAR       string  `json:"FCST_VAR"`
	FCST_UNITS     string  `json:"FCST_UNITS"`
	FCST_LEV       string  `json:"FCST_LEV"`
	OBS_VAR        string  `json:"OBS_VAR"`
	OBS_UNITS      string  `json:"OBS_UNITS"`
	OBS_LEV        string  `json:"OBS_LEV"`
	OBTYPE         string  `json:"OBTYPE"`
	VX_MASK        string  `json:"VX_MASK"`
	INTERP_MTHD    string  `json:"INTERP_MTHD"`
	INTERP_PNTS    int     `json:"INTERP_PNTS"`
	FCST_THRESH    string  `json:"FCST_THRESH"`
	OBS_THRESH     string  `json:"OBS_THRESH"`
	COV_THRESH     string  `json:"COV_THRESH"`
	ALPHA          float64 `json:"ALPHA"`
	LINE_TYPE      string  `json:"LINE_TYPE"`
}

type STAT_ORANK_header struct {
	VERSION        string  `json:"VERSION"`
	MODEL          string  `json:"MODEL"`
	DESC           string  `json:"DESC"`
	FCST_VALID_BEG int     `json:"FCST_VALID_BEG"`
	FCST_VALID_END int     `json:"FCST_VALID_END"`
	OBS_LEAD       int     `json:"OBS_LEAD"`
	OBS_VALID_BEG  int     `json:"OBS_VALID_BEG"`
	OBS_VALID_END  int     `json:"OBS_VALID_END"`
	FCST_VAR       string  `json:"FCST_VAR"`
	FCST_UNITS     string  `json:"FCST_UNITS"`
	FCST_LEV       string  `json:"FCST_LEV"`
	OBS_VAR        string  `json:"OBS_VAR"`
	OBS_UNITS      string  `json:"OBS_UNITS"`
	OBS_LEV        string  `json:"OBS_LEV"`
	OBTYPE         string  `json:"OBTYPE"`
	VX_MASK        string  `json:"VX_MASK"`
	INTERP_MTHD    string  `json:"INTERP_MTHD"`
	INTERP_PNTS    int     `json:"INTERP_PNTS"`
	FCST_THRESH    string  `json:"FCST_THRESH"`
	OBS_THRESH     string  `json:"OBS_THRESH"`
	COV_THRESH     string  `json:"COV_THRESH"`
	ALPHA          float64 `json:"ALPHA"`
	LINE_TYPE      string  `json:"LINE_TYPE"`
}

type STAT_PCT_header struct {
	VERSION        string  `json:"VERSION"`
	MODEL          string  `json:"MODEL"`
	DESC           string  `json:"DESC"`
	FCST_VALID_BEG int     `json:"FCST_VALID_BEG"`
	FCST_VALID_END int     `json:"FCST_VALID_END"`
	OBS_LEAD       int     `json:"OBS_LEAD"`
	OBS_VALID_BEG  int     `json:"OBS_VALID_BEG"`
	OBS_VALID_END  int     `json:"OBS_VALID_END"`
	FCST_VAR       string  `json:"FCST_VAR"`
	FCST_UNITS     string  `json:"FCST_UNITS"`
	FCST_LEV       string  `json:"FCST_LEV"`
	OBS_VAR        string  `json:"OBS_VAR"`
	OBS_UNITS      string  `json:"OBS_UNITS"`
	OBS_LEV        string  `json:"OBS_LEV"`
	OBTYPE         string  `json:"OBTYPE"`
	VX_MASK        string  `json:"VX_MASK"`
	INTERP_MTHD    string  `json:"INTERP_MTHD"`
	INTERP_PNTS    int     `json:"INTERP_PNTS"`
	FCST_THRESH    string  `json:"FCST_THRESH"`
	OBS_THRESH     string  `json:"OBS_THRESH"`
	COV_THRESH     string  `json:"COV_THRESH"`
	ALPHA          float64 `json:"ALPHA"`
	LINE_TYPE      string  `json:"LINE_TYPE"`
}

type STAT_PHIST_header struct {
	VERSION        string  `json:"VERSION"`
	MODEL          string  `json:"MODEL"`
	DESC           string  `json:"DESC"`
	FCST_VALID_BEG int     `json:"FCST_VALID_BEG"`
	FCST_VALID_END int     `json:"FCST_VALID_END"`
	OBS_LEAD       int     `json:"OBS_LEAD"`
	OBS_VALID_BEG  int     `json:"OBS_VALID_BEG"`
	OBS_VALID_END  int     `json:"OBS_VALID_END"`
	FCST_VAR       string  `json:"FCST_VAR"`
	FCST_UNITS     string  `json:"FCST_UNITS"`
	FCST_LEV       string  `json:"FCST_LEV"`
	OBS_VAR        string  `json:"OBS_VAR"`
	OBS_UNITS      string  `json:"OBS_UNITS"`
	OBS_LEV        string  `json:"OBS_LEV"`
	OBTYPE         string  `json:"OBTYPE"`
	VX_MASK        string  `json:"VX_MASK"`
	INTERP_MTHD    string  `json:"INTERP_MTHD"`
	INTERP_PNTS    int     `json:"INTERP_PNTS"`
	FCST_THRESH    string  `json:"FCST_THRESH"`
	OBS_THRESH     string  `json:"OBS_THRESH"`
	COV_THRESH     string  `json:"COV_THRESH"`
	ALPHA          float64 `json:"ALPHA"`
	LINE_TYPE      string  `json:"LINE_TYPE"`
}

type STAT_PJC_header struct {
	VERSION        string  `json:"VERSION"`
	MODEL          string  `json:"MODEL"`
	DESC           string  `json:"DESC"`
	FCST_VALID_BEG int     `json:"FCST_VALID_BEG"`
	FCST_VALID_END int     `json:"FCST_VALID_END"`
	OBS_LEAD       int     `json:"OBS_LEAD"`
	OBS_VALID_BEG  int     `json:"OBS_VALID_BEG"`
	OBS_VALID_END  int     `json:"OBS_VALID_END"`
	FCST_VAR       string  `json:"FCST_VAR"`
	FCST_UNITS     string  `json:"FCST_UNITS"`
	FCST_LEV       string  `json:"FCST_LEV"`
	OBS_VAR        string  `json:"OBS_VAR"`
	OBS_UNITS      string  `json:"OBS_UNITS"`
	OBS_LEV        string  `json:"OBS_LEV"`
	OBTYPE         string  `json:"OBTYPE"`
	VX_MASK        string  `json:"VX_MASK"`
	INTERP_MTHD    string  `json:"INTERP_MTHD"`
	INTERP_PNTS    int     `json:"INTERP_PNTS"`
	FCST_THRESH    string  `json:"FCST_THRESH"`
	OBS_THRESH     string  `json:"OBS_THRESH"`
	COV_THRESH     string  `json:"COV_THRESH"`
	ALPHA          float64 `json:"ALPHA"`
	LINE_TYPE      string  `json:"LINE_TYPE"`
}

type STAT_PRC_header struct {
	VERSION        string  `json:"VERSION"`
	MODEL          string  `json:"MODEL"`
	DESC           string  `json:"DESC"`
	FCST_VALID_BEG int     `json:"FCST_VALID_BEG"`
	FCST_VALID_END int     `json:"FCST_VALID_END"`
	OBS_LEAD       int     `json:"OBS_LEAD"`
	OBS_VALID_BEG  int     `json:"OBS_VALID_BEG"`
	OBS_VALID_END  int     `json:"OBS_VALID_END"`
	FCST_VAR       string  `json:"FCST_VAR"`
	FCST_UNITS     string  `json:"FCST_UNITS"`
	FCST_LEV       string  `json:"FCST_LEV"`
	OBS_VAR        string  `json:"OBS_VAR"`
	OBS_UNITS      string  `json:"OBS_UNITS"`
	OBS_LEV        string  `json:"OBS_LEV"`
	OBTYPE         string  `json:"OBTYPE"`
	VX_MASK        string  `json:"VX_MASK"`
	INTERP_MTHD    string  `json:"INTERP_MTHD"`
	INTERP_PNTS    int     `json:"INTERP_PNTS"`
	FCST_THRESH    string  `json:"FCST_THRESH"`
	OBS_THRESH     string  `json:"OBS_THRESH"`
	COV_THRESH     string  `json:"COV_THRESH"`
	ALPHA          float64 `json:"ALPHA"`
	LINE_TYPE      string  `json:"LINE_TYPE"`
}

type STAT_PSTD_header struct {
	VERSION        string  `json:"VERSION"`
	MODEL          string  `json:"MODEL"`
	DESC           string  `json:"DESC"`
	FCST_VALID_BEG int     `json:"FCST_VALID_BEG"`
	FCST_VALID_END int     `json:"FCST_VALID_END"`
	OBS_LEAD       int     `json:"OBS_LEAD"`
	OBS_VALID_BEG  int     `json:"OBS_VALID_BEG"`
	OBS_VALID_END  int     `json:"OBS_VALID_END"`
	FCST_VAR       string  `json:"FCST_VAR"`
	FCST_UNITS     string  `json:"FCST_UNITS"`
	FCST_LEV       string  `json:"FCST_LEV"`
	OBS_VAR        string  `json:"OBS_VAR"`
	OBS_UNITS      string  `json:"OBS_UNITS"`
	OBS_LEV        string  `json:"OBS_LEV"`
	OBTYPE         string  `json:"OBTYPE"`
	VX_MASK        string  `json:"VX_MASK"`
	INTERP_MTHD    string  `json:"INTERP_MTHD"`
	INTERP_PNTS    int     `json:"INTERP_PNTS"`
	FCST_THRESH    string  `json:"FCST_THRESH"`
	OBS_THRESH     string  `json:"OBS_THRESH"`
	COV_THRESH     string  `json:"COV_THRESH"`
	ALPHA          float64 `json:"ALPHA"`
	LINE_TYPE      string  `json:"LINE_TYPE"`
}

type STAT_RELP_header struct {
	VERSION        string  `json:"VERSION"`
	MODEL          string  `json:"MODEL"`
	DESC           string  `json:"DESC"`
	FCST_VALID_BEG int     `json:"FCST_VALID_BEG"`
	FCST_VALID_END int     `json:"FCST_VALID_END"`
	OBS_LEAD       int     `json:"OBS_LEAD"`
	OBS_VALID_BEG  int     `json:"OBS_VALID_BEG"`
	OBS_VALID_END  int     `json:"OBS_VALID_END"`
	FCST_VAR       string  `json:"FCST_VAR"`
	FCST_UNITS     string  `json:"FCST_UNITS"`
	FCST_LEV       string  `json:"FCST_LEV"`
	OBS_VAR        string  `json:"OBS_VAR"`
	OBS_UNITS      string  `json:"OBS_UNITS"`
	OBS_LEV        string  `json:"OBS_LEV"`
	OBTYPE         string  `json:"OBTYPE"`
	VX_MASK        string  `json:"VX_MASK"`
	INTERP_MTHD    string  `json:"INTERP_MTHD"`
	INTERP_PNTS    int     `json:"INTERP_PNTS"`
	FCST_THRESH    string  `json:"FCST_THRESH"`
	OBS_THRESH     string  `json:"OBS_THRESH"`
	COV_THRESH     string  `json:"COV_THRESH"`
	ALPHA          float64 `json:"ALPHA"`
	LINE_TYPE      string  `json:"LINE_TYPE"`
}

type STAT_RHIST_header struct {
	VERSION        string  `json:"VERSION"`
	MODEL          string  `json:"MODEL"`
	DESC           string  `json:"DESC"`
	FCST_VALID_BEG int     `json:"FCST_VALID_BEG"`
	FCST_VALID_END int     `json:"FCST_VALID_END"`
	OBS_LEAD       int     `json:"OBS_LEAD"`
	OBS_VALID_BEG  int     `json:"OBS_VALID_BEG"`
	OBS_VALID_END  int     `json:"OBS_VALID_END"`
	FCST_VAR       string  `json:"FCST_VAR"`
	FCST_UNITS     string  `json:"FCST_UNITS"`
	FCST_LEV       string  `json:"FCST_LEV"`
	OBS_VAR        string  `json:"OBS_VAR"`
	OBS_UNITS      string  `json:"OBS_UNITS"`
	OBS_LEV        string  `json:"OBS_LEV"`
	OBTYPE         string  `json:"OBTYPE"`
	VX_MASK        string  `json:"VX_MASK"`
	INTERP_MTHD    string  `json:"INTERP_MTHD"`
	INTERP_PNTS    int     `json:"INTERP_PNTS"`
	FCST_THRESH    string  `json:"FCST_THRESH"`
	OBS_THRESH     string  `json:"OBS_THRESH"`
	COV_THRESH     string  `json:"COV_THRESH"`
	ALPHA          float64 `json:"ALPHA"`
	LINE_TYPE      string  `json:"LINE_TYPE"`
}

type STAT_RPS_header struct {
	VERSION        string  `json:"VERSION"`
	MODEL          string  `json:"MODEL"`
	DESC           string  `json:"DESC"`
	FCST_VALID_BEG int     `json:"FCST_VALID_BEG"`
	FCST_VALID_END int     `json:"FCST_VALID_END"`
	OBS_LEAD       int     `json:"OBS_LEAD"`
	OBS_VALID_BEG  int     `json:"OBS_VALID_BEG"`
	OBS_VALID_END  int     `json:"OBS_VALID_END"`
	FCST_VAR       string  `json:"FCST_VAR"`
	FCST_UNITS     string  `json:"FCST_UNITS"`
	FCST_LEV       string  `json:"FCST_LEV"`
	OBS_VAR        string  `json:"OBS_VAR"`
	OBS_UNITS      string  `json:"OBS_UNITS"`
	OBS_LEV        string  `json:"OBS_LEV"`
	OBTYPE         string  `json:"OBTYPE"`
	VX_MASK        string  `json:"VX_MASK"`
	INTERP_MTHD    string  `json:"INTERP_MTHD"`
	INTERP_PNTS    int     `json:"INTERP_PNTS"`
	FCST_THRESH    string  `json:"FCST_THRESH"`
	OBS_THRESH     string  `json:"OBS_THRESH"`
	COV_THRESH     string  `json:"COV_THRESH"`
	ALPHA          float64 `json:"ALPHA"`
	LINE_TYPE      string  `json:"LINE_TYPE"`
}

type STAT_SAL1L2_header struct {
	VERSION        string  `json:"VERSION"`
	MODEL          string  `json:"MODEL"`
	DESC           string  `json:"DESC"`
	FCST_VALID_BEG int     `json:"FCST_VALID_BEG"`
	FCST_VALID_END int     `json:"FCST_VALID_END"`
	OBS_LEAD       int     `json:"OBS_LEAD"`
	OBS_VALID_BEG  int     `json:"OBS_VALID_BEG"`
	OBS_VALID_END  int     `json:"OBS_VALID_END"`
	FCST_VAR       string  `json:"FCST_VAR"`
	FCST_UNITS     string  `json:"FCST_UNITS"`
	FCST_LEV       string  `json:"FCST_LEV"`
	OBS_VAR        string  `json:"OBS_VAR"`
	OBS_UNITS      string  `json:"OBS_UNITS"`
	OBS_LEV        string  `json:"OBS_LEV"`
	OBTYPE         string  `json:"OBTYPE"`
	VX_MASK        string  `json:"VX_MASK"`
	INTERP_MTHD    string  `json:"INTERP_MTHD"`
	INTERP_PNTS    int     `json:"INTERP_PNTS"`
	FCST_THRESH    string  `json:"FCST_THRESH"`
	OBS_THRESH     string  `json:"OBS_THRESH"`
	COV_THRESH     string  `json:"COV_THRESH"`
	ALPHA          float64 `json:"ALPHA"`
	LINE_TYPE      string  `json:"LINE_TYPE"`
}

type STAT_SEEPS_MPR_header struct {
	VERSION        string  `json:"VERSION"`
	MODEL          string  `json:"MODEL"`
	DESC           string  `json:"DESC"`
	FCST_VALID_BEG int     `json:"FCST_VALID_BEG"`
	FCST_VALID_END int     `json:"FCST_VALID_END"`
	OBS_LEAD       int     `json:"OBS_LEAD"`
	OBS_VALID_BEG  int     `json:"OBS_VALID_BEG"`
	OBS_VALID_END  int     `json:"OBS_VALID_END"`
	FCST_VAR       string  `json:"FCST_VAR"`
	FCST_UNITS     string  `json:"FCST_UNITS"`
	FCST_LEV       string  `json:"FCST_LEV"`
	OBS_VAR        string  `json:"OBS_VAR"`
	OBS_UNITS      string  `json:"OBS_UNITS"`
	OBS_LEV        string  `json:"OBS_LEV"`
	OBTYPE         string  `json:"OBTYPE"`
	VX_MASK        string  `json:"VX_MASK"`
	INTERP_MTHD    string  `json:"INTERP_MTHD"`
	INTERP_PNTS    int     `json:"INTERP_PNTS"`
	FCST_THRESH    string  `json:"FCST_THRESH"`
	OBS_THRESH     string  `json:"OBS_THRESH"`
	COV_THRESH     string  `json:"COV_THRESH"`
	ALPHA          float64 `json:"ALPHA"`
	LINE_TYPE      string  `json:"LINE_TYPE"`
}

type STAT_SEEPS_header struct {
	VERSION        string  `json:"VERSION"`
	MODEL          string  `json:"MODEL"`
	DESC           string  `json:"DESC"`
	FCST_VALID_BEG int     `json:"FCST_VALID_BEG"`
	FCST_VALID_END int     `json:"FCST_VALID_END"`
	OBS_LEAD       int     `json:"OBS_LEAD"`
	OBS_VALID_BEG  int     `json:"OBS_VALID_BEG"`
	OBS_VALID_END  int     `json:"OBS_VALID_END"`
	FCST_VAR       string  `json:"FCST_VAR"`
	FCST_UNITS     string  `json:"FCST_UNITS"`
	FCST_LEV       string  `json:"FCST_LEV"`
	OBS_VAR        string  `json:"OBS_VAR"`
	OBS_UNITS      string  `json:"OBS_UNITS"`
	OBS_LEV        string  `json:"OBS_LEV"`
	OBTYPE         string  `json:"OBTYPE"`
	VX_MASK        string  `json:"VX_MASK"`
	INTERP_MTHD    string  `json:"INTERP_MTHD"`
	INTERP_PNTS    int     `json:"INTERP_PNTS"`
	FCST_THRESH    string  `json:"FCST_THRESH"`
	OBS_THRESH     string  `json:"OBS_THRESH"`
	COV_THRESH     string  `json:"COV_THRESH"`
	ALPHA          float64 `json:"ALPHA"`
	LINE_TYPE      string  `json:"LINE_TYPE"`
}

type STAT_SL1L2_header struct {
	VERSION        string  `json:"VERSION"`
	MODEL          string  `json:"MODEL"`
	DESC           string  `json:"DESC"`
	FCST_VALID_BEG int     `json:"FCST_VALID_BEG"`
	FCST_VALID_END int     `json:"FCST_VALID_END"`
	OBS_LEAD       int     `json:"OBS_LEAD"`
	OBS_VALID_BEG  int     `json:"OBS_VALID_BEG"`
	OBS_VALID_END  int     `json:"OBS_VALID_END"`
	FCST_VAR       string  `json:"FCST_VAR"`
	FCST_UNITS     string  `json:"FCST_UNITS"`
	FCST_LEV       string  `json:"FCST_LEV"`
	OBS_VAR        string  `json:"OBS_VAR"`
	OBS_UNITS      string  `json:"OBS_UNITS"`
	OBS_LEV        string  `json:"OBS_LEV"`
	OBTYPE         string  `json:"OBTYPE"`
	VX_MASK        string  `json:"VX_MASK"`
	INTERP_MTHD    string  `json:"INTERP_MTHD"`
	INTERP_PNTS    int     `json:"INTERP_PNTS"`
	FCST_THRESH    string  `json:"FCST_THRESH"`
	OBS_THRESH     string  `json:"OBS_THRESH"`
	COV_THRESH     string  `json:"COV_THRESH"`
	ALPHA          float64 `json:"ALPHA"`
	LINE_TYPE      string  `json:"LINE_TYPE"`
}

type STAT_SSIDX_header struct {
	VERSION        string  `json:"VERSION"`
	MODEL          string  `json:"MODEL"`
	DESC           string  `json:"DESC"`
	FCST_VALID_BEG int     `json:"FCST_VALID_BEG"`
	FCST_VALID_END int     `json:"FCST_VALID_END"`
	OBS_LEAD       int     `json:"OBS_LEAD"`
	OBS_VALID_BEG  int     `json:"OBS_VALID_BEG"`
	OBS_VALID_END  int     `json:"OBS_VALID_END"`
	FCST_VAR       string  `json:"FCST_VAR"`
	FCST_UNITS     string  `json:"FCST_UNITS"`
	FCST_LEV       string  `json:"FCST_LEV"`
	OBS_VAR        string  `json:"OBS_VAR"`
	OBS_UNITS      string  `json:"OBS_UNITS"`
	OBS_LEV        string  `json:"OBS_LEV"`
	OBTYPE         string  `json:"OBTYPE"`
	VX_MASK        string  `json:"VX_MASK"`
	INTERP_MTHD    string  `json:"INTERP_MTHD"`
	INTERP_PNTS    int     `json:"INTERP_PNTS"`
	FCST_THRESH    string  `json:"FCST_THRESH"`
	OBS_THRESH     string  `json:"OBS_THRESH"`
	COV_THRESH     string  `json:"COV_THRESH"`
	ALPHA          float64 `json:"ALPHA"`
	LINE_TYPE      string  `json:"LINE_TYPE"`
}

type STAT_SSVAR_header struct {
	VERSION        string  `json:"VERSION"`
	MODEL          string  `json:"MODEL"`
	DESC           string  `json:"DESC"`
	FCST_VALID_BEG int     `json:"FCST_VALID_BEG"`
	FCST_VALID_END int     `json:"FCST_VALID_END"`
	OBS_LEAD       int     `json:"OBS_LEAD"`
	OBS_VALID_BEG  int     `json:"OBS_VALID_BEG"`
	OBS_VALID_END  int     `json:"OBS_VALID_END"`
	FCST_VAR       string  `json:"FCST_VAR"`
	FCST_UNITS     string  `json:"FCST_UNITS"`
	FCST_LEV       string  `json:"FCST_LEV"`
	OBS_VAR        string  `json:"OBS_VAR"`
	OBS_UNITS      string  `json:"OBS_UNITS"`
	OBS_LEV        string  `json:"OBS_LEV"`
	OBTYPE         string  `json:"OBTYPE"`
	VX_MASK        string  `json:"VX_MASK"`
	INTERP_MTHD    string  `json:"INTERP_MTHD"`
	INTERP_PNTS    int     `json:"INTERP_PNTS"`
	FCST_THRESH    string  `json:"FCST_THRESH"`
	OBS_THRESH     string  `json:"OBS_THRESH"`
	COV_THRESH     string  `json:"COV_THRESH"`
	ALPHA          float64 `json:"ALPHA"`
	LINE_TYPE      string  `json:"LINE_TYPE"`
}

type STAT_VAL1L2_header struct {
	VERSION        string  `json:"VERSION"`
	MODEL          string  `json:"MODEL"`
	DESC           string  `json:"DESC"`
	FCST_VALID_BEG int     `json:"FCST_VALID_BEG"`
	FCST_VALID_END int     `json:"FCST_VALID_END"`
	OBS_LEAD       int     `json:"OBS_LEAD"`
	OBS_VALID_BEG  int     `json:"OBS_VALID_BEG"`
	OBS_VALID_END  int     `json:"OBS_VALID_END"`
	FCST_VAR       string  `json:"FCST_VAR"`
	FCST_UNITS     string  `json:"FCST_UNITS"`
	FCST_LEV       string  `json:"FCST_LEV"`
	OBS_VAR        string  `json:"OBS_VAR"`
	OBS_UNITS      string  `json:"OBS_UNITS"`
	OBS_LEV        string  `json:"OBS_LEV"`
	OBTYPE         string  `json:"OBTYPE"`
	VX_MASK        string  `json:"VX_MASK"`
	INTERP_MTHD    string  `json:"INTERP_MTHD"`
	INTERP_PNTS    int     `json:"INTERP_PNTS"`
	FCST_THRESH    string  `json:"FCST_THRESH"`
	OBS_THRESH     string  `json:"OBS_THRESH"`
	COV_THRESH     string  `json:"COV_THRESH"`
	ALPHA          float64 `json:"ALPHA"`
	LINE_TYPE      string  `json:"LINE_TYPE"`
}

type STAT_VCNT_header struct {
	VERSION        string  `json:"VERSION"`
	MODEL          string  `json:"MODEL"`
	DESC           string  `json:"DESC"`
	FCST_VALID_BEG int     `json:"FCST_VALID_BEG"`
	FCST_VALID_END int     `json:"FCST_VALID_END"`
	OBS_LEAD       int     `json:"OBS_LEAD"`
	OBS_VALID_BEG  int     `json:"OBS_VALID_BEG"`
	OBS_VALID_END  int     `json:"OBS_VALID_END"`
	FCST_VAR       string  `json:"FCST_VAR"`
	FCST_UNITS     string  `json:"FCST_UNITS"`
	FCST_LEV       string  `json:"FCST_LEV"`
	OBS_VAR        string  `json:"OBS_VAR"`
	OBS_UNITS      string  `json:"OBS_UNITS"`
	OBS_LEV        string  `json:"OBS_LEV"`
	OBTYPE         string  `json:"OBTYPE"`
	VX_MASK        string  `json:"VX_MASK"`
	INTERP_MTHD    string  `json:"INTERP_MTHD"`
	INTERP_PNTS    int     `json:"INTERP_PNTS"`
	FCST_THRESH    string  `json:"FCST_THRESH"`
	OBS_THRESH     string  `json:"OBS_THRESH"`
	COV_THRESH     string  `json:"COV_THRESH"`
	ALPHA          float64 `json:"ALPHA"`
	LINE_TYPE      string  `json:"LINE_TYPE"`
}

type STAT_VL1L2_header struct {
	VERSION        string  `json:"VERSION"`
	MODEL          string  `json:"MODEL"`
	DESC           string  `json:"DESC"`
	FCST_VALID_BEG int     `json:"FCST_VALID_BEG"`
	FCST_VALID_END int     `json:"FCST_VALID_END"`
	OBS_LEAD       int     `json:"OBS_LEAD"`
	OBS_VALID_BEG  int     `json:"OBS_VALID_BEG"`
	OBS_VALID_END  int     `json:"OBS_VALID_END"`
	FCST_VAR       string  `json:"FCST_VAR"`
	FCST_UNITS     string  `json:"FCST_UNITS"`
	FCST_LEV       string  `json:"FCST_LEV"`
	OBS_VAR        string  `json:"OBS_VAR"`
	OBS_UNITS      string  `json:"OBS_UNITS"`
	OBS_LEV        string  `json:"OBS_LEV"`
	OBTYPE         string  `json:"OBTYPE"`
	VX_MASK        string  `json:"VX_MASK"`
	INTERP_MTHD    string  `json:"INTERP_MTHD"`
	INTERP_PNTS    int     `json:"INTERP_PNTS"`
	FCST_THRESH    string  `json:"FCST_THRESH"`
	OBS_THRESH     string  `json:"OBS_THRESH"`
	COV_THRESH     string  `json:"COV_THRESH"`
	ALPHA          float64 `json:"ALPHA"`
	LINE_TYPE      string  `json:"LINE_TYPE"`
}

type TCST_PROBRIRW_header struct {
	VERSION    string `json:"VERSION"`
	AMODEL     string `json:"AMODEL"`
	BMODEL     string `json:"BMODEL"`
	DESC       string `json:"DESC"`
	STORM_ID   string `json:"STORM_ID"`
	BASIN      string `json:"BASIN"`
	CYCLONE    string `json:"CYCLONE"`
	STORM_NAME string `json:"STORM_NAME"`
	VALID      int    `json:"VALID"`
	INIT_MASK  string `json:"INIT_MASK"`
	VALID_MASK string `json:"VALID_MASK"`
	LINE_TYPE  string `json:"LINE_TYPE"`
}

type TCST_TCDIAG_header struct {
	VERSION    string `json:"VERSION"`
	AMODEL     string `json:"AMODEL"`
	BMODEL     string `json:"BMODEL"`
	DESC       string `json:"DESC"`
	STORM_ID   string `json:"STORM_ID"`
	BASIN      string `json:"BASIN"`
	CYCLONE    string `json:"CYCLONE"`
	STORM_NAME string `json:"STORM_NAME"`
	VALID      int    `json:"VALID"`
	INIT_MASK  string `json:"INIT_MASK"`
	VALID_MASK string `json:"VALID_MASK"`
	LINE_TYPE  string `json:"LINE_TYPE"`
}

type TCST_TCMPR_header struct {
	VERSION    string `json:"VERSION"`
	AMODEL     string `json:"AMODEL"`
	BMODEL     string `json:"BMODEL"`
	DESC       string `json:"DESC"`
	STORM_ID   string `json:"STORM_ID"`
	BASIN      string `json:"BASIN"`
	CYCLONE    string `json:"CYCLONE"`
	STORM_NAME string `json:"STORM_NAME"`
	VALID      int    `json:"VALID"`
	INIT_MASK  string `json:"INIT_MASK"`
	VALID_MASK string `json:"VALID_MASK"`
	LINE_TYPE  string `json:"LINE_TYPE"`
}

// fillHeader functions
func (s *MODE_CTS) fill_MODE_CTS_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	SetValueForField(doc, "MODE", "VERSION", i, dataLen, fields, 0, "string")
	i++
	SetValueForField(doc, "MODE", "MODEL", i, dataLen, fields, 1, "string")
	i++
	SetValueForField(doc, "MODE", "N_VALID", i, dataLen, fields, 2, "int")
	i++
	SetValueForField(doc, "MODE", "GRID_RES", i, dataLen, fields, 3, "float64")
	i++
	SetValueForField(doc, "MODE", "DESC", i, dataLen, fields, 4, "string")
	i++
	SetValueForField(doc, "MODE", "FCST_VALID", i, dataLen, fields, 6, "string")
	i++
	SetValueForField(doc, "MODE", "FCST_ACCUM", i, dataLen, fields, 7, "string")
	i++
	SetValueForField(doc, "MODE", "OBS_LEAD", i, dataLen, fields, 8, "int")
	i++
	SetValueForField(doc, "MODE", "OBS_VALID", i, dataLen, fields, 9, "string")
	i++
	SetValueForField(doc, "MODE", "OBS_ACCUM", i, dataLen, fields, 10, "string")
	i++
	SetValueForField(doc, "MODE", "FCST_RAD", i, dataLen, fields, 11, "int")
	i++
	SetValueForField(doc, "MODE", "FCST_THR", i, dataLen, fields, 12, "string")
	i++
	SetValueForField(doc, "MODE", "OBS_RAD", i, dataLen, fields, 13, "int")
	i++
	SetValueForField(doc, "MODE", "OBS_THR", i, dataLen, fields, 14, "string")
	i++
	SetValueForField(doc, "MODE", "FCST_VAR", i, dataLen, fields, 15, "string")
	i++
	SetValueForField(doc, "MODE", "FCST_UNITS", i, dataLen, fields, 16, "string")
	i++
	SetValueForField(doc, "MODE", "FCST_LEV", i, dataLen, fields, 17, "string")
	i++
	SetValueForField(doc, "MODE", "OBS_VAR", i, dataLen, fields, 18, "string")
	i++
	SetValueForField(doc, "MODE", "OBS_UNITS", i, dataLen, fields, 19, "string")
	i++
	SetValueForField(doc, "MODE", "OBS_LEV", i, dataLen, fields, 20, "string")
	i++
	SetValueForField(doc, "MODE", "OBTYPE", i, dataLen, fields, 21, "string")
	(*doc)["LINE_TYPE"] = "MODE_CTS"
}

func (s *MODE_OBJ) fill_MODE_OBJ_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	SetValueForField(doc, "MODE", "VERSION", i, dataLen, fields, 0, "string")
	i++
	SetValueForField(doc, "MODE", "MODEL", i, dataLen, fields, 1, "string")
	i++
	SetValueForField(doc, "MODE", "N_VALID", i, dataLen, fields, 2, "int")
	i++
	SetValueForField(doc, "MODE", "GRID_RES", i, dataLen, fields, 3, "float64")
	i++
	SetValueForField(doc, "MODE", "DESC", i, dataLen, fields, 4, "string")
	i++
	SetValueForField(doc, "MODE", "FCST_VALID", i, dataLen, fields, 6, "string")
	i++
	SetValueForField(doc, "MODE", "FCST_ACCUM", i, dataLen, fields, 7, "string")
	i++
	SetValueForField(doc, "MODE", "OBS_LEAD", i, dataLen, fields, 8, "int")
	i++
	SetValueForField(doc, "MODE", "OBS_VALID", i, dataLen, fields, 9, "string")
	i++
	SetValueForField(doc, "MODE", "OBS_ACCUM", i, dataLen, fields, 10, "string")
	i++
	SetValueForField(doc, "MODE", "FCST_RAD", i, dataLen, fields, 11, "int")
	i++
	SetValueForField(doc, "MODE", "FCST_THR", i, dataLen, fields, 12, "string")
	i++
	SetValueForField(doc, "MODE", "OBS_RAD", i, dataLen, fields, 13, "int")
	i++
	SetValueForField(doc, "MODE", "OBS_THR", i, dataLen, fields, 14, "string")
	i++
	SetValueForField(doc, "MODE", "FCST_VAR", i, dataLen, fields, 15, "string")
	i++
	SetValueForField(doc, "MODE", "FCST_UNITS", i, dataLen, fields, 16, "string")
	i++
	SetValueForField(doc, "MODE", "FCST_LEV", i, dataLen, fields, 17, "string")
	i++
	SetValueForField(doc, "MODE", "OBS_VAR", i, dataLen, fields, 18, "string")
	i++
	SetValueForField(doc, "MODE", "OBS_UNITS", i, dataLen, fields, 19, "string")
	i++
	SetValueForField(doc, "MODE", "OBS_LEV", i, dataLen, fields, 20, "string")
	i++
	SetValueForField(doc, "MODE", "OBTYPE", i, dataLen, fields, 21, "string")
	(*doc)["LINE_TYPE"] = "MODE_OBJ"
}

func (s *STAT_CNT) fill_STAT_CNT_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	SetValueForField(doc, "STAT", "VERSION", i, dataLen, fields, 0, "string")
	i++
	SetValueForField(doc, "STAT", "MODEL", i, dataLen, fields, 1, "string")
	i++
	SetValueForField(doc, "STAT", "DESC", i, dataLen, fields, 2, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_BEG", i, dataLen, fields, 4, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_END", i, dataLen, fields, 5, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_LEAD", i, dataLen, fields, 6, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_BEG", i, dataLen, fields, 7, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_END", i, dataLen, fields, 8, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VAR", i, dataLen, fields, 9, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_UNITS", i, dataLen, fields, 10, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_LEV", i, dataLen, fields, 11, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_VAR", i, dataLen, fields, 12, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_UNITS", i, dataLen, fields, 13, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_LEV", i, dataLen, fields, 14, "string")
	i++
	SetValueForField(doc, "STAT", "OBTYPE", i, dataLen, fields, 15, "string")
	i++
	SetValueForField(doc, "STAT", "VX_MASK", i, dataLen, fields, 16, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_MTHD", i, dataLen, fields, 17, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_PNTS", i, dataLen, fields, 18, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_THRESH", i, dataLen, fields, 19, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_THRESH", i, dataLen, fields, 20, "string")
	i++
	SetValueForField(doc, "STAT", "COV_THRESH", i, dataLen, fields, 21, "string")
	i++
	SetValueForField(doc, "STAT", "ALPHA", i, dataLen, fields, 22, "float64")
	i++
	SetValueForField(doc, "STAT", "LINE_TYPE", i, dataLen, fields, 23, "string")
}

func (s *STAT_CTC) fill_STAT_CTC_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	SetValueForField(doc, "STAT", "VERSION", i, dataLen, fields, 0, "string")
	i++
	SetValueForField(doc, "STAT", "MODEL", i, dataLen, fields, 1, "string")
	i++
	SetValueForField(doc, "STAT", "DESC", i, dataLen, fields, 2, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_BEG", i, dataLen, fields, 4, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_END", i, dataLen, fields, 5, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_LEAD", i, dataLen, fields, 6, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_BEG", i, dataLen, fields, 7, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_END", i, dataLen, fields, 8, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VAR", i, dataLen, fields, 9, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_UNITS", i, dataLen, fields, 10, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_LEV", i, dataLen, fields, 11, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_VAR", i, dataLen, fields, 12, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_UNITS", i, dataLen, fields, 13, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_LEV", i, dataLen, fields, 14, "string")
	i++
	SetValueForField(doc, "STAT", "OBTYPE", i, dataLen, fields, 15, "string")
	i++
	SetValueForField(doc, "STAT", "VX_MASK", i, dataLen, fields, 16, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_MTHD", i, dataLen, fields, 17, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_PNTS", i, dataLen, fields, 18, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_THRESH", i, dataLen, fields, 19, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_THRESH", i, dataLen, fields, 20, "string")
	i++
	SetValueForField(doc, "STAT", "COV_THRESH", i, dataLen, fields, 21, "string")
	i++
	SetValueForField(doc, "STAT", "ALPHA", i, dataLen, fields, 22, "float64")
	i++
	SetValueForField(doc, "STAT", "LINE_TYPE", i, dataLen, fields, 23, "string")
}

func (s *STAT_CTS) fill_STAT_CTS_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	SetValueForField(doc, "STAT", "VERSION", i, dataLen, fields, 0, "string")
	i++
	SetValueForField(doc, "STAT", "MODEL", i, dataLen, fields, 1, "string")
	i++
	SetValueForField(doc, "STAT", "DESC", i, dataLen, fields, 2, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_BEG", i, dataLen, fields, 4, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_END", i, dataLen, fields, 5, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_LEAD", i, dataLen, fields, 6, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_BEG", i, dataLen, fields, 7, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_END", i, dataLen, fields, 8, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VAR", i, dataLen, fields, 9, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_UNITS", i, dataLen, fields, 10, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_LEV", i, dataLen, fields, 11, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_VAR", i, dataLen, fields, 12, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_UNITS", i, dataLen, fields, 13, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_LEV", i, dataLen, fields, 14, "string")
	i++
	SetValueForField(doc, "STAT", "OBTYPE", i, dataLen, fields, 15, "string")
	i++
	SetValueForField(doc, "STAT", "VX_MASK", i, dataLen, fields, 16, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_MTHD", i, dataLen, fields, 17, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_PNTS", i, dataLen, fields, 18, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_THRESH", i, dataLen, fields, 19, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_THRESH", i, dataLen, fields, 20, "string")
	i++
	SetValueForField(doc, "STAT", "COV_THRESH", i, dataLen, fields, 21, "string")
	i++
	SetValueForField(doc, "STAT", "ALPHA", i, dataLen, fields, 22, "float64")
	i++
	SetValueForField(doc, "STAT", "LINE_TYPE", i, dataLen, fields, 23, "string")
}

func (s *STAT_DMAP) fill_STAT_DMAP_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	SetValueForField(doc, "STAT", "VERSION", i, dataLen, fields, 0, "string")
	i++
	SetValueForField(doc, "STAT", "MODEL", i, dataLen, fields, 1, "string")
	i++
	SetValueForField(doc, "STAT", "DESC", i, dataLen, fields, 2, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_BEG", i, dataLen, fields, 4, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_END", i, dataLen, fields, 5, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_LEAD", i, dataLen, fields, 6, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_BEG", i, dataLen, fields, 7, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_END", i, dataLen, fields, 8, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VAR", i, dataLen, fields, 9, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_UNITS", i, dataLen, fields, 10, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_LEV", i, dataLen, fields, 11, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_VAR", i, dataLen, fields, 12, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_UNITS", i, dataLen, fields, 13, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_LEV", i, dataLen, fields, 14, "string")
	i++
	SetValueForField(doc, "STAT", "OBTYPE", i, dataLen, fields, 15, "string")
	i++
	SetValueForField(doc, "STAT", "VX_MASK", i, dataLen, fields, 16, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_MTHD", i, dataLen, fields, 17, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_PNTS", i, dataLen, fields, 18, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_THRESH", i, dataLen, fields, 19, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_THRESH", i, dataLen, fields, 20, "string")
	i++
	SetValueForField(doc, "STAT", "COV_THRESH", i, dataLen, fields, 21, "string")
	i++
	SetValueForField(doc, "STAT", "ALPHA", i, dataLen, fields, 22, "float64")
	i++
	SetValueForField(doc, "STAT", "LINE_TYPE", i, dataLen, fields, 23, "string")
}

func (s *STAT_ECLV) fill_STAT_ECLV_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	SetValueForField(doc, "STAT", "VERSION", i, dataLen, fields, 0, "string")
	i++
	SetValueForField(doc, "STAT", "MODEL", i, dataLen, fields, 1, "string")
	i++
	SetValueForField(doc, "STAT", "DESC", i, dataLen, fields, 2, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_BEG", i, dataLen, fields, 4, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_END", i, dataLen, fields, 5, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_LEAD", i, dataLen, fields, 6, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_BEG", i, dataLen, fields, 7, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_END", i, dataLen, fields, 8, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VAR", i, dataLen, fields, 9, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_UNITS", i, dataLen, fields, 10, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_LEV", i, dataLen, fields, 11, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_VAR", i, dataLen, fields, 12, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_UNITS", i, dataLen, fields, 13, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_LEV", i, dataLen, fields, 14, "string")
	i++
	SetValueForField(doc, "STAT", "OBTYPE", i, dataLen, fields, 15, "string")
	i++
	SetValueForField(doc, "STAT", "VX_MASK", i, dataLen, fields, 16, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_MTHD", i, dataLen, fields, 17, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_PNTS", i, dataLen, fields, 18, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_THRESH", i, dataLen, fields, 19, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_THRESH", i, dataLen, fields, 20, "string")
	i++
	SetValueForField(doc, "STAT", "COV_THRESH", i, dataLen, fields, 21, "string")
	i++
	SetValueForField(doc, "STAT", "ALPHA", i, dataLen, fields, 22, "float64")
	i++
	SetValueForField(doc, "STAT", "LINE_TYPE", i, dataLen, fields, 23, "string")
}

func (s *STAT_ECNT) fill_STAT_ECNT_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	SetValueForField(doc, "STAT", "VERSION", i, dataLen, fields, 0, "string")
	i++
	SetValueForField(doc, "STAT", "MODEL", i, dataLen, fields, 1, "string")
	i++
	SetValueForField(doc, "STAT", "DESC", i, dataLen, fields, 2, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_BEG", i, dataLen, fields, 4, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_END", i, dataLen, fields, 5, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_LEAD", i, dataLen, fields, 6, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_BEG", i, dataLen, fields, 7, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_END", i, dataLen, fields, 8, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VAR", i, dataLen, fields, 9, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_UNITS", i, dataLen, fields, 10, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_LEV", i, dataLen, fields, 11, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_VAR", i, dataLen, fields, 12, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_UNITS", i, dataLen, fields, 13, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_LEV", i, dataLen, fields, 14, "string")
	i++
	SetValueForField(doc, "STAT", "OBTYPE", i, dataLen, fields, 15, "string")
	i++
	SetValueForField(doc, "STAT", "VX_MASK", i, dataLen, fields, 16, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_MTHD", i, dataLen, fields, 17, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_PNTS", i, dataLen, fields, 18, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_THRESH", i, dataLen, fields, 19, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_THRESH", i, dataLen, fields, 20, "string")
	i++
	SetValueForField(doc, "STAT", "COV_THRESH", i, dataLen, fields, 21, "string")
	i++
	SetValueForField(doc, "STAT", "ALPHA", i, dataLen, fields, 22, "float64")
	i++
	SetValueForField(doc, "STAT", "LINE_TYPE", i, dataLen, fields, 23, "string")
}

func (s *STAT_FHO) fill_STAT_FHO_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	SetValueForField(doc, "STAT", "VERSION", i, dataLen, fields, 0, "string")
	i++
	SetValueForField(doc, "STAT", "MODEL", i, dataLen, fields, 1, "string")
	i++
	SetValueForField(doc, "STAT", "DESC", i, dataLen, fields, 2, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_BEG", i, dataLen, fields, 4, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_END", i, dataLen, fields, 5, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_LEAD", i, dataLen, fields, 6, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_BEG", i, dataLen, fields, 7, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_END", i, dataLen, fields, 8, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VAR", i, dataLen, fields, 9, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_UNITS", i, dataLen, fields, 10, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_LEV", i, dataLen, fields, 11, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_VAR", i, dataLen, fields, 12, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_UNITS", i, dataLen, fields, 13, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_LEV", i, dataLen, fields, 14, "string")
	i++
	SetValueForField(doc, "STAT", "OBTYPE", i, dataLen, fields, 15, "string")
	i++
	SetValueForField(doc, "STAT", "VX_MASK", i, dataLen, fields, 16, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_MTHD", i, dataLen, fields, 17, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_PNTS", i, dataLen, fields, 18, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_THRESH", i, dataLen, fields, 19, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_THRESH", i, dataLen, fields, 20, "string")
	i++
	SetValueForField(doc, "STAT", "COV_THRESH", i, dataLen, fields, 21, "string")
	i++
	SetValueForField(doc, "STAT", "ALPHA", i, dataLen, fields, 22, "float64")
	i++
	SetValueForField(doc, "STAT", "LINE_TYPE", i, dataLen, fields, 23, "string")
}

func (s *STAT_GENMPR) fill_STAT_GENMPR_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	SetValueForField(doc, "STAT", "VERSION", i, dataLen, fields, 0, "string")
	i++
	SetValueForField(doc, "STAT", "MODEL", i, dataLen, fields, 1, "string")
	i++
	SetValueForField(doc, "STAT", "DESC", i, dataLen, fields, 2, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_BEG", i, dataLen, fields, 4, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_END", i, dataLen, fields, 5, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_LEAD", i, dataLen, fields, 6, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_BEG", i, dataLen, fields, 7, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_END", i, dataLen, fields, 8, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VAR", i, dataLen, fields, 9, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_UNITS", i, dataLen, fields, 10, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_LEV", i, dataLen, fields, 11, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_VAR", i, dataLen, fields, 12, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_UNITS", i, dataLen, fields, 13, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_LEV", i, dataLen, fields, 14, "string")
	i++
	SetValueForField(doc, "STAT", "OBTYPE", i, dataLen, fields, 15, "string")
	i++
	SetValueForField(doc, "STAT", "VX_MASK", i, dataLen, fields, 16, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_MTHD", i, dataLen, fields, 17, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_PNTS", i, dataLen, fields, 18, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_THRESH", i, dataLen, fields, 19, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_THRESH", i, dataLen, fields, 20, "string")
	i++
	SetValueForField(doc, "STAT", "COV_THRESH", i, dataLen, fields, 21, "string")
	i++
	SetValueForField(doc, "STAT", "ALPHA", i, dataLen, fields, 22, "float64")
	i++
	SetValueForField(doc, "STAT", "LINE_TYPE", i, dataLen, fields, 23, "string")
}

func (s *STAT_GRAD) fill_STAT_GRAD_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	SetValueForField(doc, "STAT", "VERSION", i, dataLen, fields, 0, "string")
	i++
	SetValueForField(doc, "STAT", "MODEL", i, dataLen, fields, 1, "string")
	i++
	SetValueForField(doc, "STAT", "DESC", i, dataLen, fields, 2, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_BEG", i, dataLen, fields, 4, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_END", i, dataLen, fields, 5, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_LEAD", i, dataLen, fields, 6, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_BEG", i, dataLen, fields, 7, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_END", i, dataLen, fields, 8, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VAR", i, dataLen, fields, 9, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_UNITS", i, dataLen, fields, 10, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_LEV", i, dataLen, fields, 11, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_VAR", i, dataLen, fields, 12, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_UNITS", i, dataLen, fields, 13, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_LEV", i, dataLen, fields, 14, "string")
	i++
	SetValueForField(doc, "STAT", "OBTYPE", i, dataLen, fields, 15, "string")
	i++
	SetValueForField(doc, "STAT", "VX_MASK", i, dataLen, fields, 16, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_MTHD", i, dataLen, fields, 17, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_PNTS", i, dataLen, fields, 18, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_THRESH", i, dataLen, fields, 19, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_THRESH", i, dataLen, fields, 20, "string")
	i++
	SetValueForField(doc, "STAT", "COV_THRESH", i, dataLen, fields, 21, "string")
	i++
	SetValueForField(doc, "STAT", "ALPHA", i, dataLen, fields, 22, "float64")
	i++
	SetValueForField(doc, "STAT", "LINE_TYPE", i, dataLen, fields, 23, "string")
}

func (s *STAT_ISC) fill_STAT_ISC_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	SetValueForField(doc, "STAT", "VERSION", i, dataLen, fields, 0, "string")
	i++
	SetValueForField(doc, "STAT", "MODEL", i, dataLen, fields, 1, "string")
	i++
	SetValueForField(doc, "STAT", "DESC", i, dataLen, fields, 2, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_BEG", i, dataLen, fields, 4, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_END", i, dataLen, fields, 5, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_LEAD", i, dataLen, fields, 6, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_BEG", i, dataLen, fields, 7, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_END", i, dataLen, fields, 8, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VAR", i, dataLen, fields, 9, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_UNITS", i, dataLen, fields, 10, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_LEV", i, dataLen, fields, 11, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_VAR", i, dataLen, fields, 12, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_UNITS", i, dataLen, fields, 13, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_LEV", i, dataLen, fields, 14, "string")
	i++
	SetValueForField(doc, "STAT", "OBTYPE", i, dataLen, fields, 15, "string")
	i++
	SetValueForField(doc, "STAT", "VX_MASK", i, dataLen, fields, 16, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_MTHD", i, dataLen, fields, 17, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_PNTS", i, dataLen, fields, 18, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_THRESH", i, dataLen, fields, 19, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_THRESH", i, dataLen, fields, 20, "string")
	i++
	SetValueForField(doc, "STAT", "COV_THRESH", i, dataLen, fields, 21, "string")
	i++
	SetValueForField(doc, "STAT", "ALPHA", i, dataLen, fields, 22, "float64")
	i++
	SetValueForField(doc, "STAT", "LINE_TYPE", i, dataLen, fields, 23, "string")
}

func (s *STAT_MCTC) fill_STAT_MCTC_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	SetValueForField(doc, "STAT", "VERSION", i, dataLen, fields, 0, "string")
	i++
	SetValueForField(doc, "STAT", "MODEL", i, dataLen, fields, 1, "string")
	i++
	SetValueForField(doc, "STAT", "DESC", i, dataLen, fields, 2, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_BEG", i, dataLen, fields, 4, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_END", i, dataLen, fields, 5, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_LEAD", i, dataLen, fields, 6, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_BEG", i, dataLen, fields, 7, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_END", i, dataLen, fields, 8, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VAR", i, dataLen, fields, 9, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_UNITS", i, dataLen, fields, 10, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_LEV", i, dataLen, fields, 11, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_VAR", i, dataLen, fields, 12, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_UNITS", i, dataLen, fields, 13, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_LEV", i, dataLen, fields, 14, "string")
	i++
	SetValueForField(doc, "STAT", "OBTYPE", i, dataLen, fields, 15, "string")
	i++
	SetValueForField(doc, "STAT", "VX_MASK", i, dataLen, fields, 16, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_MTHD", i, dataLen, fields, 17, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_PNTS", i, dataLen, fields, 18, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_THRESH", i, dataLen, fields, 19, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_THRESH", i, dataLen, fields, 20, "string")
	i++
	SetValueForField(doc, "STAT", "COV_THRESH", i, dataLen, fields, 21, "string")
	i++
	SetValueForField(doc, "STAT", "ALPHA", i, dataLen, fields, 22, "float64")
	i++
	SetValueForField(doc, "STAT", "LINE_TYPE", i, dataLen, fields, 23, "string")
}

func (s *STAT_MCTS) fill_STAT_MCTS_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	SetValueForField(doc, "STAT", "VERSION", i, dataLen, fields, 0, "string")
	i++
	SetValueForField(doc, "STAT", "MODEL", i, dataLen, fields, 1, "string")
	i++
	SetValueForField(doc, "STAT", "DESC", i, dataLen, fields, 2, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_BEG", i, dataLen, fields, 4, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_END", i, dataLen, fields, 5, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_LEAD", i, dataLen, fields, 6, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_BEG", i, dataLen, fields, 7, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_END", i, dataLen, fields, 8, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VAR", i, dataLen, fields, 9, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_UNITS", i, dataLen, fields, 10, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_LEV", i, dataLen, fields, 11, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_VAR", i, dataLen, fields, 12, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_UNITS", i, dataLen, fields, 13, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_LEV", i, dataLen, fields, 14, "string")
	i++
	SetValueForField(doc, "STAT", "OBTYPE", i, dataLen, fields, 15, "string")
	i++
	SetValueForField(doc, "STAT", "VX_MASK", i, dataLen, fields, 16, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_MTHD", i, dataLen, fields, 17, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_PNTS", i, dataLen, fields, 18, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_THRESH", i, dataLen, fields, 19, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_THRESH", i, dataLen, fields, 20, "string")
	i++
	SetValueForField(doc, "STAT", "COV_THRESH", i, dataLen, fields, 21, "string")
	i++
	SetValueForField(doc, "STAT", "ALPHA", i, dataLen, fields, 22, "float64")
	i++
	SetValueForField(doc, "STAT", "LINE_TYPE", i, dataLen, fields, 23, "string")
}

func (s *STAT_MPR) fill_STAT_MPR_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	SetValueForField(doc, "STAT", "VERSION", i, dataLen, fields, 0, "string")
	i++
	SetValueForField(doc, "STAT", "MODEL", i, dataLen, fields, 1, "string")
	i++
	SetValueForField(doc, "STAT", "DESC", i, dataLen, fields, 2, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_BEG", i, dataLen, fields, 4, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_END", i, dataLen, fields, 5, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_LEAD", i, dataLen, fields, 6, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_BEG", i, dataLen, fields, 7, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_END", i, dataLen, fields, 8, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VAR", i, dataLen, fields, 9, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_UNITS", i, dataLen, fields, 10, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_LEV", i, dataLen, fields, 11, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_VAR", i, dataLen, fields, 12, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_UNITS", i, dataLen, fields, 13, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_LEV", i, dataLen, fields, 14, "string")
	i++
	SetValueForField(doc, "STAT", "OBTYPE", i, dataLen, fields, 15, "string")
	i++
	SetValueForField(doc, "STAT", "VX_MASK", i, dataLen, fields, 16, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_MTHD", i, dataLen, fields, 17, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_PNTS", i, dataLen, fields, 18, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_THRESH", i, dataLen, fields, 19, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_THRESH", i, dataLen, fields, 20, "string")
	i++
	SetValueForField(doc, "STAT", "COV_THRESH", i, dataLen, fields, 21, "string")
	i++
	SetValueForField(doc, "STAT", "ALPHA", i, dataLen, fields, 22, "float64")
	i++
	SetValueForField(doc, "STAT", "LINE_TYPE", i, dataLen, fields, 23, "string")
}

func (s *STAT_NBRCNT) fill_STAT_NBRCNT_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	SetValueForField(doc, "STAT", "VERSION", i, dataLen, fields, 0, "string")
	i++
	SetValueForField(doc, "STAT", "MODEL", i, dataLen, fields, 1, "string")
	i++
	SetValueForField(doc, "STAT", "DESC", i, dataLen, fields, 2, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_BEG", i, dataLen, fields, 4, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_END", i, dataLen, fields, 5, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_LEAD", i, dataLen, fields, 6, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_BEG", i, dataLen, fields, 7, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_END", i, dataLen, fields, 8, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VAR", i, dataLen, fields, 9, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_UNITS", i, dataLen, fields, 10, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_LEV", i, dataLen, fields, 11, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_VAR", i, dataLen, fields, 12, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_UNITS", i, dataLen, fields, 13, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_LEV", i, dataLen, fields, 14, "string")
	i++
	SetValueForField(doc, "STAT", "OBTYPE", i, dataLen, fields, 15, "string")
	i++
	SetValueForField(doc, "STAT", "VX_MASK", i, dataLen, fields, 16, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_MTHD", i, dataLen, fields, 17, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_PNTS", i, dataLen, fields, 18, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_THRESH", i, dataLen, fields, 19, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_THRESH", i, dataLen, fields, 20, "string")
	i++
	SetValueForField(doc, "STAT", "COV_THRESH", i, dataLen, fields, 21, "string")
	i++
	SetValueForField(doc, "STAT", "ALPHA", i, dataLen, fields, 22, "float64")
	i++
	SetValueForField(doc, "STAT", "LINE_TYPE", i, dataLen, fields, 23, "string")
}

func (s *STAT_NBRCTC) fill_STAT_NBRCTC_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	SetValueForField(doc, "STAT", "VERSION", i, dataLen, fields, 0, "string")
	i++
	SetValueForField(doc, "STAT", "MODEL", i, dataLen, fields, 1, "string")
	i++
	SetValueForField(doc, "STAT", "DESC", i, dataLen, fields, 2, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_BEG", i, dataLen, fields, 4, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_END", i, dataLen, fields, 5, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_LEAD", i, dataLen, fields, 6, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_BEG", i, dataLen, fields, 7, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_END", i, dataLen, fields, 8, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VAR", i, dataLen, fields, 9, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_UNITS", i, dataLen, fields, 10, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_LEV", i, dataLen, fields, 11, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_VAR", i, dataLen, fields, 12, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_UNITS", i, dataLen, fields, 13, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_LEV", i, dataLen, fields, 14, "string")
	i++
	SetValueForField(doc, "STAT", "OBTYPE", i, dataLen, fields, 15, "string")
	i++
	SetValueForField(doc, "STAT", "VX_MASK", i, dataLen, fields, 16, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_MTHD", i, dataLen, fields, 17, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_PNTS", i, dataLen, fields, 18, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_THRESH", i, dataLen, fields, 19, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_THRESH", i, dataLen, fields, 20, "string")
	i++
	SetValueForField(doc, "STAT", "COV_THRESH", i, dataLen, fields, 21, "string")
	i++
	SetValueForField(doc, "STAT", "ALPHA", i, dataLen, fields, 22, "float64")
	i++
	SetValueForField(doc, "STAT", "LINE_TYPE", i, dataLen, fields, 23, "string")
}

func (s *STAT_NBRCTS) fill_STAT_NBRCTS_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	SetValueForField(doc, "STAT", "VERSION", i, dataLen, fields, 0, "string")
	i++
	SetValueForField(doc, "STAT", "MODEL", i, dataLen, fields, 1, "string")
	i++
	SetValueForField(doc, "STAT", "DESC", i, dataLen, fields, 2, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_BEG", i, dataLen, fields, 4, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_END", i, dataLen, fields, 5, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_LEAD", i, dataLen, fields, 6, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_BEG", i, dataLen, fields, 7, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_END", i, dataLen, fields, 8, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VAR", i, dataLen, fields, 9, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_UNITS", i, dataLen, fields, 10, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_LEV", i, dataLen, fields, 11, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_VAR", i, dataLen, fields, 12, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_UNITS", i, dataLen, fields, 13, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_LEV", i, dataLen, fields, 14, "string")
	i++
	SetValueForField(doc, "STAT", "OBTYPE", i, dataLen, fields, 15, "string")
	i++
	SetValueForField(doc, "STAT", "VX_MASK", i, dataLen, fields, 16, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_MTHD", i, dataLen, fields, 17, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_PNTS", i, dataLen, fields, 18, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_THRESH", i, dataLen, fields, 19, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_THRESH", i, dataLen, fields, 20, "string")
	i++
	SetValueForField(doc, "STAT", "COV_THRESH", i, dataLen, fields, 21, "string")
	i++
	SetValueForField(doc, "STAT", "ALPHA", i, dataLen, fields, 22, "float64")
	i++
	SetValueForField(doc, "STAT", "LINE_TYPE", i, dataLen, fields, 23, "string")
}

func (s *STAT_ORANK) fill_STAT_ORANK_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	SetValueForField(doc, "STAT", "VERSION", i, dataLen, fields, 0, "string")
	i++
	SetValueForField(doc, "STAT", "MODEL", i, dataLen, fields, 1, "string")
	i++
	SetValueForField(doc, "STAT", "DESC", i, dataLen, fields, 2, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_BEG", i, dataLen, fields, 4, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_END", i, dataLen, fields, 5, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_LEAD", i, dataLen, fields, 6, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_BEG", i, dataLen, fields, 7, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_END", i, dataLen, fields, 8, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VAR", i, dataLen, fields, 9, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_UNITS", i, dataLen, fields, 10, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_LEV", i, dataLen, fields, 11, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_VAR", i, dataLen, fields, 12, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_UNITS", i, dataLen, fields, 13, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_LEV", i, dataLen, fields, 14, "string")
	i++
	SetValueForField(doc, "STAT", "OBTYPE", i, dataLen, fields, 15, "string")
	i++
	SetValueForField(doc, "STAT", "VX_MASK", i, dataLen, fields, 16, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_MTHD", i, dataLen, fields, 17, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_PNTS", i, dataLen, fields, 18, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_THRESH", i, dataLen, fields, 19, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_THRESH", i, dataLen, fields, 20, "string")
	i++
	SetValueForField(doc, "STAT", "COV_THRESH", i, dataLen, fields, 21, "string")
	i++
	SetValueForField(doc, "STAT", "ALPHA", i, dataLen, fields, 22, "float64")
	i++
	SetValueForField(doc, "STAT", "LINE_TYPE", i, dataLen, fields, 23, "string")
}

func (s *STAT_PCT) fill_STAT_PCT_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	SetValueForField(doc, "STAT", "VERSION", i, dataLen, fields, 0, "string")
	i++
	SetValueForField(doc, "STAT", "MODEL", i, dataLen, fields, 1, "string")
	i++
	SetValueForField(doc, "STAT", "DESC", i, dataLen, fields, 2, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_BEG", i, dataLen, fields, 4, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_END", i, dataLen, fields, 5, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_LEAD", i, dataLen, fields, 6, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_BEG", i, dataLen, fields, 7, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_END", i, dataLen, fields, 8, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VAR", i, dataLen, fields, 9, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_UNITS", i, dataLen, fields, 10, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_LEV", i, dataLen, fields, 11, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_VAR", i, dataLen, fields, 12, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_UNITS", i, dataLen, fields, 13, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_LEV", i, dataLen, fields, 14, "string")
	i++
	SetValueForField(doc, "STAT", "OBTYPE", i, dataLen, fields, 15, "string")
	i++
	SetValueForField(doc, "STAT", "VX_MASK", i, dataLen, fields, 16, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_MTHD", i, dataLen, fields, 17, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_PNTS", i, dataLen, fields, 18, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_THRESH", i, dataLen, fields, 19, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_THRESH", i, dataLen, fields, 20, "string")
	i++
	SetValueForField(doc, "STAT", "COV_THRESH", i, dataLen, fields, 21, "string")
	i++
	SetValueForField(doc, "STAT", "ALPHA", i, dataLen, fields, 22, "float64")
	i++
	SetValueForField(doc, "STAT", "LINE_TYPE", i, dataLen, fields, 23, "string")
}

func (s *STAT_PHIST) fill_STAT_PHIST_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	SetValueForField(doc, "STAT", "VERSION", i, dataLen, fields, 0, "string")
	i++
	SetValueForField(doc, "STAT", "MODEL", i, dataLen, fields, 1, "string")
	i++
	SetValueForField(doc, "STAT", "DESC", i, dataLen, fields, 2, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_BEG", i, dataLen, fields, 4, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_END", i, dataLen, fields, 5, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_LEAD", i, dataLen, fields, 6, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_BEG", i, dataLen, fields, 7, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_END", i, dataLen, fields, 8, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VAR", i, dataLen, fields, 9, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_UNITS", i, dataLen, fields, 10, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_LEV", i, dataLen, fields, 11, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_VAR", i, dataLen, fields, 12, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_UNITS", i, dataLen, fields, 13, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_LEV", i, dataLen, fields, 14, "string")
	i++
	SetValueForField(doc, "STAT", "OBTYPE", i, dataLen, fields, 15, "string")
	i++
	SetValueForField(doc, "STAT", "VX_MASK", i, dataLen, fields, 16, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_MTHD", i, dataLen, fields, 17, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_PNTS", i, dataLen, fields, 18, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_THRESH", i, dataLen, fields, 19, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_THRESH", i, dataLen, fields, 20, "string")
	i++
	SetValueForField(doc, "STAT", "COV_THRESH", i, dataLen, fields, 21, "string")
	i++
	SetValueForField(doc, "STAT", "ALPHA", i, dataLen, fields, 22, "float64")
	i++
	SetValueForField(doc, "STAT", "LINE_TYPE", i, dataLen, fields, 23, "string")
}

func (s *STAT_PJC) fill_STAT_PJC_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	SetValueForField(doc, "STAT", "VERSION", i, dataLen, fields, 0, "string")
	i++
	SetValueForField(doc, "STAT", "MODEL", i, dataLen, fields, 1, "string")
	i++
	SetValueForField(doc, "STAT", "DESC", i, dataLen, fields, 2, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_BEG", i, dataLen, fields, 4, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_END", i, dataLen, fields, 5, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_LEAD", i, dataLen, fields, 6, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_BEG", i, dataLen, fields, 7, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_END", i, dataLen, fields, 8, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VAR", i, dataLen, fields, 9, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_UNITS", i, dataLen, fields, 10, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_LEV", i, dataLen, fields, 11, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_VAR", i, dataLen, fields, 12, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_UNITS", i, dataLen, fields, 13, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_LEV", i, dataLen, fields, 14, "string")
	i++
	SetValueForField(doc, "STAT", "OBTYPE", i, dataLen, fields, 15, "string")
	i++
	SetValueForField(doc, "STAT", "VX_MASK", i, dataLen, fields, 16, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_MTHD", i, dataLen, fields, 17, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_PNTS", i, dataLen, fields, 18, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_THRESH", i, dataLen, fields, 19, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_THRESH", i, dataLen, fields, 20, "string")
	i++
	SetValueForField(doc, "STAT", "COV_THRESH", i, dataLen, fields, 21, "string")
	i++
	SetValueForField(doc, "STAT", "ALPHA", i, dataLen, fields, 22, "float64")
	i++
	SetValueForField(doc, "STAT", "LINE_TYPE", i, dataLen, fields, 23, "string")
}

func (s *STAT_PRC) fill_STAT_PRC_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	SetValueForField(doc, "STAT", "VERSION", i, dataLen, fields, 0, "string")
	i++
	SetValueForField(doc, "STAT", "MODEL", i, dataLen, fields, 1, "string")
	i++
	SetValueForField(doc, "STAT", "DESC", i, dataLen, fields, 2, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_BEG", i, dataLen, fields, 4, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_END", i, dataLen, fields, 5, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_LEAD", i, dataLen, fields, 6, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_BEG", i, dataLen, fields, 7, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_END", i, dataLen, fields, 8, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VAR", i, dataLen, fields, 9, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_UNITS", i, dataLen, fields, 10, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_LEV", i, dataLen, fields, 11, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_VAR", i, dataLen, fields, 12, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_UNITS", i, dataLen, fields, 13, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_LEV", i, dataLen, fields, 14, "string")
	i++
	SetValueForField(doc, "STAT", "OBTYPE", i, dataLen, fields, 15, "string")
	i++
	SetValueForField(doc, "STAT", "VX_MASK", i, dataLen, fields, 16, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_MTHD", i, dataLen, fields, 17, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_PNTS", i, dataLen, fields, 18, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_THRESH", i, dataLen, fields, 19, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_THRESH", i, dataLen, fields, 20, "string")
	i++
	SetValueForField(doc, "STAT", "COV_THRESH", i, dataLen, fields, 21, "string")
	i++
	SetValueForField(doc, "STAT", "ALPHA", i, dataLen, fields, 22, "float64")
	i++
	SetValueForField(doc, "STAT", "LINE_TYPE", i, dataLen, fields, 23, "string")
}

func (s *STAT_PSTD) fill_STAT_PSTD_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	SetValueForField(doc, "STAT", "VERSION", i, dataLen, fields, 0, "string")
	i++
	SetValueForField(doc, "STAT", "MODEL", i, dataLen, fields, 1, "string")
	i++
	SetValueForField(doc, "STAT", "DESC", i, dataLen, fields, 2, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_BEG", i, dataLen, fields, 4, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_END", i, dataLen, fields, 5, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_LEAD", i, dataLen, fields, 6, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_BEG", i, dataLen, fields, 7, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_END", i, dataLen, fields, 8, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VAR", i, dataLen, fields, 9, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_UNITS", i, dataLen, fields, 10, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_LEV", i, dataLen, fields, 11, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_VAR", i, dataLen, fields, 12, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_UNITS", i, dataLen, fields, 13, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_LEV", i, dataLen, fields, 14, "string")
	i++
	SetValueForField(doc, "STAT", "OBTYPE", i, dataLen, fields, 15, "string")
	i++
	SetValueForField(doc, "STAT", "VX_MASK", i, dataLen, fields, 16, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_MTHD", i, dataLen, fields, 17, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_PNTS", i, dataLen, fields, 18, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_THRESH", i, dataLen, fields, 19, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_THRESH", i, dataLen, fields, 20, "string")
	i++
	SetValueForField(doc, "STAT", "COV_THRESH", i, dataLen, fields, 21, "string")
	i++
	SetValueForField(doc, "STAT", "ALPHA", i, dataLen, fields, 22, "float64")
	i++
	SetValueForField(doc, "STAT", "LINE_TYPE", i, dataLen, fields, 23, "string")
}

func (s *STAT_RELP) fill_STAT_RELP_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	SetValueForField(doc, "STAT", "VERSION", i, dataLen, fields, 0, "string")
	i++
	SetValueForField(doc, "STAT", "MODEL", i, dataLen, fields, 1, "string")
	i++
	SetValueForField(doc, "STAT", "DESC", i, dataLen, fields, 2, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_BEG", i, dataLen, fields, 4, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_END", i, dataLen, fields, 5, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_LEAD", i, dataLen, fields, 6, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_BEG", i, dataLen, fields, 7, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_END", i, dataLen, fields, 8, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VAR", i, dataLen, fields, 9, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_UNITS", i, dataLen, fields, 10, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_LEV", i, dataLen, fields, 11, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_VAR", i, dataLen, fields, 12, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_UNITS", i, dataLen, fields, 13, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_LEV", i, dataLen, fields, 14, "string")
	i++
	SetValueForField(doc, "STAT", "OBTYPE", i, dataLen, fields, 15, "string")
	i++
	SetValueForField(doc, "STAT", "VX_MASK", i, dataLen, fields, 16, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_MTHD", i, dataLen, fields, 17, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_PNTS", i, dataLen, fields, 18, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_THRESH", i, dataLen, fields, 19, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_THRESH", i, dataLen, fields, 20, "string")
	i++
	SetValueForField(doc, "STAT", "COV_THRESH", i, dataLen, fields, 21, "string")
	i++
	SetValueForField(doc, "STAT", "ALPHA", i, dataLen, fields, 22, "float64")
	i++
	SetValueForField(doc, "STAT", "LINE_TYPE", i, dataLen, fields, 23, "string")
}

func (s *STAT_RHIST) fill_STAT_RHIST_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	SetValueForField(doc, "STAT", "VERSION", i, dataLen, fields, 0, "string")
	i++
	SetValueForField(doc, "STAT", "MODEL", i, dataLen, fields, 1, "string")
	i++
	SetValueForField(doc, "STAT", "DESC", i, dataLen, fields, 2, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_BEG", i, dataLen, fields, 4, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_END", i, dataLen, fields, 5, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_LEAD", i, dataLen, fields, 6, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_BEG", i, dataLen, fields, 7, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_END", i, dataLen, fields, 8, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VAR", i, dataLen, fields, 9, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_UNITS", i, dataLen, fields, 10, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_LEV", i, dataLen, fields, 11, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_VAR", i, dataLen, fields, 12, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_UNITS", i, dataLen, fields, 13, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_LEV", i, dataLen, fields, 14, "string")
	i++
	SetValueForField(doc, "STAT", "OBTYPE", i, dataLen, fields, 15, "string")
	i++
	SetValueForField(doc, "STAT", "VX_MASK", i, dataLen, fields, 16, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_MTHD", i, dataLen, fields, 17, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_PNTS", i, dataLen, fields, 18, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_THRESH", i, dataLen, fields, 19, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_THRESH", i, dataLen, fields, 20, "string")
	i++
	SetValueForField(doc, "STAT", "COV_THRESH", i, dataLen, fields, 21, "string")
	i++
	SetValueForField(doc, "STAT", "ALPHA", i, dataLen, fields, 22, "float64")
	i++
	SetValueForField(doc, "STAT", "LINE_TYPE", i, dataLen, fields, 23, "string")
}

func (s *STAT_RPS) fill_STAT_RPS_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	SetValueForField(doc, "STAT", "VERSION", i, dataLen, fields, 0, "string")
	i++
	SetValueForField(doc, "STAT", "MODEL", i, dataLen, fields, 1, "string")
	i++
	SetValueForField(doc, "STAT", "DESC", i, dataLen, fields, 2, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_BEG", i, dataLen, fields, 4, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_END", i, dataLen, fields, 5, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_LEAD", i, dataLen, fields, 6, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_BEG", i, dataLen, fields, 7, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_END", i, dataLen, fields, 8, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VAR", i, dataLen, fields, 9, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_UNITS", i, dataLen, fields, 10, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_LEV", i, dataLen, fields, 11, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_VAR", i, dataLen, fields, 12, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_UNITS", i, dataLen, fields, 13, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_LEV", i, dataLen, fields, 14, "string")
	i++
	SetValueForField(doc, "STAT", "OBTYPE", i, dataLen, fields, 15, "string")
	i++
	SetValueForField(doc, "STAT", "VX_MASK", i, dataLen, fields, 16, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_MTHD", i, dataLen, fields, 17, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_PNTS", i, dataLen, fields, 18, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_THRESH", i, dataLen, fields, 19, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_THRESH", i, dataLen, fields, 20, "string")
	i++
	SetValueForField(doc, "STAT", "COV_THRESH", i, dataLen, fields, 21, "string")
	i++
	SetValueForField(doc, "STAT", "ALPHA", i, dataLen, fields, 22, "float64")
	i++
	SetValueForField(doc, "STAT", "LINE_TYPE", i, dataLen, fields, 23, "string")
}

func (s *STAT_SAL1L2) fill_STAT_SAL1L2_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	SetValueForField(doc, "STAT", "VERSION", i, dataLen, fields, 0, "string")
	i++
	SetValueForField(doc, "STAT", "MODEL", i, dataLen, fields, 1, "string")
	i++
	SetValueForField(doc, "STAT", "DESC", i, dataLen, fields, 2, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_BEG", i, dataLen, fields, 4, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_END", i, dataLen, fields, 5, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_LEAD", i, dataLen, fields, 6, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_BEG", i, dataLen, fields, 7, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_END", i, dataLen, fields, 8, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VAR", i, dataLen, fields, 9, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_UNITS", i, dataLen, fields, 10, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_LEV", i, dataLen, fields, 11, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_VAR", i, dataLen, fields, 12, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_UNITS", i, dataLen, fields, 13, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_LEV", i, dataLen, fields, 14, "string")
	i++
	SetValueForField(doc, "STAT", "OBTYPE", i, dataLen, fields, 15, "string")
	i++
	SetValueForField(doc, "STAT", "VX_MASK", i, dataLen, fields, 16, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_MTHD", i, dataLen, fields, 17, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_PNTS", i, dataLen, fields, 18, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_THRESH", i, dataLen, fields, 19, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_THRESH", i, dataLen, fields, 20, "string")
	i++
	SetValueForField(doc, "STAT", "COV_THRESH", i, dataLen, fields, 21, "string")
	i++
	SetValueForField(doc, "STAT", "ALPHA", i, dataLen, fields, 22, "float64")
	i++
	SetValueForField(doc, "STAT", "LINE_TYPE", i, dataLen, fields, 23, "string")
}

func (s *STAT_SEEPS_MPR) fill_STAT_SEEPS_MPR_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	SetValueForField(doc, "STAT", "VERSION", i, dataLen, fields, 0, "string")
	i++
	SetValueForField(doc, "STAT", "MODEL", i, dataLen, fields, 1, "string")
	i++
	SetValueForField(doc, "STAT", "DESC", i, dataLen, fields, 2, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_BEG", i, dataLen, fields, 4, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_END", i, dataLen, fields, 5, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_LEAD", i, dataLen, fields, 6, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_BEG", i, dataLen, fields, 7, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_END", i, dataLen, fields, 8, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VAR", i, dataLen, fields, 9, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_UNITS", i, dataLen, fields, 10, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_LEV", i, dataLen, fields, 11, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_VAR", i, dataLen, fields, 12, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_UNITS", i, dataLen, fields, 13, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_LEV", i, dataLen, fields, 14, "string")
	i++
	SetValueForField(doc, "STAT", "OBTYPE", i, dataLen, fields, 15, "string")
	i++
	SetValueForField(doc, "STAT", "VX_MASK", i, dataLen, fields, 16, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_MTHD", i, dataLen, fields, 17, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_PNTS", i, dataLen, fields, 18, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_THRESH", i, dataLen, fields, 19, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_THRESH", i, dataLen, fields, 20, "string")
	i++
	SetValueForField(doc, "STAT", "COV_THRESH", i, dataLen, fields, 21, "string")
	i++
	SetValueForField(doc, "STAT", "ALPHA", i, dataLen, fields, 22, "float64")
	i++
	SetValueForField(doc, "STAT", "LINE_TYPE", i, dataLen, fields, 23, "string")
}

func (s *STAT_SEEPS) fill_STAT_SEEPS_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	SetValueForField(doc, "STAT", "VERSION", i, dataLen, fields, 0, "string")
	i++
	SetValueForField(doc, "STAT", "MODEL", i, dataLen, fields, 1, "string")
	i++
	SetValueForField(doc, "STAT", "DESC", i, dataLen, fields, 2, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_BEG", i, dataLen, fields, 4, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_END", i, dataLen, fields, 5, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_LEAD", i, dataLen, fields, 6, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_BEG", i, dataLen, fields, 7, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_END", i, dataLen, fields, 8, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VAR", i, dataLen, fields, 9, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_UNITS", i, dataLen, fields, 10, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_LEV", i, dataLen, fields, 11, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_VAR", i, dataLen, fields, 12, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_UNITS", i, dataLen, fields, 13, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_LEV", i, dataLen, fields, 14, "string")
	i++
	SetValueForField(doc, "STAT", "OBTYPE", i, dataLen, fields, 15, "string")
	i++
	SetValueForField(doc, "STAT", "VX_MASK", i, dataLen, fields, 16, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_MTHD", i, dataLen, fields, 17, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_PNTS", i, dataLen, fields, 18, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_THRESH", i, dataLen, fields, 19, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_THRESH", i, dataLen, fields, 20, "string")
	i++
	SetValueForField(doc, "STAT", "COV_THRESH", i, dataLen, fields, 21, "string")
	i++
	SetValueForField(doc, "STAT", "ALPHA", i, dataLen, fields, 22, "float64")
	i++
	SetValueForField(doc, "STAT", "LINE_TYPE", i, dataLen, fields, 23, "string")
}

func (s *STAT_SL1L2) fill_STAT_SL1L2_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	SetValueForField(doc, "STAT", "VERSION", i, dataLen, fields, 0, "string")
	i++
	SetValueForField(doc, "STAT", "MODEL", i, dataLen, fields, 1, "string")
	i++
	SetValueForField(doc, "STAT", "DESC", i, dataLen, fields, 2, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_BEG", i, dataLen, fields, 4, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_END", i, dataLen, fields, 5, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_LEAD", i, dataLen, fields, 6, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_BEG", i, dataLen, fields, 7, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_END", i, dataLen, fields, 8, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VAR", i, dataLen, fields, 9, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_UNITS", i, dataLen, fields, 10, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_LEV", i, dataLen, fields, 11, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_VAR", i, dataLen, fields, 12, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_UNITS", i, dataLen, fields, 13, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_LEV", i, dataLen, fields, 14, "string")
	i++
	SetValueForField(doc, "STAT", "OBTYPE", i, dataLen, fields, 15, "string")
	i++
	SetValueForField(doc, "STAT", "VX_MASK", i, dataLen, fields, 16, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_MTHD", i, dataLen, fields, 17, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_PNTS", i, dataLen, fields, 18, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_THRESH", i, dataLen, fields, 19, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_THRESH", i, dataLen, fields, 20, "string")
	i++
	SetValueForField(doc, "STAT", "COV_THRESH", i, dataLen, fields, 21, "string")
	i++
	SetValueForField(doc, "STAT", "ALPHA", i, dataLen, fields, 22, "float64")
	i++
	SetValueForField(doc, "STAT", "LINE_TYPE", i, dataLen, fields, 23, "string")
}

func (s *STAT_SSIDX) fill_STAT_SSIDX_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	SetValueForField(doc, "STAT", "VERSION", i, dataLen, fields, 0, "string")
	i++
	SetValueForField(doc, "STAT", "MODEL", i, dataLen, fields, 1, "string")
	i++
	SetValueForField(doc, "STAT", "DESC", i, dataLen, fields, 2, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_BEG", i, dataLen, fields, 4, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_END", i, dataLen, fields, 5, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_LEAD", i, dataLen, fields, 6, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_BEG", i, dataLen, fields, 7, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_END", i, dataLen, fields, 8, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VAR", i, dataLen, fields, 9, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_UNITS", i, dataLen, fields, 10, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_LEV", i, dataLen, fields, 11, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_VAR", i, dataLen, fields, 12, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_UNITS", i, dataLen, fields, 13, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_LEV", i, dataLen, fields, 14, "string")
	i++
	SetValueForField(doc, "STAT", "OBTYPE", i, dataLen, fields, 15, "string")
	i++
	SetValueForField(doc, "STAT", "VX_MASK", i, dataLen, fields, 16, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_MTHD", i, dataLen, fields, 17, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_PNTS", i, dataLen, fields, 18, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_THRESH", i, dataLen, fields, 19, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_THRESH", i, dataLen, fields, 20, "string")
	i++
	SetValueForField(doc, "STAT", "COV_THRESH", i, dataLen, fields, 21, "string")
	i++
	SetValueForField(doc, "STAT", "ALPHA", i, dataLen, fields, 22, "float64")
	i++
	SetValueForField(doc, "STAT", "LINE_TYPE", i, dataLen, fields, 23, "string")
}

func (s *STAT_SSVAR) fill_STAT_SSVAR_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	SetValueForField(doc, "STAT", "VERSION", i, dataLen, fields, 0, "string")
	i++
	SetValueForField(doc, "STAT", "MODEL", i, dataLen, fields, 1, "string")
	i++
	SetValueForField(doc, "STAT", "DESC", i, dataLen, fields, 2, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_BEG", i, dataLen, fields, 4, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_END", i, dataLen, fields, 5, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_LEAD", i, dataLen, fields, 6, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_BEG", i, dataLen, fields, 7, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_END", i, dataLen, fields, 8, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VAR", i, dataLen, fields, 9, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_UNITS", i, dataLen, fields, 10, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_LEV", i, dataLen, fields, 11, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_VAR", i, dataLen, fields, 12, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_UNITS", i, dataLen, fields, 13, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_LEV", i, dataLen, fields, 14, "string")
	i++
	SetValueForField(doc, "STAT", "OBTYPE", i, dataLen, fields, 15, "string")
	i++
	SetValueForField(doc, "STAT", "VX_MASK", i, dataLen, fields, 16, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_MTHD", i, dataLen, fields, 17, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_PNTS", i, dataLen, fields, 18, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_THRESH", i, dataLen, fields, 19, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_THRESH", i, dataLen, fields, 20, "string")
	i++
	SetValueForField(doc, "STAT", "COV_THRESH", i, dataLen, fields, 21, "string")
	i++
	SetValueForField(doc, "STAT", "ALPHA", i, dataLen, fields, 22, "float64")
	i++
	SetValueForField(doc, "STAT", "LINE_TYPE", i, dataLen, fields, 23, "string")
}

func (s *STAT_VAL1L2) fill_STAT_VAL1L2_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	SetValueForField(doc, "STAT", "VERSION", i, dataLen, fields, 0, "string")
	i++
	SetValueForField(doc, "STAT", "MODEL", i, dataLen, fields, 1, "string")
	i++
	SetValueForField(doc, "STAT", "DESC", i, dataLen, fields, 2, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_BEG", i, dataLen, fields, 4, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_END", i, dataLen, fields, 5, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_LEAD", i, dataLen, fields, 6, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_BEG", i, dataLen, fields, 7, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_END", i, dataLen, fields, 8, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VAR", i, dataLen, fields, 9, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_UNITS", i, dataLen, fields, 10, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_LEV", i, dataLen, fields, 11, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_VAR", i, dataLen, fields, 12, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_UNITS", i, dataLen, fields, 13, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_LEV", i, dataLen, fields, 14, "string")
	i++
	SetValueForField(doc, "STAT", "OBTYPE", i, dataLen, fields, 15, "string")
	i++
	SetValueForField(doc, "STAT", "VX_MASK", i, dataLen, fields, 16, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_MTHD", i, dataLen, fields, 17, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_PNTS", i, dataLen, fields, 18, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_THRESH", i, dataLen, fields, 19, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_THRESH", i, dataLen, fields, 20, "string")
	i++
	SetValueForField(doc, "STAT", "COV_THRESH", i, dataLen, fields, 21, "string")
	i++
	SetValueForField(doc, "STAT", "ALPHA", i, dataLen, fields, 22, "float64")
	i++
	SetValueForField(doc, "STAT", "LINE_TYPE", i, dataLen, fields, 23, "string")
}

func (s *STAT_VCNT) fill_STAT_VCNT_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	SetValueForField(doc, "STAT", "VERSION", i, dataLen, fields, 0, "string")
	i++
	SetValueForField(doc, "STAT", "MODEL", i, dataLen, fields, 1, "string")
	i++
	SetValueForField(doc, "STAT", "DESC", i, dataLen, fields, 2, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_BEG", i, dataLen, fields, 4, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_END", i, dataLen, fields, 5, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_LEAD", i, dataLen, fields, 6, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_BEG", i, dataLen, fields, 7, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_END", i, dataLen, fields, 8, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VAR", i, dataLen, fields, 9, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_UNITS", i, dataLen, fields, 10, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_LEV", i, dataLen, fields, 11, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_VAR", i, dataLen, fields, 12, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_UNITS", i, dataLen, fields, 13, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_LEV", i, dataLen, fields, 14, "string")
	i++
	SetValueForField(doc, "STAT", "OBTYPE", i, dataLen, fields, 15, "string")
	i++
	SetValueForField(doc, "STAT", "VX_MASK", i, dataLen, fields, 16, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_MTHD", i, dataLen, fields, 17, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_PNTS", i, dataLen, fields, 18, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_THRESH", i, dataLen, fields, 19, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_THRESH", i, dataLen, fields, 20, "string")
	i++
	SetValueForField(doc, "STAT", "COV_THRESH", i, dataLen, fields, 21, "string")
	i++
	SetValueForField(doc, "STAT", "ALPHA", i, dataLen, fields, 22, "float64")
	i++
	SetValueForField(doc, "STAT", "LINE_TYPE", i, dataLen, fields, 23, "string")
}

func (s *STAT_VL1L2) fill_STAT_VL1L2_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	SetValueForField(doc, "STAT", "VERSION", i, dataLen, fields, 0, "string")
	i++
	SetValueForField(doc, "STAT", "MODEL", i, dataLen, fields, 1, "string")
	i++
	SetValueForField(doc, "STAT", "DESC", i, dataLen, fields, 2, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_BEG", i, dataLen, fields, 4, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VALID_END", i, dataLen, fields, 5, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_LEAD", i, dataLen, fields, 6, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_BEG", i, dataLen, fields, 7, "int")
	i++
	SetValueForField(doc, "STAT", "OBS_VALID_END", i, dataLen, fields, 8, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_VAR", i, dataLen, fields, 9, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_UNITS", i, dataLen, fields, 10, "string")
	i++
	SetValueForField(doc, "STAT", "FCST_LEV", i, dataLen, fields, 11, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_VAR", i, dataLen, fields, 12, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_UNITS", i, dataLen, fields, 13, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_LEV", i, dataLen, fields, 14, "string")
	i++
	SetValueForField(doc, "STAT", "OBTYPE", i, dataLen, fields, 15, "string")
	i++
	SetValueForField(doc, "STAT", "VX_MASK", i, dataLen, fields, 16, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_MTHD", i, dataLen, fields, 17, "string")
	i++
	SetValueForField(doc, "STAT", "INTERP_PNTS", i, dataLen, fields, 18, "int")
	i++
	SetValueForField(doc, "STAT", "FCST_THRESH", i, dataLen, fields, 19, "string")
	i++
	SetValueForField(doc, "STAT", "OBS_THRESH", i, dataLen, fields, 20, "string")
	i++
	SetValueForField(doc, "STAT", "COV_THRESH", i, dataLen, fields, 21, "string")
	i++
	SetValueForField(doc, "STAT", "ALPHA", i, dataLen, fields, 22, "float64")
	i++
	SetValueForField(doc, "STAT", "LINE_TYPE", i, dataLen, fields, 23, "string")
}

func (s *TCST_PROBRIRW) fill_TCST_PROBRIRW_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	SetValueForField(doc, "TCST", "VERSION", i, dataLen, fields, 0, "string")
	i++
	SetValueForField(doc, "TCST", "AMODEL", i, dataLen, fields, 1, "string")
	i++
	SetValueForField(doc, "TCST", "BMODEL", i, dataLen, fields, 2, "string")
	i++
	SetValueForField(doc, "TCST", "DESC", i, dataLen, fields, 3, "string")
	i++
	SetValueForField(doc, "TCST", "STORM_ID", i, dataLen, fields, 4, "string")
	i++
	SetValueForField(doc, "TCST", "BASIN", i, dataLen, fields, 5, "string")
	i++
	SetValueForField(doc, "TCST", "CYCLONE", i, dataLen, fields, 6, "string")
	i++
	SetValueForField(doc, "TCST", "STORM_NAME", i, dataLen, fields, 7, "string")
	i++
	SetValueForField(doc, "TCST", "VALID", i, dataLen, fields, 10, "int")
	i++
	SetValueForField(doc, "TCST", "INIT_MASK", i, dataLen, fields, 11, "string")
	i++
	SetValueForField(doc, "TCST", "VALID_MASK", i, dataLen, fields, 12, "string")
	i++
	SetValueForField(doc, "TCST", "LINE_TYPE", i, dataLen, fields, 13, "string")
}

func (s *TCST_TCDIAG) fill_TCST_TCDIAG_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	SetValueForField(doc, "TCST", "VERSION", i, dataLen, fields, 0, "string")
	i++
	SetValueForField(doc, "TCST", "AMODEL", i, dataLen, fields, 1, "string")
	i++
	SetValueForField(doc, "TCST", "BMODEL", i, dataLen, fields, 2, "string")
	i++
	SetValueForField(doc, "TCST", "DESC", i, dataLen, fields, 3, "string")
	i++
	SetValueForField(doc, "TCST", "STORM_ID", i, dataLen, fields, 4, "string")
	i++
	SetValueForField(doc, "TCST", "BASIN", i, dataLen, fields, 5, "string")
	i++
	SetValueForField(doc, "TCST", "CYCLONE", i, dataLen, fields, 6, "string")
	i++
	SetValueForField(doc, "TCST", "STORM_NAME", i, dataLen, fields, 7, "string")
	i++
	SetValueForField(doc, "TCST", "VALID", i, dataLen, fields, 10, "int")
	i++
	SetValueForField(doc, "TCST", "INIT_MASK", i, dataLen, fields, 11, "string")
	i++
	SetValueForField(doc, "TCST", "VALID_MASK", i, dataLen, fields, 12, "string")
	i++
	SetValueForField(doc, "TCST", "LINE_TYPE", i, dataLen, fields, 13, "string")
}

func (s *TCST_TCMPR) fill_TCST_TCMPR_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	SetValueForField(doc, "TCST", "VERSION", i, dataLen, fields, 0, "string")
	i++
	SetValueForField(doc, "TCST", "AMODEL", i, dataLen, fields, 1, "string")
	i++
	SetValueForField(doc, "TCST", "BMODEL", i, dataLen, fields, 2, "string")
	i++
	SetValueForField(doc, "TCST", "DESC", i, dataLen, fields, 3, "string")
	i++
	SetValueForField(doc, "TCST", "STORM_ID", i, dataLen, fields, 4, "string")
	i++
	SetValueForField(doc, "TCST", "BASIN", i, dataLen, fields, 5, "string")
	i++
	SetValueForField(doc, "TCST", "CYCLONE", i, dataLen, fields, 6, "string")
	i++
	SetValueForField(doc, "TCST", "STORM_NAME", i, dataLen, fields, 7, "string")
	i++
	SetValueForField(doc, "TCST", "VALID", i, dataLen, fields, 10, "int")
	i++
	SetValueForField(doc, "TCST", "INIT_MASK", i, dataLen, fields, 11, "string")
	i++
	SetValueForField(doc, "TCST", "VALID_MASK", i, dataLen, fields, 12, "string")
	i++
	SetValueForField(doc, "TCST", "LINE_TYPE", i, dataLen, fields, 13, "string")
}

//line data struct definitions
type MODE_CTS struct {
	FIELD string  `json:"FIELD,omitempty"`
	TOTAL int     `json:"TOTAL,omitempty"`
	FY_OY float64 `json:"FY_OY,omitempty"`
	FY_ON float64 `json:"FY_ON,omitempty"`
	FN_OY float64 `json:"FN_OY,omitempty"`
	FN_ON float64 `json:"FN_ON,omitempty"`
	BASER float64 `json:"BASER,omitempty"`
	FMEAN float64 `json:"FMEAN,omitempty"`
	ACC   float64 `json:"ACC,omitempty"`
	FBIAS float64 `json:"FBIAS,omitempty"`
	PODY  float64 `json:"PODY,omitempty"`
	PODN  float64 `json:"PODN,omitempty"`
	POFD  float64 `json:"POFD,omitempty"`
	FAR   float64 `json:"FAR,omitempty"`
	CSI   float64 `json:"CSI,omitempty"`
	GSS   float64 `json:"GSS,omitempty"`
	HK    float64 `json:"HK,omitempty"`
	HSS   float64 `json:"HSS,omitempty"`
	ODDS  float64 `json:"ODDS,omitempty"`
}

type MODE_OBJ struct {
	OBJECT_ID                  string  `json:"OBJECT_ID,omitempty"`
	OBJECT_CAT                 string  `json:"OBJECT_CAT,omitempty"`
	CENTROID_X                 float64 `json:"CENTROID_X,omitempty"`
	CENTROID_Y                 float64 `json:"CENTROID_Y,omitempty"`
	CENTROID_LAT               float64 `json:"CENTROID_LAT,omitempty"`
	CENTROID_LON               float64 `json:"CENTROID_LON,omitempty"`
	AXIS_ANG                   float64 `json:"AXIS_ANG,omitempty"`
	LENGTH                     float64 `json:"LENGTH,omitempty"`
	WIDTH                      float64 `json:"WIDTH,omitempty"`
	AREA                       int     `json:"AREA,omitempty"`
	AREA_THRESH                int     `json:"AREA_THRESH,omitempty"`
	CURVATURE                  float64 `json:"CURVATURE,omitempty"`
	CURVATURE_X                float64 `json:"CURVATURE_X,omitempty"`
	CURVATURE_Y                float64 `json:"CURVATURE_Y,omitempty"`
	COMPLEXITY                 float64 `json:"COMPLEXITY,omitempty"`
	INTENSITY_10               float64 `json:"INTENSITY_10,omitempty"`
	INTENSITY_25               float64 `json:"INTENSITY_25,omitempty"`
	INTENSITY_50               float64 `json:"INTENSITY_50,omitempty"`
	INTENSITY_75               float64 `json:"INTENSITY_75,omitempty"`
	INTENSITY_90               float64 `json:"INTENSITY_90,omitempty"`
	INTENSITY_USER             float64 `json:"INTENSITY_USER,omitempty"`
	INTENSITY_SUM              float64 `json:"INTENSITY_SUM,omitempty"`
	CENTROID_DIST              float64 `json:"CENTROID_DIST,omitempty"`
	BOUNDARY_DIST              float64 `json:"BOUNDARY_DIST,omitempty"`
	CONVEX_HULL_DIST           float64 `json:"CONVEX_HULL_DIST,omitempty"`
	ANGLE_DIFF                 float64 `json:"ANGLE_DIFF,omitempty"`
	ASPECT_DIFF                float64 `json:"ASPECT_DIFF,omitempty"`
	AREA_RATIO                 float64 `json:"AREA_RATIO,omitempty"`
	INTERSECTION_AREA          float64 `json:"INTERSECTION_AREA,omitempty"`
	UNION_AREA                 float64 `json:"UNION_AREA,omitempty"`
	SYMMETRIC_DIFF             float64 `json:"SYMMETRIC_DIFF,omitempty"`
	INTERSECTION_OVER_AREA     float64 `json:"INTERSECTION_OVER_AREA,omitempty"`
	CURVATURE_RATIO            float64 `json:"CURVATURE_RATIO,omitempty"`
	COMPLEXITY_RATIO           float64 `json:"COMPLEXITY_RATIO,omitempty"`
	PERCENTILE_INTENSITY_RATIO float64 `json:"PERCENTILE_INTENSITY_RATIO,omitempty"`
	INTEREST                   float64 `json:"INTEREST,omitempty"`
}

type STAT_CNT struct {
	TOTAL                int     `json:"TOTAL,omitempty"`
	FBAR                 float64 `json:"FBAR,omitempty"`
	FBAR_NCL             float64 `json:"FBAR_NCL,omitempty"`
	FBAR_NCU             float64 `json:"FBAR_NCU,omitempty"`
	FBAR_BCL             float64 `json:"FBAR_BCL,omitempty"`
	FBAR_BCU             float64 `json:"FBAR_BCU,omitempty"`
	FSTDEV               float64 `json:"FSTDEV,omitempty"`
	FSTDEV_NCL           float64 `json:"FSTDEV_NCL,omitempty"`
	FSTDEV_NCU           float64 `json:"FSTDEV_NCU,omitempty"`
	FSTDEV_BCL           float64 `json:"FSTDEV_BCL,omitempty"`
	FSTDEV_BCU           float64 `json:"FSTDEV_BCU,omitempty"`
	OBAR                 float64 `json:"OBAR,omitempty"`
	OBAR_NCL             float64 `json:"OBAR_NCL,omitempty"`
	OBAR_NCU             float64 `json:"OBAR_NCU,omitempty"`
	OBAR_BCL             float64 `json:"OBAR_BCL,omitempty"`
	OBAR_BCU             float64 `json:"OBAR_BCU,omitempty"`
	OSTDEV               float64 `json:"OSTDEV,omitempty"`
	OSTDEV_NCL           float64 `json:"OSTDEV_NCL,omitempty"`
	OSTDEV_NCU           float64 `json:"OSTDEV_NCU,omitempty"`
	OSTDEV_BCL           float64 `json:"OSTDEV_BCL,omitempty"`
	OSTDEV_BCU           float64 `json:"OSTDEV_BCU,omitempty"`
	PR_CORR              float64 `json:"PR_CORR,omitempty"`
	PR_CORR_NCL          float64 `json:"PR_CORR_NCL,omitempty"`
	PR_CORR_NCU          float64 `json:"PR_CORR_NCU,omitempty"`
	PR_CORR_BCL          float64 `json:"PR_CORR_BCL,omitempty"`
	PR_CORR_BCU          float64 `json:"PR_CORR_BCU,omitempty"`
	SP_CORR              float64 `json:"SP_CORR,omitempty"`
	KT_CORR              float64 `json:"KT_CORR,omitempty"`
	RANKS                int     `json:"RANKS,omitempty"`
	FRANK_TIES           int     `json:"FRANK_TIES,omitempty"`
	ORANK_TIES           int     `json:"ORANK_TIES,omitempty"`
	ME                   float64 `json:"ME,omitempty"`
	ME_NCL               float64 `json:"ME_NCL,omitempty"`
	ME_NCU               float64 `json:"ME_NCU,omitempty"`
	ME_BCL               float64 `json:"ME_BCL,omitempty"`
	ME_BCU               float64 `json:"ME_BCU,omitempty"`
	ESTDEV               float64 `json:"ESTDEV,omitempty"`
	ESTDEV_NCL           float64 `json:"ESTDEV_NCL,omitempty"`
	ESTDEV_NCU           float64 `json:"ESTDEV_NCU,omitempty"`
	ESTDEV_BCL           float64 `json:"ESTDEV_BCL,omitempty"`
	ESTDEV_BCU           float64 `json:"ESTDEV_BCU,omitempty"`
	MBIAS                float64 `json:"MBIAS,omitempty"`
	MBIAS_BCL            float64 `json:"MBIAS_BCL,omitempty"`
	MBIAS_BCU            float64 `json:"MBIAS_BCU,omitempty"`
	MAE                  float64 `json:"MAE,omitempty"`
	MAE_BCL              float64 `json:"MAE_BCL,omitempty"`
	MAE_BCU              float64 `json:"MAE_BCU,omitempty"`
	MSE                  float64 `json:"MSE,omitempty"`
	MSE_BCL              float64 `json:"MSE_BCL,omitempty"`
	MSE_BCU              float64 `json:"MSE_BCU,omitempty"`
	BCMSE                float64 `json:"BCMSE,omitempty"`
	BCMSE_BCL            float64 `json:"BCMSE_BCL,omitempty"`
	BCMSE_BCU            float64 `json:"BCMSE_BCU,omitempty"`
	RMSE                 float64 `json:"RMSE,omitempty"`
	RMSE_BCL             float64 `json:"RMSE_BCL,omitempty"`
	RMSE_BCU             float64 `json:"RMSE_BCU,omitempty"`
	E10                  float64 `json:"E10,omitempty"`
	E10_BCL              float64 `json:"E10_BCL,omitempty"`
	E10_BCU              float64 `json:"E10_BCU,omitempty"`
	E25                  float64 `json:"E25,omitempty"`
	E25_BCL              float64 `json:"E25_BCL,omitempty"`
	E25_BCU              float64 `json:"E25_BCU,omitempty"`
	E50                  float64 `json:"E50,omitempty"`
	E50_BCL              float64 `json:"E50_BCL,omitempty"`
	E50_BCU              float64 `json:"E50_BCU,omitempty"`
	E75                  float64 `json:"E75,omitempty"`
	E75_BCL              float64 `json:"E75_BCL,omitempty"`
	E75_BCU              float64 `json:"E75_BCU,omitempty"`
	E90                  float64 `json:"E90,omitempty"`
	E90_BCL              float64 `json:"E90_BCL,omitempty"`
	E90_BCU              float64 `json:"E90_BCU,omitempty"`
	EIQR                 float64 `json:"EIQR,omitempty"`
	EIQR_BCL             float64 `json:"EIQR_BCL,omitempty"`
	EIQR_BCU             float64 `json:"EIQR_BCU,omitempty"`
	MAD                  float64 `json:"MAD,omitempty"`
	MAD_BCL              float64 `json:"MAD_BCL,omitempty"`
	MAD_BCU              float64 `json:"MAD_BCU,omitempty"`
	ANOM_CORR            float64 `json:"ANOM_CORR,omitempty"`
	ANOM_CORR_NCL        float64 `json:"ANOM_CORR_NCL,omitempty"`
	ANOM_CORR_NCU        float64 `json:"ANOM_CORR_NCU,omitempty"`
	ANOM_CORR_BCL        float64 `json:"ANOM_CORR_BCL,omitempty"`
	ANOM_CORR_BCU        float64 `json:"ANOM_CORR_BCU,omitempty"`
	ME2                  float64 `json:"ME2,omitempty"`
	ME2_BCL              float64 `json:"ME2_BCL,omitempty"`
	ME2_BCU              float64 `json:"ME2_BCU,omitempty"`
	MSESS                float64 `json:"MSESS,omitempty"`
	MSESS_BCL            float64 `json:"MSESS_BCL,omitempty"`
	MSESS_BCU            float64 `json:"MSESS_BCU,omitempty"`
	RMSFA                float64 `json:"RMSFA,omitempty"`
	RMSFA_BCL            float64 `json:"RMSFA_BCL,omitempty"`
	RMSFA_BCU            float64 `json:"RMSFA_BCU,omitempty"`
	RMSOA                float64 `json:"RMSOA,omitempty"`
	RMSOA_BCL            float64 `json:"RMSOA_BCL,omitempty"`
	RMSOA_BCU            float64 `json:"RMSOA_BCU,omitempty"`
	ANOM_CORR_UNCNTR     float64 `json:"ANOM_CORR_UNCNTR,omitempty"`
	ANOM_CORR_UNCNTR_BCL float64 `json:"ANOM_CORR_UNCNTR_BCL,omitempty"`
	ANOM_CORR_UNCNTR_BCU float64 `json:"ANOM_CORR_UNCNTR_BCU,omitempty"`
	SI                   float64 `json:"SI,omitempty"`
	SI_BCL               float64 `json:"SI_BCL,omitempty"`
	SI_BCU               float64 `json:"SI_BCU,omitempty"`
}

type STAT_CTC struct {
	TOTAL    int     `json:"TOTAL,omitempty"`
	FY_OY    float64 `json:"FY_OY,omitempty"`
	FY_ON    float64 `json:"FY_ON,omitempty"`
	FN_OY    float64 `json:"FN_OY,omitempty"`
	FN_ON    float64 `json:"FN_ON,omitempty"`
	EC_VALUE float64 `json:"EC_VALUE,omitempty"`
}

type STAT_CTS struct {
	TOTAL      int     `json:"TOTAL,omitempty"`
	BASER      float64 `json:"BASER,omitempty"`
	BASER_NCL  float64 `json:"BASER_NCL,omitempty"`
	BASER_NCU  float64 `json:"BASER_NCU,omitempty"`
	BASER_BCL  float64 `json:"BASER_BCL,omitempty"`
	BASER_BCU  float64 `json:"BASER_BCU,omitempty"`
	FMEAN      float64 `json:"FMEAN,omitempty"`
	FMEAN_NCL  float64 `json:"FMEAN_NCL,omitempty"`
	FMEAN_NCU  float64 `json:"FMEAN_NCU,omitempty"`
	FMEAN_BCL  float64 `json:"FMEAN_BCL,omitempty"`
	FMEAN_BCU  float64 `json:"FMEAN_BCU,omitempty"`
	ACC        float64 `json:"ACC,omitempty"`
	ACC_NCL    float64 `json:"ACC_NCL,omitempty"`
	ACC_NCU    float64 `json:"ACC_NCU,omitempty"`
	ACC_BCL    float64 `json:"ACC_BCL,omitempty"`
	ACC_BCU    float64 `json:"ACC_BCU,omitempty"`
	FBIAS      float64 `json:"FBIAS,omitempty"`
	FBIAS_BCL  float64 `json:"FBIAS_BCL,omitempty"`
	FBIAS_BCU  float64 `json:"FBIAS_BCU,omitempty"`
	PODY       float64 `json:"PODY,omitempty"`
	PODY_NCL   float64 `json:"PODY_NCL,omitempty"`
	PODY_NCU   float64 `json:"PODY_NCU,omitempty"`
	PODY_BCL   float64 `json:"PODY_BCL,omitempty"`
	PODY_BCU   float64 `json:"PODY_BCU,omitempty"`
	PODN       float64 `json:"PODN,omitempty"`
	PODN_NCL   float64 `json:"PODN_NCL,omitempty"`
	PODN_NCU   float64 `json:"PODN_NCU,omitempty"`
	PODN_BCL   float64 `json:"PODN_BCL,omitempty"`
	PODN_BCU   float64 `json:"PODN_BCU,omitempty"`
	POFD       float64 `json:"POFD,omitempty"`
	POFD_NCL   float64 `json:"POFD_NCL,omitempty"`
	POFD_NCU   float64 `json:"POFD_NCU,omitempty"`
	POFD_BCL   float64 `json:"POFD_BCL,omitempty"`
	POFD_BCU   float64 `json:"POFD_BCU,omitempty"`
	FAR        float64 `json:"FAR,omitempty"`
	FAR_NCL    float64 `json:"FAR_NCL,omitempty"`
	FAR_NCU    float64 `json:"FAR_NCU,omitempty"`
	FAR_BCL    float64 `json:"FAR_BCL,omitempty"`
	FAR_BCU    float64 `json:"FAR_BCU,omitempty"`
	CSI        float64 `json:"CSI,omitempty"`
	CSI_NCL    float64 `json:"CSI_NCL,omitempty"`
	CSI_NCU    float64 `json:"CSI_NCU,omitempty"`
	CSI_BCL    float64 `json:"CSI_BCL,omitempty"`
	CSI_BCU    float64 `json:"CSI_BCU,omitempty"`
	GSS        float64 `json:"GSS,omitempty"`
	GSS_BCL    float64 `json:"GSS_BCL,omitempty"`
	GSS_BCU    float64 `json:"GSS_BCU,omitempty"`
	HK         float64 `json:"HK,omitempty"`
	HK_NCL     float64 `json:"HK_NCL,omitempty"`
	HK_NCU     float64 `json:"HK_NCU,omitempty"`
	HK_BCL     float64 `json:"HK_BCL,omitempty"`
	HK_BCU     float64 `json:"HK_BCU,omitempty"`
	HSS        float64 `json:"HSS,omitempty"`
	HSS_BCL    float64 `json:"HSS_BCL,omitempty"`
	HSS_BCU    float64 `json:"HSS_BCU,omitempty"`
	ODDS       float64 `json:"ODDS,omitempty"`
	ODDS_NCL   float64 `json:"ODDS_NCL,omitempty"`
	ODDS_NCU   float64 `json:"ODDS_NCU,omitempty"`
	ODDS_BCL   float64 `json:"ODDS_BCL,omitempty"`
	ODDS_BCU   float64 `json:"ODDS_BCU,omitempty"`
	LODDS      float64 `json:"LODDS,omitempty"`
	LODDS_NCL  float64 `json:"LODDS_NCL,omitempty"`
	LODDS_NCU  float64 `json:"LODDS_NCU,omitempty"`
	LODDS_BCL  float64 `json:"LODDS_BCL,omitempty"`
	LODDS_BCU  float64 `json:"LODDS_BCU,omitempty"`
	ORSS       float64 `json:"ORSS,omitempty"`
	ORSS_NCL   float64 `json:"ORSS_NCL,omitempty"`
	ORSS_NCU   float64 `json:"ORSS_NCU,omitempty"`
	ORSS_BCL   float64 `json:"ORSS_BCL,omitempty"`
	ORSS_BCU   float64 `json:"ORSS_BCU,omitempty"`
	EDS        float64 `json:"EDS,omitempty"`
	EDS_NCL    float64 `json:"EDS_NCL,omitempty"`
	EDS_NCU    float64 `json:"EDS_NCU,omitempty"`
	EDS_BCL    float64 `json:"EDS_BCL,omitempty"`
	EDS_BCU    float64 `json:"EDS_BCU,omitempty"`
	SEDS       float64 `json:"SEDS,omitempty"`
	SEDS_NCL   float64 `json:"SEDS_NCL,omitempty"`
	SEDS_NCU   float64 `json:"SEDS_NCU,omitempty"`
	SEDS_BCL   float64 `json:"SEDS_BCL,omitempty"`
	SEDS_BCU   float64 `json:"SEDS_BCU,omitempty"`
	EDI        float64 `json:"EDI,omitempty"`
	EDI_NCL    float64 `json:"EDI_NCL,omitempty"`
	EDI_NCU    float64 `json:"EDI_NCU,omitempty"`
	EDI_BCL    float64 `json:"EDI_BCL,omitempty"`
	EDI_BCU    float64 `json:"EDI_BCU,omitempty"`
	SEDI       float64 `json:"SEDI,omitempty"`
	SEDI_NCL   float64 `json:"SEDI_NCL,omitempty"`
	SEDI_NCU   float64 `json:"SEDI_NCU,omitempty"`
	SEDI_BCL   float64 `json:"SEDI_BCL,omitempty"`
	SEDI_BCU   float64 `json:"SEDI_BCU,omitempty"`
	BAGSS      float64 `json:"BAGSS,omitempty"`
	BAGSS_BCL  float64 `json:"BAGSS_BCL,omitempty"`
	BAGSS_BCU  float64 `json:"BAGSS_BCU,omitempty"`
	HSS_EC     float64 `json:"HSS_EC,omitempty"`
	HSS_EC_BCL float64 `json:"HSS_EC_BCL,omitempty"`
	HSS_EC_BCU float64 `json:"HSS_EC_BCU,omitempty"`
	EC_VALUE   float64 `json:"EC_VALUE,omitempty"`
}

type STAT_DMAP struct {
	TOTAL      int     `json:"TOTAL,omitempty"`
	FY         int     `json:"FY,omitempty"`
	OY         int     `json:"OY,omitempty"`
	FBIAS      float64 `json:"FBIAS,omitempty"`
	BADDELEY   float64 `json:"BADDELEY,omitempty"`
	HAUSDORFF  float64 `json:"HAUSDORFF,omitempty"`
	MED_FO     float64 `json:"MED_FO,omitempty"`
	MED_OF     float64 `json:"MED_OF,omitempty"`
	MED_MIN    float64 `json:"MED_MIN,omitempty"`
	MED_MAX    float64 `json:"MED_MAX,omitempty"`
	MED_MEAN   float64 `json:"MED_MEAN,omitempty"`
	FOM_FO     float64 `json:"FOM_FO,omitempty"`
	FOM_OF     float64 `json:"FOM_OF,omitempty"`
	FOM_MIN    float64 `json:"FOM_MIN,omitempty"`
	FOM_MAX    float64 `json:"FOM_MAX,omitempty"`
	FOM_MEAN   float64 `json:"FOM_MEAN,omitempty"`
	ZHU_FO     float64 `json:"ZHU_FO,omitempty"`
	ZHU_OF     float64 `json:"ZHU_OF,omitempty"`
	ZHU_MIN    float64 `json:"ZHU_MIN,omitempty"`
	ZHU_MAX    float64 `json:"ZHU_MAX,omitempty"`
	ZHU_MEAN   float64 `json:"ZHU_MEAN,omitempty"`
	G          float64 `json:"G,omitempty"`
	GBETA      float64 `json:"GBETA,omitempty"`
	BETA_VALUE float64 `json:"BETA_VALUE,omitempty"`
}

type STAT_ECLV struct {
	TOTAL       int                    `json:"TOTAL,omitempty"`
	BASER       float64                `json:"BASER,omitempty"`
	VALUE_BASER int                    `json:"VALUE_BASER,omitempty"`
	PTS         map[string]interface{} `json:"PTS,omitempty"`
}

type STAT_ECNT struct {
	TOTAL            int     `json:"TOTAL,omitempty"`
	N_ENS            int     `json:"N_ENS,omitempty"`
	CRPS             float64 `json:"CRPS,omitempty"`
	CRPSS            float64 `json:"CRPSS,omitempty"`
	IGN              float64 `json:"IGN,omitempty"`
	ME               float64 `json:"ME,omitempty"`
	RMSE             float64 `json:"RMSE,omitempty"`
	SPREAD           float64 `json:"SPREAD,omitempty"`
	ME_OERR          float64 `json:"ME_OERR,omitempty"`
	RMSE_OERR        float64 `json:"RMSE_OERR,omitempty"`
	SPREAD_OERR      float64 `json:"SPREAD_OERR,omitempty"`
	SPREAD_PLUS_OERR float64 `json:"SPREAD_PLUS_OERR,omitempty"`
	CRPSCL           float64 `json:"CRPSCL,omitempty"`
	CRPS_EMP         float64 `json:"CRPS_EMP,omitempty"`
	CRPSCL_EMP       float64 `json:"CRPSCL_EMP,omitempty"`
	CRPSS_EMP        float64 `json:"CRPSS_EMP,omitempty"`
	CRPS_EMP_FAIR    float64 `json:"CRPS_EMP_FAIR,omitempty"`
	SPREAD_MD        float64 `json:"SPREAD_MD,omitempty"`
	MAE              float64 `json:"MAE,omitempty"`
	MAE_OERR         float64 `json:"MAE_OERR,omitempty"`
	BIAS_RATIO       float64 `json:"BIAS_RATIO,omitempty"`
	N_GE_OBS         int     `json:"N_GE_OBS,omitempty"`
	ME_GE_OBS        float64 `json:"ME_GE_OBS,omitempty"`
	N_LT_OBS         int     `json:"N_LT_OBS,omitempty"`
	ME_LT_OBS        float64 `json:"ME_LT_OBS,omitempty"`
}

type STAT_FHO struct {
	TOTAL  int     `json:"TOTAL,omitempty"`
	F_RATE float64 `json:"F_RATE,omitempty"`
	H_RATE float64 `json:"H_RATE,omitempty"`
	O_RATE float64 `json:"O_RATE,omitempty"`
}

type STAT_GENMPR struct {
	TOTAL      int     `json:"TOTAL,omitempty"`
	INDEX      int     `json:"INDEX,omitempty"`
	STORM_ID   string  `json:"STORM_ID,omitempty"`
	PROB_LEAD  float64 `json:"PROB_LEAD,omitempty"`
	PROB_VAL   float64 `json:"PROB_VAL,omitempty"`
	AGEN_INIT  string  `json:"AGEN_INIT,omitempty"`
	AGEN_FHR   string  `json:"AGEN_FHR,omitempty"`
	AGEN_LAT   float64 `json:"AGEN_LAT,omitempty"`
	AGEN_LON   float64 `json:"AGEN_LON,omitempty"`
	AGEN_DLAND float64 `json:"AGEN_DLAND,omitempty"`
	BGEN_LAT   float64 `json:"BGEN_LAT,omitempty"`
	BGEN_LON   float64 `json:"BGEN_LON,omitempty"`
	BGEN_DLAND float64 `json:"BGEN_DLAND,omitempty"`
	GEN_DIST   float64 `json:"GEN_DIST,omitempty"`
	GEN_TDIFF  string  `json:"GEN_TDIFF,omitempty"`
	INIT_TDIFF string  `json:"INIT_TDIFF,omitempty"`
	DEV_CAT    string  `json:"DEV_CAT,omitempty"`
	OPS_CAT    string  `json:"OPS_CAT,omitempty"`
}

type STAT_GRAD struct {
	TOTAL      int     `json:"TOTAL,omitempty"`
	FGBAR      float64 `json:"FGBAR,omitempty"`
	OGBAR      float64 `json:"OGBAR,omitempty"`
	MGBAR      float64 `json:"MGBAR,omitempty"`
	EGBAR      float64 `json:"EGBAR,omitempty"`
	S1         float64 `json:"S1,omitempty"`
	S1_OG      float64 `json:"S1_OG,omitempty"`
	FGOG_RATIO float64 `json:"FGOG_RATIO,omitempty"`
	DX         float64 `json:"DX,omitempty"`
	DY         float64 `json:"DY,omitempty"`
}

type STAT_ISC struct {
	TOTAL    int     `json:"TOTAL,omitempty"`
	TILE_DIM int     `json:"TILE_DIM,omitempty"`
	TILE_XLL int     `json:"TILE_XLL,omitempty"`
	TILE_YLL int     `json:"TILE_YLL,omitempty"`
	NSCALE   int     `json:"NSCALE,omitempty"`
	ISCALE   int     `json:"ISCALE,omitempty"`
	MSE      float64 `json:"MSE,omitempty"`
	ISC      float64 `json:"ISC,omitempty"`
	FENERGY2 float64 `json:"FENERGY2,omitempty"`
	OENERGY2 float64 `json:"OENERGY2,omitempty"`
	BASER    float64 `json:"BASER,omitempty"`
	FBIAS    float64 `json:"FBIAS,omitempty"`
}

type STAT_MCTC struct {
	TOTAL    int                    `json:"TOTAL,omitempty"`
	CAT      map[string]interface{} `json:"CAT,omitempty"`
	EC_VALUE float64                `json:"EC_VALUE,omitempty"`
}

type STAT_MCTS struct {
	TOTAL      int     `json:"TOTAL,omitempty"`
	N_CAT      int     `json:"N_CAT,omitempty"`
	ACC        float64 `json:"ACC,omitempty"`
	ACC_NCL    float64 `json:"ACC_NCL,omitempty"`
	ACC_NCU    float64 `json:"ACC_NCU,omitempty"`
	ACC_BCL    float64 `json:"ACC_BCL,omitempty"`
	ACC_BCU    float64 `json:"ACC_BCU,omitempty"`
	HK         float64 `json:"HK,omitempty"`
	HK_BCL     float64 `json:"HK_BCL,omitempty"`
	HK_BCU     float64 `json:"HK_BCU,omitempty"`
	HSS        float64 `json:"HSS,omitempty"`
	HSS_BCL    float64 `json:"HSS_BCL,omitempty"`
	HSS_BCU    float64 `json:"HSS_BCU,omitempty"`
	GER        float64 `json:"GER,omitempty"`
	GER_BCL    float64 `json:"GER_BCL,omitempty"`
	GER_BCU    float64 `json:"GER_BCU,omitempty"`
	HSS_EC     float64 `json:"HSS_EC,omitempty"`
	HSS_EC_BCL float64 `json:"HSS_EC_BCL,omitempty"`
	HSS_EC_BCU float64 `json:"HSS_EC_BCU,omitempty"`
	EC_VALUE   float64 `json:"EC_VALUE,omitempty"`
}

type STAT_MPR struct {
	TOTAL       int     `json:"TOTAL,omitempty"`
	INDEX       int     `json:"INDEX,omitempty"`
	OBS_SID     string  `json:"OBS_SID,omitempty"`
	OBS_LAT     float64 `json:"OBS_LAT,omitempty"`
	OBS_LON     float64 `json:"OBS_LON,omitempty"`
	OBS_LVL     float64 `json:"OBS_LVL,omitempty"`
	OBS_ELV     float64 `json:"OBS_ELV,omitempty"`
	FCST        float64 `json:"FCST,omitempty"`
	OBS         float64 `json:"OBS,omitempty"`
	OBS_QC      string  `json:"OBS_QC,omitempty"`
	CLIMO_MEAN  float64 `json:"CLIMO_MEAN,omitempty"`
	CLIMO_STDEV float64 `json:"CLIMO_STDEV,omitempty"`
	CLIMO_CDF   float64 `json:"CLIMO_CDF,omitempty"`
}

type STAT_NBRCNT struct {
	TOTAL      int     `json:"TOTAL,omitempty"`
	FBS        float64 `json:"FBS,omitempty"`
	FBS_BCL    float64 `json:"FBS_BCL,omitempty"`
	FBS_BCU    float64 `json:"FBS_BCU,omitempty"`
	FSS        float64 `json:"FSS,omitempty"`
	FSS_BCL    float64 `json:"FSS_BCL,omitempty"`
	FSS_BCU    float64 `json:"FSS_BCU,omitempty"`
	AFSS       float64 `json:"AFSS,omitempty"`
	AFSS_BCL   float64 `json:"AFSS_BCL,omitempty"`
	AFSS_BCU   float64 `json:"AFSS_BCU,omitempty"`
	UFSS       float64 `json:"UFSS,omitempty"`
	UFSS_BCL   float64 `json:"UFSS_BCL,omitempty"`
	UFSS_BCU   float64 `json:"UFSS_BCU,omitempty"`
	F_RATE     float64 `json:"F_RATE,omitempty"`
	F_RATE_BCL float64 `json:"F_RATE_BCL,omitempty"`
	F_RATE_BCU float64 `json:"F_RATE_BCU,omitempty"`
	O_RATE     float64 `json:"O_RATE,omitempty"`
	O_RATE_BCL float64 `json:"O_RATE_BCL,omitempty"`
	O_RATE_BCU float64 `json:"O_RATE_BCU,omitempty"`
}

type STAT_NBRCTC struct {
	TOTAL int     `json:"TOTAL,omitempty"`
	FY_OY float64 `json:"FY_OY,omitempty"`
	FY_ON float64 `json:"FY_ON,omitempty"`
	FN_OY float64 `json:"FN_OY,omitempty"`
	FN_ON float64 `json:"FN_ON,omitempty"`
}

type STAT_NBRCTS struct {
	TOTAL     int     `json:"TOTAL,omitempty"`
	BASER     float64 `json:"BASER,omitempty"`
	BASER_NCL float64 `json:"BASER_NCL,omitempty"`
	BASER_NCU float64 `json:"BASER_NCU,omitempty"`
	BASER_BCL float64 `json:"BASER_BCL,omitempty"`
	BASER_BCU float64 `json:"BASER_BCU,omitempty"`
	FMEAN     float64 `json:"FMEAN,omitempty"`
	FMEAN_NCL float64 `json:"FMEAN_NCL,omitempty"`
	FMEAN_NCU float64 `json:"FMEAN_NCU,omitempty"`
	FMEAN_BCL float64 `json:"FMEAN_BCL,omitempty"`
	FMEAN_BCU float64 `json:"FMEAN_BCU,omitempty"`
	ACC       float64 `json:"ACC,omitempty"`
	ACC_NCL   float64 `json:"ACC_NCL,omitempty"`
	ACC_NCU   float64 `json:"ACC_NCU,omitempty"`
	ACC_BCL   float64 `json:"ACC_BCL,omitempty"`
	ACC_BCU   float64 `json:"ACC_BCU,omitempty"`
	FBIAS     float64 `json:"FBIAS,omitempty"`
	FBIAS_BCL float64 `json:"FBIAS_BCL,omitempty"`
	FBIAS_BCU float64 `json:"FBIAS_BCU,omitempty"`
	PODY      float64 `json:"PODY,omitempty"`
	PODY_NCL  float64 `json:"PODY_NCL,omitempty"`
	PODY_NCU  float64 `json:"PODY_NCU,omitempty"`
	PODY_BCL  float64 `json:"PODY_BCL,omitempty"`
	PODY_BCU  float64 `json:"PODY_BCU,omitempty"`
	PODN      float64 `json:"PODN,omitempty"`
	PODN_NCL  float64 `json:"PODN_NCL,omitempty"`
	PODN_NCU  float64 `json:"PODN_NCU,omitempty"`
	PODN_BCL  float64 `json:"PODN_BCL,omitempty"`
	PODN_BCU  float64 `json:"PODN_BCU,omitempty"`
	POFD      float64 `json:"POFD,omitempty"`
	POFD_NCL  float64 `json:"POFD_NCL,omitempty"`
	POFD_NCU  float64 `json:"POFD_NCU,omitempty"`
	POFD_BCL  float64 `json:"POFD_BCL,omitempty"`
	POFD_BCU  float64 `json:"POFD_BCU,omitempty"`
	FAR       float64 `json:"FAR,omitempty"`
	FAR_NCL   float64 `json:"FAR_NCL,omitempty"`
	FAR_NCU   float64 `json:"FAR_NCU,omitempty"`
	FAR_BCL   float64 `json:"FAR_BCL,omitempty"`
	FAR_BCU   float64 `json:"FAR_BCU,omitempty"`
	CSI       float64 `json:"CSI,omitempty"`
	CSI_NCL   float64 `json:"CSI_NCL,omitempty"`
	CSI_NCU   float64 `json:"CSI_NCU,omitempty"`
	CSI_BCL   float64 `json:"CSI_BCL,omitempty"`
	CSI_BCU   float64 `json:"CSI_BCU,omitempty"`
	GSS       float64 `json:"GSS,omitempty"`
	GSS_BCL   float64 `json:"GSS_BCL,omitempty"`
	GSS_BCU   float64 `json:"GSS_BCU,omitempty"`
	HK        float64 `json:"HK,omitempty"`
	HK_NCL    float64 `json:"HK_NCL,omitempty"`
	HK_NCU    float64 `json:"HK_NCU,omitempty"`
	HK_BCL    float64 `json:"HK_BCL,omitempty"`
	HK_BCU    float64 `json:"HK_BCU,omitempty"`
	HSS       float64 `json:"HSS,omitempty"`
	HSS_BCL   float64 `json:"HSS_BCL,omitempty"`
	HSS_BCU   float64 `json:"HSS_BCU,omitempty"`
	ODDS      float64 `json:"ODDS,omitempty"`
	ODDS_NCL  float64 `json:"ODDS_NCL,omitempty"`
	ODDS_NCU  float64 `json:"ODDS_NCU,omitempty"`
	ODDS_BCL  float64 `json:"ODDS_BCL,omitempty"`
	ODDS_BCU  float64 `json:"ODDS_BCU,omitempty"`
	LODDS     float64 `json:"LODDS,omitempty"`
	LODDS_NCL float64 `json:"LODDS_NCL,omitempty"`
	LODDS_NCU float64 `json:"LODDS_NCU,omitempty"`
	LODDS_BCL float64 `json:"LODDS_BCL,omitempty"`
	LODDS_BCU float64 `json:"LODDS_BCU,omitempty"`
	ORSS      float64 `json:"ORSS,omitempty"`
	ORSS_NCL  float64 `json:"ORSS_NCL,omitempty"`
	ORSS_NCU  float64 `json:"ORSS_NCU,omitempty"`
	ORSS_BCL  float64 `json:"ORSS_BCL,omitempty"`
	ORSS_BCU  float64 `json:"ORSS_BCU,omitempty"`
	EDS       float64 `json:"EDS,omitempty"`
	EDS_NCL   float64 `json:"EDS_NCL,omitempty"`
	EDS_NCU   float64 `json:"EDS_NCU,omitempty"`
	EDS_BCL   float64 `json:"EDS_BCL,omitempty"`
	EDS_BCU   float64 `json:"EDS_BCU,omitempty"`
	SEDS      float64 `json:"SEDS,omitempty"`
	SEDS_NCL  float64 `json:"SEDS_NCL,omitempty"`
	SEDS_NCU  float64 `json:"SEDS_NCU,omitempty"`
	SEDS_BCL  float64 `json:"SEDS_BCL,omitempty"`
	SEDS_BCU  float64 `json:"SEDS_BCU,omitempty"`
	EDI       float64 `json:"EDI,omitempty"`
	EDI_NCL   float64 `json:"EDI_NCL,omitempty"`
	EDI_NCU   float64 `json:"EDI_NCU,omitempty"`
	EDI_BCL   float64 `json:"EDI_BCL,omitempty"`
	EDI_BCU   float64 `json:"EDI_BCU,omitempty"`
	SEDI      float64 `json:"SEDI,omitempty"`
	SEDI_NCL  float64 `json:"SEDI_NCL,omitempty"`
	SEDI_NCU  float64 `json:"SEDI_NCU,omitempty"`
	SEDI_BCL  float64 `json:"SEDI_BCL,omitempty"`
	SEDI_BCU  float64 `json:"SEDI_BCU,omitempty"`
	BAGSS     float64 `json:"BAGSS,omitempty"`
	BAGSS_BCL float64 `json:"BAGSS_BCL,omitempty"`
	BAGSS_BCU float64 `json:"BAGSS_BCU,omitempty"`
}

type STAT_ORANK struct {
	TOTAL            int                    `json:"TOTAL,omitempty"`
	INDEX            int                    `json:"INDEX,omitempty"`
	OBS_SID          string                 `json:"OBS_SID,omitempty"`
	OBS_LAT          float64                `json:"OBS_LAT,omitempty"`
	OBS_LON          float64                `json:"OBS_LON,omitempty"`
	OBS_LVL          float64                `json:"OBS_LVL,omitempty"`
	OBS_ELV          float64                `json:"OBS_ELV,omitempty"`
	OBS              float64                `json:"OBS,omitempty"`
	PIT              float64                `json:"PIT,omitempty"`
	RANK             int                    `json:"RANK,omitempty"`
	N_ENS_VLD        int                    `json:"N_ENS_VLD,omitempty"`
	ENS              map[string]interface{} `json:"ENS,omitempty"`
	OBS_QC           string                 `json:"OBS_QC,omitempty"`
	ENS_MEAN         int                    `json:"ENS_MEAN,omitempty"`
	CLIMO_MEAN       float64                `json:"CLIMO_MEAN,omitempty"`
	SPREAD           float64                `json:"SPREAD,omitempty"`
	ENS_MEAN_OERR    int                    `json:"ENS_MEAN_OERR,omitempty"`
	SPREAD_OERR      float64                `json:"SPREAD_OERR,omitempty"`
	SPREAD_PLUS_OERR float64                `json:"SPREAD_PLUS_OERR,omitempty"`
	CLIMO_STDEV      float64                `json:"CLIMO_STDEV,omitempty"`
}

type STAT_PCT struct {
	TOTAL  int                    `json:"TOTAL,omitempty"`
	THRESH map[string]interface{} `json:"THRESH,omitempty"`
}

type STAT_PHIST struct {
	TOTAL    int                    `json:"TOTAL,omitempty"`
	BIN_SIZE int                    `json:"BIN_SIZE,omitempty"`
	BIN      map[string]interface{} `json:"BIN,omitempty"`
}

type STAT_PJC struct {
	TOTAL  int                    `json:"TOTAL,omitempty"`
	THRESH map[string]interface{} `json:"THRESH,omitempty"`
}

type STAT_PRC struct {
	TOTAL  int                    `json:"TOTAL,omitempty"`
	THRESH map[string]interface{} `json:"THRESH,omitempty"`
}

type STAT_PSTD struct {
	TOTAL       int                    `json:"TOTAL,omitempty"`
	THRESH      map[string]interface{} `json:"THRESH,omitempty"`
	BASER_NCL   float64                `json:"BASER_NCL,omitempty"`
	BASER_NCU   float64                `json:"BASER_NCU,omitempty"`
	RELIABILITY float64                `json:"RELIABILITY,omitempty"`
	RESOLUTION  float64                `json:"RESOLUTION,omitempty"`
	UNCERTAINTY float64                `json:"UNCERTAINTY,omitempty"`
	ROC_AUC     float64                `json:"ROC_AUC,omitempty"`
	BRIER       float64                `json:"BRIER,omitempty"`
	BRIER_NCL   float64                `json:"BRIER_NCL,omitempty"`
	BRIER_NCU   float64                `json:"BRIER_NCU,omitempty"`
	BRIERCL     float64                `json:"BRIERCL,omitempty"`
	BRIERCL_NCL float64                `json:"BRIERCL_NCL,omitempty"`
	BRIERCL_NCU float64                `json:"BRIERCL_NCU,omitempty"`
	BSS         float64                `json:"BSS,omitempty"`
	BSS_SMPL    float64                `json:"BSS_SMPL,omitempty"`
	THRESH_I    int                    `json:"THRESH_I,omitempty"`
}

type STAT_RELP struct {
	TOTAL int                    `json:"TOTAL,omitempty"`
	ENS   map[string]interface{} `json:"ENS,omitempty"`
}

type STAT_RHIST struct {
	TOTAL int                    `json:"TOTAL,omitempty"`
	RANK  map[string]interface{} `json:"RANK,omitempty"`
}

type STAT_RPS struct {
	TOTAL     int     `json:"TOTAL,omitempty"`
	N_PROB    int     `json:"N_PROB,omitempty"`
	RPS_REL   float64 `json:"RPS_REL,omitempty"`
	RPS_RES   float64 `json:"RPS_RES,omitempty"`
	RPS_UNC   float64 `json:"RPS_UNC,omitempty"`
	RPS       float64 `json:"RPS,omitempty"`
	RPSS      float64 `json:"RPSS,omitempty"`
	RPSS_SMPL float64 `json:"RPSS_SMPL,omitempty"`
	RPS_COMP  float64 `json:"RPS_COMP,omitempty"`
}

type STAT_SAL1L2 struct {
	TOTAL  int     `json:"TOTAL,omitempty"`
	FABAR  float64 `json:"FABAR,omitempty"`
	OABAR  float64 `json:"OABAR,omitempty"`
	FOABAR float64 `json:"FOABAR,omitempty"`
	FFABAR float64 `json:"FFABAR,omitempty"`
	OOABAR float64 `json:"OOABAR,omitempty"`
	MAE    float64 `json:"MAE,omitempty"`
}

type STAT_SEEPS struct {
	TOTAL     int     `json:"TOTAL,omitempty"`
	S12       string  `json:"S12,omitempty"`
	S13       string  `json:"S13,omitempty"`
	S21       string  `json:"S21,omitempty"`
	S23       string  `json:"S23,omitempty"`
	S31       string  `json:"S31,omitempty"`
	S32       string  `json:"S32,omitempty"`
	PF1       float64 `json:"PF1,omitempty"`
	PF2       float64 `json:"PF2,omitempty"`
	PF3       float64 `json:"PF3,omitempty"`
	PV1       float64 `json:"PV1,omitempty"`
	PV2       float64 `json:"PV2,omitempty"`
	PV3       float64 `json:"PV3,omitempty"`
	MEAN_FCST float64 `json:"MEAN_FCST,omitempty"`
	MEAN_OBS  float64 `json:"MEAN_OBS,omitempty"`
	SEEPS     float64 `json:"SEEPS,omitempty"`
}

type STAT_SEEPS_MPR struct {
	OBS_SID  string  `json:"OBS_SID,omitempty"`
	OBS_LAT  float64 `json:"OBS_LAT,omitempty"`
	OBS_LON  float64 `json:"OBS_LON,omitempty"`
	FCST     float64 `json:"FCST,omitempty"`
	OBS      float64 `json:"OBS,omitempty"`
	OBS_QC   string  `json:"OBS_QC,omitempty"`
	FCST_CAT int     `json:"FCST_CAT,omitempty"`
	OBS_CAT  int     `json:"OBS_CAT,omitempty"`
	P1       float64 `json:"P1,omitempty"`
	P2       float64 `json:"P2,omitempty"`
	T1       float64 `json:"T1,omitempty"`
	T2       float64 `json:"T2,omitempty"`
	SEEPS    float64 `json:"SEEPS,omitempty"`
}

type STAT_SL1L2 struct {
	TOTAL int     `json:"TOTAL,omitempty"`
	FBAR  float64 `json:"FBAR,omitempty"`
	OBAR  float64 `json:"OBAR,omitempty"`
	FOBAR float64 `json:"FOBAR,omitempty"`
	FFBAR float64 `json:"FFBAR,omitempty"`
	OOBAR float64 `json:"OOBAR,omitempty"`
	MAE   float64 `json:"MAE,omitempty"`
}

type STAT_SSIDX struct {
	FCST_MODEL string  `json:"FCST_MODEL,omitempty"`
	REF_MODEL  string  `json:"REF_MODEL,omitempty"`
	N_INIT     int     `json:"N_INIT,omitempty"`
	N_TERM     int     `json:"N_TERM,omitempty"`
	N_VLD      int     `json:"N_VLD,omitempty"`
	SS_INDEX   float64 `json:"SS_INDEX,omitempty"`
}

type STAT_SSVAR struct {
	TOTAL       int     `json:"TOTAL,omitempty"`
	N_BIN       int     `json:"N_BIN,omitempty"`
	BIN_I       int     `json:"BIN_I,omitempty"`
	BIN_N       int     `json:"BIN_N,omitempty"`
	VAR_MIN     float64 `json:"VAR_MIN,omitempty"`
	VAR_MAX     float64 `json:"VAR_MAX,omitempty"`
	VAR_MEAN    float64 `json:"VAR_MEAN,omitempty"`
	FBAR        float64 `json:"FBAR,omitempty"`
	OBAR        float64 `json:"OBAR,omitempty"`
	FOBAR       float64 `json:"FOBAR,omitempty"`
	FFBAR       float64 `json:"FFBAR,omitempty"`
	OOBAR       float64 `json:"OOBAR,omitempty"`
	FBAR_NCL    float64 `json:"FBAR_NCL,omitempty"`
	FBAR_NCU    float64 `json:"FBAR_NCU,omitempty"`
	FSTDEV      float64 `json:"FSTDEV,omitempty"`
	FSTDEV_NCL  float64 `json:"FSTDEV_NCL,omitempty"`
	FSTDEV_NCU  float64 `json:"FSTDEV_NCU,omitempty"`
	OBAR_NCL    float64 `json:"OBAR_NCL,omitempty"`
	OBAR_NCU    float64 `json:"OBAR_NCU,omitempty"`
	OSTDEV      float64 `json:"OSTDEV,omitempty"`
	OSTDEV_NCL  float64 `json:"OSTDEV_NCL,omitempty"`
	OSTDEV_NCU  float64 `json:"OSTDEV_NCU,omitempty"`
	PR_CORR     float64 `json:"PR_CORR,omitempty"`
	PR_CORR_NCL float64 `json:"PR_CORR_NCL,omitempty"`
	PR_CORR_NCU float64 `json:"PR_CORR_NCU,omitempty"`
	ME          float64 `json:"ME,omitempty"`
	ME_NCL      float64 `json:"ME_NCL,omitempty"`
	ME_NCU      float64 `json:"ME_NCU,omitempty"`
	ESTDEV      float64 `json:"ESTDEV,omitempty"`
	ESTDEV_NCL  float64 `json:"ESTDEV_NCL,omitempty"`
	ESTDEV_NCU  float64 `json:"ESTDEV_NCU,omitempty"`
	MBIAS       float64 `json:"MBIAS,omitempty"`
	MSE         float64 `json:"MSE,omitempty"`
	BCMSE       float64 `json:"BCMSE,omitempty"`
	RMSE        float64 `json:"RMSE,omitempty"`
}

type STAT_VAL1L2 struct {
	TOTAL        int     `json:"TOTAL,omitempty"`
	UFABAR       float64 `json:"UFABAR,omitempty"`
	VFABAR       float64 `json:"VFABAR,omitempty"`
	UOABAR       float64 `json:"UOABAR,omitempty"`
	VOABAR       float64 `json:"VOABAR,omitempty"`
	UVFOABAR     float64 `json:"UVFOABAR,omitempty"`
	UVFFABAR     float64 `json:"UVFFABAR,omitempty"`
	UVOOABAR     float64 `json:"UVOOABAR,omitempty"`
	FA_SPEED_BAR float64 `json:"FA_SPEED_BAR,omitempty"`
	OA_SPEED_BAR float64 `json:"OA_SPEED_BAR,omitempty"`
}

type STAT_VCNT struct {
	TOTAL                int     `json:"TOTAL,omitempty"`
	FBAR                 float64 `json:"FBAR,omitempty"`
	FBAR_BCL             float64 `json:"FBAR_BCL,omitempty"`
	FBAR_BCU             float64 `json:"FBAR_BCU,omitempty"`
	OBAR                 float64 `json:"OBAR,omitempty"`
	OBAR_BCL             float64 `json:"OBAR_BCL,omitempty"`
	OBAR_BCU             float64 `json:"OBAR_BCU,omitempty"`
	FS_RMS               float64 `json:"FS_RMS,omitempty"`
	FS_RMS_BCL           float64 `json:"FS_RMS_BCL,omitempty"`
	FS_RMS_BCU           float64 `json:"FS_RMS_BCU,omitempty"`
	OS_RMS               float64 `json:"OS_RMS,omitempty"`
	OS_RMS_BCL           float64 `json:"OS_RMS_BCL,omitempty"`
	OS_RMS_BCU           float64 `json:"OS_RMS_BCU,omitempty"`
	MSVE                 float64 `json:"MSVE,omitempty"`
	MSVE_BCL             float64 `json:"MSVE_BCL,omitempty"`
	MSVE_BCU             float64 `json:"MSVE_BCU,omitempty"`
	RMSVE                float64 `json:"RMSVE,omitempty"`
	RMSVE_BCL            float64 `json:"RMSVE_BCL,omitempty"`
	RMSVE_BCU            float64 `json:"RMSVE_BCU,omitempty"`
	FSTDEV               float64 `json:"FSTDEV,omitempty"`
	FSTDEV_BCL           float64 `json:"FSTDEV_BCL,omitempty"`
	FSTDEV_BCU           float64 `json:"FSTDEV_BCU,omitempty"`
	OSTDEV               float64 `json:"OSTDEV,omitempty"`
	OSTDEV_BCL           float64 `json:"OSTDEV_BCL,omitempty"`
	OSTDEV_BCU           float64 `json:"OSTDEV_BCU,omitempty"`
	FDIR                 float64 `json:"FDIR,omitempty"`
	FDIR_BCL             float64 `json:"FDIR_BCL,omitempty"`
	FDIR_BCU             float64 `json:"FDIR_BCU,omitempty"`
	ODIR                 float64 `json:"ODIR,omitempty"`
	ODIR_BCL             float64 `json:"ODIR_BCL,omitempty"`
	ODIR_BCU             float64 `json:"ODIR_BCU,omitempty"`
	FBAR_SPEED           float64 `json:"FBAR_SPEED,omitempty"`
	FBAR_SPEED_BCL       float64 `json:"FBAR_SPEED_BCL,omitempty"`
	FBAR_SPEED_BCU       float64 `json:"FBAR_SPEED_BCU,omitempty"`
	OBAR_SPEED           float64 `json:"OBAR_SPEED,omitempty"`
	OBAR_SPEED_BCL       float64 `json:"OBAR_SPEED_BCL,omitempty"`
	OBAR_SPEED_BCU       float64 `json:"OBAR_SPEED_BCU,omitempty"`
	VDIFF_SPEED          float64 `json:"VDIFF_SPEED,omitempty"`
	VDIFF_SPEED_BCL      float64 `json:"VDIFF_SPEED_BCL,omitempty"`
	VDIFF_SPEED_BCU      float64 `json:"VDIFF_SPEED_BCU,omitempty"`
	VDIFF_DIR            float64 `json:"VDIFF_DIR,omitempty"`
	VDIFF_DIR_BCL        float64 `json:"VDIFF_DIR_BCL,omitempty"`
	VDIFF_DIR_BCU        float64 `json:"VDIFF_DIR_BCU,omitempty"`
	SPEED_ERR            float64 `json:"SPEED_ERR,omitempty"`
	SPEED_ERR_BCL        float64 `json:"SPEED_ERR_BCL,omitempty"`
	SPEED_ERR_BCU        float64 `json:"SPEED_ERR_BCU,omitempty"`
	SPEED_ABSERR         float64 `json:"SPEED_ABSERR,omitempty"`
	SPEED_ABSERR_BCL     float64 `json:"SPEED_ABSERR_BCL,omitempty"`
	SPEED_ABSERR_BCU     float64 `json:"SPEED_ABSERR_BCU,omitempty"`
	DIR_ERR              float64 `json:"DIR_ERR,omitempty"`
	DIR_ERR_BCL          float64 `json:"DIR_ERR_BCL,omitempty"`
	DIR_ERR_BCU          float64 `json:"DIR_ERR_BCU,omitempty"`
	DIR_ABSERR           float64 `json:"DIR_ABSERR,omitempty"`
	DIR_ABSERR_BCL       float64 `json:"DIR_ABSERR_BCL,omitempty"`
	DIR_ABSERR_BCU       float64 `json:"DIR_ABSERR_BCU,omitempty"`
	ANOM_CORR            float64 `json:"ANOM_CORR,omitempty"`
	ANOM_CORR_NCL        float64 `json:"ANOM_CORR_NCL,omitempty"`
	ANOM_CORR_NCU        float64 `json:"ANOM_CORR_NCU,omitempty"`
	ANOM_CORR_BCL        float64 `json:"ANOM_CORR_BCL,omitempty"`
	ANOM_CORR_BCU        float64 `json:"ANOM_CORR_BCU,omitempty"`
	ANOM_CORR_UNCNTR     float64 `json:"ANOM_CORR_UNCNTR,omitempty"`
	ANOM_CORR_UNCNTR_BCL float64 `json:"ANOM_CORR_UNCNTR_BCL,omitempty"`
	ANOM_CORR_UNCNTR_BCU float64 `json:"ANOM_CORR_UNCNTR_BCU,omitempty"`
}

type STAT_VL1L2 struct {
	TOTAL       int     `json:"TOTAL,omitempty"`
	UFBAR       float64 `json:"UFBAR,omitempty"`
	VFBAR       float64 `json:"VFBAR,omitempty"`
	UOBAR       float64 `json:"UOBAR,omitempty"`
	VOBAR       float64 `json:"VOBAR,omitempty"`
	UVFOBAR     float64 `json:"UVFOBAR,omitempty"`
	UVFFBAR     float64 `json:"UVFFBAR,omitempty"`
	UVOOBAR     float64 `json:"UVOOBAR,omitempty"`
	F_SPEED_BAR float64 `json:"F_SPEED_BAR,omitempty"`
	O_SPEED_BAR float64 `json:"O_SPEED_BAR,omitempty"`
}

type TCST_PROBRIRW struct {
	ALAT        float64                `json:"ALAT,omitempty"`
	ALON        float64                `json:"ALON,omitempty"`
	BLAT        float64                `json:"BLAT,omitempty"`
	BLON        float64                `json:"BLON,omitempty"`
	INITIALS    string                 `json:"INITIALS,omitempty"`
	TK_ERR      float64                `json:"TK_ERR,omitempty"`
	X_ERR       float64                `json:"X_ERR,omitempty"`
	Y_ERR       float64                `json:"Y_ERR,omitempty"`
	ADLAND      float64                `json:"ADLAND,omitempty"`
	BDLAND      float64                `json:"BDLAND,omitempty"`
	RIRW_BEG    int                    `json:"RIRW_BEG,omitempty"`
	RIRW_END    int                    `json:"RIRW_END,omitempty"`
	RIRW_WINDOW int                    `json:"RIRW_WINDOW,omitempty"`
	AWIND_END   float64                `json:"AWIND_END,omitempty"`
	BWIND_BEG   float64                `json:"BWIND_BEG,omitempty"`
	BWIND_END   float64                `json:"BWIND_END,omitempty"`
	BDELTA      float64                `json:"BDELTA,omitempty"`
	BDELTA_MAX  float64                `json:"BDELTA_MAX,omitempty"`
	BLEVEL_BEG  string                 `json:"BLEVEL_BEG,omitempty"`
	BLEVEL_END  string                 `json:"BLEVEL_END,omitempty"`
	THRESH      map[string]interface{} `json:"THRESH,omitempty"`
	INIT        int                    `json:"INIT,omitempty"`
}

type TCST_TCDIAG struct {
	TOTAL        int                    `json:"TOTAL,omitempty"`
	INDEX        int                    `json:"INDEX,omitempty"`
	DIAG_SOURCE  float64                `json:"DIAG_SOURCE,omitempty"`
	TRACK_SOURCE string                 `json:"TRACK_SOURCE,omitempty"`
	FIELD_SOURCE string                 `json:"FIELD_SOURCE,omitempty"`
	DIAG         map[string]interface{} `json:"DIAG,omitempty"`
	INIT         int                    `json:"INIT,omitempty"`
}

type TCST_TCMPR struct {
	TOTAL          int     `json:"TOTAL,omitempty"`
	INDEX          int     `json:"INDEX,omitempty"`
	LEVEL          string  `json:"LEVEL,omitempty"`
	WATCH_WARN     string  `json:"WATCH_WARN,omitempty"`
	INITIALS       string  `json:"INITIALS,omitempty"`
	ALAT           float64 `json:"ALAT,omitempty"`
	ALON           float64 `json:"ALON,omitempty"`
	BLAT           float64 `json:"BLAT,omitempty"`
	BLON           float64 `json:"BLON,omitempty"`
	TK_ERR         float64 `json:"TK_ERR,omitempty"`
	X_ERR          float64 `json:"X_ERR,omitempty"`
	Y_ERR          float64 `json:"Y_ERR,omitempty"`
	ALTK_ERR       float64 `json:"ALTK_ERR,omitempty"`
	CRTK_ERR       float64 `json:"CRTK_ERR,omitempty"`
	ADLAND         float64 `json:"ADLAND,omitempty"`
	BDLAND         float64 `json:"BDLAND,omitempty"`
	AMSLP          float64 `json:"AMSLP,omitempty"`
	BMSLP          float64 `json:"BMSLP,omitempty"`
	AMAX_WIND      float64 `json:"AMAX_WIND,omitempty"`
	BMAX_WIND      float64 `json:"BMAX_WIND,omitempty"`
	AAL_WIND_34    float64 `json:"AAL_WIND_34,omitempty"`
	BAL_WIND_34    float64 `json:"BAL_WIND_34,omitempty"`
	ANE_WIND_34    float64 `json:"ANE_WIND_34,omitempty"`
	BNE_WIND_34    float64 `json:"BNE_WIND_34,omitempty"`
	ASE_WIND_34    float64 `json:"ASE_WIND_34,omitempty"`
	BSE_WIND_34    float64 `json:"BSE_WIND_34,omitempty"`
	ASW_WIND_34    float64 `json:"ASW_WIND_34,omitempty"`
	BSW_WIND_34    float64 `json:"BSW_WIND_34,omitempty"`
	ANW_WIND_34    float64 `json:"ANW_WIND_34,omitempty"`
	BNW_WIND_34    float64 `json:"BNW_WIND_34,omitempty"`
	AAL_WIND_50    float64 `json:"AAL_WIND_50,omitempty"`
	BAL_WIND_50    float64 `json:"BAL_WIND_50,omitempty"`
	ANE_WIND_50    float64 `json:"ANE_WIND_50,omitempty"`
	BNE_WIND_50    float64 `json:"BNE_WIND_50,omitempty"`
	ASE_WIND_50    float64 `json:"ASE_WIND_50,omitempty"`
	BSE_WIND_50    float64 `json:"BSE_WIND_50,omitempty"`
	ASW_WIND_50    float64 `json:"ASW_WIND_50,omitempty"`
	BSW_WIND_50    float64 `json:"BSW_WIND_50,omitempty"`
	ANW_WIND_50    float64 `json:"ANW_WIND_50,omitempty"`
	BNW_WIND_50    float64 `json:"BNW_WIND_50,omitempty"`
	AAL_WIND_64    float64 `json:"AAL_WIND_64,omitempty"`
	BAL_WIND_64    float64 `json:"BAL_WIND_64,omitempty"`
	ANE_WIND_64    float64 `json:"ANE_WIND_64,omitempty"`
	BNE_WIND_64    float64 `json:"BNE_WIND_64,omitempty"`
	ASE_WIND_64    float64 `json:"ASE_WIND_64,omitempty"`
	BSE_WIND_64    float64 `json:"BSE_WIND_64,omitempty"`
	ASW_WIND_64    float64 `json:"ASW_WIND_64,omitempty"`
	BSW_WIND_64    float64 `json:"BSW_WIND_64,omitempty"`
	ANW_WIND_64    float64 `json:"ANW_WIND_64,omitempty"`
	BNW_WIND_64    float64 `json:"BNW_WIND_64,omitempty"`
	ARADP          string  `json:"ARADP,omitempty"`
	BRADP          float64 `json:"BRADP,omitempty"`
	ARRP           int     `json:"ARRP,omitempty"`
	BRRP           float64 `json:"BRRP,omitempty"`
	AMRD           int     `json:"AMRD,omitempty"`
	BMRD           float64 `json:"BMRD,omitempty"`
	AGUSTS         int     `json:"AGUSTS,omitempty"`
	BGUSTS         float64 `json:"BGUSTS,omitempty"`
	AEYE           int     `json:"AEYE,omitempty"`
	BEYE           float64 `json:"BEYE,omitempty"`
	ADIR           int     `json:"ADIR,omitempty"`
	BDIR           float64 `json:"BDIR,omitempty"`
	ASPEED         int     `json:"ASPEED,omitempty"`
	BSPEED         float64 `json:"BSPEED,omitempty"`
	ADEPTH         int     `json:"ADEPTH,omitempty"`
	BDEPTH         float64 `json:"BDEPTH,omitempty"`
	NUM_MEMBERS    float64 `json:"NUM_MEMBERS,omitempty"`
	TRACK_SPREAD   float64 `json:"TRACK_SPREAD,omitempty"`
	TRACK_STDEV    float64 `json:"TRACK_STDEV,omitempty"`
	MSLP_STDEV     float64 `json:"MSLP_STDEV,omitempty"`
	MAX_WIND_STDEV float64 `json:"MAX_WIND_STDEV,omitempty"`
	INIT           int     `json:"INIT,omitempty"`
}

// fillStructure functions
func (s *MODE_CTS) fill_MODE_CTS(fields []string) {
	dataLen := len(fields) - 1
	i := -1
	i++
	if i <= dataLen {
		if fields[0] != "NA" {
			s.FIELD = fields[0]
		}
	}
	i++
	if i <= dataLen {
		s.TOTAL, _ = strconv.Atoi(fields[1])
	}
	i++
	if i <= dataLen {
		s.FY_OY, _ = strconv.ParseFloat(fields[2], 64)
	}
	i++
	if i <= dataLen {
		s.FY_ON, _ = strconv.ParseFloat(fields[3], 64)
	}
	i++
	if i <= dataLen {
		s.FN_OY, _ = strconv.ParseFloat(fields[4], 64)
	}
	i++
	if i <= dataLen {
		s.FN_ON, _ = strconv.ParseFloat(fields[5], 64)
	}
	i++
	if i <= dataLen {
		s.BASER, _ = strconv.ParseFloat(fields[6], 64)
	}
	i++
	if i <= dataLen {
		s.FMEAN, _ = strconv.ParseFloat(fields[7], 64)
	}
	i++
	if i <= dataLen {
		s.ACC, _ = strconv.ParseFloat(fields[8], 64)
	}
	i++
	if i <= dataLen {
		s.FBIAS, _ = strconv.ParseFloat(fields[9], 64)
	}
	i++
	if i <= dataLen {
		s.PODY, _ = strconv.ParseFloat(fields[10], 64)
	}
	i++
	if i <= dataLen {
		s.PODN, _ = strconv.ParseFloat(fields[11], 64)
	}
	i++
	if i <= dataLen {
		s.POFD, _ = strconv.ParseFloat(fields[12], 64)
	}
	i++
	if i <= dataLen {
		s.FAR, _ = strconv.ParseFloat(fields[13], 64)
	}
	i++
	if i <= dataLen {
		s.CSI, _ = strconv.ParseFloat(fields[14], 64)
	}
	i++
	if i <= dataLen {
		s.GSS, _ = strconv.ParseFloat(fields[15], 64)
	}
	i++
	if i <= dataLen {
		s.HK, _ = strconv.ParseFloat(fields[16], 64)
	}
	i++
	if i <= dataLen {
		s.HSS, _ = strconv.ParseFloat(fields[17], 64)
	}
	i++
	if i <= dataLen {
		s.ODDS, _ = strconv.ParseFloat(fields[18], 64)
	}
}

func (s *MODE_OBJ) fill_MODE_OBJ(fields []string) {
	dataLen := len(fields) - 1
	i := -1
	i++
	if i <= dataLen {
		if fields[0] != "NA" {
			s.OBJECT_ID = fields[0]
		}
	}
	i++
	if i <= dataLen {
		if fields[1] != "NA" {
			s.OBJECT_CAT = fields[1]
		}
	}
	i++
	if i <= dataLen {
		s.CENTROID_X, _ = strconv.ParseFloat(fields[2], 64)
	}
	i++
	if i <= dataLen {
		s.CENTROID_Y, _ = strconv.ParseFloat(fields[3], 64)
	}
	i++
	if i <= dataLen {
		s.CENTROID_LAT, _ = strconv.ParseFloat(fields[4], 64)
	}
	i++
	if i <= dataLen {
		s.CENTROID_LON, _ = strconv.ParseFloat(fields[5], 64)
	}
	i++
	if i <= dataLen {
		s.AXIS_ANG, _ = strconv.ParseFloat(fields[6], 64)
	}
	i++
	if i <= dataLen {
		s.LENGTH, _ = strconv.ParseFloat(fields[7], 64)
	}
	i++
	if i <= dataLen {
		s.WIDTH, _ = strconv.ParseFloat(fields[8], 64)
	}
	i++
	if i <= dataLen {
		s.AREA, _ = strconv.Atoi(fields[9])
	}
	i++
	if i <= dataLen {
		s.AREA_THRESH, _ = strconv.Atoi(fields[10])
	}
	i++
	if i <= dataLen {
		s.CURVATURE, _ = strconv.ParseFloat(fields[11], 64)
	}
	i++
	if i <= dataLen {
		s.CURVATURE_X, _ = strconv.ParseFloat(fields[12], 64)
	}
	i++
	if i <= dataLen {
		s.CURVATURE_Y, _ = strconv.ParseFloat(fields[13], 64)
	}
	i++
	if i <= dataLen {
		s.COMPLEXITY, _ = strconv.ParseFloat(fields[14], 64)
	}
	i++
	if i <= dataLen {
		s.INTENSITY_10, _ = strconv.ParseFloat(fields[15], 64)
	}
	i++
	if i <= dataLen {
		s.INTENSITY_25, _ = strconv.ParseFloat(fields[16], 64)
	}
	i++
	if i <= dataLen {
		s.INTENSITY_50, _ = strconv.ParseFloat(fields[17], 64)
	}
	i++
	if i <= dataLen {
		s.INTENSITY_75, _ = strconv.ParseFloat(fields[18], 64)
	}
	i++
	if i <= dataLen {
		s.INTENSITY_90, _ = strconv.ParseFloat(fields[19], 64)
	}
	i++
	if i <= dataLen {
		s.INTENSITY_USER, _ = strconv.ParseFloat(fields[20], 64)
	}
	i++
	if i <= dataLen {
		s.INTENSITY_SUM, _ = strconv.ParseFloat(fields[21], 64)
	}
	i++
	if i <= dataLen {
		s.CENTROID_DIST, _ = strconv.ParseFloat(fields[22], 64)
	}
	i++
	if i <= dataLen {
		s.BOUNDARY_DIST, _ = strconv.ParseFloat(fields[23], 64)
	}
	i++
	if i <= dataLen {
		s.CONVEX_HULL_DIST, _ = strconv.ParseFloat(fields[24], 64)
	}
	i++
	if i <= dataLen {
		s.ANGLE_DIFF, _ = strconv.ParseFloat(fields[25], 64)
	}
	i++
	if i <= dataLen {
		s.ASPECT_DIFF, _ = strconv.ParseFloat(fields[26], 64)
	}
	i++
	if i <= dataLen {
		s.AREA_RATIO, _ = strconv.ParseFloat(fields[27], 64)
	}
	i++
	if i <= dataLen {
		s.INTERSECTION_AREA, _ = strconv.ParseFloat(fields[28], 64)
	}
	i++
	if i <= dataLen {
		s.UNION_AREA, _ = strconv.ParseFloat(fields[29], 64)
	}
	i++
	if i <= dataLen {
		s.SYMMETRIC_DIFF, _ = strconv.ParseFloat(fields[30], 64)
	}
	i++
	if i <= dataLen {
		s.INTERSECTION_OVER_AREA, _ = strconv.ParseFloat(fields[31], 64)
	}
	i++
	if i <= dataLen {
		s.CURVATURE_RATIO, _ = strconv.ParseFloat(fields[32], 64)
	}
	i++
	if i <= dataLen {
		s.COMPLEXITY_RATIO, _ = strconv.ParseFloat(fields[33], 64)
	}
	i++
	if i <= dataLen {
		s.PERCENTILE_INTENSITY_RATIO, _ = strconv.ParseFloat(fields[34], 64)
	}
	i++
	if i <= dataLen {
		s.INTEREST, _ = strconv.ParseFloat(fields[35], 64)
	}
}

func (s *STAT_CNT) fill_STAT_CNT(fields []string) {
	dataLen := len(fields) - 1
	i := -1
	i++
	if i <= dataLen {
		s.TOTAL, _ = strconv.Atoi(fields[0])
	}
	i++
	if i <= dataLen {
		s.FBAR, _ = strconv.ParseFloat(fields[1], 64)
	}
	i++
	if i <= dataLen {
		s.FBAR_NCL, _ = strconv.ParseFloat(fields[2], 64)
	}
	i++
	if i <= dataLen {
		s.FBAR_NCU, _ = strconv.ParseFloat(fields[3], 64)
	}
	i++
	if i <= dataLen {
		s.FBAR_BCL, _ = strconv.ParseFloat(fields[4], 64)
	}
	i++
	if i <= dataLen {
		s.FBAR_BCU, _ = strconv.ParseFloat(fields[5], 64)
	}
	i++
	if i <= dataLen {
		s.FSTDEV, _ = strconv.ParseFloat(fields[6], 64)
	}
	i++
	if i <= dataLen {
		s.FSTDEV_NCL, _ = strconv.ParseFloat(fields[7], 64)
	}
	i++
	if i <= dataLen {
		s.FSTDEV_NCU, _ = strconv.ParseFloat(fields[8], 64)
	}
	i++
	if i <= dataLen {
		s.FSTDEV_BCL, _ = strconv.ParseFloat(fields[9], 64)
	}
	i++
	if i <= dataLen {
		s.FSTDEV_BCU, _ = strconv.ParseFloat(fields[10], 64)
	}
	i++
	if i <= dataLen {
		s.OBAR, _ = strconv.ParseFloat(fields[11], 64)
	}
	i++
	if i <= dataLen {
		s.OBAR_NCL, _ = strconv.ParseFloat(fields[12], 64)
	}
	i++
	if i <= dataLen {
		s.OBAR_NCU, _ = strconv.ParseFloat(fields[13], 64)
	}
	i++
	if i <= dataLen {
		s.OBAR_BCL, _ = strconv.ParseFloat(fields[14], 64)
	}
	i++
	if i <= dataLen {
		s.OBAR_BCU, _ = strconv.ParseFloat(fields[15], 64)
	}
	i++
	if i <= dataLen {
		s.OSTDEV, _ = strconv.ParseFloat(fields[16], 64)
	}
	i++
	if i <= dataLen {
		s.OSTDEV_NCL, _ = strconv.ParseFloat(fields[17], 64)
	}
	i++
	if i <= dataLen {
		s.OSTDEV_NCU, _ = strconv.ParseFloat(fields[18], 64)
	}
	i++
	if i <= dataLen {
		s.OSTDEV_BCL, _ = strconv.ParseFloat(fields[19], 64)
	}
	i++
	if i <= dataLen {
		s.OSTDEV_BCU, _ = strconv.ParseFloat(fields[20], 64)
	}
	i++
	if i <= dataLen {
		s.PR_CORR, _ = strconv.ParseFloat(fields[21], 64)
	}
	i++
	if i <= dataLen {
		s.PR_CORR_NCL, _ = strconv.ParseFloat(fields[22], 64)
	}
	i++
	if i <= dataLen {
		s.PR_CORR_NCU, _ = strconv.ParseFloat(fields[23], 64)
	}
	i++
	if i <= dataLen {
		s.PR_CORR_BCL, _ = strconv.ParseFloat(fields[24], 64)
	}
	i++
	if i <= dataLen {
		s.PR_CORR_BCU, _ = strconv.ParseFloat(fields[25], 64)
	}
	i++
	if i <= dataLen {
		s.SP_CORR, _ = strconv.ParseFloat(fields[26], 64)
	}
	i++
	if i <= dataLen {
		s.KT_CORR, _ = strconv.ParseFloat(fields[27], 64)
	}
	i++
	if i <= dataLen {
		s.RANKS, _ = strconv.Atoi(fields[28])
	}
	i++
	if i <= dataLen {
		s.FRANK_TIES, _ = strconv.Atoi(fields[29])
	}
	i++
	if i <= dataLen {
		s.ORANK_TIES, _ = strconv.Atoi(fields[30])
	}
	i++
	if i <= dataLen {
		s.ME, _ = strconv.ParseFloat(fields[31], 64)
	}
	i++
	if i <= dataLen {
		s.ME_NCL, _ = strconv.ParseFloat(fields[32], 64)
	}
	i++
	if i <= dataLen {
		s.ME_NCU, _ = strconv.ParseFloat(fields[33], 64)
	}
	i++
	if i <= dataLen {
		s.ME_BCL, _ = strconv.ParseFloat(fields[34], 64)
	}
	i++
	if i <= dataLen {
		s.ME_BCU, _ = strconv.ParseFloat(fields[35], 64)
	}
	i++
	if i <= dataLen {
		s.ESTDEV, _ = strconv.ParseFloat(fields[36], 64)
	}
	i++
	if i <= dataLen {
		s.ESTDEV_NCL, _ = strconv.ParseFloat(fields[37], 64)
	}
	i++
	if i <= dataLen {
		s.ESTDEV_NCU, _ = strconv.ParseFloat(fields[38], 64)
	}
	i++
	if i <= dataLen {
		s.ESTDEV_BCL, _ = strconv.ParseFloat(fields[39], 64)
	}
	i++
	if i <= dataLen {
		s.ESTDEV_BCU, _ = strconv.ParseFloat(fields[40], 64)
	}
	i++
	if i <= dataLen {
		s.MBIAS, _ = strconv.ParseFloat(fields[41], 64)
	}
	i++
	if i <= dataLen {
		s.MBIAS_BCL, _ = strconv.ParseFloat(fields[42], 64)
	}
	i++
	if i <= dataLen {
		s.MBIAS_BCU, _ = strconv.ParseFloat(fields[43], 64)
	}
	i++
	if i <= dataLen {
		s.MAE, _ = strconv.ParseFloat(fields[44], 64)
	}
	i++
	if i <= dataLen {
		s.MAE_BCL, _ = strconv.ParseFloat(fields[45], 64)
	}
	i++
	if i <= dataLen {
		s.MAE_BCU, _ = strconv.ParseFloat(fields[46], 64)
	}
	i++
	if i <= dataLen {
		s.MSE, _ = strconv.ParseFloat(fields[47], 64)
	}
	i++
	if i <= dataLen {
		s.MSE_BCL, _ = strconv.ParseFloat(fields[48], 64)
	}
	i++
	if i <= dataLen {
		s.MSE_BCU, _ = strconv.ParseFloat(fields[49], 64)
	}
	i++
	if i <= dataLen {
		s.BCMSE, _ = strconv.ParseFloat(fields[50], 64)
	}
	i++
	if i <= dataLen {
		s.BCMSE_BCL, _ = strconv.ParseFloat(fields[51], 64)
	}
	i++
	if i <= dataLen {
		s.BCMSE_BCU, _ = strconv.ParseFloat(fields[52], 64)
	}
	i++
	if i <= dataLen {
		s.RMSE, _ = strconv.ParseFloat(fields[53], 64)
	}
	i++
	if i <= dataLen {
		s.RMSE_BCL, _ = strconv.ParseFloat(fields[54], 64)
	}
	i++
	if i <= dataLen {
		s.RMSE_BCU, _ = strconv.ParseFloat(fields[55], 64)
	}
	i++
	if i <= dataLen {
		s.E10, _ = strconv.ParseFloat(fields[56], 64)
	}
	i++
	if i <= dataLen {
		s.E10_BCL, _ = strconv.ParseFloat(fields[57], 64)
	}
	i++
	if i <= dataLen {
		s.E10_BCU, _ = strconv.ParseFloat(fields[58], 64)
	}
	i++
	if i <= dataLen {
		s.E25, _ = strconv.ParseFloat(fields[59], 64)
	}
	i++
	if i <= dataLen {
		s.E25_BCL, _ = strconv.ParseFloat(fields[60], 64)
	}
	i++
	if i <= dataLen {
		s.E25_BCU, _ = strconv.ParseFloat(fields[61], 64)
	}
	i++
	if i <= dataLen {
		s.E50, _ = strconv.ParseFloat(fields[62], 64)
	}
	i++
	if i <= dataLen {
		s.E50_BCL, _ = strconv.ParseFloat(fields[63], 64)
	}
	i++
	if i <= dataLen {
		s.E50_BCU, _ = strconv.ParseFloat(fields[64], 64)
	}
	i++
	if i <= dataLen {
		s.E75, _ = strconv.ParseFloat(fields[65], 64)
	}
	i++
	if i <= dataLen {
		s.E75_BCL, _ = strconv.ParseFloat(fields[66], 64)
	}
	i++
	if i <= dataLen {
		s.E75_BCU, _ = strconv.ParseFloat(fields[67], 64)
	}
	i++
	if i <= dataLen {
		s.E90, _ = strconv.ParseFloat(fields[68], 64)
	}
	i++
	if i <= dataLen {
		s.E90_BCL, _ = strconv.ParseFloat(fields[69], 64)
	}
	i++
	if i <= dataLen {
		s.E90_BCU, _ = strconv.ParseFloat(fields[70], 64)
	}
	i++
	if i <= dataLen {
		s.EIQR, _ = strconv.ParseFloat(fields[71], 64)
	}
	i++
	if i <= dataLen {
		s.EIQR_BCL, _ = strconv.ParseFloat(fields[72], 64)
	}
	i++
	if i <= dataLen {
		s.EIQR_BCU, _ = strconv.ParseFloat(fields[73], 64)
	}
	i++
	if i <= dataLen {
		s.MAD, _ = strconv.ParseFloat(fields[74], 64)
	}
	i++
	if i <= dataLen {
		s.MAD_BCL, _ = strconv.ParseFloat(fields[75], 64)
	}
	i++
	if i <= dataLen {
		s.MAD_BCU, _ = strconv.ParseFloat(fields[76], 64)
	}
	i++
	if i <= dataLen {
		s.ANOM_CORR, _ = strconv.ParseFloat(fields[77], 64)
	}
	i++
	if i <= dataLen {
		s.ANOM_CORR_NCL, _ = strconv.ParseFloat(fields[78], 64)
	}
	i++
	if i <= dataLen {
		s.ANOM_CORR_NCU, _ = strconv.ParseFloat(fields[79], 64)
	}
	i++
	if i <= dataLen {
		s.ANOM_CORR_BCL, _ = strconv.ParseFloat(fields[80], 64)
	}
	i++
	if i <= dataLen {
		s.ANOM_CORR_BCU, _ = strconv.ParseFloat(fields[81], 64)
	}
	i++
	if i <= dataLen {
		s.ME2, _ = strconv.ParseFloat(fields[82], 64)
	}
	i++
	if i <= dataLen {
		s.ME2_BCL, _ = strconv.ParseFloat(fields[83], 64)
	}
	i++
	if i <= dataLen {
		s.ME2_BCU, _ = strconv.ParseFloat(fields[84], 64)
	}
	i++
	if i <= dataLen {
		s.MSESS, _ = strconv.ParseFloat(fields[85], 64)
	}
	i++
	if i <= dataLen {
		s.MSESS_BCL, _ = strconv.ParseFloat(fields[86], 64)
	}
	i++
	if i <= dataLen {
		s.MSESS_BCU, _ = strconv.ParseFloat(fields[87], 64)
	}
	i++
	if i <= dataLen {
		s.RMSFA, _ = strconv.ParseFloat(fields[88], 64)
	}
	i++
	if i <= dataLen {
		s.RMSFA_BCL, _ = strconv.ParseFloat(fields[89], 64)
	}
	i++
	if i <= dataLen {
		s.RMSFA_BCU, _ = strconv.ParseFloat(fields[90], 64)
	}
	i++
	if i <= dataLen {
		s.RMSOA, _ = strconv.ParseFloat(fields[91], 64)
	}
	i++
	if i <= dataLen {
		s.RMSOA_BCL, _ = strconv.ParseFloat(fields[92], 64)
	}
	i++
	if i <= dataLen {
		s.RMSOA_BCU, _ = strconv.ParseFloat(fields[93], 64)
	}
	i++
	if i <= dataLen {
		s.ANOM_CORR_UNCNTR, _ = strconv.ParseFloat(fields[94], 64)
	}
	i++
	if i <= dataLen {
		s.ANOM_CORR_UNCNTR_BCL, _ = strconv.ParseFloat(fields[95], 64)
	}
	i++
	if i <= dataLen {
		s.ANOM_CORR_UNCNTR_BCU, _ = strconv.ParseFloat(fields[96], 64)
	}
	i++
	if i <= dataLen {
		s.SI, _ = strconv.ParseFloat(fields[97], 64)
	}
	i++
	if i <= dataLen {
		s.SI_BCL, _ = strconv.ParseFloat(fields[98], 64)
	}
	i++
	if i <= dataLen {
		s.SI_BCU, _ = strconv.ParseFloat(fields[99], 64)
	}
}

func (s *STAT_CTC) fill_STAT_CTC(fields []string) {
	dataLen := len(fields) - 1
	i := -1
	i++
	if i <= dataLen {
		s.TOTAL, _ = strconv.Atoi(fields[0])
	}
	i++
	if i <= dataLen {
		s.FY_OY, _ = strconv.ParseFloat(fields[1], 64)
	}
	i++
	if i <= dataLen {
		s.FY_ON, _ = strconv.ParseFloat(fields[2], 64)
	}
	i++
	if i <= dataLen {
		s.FN_OY, _ = strconv.ParseFloat(fields[3], 64)
	}
	i++
	if i <= dataLen {
		s.FN_ON, _ = strconv.ParseFloat(fields[4], 64)
	}
	i++
	if i <= dataLen {
		s.EC_VALUE, _ = strconv.ParseFloat(fields[5], 64)
	}
}

func (s *STAT_CTS) fill_STAT_CTS(fields []string) {
	dataLen := len(fields) - 1
	i := -1
	i++
	if i <= dataLen {
		s.TOTAL, _ = strconv.Atoi(fields[0])
	}
	i++
	if i <= dataLen {
		s.BASER, _ = strconv.ParseFloat(fields[1], 64)
	}
	i++
	if i <= dataLen {
		s.BASER_NCL, _ = strconv.ParseFloat(fields[2], 64)
	}
	i++
	if i <= dataLen {
		s.BASER_NCU, _ = strconv.ParseFloat(fields[3], 64)
	}
	i++
	if i <= dataLen {
		s.BASER_BCL, _ = strconv.ParseFloat(fields[4], 64)
	}
	i++
	if i <= dataLen {
		s.BASER_BCU, _ = strconv.ParseFloat(fields[5], 64)
	}
	i++
	if i <= dataLen {
		s.FMEAN, _ = strconv.ParseFloat(fields[6], 64)
	}
	i++
	if i <= dataLen {
		s.FMEAN_NCL, _ = strconv.ParseFloat(fields[7], 64)
	}
	i++
	if i <= dataLen {
		s.FMEAN_NCU, _ = strconv.ParseFloat(fields[8], 64)
	}
	i++
	if i <= dataLen {
		s.FMEAN_BCL, _ = strconv.ParseFloat(fields[9], 64)
	}
	i++
	if i <= dataLen {
		s.FMEAN_BCU, _ = strconv.ParseFloat(fields[10], 64)
	}
	i++
	if i <= dataLen {
		s.ACC, _ = strconv.ParseFloat(fields[11], 64)
	}
	i++
	if i <= dataLen {
		s.ACC_NCL, _ = strconv.ParseFloat(fields[12], 64)
	}
	i++
	if i <= dataLen {
		s.ACC_NCU, _ = strconv.ParseFloat(fields[13], 64)
	}
	i++
	if i <= dataLen {
		s.ACC_BCL, _ = strconv.ParseFloat(fields[14], 64)
	}
	i++
	if i <= dataLen {
		s.ACC_BCU, _ = strconv.ParseFloat(fields[15], 64)
	}
	i++
	if i <= dataLen {
		s.FBIAS, _ = strconv.ParseFloat(fields[16], 64)
	}
	i++
	if i <= dataLen {
		s.FBIAS_BCL, _ = strconv.ParseFloat(fields[17], 64)
	}
	i++
	if i <= dataLen {
		s.FBIAS_BCU, _ = strconv.ParseFloat(fields[18], 64)
	}
	i++
	if i <= dataLen {
		s.PODY, _ = strconv.ParseFloat(fields[19], 64)
	}
	i++
	if i <= dataLen {
		s.PODY_NCL, _ = strconv.ParseFloat(fields[20], 64)
	}
	i++
	if i <= dataLen {
		s.PODY_NCU, _ = strconv.ParseFloat(fields[21], 64)
	}
	i++
	if i <= dataLen {
		s.PODY_BCL, _ = strconv.ParseFloat(fields[22], 64)
	}
	i++
	if i <= dataLen {
		s.PODY_BCU, _ = strconv.ParseFloat(fields[23], 64)
	}
	i++
	if i <= dataLen {
		s.PODN, _ = strconv.ParseFloat(fields[24], 64)
	}
	i++
	if i <= dataLen {
		s.PODN_NCL, _ = strconv.ParseFloat(fields[25], 64)
	}
	i++
	if i <= dataLen {
		s.PODN_NCU, _ = strconv.ParseFloat(fields[26], 64)
	}
	i++
	if i <= dataLen {
		s.PODN_BCL, _ = strconv.ParseFloat(fields[27], 64)
	}
	i++
	if i <= dataLen {
		s.PODN_BCU, _ = strconv.ParseFloat(fields[28], 64)
	}
	i++
	if i <= dataLen {
		s.POFD, _ = strconv.ParseFloat(fields[29], 64)
	}
	i++
	if i <= dataLen {
		s.POFD_NCL, _ = strconv.ParseFloat(fields[30], 64)
	}
	i++
	if i <= dataLen {
		s.POFD_NCU, _ = strconv.ParseFloat(fields[31], 64)
	}
	i++
	if i <= dataLen {
		s.POFD_BCL, _ = strconv.ParseFloat(fields[32], 64)
	}
	i++
	if i <= dataLen {
		s.POFD_BCU, _ = strconv.ParseFloat(fields[33], 64)
	}
	i++
	if i <= dataLen {
		s.FAR, _ = strconv.ParseFloat(fields[34], 64)
	}
	i++
	if i <= dataLen {
		s.FAR_NCL, _ = strconv.ParseFloat(fields[35], 64)
	}
	i++
	if i <= dataLen {
		s.FAR_NCU, _ = strconv.ParseFloat(fields[36], 64)
	}
	i++
	if i <= dataLen {
		s.FAR_BCL, _ = strconv.ParseFloat(fields[37], 64)
	}
	i++
	if i <= dataLen {
		s.FAR_BCU, _ = strconv.ParseFloat(fields[38], 64)
	}
	i++
	if i <= dataLen {
		s.CSI, _ = strconv.ParseFloat(fields[39], 64)
	}
	i++
	if i <= dataLen {
		s.CSI_NCL, _ = strconv.ParseFloat(fields[40], 64)
	}
	i++
	if i <= dataLen {
		s.CSI_NCU, _ = strconv.ParseFloat(fields[41], 64)
	}
	i++
	if i <= dataLen {
		s.CSI_BCL, _ = strconv.ParseFloat(fields[42], 64)
	}
	i++
	if i <= dataLen {
		s.CSI_BCU, _ = strconv.ParseFloat(fields[43], 64)
	}
	i++
	if i <= dataLen {
		s.GSS, _ = strconv.ParseFloat(fields[44], 64)
	}
	i++
	if i <= dataLen {
		s.GSS_BCL, _ = strconv.ParseFloat(fields[45], 64)
	}
	i++
	if i <= dataLen {
		s.GSS_BCU, _ = strconv.ParseFloat(fields[46], 64)
	}
	i++
	if i <= dataLen {
		s.HK, _ = strconv.ParseFloat(fields[47], 64)
	}
	i++
	if i <= dataLen {
		s.HK_NCL, _ = strconv.ParseFloat(fields[48], 64)
	}
	i++
	if i <= dataLen {
		s.HK_NCU, _ = strconv.ParseFloat(fields[49], 64)
	}
	i++
	if i <= dataLen {
		s.HK_BCL, _ = strconv.ParseFloat(fields[50], 64)
	}
	i++
	if i <= dataLen {
		s.HK_BCU, _ = strconv.ParseFloat(fields[51], 64)
	}
	i++
	if i <= dataLen {
		s.HSS, _ = strconv.ParseFloat(fields[52], 64)
	}
	i++
	if i <= dataLen {
		s.HSS_BCL, _ = strconv.ParseFloat(fields[53], 64)
	}
	i++
	if i <= dataLen {
		s.HSS_BCU, _ = strconv.ParseFloat(fields[54], 64)
	}
	i++
	if i <= dataLen {
		s.ODDS, _ = strconv.ParseFloat(fields[55], 64)
	}
	i++
	if i <= dataLen {
		s.ODDS_NCL, _ = strconv.ParseFloat(fields[56], 64)
	}
	i++
	if i <= dataLen {
		s.ODDS_NCU, _ = strconv.ParseFloat(fields[57], 64)
	}
	i++
	if i <= dataLen {
		s.ODDS_BCL, _ = strconv.ParseFloat(fields[58], 64)
	}
	i++
	if i <= dataLen {
		s.ODDS_BCU, _ = strconv.ParseFloat(fields[59], 64)
	}
	i++
	if i <= dataLen {
		s.LODDS, _ = strconv.ParseFloat(fields[60], 64)
	}
	i++
	if i <= dataLen {
		s.LODDS_NCL, _ = strconv.ParseFloat(fields[61], 64)
	}
	i++
	if i <= dataLen {
		s.LODDS_NCU, _ = strconv.ParseFloat(fields[62], 64)
	}
	i++
	if i <= dataLen {
		s.LODDS_BCL, _ = strconv.ParseFloat(fields[63], 64)
	}
	i++
	if i <= dataLen {
		s.LODDS_BCU, _ = strconv.ParseFloat(fields[64], 64)
	}
	i++
	if i <= dataLen {
		s.ORSS, _ = strconv.ParseFloat(fields[65], 64)
	}
	i++
	if i <= dataLen {
		s.ORSS_NCL, _ = strconv.ParseFloat(fields[66], 64)
	}
	i++
	if i <= dataLen {
		s.ORSS_NCU, _ = strconv.ParseFloat(fields[67], 64)
	}
	i++
	if i <= dataLen {
		s.ORSS_BCL, _ = strconv.ParseFloat(fields[68], 64)
	}
	i++
	if i <= dataLen {
		s.ORSS_BCU, _ = strconv.ParseFloat(fields[69], 64)
	}
	i++
	if i <= dataLen {
		s.EDS, _ = strconv.ParseFloat(fields[70], 64)
	}
	i++
	if i <= dataLen {
		s.EDS_NCL, _ = strconv.ParseFloat(fields[71], 64)
	}
	i++
	if i <= dataLen {
		s.EDS_NCU, _ = strconv.ParseFloat(fields[72], 64)
	}
	i++
	if i <= dataLen {
		s.EDS_BCL, _ = strconv.ParseFloat(fields[73], 64)
	}
	i++
	if i <= dataLen {
		s.EDS_BCU, _ = strconv.ParseFloat(fields[74], 64)
	}
	i++
	if i <= dataLen {
		s.SEDS, _ = strconv.ParseFloat(fields[75], 64)
	}
	i++
	if i <= dataLen {
		s.SEDS_NCL, _ = strconv.ParseFloat(fields[76], 64)
	}
	i++
	if i <= dataLen {
		s.SEDS_NCU, _ = strconv.ParseFloat(fields[77], 64)
	}
	i++
	if i <= dataLen {
		s.SEDS_BCL, _ = strconv.ParseFloat(fields[78], 64)
	}
	i++
	if i <= dataLen {
		s.SEDS_BCU, _ = strconv.ParseFloat(fields[79], 64)
	}
	i++
	if i <= dataLen {
		s.EDI, _ = strconv.ParseFloat(fields[80], 64)
	}
	i++
	if i <= dataLen {
		s.EDI_NCL, _ = strconv.ParseFloat(fields[81], 64)
	}
	i++
	if i <= dataLen {
		s.EDI_NCU, _ = strconv.ParseFloat(fields[82], 64)
	}
	i++
	if i <= dataLen {
		s.EDI_BCL, _ = strconv.ParseFloat(fields[83], 64)
	}
	i++
	if i <= dataLen {
		s.EDI_BCU, _ = strconv.ParseFloat(fields[84], 64)
	}
	i++
	if i <= dataLen {
		s.SEDI, _ = strconv.ParseFloat(fields[85], 64)
	}
	i++
	if i <= dataLen {
		s.SEDI_NCL, _ = strconv.ParseFloat(fields[86], 64)
	}
	i++
	if i <= dataLen {
		s.SEDI_NCU, _ = strconv.ParseFloat(fields[87], 64)
	}
	i++
	if i <= dataLen {
		s.SEDI_BCL, _ = strconv.ParseFloat(fields[88], 64)
	}
	i++
	if i <= dataLen {
		s.SEDI_BCU, _ = strconv.ParseFloat(fields[89], 64)
	}
	i++
	if i <= dataLen {
		s.BAGSS, _ = strconv.ParseFloat(fields[90], 64)
	}
	i++
	if i <= dataLen {
		s.BAGSS_BCL, _ = strconv.ParseFloat(fields[91], 64)
	}
	i++
	if i <= dataLen {
		s.BAGSS_BCU, _ = strconv.ParseFloat(fields[92], 64)
	}
	i++
	if i <= dataLen {
		s.HSS_EC, _ = strconv.ParseFloat(fields[93], 64)
	}
	i++
	if i <= dataLen {
		s.HSS_EC_BCL, _ = strconv.ParseFloat(fields[94], 64)
	}
	i++
	if i <= dataLen {
		s.HSS_EC_BCU, _ = strconv.ParseFloat(fields[95], 64)
	}
	i++
	if i <= dataLen {
		s.EC_VALUE, _ = strconv.ParseFloat(fields[96], 64)
	}
}

func (s *STAT_DMAP) fill_STAT_DMAP(fields []string) {
	dataLen := len(fields) - 1
	i := -1
	i++
	if i <= dataLen {
		s.TOTAL, _ = strconv.Atoi(fields[0])
	}
	i++
	if i <= dataLen {
		s.FY, _ = strconv.Atoi(fields[1])
	}
	i++
	if i <= dataLen {
		s.OY, _ = strconv.Atoi(fields[2])
	}
	i++
	if i <= dataLen {
		s.FBIAS, _ = strconv.ParseFloat(fields[3], 64)
	}
	i++
	if i <= dataLen {
		s.BADDELEY, _ = strconv.ParseFloat(fields[4], 64)
	}
	i++
	if i <= dataLen {
		s.HAUSDORFF, _ = strconv.ParseFloat(fields[5], 64)
	}
	i++
	if i <= dataLen {
		s.MED_FO, _ = strconv.ParseFloat(fields[6], 64)
	}
	i++
	if i <= dataLen {
		s.MED_OF, _ = strconv.ParseFloat(fields[7], 64)
	}
	i++
	if i <= dataLen {
		s.MED_MIN, _ = strconv.ParseFloat(fields[8], 64)
	}
	i++
	if i <= dataLen {
		s.MED_MAX, _ = strconv.ParseFloat(fields[9], 64)
	}
	i++
	if i <= dataLen {
		s.MED_MEAN, _ = strconv.ParseFloat(fields[10], 64)
	}
	i++
	if i <= dataLen {
		s.FOM_FO, _ = strconv.ParseFloat(fields[11], 64)
	}
	i++
	if i <= dataLen {
		s.FOM_OF, _ = strconv.ParseFloat(fields[12], 64)
	}
	i++
	if i <= dataLen {
		s.FOM_MIN, _ = strconv.ParseFloat(fields[13], 64)
	}
	i++
	if i <= dataLen {
		s.FOM_MAX, _ = strconv.ParseFloat(fields[14], 64)
	}
	i++
	if i <= dataLen {
		s.FOM_MEAN, _ = strconv.ParseFloat(fields[15], 64)
	}
	i++
	if i <= dataLen {
		s.ZHU_FO, _ = strconv.ParseFloat(fields[16], 64)
	}
	i++
	if i <= dataLen {
		s.ZHU_OF, _ = strconv.ParseFloat(fields[17], 64)
	}
	i++
	if i <= dataLen {
		s.ZHU_MIN, _ = strconv.ParseFloat(fields[18], 64)
	}
	i++
	if i <= dataLen {
		s.ZHU_MAX, _ = strconv.ParseFloat(fields[19], 64)
	}
	i++
	if i <= dataLen {
		s.ZHU_MEAN, _ = strconv.ParseFloat(fields[20], 64)
	}
	i++
	if i <= dataLen {
		s.G, _ = strconv.ParseFloat(fields[21], 64)
	}
	i++
	if i <= dataLen {
		s.GBETA, _ = strconv.ParseFloat(fields[22], 64)
	}
	i++
	if i <= dataLen {
		s.BETA_VALUE, _ = strconv.ParseFloat(fields[23], 64)
	}
}

func (s *STAT_ECLV) fill_STAT_ECLV(fields []string) {
	dataLen := len(fields) - 1
	i := -1
	i++
	if i <= dataLen {
		s.TOTAL, _ = strconv.Atoi(fields[0])
	}
	i++
	if i <= dataLen {
		s.BASER, _ = strconv.ParseFloat(fields[1], 64)
	}
	i++
	if i <= dataLen {
		s.VALUE_BASER, _ = strconv.Atoi(fields[2])
	}
	i++
	if i <= dataLen { // the first field of the repeating fields is the TOTAL, the second field is the 1st dimenSion of the 1st sequence (there might be only one sequence)
		var value interface{}
		count, err := strconv.Atoi(fields[3])
		if err != nil {
			count = 0
		}
		keyPrefixes := []string{"CL_", "VALUE_"}
		s.PTS = make(map[string]interface{})
		for group := 1; group <= count; group++ {
			for index := 4; index <= len(keyPrefixes); index++ {
				key := fmt.Sprintf("%s_%d", keyPrefixes[index-1], index)
				if index > len(fields) { // sometimes the data line is truncated - we will set expected data to "NA"
					value = "NA"
				} else {
					value, err = strconv.ParseFloat(fields[index], 64)
					if err != nil { // sometimes there can be these NA values in the data, which will be left out of json
						value = "NA"
					}
				}
				s.PTS[key] = value
			}
		}
	}
}

func (s *STAT_ECNT) fill_STAT_ECNT(fields []string) {
	dataLen := len(fields) - 1
	i := -1
	i++
	if i <= dataLen {
		s.TOTAL, _ = strconv.Atoi(fields[0])
	}
	i++
	if i <= dataLen {
		s.N_ENS, _ = strconv.Atoi(fields[1])
	}
	i++
	if i <= dataLen {
		s.CRPS, _ = strconv.ParseFloat(fields[2], 64)
	}
	i++
	if i <= dataLen {
		s.CRPSS, _ = strconv.ParseFloat(fields[3], 64)
	}
	i++
	if i <= dataLen {
		s.IGN, _ = strconv.ParseFloat(fields[4], 64)
	}
	i++
	if i <= dataLen {
		s.ME, _ = strconv.ParseFloat(fields[5], 64)
	}
	i++
	if i <= dataLen {
		s.RMSE, _ = strconv.ParseFloat(fields[6], 64)
	}
	i++
	if i <= dataLen {
		s.SPREAD, _ = strconv.ParseFloat(fields[7], 64)
	}
	i++
	if i <= dataLen {
		s.ME_OERR, _ = strconv.ParseFloat(fields[8], 64)
	}
	i++
	if i <= dataLen {
		s.RMSE_OERR, _ = strconv.ParseFloat(fields[9], 64)
	}
	i++
	if i <= dataLen {
		s.SPREAD_OERR, _ = strconv.ParseFloat(fields[10], 64)
	}
	i++
	if i <= dataLen {
		s.SPREAD_PLUS_OERR, _ = strconv.ParseFloat(fields[11], 64)
	}
	i++
	if i <= dataLen {
		s.CRPSCL, _ = strconv.ParseFloat(fields[12], 64)
	}
	i++
	if i <= dataLen {
		s.CRPS_EMP, _ = strconv.ParseFloat(fields[13], 64)
	}
	i++
	if i <= dataLen {
		s.CRPSCL_EMP, _ = strconv.ParseFloat(fields[14], 64)
	}
	i++
	if i <= dataLen {
		s.CRPSS_EMP, _ = strconv.ParseFloat(fields[15], 64)
	}
	i++
	if i <= dataLen {
		s.CRPS_EMP_FAIR, _ = strconv.ParseFloat(fields[16], 64)
	}
	i++
	if i <= dataLen {
		s.SPREAD_MD, _ = strconv.ParseFloat(fields[17], 64)
	}
	i++
	if i <= dataLen {
		s.MAE, _ = strconv.ParseFloat(fields[18], 64)
	}
	i++
	if i <= dataLen {
		s.MAE_OERR, _ = strconv.ParseFloat(fields[19], 64)
	}
	i++
	if i <= dataLen {
		s.BIAS_RATIO, _ = strconv.ParseFloat(fields[20], 64)
	}
	i++
	if i <= dataLen {
		s.N_GE_OBS, _ = strconv.Atoi(fields[21])
	}
	i++
	if i <= dataLen {
		s.ME_GE_OBS, _ = strconv.ParseFloat(fields[22], 64)
	}
	i++
	if i <= dataLen {
		s.N_LT_OBS, _ = strconv.Atoi(fields[23])
	}
	i++
	if i <= dataLen {
		s.ME_LT_OBS, _ = strconv.ParseFloat(fields[24], 64)
	}
}

func (s *STAT_FHO) fill_STAT_FHO(fields []string) {
	dataLen := len(fields) - 1
	i := -1
	i++
	if i <= dataLen {
		s.TOTAL, _ = strconv.Atoi(fields[0])
	}
	i++
	if i <= dataLen {
		s.F_RATE, _ = strconv.ParseFloat(fields[1], 64)
	}
	i++
	if i <= dataLen {
		s.H_RATE, _ = strconv.ParseFloat(fields[2], 64)
	}
	i++
	if i <= dataLen {
		s.O_RATE, _ = strconv.ParseFloat(fields[3], 64)
	}
}

func (s *STAT_GENMPR) fill_STAT_GENMPR(fields []string) {
	dataLen := len(fields) - 1
	i := -1
	i++
	if i <= dataLen {
		s.TOTAL, _ = strconv.Atoi(fields[0])
	}
	i++
	if i <= dataLen {
		s.INDEX, _ = strconv.Atoi(fields[1])
	}
	i++
	if i <= dataLen {
		if fields[2] != "NA" {
			s.STORM_ID = fields[2]
		}
	}
	i++
	if i <= dataLen {
		s.PROB_LEAD, _ = strconv.ParseFloat(fields[3], 64)
	}
	i++
	if i <= dataLen {
		s.PROB_VAL, _ = strconv.ParseFloat(fields[4], 64)
	}
	i++
	if i <= dataLen {
		if fields[5] != "NA" {
			s.AGEN_INIT = fields[5]
		}
	}
	i++
	if i <= dataLen {
		if fields[6] != "NA" {
			s.AGEN_FHR = fields[6]
		}
	}
	i++
	if i <= dataLen {
		s.AGEN_LAT, _ = strconv.ParseFloat(fields[7], 64)
	}
	i++
	if i <= dataLen {
		s.AGEN_LON, _ = strconv.ParseFloat(fields[8], 64)
	}
	i++
	if i <= dataLen {
		s.AGEN_DLAND, _ = strconv.ParseFloat(fields[9], 64)
	}
	i++
	if i <= dataLen {
		s.BGEN_LAT, _ = strconv.ParseFloat(fields[10], 64)
	}
	i++
	if i <= dataLen {
		s.BGEN_LON, _ = strconv.ParseFloat(fields[11], 64)
	}
	i++
	if i <= dataLen {
		s.BGEN_DLAND, _ = strconv.ParseFloat(fields[12], 64)
	}
	i++
	if i <= dataLen {
		s.GEN_DIST, _ = strconv.ParseFloat(fields[13], 64)
	}
	i++
	if i <= dataLen {
		if fields[14] != "NA" {
			s.GEN_TDIFF = fields[14]
		}
	}
	i++
	if i <= dataLen {
		if fields[15] != "NA" {
			s.INIT_TDIFF = fields[15]
		}
	}
	i++
	if i <= dataLen {
		if fields[16] != "NA" {
			s.DEV_CAT = fields[16]
		}
	}
	i++
	if i <= dataLen {
		if fields[17] != "NA" {
			s.OPS_CAT = fields[17]
		}
	}
}

func (s *STAT_GRAD) fill_STAT_GRAD(fields []string) {
	dataLen := len(fields) - 1
	i := -1
	i++
	if i <= dataLen {
		s.TOTAL, _ = strconv.Atoi(fields[0])
	}
	i++
	if i <= dataLen {
		s.FGBAR, _ = strconv.ParseFloat(fields[1], 64)
	}
	i++
	if i <= dataLen {
		s.OGBAR, _ = strconv.ParseFloat(fields[2], 64)
	}
	i++
	if i <= dataLen {
		s.MGBAR, _ = strconv.ParseFloat(fields[3], 64)
	}
	i++
	if i <= dataLen {
		s.EGBAR, _ = strconv.ParseFloat(fields[4], 64)
	}
	i++
	if i <= dataLen {
		s.S1, _ = strconv.ParseFloat(fields[5], 64)
	}
	i++
	if i <= dataLen {
		s.S1_OG, _ = strconv.ParseFloat(fields[6], 64)
	}
	i++
	if i <= dataLen {
		s.FGOG_RATIO, _ = strconv.ParseFloat(fields[7], 64)
	}
	i++
	if i <= dataLen {
		s.DX, _ = strconv.ParseFloat(fields[8], 64)
	}
	i++
	if i <= dataLen {
		s.DY, _ = strconv.ParseFloat(fields[9], 64)
	}
}

func (s *STAT_ISC) fill_STAT_ISC(fields []string) {
	dataLen := len(fields) - 1
	i := -1
	i++
	if i <= dataLen {
		s.TOTAL, _ = strconv.Atoi(fields[0])
	}
	i++
	if i <= dataLen {
		s.TILE_DIM, _ = strconv.Atoi(fields[1])
	}
	i++
	if i <= dataLen {
		s.TILE_XLL, _ = strconv.Atoi(fields[2])
	}
	i++
	if i <= dataLen {
		s.TILE_YLL, _ = strconv.Atoi(fields[3])
	}
	i++
	if i <= dataLen {
		s.NSCALE, _ = strconv.Atoi(fields[4])
	}
	i++
	if i <= dataLen {
		s.ISCALE, _ = strconv.Atoi(fields[5])
	}
	i++
	if i <= dataLen {
		s.MSE, _ = strconv.ParseFloat(fields[6], 64)
	}
	i++
	if i <= dataLen {
		s.ISC, _ = strconv.ParseFloat(fields[7], 64)
	}
	i++
	if i <= dataLen {
		s.FENERGY2, _ = strconv.ParseFloat(fields[8], 64)
	}
	i++
	if i <= dataLen {
		s.OENERGY2, _ = strconv.ParseFloat(fields[9], 64)
	}
	i++
	if i <= dataLen {
		s.BASER, _ = strconv.ParseFloat(fields[10], 64)
	}
	i++
	if i <= dataLen {
		s.FBIAS, _ = strconv.ParseFloat(fields[11], 64)
	}
}

func (s *STAT_MCTC) fill_STAT_MCTC(fields []string) {
	dataLen := len(fields) - 1
	i := -1
	i++
	if i <= dataLen {
		s.TOTAL, _ = strconv.Atoi(fields[0])
	}
	i++
	if i <= dataLen { // these values seem to always be ints (or "NA")
		var value interface{}
		count, err := strconv.Atoi(fields[1])
		if err != nil {
			count = 0
		}
		s.CAT = make(map[string]interface{})
		for i1 := 1; i1 <= count; i1++ {
			for i2 := 1; i2 <= count; i2++ {
				// generate the particular key for the map i.e. F1_O1, F1_O2, F1_O3, F1_O4, F2_O1, F2_O2, F2_O3, F2_O4, etc.
				key := fmt.Sprintf("F%d_O%d", i1, i2)
				index := (i1-1)*count + i2
				if index >= len(fields) {
					value = "NA"
				} else {
					value, err = strconv.Atoi(fields[index])
				}
				if err != nil {
					value = "NA"
				}
				s.CAT[key] = value
			}
		}
	}
	i++
	if i <= dataLen {
		s.EC_VALUE, _ = strconv.ParseFloat(fields[3], 64)
	}
}

func (s *STAT_MCTS) fill_STAT_MCTS(fields []string) {
	dataLen := len(fields) - 1
	i := -1
	i++
	if i <= dataLen {
		s.TOTAL, _ = strconv.Atoi(fields[0])
	}
	i++
	if i <= dataLen {
		s.N_CAT, _ = strconv.Atoi(fields[1])
	}
	i++
	if i <= dataLen {
		s.ACC, _ = strconv.ParseFloat(fields[2], 64)
	}
	i++
	if i <= dataLen {
		s.ACC_NCL, _ = strconv.ParseFloat(fields[3], 64)
	}
	i++
	if i <= dataLen {
		s.ACC_NCU, _ = strconv.ParseFloat(fields[4], 64)
	}
	i++
	if i <= dataLen {
		s.ACC_BCL, _ = strconv.ParseFloat(fields[5], 64)
	}
	i++
	if i <= dataLen {
		s.ACC_BCU, _ = strconv.ParseFloat(fields[6], 64)
	}
	i++
	if i <= dataLen {
		s.HK, _ = strconv.ParseFloat(fields[7], 64)
	}
	i++
	if i <= dataLen {
		s.HK_BCL, _ = strconv.ParseFloat(fields[8], 64)
	}
	i++
	if i <= dataLen {
		s.HK_BCU, _ = strconv.ParseFloat(fields[9], 64)
	}
	i++
	if i <= dataLen {
		s.HSS, _ = strconv.ParseFloat(fields[10], 64)
	}
	i++
	if i <= dataLen {
		s.HSS_BCL, _ = strconv.ParseFloat(fields[11], 64)
	}
	i++
	if i <= dataLen {
		s.HSS_BCU, _ = strconv.ParseFloat(fields[12], 64)
	}
	i++
	if i <= dataLen {
		s.GER, _ = strconv.ParseFloat(fields[13], 64)
	}
	i++
	if i <= dataLen {
		s.GER_BCL, _ = strconv.ParseFloat(fields[14], 64)
	}
	i++
	if i <= dataLen {
		s.GER_BCU, _ = strconv.ParseFloat(fields[15], 64)
	}
	i++
	if i <= dataLen {
		s.HSS_EC, _ = strconv.ParseFloat(fields[16], 64)
	}
	i++
	if i <= dataLen {
		s.HSS_EC_BCL, _ = strconv.ParseFloat(fields[17], 64)
	}
	i++
	if i <= dataLen {
		s.HSS_EC_BCU, _ = strconv.ParseFloat(fields[18], 64)
	}
	i++
	if i <= dataLen {
		s.EC_VALUE, _ = strconv.ParseFloat(fields[19], 64)
	}
}

func (s *STAT_MPR) fill_STAT_MPR(fields []string) {
	dataLen := len(fields) - 1
	i := -1
	i++
	if i <= dataLen {
		s.TOTAL, _ = strconv.Atoi(fields[0])
	}
	i++
	if i <= dataLen {
		s.INDEX, _ = strconv.Atoi(fields[1])
	}
	i++
	if i <= dataLen {
		if fields[2] != "NA" {
			s.OBS_SID = fields[2]
		}
	}
	i++
	if i <= dataLen {
		s.OBS_LAT, _ = strconv.ParseFloat(fields[3], 64)
	}
	i++
	if i <= dataLen {
		s.OBS_LON, _ = strconv.ParseFloat(fields[4], 64)
	}
	i++
	if i <= dataLen {
		s.OBS_LVL, _ = strconv.ParseFloat(fields[5], 64)
	}
	i++
	if i <= dataLen {
		s.OBS_ELV, _ = strconv.ParseFloat(fields[6], 64)
	}
	i++
	if i <= dataLen {
		s.FCST, _ = strconv.ParseFloat(fields[7], 64)
	}
	i++
	if i <= dataLen {
		s.OBS, _ = strconv.ParseFloat(fields[8], 64)
	}
	i++
	if i <= dataLen {
		if fields[9] != "NA" {
			s.OBS_QC = fields[9]
		}
	}
	i++
	if i <= dataLen {
		s.CLIMO_MEAN, _ = strconv.ParseFloat(fields[10], 64)
	}
	i++
	if i <= dataLen {
		s.CLIMO_STDEV, _ = strconv.ParseFloat(fields[11], 64)
	}
	i++
	if i <= dataLen {
		s.CLIMO_CDF, _ = strconv.ParseFloat(fields[12], 64)
	}
}

func (s *STAT_NBRCNT) fill_STAT_NBRCNT(fields []string) {
	dataLen := len(fields) - 1
	i := -1
	i++
	if i <= dataLen {
		s.TOTAL, _ = strconv.Atoi(fields[0])
	}
	i++
	if i <= dataLen {
		s.FBS, _ = strconv.ParseFloat(fields[1], 64)
	}
	i++
	if i <= dataLen {
		s.FBS_BCL, _ = strconv.ParseFloat(fields[2], 64)
	}
	i++
	if i <= dataLen {
		s.FBS_BCU, _ = strconv.ParseFloat(fields[3], 64)
	}
	i++
	if i <= dataLen {
		s.FSS, _ = strconv.ParseFloat(fields[4], 64)
	}
	i++
	if i <= dataLen {
		s.FSS_BCL, _ = strconv.ParseFloat(fields[5], 64)
	}
	i++
	if i <= dataLen {
		s.FSS_BCU, _ = strconv.ParseFloat(fields[6], 64)
	}
	i++
	if i <= dataLen {
		s.AFSS, _ = strconv.ParseFloat(fields[7], 64)
	}
	i++
	if i <= dataLen {
		s.AFSS_BCL, _ = strconv.ParseFloat(fields[8], 64)
	}
	i++
	if i <= dataLen {
		s.AFSS_BCU, _ = strconv.ParseFloat(fields[9], 64)
	}
	i++
	if i <= dataLen {
		s.UFSS, _ = strconv.ParseFloat(fields[10], 64)
	}
	i++
	if i <= dataLen {
		s.UFSS_BCL, _ = strconv.ParseFloat(fields[11], 64)
	}
	i++
	if i <= dataLen {
		s.UFSS_BCU, _ = strconv.ParseFloat(fields[12], 64)
	}
	i++
	if i <= dataLen {
		s.F_RATE, _ = strconv.ParseFloat(fields[13], 64)
	}
	i++
	if i <= dataLen {
		s.F_RATE_BCL, _ = strconv.ParseFloat(fields[14], 64)
	}
	i++
	if i <= dataLen {
		s.F_RATE_BCU, _ = strconv.ParseFloat(fields[15], 64)
	}
	i++
	if i <= dataLen {
		s.O_RATE, _ = strconv.ParseFloat(fields[16], 64)
	}
	i++
	if i <= dataLen {
		s.O_RATE_BCL, _ = strconv.ParseFloat(fields[17], 64)
	}
	i++
	if i <= dataLen {
		s.O_RATE_BCU, _ = strconv.ParseFloat(fields[18], 64)
	}
}

func (s *STAT_NBRCTC) fill_STAT_NBRCTC(fields []string) {
	dataLen := len(fields) - 1
	i := -1
	i++
	if i <= dataLen {
		s.TOTAL, _ = strconv.Atoi(fields[0])
	}
	i++
	if i <= dataLen {
		s.FY_OY, _ = strconv.ParseFloat(fields[1], 64)
	}
	i++
	if i <= dataLen {
		s.FY_ON, _ = strconv.ParseFloat(fields[2], 64)
	}
	i++
	if i <= dataLen {
		s.FN_OY, _ = strconv.ParseFloat(fields[3], 64)
	}
	i++
	if i <= dataLen {
		s.FN_ON, _ = strconv.ParseFloat(fields[4], 64)
	}
}

func (s *STAT_NBRCTS) fill_STAT_NBRCTS(fields []string) {
	dataLen := len(fields) - 1
	i := -1
	i++
	if i <= dataLen {
		s.TOTAL, _ = strconv.Atoi(fields[0])
	}
	i++
	if i <= dataLen {
		s.BASER, _ = strconv.ParseFloat(fields[1], 64)
	}
	i++
	if i <= dataLen {
		s.BASER_NCL, _ = strconv.ParseFloat(fields[2], 64)
	}
	i++
	if i <= dataLen {
		s.BASER_NCU, _ = strconv.ParseFloat(fields[3], 64)
	}
	i++
	if i <= dataLen {
		s.BASER_BCL, _ = strconv.ParseFloat(fields[4], 64)
	}
	i++
	if i <= dataLen {
		s.BASER_BCU, _ = strconv.ParseFloat(fields[5], 64)
	}
	i++
	if i <= dataLen {
		s.FMEAN, _ = strconv.ParseFloat(fields[6], 64)
	}
	i++
	if i <= dataLen {
		s.FMEAN_NCL, _ = strconv.ParseFloat(fields[7], 64)
	}
	i++
	if i <= dataLen {
		s.FMEAN_NCU, _ = strconv.ParseFloat(fields[8], 64)
	}
	i++
	if i <= dataLen {
		s.FMEAN_BCL, _ = strconv.ParseFloat(fields[9], 64)
	}
	i++
	if i <= dataLen {
		s.FMEAN_BCU, _ = strconv.ParseFloat(fields[10], 64)
	}
	i++
	if i <= dataLen {
		s.ACC, _ = strconv.ParseFloat(fields[11], 64)
	}
	i++
	if i <= dataLen {
		s.ACC_NCL, _ = strconv.ParseFloat(fields[12], 64)
	}
	i++
	if i <= dataLen {
		s.ACC_NCU, _ = strconv.ParseFloat(fields[13], 64)
	}
	i++
	if i <= dataLen {
		s.ACC_BCL, _ = strconv.ParseFloat(fields[14], 64)
	}
	i++
	if i <= dataLen {
		s.ACC_BCU, _ = strconv.ParseFloat(fields[15], 64)
	}
	i++
	if i <= dataLen {
		s.FBIAS, _ = strconv.ParseFloat(fields[16], 64)
	}
	i++
	if i <= dataLen {
		s.FBIAS_BCL, _ = strconv.ParseFloat(fields[17], 64)
	}
	i++
	if i <= dataLen {
		s.FBIAS_BCU, _ = strconv.ParseFloat(fields[18], 64)
	}
	i++
	if i <= dataLen {
		s.PODY, _ = strconv.ParseFloat(fields[19], 64)
	}
	i++
	if i <= dataLen {
		s.PODY_NCL, _ = strconv.ParseFloat(fields[20], 64)
	}
	i++
	if i <= dataLen {
		s.PODY_NCU, _ = strconv.ParseFloat(fields[21], 64)
	}
	i++
	if i <= dataLen {
		s.PODY_BCL, _ = strconv.ParseFloat(fields[22], 64)
	}
	i++
	if i <= dataLen {
		s.PODY_BCU, _ = strconv.ParseFloat(fields[23], 64)
	}
	i++
	if i <= dataLen {
		s.PODN, _ = strconv.ParseFloat(fields[24], 64)
	}
	i++
	if i <= dataLen {
		s.PODN_NCL, _ = strconv.ParseFloat(fields[25], 64)
	}
	i++
	if i <= dataLen {
		s.PODN_NCU, _ = strconv.ParseFloat(fields[26], 64)
	}
	i++
	if i <= dataLen {
		s.PODN_BCL, _ = strconv.ParseFloat(fields[27], 64)
	}
	i++
	if i <= dataLen {
		s.PODN_BCU, _ = strconv.ParseFloat(fields[28], 64)
	}
	i++
	if i <= dataLen {
		s.POFD, _ = strconv.ParseFloat(fields[29], 64)
	}
	i++
	if i <= dataLen {
		s.POFD_NCL, _ = strconv.ParseFloat(fields[30], 64)
	}
	i++
	if i <= dataLen {
		s.POFD_NCU, _ = strconv.ParseFloat(fields[31], 64)
	}
	i++
	if i <= dataLen {
		s.POFD_BCL, _ = strconv.ParseFloat(fields[32], 64)
	}
	i++
	if i <= dataLen {
		s.POFD_BCU, _ = strconv.ParseFloat(fields[33], 64)
	}
	i++
	if i <= dataLen {
		s.FAR, _ = strconv.ParseFloat(fields[34], 64)
	}
	i++
	if i <= dataLen {
		s.FAR_NCL, _ = strconv.ParseFloat(fields[35], 64)
	}
	i++
	if i <= dataLen {
		s.FAR_NCU, _ = strconv.ParseFloat(fields[36], 64)
	}
	i++
	if i <= dataLen {
		s.FAR_BCL, _ = strconv.ParseFloat(fields[37], 64)
	}
	i++
	if i <= dataLen {
		s.FAR_BCU, _ = strconv.ParseFloat(fields[38], 64)
	}
	i++
	if i <= dataLen {
		s.CSI, _ = strconv.ParseFloat(fields[39], 64)
	}
	i++
	if i <= dataLen {
		s.CSI_NCL, _ = strconv.ParseFloat(fields[40], 64)
	}
	i++
	if i <= dataLen {
		s.CSI_NCU, _ = strconv.ParseFloat(fields[41], 64)
	}
	i++
	if i <= dataLen {
		s.CSI_BCL, _ = strconv.ParseFloat(fields[42], 64)
	}
	i++
	if i <= dataLen {
		s.CSI_BCU, _ = strconv.ParseFloat(fields[43], 64)
	}
	i++
	if i <= dataLen {
		s.GSS, _ = strconv.ParseFloat(fields[44], 64)
	}
	i++
	if i <= dataLen {
		s.GSS_BCL, _ = strconv.ParseFloat(fields[45], 64)
	}
	i++
	if i <= dataLen {
		s.GSS_BCU, _ = strconv.ParseFloat(fields[46], 64)
	}
	i++
	if i <= dataLen {
		s.HK, _ = strconv.ParseFloat(fields[47], 64)
	}
	i++
	if i <= dataLen {
		s.HK_NCL, _ = strconv.ParseFloat(fields[48], 64)
	}
	i++
	if i <= dataLen {
		s.HK_NCU, _ = strconv.ParseFloat(fields[49], 64)
	}
	i++
	if i <= dataLen {
		s.HK_BCL, _ = strconv.ParseFloat(fields[50], 64)
	}
	i++
	if i <= dataLen {
		s.HK_BCU, _ = strconv.ParseFloat(fields[51], 64)
	}
	i++
	if i <= dataLen {
		s.HSS, _ = strconv.ParseFloat(fields[52], 64)
	}
	i++
	if i <= dataLen {
		s.HSS_BCL, _ = strconv.ParseFloat(fields[53], 64)
	}
	i++
	if i <= dataLen {
		s.HSS_BCU, _ = strconv.ParseFloat(fields[54], 64)
	}
	i++
	if i <= dataLen {
		s.ODDS, _ = strconv.ParseFloat(fields[55], 64)
	}
	i++
	if i <= dataLen {
		s.ODDS_NCL, _ = strconv.ParseFloat(fields[56], 64)
	}
	i++
	if i <= dataLen {
		s.ODDS_NCU, _ = strconv.ParseFloat(fields[57], 64)
	}
	i++
	if i <= dataLen {
		s.ODDS_BCL, _ = strconv.ParseFloat(fields[58], 64)
	}
	i++
	if i <= dataLen {
		s.ODDS_BCU, _ = strconv.ParseFloat(fields[59], 64)
	}
	i++
	if i <= dataLen {
		s.LODDS, _ = strconv.ParseFloat(fields[60], 64)
	}
	i++
	if i <= dataLen {
		s.LODDS_NCL, _ = strconv.ParseFloat(fields[61], 64)
	}
	i++
	if i <= dataLen {
		s.LODDS_NCU, _ = strconv.ParseFloat(fields[62], 64)
	}
	i++
	if i <= dataLen {
		s.LODDS_BCL, _ = strconv.ParseFloat(fields[63], 64)
	}
	i++
	if i <= dataLen {
		s.LODDS_BCU, _ = strconv.ParseFloat(fields[64], 64)
	}
	i++
	if i <= dataLen {
		s.ORSS, _ = strconv.ParseFloat(fields[65], 64)
	}
	i++
	if i <= dataLen {
		s.ORSS_NCL, _ = strconv.ParseFloat(fields[66], 64)
	}
	i++
	if i <= dataLen {
		s.ORSS_NCU, _ = strconv.ParseFloat(fields[67], 64)
	}
	i++
	if i <= dataLen {
		s.ORSS_BCL, _ = strconv.ParseFloat(fields[68], 64)
	}
	i++
	if i <= dataLen {
		s.ORSS_BCU, _ = strconv.ParseFloat(fields[69], 64)
	}
	i++
	if i <= dataLen {
		s.EDS, _ = strconv.ParseFloat(fields[70], 64)
	}
	i++
	if i <= dataLen {
		s.EDS_NCL, _ = strconv.ParseFloat(fields[71], 64)
	}
	i++
	if i <= dataLen {
		s.EDS_NCU, _ = strconv.ParseFloat(fields[72], 64)
	}
	i++
	if i <= dataLen {
		s.EDS_BCL, _ = strconv.ParseFloat(fields[73], 64)
	}
	i++
	if i <= dataLen {
		s.EDS_BCU, _ = strconv.ParseFloat(fields[74], 64)
	}
	i++
	if i <= dataLen {
		s.SEDS, _ = strconv.ParseFloat(fields[75], 64)
	}
	i++
	if i <= dataLen {
		s.SEDS_NCL, _ = strconv.ParseFloat(fields[76], 64)
	}
	i++
	if i <= dataLen {
		s.SEDS_NCU, _ = strconv.ParseFloat(fields[77], 64)
	}
	i++
	if i <= dataLen {
		s.SEDS_BCL, _ = strconv.ParseFloat(fields[78], 64)
	}
	i++
	if i <= dataLen {
		s.SEDS_BCU, _ = strconv.ParseFloat(fields[79], 64)
	}
	i++
	if i <= dataLen {
		s.EDI, _ = strconv.ParseFloat(fields[80], 64)
	}
	i++
	if i <= dataLen {
		s.EDI_NCL, _ = strconv.ParseFloat(fields[81], 64)
	}
	i++
	if i <= dataLen {
		s.EDI_NCU, _ = strconv.ParseFloat(fields[82], 64)
	}
	i++
	if i <= dataLen {
		s.EDI_BCL, _ = strconv.ParseFloat(fields[83], 64)
	}
	i++
	if i <= dataLen {
		s.EDI_BCU, _ = strconv.ParseFloat(fields[84], 64)
	}
	i++
	if i <= dataLen {
		s.SEDI, _ = strconv.ParseFloat(fields[85], 64)
	}
	i++
	if i <= dataLen {
		s.SEDI_NCL, _ = strconv.ParseFloat(fields[86], 64)
	}
	i++
	if i <= dataLen {
		s.SEDI_NCU, _ = strconv.ParseFloat(fields[87], 64)
	}
	i++
	if i <= dataLen {
		s.SEDI_BCL, _ = strconv.ParseFloat(fields[88], 64)
	}
	i++
	if i <= dataLen {
		s.SEDI_BCU, _ = strconv.ParseFloat(fields[89], 64)
	}
	i++
	if i <= dataLen {
		s.BAGSS, _ = strconv.ParseFloat(fields[90], 64)
	}
	i++
	if i <= dataLen {
		s.BAGSS_BCL, _ = strconv.ParseFloat(fields[91], 64)
	}
	i++
	if i <= dataLen {
		s.BAGSS_BCU, _ = strconv.ParseFloat(fields[92], 64)
	}
}

func (s *STAT_ORANK) fill_STAT_ORANK(fields []string) {
	dataLen := len(fields) - 1
	i := -1
	i++
	if i <= dataLen {
		s.TOTAL, _ = strconv.Atoi(fields[0])
	}
	i++
	if i <= dataLen {
		s.INDEX, _ = strconv.Atoi(fields[1])
	}
	i++
	if i <= dataLen {
		if fields[2] != "NA" {
			s.OBS_SID = fields[2]
		}
	}
	i++
	if i <= dataLen {
		s.OBS_LAT, _ = strconv.ParseFloat(fields[3], 64)
	}
	i++
	if i <= dataLen {
		s.OBS_LON, _ = strconv.ParseFloat(fields[4], 64)
	}
	i++
	if i <= dataLen {
		s.OBS_LVL, _ = strconv.ParseFloat(fields[5], 64)
	}
	i++
	if i <= dataLen {
		s.OBS_ELV, _ = strconv.ParseFloat(fields[6], 64)
	}
	i++
	if i <= dataLen {
		s.OBS, _ = strconv.ParseFloat(fields[7], 64)
	}
	i++
	if i <= dataLen {
		s.PIT, _ = strconv.ParseFloat(fields[8], 64)
	}
	i++
	if i <= dataLen {
		s.RANK, _ = strconv.Atoi(fields[9])
	}
	i++
	if i <= dataLen {
		s.N_ENS_VLD, _ = strconv.Atoi(fields[10])
	}
	i++
	if i <= dataLen { // the first field of the repeating fields is the TOTAL, the second field is the 1st dimenSion of the 1st sequence (there might be only one sequence)
		var value interface{}
		count, err := strconv.Atoi(fields[11])
		if err != nil {
			count = 0
		}
		keyPrefixes := []string{"ENS_"}
		s.ENS = make(map[string]interface{})
		for group := 1; group <= count; group++ {
			for index := 12; index <= len(keyPrefixes); index++ {
				key := fmt.Sprintf("%s_%d", keyPrefixes[index-1], index)
				if index > len(fields) { // sometimes the data line is truncated - we will set expected data to "NA"
					value = "NA"
				} else {
					value, err = strconv.Atoi(fields[index])
					if err != nil { // sometimes there can be these NA values in the data, which will be left out of json
						value = "NA"
					}
				}
				s.ENS[key] = value
			}
		}
	}
	i++
	if i <= dataLen {
		if fields[13] != "NA" {
			s.OBS_QC = fields[13]
		}
	}
	i++
	if i <= dataLen {
		s.ENS_MEAN, _ = strconv.Atoi(fields[14])
	}
	i++
	if i <= dataLen {
		s.CLIMO_MEAN, _ = strconv.ParseFloat(fields[15], 64)
	}
	i++
	if i <= dataLen {
		s.SPREAD, _ = strconv.ParseFloat(fields[16], 64)
	}
	i++
	if i <= dataLen {
		s.ENS_MEAN_OERR, _ = strconv.Atoi(fields[17])
	}
	i++
	if i <= dataLen {
		s.SPREAD_OERR, _ = strconv.ParseFloat(fields[18], 64)
	}
	i++
	if i <= dataLen {
		s.SPREAD_PLUS_OERR, _ = strconv.ParseFloat(fields[19], 64)
	}
	i++
	if i <= dataLen {
		s.CLIMO_STDEV, _ = strconv.ParseFloat(fields[20], 64)
	}
}

func (s *STAT_PCT) fill_STAT_PCT(fields []string) {
	dataLen := len(fields) - 1
	i := -1
	i++
	if i <= dataLen {
		s.TOTAL, _ = strconv.Atoi(fields[0])
	}
	i++
	if i <= dataLen { // the first field of the repeating fields is the TOTAL, the second field is the 1st dimenSion of the 1st sequence (there might be only one sequence)
		var value interface{}
		count, err := strconv.Atoi(fields[1])
		if err != nil {
			count = 0
		}
		keyPrefixes := []string{"THRESH_", "OY_", "ON_"}
		s.THRESH = make(map[string]interface{})
		for group := 1; group <= count; group++ {
			for index := 2; index <= len(keyPrefixes); index++ {
				key := fmt.Sprintf("%s_%d", keyPrefixes[index-1], index)
				if index > len(fields) { // sometimes the data line is truncated - we will set expected data to "NA"
					value = "NA"
				} else {
					value, err = strconv.ParseFloat(fields[index], 64)
					if err != nil { // sometimes there can be these NA values in the data, which will be left out of json
						value = "NA"
					}
				}
				s.THRESH[key] = value
			}
		}
	}
}

func (s *STAT_PHIST) fill_STAT_PHIST(fields []string) {
	dataLen := len(fields) - 1
	i := -1
	i++
	if i <= dataLen {
		s.TOTAL, _ = strconv.Atoi(fields[0])
	}
	i++
	if i <= dataLen {
		s.BIN_SIZE, _ = strconv.Atoi(fields[1])
	}
	i++
	if i <= dataLen { // the first field of the repeating fields is the TOTAL, the second field is the 1st dimenSion of the 1st sequence (there might be only one sequence)
		var value interface{}
		count, err := strconv.Atoi(fields[2])
		if err != nil {
			count = 0
		}
		keyPrefixes := []string{"BIN_"}
		s.BIN = make(map[string]interface{})
		for group := 1; group <= count; group++ {
			for index := 3; index <= len(keyPrefixes); index++ {
				key := fmt.Sprintf("%s_%d", keyPrefixes[index-1], index)
				if index > len(fields) { // sometimes the data line is truncated - we will set expected data to "NA"
					value = "NA"
				} else {
					value, err = strconv.Atoi(fields[index])
					if err != nil { // sometimes there can be these NA values in the data, which will be left out of json
						value = "NA"
					}
				}
				s.BIN[key] = value
			}
		}
	}
}

func (s *STAT_PJC) fill_STAT_PJC(fields []string) {
	dataLen := len(fields) - 1
	i := -1
	i++
	if i <= dataLen {
		s.TOTAL, _ = strconv.Atoi(fields[0])
	}
	i++
	if i <= dataLen { // the first field of the repeating fields is the TOTAL, the second field is the 1st dimenSion of the 1st sequence (there might be only one sequence)
		var value interface{}
		count, err := strconv.Atoi(fields[1])
		if err != nil {
			count = 0
		}
		keyPrefixes := []string{"THRESH_", "OY_TP_", "ON_TP_", "CALIBRATION_", "REFINEMENT", "LIKELIHOOD_", "BASER_"}
		s.THRESH = make(map[string]interface{})
		for group := 1; group <= count; group++ {
			for index := 2; index <= len(keyPrefixes); index++ {
				key := fmt.Sprintf("%s_%d", keyPrefixes[index-1], index)
				if index > len(fields) { // sometimes the data line is truncated - we will set expected data to "NA"
					value = "NA"
				} else {
					value, err = strconv.ParseFloat(fields[index], 64)
					if err != nil { // sometimes there can be these NA values in the data, which will be left out of json
						value = "NA"
					}
				}
				s.THRESH[key] = value
			}
		}
	}
}

func (s *STAT_PRC) fill_STAT_PRC(fields []string) {
	dataLen := len(fields) - 1
	i := -1
	i++
	if i <= dataLen {
		s.TOTAL, _ = strconv.Atoi(fields[0])
	}
	i++
	if i <= dataLen { // the first field of the repeating fields is the TOTAL, the second field is the 1st dimenSion of the 1st sequence (there might be only one sequence)
		var value interface{}
		count, err := strconv.Atoi(fields[1])
		if err != nil {
			count = 0
		}
		keyPrefixes := []string{"THRESH_", "PODY_", "POFD_"}
		s.THRESH = make(map[string]interface{})
		for group := 1; group <= count; group++ {
			for index := 2; index <= len(keyPrefixes); index++ {
				key := fmt.Sprintf("%s_%d", keyPrefixes[index-1], index)
				if index > len(fields) { // sometimes the data line is truncated - we will set expected data to "NA"
					value = "NA"
				} else {
					value, err = strconv.ParseFloat(fields[index], 64)
					if err != nil { // sometimes there can be these NA values in the data, which will be left out of json
						value = "NA"
					}
				}
				s.THRESH[key] = value
			}
		}
	}
}

func (s *STAT_PSTD) fill_STAT_PSTD(fields []string) {
	dataLen := len(fields) - 1
	i := -1
	i++
	if i <= dataLen {
		s.TOTAL, _ = strconv.Atoi(fields[0])
	}
	i++
	if i <= dataLen { // the first field of the repeating fields is the TOTAL, the second field is the 1st dimenSion of the 1st sequence (there might be only one sequence)
		var value interface{}
		count, err := strconv.Atoi(fields[1])
		if err != nil {
			count = 0
		}
		keyPrefixes := []string{"THRESH_"}
		s.THRESH = make(map[string]interface{})
		for group := 1; group <= count; group++ {
			for index := 2; index <= len(keyPrefixes); index++ {
				key := fmt.Sprintf("%s_%d", keyPrefixes[index-1], index)
				if index > len(fields) { // sometimes the data line is truncated - we will set expected data to "NA"
					value = "NA"
				} else {
					value, err = strconv.ParseFloat(fields[index], 64)
					if err != nil { // sometimes there can be these NA values in the data, which will be left out of json
						value = "NA"
					}
				}
				s.THRESH[key] = value
			}
		}
	}
	i++
	if i <= dataLen {
		s.BASER_NCL, _ = strconv.ParseFloat(fields[3], 64)
	}
	i++
	if i <= dataLen {
		s.BASER_NCU, _ = strconv.ParseFloat(fields[4], 64)
	}
	i++
	if i <= dataLen {
		s.RELIABILITY, _ = strconv.ParseFloat(fields[5], 64)
	}
	i++
	if i <= dataLen {
		s.RESOLUTION, _ = strconv.ParseFloat(fields[6], 64)
	}
	i++
	if i <= dataLen {
		s.UNCERTAINTY, _ = strconv.ParseFloat(fields[7], 64)
	}
	i++
	if i <= dataLen {
		s.ROC_AUC, _ = strconv.ParseFloat(fields[8], 64)
	}
	i++
	if i <= dataLen {
		s.BRIER, _ = strconv.ParseFloat(fields[9], 64)
	}
	i++
	if i <= dataLen {
		s.BRIER_NCL, _ = strconv.ParseFloat(fields[10], 64)
	}
	i++
	if i <= dataLen {
		s.BRIER_NCU, _ = strconv.ParseFloat(fields[11], 64)
	}
	i++
	if i <= dataLen {
		s.BRIERCL, _ = strconv.ParseFloat(fields[12], 64)
	}
	i++
	if i <= dataLen {
		s.BRIERCL_NCL, _ = strconv.ParseFloat(fields[13], 64)
	}
	i++
	if i <= dataLen {
		s.BRIERCL_NCU, _ = strconv.ParseFloat(fields[14], 64)
	}
	i++
	if i <= dataLen {
		s.BSS, _ = strconv.ParseFloat(fields[15], 64)
	}
	i++
	if i <= dataLen {
		s.BSS_SMPL, _ = strconv.ParseFloat(fields[16], 64)
	}
	i++
	if i <= dataLen {
		s.THRESH_I, _ = strconv.Atoi(fields[17])
	}
}

func (s *STAT_RELP) fill_STAT_RELP(fields []string) {
	dataLen := len(fields) - 1
	i := -1
	i++
	if i <= dataLen {
		s.TOTAL, _ = strconv.Atoi(fields[0])
	}
	i++
	if i <= dataLen { // the first field of the repeating fields is the TOTAL, the second field is the 1st dimenSion of the 1st sequence (there might be only one sequence)
		var value interface{}
		count, err := strconv.Atoi(fields[1])
		if err != nil {
			count = 0
		}
		keyPrefixes := []string{"RELP_"}
		s.ENS = make(map[string]interface{})
		for group := 1; group <= count; group++ {
			for index := 2; index <= len(keyPrefixes); index++ {
				key := fmt.Sprintf("%s_%d", keyPrefixes[index-1], index)
				if index > len(fields) { // sometimes the data line is truncated - we will set expected data to "NA"
					value = "NA"
				} else {
					value, err = strconv.ParseFloat(fields[index], 64)
					if err != nil { // sometimes there can be these NA values in the data, which will be left out of json
						value = "NA"
					}
				}
				s.ENS[key] = value
			}
		}
	}
}

func (s *STAT_RHIST) fill_STAT_RHIST(fields []string) {
	dataLen := len(fields) - 1
	i := -1
	i++
	if i <= dataLen {
		s.TOTAL, _ = strconv.Atoi(fields[0])
	}
	i++
	if i <= dataLen { // the first field of the repeating fields is the TOTAL, the second field is the 1st dimenSion of the 1st sequence (there might be only one sequence)
		var value interface{}
		count, err := strconv.Atoi(fields[1])
		if err != nil {
			count = 0
		}
		keyPrefixes := []string{"RANK_"}
		s.RANK = make(map[string]interface{})
		for group := 1; group <= count; group++ {
			for index := 2; index <= len(keyPrefixes); index++ {
				key := fmt.Sprintf("%s_%d", keyPrefixes[index-1], index)
				if index > len(fields) { // sometimes the data line is truncated - we will set expected data to "NA"
					value = "NA"
				} else {
					value, err = strconv.Atoi(fields[index])
					if err != nil { // sometimes there can be these NA values in the data, which will be left out of json
						value = "NA"
					}
				}
				s.RANK[key] = value
			}
		}
	}
}

func (s *STAT_RPS) fill_STAT_RPS(fields []string) {
	dataLen := len(fields) - 1
	i := -1
	i++
	if i <= dataLen {
		s.TOTAL, _ = strconv.Atoi(fields[0])
	}
	i++
	if i <= dataLen {
		s.N_PROB, _ = strconv.Atoi(fields[1])
	}
	i++
	if i <= dataLen {
		s.RPS_REL, _ = strconv.ParseFloat(fields[2], 64)
	}
	i++
	if i <= dataLen {
		s.RPS_RES, _ = strconv.ParseFloat(fields[3], 64)
	}
	i++
	if i <= dataLen {
		s.RPS_UNC, _ = strconv.ParseFloat(fields[4], 64)
	}
	i++
	if i <= dataLen {
		s.RPS, _ = strconv.ParseFloat(fields[5], 64)
	}
	i++
	if i <= dataLen {
		s.RPSS, _ = strconv.ParseFloat(fields[6], 64)
	}
	i++
	if i <= dataLen {
		s.RPSS_SMPL, _ = strconv.ParseFloat(fields[7], 64)
	}
	i++
	if i <= dataLen {
		s.RPS_COMP, _ = strconv.ParseFloat(fields[8], 64)
	}
}

func (s *STAT_SAL1L2) fill_STAT_SAL1L2(fields []string) {
	dataLen := len(fields) - 1
	i := -1
	i++
	if i <= dataLen {
		s.TOTAL, _ = strconv.Atoi(fields[0])
	}
	i++
	if i <= dataLen {
		s.FABAR, _ = strconv.ParseFloat(fields[1], 64)
	}
	i++
	if i <= dataLen {
		s.OABAR, _ = strconv.ParseFloat(fields[2], 64)
	}
	i++
	if i <= dataLen {
		s.FOABAR, _ = strconv.ParseFloat(fields[3], 64)
	}
	i++
	if i <= dataLen {
		s.FFABAR, _ = strconv.ParseFloat(fields[4], 64)
	}
	i++
	if i <= dataLen {
		s.OOABAR, _ = strconv.ParseFloat(fields[5], 64)
	}
	i++
	if i <= dataLen {
		s.MAE, _ = strconv.ParseFloat(fields[6], 64)
	}
}

func (s *STAT_SEEPS) fill_STAT_SEEPS(fields []string) {
	dataLen := len(fields) - 1
	i := -1
	i++
	if i <= dataLen {
		s.TOTAL, _ = strconv.Atoi(fields[0])
	}
	i++
	if i <= dataLen {
		if fields[1] != "NA" {
			s.S12 = fields[1]
		}
	}
	i++
	if i <= dataLen {
		if fields[2] != "NA" {
			s.S13 = fields[2]
		}
	}
	i++
	if i <= dataLen {
		if fields[3] != "NA" {
			s.S21 = fields[3]
		}
	}
	i++
	if i <= dataLen {
		if fields[4] != "NA" {
			s.S23 = fields[4]
		}
	}
	i++
	if i <= dataLen {
		if fields[5] != "NA" {
			s.S31 = fields[5]
		}
	}
	i++
	if i <= dataLen {
		if fields[6] != "NA" {
			s.S32 = fields[6]
		}
	}
	i++
	if i <= dataLen {
		s.PF1, _ = strconv.ParseFloat(fields[7], 64)
	}
	i++
	if i <= dataLen {
		s.PF2, _ = strconv.ParseFloat(fields[8], 64)
	}
	i++
	if i <= dataLen {
		s.PF3, _ = strconv.ParseFloat(fields[9], 64)
	}
	i++
	if i <= dataLen {
		s.PV1, _ = strconv.ParseFloat(fields[10], 64)
	}
	i++
	if i <= dataLen {
		s.PV2, _ = strconv.ParseFloat(fields[11], 64)
	}
	i++
	if i <= dataLen {
		s.PV3, _ = strconv.ParseFloat(fields[12], 64)
	}
	i++
	if i <= dataLen {
		s.MEAN_FCST, _ = strconv.ParseFloat(fields[13], 64)
	}
	i++
	if i <= dataLen {
		s.MEAN_OBS, _ = strconv.ParseFloat(fields[14], 64)
	}
	i++
	if i <= dataLen {
		s.SEEPS, _ = strconv.ParseFloat(fields[15], 64)
	}
}

func (s *STAT_SEEPS_MPR) fill_STAT_SEEPS_MPR(fields []string) {
	dataLen := len(fields) - 1
	i := -1
	i++
	if i <= dataLen {
		if fields[0] != "NA" {
			s.OBS_SID = fields[0]
		}
	}
	i++
	if i <= dataLen {
		s.OBS_LAT, _ = strconv.ParseFloat(fields[1], 64)
	}
	i++
	if i <= dataLen {
		s.OBS_LON, _ = strconv.ParseFloat(fields[2], 64)
	}
	i++
	if i <= dataLen {
		s.FCST, _ = strconv.ParseFloat(fields[3], 64)
	}
	i++
	if i <= dataLen {
		s.OBS, _ = strconv.ParseFloat(fields[4], 64)
	}
	i++
	if i <= dataLen {
		if fields[5] != "NA" {
			s.OBS_QC = fields[5]
		}
	}
	i++
	if i <= dataLen {
		s.FCST_CAT, _ = strconv.Atoi(fields[6])
	}
	i++
	if i <= dataLen {
		s.OBS_CAT, _ = strconv.Atoi(fields[7])
	}
	i++
	if i <= dataLen {
		s.P1, _ = strconv.ParseFloat(fields[8], 64)
	}
	i++
	if i <= dataLen {
		s.P2, _ = strconv.ParseFloat(fields[9], 64)
	}
	i++
	if i <= dataLen {
		s.T1, _ = strconv.ParseFloat(fields[10], 64)
	}
	i++
	if i <= dataLen {
		s.T2, _ = strconv.ParseFloat(fields[11], 64)
	}
	i++
	if i <= dataLen {
		s.SEEPS, _ = strconv.ParseFloat(fields[12], 64)
	}
}

func (s *STAT_SL1L2) fill_STAT_SL1L2(fields []string) {
	dataLen := len(fields) - 1
	i := -1
	i++
	if i <= dataLen {
		s.TOTAL, _ = strconv.Atoi(fields[0])
	}
	i++
	if i <= dataLen {
		s.FBAR, _ = strconv.ParseFloat(fields[1], 64)
	}
	i++
	if i <= dataLen {
		s.OBAR, _ = strconv.ParseFloat(fields[2], 64)
	}
	i++
	if i <= dataLen {
		s.FOBAR, _ = strconv.ParseFloat(fields[3], 64)
	}
	i++
	if i <= dataLen {
		s.FFBAR, _ = strconv.ParseFloat(fields[4], 64)
	}
	i++
	if i <= dataLen {
		s.OOBAR, _ = strconv.ParseFloat(fields[5], 64)
	}
	i++
	if i <= dataLen {
		s.MAE, _ = strconv.ParseFloat(fields[6], 64)
	}
}

func (s *STAT_SSIDX) fill_STAT_SSIDX(fields []string) {
	dataLen := len(fields) - 1
	i := -1
	i++
	if i <= dataLen {
		if fields[0] != "NA" {
			s.FCST_MODEL = fields[0]
		}
	}
	i++
	if i <= dataLen {
		if fields[1] != "NA" {
			s.REF_MODEL = fields[1]
		}
	}
	i++
	if i <= dataLen {
		s.N_INIT, _ = strconv.Atoi(fields[2])
	}
	i++
	if i <= dataLen {
		s.N_TERM, _ = strconv.Atoi(fields[3])
	}
	i++
	if i <= dataLen {
		s.N_VLD, _ = strconv.Atoi(fields[4])
	}
	i++
	if i <= dataLen {
		s.SS_INDEX, _ = strconv.ParseFloat(fields[5], 64)
	}
}

func (s *STAT_SSVAR) fill_STAT_SSVAR(fields []string) {
	dataLen := len(fields) - 1
	i := -1
	i++
	if i <= dataLen {
		s.TOTAL, _ = strconv.Atoi(fields[0])
	}
	i++
	if i <= dataLen {
		s.N_BIN, _ = strconv.Atoi(fields[1])
	}
	i++
	if i <= dataLen {
		s.BIN_I, _ = strconv.Atoi(fields[2])
	}
	i++
	if i <= dataLen {
		s.BIN_N, _ = strconv.Atoi(fields[3])
	}
	i++
	if i <= dataLen {
		s.VAR_MIN, _ = strconv.ParseFloat(fields[4], 64)
	}
	i++
	if i <= dataLen {
		s.VAR_MAX, _ = strconv.ParseFloat(fields[5], 64)
	}
	i++
	if i <= dataLen {
		s.VAR_MEAN, _ = strconv.ParseFloat(fields[6], 64)
	}
	i++
	if i <= dataLen {
		s.FBAR, _ = strconv.ParseFloat(fields[7], 64)
	}
	i++
	if i <= dataLen {
		s.OBAR, _ = strconv.ParseFloat(fields[8], 64)
	}
	i++
	if i <= dataLen {
		s.FOBAR, _ = strconv.ParseFloat(fields[9], 64)
	}
	i++
	if i <= dataLen {
		s.FFBAR, _ = strconv.ParseFloat(fields[10], 64)
	}
	i++
	if i <= dataLen {
		s.OOBAR, _ = strconv.ParseFloat(fields[11], 64)
	}
	i++
	if i <= dataLen {
		s.FBAR_NCL, _ = strconv.ParseFloat(fields[12], 64)
	}
	i++
	if i <= dataLen {
		s.FBAR_NCU, _ = strconv.ParseFloat(fields[13], 64)
	}
	i++
	if i <= dataLen {
		s.FSTDEV, _ = strconv.ParseFloat(fields[14], 64)
	}
	i++
	if i <= dataLen {
		s.FSTDEV_NCL, _ = strconv.ParseFloat(fields[15], 64)
	}
	i++
	if i <= dataLen {
		s.FSTDEV_NCU, _ = strconv.ParseFloat(fields[16], 64)
	}
	i++
	if i <= dataLen {
		s.OBAR_NCL, _ = strconv.ParseFloat(fields[17], 64)
	}
	i++
	if i <= dataLen {
		s.OBAR_NCU, _ = strconv.ParseFloat(fields[18], 64)
	}
	i++
	if i <= dataLen {
		s.OSTDEV, _ = strconv.ParseFloat(fields[19], 64)
	}
	i++
	if i <= dataLen {
		s.OSTDEV_NCL, _ = strconv.ParseFloat(fields[20], 64)
	}
	i++
	if i <= dataLen {
		s.OSTDEV_NCU, _ = strconv.ParseFloat(fields[21], 64)
	}
	i++
	if i <= dataLen {
		s.PR_CORR, _ = strconv.ParseFloat(fields[22], 64)
	}
	i++
	if i <= dataLen {
		s.PR_CORR_NCL, _ = strconv.ParseFloat(fields[23], 64)
	}
	i++
	if i <= dataLen {
		s.PR_CORR_NCU, _ = strconv.ParseFloat(fields[24], 64)
	}
	i++
	if i <= dataLen {
		s.ME, _ = strconv.ParseFloat(fields[25], 64)
	}
	i++
	if i <= dataLen {
		s.ME_NCL, _ = strconv.ParseFloat(fields[26], 64)
	}
	i++
	if i <= dataLen {
		s.ME_NCU, _ = strconv.ParseFloat(fields[27], 64)
	}
	i++
	if i <= dataLen {
		s.ESTDEV, _ = strconv.ParseFloat(fields[28], 64)
	}
	i++
	if i <= dataLen {
		s.ESTDEV_NCL, _ = strconv.ParseFloat(fields[29], 64)
	}
	i++
	if i <= dataLen {
		s.ESTDEV_NCU, _ = strconv.ParseFloat(fields[30], 64)
	}
	i++
	if i <= dataLen {
		s.MBIAS, _ = strconv.ParseFloat(fields[31], 64)
	}
	i++
	if i <= dataLen {
		s.MSE, _ = strconv.ParseFloat(fields[32], 64)
	}
	i++
	if i <= dataLen {
		s.BCMSE, _ = strconv.ParseFloat(fields[33], 64)
	}
	i++
	if i <= dataLen {
		s.RMSE, _ = strconv.ParseFloat(fields[34], 64)
	}
}

func (s *STAT_VAL1L2) fill_STAT_VAL1L2(fields []string) {
	dataLen := len(fields) - 1
	i := -1
	i++
	if i <= dataLen {
		s.TOTAL, _ = strconv.Atoi(fields[0])
	}
	i++
	if i <= dataLen {
		s.UFABAR, _ = strconv.ParseFloat(fields[1], 64)
	}
	i++
	if i <= dataLen {
		s.VFABAR, _ = strconv.ParseFloat(fields[2], 64)
	}
	i++
	if i <= dataLen {
		s.UOABAR, _ = strconv.ParseFloat(fields[3], 64)
	}
	i++
	if i <= dataLen {
		s.VOABAR, _ = strconv.ParseFloat(fields[4], 64)
	}
	i++
	if i <= dataLen {
		s.UVFOABAR, _ = strconv.ParseFloat(fields[5], 64)
	}
	i++
	if i <= dataLen {
		s.UVFFABAR, _ = strconv.ParseFloat(fields[6], 64)
	}
	i++
	if i <= dataLen {
		s.UVOOABAR, _ = strconv.ParseFloat(fields[7], 64)
	}
	i++
	if i <= dataLen {
		s.FA_SPEED_BAR, _ = strconv.ParseFloat(fields[8], 64)
	}
	i++
	if i <= dataLen {
		s.OA_SPEED_BAR, _ = strconv.ParseFloat(fields[9], 64)
	}
}

func (s *STAT_VCNT) fill_STAT_VCNT(fields []string) {
	dataLen := len(fields) - 1
	i := -1
	i++
	if i <= dataLen {
		s.TOTAL, _ = strconv.Atoi(fields[0])
	}
	i++
	if i <= dataLen {
		s.FBAR, _ = strconv.ParseFloat(fields[1], 64)
	}
	i++
	if i <= dataLen {
		s.FBAR_BCL, _ = strconv.ParseFloat(fields[2], 64)
	}
	i++
	if i <= dataLen {
		s.FBAR_BCU, _ = strconv.ParseFloat(fields[3], 64)
	}
	i++
	if i <= dataLen {
		s.OBAR, _ = strconv.ParseFloat(fields[4], 64)
	}
	i++
	if i <= dataLen {
		s.OBAR_BCL, _ = strconv.ParseFloat(fields[5], 64)
	}
	i++
	if i <= dataLen {
		s.OBAR_BCU, _ = strconv.ParseFloat(fields[6], 64)
	}
	i++
	if i <= dataLen {
		s.FS_RMS, _ = strconv.ParseFloat(fields[7], 64)
	}
	i++
	if i <= dataLen {
		s.FS_RMS_BCL, _ = strconv.ParseFloat(fields[8], 64)
	}
	i++
	if i <= dataLen {
		s.FS_RMS_BCU, _ = strconv.ParseFloat(fields[9], 64)
	}
	i++
	if i <= dataLen {
		s.OS_RMS, _ = strconv.ParseFloat(fields[10], 64)
	}
	i++
	if i <= dataLen {
		s.OS_RMS_BCL, _ = strconv.ParseFloat(fields[11], 64)
	}
	i++
	if i <= dataLen {
		s.OS_RMS_BCU, _ = strconv.ParseFloat(fields[12], 64)
	}
	i++
	if i <= dataLen {
		s.MSVE, _ = strconv.ParseFloat(fields[13], 64)
	}
	i++
	if i <= dataLen {
		s.MSVE_BCL, _ = strconv.ParseFloat(fields[14], 64)
	}
	i++
	if i <= dataLen {
		s.MSVE_BCU, _ = strconv.ParseFloat(fields[15], 64)
	}
	i++
	if i <= dataLen {
		s.RMSVE, _ = strconv.ParseFloat(fields[16], 64)
	}
	i++
	if i <= dataLen {
		s.RMSVE_BCL, _ = strconv.ParseFloat(fields[17], 64)
	}
	i++
	if i <= dataLen {
		s.RMSVE_BCU, _ = strconv.ParseFloat(fields[18], 64)
	}
	i++
	if i <= dataLen {
		s.FSTDEV, _ = strconv.ParseFloat(fields[19], 64)
	}
	i++
	if i <= dataLen {
		s.FSTDEV_BCL, _ = strconv.ParseFloat(fields[20], 64)
	}
	i++
	if i <= dataLen {
		s.FSTDEV_BCU, _ = strconv.ParseFloat(fields[21], 64)
	}
	i++
	if i <= dataLen {
		s.OSTDEV, _ = strconv.ParseFloat(fields[22], 64)
	}
	i++
	if i <= dataLen {
		s.OSTDEV_BCL, _ = strconv.ParseFloat(fields[23], 64)
	}
	i++
	if i <= dataLen {
		s.OSTDEV_BCU, _ = strconv.ParseFloat(fields[24], 64)
	}
	i++
	if i <= dataLen {
		s.FDIR, _ = strconv.ParseFloat(fields[25], 64)
	}
	i++
	if i <= dataLen {
		s.FDIR_BCL, _ = strconv.ParseFloat(fields[26], 64)
	}
	i++
	if i <= dataLen {
		s.FDIR_BCU, _ = strconv.ParseFloat(fields[27], 64)
	}
	i++
	if i <= dataLen {
		s.ODIR, _ = strconv.ParseFloat(fields[28], 64)
	}
	i++
	if i <= dataLen {
		s.ODIR_BCL, _ = strconv.ParseFloat(fields[29], 64)
	}
	i++
	if i <= dataLen {
		s.ODIR_BCU, _ = strconv.ParseFloat(fields[30], 64)
	}
	i++
	if i <= dataLen {
		s.FBAR_SPEED, _ = strconv.ParseFloat(fields[31], 64)
	}
	i++
	if i <= dataLen {
		s.FBAR_SPEED_BCL, _ = strconv.ParseFloat(fields[32], 64)
	}
	i++
	if i <= dataLen {
		s.FBAR_SPEED_BCU, _ = strconv.ParseFloat(fields[33], 64)
	}
	i++
	if i <= dataLen {
		s.OBAR_SPEED, _ = strconv.ParseFloat(fields[34], 64)
	}
	i++
	if i <= dataLen {
		s.OBAR_SPEED_BCL, _ = strconv.ParseFloat(fields[35], 64)
	}
	i++
	if i <= dataLen {
		s.OBAR_SPEED_BCU, _ = strconv.ParseFloat(fields[36], 64)
	}
	i++
	if i <= dataLen {
		s.VDIFF_SPEED, _ = strconv.ParseFloat(fields[37], 64)
	}
	i++
	if i <= dataLen {
		s.VDIFF_SPEED_BCL, _ = strconv.ParseFloat(fields[38], 64)
	}
	i++
	if i <= dataLen {
		s.VDIFF_SPEED_BCU, _ = strconv.ParseFloat(fields[39], 64)
	}
	i++
	if i <= dataLen {
		s.VDIFF_DIR, _ = strconv.ParseFloat(fields[40], 64)
	}
	i++
	if i <= dataLen {
		s.VDIFF_DIR_BCL, _ = strconv.ParseFloat(fields[41], 64)
	}
	i++
	if i <= dataLen {
		s.VDIFF_DIR_BCU, _ = strconv.ParseFloat(fields[42], 64)
	}
	i++
	if i <= dataLen {
		s.SPEED_ERR, _ = strconv.ParseFloat(fields[43], 64)
	}
	i++
	if i <= dataLen {
		s.SPEED_ERR_BCL, _ = strconv.ParseFloat(fields[44], 64)
	}
	i++
	if i <= dataLen {
		s.SPEED_ERR_BCU, _ = strconv.ParseFloat(fields[45], 64)
	}
	i++
	if i <= dataLen {
		s.SPEED_ABSERR, _ = strconv.ParseFloat(fields[46], 64)
	}
	i++
	if i <= dataLen {
		s.SPEED_ABSERR_BCL, _ = strconv.ParseFloat(fields[47], 64)
	}
	i++
	if i <= dataLen {
		s.SPEED_ABSERR_BCU, _ = strconv.ParseFloat(fields[48], 64)
	}
	i++
	if i <= dataLen {
		s.DIR_ERR, _ = strconv.ParseFloat(fields[49], 64)
	}
	i++
	if i <= dataLen {
		s.DIR_ERR_BCL, _ = strconv.ParseFloat(fields[50], 64)
	}
	i++
	if i <= dataLen {
		s.DIR_ERR_BCU, _ = strconv.ParseFloat(fields[51], 64)
	}
	i++
	if i <= dataLen {
		s.DIR_ABSERR, _ = strconv.ParseFloat(fields[52], 64)
	}
	i++
	if i <= dataLen {
		s.DIR_ABSERR_BCL, _ = strconv.ParseFloat(fields[53], 64)
	}
	i++
	if i <= dataLen {
		s.DIR_ABSERR_BCU, _ = strconv.ParseFloat(fields[54], 64)
	}
	i++
	if i <= dataLen {
		s.ANOM_CORR, _ = strconv.ParseFloat(fields[55], 64)
	}
	i++
	if i <= dataLen {
		s.ANOM_CORR_NCL, _ = strconv.ParseFloat(fields[56], 64)
	}
	i++
	if i <= dataLen {
		s.ANOM_CORR_NCU, _ = strconv.ParseFloat(fields[57], 64)
	}
	i++
	if i <= dataLen {
		s.ANOM_CORR_BCL, _ = strconv.ParseFloat(fields[58], 64)
	}
	i++
	if i <= dataLen {
		s.ANOM_CORR_BCU, _ = strconv.ParseFloat(fields[59], 64)
	}
	i++
	if i <= dataLen {
		s.ANOM_CORR_UNCNTR, _ = strconv.ParseFloat(fields[60], 64)
	}
	i++
	if i <= dataLen {
		s.ANOM_CORR_UNCNTR_BCL, _ = strconv.ParseFloat(fields[61], 64)
	}
	i++
	if i <= dataLen {
		s.ANOM_CORR_UNCNTR_BCU, _ = strconv.ParseFloat(fields[62], 64)
	}
}

func (s *STAT_VL1L2) fill_STAT_VL1L2(fields []string) {
	dataLen := len(fields) - 1
	i := -1
	i++
	if i <= dataLen {
		s.TOTAL, _ = strconv.Atoi(fields[0])
	}
	i++
	if i <= dataLen {
		s.UFBAR, _ = strconv.ParseFloat(fields[1], 64)
	}
	i++
	if i <= dataLen {
		s.VFBAR, _ = strconv.ParseFloat(fields[2], 64)
	}
	i++
	if i <= dataLen {
		s.UOBAR, _ = strconv.ParseFloat(fields[3], 64)
	}
	i++
	if i <= dataLen {
		s.VOBAR, _ = strconv.ParseFloat(fields[4], 64)
	}
	i++
	if i <= dataLen {
		s.UVFOBAR, _ = strconv.ParseFloat(fields[5], 64)
	}
	i++
	if i <= dataLen {
		s.UVFFBAR, _ = strconv.ParseFloat(fields[6], 64)
	}
	i++
	if i <= dataLen {
		s.UVOOBAR, _ = strconv.ParseFloat(fields[7], 64)
	}
	i++
	if i <= dataLen {
		s.F_SPEED_BAR, _ = strconv.ParseFloat(fields[8], 64)
	}
	i++
	if i <= dataLen {
		s.O_SPEED_BAR, _ = strconv.ParseFloat(fields[9], 64)
	}
}

func (s *TCST_PROBRIRW) fill_TCST_PROBRIRW(fields []string) {
	dataLen := len(fields) - 1
	i := -1
	i++
	if i <= dataLen {
		s.ALAT, _ = strconv.ParseFloat(fields[0], 64)
	}
	i++
	if i <= dataLen {
		s.ALON, _ = strconv.ParseFloat(fields[1], 64)
	}
	i++
	if i <= dataLen {
		s.BLAT, _ = strconv.ParseFloat(fields[2], 64)
	}
	i++
	if i <= dataLen {
		s.BLON, _ = strconv.ParseFloat(fields[3], 64)
	}
	i++
	if i <= dataLen {
		if fields[4] != "NA" {
			s.INITIALS = fields[4]
		}
	}
	i++
	if i <= dataLen {
		s.TK_ERR, _ = strconv.ParseFloat(fields[5], 64)
	}
	i++
	if i <= dataLen {
		s.X_ERR, _ = strconv.ParseFloat(fields[6], 64)
	}
	i++
	if i <= dataLen {
		s.Y_ERR, _ = strconv.ParseFloat(fields[7], 64)
	}
	i++
	if i <= dataLen {
		s.ADLAND, _ = strconv.ParseFloat(fields[8], 64)
	}
	i++
	if i <= dataLen {
		s.BDLAND, _ = strconv.ParseFloat(fields[9], 64)
	}
	i++
	if i <= dataLen {
		s.RIRW_BEG, _ = strconv.Atoi(fields[10])
	}
	i++
	if i <= dataLen {
		s.RIRW_END, _ = strconv.Atoi(fields[11])
	}
	i++
	if i <= dataLen {
		s.RIRW_WINDOW, _ = strconv.Atoi(fields[12])
	}
	i++
	if i <= dataLen {
		s.AWIND_END, _ = strconv.ParseFloat(fields[13], 64)
	}
	i++
	if i <= dataLen {
		s.BWIND_BEG, _ = strconv.ParseFloat(fields[14], 64)
	}
	i++
	if i <= dataLen {
		s.BWIND_END, _ = strconv.ParseFloat(fields[15], 64)
	}
	i++
	if i <= dataLen {
		s.BDELTA, _ = strconv.ParseFloat(fields[16], 64)
	}
	i++
	if i <= dataLen {
		s.BDELTA_MAX, _ = strconv.ParseFloat(fields[17], 64)
	}
	i++
	if i <= dataLen {
		if fields[18] != "NA" {
			s.BLEVEL_BEG = fields[18]
		}
	}
	i++
	if i <= dataLen {
		if fields[19] != "NA" {
			s.BLEVEL_END = fields[19]
		}
	}
	i++
	if i <= dataLen { // the first field of the repeating fields is the TOTAL, the second field is the 1st dimenSion of the 1st sequence (there might be only one sequence)
		var value interface{}
		count, err := strconv.Atoi(fields[20])
		if err != nil {
			count = 0
		}
		keyPrefixes := []string{"THRESH_", "PROB_"}
		s.THRESH = make(map[string]interface{})
		for group := 1; group <= count; group++ {
			for index := 21; index <= len(keyPrefixes); index++ {
				key := fmt.Sprintf("%s_%d", keyPrefixes[index-1], index)
				if index > len(fields) { // sometimes the data line is truncated - we will set expected data to "NA"
					value = "NA"
				} else {
					value, err = strconv.Atoi(fields[index])
					if err != nil { // sometimes there can be these NA values in the data, which will be left out of json
						value = "NA"
					}
				}
				s.THRESH[key] = value
			}
		}
	}
	i++
	if i <= dataLen {
		s.INIT, _ = strconv.Atoi(fields[23])
	}
}

func (s *TCST_TCDIAG) fill_TCST_TCDIAG(fields []string) {
	dataLen := len(fields) - 1
	i := -1
	i++
	if i <= dataLen {
		s.TOTAL, _ = strconv.Atoi(fields[0])
	}
	i++
	if i <= dataLen {
		s.INDEX, _ = strconv.Atoi(fields[1])
	}
	i++
	if i <= dataLen {
		s.DIAG_SOURCE, _ = strconv.ParseFloat(fields[2], 64)
	}
	i++
	if i <= dataLen {
		if fields[3] != "NA" {
			s.TRACK_SOURCE = fields[3]
		}
	}
	i++
	if i <= dataLen {
		if fields[4] != "NA" {
			s.FIELD_SOURCE = fields[4]
		}
	}
	i++
	if i <= dataLen { // the first field of the repeating fields is the TOTAL, the second field is the 1st dimenSion of the 1st sequence (there might be only one sequence)
		var value interface{}
		count, err := strconv.Atoi(fields[5])
		if err != nil {
			count = 0
		}
		keyPrefixes := []string{"DIAG_", "VALUE_"}
		s.DIAG = make(map[string]interface{})
		for group := 1; group <= count; group++ {
			for index := 6; index <= len(keyPrefixes); index++ {
				key := fmt.Sprintf("%s_%d", keyPrefixes[index-1], index)
				if index > len(fields) { // sometimes the data line is truncated - we will set expected data to "NA"
					value = "NA"
				} else {
					value = fields[index]
				}
				s.DIAG[key] = value
			}
		}
	}
	i++
	if i <= dataLen {
		s.INIT, _ = strconv.Atoi(fields[8])
	}
}

func (s *TCST_TCMPR) fill_TCST_TCMPR(fields []string) {
	dataLen := len(fields) - 1
	i := -1
	i++
	if i <= dataLen {
		s.TOTAL, _ = strconv.Atoi(fields[0])
	}
	i++
	if i <= dataLen {
		s.INDEX, _ = strconv.Atoi(fields[1])
	}
	i++
	if i <= dataLen {
		if fields[2] != "NA" {
			s.LEVEL = fields[2]
		}
	}
	i++
	if i <= dataLen {
		if fields[3] != "NA" {
			s.WATCH_WARN = fields[3]
		}
	}
	i++
	if i <= dataLen {
		if fields[4] != "NA" {
			s.INITIALS = fields[4]
		}
	}
	i++
	if i <= dataLen {
		s.ALAT, _ = strconv.ParseFloat(fields[5], 64)
	}
	i++
	if i <= dataLen {
		s.ALON, _ = strconv.ParseFloat(fields[6], 64)
	}
	i++
	if i <= dataLen {
		s.BLAT, _ = strconv.ParseFloat(fields[7], 64)
	}
	i++
	if i <= dataLen {
		s.BLON, _ = strconv.ParseFloat(fields[8], 64)
	}
	i++
	if i <= dataLen {
		s.TK_ERR, _ = strconv.ParseFloat(fields[9], 64)
	}
	i++
	if i <= dataLen {
		s.X_ERR, _ = strconv.ParseFloat(fields[10], 64)
	}
	i++
	if i <= dataLen {
		s.Y_ERR, _ = strconv.ParseFloat(fields[11], 64)
	}
	i++
	if i <= dataLen {
		s.ALTK_ERR, _ = strconv.ParseFloat(fields[12], 64)
	}
	i++
	if i <= dataLen {
		s.CRTK_ERR, _ = strconv.ParseFloat(fields[13], 64)
	}
	i++
	if i <= dataLen {
		s.ADLAND, _ = strconv.ParseFloat(fields[14], 64)
	}
	i++
	if i <= dataLen {
		s.BDLAND, _ = strconv.ParseFloat(fields[15], 64)
	}
	i++
	if i <= dataLen {
		s.AMSLP, _ = strconv.ParseFloat(fields[16], 64)
	}
	i++
	if i <= dataLen {
		s.BMSLP, _ = strconv.ParseFloat(fields[17], 64)
	}
	i++
	if i <= dataLen {
		s.AMAX_WIND, _ = strconv.ParseFloat(fields[18], 64)
	}
	i++
	if i <= dataLen {
		s.BMAX_WIND, _ = strconv.ParseFloat(fields[19], 64)
	}
	i++
	if i <= dataLen {
		s.AAL_WIND_34, _ = strconv.ParseFloat(fields[20], 64)
	}
	i++
	if i <= dataLen {
		s.BAL_WIND_34, _ = strconv.ParseFloat(fields[21], 64)
	}
	i++
	if i <= dataLen {
		s.ANE_WIND_34, _ = strconv.ParseFloat(fields[22], 64)
	}
	i++
	if i <= dataLen {
		s.BNE_WIND_34, _ = strconv.ParseFloat(fields[23], 64)
	}
	i++
	if i <= dataLen {
		s.ASE_WIND_34, _ = strconv.ParseFloat(fields[24], 64)
	}
	i++
	if i <= dataLen {
		s.BSE_WIND_34, _ = strconv.ParseFloat(fields[25], 64)
	}
	i++
	if i <= dataLen {
		s.ASW_WIND_34, _ = strconv.ParseFloat(fields[26], 64)
	}
	i++
	if i <= dataLen {
		s.BSW_WIND_34, _ = strconv.ParseFloat(fields[27], 64)
	}
	i++
	if i <= dataLen {
		s.ANW_WIND_34, _ = strconv.ParseFloat(fields[28], 64)
	}
	i++
	if i <= dataLen {
		s.BNW_WIND_34, _ = strconv.ParseFloat(fields[29], 64)
	}
	i++
	if i <= dataLen {
		s.AAL_WIND_50, _ = strconv.ParseFloat(fields[30], 64)
	}
	i++
	if i <= dataLen {
		s.BAL_WIND_50, _ = strconv.ParseFloat(fields[31], 64)
	}
	i++
	if i <= dataLen {
		s.ANE_WIND_50, _ = strconv.ParseFloat(fields[32], 64)
	}
	i++
	if i <= dataLen {
		s.BNE_WIND_50, _ = strconv.ParseFloat(fields[33], 64)
	}
	i++
	if i <= dataLen {
		s.ASE_WIND_50, _ = strconv.ParseFloat(fields[34], 64)
	}
	i++
	if i <= dataLen {
		s.BSE_WIND_50, _ = strconv.ParseFloat(fields[35], 64)
	}
	i++
	if i <= dataLen {
		s.ASW_WIND_50, _ = strconv.ParseFloat(fields[36], 64)
	}
	i++
	if i <= dataLen {
		s.BSW_WIND_50, _ = strconv.ParseFloat(fields[37], 64)
	}
	i++
	if i <= dataLen {
		s.ANW_WIND_50, _ = strconv.ParseFloat(fields[38], 64)
	}
	i++
	if i <= dataLen {
		s.BNW_WIND_50, _ = strconv.ParseFloat(fields[39], 64)
	}
	i++
	if i <= dataLen {
		s.AAL_WIND_64, _ = strconv.ParseFloat(fields[40], 64)
	}
	i++
	if i <= dataLen {
		s.BAL_WIND_64, _ = strconv.ParseFloat(fields[41], 64)
	}
	i++
	if i <= dataLen {
		s.ANE_WIND_64, _ = strconv.ParseFloat(fields[42], 64)
	}
	i++
	if i <= dataLen {
		s.BNE_WIND_64, _ = strconv.ParseFloat(fields[43], 64)
	}
	i++
	if i <= dataLen {
		s.ASE_WIND_64, _ = strconv.ParseFloat(fields[44], 64)
	}
	i++
	if i <= dataLen {
		s.BSE_WIND_64, _ = strconv.ParseFloat(fields[45], 64)
	}
	i++
	if i <= dataLen {
		s.ASW_WIND_64, _ = strconv.ParseFloat(fields[46], 64)
	}
	i++
	if i <= dataLen {
		s.BSW_WIND_64, _ = strconv.ParseFloat(fields[47], 64)
	}
	i++
	if i <= dataLen {
		s.ANW_WIND_64, _ = strconv.ParseFloat(fields[48], 64)
	}
	i++
	if i <= dataLen {
		s.BNW_WIND_64, _ = strconv.ParseFloat(fields[49], 64)
	}
	i++
	if i <= dataLen {
		if fields[50] != "NA" {
			s.ARADP = fields[50]
		}
	}
	i++
	if i <= dataLen {
		s.BRADP, _ = strconv.ParseFloat(fields[51], 64)
	}
	i++
	if i <= dataLen {
		s.ARRP, _ = strconv.Atoi(fields[52])
	}
	i++
	if i <= dataLen {
		s.BRRP, _ = strconv.ParseFloat(fields[53], 64)
	}
	i++
	if i <= dataLen {
		s.AMRD, _ = strconv.Atoi(fields[54])
	}
	i++
	if i <= dataLen {
		s.BMRD, _ = strconv.ParseFloat(fields[55], 64)
	}
	i++
	if i <= dataLen {
		s.AGUSTS, _ = strconv.Atoi(fields[56])
	}
	i++
	if i <= dataLen {
		s.BGUSTS, _ = strconv.ParseFloat(fields[57], 64)
	}
	i++
	if i <= dataLen {
		s.AEYE, _ = strconv.Atoi(fields[58])
	}
	i++
	if i <= dataLen {
		s.BEYE, _ = strconv.ParseFloat(fields[59], 64)
	}
	i++
	if i <= dataLen {
		s.ADIR, _ = strconv.Atoi(fields[60])
	}
	i++
	if i <= dataLen {
		s.BDIR, _ = strconv.ParseFloat(fields[61], 64)
	}
	i++
	if i <= dataLen {
		s.ASPEED, _ = strconv.Atoi(fields[62])
	}
	i++
	if i <= dataLen {
		s.BSPEED, _ = strconv.ParseFloat(fields[63], 64)
	}
	i++
	if i <= dataLen {
		s.ADEPTH, _ = strconv.Atoi(fields[64])
	}
	i++
	if i <= dataLen {
		s.BDEPTH, _ = strconv.ParseFloat(fields[65], 64)
	}
	i++
	if i <= dataLen {
		s.NUM_MEMBERS, _ = strconv.ParseFloat(fields[66], 64)
	}
	i++
	if i <= dataLen {
		s.TRACK_SPREAD, _ = strconv.ParseFloat(fields[67], 64)
	}
	i++
	if i <= dataLen {
		s.TRACK_STDEV, _ = strconv.ParseFloat(fields[68], 64)
	}
	i++
	if i <= dataLen {
		s.MSLP_STDEV, _ = strconv.ParseFloat(fields[69], 64)
	}
	i++
	if i <= dataLen {
		s.MAX_WIND_STDEV, _ = strconv.ParseFloat(fields[70], 64)
	}
	i++
	if i <= dataLen {
		s.INIT, _ = strconv.Atoi(fields[71])
	}
}

// getDocForId functions
func GetDocForId(fileLineType string, metaDataMap map[string]interface{}, headerData []string, dataData []string, dataKey string) (map[string]interface{}, error) {
	doc := make(map[string]interface{})
	// add the metadata to the doc
	for key, value := range metaDataMap {
		doc[key] = value
	}
	switch fileLineType {
	case "STAT_CNT":
		elem := STAT_CNT{}
		elem.fill_STAT_CNT_Header(headerData, &doc)
		elem.fill_STAT_CNT(dataData)
		if exists := (doc)["data"]; exists == nil {
			(doc)["data"] = make(map[string]STAT_CNT)
		}
		if val, ok := (doc)["data"].(map[string]STAT_CNT); ok {
			val[dataKey] = elem
			(doc)["data"] = val
		}
	case "STAT_CTC":
		elem := STAT_CTC{}
		elem.fill_STAT_CTC_Header(headerData, &doc)
		elem.fill_STAT_CTC(dataData)
		if exists := (doc)["data"]; exists == nil {
			(doc)["data"] = make(map[string]STAT_CTC)
		}
		if val, ok := (doc)["data"].(map[string]STAT_CTC); ok {
			val[dataKey] = elem
			(doc)["data"] = val
		}
	case "STAT_CTS":
		elem := STAT_CTS{}
		elem.fill_STAT_CTS_Header(headerData, &doc)
		elem.fill_STAT_CTS(dataData)
		if exists := (doc)["data"]; exists == nil {
			(doc)["data"] = make(map[string]STAT_CTS)
		}
		if val, ok := (doc)["data"].(map[string]STAT_CTS); ok {
			val[dataKey] = elem
			(doc)["data"] = val
		}
	case "STAT_FHO":
		elem := STAT_FHO{}
		elem.fill_STAT_FHO_Header(headerData, &doc)
		elem.fill_STAT_FHO(dataData)
		if exists := (doc)["data"]; exists == nil {
			(doc)["data"] = make(map[string]STAT_FHO)
		}
		if val, ok := (doc)["data"].(map[string]STAT_FHO); ok {
			val[dataKey] = elem
			(doc)["data"] = val
		}
	case "STAT_ISC":
		elem := STAT_ISC{}
		elem.fill_STAT_ISC_Header(headerData, &doc)
		elem.fill_STAT_ISC(dataData)
		if exists := (doc)["data"]; exists == nil {
			(doc)["data"] = make(map[string]STAT_ISC)
		}
		if val, ok := (doc)["data"].(map[string]STAT_ISC); ok {
			val[dataKey] = elem
			(doc)["data"] = val
		}
	case "STAT_MCTC":
		elem := STAT_MCTC{}
		elem.fill_STAT_MCTC_Header(headerData, &doc)
		elem.fill_STAT_MCTC(dataData)
		if exists := (doc)["data"]; exists == nil {
			(doc)["data"] = make(map[string]STAT_MCTC)
		}
		if val, ok := (doc)["data"].(map[string]STAT_MCTC); ok {
			val[dataKey] = elem
			(doc)["data"] = val
		}
	case "STAT_MCTS":
		elem := STAT_MCTS{}
		elem.fill_STAT_MCTS_Header(headerData, &doc)
		elem.fill_STAT_MCTS(dataData)
		if exists := (doc)["data"]; exists == nil {
			(doc)["data"] = make(map[string]STAT_MCTS)
		}
		if val, ok := (doc)["data"].(map[string]STAT_MCTS); ok {
			val[dataKey] = elem
			(doc)["data"] = val
		}
	case "STAT_MPR":
		elem := STAT_MPR{}
		elem.fill_STAT_MPR_Header(headerData, &doc)
		elem.fill_STAT_MPR(dataData)
		if exists := (doc)["data"]; exists == nil {
			(doc)["data"] = make(map[string]STAT_MPR)
		}
		if val, ok := (doc)["data"].(map[string]STAT_MPR); ok {
			val[dataKey] = elem
			(doc)["data"] = val
		}
	case "STAT_SEEPS":
		elem := STAT_SEEPS{}
		elem.fill_STAT_SEEPS_Header(headerData, &doc)
		elem.fill_STAT_SEEPS(dataData)
		if exists := (doc)["data"]; exists == nil {
			(doc)["data"] = make(map[string]STAT_SEEPS)
		}
		if val, ok := (doc)["data"].(map[string]STAT_SEEPS); ok {
			val[dataKey] = elem
			(doc)["data"] = val
		}
	case "STAT_SEEPS_MPR":
		elem := STAT_SEEPS_MPR{}
		elem.fill_STAT_SEEPS_MPR_Header(headerData, &doc)
		elem.fill_STAT_SEEPS_MPR(dataData)
		if exists := (doc)["data"]; exists == nil {
			(doc)["data"] = make(map[string]STAT_SEEPS_MPR)
		}
		if val, ok := (doc)["data"].(map[string]STAT_SEEPS_MPR); ok {
			val[dataKey] = elem
			(doc)["data"] = val
		}
	case "STAT_NBRCNT":
		elem := STAT_NBRCNT{}
		elem.fill_STAT_NBRCNT_Header(headerData, &doc)
		elem.fill_STAT_NBRCNT(dataData)
		if exists := (doc)["data"]; exists == nil {
			(doc)["data"] = make(map[string]STAT_NBRCNT)
		}
		if val, ok := (doc)["data"].(map[string]STAT_NBRCNT); ok {
			val[dataKey] = elem
			(doc)["data"] = val
		}
	case "STAT_NBRCTC":
		elem := STAT_NBRCTC{}
		elem.fill_STAT_NBRCTC_Header(headerData, &doc)
		elem.fill_STAT_NBRCTC(dataData)
		if exists := (doc)["data"]; exists == nil {
			(doc)["data"] = make(map[string]STAT_NBRCTC)
		}
		if val, ok := (doc)["data"].(map[string]STAT_NBRCTC); ok {
			val[dataKey] = elem
			(doc)["data"] = val
		}
	case "STAT_NBRCTS":
		elem := STAT_NBRCTS{}
		elem.fill_STAT_NBRCTS_Header(headerData, &doc)
		elem.fill_STAT_NBRCTS(dataData)
		if exists := (doc)["data"]; exists == nil {
			(doc)["data"] = make(map[string]STAT_NBRCTS)
		}
		if val, ok := (doc)["data"].(map[string]STAT_NBRCTS); ok {
			val[dataKey] = elem
			(doc)["data"] = val
		}
	case "STAT_GRAD":
		elem := STAT_GRAD{}
		elem.fill_STAT_GRAD_Header(headerData, &doc)
		elem.fill_STAT_GRAD(dataData)
		if exists := (doc)["data"]; exists == nil {
			(doc)["data"] = make(map[string]STAT_GRAD)
		}
		if val, ok := (doc)["data"].(map[string]STAT_GRAD); ok {
			val[dataKey] = elem
			(doc)["data"] = val
		}
	case "STAT_DMAP":
		elem := STAT_DMAP{}
		elem.fill_STAT_DMAP_Header(headerData, &doc)
		elem.fill_STAT_DMAP(dataData)
		if exists := (doc)["data"]; exists == nil {
			(doc)["data"] = make(map[string]STAT_DMAP)
		}
		if val, ok := (doc)["data"].(map[string]STAT_DMAP); ok {
			val[dataKey] = elem
			(doc)["data"] = val
		}
	case "STAT_ORANK":
		elem := STAT_ORANK{}
		elem.fill_STAT_ORANK_Header(headerData, &doc)
		elem.fill_STAT_ORANK(dataData)
		if exists := (doc)["data"]; exists == nil {
			(doc)["data"] = make(map[string]STAT_ORANK)
		}
		if val, ok := (doc)["data"].(map[string]STAT_ORANK); ok {
			val[dataKey] = elem
			(doc)["data"] = val
		}
	case "STAT_PCT":
		elem := STAT_PCT{}
		elem.fill_STAT_PCT_Header(headerData, &doc)
		elem.fill_STAT_PCT(dataData)
		if exists := (doc)["data"]; exists == nil {
			(doc)["data"] = make(map[string]STAT_PCT)
		}
		if val, ok := (doc)["data"].(map[string]STAT_PCT); ok {
			val[dataKey] = elem
			(doc)["data"] = val
		}
	case "STAT_PJC":
		elem := STAT_PJC{}
		elem.fill_STAT_PJC_Header(headerData, &doc)
		elem.fill_STAT_PJC(dataData)
		if exists := (doc)["data"]; exists == nil {
			(doc)["data"] = make(map[string]STAT_PJC)
		}
		if val, ok := (doc)["data"].(map[string]STAT_PJC); ok {
			val[dataKey] = elem
			(doc)["data"] = val
		}
	case "STAT_PRC":
		elem := STAT_PRC{}
		elem.fill_STAT_PRC_Header(headerData, &doc)
		elem.fill_STAT_PRC(dataData)
		if exists := (doc)["data"]; exists == nil {
			(doc)["data"] = make(map[string]STAT_PRC)
		}
		if val, ok := (doc)["data"].(map[string]STAT_PRC); ok {
			val[dataKey] = elem
			(doc)["data"] = val
		}
	case "STAT_PSTD":
		elem := STAT_PSTD{}
		elem.fill_STAT_PSTD_Header(headerData, &doc)
		elem.fill_STAT_PSTD(dataData)
		if exists := (doc)["data"]; exists == nil {
			(doc)["data"] = make(map[string]STAT_PSTD)
		}
		if val, ok := (doc)["data"].(map[string]STAT_PSTD); ok {
			val[dataKey] = elem
			(doc)["data"] = val
		}
	case "STAT_ECLV":
		elem := STAT_ECLV{}
		elem.fill_STAT_ECLV_Header(headerData, &doc)
		elem.fill_STAT_ECLV(dataData)
		if exists := (doc)["data"]; exists == nil {
			(doc)["data"] = make(map[string]STAT_ECLV)
		}
		if val, ok := (doc)["data"].(map[string]STAT_ECLV); ok {
			val[dataKey] = elem
			(doc)["data"] = val
		}
	case "STAT_ECNT":
		elem := STAT_ECNT{}
		elem.fill_STAT_ECNT_Header(headerData, &doc)
		elem.fill_STAT_ECNT(dataData)
		if exists := (doc)["data"]; exists == nil {
			(doc)["data"] = make(map[string]STAT_ECNT)
		}
		if val, ok := (doc)["data"].(map[string]STAT_ECNT); ok {
			val[dataKey] = elem
			(doc)["data"] = val
		}
	case "STAT_RPS":
		elem := STAT_RPS{}
		elem.fill_STAT_RPS_Header(headerData, &doc)
		elem.fill_STAT_RPS(dataData)
		if exists := (doc)["data"]; exists == nil {
			(doc)["data"] = make(map[string]STAT_RPS)
		}
		if val, ok := (doc)["data"].(map[string]STAT_RPS); ok {
			val[dataKey] = elem
			(doc)["data"] = val
		}
	case "STAT_RHIST":
		elem := STAT_RHIST{}
		elem.fill_STAT_RHIST_Header(headerData, &doc)
		elem.fill_STAT_RHIST(dataData)
		if exists := (doc)["data"]; exists == nil {
			(doc)["data"] = make(map[string]STAT_RHIST)
		}
		if val, ok := (doc)["data"].(map[string]STAT_RHIST); ok {
			val[dataKey] = elem
			(doc)["data"] = val
		}
	case "STAT_PHIST":
		elem := STAT_PHIST{}
		elem.fill_STAT_PHIST_Header(headerData, &doc)
		elem.fill_STAT_PHIST(dataData)
		if exists := (doc)["data"]; exists == nil {
			(doc)["data"] = make(map[string]STAT_PHIST)
		}
		if val, ok := (doc)["data"].(map[string]STAT_PHIST); ok {
			val[dataKey] = elem
			(doc)["data"] = val
		}
	case "STAT_RELP":
		elem := STAT_RELP{}
		elem.fill_STAT_RELP_Header(headerData, &doc)
		elem.fill_STAT_RELP(dataData)
		if exists := (doc)["data"]; exists == nil {
			(doc)["data"] = make(map[string]STAT_RELP)
		}
		if val, ok := (doc)["data"].(map[string]STAT_RELP); ok {
			val[dataKey] = elem
			(doc)["data"] = val
		}
	case "STAT_SAL1L2":
		elem := STAT_SAL1L2{}
		elem.fill_STAT_SAL1L2_Header(headerData, &doc)
		elem.fill_STAT_SAL1L2(dataData)
		if exists := (doc)["data"]; exists == nil {
			(doc)["data"] = make(map[string]STAT_SAL1L2)
		}
		if val, ok := (doc)["data"].(map[string]STAT_SAL1L2); ok {
			val[dataKey] = elem
			(doc)["data"] = val
		}
	case "STAT_SL1L2":
		elem := STAT_SL1L2{}
		elem.fill_STAT_SL1L2_Header(headerData, &doc)
		elem.fill_STAT_SL1L2(dataData)
		if exists := (doc)["data"]; exists == nil {
			(doc)["data"] = make(map[string]STAT_SL1L2)
		}
		if val, ok := (doc)["data"].(map[string]STAT_SL1L2); ok {
			val[dataKey] = elem
			(doc)["data"] = val
		}
	case "STAT_SSVAR":
		elem := STAT_SSVAR{}
		elem.fill_STAT_SSVAR_Header(headerData, &doc)
		elem.fill_STAT_SSVAR(dataData)
		if exists := (doc)["data"]; exists == nil {
			(doc)["data"] = make(map[string]STAT_SSVAR)
		}
		if val, ok := (doc)["data"].(map[string]STAT_SSVAR); ok {
			val[dataKey] = elem
			(doc)["data"] = val
		}
	case "STAT_VAL1L2":
		elem := STAT_VAL1L2{}
		elem.fill_STAT_VAL1L2_Header(headerData, &doc)
		elem.fill_STAT_VAL1L2(dataData)
		if exists := (doc)["data"]; exists == nil {
			(doc)["data"] = make(map[string]STAT_VAL1L2)
		}
		if val, ok := (doc)["data"].(map[string]STAT_VAL1L2); ok {
			val[dataKey] = elem
			(doc)["data"] = val
		}
	case "STAT_VL1L2":
		elem := STAT_VL1L2{}
		elem.fill_STAT_VL1L2_Header(headerData, &doc)
		elem.fill_STAT_VL1L2(dataData)
		if exists := (doc)["data"]; exists == nil {
			(doc)["data"] = make(map[string]STAT_VL1L2)
		}
		if val, ok := (doc)["data"].(map[string]STAT_VL1L2); ok {
			val[dataKey] = elem
			(doc)["data"] = val
		}
	case "STAT_VCNT":
		elem := STAT_VCNT{}
		elem.fill_STAT_VCNT_Header(headerData, &doc)
		elem.fill_STAT_VCNT(dataData)
		if exists := (doc)["data"]; exists == nil {
			(doc)["data"] = make(map[string]STAT_VCNT)
		}
		if val, ok := (doc)["data"].(map[string]STAT_VCNT); ok {
			val[dataKey] = elem
			(doc)["data"] = val
		}
	case "STAT_GENMPR":
		elem := STAT_GENMPR{}
		elem.fill_STAT_GENMPR_Header(headerData, &doc)
		elem.fill_STAT_GENMPR(dataData)
		if exists := (doc)["data"]; exists == nil {
			(doc)["data"] = make(map[string]STAT_GENMPR)
		}
		if val, ok := (doc)["data"].(map[string]STAT_GENMPR); ok {
			val[dataKey] = elem
			(doc)["data"] = val
		}
	case "STAT_SSIDX":
		elem := STAT_SSIDX{}
		elem.fill_STAT_SSIDX_Header(headerData, &doc)
		elem.fill_STAT_SSIDX(dataData)
		if exists := (doc)["data"]; exists == nil {
			(doc)["data"] = make(map[string]STAT_SSIDX)
		}
		if val, ok := (doc)["data"].(map[string]STAT_SSIDX); ok {
			val[dataKey] = elem
			(doc)["data"] = val
		}
	case "MODE_OBJ":
		elem := MODE_OBJ{}
		elem.fill_MODE_OBJ_Header(headerData, &doc)
		elem.fill_MODE_OBJ(dataData)
		if exists := (doc)["data"]; exists == nil {
			(doc)["data"] = make(map[string]MODE_OBJ)
		}
		if val, ok := (doc)["data"].(map[string]MODE_OBJ); ok {
			val[dataKey] = elem
			(doc)["data"] = val
		}
	case "MODE_CTS":
		elem := MODE_CTS{}
		elem.fill_MODE_CTS_Header(headerData, &doc)
		elem.fill_MODE_CTS(dataData)
		if exists := (doc)["data"]; exists == nil {
			(doc)["data"] = make(map[string]MODE_CTS)
		}
		if val, ok := (doc)["data"].(map[string]MODE_CTS); ok {
			val[dataKey] = elem
			(doc)["data"] = val
		}
	case "TCST_TCMPR":
		elem := TCST_TCMPR{}
		elem.fill_TCST_TCMPR_Header(headerData, &doc)
		elem.fill_TCST_TCMPR(dataData)
		if exists := (doc)["data"]; exists == nil {
			(doc)["data"] = make(map[string]TCST_TCMPR)
		}
		if val, ok := (doc)["data"].(map[string]TCST_TCMPR); ok {
			val[dataKey] = elem
			(doc)["data"] = val
		}
	case "TCST_TCDIAG":
		elem := TCST_TCDIAG{}
		elem.fill_TCST_TCDIAG_Header(headerData, &doc)
		elem.fill_TCST_TCDIAG(dataData)
		if exists := (doc)["data"]; exists == nil {
			(doc)["data"] = make(map[string]TCST_TCDIAG)
		}
		if val, ok := (doc)["data"].(map[string]TCST_TCDIAG); ok {
			val[dataKey] = elem
			(doc)["data"] = val
		}
	case "TCST_PROBRIRW":
		elem := TCST_PROBRIRW{}
		elem.fill_TCST_PROBRIRW_Header(headerData, &doc)
		elem.fill_TCST_PROBRIRW(dataData)
		if exists := (doc)["data"]; exists == nil {
			(doc)["data"] = make(map[string]TCST_PROBRIRW)
		}
		if val, ok := (doc)["data"].(map[string]TCST_PROBRIRW); ok {
			val[dataKey] = elem
			(doc)["data"] = val
		}
	default:
		return nil, errors.New("GetDocForId: Unknown file_line type:" + fileLineType)
	}
	return doc, nil
}

// addDataElement functions
func AddDataElement(dataKey string, fileLineType string, dataData []string, doc *map[string]interface{}) (map[string]interface{}, error) {
	switch fileLineType {
	case "STAT_CNT":
		elem := STAT_CNT{}
		elem.fill_STAT_CNT(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_CNT); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
	case "STAT_CTC":
		elem := STAT_CTC{}
		elem.fill_STAT_CTC(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_CTC); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
	case "STAT_CTS":
		elem := STAT_CTS{}
		elem.fill_STAT_CTS(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_CTS); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
	case "STAT_FHO":
		elem := STAT_FHO{}
		elem.fill_STAT_FHO(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_FHO); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
	case "STAT_ISC":
		elem := STAT_ISC{}
		elem.fill_STAT_ISC(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_ISC); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
	case "STAT_MCTC":
		elem := STAT_MCTC{}
		elem.fill_STAT_MCTC(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_MCTC); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
	case "STAT_MCTS":
		elem := STAT_MCTS{}
		elem.fill_STAT_MCTS(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_MCTS); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
	case "STAT_MPR":
		elem := STAT_MPR{}
		elem.fill_STAT_MPR(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_MPR); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
	case "STAT_SEEPS":
		elem := STAT_SEEPS{}
		elem.fill_STAT_SEEPS(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_SEEPS); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
	case "STAT_SEEPS_MPR":
		elem := STAT_SEEPS_MPR{}
		elem.fill_STAT_SEEPS_MPR(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_SEEPS_MPR); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
	case "STAT_NBRCNT":
		elem := STAT_NBRCNT{}
		elem.fill_STAT_NBRCNT(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_NBRCNT); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
	case "STAT_NBRCTC":
		elem := STAT_NBRCTC{}
		elem.fill_STAT_NBRCTC(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_NBRCTC); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
	case "STAT_NBRCTS":
		elem := STAT_NBRCTS{}
		elem.fill_STAT_NBRCTS(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_NBRCTS); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
	case "STAT_GRAD":
		elem := STAT_GRAD{}
		elem.fill_STAT_GRAD(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_GRAD); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
	case "STAT_DMAP":
		elem := STAT_DMAP{}
		elem.fill_STAT_DMAP(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_DMAP); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
	case "STAT_ORANK":
		elem := STAT_ORANK{}
		elem.fill_STAT_ORANK(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_ORANK); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
	case "STAT_PCT":
		elem := STAT_PCT{}
		elem.fill_STAT_PCT(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_PCT); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
	case "STAT_PJC":
		elem := STAT_PJC{}
		elem.fill_STAT_PJC(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_PJC); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
	case "STAT_PRC":
		elem := STAT_PRC{}
		elem.fill_STAT_PRC(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_PRC); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
	case "STAT_PSTD":
		elem := STAT_PSTD{}
		elem.fill_STAT_PSTD(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_PSTD); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
	case "STAT_ECLV":
		elem := STAT_ECLV{}
		elem.fill_STAT_ECLV(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_ECLV); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
	case "STAT_ECNT":
		elem := STAT_ECNT{}
		elem.fill_STAT_ECNT(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_ECNT); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
	case "STAT_RPS":
		elem := STAT_RPS{}
		elem.fill_STAT_RPS(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_RPS); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
	case "STAT_RHIST":
		elem := STAT_RHIST{}
		elem.fill_STAT_RHIST(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_RHIST); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
	case "STAT_PHIST":
		elem := STAT_PHIST{}
		elem.fill_STAT_PHIST(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_PHIST); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
	case "STAT_RELP":
		elem := STAT_RELP{}
		elem.fill_STAT_RELP(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_RELP); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
	case "STAT_SAL1L2":
		elem := STAT_SAL1L2{}
		elem.fill_STAT_SAL1L2(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_SAL1L2); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
	case "STAT_SL1L2":
		elem := STAT_SL1L2{}
		elem.fill_STAT_SL1L2(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_SL1L2); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
	case "STAT_SSVAR":
		elem := STAT_SSVAR{}
		elem.fill_STAT_SSVAR(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_SSVAR); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
	case "STAT_VAL1L2":
		elem := STAT_VAL1L2{}
		elem.fill_STAT_VAL1L2(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_VAL1L2); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
	case "STAT_VL1L2":
		elem := STAT_VL1L2{}
		elem.fill_STAT_VL1L2(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_VL1L2); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
	case "STAT_VCNT":
		elem := STAT_VCNT{}
		elem.fill_STAT_VCNT(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_VCNT); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
	case "STAT_GENMPR":
		elem := STAT_GENMPR{}
		elem.fill_STAT_GENMPR(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_GENMPR); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
	case "STAT_SSIDX":
		elem := STAT_SSIDX{}
		elem.fill_STAT_SSIDX(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_SSIDX); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
	case "MODE_OBJ":
		elem := MODE_OBJ{}
		elem.fill_MODE_OBJ(dataData)
		if val, ok := (*doc)["data"].(map[string]MODE_OBJ); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
	case "MODE_CTS":
		elem := MODE_CTS{}
		elem.fill_MODE_CTS(dataData)
		if val, ok := (*doc)["data"].(map[string]MODE_CTS); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
	case "TCST_TCMPR":
		elem := TCST_TCMPR{}
		elem.fill_TCST_TCMPR(dataData)
		if val, ok := (*doc)["data"].(map[string]TCST_TCMPR); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
	case "TCST_TCDIAG":
		elem := TCST_TCDIAG{}
		elem.fill_TCST_TCDIAG(dataData)
		if val, ok := (*doc)["data"].(map[string]TCST_TCDIAG); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
	case "TCST_PROBRIRW":
		elem := TCST_PROBRIRW{}
		elem.fill_TCST_PROBRIRW(dataData)
		if val, ok := (*doc)["data"].(map[string]TCST_PROBRIRW); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
	default:
		return nil, errors.New("AddDataElement: Unknown file_line type:" + fileLineType)
	}
	return *doc, nil
}

var MetHeaderColumnsFileUrl = "https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/data/table_files/met_header_columns_V11.0.txt"
