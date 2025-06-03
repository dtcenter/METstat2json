package v10_1

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
	VERSION    string  `json:"version"`
	MODEL      string  `json:"model"`
	N_VALID    int     `json:"nValid"`
	GRID_RES   float64 `json:"gridRes"`
	DESC       string  `json:"desc"`
	FCST_VALID string  `json:"fcstValid"`
	FCST_ACCUM string  `json:"fcstAccum"`
	OBS_LEAD   int     `json:"obsLead"`
	OBS_VALID  string  `json:"obsValid"`
	OBS_ACCUM  string  `json:"obsAccum"`
	FCST_RAD   int     `json:"fcstRad"`
	FCST_THR   string  `json:"fcstThr"`
	OBS_RAD    int     `json:"obsRad"`
	OBS_THR    string  `json:"obsThr"`
	FCST_VAR   string  `json:"fcstVar"`
	FCST_UNITS string  `json:"fcstUnits"`
	FCST_LEV   string  `json:"fcstLev"`
	OBS_VAR    string  `json:"obsVar"`
	OBS_UNITS  string  `json:"obsUnits"`
	OBS_LEV    string  `json:"obsLev"`
	OBTYPE     string  `json:"obtype"`
	LINE_TYPE  string  `json:"lineType"`
}

type MODE_OBJ_header struct {
	VERSION    string  `json:"version"`
	MODEL      string  `json:"model"`
	N_VALID    int     `json:"nValid"`
	GRID_RES   float64 `json:"gridRes"`
	DESC       string  `json:"desc"`
	FCST_VALID string  `json:"fcstValid"`
	FCST_ACCUM string  `json:"fcstAccum"`
	OBS_LEAD   int     `json:"obsLead"`
	OBS_VALID  string  `json:"obsValid"`
	OBS_ACCUM  string  `json:"obsAccum"`
	FCST_RAD   int     `json:"fcstRad"`
	FCST_THR   string  `json:"fcstThr"`
	OBS_RAD    int     `json:"obsRad"`
	OBS_THR    string  `json:"obsThr"`
	FCST_VAR   string  `json:"fcstVar"`
	FCST_UNITS string  `json:"fcstUnits"`
	FCST_LEV   string  `json:"fcstLev"`
	OBS_VAR    string  `json:"obsVar"`
	OBS_UNITS  string  `json:"obsUnits"`
	OBS_LEV    string  `json:"obsLev"`
	OBTYPE     string  `json:"obtype"`
	LINE_TYPE  string  `json:"lineType"`
}

type STAT_CNT_header struct {
	VERSION        string  `json:"version"`
	MODEL          string  `json:"model"`
	DESC           string  `json:"desc"`
	FCST_VALID_BEG int     `json:"fcstValidBeg"`
	FCST_VALID_END int     `json:"fcstValidEnd"`
	OBS_LEAD       int     `json:"obsLead"`
	OBS_VALID_BEG  int     `json:"obsValidBeg"`
	OBS_VALID_END  int     `json:"obsValidEnd"`
	FCST_VAR       string  `json:"fcstVar"`
	FCST_UNITS     string  `json:"fcstUnits"`
	FCST_LEV       string  `json:"fcstLev"`
	OBS_VAR        string  `json:"obsVar"`
	OBS_UNITS      string  `json:"obsUnits"`
	OBS_LEV        string  `json:"obsLev"`
	OBTYPE         string  `json:"obtype"`
	VX_MASK        string  `json:"vxMask"`
	INTERP_MTHD    string  `json:"interpMthd"`
	INTERP_PNTS    int     `json:"interpPnts"`
	FCST_THRESH    string  `json:"fcstThresh"`
	OBS_THRESH     string  `json:"obsThresh"`
	COV_THRESH     string  `json:"covThresh"`
	ALPHA          float64 `json:"alpha"`
	LINE_TYPE      string  `json:"lineType"`
}

type STAT_CTC_header struct {
	VERSION        string  `json:"version"`
	MODEL          string  `json:"model"`
	DESC           string  `json:"desc"`
	FCST_VALID_BEG int     `json:"fcstValidBeg"`
	FCST_VALID_END int     `json:"fcstValidEnd"`
	OBS_LEAD       int     `json:"obsLead"`
	OBS_VALID_BEG  int     `json:"obsValidBeg"`
	OBS_VALID_END  int     `json:"obsValidEnd"`
	FCST_VAR       string  `json:"fcstVar"`
	FCST_UNITS     string  `json:"fcstUnits"`
	FCST_LEV       string  `json:"fcstLev"`
	OBS_VAR        string  `json:"obsVar"`
	OBS_UNITS      string  `json:"obsUnits"`
	OBS_LEV        string  `json:"obsLev"`
	OBTYPE         string  `json:"obtype"`
	VX_MASK        string  `json:"vxMask"`
	INTERP_MTHD    string  `json:"interpMthd"`
	INTERP_PNTS    int     `json:"interpPnts"`
	FCST_THRESH    string  `json:"fcstThresh"`
	OBS_THRESH     string  `json:"obsThresh"`
	COV_THRESH     string  `json:"covThresh"`
	ALPHA          float64 `json:"alpha"`
	LINE_TYPE      string  `json:"lineType"`
}

type STAT_CTS_header struct {
	VERSION        string  `json:"version"`
	MODEL          string  `json:"model"`
	DESC           string  `json:"desc"`
	FCST_VALID_BEG int     `json:"fcstValidBeg"`
	FCST_VALID_END int     `json:"fcstValidEnd"`
	OBS_LEAD       int     `json:"obsLead"`
	OBS_VALID_BEG  int     `json:"obsValidBeg"`
	OBS_VALID_END  int     `json:"obsValidEnd"`
	FCST_VAR       string  `json:"fcstVar"`
	FCST_UNITS     string  `json:"fcstUnits"`
	FCST_LEV       string  `json:"fcstLev"`
	OBS_VAR        string  `json:"obsVar"`
	OBS_UNITS      string  `json:"obsUnits"`
	OBS_LEV        string  `json:"obsLev"`
	OBTYPE         string  `json:"obtype"`
	VX_MASK        string  `json:"vxMask"`
	INTERP_MTHD    string  `json:"interpMthd"`
	INTERP_PNTS    int     `json:"interpPnts"`
	FCST_THRESH    string  `json:"fcstThresh"`
	OBS_THRESH     string  `json:"obsThresh"`
	COV_THRESH     string  `json:"covThresh"`
	ALPHA          float64 `json:"alpha"`
	LINE_TYPE      string  `json:"lineType"`
}

type STAT_DMAP_header struct {
	VERSION        string  `json:"version"`
	MODEL          string  `json:"model"`
	DESC           string  `json:"desc"`
	FCST_VALID_BEG int     `json:"fcstValidBeg"`
	FCST_VALID_END int     `json:"fcstValidEnd"`
	OBS_LEAD       int     `json:"obsLead"`
	OBS_VALID_BEG  int     `json:"obsValidBeg"`
	OBS_VALID_END  int     `json:"obsValidEnd"`
	FCST_VAR       string  `json:"fcstVar"`
	FCST_UNITS     string  `json:"fcstUnits"`
	FCST_LEV       string  `json:"fcstLev"`
	OBS_VAR        string  `json:"obsVar"`
	OBS_UNITS      string  `json:"obsUnits"`
	OBS_LEV        string  `json:"obsLev"`
	OBTYPE         string  `json:"obtype"`
	VX_MASK        string  `json:"vxMask"`
	INTERP_MTHD    string  `json:"interpMthd"`
	INTERP_PNTS    int     `json:"interpPnts"`
	FCST_THRESH    string  `json:"fcstThresh"`
	OBS_THRESH     string  `json:"obsThresh"`
	COV_THRESH     string  `json:"covThresh"`
	ALPHA          float64 `json:"alpha"`
	LINE_TYPE      string  `json:"lineType"`
}

type STAT_ECLV_header struct {
	VERSION        string  `json:"version"`
	MODEL          string  `json:"model"`
	DESC           string  `json:"desc"`
	FCST_VALID_BEG int     `json:"fcstValidBeg"`
	FCST_VALID_END int     `json:"fcstValidEnd"`
	OBS_LEAD       int     `json:"obsLead"`
	OBS_VALID_BEG  int     `json:"obsValidBeg"`
	OBS_VALID_END  int     `json:"obsValidEnd"`
	FCST_VAR       string  `json:"fcstVar"`
	FCST_UNITS     string  `json:"fcstUnits"`
	FCST_LEV       string  `json:"fcstLev"`
	OBS_VAR        string  `json:"obsVar"`
	OBS_UNITS      string  `json:"obsUnits"`
	OBS_LEV        string  `json:"obsLev"`
	OBTYPE         string  `json:"obtype"`
	VX_MASK        string  `json:"vxMask"`
	INTERP_MTHD    string  `json:"interpMthd"`
	INTERP_PNTS    int     `json:"interpPnts"`
	FCST_THRESH    string  `json:"fcstThresh"`
	OBS_THRESH     string  `json:"obsThresh"`
	COV_THRESH     string  `json:"covThresh"`
	ALPHA          float64 `json:"alpha"`
	LINE_TYPE      string  `json:"lineType"`
}

type STAT_ECNT_header struct {
	VERSION        string  `json:"version"`
	MODEL          string  `json:"model"`
	DESC           string  `json:"desc"`
	FCST_VALID_BEG int     `json:"fcstValidBeg"`
	FCST_VALID_END int     `json:"fcstValidEnd"`
	OBS_LEAD       int     `json:"obsLead"`
	OBS_VALID_BEG  int     `json:"obsValidBeg"`
	OBS_VALID_END  int     `json:"obsValidEnd"`
	FCST_VAR       string  `json:"fcstVar"`
	FCST_UNITS     string  `json:"fcstUnits"`
	FCST_LEV       string  `json:"fcstLev"`
	OBS_VAR        string  `json:"obsVar"`
	OBS_UNITS      string  `json:"obsUnits"`
	OBS_LEV        string  `json:"obsLev"`
	OBTYPE         string  `json:"obtype"`
	VX_MASK        string  `json:"vxMask"`
	INTERP_MTHD    string  `json:"interpMthd"`
	INTERP_PNTS    int     `json:"interpPnts"`
	FCST_THRESH    string  `json:"fcstThresh"`
	OBS_THRESH     string  `json:"obsThresh"`
	COV_THRESH     string  `json:"covThresh"`
	ALPHA          float64 `json:"alpha"`
	LINE_TYPE      string  `json:"lineType"`
}

type STAT_FHO_header struct {
	VERSION        string  `json:"version"`
	MODEL          string  `json:"model"`
	DESC           string  `json:"desc"`
	FCST_VALID_BEG int     `json:"fcstValidBeg"`
	FCST_VALID_END int     `json:"fcstValidEnd"`
	OBS_LEAD       int     `json:"obsLead"`
	OBS_VALID_BEG  int     `json:"obsValidBeg"`
	OBS_VALID_END  int     `json:"obsValidEnd"`
	FCST_VAR       string  `json:"fcstVar"`
	FCST_UNITS     string  `json:"fcstUnits"`
	FCST_LEV       string  `json:"fcstLev"`
	OBS_VAR        string  `json:"obsVar"`
	OBS_UNITS      string  `json:"obsUnits"`
	OBS_LEV        string  `json:"obsLev"`
	OBTYPE         string  `json:"obtype"`
	VX_MASK        string  `json:"vxMask"`
	INTERP_MTHD    string  `json:"interpMthd"`
	INTERP_PNTS    int     `json:"interpPnts"`
	FCST_THRESH    string  `json:"fcstThresh"`
	OBS_THRESH     string  `json:"obsThresh"`
	COV_THRESH     string  `json:"covThresh"`
	ALPHA          float64 `json:"alpha"`
	LINE_TYPE      string  `json:"lineType"`
}

type STAT_GENMPR_header struct {
	VERSION        string  `json:"version"`
	MODEL          string  `json:"model"`
	DESC           string  `json:"desc"`
	FCST_VALID_BEG int     `json:"fcstValidBeg"`
	FCST_VALID_END int     `json:"fcstValidEnd"`
	OBS_LEAD       int     `json:"obsLead"`
	OBS_VALID_BEG  int     `json:"obsValidBeg"`
	OBS_VALID_END  int     `json:"obsValidEnd"`
	FCST_VAR       string  `json:"fcstVar"`
	FCST_UNITS     string  `json:"fcstUnits"`
	FCST_LEV       string  `json:"fcstLev"`
	OBS_VAR        string  `json:"obsVar"`
	OBS_UNITS      string  `json:"obsUnits"`
	OBS_LEV        string  `json:"obsLev"`
	OBTYPE         string  `json:"obtype"`
	VX_MASK        string  `json:"vxMask"`
	INTERP_MTHD    string  `json:"interpMthd"`
	INTERP_PNTS    int     `json:"interpPnts"`
	FCST_THRESH    string  `json:"fcstThresh"`
	OBS_THRESH     string  `json:"obsThresh"`
	COV_THRESH     string  `json:"covThresh"`
	ALPHA          float64 `json:"alpha"`
	LINE_TYPE      string  `json:"lineType"`
}

type STAT_GRAD_header struct {
	VERSION        string  `json:"version"`
	MODEL          string  `json:"model"`
	DESC           string  `json:"desc"`
	FCST_VALID_BEG int     `json:"fcstValidBeg"`
	FCST_VALID_END int     `json:"fcstValidEnd"`
	OBS_LEAD       int     `json:"obsLead"`
	OBS_VALID_BEG  int     `json:"obsValidBeg"`
	OBS_VALID_END  int     `json:"obsValidEnd"`
	FCST_VAR       string  `json:"fcstVar"`
	FCST_UNITS     string  `json:"fcstUnits"`
	FCST_LEV       string  `json:"fcstLev"`
	OBS_VAR        string  `json:"obsVar"`
	OBS_UNITS      string  `json:"obsUnits"`
	OBS_LEV        string  `json:"obsLev"`
	OBTYPE         string  `json:"obtype"`
	VX_MASK        string  `json:"vxMask"`
	INTERP_MTHD    string  `json:"interpMthd"`
	INTERP_PNTS    int     `json:"interpPnts"`
	FCST_THRESH    string  `json:"fcstThresh"`
	OBS_THRESH     string  `json:"obsThresh"`
	COV_THRESH     string  `json:"covThresh"`
	ALPHA          float64 `json:"alpha"`
	LINE_TYPE      string  `json:"lineType"`
}

type STAT_ISC_header struct {
	VERSION        string  `json:"version"`
	MODEL          string  `json:"model"`
	DESC           string  `json:"desc"`
	FCST_VALID_BEG int     `json:"fcstValidBeg"`
	FCST_VALID_END int     `json:"fcstValidEnd"`
	OBS_LEAD       int     `json:"obsLead"`
	OBS_VALID_BEG  int     `json:"obsValidBeg"`
	OBS_VALID_END  int     `json:"obsValidEnd"`
	FCST_VAR       string  `json:"fcstVar"`
	FCST_UNITS     string  `json:"fcstUnits"`
	FCST_LEV       string  `json:"fcstLev"`
	OBS_VAR        string  `json:"obsVar"`
	OBS_UNITS      string  `json:"obsUnits"`
	OBS_LEV        string  `json:"obsLev"`
	OBTYPE         string  `json:"obtype"`
	VX_MASK        string  `json:"vxMask"`
	INTERP_MTHD    string  `json:"interpMthd"`
	INTERP_PNTS    int     `json:"interpPnts"`
	FCST_THRESH    string  `json:"fcstThresh"`
	OBS_THRESH     string  `json:"obsThresh"`
	COV_THRESH     string  `json:"covThresh"`
	ALPHA          float64 `json:"alpha"`
	LINE_TYPE      string  `json:"lineType"`
}

type STAT_MCTC_header struct {
	VERSION        string  `json:"version"`
	MODEL          string  `json:"model"`
	DESC           string  `json:"desc"`
	FCST_VALID_BEG int     `json:"fcstValidBeg"`
	FCST_VALID_END int     `json:"fcstValidEnd"`
	OBS_LEAD       int     `json:"obsLead"`
	OBS_VALID_BEG  int     `json:"obsValidBeg"`
	OBS_VALID_END  int     `json:"obsValidEnd"`
	FCST_VAR       string  `json:"fcstVar"`
	FCST_UNITS     string  `json:"fcstUnits"`
	FCST_LEV       string  `json:"fcstLev"`
	OBS_VAR        string  `json:"obsVar"`
	OBS_UNITS      string  `json:"obsUnits"`
	OBS_LEV        string  `json:"obsLev"`
	OBTYPE         string  `json:"obtype"`
	VX_MASK        string  `json:"vxMask"`
	INTERP_MTHD    string  `json:"interpMthd"`
	INTERP_PNTS    int     `json:"interpPnts"`
	FCST_THRESH    string  `json:"fcstThresh"`
	OBS_THRESH     string  `json:"obsThresh"`
	COV_THRESH     string  `json:"covThresh"`
	ALPHA          float64 `json:"alpha"`
	LINE_TYPE      string  `json:"lineType"`
}

type STAT_MCTS_header struct {
	VERSION        string  `json:"version"`
	MODEL          string  `json:"model"`
	DESC           string  `json:"desc"`
	FCST_VALID_BEG int     `json:"fcstValidBeg"`
	FCST_VALID_END int     `json:"fcstValidEnd"`
	OBS_LEAD       int     `json:"obsLead"`
	OBS_VALID_BEG  int     `json:"obsValidBeg"`
	OBS_VALID_END  int     `json:"obsValidEnd"`
	FCST_VAR       string  `json:"fcstVar"`
	FCST_UNITS     string  `json:"fcstUnits"`
	FCST_LEV       string  `json:"fcstLev"`
	OBS_VAR        string  `json:"obsVar"`
	OBS_UNITS      string  `json:"obsUnits"`
	OBS_LEV        string  `json:"obsLev"`
	OBTYPE         string  `json:"obtype"`
	VX_MASK        string  `json:"vxMask"`
	INTERP_MTHD    string  `json:"interpMthd"`
	INTERP_PNTS    int     `json:"interpPnts"`
	FCST_THRESH    string  `json:"fcstThresh"`
	OBS_THRESH     string  `json:"obsThresh"`
	COV_THRESH     string  `json:"covThresh"`
	ALPHA          float64 `json:"alpha"`
	LINE_TYPE      string  `json:"lineType"`
}

type STAT_MPR_header struct {
	VERSION        string  `json:"version"`
	MODEL          string  `json:"model"`
	DESC           string  `json:"desc"`
	FCST_VALID_BEG int     `json:"fcstValidBeg"`
	FCST_VALID_END int     `json:"fcstValidEnd"`
	OBS_LEAD       int     `json:"obsLead"`
	OBS_VALID_BEG  int     `json:"obsValidBeg"`
	OBS_VALID_END  int     `json:"obsValidEnd"`
	FCST_VAR       string  `json:"fcstVar"`
	FCST_UNITS     string  `json:"fcstUnits"`
	FCST_LEV       string  `json:"fcstLev"`
	OBS_VAR        string  `json:"obsVar"`
	OBS_UNITS      string  `json:"obsUnits"`
	OBS_LEV        string  `json:"obsLev"`
	OBTYPE         string  `json:"obtype"`
	VX_MASK        string  `json:"vxMask"`
	INTERP_MTHD    string  `json:"interpMthd"`
	INTERP_PNTS    int     `json:"interpPnts"`
	FCST_THRESH    string  `json:"fcstThresh"`
	OBS_THRESH     string  `json:"obsThresh"`
	COV_THRESH     string  `json:"covThresh"`
	ALPHA          float64 `json:"alpha"`
	LINE_TYPE      string  `json:"lineType"`
}

type STAT_NBRCNT_header struct {
	VERSION        string  `json:"version"`
	MODEL          string  `json:"model"`
	DESC           string  `json:"desc"`
	FCST_VALID_BEG int     `json:"fcstValidBeg"`
	FCST_VALID_END int     `json:"fcstValidEnd"`
	OBS_LEAD       int     `json:"obsLead"`
	OBS_VALID_BEG  int     `json:"obsValidBeg"`
	OBS_VALID_END  int     `json:"obsValidEnd"`
	FCST_VAR       string  `json:"fcstVar"`
	FCST_UNITS     string  `json:"fcstUnits"`
	FCST_LEV       string  `json:"fcstLev"`
	OBS_VAR        string  `json:"obsVar"`
	OBS_UNITS      string  `json:"obsUnits"`
	OBS_LEV        string  `json:"obsLev"`
	OBTYPE         string  `json:"obtype"`
	VX_MASK        string  `json:"vxMask"`
	INTERP_MTHD    string  `json:"interpMthd"`
	INTERP_PNTS    int     `json:"interpPnts"`
	FCST_THRESH    string  `json:"fcstThresh"`
	OBS_THRESH     string  `json:"obsThresh"`
	COV_THRESH     string  `json:"covThresh"`
	ALPHA          float64 `json:"alpha"`
	LINE_TYPE      string  `json:"lineType"`
}

type STAT_NBRCTC_header struct {
	VERSION        string  `json:"version"`
	MODEL          string  `json:"model"`
	DESC           string  `json:"desc"`
	FCST_VALID_BEG int     `json:"fcstValidBeg"`
	FCST_VALID_END int     `json:"fcstValidEnd"`
	OBS_LEAD       int     `json:"obsLead"`
	OBS_VALID_BEG  int     `json:"obsValidBeg"`
	OBS_VALID_END  int     `json:"obsValidEnd"`
	FCST_VAR       string  `json:"fcstVar"`
	FCST_UNITS     string  `json:"fcstUnits"`
	FCST_LEV       string  `json:"fcstLev"`
	OBS_VAR        string  `json:"obsVar"`
	OBS_UNITS      string  `json:"obsUnits"`
	OBS_LEV        string  `json:"obsLev"`
	OBTYPE         string  `json:"obtype"`
	VX_MASK        string  `json:"vxMask"`
	INTERP_MTHD    string  `json:"interpMthd"`
	INTERP_PNTS    int     `json:"interpPnts"`
	FCST_THRESH    string  `json:"fcstThresh"`
	OBS_THRESH     string  `json:"obsThresh"`
	COV_THRESH     string  `json:"covThresh"`
	ALPHA          float64 `json:"alpha"`
	LINE_TYPE      string  `json:"lineType"`
}

type STAT_NBRCTS_header struct {
	VERSION        string  `json:"version"`
	MODEL          string  `json:"model"`
	DESC           string  `json:"desc"`
	FCST_VALID_BEG int     `json:"fcstValidBeg"`
	FCST_VALID_END int     `json:"fcstValidEnd"`
	OBS_LEAD       int     `json:"obsLead"`
	OBS_VALID_BEG  int     `json:"obsValidBeg"`
	OBS_VALID_END  int     `json:"obsValidEnd"`
	FCST_VAR       string  `json:"fcstVar"`
	FCST_UNITS     string  `json:"fcstUnits"`
	FCST_LEV       string  `json:"fcstLev"`
	OBS_VAR        string  `json:"obsVar"`
	OBS_UNITS      string  `json:"obsUnits"`
	OBS_LEV        string  `json:"obsLev"`
	OBTYPE         string  `json:"obtype"`
	VX_MASK        string  `json:"vxMask"`
	INTERP_MTHD    string  `json:"interpMthd"`
	INTERP_PNTS    int     `json:"interpPnts"`
	FCST_THRESH    string  `json:"fcstThresh"`
	OBS_THRESH     string  `json:"obsThresh"`
	COV_THRESH     string  `json:"covThresh"`
	ALPHA          float64 `json:"alpha"`
	LINE_TYPE      string  `json:"lineType"`
}

type STAT_ORANK_header struct {
	VERSION        string  `json:"version"`
	MODEL          string  `json:"model"`
	DESC           string  `json:"desc"`
	FCST_VALID_BEG int     `json:"fcstValidBeg"`
	FCST_VALID_END int     `json:"fcstValidEnd"`
	OBS_LEAD       int     `json:"obsLead"`
	OBS_VALID_BEG  int     `json:"obsValidBeg"`
	OBS_VALID_END  int     `json:"obsValidEnd"`
	FCST_VAR       string  `json:"fcstVar"`
	FCST_UNITS     string  `json:"fcstUnits"`
	FCST_LEV       string  `json:"fcstLev"`
	OBS_VAR        string  `json:"obsVar"`
	OBS_UNITS      string  `json:"obsUnits"`
	OBS_LEV        string  `json:"obsLev"`
	OBTYPE         string  `json:"obtype"`
	VX_MASK        string  `json:"vxMask"`
	INTERP_MTHD    string  `json:"interpMthd"`
	INTERP_PNTS    int     `json:"interpPnts"`
	FCST_THRESH    string  `json:"fcstThresh"`
	OBS_THRESH     string  `json:"obsThresh"`
	COV_THRESH     string  `json:"covThresh"`
	ALPHA          float64 `json:"alpha"`
	LINE_TYPE      string  `json:"lineType"`
}

type STAT_PCT_header struct {
	VERSION        string  `json:"version"`
	MODEL          string  `json:"model"`
	DESC           string  `json:"desc"`
	FCST_VALID_BEG int     `json:"fcstValidBeg"`
	FCST_VALID_END int     `json:"fcstValidEnd"`
	OBS_LEAD       int     `json:"obsLead"`
	OBS_VALID_BEG  int     `json:"obsValidBeg"`
	OBS_VALID_END  int     `json:"obsValidEnd"`
	FCST_VAR       string  `json:"fcstVar"`
	FCST_UNITS     string  `json:"fcstUnits"`
	FCST_LEV       string  `json:"fcstLev"`
	OBS_VAR        string  `json:"obsVar"`
	OBS_UNITS      string  `json:"obsUnits"`
	OBS_LEV        string  `json:"obsLev"`
	OBTYPE         string  `json:"obtype"`
	VX_MASK        string  `json:"vxMask"`
	INTERP_MTHD    string  `json:"interpMthd"`
	INTERP_PNTS    int     `json:"interpPnts"`
	FCST_THRESH    string  `json:"fcstThresh"`
	OBS_THRESH     string  `json:"obsThresh"`
	COV_THRESH     string  `json:"covThresh"`
	ALPHA          float64 `json:"alpha"`
	LINE_TYPE      string  `json:"lineType"`
}

type STAT_PHIST_header struct {
	VERSION        string  `json:"version"`
	MODEL          string  `json:"model"`
	DESC           string  `json:"desc"`
	FCST_VALID_BEG int     `json:"fcstValidBeg"`
	FCST_VALID_END int     `json:"fcstValidEnd"`
	OBS_LEAD       int     `json:"obsLead"`
	OBS_VALID_BEG  int     `json:"obsValidBeg"`
	OBS_VALID_END  int     `json:"obsValidEnd"`
	FCST_VAR       string  `json:"fcstVar"`
	FCST_UNITS     string  `json:"fcstUnits"`
	FCST_LEV       string  `json:"fcstLev"`
	OBS_VAR        string  `json:"obsVar"`
	OBS_UNITS      string  `json:"obsUnits"`
	OBS_LEV        string  `json:"obsLev"`
	OBTYPE         string  `json:"obtype"`
	VX_MASK        string  `json:"vxMask"`
	INTERP_MTHD    string  `json:"interpMthd"`
	INTERP_PNTS    int     `json:"interpPnts"`
	FCST_THRESH    string  `json:"fcstThresh"`
	OBS_THRESH     string  `json:"obsThresh"`
	COV_THRESH     string  `json:"covThresh"`
	ALPHA          float64 `json:"alpha"`
	LINE_TYPE      string  `json:"lineType"`
}

type STAT_PJC_header struct {
	VERSION        string  `json:"version"`
	MODEL          string  `json:"model"`
	DESC           string  `json:"desc"`
	FCST_VALID_BEG int     `json:"fcstValidBeg"`
	FCST_VALID_END int     `json:"fcstValidEnd"`
	OBS_LEAD       int     `json:"obsLead"`
	OBS_VALID_BEG  int     `json:"obsValidBeg"`
	OBS_VALID_END  int     `json:"obsValidEnd"`
	FCST_VAR       string  `json:"fcstVar"`
	FCST_UNITS     string  `json:"fcstUnits"`
	FCST_LEV       string  `json:"fcstLev"`
	OBS_VAR        string  `json:"obsVar"`
	OBS_UNITS      string  `json:"obsUnits"`
	OBS_LEV        string  `json:"obsLev"`
	OBTYPE         string  `json:"obtype"`
	VX_MASK        string  `json:"vxMask"`
	INTERP_MTHD    string  `json:"interpMthd"`
	INTERP_PNTS    int     `json:"interpPnts"`
	FCST_THRESH    string  `json:"fcstThresh"`
	OBS_THRESH     string  `json:"obsThresh"`
	COV_THRESH     string  `json:"covThresh"`
	ALPHA          float64 `json:"alpha"`
	LINE_TYPE      string  `json:"lineType"`
}

type STAT_PRC_header struct {
	VERSION        string  `json:"version"`
	MODEL          string  `json:"model"`
	DESC           string  `json:"desc"`
	FCST_VALID_BEG int     `json:"fcstValidBeg"`
	FCST_VALID_END int     `json:"fcstValidEnd"`
	OBS_LEAD       int     `json:"obsLead"`
	OBS_VALID_BEG  int     `json:"obsValidBeg"`
	OBS_VALID_END  int     `json:"obsValidEnd"`
	FCST_VAR       string  `json:"fcstVar"`
	FCST_UNITS     string  `json:"fcstUnits"`
	FCST_LEV       string  `json:"fcstLev"`
	OBS_VAR        string  `json:"obsVar"`
	OBS_UNITS      string  `json:"obsUnits"`
	OBS_LEV        string  `json:"obsLev"`
	OBTYPE         string  `json:"obtype"`
	VX_MASK        string  `json:"vxMask"`
	INTERP_MTHD    string  `json:"interpMthd"`
	INTERP_PNTS    int     `json:"interpPnts"`
	FCST_THRESH    string  `json:"fcstThresh"`
	OBS_THRESH     string  `json:"obsThresh"`
	COV_THRESH     string  `json:"covThresh"`
	ALPHA          float64 `json:"alpha"`
	LINE_TYPE      string  `json:"lineType"`
}

type STAT_PSTD_header struct {
	VERSION        string  `json:"version"`
	MODEL          string  `json:"model"`
	DESC           string  `json:"desc"`
	FCST_VALID_BEG int     `json:"fcstValidBeg"`
	FCST_VALID_END int     `json:"fcstValidEnd"`
	OBS_LEAD       int     `json:"obsLead"`
	OBS_VALID_BEG  int     `json:"obsValidBeg"`
	OBS_VALID_END  int     `json:"obsValidEnd"`
	FCST_VAR       string  `json:"fcstVar"`
	FCST_UNITS     string  `json:"fcstUnits"`
	FCST_LEV       string  `json:"fcstLev"`
	OBS_VAR        string  `json:"obsVar"`
	OBS_UNITS      string  `json:"obsUnits"`
	OBS_LEV        string  `json:"obsLev"`
	OBTYPE         string  `json:"obtype"`
	VX_MASK        string  `json:"vxMask"`
	INTERP_MTHD    string  `json:"interpMthd"`
	INTERP_PNTS    int     `json:"interpPnts"`
	FCST_THRESH    string  `json:"fcstThresh"`
	OBS_THRESH     string  `json:"obsThresh"`
	COV_THRESH     string  `json:"covThresh"`
	ALPHA          float64 `json:"alpha"`
	LINE_TYPE      string  `json:"lineType"`
}

type STAT_RELP_header struct {
	VERSION        string  `json:"version"`
	MODEL          string  `json:"model"`
	DESC           string  `json:"desc"`
	FCST_VALID_BEG int     `json:"fcstValidBeg"`
	FCST_VALID_END int     `json:"fcstValidEnd"`
	OBS_LEAD       int     `json:"obsLead"`
	OBS_VALID_BEG  int     `json:"obsValidBeg"`
	OBS_VALID_END  int     `json:"obsValidEnd"`
	FCST_VAR       string  `json:"fcstVar"`
	FCST_UNITS     string  `json:"fcstUnits"`
	FCST_LEV       string  `json:"fcstLev"`
	OBS_VAR        string  `json:"obsVar"`
	OBS_UNITS      string  `json:"obsUnits"`
	OBS_LEV        string  `json:"obsLev"`
	OBTYPE         string  `json:"obtype"`
	VX_MASK        string  `json:"vxMask"`
	INTERP_MTHD    string  `json:"interpMthd"`
	INTERP_PNTS    int     `json:"interpPnts"`
	FCST_THRESH    string  `json:"fcstThresh"`
	OBS_THRESH     string  `json:"obsThresh"`
	COV_THRESH     string  `json:"covThresh"`
	ALPHA          float64 `json:"alpha"`
	LINE_TYPE      string  `json:"lineType"`
}

type STAT_RHIST_header struct {
	VERSION        string  `json:"version"`
	MODEL          string  `json:"model"`
	DESC           string  `json:"desc"`
	FCST_VALID_BEG int     `json:"fcstValidBeg"`
	FCST_VALID_END int     `json:"fcstValidEnd"`
	OBS_LEAD       int     `json:"obsLead"`
	OBS_VALID_BEG  int     `json:"obsValidBeg"`
	OBS_VALID_END  int     `json:"obsValidEnd"`
	FCST_VAR       string  `json:"fcstVar"`
	FCST_UNITS     string  `json:"fcstUnits"`
	FCST_LEV       string  `json:"fcstLev"`
	OBS_VAR        string  `json:"obsVar"`
	OBS_UNITS      string  `json:"obsUnits"`
	OBS_LEV        string  `json:"obsLev"`
	OBTYPE         string  `json:"obtype"`
	VX_MASK        string  `json:"vxMask"`
	INTERP_MTHD    string  `json:"interpMthd"`
	INTERP_PNTS    int     `json:"interpPnts"`
	FCST_THRESH    string  `json:"fcstThresh"`
	OBS_THRESH     string  `json:"obsThresh"`
	COV_THRESH     string  `json:"covThresh"`
	ALPHA          float64 `json:"alpha"`
	LINE_TYPE      string  `json:"lineType"`
}

type STAT_RPS_header struct {
	VERSION        string  `json:"version"`
	MODEL          string  `json:"model"`
	DESC           string  `json:"desc"`
	FCST_VALID_BEG int     `json:"fcstValidBeg"`
	FCST_VALID_END int     `json:"fcstValidEnd"`
	OBS_LEAD       int     `json:"obsLead"`
	OBS_VALID_BEG  int     `json:"obsValidBeg"`
	OBS_VALID_END  int     `json:"obsValidEnd"`
	FCST_VAR       string  `json:"fcstVar"`
	FCST_UNITS     string  `json:"fcstUnits"`
	FCST_LEV       string  `json:"fcstLev"`
	OBS_VAR        string  `json:"obsVar"`
	OBS_UNITS      string  `json:"obsUnits"`
	OBS_LEV        string  `json:"obsLev"`
	OBTYPE         string  `json:"obtype"`
	VX_MASK        string  `json:"vxMask"`
	INTERP_MTHD    string  `json:"interpMthd"`
	INTERP_PNTS    int     `json:"interpPnts"`
	FCST_THRESH    string  `json:"fcstThresh"`
	OBS_THRESH     string  `json:"obsThresh"`
	COV_THRESH     string  `json:"covThresh"`
	ALPHA          float64 `json:"alpha"`
	LINE_TYPE      string  `json:"lineType"`
}

type STAT_SAL1L2_header struct {
	VERSION        string  `json:"version"`
	MODEL          string  `json:"model"`
	DESC           string  `json:"desc"`
	FCST_VALID_BEG int     `json:"fcstValidBeg"`
	FCST_VALID_END int     `json:"fcstValidEnd"`
	OBS_LEAD       int     `json:"obsLead"`
	OBS_VALID_BEG  int     `json:"obsValidBeg"`
	OBS_VALID_END  int     `json:"obsValidEnd"`
	FCST_VAR       string  `json:"fcstVar"`
	FCST_UNITS     string  `json:"fcstUnits"`
	FCST_LEV       string  `json:"fcstLev"`
	OBS_VAR        string  `json:"obsVar"`
	OBS_UNITS      string  `json:"obsUnits"`
	OBS_LEV        string  `json:"obsLev"`
	OBTYPE         string  `json:"obtype"`
	VX_MASK        string  `json:"vxMask"`
	INTERP_MTHD    string  `json:"interpMthd"`
	INTERP_PNTS    int     `json:"interpPnts"`
	FCST_THRESH    string  `json:"fcstThresh"`
	OBS_THRESH     string  `json:"obsThresh"`
	COV_THRESH     string  `json:"covThresh"`
	ALPHA          float64 `json:"alpha"`
	LINE_TYPE      string  `json:"lineType"`
}

type STAT_SL1L2_header struct {
	VERSION        string  `json:"version"`
	MODEL          string  `json:"model"`
	DESC           string  `json:"desc"`
	FCST_VALID_BEG int     `json:"fcstValidBeg"`
	FCST_VALID_END int     `json:"fcstValidEnd"`
	OBS_LEAD       int     `json:"obsLead"`
	OBS_VALID_BEG  int     `json:"obsValidBeg"`
	OBS_VALID_END  int     `json:"obsValidEnd"`
	FCST_VAR       string  `json:"fcstVar"`
	FCST_UNITS     string  `json:"fcstUnits"`
	FCST_LEV       string  `json:"fcstLev"`
	OBS_VAR        string  `json:"obsVar"`
	OBS_UNITS      string  `json:"obsUnits"`
	OBS_LEV        string  `json:"obsLev"`
	OBTYPE         string  `json:"obtype"`
	VX_MASK        string  `json:"vxMask"`
	INTERP_MTHD    string  `json:"interpMthd"`
	INTERP_PNTS    int     `json:"interpPnts"`
	FCST_THRESH    string  `json:"fcstThresh"`
	OBS_THRESH     string  `json:"obsThresh"`
	COV_THRESH     string  `json:"covThresh"`
	ALPHA          float64 `json:"alpha"`
	LINE_TYPE      string  `json:"lineType"`
}

type STAT_SSIDX_header struct {
	VERSION        string  `json:"version"`
	MODEL          string  `json:"model"`
	DESC           string  `json:"desc"`
	FCST_VALID_BEG int     `json:"fcstValidBeg"`
	FCST_VALID_END int     `json:"fcstValidEnd"`
	OBS_LEAD       int     `json:"obsLead"`
	OBS_VALID_BEG  int     `json:"obsValidBeg"`
	OBS_VALID_END  int     `json:"obsValidEnd"`
	FCST_VAR       string  `json:"fcstVar"`
	FCST_UNITS     string  `json:"fcstUnits"`
	FCST_LEV       string  `json:"fcstLev"`
	OBS_VAR        string  `json:"obsVar"`
	OBS_UNITS      string  `json:"obsUnits"`
	OBS_LEV        string  `json:"obsLev"`
	OBTYPE         string  `json:"obtype"`
	VX_MASK        string  `json:"vxMask"`
	INTERP_MTHD    string  `json:"interpMthd"`
	INTERP_PNTS    int     `json:"interpPnts"`
	FCST_THRESH    string  `json:"fcstThresh"`
	OBS_THRESH     string  `json:"obsThresh"`
	COV_THRESH     string  `json:"covThresh"`
	ALPHA          float64 `json:"alpha"`
	LINE_TYPE      string  `json:"lineType"`
}

type STAT_SSVAR_header struct {
	VERSION        string  `json:"version"`
	MODEL          string  `json:"model"`
	DESC           string  `json:"desc"`
	FCST_VALID_BEG int     `json:"fcstValidBeg"`
	FCST_VALID_END int     `json:"fcstValidEnd"`
	OBS_LEAD       int     `json:"obsLead"`
	OBS_VALID_BEG  int     `json:"obsValidBeg"`
	OBS_VALID_END  int     `json:"obsValidEnd"`
	FCST_VAR       string  `json:"fcstVar"`
	FCST_UNITS     string  `json:"fcstUnits"`
	FCST_LEV       string  `json:"fcstLev"`
	OBS_VAR        string  `json:"obsVar"`
	OBS_UNITS      string  `json:"obsUnits"`
	OBS_LEV        string  `json:"obsLev"`
	OBTYPE         string  `json:"obtype"`
	VX_MASK        string  `json:"vxMask"`
	INTERP_MTHD    string  `json:"interpMthd"`
	INTERP_PNTS    int     `json:"interpPnts"`
	FCST_THRESH    string  `json:"fcstThresh"`
	OBS_THRESH     string  `json:"obsThresh"`
	COV_THRESH     string  `json:"covThresh"`
	ALPHA          float64 `json:"alpha"`
	LINE_TYPE      string  `json:"lineType"`
}

type STAT_VAL1L2_header struct {
	VERSION        string  `json:"version"`
	MODEL          string  `json:"model"`
	DESC           string  `json:"desc"`
	FCST_VALID_BEG int     `json:"fcstValidBeg"`
	FCST_VALID_END int     `json:"fcstValidEnd"`
	OBS_LEAD       int     `json:"obsLead"`
	OBS_VALID_BEG  int     `json:"obsValidBeg"`
	OBS_VALID_END  int     `json:"obsValidEnd"`
	FCST_VAR       string  `json:"fcstVar"`
	FCST_UNITS     string  `json:"fcstUnits"`
	FCST_LEV       string  `json:"fcstLev"`
	OBS_VAR        string  `json:"obsVar"`
	OBS_UNITS      string  `json:"obsUnits"`
	OBS_LEV        string  `json:"obsLev"`
	OBTYPE         string  `json:"obtype"`
	VX_MASK        string  `json:"vxMask"`
	INTERP_MTHD    string  `json:"interpMthd"`
	INTERP_PNTS    int     `json:"interpPnts"`
	FCST_THRESH    string  `json:"fcstThresh"`
	OBS_THRESH     string  `json:"obsThresh"`
	COV_THRESH     string  `json:"covThresh"`
	ALPHA          float64 `json:"alpha"`
	LINE_TYPE      string  `json:"lineType"`
}

type STAT_VCNT_header struct {
	VERSION        string  `json:"version"`
	MODEL          string  `json:"model"`
	DESC           string  `json:"desc"`
	FCST_VALID_BEG int     `json:"fcstValidBeg"`
	FCST_VALID_END int     `json:"fcstValidEnd"`
	OBS_LEAD       int     `json:"obsLead"`
	OBS_VALID_BEG  int     `json:"obsValidBeg"`
	OBS_VALID_END  int     `json:"obsValidEnd"`
	FCST_VAR       string  `json:"fcstVar"`
	FCST_UNITS     string  `json:"fcstUnits"`
	FCST_LEV       string  `json:"fcstLev"`
	OBS_VAR        string  `json:"obsVar"`
	OBS_UNITS      string  `json:"obsUnits"`
	OBS_LEV        string  `json:"obsLev"`
	OBTYPE         string  `json:"obtype"`
	VX_MASK        string  `json:"vxMask"`
	INTERP_MTHD    string  `json:"interpMthd"`
	INTERP_PNTS    int     `json:"interpPnts"`
	FCST_THRESH    string  `json:"fcstThresh"`
	OBS_THRESH     string  `json:"obsThresh"`
	COV_THRESH     string  `json:"covThresh"`
	ALPHA          float64 `json:"alpha"`
	LINE_TYPE      string  `json:"lineType"`
}

type STAT_VL1L2_header struct {
	VERSION        string  `json:"version"`
	MODEL          string  `json:"model"`
	DESC           string  `json:"desc"`
	FCST_VALID_BEG int     `json:"fcstValidBeg"`
	FCST_VALID_END int     `json:"fcstValidEnd"`
	OBS_LEAD       int     `json:"obsLead"`
	OBS_VALID_BEG  int     `json:"obsValidBeg"`
	OBS_VALID_END  int     `json:"obsValidEnd"`
	FCST_VAR       string  `json:"fcstVar"`
	FCST_UNITS     string  `json:"fcstUnits"`
	FCST_LEV       string  `json:"fcstLev"`
	OBS_VAR        string  `json:"obsVar"`
	OBS_UNITS      string  `json:"obsUnits"`
	OBS_LEV        string  `json:"obsLev"`
	OBTYPE         string  `json:"obtype"`
	VX_MASK        string  `json:"vxMask"`
	INTERP_MTHD    string  `json:"interpMthd"`
	INTERP_PNTS    int     `json:"interpPnts"`
	FCST_THRESH    string  `json:"fcstThresh"`
	OBS_THRESH     string  `json:"obsThresh"`
	COV_THRESH     string  `json:"covThresh"`
	ALPHA          float64 `json:"alpha"`
	LINE_TYPE      string  `json:"lineType"`
}

type TCST_PROBRIRW_header struct {
	VERSION    string `json:"version"`
	AMODEL     string `json:"amodel"`
	BMODEL     string `json:"bmodel"`
	DESC       string `json:"desc"`
	STORM_ID   string `json:"stormId"`
	BASIN      string `json:"basin"`
	CYCLONE    string `json:"cyclone"`
	STORM_NAME string `json:"stormName"`
	VALID      int    `json:"valid"`
	INIT_MASK  string `json:"initMask"`
	VALID_MASK string `json:"validMask"`
	LINE_TYPE  string `json:"lineType"`
}

type TCST_TCMPR_header struct {
	VERSION    string `json:"version"`
	AMODEL     string `json:"amodel"`
	BMODEL     string `json:"bmodel"`
	DESC       string `json:"desc"`
	STORM_ID   string `json:"stormId"`
	BASIN      string `json:"basin"`
	CYCLONE    string `json:"cyclone"`
	STORM_NAME string `json:"stormName"`
	VALID      int    `json:"valid"`
	INIT_MASK  string `json:"initMask"`
	VALID_MASK string `json:"validMask"`
	LINE_TYPE  string `json:"lineType"`
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
	FIELD string  `json:"field,omitempty"`
	TOTAL int     `json:"total,omitempty"`
	FY_OY float64 `json:"fyOy,omitempty"`
	FY_ON float64 `json:"fyOn,omitempty"`
	FN_OY float64 `json:"fnOy,omitempty"`
	FN_ON float64 `json:"fnOn,omitempty"`
	BASER float64 `json:"baser,omitempty"`
	FMEAN float64 `json:"fmean,omitempty"`
	ACC   float64 `json:"acc,omitempty"`
	FBIAS float64 `json:"fbias,omitempty"`
	PODY  float64 `json:"pody,omitempty"`
	PODN  float64 `json:"podn,omitempty"`
	POFD  float64 `json:"pofd,omitempty"`
	FAR   float64 `json:"far,omitempty"`
	CSI   float64 `json:"csi,omitempty"`
	GSS   float64 `json:"gss,omitempty"`
	HK    float64 `json:"hk,omitempty"`
	HSS   float64 `json:"hss,omitempty"`
	ODDS  float64 `json:"odds,omitempty"`
}

type MODE_OBJ struct {
	OBJECT_ID                  string  `json:"objectId,omitempty"`
	OBJECT_CAT                 string  `json:"objectCat,omitempty"`
	CENTROID_X                 float64 `json:"centroidX,omitempty"`
	CENTROID_Y                 float64 `json:"centroidY,omitempty"`
	CENTROID_LAT               float64 `json:"centroidLat,omitempty"`
	CENTROID_LON               float64 `json:"centroidLon,omitempty"`
	AXIS_ANG                   float64 `json:"axisAng,omitempty"`
	LENGTH                     float64 `json:"length,omitempty"`
	WIDTH                      float64 `json:"width,omitempty"`
	AREA                       int     `json:"area,omitempty"`
	AREA_THRESH                int     `json:"areaThresh,omitempty"`
	CURVATURE                  float64 `json:"curvature,omitempty"`
	CURVATURE_X                float64 `json:"curvatureX,omitempty"`
	CURVATURE_Y                float64 `json:"curvatureY,omitempty"`
	COMPLEXITY                 float64 `json:"complexity,omitempty"`
	INTENSITY_10               float64 `json:"intensity10,omitempty"`
	INTENSITY_25               float64 `json:"intensity25,omitempty"`
	INTENSITY_50               float64 `json:"intensity50,omitempty"`
	INTENSITY_75               float64 `json:"intensity75,omitempty"`
	INTENSITY_90               float64 `json:"intensity90,omitempty"`
	INTENSITY_USER             float64 `json:"intensityUser,omitempty"`
	INTENSITY_SUM              float64 `json:"intensitySum,omitempty"`
	CENTROID_DIST              float64 `json:"centroidDist,omitempty"`
	BOUNDARY_DIST              float64 `json:"boundaryDist,omitempty"`
	CONVEX_HULL_DIST           float64 `json:"convexHullDist,omitempty"`
	ANGLE_DIFF                 float64 `json:"angleDiff,omitempty"`
	ASPECT_DIFF                float64 `json:"aspectDiff,omitempty"`
	AREA_RATIO                 float64 `json:"areaRatio,omitempty"`
	INTERSECTION_AREA          float64 `json:"intersectionArea,omitempty"`
	UNION_AREA                 float64 `json:"unionArea,omitempty"`
	SYMMETRIC_DIFF             float64 `json:"symmetricDiff,omitempty"`
	INTERSECTION_OVER_AREA     float64 `json:"intersectionOverArea,omitempty"`
	CURVATURE_RATIO            float64 `json:"curvatureRatio,omitempty"`
	COMPLEXITY_RATIO           float64 `json:"complexityRatio,omitempty"`
	PERCENTILE_INTENSITY_RATIO float64 `json:"percentileIntensityRatio,omitempty"`
	INTEREST                   float64 `json:"interest,omitempty"`
}

type STAT_CNT struct {
	TOTAL                int     `json:"total,omitempty"`
	FBAR                 float64 `json:"fbar,omitempty"`
	FBAR_NCL             float64 `json:"fbarNcl,omitempty"`
	FBAR_NCU             float64 `json:"fbarNcu,omitempty"`
	FBAR_BCL             float64 `json:"fbarBcl,omitempty"`
	FBAR_BCU             float64 `json:"fbarBcu,omitempty"`
	FSTDEV               float64 `json:"fstdev,omitempty"`
	FSTDEV_NCL           float64 `json:"fstdevNcl,omitempty"`
	FSTDEV_NCU           float64 `json:"fstdevNcu,omitempty"`
	FSTDEV_BCL           float64 `json:"fstdevBcl,omitempty"`
	FSTDEV_BCU           float64 `json:"fstdevBcu,omitempty"`
	OBAR                 float64 `json:"obar,omitempty"`
	OBAR_NCL             float64 `json:"obarNcl,omitempty"`
	OBAR_NCU             float64 `json:"obarNcu,omitempty"`
	OBAR_BCL             float64 `json:"obarBcl,omitempty"`
	OBAR_BCU             float64 `json:"obarBcu,omitempty"`
	OSTDEV               float64 `json:"ostdev,omitempty"`
	OSTDEV_NCL           float64 `json:"ostdevNcl,omitempty"`
	OSTDEV_NCU           float64 `json:"ostdevNcu,omitempty"`
	OSTDEV_BCL           float64 `json:"ostdevBcl,omitempty"`
	OSTDEV_BCU           float64 `json:"ostdevBcu,omitempty"`
	PR_CORR              float64 `json:"prCorr,omitempty"`
	PR_CORR_NCL          float64 `json:"prCorrNcl,omitempty"`
	PR_CORR_NCU          float64 `json:"prCorrNcu,omitempty"`
	PR_CORR_BCL          float64 `json:"prCorrBcl,omitempty"`
	PR_CORR_BCU          float64 `json:"prCorrBcu,omitempty"`
	SP_CORR              float64 `json:"spCorr,omitempty"`
	KT_CORR              float64 `json:"ktCorr,omitempty"`
	RANKS                int     `json:"ranks,omitempty"`
	FRANK_TIES           int     `json:"frankTies,omitempty"`
	ORANK_TIES           int     `json:"orankTies,omitempty"`
	ME                   float64 `json:"me,omitempty"`
	ME_NCL               float64 `json:"meNcl,omitempty"`
	ME_NCU               float64 `json:"meNcu,omitempty"`
	ME_BCL               float64 `json:"meBcl,omitempty"`
	ME_BCU               float64 `json:"meBcu,omitempty"`
	ESTDEV               float64 `json:"estdev,omitempty"`
	ESTDEV_NCL           float64 `json:"estdevNcl,omitempty"`
	ESTDEV_NCU           float64 `json:"estdevNcu,omitempty"`
	ESTDEV_BCL           float64 `json:"estdevBcl,omitempty"`
	ESTDEV_BCU           float64 `json:"estdevBcu,omitempty"`
	MBIAS                float64 `json:"mbias,omitempty"`
	MBIAS_BCL            float64 `json:"mbiasBcl,omitempty"`
	MBIAS_BCU            float64 `json:"mbiasBcu,omitempty"`
	MAE                  float64 `json:"mae,omitempty"`
	MAE_BCL              float64 `json:"maeBcl,omitempty"`
	MAE_BCU              float64 `json:"maeBcu,omitempty"`
	MSE                  float64 `json:"mse,omitempty"`
	MSE_BCL              float64 `json:"mseBcl,omitempty"`
	MSE_BCU              float64 `json:"mseBcu,omitempty"`
	BCMSE                float64 `json:"bcmse,omitempty"`
	BCMSE_BCL            float64 `json:"bcmseBcl,omitempty"`
	BCMSE_BCU            float64 `json:"bcmseBcu,omitempty"`
	RMSE                 float64 `json:"rmse,omitempty"`
	RMSE_BCL             float64 `json:"rmseBcl,omitempty"`
	RMSE_BCU             float64 `json:"rmseBcu,omitempty"`
	E10                  float64 `json:"e10,omitempty"`
	E10_BCL              float64 `json:"e10Bcl,omitempty"`
	E10_BCU              float64 `json:"e10Bcu,omitempty"`
	E25                  float64 `json:"e25,omitempty"`
	E25_BCL              float64 `json:"e25Bcl,omitempty"`
	E25_BCU              float64 `json:"e25Bcu,omitempty"`
	E50                  float64 `json:"e50,omitempty"`
	E50_BCL              float64 `json:"e50Bcl,omitempty"`
	E50_BCU              float64 `json:"e50Bcu,omitempty"`
	E75                  float64 `json:"e75,omitempty"`
	E75_BCL              float64 `json:"e75Bcl,omitempty"`
	E75_BCU              float64 `json:"e75Bcu,omitempty"`
	E90                  float64 `json:"e90,omitempty"`
	E90_BCL              float64 `json:"e90Bcl,omitempty"`
	E90_BCU              float64 `json:"e90Bcu,omitempty"`
	EIQR                 float64 `json:"eiqr,omitempty"`
	EIQR_BCL             float64 `json:"eiqrBcl,omitempty"`
	EIQR_BCU             float64 `json:"eiqrBcu,omitempty"`
	MAD                  float64 `json:"mad,omitempty"`
	MAD_BCL              float64 `json:"madBcl,omitempty"`
	MAD_BCU              float64 `json:"madBcu,omitempty"`
	ANOM_CORR            float64 `json:"anomCorr,omitempty"`
	ANOM_CORR_NCL        float64 `json:"anomCorrNcl,omitempty"`
	ANOM_CORR_NCU        float64 `json:"anomCorrNcu,omitempty"`
	ANOM_CORR_BCL        float64 `json:"anomCorrBcl,omitempty"`
	ANOM_CORR_BCU        float64 `json:"anomCorrBcu,omitempty"`
	ME2                  float64 `json:"me2,omitempty"`
	ME2_BCL              float64 `json:"me2Bcl,omitempty"`
	ME2_BCU              float64 `json:"me2Bcu,omitempty"`
	MSESS                float64 `json:"msess,omitempty"`
	MSESS_BCL            float64 `json:"msessBcl,omitempty"`
	MSESS_BCU            float64 `json:"msessBcu,omitempty"`
	RMSFA                float64 `json:"rmsfa,omitempty"`
	RMSFA_BCL            float64 `json:"rmsfaBcl,omitempty"`
	RMSFA_BCU            float64 `json:"rmsfaBcu,omitempty"`
	RMSOA                float64 `json:"rmsoa,omitempty"`
	RMSOA_BCL            float64 `json:"rmsoaBcl,omitempty"`
	RMSOA_BCU            float64 `json:"rmsoaBcu,omitempty"`
	ANOM_CORR_UNCNTR     float64 `json:"anomCorrUncntr,omitempty"`
	ANOM_CORR_UNCNTR_BCL float64 `json:"anomCorrUncntrBcl,omitempty"`
	ANOM_CORR_UNCNTR_BCU float64 `json:"anomCorrUncntrBcu,omitempty"`
	SI                   float64 `json:"si,omitempty"`
	SI_BCL               float64 `json:"siBcl,omitempty"`
	SI_BCU               float64 `json:"siBcu,omitempty"`
}

type STAT_CTC struct {
	TOTAL int     `json:"total,omitempty"`
	FY_OY float64 `json:"fyOy,omitempty"`
	FY_ON float64 `json:"fyOn,omitempty"`
	FN_OY float64 `json:"fnOy,omitempty"`
	FN_ON float64 `json:"fnOn,omitempty"`
}

type STAT_CTS struct {
	TOTAL     int     `json:"total,omitempty"`
	BASER     float64 `json:"baser,omitempty"`
	BASER_NCL float64 `json:"baserNcl,omitempty"`
	BASER_NCU float64 `json:"baserNcu,omitempty"`
	BASER_BCL float64 `json:"baserBcl,omitempty"`
	BASER_BCU float64 `json:"baserBcu,omitempty"`
	FMEAN     float64 `json:"fmean,omitempty"`
	FMEAN_NCL float64 `json:"fmeanNcl,omitempty"`
	FMEAN_NCU float64 `json:"fmeanNcu,omitempty"`
	FMEAN_BCL float64 `json:"fmeanBcl,omitempty"`
	FMEAN_BCU float64 `json:"fmeanBcu,omitempty"`
	ACC       float64 `json:"acc,omitempty"`
	ACC_NCL   float64 `json:"accNcl,omitempty"`
	ACC_NCU   float64 `json:"accNcu,omitempty"`
	ACC_BCL   float64 `json:"accBcl,omitempty"`
	ACC_BCU   float64 `json:"accBcu,omitempty"`
	FBIAS     float64 `json:"fbias,omitempty"`
	FBIAS_BCL float64 `json:"fbiasBcl,omitempty"`
	FBIAS_BCU float64 `json:"fbiasBcu,omitempty"`
	PODY      float64 `json:"pody,omitempty"`
	PODY_NCL  float64 `json:"podyNcl,omitempty"`
	PODY_NCU  float64 `json:"podyNcu,omitempty"`
	PODY_BCL  float64 `json:"podyBcl,omitempty"`
	PODY_BCU  float64 `json:"podyBcu,omitempty"`
	PODN      float64 `json:"podn,omitempty"`
	PODN_NCL  float64 `json:"podnNcl,omitempty"`
	PODN_NCU  float64 `json:"podnNcu,omitempty"`
	PODN_BCL  float64 `json:"podnBcl,omitempty"`
	PODN_BCU  float64 `json:"podnBcu,omitempty"`
	POFD      float64 `json:"pofd,omitempty"`
	POFD_NCL  float64 `json:"pofdNcl,omitempty"`
	POFD_NCU  float64 `json:"pofdNcu,omitempty"`
	POFD_BCL  float64 `json:"pofdBcl,omitempty"`
	POFD_BCU  float64 `json:"pofdBcu,omitempty"`
	FAR       float64 `json:"far,omitempty"`
	FAR_NCL   float64 `json:"farNcl,omitempty"`
	FAR_NCU   float64 `json:"farNcu,omitempty"`
	FAR_BCL   float64 `json:"farBcl,omitempty"`
	FAR_BCU   float64 `json:"farBcu,omitempty"`
	CSI       float64 `json:"csi,omitempty"`
	CSI_NCL   float64 `json:"csiNcl,omitempty"`
	CSI_NCU   float64 `json:"csiNcu,omitempty"`
	CSI_BCL   float64 `json:"csiBcl,omitempty"`
	CSI_BCU   float64 `json:"csiBcu,omitempty"`
	GSS       float64 `json:"gss,omitempty"`
	GSS_BCL   float64 `json:"gssBcl,omitempty"`
	GSS_BCU   float64 `json:"gssBcu,omitempty"`
	HK        float64 `json:"hk,omitempty"`
	HK_NCL    float64 `json:"hkNcl,omitempty"`
	HK_NCU    float64 `json:"hkNcu,omitempty"`
	HK_BCL    float64 `json:"hkBcl,omitempty"`
	HK_BCU    float64 `json:"hkBcu,omitempty"`
	HSS       float64 `json:"hss,omitempty"`
	HSS_BCL   float64 `json:"hssBcl,omitempty"`
	HSS_BCU   float64 `json:"hssBcu,omitempty"`
	ODDS      float64 `json:"odds,omitempty"`
	ODDS_NCL  float64 `json:"oddsNcl,omitempty"`
	ODDS_NCU  float64 `json:"oddsNcu,omitempty"`
	ODDS_BCL  float64 `json:"oddsBcl,omitempty"`
	ODDS_BCU  float64 `json:"oddsBcu,omitempty"`
	LODDS     float64 `json:"lodds,omitempty"`
	LODDS_NCL float64 `json:"loddsNcl,omitempty"`
	LODDS_NCU float64 `json:"loddsNcu,omitempty"`
	LODDS_BCL float64 `json:"loddsBcl,omitempty"`
	LODDS_BCU float64 `json:"loddsBcu,omitempty"`
	ORSS      float64 `json:"orss,omitempty"`
	ORSS_NCL  float64 `json:"orssNcl,omitempty"`
	ORSS_NCU  float64 `json:"orssNcu,omitempty"`
	ORSS_BCL  float64 `json:"orssBcl,omitempty"`
	ORSS_BCU  float64 `json:"orssBcu,omitempty"`
	EDS       float64 `json:"eds,omitempty"`
	EDS_NCL   float64 `json:"edsNcl,omitempty"`
	EDS_NCU   float64 `json:"edsNcu,omitempty"`
	EDS_BCL   float64 `json:"edsBcl,omitempty"`
	EDS_BCU   float64 `json:"edsBcu,omitempty"`
	SEDS      float64 `json:"seds,omitempty"`
	SEDS_NCL  float64 `json:"sedsNcl,omitempty"`
	SEDS_NCU  float64 `json:"sedsNcu,omitempty"`
	SEDS_BCL  float64 `json:"sedsBcl,omitempty"`
	SEDS_BCU  float64 `json:"sedsBcu,omitempty"`
	EDI       float64 `json:"edi,omitempty"`
	EDI_NCL   float64 `json:"ediNcl,omitempty"`
	EDI_NCU   float64 `json:"ediNcu,omitempty"`
	EDI_BCL   float64 `json:"ediBcl,omitempty"`
	EDI_BCU   float64 `json:"ediBcu,omitempty"`
	SEDI      float64 `json:"sedi,omitempty"`
	SEDI_NCL  float64 `json:"sediNcl,omitempty"`
	SEDI_NCU  float64 `json:"sediNcu,omitempty"`
	SEDI_BCL  float64 `json:"sediBcl,omitempty"`
	SEDI_BCU  float64 `json:"sediBcu,omitempty"`
	BAGSS     float64 `json:"bagss,omitempty"`
	BAGSS_BCL float64 `json:"bagssBcl,omitempty"`
	BAGSS_BCU float64 `json:"bagssBcu,omitempty"`
}

type STAT_DMAP struct {
	TOTAL      int     `json:"total,omitempty"`
	FY         int     `json:"fy,omitempty"`
	OY         int     `json:"oy,omitempty"`
	FBIAS      float64 `json:"fbias,omitempty"`
	BADDELEY   float64 `json:"baddeley,omitempty"`
	HAUSDORFF  float64 `json:"hausdorff,omitempty"`
	MED_FO     float64 `json:"medFo,omitempty"`
	MED_OF     float64 `json:"medOf,omitempty"`
	MED_MIN    float64 `json:"medMin,omitempty"`
	MED_MAX    float64 `json:"medMax,omitempty"`
	MED_MEAN   float64 `json:"medMean,omitempty"`
	FOM_FO     float64 `json:"fomFo,omitempty"`
	FOM_OF     float64 `json:"fomOf,omitempty"`
	FOM_MIN    float64 `json:"fomMin,omitempty"`
	FOM_MAX    float64 `json:"fomMax,omitempty"`
	FOM_MEAN   float64 `json:"fomMean,omitempty"`
	ZHU_FO     float64 `json:"zhuFo,omitempty"`
	ZHU_OF     float64 `json:"zhuOf,omitempty"`
	ZHU_MIN    float64 `json:"zhuMin,omitempty"`
	ZHU_MAX    float64 `json:"zhuMax,omitempty"`
	ZHU_MEAN   float64 `json:"zhuMean,omitempty"`
	G          float64 `json:"g,omitempty"`
	GBETA      float64 `json:"gbeta,omitempty"`
	BETA_VALUE float64 `json:"betaValue,omitempty"`
}

type STAT_ECLV struct {
	TOTAL       int                    `json:"total,omitempty"`
	BASER       float64                `json:"baser,omitempty"`
	VALUE_BASER int                    `json:"valueBaser,omitempty"`
	PTS         map[string]interface{} `json:"pts,omitempty"`
}

type STAT_ECNT struct {
	TOTAL            int     `json:"total,omitempty"`
	N_ENS            int     `json:"nEns,omitempty"`
	CRPS             float64 `json:"crps,omitempty"`
	CRPSS            float64 `json:"crpss,omitempty"`
	IGN              float64 `json:"ign,omitempty"`
	ME               float64 `json:"me,omitempty"`
	RMSE             float64 `json:"rmse,omitempty"`
	SPREAD           float64 `json:"spread,omitempty"`
	ME_OERR          float64 `json:"meOerr,omitempty"`
	RMSE_OERR        float64 `json:"rmseOerr,omitempty"`
	SPREAD_OERR      float64 `json:"spreadOerr,omitempty"`
	SPREAD_PLUS_OERR float64 `json:"spreadPlusOerr,omitempty"`
	CRPSCL           float64 `json:"crpscl,omitempty"`
	CRPS_EMP         float64 `json:"crpsEmp,omitempty"`
	CRPSCL_EMP       float64 `json:"crpsclEmp,omitempty"`
	CRPSS_EMP        float64 `json:"crpssEmp,omitempty"`
}

type STAT_FHO struct {
	TOTAL  int     `json:"total,omitempty"`
	F_RATE float64 `json:"fRate,omitempty"`
	H_RATE float64 `json:"hRate,omitempty"`
	O_RATE float64 `json:"oRate,omitempty"`
}

type STAT_GENMPR struct {
	TOTAL      int     `json:"total,omitempty"`
	INDEX      int     `json:"index,omitempty"`
	STORM_ID   string  `json:"stormId,omitempty"`
	PROB_LEAD  float64 `json:"probLead,omitempty"`
	PROB_VAL   float64 `json:"probVal,omitempty"`
	AGEN_INIT  string  `json:"agenInit,omitempty"`
	AGEN_FHR   string  `json:"agenFhr,omitempty"`
	AGEN_LAT   float64 `json:"agenLat,omitempty"`
	AGEN_LON   float64 `json:"agenLon,omitempty"`
	AGEN_DLAND float64 `json:"agenDland,omitempty"`
	BGEN_LAT   float64 `json:"bgenLat,omitempty"`
	BGEN_LON   float64 `json:"bgenLon,omitempty"`
	BGEN_DLAND float64 `json:"bgenDland,omitempty"`
	GEN_DIST   float64 `json:"genDist,omitempty"`
	GEN_TDIFF  string  `json:"genTdiff,omitempty"`
	INIT_TDIFF string  `json:"initTdiff,omitempty"`
	DEV_CAT    string  `json:"devCat,omitempty"`
	OPS_CAT    string  `json:"opsCat,omitempty"`
}

type STAT_GRAD struct {
	TOTAL      int     `json:"total,omitempty"`
	FGBAR      float64 `json:"fgbar,omitempty"`
	OGBAR      float64 `json:"ogbar,omitempty"`
	MGBAR      float64 `json:"mgbar,omitempty"`
	EGBAR      float64 `json:"egbar,omitempty"`
	S1         float64 `json:"s1,omitempty"`
	S1_OG      float64 `json:"s1Og,omitempty"`
	FGOG_RATIO float64 `json:"fgogRatio,omitempty"`
	DX         float64 `json:"dx,omitempty"`
	DY         float64 `json:"dy,omitempty"`
}

type STAT_ISC struct {
	TOTAL    int     `json:"total,omitempty"`
	TILE_DIM int     `json:"tileDim,omitempty"`
	TILE_XLL int     `json:"tileXll,omitempty"`
	TILE_YLL int     `json:"tileYll,omitempty"`
	NSCALE   int     `json:"nscale,omitempty"`
	ISCALE   int     `json:"iscale,omitempty"`
	MSE      float64 `json:"mse,omitempty"`
	ISC      float64 `json:"isc,omitempty"`
	FENERGY2 float64 `json:"fenergy2,omitempty"`
	OENERGY2 float64 `json:"oenergy2,omitempty"`
	BASER    float64 `json:"baser,omitempty"`
	FBIAS    float64 `json:"fbias,omitempty"`
}

type STAT_MCTC struct {
	TOTAL    int                    `json:"total,omitempty"`
	CAT      map[string]interface{} `json:"cat,omitempty"`
	EC_VALUE float64                `json:"ecValue,omitempty"`
}

type STAT_MCTS struct {
	TOTAL      int     `json:"total,omitempty"`
	N_CAT      int     `json:"nCat,omitempty"`
	ACC        float64 `json:"acc,omitempty"`
	ACC_NCL    float64 `json:"accNcl,omitempty"`
	ACC_NCU    float64 `json:"accNcu,omitempty"`
	ACC_BCL    float64 `json:"accBcl,omitempty"`
	ACC_BCU    float64 `json:"accBcu,omitempty"`
	HK         float64 `json:"hk,omitempty"`
	HK_BCL     float64 `json:"hkBcl,omitempty"`
	HK_BCU     float64 `json:"hkBcu,omitempty"`
	HSS        float64 `json:"hss,omitempty"`
	HSS_BCL    float64 `json:"hssBcl,omitempty"`
	HSS_BCU    float64 `json:"hssBcu,omitempty"`
	GER        float64 `json:"ger,omitempty"`
	GER_BCL    float64 `json:"gerBcl,omitempty"`
	GER_BCU    float64 `json:"gerBcu,omitempty"`
	HSS_EC     float64 `json:"hssEc,omitempty"`
	HSS_EC_BCL float64 `json:"hssEcBcl,omitempty"`
	HSS_EC_BCU float64 `json:"hssEcBcu,omitempty"`
	EC_VALUE   float64 `json:"ecValue,omitempty"`
}

type STAT_MPR struct {
	TOTAL       int     `json:"total,omitempty"`
	INDEX       int     `json:"index,omitempty"`
	OBS_SID     string  `json:"obsSid,omitempty"`
	OBS_LAT     float64 `json:"obsLat,omitempty"`
	OBS_LON     float64 `json:"obsLon,omitempty"`
	OBS_LVL     float64 `json:"obsLvl,omitempty"`
	OBS_ELV     float64 `json:"obsElv,omitempty"`
	FCST        float64 `json:"fcst,omitempty"`
	OBS         float64 `json:"obs,omitempty"`
	OBS_QC      string  `json:"obsQc,omitempty"`
	CLIMO_MEAN  float64 `json:"climoMean,omitempty"`
	CLIMO_STDEV float64 `json:"climoStdev,omitempty"`
	CLIMO_CDF   float64 `json:"climoCdf,omitempty"`
}

type STAT_NBRCNT struct {
	TOTAL      int     `json:"total,omitempty"`
	FBS        float64 `json:"fbs,omitempty"`
	FBS_BCL    float64 `json:"fbsBcl,omitempty"`
	FBS_BCU    float64 `json:"fbsBcu,omitempty"`
	FSS        float64 `json:"fss,omitempty"`
	FSS_BCL    float64 `json:"fssBcl,omitempty"`
	FSS_BCU    float64 `json:"fssBcu,omitempty"`
	AFSS       float64 `json:"afss,omitempty"`
	AFSS_BCL   float64 `json:"afssBcl,omitempty"`
	AFSS_BCU   float64 `json:"afssBcu,omitempty"`
	UFSS       float64 `json:"ufss,omitempty"`
	UFSS_BCL   float64 `json:"ufssBcl,omitempty"`
	UFSS_BCU   float64 `json:"ufssBcu,omitempty"`
	F_RATE     float64 `json:"fRate,omitempty"`
	F_RATE_BCL float64 `json:"fRateBcl,omitempty"`
	F_RATE_BCU float64 `json:"fRateBcu,omitempty"`
	O_RATE     float64 `json:"oRate,omitempty"`
	O_RATE_BCL float64 `json:"oRateBcl,omitempty"`
	O_RATE_BCU float64 `json:"oRateBcu,omitempty"`
}

type STAT_NBRCTC struct {
	TOTAL int     `json:"total,omitempty"`
	FY_OY float64 `json:"fyOy,omitempty"`
	FY_ON float64 `json:"fyOn,omitempty"`
	FN_OY float64 `json:"fnOy,omitempty"`
	FN_ON float64 `json:"fnOn,omitempty"`
}

type STAT_NBRCTS struct {
	TOTAL     int     `json:"total,omitempty"`
	BASER     float64 `json:"baser,omitempty"`
	BASER_NCL float64 `json:"baserNcl,omitempty"`
	BASER_NCU float64 `json:"baserNcu,omitempty"`
	BASER_BCL float64 `json:"baserBcl,omitempty"`
	BASER_BCU float64 `json:"baserBcu,omitempty"`
	FMEAN     float64 `json:"fmean,omitempty"`
	FMEAN_NCL float64 `json:"fmeanNcl,omitempty"`
	FMEAN_NCU float64 `json:"fmeanNcu,omitempty"`
	FMEAN_BCL float64 `json:"fmeanBcl,omitempty"`
	FMEAN_BCU float64 `json:"fmeanBcu,omitempty"`
	ACC       float64 `json:"acc,omitempty"`
	ACC_NCL   float64 `json:"accNcl,omitempty"`
	ACC_NCU   float64 `json:"accNcu,omitempty"`
	ACC_BCL   float64 `json:"accBcl,omitempty"`
	ACC_BCU   float64 `json:"accBcu,omitempty"`
	FBIAS     float64 `json:"fbias,omitempty"`
	FBIAS_BCL float64 `json:"fbiasBcl,omitempty"`
	FBIAS_BCU float64 `json:"fbiasBcu,omitempty"`
	PODY      float64 `json:"pody,omitempty"`
	PODY_NCL  float64 `json:"podyNcl,omitempty"`
	PODY_NCU  float64 `json:"podyNcu,omitempty"`
	PODY_BCL  float64 `json:"podyBcl,omitempty"`
	PODY_BCU  float64 `json:"podyBcu,omitempty"`
	PODN      float64 `json:"podn,omitempty"`
	PODN_NCL  float64 `json:"podnNcl,omitempty"`
	PODN_NCU  float64 `json:"podnNcu,omitempty"`
	PODN_BCL  float64 `json:"podnBcl,omitempty"`
	PODN_BCU  float64 `json:"podnBcu,omitempty"`
	POFD      float64 `json:"pofd,omitempty"`
	POFD_NCL  float64 `json:"pofdNcl,omitempty"`
	POFD_NCU  float64 `json:"pofdNcu,omitempty"`
	POFD_BCL  float64 `json:"pofdBcl,omitempty"`
	POFD_BCU  float64 `json:"pofdBcu,omitempty"`
	FAR       float64 `json:"far,omitempty"`
	FAR_NCL   float64 `json:"farNcl,omitempty"`
	FAR_NCU   float64 `json:"farNcu,omitempty"`
	FAR_BCL   float64 `json:"farBcl,omitempty"`
	FAR_BCU   float64 `json:"farBcu,omitempty"`
	CSI       float64 `json:"csi,omitempty"`
	CSI_NCL   float64 `json:"csiNcl,omitempty"`
	CSI_NCU   float64 `json:"csiNcu,omitempty"`
	CSI_BCL   float64 `json:"csiBcl,omitempty"`
	CSI_BCU   float64 `json:"csiBcu,omitempty"`
	GSS       float64 `json:"gss,omitempty"`
	GSS_BCL   float64 `json:"gssBcl,omitempty"`
	GSS_BCU   float64 `json:"gssBcu,omitempty"`
	HK        float64 `json:"hk,omitempty"`
	HK_NCL    float64 `json:"hkNcl,omitempty"`
	HK_NCU    float64 `json:"hkNcu,omitempty"`
	HK_BCL    float64 `json:"hkBcl,omitempty"`
	HK_BCU    float64 `json:"hkBcu,omitempty"`
	HSS       float64 `json:"hss,omitempty"`
	HSS_BCL   float64 `json:"hssBcl,omitempty"`
	HSS_BCU   float64 `json:"hssBcu,omitempty"`
	ODDS      float64 `json:"odds,omitempty"`
	ODDS_NCL  float64 `json:"oddsNcl,omitempty"`
	ODDS_NCU  float64 `json:"oddsNcu,omitempty"`
	ODDS_BCL  float64 `json:"oddsBcl,omitempty"`
	ODDS_BCU  float64 `json:"oddsBcu,omitempty"`
	LODDS     float64 `json:"lodds,omitempty"`
	LODDS_NCL float64 `json:"loddsNcl,omitempty"`
	LODDS_NCU float64 `json:"loddsNcu,omitempty"`
	LODDS_BCL float64 `json:"loddsBcl,omitempty"`
	LODDS_BCU float64 `json:"loddsBcu,omitempty"`
	ORSS      float64 `json:"orss,omitempty"`
	ORSS_NCL  float64 `json:"orssNcl,omitempty"`
	ORSS_NCU  float64 `json:"orssNcu,omitempty"`
	ORSS_BCL  float64 `json:"orssBcl,omitempty"`
	ORSS_BCU  float64 `json:"orssBcu,omitempty"`
	EDS       float64 `json:"eds,omitempty"`
	EDS_NCL   float64 `json:"edsNcl,omitempty"`
	EDS_NCU   float64 `json:"edsNcu,omitempty"`
	EDS_BCL   float64 `json:"edsBcl,omitempty"`
	EDS_BCU   float64 `json:"edsBcu,omitempty"`
	SEDS      float64 `json:"seds,omitempty"`
	SEDS_NCL  float64 `json:"sedsNcl,omitempty"`
	SEDS_NCU  float64 `json:"sedsNcu,omitempty"`
	SEDS_BCL  float64 `json:"sedsBcl,omitempty"`
	SEDS_BCU  float64 `json:"sedsBcu,omitempty"`
	EDI       float64 `json:"edi,omitempty"`
	EDI_NCL   float64 `json:"ediNcl,omitempty"`
	EDI_NCU   float64 `json:"ediNcu,omitempty"`
	EDI_BCL   float64 `json:"ediBcl,omitempty"`
	EDI_BCU   float64 `json:"ediBcu,omitempty"`
	SEDI      float64 `json:"sedi,omitempty"`
	SEDI_NCL  float64 `json:"sediNcl,omitempty"`
	SEDI_NCU  float64 `json:"sediNcu,omitempty"`
	SEDI_BCL  float64 `json:"sediBcl,omitempty"`
	SEDI_BCU  float64 `json:"sediBcu,omitempty"`
	BAGSS     float64 `json:"bagss,omitempty"`
	BAGSS_BCL float64 `json:"bagssBcl,omitempty"`
	BAGSS_BCU float64 `json:"bagssBcu,omitempty"`
}

type STAT_ORANK struct {
	TOTAL            int                    `json:"total,omitempty"`
	INDEX            int                    `json:"index,omitempty"`
	OBS_SID          string                 `json:"obsSid,omitempty"`
	OBS_LAT          float64                `json:"obsLat,omitempty"`
	OBS_LON          float64                `json:"obsLon,omitempty"`
	OBS_LVL          float64                `json:"obsLvl,omitempty"`
	OBS_ELV          float64                `json:"obsElv,omitempty"`
	OBS              float64                `json:"obs,omitempty"`
	PIT              float64                `json:"pit,omitempty"`
	RANK             int                    `json:"rank,omitempty"`
	N_ENS_VLD        int                    `json:"nEnsVld,omitempty"`
	ENS              map[string]interface{} `json:"ens,omitempty"`
	OBS_QC           string                 `json:"obsQc,omitempty"`
	ENS_MEAN         int                    `json:"ensMean,omitempty"`
	CLIMO_MEAN       float64                `json:"climoMean,omitempty"`
	SPREAD           float64                `json:"spread,omitempty"`
	ENS_MEAN_OERR    int                    `json:"ensMeanOerr,omitempty"`
	SPREAD_OERR      float64                `json:"spreadOerr,omitempty"`
	SPREAD_PLUS_OERR float64                `json:"spreadPlusOerr,omitempty"`
	CLIMO_STDEV      float64                `json:"climoStdev,omitempty"`
}

type STAT_PCT struct {
	TOTAL  int                    `json:"total,omitempty"`
	THRESH map[string]interface{} `json:"thresh,omitempty"`
}

type STAT_PHIST struct {
	TOTAL    int                    `json:"total,omitempty"`
	BIN_SIZE int                    `json:"binSize,omitempty"`
	BIN      map[string]interface{} `json:"bin,omitempty"`
}

type STAT_PJC struct {
	TOTAL  int                    `json:"total,omitempty"`
	THRESH map[string]interface{} `json:"thresh,omitempty"`
}

type STAT_PRC struct {
	TOTAL  int                    `json:"total,omitempty"`
	THRESH map[string]interface{} `json:"thresh,omitempty"`
}

type STAT_PSTD struct {
	TOTAL       int                    `json:"total,omitempty"`
	THRESH      map[string]interface{} `json:"thresh,omitempty"`
	BASER_NCL   float64                `json:"baserNcl,omitempty"`
	BASER_NCU   float64                `json:"baserNcu,omitempty"`
	RELIABILITY float64                `json:"reliability,omitempty"`
	RESOLUTION  float64                `json:"resolution,omitempty"`
	UNCERTAINTY float64                `json:"uncertainty,omitempty"`
	ROC_AUC     float64                `json:"rocAuc,omitempty"`
	BRIER       float64                `json:"brier,omitempty"`
	BRIER_NCL   float64                `json:"brierNcl,omitempty"`
	BRIER_NCU   float64                `json:"brierNcu,omitempty"`
	BRIERCL     float64                `json:"briercl,omitempty"`
	BRIERCL_NCL float64                `json:"brierclNcl,omitempty"`
	BRIERCL_NCU float64                `json:"brierclNcu,omitempty"`
	BSS         float64                `json:"bss,omitempty"`
	BSS_SMPL    float64                `json:"bssSmpl,omitempty"`
	THRESH_I    int                    `json:"threshI,omitempty"`
}

type STAT_RELP struct {
	TOTAL int                    `json:"total,omitempty"`
	ENS   map[string]interface{} `json:"ens,omitempty"`
}

type STAT_RHIST struct {
	TOTAL int                    `json:"total,omitempty"`
	RANK  map[string]interface{} `json:"rank,omitempty"`
}

type STAT_RPS struct {
	TOTAL     int     `json:"total,omitempty"`
	N_PROB    int     `json:"nProb,omitempty"`
	RPS_REL   float64 `json:"rpsRel,omitempty"`
	RPS_RES   float64 `json:"rpsRes,omitempty"`
	RPS_UNC   float64 `json:"rpsUnc,omitempty"`
	RPS       float64 `json:"rps,omitempty"`
	RPSS      float64 `json:"rpss,omitempty"`
	RPSS_SMPL float64 `json:"rpssSmpl,omitempty"`
	RPS_COMP  float64 `json:"rpsComp,omitempty"`
}

type STAT_SAL1L2 struct {
	TOTAL  int     `json:"total,omitempty"`
	FABAR  float64 `json:"fabar,omitempty"`
	OABAR  float64 `json:"oabar,omitempty"`
	FOABAR float64 `json:"foabar,omitempty"`
	FFABAR float64 `json:"ffabar,omitempty"`
	OOABAR float64 `json:"ooabar,omitempty"`
	MAE    float64 `json:"mae,omitempty"`
}

type STAT_SL1L2 struct {
	TOTAL int     `json:"total,omitempty"`
	FBAR  float64 `json:"fbar,omitempty"`
	OBAR  float64 `json:"obar,omitempty"`
	FOBAR float64 `json:"fobar,omitempty"`
	FFBAR float64 `json:"ffbar,omitempty"`
	OOBAR float64 `json:"oobar,omitempty"`
	MAE   float64 `json:"mae,omitempty"`
}

type STAT_SSIDX struct {
	FCST_MODEL string  `json:"fcstModel,omitempty"`
	REF_MODEL  string  `json:"refModel,omitempty"`
	N_INIT     int     `json:"nInit,omitempty"`
	N_TERM     int     `json:"nTerm,omitempty"`
	N_VLD      int     `json:"nVld,omitempty"`
	SS_INDEX   float64 `json:"ssIndex,omitempty"`
}

type STAT_SSVAR struct {
	TOTAL       int     `json:"total,omitempty"`
	N_BIN       int     `json:"nBin,omitempty"`
	BIN_I       int     `json:"binI,omitempty"`
	BIN_N       int     `json:"binN,omitempty"`
	VAR_MIN     float64 `json:"varMin,omitempty"`
	VAR_MAX     float64 `json:"varMax,omitempty"`
	VAR_MEAN    float64 `json:"varMean,omitempty"`
	FBAR        float64 `json:"fbar,omitempty"`
	OBAR        float64 `json:"obar,omitempty"`
	FOBAR       float64 `json:"fobar,omitempty"`
	FFBAR       float64 `json:"ffbar,omitempty"`
	OOBAR       float64 `json:"oobar,omitempty"`
	FBAR_NCL    float64 `json:"fbarNcl,omitempty"`
	FBAR_NCU    float64 `json:"fbarNcu,omitempty"`
	FSTDEV      float64 `json:"fstdev,omitempty"`
	FSTDEV_NCL  float64 `json:"fstdevNcl,omitempty"`
	FSTDEV_NCU  float64 `json:"fstdevNcu,omitempty"`
	OBAR_NCL    float64 `json:"obarNcl,omitempty"`
	OBAR_NCU    float64 `json:"obarNcu,omitempty"`
	OSTDEV      float64 `json:"ostdev,omitempty"`
	OSTDEV_NCL  float64 `json:"ostdevNcl,omitempty"`
	OSTDEV_NCU  float64 `json:"ostdevNcu,omitempty"`
	PR_CORR     float64 `json:"prCorr,omitempty"`
	PR_CORR_NCL float64 `json:"prCorrNcl,omitempty"`
	PR_CORR_NCU float64 `json:"prCorrNcu,omitempty"`
	ME          float64 `json:"me,omitempty"`
	ME_NCL      float64 `json:"meNcl,omitempty"`
	ME_NCU      float64 `json:"meNcu,omitempty"`
	ESTDEV      float64 `json:"estdev,omitempty"`
	ESTDEV_NCL  float64 `json:"estdevNcl,omitempty"`
	ESTDEV_NCU  float64 `json:"estdevNcu,omitempty"`
	MBIAS       float64 `json:"mbias,omitempty"`
	MSE         float64 `json:"mse,omitempty"`
	BCMSE       float64 `json:"bcmse,omitempty"`
	RMSE        float64 `json:"rmse,omitempty"`
}

type STAT_VAL1L2 struct {
	TOTAL    int     `json:"total,omitempty"`
	UFABAR   float64 `json:"ufabar,omitempty"`
	VFABAR   float64 `json:"vfabar,omitempty"`
	UOABAR   float64 `json:"uoabar,omitempty"`
	VOABAR   float64 `json:"voabar,omitempty"`
	UVFOABAR float64 `json:"uvfoabar,omitempty"`
	UVFFABAR float64 `json:"uvffabar,omitempty"`
	UVOOABAR float64 `json:"uvooabar,omitempty"`
}

type STAT_VCNT struct {
	TOTAL            int     `json:"total,omitempty"`
	FBAR             float64 `json:"fbar,omitempty"`
	FBAR_BCL         float64 `json:"fbarBcl,omitempty"`
	FBAR_BCU         float64 `json:"fbarBcu,omitempty"`
	OBAR             float64 `json:"obar,omitempty"`
	OBAR_BCL         float64 `json:"obarBcl,omitempty"`
	OBAR_BCU         float64 `json:"obarBcu,omitempty"`
	FS_RMS           float64 `json:"fsRms,omitempty"`
	FS_RMS_BCL       float64 `json:"fsRmsBcl,omitempty"`
	FS_RMS_BCU       float64 `json:"fsRmsBcu,omitempty"`
	OS_RMS           float64 `json:"osRms,omitempty"`
	OS_RMS_BCL       float64 `json:"osRmsBcl,omitempty"`
	OS_RMS_BCU       float64 `json:"osRmsBcu,omitempty"`
	MSVE             float64 `json:"msve,omitempty"`
	MSVE_BCL         float64 `json:"msveBcl,omitempty"`
	MSVE_BCU         float64 `json:"msveBcu,omitempty"`
	RMSVE            float64 `json:"rmsve,omitempty"`
	RMSVE_BCL        float64 `json:"rmsveBcl,omitempty"`
	RMSVE_BCU        float64 `json:"rmsveBcu,omitempty"`
	FSTDEV           float64 `json:"fstdev,omitempty"`
	FSTDEV_BCL       float64 `json:"fstdevBcl,omitempty"`
	FSTDEV_BCU       float64 `json:"fstdevBcu,omitempty"`
	OSTDEV           float64 `json:"ostdev,omitempty"`
	OSTDEV_BCL       float64 `json:"ostdevBcl,omitempty"`
	OSTDEV_BCU       float64 `json:"ostdevBcu,omitempty"`
	FDIR             float64 `json:"fdir,omitempty"`
	FDIR_BCL         float64 `json:"fdirBcl,omitempty"`
	FDIR_BCU         float64 `json:"fdirBcu,omitempty"`
	ODIR             float64 `json:"odir,omitempty"`
	ODIR_BCL         float64 `json:"odirBcl,omitempty"`
	ODIR_BCU         float64 `json:"odirBcu,omitempty"`
	FBAR_SPEED       float64 `json:"fbarSpeed,omitempty"`
	FBAR_SPEED_BCL   float64 `json:"fbarSpeedBcl,omitempty"`
	FBAR_SPEED_BCU   float64 `json:"fbarSpeedBcu,omitempty"`
	OBAR_SPEED       float64 `json:"obarSpeed,omitempty"`
	OBAR_SPEED_BCL   float64 `json:"obarSpeedBcl,omitempty"`
	OBAR_SPEED_BCU   float64 `json:"obarSpeedBcu,omitempty"`
	VDIFF_SPEED      float64 `json:"vdiffSpeed,omitempty"`
	VDIFF_SPEED_BCL  float64 `json:"vdiffSpeedBcl,omitempty"`
	VDIFF_SPEED_BCU  float64 `json:"vdiffSpeedBcu,omitempty"`
	VDIFF_DIR        float64 `json:"vdiffDir,omitempty"`
	VDIFF_DIR_BCL    float64 `json:"vdiffDirBcl,omitempty"`
	VDIFF_DIR_BCU    float64 `json:"vdiffDirBcu,omitempty"`
	SPEED_ERR        float64 `json:"speedErr,omitempty"`
	SPEED_ERR_BCL    float64 `json:"speedErrBcl,omitempty"`
	SPEED_ERR_BCU    float64 `json:"speedErrBcu,omitempty"`
	SPEED_ABSERR     float64 `json:"speedAbserr,omitempty"`
	SPEED_ABSERR_BCL float64 `json:"speedAbserrBcl,omitempty"`
	SPEED_ABSERR_BCU float64 `json:"speedAbserrBcu,omitempty"`
	DIR_ERR          float64 `json:"dirErr,omitempty"`
	DIR_ERR_BCL      float64 `json:"dirErrBcl,omitempty"`
	DIR_ERR_BCU      float64 `json:"dirErrBcu,omitempty"`
	DIR_ABSERR       float64 `json:"dirAbserr,omitempty"`
	DIR_ABSERR_BCL   float64 `json:"dirAbserrBcl,omitempty"`
	DIR_ABSERR_BCU   float64 `json:"dirAbserrBcu,omitempty"`
}

type STAT_VL1L2 struct {
	TOTAL       int     `json:"total,omitempty"`
	UFBAR       float64 `json:"ufbar,omitempty"`
	VFBAR       float64 `json:"vfbar,omitempty"`
	UOBAR       float64 `json:"uobar,omitempty"`
	VOBAR       float64 `json:"vobar,omitempty"`
	UVFOBAR     float64 `json:"uvfobar,omitempty"`
	UVFFBAR     float64 `json:"uvffbar,omitempty"`
	UVOOBAR     float64 `json:"uvoobar,omitempty"`
	F_SPEED_BAR float64 `json:"fSpeedBar,omitempty"`
	O_SPEED_BAR float64 `json:"oSpeedBar,omitempty"`
}

type TCST_PROBRIRW struct {
	ALAT        float64                `json:"alat,omitempty"`
	ALON        float64                `json:"alon,omitempty"`
	BLAT        float64                `json:"blat,omitempty"`
	BLON        float64                `json:"blon,omitempty"`
	INITIALS    string                 `json:"initials,omitempty"`
	TK_ERR      float64                `json:"tkErr,omitempty"`
	X_ERR       float64                `json:"xErr,omitempty"`
	Y_ERR       float64                `json:"yErr,omitempty"`
	ADLAND      float64                `json:"adland,omitempty"`
	BDLAND      float64                `json:"bdland,omitempty"`
	RIRW_BEG    int                    `json:"rirwBeg,omitempty"`
	RIRW_END    int                    `json:"rirwEnd,omitempty"`
	RIRW_WINDOW int                    `json:"rirwWindow,omitempty"`
	AWIND_END   float64                `json:"awindEnd,omitempty"`
	BWIND_BEG   float64                `json:"bwindBeg,omitempty"`
	BWIND_END   float64                `json:"bwindEnd,omitempty"`
	BDELTA      float64                `json:"bdelta,omitempty"`
	BDELTA_MAX  float64                `json:"bdeltaMax,omitempty"`
	BLEVEL_BEG  string                 `json:"blevelBeg,omitempty"`
	BLEVEL_END  string                 `json:"blevelEnd,omitempty"`
	THRESH      map[string]interface{} `json:"thresh,omitempty"`
	INIT        int                    `json:"init,omitempty"`
}

type TCST_TCMPR struct {
	TOTAL       int     `json:"total,omitempty"`
	INDEX       int     `json:"index,omitempty"`
	LEVEL       string  `json:"level,omitempty"`
	WATCH_WARN  string  `json:"watchWarn,omitempty"`
	INITIALS    string  `json:"initials,omitempty"`
	ALAT        float64 `json:"alat,omitempty"`
	ALON        float64 `json:"alon,omitempty"`
	BLAT        float64 `json:"blat,omitempty"`
	BLON        float64 `json:"blon,omitempty"`
	TK_ERR      float64 `json:"tkErr,omitempty"`
	X_ERR       float64 `json:"xErr,omitempty"`
	Y_ERR       float64 `json:"yErr,omitempty"`
	ALTK_ERR    float64 `json:"altkErr,omitempty"`
	CRTK_ERR    float64 `json:"crtkErr,omitempty"`
	ADLAND      float64 `json:"adland,omitempty"`
	BDLAND      float64 `json:"bdland,omitempty"`
	AMSLP       float64 `json:"amslp,omitempty"`
	BMSLP       float64 `json:"bmslp,omitempty"`
	AMAX_WIND   float64 `json:"amaxWind,omitempty"`
	BMAX_WIND   float64 `json:"bmaxWind,omitempty"`
	AAL_WIND_34 float64 `json:"aalWind34,omitempty"`
	BAL_WIND_34 float64 `json:"balWind34,omitempty"`
	ANE_WIND_34 float64 `json:"aneWind34,omitempty"`
	BNE_WIND_34 float64 `json:"bneWind34,omitempty"`
	ASE_WIND_34 float64 `json:"aseWind34,omitempty"`
	BSE_WIND_34 float64 `json:"bseWind34,omitempty"`
	ASW_WIND_34 float64 `json:"aswWind34,omitempty"`
	BSW_WIND_34 float64 `json:"bswWind34,omitempty"`
	ANW_WIND_34 float64 `json:"anwWind34,omitempty"`
	BNW_WIND_34 float64 `json:"bnwWind34,omitempty"`
	AAL_WIND_50 float64 `json:"aalWind50,omitempty"`
	BAL_WIND_50 float64 `json:"balWind50,omitempty"`
	ANE_WIND_50 float64 `json:"aneWind50,omitempty"`
	BNE_WIND_50 float64 `json:"bneWind50,omitempty"`
	ASE_WIND_50 float64 `json:"aseWind50,omitempty"`
	BSE_WIND_50 float64 `json:"bseWind50,omitempty"`
	ASW_WIND_50 float64 `json:"aswWind50,omitempty"`
	BSW_WIND_50 float64 `json:"bswWind50,omitempty"`
	ANW_WIND_50 float64 `json:"anwWind50,omitempty"`
	BNW_WIND_50 float64 `json:"bnwWind50,omitempty"`
	AAL_WIND_64 float64 `json:"aalWind64,omitempty"`
	BAL_WIND_64 float64 `json:"balWind64,omitempty"`
	ANE_WIND_64 float64 `json:"aneWind64,omitempty"`
	BNE_WIND_64 float64 `json:"bneWind64,omitempty"`
	ASE_WIND_64 float64 `json:"aseWind64,omitempty"`
	BSE_WIND_64 float64 `json:"bseWind64,omitempty"`
	ASW_WIND_64 float64 `json:"aswWind64,omitempty"`
	BSW_WIND_64 float64 `json:"bswWind64,omitempty"`
	ANW_WIND_64 float64 `json:"anwWind64,omitempty"`
	BNW_WIND_64 float64 `json:"bnwWind64,omitempty"`
	ARADP       string  `json:"aradp,omitempty"`
	BRADP       float64 `json:"bradp,omitempty"`
	ARRP        int     `json:"arrp,omitempty"`
	BRRP        float64 `json:"brrp,omitempty"`
	AMRD        int     `json:"amrd,omitempty"`
	BMRD        float64 `json:"bmrd,omitempty"`
	AGUSTS      int     `json:"agusts,omitempty"`
	BGUSTS      float64 `json:"bgusts,omitempty"`
	AEYE        int     `json:"aeye,omitempty"`
	BEYE        float64 `json:"beye,omitempty"`
	ADIR        int     `json:"adir,omitempty"`
	BDIR        float64 `json:"bdir,omitempty"`
	ASPEED      int     `json:"aspeed,omitempty"`
	BSPEED      float64 `json:"bspeed,omitempty"`
	ADEPTH      int     `json:"adepth,omitempty"`
	BDEPTH      float64 `json:"bdepth,omitempty"`
	INIT        int     `json:"init,omitempty"`
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
		s.INIT, _ = strconv.Atoi(fields[66])
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

var MetHeaderColumnsFileUrl = "https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/data/table_files/met_header_columns_V10.1.txt"
