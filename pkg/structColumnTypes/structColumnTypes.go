package structColumnTypes

import (
	"errors"
	"strconv"
)

// vxMetadata struct definition
type VxMetadata struct {
	ID      string `json:"id"`
	Subset  string `json:"subset"`
	Type    string `json:"type"`
	SubType string `json:"subtype"`
}

func (m *VxMetadata) fillMetadata(metaData VxMetadata, doc *map[string]interface{}) {
	// these must match the fields in VxMetdata!
	(*doc)["ID"] = metaData.ID
	(*doc)["Subset"] = metaData.Subset
	(*doc)["Type"] = metaData.Type
	(*doc)["SubType"] = metaData.SubType
}

var DateFieldNames = [...]string{"FCST_VALID_BEG", "FCST_VALID_END", "OBS_VALID_BEG", "OBS_VALID_END"}

// Header struct definitions
type STAT_PJC_header struct {
	Version        string `json:"version"`
	Model          string `json:"model"`
	Desc           string `json:"desc"`
	Fcst_lead      string `json:"fcst_lead"`
	Fcst_valid_beg string `json:"fcst_valid_beg"`
	Fcst_valid_end string `json:"fcst_valid_end"`
	Obs_lead       string `json:"obs_lead"`
	Obs_valid_beg  string `json:"obs_valid_beg"`
	Obs_valid_end  string `json:"obs_valid_end"`
	Fcst_var       string `json:"fcst_var"`
	Fcst_units     string `json:"fcst_units"`
	Fcst_lev       string `json:"fcst_lev"`
	Obs_var        string `json:"obs_var"`
	Obs_units      string `json:"obs_units"`
	Obs_lev        string `json:"obs_lev"`
	Obtype         string `json:"obtype"`
	Vx_mask        string `json:"vx_mask"`
	Interp_mthd    string `json:"interp_mthd"`
	Interp_pnts    string `json:"interp_pnts"`
	Fcst_thresh    string `json:"icst_thresh"`
	Obs_thresh     string `json:"obs_thresh"`
	Cov_thresh     string `json:"cov_thresh"`
	Alpha          string `json:"alpha"`
	Line_type      string `json:"line_type"`
}

type STAT_SSIDX_header struct {
	Version        string `json:"version"`
	Model          string `json:"model"`
	Desc           string `json:"desc"`
	Fcst_lead      string `json:"fcst_lead"`
	Fcst_valid_beg string `json:"fcst_valid_beg"`
	Fcst_valid_end string `json:"fcst_valid_end"`
	Obs_lead       string `json:"obs_lead"`
	Obs_valid_beg  string `json:"obs_valid_beg"`
	Obs_valid_end  string `json:"obs_valid_end"`
	Fcst_var       string `json:"fcst_var"`
	Fcst_units     string `json:"fcst_units"`
	Fcst_lev       string `json:"fcst_lev"`
	Obs_var        string `json:"obs_var"`
	Obs_units      string `json:"obs_units"`
	Obs_lev        string `json:"obs_lev"`
	Obtype         string `json:"obtype"`
	Vx_mask        string `json:"vx_mask"`
	Interp_mthd    string `json:"interp_mthd"`
	Interp_pnts    string `json:"interp_pnts"`
	Fcst_thresh    string `json:"fcst_thresh"`
	Obs_thresh     string `json:"obs_thresh"`
	Cov_thresh     string `json:"cov_thresh"`
	Alpha          string `json:"alpha"`
	Line_type      string `json:"line_type"`
}

type MODE_OBJ_header struct {
	Version    string `json:"version"`
	Model      string `json:"model"`
	N_valid    string `json:"n_valid"`
	Grid_res   string `json:"grid_res"`
	Desc       string `json:"desc"`
	Fcst_lead  string `json:"fcst_lead"`
	Fcst_valid string `json:"fcst_valid"`
	Fcst_accum string `json:"fcst_accum"`
	Obs_lead   string `json:"obs_lead"`
	Obs_valid  string `json:"obs_valid"`
	Obs_accum  string `json:"obs_accum"`
	Fcst_rad   string `json:"fcst_rad"`
	Fcst_thr   string `json:"fcst_thr"`
	Obs_rad    string `json:"obs_rad"`
	Obs_thr    string `json:"obs_thr"`
	Fcst_var   string `json:"fcst_var"`
	Fcst_units string `json:"fcst_units"`
	Fcst_lev   string `json:"fcst_lev"`
	Obs_var    string `json:"obs_var"`
	Obs_units  string `json:"obs_units"`
	Obs_lev    string `json:"obs_lev"`
	Obtype     string `json:"obtype"`
}

type TCST_TCDIAG_header struct {
	Version    string `json:"version"`
	Amodel     string `json:"amodel"`
	Bmodel     string `json:"bmodel"`
	Desc       string `json:"desc"`
	Storm_id   string `json:"storm_id"`
	Basin      string `json:"basin"`
	Cyclone    string `json:"cyclone"`
	Storm_name string `json:"storm_name"`
	Init       string `json:"init"`
	Lead       string `json:"lead"`
	Valid      string `json:"valid"`
	Init_mask  string `json:"init_mask"`
	Valid_mask string `json:"valid_mask"`
	Line_type  string `json:"line_type"`
}

type STAT_CTC_header struct {
	Version        string `json:"version"`
	Model          string `json:"model"`
	Desc           string `json:"desc"`
	Fcst_lead      string `json:"fcst_lead"`
	Fcst_valid_beg string `json:"fcst_valid_beg"`
	Fcst_valid_end string `json:"fcst_valid_end"`
	Obs_lead       string `json:"obs_lead"`
	Obs_valid_beg  string `json:"obs_valid_beg"`
	Obs_valid_end  string `json:"obs_valid_end"`
	Fcst_var       string `json:"fcst_var"`
	Fcst_units     string `json:"fcst_units"`
	Fcst_lev       string `json:"fcst_lev"`
	Obs_var        string `json:"obs_var"`
	Obs_units      string `json:"obs_units"`
	Obs_lev        string `json:"obs_lev"`
	Obtype         string `json:"obtype"`
	Vx_mask        string `json:"vx_mask"`
	Interp_mthd    string `json:"interp_mthd"`
	Interp_pnts    string `json:"interp_pnts"`
	Fcst_thresh    string `json:"fcst_thresh"`
	Obs_thresh     string `json:"obs_thresh"`
	Cov_thresh     string `json:"cov_thresh"`
	Alpha          string `json:"alpha"`
	Line_type      string `json:"line_type"`
}

type STAT_MCTC_header struct {
	Version        string `json:"version"`
	Model          string `json:"model"`
	Desc           string `json:"desc"`
	Fcst_lead      string `json:"fcst_lead"`
	Fcst_valid_beg string `json:"fcst_valid_beg"`
	Fcst_valid_end string `json:"fcst_valid_end"`
	Obs_lead       string `json:"obs_lead"`
	Obs_valid_beg  string `json:"obs_valid_beg"`
	Obs_valid_end  string `json:"obs_valid_end"`
	Fcst_var       string `json:"fcst_var"`
	Fcst_units     string `json:"fcst_units"`
	Fcst_lev       string `json:"fcst_lev"`
	Obs_var        string `json:"obs_var"`
	Obs_units      string `json:"obs_units"`
	Obs_lev        string `json:"obs_lev"`
	Obtype         string `json:"obtype"`
	Vx_mask        string `json:"vx_mask"`
	Interp_mthd    string `json:"interp_mthd"`
	Interp_pnts    string `json:"interp_pnts"`
	Fcst_thresh    string `json:"fcst_thresh"`
	Obs_thresh     string `json:"obs_thresh"`
	Cov_thresh     string `json:"cov_thresh"`
	Alpha          string `json:"alpha"`
	Line_type      string `json:"line_type"`
}

type STAT_MPR_header struct {
	Version        string `json:"version"`
	Model          string `json:"model"`
	Desc           string `json:"desc"`
	Fcst_lead      string `json:"fcst_lead"`
	Fcst_valid_beg string `json:"fcst_valid_beg"`
	Fcst_valid_end string `json:"fcst_valid_end"`
	Obs_lead       string `json:"obs_lead"`
	Obs_valid_beg  string `json:"obs_valid_beg"`
	Obs_valid_end  string `json:"obs_valid_end"`
	Fcst_var       string `json:"fcst_var"`
	Fcst_units     string `json:"fcst_units"`
	Fcst_lev       string `json:"fcst_lev"`
	Obs_var        string `json:"obs_var"`
	Obs_units      string `json:"obs_units"`
	Obs_lev        string `json:"obs_lev"`
	Obtype         string `json:"obtype"`
	Vx_mask        string `json:"vx_mask"`
	Interp_mthd    string `json:"interp_mthd"`
	Interp_pnts    string `json:"interp_pnts"`
	Fcst_thresh    string `json:"fcst_thresh"`
	Obs_thresh     string `json:"obs_thresh"`
	Cov_thresh     string `json:"cov_thresh"`
	Alpha          string `json:"alpha"`
	Line_type      string `json:"line_type"`
}

type STAT_DMAP_header struct {
	Version        string `json:"version"`
	Model          string `json:"model"`
	Desc           string `json:"desc"`
	Fcst_lead      string `json:"fcst_lead"`
	Fcst_valid_beg string `json:"fcst_valid_beg"`
	Fcst_valid_end string `json:"fcst_valid_end"`
	Obs_lead       string `json:"obs_lead"`
	Obs_valid_beg  string `json:"obs_valid_beg"`
	Obs_valid_end  string `json:"obs_valid_end"`
	Fcst_var       string `json:"fcst_var"`
	Fcst_units     string `json:"fcst_units"`
	Fcst_lev       string `json:"fcst_lev"`
	Obs_var        string `json:"obs_var"`
	Obs_units      string `json:"obs_units"`
	Obs_lev        string `json:"obs_lev"`
	Obtype         string `json:"obtype"`
	Vx_mask        string `json:"vx_mask"`
	Interp_mthd    string `json:"interp_mthd"`
	Interp_pnts    string `json:"interp_pnts"`
	Fcst_thresh    string `json:"fcst_thresh"`
	Obs_thresh     string `json:"obs_thresh"`
	Cov_thresh     string `json:"cov_thresh"`
	Alpha          string `json:"alpha"`
	Line_type      string `json:"line_type"`
}

type STAT_VL1L2_header struct {
	Version        string `json:"version"`
	Model          string `json:"model"`
	Desc           string `json:"desc"`
	Fcst_lead      string `json:"fcst_lead"`
	Fcst_valid_beg string `json:"fcst_valid_beg"`
	Fcst_valid_end string `json:"fcst_valid_end"`
	Obs_lead       string `json:"obs_lead"`
	Obs_valid_beg  string `json:"obs_valid_beg"`
	Obs_valid_end  string `json:"obs_valid_end"`
	Fcst_var       string `json:"fcst_var"`
	Fcst_units     string `json:"fcst_units"`
	Fcst_lev       string `json:"fcst_lev"`
	Obs_var        string `json:"obs_var"`
	Obs_units      string `json:"obs_units"`
	Obs_lev        string `json:"obs_lev"`
	Obtype         string `json:"obtype"`
	Vx_mask        string `json:"vx_mask"`
	Interp_mthd    string `json:"interp_mthd"`
	Interp_pnts    string `json:"interp_pnts"`
	Fcst_thresh    string `json:"fcst_thresh"`
	Obs_thresh     string `json:"obs_thresh"`
	Cov_thresh     string `json:"cov_thresh"`
	Alpha          string `json:"alpha"`
	Line_type      string `json:"line_type"`
}

type MODE_CTS_header struct {
	Version    string `json:"version"`
	Model      string `json:"model"`
	N_valid    string `json:"n_valid"`
	Grid_res   string `json:"grid_res"`
	Desc       string `json:"desc"`
	Fcst_lead  string `json:"fcst_lead"`
	Fcst_valid string `json:"fcst_valid"`
	Fcst_accum string `json:"fcst_accum"`
	Obs_lead   string `json:"obs_lead"`
	Obs_valid  string `json:"obs_valid"`
	Obs_accum  string `json:"obs_accum"`
	Fcst_rad   string `json:"fcst_rad"`
	Fcst_thr   string `json:"fcst_thr"`
	Obs_rad    string `json:"obs_rad"`
	Obs_thr    string `json:"obs_thr"`
	Fcst_var   string `json:"fcst_var"`
	Fcst_units string `json:"fcst_units"`
	Fcst_lev   string `json:"fcst_lev"`
	Obs_var    string `json:"obs_var"`
	Obs_units  string `json:"obs_units"`
	Obs_lev    string `json:"obs_lev"`
	Obtype     string `json:"obtype"`
}

type STAT_CTS_header struct {
	Version        string `json:"version"`
	Model          string `json:"model"`
	Desc           string `json:"desc"`
	Fcst_lead      string `json:"fcst_lead"`
	Fcst_valid_beg string `json:"fcst_valid_beg"`
	Fcst_valid_end string `json:"fcst_valid_end"`
	Obs_lead       string `json:"obs_lead"`
	Obs_valid_beg  string `json:"obs_valid_beg"`
	Obs_valid_end  string `json:"obs_valid_end"`
	Fcst_var       string `json:"fcst_var"`
	Fcst_units     string `json:"fcst_units"`
	Fcst_lev       string `json:"fcst_lev"`
	Obs_var        string `json:"obs_var"`
	Obs_units      string `json:"obs_units"`
	Obs_lev        string `json:"obs_lev"`
	Obtype         string `json:"obtype"`
	Vx_mask        string `json:"vx_mask"`
	Interp_mthd    string `json:"interp_mthd"`
	Interp_pnts    string `json:"interp_pnts"`
	Fcst_thresh    string `json:"fcst_thresh"`
	Obs_thresh     string `json:"obs_thresh"`
	Cov_thresh     string `json:"cov_thresh"`
	Alpha          string `json:"alpha"`
	Line_type      string `json:"line_type"`
}

type STAT_NBRCNT_header struct {
	Version        string `json:"version"`
	Model          string `json:"model"`
	Desc           string `json:"desc"`
	Fcst_lead      string `json:"fcst_lead"`
	Fcst_valid_beg string `json:"fcst_valid_beg"`
	Fcst_valid_end string `json:"fcst_valid_end"`
	Obs_lead       string `json:"obs_lead"`
	Obs_valid_beg  string `json:"obs_valid_beg"`
	Obs_valid_end  string `json:"obs_valid_end"`
	Fcst_var       string `json:"fcst_var"`
	Fcst_units     string `json:"fcst_units"`
	Fcst_lev       string `json:"fcst_lev"`
	Obs_var        string `json:"obs_var"`
	Obs_units      string `json:"obs_units"`
	Obs_lev        string `json:"obs_lev"`
	Obtype         string `json:"obtype"`
	Vx_mask        string `json:"vx_mask"`
	Interp_mthd    string `json:"interp_mthd"`
	Interp_pnts    string `json:"interp_pnts"`
	Fcst_thresh    string `json:"fcst_thresh"`
	Obs_thresh     string `json:"obs_thresh"`
	Cov_thresh     string `json:"cov_thresh"`
	Alpha          string `json:"alpha"`
	Line_type      string `json:"line_type"`
}

type STAT_PHIST_header struct {
	Version        string `json:"version"`
	Model          string `json:"model"`
	Desc           string `json:"desc"`
	Fcst_lead      string `json:"fcst_lead"`
	Fcst_valid_beg string `json:"fcst_valid_beg"`
	Fcst_valid_end string `json:"fcst_valid_end"`
	Obs_lead       string `json:"obs_lead"`
	Obs_valid_beg  string `json:"obs_valid_beg"`
	Obs_valid_end  string `json:"obs_valid_end"`
	Fcst_var       string `json:"fcst_var"`
	Fcst_units     string `json:"fcst_units"`
	Fcst_lev       string `json:"fcst_lev"`
	Obs_var        string `json:"obs_var"`
	Obs_units      string `json:"obs_units"`
	Obs_lev        string `json:"obs_lev"`
	Obtype         string `json:"obtype"`
	Vx_mask        string `json:"vx_mask"`
	Interp_mthd    string `json:"interp_mthd"`
	Interp_pnts    string `json:"interp_pnts"`
	Fcst_thresh    string `json:"fcst_thresh"`
	Obs_thresh     string `json:"obs_thresh"`
	Cov_thresh     string `json:"cov_thresh"`
	Alpha          string `json:"alpha"`
	Line_type      string `json:"line_type"`
}

type STAT_VAL1L2_header struct {
	Version        string `json:"version"`
	Model          string `json:"model"`
	Desc           string `json:"desc"`
	Fcst_valid_beg string `json:"fcst_valid_beg"`
	Fcst_valid_end string `json:"fcst_valid_end"`
	Obs_lead       string `json:"obs_lead"`
	Obs_valid_beg  string `json:"obs_valid_beg"`
	Obs_valid_end  string `json:"obs_valid_end"`
	Fcst_var       string `json:"fcst_var"`
	Fcst_units     string `json:"fcst_units"`
	Fcst_lev       string `json:"fcst_lev"`
	Obs_var        string `json:"obs_var"`
	Obs_units      string `json:"obs_units"`
	Obs_lev        string `json:"obs_lev"`
	Obtype         string `json:"obtype"`
	Vx_mask        string `json:"vx_mask"`
	Interp_mthd    string `json:"interp_mthd"`
	Interp_pnts    string `json:"interp_pnts"`
	Fcst_thresh    string `json:"fcst_thresh"`
	Obs_thresh     string `json:"obs_thresh"`
	Cov_thresh     string `json:"cov_thresh"`
	Alpha          string `json:"alpha"`
	Line_type      string `json:"line_type"`
}

type STAT_CNT_header struct {
	Version        string `json:"version"`
	Model          string `json:"model"`
	Desc           string `json:"desc"`
	Fcst_lead      string `json:"fcst_lead"`
	Fcst_valid_beg string `json:"fcst_valid_beg"`
	Fcst_valid_end string `json:"fcst_valid_end"`
	Obs_lead       string `json:"obs_lead"`
	Obs_valid_beg  string `json:"obs_valid_beg"`
	Obs_valid_end  string `json:"obs_valid_end"`
	Fcst_var       string `json:"fcst_var"`
	Fcst_units     string `json:"fcst_units"`
	Fcst_lev       string `json:"fcst_lev"`
	Obs_var        string `json:"obs_var"`
	Obs_units      string `json:"obs_units"`
	Obs_lev        string `json:"obs_lev"`
	Obtype         string `json:"obtype"`
	Vx_mask        string `json:"vx_mask"`
	Interp_mthd    string `json:"interp_mthd"`
	Interp_pnts    string `json:"interp_pnts"`
	Fcst_thresh    string `json:"fcst_thresh"`
	Obs_thresh     string `json:"obs_thresh"`
	Cov_thresh     string `json:"cov_thresh"`
	Alpha          string `json:"alpha"`
	Line_type      string `json:"line_type"`
}

type STAT_PRC_header struct {
	Version        string `json:"version"`
	Model          string `json:"model"`
	Desc           string `json:"desc"`
	Fcst_lead      string `json:"fcst_lead"`
	Fcst_valid_beg string `json:"fcst_valid_beg"`
	Fcst_valid_end string `json:"fcst_valid_end"`
	Obs_lead       string `json:"obs_lead"`
	Obs_valid_beg  string `json:"obs_valid_beg"`
	Obs_valid_end  string `json:"obs_valid_end"`
	Fcst_var       string `json:"fcst_var"`
	Fcst_units     string `json:"fcst_units"`
	Fcst_lev       string `json:"fcst_lev"`
	Obs_var        string `json:"obs_var"`
	Obs_units      string `json:"obs_units"`
	Obs_lev        string `json:"obs_lev"`
	Obtype         string `json:"obtype"`
	Vx_mask        string `json:"vx_mask"`
	Interp_mthd    string `json:"interp_mthd"`
	Interp_pnts    string `json:"interp_pnts"`
	Fcst_thresh    string `json:"fcst_thresh"`
	Obs_thresh     string `json:"obs_thresh"`
	Cov_thresh     string `json:"cov_thresh"`
	Alpha          string `json:"alpha"`
	Line_type      string `json:"line_type"`
}

type STAT_SL1L2_header struct {
	Version        string `json:"version"`
	Model          string `json:"model"`
	Desc           string `json:"desc"`
	Fcst_lead      string `json:"fcst_lead"`
	Fcst_valid_beg string `json:"fcst_valid_beg"`
	Fcst_valid_end string `json:"fcst_valid_end"`
	Obs_lead       string `json:"obs_lead"`
	Obs_valid_beg  string `json:"obs_valid_beg"`
	Obs_valid_end  string `json:"obs_valid_end"`
	Fcst_var       string `json:"fcst_var"`
	Fcst_units     string `json:"fcst_units"`
	Fcst_lev       string `json:"fcst_lev"`
	Obs_var        string `json:"obs_var"`
	Obs_units      string `json:"obs_units"`
	Obs_lev        string `json:"obs_lev"`
	Obtype         string `json:"obtype"`
	Vx_mask        string `json:"vx_mask"`
	Interp_mthd    string `json:"interp_mthd"`
	Interp_pnts    string `json:"interp_pnts"`
	Fcst_thresh    string `json:"fcst_thresh"`
	Obs_thresh     string `json:"obs_thresh"`
	Cov_thresh     string `json:"cov_thresh"`
	Alpha          string `json:"alpha"`
	Line_type      string `json:"line_type"`
}

type TCST_TCMPR_header struct {
	Version    string `json:"version"`
	Amodel     string `json:"amodel"`
	Bmodel     string `json:"bmodel"`
	Desc       string `json:"desc"`
	Storm_id   string `json:"storm_id"`
	Basin      string `json:"basin"`
	Cyclone    string `json:"cyclone"`
	Storm_name string `json:"storm_name"`
	Init       string `json:"init"`
	Lead       string `json:"lead"`
	Valid      string `json:"valid"`
	Init_mask  string `json:"init_mask"`
	Valid_mask string `json:"valid_mask"`
	Line_type  string `json:"line_type"`
}

type STAT_FHO_header struct {
	Version        string `json:"version"`
	Model          string `json:"model"`
	Desc           string `json:"desc"`
	Fcst_lead      string `json:"fcst_lead"`
	Fcst_valid_beg string `json:"fcst_valid_beg"`
	Fcst_valid_end string `json:"fcst_valid_end"`
	Obs_lead       string `json:"obs_lead"`
	Obs_valid_beg  string `json:"obs_valid_beg"`
	Obs_valid_end  string `json:"obs_valid_end"`
	Fcst_var       string `json:"fcst_var"`
	Fcst_units     string `json:"fcst_units"`
	Fcst_lev       string `json:"fcst_lev"`
	Obs_var        string `json:"obs_var"`
	Obs_units      string `json:"obs_units"`
	Obs_lev        string `json:"obs_lev"`
	Obtype         string `json:"obtype"`
	Vx_mask        string `json:"vx_mask"`
	Interp_mthd    string `json:"interp_mthd"`
	Interp_pnts    string `json:"interp_pnts"`
	Fcst_thresh    string `json:"fcst_thresh"`
	Obs_thresh     string `json:"obs_thresh"`
	Cov_thresh     string `json:"cov_thresh"`
	Alpha          string `json:"alpha"`
	Line_type      string `json:"line_type"`
}

type STAT_ISC_header struct {
	Version        string `json:"version"`
	Model          string `json:"model"`
	Desc           string `json:"desc"`
	Fcst_lead      string `json:"fcst_lead"`
	Fcst_valid_beg string `json:"fcst_valid_beg"`
	Fcst_valid_end string `json:"fcst_valid_end"`
	Obs_lead       string `json:"obs_lead"`
	Obs_valid_beg  string `json:"obs_valid_beg"`
	Obs_valid_end  string `json:"obs_valid_end"`
	Fcst_var       string `json:"fcst_var"`
	Fcst_units     string `json:"fcst_units"`
	Fcst_lev       string `json:"fcst_lev"`
	Obs_var        string `json:"obs_var"`
	Obs_units      string `json:"obs_units"`
	Obs_lev        string `json:"obs_lev"`
	Obtype         string `json:"obtype"`
	Vx_mask        string `json:"vx_mask"`
	Interp_mthd    string `json:"interp_mthd"`
	Interp_pnts    string `json:"interp_pnts"`
	Fcst_thresh    string `json:"fcst_thresh"`
	Obs_thresh     string `json:"obs_thresh"`
	Cov_thresh     string `json:"cov_thresh"`
	Alpha          string `json:"alpha"`
	Line_type      string `json:"line_type"`
}

type STAT_SEEPS_MPR_header struct {
	Version        string `json:"version"`
	Model          string `json:"model"`
	Desc           string `json:"desc"`
	Fcst_lead      string `json:"fcst_lead"`
	Fcst_valid_beg string `json:"fcst_valid_beg"`
	Fcst_valid_end string `json:"fcst_valid_end"`
	Obs_lead       string `json:"obs_lead"`
	Obs_valid_beg  string `json:"obs_valid_beg"`
	Obs_valid_end  string `json:"obs_valid_end"`
	Fcst_var       string `json:"fcst_var"`
	Fcst_units     string `json:"fcst_units"`
	Fcst_lev       string `json:"fcst_lev"`
	Obs_var        string `json:"obs_var"`
	Obs_units      string `json:"obs_units"`
	Obs_lev        string `json:"obs_lev"`
	Obtype         string `json:"obtype"`
	Vx_mask        string `json:"vx_mask"`
	Interp_mthd    string `json:"interp_mthd"`
	Interp_pnts    string `json:"interp_pnts"`
	Fcst_thresh    string `json:"fcst_thresh"`
	Obs_thresh     string `json:"obs_thresh"`
	Cov_thresh     string `json:"cov_thresh"`
	Alpha          string `json:"alpha"`
	Line_type      string `json:"line_type"`
}

type STAT_SSVAR_header struct {
	Version        string `json:"version"`
	Model          string `json:"model"`
	Desc           string `json:"desc"`
	Fcst_lead      string `json:"fcst_lead"`
	Fcst_valid_beg string `json:"fcst_valid_beg"`
	Fcst_valid_end string `json:"fcst_valid_end"`
	Obs_lead       string `json:"obs_lead"`
	Obs_valid_beg  string `json:"obs_valid_beg"`
	Obs_valid_end  string `json:"obs_valid_end"`
	Fcst_var       string `json:"fcst_var"`
	Fcst_units     string `json:"fcst_units"`
	Fcst_lev       string `json:"fcst_lev"`
	Obs_var        string `json:"obs_var"`
	Obs_units      string `json:"obs_units"`
	Obs_lev        string `json:"obs_lev"`
	Obtype         string `json:"obtype"`
	Vx_mask        string `json:"vx_mask"`
	Interp_mthd    string `json:"interp_mthd"`
	Interp_pnts    string `json:"interp_pnts"`
	Fcst_thresh    string `json:"fcst_thresh"`
	Obs_thresh     string `json:"obs_thresh"`
	Cov_thresh     string `json:"cov_thresh"`
	Alpha          string `json:"alpha"`
	Line_type      string `json:"line_type"`
}

type STAT_GENMPR_header struct {
	Version        string `json:"version"`
	Model          string `json:"model"`
	Desc           string `json:"desc"`
	Fcst_lead      string `json:"fcst_lead"`
	Fcst_valid_beg string `json:"fcst_valid_beg"`
	Fcst_valid_end string `json:"fcst_valid_end"`
	Obs_lead       string `json:"obs_lead"`
	Obs_valid_beg  string `json:"obs_valid_beg"`
	Obs_valid_end  string `json:"obs_valid_end"`
	Fcst_var       string `json:"fcst_var"`
	Fcst_units     string `json:"fcst_units"`
	Fcst_lev       string `json:"fcst_lev"`
	Obs_var        string `json:"obs_var"`
	Obs_units      string `json:"obs_units"`
	Obs_lev        string `json:"obs_lev"`
	Obtype         string `json:"obtype"`
	Vx_mask        string `json:"vx_mask"`
	Interp_mthd    string `json:"interp_mthd"`
	Interp_pnts    string `json:"interp_pnts"`
	Fcst_thresh    string `json:"fcst_thresh"`
	Obs_thresh     string `json:"obs_thresh"`
	Cov_thresh     string `json:"cov_thresh"`
	Alpha          string `json:"alpha"`
	Line_type      string `json:"line_type"`
}

type TCST_PROBRIRW_header struct {
	Version    string `json:"version"`
	Amodel     string `json:"amodel"`
	Bmodel     string `json:"bmodel"`
	Desc       string `json:"desc"`
	Storm_id   string `json:"storm_id"`
	Basin      string `json:"basin"`
	Cyclone    string `json:"cyclone"`
	Storm_name string `json:"storm_name"`
	Init       string `json:"init"`
	Lead       string `json:"lead"`
	Valid      string `json:"valid"`
	Init_mask  string `json:"init_mask"`
	Valid_mask string `json:"valid_mask"`
	Line_type  string `json:"line_type"`
}

type STAT_NBRCTS_header struct {
	Version        string `json:"version"`
	Model          string `json:"model"`
	Desc           string `json:"desc"`
	Fcst_lead      string `json:"fcst_lead"`
	Fcst_valid_beg string `json:"fcst_valid_beg"`
	Fcst_valid_end string `json:"fcst_valid_end"`
	Obs_lead       string `json:"obs_lead"`
	Obs_valid_beg  string `json:"obs_valid_beg"`
	Obs_valid_end  string `json:"obs_valid_end"`
	Fcst_var       string `json:"fcst_var"`
	Fcst_units     string `json:"fcst_units"`
	Fcst_lev       string `json:"fcst_lev"`
	Obs_var        string `json:"obs_var"`
	Obs_units      string `json:"obs_units"`
	Obs_lev        string `json:"obs_lev"`
	Obtype         string `json:"obtype"`
	Vx_mask        string `json:"vx_mask"`
	Interp_mthd    string `json:"interp_mthd"`
	Interp_pnts    string `json:"interp_pnts"`
	Fcst_thresh    string `json:"fcst_thresh"`
	Obs_thresh     string `json:"obs_thresh"`
	Cov_thresh     string `json:"cov_thresh"`
	Alpha          string `json:"alpha"`
	Line_type      string `json:"line_type"`
}

type STAT_ORANK_header struct {
	Version        string `json:"version"`
	Model          string `json:"model"`
	Desc           string `json:"desc"`
	Fcst_lead      string `json:"fcst_lead"`
	Fcst_valid_beg string `json:"fcst_valid_beg"`
	Fcst_valid_end string `json:"fcst_valid_end"`
	Obs_lead       string `json:"obs_lead"`
	Obs_valid_beg  string `json:"obs_valid_beg"`
	Obs_valid_end  string `json:"obs_valid_end"`
	Fcst_var       string `json:"fcst_var"`
	Fcst_units     string `json:"fcst_units"`
	Fcst_lev       string `json:"fcst_lev"`
	Obs_var        string `json:"obs_var"`
	Obs_units      string `json:"obs_units"`
	Obs_lev        string `json:"obs_lev"`
	Obtype         string `json:"obtype"`
	Vx_mask        string `json:"vx_mask"`
	Interp_mthd    string `json:"interp_mthd"`
	Interp_pnts    string `json:"interp_pnts"`
	Fcst_thresh    string `json:"fcst_thresh"`
	Obs_thresh     string `json:"obs_thresh"`
	Cov_thresh     string `json:"cov_thresh"`
	Alpha          string `json:"alpha"`
	Line_type      string `json:"line_type"`
}

type STAT_ECNT_header struct {
	Version        string `json:"version"`
	Model          string `json:"model"`
	Desc           string `json:"desc"`
	Fcst_lead      string `json:"fcst_lead"`
	Fcst_valid_beg string `json:"fcst_valid_beg"`
	Fcst_valid_end string `json:"fcst_valid_end"`
	Obs_lead       string `json:"obs_lead"`
	Obs_valid_beg  string `json:"obs_valid_beg"`
	Obs_valid_end  string `json:"obs_valid_end"`
	Fcst_var       string `json:"fcst_var"`
	Fcst_units     string `json:"fcst_units"`
	Fcst_lev       string `json:"fcst_lev"`
	Obs_var        string `json:"obs_var"`
	Obs_units      string `json:"obs_units"`
	Obs_lev        string `json:"obs_lev"`
	Obtype         string `json:"obtype"`
	Vx_mask        string `json:"vx_mask"`
	Interp_mthd    string `json:"interp_mthd"`
	Interp_pnts    string `json:"interp_pnts"`
	Fcst_thresh    string `json:"fcst_thresh"`
	Obs_thresh     string `json:"obs_thresh"`
	Cov_thresh     string `json:"cov_thresh"`
	Alpha          string `json:"alpha"`
	Line_type      string `json:"line_type"`
}

type STAT_RPS_header struct {
	Version        string `json:"version"`
	Model          string `json:"model"`
	Desc           string `json:"desc"`
	Fcst_lead      string `json:"fcst_lead"`
	Fcst_valid_beg string `json:"fcst_valid_beg"`
	Fcst_valid_end string `json:"fcst_valid_end"`
	Obs_lead       string `json:"obs_lead"`
	Obs_valid_beg  string `json:"obs_valid_beg"`
	Obs_valid_end  string `json:"obs_valid_end"`
	Fcst_var       string `json:"fcst_var"`
	Fcst_units     string `json:"fcst_units"`
	Fcst_lev       string `json:"fcst_lev"`
	Obs_var        string `json:"obs_var"`
	Obs_units      string `json:"obs_units"`
	Obs_lev        string `json:"obs_lev"`
	Obtype         string `json:"obtype"`
	Vx_mask        string `json:"vx_mask"`
	Interp_mthd    string `json:"interp_mthd"`
	Interp_pnts    string `json:"interp_pnts"`
	Fcst_thresh    string `json:"fcst_thresh"`
	Obs_thresh     string `json:"obs_thresh"`
	Cov_thresh     string `json:"cov_thresh"`
	Alpha          string `json:"alpha"`
	Line_type      string `json:"line_type"`
}

type STAT_RHIST_header struct {
	Version        string `json:"version"`
	Model          string `json:"model"`
	Desc           string `json:"desc"`
	Fcst_lead      string `json:"fcst_lead"`
	Fcst_valid_beg string `json:"fcst_valid_beg"`
	Fcst_valid_end string `json:"fcst_valid_end"`
	Obs_lead       string `json:"obs_lead"`
	Obs_valid_beg  string `json:"obs_valid_beg"`
	Obs_valid_end  string `json:"obs_valid_end"`
	Fcst_var       string `json:"fcst_var"`
	Fcst_units     string `json:"fcst_units"`
	Fcst_lev       string `json:"fcst_lev"`
	Obs_var        string `json:"obs_var"`
	Obs_units      string `json:"obs_units"`
	Obs_lev        string `json:"obs_lev"`
	Obtype         string `json:"obtype"`
	Vx_mask        string `json:"vx_mask"`
	Interp_mthd    string `json:"interp_mthd"`
	Interp_pnts    string `json:"interp_pnts"`
	Fcst_thresh    string `json:"fcst_thresh"`
	Obs_thresh     string `json:"obs_thresh"`
	Cov_thresh     string `json:"cov_thresh"`
	Alpha          string `json:"alpha"`
	Line_type      string `json:"line_type"`
}

type STAT_RELP_header struct {
	Version        string `json:"version"`
	Model          string `json:"model"`
	Desc           string `json:"desc"`
	Fcst_lead      string `json:"fcst_lead"`
	Fcst_valid_beg string `json:"fcst_valid_beg"`
	Fcst_valid_end string `json:"fcst_valid_end"`
	Obs_lead       string `json:"obs_lead"`
	Obs_valid_beg  string `json:"obs_valid_beg"`
	Obs_valid_end  string `json:"obs_valid_end"`
	Fcst_var       string `json:"fcst_var"`
	Fcst_units     string `json:"fcst_units"`
	Fcst_lev       string `json:"fcst_lev"`
	Obs_var        string `json:"obs_var"`
	Obs_units      string `json:"obs_units"`
	Obs_lev        string `json:"obs_lev"`
	Obtype         string `json:"obtype"`
	Vx_mask        string `json:"vx_mask"`
	Interp_mthd    string `json:"interp_mthd"`
	Interp_pnts    string `json:"interp_pnts"`
	Fcst_thresh    string `json:"fcst_thresh"`
	Obs_thresh     string `json:"obs_thresh"`
	Cov_thresh     string `json:"cov_thresh"`
	Alpha          string `json:"alpha"`
	Line_type      string `json:"line_type"`
}

type STAT_PSTD_header struct {
	Version        string `json:"version"`
	Model          string `json:"model"`
	Desc           string `json:"desc"`
	Fcst_lead      string `json:"fcst_lead"`
	Fcst_valid_beg string `json:"fcst_valid_beg"`
	Fcst_valid_end string `json:"fcst_valid_end"`
	Obs_lead       string `json:"obs_lead"`
	Obs_valid_beg  string `json:"obs_valid_beg"`
	Obs_valid_end  string `json:"obs_valid_end"`
	Fcst_var       string `json:"fcst_var"`
	Fcst_units     string `json:"fcst_units"`
	Fcst_lev       string `json:"fcst_lev"`
	Obs_var        string `json:"obs_var"`
	Obs_units      string `json:"obs_units"`
	Obs_lev        string `json:"obs_lev"`
	Obtype         string `json:"obtype"`
	Vx_mask        string `json:"vx_mask"`
	Interp_mthd    string `json:"interp_mthd"`
	Interp_pnts    string `json:"interp_pnts"`
	Fcst_thresh    string `json:"fcst_thresh"`
	Obs_thresh     string `json:"obs_thresh"`
	Cov_thresh     string `json:"cov_thresh"`
	Alpha          string `json:"alpha"`
	Line_type      string `json:"line_type"`
}

type STAT_SEEPS_header struct {
	Version        string `json:"version"`
	Model          string `json:"model"`
	Desc           string `json:"desc"`
	Fcst_lead      string `json:"fcst_lead"`
	Fcst_valid_beg string `json:"fcst_valid_beg"`
	Fcst_valid_end string `json:"fcst_valid_end"`
	Obs_lead       string `json:"obs_lead"`
	Obs_valid_beg  string `json:"obs_valid_beg"`
	Obs_valid_end  string `json:"obs_valid_end"`
	Fcst_var       string `json:"fcst_var"`
	Fcst_units     string `json:"fcst_units"`
	Fcst_lev       string `json:"fcst_lev"`
	Obs_var        string `json:"obs_var"`
	Obs_units      string `json:"obs_units"`
	Obs_lev        string `json:"obs_lev"`
	Obtype         string `json:"obtype"`
	Vx_mask        string `json:"vx_mask"`
	Interp_mthd    string `json:"interp_mthd"`
	Interp_pnts    string `json:"interp_pnts"`
	Fcst_thresh    string `json:"fcst_thresh"`
	Obs_thresh     string `json:"obs_thresh"`
	Cov_thresh     string `json:"cov_thresh"`
	Alpha          string `json:"alpha"`
	Line_type      string `json:"line_type"`
}

type STAT_NBRCTC_header struct {
	Version        string `json:"version"`
	Model          string `json:"model"`
	Desc           string `json:"desc"`
	Fcst_lead      string `json:"fcst_lead"`
	Fcst_valid_beg string `json:"fcst_valid_beg"`
	Fcst_valid_end string `json:"fcst_valid_end"`
	Obs_lead       string `json:"obs_lead"`
	Obs_valid_beg  string `json:"obs_valid_beg"`
	Obs_valid_end  string `json:"obs_valid_end"`
	Fcst_var       string `json:"fcst_var"`
	Fcst_units     string `json:"fcst_units"`
	Fcst_lev       string `json:"fcst_lev"`
	Obs_var        string `json:"obs_var"`
	Obs_units      string `json:"obs_units"`
	Obs_lev        string `json:"obs_lev"`
	Obtype         string `json:"obtype"`
	Vx_mask        string `json:"vx_mask"`
	Interp_mthd    string `json:"interp_mthd"`
	Interp_pnts    string `json:"interp_pnts"`
	Fcst_thresh    string `json:"fcst_thresh"`
	Obs_thresh     string `json:"obs_thresh"`
	Cov_thresh     string `json:"cov_thresh"`
	Alpha          string `json:"alpha"`
	Line_type      string `json:"line_type"`
}

type STAT_GRAD_header struct {
	Version        string `json:"version"`
	Model          string `json:"model"`
	Desc           string `json:"desc"`
	Fcst_lead      string `json:"fcst_lead"`
	Fcst_valid_beg string `json:"fcst_valid_beg"`
	Fcst_valid_end string `json:"fcst_valid_end"`
	Obs_lead       string `json:"obs_lead"`
	Obs_valid_beg  string `json:"obs_valid_beg"`
	Obs_valid_end  string `json:"obs_valid_end"`
	Fcst_var       string `json:"fcst_var"`
	Fcst_units     string `json:"fcst_units"`
	Fcst_lev       string `json:"fcst_lev"`
	Obs_var        string `json:"obs_var"`
	Obs_units      string `json:"obs_units"`
	Obs_lev        string `json:"obs_lev"`
	Obtype         string `json:"obtype"`
	Vx_mask        string `json:"vx_mask"`
	Interp_mthd    string `json:"interp_mthd"`
	Interp_pnts    string `json:"interp_pnts"`
	Fcst_thresh    string `json:"fcst_thresh"`
	Obs_thresh     string `json:"obs_thresh"`
	Cov_thresh     string `json:"cov_thresh"`
	Alpha          string `json:"alpha"`
	Line_type      string `json:"line_type"`
}

type STAT_PCT_header struct {
	Version        string `json:"version"`
	Model          string `json:"model"`
	Desc           string `json:"desc"`
	Fcst_lead      string `json:"fcst_lead"`
	Fcst_valid_beg string `json:"fcst_valid_beg"`
	Fcst_valid_end string `json:"fcst_valid_end"`
	Obs_lead       string `json:"obs_lead"`
	Obs_valid_beg  string `json:"obs_valid_beg"`
	Obs_valid_end  string `json:"obs_valid_end"`
	Fcst_var       string `json:"fcst_var"`
	Fcst_units     string `json:"fcst_units"`
	Fcst_lev       string `json:"fcst_lev"`
	Obs_var        string `json:"obs_var"`
	Obs_units      string `json:"obs_units"`
	Obs_lev        string `json:"obs_lev"`
	Obtype         string `json:"obtype"`
	Vx_mask        string `json:"vx_mask"`
	Interp_mthd    string `json:"interp_mthd"`
	Interp_pnts    string `json:"interp_pnts"`
	Fcst_thresh    string `json:"fcst_thresh"`
	Obs_thresh     string `json:"obs_thresh"`
	Cov_thresh     string `json:"cov_thresh"`
	Alpha          string `json:"alpha"`
	Line_type      string `json:"line_type"`
}

type STAT_MCTS_header struct {
	Version        string `json:"version"`
	Model          string `json:"model"`
	Desc           string `json:"desc"`
	Fcst_lead      string `json:"fcst_lead"`
	Fcst_valid_beg string `json:"fcst_valid_beg"`
	Fcst_valid_end string `json:"fcst_valid_end"`
	Obs_lead       string `json:"obs_lead"`
	Obs_valid_beg  string `json:"obs_valid_beg"`
	Obs_valid_end  string `json:"obs_valid_end"`
	Fcst_var       string `json:"fcst_var"`
	Fcst_units     string `json:"fcst_units"`
	Fcst_lev       string `json:"fcst_lev"`
	Obs_var        string `json:"obs_var"`
	Obs_units      string `json:"obs_units"`
	Obs_lev        string `json:"obs_lev"`
	Obtype         string `json:"obtype"`
	Vx_mask        string `json:"vx_mask"`
	Interp_mthd    string `json:"interp_mthd"`
	Interp_pnts    string `json:"interp_pnts"`
	Fcst_thresh    string `json:"fcst_thresh"`
	Obs_thresh     string `json:"obs_thresh"`
	Cov_thresh     string `json:"cov_thresh"`
	Alpha          string `json:"alpha"`
	Line_type      string `json:"line_type"`
}

type STAT_ECLV_header struct {
	Version        string `json:"version"`
	Model          string `json:"model"`
	Desc           string `json:"desc"`
	Fcst_lead      string `json:"fcst_lead"`
	Fcst_valid_beg string `json:"fcst_valid_beg"`
	Fcst_valid_end string `json:"fcst_valid_end"`
	Obs_lead       string `json:"obs_lead"`
	Obs_valid_beg  string `json:"obs_valid_beg"`
	Obs_valid_end  string `json:"obs_valid_end"`
	Fcst_var       string `json:"fcst_var"`
	Fcst_units     string `json:"fcst_units"`
	Fcst_lev       string `json:"fcst_lev"`
	Obs_var        string `json:"obs_var"`
	Obs_units      string `json:"obs_units"`
	Obs_lev        string `json:"obs_lev"`
	Obtype         string `json:"obtype"`
	Vx_mask        string `json:"vx_mask"`
	Interp_mthd    string `json:"interp_mthd"`
	Interp_pnts    string `json:"interp_pnts"`
	Fcst_thresh    string `json:"fcst_thresh"`
	Obs_thresh     string `json:"obs_thresh"`
	Cov_thresh     string `json:"cov_thresh"`
	Alpha          string `json:"alpha"`
	Line_type      string `json:"line_type"`
}

type STAT_SAL1L2_header struct {
	Version        string `json:"version"`
	Model          string `json:"model"`
	Desc           string `json:"desc"`
	Fcst_lead      string `json:"fcst_lead"`
	Fcst_valid_beg string `json:"fcst_valid_beg"`
	Fcst_valid_end string `json:"fcst_valid_end"`
	Obs_lead       string `json:"obs_lead"`
	Obs_valid_beg  string `json:"obs_valid_beg"`
	Obs_valid_end  string `json:"obs_valid_end"`
	Fcst_var       string `json:"fcst_var"`
	Fcst_units     string `json:"fcst_units"`
	Fcst_lev       string `json:"fcst_lev"`
	Obs_var        string `json:"obs_var"`
	Obs_units      string `json:"obs_units"`
	Obs_lev        string `json:"obs_lev"`
	Obtype         string `json:"obtype"`
	Vx_mask        string `json:"vx_mask"`
	Interp_mthd    string `json:"interp_mthd"`
	Interp_pnts    string `json:"interp_pnts"`
	Fcst_thresh    string `json:"fcst_thresh"`
	Obs_thresh     string `json:"obs_thresh"`
	Cov_thresh     string `json:"cov_thresh"`
	Alpha          string `json:"alpha"`
	Line_type      string `json:"line_type"`
}

type STAT_VCNT_header struct {
	Version        string `json:"version"`
	Model          string `json:"model"`
	Desc           string `json:"desc"`
	Fcst_lead      string `json:"fcst_lead"`
	Fcst_valid_beg string `json:"fcst_valid_beg"`
	Fcst_valid_end string `json:"fcst_valid_end"`
	Obs_lead       string `json:"obs_lead"`
	Obs_valid_beg  string `json:"obs_valid_beg"`
	Obs_valid_end  string `json:"obs_valid_end"`
	Fcst_var       string `json:"fcst_var"`
	Fcst_units     string `json:"fcst_units"`
	Fcst_lev       string `json:"fcst_lev"`
	Obs_var        string `json:"obs_var"`
	Obs_units      string `json:"obs_units"`
	Obs_lev        string `json:"obs_lev"`
	Obtype         string `json:"obtype"`
	Vx_mask        string `json:"vx_mask"`
	Interp_mthd    string `json:"interp_mthd"`
	Interp_pnts    string `json:"interp_pnts"`
	Fcst_thresh    string `json:"fcst_thresh"`
	Obs_thresh     string `json:"obs_thresh"`
	Cov_thresh     string `json:"cov_thresh"`
	Alpha          string `json:"alpha"`
	Line_type      string `json:"line_type"`
}

//line data struct definitions
type STAT_ISC struct {
	Total    int     `json:"total"`
	Tile_dim int     `json:"tile_dim"`
	Tile_xll int     `json:"tile_xll"`
	Tile_yll int     `json:"tile_yll"`
	Nscale   int     `json:"nscale"`
	Iscale   int     `json:"iscale"`
	Mse      float64 `json:"mse"`
	Isc      int     `json:"isc"`
	Fenergy2 float64 `json:"fenergy2"`
	Oenergy2 float64 `json:"oenergy2"`
	Baser    float64 `json:"baser"`
	Fbias    float64 `json:"fbias"`
}

type STAT_GRAD struct {
	Total      int     `json:"total"`
	Fgbar      float64 `json:"fgbar"`
	Ogbar      float64 `json:"ogbar"`
	Mgbar      float64 `json:"mgbar"`
	Egbar      float64 `json:"egbar"`
	S1         string  `json:"s1"`
	S1_og      string  `json:"s1_og"`
	Fgog_ratio string  `json:"fgog_ratio"`
	Dx         int     `json:"dx"`
	Dy         int     `json:"dy"`
}

type STAT_PSTD struct {
	Total       int     `json:"total"`
	N_thresh    int     `json:"n_thresh"`
	Baser       float64 `json:"baser"`
	Baser_ncl   string  `json:"baser_ncl"`
	Baser_ncu   string  `json:"baser_ncu"`
	Reliability string  `json:"reliability"`
	Resolution  string  `json:"resolution"`
	Uncertainty string  `json:"uncertainty"`
	Roc_auc     string  `json:"roc_auc"`
	Brier       string  `json:"brier"`
	Brier_ncl   string  `json:"brier_ncl"`
	Brier_ncu   string  `json:"brier_ncu"`
	Briercl     string  `json:"briercl"`
	Briercl_ncl string  `json:"briercl_ncl"`
	Briercl_ncu string  `json:"briercl_ncu"`
	Bss         string  `json:"bss"`
	Bss_smpl    string  `json:"bss_smpl"`
	Thresh_i    string  `json:"thresh_i"`
}

type STAT_SAL1L2 struct {
	Total  int     `json:"total"`
	Fabar  float64 `json:"fabar"`
	Oabar  float64 `json:"oabar"`
	Foabar float64 `json:"foabar"`
	Ffabar float64 `json:"ffabar"`
	Ooabar float64 `json:"ooabar"`
	Mae    float64 `json:"mae"`
}

type STAT_CTS struct {
	Total      int     `json:"total"`
	Baser      float64 `json:"baser"`
	Baser_ncl  string  `json:"baser_ncl"`
	Baser_ncu  string  `json:"baser_ncu"`
	Baser_bcl  string  `json:"baser_bcl"`
	Baser_bcu  string  `json:"baser_bcu"`
	Fmean      string  `json:"fmean"`
	Fmean_ncl  string  `json:"fmean_ncl"`
	Fmean_ncu  string  `json:"fmean_ncu"`
	Fmean_bcl  string  `json:"fmean_bcl"`
	Fmean_bcu  string  `json:"fmean_bcu"`
	Acc        string  `json:"acc"`
	Acc_ncl    string  `json:"acc_ncl"`
	Acc_ncu    string  `json:"acc_ncu"`
	Acc_bcl    string  `json:"acc_bcl"`
	Acc_bcu    string  `json:"acc_bcu"`
	Fbias      float64 `json:"fbias"`
	Fbias_bcl  string  `json:"fbias_bcl"`
	Fbias_bcu  string  `json:"fbias_bcu"`
	Pody       string  `json:"pody"`
	Pody_ncl   string  `json:"pody_ncl"`
	Pody_ncu   string  `json:"pody_ncu"`
	Pody_bcl   string  `json:"pody_bcl"`
	Pody_bcu   string  `json:"pody_bcu"`
	Podn       string  `json:"podn"`
	Podn_ncl   string  `json:"podn_ncl"`
	Podn_ncu   string  `json:"podn_ncu"`
	Podn_bcl   string  `json:"podn_bcl"`
	Podn_bcu   string  `json:"podn_bcu"`
	Pofd       string  `json:"pofd"`
	Pofd_ncl   string  `json:"pofd_ncl"`
	Pofd_ncu   string  `json:"pofd_ncu"`
	Pofd_bcl   string  `json:"pofd_bcl"`
	Pofd_bcu   string  `json:"pofd_bcu"`
	Far        string  `json:"far"`
	Far_ncl    string  `json:"far_ncl"`
	Far_ncu    string  `json:"far_ncu"`
	Far_bcl    string  `json:"far_bcl"`
	Far_bcu    string  `json:"far_bcu"`
	Csi        string  `json:"csi"`
	Csi_ncl    string  `json:"csi_ncl"`
	Csi_ncu    string  `json:"csi_ncu"`
	Csi_bcl    string  `json:"csi_bcl"`
	Csi_bcu    string  `json:"csi_bcu"`
	Gss        string  `json:"gss"`
	Gss_bcl    string  `json:"gss_bcl"`
	Gss_bcu    string  `json:"gss_bcu"`
	Hk         string  `json:"hk"`
	Hk_ncl     string  `json:"hk_ncl"`
	Hk_ncu     string  `json:"hk_ncu"`
	Hk_bcl     string  `json:"hk_bcl"`
	Hk_bcu     string  `json:"hk_bcu"`
	Hss        string  `json:"hss"`
	Hss_bcl    string  `json:"hss_bcl"`
	Hss_bcu    string  `json:"hss_bcu"`
	Odds       string  `json:"odds"`
	Odds_ncl   string  `json:"odds_ncl"`
	Odds_ncu   string  `json:"odds_ncu"`
	Odds_bcl   string  `json:"odds_bcl"`
	Odds_bcu   string  `json:"odds_bcu"`
	Lodds      string  `json:"lodds"`
	Lodds_ncl  string  `json:"lodds_ncl"`
	Lodds_ncu  string  `json:"lodds_ncu"`
	Lodds_bcl  string  `json:"lodds_bcl"`
	Lodds_bcu  string  `json:"lodds_bcu"`
	Orss       string  `json:"orss"`
	Orss_ncl   string  `json:"orss_ncl"`
	Orss_ncu   string  `json:"orss_ncu"`
	Orss_bcl   string  `json:"orss_bcl"`
	Orss_bcu   string  `json:"orss_bcu"`
	Eds        string  `json:"eds"`
	Eds_ncl    string  `json:"eds_ncl"`
	Eds_ncu    string  `json:"eds_ncu"`
	Eds_bcl    string  `json:"eds_bcl"`
	Eds_bcu    string  `json:"eds_bcu"`
	Seds       string  `json:"seds"`
	Seds_ncl   string  `json:"seds_ncl"`
	Seds_ncu   string  `json:"seds_ncu"`
	Seds_bcl   string  `json:"seds_bcl"`
	Seds_bcu   string  `json:"seds_bcu"`
	Edi        string  `json:"edi"`
	Edi_ncl    string  `json:"edi_ncl"`
	Edi_ncu    string  `json:"edi_ncu"`
	Edi_bcl    string  `json:"edi_bcl"`
	Edi_bcu    string  `json:"edi_bcu"`
	Sedi       string  `json:"sedi"`
	Sedi_ncl   string  `json:"sedi_ncl"`
	Sedi_ncu   string  `json:"sedi_ncu"`
	Sedi_bcl   string  `json:"sedi_bcl"`
	Sedi_bcu   string  `json:"sedi_bcu"`
	Bagss      string  `json:"bagss"`
	Bagss_bcl  string  `json:"bagss_bcl"`
	Bagss_bcu  string  `json:"bagss_bcu"`
	Hss_ec     string  `json:"hss_ec"`
	Hss_ec_bcl string  `json:"hss_ec_bcl"`
	Hss_ec_bcu string  `json:"hss_ec_bcu"`
	Ec_value   float64 `json:"ec_value"`
}

type STAT_FHO struct {
	Total  int     `json:"total"`
	F_rate float64 `json:"f_rate"`
	H_rate float64 `json:"h_rate"`
	O_rate float64 `json:"o_rate"`
}

type STAT_MPR struct {
	Total            int     `json:"total"`
	Index            int     `json:"index"`
	Obs_sid          string  `json:"obs_sid"`
	Obs_lat          float64 `json:"obs_lat"`
	Obs_lon          float64 `json:"obs_lon"`
	Obs_lvl          float64 `json:"obs_lvl"`
	Obs_elv          float64 `json:"obs_elv"`
	Fcst             float64 `json:"fcst"`
	Obs              float64 `json:"obs"`
	Obs_qc           string  `json:"obs_qc"`
	Obs_climo_mean   float64 `json:"obs_climo_mean"`
	Obs_climo_stdev  float64 `json:"obs_climo_stdev"`
	Obs_climo_cdf    float64 `json:"obs_climo_cdf"`
	Fcst_climo_mean  float64 `json:"fcst_climo_mean"`
	Fcst_climo_stdev float64 `json:"fcst_climo_stdev"`
}

type STAT_DMAP struct {
	Total      int     `json:"total"`
	Fy         float64 `json:"fy"`
	Oy         float64 `json:"oy"`
	Fbias      float64 `json:"fbias"`
	Baddeley   string  `json:"baddeley"`
	Hausdorff  string  `json:"hausdorff"`
	Med_fo     string  `json:"med_fo"`
	Med_of     string  `json:"med_of"`
	Med_min    string  `json:"med_min"`
	Med_max    string  `json:"med_max"`
	Med_mean   string  `json:"med_mean"`
	Fom_fo     string  `json:"fom_fo"`
	Fom_of     string  `json:"fom_of"`
	Fom_min    string  `json:"fom_min"`
	Fom_max    string  `json:"fom_max"`
	Fom_mean   string  `json:"fom_mean"`
	Zhu_fo     string  `json:"zhu_fo"`
	Zhu_of     string  `json:"zhu_of"`
	Zhu_min    string  `json:"zhu_min"`
	Zhu_max    string  `json:"zhu_max"`
	Zhu_mean   string  `json:"zhu_mean"`
	G          float64 `json:"g"`
	Gbeta      string  `json:"gbeta"`
	Beta_value string  `json:"beta_value"`
}

type STAT_ORANK struct {
	Total            int     `json:"total"`
	Index            int     `json:"index"`
	Obs_sid          string  `json:"obs_sid"`
	Obs_lat          float64 `json:"obs_lat"`
	Obs_lon          float64 `json:"obs_lon"`
	Obs_lvl          float64 `json:"obs_lvl"`
	Obs_elv          float64 `json:"obs_elv"`
	Obs              float64 `json:"obs"`
	Pit              float64 `json:"pit"`
	Rank             int     `json:"rank"`
	N_ens_vld        int     `json:"n_ens_vld"`
	N_ens            float64 `json:"n_ens"`
	Ens_i            string  `json:"ens_i"`
	Obs_qc           string  `json:"obs_qc"`
	Ens_mean         float64 `json:"ens_mean"`
	Obs_climo_mean   float64 `json:"obs_climo_mean"`
	Spread           float64 `json:"spread"`
	Ens_mean_oerr    float64 `json:"ens_mean_oerr"`
	Spread_oerr      float64 `json:"spread_oerr"`
	Spread_plus_oerr float64 `json:"spread_plus_oerr"`
	Obs_climo_stdev  float64 `json:"obs_climo_stdev"`
	Fcst_climo_mean  float64 `json:"fcst_climo_mean"`
	Fcst_climo_stdev float64 `json:"fcst_climo_stdev"`
}

type STAT_RPS struct {
	Total     int     `json:"total"`
	N_prob    float64 `json:"n_prob"`
	Rps_rel   float64 `json:"rps_rel"`
	Rps_res   float64 `json:"rps_res"`
	Rps_unc   float64 `json:"rps_unc"`
	Rps       float64 `json:"rps"`
	Rpss      float64 `json:"rpss"`
	Rpss_smpl float64 `json:"rpss_smpl"`
	Rps_comp  string  `json:"rps_comp"`
}

type STAT_VL1L2 struct {
	Total       int     `json:"total"`
	Ufbar       float64 `json:"ufbar"`
	Vfbar       float64 `json:"vfbar"`
	Uobar       float64 `json:"uobar"`
	Vobar       float64 `json:"vobar"`
	Uvfobar     float64 `json:"uvfobar"`
	Uvffbar     float64 `json:"uvffbar"`
	Uvoobar     float64 `json:"uvoobar"`
	F_speed_bar float64 `json:"f_speed_bar"`
	O_speed_bar float64 `json:"o_speed_bar"`
	Total_dir   int     `json:"total_dir"`
	Dir_me      float64 `json:"dir_me"`
	Dir_mae     float64 `json:"dir_mae"`
	Dir_mse     float64 `json:"dir_mse"`
}

type MODE_OBJ struct {
	Object_id                  string `json:"object_id"`
	Object_cat                 string `json:"object_cat"`
	Centroid_x                 string `json:"centroid_x"`
	Centroid_y                 string `json:"centroid_y"`
	Centroid_lat               string `json:"centroid_lat"`
	Centroid_lon               string `json:"centroid_lon"`
	Axis_ang                   string `json:"axis_ang"`
	Length                     string `json:"length"`
	Width                      string `json:"width"`
	Area                       string `json:"area"`
	Area_thresh                string `json:"area_thresh"`
	Curvature                  string `json:"curvature"`
	Curvature_x                string `json:"curvature_x"`
	Curvature_y                string `json:"curvature_y"`
	Complexity                 string `json:"complexity"`
	Intensity_10               string `json:"intensity_10"`
	Intensity_25               string `json:"intensity_25"`
	Intensity_50               string `json:"intensity_50"`
	Intensity_75               string `json:"intensity_75"`
	Intensity_90               string `json:"intensity_90"`
	Intensity_user             string `json:"intensity_user"`
	Intensity_sum              string `json:"intensity_sum"`
	Centroid_dist              string `json:"centroid_dist"`
	Boundary_dist              string `json:"boundary_dist"`
	Convex_hull_dist           string `json:"convex_hull_dist"`
	Angle_diff                 string `json:"angle_diff"`
	Aspect_diff                string `json:"aspect_diff"`
	Area_ratio                 string `json:"area_ratio"`
	Intersection_area          string `json:"intersection_area"`
	Union_area                 string `json:"union_area"`
	Symmetric_diff             string `json:"symmetric_diff"`
	Intersection_over_area     string `json:"intersection_over_area"`
	Curvature_ratio            string `json:"curvature_ratio"`
	Complexity_ratio           string `json:"complexity_ratio"`
	Percentile_intensity_ratio string `json:"percentile_intensity_ratio"`
	Interest                   string `json:"interest"`
}

type STAT_CTC struct {
	Total    int     `json:"total"`
	Fy_oy    float64 `json:"fy_oy"`
	Fy_on    float64 `json:"fy_on"`
	Fn_oy    float64 `json:"fn_oy"`
	Fn_on    float64 `json:"fn_on"`
	Ec_value float64 `json:"ec_value"`
}

type TCST_PROBRIRW struct {
	Alat        string `json:"alat"`
	Alon        string `json:"alon"`
	Blat        string `json:"blat"`
	Blon        string `json:"blon"`
	Initials    string `json:"initials"`
	Tk_err      string `json:"tk_err"`
	X_err       string `json:"x_err"`
	Y_err       string `json:"y_err"`
	Adland      string `json:"adland"`
	Bdland      string `json:"bdland"`
	Rirw_beg    string `json:"rirw_beg"`
	Rirw_end    string `json:"rirw_end"`
	Rirw_window string `json:"rirw_window"`
	Awind_end   string `json:"awind_end"`
	Bwind_beg   string `json:"bwind_beg"`
	Bwind_end   string `json:"bwind_end"`
	Bdelta      string `json:"bdelta"`
	Bdelta_max  string `json:"bdelta_max"`
	Blevel_beg  string `json:"blevel_beg"`
	Blevel_end  string `json:"blevel_end"`
	N_thresh    int    `json:"n_thresh"`
	Thresh_i    string `json:"thresh_i"`
	Prob_i      string `json:"prob_i"`
}

type TCST_TCMPR struct {
	Total          int    `json:"total"`
	Index          int    `json:"index"`
	Level          string `json:"level"`
	Watch_warn     string `json:"watch_warn"`
	Initials       string `json:"initials"`
	Alat           string `json:"alat"`
	Alon           string `json:"alon"`
	Blat           string `json:"blat"`
	Blon           string `json:"blon"`
	Tk_err         string `json:"tk_err"`
	X_err          string `json:"x_err"`
	Y_err          string `json:"y_err"`
	Altk_err       string `json:"altk_err"`
	Crtk_err       string `json:"crtk_err"`
	Adland         string `json:"adland"`
	Bdland         string `json:"bdland"`
	Amslp          string `json:"amslp"`
	Bmslp          string `json:"bmslp"`
	Amax_wind      string `json:"amax_wind"`
	Bmax_wind      string `json:"bmax_wind"`
	Aal_wind_34    string `json:"aal_wind_34"`
	Bal_wind_34    string `json:"bal_wind_34"`
	Ane_wind_34    string `json:"ane_wind_34"`
	Bne_wind_34    string `json:"bne_wind_34"`
	Ase_wind_34    string `json:"ase_wind_34"`
	Bse_wind_34    string `json:"bse_wind_34"`
	Asw_wind_34    string `json:"asw_wind_34"`
	Bsw_wind_34    string `json:"bsw_wind_34"`
	Anw_wind_34    string `json:"anw_wind_34"`
	Bnw_wind_34    string `json:"bnw_wind_34"`
	Aal_wind_50    string `json:"aal_wind_50"`
	Bal_wind_50    string `json:"bal_wind_50"`
	Ane_wind_50    string `json:"ane_wind_50"`
	Bne_wind_50    string `json:"bne_wind_50"`
	Ase_wind_50    string `json:"ase_wind_50"`
	Bse_wind_50    string `json:"bse_wind_50"`
	Asw_wind_50    string `json:"asw_wind_50"`
	Bsw_wind_50    string `json:"bsw_wind_50"`
	Anw_wind_50    string `json:"anw_wind_50"`
	Bnw_wind_50    string `json:"bnw_wind_50"`
	Aal_wind_64    string `json:"aal_wind_64"`
	Bal_wind_64    string `json:"bal_wind_64"`
	Ane_wind_64    string `json:"ane_wind_64"`
	Bne_wind_64    string `json:"bne_wind_64"`
	Ase_wind_64    string `json:"ase_wind_64"`
	Bse_wind_64    string `json:"bse_wind_64"`
	Asw_wind_64    string `json:"asw_wind_64"`
	Bsw_wind_64    string `json:"bsw_wind_64"`
	Anw_wind_64    string `json:"anw_wind_64"`
	Bnw_wind_64    string `json:"bnw_wind_64"`
	Aradp          string `json:"aradp"`
	Bradp          string `json:"bradp"`
	Arrp           string `json:"arrp"`
	Brrp           string `json:"brrp"`
	Amrd           string `json:"amrd"`
	Bmrd           string `json:"bmrd"`
	Agusts         string `json:"agusts"`
	Bgusts         string `json:"bgusts"`
	Aeye           string `json:"aeye"`
	Beye           string `json:"beye"`
	Adir           string `json:"adir"`
	Bdir           string `json:"bdir"`
	Aspeed         string `json:"aspeed"`
	Bspeed         string `json:"bspeed"`
	Adepth         string `json:"adepth"`
	Bdepth         string `json:"bdepth"`
	Num_members    string `json:"num_members"`
	Track_spread   string `json:"track_spread"`
	Track_stdev    string `json:"track_stdev"`
	Mslp_stdev     string `json:"mslp_stdev"`
	Max_wind_stdev string `json:"max_wind_stdev"`
}

type STAT_CNT struct {
	Total                int     `json:"total"`
	Fbar                 float64 `json:"fbar"`
	Fbar_ncl             string  `json:"fbar_ncl"`
	Fbar_ncu             string  `json:"fbar_ncu"`
	Fbar_bcl             string  `json:"fbar_bcl"`
	Fbar_bcu             string  `json:"fbar_bcu"`
	Fstdev               string  `json:"fstdev"`
	Fstdev_ncl           string  `json:"fstdev_ncl"`
	Fstdev_ncu           string  `json:"fstdev_ncu"`
	Fstdev_bcl           string  `json:"fstdev_bcl"`
	Fstdev_bcu           string  `json:"fstdev_bcu"`
	Obar                 float64 `json:"obar"`
	Obar_ncl             string  `json:"obar_ncl"`
	Obar_ncu             string  `json:"obar_ncu"`
	Obar_bcl             string  `json:"obar_bcl"`
	Obar_bcu             string  `json:"obar_bcu"`
	Ostdev               string  `json:"ostdev"`
	Ostdev_ncl           string  `json:"ostdev_ncl"`
	Ostdev_ncu           string  `json:"ostdev_ncu"`
	Ostdev_bcl           string  `json:"ostdev_bcl"`
	Ostdev_bcu           string  `json:"ostdev_bcu"`
	Pr_corr              string  `json:"pr_corr"`
	Pr_corr_ncl          string  `json:"pr_corr_ncl"`
	Pr_corr_ncu          string  `json:"pr_corr_ncu"`
	Pr_corr_bcl          string  `json:"pr_corr_bcl"`
	Pr_corr_bcu          string  `json:"pr_corr_bcu"`
	Sp_corr              string  `json:"sp_corr"`
	Kt_corr              string  `json:"kt_corr"`
	Ranks                string  `json:"ranks"`
	Frank_ties           string  `json:"frank_ties"`
	Orank_ties           string  `json:"orank_ties"`
	Me                   float64 `json:"me"`
	Me_ncl               string  `json:"me_ncl"`
	Me_ncu               string  `json:"me_ncu"`
	Me_bcl               string  `json:"me_bcl"`
	Me_bcu               string  `json:"me_bcu"`
	Estdev               string  `json:"estdev"`
	Estdev_ncl           string  `json:"estdev_ncl"`
	Estdev_ncu           string  `json:"estdev_ncu"`
	Estdev_bcl           string  `json:"estdev_bcl"`
	Estdev_bcu           string  `json:"estdev_bcu"`
	Mbias                string  `json:"mbias"`
	Mbias_bcl            string  `json:"mbias_bcl"`
	Mbias_bcu            string  `json:"mbias_bcu"`
	Mae                  float64 `json:"mae"`
	Mae_bcl              string  `json:"mae_bcl"`
	Mae_bcu              string  `json:"mae_bcu"`
	Mse                  float64 `json:"mse"`
	Mse_bcl              string  `json:"mse_bcl"`
	Mse_bcu              string  `json:"mse_bcu"`
	Bcmse                string  `json:"bcmse"`
	Bcmse_bcl            string  `json:"bcmse_bcl"`
	Bcmse_bcu            string  `json:"bcmse_bcu"`
	Rmse                 float64 `json:"rmse"`
	Rmse_bcl             string  `json:"rmse_bcl"`
	Rmse_bcu             string  `json:"rmse_bcu"`
	E10                  string  `json:"e10"`
	E10_bcl              string  `json:"e10_bcl"`
	E10_bcu              string  `json:"e10_bcu"`
	E25                  string  `json:"e25"`
	E25_bcl              string  `json:"e25_bcl"`
	E25_bcu              string  `json:"e25_bcu"`
	E50                  string  `json:"e50"`
	E50_bcl              string  `json:"e50_bcl"`
	E50_bcu              string  `json:"e50_bcu"`
	E75                  string  `json:"e75"`
	E75_bcl              string  `json:"e75_bcl"`
	E75_bcu              string  `json:"e75_bcu"`
	E90                  string  `json:"e90"`
	E90_bcl              string  `json:"e90_bcl"`
	E90_bcu              string  `json:"e90_bcu"`
	Eiqr                 string  `json:"eiqr"`
	Eiqr_bcl             string  `json:"eiqr_bcl"`
	Eiqr_bcu             string  `json:"eiqr_bcu"`
	Mad                  string  `json:"mad"`
	Mad_bcl              string  `json:"mad_bcl"`
	Mad_bcu              string  `json:"mad_bcu"`
	Anom_corr            string  `json:"anom_corr"`
	Anom_corr_ncl        string  `json:"anom_corr_ncl"`
	Anom_corr_ncu        string  `json:"anom_corr_ncu"`
	Anom_corr_bcl        string  `json:"anom_corr_bcl"`
	Anom_corr_bcu        string  `json:"anom_corr_bcu"`
	Me2                  string  `json:"me2"`
	Me2_bcl              string  `json:"me2_bcl"`
	Me2_bcu              string  `json:"me2_bcu"`
	Msess                string  `json:"msess"`
	Msess_bcl            string  `json:"msess_bcl"`
	Msess_bcu            string  `json:"msess_bcu"`
	Rmsfa                string  `json:"rmsfa"`
	Rmsfa_bcl            string  `json:"rmsfa_bcl"`
	Rmsfa_bcu            string  `json:"rmsfa_bcu"`
	Rmsoa                string  `json:"rmsoa"`
	Rmsoa_bcl            string  `json:"rmsoa_bcl"`
	Rmsoa_bcu            string  `json:"rmsoa_bcu"`
	Anom_corr_uncntr     string  `json:"anom_corr_uncntr"`
	Anom_corr_uncntr_bcl string  `json:"anom_corr_uncntr_bcl"`
	Anom_corr_uncntr_bcu string  `json:"anom_corr_uncntr_bcu"`
	Si                   float64 `json:"si"`
	Si_bcl               string  `json:"si_bcl"`
	Si_bcu               string  `json:"si_bcu"`
}

type TCST_TCDIAG struct {
	Total        int    `json:"total"`
	Index        int    `json:"index"`
	Diag_source  string `json:"diag_source"`
	Track_source string `json:"track_source"`
	Field_source string `json:"field_source"`
	N_diag       string `json:"n_diag"`
	Diag_i       string `json:"diag_i"`
	Value_i      string `json:"value_i"`
}

type STAT_VCNT struct {
	Total                int     `json:"total"`
	Fbar                 float64 `json:"fbar"`
	Fbar_bcl             string  `json:"fbar_bcl"`
	Fbar_bcu             string  `json:"fbar_bcu"`
	Obar                 float64 `json:"obar"`
	Obar_bcl             string  `json:"obar_bcl"`
	Obar_bcu             string  `json:"obar_bcu"`
	Fs_rms               string  `json:"fs_rms"`
	Fs_rms_bcl           string  `json:"fs_rms_bcl"`
	Fs_rms_bcu           string  `json:"fs_rms_bcu"`
	Os_rms               string  `json:"os_rms"`
	Os_rms_bcl           string  `json:"os_rms_bcl"`
	Os_rms_bcu           string  `json:"os_rms_bcu"`
	Msve                 string  `json:"msve"`
	Msve_bcl             string  `json:"msve_bcl"`
	Msve_bcu             string  `json:"msve_bcu"`
	Rmsve                string  `json:"rmsve"`
	Rmsve_bcl            string  `json:"rmsve_bcl"`
	Rmsve_bcu            string  `json:"rmsve_bcu"`
	Fstdev               string  `json:"fstdev"`
	Fstdev_bcl           string  `json:"fstdev_bcl"`
	Fstdev_bcu           string  `json:"fstdev_bcu"`
	Ostdev               string  `json:"ostdev"`
	Ostdev_bcl           string  `json:"ostdev_bcl"`
	Ostdev_bcu           string  `json:"ostdev_bcu"`
	Fdir                 string  `json:"fdir"`
	Fdir_bcl             string  `json:"fdir_bcl"`
	Fdir_bcu             string  `json:"fdir_bcu"`
	Odir                 string  `json:"odir"`
	Odir_bcl             string  `json:"odir_bcl"`
	Odir_bcu             string  `json:"odir_bcu"`
	Fbar_speed           string  `json:"fbar_speed"`
	Fbar_speed_bcl       string  `json:"fbar_speed_bcl"`
	Fbar_speed_bcu       string  `json:"fbar_speed_bcu"`
	Obar_speed           string  `json:"obar_speed"`
	Obar_speed_bcl       string  `json:"obar_speed_bcl"`
	Obar_speed_bcu       string  `json:"obar_speed_bcu"`
	Vdiff_speed          string  `json:"vdiff_speed"`
	Vdiff_speed_bcl      string  `json:"vdiff_speed_bcl"`
	Vdiff_speed_bcu      string  `json:"vdiff_speed_bcu"`
	Vdiff_dir            string  `json:"vdiff_dir"`
	Vdiff_dir_bcl        string  `json:"vdiff_dir_bcl"`
	Vdiff_dir_bcu        string  `json:"vdiff_dir_bcu"`
	Speed_err            string  `json:"speed_err"`
	Speed_err_bcl        string  `json:"speed_err_bcl"`
	Speed_err_bcu        string  `json:"speed_err_bcu"`
	Speed_abserr         string  `json:"speed_abserr"`
	Speed_abserr_bcl     string  `json:"speed_abserr_bcl"`
	Speed_abserr_bcu     string  `json:"speed_abserr_bcu"`
	Dir_err              string  `json:"dir_err"`
	Dir_err_bcl          string  `json:"dir_err_bcl"`
	Dir_err_bcu          string  `json:"dir_err_bcu"`
	Dir_abserr           string  `json:"dir_abserr"`
	Dir_abserr_bcl       string  `json:"dir_abserr_bcl"`
	Dir_abserr_bcu       string  `json:"dir_abserr_bcu"`
	Anom_corr            string  `json:"anom_corr"`
	Anom_corr_ncl        string  `json:"anom_corr_ncl"`
	Anom_corr_ncu        string  `json:"anom_corr_ncu"`
	Anom_corr_bcl        string  `json:"anom_corr_bcl"`
	Anom_corr_bcu        string  `json:"anom_corr_bcu"`
	Anom_corr_uncntr     string  `json:"anom_corr_uncntr"`
	Anom_corr_uncntr_bcl string  `json:"anom_corr_uncntr_bcl"`
	Anom_corr_uncntr_bcu string  `json:"anom_corr_uncntr_bcu"`
	Total_dir            int     `json:"total_dir"`
	Dir_me               float64 `json:"dir_me"`
	Dir_me_bcl           string  `json:"dir_me_bcl"`
	Dir_me_bcu           string  `json:"dir_me_bcu"`
	Dir_mae              float64 `json:"dir_mae"`
	Dir_mae_bcl          string  `json:"dir_mae_bcl"`
	Dir_mae_bcu          string  `json:"dir_mae_bcu"`
	Dir_mse              float64 `json:"dir_mse"`
	Dir_mse_bcl          string  `json:"dir_mse_bcl"`
	Dir_mse_bcu          string  `json:"dir_mse_bcu"`
	Dir_rmse             string  `json:"dir_rmse"`
	Dir_rmse_bcl         string  `json:"dir_rmse_bcl"`
	Dir_rmse_bcu         string  `json:"dir_rmse_bcu"`
}

type STAT_RELP struct {
	Total  int     `json:"total"`
	N_ens  float64 `json:"n_ens"`
	Relp_i string  `json:"relp_i"`
}

type STAT_SL1L2 struct {
	Total int     `json:"total"`
	Fbar  float64 `json:"fbar"`
	Obar  float64 `json:"obar"`
	Fobar float64 `json:"fobar"`
	Ffbar float64 `json:"ffbar"`
	Oobar float64 `json:"oobar"`
	Mae   float64 `json:"mae"`
}

type STAT_PJC struct {
	Total         int    `json:"total"`
	N_thresh      int    `json:"n_thresh"`
	Thresh_i      string `json:"thresh_i"`
	Oy_tp_i       string `json:"oy_tp_i"`
	On_tp_i       string `json:"on_tp_i"`
	Calibration_i string `json:"calibration_i"`
	Refinement_i  string `json:"refinement_i"`
	Likelihood_i  string `json:"likelihood_i"`
	Baser_i       string `json:"baser_i"`
}

type STAT_MCTS struct {
	Total      int     `json:"total"`
	N_cat      int     `json:"n_cat"`
	Acc        string  `json:"acc"`
	Acc_ncl    string  `json:"acc_ncl"`
	Acc_ncu    string  `json:"acc_ncu"`
	Acc_bcl    string  `json:"acc_bcl"`
	Acc_bcu    string  `json:"acc_bcu"`
	Hk         string  `json:"hk"`
	Hk_bcl     string  `json:"hk_bcl"`
	Hk_bcu     string  `json:"hk_bcu"`
	Hss        string  `json:"hss"`
	Hss_bcl    string  `json:"hss_bcl"`
	Hss_bcu    string  `json:"hss_bcu"`
	Ger        string  `json:"ger"`
	Ger_bcl    string  `json:"ger_bcl"`
	Ger_bcu    string  `json:"ger_bcu"`
	Hss_ec     string  `json:"hss_ec"`
	Hss_ec_bcl string  `json:"hss_ec_bcl"`
	Hss_ec_bcu string  `json:"hss_ec_bcu"`
	Ec_value   float64 `json:"ec_value"`
}

type STAT_RHIST struct {
	Total  int    `json:"total"`
	N_rank int    `json:"n_rank"`
	Rank_i string `json:"rank_i"`
}

type STAT_PHIST struct {
	Total    int     `json:"total"`
	Bin_size float64 `json:"bin_size"`
	N_bin    int     `json:"n_bin"`
	Bin_i    string  `json:"bin_i"`
}

type STAT_SSVAR struct {
	Total       int     `json:"total"`
	N_bin       int     `json:"n_bin"`
	Bin_i       string  `json:"bin_i"`
	Bin_n       int     `json:"bin_n"`
	Var_min     float64 `json:"var_min"`
	Var_max     float64 `json:"var_max"`
	Var_mean    float64 `json:"var_mean"`
	Fbar        float64 `json:"fbar"`
	Obar        float64 `json:"obar"`
	Fobar       float64 `json:"fobar"`
	Ffbar       float64 `json:"ffbar"`
	Oobar       float64 `json:"oobar"`
	Fbar_ncl    string  `json:"fbar_ncl"`
	Fbar_ncu    string  `json:"fbar_ncu"`
	Fstdev      string  `json:"fstdev"`
	Fstdev_ncl  string  `json:"fstdev_ncl"`
	Fstdev_ncu  string  `json:"fstdev_ncu"`
	Obar_ncl    string  `json:"obar_ncl"`
	Obar_ncu    string  `json:"obar_ncu"`
	Ostdev      string  `json:"ostdev"`
	Ostdev_ncl  string  `json:"ostdev_ncl"`
	Ostdev_ncu  string  `json:"ostdev_ncu"`
	Pr_corr     string  `json:"pr_corr"`
	Pr_corr_ncl string  `json:"pr_corr_ncl"`
	Pr_corr_ncu string  `json:"pr_corr_ncu"`
	Me          float64 `json:"me"`
	Me_ncl      string  `json:"me_ncl"`
	Me_ncu      string  `json:"me_ncu"`
	Estdev      string  `json:"estdev"`
	Estdev_ncl  string  `json:"estdev_ncl"`
	Estdev_ncu  string  `json:"estdev_ncu"`
	Mbias       string  `json:"mbias"`
	Mse         float64 `json:"mse"`
	Bcmse       string  `json:"bcmse"`
	Rmse        float64 `json:"rmse"`
}

type STAT_SSIDX struct {
	Fcst_model string `json:"fcst_model"`
	Ref_model  string `json:"ref_model"`
	N_init     string `json:"n_init"`
	N_term     string `json:"n_term"`
	N_vld      string `json:"n_vld"`
	Ss_index   string `json:"ss_index"`
}

type STAT_MCTC struct {
	Total    int     `json:"total"`
	N_cat    int     `json:"n_cat"`
	Fi_oi    string  `json:"fi_oi"`
	Ec_value float64 `json:"ec_value"`
}

type STAT_SEEPS_MPR struct {
	Obs_sid  string  `json:"obs_sid"`
	Obs_lat  float64 `json:"obs_lat"`
	Obs_lon  float64 `json:"obs_lon"`
	Fcst     float64 `json:"fcst"`
	Obs      float64 `json:"obs"`
	Obs_qc   string  `json:"obs_qc"`
	Fcst_cat int     `json:"fcst_cat"`
	Obs_cat  int     `json:"obs_cat"`
	P1       float64 `json:"p1"`
	P2       float64 `json:"p2"`
	T1       float64 `json:"t1"`
	T2       float64 `json:"t2"`
	Seeps    float64 `json:"seeps"`
}

type STAT_NBRCNT struct {
	Total      int     `json:"total"`
	Fbs        float64 `json:"fbs"`
	Fbs_bcl    string  `json:"fbs_bcl"`
	Fbs_bcu    string  `json:"fbs_bcu"`
	Fss        float64 `json:"fss"`
	Fss_bcl    string  `json:"fss_bcl"`
	Fss_bcu    string  `json:"fss_bcu"`
	Afss       float64 `json:"afss"`
	Afss_bcl   string  `json:"afss_bcl"`
	Afss_bcu   string  `json:"afss_bcu"`
	Ufss       float64 `json:"ufss"`
	Ufss_bcl   string  `json:"ufss_bcl"`
	Ufss_bcu   string  `json:"ufss_bcu"`
	F_rate     float64 `json:"f_rate"`
	F_rate_bcl string  `json:"f_rate_bcl"`
	F_rate_bcu string  `json:"f_rate_bcu"`
	O_rate     float64 `json:"o_rate"`
	O_rate_bcl string  `json:"o_rate_bcl"`
	O_rate_bcu string  `json:"o_rate_bcu"`
}

type STAT_NBRCTC struct {
	Total int     `json:"total"`
	Fy_oy float64 `json:"fy_oy"`
	Fy_on float64 `json:"fy_on"`
	Fn_oy float64 `json:"fn_oy"`
	Fn_on float64 `json:"fn_on"`
}

type STAT_PCT struct {
	Total    int    `json:"total"`
	N_thresh int    `json:"n_thresh"`
	Thresh_i string `json:"thresh_i"`
	Oy_i     string `json:"oy_i"`
	On_i     string `json:"on_i"`
}

type STAT_PRC struct {
	Total    int    `json:"total"`
	N_thresh int    `json:"n_thresh"`
	Thresh_i string `json:"thresh_i"`
	Pody_i   string `json:"pody_i"`
	Pofd_i   string `json:"pofd_i"`
}

type STAT_ECNT struct {
	Total            int     `json:"total"`
	N_ens            float64 `json:"n_ens"`
	Crps             float64 `json:"crps"`
	Crpss            float64 `json:"crpss"`
	Ign              float64 `json:"ign"`
	Me               float64 `json:"me"`
	Rmse             float64 `json:"rmse"`
	Spread           float64 `json:"spread"`
	Me_oerr          float64 `json:"me_oerr"`
	Rmse_oerr        float64 `json:"rmse_oerr"`
	Spread_oerr      float64 `json:"spread_oerr"`
	Spread_plus_oerr float64 `json:"spread_plus_oerr"`
	Crpscl           float64 `json:"crpscl"`
	Crps_emp         float64 `json:"crps_emp"`
	Crpscl_emp       float64 `json:"crpscl_emp"`
	Crpss_emp        float64 `json:"crpss_emp"`
	Crps_emp_fair    float64 `json:"crps_emp_fair"`
	Spread_md        float64 `json:"spread_md"`
	Mae              float64 `json:"mae"`
	Mae_oerr         float64 `json:"mae_oerr"`
	Bias_ratio       float64 `json:"bias_ratio"`
	N_ge_obs         int     `json:"n_ge_obs"`
	Me_ge_obs        float64 `json:"me_ge_obs"`
	N_lt_obs         int     `json:"n_lt_obs"`
	Me_lt_obs        float64 `json:"me_lt_obs"`
	Ign_conv_oerr    float64 `json:"ign_conv_oerr"`
	Ign_corr_oerr    float64 `json:"ign_corr_oerr"`
}

type STAT_GENMPR struct {
	Total      int    `json:"total"`
	Index      int    `json:"index"`
	Storm_id   string `json:"storm_id"`
	Prob_lead  string `json:"prob_lead"`
	Prob_val   string `json:"prob_val"`
	Agen_init  string `json:"agen_init"`
	Agen_fhr   string `json:"agen_fhr"`
	Agen_lat   string `json:"agen_lat"`
	Agen_lon   string `json:"agen_lon"`
	Agen_dland string `json:"agen_dland"`
	Bgen_lat   string `json:"bgen_lat"`
	Bgen_lon   string `json:"bgen_lon"`
	Bgen_dland string `json:"bgen_dland"`
	Gen_dist   string `json:"gen_dist"`
	Gen_tdiff  string `json:"gen_tdiff"`
	Init_tdiff string `json:"init_tdiff"`
	Dev_cat    string `json:"dev_cat"`
	Ops_cat    string `json:"ops_cat"`
}

type STAT_SEEPS struct {
	Total     int     `json:"total"`
	Odfl      float64 `json:"odfl"`
	Odfh      float64 `json:"odfh"`
	Olfd      float64 `json:"olfd"`
	Olfh      float64 `json:"olfh"`
	Ohfd      float64 `json:"ohfd"`
	Ohfl      float64 `json:"ohfl"`
	Pf1       float64 `json:"pf1"`
	Pf2       float64 `json:"pf2"`
	Pf3       float64 `json:"pf3"`
	Pv1       float64 `json:"pv1"`
	Pv2       float64 `json:"pv2"`
	Pv3       float64 `json:"pv3"`
	Mean_fcst float64 `json:"mean_fcst"`
	Mean_obs  float64 `json:"mean_obs"`
	Seeps     float64 `json:"seeps"`
}

type MODE_CTS struct {
	Field string  `json:"field"`
	Total int     `json:"total"`
	Fy_oy float64 `json:"fy_oy"`
	Fy_on float64 `json:"fy_on"`
	Fn_oy float64 `json:"fn_oy"`
	Fn_on float64 `json:"fn_on"`
	Baser float64 `json:"baser"`
	Fmean string  `json:"fmean"`
	Acc   string  `json:"acc"`
	Fbias float64 `json:"fbias"`
	Pody  string  `json:"pody"`
	Podn  string  `json:"podn"`
	Pofd  string  `json:"pofd"`
	Far   string  `json:"far"`
	Csi   string  `json:"csi"`
	Gss   string  `json:"gss"`
	Hk    string  `json:"hk"`
	Hss   string  `json:"hss"`
	Odds  string  `json:"odds"`
	Lodds string  `json:"lodds"`
	Orss  string  `json:"orss"`
	Eds   string  `json:"eds"`
	Seds  string  `json:"seds"`
	Edi   string  `json:"edi"`
	Sedi  string  `json:"sedi"`
	Bagss string  `json:"bagss"`
}

type STAT_ECLV struct {
	Total       int     `json:"total"`
	Baser       float64 `json:"baser"`
	Value_baser string  `json:"value_baser"`
	N_pts       string  `json:"n_pts"`
	Cl_i        string  `json:"cl_i"`
	Value_i     string  `json:"value_i"`
}

type STAT_VAL1L2 struct {
	Total        int     `json:"total"`
	Ufabar       float64 `json:"ufabar"`
	Vfabar       float64 `json:"vfabar"`
	Uoabar       float64 `json:"uoabar"`
	Voabar       float64 `json:"voabar"`
	Uvfoabar     float64 `json:"uvfoabar"`
	Uvffabar     float64 `json:"uvffabar"`
	Uvooabar     float64 `json:"uvooabar"`
	Fa_speed_bar float64 `json:"fa_speed_bar"`
	Oa_speed_bar float64 `json:"oa_speed_bar"`
	Total_dir    int     `json:"total_dir"`
	Dira_me      float64 `json:"dira_me"`
	Dira_mae     float64 `json:"dira_mae"`
	Dira_mse     float64 `json:"dira_mse"`
}

type STAT_NBRCTS struct {
	Total     int     `json:"total"`
	Baser     float64 `json:"baser"`
	Baser_ncl string  `json:"baser_ncl"`
	Baser_ncu string  `json:"baser_ncu"`
	Baser_bcl string  `json:"baser_bcl"`
	Baser_bcu string  `json:"baser_bcu"`
	Fmean     string  `json:"fmean"`
	Fmean_ncl string  `json:"fmean_ncl"`
	Fmean_ncu string  `json:"fmean_ncu"`
	Fmean_bcl string  `json:"fmean_bcl"`
	Fmean_bcu string  `json:"fmean_bcu"`
	Acc       string  `json:"acc"`
	Acc_ncl   string  `json:"acc_ncl"`
	Acc_ncu   string  `json:"acc_ncu"`
	Acc_bcl   string  `json:"acc_bcl"`
	Acc_bcu   string  `json:"acc_bcu"`
	Fbias     float64 `json:"fbias"`
	Fbias_bcl string  `json:"fbias_bcl"`
	Fbias_bcu string  `json:"fbias_bcu"`
	Pody      string  `json:"pody"`
	Pody_ncl  string  `json:"pody_ncl"`
	Pody_ncu  string  `json:"pody_ncu"`
	Pody_bcl  string  `json:"pody_bcl"`
	Pody_bcu  string  `json:"pody_bcu"`
	Podn      string  `json:"podn"`
	Podn_ncl  string  `json:"podn_ncl"`
	Podn_ncu  string  `json:"podn_ncu"`
	Podn_bcl  string  `json:"podn_bcl"`
	Podn_bcu  string  `json:"podn_bcu"`
	Pofd      string  `json:"pofd"`
	Pofd_ncl  string  `json:"pofd_ncl"`
	Pofd_ncu  string  `json:"pofd_ncu"`
	Pofd_bcl  string  `json:"pofd_bcl"`
	Pofd_bcu  string  `json:"pofd_bcu"`
	Far       string  `json:"far"`
	Far_ncl   string  `json:"far_ncl"`
	Far_ncu   string  `json:"far_ncu"`
	Far_bcl   string  `json:"far_bcl"`
	Far_bcu   string  `json:"far_bcu"`
	Csi       string  `json:"csi"`
	Csi_ncl   string  `json:"csi_ncl"`
	Csi_ncu   string  `json:"csi_ncu"`
	Csi_bcl   string  `json:"csi_bcl"`
	Csi_bcu   string  `json:"csi_bcu"`
	Gss       string  `json:"gss"`
	Gss_bcl   string  `json:"gss_bcl"`
	Gss_bcu   string  `json:"gss_bcu"`
	Hk        string  `json:"hk"`
	Hk_ncl    string  `json:"hk_ncl"`
	Hk_ncu    string  `json:"hk_ncu"`
	Hk_bcl    string  `json:"hk_bcl"`
	Hk_bcu    string  `json:"hk_bcu"`
	Hss       string  `json:"hss"`
	Hss_bcl   string  `json:"hss_bcl"`
	Hss_bcu   string  `json:"hss_bcu"`
	Odds      string  `json:"odds"`
	Odds_ncl  string  `json:"odds_ncl"`
	Odds_ncu  string  `json:"odds_ncu"`
	Odds_bcl  string  `json:"odds_bcl"`
	Odds_bcu  string  `json:"odds_bcu"`
	Lodds     string  `json:"lodds"`
	Lodds_ncl string  `json:"lodds_ncl"`
	Lodds_ncu string  `json:"lodds_ncu"`
	Lodds_bcl string  `json:"lodds_bcl"`
	Lodds_bcu string  `json:"lodds_bcu"`
	Orss      string  `json:"orss"`
	Orss_ncl  string  `json:"orss_ncl"`
	Orss_ncu  string  `json:"orss_ncu"`
	Orss_bcl  string  `json:"orss_bcl"`
	Orss_bcu  string  `json:"orss_bcu"`
	Eds       string  `json:"eds"`
	Eds_ncl   string  `json:"eds_ncl"`
	Eds_ncu   string  `json:"eds_ncu"`
	Eds_bcl   string  `json:"eds_bcl"`
	Eds_bcu   string  `json:"eds_bcu"`
	Seds      string  `json:"seds"`
	Seds_ncl  string  `json:"seds_ncl"`
	Seds_ncu  string  `json:"seds_ncu"`
	Seds_bcl  string  `json:"seds_bcl"`
	Seds_bcu  string  `json:"seds_bcu"`
	Edi       string  `json:"edi"`
	Edi_ncl   string  `json:"edi_ncl"`
	Edi_ncu   string  `json:"edi_ncu"`
	Edi_bcl   string  `json:"edi_bcl"`
	Edi_bcu   string  `json:"edi_bcu"`
	Sedi      string  `json:"sedi"`
	Sedi_ncl  string  `json:"sedi_ncl"`
	Sedi_ncu  string  `json:"sedi_ncu"`
	Sedi_bcl  string  `json:"sedi_bcl"`
	Sedi_bcu  string  `json:"sedi_bcu"`
	Bagss     string  `json:"bagss"`
	Bagss_bcl string  `json:"bagss_bcl"`
	Bagss_bcu string  `json:"bagss_bcu"`
}

// doc struct definitions
type STAT_ISC_doc struct {
	VxMetadata      `json:"vx_metadata"`
	STAT_ISC_header `json:"met_header"`
	Data            map[string]STAT_ISC `json:"data"`
}
type STAT_ECLV_doc struct {
	VxMetadata       `json:"vx_metadata"`
	STAT_ECLV_header `json:"met_header"`
	Data             map[string]STAT_ECLV `json:"data"`
}
type STAT_RHIST_doc struct {
	VxMetadata        `json:"vx_metadata"`
	STAT_RHIST_header `json:"met_header"`
	Data              map[string]STAT_RHIST `json:"data"`
}
type STAT_SL1L2_doc struct {
	VxMetadata        `json:"vx_metadata"`
	STAT_SL1L2_header `json:"met_header"`
	Data              map[string]STAT_SL1L2 `json:"data"`
}
type STAT_GENMPR_doc struct {
	VxMetadata         `json:"vx_metadata"`
	STAT_GENMPR_header `json:"met_header"`
	Data               map[string]STAT_GENMPR `json:"data"`
}
type TCST_PROBRIRW_doc struct {
	VxMetadata           `json:"vx_metadata"`
	TCST_PROBRIRW_header `json:"met_header"`
	Data                 map[string]TCST_PROBRIRW `json:"data"`
}
type STAT_CNT_doc struct {
	VxMetadata      `json:"vx_metadata"`
	STAT_CNT_header `json:"met_header"`
	Data            map[string]STAT_CNT `json:"data"`
}
type STAT_NBRCTS_doc struct {
	VxMetadata         `json:"vx_metadata"`
	STAT_NBRCTS_header `json:"met_header"`
	Data               map[string]STAT_NBRCTS `json:"data"`
}
type STAT_PJC_doc struct {
	VxMetadata      `json:"vx_metadata"`
	STAT_PJC_header `json:"met_header"`
	Data            map[string]STAT_PJC `json:"data"`
}
type STAT_SAL1L2_doc struct {
	VxMetadata         `json:"vx_metadata"`
	STAT_SAL1L2_header `json:"met_header"`
	Data               map[string]STAT_SAL1L2 `json:"data"`
}
type MODE_OBJ_doc struct {
	VxMetadata      `json:"vx_metadata"`
	MODE_OBJ_header `json:"met_header"`
	Data            map[string]MODE_OBJ `json:"data"`
}
type STAT_CTS_doc struct {
	VxMetadata      `json:"vx_metadata"`
	STAT_CTS_header `json:"met_header"`
	Data            map[string]STAT_CTS `json:"data"`
}
type STAT_NBRCTC_doc struct {
	VxMetadata         `json:"vx_metadata"`
	STAT_NBRCTC_header `json:"met_header"`
	Data               map[string]STAT_NBRCTC `json:"data"`
}
type STAT_PSTD_doc struct {
	VxMetadata       `json:"vx_metadata"`
	STAT_PSTD_header `json:"met_header"`
	Data             map[string]STAT_PSTD `json:"data"`
}
type STAT_RPS_doc struct {
	VxMetadata      `json:"vx_metadata"`
	STAT_RPS_header `json:"met_header"`
	Data            map[string]STAT_RPS `json:"data"`
}
type STAT_PHIST_doc struct {
	VxMetadata        `json:"vx_metadata"`
	STAT_PHIST_header `json:"met_header"`
	Data              map[string]STAT_PHIST `json:"data"`
}
type STAT_VL1L2_doc struct {
	VxMetadata        `json:"vx_metadata"`
	STAT_VL1L2_header `json:"met_header"`
	Data              map[string]STAT_VL1L2 `json:"Data"`
}
type STAT_VCNT_doc struct {
	VxMetadata       `json:"vx_metadata"`
	STAT_VCNT_header `json:"met_header"`
	Data             map[string]STAT_VCNT `json:"data"`
}
type STAT_FHO_doc struct {
	VxMetadata      `json:"vx_metadata"`
	STAT_FHO_header `json:"met_header"`
	Data            map[string]STAT_FHO `json:"data"`
}
type STAT_DMAP_doc struct {
	VxMetadata       `json:"vx_metadata"`
	STAT_DMAP_header `json:"met_header"`
	Data             map[string]STAT_DMAP `json:"data"`
}
type STAT_ORANK_doc struct {
	VxMetadata        `json:"vx_metadata"`
	STAT_ORANK_header `json:"met_header"`
	Data              map[string]STAT_ORANK `json:"data"`
}
type STAT_RELP_doc struct {
	VxMetadata       `json:"vx_metadata"`
	STAT_RELP_header `json:"met_header"`
	Data             map[string]STAT_RELP `json:"data"`
}
type STAT_SSVAR_doc struct {
	VxMetadata        `json:"vx_metadata"`
	STAT_SSVAR_header `json:"met_header"`
	Data              map[string]STAT_SSVAR `json:"data"`
}
type STAT_SSIDX_doc struct {
	VxMetadata        `json:"vx_metadata"`
	STAT_SSIDX_header `json:"met_header"`
	Data              map[string]STAT_SSIDX `json:"data"`
}
type STAT_SEEPS_MPR_doc struct {
	VxMetadata            `json:"vx_metadata"`
	STAT_SEEPS_MPR_header `json:"met_header"`
	Data                  map[string]STAT_SEEPS_MPR `json:"data"`
}
type STAT_GRAD_doc struct {
	VxMetadata       `json:"vx_metadata"`
	STAT_GRAD_header `json:"met_header"`
	Data             map[string]STAT_GRAD `json:"data"`
}
type STAT_PCT_doc struct {
	VxMetadata      `json:"vx_metadata"`
	STAT_PCT_header `json:"met_header"`
	Data            map[string]STAT_PCT `json:"data"`
}
type MODE_CTS_doc struct {
	VxMetadata      `json:"vx_metadata"`
	MODE_CTS_header `json:"met_header"`
	Data            map[string]MODE_CTS `json:"data"`
}
type STAT_MCTC_doc struct {
	VxMetadata       `json:"vx_metadata"`
	STAT_MCTC_header `json:"met_header"`
	Data             map[string]STAT_MCTC `json:"data"`
}
type STAT_MPR_doc struct {
	VxMetadata      `json:"vx_metadata"`
	STAT_MPR_header `json:"met_header"`
	Data            map[string]STAT_MPR `json:"data"`
}
type STAT_PRC_doc struct {
	VxMetadata      `json:"vx_metadata"`
	STAT_PRC_header `json:"met_header"`
	Data            map[string]STAT_PRC `json:"data"`
}
type TCST_TCMPR_doc struct {
	VxMetadata        `json:"vx_metadata"`
	TCST_TCMPR_header `json:"met_header"`
	Data              map[string]TCST_TCMPR `json:"data"`
}
type STAT_MCTS_doc struct {
	VxMetadata       `json:"vx_metadata"`
	STAT_MCTS_header `json:"met_header"`
	Data             map[string]STAT_MCTS `json:"data"`
}
type STAT_SEEPS_doc struct {
	VxMetadata        `json:"vx_metadata"`
	STAT_SEEPS_header `json:"met_header"`
	Data              map[string]STAT_SEEPS `json:"data"`
}
type STAT_ECNT_doc struct {
	VxMetadata       `json:"vx_metadata"`
	STAT_ECNT_header `json:"met_header"`
	Data             map[string]STAT_ECNT `json:"data"`
}
type TCST_TCDIAG_doc struct {
	VxMetadata         `json:"vx_metadata"`
	TCST_TCDIAG_header `json:"met_header"`
	Data               map[string]TCST_TCDIAG `json:"data"`
}
type STAT_CTC_doc struct {
	VxMetadata      `json:"vx_metadata"`
	STAT_CTC_header `json:"met_header"`
	Data            map[string]STAT_CTC `json:"data"`
}
type STAT_NBRCNT_doc struct {
	VxMetadata         `json:"vx_metadata"`
	STAT_NBRCNT_header `json:"met_header"`
	Data               map[string]STAT_NBRCNT `json:"data"`
}

// type STAT_VAL1L2_doc struct {
// 	VxMetadata         `json:"vx_metadata"`
// 	STAT_VAL1L2_header `json:"met_header"`
// 	Data               map[string]STAT_VAL1L2 `json:"Data"`
// }

func (s *STAT_VAL1L2) fillHeader(metaData VxMetadata, fields []string, doc *map[string]interface{}) {
	// fill the metadata
	metaData.fillMetadata(metaData, doc)
	// fill the met fields
	(*doc)["ID"] = metaData.ID
	(*doc)["Subset"] = metaData.Subset
	(*doc)["Type"] = metaData.Type
	(*doc)["SubType"] = metaData.SubType

	(*doc)["Version"] = fields[0]
	(*doc)["Model"] = fields[1]
	(*doc)["Desc"] = fields[2]
	(*doc)["Fcst_valid_beg"] = fields[3]
	(*doc)["Fcst_valid_end"] = fields[4]
	(*doc)["Obs_lead"] = fields[5]
	(*doc)["Obs_valid_beg"] = fields[6]
	(*doc)["Obs_valid_end"] = fields[7]
	(*doc)["Fcst_var"] = fields[8]
	(*doc)["Fcst_units"] = fields[9]
	(*doc)["Fcst_lev"] = fields[10]
	(*doc)["Obs_var"] = fields[11]
	(*doc)["Obs_units"] = fields[12]
	(*doc)["Obs_lev"] = fields[13]
	(*doc)["Obtype"] = fields[14]
	(*doc)["Vx_mask"] = fields[15]
	(*doc)["Interp_mthd"] = fields[16]
	(*doc)["Interp_pnts"] = fields[17]
	(*doc)["Fcst_thresh"] = fields[18]
	(*doc)["Obs_thresh"] = fields[19]
	(*doc)["Cov_thresh"] = fields[20]
	(*doc)["Alpha"] = fields[21]
	(*doc)["Line_type"] = fields[22]
}

func (s *STAT_VAL1L2) fillStructure(fields []string) {
	s.Total, _ = strconv.Atoi(fields[0])
	s.Ufabar, _ = strconv.ParseFloat(fields[1], 64)
	s.Vfabar, _ = strconv.ParseFloat(fields[2], 64)
	s.Uoabar, _ = strconv.ParseFloat(fields[3], 64)
	s.Voabar, _ = strconv.ParseFloat(fields[4], 64)
	s.Uvfoabar, _ = strconv.ParseFloat(fields[5], 64)
	s.Uvffabar, _ = strconv.ParseFloat(fields[6], 64)
	s.Uvooabar, _ = strconv.ParseFloat(fields[7], 64)
	s.Fa_speed_bar, _ = strconv.ParseFloat(fields[8], 64)
	s.Oa_speed_bar, _ = strconv.ParseFloat(fields[9], 64)
	s.Total_dir, _ = strconv.Atoi(fields[10])
	s.Dira_me, _ = strconv.ParseFloat(fields[11], 64)
	s.Dira_mae, _ = strconv.ParseFloat(fields[12], 64)
	s.Dira_mse, _ = strconv.ParseFloat(fields[13], 64)
}

// data key definitions
var DataKeyMap = map[string][]string{
	"STAT_CNT":       {"FCST_LEAD"},
	"STAT_CTC":       {"FCST_LEAD"},
	"STAT_CTS":       {"FCST_LEAD"},
	"STAT_FHO":       {"FCST_LEAD"},
	"STAT_ISC":       {"FCST_LEAD"},
	"STAT_MCTC":      {"FCST_LEAD"},
	"STAT_MCTS":      {"FCST_LEAD"},
	"STAT_MPR":       {"FCST_LEAD"},
	"STAT_SEEPS":     {"FCST_LEAD"},
	"STAT_SEEPS_MPR": {"FCST_LEAD"},
	"STAT_NBRCNT":    {"FCST_LEAD"},
	"STAT_NBRCTC":    {"FCST_LEAD"},
	"STAT_NBRCTS":    {"FCST_LEAD"},
	"STAT_GRAD":      {"FCST_LEAD"},
	"STAT_DMAP":      {"FCST_LEAD"},
	"STAT_ORANK":     {"FCST_LEAD"},
	"STAT_PCT":       {"FCST_LEAD"},
	"STAT_PJC":       {"FCST_LEAD"},
	"STAT_PRC":       {"FCST_LEAD"},
	"STAT_PSTD":      {"FCST_LEAD"},
	"STAT_ECLV":      {"FCST_LEAD"},
	"STAT_ECNT":      {"FCST_LEAD"},
	"STAT_RPS":       {"FCST_LEAD"},
	"STAT_RHIST":     {"FCST_LEAD"},
	"STAT_PHIST":     {"FCST_LEAD"},
	"STAT_RELP":      {"FCST_LEAD"},
	"STAT_SAL1L2":    {"FCST_LEAD"},
	"STAT_SL1L2":     {"FCST_LEAD"},
	"STAT_SSVAR":     {"FCST_LEAD"},
	"STAT_VAL1L2":    {"FCST_LEAD"},
	"STAT_VL1L2":     {"FCST_LEAD"},
	"STAT_VCNT":      {"FCST_LEAD"},
	"STAT_GENMPR":    {"FCST_LEAD"},
	"STAT_SSIDX":     {"FCST_LEAD"},
	"MODE_OBJ":       {"FCST_LEAD, FCST_LEV"},
	"MODE_CTS":       {"FCST_LEAD, FCST_LEV"},
	"TCST_TCMPR":     {"FCST_LEAD"},
	"TCST_TCDIAG":    {"FCST_LEAD"},
	"TCST_PROBRIRW":  {"FCST_LEAD"},
}

/*
Returns a document interface for a given line type, creating it if necessary
A fileLineType looks loike fileType+"_"+lineType e.g. "stat_val1l2"
*/
func GetDocForId(fileLineType string, metaData VxMetadata, headerData []string, dataData []string, dataKey string) (interface{}, error) {
	switch fileLineType {
	case "STAT_VAL1L2":
		{
			doc := make(map[string]interface{})
			//lineDoc := STAT_VAL1L2_doc{VxMetadata: metaData, STAT_VAL1L2_header: STAT_VAL1L2_header{}, Data: make(map[string]STAT_VAL1L2)}
			dataDoc := STAT_VAL1L2{}
			dataDoc.fillHeader(metaData, headerData, &doc)
			err := AddDataElement(dataKey, fileLineType, dataData, &doc)
			return doc, err
		}
	default:
		return nil, errors.New("Unknonwn file_line type:" + fileLineType)
	}
}

/*
Adds a data element to an existing document Data map based on the line type.
*/
func AddDataElement(dataKey string, fileLineType string, data []string, doc *map[string]interface{}) error {
	switch fileLineType {
	case "STAT_VAL1L2":
		{
			elem := STAT_VAL1L2{}
			elem.fillStructure(data)
			if exists := (*doc)["data"]; exists == nil {
				(*doc)["data"] = make(map[string]STAT_VAL1L2)
			}
			if val, ok := (*doc)["data"].(map[string]STAT_VAL1L2); ok {
				val[dataKey] = elem
				(*doc)["data"] = val
			}
		}
	default:
		return errors.New("Unknonwn file_line type:" + fileLineType)
	}
	return nil
}
