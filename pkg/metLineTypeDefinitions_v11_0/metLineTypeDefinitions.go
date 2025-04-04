/*undefined: [S12 S32 S13 S31 S21 S23]*/
package metLineTypeDefinitions_v11_0

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

/*
THIS CODE IS AUTOMATICALLY GENERATED - DO NOT EDIT THIS CODE
To modify this code - modify the buildHeaderLineTypes.go file and run the buildHeaderLineTypes.go program
cd  <repo path>/metlinetypes/pkg/buildHeaderLineTypes
go run . > /tmp/types.go
cp /tmp/types.go ../metLineTypeDefinitions/metLineTypeDefinitions.go
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
type STAT_MPR_header struct {
	VERSION        string  `json:"version"`
	MODEL          string  `json:"model"`
	DESC           string  `json:"desc"`
	FCST_VALID_BEG int     `json:"fcst_valid_beg"`
	FCST_VALID_END int     `json:"fcst_valid_end"`
	OBS_LEAD       int     `json:"obs_lead"`
	OBS_VALID_BEG  int     `json:"obs_valid_beg"`
	OBS_VALID_END  int     `json:"obs_valid_end"`
	FCST_VAR       string  `json:"fcst_var"`
	FCST_UNITS     string  `json:"fcst_units"`
	FCST_LEV       string  `json:"fcst_lev"`
	OBS_VAR        string  `json:"obs_var"`
	OBS_UNITS      string  `json:"obs_units"`
	OBS_LEV        string  `json:"obs_lev"`
	OBTYPE         string  `json:"obtype"`
	VX_MASK        string  `json:"vx_mask"`
	INTERP_MTHD    string  `json:"interp_mthd"`
	INTERP_PNTS    int     `json:"interp_pnts"`
	FCST_THRESH    string  `json:"fcst_thresh"`
	OBS_THRESH     string  `json:"obs_thresh"`
	COV_THRESH     string  `json:"cov_thresh"`
	ALPHA          float64 `json:"alpha"`
	LINE_TYPE      string  `json:"line_type"`
}

type STAT_SAL1L2_header struct {
	VERSION        string  `json:"version"`
	MODEL          string  `json:"model"`
	DESC           string  `json:"desc"`
	FCST_VALID_BEG int     `json:"fcst_valid_beg"`
	FCST_VALID_END int     `json:"fcst_valid_end"`
	OBS_LEAD       int     `json:"obs_lead"`
	OBS_VALID_BEG  int     `json:"obs_valid_beg"`
	OBS_VALID_END  int     `json:"obs_valid_end"`
	FCST_VAR       string  `json:"fcst_var"`
	FCST_UNITS     string  `json:"fcst_units"`
	FCST_LEV       string  `json:"fcst_lev"`
	OBS_VAR        string  `json:"obs_var"`
	OBS_UNITS      string  `json:"obs_units"`
	OBS_LEV        string  `json:"obs_lev"`
	OBTYPE         string  `json:"obtype"`
	VX_MASK        string  `json:"vx_mask"`
	INTERP_MTHD    string  `json:"interp_mthd"`
	INTERP_PNTS    int     `json:"interp_pnts"`
	FCST_THRESH    string  `json:"fcst_thresh"`
	OBS_THRESH     string  `json:"obs_thresh"`
	COV_THRESH     string  `json:"cov_thresh"`
	ALPHA          float64 `json:"alpha"`
	LINE_TYPE      string  `json:"line_type"`
}

type STAT_ECNT_header struct {
	VERSION        string  `json:"version"`
	MODEL          string  `json:"model"`
	DESC           string  `json:"desc"`
	FCST_VALID_BEG int     `json:"fcst_valid_beg"`
	FCST_VALID_END int     `json:"fcst_valid_end"`
	OBS_LEAD       int     `json:"obs_lead"`
	OBS_VALID_BEG  int     `json:"obs_valid_beg"`
	OBS_VALID_END  int     `json:"obs_valid_end"`
	FCST_VAR       string  `json:"fcst_var"`
	FCST_UNITS     string  `json:"fcst_units"`
	FCST_LEV       string  `json:"fcst_lev"`
	OBS_VAR        string  `json:"obs_var"`
	OBS_UNITS      string  `json:"obs_units"`
	OBS_LEV        string  `json:"obs_lev"`
	OBTYPE         string  `json:"obtype"`
	VX_MASK        string  `json:"vx_mask"`
	INTERP_MTHD    string  `json:"interp_mthd"`
	INTERP_PNTS    int     `json:"interp_pnts"`
	FCST_THRESH    string  `json:"fcst_thresh"`
	OBS_THRESH     string  `json:"obs_thresh"`
	COV_THRESH     string  `json:"cov_thresh"`
	ALPHA          float64 `json:"alpha"`
	LINE_TYPE      string  `json:"line_type"`
}

type STAT_PHIST_header struct {
	VERSION        string  `json:"version"`
	MODEL          string  `json:"model"`
	DESC           string  `json:"desc"`
	FCST_VALID_BEG int     `json:"fcst_valid_beg"`
	FCST_VALID_END int     `json:"fcst_valid_end"`
	OBS_LEAD       int     `json:"obs_lead"`
	OBS_VALID_BEG  int     `json:"obs_valid_beg"`
	OBS_VALID_END  int     `json:"obs_valid_end"`
	FCST_VAR       string  `json:"fcst_var"`
	FCST_UNITS     string  `json:"fcst_units"`
	FCST_LEV       string  `json:"fcst_lev"`
	OBS_VAR        string  `json:"obs_var"`
	OBS_UNITS      string  `json:"obs_units"`
	OBS_LEV        string  `json:"obs_lev"`
	OBTYPE         string  `json:"obtype"`
	VX_MASK        string  `json:"vx_mask"`
	INTERP_MTHD    string  `json:"interp_mthd"`
	INTERP_PNTS    int     `json:"interp_pnts"`
	FCST_THRESH    string  `json:"fcst_thresh"`
	OBS_THRESH     string  `json:"obs_thresh"`
	COV_THRESH     string  `json:"cov_thresh"`
	ALPHA          float64 `json:"alpha"`
	LINE_TYPE      string  `json:"line_type"`
}

type STAT_VCNT_header struct {
	VERSION        string  `json:"version"`
	MODEL          string  `json:"model"`
	DESC           string  `json:"desc"`
	FCST_VALID_BEG int     `json:"fcst_valid_beg"`
	FCST_VALID_END int     `json:"fcst_valid_end"`
	OBS_LEAD       int     `json:"obs_lead"`
	OBS_VALID_BEG  int     `json:"obs_valid_beg"`
	OBS_VALID_END  int     `json:"obs_valid_end"`
	FCST_VAR       string  `json:"fcst_var"`
	FCST_UNITS     string  `json:"fcst_units"`
	FCST_LEV       string  `json:"fcst_lev"`
	OBS_VAR        string  `json:"obs_var"`
	OBS_UNITS      string  `json:"obs_units"`
	OBS_LEV        string  `json:"obs_lev"`
	OBTYPE         string  `json:"obtype"`
	VX_MASK        string  `json:"vx_mask"`
	INTERP_MTHD    string  `json:"interp_mthd"`
	INTERP_PNTS    int     `json:"interp_pnts"`
	FCST_THRESH    string  `json:"fcst_thresh"`
	OBS_THRESH     string  `json:"obs_thresh"`
	COV_THRESH     string  `json:"cov_thresh"`
	ALPHA          float64 `json:"alpha"`
	LINE_TYPE      string  `json:"line_type"`
}

type MODE_CTS_header struct {
	VERSION    string  `json:"version"`
	MODEL      string  `json:"model"`
	N_VALID    int     `json:"n_valid"`
	GRID_RES   float64 `json:"grid_res"`
	DESC       string  `json:"desc"`
	FCST_VALID string  `json:"fcst_valid"`
	FCST_ACCUM string  `json:"fcst_accum"`
	OBS_LEAD   int     `json:"obs_lead"`
	OBS_VALID  string  `json:"obs_valid"`
	OBS_ACCUM  string  `json:"obs_accum"`
	FCST_RAD   int     `json:"fcst_rad"`
	FCST_THR   string  `json:"fcst_thr"`
	OBS_RAD    int     `json:"obs_rad"`
	OBS_THR    string  `json:"obs_thr"`
	FCST_VAR   string  `json:"fcst_var"`
	FCST_UNITS string  `json:"fcst_units"`
	FCST_LEV   string  `json:"fcst_lev"`
	OBS_VAR    string  `json:"obs_var"`
	OBS_UNITS  string  `json:"obs_units"`
	OBS_LEV    string  `json:"obs_lev"`
	OBTYPE     string  `json:"obtype"`
	LINE_TYPE  string  `json:"line_type"`
}

type TCST_PROBRIRW_header struct {
	VERSION    string `json:"version"`
	AMODEL     string `json:"amodel"`
	BMODEL     string `json:"bmodel"`
	DESC       string `json:"desc"`
	STORM_ID   string `json:"storm_id"`
	BASIN      string `json:"basin"`
	CYCLONE    string `json:"cyclone"`
	STORM_NAME string `json:"storm_name"`
	VALID      int    `json:"valid"`
	INIT_MASK  string `json:"init_mask"`
	VALID_MASK string `json:"valid_mask"`
	LINE_TYPE  string `json:"line_type"`
}

type TCST_TCMPR_header struct {
	VERSION    string `json:"version"`
	AMODEL     string `json:"amodel"`
	BMODEL     string `json:"bmodel"`
	DESC       string `json:"desc"`
	STORM_ID   string `json:"storm_id"`
	BASIN      string `json:"basin"`
	CYCLONE    string `json:"cyclone"`
	STORM_NAME string `json:"storm_name"`
	VALID      int    `json:"valid"`
	INIT_MASK  string `json:"init_mask"`
	VALID_MASK string `json:"valid_mask"`
	LINE_TYPE  string `json:"line_type"`
}

type STAT_ISC_header struct {
	VERSION        string  `json:"version"`
	MODEL          string  `json:"model"`
	DESC           string  `json:"desc"`
	FCST_VALID_BEG int     `json:"fcst_valid_beg"`
	FCST_VALID_END int     `json:"fcst_valid_end"`
	OBS_LEAD       int     `json:"obs_lead"`
	OBS_VALID_BEG  int     `json:"obs_valid_beg"`
	OBS_VALID_END  int     `json:"obs_valid_end"`
	FCST_VAR       string  `json:"fcst_var"`
	FCST_UNITS     string  `json:"fcst_units"`
	FCST_LEV       string  `json:"fcst_lev"`
	OBS_VAR        string  `json:"obs_var"`
	OBS_UNITS      string  `json:"obs_units"`
	OBS_LEV        string  `json:"obs_lev"`
	OBTYPE         string  `json:"obtype"`
	VX_MASK        string  `json:"vx_mask"`
	INTERP_MTHD    string  `json:"interp_mthd"`
	INTERP_PNTS    int     `json:"interp_pnts"`
	FCST_THRESH    string  `json:"fcst_thresh"`
	OBS_THRESH     string  `json:"obs_thresh"`
	COV_THRESH     string  `json:"cov_thresh"`
	ALPHA          float64 `json:"alpha"`
	LINE_TYPE      string  `json:"line_type"`
}

type STAT_NBRCNT_header struct {
	VERSION        string  `json:"version"`
	MODEL          string  `json:"model"`
	DESC           string  `json:"desc"`
	FCST_VALID_BEG int     `json:"fcst_valid_beg"`
	FCST_VALID_END int     `json:"fcst_valid_end"`
	OBS_LEAD       int     `json:"obs_lead"`
	OBS_VALID_BEG  int     `json:"obs_valid_beg"`
	OBS_VALID_END  int     `json:"obs_valid_end"`
	FCST_VAR       string  `json:"fcst_var"`
	FCST_UNITS     string  `json:"fcst_units"`
	FCST_LEV       string  `json:"fcst_lev"`
	OBS_VAR        string  `json:"obs_var"`
	OBS_UNITS      string  `json:"obs_units"`
	OBS_LEV        string  `json:"obs_lev"`
	OBTYPE         string  `json:"obtype"`
	VX_MASK        string  `json:"vx_mask"`
	INTERP_MTHD    string  `json:"interp_mthd"`
	INTERP_PNTS    int     `json:"interp_pnts"`
	FCST_THRESH    string  `json:"fcst_thresh"`
	OBS_THRESH     string  `json:"obs_thresh"`
	COV_THRESH     string  `json:"cov_thresh"`
	ALPHA          float64 `json:"alpha"`
	LINE_TYPE      string  `json:"line_type"`
}

type STAT_PJC_header struct {
	VERSION        string  `json:"version"`
	MODEL          string  `json:"model"`
	DESC           string  `json:"desc"`
	FCST_VALID_BEG int     `json:"fcst_valid_beg"`
	FCST_VALID_END int     `json:"fcst_valid_end"`
	OBS_LEAD       int     `json:"obs_lead"`
	OBS_VALID_BEG  int     `json:"obs_valid_beg"`
	OBS_VALID_END  int     `json:"obs_valid_end"`
	FCST_VAR       string  `json:"fcst_var"`
	FCST_UNITS     string  `json:"fcst_units"`
	FCST_LEV       string  `json:"fcst_lev"`
	OBS_VAR        string  `json:"obs_var"`
	OBS_UNITS      string  `json:"obs_units"`
	OBS_LEV        string  `json:"obs_lev"`
	OBTYPE         string  `json:"obtype"`
	VX_MASK        string  `json:"vx_mask"`
	INTERP_MTHD    string  `json:"interp_mthd"`
	INTERP_PNTS    int     `json:"interp_pnts"`
	FCST_THRESH    string  `json:"fcst_thresh"`
	OBS_THRESH     string  `json:"obs_thresh"`
	COV_THRESH     string  `json:"cov_thresh"`
	ALPHA          float64 `json:"alpha"`
	LINE_TYPE      string  `json:"line_type"`
}

type STAT_SL1L2_header struct {
	VERSION        string  `json:"version"`
	MODEL          string  `json:"model"`
	DESC           string  `json:"desc"`
	FCST_VALID_BEG int     `json:"fcst_valid_beg"`
	FCST_VALID_END int     `json:"fcst_valid_end"`
	OBS_LEAD       int     `json:"obs_lead"`
	OBS_VALID_BEG  int     `json:"obs_valid_beg"`
	OBS_VALID_END  int     `json:"obs_valid_end"`
	FCST_VAR       string  `json:"fcst_var"`
	FCST_UNITS     string  `json:"fcst_units"`
	FCST_LEV       string  `json:"fcst_lev"`
	OBS_VAR        string  `json:"obs_var"`
	OBS_UNITS      string  `json:"obs_units"`
	OBS_LEV        string  `json:"obs_lev"`
	OBTYPE         string  `json:"obtype"`
	VX_MASK        string  `json:"vx_mask"`
	INTERP_MTHD    string  `json:"interp_mthd"`
	INTERP_PNTS    int     `json:"interp_pnts"`
	FCST_THRESH    string  `json:"fcst_thresh"`
	OBS_THRESH     string  `json:"obs_thresh"`
	COV_THRESH     string  `json:"cov_thresh"`
	ALPHA          float64 `json:"alpha"`
	LINE_TYPE      string  `json:"line_type"`
}

type STAT_SSIDX_header struct {
	VERSION        string  `json:"version"`
	MODEL          string  `json:"model"`
	DESC           string  `json:"desc"`
	FCST_VALID_BEG int     `json:"fcst_valid_beg"`
	FCST_VALID_END int     `json:"fcst_valid_end"`
	OBS_LEAD       int     `json:"obs_lead"`
	OBS_VALID_BEG  int     `json:"obs_valid_beg"`
	OBS_VALID_END  int     `json:"obs_valid_end"`
	FCST_VAR       string  `json:"fcst_var"`
	FCST_UNITS     string  `json:"fcst_units"`
	FCST_LEV       string  `json:"fcst_lev"`
	OBS_VAR        string  `json:"obs_var"`
	OBS_UNITS      string  `json:"obs_units"`
	OBS_LEV        string  `json:"obs_lev"`
	OBTYPE         string  `json:"obtype"`
	VX_MASK        string  `json:"vx_mask"`
	INTERP_MTHD    string  `json:"interp_mthd"`
	INTERP_PNTS    int     `json:"interp_pnts"`
	FCST_THRESH    string  `json:"fcst_thresh"`
	OBS_THRESH     string  `json:"obs_thresh"`
	COV_THRESH     string  `json:"cov_thresh"`
	ALPHA          float64 `json:"alpha"`
	LINE_TYPE      string  `json:"line_type"`
}

type STAT_DMAP_header struct {
	VERSION        string  `json:"version"`
	MODEL          string  `json:"model"`
	DESC           string  `json:"desc"`
	FCST_VALID_BEG int     `json:"fcst_valid_beg"`
	FCST_VALID_END int     `json:"fcst_valid_end"`
	OBS_LEAD       int     `json:"obs_lead"`
	OBS_VALID_BEG  int     `json:"obs_valid_beg"`
	OBS_VALID_END  int     `json:"obs_valid_end"`
	FCST_VAR       string  `json:"fcst_var"`
	FCST_UNITS     string  `json:"fcst_units"`
	FCST_LEV       string  `json:"fcst_lev"`
	OBS_VAR        string  `json:"obs_var"`
	OBS_UNITS      string  `json:"obs_units"`
	OBS_LEV        string  `json:"obs_lev"`
	OBTYPE         string  `json:"obtype"`
	VX_MASK        string  `json:"vx_mask"`
	INTERP_MTHD    string  `json:"interp_mthd"`
	INTERP_PNTS    int     `json:"interp_pnts"`
	FCST_THRESH    string  `json:"fcst_thresh"`
	OBS_THRESH     string  `json:"obs_thresh"`
	COV_THRESH     string  `json:"cov_thresh"`
	ALPHA          float64 `json:"alpha"`
	LINE_TYPE      string  `json:"line_type"`
}

type STAT_ECLV_header struct {
	VERSION        string  `json:"version"`
	MODEL          string  `json:"model"`
	DESC           string  `json:"desc"`
	FCST_VALID_BEG int     `json:"fcst_valid_beg"`
	FCST_VALID_END int     `json:"fcst_valid_end"`
	OBS_LEAD       int     `json:"obs_lead"`
	OBS_VALID_BEG  int     `json:"obs_valid_beg"`
	OBS_VALID_END  int     `json:"obs_valid_end"`
	FCST_VAR       string  `json:"fcst_var"`
	FCST_UNITS     string  `json:"fcst_units"`
	FCST_LEV       string  `json:"fcst_lev"`
	OBS_VAR        string  `json:"obs_var"`
	OBS_UNITS      string  `json:"obs_units"`
	OBS_LEV        string  `json:"obs_lev"`
	OBTYPE         string  `json:"obtype"`
	VX_MASK        string  `json:"vx_mask"`
	INTERP_MTHD    string  `json:"interp_mthd"`
	INTERP_PNTS    int     `json:"interp_pnts"`
	FCST_THRESH    string  `json:"fcst_thresh"`
	OBS_THRESH     string  `json:"obs_thresh"`
	COV_THRESH     string  `json:"cov_thresh"`
	ALPHA          float64 `json:"alpha"`
	LINE_TYPE      string  `json:"line_type"`
}

type STAT_RHIST_header struct {
	VERSION        string  `json:"version"`
	MODEL          string  `json:"model"`
	DESC           string  `json:"desc"`
	FCST_VALID_BEG int     `json:"fcst_valid_beg"`
	FCST_VALID_END int     `json:"fcst_valid_end"`
	OBS_LEAD       int     `json:"obs_lead"`
	OBS_VALID_BEG  int     `json:"obs_valid_beg"`
	OBS_VALID_END  int     `json:"obs_valid_end"`
	FCST_VAR       string  `json:"fcst_var"`
	FCST_UNITS     string  `json:"fcst_units"`
	FCST_LEV       string  `json:"fcst_lev"`
	OBS_VAR        string  `json:"obs_var"`
	OBS_UNITS      string  `json:"obs_units"`
	OBS_LEV        string  `json:"obs_lev"`
	OBTYPE         string  `json:"obtype"`
	VX_MASK        string  `json:"vx_mask"`
	INTERP_MTHD    string  `json:"interp_mthd"`
	INTERP_PNTS    int     `json:"interp_pnts"`
	FCST_THRESH    string  `json:"fcst_thresh"`
	OBS_THRESH     string  `json:"obs_thresh"`
	COV_THRESH     string  `json:"cov_thresh"`
	ALPHA          float64 `json:"alpha"`
	LINE_TYPE      string  `json:"line_type"`
}

type STAT_SSVAR_header struct {
	VERSION        string  `json:"version"`
	MODEL          string  `json:"model"`
	DESC           string  `json:"desc"`
	FCST_VALID_BEG int     `json:"fcst_valid_beg"`
	FCST_VALID_END int     `json:"fcst_valid_end"`
	OBS_LEAD       int     `json:"obs_lead"`
	OBS_VALID_BEG  int     `json:"obs_valid_beg"`
	OBS_VALID_END  int     `json:"obs_valid_end"`
	FCST_VAR       string  `json:"fcst_var"`
	FCST_UNITS     string  `json:"fcst_units"`
	FCST_LEV       string  `json:"fcst_lev"`
	OBS_VAR        string  `json:"obs_var"`
	OBS_UNITS      string  `json:"obs_units"`
	OBS_LEV        string  `json:"obs_lev"`
	OBTYPE         string  `json:"obtype"`
	VX_MASK        string  `json:"vx_mask"`
	INTERP_MTHD    string  `json:"interp_mthd"`
	INTERP_PNTS    int     `json:"interp_pnts"`
	FCST_THRESH    string  `json:"fcst_thresh"`
	OBS_THRESH     string  `json:"obs_thresh"`
	COV_THRESH     string  `json:"cov_thresh"`
	ALPHA          float64 `json:"alpha"`
	LINE_TYPE      string  `json:"line_type"`
}

type STAT_VAL1L2_header struct {
	VERSION        string  `json:"version"`
	MODEL          string  `json:"model"`
	DESC           string  `json:"desc"`
	FCST_VALID_BEG int     `json:"fcst_valid_beg"`
	FCST_VALID_END int     `json:"fcst_valid_end"`
	OBS_LEAD       int     `json:"obs_lead"`
	OBS_VALID_BEG  int     `json:"obs_valid_beg"`
	OBS_VALID_END  int     `json:"obs_valid_end"`
	FCST_VAR       string  `json:"fcst_var"`
	FCST_UNITS     string  `json:"fcst_units"`
	FCST_LEV       string  `json:"fcst_lev"`
	OBS_VAR        string  `json:"obs_var"`
	OBS_UNITS      string  `json:"obs_units"`
	OBS_LEV        string  `json:"obs_lev"`
	OBTYPE         string  `json:"obtype"`
	VX_MASK        string  `json:"vx_mask"`
	INTERP_MTHD    string  `json:"interp_mthd"`
	INTERP_PNTS    int     `json:"interp_pnts"`
	FCST_THRESH    string  `json:"fcst_thresh"`
	OBS_THRESH     string  `json:"obs_thresh"`
	COV_THRESH     string  `json:"cov_thresh"`
	ALPHA          float64 `json:"alpha"`
	LINE_TYPE      string  `json:"line_type"`
}

type STAT_CNT_header struct {
	VERSION        string  `json:"version"`
	MODEL          string  `json:"model"`
	DESC           string  `json:"desc"`
	FCST_VALID_BEG int     `json:"fcst_valid_beg"`
	FCST_VALID_END int     `json:"fcst_valid_end"`
	OBS_LEAD       int     `json:"obs_lead"`
	OBS_VALID_BEG  int     `json:"obs_valid_beg"`
	OBS_VALID_END  int     `json:"obs_valid_end"`
	FCST_VAR       string  `json:"fcst_var"`
	FCST_UNITS     string  `json:"fcst_units"`
	FCST_LEV       string  `json:"fcst_lev"`
	OBS_VAR        string  `json:"obs_var"`
	OBS_UNITS      string  `json:"obs_units"`
	OBS_LEV        string  `json:"obs_lev"`
	OBTYPE         string  `json:"obtype"`
	VX_MASK        string  `json:"vx_mask"`
	INTERP_MTHD    string  `json:"interp_mthd"`
	INTERP_PNTS    int     `json:"interp_pnts"`
	FCST_THRESH    string  `json:"fcst_thresh"`
	OBS_THRESH     string  `json:"obs_thresh"`
	COV_THRESH     string  `json:"cov_thresh"`
	ALPHA          float64 `json:"alpha"`
	LINE_TYPE      string  `json:"line_type"`
}

type STAT_MCTS_header struct {
	VERSION        string  `json:"version"`
	MODEL          string  `json:"model"`
	DESC           string  `json:"desc"`
	FCST_VALID_BEG int     `json:"fcst_valid_beg"`
	FCST_VALID_END int     `json:"fcst_valid_end"`
	OBS_LEAD       int     `json:"obs_lead"`
	OBS_VALID_BEG  int     `json:"obs_valid_beg"`
	OBS_VALID_END  int     `json:"obs_valid_end"`
	FCST_VAR       string  `json:"fcst_var"`
	FCST_UNITS     string  `json:"fcst_units"`
	FCST_LEV       string  `json:"fcst_lev"`
	OBS_VAR        string  `json:"obs_var"`
	OBS_UNITS      string  `json:"obs_units"`
	OBS_LEV        string  `json:"obs_lev"`
	OBTYPE         string  `json:"obtype"`
	VX_MASK        string  `json:"vx_mask"`
	INTERP_MTHD    string  `json:"interp_mthd"`
	INTERP_PNTS    int     `json:"interp_pnts"`
	FCST_THRESH    string  `json:"fcst_thresh"`
	OBS_THRESH     string  `json:"obs_thresh"`
	COV_THRESH     string  `json:"cov_thresh"`
	ALPHA          float64 `json:"alpha"`
	LINE_TYPE      string  `json:"line_type"`
}

type STAT_SEEPS_header struct {
	VERSION        string  `json:"version"`
	MODEL          string  `json:"model"`
	DESC           string  `json:"desc"`
	FCST_VALID_BEG int     `json:"fcst_valid_beg"`
	FCST_VALID_END int     `json:"fcst_valid_end"`
	OBS_LEAD       int     `json:"obs_lead"`
	OBS_VALID_BEG  int     `json:"obs_valid_beg"`
	OBS_VALID_END  int     `json:"obs_valid_end"`
	FCST_VAR       string  `json:"fcst_var"`
	FCST_UNITS     string  `json:"fcst_units"`
	FCST_LEV       string  `json:"fcst_lev"`
	OBS_VAR        string  `json:"obs_var"`
	OBS_UNITS      string  `json:"obs_units"`
	OBS_LEV        string  `json:"obs_lev"`
	OBTYPE         string  `json:"obtype"`
	VX_MASK        string  `json:"vx_mask"`
	INTERP_MTHD    string  `json:"interp_mthd"`
	INTERP_PNTS    int     `json:"interp_pnts"`
	FCST_THRESH    string  `json:"fcst_thresh"`
	OBS_THRESH     string  `json:"obs_thresh"`
	COV_THRESH     string  `json:"cov_thresh"`
	ALPHA          float64 `json:"alpha"`
	LINE_TYPE      string  `json:"line_type"`
}

type STAT_GRAD_header struct {
	VERSION        string  `json:"version"`
	MODEL          string  `json:"model"`
	DESC           string  `json:"desc"`
	FCST_VALID_BEG int     `json:"fcst_valid_beg"`
	FCST_VALID_END int     `json:"fcst_valid_end"`
	OBS_LEAD       int     `json:"obs_lead"`
	OBS_VALID_BEG  int     `json:"obs_valid_beg"`
	OBS_VALID_END  int     `json:"obs_valid_end"`
	FCST_VAR       string  `json:"fcst_var"`
	FCST_UNITS     string  `json:"fcst_units"`
	FCST_LEV       string  `json:"fcst_lev"`
	OBS_VAR        string  `json:"obs_var"`
	OBS_UNITS      string  `json:"obs_units"`
	OBS_LEV        string  `json:"obs_lev"`
	OBTYPE         string  `json:"obtype"`
	VX_MASK        string  `json:"vx_mask"`
	INTERP_MTHD    string  `json:"interp_mthd"`
	INTERP_PNTS    int     `json:"interp_pnts"`
	FCST_THRESH    string  `json:"fcst_thresh"`
	OBS_THRESH     string  `json:"obs_thresh"`
	COV_THRESH     string  `json:"cov_thresh"`
	ALPHA          float64 `json:"alpha"`
	LINE_TYPE      string  `json:"line_type"`
}

type STAT_PCT_header struct {
	VERSION        string  `json:"version"`
	MODEL          string  `json:"model"`
	DESC           string  `json:"desc"`
	FCST_VALID_BEG int     `json:"fcst_valid_beg"`
	FCST_VALID_END int     `json:"fcst_valid_end"`
	OBS_LEAD       int     `json:"obs_lead"`
	OBS_VALID_BEG  int     `json:"obs_valid_beg"`
	OBS_VALID_END  int     `json:"obs_valid_end"`
	FCST_VAR       string  `json:"fcst_var"`
	FCST_UNITS     string  `json:"fcst_units"`
	FCST_LEV       string  `json:"fcst_lev"`
	OBS_VAR        string  `json:"obs_var"`
	OBS_UNITS      string  `json:"obs_units"`
	OBS_LEV        string  `json:"obs_lev"`
	OBTYPE         string  `json:"obtype"`
	VX_MASK        string  `json:"vx_mask"`
	INTERP_MTHD    string  `json:"interp_mthd"`
	INTERP_PNTS    int     `json:"interp_pnts"`
	FCST_THRESH    string  `json:"fcst_thresh"`
	OBS_THRESH     string  `json:"obs_thresh"`
	COV_THRESH     string  `json:"cov_thresh"`
	ALPHA          float64 `json:"alpha"`
	LINE_TYPE      string  `json:"line_type"`
}

type STAT_RPS_header struct {
	VERSION        string  `json:"version"`
	MODEL          string  `json:"model"`
	DESC           string  `json:"desc"`
	FCST_VALID_BEG int     `json:"fcst_valid_beg"`
	FCST_VALID_END int     `json:"fcst_valid_end"`
	OBS_LEAD       int     `json:"obs_lead"`
	OBS_VALID_BEG  int     `json:"obs_valid_beg"`
	OBS_VALID_END  int     `json:"obs_valid_end"`
	FCST_VAR       string  `json:"fcst_var"`
	FCST_UNITS     string  `json:"fcst_units"`
	FCST_LEV       string  `json:"fcst_lev"`
	OBS_VAR        string  `json:"obs_var"`
	OBS_UNITS      string  `json:"obs_units"`
	OBS_LEV        string  `json:"obs_lev"`
	OBTYPE         string  `json:"obtype"`
	VX_MASK        string  `json:"vx_mask"`
	INTERP_MTHD    string  `json:"interp_mthd"`
	INTERP_PNTS    int     `json:"interp_pnts"`
	FCST_THRESH    string  `json:"fcst_thresh"`
	OBS_THRESH     string  `json:"obs_thresh"`
	COV_THRESH     string  `json:"cov_thresh"`
	ALPHA          float64 `json:"alpha"`
	LINE_TYPE      string  `json:"line_type"`
}

type STAT_RELP_header struct {
	VERSION        string  `json:"version"`
	MODEL          string  `json:"model"`
	DESC           string  `json:"desc"`
	FCST_VALID_BEG int     `json:"fcst_valid_beg"`
	FCST_VALID_END int     `json:"fcst_valid_end"`
	OBS_LEAD       int     `json:"obs_lead"`
	OBS_VALID_BEG  int     `json:"obs_valid_beg"`
	OBS_VALID_END  int     `json:"obs_valid_end"`
	FCST_VAR       string  `json:"fcst_var"`
	FCST_UNITS     string  `json:"fcst_units"`
	FCST_LEV       string  `json:"fcst_lev"`
	OBS_VAR        string  `json:"obs_var"`
	OBS_UNITS      string  `json:"obs_units"`
	OBS_LEV        string  `json:"obs_lev"`
	OBTYPE         string  `json:"obtype"`
	VX_MASK        string  `json:"vx_mask"`
	INTERP_MTHD    string  `json:"interp_mthd"`
	INTERP_PNTS    int     `json:"interp_pnts"`
	FCST_THRESH    string  `json:"fcst_thresh"`
	OBS_THRESH     string  `json:"obs_thresh"`
	COV_THRESH     string  `json:"cov_thresh"`
	ALPHA          float64 `json:"alpha"`
	LINE_TYPE      string  `json:"line_type"`
}

type STAT_VL1L2_header struct {
	VERSION        string  `json:"version"`
	MODEL          string  `json:"model"`
	DESC           string  `json:"desc"`
	FCST_VALID_BEG int     `json:"fcst_valid_beg"`
	FCST_VALID_END int     `json:"fcst_valid_end"`
	OBS_LEAD       int     `json:"obs_lead"`
	OBS_VALID_BEG  int     `json:"obs_valid_beg"`
	OBS_VALID_END  int     `json:"obs_valid_end"`
	FCST_VAR       string  `json:"fcst_var"`
	FCST_UNITS     string  `json:"fcst_units"`
	FCST_LEV       string  `json:"fcst_lev"`
	OBS_VAR        string  `json:"obs_var"`
	OBS_UNITS      string  `json:"obs_units"`
	OBS_LEV        string  `json:"obs_lev"`
	OBTYPE         string  `json:"obtype"`
	VX_MASK        string  `json:"vx_mask"`
	INTERP_MTHD    string  `json:"interp_mthd"`
	INTERP_PNTS    int     `json:"interp_pnts"`
	FCST_THRESH    string  `json:"fcst_thresh"`
	OBS_THRESH     string  `json:"obs_thresh"`
	COV_THRESH     string  `json:"cov_thresh"`
	ALPHA          float64 `json:"alpha"`
	LINE_TYPE      string  `json:"line_type"`
}

type STAT_MCTC_header struct {
	VERSION        string  `json:"version"`
	MODEL          string  `json:"model"`
	DESC           string  `json:"desc"`
	FCST_VALID_BEG int     `json:"fcst_valid_beg"`
	FCST_VALID_END int     `json:"fcst_valid_end"`
	OBS_LEAD       int     `json:"obs_lead"`
	OBS_VALID_BEG  int     `json:"obs_valid_beg"`
	OBS_VALID_END  int     `json:"obs_valid_end"`
	FCST_VAR       string  `json:"fcst_var"`
	FCST_UNITS     string  `json:"fcst_units"`
	FCST_LEV       string  `json:"fcst_lev"`
	OBS_VAR        string  `json:"obs_var"`
	OBS_UNITS      string  `json:"obs_units"`
	OBS_LEV        string  `json:"obs_lev"`
	OBTYPE         string  `json:"obtype"`
	VX_MASK        string  `json:"vx_mask"`
	INTERP_MTHD    string  `json:"interp_mthd"`
	INTERP_PNTS    int     `json:"interp_pnts"`
	FCST_THRESH    string  `json:"fcst_thresh"`
	OBS_THRESH     string  `json:"obs_thresh"`
	COV_THRESH     string  `json:"cov_thresh"`
	ALPHA          float64 `json:"alpha"`
	LINE_TYPE      string  `json:"line_type"`
}

type STAT_NBRCTS_header struct {
	VERSION        string  `json:"version"`
	MODEL          string  `json:"model"`
	DESC           string  `json:"desc"`
	FCST_VALID_BEG int     `json:"fcst_valid_beg"`
	FCST_VALID_END int     `json:"fcst_valid_end"`
	OBS_LEAD       int     `json:"obs_lead"`
	OBS_VALID_BEG  int     `json:"obs_valid_beg"`
	OBS_VALID_END  int     `json:"obs_valid_end"`
	FCST_VAR       string  `json:"fcst_var"`
	FCST_UNITS     string  `json:"fcst_units"`
	FCST_LEV       string  `json:"fcst_lev"`
	OBS_VAR        string  `json:"obs_var"`
	OBS_UNITS      string  `json:"obs_units"`
	OBS_LEV        string  `json:"obs_lev"`
	OBTYPE         string  `json:"obtype"`
	VX_MASK        string  `json:"vx_mask"`
	INTERP_MTHD    string  `json:"interp_mthd"`
	INTERP_PNTS    int     `json:"interp_pnts"`
	FCST_THRESH    string  `json:"fcst_thresh"`
	OBS_THRESH     string  `json:"obs_thresh"`
	COV_THRESH     string  `json:"cov_thresh"`
	ALPHA          float64 `json:"alpha"`
	LINE_TYPE      string  `json:"line_type"`
}

type STAT_ORANK_header struct {
	VERSION        string  `json:"version"`
	MODEL          string  `json:"model"`
	DESC           string  `json:"desc"`
	FCST_VALID_BEG int     `json:"fcst_valid_beg"`
	FCST_VALID_END int     `json:"fcst_valid_end"`
	OBS_LEAD       int     `json:"obs_lead"`
	OBS_VALID_BEG  int     `json:"obs_valid_beg"`
	OBS_VALID_END  int     `json:"obs_valid_end"`
	FCST_VAR       string  `json:"fcst_var"`
	FCST_UNITS     string  `json:"fcst_units"`
	FCST_LEV       string  `json:"fcst_lev"`
	OBS_VAR        string  `json:"obs_var"`
	OBS_UNITS      string  `json:"obs_units"`
	OBS_LEV        string  `json:"obs_lev"`
	OBTYPE         string  `json:"obtype"`
	VX_MASK        string  `json:"vx_mask"`
	INTERP_MTHD    string  `json:"interp_mthd"`
	INTERP_PNTS    int     `json:"interp_pnts"`
	FCST_THRESH    string  `json:"fcst_thresh"`
	OBS_THRESH     string  `json:"obs_thresh"`
	COV_THRESH     string  `json:"cov_thresh"`
	ALPHA          float64 `json:"alpha"`
	LINE_TYPE      string  `json:"line_type"`
}

type STAT_PRC_header struct {
	VERSION        string  `json:"version"`
	MODEL          string  `json:"model"`
	DESC           string  `json:"desc"`
	FCST_VALID_BEG int     `json:"fcst_valid_beg"`
	FCST_VALID_END int     `json:"fcst_valid_end"`
	OBS_LEAD       int     `json:"obs_lead"`
	OBS_VALID_BEG  int     `json:"obs_valid_beg"`
	OBS_VALID_END  int     `json:"obs_valid_end"`
	FCST_VAR       string  `json:"fcst_var"`
	FCST_UNITS     string  `json:"fcst_units"`
	FCST_LEV       string  `json:"fcst_lev"`
	OBS_VAR        string  `json:"obs_var"`
	OBS_UNITS      string  `json:"obs_units"`
	OBS_LEV        string  `json:"obs_lev"`
	OBTYPE         string  `json:"obtype"`
	VX_MASK        string  `json:"vx_mask"`
	INTERP_MTHD    string  `json:"interp_mthd"`
	INTERP_PNTS    int     `json:"interp_pnts"`
	FCST_THRESH    string  `json:"fcst_thresh"`
	OBS_THRESH     string  `json:"obs_thresh"`
	COV_THRESH     string  `json:"cov_thresh"`
	ALPHA          float64 `json:"alpha"`
	LINE_TYPE      string  `json:"line_type"`
}

type STAT_PSTD_header struct {
	VERSION        string  `json:"version"`
	MODEL          string  `json:"model"`
	DESC           string  `json:"desc"`
	FCST_VALID_BEG int     `json:"fcst_valid_beg"`
	FCST_VALID_END int     `json:"fcst_valid_end"`
	OBS_LEAD       int     `json:"obs_lead"`
	OBS_VALID_BEG  int     `json:"obs_valid_beg"`
	OBS_VALID_END  int     `json:"obs_valid_end"`
	FCST_VAR       string  `json:"fcst_var"`
	FCST_UNITS     string  `json:"fcst_units"`
	FCST_LEV       string  `json:"fcst_lev"`
	OBS_VAR        string  `json:"obs_var"`
	OBS_UNITS      string  `json:"obs_units"`
	OBS_LEV        string  `json:"obs_lev"`
	OBTYPE         string  `json:"obtype"`
	VX_MASK        string  `json:"vx_mask"`
	INTERP_MTHD    string  `json:"interp_mthd"`
	INTERP_PNTS    int     `json:"interp_pnts"`
	FCST_THRESH    string  `json:"fcst_thresh"`
	OBS_THRESH     string  `json:"obs_thresh"`
	COV_THRESH     string  `json:"cov_thresh"`
	ALPHA          float64 `json:"alpha"`
	LINE_TYPE      string  `json:"line_type"`
}

type TCST_TCDIAG_header struct {
	VERSION    string `json:"version"`
	AMODEL     string `json:"amodel"`
	BMODEL     string `json:"bmodel"`
	DESC       string `json:"desc"`
	STORM_ID   string `json:"storm_id"`
	BASIN      string `json:"basin"`
	CYCLONE    string `json:"cyclone"`
	STORM_NAME string `json:"storm_name"`
	VALID      int    `json:"valid"`
	INIT_MASK  string `json:"init_mask"`
	VALID_MASK string `json:"valid_mask"`
	LINE_TYPE  string `json:"line_type"`
}

type STAT_CTS_header struct {
	VERSION        string  `json:"version"`
	MODEL          string  `json:"model"`
	DESC           string  `json:"desc"`
	FCST_VALID_BEG int     `json:"fcst_valid_beg"`
	FCST_VALID_END int     `json:"fcst_valid_end"`
	OBS_LEAD       int     `json:"obs_lead"`
	OBS_VALID_BEG  int     `json:"obs_valid_beg"`
	OBS_VALID_END  int     `json:"obs_valid_end"`
	FCST_VAR       string  `json:"fcst_var"`
	FCST_UNITS     string  `json:"fcst_units"`
	FCST_LEV       string  `json:"fcst_lev"`
	OBS_VAR        string  `json:"obs_var"`
	OBS_UNITS      string  `json:"obs_units"`
	OBS_LEV        string  `json:"obs_lev"`
	OBTYPE         string  `json:"obtype"`
	VX_MASK        string  `json:"vx_mask"`
	INTERP_MTHD    string  `json:"interp_mthd"`
	INTERP_PNTS    int     `json:"interp_pnts"`
	FCST_THRESH    string  `json:"fcst_thresh"`
	OBS_THRESH     string  `json:"obs_thresh"`
	COV_THRESH     string  `json:"cov_thresh"`
	ALPHA          float64 `json:"alpha"`
	LINE_TYPE      string  `json:"line_type"`
}

type STAT_FHO_header struct {
	VERSION        string  `json:"version"`
	MODEL          string  `json:"model"`
	DESC           string  `json:"desc"`
	FCST_VALID_BEG int     `json:"fcst_valid_beg"`
	FCST_VALID_END int     `json:"fcst_valid_end"`
	OBS_LEAD       int     `json:"obs_lead"`
	OBS_VALID_BEG  int     `json:"obs_valid_beg"`
	OBS_VALID_END  int     `json:"obs_valid_end"`
	FCST_VAR       string  `json:"fcst_var"`
	FCST_UNITS     string  `json:"fcst_units"`
	FCST_LEV       string  `json:"fcst_lev"`
	OBS_VAR        string  `json:"obs_var"`
	OBS_UNITS      string  `json:"obs_units"`
	OBS_LEV        string  `json:"obs_lev"`
	OBTYPE         string  `json:"obtype"`
	VX_MASK        string  `json:"vx_mask"`
	INTERP_MTHD    string  `json:"interp_mthd"`
	INTERP_PNTS    int     `json:"interp_pnts"`
	FCST_THRESH    string  `json:"fcst_thresh"`
	OBS_THRESH     string  `json:"obs_thresh"`
	COV_THRESH     string  `json:"cov_thresh"`
	ALPHA          float64 `json:"alpha"`
	LINE_TYPE      string  `json:"line_type"`
}

type MODE_OBJ_header struct {
	VERSION    string  `json:"version"`
	MODEL      string  `json:"model"`
	N_VALID    int     `json:"n_valid"`
	GRID_RES   float64 `json:"grid_res"`
	DESC       string  `json:"desc"`
	FCST_VALID string  `json:"fcst_valid"`
	FCST_ACCUM string  `json:"fcst_accum"`
	OBS_LEAD   int     `json:"obs_lead"`
	OBS_VALID  string  `json:"obs_valid"`
	OBS_ACCUM  string  `json:"obs_accum"`
	FCST_RAD   int     `json:"fcst_rad"`
	FCST_THR   string  `json:"fcst_thr"`
	OBS_RAD    int     `json:"obs_rad"`
	OBS_THR    string  `json:"obs_thr"`
	FCST_VAR   string  `json:"fcst_var"`
	FCST_UNITS string  `json:"fcst_units"`
	FCST_LEV   string  `json:"fcst_lev"`
	OBS_VAR    string  `json:"obs_var"`
	OBS_UNITS  string  `json:"obs_units"`
	OBS_LEV    string  `json:"obs_lev"`
	OBTYPE     string  `json:"obtype"`
	LINE_TYPE  string  `json:"line_type"`
}

type STAT_CTC_header struct {
	VERSION        string  `json:"version"`
	MODEL          string  `json:"model"`
	DESC           string  `json:"desc"`
	FCST_VALID_BEG int     `json:"fcst_valid_beg"`
	FCST_VALID_END int     `json:"fcst_valid_end"`
	OBS_LEAD       int     `json:"obs_lead"`
	OBS_VALID_BEG  int     `json:"obs_valid_beg"`
	OBS_VALID_END  int     `json:"obs_valid_end"`
	FCST_VAR       string  `json:"fcst_var"`
	FCST_UNITS     string  `json:"fcst_units"`
	FCST_LEV       string  `json:"fcst_lev"`
	OBS_VAR        string  `json:"obs_var"`
	OBS_UNITS      string  `json:"obs_units"`
	OBS_LEV        string  `json:"obs_lev"`
	OBTYPE         string  `json:"obtype"`
	VX_MASK        string  `json:"vx_mask"`
	INTERP_MTHD    string  `json:"interp_mthd"`
	INTERP_PNTS    int     `json:"interp_pnts"`
	FCST_THRESH    string  `json:"fcst_thresh"`
	OBS_THRESH     string  `json:"obs_thresh"`
	COV_THRESH     string  `json:"cov_thresh"`
	ALPHA          float64 `json:"alpha"`
	LINE_TYPE      string  `json:"line_type"`
}

type STAT_SEEPS_MPR_header struct {
	VERSION        string  `json:"version"`
	MODEL          string  `json:"model"`
	DESC           string  `json:"desc"`
	FCST_VALID_BEG int     `json:"fcst_valid_beg"`
	FCST_VALID_END int     `json:"fcst_valid_end"`
	OBS_LEAD       int     `json:"obs_lead"`
	OBS_VALID_BEG  int     `json:"obs_valid_beg"`
	OBS_VALID_END  int     `json:"obs_valid_end"`
	FCST_VAR       string  `json:"fcst_var"`
	FCST_UNITS     string  `json:"fcst_units"`
	FCST_LEV       string  `json:"fcst_lev"`
	OBS_VAR        string  `json:"obs_var"`
	OBS_UNITS      string  `json:"obs_units"`
	OBS_LEV        string  `json:"obs_lev"`
	OBTYPE         string  `json:"obtype"`
	VX_MASK        string  `json:"vx_mask"`
	INTERP_MTHD    string  `json:"interp_mthd"`
	INTERP_PNTS    int     `json:"interp_pnts"`
	FCST_THRESH    string  `json:"fcst_thresh"`
	OBS_THRESH     string  `json:"obs_thresh"`
	COV_THRESH     string  `json:"cov_thresh"`
	ALPHA          float64 `json:"alpha"`
	LINE_TYPE      string  `json:"line_type"`
}

type STAT_NBRCTC_header struct {
	VERSION        string  `json:"version"`
	MODEL          string  `json:"model"`
	DESC           string  `json:"desc"`
	FCST_VALID_BEG int     `json:"fcst_valid_beg"`
	FCST_VALID_END int     `json:"fcst_valid_end"`
	OBS_LEAD       int     `json:"obs_lead"`
	OBS_VALID_BEG  int     `json:"obs_valid_beg"`
	OBS_VALID_END  int     `json:"obs_valid_end"`
	FCST_VAR       string  `json:"fcst_var"`
	FCST_UNITS     string  `json:"fcst_units"`
	FCST_LEV       string  `json:"fcst_lev"`
	OBS_VAR        string  `json:"obs_var"`
	OBS_UNITS      string  `json:"obs_units"`
	OBS_LEV        string  `json:"obs_lev"`
	OBTYPE         string  `json:"obtype"`
	VX_MASK        string  `json:"vx_mask"`
	INTERP_MTHD    string  `json:"interp_mthd"`
	INTERP_PNTS    int     `json:"interp_pnts"`
	FCST_THRESH    string  `json:"fcst_thresh"`
	OBS_THRESH     string  `json:"obs_thresh"`
	COV_THRESH     string  `json:"cov_thresh"`
	ALPHA          float64 `json:"alpha"`
	LINE_TYPE      string  `json:"line_type"`
}

type STAT_GENMPR_header struct {
	VERSION        string  `json:"version"`
	MODEL          string  `json:"model"`
	DESC           string  `json:"desc"`
	FCST_VALID_BEG int     `json:"fcst_valid_beg"`
	FCST_VALID_END int     `json:"fcst_valid_end"`
	OBS_LEAD       int     `json:"obs_lead"`
	OBS_VALID_BEG  int     `json:"obs_valid_beg"`
	OBS_VALID_END  int     `json:"obs_valid_end"`
	FCST_VAR       string  `json:"fcst_var"`
	FCST_UNITS     string  `json:"fcst_units"`
	FCST_LEV       string  `json:"fcst_lev"`
	OBS_VAR        string  `json:"obs_var"`
	OBS_UNITS      string  `json:"obs_units"`
	OBS_LEV        string  `json:"obs_lev"`
	OBTYPE         string  `json:"obtype"`
	VX_MASK        string  `json:"vx_mask"`
	INTERP_MTHD    string  `json:"interp_mthd"`
	INTERP_PNTS    int     `json:"interp_pnts"`
	FCST_THRESH    string  `json:"fcst_thresh"`
	OBS_THRESH     string  `json:"obs_thresh"`
	COV_THRESH     string  `json:"cov_thresh"`
	ALPHA          float64 `json:"alpha"`
	LINE_TYPE      string  `json:"line_type"`
}

// fillHeader functions
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

//line data struct definitions
type STAT_SEEPS_MPR struct {
	OBS_SID  string  `json:"obs_sid,omitempty"`
	OBS_LAT  float64 `json:"obs_lat,omitempty"`
	OBS_LON  float64 `json:"obs_lon,omitempty"`
	FCST     float64 `json:"fcst,omitempty"`
	OBS      float64 `json:"obs,omitempty"`
	OBS_QC   string  `json:"obs_qc,omitempty"`
	FCST_CAT int     `json:"fcst_cat,omitempty"`
	OBS_CAT  int     `json:"obs_cat,omitempty"`
	P1       float64 `json:"p1,omitempty"`
	P2       float64 `json:"p2,omitempty"`
	T1       float64 `json:"t1,omitempty"`
	T2       float64 `json:"t2,omitempty"`
	SEEPS    float64 `json:"seeps,omitempty"`
}

type STAT_RPS struct {
	TOTAL     int     `json:"total,omitempty"`
	N_PROB    int     `json:"n_prob,omitempty"`
	RPS_REL   float64 `json:"rps_rel,omitempty"`
	RPS_RES   float64 `json:"rps_res,omitempty"`
	RPS_UNC   float64 `json:"rps_unc,omitempty"`
	RPS       float64 `json:"rps,omitempty"`
	RPSS      float64 `json:"rpss,omitempty"`
	RPSS_SMPL float64 `json:"rpss_smpl,omitempty"`
	RPS_COMP  float64 `json:"rps_comp,omitempty"`
}

type STAT_PHIST struct {
	TOTAL    int                    `json:"total,omitempty"`
	BIN_SIZE int                    `json:"bin_size,omitempty"`
	BIN      map[string]interface{} `json:"bin,omitempty"`
}

type STAT_CNT struct {
	TOTAL                int     `json:"total,omitempty"`
	FBAR                 float64 `json:"fbar,omitempty"`
	FBAR_NCL             float64 `json:"fbar_ncl,omitempty"`
	FBAR_NCU             float64 `json:"fbar_ncu,omitempty"`
	FBAR_BCL             float64 `json:"fbar_bcl,omitempty"`
	FBAR_BCU             float64 `json:"fbar_bcu,omitempty"`
	FSTDEV               float64 `json:"fstdev,omitempty"`
	FSTDEV_NCL           float64 `json:"fstdev_ncl,omitempty"`
	FSTDEV_NCU           float64 `json:"fstdev_ncu,omitempty"`
	FSTDEV_BCL           float64 `json:"fstdev_bcl,omitempty"`
	FSTDEV_BCU           float64 `json:"fstdev_bcu,omitempty"`
	OBAR                 float64 `json:"obar,omitempty"`
	OBAR_NCL             float64 `json:"obar_ncl,omitempty"`
	OBAR_NCU             float64 `json:"obar_ncu,omitempty"`
	OBAR_BCL             float64 `json:"obar_bcl,omitempty"`
	OBAR_BCU             float64 `json:"obar_bcu,omitempty"`
	OSTDEV               float64 `json:"ostdev,omitempty"`
	OSTDEV_NCL           float64 `json:"ostdev_ncl,omitempty"`
	OSTDEV_NCU           float64 `json:"ostdev_ncu,omitempty"`
	OSTDEV_BCL           float64 `json:"ostdev_bcl,omitempty"`
	OSTDEV_BCU           float64 `json:"ostdev_bcu,omitempty"`
	PR_CORR              float64 `json:"pr_corr,omitempty"`
	PR_CORR_NCL          float64 `json:"pr_corr_ncl,omitempty"`
	PR_CORR_NCU          float64 `json:"pr_corr_ncu,omitempty"`
	PR_CORR_BCL          float64 `json:"pr_corr_bcl,omitempty"`
	PR_CORR_BCU          float64 `json:"pr_corr_bcu,omitempty"`
	SP_CORR              float64 `json:"sp_corr,omitempty"`
	KT_CORR              float64 `json:"kt_corr,omitempty"`
	RANKS                int     `json:"ranks,omitempty"`
	FRANK_TIES           int     `json:"frank_ties,omitempty"`
	ORANK_TIES           int     `json:"orank_ties,omitempty"`
	ME                   float64 `json:"me,omitempty"`
	ME_NCL               float64 `json:"me_ncl,omitempty"`
	ME_NCU               float64 `json:"me_ncu,omitempty"`
	ME_BCL               float64 `json:"me_bcl,omitempty"`
	ME_BCU               float64 `json:"me_bcu,omitempty"`
	ESTDEV               float64 `json:"estdev,omitempty"`
	ESTDEV_NCL           float64 `json:"estdev_ncl,omitempty"`
	ESTDEV_NCU           float64 `json:"estdev_ncu,omitempty"`
	ESTDEV_BCL           float64 `json:"estdev_bcl,omitempty"`
	ESTDEV_BCU           float64 `json:"estdev_bcu,omitempty"`
	MBIAS                float64 `json:"mbias,omitempty"`
	MBIAS_BCL            float64 `json:"mbias_bcl,omitempty"`
	MBIAS_BCU            float64 `json:"mbias_bcu,omitempty"`
	MAE                  float64 `json:"mae,omitempty"`
	MAE_BCL              float64 `json:"mae_bcl,omitempty"`
	MAE_BCU              float64 `json:"mae_bcu,omitempty"`
	MSE                  float64 `json:"mse,omitempty"`
	MSE_BCL              float64 `json:"mse_bcl,omitempty"`
	MSE_BCU              float64 `json:"mse_bcu,omitempty"`
	BCMSE                float64 `json:"bcmse,omitempty"`
	BCMSE_BCL            float64 `json:"bcmse_bcl,omitempty"`
	BCMSE_BCU            float64 `json:"bcmse_bcu,omitempty"`
	RMSE                 float64 `json:"rmse,omitempty"`
	RMSE_BCL             float64 `json:"rmse_bcl,omitempty"`
	RMSE_BCU             float64 `json:"rmse_bcu,omitempty"`
	E10                  float64 `json:"e10,omitempty"`
	E10_BCL              float64 `json:"e10_bcl,omitempty"`
	E10_BCU              float64 `json:"e10_bcu,omitempty"`
	E25                  float64 `json:"e25,omitempty"`
	E25_BCL              float64 `json:"e25_bcl,omitempty"`
	E25_BCU              float64 `json:"e25_bcu,omitempty"`
	E50                  float64 `json:"e50,omitempty"`
	E50_BCL              float64 `json:"e50_bcl,omitempty"`
	E50_BCU              float64 `json:"e50_bcu,omitempty"`
	E75                  float64 `json:"e75,omitempty"`
	E75_BCL              float64 `json:"e75_bcl,omitempty"`
	E75_BCU              float64 `json:"e75_bcu,omitempty"`
	E90                  float64 `json:"e90,omitempty"`
	E90_BCL              float64 `json:"e90_bcl,omitempty"`
	E90_BCU              float64 `json:"e90_bcu,omitempty"`
	EIQR                 float64 `json:"eiqr,omitempty"`
	EIQR_BCL             float64 `json:"eiqr_bcl,omitempty"`
	EIQR_BCU             float64 `json:"eiqr_bcu,omitempty"`
	MAD                  float64 `json:"mad,omitempty"`
	MAD_BCL              float64 `json:"mad_bcl,omitempty"`
	MAD_BCU              float64 `json:"mad_bcu,omitempty"`
	ANOM_CORR            float64 `json:"anom_corr,omitempty"`
	ANOM_CORR_NCL        float64 `json:"anom_corr_ncl,omitempty"`
	ANOM_CORR_NCU        float64 `json:"anom_corr_ncu,omitempty"`
	ANOM_CORR_BCL        float64 `json:"anom_corr_bcl,omitempty"`
	ANOM_CORR_BCU        float64 `json:"anom_corr_bcu,omitempty"`
	ME2                  float64 `json:"me2,omitempty"`
	ME2_BCL              float64 `json:"me2_bcl,omitempty"`
	ME2_BCU              float64 `json:"me2_bcu,omitempty"`
	MSESS                float64 `json:"msess,omitempty"`
	MSESS_BCL            float64 `json:"msess_bcl,omitempty"`
	MSESS_BCU            float64 `json:"msess_bcu,omitempty"`
	RMSFA                float64 `json:"rmsfa,omitempty"`
	RMSFA_BCL            float64 `json:"rmsfa_bcl,omitempty"`
	RMSFA_BCU            float64 `json:"rmsfa_bcu,omitempty"`
	RMSOA                float64 `json:"rmsoa,omitempty"`
	RMSOA_BCL            float64 `json:"rmsoa_bcl,omitempty"`
	RMSOA_BCU            float64 `json:"rmsoa_bcu,omitempty"`
	ANOM_CORR_UNCNTR     float64 `json:"anom_corr_uncntr,omitempty"`
	ANOM_CORR_UNCNTR_BCL float64 `json:"anom_corr_uncntr_bcl,omitempty"`
	ANOM_CORR_UNCNTR_BCU float64 `json:"anom_corr_uncntr_bcu,omitempty"`
	SI                   float64 `json:"si,omitempty"`
	SI_BCL               float64 `json:"si_bcl,omitempty"`
	SI_BCU               float64 `json:"si_bcu,omitempty"`
}

type STAT_CTC struct {
	TOTAL    int     `json:"total,omitempty"`
	FY_OY    float64 `json:"fy_oy,omitempty"`
	FY_ON    float64 `json:"fy_on,omitempty"`
	FN_OY    float64 `json:"fn_oy,omitempty"`
	FN_ON    float64 `json:"fn_on,omitempty"`
	EC_VALUE float64 `json:"ec_value,omitempty"`
}

type STAT_SEEPS struct {
	TOTAL     int     `json:"total,omitempty"`
	S12       string  `json:"s12,omitempty"`
	S13       string  `json:"s13,omitempty"`
	S21       string  `json:"s21,omitempty"`
	S23       string  `json:"s23,omitempty"`
	S31       string  `json:"s31,omitempty"`
	S32       string  `json:"s32,omitempty"`
	PF1       float64 `json:"pf1,omitempty"`
	PF2       float64 `json:"pf2,omitempty"`
	PF3       float64 `json:"pf3,omitempty"`
	PV1       float64 `json:"pv1,omitempty"`
	PV2       float64 `json:"pv2,omitempty"`
	PV3       float64 `json:"pv3,omitempty"`
	MEAN_FCST float64 `json:"mean_fcst,omitempty"`
	MEAN_OBS  float64 `json:"mean_obs,omitempty"`
	SEEPS     float64 `json:"seeps,omitempty"`
}

type STAT_GRAD struct {
	TOTAL      int     `json:"total,omitempty"`
	FGBAR      float64 `json:"fgbar,omitempty"`
	OGBAR      float64 `json:"ogbar,omitempty"`
	MGBAR      float64 `json:"mgbar,omitempty"`
	EGBAR      float64 `json:"egbar,omitempty"`
	S1         float64 `json:"s1,omitempty"`
	S1_OG      float64 `json:"s1_og,omitempty"`
	FGOG_RATIO float64 `json:"fgog_ratio,omitempty"`
	DX         float64 `json:"dx,omitempty"`
	DY         float64 `json:"dy,omitempty"`
}

type TCST_TCDIAG struct {
	TOTAL        int                    `json:"total,omitempty"`
	INDEX        int                    `json:"index,omitempty"`
	DIAG_SOURCE  float64                `json:"diag_source,omitempty"`
	TRACK_SOURCE string                 `json:"track_source,omitempty"`
	FIELD_SOURCE string                 `json:"field_source,omitempty"`
	DIAG         map[string]interface{} `json:"diag,omitempty"`
	INIT         int                    `json:"init,omitempty"`
}

type STAT_ISC struct {
	TOTAL    int     `json:"total,omitempty"`
	TILE_DIM int     `json:"tile_dim,omitempty"`
	TILE_XLL int     `json:"tile_xll,omitempty"`
	TILE_YLL int     `json:"tile_yll,omitempty"`
	NSCALE   int     `json:"nscale,omitempty"`
	ISCALE   int     `json:"iscale,omitempty"`
	MSE      float64 `json:"mse,omitempty"`
	ISC      float64 `json:"isc,omitempty"`
	FENERGY2 float64 `json:"fenergy2,omitempty"`
	OENERGY2 float64 `json:"oenergy2,omitempty"`
	BASER    float64 `json:"baser,omitempty"`
	FBIAS    float64 `json:"fbias,omitempty"`
}

type STAT_ORANK struct {
	TOTAL            int                    `json:"total,omitempty"`
	INDEX            int                    `json:"index,omitempty"`
	OBS_SID          string                 `json:"obs_sid,omitempty"`
	OBS_LAT          float64                `json:"obs_lat,omitempty"`
	OBS_LON          float64                `json:"obs_lon,omitempty"`
	OBS_LVL          float64                `json:"obs_lvl,omitempty"`
	OBS_ELV          float64                `json:"obs_elv,omitempty"`
	OBS              float64                `json:"obs,omitempty"`
	PIT              float64                `json:"pit,omitempty"`
	RANK             int                    `json:"rank,omitempty"`
	N_ENS_VLD        int                    `json:"n_ens_vld,omitempty"`
	ENS              map[string]interface{} `json:"ens,omitempty"`
	OBS_QC           string                 `json:"obs_qc,omitempty"`
	ENS_MEAN         int                    `json:"ens_mean,omitempty"`
	CLIMO_MEAN       float64                `json:"climo_mean,omitempty"`
	SPREAD           float64                `json:"spread,omitempty"`
	ENS_MEAN_OERR    int                    `json:"ens_mean_oerr,omitempty"`
	SPREAD_OERR      float64                `json:"spread_oerr,omitempty"`
	SPREAD_PLUS_OERR float64                `json:"spread_plus_oerr,omitempty"`
	CLIMO_STDEV      float64                `json:"climo_stdev,omitempty"`
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
	F_SPEED_BAR float64 `json:"f_speed_bar,omitempty"`
	O_SPEED_BAR float64 `json:"o_speed_bar,omitempty"`
}

type MODE_OBJ struct {
	OBJECT_ID                  string  `json:"object_id,omitempty"`
	OBJECT_CAT                 string  `json:"object_cat,omitempty"`
	CENTROID_X                 float64 `json:"centroid_x,omitempty"`
	CENTROID_Y                 float64 `json:"centroid_y,omitempty"`
	CENTROID_LAT               float64 `json:"centroid_lat,omitempty"`
	CENTROID_LON               float64 `json:"centroid_lon,omitempty"`
	AXIS_ANG                   float64 `json:"axis_ang,omitempty"`
	LENGTH                     float64 `json:"length,omitempty"`
	WIDTH                      float64 `json:"width,omitempty"`
	AREA                       int     `json:"area,omitempty"`
	AREA_THRESH                int     `json:"area_thresh,omitempty"`
	CURVATURE                  float64 `json:"curvature,omitempty"`
	CURVATURE_X                float64 `json:"curvature_x,omitempty"`
	CURVATURE_Y                float64 `json:"curvature_y,omitempty"`
	COMPLEXITY                 float64 `json:"complexity,omitempty"`
	INTENSITY_10               float64 `json:"intensity_10,omitempty"`
	INTENSITY_25               float64 `json:"intensity_25,omitempty"`
	INTENSITY_50               float64 `json:"intensity_50,omitempty"`
	INTENSITY_75               float64 `json:"intensity_75,omitempty"`
	INTENSITY_90               float64 `json:"intensity_90,omitempty"`
	INTENSITY_USER             float64 `json:"intensity_user,omitempty"`
	INTENSITY_SUM              float64 `json:"intensity_sum,omitempty"`
	CENTROID_DIST              float64 `json:"centroid_dist,omitempty"`
	BOUNDARY_DIST              float64 `json:"boundary_dist,omitempty"`
	CONVEX_HULL_DIST           float64 `json:"convex_hull_dist,omitempty"`
	ANGLE_DIFF                 float64 `json:"angle_diff,omitempty"`
	ASPECT_DIFF                float64 `json:"aspect_diff,omitempty"`
	AREA_RATIO                 float64 `json:"area_ratio,omitempty"`
	INTERSECTION_AREA          float64 `json:"intersection_area,omitempty"`
	UNION_AREA                 float64 `json:"union_area,omitempty"`
	SYMMETRIC_DIFF             float64 `json:"symmetric_diff,omitempty"`
	INTERSECTION_OVER_AREA     float64 `json:"intersection_over_area,omitempty"`
	CURVATURE_RATIO            float64 `json:"curvature_ratio,omitempty"`
	COMPLEXITY_RATIO           float64 `json:"complexity_ratio,omitempty"`
	PERCENTILE_INTENSITY_RATIO float64 `json:"percentile_intensity_ratio,omitempty"`
	INTEREST                   float64 `json:"interest,omitempty"`
}

type STAT_CTS struct {
	TOTAL      int     `json:"total,omitempty"`
	BASER      float64 `json:"baser,omitempty"`
	BASER_NCL  float64 `json:"baser_ncl,omitempty"`
	BASER_NCU  float64 `json:"baser_ncu,omitempty"`
	BASER_BCL  float64 `json:"baser_bcl,omitempty"`
	BASER_BCU  float64 `json:"baser_bcu,omitempty"`
	FMEAN      float64 `json:"fmean,omitempty"`
	FMEAN_NCL  float64 `json:"fmean_ncl,omitempty"`
	FMEAN_NCU  float64 `json:"fmean_ncu,omitempty"`
	FMEAN_BCL  float64 `json:"fmean_bcl,omitempty"`
	FMEAN_BCU  float64 `json:"fmean_bcu,omitempty"`
	ACC        float64 `json:"acc,omitempty"`
	ACC_NCL    float64 `json:"acc_ncl,omitempty"`
	ACC_NCU    float64 `json:"acc_ncu,omitempty"`
	ACC_BCL    float64 `json:"acc_bcl,omitempty"`
	ACC_BCU    float64 `json:"acc_bcu,omitempty"`
	FBIAS      float64 `json:"fbias,omitempty"`
	FBIAS_BCL  float64 `json:"fbias_bcl,omitempty"`
	FBIAS_BCU  float64 `json:"fbias_bcu,omitempty"`
	PODY       float64 `json:"pody,omitempty"`
	PODY_NCL   float64 `json:"pody_ncl,omitempty"`
	PODY_NCU   float64 `json:"pody_ncu,omitempty"`
	PODY_BCL   float64 `json:"pody_bcl,omitempty"`
	PODY_BCU   float64 `json:"pody_bcu,omitempty"`
	PODN       float64 `json:"podn,omitempty"`
	PODN_NCL   float64 `json:"podn_ncl,omitempty"`
	PODN_NCU   float64 `json:"podn_ncu,omitempty"`
	PODN_BCL   float64 `json:"podn_bcl,omitempty"`
	PODN_BCU   float64 `json:"podn_bcu,omitempty"`
	POFD       float64 `json:"pofd,omitempty"`
	POFD_NCL   float64 `json:"pofd_ncl,omitempty"`
	POFD_NCU   float64 `json:"pofd_ncu,omitempty"`
	POFD_BCL   float64 `json:"pofd_bcl,omitempty"`
	POFD_BCU   float64 `json:"pofd_bcu,omitempty"`
	FAR        float64 `json:"far,omitempty"`
	FAR_NCL    float64 `json:"far_ncl,omitempty"`
	FAR_NCU    float64 `json:"far_ncu,omitempty"`
	FAR_BCL    float64 `json:"far_bcl,omitempty"`
	FAR_BCU    float64 `json:"far_bcu,omitempty"`
	CSI        float64 `json:"csi,omitempty"`
	CSI_NCL    float64 `json:"csi_ncl,omitempty"`
	CSI_NCU    float64 `json:"csi_ncu,omitempty"`
	CSI_BCL    float64 `json:"csi_bcl,omitempty"`
	CSI_BCU    float64 `json:"csi_bcu,omitempty"`
	GSS        float64 `json:"gss,omitempty"`
	GSS_BCL    float64 `json:"gss_bcl,omitempty"`
	GSS_BCU    float64 `json:"gss_bcu,omitempty"`
	HK         float64 `json:"hk,omitempty"`
	HK_NCL     float64 `json:"hk_ncl,omitempty"`
	HK_NCU     float64 `json:"hk_ncu,omitempty"`
	HK_BCL     float64 `json:"hk_bcl,omitempty"`
	HK_BCU     float64 `json:"hk_bcu,omitempty"`
	HSS        float64 `json:"hss,omitempty"`
	HSS_BCL    float64 `json:"hss_bcl,omitempty"`
	HSS_BCU    float64 `json:"hss_bcu,omitempty"`
	ODDS       float64 `json:"odds,omitempty"`
	ODDS_NCL   float64 `json:"odds_ncl,omitempty"`
	ODDS_NCU   float64 `json:"odds_ncu,omitempty"`
	ODDS_BCL   float64 `json:"odds_bcl,omitempty"`
	ODDS_BCU   float64 `json:"odds_bcu,omitempty"`
	LODDS      float64 `json:"lodds,omitempty"`
	LODDS_NCL  float64 `json:"lodds_ncl,omitempty"`
	LODDS_NCU  float64 `json:"lodds_ncu,omitempty"`
	LODDS_BCL  float64 `json:"lodds_bcl,omitempty"`
	LODDS_BCU  float64 `json:"lodds_bcu,omitempty"`
	ORSS       float64 `json:"orss,omitempty"`
	ORSS_NCL   float64 `json:"orss_ncl,omitempty"`
	ORSS_NCU   float64 `json:"orss_ncu,omitempty"`
	ORSS_BCL   float64 `json:"orss_bcl,omitempty"`
	ORSS_BCU   float64 `json:"orss_bcu,omitempty"`
	EDS        float64 `json:"eds,omitempty"`
	EDS_NCL    float64 `json:"eds_ncl,omitempty"`
	EDS_NCU    float64 `json:"eds_ncu,omitempty"`
	EDS_BCL    float64 `json:"eds_bcl,omitempty"`
	EDS_BCU    float64 `json:"eds_bcu,omitempty"`
	SEDS       float64 `json:"seds,omitempty"`
	SEDS_NCL   float64 `json:"seds_ncl,omitempty"`
	SEDS_NCU   float64 `json:"seds_ncu,omitempty"`
	SEDS_BCL   float64 `json:"seds_bcl,omitempty"`
	SEDS_BCU   float64 `json:"seds_bcu,omitempty"`
	EDI        float64 `json:"edi,omitempty"`
	EDI_NCL    float64 `json:"edi_ncl,omitempty"`
	EDI_NCU    float64 `json:"edi_ncu,omitempty"`
	EDI_BCL    float64 `json:"edi_bcl,omitempty"`
	EDI_BCU    float64 `json:"edi_bcu,omitempty"`
	SEDI       float64 `json:"sedi,omitempty"`
	SEDI_NCL   float64 `json:"sedi_ncl,omitempty"`
	SEDI_NCU   float64 `json:"sedi_ncu,omitempty"`
	SEDI_BCL   float64 `json:"sedi_bcl,omitempty"`
	SEDI_BCU   float64 `json:"sedi_bcu,omitempty"`
	BAGSS      float64 `json:"bagss,omitempty"`
	BAGSS_BCL  float64 `json:"bagss_bcl,omitempty"`
	BAGSS_BCU  float64 `json:"bagss_bcu,omitempty"`
	HSS_EC     float64 `json:"hss_ec,omitempty"`
	HSS_EC_BCL float64 `json:"hss_ec_bcl,omitempty"`
	HSS_EC_BCU float64 `json:"hss_ec_bcu,omitempty"`
	EC_VALUE   float64 `json:"ec_value,omitempty"`
}

type STAT_MCTS struct {
	TOTAL      int     `json:"total,omitempty"`
	N_CAT      int     `json:"n_cat,omitempty"`
	ACC        float64 `json:"acc,omitempty"`
	ACC_NCL    float64 `json:"acc_ncl,omitempty"`
	ACC_NCU    float64 `json:"acc_ncu,omitempty"`
	ACC_BCL    float64 `json:"acc_bcl,omitempty"`
	ACC_BCU    float64 `json:"acc_bcu,omitempty"`
	HK         float64 `json:"hk,omitempty"`
	HK_BCL     float64 `json:"hk_bcl,omitempty"`
	HK_BCU     float64 `json:"hk_bcu,omitempty"`
	HSS        float64 `json:"hss,omitempty"`
	HSS_BCL    float64 `json:"hss_bcl,omitempty"`
	HSS_BCU    float64 `json:"hss_bcu,omitempty"`
	GER        float64 `json:"ger,omitempty"`
	GER_BCL    float64 `json:"ger_bcl,omitempty"`
	GER_BCU    float64 `json:"ger_bcu,omitempty"`
	HSS_EC     float64 `json:"hss_ec,omitempty"`
	HSS_EC_BCL float64 `json:"hss_ec_bcl,omitempty"`
	HSS_EC_BCU float64 `json:"hss_ec_bcu,omitempty"`
	EC_VALUE   float64 `json:"ec_value,omitempty"`
}

type STAT_PRC struct {
	TOTAL  int                    `json:"total,omitempty"`
	THRESH map[string]interface{} `json:"thresh,omitempty"`
}

type STAT_ECLV struct {
	TOTAL       int                    `json:"total,omitempty"`
	BASER       float64                `json:"baser,omitempty"`
	VALUE_BASER int                    `json:"value_baser,omitempty"`
	PTS         map[string]interface{} `json:"pts,omitempty"`
}

type STAT_RHIST struct {
	TOTAL int                    `json:"total,omitempty"`
	RANK  map[string]interface{} `json:"rank,omitempty"`
}

type STAT_SSVAR struct {
	TOTAL       int     `json:"total,omitempty"`
	N_BIN       int     `json:"n_bin,omitempty"`
	BIN_I       int     `json:"bin_i,omitempty"`
	BIN_N       int     `json:"bin_n,omitempty"`
	VAR_MIN     float64 `json:"var_min,omitempty"`
	VAR_MAX     float64 `json:"var_max,omitempty"`
	VAR_MEAN    float64 `json:"var_mean,omitempty"`
	FBAR        float64 `json:"fbar,omitempty"`
	OBAR        float64 `json:"obar,omitempty"`
	FOBAR       float64 `json:"fobar,omitempty"`
	FFBAR       float64 `json:"ffbar,omitempty"`
	OOBAR       float64 `json:"oobar,omitempty"`
	FBAR_NCL    float64 `json:"fbar_ncl,omitempty"`
	FBAR_NCU    float64 `json:"fbar_ncu,omitempty"`
	FSTDEV      float64 `json:"fstdev,omitempty"`
	FSTDEV_NCL  float64 `json:"fstdev_ncl,omitempty"`
	FSTDEV_NCU  float64 `json:"fstdev_ncu,omitempty"`
	OBAR_NCL    float64 `json:"obar_ncl,omitempty"`
	OBAR_NCU    float64 `json:"obar_ncu,omitempty"`
	OSTDEV      float64 `json:"ostdev,omitempty"`
	OSTDEV_NCL  float64 `json:"ostdev_ncl,omitempty"`
	OSTDEV_NCU  float64 `json:"ostdev_ncu,omitempty"`
	PR_CORR     float64 `json:"pr_corr,omitempty"`
	PR_CORR_NCL float64 `json:"pr_corr_ncl,omitempty"`
	PR_CORR_NCU float64 `json:"pr_corr_ncu,omitempty"`
	ME          float64 `json:"me,omitempty"`
	ME_NCL      float64 `json:"me_ncl,omitempty"`
	ME_NCU      float64 `json:"me_ncu,omitempty"`
	ESTDEV      float64 `json:"estdev,omitempty"`
	ESTDEV_NCL  float64 `json:"estdev_ncl,omitempty"`
	ESTDEV_NCU  float64 `json:"estdev_ncu,omitempty"`
	MBIAS       float64 `json:"mbias,omitempty"`
	MSE         float64 `json:"mse,omitempty"`
	BCMSE       float64 `json:"bcmse,omitempty"`
	RMSE        float64 `json:"rmse,omitempty"`
}

type TCST_TCMPR struct {
	TOTAL          int     `json:"total,omitempty"`
	INDEX          int     `json:"index,omitempty"`
	LEVEL          string  `json:"level,omitempty"`
	WATCH_WARN     string  `json:"watch_warn,omitempty"`
	INITIALS       string  `json:"initials,omitempty"`
	ALAT           float64 `json:"alat,omitempty"`
	ALON           float64 `json:"alon,omitempty"`
	BLAT           float64 `json:"blat,omitempty"`
	BLON           float64 `json:"blon,omitempty"`
	TK_ERR         float64 `json:"tk_err,omitempty"`
	X_ERR          float64 `json:"x_err,omitempty"`
	Y_ERR          float64 `json:"y_err,omitempty"`
	ALTK_ERR       float64 `json:"altk_err,omitempty"`
	CRTK_ERR       float64 `json:"crtk_err,omitempty"`
	ADLAND         float64 `json:"adland,omitempty"`
	BDLAND         float64 `json:"bdland,omitempty"`
	AMSLP          float64 `json:"amslp,omitempty"`
	BMSLP          float64 `json:"bmslp,omitempty"`
	AMAX_WIND      float64 `json:"amax_wind,omitempty"`
	BMAX_WIND      float64 `json:"bmax_wind,omitempty"`
	AAL_WIND_34    float64 `json:"aal_wind_34,omitempty"`
	BAL_WIND_34    float64 `json:"bal_wind_34,omitempty"`
	ANE_WIND_34    float64 `json:"ane_wind_34,omitempty"`
	BNE_WIND_34    float64 `json:"bne_wind_34,omitempty"`
	ASE_WIND_34    float64 `json:"ase_wind_34,omitempty"`
	BSE_WIND_34    float64 `json:"bse_wind_34,omitempty"`
	ASW_WIND_34    float64 `json:"asw_wind_34,omitempty"`
	BSW_WIND_34    float64 `json:"bsw_wind_34,omitempty"`
	ANW_WIND_34    float64 `json:"anw_wind_34,omitempty"`
	BNW_WIND_34    float64 `json:"bnw_wind_34,omitempty"`
	AAL_WIND_50    float64 `json:"aal_wind_50,omitempty"`
	BAL_WIND_50    float64 `json:"bal_wind_50,omitempty"`
	ANE_WIND_50    float64 `json:"ane_wind_50,omitempty"`
	BNE_WIND_50    float64 `json:"bne_wind_50,omitempty"`
	ASE_WIND_50    float64 `json:"ase_wind_50,omitempty"`
	BSE_WIND_50    float64 `json:"bse_wind_50,omitempty"`
	ASW_WIND_50    float64 `json:"asw_wind_50,omitempty"`
	BSW_WIND_50    float64 `json:"bsw_wind_50,omitempty"`
	ANW_WIND_50    float64 `json:"anw_wind_50,omitempty"`
	BNW_WIND_50    float64 `json:"bnw_wind_50,omitempty"`
	AAL_WIND_64    float64 `json:"aal_wind_64,omitempty"`
	BAL_WIND_64    float64 `json:"bal_wind_64,omitempty"`
	ANE_WIND_64    float64 `json:"ane_wind_64,omitempty"`
	BNE_WIND_64    float64 `json:"bne_wind_64,omitempty"`
	ASE_WIND_64    float64 `json:"ase_wind_64,omitempty"`
	BSE_WIND_64    float64 `json:"bse_wind_64,omitempty"`
	ASW_WIND_64    float64 `json:"asw_wind_64,omitempty"`
	BSW_WIND_64    float64 `json:"bsw_wind_64,omitempty"`
	ANW_WIND_64    float64 `json:"anw_wind_64,omitempty"`
	BNW_WIND_64    float64 `json:"bnw_wind_64,omitempty"`
	ARADP          string  `json:"aradp,omitempty"`
	BRADP          float64 `json:"bradp,omitempty"`
	ARRP           int     `json:"arrp,omitempty"`
	BRRP           float64 `json:"brrp,omitempty"`
	AMRD           int     `json:"amrd,omitempty"`
	BMRD           float64 `json:"bmrd,omitempty"`
	AGUSTS         int     `json:"agusts,omitempty"`
	BGUSTS         float64 `json:"bgusts,omitempty"`
	AEYE           int     `json:"aeye,omitempty"`
	BEYE           float64 `json:"beye,omitempty"`
	ADIR           int     `json:"adir,omitempty"`
	BDIR           float64 `json:"bdir,omitempty"`
	ASPEED         int     `json:"aspeed,omitempty"`
	BSPEED         float64 `json:"bspeed,omitempty"`
	ADEPTH         int     `json:"adepth,omitempty"`
	BDEPTH         float64 `json:"bdepth,omitempty"`
	NUM_MEMBERS    float64 `json:"num_members,omitempty"`
	TRACK_SPREAD   float64 `json:"track_spread,omitempty"`
	TRACK_STDEV    float64 `json:"track_stdev,omitempty"`
	MSLP_STDEV     float64 `json:"mslp_stdev,omitempty"`
	MAX_WIND_STDEV float64 `json:"max_wind_stdev,omitempty"`
	INIT           int     `json:"init,omitempty"`
}

type TCST_PROBRIRW struct {
	ALAT        float64                `json:"alat,omitempty"`
	ALON        float64                `json:"alon,omitempty"`
	BLAT        float64                `json:"blat,omitempty"`
	BLON        float64                `json:"blon,omitempty"`
	INITIALS    string                 `json:"initials,omitempty"`
	TK_ERR      float64                `json:"tk_err,omitempty"`
	X_ERR       float64                `json:"x_err,omitempty"`
	Y_ERR       float64                `json:"y_err,omitempty"`
	ADLAND      float64                `json:"adland,omitempty"`
	BDLAND      float64                `json:"bdland,omitempty"`
	RIRW_BEG    int                    `json:"rirw_beg,omitempty"`
	RIRW_END    int                    `json:"rirw_end,omitempty"`
	RIRW_WINDOW int                    `json:"rirw_window,omitempty"`
	AWIND_END   float64                `json:"awind_end,omitempty"`
	BWIND_BEG   float64                `json:"bwind_beg,omitempty"`
	BWIND_END   float64                `json:"bwind_end,omitempty"`
	BDELTA      float64                `json:"bdelta,omitempty"`
	BDELTA_MAX  float64                `json:"bdelta_max,omitempty"`
	BLEVEL_BEG  string                 `json:"blevel_beg,omitempty"`
	BLEVEL_END  string                 `json:"blevel_end,omitempty"`
	THRESH      map[string]interface{} `json:"thresh,omitempty"`
	INIT        int                    `json:"init,omitempty"`
}

type STAT_NBRCNT struct {
	TOTAL      int     `json:"total,omitempty"`
	FBS        float64 `json:"fbs,omitempty"`
	FBS_BCL    float64 `json:"fbs_bcl,omitempty"`
	FBS_BCU    float64 `json:"fbs_bcu,omitempty"`
	FSS        float64 `json:"fss,omitempty"`
	FSS_BCL    float64 `json:"fss_bcl,omitempty"`
	FSS_BCU    float64 `json:"fss_bcu,omitempty"`
	AFSS       float64 `json:"afss,omitempty"`
	AFSS_BCL   float64 `json:"afss_bcl,omitempty"`
	AFSS_BCU   float64 `json:"afss_bcu,omitempty"`
	UFSS       float64 `json:"ufss,omitempty"`
	UFSS_BCL   float64 `json:"ufss_bcl,omitempty"`
	UFSS_BCU   float64 `json:"ufss_bcu,omitempty"`
	F_RATE     float64 `json:"f_rate,omitempty"`
	F_RATE_BCL float64 `json:"f_rate_bcl,omitempty"`
	F_RATE_BCU float64 `json:"f_rate_bcu,omitempty"`
	O_RATE     float64 `json:"o_rate,omitempty"`
	O_RATE_BCL float64 `json:"o_rate_bcl,omitempty"`
	O_RATE_BCU float64 `json:"o_rate_bcu,omitempty"`
}

type STAT_NBRCTC struct {
	TOTAL int     `json:"total,omitempty"`
	FY_OY float64 `json:"fy_oy,omitempty"`
	FY_ON float64 `json:"fy_on,omitempty"`
	FN_OY float64 `json:"fn_oy,omitempty"`
	FN_ON float64 `json:"fn_on,omitempty"`
}

type STAT_ECNT struct {
	TOTAL            int     `json:"total,omitempty"`
	N_ENS            int     `json:"n_ens,omitempty"`
	CRPS             float64 `json:"crps,omitempty"`
	CRPSS            float64 `json:"crpss,omitempty"`
	IGN              float64 `json:"ign,omitempty"`
	ME               float64 `json:"me,omitempty"`
	RMSE             float64 `json:"rmse,omitempty"`
	SPREAD           float64 `json:"spread,omitempty"`
	ME_OERR          float64 `json:"me_oerr,omitempty"`
	RMSE_OERR        float64 `json:"rmse_oerr,omitempty"`
	SPREAD_OERR      float64 `json:"spread_oerr,omitempty"`
	SPREAD_PLUS_OERR float64 `json:"spread_plus_oerr,omitempty"`
	CRPSCL           float64 `json:"crpscl,omitempty"`
	CRPS_EMP         float64 `json:"crps_emp,omitempty"`
	CRPSCL_EMP       float64 `json:"crpscl_emp,omitempty"`
	CRPSS_EMP        float64 `json:"crpss_emp,omitempty"`
	CRPS_EMP_FAIR    float64 `json:"crps_emp_fair,omitempty"`
	SPREAD_MD        float64 `json:"spread_md,omitempty"`
	MAE              float64 `json:"mae,omitempty"`
	MAE_OERR         float64 `json:"mae_oerr,omitempty"`
	BIAS_RATIO       float64 `json:"bias_ratio,omitempty"`
	N_GE_OBS         int     `json:"n_ge_obs,omitempty"`
	ME_GE_OBS        float64 `json:"me_ge_obs,omitempty"`
	N_LT_OBS         int     `json:"n_lt_obs,omitempty"`
	ME_LT_OBS        float64 `json:"me_lt_obs,omitempty"`
}

type STAT_RELP struct {
	TOTAL int                    `json:"total,omitempty"`
	ENS   map[string]interface{} `json:"ens,omitempty"`
}

type STAT_GENMPR struct {
	TOTAL      int     `json:"total,omitempty"`
	INDEX      int     `json:"index,omitempty"`
	STORM_ID   string  `json:"storm_id,omitempty"`
	PROB_LEAD  float64 `json:"prob_lead,omitempty"`
	PROB_VAL   float64 `json:"prob_val,omitempty"`
	AGEN_INIT  string  `json:"agen_init,omitempty"`
	AGEN_FHR   string  `json:"agen_fhr,omitempty"`
	AGEN_LAT   float64 `json:"agen_lat,omitempty"`
	AGEN_LON   float64 `json:"agen_lon,omitempty"`
	AGEN_DLAND float64 `json:"agen_dland,omitempty"`
	BGEN_LAT   float64 `json:"bgen_lat,omitempty"`
	BGEN_LON   float64 `json:"bgen_lon,omitempty"`
	BGEN_DLAND float64 `json:"bgen_dland,omitempty"`
	GEN_DIST   float64 `json:"gen_dist,omitempty"`
	GEN_TDIFF  string  `json:"gen_tdiff,omitempty"`
	INIT_TDIFF string  `json:"init_tdiff,omitempty"`
	DEV_CAT    string  `json:"dev_cat,omitempty"`
	OPS_CAT    string  `json:"ops_cat,omitempty"`
}

type STAT_SSIDX struct {
	FCST_MODEL string  `json:"fcst_model,omitempty"`
	REF_MODEL  string  `json:"ref_model,omitempty"`
	N_INIT     int     `json:"n_init,omitempty"`
	N_TERM     int     `json:"n_term,omitempty"`
	N_VLD      int     `json:"n_vld,omitempty"`
	SS_INDEX   float64 `json:"ss_index,omitempty"`
}

type STAT_DMAP struct {
	TOTAL      int     `json:"total,omitempty"`
	FY         int     `json:"fy,omitempty"`
	OY         int     `json:"oy,omitempty"`
	FBIAS      float64 `json:"fbias,omitempty"`
	BADDELEY   float64 `json:"baddeley,omitempty"`
	HAUSDORFF  float64 `json:"hausdorff,omitempty"`
	MED_FO     float64 `json:"med_fo,omitempty"`
	MED_OF     float64 `json:"med_of,omitempty"`
	MED_MIN    float64 `json:"med_min,omitempty"`
	MED_MAX    float64 `json:"med_max,omitempty"`
	MED_MEAN   float64 `json:"med_mean,omitempty"`
	FOM_FO     float64 `json:"fom_fo,omitempty"`
	FOM_OF     float64 `json:"fom_of,omitempty"`
	FOM_MIN    float64 `json:"fom_min,omitempty"`
	FOM_MAX    float64 `json:"fom_max,omitempty"`
	FOM_MEAN   float64 `json:"fom_mean,omitempty"`
	ZHU_FO     float64 `json:"zhu_fo,omitempty"`
	ZHU_OF     float64 `json:"zhu_of,omitempty"`
	ZHU_MIN    float64 `json:"zhu_min,omitempty"`
	ZHU_MAX    float64 `json:"zhu_max,omitempty"`
	ZHU_MEAN   float64 `json:"zhu_mean,omitempty"`
	G          float64 `json:"g,omitempty"`
	GBETA      float64 `json:"gbeta,omitempty"`
	BETA_VALUE float64 `json:"beta_value,omitempty"`
}

type STAT_PSTD struct {
	TOTAL       int                    `json:"total,omitempty"`
	THRESH      map[string]interface{} `json:"thresh,omitempty"`
	BASER_NCL   float64                `json:"baser_ncl,omitempty"`
	BASER_NCU   float64                `json:"baser_ncu,omitempty"`
	RELIABILITY float64                `json:"reliability,omitempty"`
	RESOLUTION  float64                `json:"resolution,omitempty"`
	UNCERTAINTY float64                `json:"uncertainty,omitempty"`
	ROC_AUC     float64                `json:"roc_auc,omitempty"`
	BRIER       float64                `json:"brier,omitempty"`
	BRIER_NCL   float64                `json:"brier_ncl,omitempty"`
	BRIER_NCU   float64                `json:"brier_ncu,omitempty"`
	BRIERCL     float64                `json:"briercl,omitempty"`
	BRIERCL_NCL float64                `json:"briercl_ncl,omitempty"`
	BRIERCL_NCU float64                `json:"briercl_ncu,omitempty"`
	BSS         float64                `json:"bss,omitempty"`
	BSS_SMPL    float64                `json:"bss_smpl,omitempty"`
	THRESH_I    int                    `json:"thresh_i,omitempty"`
}

type STAT_PJC struct {
	TOTAL  int                    `json:"total,omitempty"`
	THRESH map[string]interface{} `json:"thresh,omitempty"`
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

type STAT_VAL1L2 struct {
	TOTAL        int     `json:"total,omitempty"`
	UFABAR       float64 `json:"ufabar,omitempty"`
	VFABAR       float64 `json:"vfabar,omitempty"`
	UOABAR       float64 `json:"uoabar,omitempty"`
	VOABAR       float64 `json:"voabar,omitempty"`
	UVFOABAR     float64 `json:"uvfoabar,omitempty"`
	UVFFABAR     float64 `json:"uvffabar,omitempty"`
	UVOOABAR     float64 `json:"uvooabar,omitempty"`
	FA_SPEED_BAR float64 `json:"fa_speed_bar,omitempty"`
	OA_SPEED_BAR float64 `json:"oa_speed_bar,omitempty"`
}

type MODE_CTS struct {
	FIELD string  `json:"field,omitempty"`
	TOTAL int     `json:"total,omitempty"`
	FY_OY float64 `json:"fy_oy,omitempty"`
	FY_ON float64 `json:"fy_on,omitempty"`
	FN_OY float64 `json:"fn_oy,omitempty"`
	FN_ON float64 `json:"fn_on,omitempty"`
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

type STAT_FHO struct {
	TOTAL  int     `json:"total,omitempty"`
	F_RATE float64 `json:"f_rate,omitempty"`
	H_RATE float64 `json:"h_rate,omitempty"`
	O_RATE float64 `json:"o_rate,omitempty"`
}

type STAT_MCTC struct {
	TOTAL    int                    `json:"total,omitempty"`
	CAT      map[string]interface{} `json:"cat,omitempty"`
	EC_VALUE float64                `json:"ec_value,omitempty"`
}

type STAT_MPR struct {
	TOTAL       int     `json:"total,omitempty"`
	INDEX       int     `json:"index,omitempty"`
	OBS_SID     string  `json:"obs_sid,omitempty"`
	OBS_LAT     float64 `json:"obs_lat,omitempty"`
	OBS_LON     float64 `json:"obs_lon,omitempty"`
	OBS_LVL     float64 `json:"obs_lvl,omitempty"`
	OBS_ELV     float64 `json:"obs_elv,omitempty"`
	FCST        float64 `json:"fcst,omitempty"`
	OBS         float64 `json:"obs,omitempty"`
	OBS_QC      string  `json:"obs_qc,omitempty"`
	CLIMO_MEAN  float64 `json:"climo_mean,omitempty"`
	CLIMO_STDEV float64 `json:"climo_stdev,omitempty"`
	CLIMO_CDF   float64 `json:"climo_cdf,omitempty"`
}

type STAT_NBRCTS struct {
	TOTAL     int     `json:"total,omitempty"`
	BASER     float64 `json:"baser,omitempty"`
	BASER_NCL float64 `json:"baser_ncl,omitempty"`
	BASER_NCU float64 `json:"baser_ncu,omitempty"`
	BASER_BCL float64 `json:"baser_bcl,omitempty"`
	BASER_BCU float64 `json:"baser_bcu,omitempty"`
	FMEAN     float64 `json:"fmean,omitempty"`
	FMEAN_NCL float64 `json:"fmean_ncl,omitempty"`
	FMEAN_NCU float64 `json:"fmean_ncu,omitempty"`
	FMEAN_BCL float64 `json:"fmean_bcl,omitempty"`
	FMEAN_BCU float64 `json:"fmean_bcu,omitempty"`
	ACC       float64 `json:"acc,omitempty"`
	ACC_NCL   float64 `json:"acc_ncl,omitempty"`
	ACC_NCU   float64 `json:"acc_ncu,omitempty"`
	ACC_BCL   float64 `json:"acc_bcl,omitempty"`
	ACC_BCU   float64 `json:"acc_bcu,omitempty"`
	FBIAS     float64 `json:"fbias,omitempty"`
	FBIAS_BCL float64 `json:"fbias_bcl,omitempty"`
	FBIAS_BCU float64 `json:"fbias_bcu,omitempty"`
	PODY      float64 `json:"pody,omitempty"`
	PODY_NCL  float64 `json:"pody_ncl,omitempty"`
	PODY_NCU  float64 `json:"pody_ncu,omitempty"`
	PODY_BCL  float64 `json:"pody_bcl,omitempty"`
	PODY_BCU  float64 `json:"pody_bcu,omitempty"`
	PODN      float64 `json:"podn,omitempty"`
	PODN_NCL  float64 `json:"podn_ncl,omitempty"`
	PODN_NCU  float64 `json:"podn_ncu,omitempty"`
	PODN_BCL  float64 `json:"podn_bcl,omitempty"`
	PODN_BCU  float64 `json:"podn_bcu,omitempty"`
	POFD      float64 `json:"pofd,omitempty"`
	POFD_NCL  float64 `json:"pofd_ncl,omitempty"`
	POFD_NCU  float64 `json:"pofd_ncu,omitempty"`
	POFD_BCL  float64 `json:"pofd_bcl,omitempty"`
	POFD_BCU  float64 `json:"pofd_bcu,omitempty"`
	FAR       float64 `json:"far,omitempty"`
	FAR_NCL   float64 `json:"far_ncl,omitempty"`
	FAR_NCU   float64 `json:"far_ncu,omitempty"`
	FAR_BCL   float64 `json:"far_bcl,omitempty"`
	FAR_BCU   float64 `json:"far_bcu,omitempty"`
	CSI       float64 `json:"csi,omitempty"`
	CSI_NCL   float64 `json:"csi_ncl,omitempty"`
	CSI_NCU   float64 `json:"csi_ncu,omitempty"`
	CSI_BCL   float64 `json:"csi_bcl,omitempty"`
	CSI_BCU   float64 `json:"csi_bcu,omitempty"`
	GSS       float64 `json:"gss,omitempty"`
	GSS_BCL   float64 `json:"gss_bcl,omitempty"`
	GSS_BCU   float64 `json:"gss_bcu,omitempty"`
	HK        float64 `json:"hk,omitempty"`
	HK_NCL    float64 `json:"hk_ncl,omitempty"`
	HK_NCU    float64 `json:"hk_ncu,omitempty"`
	HK_BCL    float64 `json:"hk_bcl,omitempty"`
	HK_BCU    float64 `json:"hk_bcu,omitempty"`
	HSS       float64 `json:"hss,omitempty"`
	HSS_BCL   float64 `json:"hss_bcl,omitempty"`
	HSS_BCU   float64 `json:"hss_bcu,omitempty"`
	ODDS      float64 `json:"odds,omitempty"`
	ODDS_NCL  float64 `json:"odds_ncl,omitempty"`
	ODDS_NCU  float64 `json:"odds_ncu,omitempty"`
	ODDS_BCL  float64 `json:"odds_bcl,omitempty"`
	ODDS_BCU  float64 `json:"odds_bcu,omitempty"`
	LODDS     float64 `json:"lodds,omitempty"`
	LODDS_NCL float64 `json:"lodds_ncl,omitempty"`
	LODDS_NCU float64 `json:"lodds_ncu,omitempty"`
	LODDS_BCL float64 `json:"lodds_bcl,omitempty"`
	LODDS_BCU float64 `json:"lodds_bcu,omitempty"`
	ORSS      float64 `json:"orss,omitempty"`
	ORSS_NCL  float64 `json:"orss_ncl,omitempty"`
	ORSS_NCU  float64 `json:"orss_ncu,omitempty"`
	ORSS_BCL  float64 `json:"orss_bcl,omitempty"`
	ORSS_BCU  float64 `json:"orss_bcu,omitempty"`
	EDS       float64 `json:"eds,omitempty"`
	EDS_NCL   float64 `json:"eds_ncl,omitempty"`
	EDS_NCU   float64 `json:"eds_ncu,omitempty"`
	EDS_BCL   float64 `json:"eds_bcl,omitempty"`
	EDS_BCU   float64 `json:"eds_bcu,omitempty"`
	SEDS      float64 `json:"seds,omitempty"`
	SEDS_NCL  float64 `json:"seds_ncl,omitempty"`
	SEDS_NCU  float64 `json:"seds_ncu,omitempty"`
	SEDS_BCL  float64 `json:"seds_bcl,omitempty"`
	SEDS_BCU  float64 `json:"seds_bcu,omitempty"`
	EDI       float64 `json:"edi,omitempty"`
	EDI_NCL   float64 `json:"edi_ncl,omitempty"`
	EDI_NCU   float64 `json:"edi_ncu,omitempty"`
	EDI_BCL   float64 `json:"edi_bcl,omitempty"`
	EDI_BCU   float64 `json:"edi_bcu,omitempty"`
	SEDI      float64 `json:"sedi,omitempty"`
	SEDI_NCL  float64 `json:"sedi_ncl,omitempty"`
	SEDI_NCU  float64 `json:"sedi_ncu,omitempty"`
	SEDI_BCL  float64 `json:"sedi_bcl,omitempty"`
	SEDI_BCU  float64 `json:"sedi_bcu,omitempty"`
	BAGSS     float64 `json:"bagss,omitempty"`
	BAGSS_BCL float64 `json:"bagss_bcl,omitempty"`
	BAGSS_BCU float64 `json:"bagss_bcu,omitempty"`
}

type STAT_PCT struct {
	TOTAL  int                    `json:"total,omitempty"`
	THRESH map[string]interface{} `json:"thresh,omitempty"`
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

type STAT_VCNT struct {
	TOTAL                int     `json:"total,omitempty"`
	FBAR                 float64 `json:"fbar,omitempty"`
	FBAR_BCL             float64 `json:"fbar_bcl,omitempty"`
	FBAR_BCU             float64 `json:"fbar_bcu,omitempty"`
	OBAR                 float64 `json:"obar,omitempty"`
	OBAR_BCL             float64 `json:"obar_bcl,omitempty"`
	OBAR_BCU             float64 `json:"obar_bcu,omitempty"`
	FS_RMS               float64 `json:"fs_rms,omitempty"`
	FS_RMS_BCL           float64 `json:"fs_rms_bcl,omitempty"`
	FS_RMS_BCU           float64 `json:"fs_rms_bcu,omitempty"`
	OS_RMS               float64 `json:"os_rms,omitempty"`
	OS_RMS_BCL           float64 `json:"os_rms_bcl,omitempty"`
	OS_RMS_BCU           float64 `json:"os_rms_bcu,omitempty"`
	MSVE                 float64 `json:"msve,omitempty"`
	MSVE_BCL             float64 `json:"msve_bcl,omitempty"`
	MSVE_BCU             float64 `json:"msve_bcu,omitempty"`
	RMSVE                float64 `json:"rmsve,omitempty"`
	RMSVE_BCL            float64 `json:"rmsve_bcl,omitempty"`
	RMSVE_BCU            float64 `json:"rmsve_bcu,omitempty"`
	FSTDEV               float64 `json:"fstdev,omitempty"`
	FSTDEV_BCL           float64 `json:"fstdev_bcl,omitempty"`
	FSTDEV_BCU           float64 `json:"fstdev_bcu,omitempty"`
	OSTDEV               float64 `json:"ostdev,omitempty"`
	OSTDEV_BCL           float64 `json:"ostdev_bcl,omitempty"`
	OSTDEV_BCU           float64 `json:"ostdev_bcu,omitempty"`
	FDIR                 float64 `json:"fdir,omitempty"`
	FDIR_BCL             float64 `json:"fdir_bcl,omitempty"`
	FDIR_BCU             float64 `json:"fdir_bcu,omitempty"`
	ODIR                 float64 `json:"odir,omitempty"`
	ODIR_BCL             float64 `json:"odir_bcl,omitempty"`
	ODIR_BCU             float64 `json:"odir_bcu,omitempty"`
	FBAR_SPEED           float64 `json:"fbar_speed,omitempty"`
	FBAR_SPEED_BCL       float64 `json:"fbar_speed_bcl,omitempty"`
	FBAR_SPEED_BCU       float64 `json:"fbar_speed_bcu,omitempty"`
	OBAR_SPEED           float64 `json:"obar_speed,omitempty"`
	OBAR_SPEED_BCL       float64 `json:"obar_speed_bcl,omitempty"`
	OBAR_SPEED_BCU       float64 `json:"obar_speed_bcu,omitempty"`
	VDIFF_SPEED          float64 `json:"vdiff_speed,omitempty"`
	VDIFF_SPEED_BCL      float64 `json:"vdiff_speed_bcl,omitempty"`
	VDIFF_SPEED_BCU      float64 `json:"vdiff_speed_bcu,omitempty"`
	VDIFF_DIR            float64 `json:"vdiff_dir,omitempty"`
	VDIFF_DIR_BCL        float64 `json:"vdiff_dir_bcl,omitempty"`
	VDIFF_DIR_BCU        float64 `json:"vdiff_dir_bcu,omitempty"`
	SPEED_ERR            float64 `json:"speed_err,omitempty"`
	SPEED_ERR_BCL        float64 `json:"speed_err_bcl,omitempty"`
	SPEED_ERR_BCU        float64 `json:"speed_err_bcu,omitempty"`
	SPEED_ABSERR         float64 `json:"speed_abserr,omitempty"`
	SPEED_ABSERR_BCL     float64 `json:"speed_abserr_bcl,omitempty"`
	SPEED_ABSERR_BCU     float64 `json:"speed_abserr_bcu,omitempty"`
	DIR_ERR              float64 `json:"dir_err,omitempty"`
	DIR_ERR_BCL          float64 `json:"dir_err_bcl,omitempty"`
	DIR_ERR_BCU          float64 `json:"dir_err_bcu,omitempty"`
	DIR_ABSERR           float64 `json:"dir_abserr,omitempty"`
	DIR_ABSERR_BCL       float64 `json:"dir_abserr_bcl,omitempty"`
	DIR_ABSERR_BCU       float64 `json:"dir_abserr_bcu,omitempty"`
	ANOM_CORR            float64 `json:"anom_corr,omitempty"`
	ANOM_CORR_NCL        float64 `json:"anom_corr_ncl,omitempty"`
	ANOM_CORR_NCU        float64 `json:"anom_corr_ncu,omitempty"`
	ANOM_CORR_BCL        float64 `json:"anom_corr_bcl,omitempty"`
	ANOM_CORR_BCU        float64 `json:"anom_corr_bcu,omitempty"`
	ANOM_CORR_UNCNTR     float64 `json:"anom_corr_uncntr,omitempty"`
	ANOM_CORR_UNCNTR_BCL float64 `json:"anom_corr_uncntr_bcl,omitempty"`
	ANOM_CORR_UNCNTR_BCU float64 `json:"anom_corr_uncntr_bcu,omitempty"`
}

// fillStructure functions
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
		s.LEVEL = fields[2]
	}
	i++
	if i <= dataLen {
		s.WATCH_WARN = fields[3]
	}
	i++
	if i <= dataLen {
		s.INITIALS = fields[4]
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
		s.ARADP = fields[50]
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
		s.STORM_ID = fields[2]
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
		s.AGEN_INIT = fields[5]
	}
	i++
	if i <= dataLen {
		s.AGEN_FHR = fields[6]
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
		s.GEN_TDIFF = fields[14]
	}
	i++
	if i <= dataLen {
		s.INIT_TDIFF = fields[15]
	}
	i++
	if i <= dataLen {
		s.DEV_CAT = fields[16]
	}
	i++
	if i <= dataLen {
		s.OPS_CAT = fields[17]
	}
}

func (s *STAT_SSIDX) fill_STAT_SSIDX(fields []string) {
	dataLen := len(fields) - 1
	i := -1
	i++
	if i <= dataLen {
		s.FCST_MODEL = fields[0]
	}
	i++
	if i <= dataLen {
		s.REF_MODEL = fields[1]
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
		s.INITIALS = fields[4]
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
		s.BLEVEL_BEG = fields[18]
	}
	i++
	if i <= dataLen {
		s.BLEVEL_END = fields[19]
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
		s.TRACK_SOURCE = fields[3]
	}
	i++
	if i <= dataLen {
		s.FIELD_SOURCE = fields[4]
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

func (s *MODE_CTS) fill_MODE_CTS(fields []string) {
	dataLen := len(fields) - 1
	i := -1
	i++
	if i <= dataLen {
		s.FIELD = fields[0]
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
		s.OBS_SID = fields[2]
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
		s.OBS_QC = fields[9]
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

func (s *STAT_SEEPS_MPR) fill_STAT_SEEPS_MPR(fields []string) {
	dataLen := len(fields) - 1
	i := -1
	i++
	if i <= dataLen {
		s.OBS_SID = fields[0]
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
		s.OBS_QC = fields[5]
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

func (s *MODE_OBJ) fill_MODE_OBJ(fields []string) {
	dataLen := len(fields) - 1
	i := -1
	i++
	if i <= dataLen {
		s.OBJECT_ID = fields[0]
	}
	i++
	if i <= dataLen {
		s.OBJECT_CAT = fields[1]
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
		s.OBS_SID = fields[2]
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
		s.OBS_QC = fields[13]
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

func (s *STAT_SEEPS) fill_STAT_SEEPS(fields []string) {
	dataLen := len(fields) - 1
	i := -1
	i++
	if i <= dataLen {
		s.TOTAL, _ = strconv.Atoi(fields[0])
	}
	i++
	if i <= dataLen {
		s.S12 = fields[1]
	}
	i++
	if i <= dataLen {
		s.S13 = fields[2]
	}
	i++
	if i <= dataLen {
		s.S21 = fields[3]
	}
	i++
	if i <= dataLen {
		s.S23 = fields[4]
	}
	i++
	if i <= dataLen {
		s.S31 = fields[5]
	}
	i++
	if i <= dataLen {
		s.S32 = fields[6]
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
