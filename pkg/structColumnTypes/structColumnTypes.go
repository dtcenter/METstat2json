package structColumnTypes

import (
	"errors"
	"fmt"
	"strconv"
)

/*
THIS CODE IS AUTOMATICALLY GENERATED - DO NOT EDIT THIS CODE
To modify this code - modify the buildHeaderLineTypes.go file and run the buildHeaderLineTypes.go program
cd  <repo path>/metlinetypes/pkg/buildHeaderLineTypes
go run . > /tmp/types.go
cp /tmp/types.go ../structColumnTypes/structColumnTypes.go
*/

const DOC_NOT_FOUND = "Document not found:"

// Header struct definitions
type STAT_ISC_header struct {
	VERSION        string `json:"version"`
	MODEL          string `json:"model"`
	DESC           string `json:"desc"`
	FCST_VALID_BEG string `json:"fcst_valid_beg"`
	FCST_VALID_END string `json:"fcst_valid_end"`
	OBS_LEAD       string `json:"obs_lead"`
	OBS_VALID_BEG  string `json:"obs_valid_beg"`
	OBS_VALID_END  string `json:"obs_valid_end"`
	FCST_VAR       string `json:"fcst_var"`
	FCST_UNITS     string `json:"fcst_units"`
	FCST_LEV       string `json:"fcst_lev"`
	OBS_VAR        string `json:"obs_var"`
	OBS_UNITS      string `json:"obs_units"`
	OBS_LEV        string `json:"obs_lev"`
	OBTYPE         string `json:"obtype"`
	VX_MASK        string `json:"vx_mask"`
	INTERP_MTHD    string `json:"interp_mthd"`
	INTERP_PNTS    string `json:"interp_pnts"`
	FCST_THRESH    string `json:"fcst_thresh"`
	OBS_THRESH     string `json:"obs_thresh"`
	COV_THRESH     string `json:"cov_thresh"`
	ALPHA          string `json:"alpha"`
	LINE_TYPE      string `json:"line_type"`
}

type STAT_NBRCNT_header struct {
	VERSION        string `json:"version"`
	MODEL          string `json:"model"`
	DESC           string `json:"desc"`
	FCST_VALID_BEG string `json:"fcst_valid_beg"`
	FCST_VALID_END string `json:"fcst_valid_end"`
	OBS_LEAD       string `json:"obs_lead"`
	OBS_VALID_BEG  string `json:"obs_valid_beg"`
	OBS_VALID_END  string `json:"obs_valid_end"`
	FCST_VAR       string `json:"fcst_var"`
	FCST_UNITS     string `json:"fcst_units"`
	FCST_LEV       string `json:"fcst_lev"`
	OBS_VAR        string `json:"obs_var"`
	OBS_UNITS      string `json:"obs_units"`
	OBS_LEV        string `json:"obs_lev"`
	OBTYPE         string `json:"obtype"`
	VX_MASK        string `json:"vx_mask"`
	INTERP_MTHD    string `json:"interp_mthd"`
	INTERP_PNTS    string `json:"interp_pnts"`
	FCST_THRESH    string `json:"fcst_thresh"`
	OBS_THRESH     string `json:"obs_thresh"`
	COV_THRESH     string `json:"cov_thresh"`
	ALPHA          string `json:"alpha"`
	LINE_TYPE      string `json:"line_type"`
}

type STAT_ORANK_header struct {
	VERSION        string `json:"version"`
	MODEL          string `json:"model"`
	DESC           string `json:"desc"`
	FCST_VALID_BEG string `json:"fcst_valid_beg"`
	FCST_VALID_END string `json:"fcst_valid_end"`
	OBS_LEAD       string `json:"obs_lead"`
	OBS_VALID_BEG  string `json:"obs_valid_beg"`
	OBS_VALID_END  string `json:"obs_valid_end"`
	FCST_VAR       string `json:"fcst_var"`
	FCST_UNITS     string `json:"fcst_units"`
	FCST_LEV       string `json:"fcst_lev"`
	OBS_VAR        string `json:"obs_var"`
	OBS_UNITS      string `json:"obs_units"`
	OBS_LEV        string `json:"obs_lev"`
	OBTYPE         string `json:"obtype"`
	VX_MASK        string `json:"vx_mask"`
	INTERP_MTHD    string `json:"interp_mthd"`
	INTERP_PNTS    string `json:"interp_pnts"`
	FCST_THRESH    string `json:"fcst_thresh"`
	OBS_THRESH     string `json:"obs_thresh"`
	COV_THRESH     string `json:"cov_thresh"`
	ALPHA          string `json:"alpha"`
	LINE_TYPE      string `json:"line_type"`
}

type STAT_SSIDX_header struct {
	VERSION        string `json:"version"`
	MODEL          string `json:"model"`
	DESC           string `json:"desc"`
	FCST_VALID_BEG string `json:"fcst_valid_beg"`
	FCST_VALID_END string `json:"fcst_valid_end"`
	OBS_LEAD       string `json:"obs_lead"`
	OBS_VALID_BEG  string `json:"obs_valid_beg"`
	OBS_VALID_END  string `json:"obs_valid_end"`
	FCST_VAR       string `json:"fcst_var"`
	FCST_UNITS     string `json:"fcst_units"`
	FCST_LEV       string `json:"fcst_lev"`
	OBS_VAR        string `json:"obs_var"`
	OBS_UNITS      string `json:"obs_units"`
	OBS_LEV        string `json:"obs_lev"`
	OBTYPE         string `json:"obtype"`
	VX_MASK        string `json:"vx_mask"`
	INTERP_MTHD    string `json:"interp_mthd"`
	INTERP_PNTS    string `json:"interp_pnts"`
	FCST_THRESH    string `json:"fcst_thresh"`
	OBS_THRESH     string `json:"obs_thresh"`
	COV_THRESH     string `json:"cov_thresh"`
	ALPHA          string `json:"alpha"`
	LINE_TYPE      string `json:"line_type"`
}

type STAT_CNT_header struct {
	VERSION        string `json:"version"`
	MODEL          string `json:"model"`
	DESC           string `json:"desc"`
	FCST_VALID_BEG string `json:"fcst_valid_beg"`
	FCST_VALID_END string `json:"fcst_valid_end"`
	OBS_LEAD       string `json:"obs_lead"`
	OBS_VALID_BEG  string `json:"obs_valid_beg"`
	OBS_VALID_END  string `json:"obs_valid_end"`
	FCST_VAR       string `json:"fcst_var"`
	FCST_UNITS     string `json:"fcst_units"`
	FCST_LEV       string `json:"fcst_lev"`
	OBS_VAR        string `json:"obs_var"`
	OBS_UNITS      string `json:"obs_units"`
	OBS_LEV        string `json:"obs_lev"`
	OBTYPE         string `json:"obtype"`
	VX_MASK        string `json:"vx_mask"`
	INTERP_MTHD    string `json:"interp_mthd"`
	INTERP_PNTS    string `json:"interp_pnts"`
	FCST_THRESH    string `json:"fcst_thresh"`
	OBS_THRESH     string `json:"obs_thresh"`
	COV_THRESH     string `json:"cov_thresh"`
	ALPHA          string `json:"alpha"`
	LINE_TYPE      string `json:"line_type"`
}

type STAT_GRAD_header struct {
	VERSION        string `json:"version"`
	MODEL          string `json:"model"`
	DESC           string `json:"desc"`
	FCST_VALID_BEG string `json:"fcst_valid_beg"`
	FCST_VALID_END string `json:"fcst_valid_end"`
	OBS_LEAD       string `json:"obs_lead"`
	OBS_VALID_BEG  string `json:"obs_valid_beg"`
	OBS_VALID_END  string `json:"obs_valid_end"`
	FCST_VAR       string `json:"fcst_var"`
	FCST_UNITS     string `json:"fcst_units"`
	FCST_LEV       string `json:"fcst_lev"`
	OBS_VAR        string `json:"obs_var"`
	OBS_UNITS      string `json:"obs_units"`
	OBS_LEV        string `json:"obs_lev"`
	OBTYPE         string `json:"obtype"`
	VX_MASK        string `json:"vx_mask"`
	INTERP_MTHD    string `json:"interp_mthd"`
	INTERP_PNTS    string `json:"interp_pnts"`
	FCST_THRESH    string `json:"fcst_thresh"`
	OBS_THRESH     string `json:"obs_thresh"`
	COV_THRESH     string `json:"cov_thresh"`
	ALPHA          string `json:"alpha"`
	LINE_TYPE      string `json:"line_type"`
}

type STAT_PHIST_header struct {
	VERSION        string `json:"version"`
	MODEL          string `json:"model"`
	DESC           string `json:"desc"`
	FCST_VALID_BEG string `json:"fcst_valid_beg"`
	FCST_VALID_END string `json:"fcst_valid_end"`
	OBS_LEAD       string `json:"obs_lead"`
	OBS_VALID_BEG  string `json:"obs_valid_beg"`
	OBS_VALID_END  string `json:"obs_valid_end"`
	FCST_VAR       string `json:"fcst_var"`
	FCST_UNITS     string `json:"fcst_units"`
	FCST_LEV       string `json:"fcst_lev"`
	OBS_VAR        string `json:"obs_var"`
	OBS_UNITS      string `json:"obs_units"`
	OBS_LEV        string `json:"obs_lev"`
	OBTYPE         string `json:"obtype"`
	VX_MASK        string `json:"vx_mask"`
	INTERP_MTHD    string `json:"interp_mthd"`
	INTERP_PNTS    string `json:"interp_pnts"`
	FCST_THRESH    string `json:"fcst_thresh"`
	OBS_THRESH     string `json:"obs_thresh"`
	COV_THRESH     string `json:"cov_thresh"`
	ALPHA          string `json:"alpha"`
	LINE_TYPE      string `json:"line_type"`
}

type STAT_SAL1L2_header struct {
	VERSION        string `json:"version"`
	MODEL          string `json:"model"`
	DESC           string `json:"desc"`
	FCST_VALID_BEG string `json:"fcst_valid_beg"`
	FCST_VALID_END string `json:"fcst_valid_end"`
	OBS_LEAD       string `json:"obs_lead"`
	OBS_VALID_BEG  string `json:"obs_valid_beg"`
	OBS_VALID_END  string `json:"obs_valid_end"`
	FCST_VAR       string `json:"fcst_var"`
	FCST_UNITS     string `json:"fcst_units"`
	FCST_LEV       string `json:"fcst_lev"`
	OBS_VAR        string `json:"obs_var"`
	OBS_UNITS      string `json:"obs_units"`
	OBS_LEV        string `json:"obs_lev"`
	OBTYPE         string `json:"obtype"`
	VX_MASK        string `json:"vx_mask"`
	INTERP_MTHD    string `json:"interp_mthd"`
	INTERP_PNTS    string `json:"interp_pnts"`
	FCST_THRESH    string `json:"fcst_thresh"`
	OBS_THRESH     string `json:"obs_thresh"`
	COV_THRESH     string `json:"cov_thresh"`
	ALPHA          string `json:"alpha"`
	LINE_TYPE      string `json:"line_type"`
}

type STAT_RHIST_header struct {
	VERSION        string `json:"version"`
	MODEL          string `json:"model"`
	DESC           string `json:"desc"`
	FCST_VALID_BEG string `json:"fcst_valid_beg"`
	FCST_VALID_END string `json:"fcst_valid_end"`
	OBS_LEAD       string `json:"obs_lead"`
	OBS_VALID_BEG  string `json:"obs_valid_beg"`
	OBS_VALID_END  string `json:"obs_valid_end"`
	FCST_VAR       string `json:"fcst_var"`
	FCST_UNITS     string `json:"fcst_units"`
	FCST_LEV       string `json:"fcst_lev"`
	OBS_VAR        string `json:"obs_var"`
	OBS_UNITS      string `json:"obs_units"`
	OBS_LEV        string `json:"obs_lev"`
	OBTYPE         string `json:"obtype"`
	VX_MASK        string `json:"vx_mask"`
	INTERP_MTHD    string `json:"interp_mthd"`
	INTERP_PNTS    string `json:"interp_pnts"`
	FCST_THRESH    string `json:"fcst_thresh"`
	OBS_THRESH     string `json:"obs_thresh"`
	COV_THRESH     string `json:"cov_thresh"`
	ALPHA          string `json:"alpha"`
	LINE_TYPE      string `json:"line_type"`
}

type MTD_3DSINGLE_header struct {
	VERSION    string `json:"version"`
	MODEL      string `json:"model"`
	DESC       string `json:"desc"`
	FCST_VALID string `json:"fcst_valid"`
	OBS_LEAD   string `json:"obs_lead"`
	OBS_VALID  string `json:"obs_valid"`
	T_DELTA    string `json:"t_delta"`
	FCST_T_BEG string `json:"fcst_t_beg"`
	FCST_T_END string `json:"fcst_t_end"`
	FCST_RAD   string `json:"fcst_rad"`
	FCST_THR   string `json:"fcst_thr"`
	OBS_T_BEG  string `json:"obs_t_beg"`
	OBS_T_END  string `json:"obs_t_end"`
	OBS_RAD    string `json:"obs_rad"`
	OBS_THR    string `json:"obs_thr"`
	FCST_VAR   string `json:"fcst_var"`
	FCST_UNITS string `json:"fcst_units"`
	OBS_VAR    string `json:"obs_var"`
	OBS_UNITS  string `json:"obs_units"`
	OBS_LEV    string `json:"obs_lev"`
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
	INIT       string `json:"init"`
	LEAD       string `json:"lead"`
	VALID      string `json:"valid"`
	INIT_MASK  string `json:"init_mask"`
	VALID_MASK string `json:"valid_mask"`
	LINE_TYPE  string `json:"line_type"`
}

type STAT_SEEPS_header struct {
	VERSION        string `json:"version"`
	MODEL          string `json:"model"`
	DESC           string `json:"desc"`
	FCST_VALID_BEG string `json:"fcst_valid_beg"`
	FCST_VALID_END string `json:"fcst_valid_end"`
	OBS_LEAD       string `json:"obs_lead"`
	OBS_VALID_BEG  string `json:"obs_valid_beg"`
	OBS_VALID_END  string `json:"obs_valid_end"`
	FCST_VAR       string `json:"fcst_var"`
	FCST_UNITS     string `json:"fcst_units"`
	FCST_LEV       string `json:"fcst_lev"`
	OBS_VAR        string `json:"obs_var"`
	OBS_UNITS      string `json:"obs_units"`
	OBS_LEV        string `json:"obs_lev"`
	OBTYPE         string `json:"obtype"`
	VX_MASK        string `json:"vx_mask"`
	INTERP_MTHD    string `json:"interp_mthd"`
	INTERP_PNTS    string `json:"interp_pnts"`
	FCST_THRESH    string `json:"fcst_thresh"`
	OBS_THRESH     string `json:"obs_thresh"`
	COV_THRESH     string `json:"cov_thresh"`
	ALPHA          string `json:"alpha"`
	LINE_TYPE      string `json:"line_type"`
}

type STAT_PJC_header struct {
	VERSION        string `json:"version"`
	MODEL          string `json:"model"`
	DESC           string `json:"desc"`
	FCST_VALID_BEG string `json:"fcst_valid_beg"`
	FCST_VALID_END string `json:"fcst_valid_end"`
	OBS_LEAD       string `json:"obs_lead"`
	OBS_VALID_BEG  string `json:"obs_valid_beg"`
	OBS_VALID_END  string `json:"obs_valid_end"`
	FCST_VAR       string `json:"fcst_var"`
	FCST_UNITS     string `json:"fcst_units"`
	FCST_LEV       string `json:"fcst_lev"`
	OBS_VAR        string `json:"obs_var"`
	OBS_UNITS      string `json:"obs_units"`
	OBS_LEV        string `json:"obs_lev"`
	OBTYPE         string `json:"obtype"`
	VX_MASK        string `json:"vx_mask"`
	INTERP_MTHD    string `json:"interp_mthd"`
	INTERP_PNTS    string `json:"interp_pnts"`
	FCST_THRESH    string `json:"fcst_thresh"`
	OBS_THRESH     string `json:"obs_thresh"`
	COV_THRESH     string `json:"cov_thresh"`
	ALPHA          string `json:"alpha"`
	LINE_TYPE      string `json:"line_type"`
}

type STAT_PSTD_header struct {
	VERSION        string `json:"version"`
	MODEL          string `json:"model"`
	DESC           string `json:"desc"`
	FCST_VALID_BEG string `json:"fcst_valid_beg"`
	FCST_VALID_END string `json:"fcst_valid_end"`
	OBS_LEAD       string `json:"obs_lead"`
	OBS_VALID_BEG  string `json:"obs_valid_beg"`
	OBS_VALID_END  string `json:"obs_valid_end"`
	FCST_VAR       string `json:"fcst_var"`
	FCST_UNITS     string `json:"fcst_units"`
	FCST_LEV       string `json:"fcst_lev"`
	OBS_VAR        string `json:"obs_var"`
	OBS_UNITS      string `json:"obs_units"`
	OBS_LEV        string `json:"obs_lev"`
	OBTYPE         string `json:"obtype"`
	VX_MASK        string `json:"vx_mask"`
	INTERP_MTHD    string `json:"interp_mthd"`
	INTERP_PNTS    string `json:"interp_pnts"`
	FCST_THRESH    string `json:"fcst_thresh"`
	OBS_THRESH     string `json:"obs_thresh"`
	COV_THRESH     string `json:"cov_thresh"`
	ALPHA          string `json:"alpha"`
	LINE_TYPE      string `json:"line_type"`
}

type STAT_RPS_header struct {
	VERSION        string `json:"version"`
	MODEL          string `json:"model"`
	DESC           string `json:"desc"`
	FCST_VALID_BEG string `json:"fcst_valid_beg"`
	FCST_VALID_END string `json:"fcst_valid_end"`
	OBS_LEAD       string `json:"obs_lead"`
	OBS_VALID_BEG  string `json:"obs_valid_beg"`
	OBS_VALID_END  string `json:"obs_valid_end"`
	FCST_VAR       string `json:"fcst_var"`
	FCST_UNITS     string `json:"fcst_units"`
	FCST_LEV       string `json:"fcst_lev"`
	OBS_VAR        string `json:"obs_var"`
	OBS_UNITS      string `json:"obs_units"`
	OBS_LEV        string `json:"obs_lev"`
	OBTYPE         string `json:"obtype"`
	VX_MASK        string `json:"vx_mask"`
	INTERP_MTHD    string `json:"interp_mthd"`
	INTERP_PNTS    string `json:"interp_pnts"`
	FCST_THRESH    string `json:"fcst_thresh"`
	OBS_THRESH     string `json:"obs_thresh"`
	COV_THRESH     string `json:"cov_thresh"`
	ALPHA          string `json:"alpha"`
	LINE_TYPE      string `json:"line_type"`
}

type STAT_SSVAR_header struct {
	VERSION        string `json:"version"`
	MODEL          string `json:"model"`
	DESC           string `json:"desc"`
	FCST_VALID_BEG string `json:"fcst_valid_beg"`
	FCST_VALID_END string `json:"fcst_valid_end"`
	OBS_LEAD       string `json:"obs_lead"`
	OBS_VALID_BEG  string `json:"obs_valid_beg"`
	OBS_VALID_END  string `json:"obs_valid_end"`
	FCST_VAR       string `json:"fcst_var"`
	FCST_UNITS     string `json:"fcst_units"`
	FCST_LEV       string `json:"fcst_lev"`
	OBS_VAR        string `json:"obs_var"`
	OBS_UNITS      string `json:"obs_units"`
	OBS_LEV        string `json:"obs_lev"`
	OBTYPE         string `json:"obtype"`
	VX_MASK        string `json:"vx_mask"`
	INTERP_MTHD    string `json:"interp_mthd"`
	INTERP_PNTS    string `json:"interp_pnts"`
	FCST_THRESH    string `json:"fcst_thresh"`
	OBS_THRESH     string `json:"obs_thresh"`
	COV_THRESH     string `json:"cov_thresh"`
	ALPHA          string `json:"alpha"`
	LINE_TYPE      string `json:"line_type"`
}

type MTD_3DPAIR_header struct {
	VERSION    string `json:"version"`
	MODEL      string `json:"model"`
	DESC       string `json:"desc"`
	FCST_VALID string `json:"fcst_valid"`
	OBS_LEAD   string `json:"obs_lead"`
	OBS_VALID  string `json:"obs_valid"`
	T_DELTA    string `json:"t_delta"`
	FCST_T_BEG string `json:"fcst_t_beg"`
	FCST_T_END string `json:"fcst_t_end"`
	FCST_RAD   string `json:"fcst_rad"`
	FCST_THR   string `json:"fcst_thr"`
	OBS_T_BEG  string `json:"obs_t_beg"`
	OBS_T_END  string `json:"obs_t_end"`
	OBS_RAD    string `json:"obs_rad"`
	OBS_THR    string `json:"obs_thr"`
	FCST_VAR   string `json:"fcst_var"`
	FCST_UNITS string `json:"fcst_units"`
	OBS_VAR    string `json:"obs_var"`
	OBS_UNITS  string `json:"obs_units"`
	OBS_LEV    string `json:"obs_lev"`
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
	INIT       string `json:"init"`
	LEAD       string `json:"lead"`
	VALID      string `json:"valid"`
	INIT_MASK  string `json:"init_mask"`
	VALID_MASK string `json:"valid_mask"`
	LINE_TYPE  string `json:"line_type"`
}

type STAT_NBRCTC_header struct {
	VERSION        string `json:"version"`
	MODEL          string `json:"model"`
	DESC           string `json:"desc"`
	FCST_VALID_BEG string `json:"fcst_valid_beg"`
	FCST_VALID_END string `json:"fcst_valid_end"`
	OBS_LEAD       string `json:"obs_lead"`
	OBS_VALID_BEG  string `json:"obs_valid_beg"`
	OBS_VALID_END  string `json:"obs_valid_end"`
	FCST_VAR       string `json:"fcst_var"`
	FCST_UNITS     string `json:"fcst_units"`
	FCST_LEV       string `json:"fcst_lev"`
	OBS_VAR        string `json:"obs_var"`
	OBS_UNITS      string `json:"obs_units"`
	OBS_LEV        string `json:"obs_lev"`
	OBTYPE         string `json:"obtype"`
	VX_MASK        string `json:"vx_mask"`
	INTERP_MTHD    string `json:"interp_mthd"`
	INTERP_PNTS    string `json:"interp_pnts"`
	FCST_THRESH    string `json:"fcst_thresh"`
	OBS_THRESH     string `json:"obs_thresh"`
	COV_THRESH     string `json:"cov_thresh"`
	ALPHA          string `json:"alpha"`
	LINE_TYPE      string `json:"line_type"`
}

type STAT_ECNT_header struct {
	VERSION        string `json:"version"`
	MODEL          string `json:"model"`
	DESC           string `json:"desc"`
	FCST_VALID_BEG string `json:"fcst_valid_beg"`
	FCST_VALID_END string `json:"fcst_valid_end"`
	OBS_LEAD       string `json:"obs_lead"`
	OBS_VALID_BEG  string `json:"obs_valid_beg"`
	OBS_VALID_END  string `json:"obs_valid_end"`
	FCST_VAR       string `json:"fcst_var"`
	FCST_UNITS     string `json:"fcst_units"`
	FCST_LEV       string `json:"fcst_lev"`
	OBS_VAR        string `json:"obs_var"`
	OBS_UNITS      string `json:"obs_units"`
	OBS_LEV        string `json:"obs_lev"`
	OBTYPE         string `json:"obtype"`
	VX_MASK        string `json:"vx_mask"`
	INTERP_MTHD    string `json:"interp_mthd"`
	INTERP_PNTS    string `json:"interp_pnts"`
	FCST_THRESH    string `json:"fcst_thresh"`
	OBS_THRESH     string `json:"obs_thresh"`
	COV_THRESH     string `json:"cov_thresh"`
	ALPHA          string `json:"alpha"`
	LINE_TYPE      string `json:"line_type"`
}

type STAT_RELP_header struct {
	VERSION        string `json:"version"`
	MODEL          string `json:"model"`
	DESC           string `json:"desc"`
	FCST_VALID_BEG string `json:"fcst_valid_beg"`
	FCST_VALID_END string `json:"fcst_valid_end"`
	OBS_LEAD       string `json:"obs_lead"`
	OBS_VALID_BEG  string `json:"obs_valid_beg"`
	OBS_VALID_END  string `json:"obs_valid_end"`
	FCST_VAR       string `json:"fcst_var"`
	FCST_UNITS     string `json:"fcst_units"`
	FCST_LEV       string `json:"fcst_lev"`
	OBS_VAR        string `json:"obs_var"`
	OBS_UNITS      string `json:"obs_units"`
	OBS_LEV        string `json:"obs_lev"`
	OBTYPE         string `json:"obtype"`
	VX_MASK        string `json:"vx_mask"`
	INTERP_MTHD    string `json:"interp_mthd"`
	INTERP_PNTS    string `json:"interp_pnts"`
	FCST_THRESH    string `json:"fcst_thresh"`
	OBS_THRESH     string `json:"obs_thresh"`
	COV_THRESH     string `json:"cov_thresh"`
	ALPHA          string `json:"alpha"`
	LINE_TYPE      string `json:"line_type"`
}

type STAT_VL1L2_header struct {
	VERSION        string `json:"version"`
	MODEL          string `json:"model"`
	DESC           string `json:"desc"`
	FCST_VALID_BEG string `json:"fcst_valid_beg"`
	FCST_VALID_END string `json:"fcst_valid_end"`
	OBS_LEAD       string `json:"obs_lead"`
	OBS_VALID_BEG  string `json:"obs_valid_beg"`
	OBS_VALID_END  string `json:"obs_valid_end"`
	FCST_VAR       string `json:"fcst_var"`
	FCST_UNITS     string `json:"fcst_units"`
	FCST_LEV       string `json:"fcst_lev"`
	OBS_VAR        string `json:"obs_var"`
	OBS_UNITS      string `json:"obs_units"`
	OBS_LEV        string `json:"obs_lev"`
	OBTYPE         string `json:"obtype"`
	VX_MASK        string `json:"vx_mask"`
	INTERP_MTHD    string `json:"interp_mthd"`
	INTERP_PNTS    string `json:"interp_pnts"`
	FCST_THRESH    string `json:"fcst_thresh"`
	OBS_THRESH     string `json:"obs_thresh"`
	COV_THRESH     string `json:"cov_thresh"`
	ALPHA          string `json:"alpha"`
	LINE_TYPE      string `json:"line_type"`
}

type STAT_CTC_header struct {
	VERSION        string `json:"version"`
	MODEL          string `json:"model"`
	DESC           string `json:"desc"`
	FCST_VALID_BEG string `json:"fcst_valid_beg"`
	FCST_VALID_END string `json:"fcst_valid_end"`
	OBS_LEAD       string `json:"obs_lead"`
	OBS_VALID_BEG  string `json:"obs_valid_beg"`
	OBS_VALID_END  string `json:"obs_valid_end"`
	FCST_VAR       string `json:"fcst_var"`
	FCST_UNITS     string `json:"fcst_units"`
	FCST_LEV       string `json:"fcst_lev"`
	OBS_VAR        string `json:"obs_var"`
	OBS_UNITS      string `json:"obs_units"`
	OBS_LEV        string `json:"obs_lev"`
	OBTYPE         string `json:"obtype"`
	VX_MASK        string `json:"vx_mask"`
	INTERP_MTHD    string `json:"interp_mthd"`
	INTERP_PNTS    string `json:"interp_pnts"`
	FCST_THRESH    string `json:"fcst_thresh"`
	OBS_THRESH     string `json:"obs_thresh"`
	COV_THRESH     string `json:"cov_thresh"`
	ALPHA          string `json:"alpha"`
	LINE_TYPE      string `json:"line_type"`
}

type STAT_FHO_header struct {
	VERSION        string `json:"version"`
	MODEL          string `json:"model"`
	DESC           string `json:"desc"`
	FCST_VALID_BEG string `json:"fcst_valid_beg"`
	FCST_VALID_END string `json:"fcst_valid_end"`
	OBS_LEAD       string `json:"obs_lead"`
	OBS_VALID_BEG  string `json:"obs_valid_beg"`
	OBS_VALID_END  string `json:"obs_valid_end"`
	FCST_VAR       string `json:"fcst_var"`
	FCST_UNITS     string `json:"fcst_units"`
	FCST_LEV       string `json:"fcst_lev"`
	OBS_VAR        string `json:"obs_var"`
	OBS_UNITS      string `json:"obs_units"`
	OBS_LEV        string `json:"obs_lev"`
	OBTYPE         string `json:"obtype"`
	VX_MASK        string `json:"vx_mask"`
	INTERP_MTHD    string `json:"interp_mthd"`
	INTERP_PNTS    string `json:"interp_pnts"`
	FCST_THRESH    string `json:"fcst_thresh"`
	OBS_THRESH     string `json:"obs_thresh"`
	COV_THRESH     string `json:"cov_thresh"`
	ALPHA          string `json:"alpha"`
	LINE_TYPE      string `json:"line_type"`
}

type STAT_MCTC_header struct {
	VERSION        string `json:"version"`
	MODEL          string `json:"model"`
	DESC           string `json:"desc"`
	FCST_VALID_BEG string `json:"fcst_valid_beg"`
	FCST_VALID_END string `json:"fcst_valid_end"`
	OBS_LEAD       string `json:"obs_lead"`
	OBS_VALID_BEG  string `json:"obs_valid_beg"`
	OBS_VALID_END  string `json:"obs_valid_end"`
	FCST_VAR       string `json:"fcst_var"`
	FCST_UNITS     string `json:"fcst_units"`
	FCST_LEV       string `json:"fcst_lev"`
	OBS_VAR        string `json:"obs_var"`
	OBS_UNITS      string `json:"obs_units"`
	OBS_LEV        string `json:"obs_lev"`
	OBTYPE         string `json:"obtype"`
	VX_MASK        string `json:"vx_mask"`
	INTERP_MTHD    string `json:"interp_mthd"`
	INTERP_PNTS    string `json:"interp_pnts"`
	FCST_THRESH    string `json:"fcst_thresh"`
	OBS_THRESH     string `json:"obs_thresh"`
	COV_THRESH     string `json:"cov_thresh"`
	ALPHA          string `json:"alpha"`
	LINE_TYPE      string `json:"line_type"`
}

type STAT_MPR_header struct {
	VERSION        string `json:"version"`
	MODEL          string `json:"model"`
	DESC           string `json:"desc"`
	FCST_VALID_BEG string `json:"fcst_valid_beg"`
	FCST_VALID_END string `json:"fcst_valid_end"`
	OBS_LEAD       string `json:"obs_lead"`
	OBS_VALID_BEG  string `json:"obs_valid_beg"`
	OBS_VALID_END  string `json:"obs_valid_end"`
	FCST_VAR       string `json:"fcst_var"`
	FCST_UNITS     string `json:"fcst_units"`
	FCST_LEV       string `json:"fcst_lev"`
	OBS_VAR        string `json:"obs_var"`
	OBS_UNITS      string `json:"obs_units"`
	OBS_LEV        string `json:"obs_lev"`
	OBTYPE         string `json:"obtype"`
	VX_MASK        string `json:"vx_mask"`
	INTERP_MTHD    string `json:"interp_mthd"`
	INTERP_PNTS    string `json:"interp_pnts"`
	FCST_THRESH    string `json:"fcst_thresh"`
	OBS_THRESH     string `json:"obs_thresh"`
	COV_THRESH     string `json:"cov_thresh"`
	ALPHA          string `json:"alpha"`
	LINE_TYPE      string `json:"line_type"`
}

type MTD_2DSINGLE_header struct {
	VERSION    string `json:"version"`
	MODEL      string `json:"model"`
	DESC       string `json:"desc"`
	FCST_VALID string `json:"fcst_valid"`
	OBS_LEAD   string `json:"obs_lead"`
	OBS_VALID  string `json:"obs_valid"`
	T_DELTA    string `json:"t_delta"`
	FCST_T_BEG string `json:"fcst_t_beg"`
	FCST_T_END string `json:"fcst_t_end"`
	FCST_RAD   string `json:"fcst_rad"`
	FCST_THR   string `json:"fcst_thr"`
	OBS_T_BEG  string `json:"obs_t_beg"`
	OBS_T_END  string `json:"obs_t_end"`
	OBS_RAD    string `json:"obs_rad"`
	OBS_THR    string `json:"obs_thr"`
	FCST_VAR   string `json:"fcst_var"`
	FCST_UNITS string `json:"fcst_units"`
	OBS_VAR    string `json:"obs_var"`
	OBS_UNITS  string `json:"obs_units"`
	OBS_LEV    string `json:"obs_lev"`
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
	INIT       string `json:"init"`
	LEAD       string `json:"lead"`
	VALID      string `json:"valid"`
	INIT_MASK  string `json:"init_mask"`
	VALID_MASK string `json:"valid_mask"`
	LINE_TYPE  string `json:"line_type"`
}

type STAT_ECLV_header struct {
	VERSION        string `json:"version"`
	MODEL          string `json:"model"`
	DESC           string `json:"desc"`
	FCST_VALID_BEG string `json:"fcst_valid_beg"`
	FCST_VALID_END string `json:"fcst_valid_end"`
	OBS_LEAD       string `json:"obs_lead"`
	OBS_VALID_BEG  string `json:"obs_valid_beg"`
	OBS_VALID_END  string `json:"obs_valid_end"`
	FCST_VAR       string `json:"fcst_var"`
	FCST_UNITS     string `json:"fcst_units"`
	FCST_LEV       string `json:"fcst_lev"`
	OBS_VAR        string `json:"obs_var"`
	OBS_UNITS      string `json:"obs_units"`
	OBS_LEV        string `json:"obs_lev"`
	OBTYPE         string `json:"obtype"`
	VX_MASK        string `json:"vx_mask"`
	INTERP_MTHD    string `json:"interp_mthd"`
	INTERP_PNTS    string `json:"interp_pnts"`
	FCST_THRESH    string `json:"fcst_thresh"`
	OBS_THRESH     string `json:"obs_thresh"`
	COV_THRESH     string `json:"cov_thresh"`
	ALPHA          string `json:"alpha"`
	LINE_TYPE      string `json:"line_type"`
}

type STAT_GENMPR_header struct {
	VERSION        string `json:"version"`
	MODEL          string `json:"model"`
	DESC           string `json:"desc"`
	FCST_VALID_BEG string `json:"fcst_valid_beg"`
	FCST_VALID_END string `json:"fcst_valid_end"`
	OBS_LEAD       string `json:"obs_lead"`
	OBS_VALID_BEG  string `json:"obs_valid_beg"`
	OBS_VALID_END  string `json:"obs_valid_end"`
	FCST_VAR       string `json:"fcst_var"`
	FCST_UNITS     string `json:"fcst_units"`
	FCST_LEV       string `json:"fcst_lev"`
	OBS_VAR        string `json:"obs_var"`
	OBS_UNITS      string `json:"obs_units"`
	OBS_LEV        string `json:"obs_lev"`
	OBTYPE         string `json:"obtype"`
	VX_MASK        string `json:"vx_mask"`
	INTERP_MTHD    string `json:"interp_mthd"`
	INTERP_PNTS    string `json:"interp_pnts"`
	FCST_THRESH    string `json:"fcst_thresh"`
	OBS_THRESH     string `json:"obs_thresh"`
	COV_THRESH     string `json:"cov_thresh"`
	ALPHA          string `json:"alpha"`
	LINE_TYPE      string `json:"line_type"`
}

type MODE_CTS_header struct {
	VERSION    string `json:"version"`
	MODEL      string `json:"model"`
	N_VALID    string `json:"n_valid"`
	GRID_RES   string `json:"grid_res"`
	DESC       string `json:"desc"`
	FCST_VALID string `json:"fcst_valid"`
	FCST_ACCUM string `json:"fcst_accum"`
	OBS_LEAD   string `json:"obs_lead"`
	OBS_VALID  string `json:"obs_valid"`
	OBS_ACCUM  string `json:"obs_accum"`
	FCST_RAD   string `json:"fcst_rad"`
	FCST_THR   string `json:"fcst_thr"`
	OBS_RAD    string `json:"obs_rad"`
	OBS_THR    string `json:"obs_thr"`
	FCST_VAR   string `json:"fcst_var"`
	FCST_UNITS string `json:"fcst_units"`
	OBS_VAR    string `json:"obs_var"`
	OBS_UNITS  string `json:"obs_units"`
	OBS_LEV    string `json:"obs_lev"`
	OBTYPE     string `json:"obtype"`
}

type STAT_MCTS_header struct {
	VERSION        string `json:"version"`
	MODEL          string `json:"model"`
	DESC           string `json:"desc"`
	FCST_VALID_BEG string `json:"fcst_valid_beg"`
	FCST_VALID_END string `json:"fcst_valid_end"`
	OBS_LEAD       string `json:"obs_lead"`
	OBS_VALID_BEG  string `json:"obs_valid_beg"`
	OBS_VALID_END  string `json:"obs_valid_end"`
	FCST_VAR       string `json:"fcst_var"`
	FCST_UNITS     string `json:"fcst_units"`
	FCST_LEV       string `json:"fcst_lev"`
	OBS_VAR        string `json:"obs_var"`
	OBS_UNITS      string `json:"obs_units"`
	OBS_LEV        string `json:"obs_lev"`
	OBTYPE         string `json:"obtype"`
	VX_MASK        string `json:"vx_mask"`
	INTERP_MTHD    string `json:"interp_mthd"`
	INTERP_PNTS    string `json:"interp_pnts"`
	FCST_THRESH    string `json:"fcst_thresh"`
	OBS_THRESH     string `json:"obs_thresh"`
	COV_THRESH     string `json:"cov_thresh"`
	ALPHA          string `json:"alpha"`
	LINE_TYPE      string `json:"line_type"`
}

type STAT_NBRCTS_header struct {
	VERSION        string `json:"version"`
	MODEL          string `json:"model"`
	DESC           string `json:"desc"`
	FCST_VALID_BEG string `json:"fcst_valid_beg"`
	FCST_VALID_END string `json:"fcst_valid_end"`
	OBS_LEAD       string `json:"obs_lead"`
	OBS_VALID_BEG  string `json:"obs_valid_beg"`
	OBS_VALID_END  string `json:"obs_valid_end"`
	FCST_VAR       string `json:"fcst_var"`
	FCST_UNITS     string `json:"fcst_units"`
	FCST_LEV       string `json:"fcst_lev"`
	OBS_VAR        string `json:"obs_var"`
	OBS_UNITS      string `json:"obs_units"`
	OBS_LEV        string `json:"obs_lev"`
	OBTYPE         string `json:"obtype"`
	VX_MASK        string `json:"vx_mask"`
	INTERP_MTHD    string `json:"interp_mthd"`
	INTERP_PNTS    string `json:"interp_pnts"`
	FCST_THRESH    string `json:"fcst_thresh"`
	OBS_THRESH     string `json:"obs_thresh"`
	COV_THRESH     string `json:"cov_thresh"`
	ALPHA          string `json:"alpha"`
	LINE_TYPE      string `json:"line_type"`
}

type STAT_DMAP_header struct {
	VERSION        string `json:"version"`
	MODEL          string `json:"model"`
	DESC           string `json:"desc"`
	FCST_VALID_BEG string `json:"fcst_valid_beg"`
	FCST_VALID_END string `json:"fcst_valid_end"`
	OBS_LEAD       string `json:"obs_lead"`
	OBS_VALID_BEG  string `json:"obs_valid_beg"`
	OBS_VALID_END  string `json:"obs_valid_end"`
	FCST_VAR       string `json:"fcst_var"`
	FCST_UNITS     string `json:"fcst_units"`
	FCST_LEV       string `json:"fcst_lev"`
	OBS_VAR        string `json:"obs_var"`
	OBS_UNITS      string `json:"obs_units"`
	OBS_LEV        string `json:"obs_lev"`
	OBTYPE         string `json:"obtype"`
	VX_MASK        string `json:"vx_mask"`
	INTERP_MTHD    string `json:"interp_mthd"`
	INTERP_PNTS    string `json:"interp_pnts"`
	FCST_THRESH    string `json:"fcst_thresh"`
	OBS_THRESH     string `json:"obs_thresh"`
	COV_THRESH     string `json:"cov_thresh"`
	ALPHA          string `json:"alpha"`
	LINE_TYPE      string `json:"line_type"`
}

type STAT_PRC_header struct {
	VERSION        string `json:"version"`
	MODEL          string `json:"model"`
	DESC           string `json:"desc"`
	FCST_VALID_BEG string `json:"fcst_valid_beg"`
	FCST_VALID_END string `json:"fcst_valid_end"`
	OBS_LEAD       string `json:"obs_lead"`
	OBS_VALID_BEG  string `json:"obs_valid_beg"`
	OBS_VALID_END  string `json:"obs_valid_end"`
	FCST_VAR       string `json:"fcst_var"`
	FCST_UNITS     string `json:"fcst_units"`
	FCST_LEV       string `json:"fcst_lev"`
	OBS_VAR        string `json:"obs_var"`
	OBS_UNITS      string `json:"obs_units"`
	OBS_LEV        string `json:"obs_lev"`
	OBTYPE         string `json:"obtype"`
	VX_MASK        string `json:"vx_mask"`
	INTERP_MTHD    string `json:"interp_mthd"`
	INTERP_PNTS    string `json:"interp_pnts"`
	FCST_THRESH    string `json:"fcst_thresh"`
	OBS_THRESH     string `json:"obs_thresh"`
	COV_THRESH     string `json:"cov_thresh"`
	ALPHA          string `json:"alpha"`
	LINE_TYPE      string `json:"line_type"`
}

type STAT_VAL1L2_header struct {
	VERSION        string `json:"version"`
	MODEL          string `json:"model"`
	DESC           string `json:"desc"`
	FCST_VALID_BEG string `json:"fcst_valid_beg"`
	FCST_VALID_END string `json:"fcst_valid_end"`
	OBS_LEAD       string `json:"obs_lead"`
	OBS_VALID_BEG  string `json:"obs_valid_beg"`
	OBS_VALID_END  string `json:"obs_valid_end"`
	FCST_VAR       string `json:"fcst_var"`
	FCST_UNITS     string `json:"fcst_units"`
	FCST_LEV       string `json:"fcst_lev"`
	OBS_VAR        string `json:"obs_var"`
	OBS_UNITS      string `json:"obs_units"`
	OBS_LEV        string `json:"obs_lev"`
	OBTYPE         string `json:"obtype"`
	VX_MASK        string `json:"vx_mask"`
	INTERP_MTHD    string `json:"interp_mthd"`
	INTERP_PNTS    string `json:"interp_pnts"`
	FCST_THRESH    string `json:"fcst_thresh"`
	OBS_THRESH     string `json:"obs_thresh"`
	COV_THRESH     string `json:"cov_thresh"`
	ALPHA          string `json:"alpha"`
	LINE_TYPE      string `json:"line_type"`
}

type STAT_VCNT_header struct {
	VERSION        string `json:"version"`
	MODEL          string `json:"model"`
	DESC           string `json:"desc"`
	FCST_VALID_BEG string `json:"fcst_valid_beg"`
	FCST_VALID_END string `json:"fcst_valid_end"`
	OBS_LEAD       string `json:"obs_lead"`
	OBS_VALID_BEG  string `json:"obs_valid_beg"`
	OBS_VALID_END  string `json:"obs_valid_end"`
	FCST_VAR       string `json:"fcst_var"`
	FCST_UNITS     string `json:"fcst_units"`
	FCST_LEV       string `json:"fcst_lev"`
	OBS_VAR        string `json:"obs_var"`
	OBS_UNITS      string `json:"obs_units"`
	OBS_LEV        string `json:"obs_lev"`
	OBTYPE         string `json:"obtype"`
	VX_MASK        string `json:"vx_mask"`
	INTERP_MTHD    string `json:"interp_mthd"`
	INTERP_PNTS    string `json:"interp_pnts"`
	FCST_THRESH    string `json:"fcst_thresh"`
	OBS_THRESH     string `json:"obs_thresh"`
	COV_THRESH     string `json:"cov_thresh"`
	ALPHA          string `json:"alpha"`
	LINE_TYPE      string `json:"line_type"`
}

type MODE_OBJ_header struct {
	VERSION    string `json:"version"`
	MODEL      string `json:"model"`
	N_VALID    string `json:"n_valid"`
	GRID_RES   string `json:"grid_res"`
	DESC       string `json:"desc"`
	FCST_VALID string `json:"fcst_valid"`
	FCST_ACCUM string `json:"fcst_accum"`
	OBS_LEAD   string `json:"obs_lead"`
	OBS_VALID  string `json:"obs_valid"`
	OBS_ACCUM  string `json:"obs_accum"`
	FCST_RAD   string `json:"fcst_rad"`
	FCST_THR   string `json:"fcst_thr"`
	OBS_RAD    string `json:"obs_rad"`
	OBS_THR    string `json:"obs_thr"`
	FCST_VAR   string `json:"fcst_var"`
	FCST_UNITS string `json:"fcst_units"`
	OBS_VAR    string `json:"obs_var"`
	OBS_UNITS  string `json:"obs_units"`
	OBS_LEV    string `json:"obs_lev"`
	OBTYPE     string `json:"obtype"`
}

type STAT_CTS_header struct {
	VERSION        string `json:"version"`
	MODEL          string `json:"model"`
	DESC           string `json:"desc"`
	FCST_VALID_BEG string `json:"fcst_valid_beg"`
	FCST_VALID_END string `json:"fcst_valid_end"`
	OBS_LEAD       string `json:"obs_lead"`
	OBS_VALID_BEG  string `json:"obs_valid_beg"`
	OBS_VALID_END  string `json:"obs_valid_end"`
	FCST_VAR       string `json:"fcst_var"`
	FCST_UNITS     string `json:"fcst_units"`
	FCST_LEV       string `json:"fcst_lev"`
	OBS_VAR        string `json:"obs_var"`
	OBS_UNITS      string `json:"obs_units"`
	OBS_LEV        string `json:"obs_lev"`
	OBTYPE         string `json:"obtype"`
	VX_MASK        string `json:"vx_mask"`
	INTERP_MTHD    string `json:"interp_mthd"`
	INTERP_PNTS    string `json:"interp_pnts"`
	FCST_THRESH    string `json:"fcst_thresh"`
	OBS_THRESH     string `json:"obs_thresh"`
	COV_THRESH     string `json:"cov_thresh"`
	ALPHA          string `json:"alpha"`
	LINE_TYPE      string `json:"line_type"`
}

type STAT_SEEPS_MPR_header struct {
	VERSION        string `json:"version"`
	MODEL          string `json:"model"`
	DESC           string `json:"desc"`
	FCST_VALID_BEG string `json:"fcst_valid_beg"`
	FCST_VALID_END string `json:"fcst_valid_end"`
	OBS_LEAD       string `json:"obs_lead"`
	OBS_VALID_BEG  string `json:"obs_valid_beg"`
	OBS_VALID_END  string `json:"obs_valid_end"`
	FCST_VAR       string `json:"fcst_var"`
	FCST_UNITS     string `json:"fcst_units"`
	FCST_LEV       string `json:"fcst_lev"`
	OBS_VAR        string `json:"obs_var"`
	OBS_UNITS      string `json:"obs_units"`
	OBS_LEV        string `json:"obs_lev"`
	OBTYPE         string `json:"obtype"`
	VX_MASK        string `json:"vx_mask"`
	INTERP_MTHD    string `json:"interp_mthd"`
	INTERP_PNTS    string `json:"interp_pnts"`
	FCST_THRESH    string `json:"fcst_thresh"`
	OBS_THRESH     string `json:"obs_thresh"`
	COV_THRESH     string `json:"cov_thresh"`
	ALPHA          string `json:"alpha"`
	LINE_TYPE      string `json:"line_type"`
}

type STAT_PCT_header struct {
	VERSION        string `json:"version"`
	MODEL          string `json:"model"`
	DESC           string `json:"desc"`
	FCST_VALID_BEG string `json:"fcst_valid_beg"`
	FCST_VALID_END string `json:"fcst_valid_end"`
	OBS_LEAD       string `json:"obs_lead"`
	OBS_VALID_BEG  string `json:"obs_valid_beg"`
	OBS_VALID_END  string `json:"obs_valid_end"`
	FCST_VAR       string `json:"fcst_var"`
	FCST_UNITS     string `json:"fcst_units"`
	FCST_LEV       string `json:"fcst_lev"`
	OBS_VAR        string `json:"obs_var"`
	OBS_UNITS      string `json:"obs_units"`
	OBS_LEV        string `json:"obs_lev"`
	OBTYPE         string `json:"obtype"`
	VX_MASK        string `json:"vx_mask"`
	INTERP_MTHD    string `json:"interp_mthd"`
	INTERP_PNTS    string `json:"interp_pnts"`
	FCST_THRESH    string `json:"fcst_thresh"`
	OBS_THRESH     string `json:"obs_thresh"`
	COV_THRESH     string `json:"cov_thresh"`
	ALPHA          string `json:"alpha"`
	LINE_TYPE      string `json:"line_type"`
}

type STAT_SL1L2_header struct {
	VERSION        string `json:"version"`
	MODEL          string `json:"model"`
	DESC           string `json:"desc"`
	FCST_VALID_BEG string `json:"fcst_valid_beg"`
	FCST_VALID_END string `json:"fcst_valid_end"`
	OBS_LEAD       string `json:"obs_lead"`
	OBS_VALID_BEG  string `json:"obs_valid_beg"`
	OBS_VALID_END  string `json:"obs_valid_end"`
	FCST_VAR       string `json:"fcst_var"`
	FCST_UNITS     string `json:"fcst_units"`
	FCST_LEV       string `json:"fcst_lev"`
	OBS_VAR        string `json:"obs_var"`
	OBS_UNITS      string `json:"obs_units"`
	OBS_LEV        string `json:"obs_lev"`
	OBTYPE         string `json:"obtype"`
	VX_MASK        string `json:"vx_mask"`
	INTERP_MTHD    string `json:"interp_mthd"`
	INTERP_PNTS    string `json:"interp_pnts"`
	FCST_THRESH    string `json:"fcst_thresh"`
	OBS_THRESH     string `json:"obs_thresh"`
	COV_THRESH     string `json:"cov_thresh"`
	ALPHA          string `json:"alpha"`
	LINE_TYPE      string `json:"line_type"`
}

// fillHeader functions
func (s *STAT_FHO) fill_STAT_FHO_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	if i <= dataLen && fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	i++
	if i <= dataLen && fields[1] != "" && fields[1] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	i++
	if i <= dataLen && fields[2] != "" && fields[2] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	i++
	if i <= dataLen && fields[4] != "" && fields[4] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	i++
	if i <= dataLen && fields[5] != "" && fields[5] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	i++
	if i <= dataLen && fields[6] != "" && fields[6] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	i++
	if i <= dataLen && fields[7] != "" && fields[7] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	i++
	if i <= dataLen && fields[8] != "" && fields[8] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	i++
	if i <= dataLen && fields[9] != "" && fields[9] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	i++
	if i <= dataLen && fields[10] != "" && fields[10] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	i++
	if i <= dataLen && fields[11] != "" && fields[11] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	i++
	if i <= dataLen && fields[12] != "" && fields[12] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	i++
	if i <= dataLen && fields[13] != "" && fields[13] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	i++
	if i <= dataLen && fields[14] != "" && fields[14] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	i++
	if i <= dataLen && fields[15] != "" && fields[15] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	i++
	if i <= dataLen && fields[16] != "" && fields[16] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	i++
	if i <= dataLen && fields[17] != "" && fields[17] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	i++
	if i <= dataLen && fields[18] != "" && fields[18] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	i++
	if i <= dataLen && fields[19] != "" && fields[19] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	i++
	if i <= dataLen && fields[20] != "" && fields[20] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	i++
	if i <= dataLen && fields[21] != "" && fields[21] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	i++
	if i <= dataLen && fields[22] != "" && fields[22] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	i++
	if i <= dataLen && fields[23] != "" && fields[23] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_MCTC) fill_STAT_MCTC_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	if i <= dataLen && fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	i++
	if i <= dataLen && fields[1] != "" && fields[1] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	i++
	if i <= dataLen && fields[2] != "" && fields[2] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	i++
	if i <= dataLen && fields[4] != "" && fields[4] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	i++
	if i <= dataLen && fields[5] != "" && fields[5] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	i++
	if i <= dataLen && fields[6] != "" && fields[6] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	i++
	if i <= dataLen && fields[7] != "" && fields[7] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	i++
	if i <= dataLen && fields[8] != "" && fields[8] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	i++
	if i <= dataLen && fields[9] != "" && fields[9] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	i++
	if i <= dataLen && fields[10] != "" && fields[10] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	i++
	if i <= dataLen && fields[11] != "" && fields[11] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	i++
	if i <= dataLen && fields[12] != "" && fields[12] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	i++
	if i <= dataLen && fields[13] != "" && fields[13] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	i++
	if i <= dataLen && fields[14] != "" && fields[14] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	i++
	if i <= dataLen && fields[15] != "" && fields[15] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	i++
	if i <= dataLen && fields[16] != "" && fields[16] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	i++
	if i <= dataLen && fields[17] != "" && fields[17] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	i++
	if i <= dataLen && fields[18] != "" && fields[18] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	i++
	if i <= dataLen && fields[19] != "" && fields[19] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	i++
	if i <= dataLen && fields[20] != "" && fields[20] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	i++
	if i <= dataLen && fields[21] != "" && fields[21] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	i++
	if i <= dataLen && fields[22] != "" && fields[22] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	i++
	if i <= dataLen && fields[23] != "" && fields[23] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_RELP) fill_STAT_RELP_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	if i <= dataLen && fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	i++
	if i <= dataLen && fields[1] != "" && fields[1] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	i++
	if i <= dataLen && fields[2] != "" && fields[2] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	i++
	if i <= dataLen && fields[4] != "" && fields[4] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	i++
	if i <= dataLen && fields[5] != "" && fields[5] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	i++
	if i <= dataLen && fields[6] != "" && fields[6] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	i++
	if i <= dataLen && fields[7] != "" && fields[7] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	i++
	if i <= dataLen && fields[8] != "" && fields[8] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	i++
	if i <= dataLen && fields[9] != "" && fields[9] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	i++
	if i <= dataLen && fields[10] != "" && fields[10] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	i++
	if i <= dataLen && fields[11] != "" && fields[11] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	i++
	if i <= dataLen && fields[12] != "" && fields[12] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	i++
	if i <= dataLen && fields[13] != "" && fields[13] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	i++
	if i <= dataLen && fields[14] != "" && fields[14] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	i++
	if i <= dataLen && fields[15] != "" && fields[15] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	i++
	if i <= dataLen && fields[16] != "" && fields[16] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	i++
	if i <= dataLen && fields[17] != "" && fields[17] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	i++
	if i <= dataLen && fields[18] != "" && fields[18] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	i++
	if i <= dataLen && fields[19] != "" && fields[19] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	i++
	if i <= dataLen && fields[20] != "" && fields[20] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	i++
	if i <= dataLen && fields[21] != "" && fields[21] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	i++
	if i <= dataLen && fields[22] != "" && fields[22] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	i++
	if i <= dataLen && fields[23] != "" && fields[23] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_SL1L2) fill_STAT_SL1L2_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	if i <= dataLen && fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	i++
	if i <= dataLen && fields[1] != "" && fields[1] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	i++
	if i <= dataLen && fields[2] != "" && fields[2] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	i++
	if i <= dataLen && fields[4] != "" && fields[4] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	i++
	if i <= dataLen && fields[5] != "" && fields[5] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	i++
	if i <= dataLen && fields[6] != "" && fields[6] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	i++
	if i <= dataLen && fields[7] != "" && fields[7] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	i++
	if i <= dataLen && fields[8] != "" && fields[8] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	i++
	if i <= dataLen && fields[9] != "" && fields[9] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	i++
	if i <= dataLen && fields[10] != "" && fields[10] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	i++
	if i <= dataLen && fields[11] != "" && fields[11] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	i++
	if i <= dataLen && fields[12] != "" && fields[12] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	i++
	if i <= dataLen && fields[13] != "" && fields[13] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	i++
	if i <= dataLen && fields[14] != "" && fields[14] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	i++
	if i <= dataLen && fields[15] != "" && fields[15] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	i++
	if i <= dataLen && fields[16] != "" && fields[16] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	i++
	if i <= dataLen && fields[17] != "" && fields[17] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	i++
	if i <= dataLen && fields[18] != "" && fields[18] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	i++
	if i <= dataLen && fields[19] != "" && fields[19] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	i++
	if i <= dataLen && fields[20] != "" && fields[20] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	i++
	if i <= dataLen && fields[21] != "" && fields[21] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	i++
	if i <= dataLen && fields[22] != "" && fields[22] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	i++
	if i <= dataLen && fields[23] != "" && fields[23] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_SSVAR) fill_STAT_SSVAR_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	if i <= dataLen && fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	i++
	if i <= dataLen && fields[1] != "" && fields[1] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	i++
	if i <= dataLen && fields[2] != "" && fields[2] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	i++
	if i <= dataLen && fields[4] != "" && fields[4] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	i++
	if i <= dataLen && fields[5] != "" && fields[5] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	i++
	if i <= dataLen && fields[6] != "" && fields[6] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	i++
	if i <= dataLen && fields[7] != "" && fields[7] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	i++
	if i <= dataLen && fields[8] != "" && fields[8] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	i++
	if i <= dataLen && fields[9] != "" && fields[9] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	i++
	if i <= dataLen && fields[10] != "" && fields[10] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	i++
	if i <= dataLen && fields[11] != "" && fields[11] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	i++
	if i <= dataLen && fields[12] != "" && fields[12] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	i++
	if i <= dataLen && fields[13] != "" && fields[13] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	i++
	if i <= dataLen && fields[14] != "" && fields[14] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	i++
	if i <= dataLen && fields[15] != "" && fields[15] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	i++
	if i <= dataLen && fields[16] != "" && fields[16] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	i++
	if i <= dataLen && fields[17] != "" && fields[17] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	i++
	if i <= dataLen && fields[18] != "" && fields[18] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	i++
	if i <= dataLen && fields[19] != "" && fields[19] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	i++
	if i <= dataLen && fields[20] != "" && fields[20] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	i++
	if i <= dataLen && fields[21] != "" && fields[21] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	i++
	if i <= dataLen && fields[22] != "" && fields[22] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	i++
	if i <= dataLen && fields[23] != "" && fields[23] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_VL1L2) fill_STAT_VL1L2_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	if i <= dataLen && fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	i++
	if i <= dataLen && fields[1] != "" && fields[1] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	i++
	if i <= dataLen && fields[2] != "" && fields[2] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	i++
	if i <= dataLen && fields[4] != "" && fields[4] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	i++
	if i <= dataLen && fields[5] != "" && fields[5] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	i++
	if i <= dataLen && fields[6] != "" && fields[6] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	i++
	if i <= dataLen && fields[7] != "" && fields[7] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	i++
	if i <= dataLen && fields[8] != "" && fields[8] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	i++
	if i <= dataLen && fields[9] != "" && fields[9] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	i++
	if i <= dataLen && fields[10] != "" && fields[10] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	i++
	if i <= dataLen && fields[11] != "" && fields[11] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	i++
	if i <= dataLen && fields[12] != "" && fields[12] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	i++
	if i <= dataLen && fields[13] != "" && fields[13] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	i++
	if i <= dataLen && fields[14] != "" && fields[14] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	i++
	if i <= dataLen && fields[15] != "" && fields[15] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	i++
	if i <= dataLen && fields[16] != "" && fields[16] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	i++
	if i <= dataLen && fields[17] != "" && fields[17] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	i++
	if i <= dataLen && fields[18] != "" && fields[18] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	i++
	if i <= dataLen && fields[19] != "" && fields[19] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	i++
	if i <= dataLen && fields[20] != "" && fields[20] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	i++
	if i <= dataLen && fields[21] != "" && fields[21] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	i++
	if i <= dataLen && fields[22] != "" && fields[22] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	i++
	if i <= dataLen && fields[23] != "" && fields[23] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_ECNT) fill_STAT_ECNT_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	if i <= dataLen && fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	i++
	if i <= dataLen && fields[1] != "" && fields[1] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	i++
	if i <= dataLen && fields[2] != "" && fields[2] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	i++
	if i <= dataLen && fields[4] != "" && fields[4] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	i++
	if i <= dataLen && fields[5] != "" && fields[5] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	i++
	if i <= dataLen && fields[6] != "" && fields[6] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	i++
	if i <= dataLen && fields[7] != "" && fields[7] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	i++
	if i <= dataLen && fields[8] != "" && fields[8] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	i++
	if i <= dataLen && fields[9] != "" && fields[9] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	i++
	if i <= dataLen && fields[10] != "" && fields[10] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	i++
	if i <= dataLen && fields[11] != "" && fields[11] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	i++
	if i <= dataLen && fields[12] != "" && fields[12] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	i++
	if i <= dataLen && fields[13] != "" && fields[13] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	i++
	if i <= dataLen && fields[14] != "" && fields[14] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	i++
	if i <= dataLen && fields[15] != "" && fields[15] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	i++
	if i <= dataLen && fields[16] != "" && fields[16] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	i++
	if i <= dataLen && fields[17] != "" && fields[17] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	i++
	if i <= dataLen && fields[18] != "" && fields[18] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	i++
	if i <= dataLen && fields[19] != "" && fields[19] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	i++
	if i <= dataLen && fields[20] != "" && fields[20] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	i++
	if i <= dataLen && fields[21] != "" && fields[21] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	i++
	if i <= dataLen && fields[22] != "" && fields[22] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	i++
	if i <= dataLen && fields[23] != "" && fields[23] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_RPS) fill_STAT_RPS_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	if i <= dataLen && fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	i++
	if i <= dataLen && fields[1] != "" && fields[1] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	i++
	if i <= dataLen && fields[2] != "" && fields[2] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	i++
	if i <= dataLen && fields[4] != "" && fields[4] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	i++
	if i <= dataLen && fields[5] != "" && fields[5] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	i++
	if i <= dataLen && fields[6] != "" && fields[6] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	i++
	if i <= dataLen && fields[7] != "" && fields[7] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	i++
	if i <= dataLen && fields[8] != "" && fields[8] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	i++
	if i <= dataLen && fields[9] != "" && fields[9] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	i++
	if i <= dataLen && fields[10] != "" && fields[10] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	i++
	if i <= dataLen && fields[11] != "" && fields[11] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	i++
	if i <= dataLen && fields[12] != "" && fields[12] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	i++
	if i <= dataLen && fields[13] != "" && fields[13] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	i++
	if i <= dataLen && fields[14] != "" && fields[14] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	i++
	if i <= dataLen && fields[15] != "" && fields[15] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	i++
	if i <= dataLen && fields[16] != "" && fields[16] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	i++
	if i <= dataLen && fields[17] != "" && fields[17] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	i++
	if i <= dataLen && fields[18] != "" && fields[18] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	i++
	if i <= dataLen && fields[19] != "" && fields[19] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	i++
	if i <= dataLen && fields[20] != "" && fields[20] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	i++
	if i <= dataLen && fields[21] != "" && fields[21] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	i++
	if i <= dataLen && fields[22] != "" && fields[22] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	i++
	if i <= dataLen && fields[23] != "" && fields[23] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_CTC) fill_STAT_CTC_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	if i <= dataLen && fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	i++
	if i <= dataLen && fields[1] != "" && fields[1] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	i++
	if i <= dataLen && fields[2] != "" && fields[2] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	i++
	if i <= dataLen && fields[4] != "" && fields[4] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	i++
	if i <= dataLen && fields[5] != "" && fields[5] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	i++
	if i <= dataLen && fields[6] != "" && fields[6] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	i++
	if i <= dataLen && fields[7] != "" && fields[7] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	i++
	if i <= dataLen && fields[8] != "" && fields[8] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	i++
	if i <= dataLen && fields[9] != "" && fields[9] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	i++
	if i <= dataLen && fields[10] != "" && fields[10] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	i++
	if i <= dataLen && fields[11] != "" && fields[11] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	i++
	if i <= dataLen && fields[12] != "" && fields[12] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	i++
	if i <= dataLen && fields[13] != "" && fields[13] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	i++
	if i <= dataLen && fields[14] != "" && fields[14] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	i++
	if i <= dataLen && fields[15] != "" && fields[15] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	i++
	if i <= dataLen && fields[16] != "" && fields[16] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	i++
	if i <= dataLen && fields[17] != "" && fields[17] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	i++
	if i <= dataLen && fields[18] != "" && fields[18] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	i++
	if i <= dataLen && fields[19] != "" && fields[19] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	i++
	if i <= dataLen && fields[20] != "" && fields[20] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	i++
	if i <= dataLen && fields[21] != "" && fields[21] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	i++
	if i <= dataLen && fields[22] != "" && fields[22] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	i++
	if i <= dataLen && fields[23] != "" && fields[23] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_SEEPS) fill_STAT_SEEPS_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	if i <= dataLen && fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	i++
	if i <= dataLen && fields[1] != "" && fields[1] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	i++
	if i <= dataLen && fields[2] != "" && fields[2] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	i++
	if i <= dataLen && fields[4] != "" && fields[4] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	i++
	if i <= dataLen && fields[5] != "" && fields[5] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	i++
	if i <= dataLen && fields[6] != "" && fields[6] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	i++
	if i <= dataLen && fields[7] != "" && fields[7] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	i++
	if i <= dataLen && fields[8] != "" && fields[8] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	i++
	if i <= dataLen && fields[9] != "" && fields[9] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	i++
	if i <= dataLen && fields[10] != "" && fields[10] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	i++
	if i <= dataLen && fields[11] != "" && fields[11] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	i++
	if i <= dataLen && fields[12] != "" && fields[12] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	i++
	if i <= dataLen && fields[13] != "" && fields[13] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	i++
	if i <= dataLen && fields[14] != "" && fields[14] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	i++
	if i <= dataLen && fields[15] != "" && fields[15] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	i++
	if i <= dataLen && fields[16] != "" && fields[16] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	i++
	if i <= dataLen && fields[17] != "" && fields[17] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	i++
	if i <= dataLen && fields[18] != "" && fields[18] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	i++
	if i <= dataLen && fields[19] != "" && fields[19] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	i++
	if i <= dataLen && fields[20] != "" && fields[20] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	i++
	if i <= dataLen && fields[21] != "" && fields[21] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	i++
	if i <= dataLen && fields[22] != "" && fields[22] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	i++
	if i <= dataLen && fields[23] != "" && fields[23] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_NBRCTS) fill_STAT_NBRCTS_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	if i <= dataLen && fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	i++
	if i <= dataLen && fields[1] != "" && fields[1] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	i++
	if i <= dataLen && fields[2] != "" && fields[2] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	i++
	if i <= dataLen && fields[4] != "" && fields[4] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	i++
	if i <= dataLen && fields[5] != "" && fields[5] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	i++
	if i <= dataLen && fields[6] != "" && fields[6] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	i++
	if i <= dataLen && fields[7] != "" && fields[7] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	i++
	if i <= dataLen && fields[8] != "" && fields[8] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	i++
	if i <= dataLen && fields[9] != "" && fields[9] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	i++
	if i <= dataLen && fields[10] != "" && fields[10] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	i++
	if i <= dataLen && fields[11] != "" && fields[11] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	i++
	if i <= dataLen && fields[12] != "" && fields[12] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	i++
	if i <= dataLen && fields[13] != "" && fields[13] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	i++
	if i <= dataLen && fields[14] != "" && fields[14] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	i++
	if i <= dataLen && fields[15] != "" && fields[15] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	i++
	if i <= dataLen && fields[16] != "" && fields[16] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	i++
	if i <= dataLen && fields[17] != "" && fields[17] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	i++
	if i <= dataLen && fields[18] != "" && fields[18] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	i++
	if i <= dataLen && fields[19] != "" && fields[19] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	i++
	if i <= dataLen && fields[20] != "" && fields[20] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	i++
	if i <= dataLen && fields[21] != "" && fields[21] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	i++
	if i <= dataLen && fields[22] != "" && fields[22] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	i++
	if i <= dataLen && fields[23] != "" && fields[23] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_PJC) fill_STAT_PJC_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	if i <= dataLen && fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	i++
	if i <= dataLen && fields[1] != "" && fields[1] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	i++
	if i <= dataLen && fields[2] != "" && fields[2] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	i++
	if i <= dataLen && fields[4] != "" && fields[4] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	i++
	if i <= dataLen && fields[5] != "" && fields[5] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	i++
	if i <= dataLen && fields[6] != "" && fields[6] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	i++
	if i <= dataLen && fields[7] != "" && fields[7] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	i++
	if i <= dataLen && fields[8] != "" && fields[8] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	i++
	if i <= dataLen && fields[9] != "" && fields[9] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	i++
	if i <= dataLen && fields[10] != "" && fields[10] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	i++
	if i <= dataLen && fields[11] != "" && fields[11] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	i++
	if i <= dataLen && fields[12] != "" && fields[12] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	i++
	if i <= dataLen && fields[13] != "" && fields[13] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	i++
	if i <= dataLen && fields[14] != "" && fields[14] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	i++
	if i <= dataLen && fields[15] != "" && fields[15] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	i++
	if i <= dataLen && fields[16] != "" && fields[16] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	i++
	if i <= dataLen && fields[17] != "" && fields[17] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	i++
	if i <= dataLen && fields[18] != "" && fields[18] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	i++
	if i <= dataLen && fields[19] != "" && fields[19] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	i++
	if i <= dataLen && fields[20] != "" && fields[20] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	i++
	if i <= dataLen && fields[21] != "" && fields[21] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	i++
	if i <= dataLen && fields[22] != "" && fields[22] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	i++
	if i <= dataLen && fields[23] != "" && fields[23] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_PSTD) fill_STAT_PSTD_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	if i <= dataLen && fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	i++
	if i <= dataLen && fields[1] != "" && fields[1] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	i++
	if i <= dataLen && fields[2] != "" && fields[2] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	i++
	if i <= dataLen && fields[4] != "" && fields[4] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	i++
	if i <= dataLen && fields[5] != "" && fields[5] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	i++
	if i <= dataLen && fields[6] != "" && fields[6] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	i++
	if i <= dataLen && fields[7] != "" && fields[7] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	i++
	if i <= dataLen && fields[8] != "" && fields[8] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	i++
	if i <= dataLen && fields[9] != "" && fields[9] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	i++
	if i <= dataLen && fields[10] != "" && fields[10] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	i++
	if i <= dataLen && fields[11] != "" && fields[11] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	i++
	if i <= dataLen && fields[12] != "" && fields[12] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	i++
	if i <= dataLen && fields[13] != "" && fields[13] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	i++
	if i <= dataLen && fields[14] != "" && fields[14] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	i++
	if i <= dataLen && fields[15] != "" && fields[15] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	i++
	if i <= dataLen && fields[16] != "" && fields[16] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	i++
	if i <= dataLen && fields[17] != "" && fields[17] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	i++
	if i <= dataLen && fields[18] != "" && fields[18] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	i++
	if i <= dataLen && fields[19] != "" && fields[19] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	i++
	if i <= dataLen && fields[20] != "" && fields[20] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	i++
	if i <= dataLen && fields[21] != "" && fields[21] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	i++
	if i <= dataLen && fields[22] != "" && fields[22] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	i++
	if i <= dataLen && fields[23] != "" && fields[23] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_ECLV) fill_STAT_ECLV_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	if i <= dataLen && fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	i++
	if i <= dataLen && fields[1] != "" && fields[1] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	i++
	if i <= dataLen && fields[2] != "" && fields[2] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	i++
	if i <= dataLen && fields[4] != "" && fields[4] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	i++
	if i <= dataLen && fields[5] != "" && fields[5] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	i++
	if i <= dataLen && fields[6] != "" && fields[6] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	i++
	if i <= dataLen && fields[7] != "" && fields[7] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	i++
	if i <= dataLen && fields[8] != "" && fields[8] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	i++
	if i <= dataLen && fields[9] != "" && fields[9] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	i++
	if i <= dataLen && fields[10] != "" && fields[10] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	i++
	if i <= dataLen && fields[11] != "" && fields[11] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	i++
	if i <= dataLen && fields[12] != "" && fields[12] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	i++
	if i <= dataLen && fields[13] != "" && fields[13] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	i++
	if i <= dataLen && fields[14] != "" && fields[14] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	i++
	if i <= dataLen && fields[15] != "" && fields[15] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	i++
	if i <= dataLen && fields[16] != "" && fields[16] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	i++
	if i <= dataLen && fields[17] != "" && fields[17] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	i++
	if i <= dataLen && fields[18] != "" && fields[18] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	i++
	if i <= dataLen && fields[19] != "" && fields[19] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	i++
	if i <= dataLen && fields[20] != "" && fields[20] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	i++
	if i <= dataLen && fields[21] != "" && fields[21] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	i++
	if i <= dataLen && fields[22] != "" && fields[22] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	i++
	if i <= dataLen && fields[23] != "" && fields[23] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_VAL1L2) fill_STAT_VAL1L2_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	if i <= dataLen && fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	i++
	if i <= dataLen && fields[1] != "" && fields[1] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	i++
	if i <= dataLen && fields[2] != "" && fields[2] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	i++
	if i <= dataLen && fields[4] != "" && fields[4] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	i++
	if i <= dataLen && fields[5] != "" && fields[5] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	i++
	if i <= dataLen && fields[6] != "" && fields[6] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	i++
	if i <= dataLen && fields[7] != "" && fields[7] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	i++
	if i <= dataLen && fields[8] != "" && fields[8] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	i++
	if i <= dataLen && fields[9] != "" && fields[9] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	i++
	if i <= dataLen && fields[10] != "" && fields[10] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	i++
	if i <= dataLen && fields[11] != "" && fields[11] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	i++
	if i <= dataLen && fields[12] != "" && fields[12] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	i++
	if i <= dataLen && fields[13] != "" && fields[13] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	i++
	if i <= dataLen && fields[14] != "" && fields[14] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	i++
	if i <= dataLen && fields[15] != "" && fields[15] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	i++
	if i <= dataLen && fields[16] != "" && fields[16] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	i++
	if i <= dataLen && fields[17] != "" && fields[17] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	i++
	if i <= dataLen && fields[18] != "" && fields[18] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	i++
	if i <= dataLen && fields[19] != "" && fields[19] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	i++
	if i <= dataLen && fields[20] != "" && fields[20] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	i++
	if i <= dataLen && fields[21] != "" && fields[21] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	i++
	if i <= dataLen && fields[22] != "" && fields[22] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	i++
	if i <= dataLen && fields[23] != "" && fields[23] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_SSIDX) fill_STAT_SSIDX_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	if i <= dataLen && fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	i++
	if i <= dataLen && fields[1] != "" && fields[1] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	i++
	if i <= dataLen && fields[2] != "" && fields[2] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	i++
	if i <= dataLen && fields[4] != "" && fields[4] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	i++
	if i <= dataLen && fields[5] != "" && fields[5] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	i++
	if i <= dataLen && fields[6] != "" && fields[6] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	i++
	if i <= dataLen && fields[7] != "" && fields[7] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	i++
	if i <= dataLen && fields[8] != "" && fields[8] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	i++
	if i <= dataLen && fields[9] != "" && fields[9] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	i++
	if i <= dataLen && fields[10] != "" && fields[10] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	i++
	if i <= dataLen && fields[11] != "" && fields[11] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	i++
	if i <= dataLen && fields[12] != "" && fields[12] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	i++
	if i <= dataLen && fields[13] != "" && fields[13] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	i++
	if i <= dataLen && fields[14] != "" && fields[14] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	i++
	if i <= dataLen && fields[15] != "" && fields[15] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	i++
	if i <= dataLen && fields[16] != "" && fields[16] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	i++
	if i <= dataLen && fields[17] != "" && fields[17] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	i++
	if i <= dataLen && fields[18] != "" && fields[18] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	i++
	if i <= dataLen && fields[19] != "" && fields[19] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	i++
	if i <= dataLen && fields[20] != "" && fields[20] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	i++
	if i <= dataLen && fields[21] != "" && fields[21] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	i++
	if i <= dataLen && fields[22] != "" && fields[22] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	i++
	if i <= dataLen && fields[23] != "" && fields[23] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *MODE_OBJ) fill_MODE_OBJ_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	if i <= dataLen && fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	i++
	if i <= dataLen && fields[1] != "" && fields[1] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	i++
	if i <= dataLen && fields[2] != "" && fields[2] != "NA" {
		(*doc)["N_VALID"] = fields[2]
	}
	i++
	if i <= dataLen && fields[3] != "" && fields[3] != "NA" {
		(*doc)["GRID_RES"] = fields[3]
	}
	i++
	if i <= dataLen && fields[4] != "" && fields[4] != "NA" {
		(*doc)["DESC"] = fields[4]
	}
	i++
	if i <= dataLen && fields[6] != "" && fields[6] != "NA" {
		(*doc)["FCST_VALID"] = fields[6]
	}
	i++
	if i <= dataLen && fields[7] != "" && fields[7] != "NA" {
		(*doc)["FCST_ACCUM"] = fields[7]
	}
	i++
	if i <= dataLen && fields[8] != "" && fields[8] != "NA" {
		(*doc)["OBS_LEAD"] = fields[8]
	}
	i++
	if i <= dataLen && fields[9] != "" && fields[9] != "NA" {
		(*doc)["OBS_VALID"] = fields[9]
	}
	i++
	if i <= dataLen && fields[10] != "" && fields[10] != "NA" {
		(*doc)["OBS_ACCUM"] = fields[10]
	}
	i++
	if i <= dataLen && fields[11] != "" && fields[11] != "NA" {
		(*doc)["FCST_RAD"] = fields[11]
	}
	i++
	if i <= dataLen && fields[12] != "" && fields[12] != "NA" {
		(*doc)["FCST_THR"] = fields[12]
	}
	i++
	if i <= dataLen && fields[13] != "" && fields[13] != "NA" {
		(*doc)["OBS_RAD"] = fields[13]
	}
	i++
	if i <= dataLen && fields[14] != "" && fields[14] != "NA" {
		(*doc)["OBS_THR"] = fields[14]
	}
	i++
	if i <= dataLen && fields[15] != "" && fields[15] != "NA" {
		(*doc)["FCST_VAR"] = fields[15]
	}
	i++
	if i <= dataLen && fields[16] != "" && fields[16] != "NA" {
		(*doc)["FCST_UNITS"] = fields[16]
	}
	i++
	if i <= dataLen && fields[18] != "" && fields[18] != "NA" {
		(*doc)["OBS_VAR"] = fields[18]
	}
	i++
	if i <= dataLen && fields[19] != "" && fields[19] != "NA" {
		(*doc)["OBS_UNITS"] = fields[19]
	}
	i++
	if i <= dataLen && fields[20] != "" && fields[20] != "NA" {
		(*doc)["OBS_LEV"] = fields[20]
	}
	i++
	if i <= dataLen && fields[21] != "" && fields[21] != "NA" {
		(*doc)["OBTYPE"] = fields[21]
	}
}

func (s *TCST_TCMPR) fill_TCST_TCMPR_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	if i <= dataLen && fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	i++
	if i <= dataLen && fields[1] != "" && fields[1] != "NA" {
		(*doc)["AMODEL"] = fields[1]
	}
	i++
	if i <= dataLen && fields[2] != "" && fields[2] != "NA" {
		(*doc)["BMODEL"] = fields[2]
	}
	i++
	if i <= dataLen && fields[3] != "" && fields[3] != "NA" {
		(*doc)["DESC"] = fields[3]
	}
	i++
	if i <= dataLen && fields[4] != "" && fields[4] != "NA" {
		(*doc)["STORM_ID"] = fields[4]
	}
	i++
	if i <= dataLen && fields[5] != "" && fields[5] != "NA" {
		(*doc)["BASIN"] = fields[5]
	}
	i++
	if i <= dataLen && fields[6] != "" && fields[6] != "NA" {
		(*doc)["CYCLONE"] = fields[6]
	}
	i++
	if i <= dataLen && fields[7] != "" && fields[7] != "NA" {
		(*doc)["STORM_NAME"] = fields[7]
	}
	i++
	if i <= dataLen && fields[8] != "" && fields[8] != "NA" {
		(*doc)["INIT"] = fields[8]
	}
	i++
	if i <= dataLen && fields[9] != "" && fields[9] != "NA" {
		(*doc)["LEAD"] = fields[9]
	}
	i++
	if i <= dataLen && fields[10] != "" && fields[10] != "NA" {
		(*doc)["VALID"] = fields[10]
	}
	i++
	if i <= dataLen && fields[11] != "" && fields[11] != "NA" {
		(*doc)["INIT_MASK"] = fields[11]
	}
	i++
	if i <= dataLen && fields[12] != "" && fields[12] != "NA" {
		(*doc)["VALID_MASK"] = fields[12]
	}
	i++
	if i <= dataLen && fields[13] != "" && fields[13] != "NA" {
		(*doc)["LINE_TYPE"] = fields[13]
	}
}

func (s *STAT_RHIST) fill_STAT_RHIST_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	if i <= dataLen && fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	i++
	if i <= dataLen && fields[1] != "" && fields[1] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	i++
	if i <= dataLen && fields[2] != "" && fields[2] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	i++
	if i <= dataLen && fields[4] != "" && fields[4] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	i++
	if i <= dataLen && fields[5] != "" && fields[5] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	i++
	if i <= dataLen && fields[6] != "" && fields[6] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	i++
	if i <= dataLen && fields[7] != "" && fields[7] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	i++
	if i <= dataLen && fields[8] != "" && fields[8] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	i++
	if i <= dataLen && fields[9] != "" && fields[9] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	i++
	if i <= dataLen && fields[10] != "" && fields[10] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	i++
	if i <= dataLen && fields[11] != "" && fields[11] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	i++
	if i <= dataLen && fields[12] != "" && fields[12] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	i++
	if i <= dataLen && fields[13] != "" && fields[13] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	i++
	if i <= dataLen && fields[14] != "" && fields[14] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	i++
	if i <= dataLen && fields[15] != "" && fields[15] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	i++
	if i <= dataLen && fields[16] != "" && fields[16] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	i++
	if i <= dataLen && fields[17] != "" && fields[17] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	i++
	if i <= dataLen && fields[18] != "" && fields[18] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	i++
	if i <= dataLen && fields[19] != "" && fields[19] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	i++
	if i <= dataLen && fields[20] != "" && fields[20] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	i++
	if i <= dataLen && fields[21] != "" && fields[21] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	i++
	if i <= dataLen && fields[22] != "" && fields[22] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	i++
	if i <= dataLen && fields[23] != "" && fields[23] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_VCNT) fill_STAT_VCNT_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	if i <= dataLen && fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	i++
	if i <= dataLen && fields[1] != "" && fields[1] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	i++
	if i <= dataLen && fields[2] != "" && fields[2] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	i++
	if i <= dataLen && fields[4] != "" && fields[4] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	i++
	if i <= dataLen && fields[5] != "" && fields[5] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	i++
	if i <= dataLen && fields[6] != "" && fields[6] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	i++
	if i <= dataLen && fields[7] != "" && fields[7] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	i++
	if i <= dataLen && fields[8] != "" && fields[8] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	i++
	if i <= dataLen && fields[9] != "" && fields[9] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	i++
	if i <= dataLen && fields[10] != "" && fields[10] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	i++
	if i <= dataLen && fields[11] != "" && fields[11] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	i++
	if i <= dataLen && fields[12] != "" && fields[12] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	i++
	if i <= dataLen && fields[13] != "" && fields[13] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	i++
	if i <= dataLen && fields[14] != "" && fields[14] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	i++
	if i <= dataLen && fields[15] != "" && fields[15] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	i++
	if i <= dataLen && fields[16] != "" && fields[16] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	i++
	if i <= dataLen && fields[17] != "" && fields[17] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	i++
	if i <= dataLen && fields[18] != "" && fields[18] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	i++
	if i <= dataLen && fields[19] != "" && fields[19] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	i++
	if i <= dataLen && fields[20] != "" && fields[20] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	i++
	if i <= dataLen && fields[21] != "" && fields[21] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	i++
	if i <= dataLen && fields[22] != "" && fields[22] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	i++
	if i <= dataLen && fields[23] != "" && fields[23] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_ISC) fill_STAT_ISC_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	if i <= dataLen && fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	i++
	if i <= dataLen && fields[1] != "" && fields[1] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	i++
	if i <= dataLen && fields[2] != "" && fields[2] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	i++
	if i <= dataLen && fields[4] != "" && fields[4] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	i++
	if i <= dataLen && fields[5] != "" && fields[5] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	i++
	if i <= dataLen && fields[6] != "" && fields[6] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	i++
	if i <= dataLen && fields[7] != "" && fields[7] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	i++
	if i <= dataLen && fields[8] != "" && fields[8] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	i++
	if i <= dataLen && fields[9] != "" && fields[9] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	i++
	if i <= dataLen && fields[10] != "" && fields[10] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	i++
	if i <= dataLen && fields[11] != "" && fields[11] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	i++
	if i <= dataLen && fields[12] != "" && fields[12] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	i++
	if i <= dataLen && fields[13] != "" && fields[13] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	i++
	if i <= dataLen && fields[14] != "" && fields[14] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	i++
	if i <= dataLen && fields[15] != "" && fields[15] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	i++
	if i <= dataLen && fields[16] != "" && fields[16] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	i++
	if i <= dataLen && fields[17] != "" && fields[17] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	i++
	if i <= dataLen && fields[18] != "" && fields[18] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	i++
	if i <= dataLen && fields[19] != "" && fields[19] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	i++
	if i <= dataLen && fields[20] != "" && fields[20] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	i++
	if i <= dataLen && fields[21] != "" && fields[21] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	i++
	if i <= dataLen && fields[22] != "" && fields[22] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	i++
	if i <= dataLen && fields[23] != "" && fields[23] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_MCTS) fill_STAT_MCTS_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	if i <= dataLen && fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	i++
	if i <= dataLen && fields[1] != "" && fields[1] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	i++
	if i <= dataLen && fields[2] != "" && fields[2] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	i++
	if i <= dataLen && fields[4] != "" && fields[4] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	i++
	if i <= dataLen && fields[5] != "" && fields[5] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	i++
	if i <= dataLen && fields[6] != "" && fields[6] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	i++
	if i <= dataLen && fields[7] != "" && fields[7] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	i++
	if i <= dataLen && fields[8] != "" && fields[8] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	i++
	if i <= dataLen && fields[9] != "" && fields[9] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	i++
	if i <= dataLen && fields[10] != "" && fields[10] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	i++
	if i <= dataLen && fields[11] != "" && fields[11] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	i++
	if i <= dataLen && fields[12] != "" && fields[12] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	i++
	if i <= dataLen && fields[13] != "" && fields[13] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	i++
	if i <= dataLen && fields[14] != "" && fields[14] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	i++
	if i <= dataLen && fields[15] != "" && fields[15] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	i++
	if i <= dataLen && fields[16] != "" && fields[16] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	i++
	if i <= dataLen && fields[17] != "" && fields[17] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	i++
	if i <= dataLen && fields[18] != "" && fields[18] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	i++
	if i <= dataLen && fields[19] != "" && fields[19] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	i++
	if i <= dataLen && fields[20] != "" && fields[20] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	i++
	if i <= dataLen && fields[21] != "" && fields[21] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	i++
	if i <= dataLen && fields[22] != "" && fields[22] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	i++
	if i <= dataLen && fields[23] != "" && fields[23] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_ORANK) fill_STAT_ORANK_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	if i <= dataLen && fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	i++
	if i <= dataLen && fields[1] != "" && fields[1] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	i++
	if i <= dataLen && fields[2] != "" && fields[2] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	i++
	if i <= dataLen && fields[4] != "" && fields[4] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	i++
	if i <= dataLen && fields[5] != "" && fields[5] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	i++
	if i <= dataLen && fields[6] != "" && fields[6] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	i++
	if i <= dataLen && fields[7] != "" && fields[7] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	i++
	if i <= dataLen && fields[8] != "" && fields[8] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	i++
	if i <= dataLen && fields[9] != "" && fields[9] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	i++
	if i <= dataLen && fields[10] != "" && fields[10] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	i++
	if i <= dataLen && fields[11] != "" && fields[11] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	i++
	if i <= dataLen && fields[12] != "" && fields[12] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	i++
	if i <= dataLen && fields[13] != "" && fields[13] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	i++
	if i <= dataLen && fields[14] != "" && fields[14] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	i++
	if i <= dataLen && fields[15] != "" && fields[15] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	i++
	if i <= dataLen && fields[16] != "" && fields[16] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	i++
	if i <= dataLen && fields[17] != "" && fields[17] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	i++
	if i <= dataLen && fields[18] != "" && fields[18] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	i++
	if i <= dataLen && fields[19] != "" && fields[19] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	i++
	if i <= dataLen && fields[20] != "" && fields[20] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	i++
	if i <= dataLen && fields[21] != "" && fields[21] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	i++
	if i <= dataLen && fields[22] != "" && fields[22] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	i++
	if i <= dataLen && fields[23] != "" && fields[23] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_PRC) fill_STAT_PRC_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	if i <= dataLen && fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	i++
	if i <= dataLen && fields[1] != "" && fields[1] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	i++
	if i <= dataLen && fields[2] != "" && fields[2] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	i++
	if i <= dataLen && fields[4] != "" && fields[4] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	i++
	if i <= dataLen && fields[5] != "" && fields[5] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	i++
	if i <= dataLen && fields[6] != "" && fields[6] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	i++
	if i <= dataLen && fields[7] != "" && fields[7] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	i++
	if i <= dataLen && fields[8] != "" && fields[8] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	i++
	if i <= dataLen && fields[9] != "" && fields[9] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	i++
	if i <= dataLen && fields[10] != "" && fields[10] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	i++
	if i <= dataLen && fields[11] != "" && fields[11] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	i++
	if i <= dataLen && fields[12] != "" && fields[12] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	i++
	if i <= dataLen && fields[13] != "" && fields[13] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	i++
	if i <= dataLen && fields[14] != "" && fields[14] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	i++
	if i <= dataLen && fields[15] != "" && fields[15] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	i++
	if i <= dataLen && fields[16] != "" && fields[16] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	i++
	if i <= dataLen && fields[17] != "" && fields[17] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	i++
	if i <= dataLen && fields[18] != "" && fields[18] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	i++
	if i <= dataLen && fields[19] != "" && fields[19] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	i++
	if i <= dataLen && fields[20] != "" && fields[20] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	i++
	if i <= dataLen && fields[21] != "" && fields[21] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	i++
	if i <= dataLen && fields[22] != "" && fields[22] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	i++
	if i <= dataLen && fields[23] != "" && fields[23] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_CNT) fill_STAT_CNT_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	if i <= dataLen && fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	i++
	if i <= dataLen && fields[1] != "" && fields[1] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	i++
	if i <= dataLen && fields[2] != "" && fields[2] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	i++
	if i <= dataLen && fields[4] != "" && fields[4] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	i++
	if i <= dataLen && fields[5] != "" && fields[5] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	i++
	if i <= dataLen && fields[6] != "" && fields[6] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	i++
	if i <= dataLen && fields[7] != "" && fields[7] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	i++
	if i <= dataLen && fields[8] != "" && fields[8] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	i++
	if i <= dataLen && fields[9] != "" && fields[9] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	i++
	if i <= dataLen && fields[10] != "" && fields[10] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	i++
	if i <= dataLen && fields[11] != "" && fields[11] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	i++
	if i <= dataLen && fields[12] != "" && fields[12] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	i++
	if i <= dataLen && fields[13] != "" && fields[13] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	i++
	if i <= dataLen && fields[14] != "" && fields[14] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	i++
	if i <= dataLen && fields[15] != "" && fields[15] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	i++
	if i <= dataLen && fields[16] != "" && fields[16] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	i++
	if i <= dataLen && fields[17] != "" && fields[17] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	i++
	if i <= dataLen && fields[18] != "" && fields[18] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	i++
	if i <= dataLen && fields[19] != "" && fields[19] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	i++
	if i <= dataLen && fields[20] != "" && fields[20] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	i++
	if i <= dataLen && fields[21] != "" && fields[21] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	i++
	if i <= dataLen && fields[22] != "" && fields[22] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	i++
	if i <= dataLen && fields[23] != "" && fields[23] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_NBRCTC) fill_STAT_NBRCTC_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	if i <= dataLen && fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	i++
	if i <= dataLen && fields[1] != "" && fields[1] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	i++
	if i <= dataLen && fields[2] != "" && fields[2] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	i++
	if i <= dataLen && fields[4] != "" && fields[4] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	i++
	if i <= dataLen && fields[5] != "" && fields[5] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	i++
	if i <= dataLen && fields[6] != "" && fields[6] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	i++
	if i <= dataLen && fields[7] != "" && fields[7] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	i++
	if i <= dataLen && fields[8] != "" && fields[8] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	i++
	if i <= dataLen && fields[9] != "" && fields[9] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	i++
	if i <= dataLen && fields[10] != "" && fields[10] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	i++
	if i <= dataLen && fields[11] != "" && fields[11] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	i++
	if i <= dataLen && fields[12] != "" && fields[12] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	i++
	if i <= dataLen && fields[13] != "" && fields[13] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	i++
	if i <= dataLen && fields[14] != "" && fields[14] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	i++
	if i <= dataLen && fields[15] != "" && fields[15] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	i++
	if i <= dataLen && fields[16] != "" && fields[16] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	i++
	if i <= dataLen && fields[17] != "" && fields[17] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	i++
	if i <= dataLen && fields[18] != "" && fields[18] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	i++
	if i <= dataLen && fields[19] != "" && fields[19] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	i++
	if i <= dataLen && fields[20] != "" && fields[20] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	i++
	if i <= dataLen && fields[21] != "" && fields[21] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	i++
	if i <= dataLen && fields[22] != "" && fields[22] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	i++
	if i <= dataLen && fields[23] != "" && fields[23] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_GENMPR) fill_STAT_GENMPR_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	if i <= dataLen && fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	i++
	if i <= dataLen && fields[1] != "" && fields[1] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	i++
	if i <= dataLen && fields[2] != "" && fields[2] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	i++
	if i <= dataLen && fields[4] != "" && fields[4] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	i++
	if i <= dataLen && fields[5] != "" && fields[5] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	i++
	if i <= dataLen && fields[6] != "" && fields[6] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	i++
	if i <= dataLen && fields[7] != "" && fields[7] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	i++
	if i <= dataLen && fields[8] != "" && fields[8] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	i++
	if i <= dataLen && fields[9] != "" && fields[9] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	i++
	if i <= dataLen && fields[10] != "" && fields[10] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	i++
	if i <= dataLen && fields[11] != "" && fields[11] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	i++
	if i <= dataLen && fields[12] != "" && fields[12] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	i++
	if i <= dataLen && fields[13] != "" && fields[13] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	i++
	if i <= dataLen && fields[14] != "" && fields[14] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	i++
	if i <= dataLen && fields[15] != "" && fields[15] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	i++
	if i <= dataLen && fields[16] != "" && fields[16] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	i++
	if i <= dataLen && fields[17] != "" && fields[17] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	i++
	if i <= dataLen && fields[18] != "" && fields[18] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	i++
	if i <= dataLen && fields[19] != "" && fields[19] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	i++
	if i <= dataLen && fields[20] != "" && fields[20] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	i++
	if i <= dataLen && fields[21] != "" && fields[21] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	i++
	if i <= dataLen && fields[22] != "" && fields[22] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	i++
	if i <= dataLen && fields[23] != "" && fields[23] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *MODE_CTS) fill_MODE_CTS_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	if i <= dataLen && fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	i++
	if i <= dataLen && fields[1] != "" && fields[1] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	i++
	if i <= dataLen && fields[2] != "" && fields[2] != "NA" {
		(*doc)["N_VALID"] = fields[2]
	}
	i++
	if i <= dataLen && fields[3] != "" && fields[3] != "NA" {
		(*doc)["GRID_RES"] = fields[3]
	}
	i++
	if i <= dataLen && fields[4] != "" && fields[4] != "NA" {
		(*doc)["DESC"] = fields[4]
	}
	i++
	if i <= dataLen && fields[6] != "" && fields[6] != "NA" {
		(*doc)["FCST_VALID"] = fields[6]
	}
	i++
	if i <= dataLen && fields[7] != "" && fields[7] != "NA" {
		(*doc)["FCST_ACCUM"] = fields[7]
	}
	i++
	if i <= dataLen && fields[8] != "" && fields[8] != "NA" {
		(*doc)["OBS_LEAD"] = fields[8]
	}
	i++
	if i <= dataLen && fields[9] != "" && fields[9] != "NA" {
		(*doc)["OBS_VALID"] = fields[9]
	}
	i++
	if i <= dataLen && fields[10] != "" && fields[10] != "NA" {
		(*doc)["OBS_ACCUM"] = fields[10]
	}
	i++
	if i <= dataLen && fields[11] != "" && fields[11] != "NA" {
		(*doc)["FCST_RAD"] = fields[11]
	}
	i++
	if i <= dataLen && fields[12] != "" && fields[12] != "NA" {
		(*doc)["FCST_THR"] = fields[12]
	}
	i++
	if i <= dataLen && fields[13] != "" && fields[13] != "NA" {
		(*doc)["OBS_RAD"] = fields[13]
	}
	i++
	if i <= dataLen && fields[14] != "" && fields[14] != "NA" {
		(*doc)["OBS_THR"] = fields[14]
	}
	i++
	if i <= dataLen && fields[15] != "" && fields[15] != "NA" {
		(*doc)["FCST_VAR"] = fields[15]
	}
	i++
	if i <= dataLen && fields[16] != "" && fields[16] != "NA" {
		(*doc)["FCST_UNITS"] = fields[16]
	}
	i++
	if i <= dataLen && fields[18] != "" && fields[18] != "NA" {
		(*doc)["OBS_VAR"] = fields[18]
	}
	i++
	if i <= dataLen && fields[19] != "" && fields[19] != "NA" {
		(*doc)["OBS_UNITS"] = fields[19]
	}
	i++
	if i <= dataLen && fields[20] != "" && fields[20] != "NA" {
		(*doc)["OBS_LEV"] = fields[20]
	}
	i++
	if i <= dataLen && fields[21] != "" && fields[21] != "NA" {
		(*doc)["OBTYPE"] = fields[21]
	}
}

func (s *STAT_DMAP) fill_STAT_DMAP_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	if i <= dataLen && fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	i++
	if i <= dataLen && fields[1] != "" && fields[1] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	i++
	if i <= dataLen && fields[2] != "" && fields[2] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	i++
	if i <= dataLen && fields[4] != "" && fields[4] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	i++
	if i <= dataLen && fields[5] != "" && fields[5] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	i++
	if i <= dataLen && fields[6] != "" && fields[6] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	i++
	if i <= dataLen && fields[7] != "" && fields[7] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	i++
	if i <= dataLen && fields[8] != "" && fields[8] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	i++
	if i <= dataLen && fields[9] != "" && fields[9] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	i++
	if i <= dataLen && fields[10] != "" && fields[10] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	i++
	if i <= dataLen && fields[11] != "" && fields[11] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	i++
	if i <= dataLen && fields[12] != "" && fields[12] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	i++
	if i <= dataLen && fields[13] != "" && fields[13] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	i++
	if i <= dataLen && fields[14] != "" && fields[14] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	i++
	if i <= dataLen && fields[15] != "" && fields[15] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	i++
	if i <= dataLen && fields[16] != "" && fields[16] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	i++
	if i <= dataLen && fields[17] != "" && fields[17] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	i++
	if i <= dataLen && fields[18] != "" && fields[18] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	i++
	if i <= dataLen && fields[19] != "" && fields[19] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	i++
	if i <= dataLen && fields[20] != "" && fields[20] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	i++
	if i <= dataLen && fields[21] != "" && fields[21] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	i++
	if i <= dataLen && fields[22] != "" && fields[22] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	i++
	if i <= dataLen && fields[23] != "" && fields[23] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_PCT) fill_STAT_PCT_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	if i <= dataLen && fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	i++
	if i <= dataLen && fields[1] != "" && fields[1] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	i++
	if i <= dataLen && fields[2] != "" && fields[2] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	i++
	if i <= dataLen && fields[4] != "" && fields[4] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	i++
	if i <= dataLen && fields[5] != "" && fields[5] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	i++
	if i <= dataLen && fields[6] != "" && fields[6] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	i++
	if i <= dataLen && fields[7] != "" && fields[7] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	i++
	if i <= dataLen && fields[8] != "" && fields[8] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	i++
	if i <= dataLen && fields[9] != "" && fields[9] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	i++
	if i <= dataLen && fields[10] != "" && fields[10] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	i++
	if i <= dataLen && fields[11] != "" && fields[11] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	i++
	if i <= dataLen && fields[12] != "" && fields[12] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	i++
	if i <= dataLen && fields[13] != "" && fields[13] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	i++
	if i <= dataLen && fields[14] != "" && fields[14] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	i++
	if i <= dataLen && fields[15] != "" && fields[15] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	i++
	if i <= dataLen && fields[16] != "" && fields[16] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	i++
	if i <= dataLen && fields[17] != "" && fields[17] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	i++
	if i <= dataLen && fields[18] != "" && fields[18] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	i++
	if i <= dataLen && fields[19] != "" && fields[19] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	i++
	if i <= dataLen && fields[20] != "" && fields[20] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	i++
	if i <= dataLen && fields[21] != "" && fields[21] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	i++
	if i <= dataLen && fields[22] != "" && fields[22] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	i++
	if i <= dataLen && fields[23] != "" && fields[23] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_SAL1L2) fill_STAT_SAL1L2_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	if i <= dataLen && fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	i++
	if i <= dataLen && fields[1] != "" && fields[1] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	i++
	if i <= dataLen && fields[2] != "" && fields[2] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	i++
	if i <= dataLen && fields[4] != "" && fields[4] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	i++
	if i <= dataLen && fields[5] != "" && fields[5] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	i++
	if i <= dataLen && fields[6] != "" && fields[6] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	i++
	if i <= dataLen && fields[7] != "" && fields[7] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	i++
	if i <= dataLen && fields[8] != "" && fields[8] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	i++
	if i <= dataLen && fields[9] != "" && fields[9] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	i++
	if i <= dataLen && fields[10] != "" && fields[10] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	i++
	if i <= dataLen && fields[11] != "" && fields[11] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	i++
	if i <= dataLen && fields[12] != "" && fields[12] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	i++
	if i <= dataLen && fields[13] != "" && fields[13] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	i++
	if i <= dataLen && fields[14] != "" && fields[14] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	i++
	if i <= dataLen && fields[15] != "" && fields[15] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	i++
	if i <= dataLen && fields[16] != "" && fields[16] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	i++
	if i <= dataLen && fields[17] != "" && fields[17] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	i++
	if i <= dataLen && fields[18] != "" && fields[18] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	i++
	if i <= dataLen && fields[19] != "" && fields[19] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	i++
	if i <= dataLen && fields[20] != "" && fields[20] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	i++
	if i <= dataLen && fields[21] != "" && fields[21] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	i++
	if i <= dataLen && fields[22] != "" && fields[22] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	i++
	if i <= dataLen && fields[23] != "" && fields[23] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *MTD_2DSINGLE) fill_MTD_2DSINGLE_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	if i <= dataLen && fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	i++
	if i <= dataLen && fields[1] != "" && fields[1] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	i++
	if i <= dataLen && fields[2] != "" && fields[2] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	i++
	if i <= dataLen && fields[4] != "" && fields[4] != "NA" {
		(*doc)["FCST_VALID"] = fields[4]
	}
	i++
	if i <= dataLen && fields[5] != "" && fields[5] != "NA" {
		(*doc)["OBS_LEAD"] = fields[5]
	}
	i++
	if i <= dataLen && fields[6] != "" && fields[6] != "NA" {
		(*doc)["OBS_VALID"] = fields[6]
	}
	i++
	if i <= dataLen && fields[7] != "" && fields[7] != "NA" {
		(*doc)["T_DELTA"] = fields[7]
	}
	i++
	if i <= dataLen && fields[8] != "" && fields[8] != "NA" {
		(*doc)["FCST_T_BEG"] = fields[8]
	}
	i++
	if i <= dataLen && fields[9] != "" && fields[9] != "NA" {
		(*doc)["FCST_T_END"] = fields[9]
	}
	i++
	if i <= dataLen && fields[10] != "" && fields[10] != "NA" {
		(*doc)["FCST_RAD"] = fields[10]
	}
	i++
	if i <= dataLen && fields[11] != "" && fields[11] != "NA" {
		(*doc)["FCST_THR"] = fields[11]
	}
	i++
	if i <= dataLen && fields[12] != "" && fields[12] != "NA" {
		(*doc)["OBS_T_BEG"] = fields[12]
	}
	i++
	if i <= dataLen && fields[13] != "" && fields[13] != "NA" {
		(*doc)["OBS_T_END"] = fields[13]
	}
	i++
	if i <= dataLen && fields[14] != "" && fields[14] != "NA" {
		(*doc)["OBS_RAD"] = fields[14]
	}
	i++
	if i <= dataLen && fields[15] != "" && fields[15] != "NA" {
		(*doc)["OBS_THR"] = fields[15]
	}
	i++
	if i <= dataLen && fields[16] != "" && fields[16] != "NA" {
		(*doc)["FCST_VAR"] = fields[16]
	}
	i++
	if i <= dataLen && fields[17] != "" && fields[17] != "NA" {
		(*doc)["FCST_UNITS"] = fields[17]
	}
	i++
	if i <= dataLen && fields[19] != "" && fields[19] != "NA" {
		(*doc)["OBS_VAR"] = fields[19]
	}
	i++
	if i <= dataLen && fields[20] != "" && fields[20] != "NA" {
		(*doc)["OBS_UNITS"] = fields[20]
	}
	i++
	if i <= dataLen && fields[21] != "" && fields[21] != "NA" {
		(*doc)["OBS_LEV"] = fields[21]
	}
}

func (s *MTD_3DSINGLE) fill_MTD_3DSINGLE_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	if i <= dataLen && fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	i++
	if i <= dataLen && fields[1] != "" && fields[1] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	i++
	if i <= dataLen && fields[2] != "" && fields[2] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	i++
	if i <= dataLen && fields[4] != "" && fields[4] != "NA" {
		(*doc)["FCST_VALID"] = fields[4]
	}
	i++
	if i <= dataLen && fields[5] != "" && fields[5] != "NA" {
		(*doc)["OBS_LEAD"] = fields[5]
	}
	i++
	if i <= dataLen && fields[6] != "" && fields[6] != "NA" {
		(*doc)["OBS_VALID"] = fields[6]
	}
	i++
	if i <= dataLen && fields[7] != "" && fields[7] != "NA" {
		(*doc)["T_DELTA"] = fields[7]
	}
	i++
	if i <= dataLen && fields[8] != "" && fields[8] != "NA" {
		(*doc)["FCST_T_BEG"] = fields[8]
	}
	i++
	if i <= dataLen && fields[9] != "" && fields[9] != "NA" {
		(*doc)["FCST_T_END"] = fields[9]
	}
	i++
	if i <= dataLen && fields[10] != "" && fields[10] != "NA" {
		(*doc)["FCST_RAD"] = fields[10]
	}
	i++
	if i <= dataLen && fields[11] != "" && fields[11] != "NA" {
		(*doc)["FCST_THR"] = fields[11]
	}
	i++
	if i <= dataLen && fields[12] != "" && fields[12] != "NA" {
		(*doc)["OBS_T_BEG"] = fields[12]
	}
	i++
	if i <= dataLen && fields[13] != "" && fields[13] != "NA" {
		(*doc)["OBS_T_END"] = fields[13]
	}
	i++
	if i <= dataLen && fields[14] != "" && fields[14] != "NA" {
		(*doc)["OBS_RAD"] = fields[14]
	}
	i++
	if i <= dataLen && fields[15] != "" && fields[15] != "NA" {
		(*doc)["OBS_THR"] = fields[15]
	}
	i++
	if i <= dataLen && fields[16] != "" && fields[16] != "NA" {
		(*doc)["FCST_VAR"] = fields[16]
	}
	i++
	if i <= dataLen && fields[17] != "" && fields[17] != "NA" {
		(*doc)["FCST_UNITS"] = fields[17]
	}
	i++
	if i <= dataLen && fields[19] != "" && fields[19] != "NA" {
		(*doc)["OBS_VAR"] = fields[19]
	}
	i++
	if i <= dataLen && fields[20] != "" && fields[20] != "NA" {
		(*doc)["OBS_UNITS"] = fields[20]
	}
	i++
	if i <= dataLen && fields[21] != "" && fields[21] != "NA" {
		(*doc)["OBS_LEV"] = fields[21]
	}
}

func (s *MTD_3DPAIR) fill_MTD_3DPAIR_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	if i <= dataLen && fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	i++
	if i <= dataLen && fields[1] != "" && fields[1] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	i++
	if i <= dataLen && fields[2] != "" && fields[2] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	i++
	if i <= dataLen && fields[4] != "" && fields[4] != "NA" {
		(*doc)["FCST_VALID"] = fields[4]
	}
	i++
	if i <= dataLen && fields[5] != "" && fields[5] != "NA" {
		(*doc)["OBS_LEAD"] = fields[5]
	}
	i++
	if i <= dataLen && fields[6] != "" && fields[6] != "NA" {
		(*doc)["OBS_VALID"] = fields[6]
	}
	i++
	if i <= dataLen && fields[7] != "" && fields[7] != "NA" {
		(*doc)["T_DELTA"] = fields[7]
	}
	i++
	if i <= dataLen && fields[8] != "" && fields[8] != "NA" {
		(*doc)["FCST_T_BEG"] = fields[8]
	}
	i++
	if i <= dataLen && fields[9] != "" && fields[9] != "NA" {
		(*doc)["FCST_T_END"] = fields[9]
	}
	i++
	if i <= dataLen && fields[10] != "" && fields[10] != "NA" {
		(*doc)["FCST_RAD"] = fields[10]
	}
	i++
	if i <= dataLen && fields[11] != "" && fields[11] != "NA" {
		(*doc)["FCST_THR"] = fields[11]
	}
	i++
	if i <= dataLen && fields[12] != "" && fields[12] != "NA" {
		(*doc)["OBS_T_BEG"] = fields[12]
	}
	i++
	if i <= dataLen && fields[13] != "" && fields[13] != "NA" {
		(*doc)["OBS_T_END"] = fields[13]
	}
	i++
	if i <= dataLen && fields[14] != "" && fields[14] != "NA" {
		(*doc)["OBS_RAD"] = fields[14]
	}
	i++
	if i <= dataLen && fields[15] != "" && fields[15] != "NA" {
		(*doc)["OBS_THR"] = fields[15]
	}
	i++
	if i <= dataLen && fields[16] != "" && fields[16] != "NA" {
		(*doc)["FCST_VAR"] = fields[16]
	}
	i++
	if i <= dataLen && fields[17] != "" && fields[17] != "NA" {
		(*doc)["FCST_UNITS"] = fields[17]
	}
	i++
	if i <= dataLen && fields[19] != "" && fields[19] != "NA" {
		(*doc)["OBS_VAR"] = fields[19]
	}
	i++
	if i <= dataLen && fields[20] != "" && fields[20] != "NA" {
		(*doc)["OBS_UNITS"] = fields[20]
	}
	i++
	if i <= dataLen && fields[21] != "" && fields[21] != "NA" {
		(*doc)["OBS_LEV"] = fields[21]
	}
}

func (s *STAT_CTS) fill_STAT_CTS_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	if i <= dataLen && fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	i++
	if i <= dataLen && fields[1] != "" && fields[1] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	i++
	if i <= dataLen && fields[2] != "" && fields[2] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	i++
	if i <= dataLen && fields[4] != "" && fields[4] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	i++
	if i <= dataLen && fields[5] != "" && fields[5] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	i++
	if i <= dataLen && fields[6] != "" && fields[6] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	i++
	if i <= dataLen && fields[7] != "" && fields[7] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	i++
	if i <= dataLen && fields[8] != "" && fields[8] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	i++
	if i <= dataLen && fields[9] != "" && fields[9] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	i++
	if i <= dataLen && fields[10] != "" && fields[10] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	i++
	if i <= dataLen && fields[11] != "" && fields[11] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	i++
	if i <= dataLen && fields[12] != "" && fields[12] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	i++
	if i <= dataLen && fields[13] != "" && fields[13] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	i++
	if i <= dataLen && fields[14] != "" && fields[14] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	i++
	if i <= dataLen && fields[15] != "" && fields[15] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	i++
	if i <= dataLen && fields[16] != "" && fields[16] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	i++
	if i <= dataLen && fields[17] != "" && fields[17] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	i++
	if i <= dataLen && fields[18] != "" && fields[18] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	i++
	if i <= dataLen && fields[19] != "" && fields[19] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	i++
	if i <= dataLen && fields[20] != "" && fields[20] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	i++
	if i <= dataLen && fields[21] != "" && fields[21] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	i++
	if i <= dataLen && fields[22] != "" && fields[22] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	i++
	if i <= dataLen && fields[23] != "" && fields[23] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_SEEPS_MPR) fill_STAT_SEEPS_MPR_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	if i <= dataLen && fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	i++
	if i <= dataLen && fields[1] != "" && fields[1] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	i++
	if i <= dataLen && fields[2] != "" && fields[2] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	i++
	if i <= dataLen && fields[4] != "" && fields[4] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	i++
	if i <= dataLen && fields[5] != "" && fields[5] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	i++
	if i <= dataLen && fields[6] != "" && fields[6] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	i++
	if i <= dataLen && fields[7] != "" && fields[7] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	i++
	if i <= dataLen && fields[8] != "" && fields[8] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	i++
	if i <= dataLen && fields[9] != "" && fields[9] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	i++
	if i <= dataLen && fields[10] != "" && fields[10] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	i++
	if i <= dataLen && fields[11] != "" && fields[11] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	i++
	if i <= dataLen && fields[12] != "" && fields[12] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	i++
	if i <= dataLen && fields[13] != "" && fields[13] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	i++
	if i <= dataLen && fields[14] != "" && fields[14] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	i++
	if i <= dataLen && fields[15] != "" && fields[15] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	i++
	if i <= dataLen && fields[16] != "" && fields[16] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	i++
	if i <= dataLen && fields[17] != "" && fields[17] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	i++
	if i <= dataLen && fields[18] != "" && fields[18] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	i++
	if i <= dataLen && fields[19] != "" && fields[19] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	i++
	if i <= dataLen && fields[20] != "" && fields[20] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	i++
	if i <= dataLen && fields[21] != "" && fields[21] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	i++
	if i <= dataLen && fields[22] != "" && fields[22] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	i++
	if i <= dataLen && fields[23] != "" && fields[23] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_NBRCNT) fill_STAT_NBRCNT_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	if i <= dataLen && fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	i++
	if i <= dataLen && fields[1] != "" && fields[1] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	i++
	if i <= dataLen && fields[2] != "" && fields[2] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	i++
	if i <= dataLen && fields[4] != "" && fields[4] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	i++
	if i <= dataLen && fields[5] != "" && fields[5] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	i++
	if i <= dataLen && fields[6] != "" && fields[6] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	i++
	if i <= dataLen && fields[7] != "" && fields[7] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	i++
	if i <= dataLen && fields[8] != "" && fields[8] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	i++
	if i <= dataLen && fields[9] != "" && fields[9] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	i++
	if i <= dataLen && fields[10] != "" && fields[10] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	i++
	if i <= dataLen && fields[11] != "" && fields[11] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	i++
	if i <= dataLen && fields[12] != "" && fields[12] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	i++
	if i <= dataLen && fields[13] != "" && fields[13] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	i++
	if i <= dataLen && fields[14] != "" && fields[14] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	i++
	if i <= dataLen && fields[15] != "" && fields[15] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	i++
	if i <= dataLen && fields[16] != "" && fields[16] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	i++
	if i <= dataLen && fields[17] != "" && fields[17] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	i++
	if i <= dataLen && fields[18] != "" && fields[18] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	i++
	if i <= dataLen && fields[19] != "" && fields[19] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	i++
	if i <= dataLen && fields[20] != "" && fields[20] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	i++
	if i <= dataLen && fields[21] != "" && fields[21] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	i++
	if i <= dataLen && fields[22] != "" && fields[22] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	i++
	if i <= dataLen && fields[23] != "" && fields[23] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_PHIST) fill_STAT_PHIST_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	if i <= dataLen && fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	i++
	if i <= dataLen && fields[1] != "" && fields[1] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	i++
	if i <= dataLen && fields[2] != "" && fields[2] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	i++
	if i <= dataLen && fields[4] != "" && fields[4] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	i++
	if i <= dataLen && fields[5] != "" && fields[5] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	i++
	if i <= dataLen && fields[6] != "" && fields[6] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	i++
	if i <= dataLen && fields[7] != "" && fields[7] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	i++
	if i <= dataLen && fields[8] != "" && fields[8] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	i++
	if i <= dataLen && fields[9] != "" && fields[9] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	i++
	if i <= dataLen && fields[10] != "" && fields[10] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	i++
	if i <= dataLen && fields[11] != "" && fields[11] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	i++
	if i <= dataLen && fields[12] != "" && fields[12] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	i++
	if i <= dataLen && fields[13] != "" && fields[13] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	i++
	if i <= dataLen && fields[14] != "" && fields[14] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	i++
	if i <= dataLen && fields[15] != "" && fields[15] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	i++
	if i <= dataLen && fields[16] != "" && fields[16] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	i++
	if i <= dataLen && fields[17] != "" && fields[17] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	i++
	if i <= dataLen && fields[18] != "" && fields[18] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	i++
	if i <= dataLen && fields[19] != "" && fields[19] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	i++
	if i <= dataLen && fields[20] != "" && fields[20] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	i++
	if i <= dataLen && fields[21] != "" && fields[21] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	i++
	if i <= dataLen && fields[22] != "" && fields[22] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	i++
	if i <= dataLen && fields[23] != "" && fields[23] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *TCST_TCDIAG) fill_TCST_TCDIAG_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	if i <= dataLen && fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	i++
	if i <= dataLen && fields[1] != "" && fields[1] != "NA" {
		(*doc)["AMODEL"] = fields[1]
	}
	i++
	if i <= dataLen && fields[2] != "" && fields[2] != "NA" {
		(*doc)["BMODEL"] = fields[2]
	}
	i++
	if i <= dataLen && fields[3] != "" && fields[3] != "NA" {
		(*doc)["DESC"] = fields[3]
	}
	i++
	if i <= dataLen && fields[4] != "" && fields[4] != "NA" {
		(*doc)["STORM_ID"] = fields[4]
	}
	i++
	if i <= dataLen && fields[5] != "" && fields[5] != "NA" {
		(*doc)["BASIN"] = fields[5]
	}
	i++
	if i <= dataLen && fields[6] != "" && fields[6] != "NA" {
		(*doc)["CYCLONE"] = fields[6]
	}
	i++
	if i <= dataLen && fields[7] != "" && fields[7] != "NA" {
		(*doc)["STORM_NAME"] = fields[7]
	}
	i++
	if i <= dataLen && fields[8] != "" && fields[8] != "NA" {
		(*doc)["INIT"] = fields[8]
	}
	i++
	if i <= dataLen && fields[9] != "" && fields[9] != "NA" {
		(*doc)["LEAD"] = fields[9]
	}
	i++
	if i <= dataLen && fields[10] != "" && fields[10] != "NA" {
		(*doc)["VALID"] = fields[10]
	}
	i++
	if i <= dataLen && fields[11] != "" && fields[11] != "NA" {
		(*doc)["INIT_MASK"] = fields[11]
	}
	i++
	if i <= dataLen && fields[12] != "" && fields[12] != "NA" {
		(*doc)["VALID_MASK"] = fields[12]
	}
	i++
	if i <= dataLen && fields[13] != "" && fields[13] != "NA" {
		(*doc)["LINE_TYPE"] = fields[13]
	}
}

func (s *STAT_MPR) fill_STAT_MPR_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	if i <= dataLen && fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	i++
	if i <= dataLen && fields[1] != "" && fields[1] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	i++
	if i <= dataLen && fields[2] != "" && fields[2] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	i++
	if i <= dataLen && fields[4] != "" && fields[4] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	i++
	if i <= dataLen && fields[5] != "" && fields[5] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	i++
	if i <= dataLen && fields[6] != "" && fields[6] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	i++
	if i <= dataLen && fields[7] != "" && fields[7] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	i++
	if i <= dataLen && fields[8] != "" && fields[8] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	i++
	if i <= dataLen && fields[9] != "" && fields[9] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	i++
	if i <= dataLen && fields[10] != "" && fields[10] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	i++
	if i <= dataLen && fields[11] != "" && fields[11] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	i++
	if i <= dataLen && fields[12] != "" && fields[12] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	i++
	if i <= dataLen && fields[13] != "" && fields[13] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	i++
	if i <= dataLen && fields[14] != "" && fields[14] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	i++
	if i <= dataLen && fields[15] != "" && fields[15] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	i++
	if i <= dataLen && fields[16] != "" && fields[16] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	i++
	if i <= dataLen && fields[17] != "" && fields[17] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	i++
	if i <= dataLen && fields[18] != "" && fields[18] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	i++
	if i <= dataLen && fields[19] != "" && fields[19] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	i++
	if i <= dataLen && fields[20] != "" && fields[20] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	i++
	if i <= dataLen && fields[21] != "" && fields[21] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	i++
	if i <= dataLen && fields[22] != "" && fields[22] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	i++
	if i <= dataLen && fields[23] != "" && fields[23] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_GRAD) fill_STAT_GRAD_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	if i <= dataLen && fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	i++
	if i <= dataLen && fields[1] != "" && fields[1] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	i++
	if i <= dataLen && fields[2] != "" && fields[2] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	i++
	if i <= dataLen && fields[4] != "" && fields[4] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	i++
	if i <= dataLen && fields[5] != "" && fields[5] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	i++
	if i <= dataLen && fields[6] != "" && fields[6] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	i++
	if i <= dataLen && fields[7] != "" && fields[7] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	i++
	if i <= dataLen && fields[8] != "" && fields[8] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	i++
	if i <= dataLen && fields[9] != "" && fields[9] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	i++
	if i <= dataLen && fields[10] != "" && fields[10] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	i++
	if i <= dataLen && fields[11] != "" && fields[11] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	i++
	if i <= dataLen && fields[12] != "" && fields[12] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	i++
	if i <= dataLen && fields[13] != "" && fields[13] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	i++
	if i <= dataLen && fields[14] != "" && fields[14] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	i++
	if i <= dataLen && fields[15] != "" && fields[15] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	i++
	if i <= dataLen && fields[16] != "" && fields[16] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	i++
	if i <= dataLen && fields[17] != "" && fields[17] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	i++
	if i <= dataLen && fields[18] != "" && fields[18] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	i++
	if i <= dataLen && fields[19] != "" && fields[19] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	i++
	if i <= dataLen && fields[20] != "" && fields[20] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	i++
	if i <= dataLen && fields[21] != "" && fields[21] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	i++
	if i <= dataLen && fields[22] != "" && fields[22] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	i++
	if i <= dataLen && fields[23] != "" && fields[23] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *TCST_PROBRIRW) fill_TCST_PROBRIRW_Header(fields []string, doc *map[string]interface{}) {
	dataLen := len(fields)
	i := -1
	// fill the met fields leaving out "" and NA values
	i++
	if i <= dataLen && fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	i++
	if i <= dataLen && fields[1] != "" && fields[1] != "NA" {
		(*doc)["AMODEL"] = fields[1]
	}
	i++
	if i <= dataLen && fields[2] != "" && fields[2] != "NA" {
		(*doc)["BMODEL"] = fields[2]
	}
	i++
	if i <= dataLen && fields[3] != "" && fields[3] != "NA" {
		(*doc)["DESC"] = fields[3]
	}
	i++
	if i <= dataLen && fields[4] != "" && fields[4] != "NA" {
		(*doc)["STORM_ID"] = fields[4]
	}
	i++
	if i <= dataLen && fields[5] != "" && fields[5] != "NA" {
		(*doc)["BASIN"] = fields[5]
	}
	i++
	if i <= dataLen && fields[6] != "" && fields[6] != "NA" {
		(*doc)["CYCLONE"] = fields[6]
	}
	i++
	if i <= dataLen && fields[7] != "" && fields[7] != "NA" {
		(*doc)["STORM_NAME"] = fields[7]
	}
	i++
	if i <= dataLen && fields[8] != "" && fields[8] != "NA" {
		(*doc)["INIT"] = fields[8]
	}
	i++
	if i <= dataLen && fields[9] != "" && fields[9] != "NA" {
		(*doc)["LEAD"] = fields[9]
	}
	i++
	if i <= dataLen && fields[10] != "" && fields[10] != "NA" {
		(*doc)["VALID"] = fields[10]
	}
	i++
	if i <= dataLen && fields[11] != "" && fields[11] != "NA" {
		(*doc)["INIT_MASK"] = fields[11]
	}
	i++
	if i <= dataLen && fields[12] != "" && fields[12] != "NA" {
		(*doc)["VALID_MASK"] = fields[12]
	}
	i++
	if i <= dataLen && fields[13] != "" && fields[13] != "NA" {
		(*doc)["LINE_TYPE"] = fields[13]
	}
}

//line data struct definitions
type STAT_MPR struct {
	TOTAL            int     `json:"total"`
	INDEX            int     `json:"index"`
	OBS_SID          string  `json:"obs_sid"`
	OBS_LAT          float64 `json:"obs_lat"`
	OBS_LON          float64 `json:"obs_lon"`
	OBS_LVL          float64 `json:"obs_lvl"`
	OBS_ELV          float64 `json:"obs_elv"`
	FCST             float64 `json:"fcst"`
	OBS              float64 `json:"obs"`
	OBS_QC           string  `json:"obs_qc"`
	OBS_CLIMO_MEAN   float64 `json:"obs_climo_mean"`
	OBS_CLIMO_STDEV  float64 `json:"obs_climo_stdev"`
	OBS_CLIMO_CDF    float64 `json:"obs_climo_cdf"`
	FCST_CLIMO_MEAN  float64 `json:"fcst_climo_mean"`
	FCST_CLIMO_STDEV float64 `json:"fcst_climo_stdev"`
}

type STAT_DMAP struct {
	TOTAL      int     `json:"total"`
	FY         int     `json:"fy"`
	OY         int     `json:"oy"`
	FBIAS      float64 `json:"fbias"`
	BADDELEY   float64 `json:"baddeley"`
	HAUSDORFF  float64 `json:"hausdorff"`
	MED_FO     float64 `json:"med_fo"`
	MED_OF     float64 `json:"med_of"`
	MED_MIN    float64 `json:"med_min"`
	MED_MAX    float64 `json:"med_max"`
	MED_MEAN   float64 `json:"med_mean"`
	FOM_FO     float64 `json:"fom_fo"`
	FOM_OF     float64 `json:"fom_of"`
	FOM_MIN    float64 `json:"fom_min"`
	FOM_MAX    float64 `json:"fom_max"`
	FOM_MEAN   float64 `json:"fom_mean"`
	ZHU_FO     float64 `json:"zhu_fo"`
	ZHU_OF     float64 `json:"zhu_of"`
	ZHU_MIN    float64 `json:"zhu_min"`
	ZHU_MAX    float64 `json:"zhu_max"`
	ZHU_MEAN   float64 `json:"zhu_mean"`
	G          float64 `json:"g"`
	GBETA      float64 `json:"gbeta"`
	BETA_VALUE float64 `json:"beta_value"`
}

type STAT_SAL1L2 struct {
	TOTAL  int     `json:"total"`
	FABAR  float64 `json:"fabar"`
	OABAR  float64 `json:"oabar"`
	FOABAR float64 `json:"foabar"`
	FFABAR float64 `json:"ffabar"`
	OOABAR float64 `json:"ooabar"`
	MAE    float64 `json:"mae"`
}

type MTD_2DSINGLE struct {
	OBJECT_ID      string  `json:"object_id"`
	OBJECT_CAT     string  `json:"object_cat"`
	TIME_INDEX     int     `json:"time_index"`
	AREA           int     `json:"area"`
	CENTROID_X     float64 `json:"centroid_x"`
	CENTROID_Y     float64 `json:"centroid_y"`
	CENTROID_LAT   float64 `json:"centroid_lat"`
	CENTROID_LON   float64 `json:"centroid_lon"`
	AXIS_ANG       float64 `json:"axis_ang"`
	INTENSITY_10   float64 `json:"intensity_10"`
	INTENSITY_25   float64 `json:"intensity_25"`
	INTENSITY_50   float64 `json:"intensity_50"`
	INTENSITY_75   float64 `json:"intensity_75"`
	INTENSITY_90   float64 `json:"intensity_90"`
	INTENSITY_USER float64 `json:"intensity_user"`
}

type STAT_CTC struct {
	TOTAL    int     `json:"total"`
	FY_OY    float64 `json:"fy_oy"`
	FY_ON    float64 `json:"fy_on"`
	FN_OY    float64 `json:"fn_oy"`
	FN_ON    float64 `json:"fn_on"`
	EC_VALUE float64 `json:"ec_value"`
}

type STAT_ISC struct {
	TOTAL    int     `json:"total"`
	TILE_DIM int     `json:"tile_dim"`
	TILE_XLL int     `json:"tile_xll"`
	TILE_YLL int     `json:"tile_yll"`
	NSCALE   int     `json:"nscale"`
	ISCALE   int     `json:"iscale"`
	MSE      float64 `json:"mse"`
	ISC      float64 `json:"isc"`
	FENERGY2 float64 `json:"fenergy2"`
	OENERGY2 float64 `json:"oenergy2"`
	BASER    float64 `json:"baser"`
	FBIAS    float64 `json:"fbias"`
}

type STAT_MCTC struct {
	TOTAL    int                    `json:"total"`
	CAT      map[string]interface{} `json:"cat"`
	EC_VALUE float64                `json:"ec_value"`
}

type STAT_PJC struct {
	TOTAL  int                    `json:"total"`
	THRESH map[string]interface{} `json:"thresh"`
}

type STAT_ECLV struct {
	TOTAL       int                    `json:"total"`
	BASER       float64                `json:"baser"`
	VALUE_BASER int                    `json:"value_baser"`
	PTS         map[string]interface{} `json:"pts"`
}

type STAT_NBRCTS struct {
	TOTAL     int     `json:"total"`
	BASER     float64 `json:"baser"`
	BASER_NCL float64 `json:"baser_ncl"`
	BASER_NCU float64 `json:"baser_ncu"`
	BASER_BCL float64 `json:"baser_bcl"`
	BASER_BCU float64 `json:"baser_bcu"`
	FMEAN     float64 `json:"fmean"`
	FMEAN_NCL float64 `json:"fmean_ncl"`
	FMEAN_NCU float64 `json:"fmean_ncu"`
	FMEAN_BCL float64 `json:"fmean_bcl"`
	FMEAN_BCU float64 `json:"fmean_bcu"`
	ACC       float64 `json:"acc"`
	ACC_NCL   float64 `json:"acc_ncl"`
	ACC_NCU   float64 `json:"acc_ncu"`
	ACC_BCL   float64 `json:"acc_bcl"`
	ACC_BCU   float64 `json:"acc_bcu"`
	FBIAS     float64 `json:"fbias"`
	FBIAS_BCL float64 `json:"fbias_bcl"`
	FBIAS_BCU float64 `json:"fbias_bcu"`
	PODY      float64 `json:"pody"`
	PODY_NCL  float64 `json:"pody_ncl"`
	PODY_NCU  float64 `json:"pody_ncu"`
	PODY_BCL  float64 `json:"pody_bcl"`
	PODY_BCU  float64 `json:"pody_bcu"`
	PODN      float64 `json:"podn"`
	PODN_NCL  float64 `json:"podn_ncl"`
	PODN_NCU  float64 `json:"podn_ncu"`
	PODN_BCL  float64 `json:"podn_bcl"`
	PODN_BCU  float64 `json:"podn_bcu"`
	POFD      float64 `json:"pofd"`
	POFD_NCL  float64 `json:"pofd_ncl"`
	POFD_NCU  float64 `json:"pofd_ncu"`
	POFD_BCL  float64 `json:"pofd_bcl"`
	POFD_BCU  float64 `json:"pofd_bcu"`
	FAR       float64 `json:"far"`
	FAR_NCL   float64 `json:"far_ncl"`
	FAR_NCU   float64 `json:"far_ncu"`
	FAR_BCL   float64 `json:"far_bcl"`
	FAR_BCU   float64 `json:"far_bcu"`
	CSI       float64 `json:"csi"`
	CSI_NCL   float64 `json:"csi_ncl"`
	CSI_NCU   float64 `json:"csi_ncu"`
	CSI_BCL   float64 `json:"csi_bcl"`
	CSI_BCU   float64 `json:"csi_bcu"`
	GSS       float64 `json:"gss"`
	GSS_BCL   float64 `json:"gss_bcl"`
	GSS_BCU   float64 `json:"gss_bcu"`
	HK        float64 `json:"hk"`
	HK_NCL    float64 `json:"hk_ncl"`
	HK_NCU    float64 `json:"hk_ncu"`
	HK_BCL    float64 `json:"hk_bcl"`
	HK_BCU    float64 `json:"hk_bcu"`
	HSS       float64 `json:"hss"`
	HSS_BCL   float64 `json:"hss_bcl"`
	HSS_BCU   float64 `json:"hss_bcu"`
	ODDS      float64 `json:"odds"`
	ODDS_NCL  float64 `json:"odds_ncl"`
	ODDS_NCU  float64 `json:"odds_ncu"`
	ODDS_BCL  float64 `json:"odds_bcl"`
	ODDS_BCU  float64 `json:"odds_bcu"`
	LODDS     float64 `json:"lodds"`
	LODDS_NCL float64 `json:"lodds_ncl"`
	LODDS_NCU float64 `json:"lodds_ncu"`
	LODDS_BCL float64 `json:"lodds_bcl"`
	LODDS_BCU float64 `json:"lodds_bcu"`
	ORSS      float64 `json:"orss"`
	ORSS_NCL  float64 `json:"orss_ncl"`
	ORSS_NCU  float64 `json:"orss_ncu"`
	ORSS_BCL  float64 `json:"orss_bcl"`
	ORSS_BCU  float64 `json:"orss_bcu"`
	EDS       float64 `json:"eds"`
	EDS_NCL   float64 `json:"eds_ncl"`
	EDS_NCU   float64 `json:"eds_ncu"`
	EDS_BCL   float64 `json:"eds_bcl"`
	EDS_BCU   float64 `json:"eds_bcu"`
	SEDS      float64 `json:"seds"`
	SEDS_NCL  float64 `json:"seds_ncl"`
	SEDS_NCU  float64 `json:"seds_ncu"`
	SEDS_BCL  float64 `json:"seds_bcl"`
	SEDS_BCU  float64 `json:"seds_bcu"`
	EDI       float64 `json:"edi"`
	EDI_NCL   float64 `json:"edi_ncl"`
	EDI_NCU   float64 `json:"edi_ncu"`
	EDI_BCL   float64 `json:"edi_bcl"`
	EDI_BCU   float64 `json:"edi_bcu"`
	SEDI      float64 `json:"sedi"`
	SEDI_NCL  float64 `json:"sedi_ncl"`
	SEDI_NCU  float64 `json:"sedi_ncu"`
	SEDI_BCL  float64 `json:"sedi_bcl"`
	SEDI_BCU  float64 `json:"sedi_bcu"`
	BAGSS     float64 `json:"bagss"`
	BAGSS_BCL float64 `json:"bagss_bcl"`
	BAGSS_BCU float64 `json:"bagss_bcu"`
}

type STAT_GRAD struct {
	TOTAL      int     `json:"total"`
	FGBAR      float64 `json:"fgbar"`
	OGBAR      float64 `json:"ogbar"`
	MGBAR      float64 `json:"mgbar"`
	EGBAR      float64 `json:"egbar"`
	S1         float64 `json:"s1"`
	S1_OG      float64 `json:"s1_og"`
	FGOG_RATIO float64 `json:"fgog_ratio"`
	DX         float64 `json:"dx"`
	DY         float64 `json:"dy"`
}

type TCST_TCDIAG struct {
	TOTAL        int                    `json:"total"`
	INDEX        int                    `json:"index"`
	DIAG_SOURCE  float64                `json:"diag_source"`
	TRACK_SOURCE string                 `json:"track_source"`
	FIELD_SOURCE string                 `json:"field_source"`
	DIAG         map[string]interface{} `json:"diag"`
}

type STAT_RPS struct {
	TOTAL     int     `json:"total"`
	N_PROB    int     `json:"n_prob"`
	RPS_REL   float64 `json:"rps_rel"`
	RPS_RES   float64 `json:"rps_res"`
	RPS_UNC   float64 `json:"rps_unc"`
	RPS       float64 `json:"rps"`
	RPSS      float64 `json:"rpss"`
	RPSS_SMPL float64 `json:"rpss_smpl"`
	RPS_COMP  float64 `json:"rps_comp"`
}

type STAT_RELP struct {
	TOTAL int                    `json:"total"`
	ENS   map[string]interface{} `json:"ens"`
}

type STAT_SSVAR struct {
	TOTAL       int     `json:"total"`
	N_BIN       int     `json:"n_bin"`
	BIN_I       int     `json:"bin_i"`
	BIN_N       int     `json:"bin_n"`
	VAR_MIN     float64 `json:"var_min"`
	VAR_MAX     float64 `json:"var_max"`
	VAR_MEAN    float64 `json:"var_mean"`
	FBAR        float64 `json:"fbar"`
	OBAR        float64 `json:"obar"`
	FOBAR       float64 `json:"fobar"`
	FFBAR       float64 `json:"ffbar"`
	OOBAR       float64 `json:"oobar"`
	FBAR_NCL    float64 `json:"fbar_ncl"`
	FBAR_NCU    float64 `json:"fbar_ncu"`
	FSTDEV      float64 `json:"fstdev"`
	FSTDEV_NCL  float64 `json:"fstdev_ncl"`
	FSTDEV_NCU  float64 `json:"fstdev_ncu"`
	OBAR_NCL    float64 `json:"obar_ncl"`
	OBAR_NCU    float64 `json:"obar_ncu"`
	OSTDEV      float64 `json:"ostdev"`
	OSTDEV_NCL  float64 `json:"ostdev_ncl"`
	OSTDEV_NCU  float64 `json:"ostdev_ncu"`
	PR_CORR     float64 `json:"pr_corr"`
	PR_CORR_NCL float64 `json:"pr_corr_ncl"`
	PR_CORR_NCU float64 `json:"pr_corr_ncu"`
	ME          float64 `json:"me"`
	ME_NCL      float64 `json:"me_ncl"`
	ME_NCU      float64 `json:"me_ncu"`
	ESTDEV      float64 `json:"estdev"`
	ESTDEV_NCL  float64 `json:"estdev_ncl"`
	ESTDEV_NCU  float64 `json:"estdev_ncu"`
	MBIAS       float64 `json:"mbias"`
	MSE         float64 `json:"mse"`
	BCMSE       float64 `json:"bcmse"`
	RMSE        float64 `json:"rmse"`
}

type STAT_CNT struct {
	TOTAL                int     `json:"total"`
	FBAR                 float64 `json:"fbar"`
	FBAR_NCL             float64 `json:"fbar_ncl"`
	FBAR_NCU             float64 `json:"fbar_ncu"`
	FBAR_BCL             float64 `json:"fbar_bcl"`
	FBAR_BCU             float64 `json:"fbar_bcu"`
	FSTDEV               float64 `json:"fstdev"`
	FSTDEV_NCL           float64 `json:"fstdev_ncl"`
	FSTDEV_NCU           float64 `json:"fstdev_ncu"`
	FSTDEV_BCL           float64 `json:"fstdev_bcl"`
	FSTDEV_BCU           float64 `json:"fstdev_bcu"`
	OBAR                 float64 `json:"obar"`
	OBAR_NCL             float64 `json:"obar_ncl"`
	OBAR_NCU             float64 `json:"obar_ncu"`
	OBAR_BCL             float64 `json:"obar_bcl"`
	OBAR_BCU             float64 `json:"obar_bcu"`
	OSTDEV               float64 `json:"ostdev"`
	OSTDEV_NCL           float64 `json:"ostdev_ncl"`
	OSTDEV_NCU           float64 `json:"ostdev_ncu"`
	OSTDEV_BCL           float64 `json:"ostdev_bcl"`
	OSTDEV_BCU           float64 `json:"ostdev_bcu"`
	PR_CORR              float64 `json:"pr_corr"`
	PR_CORR_NCL          float64 `json:"pr_corr_ncl"`
	PR_CORR_NCU          float64 `json:"pr_corr_ncu"`
	PR_CORR_BCL          float64 `json:"pr_corr_bcl"`
	PR_CORR_BCU          float64 `json:"pr_corr_bcu"`
	SP_CORR              float64 `json:"sp_corr"`
	KT_CORR              float64 `json:"kt_corr"`
	RANKS                int     `json:"ranks"`
	FRANK_TIES           int     `json:"frank_ties"`
	ORANK_TIES           int     `json:"orank_ties"`
	ME                   float64 `json:"me"`
	ME_NCL               float64 `json:"me_ncl"`
	ME_NCU               float64 `json:"me_ncu"`
	ME_BCL               float64 `json:"me_bcl"`
	ME_BCU               float64 `json:"me_bcu"`
	ESTDEV               float64 `json:"estdev"`
	ESTDEV_NCL           float64 `json:"estdev_ncl"`
	ESTDEV_NCU           float64 `json:"estdev_ncu"`
	ESTDEV_BCL           float64 `json:"estdev_bcl"`
	ESTDEV_BCU           float64 `json:"estdev_bcu"`
	MBIAS                float64 `json:"mbias"`
	MBIAS_BCL            float64 `json:"mbias_bcl"`
	MBIAS_BCU            float64 `json:"mbias_bcu"`
	MAE                  float64 `json:"mae"`
	MAE_BCL              float64 `json:"mae_bcl"`
	MAE_BCU              float64 `json:"mae_bcu"`
	MSE                  float64 `json:"mse"`
	MSE_BCL              float64 `json:"mse_bcl"`
	MSE_BCU              float64 `json:"mse_bcu"`
	BCMSE                float64 `json:"bcmse"`
	BCMSE_BCL            float64 `json:"bcmse_bcl"`
	BCMSE_BCU            float64 `json:"bcmse_bcu"`
	RMSE                 float64 `json:"rmse"`
	RMSE_BCL             float64 `json:"rmse_bcl"`
	RMSE_BCU             float64 `json:"rmse_bcu"`
	E10                  float64 `json:"e10"`
	E10_BCL              float64 `json:"e10_bcl"`
	E10_BCU              float64 `json:"e10_bcu"`
	E25                  float64 `json:"e25"`
	E25_BCL              float64 `json:"e25_bcl"`
	E25_BCU              float64 `json:"e25_bcu"`
	E50                  float64 `json:"e50"`
	E50_BCL              float64 `json:"e50_bcl"`
	E50_BCU              float64 `json:"e50_bcu"`
	E75                  float64 `json:"e75"`
	E75_BCL              float64 `json:"e75_bcl"`
	E75_BCU              float64 `json:"e75_bcu"`
	E90                  float64 `json:"e90"`
	E90_BCL              float64 `json:"e90_bcl"`
	E90_BCU              float64 `json:"e90_bcu"`
	EIQR                 float64 `json:"eiqr"`
	EIQR_BCL             float64 `json:"eiqr_bcl"`
	EIQR_BCU             float64 `json:"eiqr_bcu"`
	MAD                  float64 `json:"mad"`
	MAD_BCL              float64 `json:"mad_bcl"`
	MAD_BCU              float64 `json:"mad_bcu"`
	ANOM_CORR            float64 `json:"anom_corr"`
	ANOM_CORR_NCL        float64 `json:"anom_corr_ncl"`
	ANOM_CORR_NCU        float64 `json:"anom_corr_ncu"`
	ANOM_CORR_BCL        float64 `json:"anom_corr_bcl"`
	ANOM_CORR_BCU        float64 `json:"anom_corr_bcu"`
	ME2                  float64 `json:"me2"`
	ME2_BCL              float64 `json:"me2_bcl"`
	ME2_BCU              float64 `json:"me2_bcu"`
	MSESS                float64 `json:"msess"`
	MSESS_BCL            float64 `json:"msess_bcl"`
	MSESS_BCU            float64 `json:"msess_bcu"`
	RMSFA                float64 `json:"rmsfa"`
	RMSFA_BCL            float64 `json:"rmsfa_bcl"`
	RMSFA_BCU            float64 `json:"rmsfa_bcu"`
	RMSOA                float64 `json:"rmsoa"`
	RMSOA_BCL            float64 `json:"rmsoa_bcl"`
	RMSOA_BCU            float64 `json:"rmsoa_bcu"`
	ANOM_CORR_UNCNTR     float64 `json:"anom_corr_uncntr"`
	ANOM_CORR_UNCNTR_BCL float64 `json:"anom_corr_uncntr_bcl"`
	ANOM_CORR_UNCNTR_BCU float64 `json:"anom_corr_uncntr_bcu"`
	SI                   float64 `json:"si"`
	SI_BCL               float64 `json:"si_bcl"`
	SI_BCU               float64 `json:"si_bcu"`
}

type STAT_MCTS struct {
	TOTAL      int     `json:"total"`
	N_CAT      int     `json:"n_cat"`
	ACC        float64 `json:"acc"`
	ACC_NCL    float64 `json:"acc_ncl"`
	ACC_NCU    float64 `json:"acc_ncu"`
	ACC_BCL    float64 `json:"acc_bcl"`
	ACC_BCU    float64 `json:"acc_bcu"`
	HK         float64 `json:"hk"`
	HK_BCL     float64 `json:"hk_bcl"`
	HK_BCU     float64 `json:"hk_bcu"`
	HSS        float64 `json:"hss"`
	HSS_BCL    float64 `json:"hss_bcl"`
	HSS_BCU    float64 `json:"hss_bcu"`
	GER        float64 `json:"ger"`
	GER_BCL    float64 `json:"ger_bcl"`
	GER_BCU    float64 `json:"ger_bcu"`
	HSS_EC     float64 `json:"hss_ec"`
	HSS_EC_BCL float64 `json:"hss_ec_bcl"`
	HSS_EC_BCU float64 `json:"hss_ec_bcu"`
	EC_VALUE   float64 `json:"ec_value"`
}

type STAT_SEEPS struct {
	TOTAL     int     `json:"total"`
	ODFL      float64 `json:"odfl"`
	ODFH      float64 `json:"odfh"`
	OLFD      float64 `json:"olfd"`
	OLFH      float64 `json:"olfh"`
	OHFD      float64 `json:"ohfd"`
	OHFL      float64 `json:"ohfl"`
	PF1       float64 `json:"pf1"`
	PF2       float64 `json:"pf2"`
	PF3       float64 `json:"pf3"`
	PV1       float64 `json:"pv1"`
	PV2       float64 `json:"pv2"`
	PV3       float64 `json:"pv3"`
	MEAN_FCST float64 `json:"mean_fcst"`
	MEAN_OBS  float64 `json:"mean_obs"`
	SEEPS     float64 `json:"seeps"`
}

type STAT_SEEPS_MPR struct {
	OBS_SID  string  `json:"obs_sid"`
	OBS_LAT  float64 `json:"obs_lat"`
	OBS_LON  float64 `json:"obs_lon"`
	FCST     float64 `json:"fcst"`
	OBS      float64 `json:"obs"`
	OBS_QC   string  `json:"obs_qc"`
	FCST_CAT int     `json:"fcst_cat"`
	OBS_CAT  int     `json:"obs_cat"`
	P1       float64 `json:"p1"`
	P2       float64 `json:"p2"`
	T1       float64 `json:"t1"`
	T2       float64 `json:"t2"`
	SEEPS    float64 `json:"seeps"`
}

type STAT_NBRCTC struct {
	TOTAL int     `json:"total"`
	FY_OY float64 `json:"fy_oy"`
	FY_ON float64 `json:"fy_on"`
	FN_OY float64 `json:"fn_oy"`
	FN_ON float64 `json:"fn_on"`
}

type MODE_CTS struct {
	FIELD string  `json:"field"`
	TOTAL int     `json:"total"`
	FY_OY float64 `json:"fy_oy"`
	FY_ON float64 `json:"fy_on"`
	FN_OY float64 `json:"fn_oy"`
	FN_ON float64 `json:"fn_on"`
	BASER float64 `json:"baser"`
	FMEAN float64 `json:"fmean"`
	ACC   float64 `json:"acc"`
	FBIAS float64 `json:"fbias"`
	PODY  float64 `json:"pody"`
	PODN  float64 `json:"podn"`
	POFD  float64 `json:"pofd"`
	FAR   float64 `json:"far"`
	CSI   float64 `json:"csi"`
	GSS   float64 `json:"gss"`
	HK    float64 `json:"hk"`
	HSS   float64 `json:"hss"`
	ODDS  float64 `json:"odds"`
	LODDS float64 `json:"lodds"`
	ORSS  float64 `json:"orss"`
	EDS   float64 `json:"eds"`
	SEDS  float64 `json:"seds"`
	EDI   float64 `json:"edi"`
	SEDI  float64 `json:"sedi"`
	BAGSS float64 `json:"bagss"`
}

type MTD_3DSINGLE struct {
	OBJECT_ID       string  `json:"object_id"`
	OBJECT_CAT      string  `json:"object_cat"`
	CENTROID_X      float64 `json:"centroid_x"`
	CENTROID_Y      float64 `json:"centroid_y"`
	CENTROID_T      float64 `json:"centroid_t"`
	CENTROID_LAT    float64 `json:"centroid_lat"`
	CENTROID_LON    float64 `json:"centroid_lon"`
	X_DOT           float64 `json:"x_dot"`
	Y_DOT           float64 `json:"y_dot"`
	AXIS_ANG        float64 `json:"axis_ang"`
	VOLUME          int     `json:"volume"`
	START_TIME      int     `json:"start_time"`
	END_TIME        int     `json:"end_time"`
	CDIST_TRAVELLED float64 `json:"cdist_travelled"`
	INTENSITY_10    float64 `json:"intensity_10"`
	INTENSITY_25    float64 `json:"intensity_25"`
	INTENSITY_50    float64 `json:"intensity_50"`
	INTENSITY_75    float64 `json:"intensity_75"`
	INTENSITY_90    float64 `json:"intensity_90"`
	INTENSITY_USER  float64 `json:"intensity_user"`
}

type TCST_PROBRIRW struct {
	ALAT        float64                `json:"alat"`
	ALON        float64                `json:"alon"`
	BLAT        float64                `json:"blat"`
	BLON        float64                `json:"blon"`
	INITIALS    string                 `json:"initials"`
	TK_ERR      float64                `json:"tk_err"`
	X_ERR       float64                `json:"x_err"`
	Y_ERR       float64                `json:"y_err"`
	ADLAND      float64                `json:"adland"`
	BDLAND      float64                `json:"bdland"`
	RIRW_BEG    int                    `json:"rirw_beg"`
	RIRW_END    int                    `json:"rirw_end"`
	RIRW_WINDOW int                    `json:"rirw_window"`
	AWIND_END   float64                `json:"awind_end"`
	BWIND_BEG   float64                `json:"bwind_beg"`
	BWIND_END   float64                `json:"bwind_end"`
	BDELTA      float64                `json:"bdelta"`
	BDELTA_MAX  float64                `json:"bdelta_max"`
	BLEVEL_BEG  string                 `json:"blevel_beg"`
	BLEVEL_END  string                 `json:"blevel_end"`
	THRESH      map[string]interface{} `json:"thresh"`
}

type STAT_PCT struct {
	TOTAL  int                    `json:"total"`
	THRESH map[string]interface{} `json:"thresh"`
}

type STAT_PRC struct {
	TOTAL  int                    `json:"total"`
	THRESH map[string]interface{} `json:"thresh"`
}

type STAT_RHIST struct {
	TOTAL int                    `json:"total"`
	RANK  map[string]interface{} `json:"rank"`
}

type STAT_VL1L2 struct {
	TOTAL       int     `json:"total"`
	UFBAR       float64 `json:"ufbar"`
	VFBAR       float64 `json:"vfbar"`
	UOBAR       float64 `json:"uobar"`
	VOBAR       float64 `json:"vobar"`
	UVFOBAR     float64 `json:"uvfobar"`
	UVFFBAR     float64 `json:"uvffbar"`
	UVOOBAR     float64 `json:"uvoobar"`
	F_SPEED_BAR float64 `json:"f_speed_bar"`
	O_SPEED_BAR float64 `json:"o_speed_bar"`
	TOTAL_DIR   float64 `json:"total_dir"`
	DIR_ME      float64 `json:"dir_me"`
	DIR_MAE     float64 `json:"dir_mae"`
	DIR_MSE     float64 `json:"dir_mse"`
}

type MTD_3DPAIR struct {
	OBJECT_ID           string  `json:"object_id"`
	OBJECT_CAT          string  `json:"object_cat"`
	SPACE_CENTROID_DIST float64 `json:"space_centroid_dist"`
	TIME_CENTROID_DELTA float64 `json:"time_centroid_delta"`
	AXIS_DIFF           float64 `json:"axis_diff"`
	SPEED_DELTA         float64 `json:"speed_delta"`
	DIRECTION_DIFF      float64 `json:"direction_diff"`
	VOLUME_RATIO        float64 `json:"volume_ratio"`
	START_TIME_DELTA    int     `json:"start_time_delta"`
	END_TIME_DELTA      int     `json:"end_time_delta"`
	INTERSECTION_VOLUME float64 `json:"intersection_volume"`
	DURATION_DIFF       float64 `json:"duration_diff"`
	INTEREST            float64 `json:"interest"`
}

type STAT_PHIST struct {
	TOTAL    int                    `json:"total"`
	BIN_SIZE int                    `json:"bin_size"`
	BIN      map[string]interface{} `json:"bin"`
}

type STAT_GENMPR struct {
	TOTAL      int     `json:"total"`
	INDEX      int     `json:"index"`
	STORM_ID   string  `json:"storm_id"`
	PROB_LEAD  float64 `json:"prob_lead"`
	PROB_VAL   float64 `json:"prob_val"`
	AGEN_INIT  string  `json:"agen_init"`
	AGEN_FHR   string  `json:"agen_fhr"`
	AGEN_LAT   float64 `json:"agen_lat"`
	AGEN_LON   float64 `json:"agen_lon"`
	AGEN_DLAND float64 `json:"agen_dland"`
	BGEN_LAT   float64 `json:"bgen_lat"`
	BGEN_LON   float64 `json:"bgen_lon"`
	BGEN_DLAND float64 `json:"bgen_dland"`
	GEN_DIST   float64 `json:"gen_dist"`
	GEN_TDIFF  string  `json:"gen_tdiff"`
	INIT_TDIFF string  `json:"init_tdiff"`
	DEV_CAT    string  `json:"dev_cat"`
	OPS_CAT    string  `json:"ops_cat"`
}

type MODE_OBJ struct {
	OBJECT_ID                  string  `json:"object_id"`
	OBJECT_CAT                 string  `json:"object_cat"`
	CENTROID_X                 float64 `json:"centroid_x"`
	CENTROID_Y                 float64 `json:"centroid_y"`
	CENTROID_LAT               float64 `json:"centroid_lat"`
	CENTROID_LON               float64 `json:"centroid_lon"`
	AXIS_ANG                   float64 `json:"axis_ang"`
	LENGTH                     float64 `json:"length"`
	WIDTH                      float64 `json:"width"`
	AREA                       int     `json:"area"`
	AREA_THRESH                int     `json:"area_thresh"`
	CURVATURE                  float64 `json:"curvature"`
	CURVATURE_X                float64 `json:"curvature_x"`
	CURVATURE_Y                float64 `json:"curvature_y"`
	COMPLEXITY                 float64 `json:"complexity"`
	INTENSITY_10               float64 `json:"intensity_10"`
	INTENSITY_25               float64 `json:"intensity_25"`
	INTENSITY_50               float64 `json:"intensity_50"`
	INTENSITY_75               float64 `json:"intensity_75"`
	INTENSITY_90               float64 `json:"intensity_90"`
	INTENSITY_USER             float64 `json:"intensity_user"`
	INTENSITY_SUM              float64 `json:"intensity_sum"`
	CENTROID_DIST              float64 `json:"centroid_dist"`
	BOUNDARY_DIST              float64 `json:"boundary_dist"`
	CONVEX_HULL_DIST           float64 `json:"convex_hull_dist"`
	ANGLE_DIFF                 float64 `json:"angle_diff"`
	ASPECT_DIFF                float64 `json:"aspect_diff"`
	AREA_RATIO                 float64 `json:"area_ratio"`
	INTERSECTION_AREA          float64 `json:"intersection_area"`
	UNION_AREA                 float64 `json:"union_area"`
	SYMMETRIC_DIFF             float64 `json:"symmetric_diff"`
	INTERSECTION_OVER_AREA     float64 `json:"intersection_over_area"`
	CURVATURE_RATIO            float64 `json:"curvature_ratio"`
	COMPLEXITY_RATIO           float64 `json:"complexity_ratio"`
	PERCENTILE_INTENSITY_RATIO float64 `json:"percentile_intensity_ratio"`
	INTEREST                   float64 `json:"interest"`
}

type STAT_ORANK struct {
	TOTAL            int                    `json:"total"`
	INDEX            int                    `json:"index"`
	OBS_SID          string                 `json:"obs_sid"`
	OBS_LAT          float64                `json:"obs_lat"`
	OBS_LON          float64                `json:"obs_lon"`
	OBS_LVL          float64                `json:"obs_lvl"`
	OBS_ELV          float64                `json:"obs_elv"`
	OBS              float64                `json:"obs"`
	PIT              float64                `json:"pit"`
	RANK             int                    `json:"rank"`
	N_ENS_VLD        int                    `json:"n_ens_vld"`
	ENS              map[string]interface{} `json:"ens"`
	OBS_QC           string                 `json:"obs_qc"`
	ENS_MEAN         int                    `json:"ens_mean"`
	OBS_CLIMO_MEAN   float64                `json:"obs_climo_mean"`
	SPREAD           float64                `json:"spread"`
	ENS_MEAN_OERR    int                    `json:"ens_mean_oerr"`
	SPREAD_OERR      float64                `json:"spread_oerr"`
	SPREAD_PLUS_OERR float64                `json:"spread_plus_oerr"`
	OBS_CLIMO_STDEV  float64                `json:"obs_climo_stdev"`
	FCST_CLIMO_MEAN  float64                `json:"fcst_climo_mean"`
	FCST_CLIMO_STDEV float64                `json:"fcst_climo_stdev"`
}

type STAT_PSTD struct {
	TOTAL       int                    `json:"total"`
	THRESH      map[string]interface{} `json:"thresh"`
	BASER_NCL   float64                `json:"baser_ncl"`
	BASER_NCU   float64                `json:"baser_ncu"`
	RELIABILITY float64                `json:"reliability"`
	RESOLUTION  float64                `json:"resolution"`
	UNCERTAINTY float64                `json:"uncertainty"`
	ROC_AUC     float64                `json:"roc_auc"`
	BRIER       float64                `json:"brier"`
	BRIER_NCL   float64                `json:"brier_ncl"`
	BRIER_NCU   float64                `json:"brier_ncu"`
	BRIERCL     float64                `json:"briercl"`
	BRIERCL_NCL float64                `json:"briercl_ncl"`
	BRIERCL_NCU float64                `json:"briercl_ncu"`
	BSS         float64                `json:"bss"`
	BSS_SMPL    float64                `json:"bss_smpl"`
	THRESH_I    int                    `json:"thresh_i"`
}

type STAT_VAL1L2 struct {
	TOTAL        int     `json:"total"`
	UFABAR       float64 `json:"ufabar"`
	VFABAR       float64 `json:"vfabar"`
	UOABAR       float64 `json:"uoabar"`
	VOABAR       float64 `json:"voabar"`
	UVFOABAR     float64 `json:"uvfoabar"`
	UVFFABAR     float64 `json:"uvffabar"`
	UVOOABAR     float64 `json:"uvooabar"`
	FA_SPEED_BAR float64 `json:"fa_speed_bar"`
	OA_SPEED_BAR float64 `json:"oa_speed_bar"`
	TOTAL_DIR    float64 `json:"total_dir"`
	DIRA_ME      float64 `json:"dira_me"`
	DIRA_MAE     float64 `json:"dira_mae"`
	DIRA_MSE     float64 `json:"dira_mse"`
}

type STAT_VCNT struct {
	TOTAL                int     `json:"total"`
	FBAR                 float64 `json:"fbar"`
	FBAR_BCL             float64 `json:"fbar_bcl"`
	FBAR_BCU             float64 `json:"fbar_bcu"`
	OBAR                 float64 `json:"obar"`
	OBAR_BCL             float64 `json:"obar_bcl"`
	OBAR_BCU             float64 `json:"obar_bcu"`
	FS_RMS               float64 `json:"fs_rms"`
	FS_RMS_BCL           float64 `json:"fs_rms_bcl"`
	FS_RMS_BCU           float64 `json:"fs_rms_bcu"`
	OS_RMS               float64 `json:"os_rms"`
	OS_RMS_BCL           float64 `json:"os_rms_bcl"`
	OS_RMS_BCU           float64 `json:"os_rms_bcu"`
	MSVE                 float64 `json:"msve"`
	MSVE_BCL             float64 `json:"msve_bcl"`
	MSVE_BCU             float64 `json:"msve_bcu"`
	RMSVE                float64 `json:"rmsve"`
	RMSVE_BCL            float64 `json:"rmsve_bcl"`
	RMSVE_BCU            float64 `json:"rmsve_bcu"`
	FSTDEV               float64 `json:"fstdev"`
	FSTDEV_BCL           float64 `json:"fstdev_bcl"`
	FSTDEV_BCU           float64 `json:"fstdev_bcu"`
	OSTDEV               float64 `json:"ostdev"`
	OSTDEV_BCL           float64 `json:"ostdev_bcl"`
	OSTDEV_BCU           float64 `json:"ostdev_bcu"`
	FDIR                 float64 `json:"fdir"`
	FDIR_BCL             float64 `json:"fdir_bcl"`
	FDIR_BCU             float64 `json:"fdir_bcu"`
	ODIR                 float64 `json:"odir"`
	ODIR_BCL             float64 `json:"odir_bcl"`
	ODIR_BCU             float64 `json:"odir_bcu"`
	FBAR_SPEED           float64 `json:"fbar_speed"`
	FBAR_SPEED_BCL       float64 `json:"fbar_speed_bcl"`
	FBAR_SPEED_BCU       float64 `json:"fbar_speed_bcu"`
	OBAR_SPEED           float64 `json:"obar_speed"`
	OBAR_SPEED_BCL       float64 `json:"obar_speed_bcl"`
	OBAR_SPEED_BCU       float64 `json:"obar_speed_bcu"`
	VDIFF_SPEED          float64 `json:"vdiff_speed"`
	VDIFF_SPEED_BCL      float64 `json:"vdiff_speed_bcl"`
	VDIFF_SPEED_BCU      float64 `json:"vdiff_speed_bcu"`
	VDIFF_DIR            float64 `json:"vdiff_dir"`
	VDIFF_DIR_BCL        float64 `json:"vdiff_dir_bcl"`
	VDIFF_DIR_BCU        float64 `json:"vdiff_dir_bcu"`
	SPEED_ERR            float64 `json:"speed_err"`
	SPEED_ERR_BCL        float64 `json:"speed_err_bcl"`
	SPEED_ERR_BCU        float64 `json:"speed_err_bcu"`
	SPEED_ABSERR         float64 `json:"speed_abserr"`
	SPEED_ABSERR_BCL     float64 `json:"speed_abserr_bcl"`
	SPEED_ABSERR_BCU     float64 `json:"speed_abserr_bcu"`
	DIR_ERR              float64 `json:"dir_err"`
	DIR_ERR_BCL          float64 `json:"dir_err_bcl"`
	DIR_ERR_BCU          float64 `json:"dir_err_bcu"`
	DIR_ABSERR           float64 `json:"dir_abserr"`
	DIR_ABSERR_BCL       float64 `json:"dir_abserr_bcl"`
	DIR_ABSERR_BCU       float64 `json:"dir_abserr_bcu"`
	ANOM_CORR            float64 `json:"anom_corr"`
	ANOM_CORR_NCL        float64 `json:"anom_corr_ncl"`
	ANOM_CORR_NCU        float64 `json:"anom_corr_ncu"`
	ANOM_CORR_BCL        float64 `json:"anom_corr_bcl"`
	ANOM_CORR_BCU        float64 `json:"anom_corr_bcu"`
	ANOM_CORR_UNCNTR     float64 `json:"anom_corr_uncntr"`
	ANOM_CORR_UNCNTR_BCL float64 `json:"anom_corr_uncntr_bcl"`
	ANOM_CORR_UNCNTR_BCU float64 `json:"anom_corr_uncntr_bcu"`
	TOTAL_DIR            float64 `json:"total_dir"`
	DIR_ME               float64 `json:"dir_me"`
	DIR_ME_BCL           float64 `json:"dir_me_bcl"`
	DIR_ME_BCU           float64 `json:"dir_me_bcu"`
	DIR_MAE              float64 `json:"dir_mae"`
	DIR_MAE_BCL          float64 `json:"dir_mae_bcl"`
	DIR_MAE_BCU          float64 `json:"dir_mae_bcu"`
	DIR_MSE              float64 `json:"dir_mse"`
	DIR_MSE_BCL          float64 `json:"dir_mse_bcl"`
	DIR_MSE_BCU          float64 `json:"dir_mse_bcu"`
	DIR_RMSE             float64 `json:"dir_rmse"`
	DIR_RMSE_BCL         float64 `json:"dir_rmse_bcl"`
	DIR_RMSE_BCU         float64 `json:"dir_rmse_bcu"`
}

type TCST_TCMPR struct {
	TOTAL          int     `json:"total"`
	INDEX          int     `json:"index"`
	LEVEL          string  `json:"level"`
	WATCH_WARN     string  `json:"watch_warn"`
	INITIALS       string  `json:"initials"`
	ALAT           float64 `json:"alat"`
	ALON           float64 `json:"alon"`
	BLAT           float64 `json:"blat"`
	BLON           float64 `json:"blon"`
	TK_ERR         float64 `json:"tk_err"`
	X_ERR          float64 `json:"x_err"`
	Y_ERR          float64 `json:"y_err"`
	ALTK_ERR       float64 `json:"altk_err"`
	CRTK_ERR       float64 `json:"crtk_err"`
	ADLAND         float64 `json:"adland"`
	BDLAND         float64 `json:"bdland"`
	AMSLP          float64 `json:"amslp"`
	BMSLP          float64 `json:"bmslp"`
	AMAX_WIND      float64 `json:"amax_wind"`
	BMAX_WIND      float64 `json:"bmax_wind"`
	AAL_WIND_34    float64 `json:"aal_wind_34"`
	BAL_WIND_34    float64 `json:"bal_wind_34"`
	ANE_WIND_34    float64 `json:"ane_wind_34"`
	BNE_WIND_34    float64 `json:"bne_wind_34"`
	ASE_WIND_34    float64 `json:"ase_wind_34"`
	BSE_WIND_34    float64 `json:"bse_wind_34"`
	ASW_WIND_34    float64 `json:"asw_wind_34"`
	BSW_WIND_34    float64 `json:"bsw_wind_34"`
	ANW_WIND_34    float64 `json:"anw_wind_34"`
	BNW_WIND_34    float64 `json:"bnw_wind_34"`
	AAL_WIND_50    float64 `json:"aal_wind_50"`
	BAL_WIND_50    float64 `json:"bal_wind_50"`
	ANE_WIND_50    float64 `json:"ane_wind_50"`
	BNE_WIND_50    float64 `json:"bne_wind_50"`
	ASE_WIND_50    float64 `json:"ase_wind_50"`
	BSE_WIND_50    float64 `json:"bse_wind_50"`
	ASW_WIND_50    float64 `json:"asw_wind_50"`
	BSW_WIND_50    float64 `json:"bsw_wind_50"`
	ANW_WIND_50    float64 `json:"anw_wind_50"`
	BNW_WIND_50    float64 `json:"bnw_wind_50"`
	AAL_WIND_64    float64 `json:"aal_wind_64"`
	BAL_WIND_64    float64 `json:"bal_wind_64"`
	ANE_WIND_64    float64 `json:"ane_wind_64"`
	BNE_WIND_64    float64 `json:"bne_wind_64"`
	ASE_WIND_64    float64 `json:"ase_wind_64"`
	BSE_WIND_64    float64 `json:"bse_wind_64"`
	ASW_WIND_64    float64 `json:"asw_wind_64"`
	BSW_WIND_64    float64 `json:"bsw_wind_64"`
	ANW_WIND_64    float64 `json:"anw_wind_64"`
	BNW_WIND_64    float64 `json:"bnw_wind_64"`
	ARADP          string  `json:"aradp"`
	BRADP          float64 `json:"bradp"`
	ARRP           int     `json:"arrp"`
	BRRP           float64 `json:"brrp"`
	AMRD           int     `json:"amrd"`
	BMRD           float64 `json:"bmrd"`
	AGUSTS         int     `json:"agusts"`
	BGUSTS         float64 `json:"bgusts"`
	AEYE           int     `json:"aeye"`
	BEYE           float64 `json:"beye"`
	ADIR           int     `json:"adir"`
	BDIR           float64 `json:"bdir"`
	ASPEED         int     `json:"aspeed"`
	BSPEED         float64 `json:"bspeed"`
	ADEPTH         int     `json:"adepth"`
	BDEPTH         float64 `json:"bdepth"`
	NUM_MEMBERS    float64 `json:"num_members"`
	TRACK_SPREAD   float64 `json:"track_spread"`
	TRACK_STDEV    float64 `json:"track_stdev"`
	MSLP_STDEV     float64 `json:"mslp_stdev"`
	MAX_WIND_STDEV float64 `json:"max_wind_stdev"`
}

type STAT_SSIDX struct {
	FCST_MODEL string  `json:"fcst_model"`
	REF_MODEL  string  `json:"ref_model"`
	N_INIT     int     `json:"n_init"`
	N_TERM     int     `json:"n_term"`
	N_VLD      int     `json:"n_vld"`
	SS_INDEX   float64 `json:"ss_index"`
}

type STAT_CTS struct {
	TOTAL      int     `json:"total"`
	BASER      float64 `json:"baser"`
	BASER_NCL  float64 `json:"baser_ncl"`
	BASER_NCU  float64 `json:"baser_ncu"`
	BASER_BCL  float64 `json:"baser_bcl"`
	BASER_BCU  float64 `json:"baser_bcu"`
	FMEAN      float64 `json:"fmean"`
	FMEAN_NCL  float64 `json:"fmean_ncl"`
	FMEAN_NCU  float64 `json:"fmean_ncu"`
	FMEAN_BCL  float64 `json:"fmean_bcl"`
	FMEAN_BCU  float64 `json:"fmean_bcu"`
	ACC        float64 `json:"acc"`
	ACC_NCL    float64 `json:"acc_ncl"`
	ACC_NCU    float64 `json:"acc_ncu"`
	ACC_BCL    float64 `json:"acc_bcl"`
	ACC_BCU    float64 `json:"acc_bcu"`
	FBIAS      float64 `json:"fbias"`
	FBIAS_BCL  float64 `json:"fbias_bcl"`
	FBIAS_BCU  float64 `json:"fbias_bcu"`
	PODY       float64 `json:"pody"`
	PODY_NCL   float64 `json:"pody_ncl"`
	PODY_NCU   float64 `json:"pody_ncu"`
	PODY_BCL   float64 `json:"pody_bcl"`
	PODY_BCU   float64 `json:"pody_bcu"`
	PODN       float64 `json:"podn"`
	PODN_NCL   float64 `json:"podn_ncl"`
	PODN_NCU   float64 `json:"podn_ncu"`
	PODN_BCL   float64 `json:"podn_bcl"`
	PODN_BCU   float64 `json:"podn_bcu"`
	POFD       float64 `json:"pofd"`
	POFD_NCL   float64 `json:"pofd_ncl"`
	POFD_NCU   float64 `json:"pofd_ncu"`
	POFD_BCL   float64 `json:"pofd_bcl"`
	POFD_BCU   float64 `json:"pofd_bcu"`
	FAR        float64 `json:"far"`
	FAR_NCL    float64 `json:"far_ncl"`
	FAR_NCU    float64 `json:"far_ncu"`
	FAR_BCL    float64 `json:"far_bcl"`
	FAR_BCU    float64 `json:"far_bcu"`
	CSI        float64 `json:"csi"`
	CSI_NCL    float64 `json:"csi_ncl"`
	CSI_NCU    float64 `json:"csi_ncu"`
	CSI_BCL    float64 `json:"csi_bcl"`
	CSI_BCU    float64 `json:"csi_bcu"`
	GSS        float64 `json:"gss"`
	GSS_BCL    float64 `json:"gss_bcl"`
	GSS_BCU    float64 `json:"gss_bcu"`
	HK         float64 `json:"hk"`
	HK_NCL     float64 `json:"hk_ncl"`
	HK_NCU     float64 `json:"hk_ncu"`
	HK_BCL     float64 `json:"hk_bcl"`
	HK_BCU     float64 `json:"hk_bcu"`
	HSS        float64 `json:"hss"`
	HSS_BCL    float64 `json:"hss_bcl"`
	HSS_BCU    float64 `json:"hss_bcu"`
	ODDS       float64 `json:"odds"`
	ODDS_NCL   float64 `json:"odds_ncl"`
	ODDS_NCU   float64 `json:"odds_ncu"`
	ODDS_BCL   float64 `json:"odds_bcl"`
	ODDS_BCU   float64 `json:"odds_bcu"`
	LODDS      float64 `json:"lodds"`
	LODDS_NCL  float64 `json:"lodds_ncl"`
	LODDS_NCU  float64 `json:"lodds_ncu"`
	LODDS_BCL  float64 `json:"lodds_bcl"`
	LODDS_BCU  float64 `json:"lodds_bcu"`
	ORSS       float64 `json:"orss"`
	ORSS_NCL   float64 `json:"orss_ncl"`
	ORSS_NCU   float64 `json:"orss_ncu"`
	ORSS_BCL   float64 `json:"orss_bcl"`
	ORSS_BCU   float64 `json:"orss_bcu"`
	EDS        float64 `json:"eds"`
	EDS_NCL    float64 `json:"eds_ncl"`
	EDS_NCU    float64 `json:"eds_ncu"`
	EDS_BCL    float64 `json:"eds_bcl"`
	EDS_BCU    float64 `json:"eds_bcu"`
	SEDS       float64 `json:"seds"`
	SEDS_NCL   float64 `json:"seds_ncl"`
	SEDS_NCU   float64 `json:"seds_ncu"`
	SEDS_BCL   float64 `json:"seds_bcl"`
	SEDS_BCU   float64 `json:"seds_bcu"`
	EDI        float64 `json:"edi"`
	EDI_NCL    float64 `json:"edi_ncl"`
	EDI_NCU    float64 `json:"edi_ncu"`
	EDI_BCL    float64 `json:"edi_bcl"`
	EDI_BCU    float64 `json:"edi_bcu"`
	SEDI       float64 `json:"sedi"`
	SEDI_NCL   float64 `json:"sedi_ncl"`
	SEDI_NCU   float64 `json:"sedi_ncu"`
	SEDI_BCL   float64 `json:"sedi_bcl"`
	SEDI_BCU   float64 `json:"sedi_bcu"`
	BAGSS      float64 `json:"bagss"`
	BAGSS_BCL  float64 `json:"bagss_bcl"`
	BAGSS_BCU  float64 `json:"bagss_bcu"`
	HSS_EC     float64 `json:"hss_ec"`
	HSS_EC_BCL float64 `json:"hss_ec_bcl"`
	HSS_EC_BCU float64 `json:"hss_ec_bcu"`
	EC_VALUE   float64 `json:"ec_value"`
}

type STAT_FHO struct {
	TOTAL  int     `json:"total"`
	F_RATE float64 `json:"f_rate"`
	H_RATE float64 `json:"h_rate"`
	O_RATE float64 `json:"o_rate"`
}

type STAT_NBRCNT struct {
	TOTAL      int     `json:"total"`
	FBS        float64 `json:"fbs"`
	FBS_BCL    float64 `json:"fbs_bcl"`
	FBS_BCU    float64 `json:"fbs_bcu"`
	FSS        float64 `json:"fss"`
	FSS_BCL    float64 `json:"fss_bcl"`
	FSS_BCU    float64 `json:"fss_bcu"`
	AFSS       float64 `json:"afss"`
	AFSS_BCL   float64 `json:"afss_bcl"`
	AFSS_BCU   float64 `json:"afss_bcu"`
	UFSS       float64 `json:"ufss"`
	UFSS_BCL   float64 `json:"ufss_bcl"`
	UFSS_BCU   float64 `json:"ufss_bcu"`
	F_RATE     float64 `json:"f_rate"`
	F_RATE_BCL float64 `json:"f_rate_bcl"`
	F_RATE_BCU float64 `json:"f_rate_bcu"`
	O_RATE     float64 `json:"o_rate"`
	O_RATE_BCL float64 `json:"o_rate_bcl"`
	O_RATE_BCU float64 `json:"o_rate_bcu"`
}

type STAT_ECNT struct {
	TOTAL            int     `json:"total"`
	N_ENS            int     `json:"n_ens"`
	CRPS             float64 `json:"crps"`
	CRPSS            float64 `json:"crpss"`
	IGN              float64 `json:"ign"`
	ME               float64 `json:"me"`
	RMSE             float64 `json:"rmse"`
	SPREAD           float64 `json:"spread"`
	ME_OERR          float64 `json:"me_oerr"`
	RMSE_OERR        float64 `json:"rmse_oerr"`
	SPREAD_OERR      float64 `json:"spread_oerr"`
	SPREAD_PLUS_OERR float64 `json:"spread_plus_oerr"`
	CRPSCL           float64 `json:"crpscl"`
	CRPS_EMP         float64 `json:"crps_emp"`
	CRPSCL_EMP       float64 `json:"crpscl_emp"`
	CRPSS_EMP        float64 `json:"crpss_emp"`
	CRPS_EMP_FAIR    float64 `json:"crps_emp_fair"`
	SPREAD_MD        float64 `json:"spread_md"`
	MAE              float64 `json:"mae"`
	MAE_OERR         float64 `json:"mae_oerr"`
	BIAS_RATIO       float64 `json:"bias_ratio"`
	N_GE_OBS         int     `json:"n_ge_obs"`
	ME_GE_OBS        float64 `json:"me_ge_obs"`
	N_LT_OBS         int     `json:"n_lt_obs"`
	ME_LT_OBS        float64 `json:"me_lt_obs"`
	IGN_CONV_OERR    float64 `json:"ign_conv_oerr"`
	IGN_CORR_OERR    float64 `json:"ign_corr_oerr"`
}

type STAT_SL1L2 struct {
	TOTAL int     `json:"total"`
	FBAR  float64 `json:"fbar"`
	OBAR  float64 `json:"obar"`
	FOBAR float64 `json:"fobar"`
	FFBAR float64 `json:"ffbar"`
	OOBAR float64 `json:"oobar"`
	MAE   float64 `json:"mae"`
}

// fillStructure functions
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
	i++
	if i <= dataLen {
		s.TOTAL_DIR, _ = strconv.ParseFloat(fields[10], 64)
	}
	i++
	if i <= dataLen {
		s.DIRA_ME, _ = strconv.ParseFloat(fields[11], 64)
	}
	i++
	if i <= dataLen {
		s.DIRA_MAE, _ = strconv.ParseFloat(fields[12], 64)
	}
	i++
	if i <= dataLen {
		s.DIRA_MSE, _ = strconv.ParseFloat(fields[13], 64)
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

func (s *STAT_SEEPS) fill_STAT_SEEPS(fields []string) {
	dataLen := len(fields) - 1
	i := -1
	i++
	if i <= dataLen {
		s.TOTAL, _ = strconv.Atoi(fields[0])
	}
	i++
	if i <= dataLen {
		s.ODFL, _ = strconv.ParseFloat(fields[1], 64)
	}
	i++
	if i <= dataLen {
		s.ODFH, _ = strconv.ParseFloat(fields[2], 64)
	}
	i++
	if i <= dataLen {
		s.OLFD, _ = strconv.ParseFloat(fields[3], 64)
	}
	i++
	if i <= dataLen {
		s.OLFH, _ = strconv.ParseFloat(fields[4], 64)
	}
	i++
	if i <= dataLen {
		s.OHFD, _ = strconv.ParseFloat(fields[5], 64)
	}
	i++
	if i <= dataLen {
		s.OHFL, _ = strconv.ParseFloat(fields[6], 64)
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

func (s *MTD_3DPAIR) fill_MTD_3DPAIR(fields []string) {
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
		s.SPACE_CENTROID_DIST, _ = strconv.ParseFloat(fields[2], 64)
	}
	i++
	if i <= dataLen {
		s.TIME_CENTROID_DELTA, _ = strconv.ParseFloat(fields[3], 64)
	}
	i++
	if i <= dataLen {
		s.AXIS_DIFF, _ = strconv.ParseFloat(fields[4], 64)
	}
	i++
	if i <= dataLen {
		s.SPEED_DELTA, _ = strconv.ParseFloat(fields[5], 64)
	}
	i++
	if i <= dataLen {
		s.DIRECTION_DIFF, _ = strconv.ParseFloat(fields[6], 64)
	}
	i++
	if i <= dataLen {
		s.VOLUME_RATIO, _ = strconv.ParseFloat(fields[7], 64)
	}
	i++
	if i <= dataLen {
		s.START_TIME_DELTA, _ = strconv.Atoi(fields[8])
	}
	i++
	if i <= dataLen {
		s.END_TIME_DELTA, _ = strconv.Atoi(fields[9])
	}
	i++
	if i <= dataLen {
		s.INTERSECTION_VOLUME, _ = strconv.ParseFloat(fields[10], 64)
	}
	i++
	if i <= dataLen {
		s.DURATION_DIFF, _ = strconv.ParseFloat(fields[11], 64)
	}
	i++
	if i <= dataLen {
		s.INTEREST, _ = strconv.ParseFloat(fields[12], 64)
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
	i++
	if i <= dataLen {
		s.LODDS, _ = strconv.ParseFloat(fields[19], 64)
	}
	i++
	if i <= dataLen {
		s.ORSS, _ = strconv.ParseFloat(fields[20], 64)
	}
	i++
	if i <= dataLen {
		s.EDS, _ = strconv.ParseFloat(fields[21], 64)
	}
	i++
	if i <= dataLen {
		s.SEDS, _ = strconv.ParseFloat(fields[22], 64)
	}
	i++
	if i <= dataLen {
		s.EDI, _ = strconv.ParseFloat(fields[23], 64)
	}
	i++
	if i <= dataLen {
		s.SEDI, _ = strconv.ParseFloat(fields[24], 64)
	}
	i++
	if i <= dataLen {
		s.BAGSS, _ = strconv.ParseFloat(fields[25], 64)
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
		s.OBS_CLIMO_MEAN, _ = strconv.ParseFloat(fields[10], 64)
	}
	i++
	if i <= dataLen {
		s.OBS_CLIMO_STDEV, _ = strconv.ParseFloat(fields[11], 64)
	}
	i++
	if i <= dataLen {
		s.OBS_CLIMO_CDF, _ = strconv.ParseFloat(fields[12], 64)
	}
	i++
	if i <= dataLen {
		s.FCST_CLIMO_MEAN, _ = strconv.ParseFloat(fields[13], 64)
	}
	i++
	if i <= dataLen {
		s.FCST_CLIMO_STDEV, _ = strconv.ParseFloat(fields[14], 64)
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
	i++
	if i <= dataLen {
		s.TOTAL_DIR, _ = strconv.ParseFloat(fields[10], 64)
	}
	i++
	if i <= dataLen {
		s.DIR_ME, _ = strconv.ParseFloat(fields[11], 64)
	}
	i++
	if i <= dataLen {
		s.DIR_MAE, _ = strconv.ParseFloat(fields[12], 64)
	}
	i++
	if i <= dataLen {
		s.DIR_MSE, _ = strconv.ParseFloat(fields[13], 64)
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
	i++
	if i <= dataLen {
		s.TOTAL_DIR, _ = strconv.ParseFloat(fields[63], 64)
	}
	i++
	if i <= dataLen {
		s.DIR_ME, _ = strconv.ParseFloat(fields[64], 64)
	}
	i++
	if i <= dataLen {
		s.DIR_ME_BCL, _ = strconv.ParseFloat(fields[65], 64)
	}
	i++
	if i <= dataLen {
		s.DIR_ME_BCU, _ = strconv.ParseFloat(fields[66], 64)
	}
	i++
	if i <= dataLen {
		s.DIR_MAE, _ = strconv.ParseFloat(fields[67], 64)
	}
	i++
	if i <= dataLen {
		s.DIR_MAE_BCL, _ = strconv.ParseFloat(fields[68], 64)
	}
	i++
	if i <= dataLen {
		s.DIR_MAE_BCU, _ = strconv.ParseFloat(fields[69], 64)
	}
	i++
	if i <= dataLen {
		s.DIR_MSE, _ = strconv.ParseFloat(fields[70], 64)
	}
	i++
	if i <= dataLen {
		s.DIR_MSE_BCL, _ = strconv.ParseFloat(fields[71], 64)
	}
	i++
	if i <= dataLen {
		s.DIR_MSE_BCU, _ = strconv.ParseFloat(fields[72], 64)
	}
	i++
	if i <= dataLen {
		s.DIR_RMSE, _ = strconv.ParseFloat(fields[73], 64)
	}
	i++
	if i <= dataLen {
		s.DIR_RMSE_BCL, _ = strconv.ParseFloat(fields[74], 64)
	}
	i++
	if i <= dataLen {
		s.DIR_RMSE_BCU, _ = strconv.ParseFloat(fields[75], 64)
	}
}

func (s *MTD_2DSINGLE) fill_MTD_2DSINGLE(fields []string) {
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
		s.TIME_INDEX, _ = strconv.Atoi(fields[2])
	}
	i++
	if i <= dataLen {
		s.AREA, _ = strconv.Atoi(fields[3])
	}
	i++
	if i <= dataLen {
		s.CENTROID_X, _ = strconv.ParseFloat(fields[4], 64)
	}
	i++
	if i <= dataLen {
		s.CENTROID_Y, _ = strconv.ParseFloat(fields[5], 64)
	}
	i++
	if i <= dataLen {
		s.CENTROID_LAT, _ = strconv.ParseFloat(fields[6], 64)
	}
	i++
	if i <= dataLen {
		s.CENTROID_LON, _ = strconv.ParseFloat(fields[7], 64)
	}
	i++
	if i <= dataLen {
		s.AXIS_ANG, _ = strconv.ParseFloat(fields[8], 64)
	}
	i++
	if i <= dataLen {
		s.INTENSITY_10, _ = strconv.ParseFloat(fields[9], 64)
	}
	i++
	if i <= dataLen {
		s.INTENSITY_25, _ = strconv.ParseFloat(fields[10], 64)
	}
	i++
	if i <= dataLen {
		s.INTENSITY_50, _ = strconv.ParseFloat(fields[11], 64)
	}
	i++
	if i <= dataLen {
		s.INTENSITY_75, _ = strconv.ParseFloat(fields[12], 64)
	}
	i++
	if i <= dataLen {
		s.INTENSITY_90, _ = strconv.ParseFloat(fields[13], 64)
	}
	i++
	if i <= dataLen {
		s.INTENSITY_USER, _ = strconv.ParseFloat(fields[14], 64)
	}
}

func (s *MTD_3DSINGLE) fill_MTD_3DSINGLE(fields []string) {
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
		s.CENTROID_T, _ = strconv.ParseFloat(fields[4], 64)
	}
	i++
	if i <= dataLen {
		s.CENTROID_LAT, _ = strconv.ParseFloat(fields[5], 64)
	}
	i++
	if i <= dataLen {
		s.CENTROID_LON, _ = strconv.ParseFloat(fields[6], 64)
	}
	i++
	if i <= dataLen {
		s.X_DOT, _ = strconv.ParseFloat(fields[7], 64)
	}
	i++
	if i <= dataLen {
		s.Y_DOT, _ = strconv.ParseFloat(fields[8], 64)
	}
	i++
	if i <= dataLen {
		s.AXIS_ANG, _ = strconv.ParseFloat(fields[9], 64)
	}
	i++
	if i <= dataLen {
		s.VOLUME, _ = strconv.Atoi(fields[10])
	}
	i++
	if i <= dataLen {
		s.START_TIME, _ = strconv.Atoi(fields[11])
	}
	i++
	if i <= dataLen {
		s.END_TIME, _ = strconv.Atoi(fields[12])
	}
	i++
	if i <= dataLen {
		s.CDIST_TRAVELLED, _ = strconv.ParseFloat(fields[13], 64)
	}
	i++
	if i <= dataLen {
		s.INTENSITY_10, _ = strconv.ParseFloat(fields[14], 64)
	}
	i++
	if i <= dataLen {
		s.INTENSITY_25, _ = strconv.ParseFloat(fields[15], 64)
	}
	i++
	if i <= dataLen {
		s.INTENSITY_50, _ = strconv.ParseFloat(fields[16], 64)
	}
	i++
	if i <= dataLen {
		s.INTENSITY_75, _ = strconv.ParseFloat(fields[17], 64)
	}
	i++
	if i <= dataLen {
		s.INTENSITY_90, _ = strconv.ParseFloat(fields[18], 64)
	}
	i++
	if i <= dataLen {
		s.INTENSITY_USER, _ = strconv.ParseFloat(fields[19], 64)
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
		s.OBS_CLIMO_MEAN, _ = strconv.ParseFloat(fields[15], 64)
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
		s.OBS_CLIMO_STDEV, _ = strconv.ParseFloat(fields[20], 64)
	}
	i++
	if i <= dataLen {
		s.FCST_CLIMO_MEAN, _ = strconv.ParseFloat(fields[21], 64)
	}
	i++
	if i <= dataLen {
		s.FCST_CLIMO_STDEV, _ = strconv.ParseFloat(fields[22], 64)
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
	i++
	if i <= dataLen {
		s.IGN_CONV_OERR, _ = strconv.ParseFloat(fields[25], 64)
	}
	i++
	if i <= dataLen {
		s.IGN_CORR_OERR, _ = strconv.ParseFloat(fields[26], 64)
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
	case "MTD_2DSINGLE":
		elem := MTD_2DSINGLE{}
		elem.fill_MTD_2DSINGLE_Header(headerData, &doc)
		elem.fill_MTD_2DSINGLE(dataData)
		if exists := (doc)["data"]; exists == nil {
			(doc)["data"] = make(map[string]MTD_2DSINGLE)
		}
		if val, ok := (doc)["data"].(map[string]MTD_2DSINGLE); ok {
			val[dataKey] = elem
			(doc)["data"] = val
		}
	case "MTD_3DSINGLE":
		elem := MTD_3DSINGLE{}
		elem.fill_MTD_3DSINGLE_Header(headerData, &doc)
		elem.fill_MTD_3DSINGLE(dataData)
		if exists := (doc)["data"]; exists == nil {
			(doc)["data"] = make(map[string]MTD_3DSINGLE)
		}
		if val, ok := (doc)["data"].(map[string]MTD_3DSINGLE); ok {
			val[dataKey] = elem
			(doc)["data"] = val
		}
	case "MTD_3DPAIR":
		elem := MTD_3DPAIR{}
		elem.fill_MTD_3DPAIR_Header(headerData, &doc)
		elem.fill_MTD_3DPAIR(dataData)
		if exists := (doc)["data"]; exists == nil {
			(doc)["data"] = make(map[string]MTD_3DPAIR)
		}
		if val, ok := (doc)["data"].(map[string]MTD_3DPAIR); ok {
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
	case "MTD_2DSINGLE":
		elem := MTD_2DSINGLE{}
		elem.fill_MTD_2DSINGLE(dataData)
		if val, ok := (*doc)["data"].(map[string]MTD_2DSINGLE); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
	case "MTD_3DSINGLE":
		elem := MTD_3DSINGLE{}
		elem.fill_MTD_3DSINGLE(dataData)
		if val, ok := (*doc)["data"].(map[string]MTD_3DSINGLE); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
	case "MTD_3DPAIR":
		elem := MTD_3DPAIR{}
		elem.fill_MTD_3DPAIR(dataData)
		if val, ok := (*doc)["data"].(map[string]MTD_3DPAIR); ok {
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

// DateFieldNames is a list of the date fields that need to be converted to epochs
var (
	DateFieldNames          = []string{"FCST_VALID_BEG", "FCST_VALID_END", "OBS_VALID_BEG", "OBS_VALID_END"}
	MetHeaderColumnsFileUrl = "https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/data/table_files/met_header_columns_V12.0.txt"
)
