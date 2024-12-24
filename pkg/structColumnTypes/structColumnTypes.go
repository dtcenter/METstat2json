package structColumnTypes

import
(
"strconv"
"errors"
)
//vxMetadata struct definition
type VxMetadata struct {
    ID      string `json:"id"`
    Subset  string `json:"subset"`
    Type    string `json:"type"`
    SubType string `json:"subtype"`
}


//Header struct definitions
type STAT_GENMPR_header struct {
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

type STAT_NBRCTC_header struct {
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

type STAT_NBRCTS_header struct {
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

type STAT_GRAD_header struct {
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

type STAT_PHIST_header struct {
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

type MTD_2DSINGLE_header struct {
    Version    string `json:"version"`
    Model      string `json:"model"`
    Desc       string `json:"desc"`
    Fcst_valid string `json:"fcst_valid"`
    Obs_lead   string `json:"obs_lead"`
    Obs_valid  string `json:"obs_valid"`
    T_delta    string `json:"t_delta"`
    Fcst_t_beg string `json:"fcst_t_beg"`
    Fcst_t_end string `json:"fcst_t_end"`
    Fcst_rad   string `json:"fcst_rad"`
    Fcst_thr   string `json:"fcst_thr"`
    Obs_t_beg  string `json:"obs_t_beg"`
    Obs_t_end  string `json:"obs_t_end"`
    Obs_rad    string `json:"obs_rad"`
    Obs_thr    string `json:"obs_thr"`
    Fcst_var   string `json:"fcst_var"`
    Fcst_units string `json:"fcst_units"`
    Obs_var    string `json:"obs_var"`
    Obs_units  string `json:"obs_units"`
    Obs_lev    string `json:"obs_lev"`
}

type STAT_CNT_header struct {
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

type STAT_PCT_header struct {
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

type STAT_ECLV_header struct {
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

type STAT_RHIST_header struct {
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

type STAT_SAL1L2_header struct {
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

type MODE_CTS_header struct {
    Version    string `json:"version"`
    Model      string `json:"model"`
    N_valid    string `json:"n_valid"`
    Grid_res   string `json:"grid_res"`
    Desc       string `json:"desc"`
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
    Obs_var    string `json:"obs_var"`
    Obs_units  string `json:"obs_units"`
    Obs_lev    string `json:"obs_lev"`
}

type STAT_CTC_header struct {
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

type STAT_RELP_header struct {
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

type MTD_3DSINGLE_header struct {
    Version    string `json:"version"`
    Model      string `json:"model"`
    Desc       string `json:"desc"`
    Fcst_valid string `json:"fcst_valid"`
    Obs_lead   string `json:"obs_lead"`
    Obs_valid  string `json:"obs_valid"`
    T_delta    string `json:"t_delta"`
    Fcst_t_beg string `json:"fcst_t_beg"`
    Fcst_t_end string `json:"fcst_t_end"`
    Fcst_rad   string `json:"fcst_rad"`
    Fcst_thr   string `json:"fcst_thr"`
    Obs_t_beg  string `json:"obs_t_beg"`
    Obs_t_end  string `json:"obs_t_end"`
    Obs_rad    string `json:"obs_rad"`
    Obs_thr    string `json:"obs_thr"`
    Fcst_var   string `json:"fcst_var"`
    Fcst_units string `json:"fcst_units"`
    Obs_var    string `json:"obs_var"`
    Obs_units  string `json:"obs_units"`
    Obs_lev    string `json:"obs_lev"`
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

type STAT_SSIDX_header struct {
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

type MODE_OBJ_header struct {
    Version    string `json:"version"`
    Model      string `json:"model"`
    N_valid    string `json:"n_valid"`
    Grid_res   string `json:"grid_res"`
    Desc       string `json:"desc"`
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
    Obs_var    string `json:"obs_var"`
    Obs_units  string `json:"obs_units"`
    Obs_lev    string `json:"obs_lev"`
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

type STAT_MCTS_header struct {
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

type STAT_ORANK_header struct {
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

type STAT_VL1L2_header struct {
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

type STAT_VCNT_header struct {
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

type MTD_3DPAIR_header struct {
    Version    string `json:"version"`
    Model      string `json:"model"`
    Desc       string `json:"desc"`
    Fcst_valid string `json:"fcst_valid"`
    Obs_lead   string `json:"obs_lead"`
    Obs_valid  string `json:"obs_valid"`
    T_delta    string `json:"t_delta"`
    Fcst_t_beg string `json:"fcst_t_beg"`
    Fcst_t_end string `json:"fcst_t_end"`
    Fcst_rad   string `json:"fcst_rad"`
    Fcst_thr   string `json:"fcst_thr"`
    Obs_t_beg  string `json:"obs_t_beg"`
    Obs_t_end  string `json:"obs_t_end"`
    Obs_rad    string `json:"obs_rad"`
    Obs_thr    string `json:"obs_thr"`
    Fcst_var   string `json:"fcst_var"`
    Fcst_units string `json:"fcst_units"`
    Obs_var    string `json:"obs_var"`
    Obs_units  string `json:"obs_units"`
    Obs_lev    string `json:"obs_lev"`
}

type STAT_CTS_header struct {
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

type STAT_MCTC_header struct {
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

type STAT_PJC_header struct {
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

type STAT_PRC_header struct {
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

type STAT_PSTD_header struct {
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

type STAT_ECNT_header struct {
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

type STAT_SL1L2_header struct {
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

type STAT_ISC_header struct {
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

type STAT_MPR_header struct {
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

type STAT_SEEPS_MPR_header struct {
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

type STAT_RPS_header struct {
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

type STAT_SSVAR_header struct {
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


//fillHeader functions
func (s *STAT_CTC) fill_STAT_CTC_Header(metaData VxMetadata, fields []string, doc *map[string]interface{}) {
// fill the met fields
(*doc)["ID"] = metaData.ID
(*doc)["Subset"] = metaData.Subset
(*doc)["Type"] = metaData.Type
(*doc)["SubType"] = metaData.SubType
(*doc)["version"] = fields[0]
(*doc)["model"] = fields[1]
(*doc)["desc"] = fields[2]
(*doc)["fcst_valid_beg"] = fields[4]
(*doc)["fcst_valid_end"] = fields[5]
(*doc)["obs_lead"] = fields[6]
(*doc)["obs_valid_beg"] = fields[7]
(*doc)["obs_valid_end"] = fields[8]
(*doc)["fcst_var"] = fields[9]
(*doc)["fcst_units"] = fields[10]
(*doc)["fcst_lev"] = fields[11]
(*doc)["obs_var"] = fields[12]
(*doc)["obs_units"] = fields[13]
(*doc)["obs_lev"] = fields[14]
(*doc)["obtype"] = fields[15]
(*doc)["vx_mask"] = fields[16]
(*doc)["interp_mthd"] = fields[17]
(*doc)["interp_pnts"] = fields[18]
(*doc)["fcst_thresh"] = fields[19]
(*doc)["obs_thresh"] = fields[20]
(*doc)["cov_thresh"] = fields[21]
(*doc)["alpha"] = fields[22]
(*doc)["line_type"] = fields[23]
}

func (s *STAT_SEEPS) fill_STAT_SEEPS_Header(metaData VxMetadata, fields []string, doc *map[string]interface{}) {
// fill the met fields
(*doc)["ID"] = metaData.ID
(*doc)["Subset"] = metaData.Subset
(*doc)["Type"] = metaData.Type
(*doc)["SubType"] = metaData.SubType
(*doc)["version"] = fields[0]
(*doc)["model"] = fields[1]
(*doc)["desc"] = fields[2]
(*doc)["fcst_valid_beg"] = fields[4]
(*doc)["fcst_valid_end"] = fields[5]
(*doc)["obs_lead"] = fields[6]
(*doc)["obs_valid_beg"] = fields[7]
(*doc)["obs_valid_end"] = fields[8]
(*doc)["fcst_var"] = fields[9]
(*doc)["fcst_units"] = fields[10]
(*doc)["fcst_lev"] = fields[11]
(*doc)["obs_var"] = fields[12]
(*doc)["obs_units"] = fields[13]
(*doc)["obs_lev"] = fields[14]
(*doc)["obtype"] = fields[15]
(*doc)["vx_mask"] = fields[16]
(*doc)["interp_mthd"] = fields[17]
(*doc)["interp_pnts"] = fields[18]
(*doc)["fcst_thresh"] = fields[19]
(*doc)["obs_thresh"] = fields[20]
(*doc)["cov_thresh"] = fields[21]
(*doc)["alpha"] = fields[22]
(*doc)["line_type"] = fields[23]
}

func (s *STAT_NBRCTC) fill_STAT_NBRCTC_Header(metaData VxMetadata, fields []string, doc *map[string]interface{}) {
// fill the met fields
(*doc)["ID"] = metaData.ID
(*doc)["Subset"] = metaData.Subset
(*doc)["Type"] = metaData.Type
(*doc)["SubType"] = metaData.SubType
(*doc)["version"] = fields[0]
(*doc)["model"] = fields[1]
(*doc)["desc"] = fields[2]
(*doc)["fcst_valid_beg"] = fields[4]
(*doc)["fcst_valid_end"] = fields[5]
(*doc)["obs_lead"] = fields[6]
(*doc)["obs_valid_beg"] = fields[7]
(*doc)["obs_valid_end"] = fields[8]
(*doc)["fcst_var"] = fields[9]
(*doc)["fcst_units"] = fields[10]
(*doc)["fcst_lev"] = fields[11]
(*doc)["obs_var"] = fields[12]
(*doc)["obs_units"] = fields[13]
(*doc)["obs_lev"] = fields[14]
(*doc)["obtype"] = fields[15]
(*doc)["vx_mask"] = fields[16]
(*doc)["interp_mthd"] = fields[17]
(*doc)["interp_pnts"] = fields[18]
(*doc)["fcst_thresh"] = fields[19]
(*doc)["obs_thresh"] = fields[20]
(*doc)["cov_thresh"] = fields[21]
(*doc)["alpha"] = fields[22]
(*doc)["line_type"] = fields[23]
}

func (s *STAT_PJC) fill_STAT_PJC_Header(metaData VxMetadata, fields []string, doc *map[string]interface{}) {
// fill the met fields
(*doc)["ID"] = metaData.ID
(*doc)["Subset"] = metaData.Subset
(*doc)["Type"] = metaData.Type
(*doc)["SubType"] = metaData.SubType
(*doc)["version"] = fields[0]
(*doc)["model"] = fields[1]
(*doc)["desc"] = fields[2]
(*doc)["fcst_valid_beg"] = fields[4]
(*doc)["fcst_valid_end"] = fields[5]
(*doc)["obs_lead"] = fields[6]
(*doc)["obs_valid_beg"] = fields[7]
(*doc)["obs_valid_end"] = fields[8]
(*doc)["fcst_var"] = fields[9]
(*doc)["fcst_units"] = fields[10]
(*doc)["fcst_lev"] = fields[11]
(*doc)["obs_var"] = fields[12]
(*doc)["obs_units"] = fields[13]
(*doc)["obs_lev"] = fields[14]
(*doc)["obtype"] = fields[15]
(*doc)["vx_mask"] = fields[16]
(*doc)["interp_mthd"] = fields[17]
(*doc)["interp_pnts"] = fields[18]
(*doc)["fcst_thresh"] = fields[19]
(*doc)["obs_thresh"] = fields[20]
(*doc)["cov_thresh"] = fields[21]
(*doc)["alpha"] = fields[22]
(*doc)["line_type"] = fields[23]
}

func (s *STAT_PRC) fill_STAT_PRC_Header(metaData VxMetadata, fields []string, doc *map[string]interface{}) {
// fill the met fields
(*doc)["ID"] = metaData.ID
(*doc)["Subset"] = metaData.Subset
(*doc)["Type"] = metaData.Type
(*doc)["SubType"] = metaData.SubType
(*doc)["version"] = fields[0]
(*doc)["model"] = fields[1]
(*doc)["desc"] = fields[2]
(*doc)["fcst_valid_beg"] = fields[4]
(*doc)["fcst_valid_end"] = fields[5]
(*doc)["obs_lead"] = fields[6]
(*doc)["obs_valid_beg"] = fields[7]
(*doc)["obs_valid_end"] = fields[8]
(*doc)["fcst_var"] = fields[9]
(*doc)["fcst_units"] = fields[10]
(*doc)["fcst_lev"] = fields[11]
(*doc)["obs_var"] = fields[12]
(*doc)["obs_units"] = fields[13]
(*doc)["obs_lev"] = fields[14]
(*doc)["obtype"] = fields[15]
(*doc)["vx_mask"] = fields[16]
(*doc)["interp_mthd"] = fields[17]
(*doc)["interp_pnts"] = fields[18]
(*doc)["fcst_thresh"] = fields[19]
(*doc)["obs_thresh"] = fields[20]
(*doc)["cov_thresh"] = fields[21]
(*doc)["alpha"] = fields[22]
(*doc)["line_type"] = fields[23]
}

func (s *STAT_PSTD) fill_STAT_PSTD_Header(metaData VxMetadata, fields []string, doc *map[string]interface{}) {
// fill the met fields
(*doc)["ID"] = metaData.ID
(*doc)["Subset"] = metaData.Subset
(*doc)["Type"] = metaData.Type
(*doc)["SubType"] = metaData.SubType
(*doc)["version"] = fields[0]
(*doc)["model"] = fields[1]
(*doc)["desc"] = fields[2]
(*doc)["fcst_valid_beg"] = fields[4]
(*doc)["fcst_valid_end"] = fields[5]
(*doc)["obs_lead"] = fields[6]
(*doc)["obs_valid_beg"] = fields[7]
(*doc)["obs_valid_end"] = fields[8]
(*doc)["fcst_var"] = fields[9]
(*doc)["fcst_units"] = fields[10]
(*doc)["fcst_lev"] = fields[11]
(*doc)["obs_var"] = fields[12]
(*doc)["obs_units"] = fields[13]
(*doc)["obs_lev"] = fields[14]
(*doc)["obtype"] = fields[15]
(*doc)["vx_mask"] = fields[16]
(*doc)["interp_mthd"] = fields[17]
(*doc)["interp_pnts"] = fields[18]
(*doc)["fcst_thresh"] = fields[19]
(*doc)["obs_thresh"] = fields[20]
(*doc)["cov_thresh"] = fields[21]
(*doc)["alpha"] = fields[22]
(*doc)["line_type"] = fields[23]
}

func (s *STAT_SL1L2) fill_STAT_SL1L2_Header(metaData VxMetadata, fields []string, doc *map[string]interface{}) {
// fill the met fields
(*doc)["ID"] = metaData.ID
(*doc)["Subset"] = metaData.Subset
(*doc)["Type"] = metaData.Type
(*doc)["SubType"] = metaData.SubType
(*doc)["version"] = fields[0]
(*doc)["model"] = fields[1]
(*doc)["desc"] = fields[2]
(*doc)["fcst_valid_beg"] = fields[4]
(*doc)["fcst_valid_end"] = fields[5]
(*doc)["obs_lead"] = fields[6]
(*doc)["obs_valid_beg"] = fields[7]
(*doc)["obs_valid_end"] = fields[8]
(*doc)["fcst_var"] = fields[9]
(*doc)["fcst_units"] = fields[10]
(*doc)["fcst_lev"] = fields[11]
(*doc)["obs_var"] = fields[12]
(*doc)["obs_units"] = fields[13]
(*doc)["obs_lev"] = fields[14]
(*doc)["obtype"] = fields[15]
(*doc)["vx_mask"] = fields[16]
(*doc)["interp_mthd"] = fields[17]
(*doc)["interp_pnts"] = fields[18]
(*doc)["fcst_thresh"] = fields[19]
(*doc)["obs_thresh"] = fields[20]
(*doc)["cov_thresh"] = fields[21]
(*doc)["alpha"] = fields[22]
(*doc)["line_type"] = fields[23]
}

func (s *STAT_MCTC) fill_STAT_MCTC_Header(metaData VxMetadata, fields []string, doc *map[string]interface{}) {
// fill the met fields
(*doc)["ID"] = metaData.ID
(*doc)["Subset"] = metaData.Subset
(*doc)["Type"] = metaData.Type
(*doc)["SubType"] = metaData.SubType
(*doc)["version"] = fields[0]
(*doc)["model"] = fields[1]
(*doc)["desc"] = fields[2]
(*doc)["fcst_valid_beg"] = fields[4]
(*doc)["fcst_valid_end"] = fields[5]
(*doc)["obs_lead"] = fields[6]
(*doc)["obs_valid_beg"] = fields[7]
(*doc)["obs_valid_end"] = fields[8]
(*doc)["fcst_var"] = fields[9]
(*doc)["fcst_units"] = fields[10]
(*doc)["fcst_lev"] = fields[11]
(*doc)["obs_var"] = fields[12]
(*doc)["obs_units"] = fields[13]
(*doc)["obs_lev"] = fields[14]
(*doc)["obtype"] = fields[15]
(*doc)["vx_mask"] = fields[16]
(*doc)["interp_mthd"] = fields[17]
(*doc)["interp_pnts"] = fields[18]
(*doc)["fcst_thresh"] = fields[19]
(*doc)["obs_thresh"] = fields[20]
(*doc)["cov_thresh"] = fields[21]
(*doc)["alpha"] = fields[22]
(*doc)["line_type"] = fields[23]
}

func (s *STAT_MPR) fill_STAT_MPR_Header(metaData VxMetadata, fields []string, doc *map[string]interface{}) {
// fill the met fields
(*doc)["ID"] = metaData.ID
(*doc)["Subset"] = metaData.Subset
(*doc)["Type"] = metaData.Type
(*doc)["SubType"] = metaData.SubType
(*doc)["version"] = fields[0]
(*doc)["model"] = fields[1]
(*doc)["desc"] = fields[2]
(*doc)["fcst_valid_beg"] = fields[4]
(*doc)["fcst_valid_end"] = fields[5]
(*doc)["obs_lead"] = fields[6]
(*doc)["obs_valid_beg"] = fields[7]
(*doc)["obs_valid_end"] = fields[8]
(*doc)["fcst_var"] = fields[9]
(*doc)["fcst_units"] = fields[10]
(*doc)["fcst_lev"] = fields[11]
(*doc)["obs_var"] = fields[12]
(*doc)["obs_units"] = fields[13]
(*doc)["obs_lev"] = fields[14]
(*doc)["obtype"] = fields[15]
(*doc)["vx_mask"] = fields[16]
(*doc)["interp_mthd"] = fields[17]
(*doc)["interp_pnts"] = fields[18]
(*doc)["fcst_thresh"] = fields[19]
(*doc)["obs_thresh"] = fields[20]
(*doc)["cov_thresh"] = fields[21]
(*doc)["alpha"] = fields[22]
(*doc)["line_type"] = fields[23]
}

func (s *STAT_ORANK) fill_STAT_ORANK_Header(metaData VxMetadata, fields []string, doc *map[string]interface{}) {
// fill the met fields
(*doc)["ID"] = metaData.ID
(*doc)["Subset"] = metaData.Subset
(*doc)["Type"] = metaData.Type
(*doc)["SubType"] = metaData.SubType
(*doc)["version"] = fields[0]
(*doc)["model"] = fields[1]
(*doc)["desc"] = fields[2]
(*doc)["fcst_valid_beg"] = fields[4]
(*doc)["fcst_valid_end"] = fields[5]
(*doc)["obs_lead"] = fields[6]
(*doc)["obs_valid_beg"] = fields[7]
(*doc)["obs_valid_end"] = fields[8]
(*doc)["fcst_var"] = fields[9]
(*doc)["fcst_units"] = fields[10]
(*doc)["fcst_lev"] = fields[11]
(*doc)["obs_var"] = fields[12]
(*doc)["obs_units"] = fields[13]
(*doc)["obs_lev"] = fields[14]
(*doc)["obtype"] = fields[15]
(*doc)["vx_mask"] = fields[16]
(*doc)["interp_mthd"] = fields[17]
(*doc)["interp_pnts"] = fields[18]
(*doc)["fcst_thresh"] = fields[19]
(*doc)["obs_thresh"] = fields[20]
(*doc)["cov_thresh"] = fields[21]
(*doc)["alpha"] = fields[22]
(*doc)["line_type"] = fields[23]
}

func (s *STAT_ECNT) fill_STAT_ECNT_Header(metaData VxMetadata, fields []string, doc *map[string]interface{}) {
// fill the met fields
(*doc)["ID"] = metaData.ID
(*doc)["Subset"] = metaData.Subset
(*doc)["Type"] = metaData.Type
(*doc)["SubType"] = metaData.SubType
(*doc)["version"] = fields[0]
(*doc)["model"] = fields[1]
(*doc)["desc"] = fields[2]
(*doc)["fcst_valid_beg"] = fields[4]
(*doc)["fcst_valid_end"] = fields[5]
(*doc)["obs_lead"] = fields[6]
(*doc)["obs_valid_beg"] = fields[7]
(*doc)["obs_valid_end"] = fields[8]
(*doc)["fcst_var"] = fields[9]
(*doc)["fcst_units"] = fields[10]
(*doc)["fcst_lev"] = fields[11]
(*doc)["obs_var"] = fields[12]
(*doc)["obs_units"] = fields[13]
(*doc)["obs_lev"] = fields[14]
(*doc)["obtype"] = fields[15]
(*doc)["vx_mask"] = fields[16]
(*doc)["interp_mthd"] = fields[17]
(*doc)["interp_pnts"] = fields[18]
(*doc)["fcst_thresh"] = fields[19]
(*doc)["obs_thresh"] = fields[20]
(*doc)["cov_thresh"] = fields[21]
(*doc)["alpha"] = fields[22]
(*doc)["line_type"] = fields[23]
}

func (s *STAT_VL1L2) fill_STAT_VL1L2_Header(metaData VxMetadata, fields []string, doc *map[string]interface{}) {
// fill the met fields
(*doc)["ID"] = metaData.ID
(*doc)["Subset"] = metaData.Subset
(*doc)["Type"] = metaData.Type
(*doc)["SubType"] = metaData.SubType
(*doc)["version"] = fields[0]
(*doc)["model"] = fields[1]
(*doc)["desc"] = fields[2]
(*doc)["fcst_valid_beg"] = fields[4]
(*doc)["fcst_valid_end"] = fields[5]
(*doc)["obs_lead"] = fields[6]
(*doc)["obs_valid_beg"] = fields[7]
(*doc)["obs_valid_end"] = fields[8]
(*doc)["fcst_var"] = fields[9]
(*doc)["fcst_units"] = fields[10]
(*doc)["fcst_lev"] = fields[11]
(*doc)["obs_var"] = fields[12]
(*doc)["obs_units"] = fields[13]
(*doc)["obs_lev"] = fields[14]
(*doc)["obtype"] = fields[15]
(*doc)["vx_mask"] = fields[16]
(*doc)["interp_mthd"] = fields[17]
(*doc)["interp_pnts"] = fields[18]
(*doc)["fcst_thresh"] = fields[19]
(*doc)["obs_thresh"] = fields[20]
(*doc)["cov_thresh"] = fields[21]
(*doc)["alpha"] = fields[22]
(*doc)["line_type"] = fields[23]
}

func (s *STAT_SSIDX) fill_STAT_SSIDX_Header(metaData VxMetadata, fields []string, doc *map[string]interface{}) {
// fill the met fields
(*doc)["ID"] = metaData.ID
(*doc)["Subset"] = metaData.Subset
(*doc)["Type"] = metaData.Type
(*doc)["SubType"] = metaData.SubType
(*doc)["version"] = fields[0]
(*doc)["model"] = fields[1]
(*doc)["desc"] = fields[2]
(*doc)["fcst_valid_beg"] = fields[4]
(*doc)["fcst_valid_end"] = fields[5]
(*doc)["obs_lead"] = fields[6]
(*doc)["obs_valid_beg"] = fields[7]
(*doc)["obs_valid_end"] = fields[8]
(*doc)["fcst_var"] = fields[9]
(*doc)["fcst_units"] = fields[10]
(*doc)["fcst_lev"] = fields[11]
(*doc)["obs_var"] = fields[12]
(*doc)["obs_units"] = fields[13]
(*doc)["obs_lev"] = fields[14]
(*doc)["obtype"] = fields[15]
(*doc)["vx_mask"] = fields[16]
(*doc)["interp_mthd"] = fields[17]
(*doc)["interp_pnts"] = fields[18]
(*doc)["fcst_thresh"] = fields[19]
(*doc)["obs_thresh"] = fields[20]
(*doc)["cov_thresh"] = fields[21]
(*doc)["alpha"] = fields[22]
(*doc)["line_type"] = fields[23]
}

func (s *STAT_NBRCNT) fill_STAT_NBRCNT_Header(metaData VxMetadata, fields []string, doc *map[string]interface{}) {
// fill the met fields
(*doc)["ID"] = metaData.ID
(*doc)["Subset"] = metaData.Subset
(*doc)["Type"] = metaData.Type
(*doc)["SubType"] = metaData.SubType
(*doc)["version"] = fields[0]
(*doc)["model"] = fields[1]
(*doc)["desc"] = fields[2]
(*doc)["fcst_valid_beg"] = fields[4]
(*doc)["fcst_valid_end"] = fields[5]
(*doc)["obs_lead"] = fields[6]
(*doc)["obs_valid_beg"] = fields[7]
(*doc)["obs_valid_end"] = fields[8]
(*doc)["fcst_var"] = fields[9]
(*doc)["fcst_units"] = fields[10]
(*doc)["fcst_lev"] = fields[11]
(*doc)["obs_var"] = fields[12]
(*doc)["obs_units"] = fields[13]
(*doc)["obs_lev"] = fields[14]
(*doc)["obtype"] = fields[15]
(*doc)["vx_mask"] = fields[16]
(*doc)["interp_mthd"] = fields[17]
(*doc)["interp_pnts"] = fields[18]
(*doc)["fcst_thresh"] = fields[19]
(*doc)["obs_thresh"] = fields[20]
(*doc)["cov_thresh"] = fields[21]
(*doc)["alpha"] = fields[22]
(*doc)["line_type"] = fields[23]
}

func (s *TCST_TCMPR) fill_TCST_TCMPR_Header(metaData VxMetadata, fields []string, doc *map[string]interface{}) {
// fill the met fields
(*doc)["ID"] = metaData.ID
(*doc)["Subset"] = metaData.Subset
(*doc)["Type"] = metaData.Type
(*doc)["SubType"] = metaData.SubType
(*doc)["version"] = fields[0]
(*doc)["amodel"] = fields[1]
(*doc)["bmodel"] = fields[2]
(*doc)["desc"] = fields[3]
(*doc)["storm_id"] = fields[4]
(*doc)["basin"] = fields[5]
(*doc)["cyclone"] = fields[6]
(*doc)["storm_name"] = fields[7]
(*doc)["init"] = fields[8]
(*doc)["lead"] = fields[9]
(*doc)["valid"] = fields[10]
(*doc)["init_mask"] = fields[11]
(*doc)["valid_mask"] = fields[12]
(*doc)["line_type"] = fields[13]
}

func (s *TCST_PROBRIRW) fill_TCST_PROBRIRW_Header(metaData VxMetadata, fields []string, doc *map[string]interface{}) {
// fill the met fields
(*doc)["ID"] = metaData.ID
(*doc)["Subset"] = metaData.Subset
(*doc)["Type"] = metaData.Type
(*doc)["SubType"] = metaData.SubType
(*doc)["version"] = fields[0]
(*doc)["amodel"] = fields[1]
(*doc)["bmodel"] = fields[2]
(*doc)["desc"] = fields[3]
(*doc)["storm_id"] = fields[4]
(*doc)["basin"] = fields[5]
(*doc)["cyclone"] = fields[6]
(*doc)["storm_name"] = fields[7]
(*doc)["init"] = fields[8]
(*doc)["lead"] = fields[9]
(*doc)["valid"] = fields[10]
(*doc)["init_mask"] = fields[11]
(*doc)["valid_mask"] = fields[12]
(*doc)["line_type"] = fields[13]
}

func (s *STAT_CNT) fill_STAT_CNT_Header(metaData VxMetadata, fields []string, doc *map[string]interface{}) {
// fill the met fields
(*doc)["ID"] = metaData.ID
(*doc)["Subset"] = metaData.Subset
(*doc)["Type"] = metaData.Type
(*doc)["SubType"] = metaData.SubType
(*doc)["version"] = fields[0]
(*doc)["model"] = fields[1]
(*doc)["desc"] = fields[2]
(*doc)["fcst_valid_beg"] = fields[4]
(*doc)["fcst_valid_end"] = fields[5]
(*doc)["obs_lead"] = fields[6]
(*doc)["obs_valid_beg"] = fields[7]
(*doc)["obs_valid_end"] = fields[8]
(*doc)["fcst_var"] = fields[9]
(*doc)["fcst_units"] = fields[10]
(*doc)["fcst_lev"] = fields[11]
(*doc)["obs_var"] = fields[12]
(*doc)["obs_units"] = fields[13]
(*doc)["obs_lev"] = fields[14]
(*doc)["obtype"] = fields[15]
(*doc)["vx_mask"] = fields[16]
(*doc)["interp_mthd"] = fields[17]
(*doc)["interp_pnts"] = fields[18]
(*doc)["fcst_thresh"] = fields[19]
(*doc)["obs_thresh"] = fields[20]
(*doc)["cov_thresh"] = fields[21]
(*doc)["alpha"] = fields[22]
(*doc)["line_type"] = fields[23]
}

func (s *STAT_MCTS) fill_STAT_MCTS_Header(metaData VxMetadata, fields []string, doc *map[string]interface{}) {
// fill the met fields
(*doc)["ID"] = metaData.ID
(*doc)["Subset"] = metaData.Subset
(*doc)["Type"] = metaData.Type
(*doc)["SubType"] = metaData.SubType
(*doc)["version"] = fields[0]
(*doc)["model"] = fields[1]
(*doc)["desc"] = fields[2]
(*doc)["fcst_valid_beg"] = fields[4]
(*doc)["fcst_valid_end"] = fields[5]
(*doc)["obs_lead"] = fields[6]
(*doc)["obs_valid_beg"] = fields[7]
(*doc)["obs_valid_end"] = fields[8]
(*doc)["fcst_var"] = fields[9]
(*doc)["fcst_units"] = fields[10]
(*doc)["fcst_lev"] = fields[11]
(*doc)["obs_var"] = fields[12]
(*doc)["obs_units"] = fields[13]
(*doc)["obs_lev"] = fields[14]
(*doc)["obtype"] = fields[15]
(*doc)["vx_mask"] = fields[16]
(*doc)["interp_mthd"] = fields[17]
(*doc)["interp_pnts"] = fields[18]
(*doc)["fcst_thresh"] = fields[19]
(*doc)["obs_thresh"] = fields[20]
(*doc)["cov_thresh"] = fields[21]
(*doc)["alpha"] = fields[22]
(*doc)["line_type"] = fields[23]
}

func (s *STAT_PCT) fill_STAT_PCT_Header(metaData VxMetadata, fields []string, doc *map[string]interface{}) {
// fill the met fields
(*doc)["ID"] = metaData.ID
(*doc)["Subset"] = metaData.Subset
(*doc)["Type"] = metaData.Type
(*doc)["SubType"] = metaData.SubType
(*doc)["version"] = fields[0]
(*doc)["model"] = fields[1]
(*doc)["desc"] = fields[2]
(*doc)["fcst_valid_beg"] = fields[4]
(*doc)["fcst_valid_end"] = fields[5]
(*doc)["obs_lead"] = fields[6]
(*doc)["obs_valid_beg"] = fields[7]
(*doc)["obs_valid_end"] = fields[8]
(*doc)["fcst_var"] = fields[9]
(*doc)["fcst_units"] = fields[10]
(*doc)["fcst_lev"] = fields[11]
(*doc)["obs_var"] = fields[12]
(*doc)["obs_units"] = fields[13]
(*doc)["obs_lev"] = fields[14]
(*doc)["obtype"] = fields[15]
(*doc)["vx_mask"] = fields[16]
(*doc)["interp_mthd"] = fields[17]
(*doc)["interp_pnts"] = fields[18]
(*doc)["fcst_thresh"] = fields[19]
(*doc)["obs_thresh"] = fields[20]
(*doc)["cov_thresh"] = fields[21]
(*doc)["alpha"] = fields[22]
(*doc)["line_type"] = fields[23]
}

func (s *STAT_ECLV) fill_STAT_ECLV_Header(metaData VxMetadata, fields []string, doc *map[string]interface{}) {
// fill the met fields
(*doc)["ID"] = metaData.ID
(*doc)["Subset"] = metaData.Subset
(*doc)["Type"] = metaData.Type
(*doc)["SubType"] = metaData.SubType
(*doc)["version"] = fields[0]
(*doc)["model"] = fields[1]
(*doc)["desc"] = fields[2]
(*doc)["fcst_valid_beg"] = fields[4]
(*doc)["fcst_valid_end"] = fields[5]
(*doc)["obs_lead"] = fields[6]
(*doc)["obs_valid_beg"] = fields[7]
(*doc)["obs_valid_end"] = fields[8]
(*doc)["fcst_var"] = fields[9]
(*doc)["fcst_units"] = fields[10]
(*doc)["fcst_lev"] = fields[11]
(*doc)["obs_var"] = fields[12]
(*doc)["obs_units"] = fields[13]
(*doc)["obs_lev"] = fields[14]
(*doc)["obtype"] = fields[15]
(*doc)["vx_mask"] = fields[16]
(*doc)["interp_mthd"] = fields[17]
(*doc)["interp_pnts"] = fields[18]
(*doc)["fcst_thresh"] = fields[19]
(*doc)["obs_thresh"] = fields[20]
(*doc)["cov_thresh"] = fields[21]
(*doc)["alpha"] = fields[22]
(*doc)["line_type"] = fields[23]
}

func (s *STAT_RHIST) fill_STAT_RHIST_Header(metaData VxMetadata, fields []string, doc *map[string]interface{}) {
// fill the met fields
(*doc)["ID"] = metaData.ID
(*doc)["Subset"] = metaData.Subset
(*doc)["Type"] = metaData.Type
(*doc)["SubType"] = metaData.SubType
(*doc)["version"] = fields[0]
(*doc)["model"] = fields[1]
(*doc)["desc"] = fields[2]
(*doc)["fcst_valid_beg"] = fields[4]
(*doc)["fcst_valid_end"] = fields[5]
(*doc)["obs_lead"] = fields[6]
(*doc)["obs_valid_beg"] = fields[7]
(*doc)["obs_valid_end"] = fields[8]
(*doc)["fcst_var"] = fields[9]
(*doc)["fcst_units"] = fields[10]
(*doc)["fcst_lev"] = fields[11]
(*doc)["obs_var"] = fields[12]
(*doc)["obs_units"] = fields[13]
(*doc)["obs_lev"] = fields[14]
(*doc)["obtype"] = fields[15]
(*doc)["vx_mask"] = fields[16]
(*doc)["interp_mthd"] = fields[17]
(*doc)["interp_pnts"] = fields[18]
(*doc)["fcst_thresh"] = fields[19]
(*doc)["obs_thresh"] = fields[20]
(*doc)["cov_thresh"] = fields[21]
(*doc)["alpha"] = fields[22]
(*doc)["line_type"] = fields[23]
}

func (s *MTD_3DSINGLE) fill_MTD_3DSINGLE_Header(metaData VxMetadata, fields []string, doc *map[string]interface{}) {
// fill the met fields
(*doc)["ID"] = metaData.ID
(*doc)["Subset"] = metaData.Subset
(*doc)["Type"] = metaData.Type
(*doc)["SubType"] = metaData.SubType
(*doc)["version"] = fields[0]
(*doc)["model"] = fields[1]
(*doc)["desc"] = fields[2]
(*doc)["fcst_valid"] = fields[4]
(*doc)["obs_lead"] = fields[5]
(*doc)["obs_valid"] = fields[6]
(*doc)["t_delta"] = fields[7]
(*doc)["fcst_t_beg"] = fields[8]
(*doc)["fcst_t_end"] = fields[9]
(*doc)["fcst_rad"] = fields[10]
(*doc)["fcst_thr"] = fields[11]
(*doc)["obs_t_beg"] = fields[12]
(*doc)["obs_t_end"] = fields[13]
(*doc)["obs_rad"] = fields[14]
(*doc)["obs_thr"] = fields[15]
(*doc)["fcst_var"] = fields[16]
(*doc)["fcst_units"] = fields[17]
(*doc)["obs_var"] = fields[19]
(*doc)["obs_units"] = fields[20]
(*doc)["obs_lev"] = fields[21]
}

func (s *STAT_GRAD) fill_STAT_GRAD_Header(metaData VxMetadata, fields []string, doc *map[string]interface{}) {
// fill the met fields
(*doc)["ID"] = metaData.ID
(*doc)["Subset"] = metaData.Subset
(*doc)["Type"] = metaData.Type
(*doc)["SubType"] = metaData.SubType
(*doc)["version"] = fields[0]
(*doc)["model"] = fields[1]
(*doc)["desc"] = fields[2]
(*doc)["fcst_valid_beg"] = fields[4]
(*doc)["fcst_valid_end"] = fields[5]
(*doc)["obs_lead"] = fields[6]
(*doc)["obs_valid_beg"] = fields[7]
(*doc)["obs_valid_end"] = fields[8]
(*doc)["fcst_var"] = fields[9]
(*doc)["fcst_units"] = fields[10]
(*doc)["fcst_lev"] = fields[11]
(*doc)["obs_var"] = fields[12]
(*doc)["obs_units"] = fields[13]
(*doc)["obs_lev"] = fields[14]
(*doc)["obtype"] = fields[15]
(*doc)["vx_mask"] = fields[16]
(*doc)["interp_mthd"] = fields[17]
(*doc)["interp_pnts"] = fields[18]
(*doc)["fcst_thresh"] = fields[19]
(*doc)["obs_thresh"] = fields[20]
(*doc)["cov_thresh"] = fields[21]
(*doc)["alpha"] = fields[22]
(*doc)["line_type"] = fields[23]
}

func (s *STAT_GENMPR) fill_STAT_GENMPR_Header(metaData VxMetadata, fields []string, doc *map[string]interface{}) {
// fill the met fields
(*doc)["ID"] = metaData.ID
(*doc)["Subset"] = metaData.Subset
(*doc)["Type"] = metaData.Type
(*doc)["SubType"] = metaData.SubType
(*doc)["version"] = fields[0]
(*doc)["model"] = fields[1]
(*doc)["desc"] = fields[2]
(*doc)["fcst_valid_beg"] = fields[4]
(*doc)["fcst_valid_end"] = fields[5]
(*doc)["obs_lead"] = fields[6]
(*doc)["obs_valid_beg"] = fields[7]
(*doc)["obs_valid_end"] = fields[8]
(*doc)["fcst_var"] = fields[9]
(*doc)["fcst_units"] = fields[10]
(*doc)["fcst_lev"] = fields[11]
(*doc)["obs_var"] = fields[12]
(*doc)["obs_units"] = fields[13]
(*doc)["obs_lev"] = fields[14]
(*doc)["obtype"] = fields[15]
(*doc)["vx_mask"] = fields[16]
(*doc)["interp_mthd"] = fields[17]
(*doc)["interp_pnts"] = fields[18]
(*doc)["fcst_thresh"] = fields[19]
(*doc)["obs_thresh"] = fields[20]
(*doc)["cov_thresh"] = fields[21]
(*doc)["alpha"] = fields[22]
(*doc)["line_type"] = fields[23]
}

func (s *MODE_OBJ) fill_MODE_OBJ_Header(metaData VxMetadata, fields []string, doc *map[string]interface{}) {
// fill the met fields
(*doc)["ID"] = metaData.ID
(*doc)["Subset"] = metaData.Subset
(*doc)["Type"] = metaData.Type
(*doc)["SubType"] = metaData.SubType
(*doc)["version"] = fields[0]
(*doc)["model"] = fields[1]
(*doc)["n_valid"] = fields[2]
(*doc)["grid_res"] = fields[3]
(*doc)["desc"] = fields[4]
(*doc)["fcst_valid"] = fields[6]
(*doc)["fcst_accum"] = fields[7]
(*doc)["obs_lead"] = fields[8]
(*doc)["obs_valid"] = fields[9]
(*doc)["obs_accum"] = fields[10]
(*doc)["fcst_rad"] = fields[11]
(*doc)["fcst_thr"] = fields[12]
(*doc)["obs_rad"] = fields[13]
(*doc)["obs_thr"] = fields[14]
(*doc)["fcst_var"] = fields[15]
(*doc)["fcst_units"] = fields[16]
(*doc)["obs_var"] = fields[18]
(*doc)["obs_units"] = fields[19]
(*doc)["obs_lev"] = fields[20]
}

func (s *MTD_3DPAIR) fill_MTD_3DPAIR_Header(metaData VxMetadata, fields []string, doc *map[string]interface{}) {
// fill the met fields
(*doc)["ID"] = metaData.ID
(*doc)["Subset"] = metaData.Subset
(*doc)["Type"] = metaData.Type
(*doc)["SubType"] = metaData.SubType
(*doc)["version"] = fields[0]
(*doc)["model"] = fields[1]
(*doc)["desc"] = fields[2]
(*doc)["fcst_valid"] = fields[4]
(*doc)["obs_lead"] = fields[5]
(*doc)["obs_valid"] = fields[6]
(*doc)["t_delta"] = fields[7]
(*doc)["fcst_t_beg"] = fields[8]
(*doc)["fcst_t_end"] = fields[9]
(*doc)["fcst_rad"] = fields[10]
(*doc)["fcst_thr"] = fields[11]
(*doc)["obs_t_beg"] = fields[12]
(*doc)["obs_t_end"] = fields[13]
(*doc)["obs_rad"] = fields[14]
(*doc)["obs_thr"] = fields[15]
(*doc)["fcst_var"] = fields[16]
(*doc)["fcst_units"] = fields[17]
(*doc)["obs_var"] = fields[19]
(*doc)["obs_units"] = fields[20]
(*doc)["obs_lev"] = fields[21]
}

func (s *STAT_NBRCTS) fill_STAT_NBRCTS_Header(metaData VxMetadata, fields []string, doc *map[string]interface{}) {
// fill the met fields
(*doc)["ID"] = metaData.ID
(*doc)["Subset"] = metaData.Subset
(*doc)["Type"] = metaData.Type
(*doc)["SubType"] = metaData.SubType
(*doc)["version"] = fields[0]
(*doc)["model"] = fields[1]
(*doc)["desc"] = fields[2]
(*doc)["fcst_valid_beg"] = fields[4]
(*doc)["fcst_valid_end"] = fields[5]
(*doc)["obs_lead"] = fields[6]
(*doc)["obs_valid_beg"] = fields[7]
(*doc)["obs_valid_end"] = fields[8]
(*doc)["fcst_var"] = fields[9]
(*doc)["fcst_units"] = fields[10]
(*doc)["fcst_lev"] = fields[11]
(*doc)["obs_var"] = fields[12]
(*doc)["obs_units"] = fields[13]
(*doc)["obs_lev"] = fields[14]
(*doc)["obtype"] = fields[15]
(*doc)["vx_mask"] = fields[16]
(*doc)["interp_mthd"] = fields[17]
(*doc)["interp_pnts"] = fields[18]
(*doc)["fcst_thresh"] = fields[19]
(*doc)["obs_thresh"] = fields[20]
(*doc)["cov_thresh"] = fields[21]
(*doc)["alpha"] = fields[22]
(*doc)["line_type"] = fields[23]
}

func (s *STAT_SAL1L2) fill_STAT_SAL1L2_Header(metaData VxMetadata, fields []string, doc *map[string]interface{}) {
// fill the met fields
(*doc)["ID"] = metaData.ID
(*doc)["Subset"] = metaData.Subset
(*doc)["Type"] = metaData.Type
(*doc)["SubType"] = metaData.SubType
(*doc)["version"] = fields[0]
(*doc)["model"] = fields[1]
(*doc)["desc"] = fields[2]
(*doc)["fcst_valid_beg"] = fields[4]
(*doc)["fcst_valid_end"] = fields[5]
(*doc)["obs_lead"] = fields[6]
(*doc)["obs_valid_beg"] = fields[7]
(*doc)["obs_valid_end"] = fields[8]
(*doc)["fcst_var"] = fields[9]
(*doc)["fcst_units"] = fields[10]
(*doc)["fcst_lev"] = fields[11]
(*doc)["obs_var"] = fields[12]
(*doc)["obs_units"] = fields[13]
(*doc)["obs_lev"] = fields[14]
(*doc)["obtype"] = fields[15]
(*doc)["vx_mask"] = fields[16]
(*doc)["interp_mthd"] = fields[17]
(*doc)["interp_pnts"] = fields[18]
(*doc)["fcst_thresh"] = fields[19]
(*doc)["obs_thresh"] = fields[20]
(*doc)["cov_thresh"] = fields[21]
(*doc)["alpha"] = fields[22]
(*doc)["line_type"] = fields[23]
}

func (s *STAT_SSVAR) fill_STAT_SSVAR_Header(metaData VxMetadata, fields []string, doc *map[string]interface{}) {
// fill the met fields
(*doc)["ID"] = metaData.ID
(*doc)["Subset"] = metaData.Subset
(*doc)["Type"] = metaData.Type
(*doc)["SubType"] = metaData.SubType
(*doc)["version"] = fields[0]
(*doc)["model"] = fields[1]
(*doc)["desc"] = fields[2]
(*doc)["fcst_valid_beg"] = fields[4]
(*doc)["fcst_valid_end"] = fields[5]
(*doc)["obs_lead"] = fields[6]
(*doc)["obs_valid_beg"] = fields[7]
(*doc)["obs_valid_end"] = fields[8]
(*doc)["fcst_var"] = fields[9]
(*doc)["fcst_units"] = fields[10]
(*doc)["fcst_lev"] = fields[11]
(*doc)["obs_var"] = fields[12]
(*doc)["obs_units"] = fields[13]
(*doc)["obs_lev"] = fields[14]
(*doc)["obtype"] = fields[15]
(*doc)["vx_mask"] = fields[16]
(*doc)["interp_mthd"] = fields[17]
(*doc)["interp_pnts"] = fields[18]
(*doc)["fcst_thresh"] = fields[19]
(*doc)["obs_thresh"] = fields[20]
(*doc)["cov_thresh"] = fields[21]
(*doc)["alpha"] = fields[22]
(*doc)["line_type"] = fields[23]
}

func (s *MODE_CTS) fill_MODE_CTS_Header(metaData VxMetadata, fields []string, doc *map[string]interface{}) {
// fill the met fields
(*doc)["ID"] = metaData.ID
(*doc)["Subset"] = metaData.Subset
(*doc)["Type"] = metaData.Type
(*doc)["SubType"] = metaData.SubType
(*doc)["version"] = fields[0]
(*doc)["model"] = fields[1]
(*doc)["n_valid"] = fields[2]
(*doc)["grid_res"] = fields[3]
(*doc)["desc"] = fields[4]
(*doc)["fcst_valid"] = fields[6]
(*doc)["fcst_accum"] = fields[7]
(*doc)["obs_lead"] = fields[8]
(*doc)["obs_valid"] = fields[9]
(*doc)["obs_accum"] = fields[10]
(*doc)["fcst_rad"] = fields[11]
(*doc)["fcst_thr"] = fields[12]
(*doc)["obs_rad"] = fields[13]
(*doc)["obs_thr"] = fields[14]
(*doc)["fcst_var"] = fields[15]
(*doc)["fcst_units"] = fields[16]
(*doc)["obs_var"] = fields[18]
(*doc)["obs_units"] = fields[19]
(*doc)["obs_lev"] = fields[20]
}

func (s *STAT_CTS) fill_STAT_CTS_Header(metaData VxMetadata, fields []string, doc *map[string]interface{}) {
// fill the met fields
(*doc)["ID"] = metaData.ID
(*doc)["Subset"] = metaData.Subset
(*doc)["Type"] = metaData.Type
(*doc)["SubType"] = metaData.SubType
(*doc)["version"] = fields[0]
(*doc)["model"] = fields[1]
(*doc)["desc"] = fields[2]
(*doc)["fcst_valid_beg"] = fields[4]
(*doc)["fcst_valid_end"] = fields[5]
(*doc)["obs_lead"] = fields[6]
(*doc)["obs_valid_beg"] = fields[7]
(*doc)["obs_valid_end"] = fields[8]
(*doc)["fcst_var"] = fields[9]
(*doc)["fcst_units"] = fields[10]
(*doc)["fcst_lev"] = fields[11]
(*doc)["obs_var"] = fields[12]
(*doc)["obs_units"] = fields[13]
(*doc)["obs_lev"] = fields[14]
(*doc)["obtype"] = fields[15]
(*doc)["vx_mask"] = fields[16]
(*doc)["interp_mthd"] = fields[17]
(*doc)["interp_pnts"] = fields[18]
(*doc)["fcst_thresh"] = fields[19]
(*doc)["obs_thresh"] = fields[20]
(*doc)["cov_thresh"] = fields[21]
(*doc)["alpha"] = fields[22]
(*doc)["line_type"] = fields[23]
}

func (s *STAT_RPS) fill_STAT_RPS_Header(metaData VxMetadata, fields []string, doc *map[string]interface{}) {
// fill the met fields
(*doc)["ID"] = metaData.ID
(*doc)["Subset"] = metaData.Subset
(*doc)["Type"] = metaData.Type
(*doc)["SubType"] = metaData.SubType
(*doc)["version"] = fields[0]
(*doc)["model"] = fields[1]
(*doc)["desc"] = fields[2]
(*doc)["fcst_valid_beg"] = fields[4]
(*doc)["fcst_valid_end"] = fields[5]
(*doc)["obs_lead"] = fields[6]
(*doc)["obs_valid_beg"] = fields[7]
(*doc)["obs_valid_end"] = fields[8]
(*doc)["fcst_var"] = fields[9]
(*doc)["fcst_units"] = fields[10]
(*doc)["fcst_lev"] = fields[11]
(*doc)["obs_var"] = fields[12]
(*doc)["obs_units"] = fields[13]
(*doc)["obs_lev"] = fields[14]
(*doc)["obtype"] = fields[15]
(*doc)["vx_mask"] = fields[16]
(*doc)["interp_mthd"] = fields[17]
(*doc)["interp_pnts"] = fields[18]
(*doc)["fcst_thresh"] = fields[19]
(*doc)["obs_thresh"] = fields[20]
(*doc)["cov_thresh"] = fields[21]
(*doc)["alpha"] = fields[22]
(*doc)["line_type"] = fields[23]
}

func (s *STAT_PHIST) fill_STAT_PHIST_Header(metaData VxMetadata, fields []string, doc *map[string]interface{}) {
// fill the met fields
(*doc)["ID"] = metaData.ID
(*doc)["Subset"] = metaData.Subset
(*doc)["Type"] = metaData.Type
(*doc)["SubType"] = metaData.SubType
(*doc)["version"] = fields[0]
(*doc)["model"] = fields[1]
(*doc)["desc"] = fields[2]
(*doc)["fcst_valid_beg"] = fields[4]
(*doc)["fcst_valid_end"] = fields[5]
(*doc)["obs_lead"] = fields[6]
(*doc)["obs_valid_beg"] = fields[7]
(*doc)["obs_valid_end"] = fields[8]
(*doc)["fcst_var"] = fields[9]
(*doc)["fcst_units"] = fields[10]
(*doc)["fcst_lev"] = fields[11]
(*doc)["obs_var"] = fields[12]
(*doc)["obs_units"] = fields[13]
(*doc)["obs_lev"] = fields[14]
(*doc)["obtype"] = fields[15]
(*doc)["vx_mask"] = fields[16]
(*doc)["interp_mthd"] = fields[17]
(*doc)["interp_pnts"] = fields[18]
(*doc)["fcst_thresh"] = fields[19]
(*doc)["obs_thresh"] = fields[20]
(*doc)["cov_thresh"] = fields[21]
(*doc)["alpha"] = fields[22]
(*doc)["line_type"] = fields[23]
}

func (s *STAT_VAL1L2) fill_STAT_VAL1L2_Header(metaData VxMetadata, fields []string, doc *map[string]interface{}) {
// fill the met fields
(*doc)["ID"] = metaData.ID
(*doc)["Subset"] = metaData.Subset
(*doc)["Type"] = metaData.Type
(*doc)["SubType"] = metaData.SubType
(*doc)["version"] = fields[0]
(*doc)["model"] = fields[1]
(*doc)["desc"] = fields[2]
(*doc)["fcst_valid_beg"] = fields[4]
(*doc)["fcst_valid_end"] = fields[5]
(*doc)["obs_lead"] = fields[6]
(*doc)["obs_valid_beg"] = fields[7]
(*doc)["obs_valid_end"] = fields[8]
(*doc)["fcst_var"] = fields[9]
(*doc)["fcst_units"] = fields[10]
(*doc)["fcst_lev"] = fields[11]
(*doc)["obs_var"] = fields[12]
(*doc)["obs_units"] = fields[13]
(*doc)["obs_lev"] = fields[14]
(*doc)["obtype"] = fields[15]
(*doc)["vx_mask"] = fields[16]
(*doc)["interp_mthd"] = fields[17]
(*doc)["interp_pnts"] = fields[18]
(*doc)["fcst_thresh"] = fields[19]
(*doc)["obs_thresh"] = fields[20]
(*doc)["cov_thresh"] = fields[21]
(*doc)["alpha"] = fields[22]
(*doc)["line_type"] = fields[23]
}

func (s *STAT_FHO) fill_STAT_FHO_Header(metaData VxMetadata, fields []string, doc *map[string]interface{}) {
// fill the met fields
(*doc)["ID"] = metaData.ID
(*doc)["Subset"] = metaData.Subset
(*doc)["Type"] = metaData.Type
(*doc)["SubType"] = metaData.SubType
(*doc)["version"] = fields[0]
(*doc)["model"] = fields[1]
(*doc)["desc"] = fields[2]
(*doc)["fcst_valid_beg"] = fields[4]
(*doc)["fcst_valid_end"] = fields[5]
(*doc)["obs_lead"] = fields[6]
(*doc)["obs_valid_beg"] = fields[7]
(*doc)["obs_valid_end"] = fields[8]
(*doc)["fcst_var"] = fields[9]
(*doc)["fcst_units"] = fields[10]
(*doc)["fcst_lev"] = fields[11]
(*doc)["obs_var"] = fields[12]
(*doc)["obs_units"] = fields[13]
(*doc)["obs_lev"] = fields[14]
(*doc)["obtype"] = fields[15]
(*doc)["vx_mask"] = fields[16]
(*doc)["interp_mthd"] = fields[17]
(*doc)["interp_pnts"] = fields[18]
(*doc)["fcst_thresh"] = fields[19]
(*doc)["obs_thresh"] = fields[20]
(*doc)["cov_thresh"] = fields[21]
(*doc)["alpha"] = fields[22]
(*doc)["line_type"] = fields[23]
}

func (s *STAT_ISC) fill_STAT_ISC_Header(metaData VxMetadata, fields []string, doc *map[string]interface{}) {
// fill the met fields
(*doc)["ID"] = metaData.ID
(*doc)["Subset"] = metaData.Subset
(*doc)["Type"] = metaData.Type
(*doc)["SubType"] = metaData.SubType
(*doc)["version"] = fields[0]
(*doc)["model"] = fields[1]
(*doc)["desc"] = fields[2]
(*doc)["fcst_valid_beg"] = fields[4]
(*doc)["fcst_valid_end"] = fields[5]
(*doc)["obs_lead"] = fields[6]
(*doc)["obs_valid_beg"] = fields[7]
(*doc)["obs_valid_end"] = fields[8]
(*doc)["fcst_var"] = fields[9]
(*doc)["fcst_units"] = fields[10]
(*doc)["fcst_lev"] = fields[11]
(*doc)["obs_var"] = fields[12]
(*doc)["obs_units"] = fields[13]
(*doc)["obs_lev"] = fields[14]
(*doc)["obtype"] = fields[15]
(*doc)["vx_mask"] = fields[16]
(*doc)["interp_mthd"] = fields[17]
(*doc)["interp_pnts"] = fields[18]
(*doc)["fcst_thresh"] = fields[19]
(*doc)["obs_thresh"] = fields[20]
(*doc)["cov_thresh"] = fields[21]
(*doc)["alpha"] = fields[22]
(*doc)["line_type"] = fields[23]
}

func (s *STAT_SEEPS_MPR) fill_STAT_SEEPS_MPR_Header(metaData VxMetadata, fields []string, doc *map[string]interface{}) {
// fill the met fields
(*doc)["ID"] = metaData.ID
(*doc)["Subset"] = metaData.Subset
(*doc)["Type"] = metaData.Type
(*doc)["SubType"] = metaData.SubType
(*doc)["version"] = fields[0]
(*doc)["model"] = fields[1]
(*doc)["desc"] = fields[2]
(*doc)["fcst_valid_beg"] = fields[4]
(*doc)["fcst_valid_end"] = fields[5]
(*doc)["obs_lead"] = fields[6]
(*doc)["obs_valid_beg"] = fields[7]
(*doc)["obs_valid_end"] = fields[8]
(*doc)["fcst_var"] = fields[9]
(*doc)["fcst_units"] = fields[10]
(*doc)["fcst_lev"] = fields[11]
(*doc)["obs_var"] = fields[12]
(*doc)["obs_units"] = fields[13]
(*doc)["obs_lev"] = fields[14]
(*doc)["obtype"] = fields[15]
(*doc)["vx_mask"] = fields[16]
(*doc)["interp_mthd"] = fields[17]
(*doc)["interp_pnts"] = fields[18]
(*doc)["fcst_thresh"] = fields[19]
(*doc)["obs_thresh"] = fields[20]
(*doc)["cov_thresh"] = fields[21]
(*doc)["alpha"] = fields[22]
(*doc)["line_type"] = fields[23]
}

func (s *STAT_DMAP) fill_STAT_DMAP_Header(metaData VxMetadata, fields []string, doc *map[string]interface{}) {
// fill the met fields
(*doc)["ID"] = metaData.ID
(*doc)["Subset"] = metaData.Subset
(*doc)["Type"] = metaData.Type
(*doc)["SubType"] = metaData.SubType
(*doc)["version"] = fields[0]
(*doc)["model"] = fields[1]
(*doc)["desc"] = fields[2]
(*doc)["fcst_valid_beg"] = fields[4]
(*doc)["fcst_valid_end"] = fields[5]
(*doc)["obs_lead"] = fields[6]
(*doc)["obs_valid_beg"] = fields[7]
(*doc)["obs_valid_end"] = fields[8]
(*doc)["fcst_var"] = fields[9]
(*doc)["fcst_units"] = fields[10]
(*doc)["fcst_lev"] = fields[11]
(*doc)["obs_var"] = fields[12]
(*doc)["obs_units"] = fields[13]
(*doc)["obs_lev"] = fields[14]
(*doc)["obtype"] = fields[15]
(*doc)["vx_mask"] = fields[16]
(*doc)["interp_mthd"] = fields[17]
(*doc)["interp_pnts"] = fields[18]
(*doc)["fcst_thresh"] = fields[19]
(*doc)["obs_thresh"] = fields[20]
(*doc)["cov_thresh"] = fields[21]
(*doc)["alpha"] = fields[22]
(*doc)["line_type"] = fields[23]
}

func (s *STAT_RELP) fill_STAT_RELP_Header(metaData VxMetadata, fields []string, doc *map[string]interface{}) {
// fill the met fields
(*doc)["ID"] = metaData.ID
(*doc)["Subset"] = metaData.Subset
(*doc)["Type"] = metaData.Type
(*doc)["SubType"] = metaData.SubType
(*doc)["version"] = fields[0]
(*doc)["model"] = fields[1]
(*doc)["desc"] = fields[2]
(*doc)["fcst_valid_beg"] = fields[4]
(*doc)["fcst_valid_end"] = fields[5]
(*doc)["obs_lead"] = fields[6]
(*doc)["obs_valid_beg"] = fields[7]
(*doc)["obs_valid_end"] = fields[8]
(*doc)["fcst_var"] = fields[9]
(*doc)["fcst_units"] = fields[10]
(*doc)["fcst_lev"] = fields[11]
(*doc)["obs_var"] = fields[12]
(*doc)["obs_units"] = fields[13]
(*doc)["obs_lev"] = fields[14]
(*doc)["obtype"] = fields[15]
(*doc)["vx_mask"] = fields[16]
(*doc)["interp_mthd"] = fields[17]
(*doc)["interp_pnts"] = fields[18]
(*doc)["fcst_thresh"] = fields[19]
(*doc)["obs_thresh"] = fields[20]
(*doc)["cov_thresh"] = fields[21]
(*doc)["alpha"] = fields[22]
(*doc)["line_type"] = fields[23]
}

func (s *STAT_VCNT) fill_STAT_VCNT_Header(metaData VxMetadata, fields []string, doc *map[string]interface{}) {
// fill the met fields
(*doc)["ID"] = metaData.ID
(*doc)["Subset"] = metaData.Subset
(*doc)["Type"] = metaData.Type
(*doc)["SubType"] = metaData.SubType
(*doc)["version"] = fields[0]
(*doc)["model"] = fields[1]
(*doc)["desc"] = fields[2]
(*doc)["fcst_valid_beg"] = fields[4]
(*doc)["fcst_valid_end"] = fields[5]
(*doc)["obs_lead"] = fields[6]
(*doc)["obs_valid_beg"] = fields[7]
(*doc)["obs_valid_end"] = fields[8]
(*doc)["fcst_var"] = fields[9]
(*doc)["fcst_units"] = fields[10]
(*doc)["fcst_lev"] = fields[11]
(*doc)["obs_var"] = fields[12]
(*doc)["obs_units"] = fields[13]
(*doc)["obs_lev"] = fields[14]
(*doc)["obtype"] = fields[15]
(*doc)["vx_mask"] = fields[16]
(*doc)["interp_mthd"] = fields[17]
(*doc)["interp_pnts"] = fields[18]
(*doc)["fcst_thresh"] = fields[19]
(*doc)["obs_thresh"] = fields[20]
(*doc)["cov_thresh"] = fields[21]
(*doc)["alpha"] = fields[22]
(*doc)["line_type"] = fields[23]
}

func (s *MTD_2DSINGLE) fill_MTD_2DSINGLE_Header(metaData VxMetadata, fields []string, doc *map[string]interface{}) {
// fill the met fields
(*doc)["ID"] = metaData.ID
(*doc)["Subset"] = metaData.Subset
(*doc)["Type"] = metaData.Type
(*doc)["SubType"] = metaData.SubType
(*doc)["version"] = fields[0]
(*doc)["model"] = fields[1]
(*doc)["desc"] = fields[2]
(*doc)["fcst_valid"] = fields[4]
(*doc)["obs_lead"] = fields[5]
(*doc)["obs_valid"] = fields[6]
(*doc)["t_delta"] = fields[7]
(*doc)["fcst_t_beg"] = fields[8]
(*doc)["fcst_t_end"] = fields[9]
(*doc)["fcst_rad"] = fields[10]
(*doc)["fcst_thr"] = fields[11]
(*doc)["obs_t_beg"] = fields[12]
(*doc)["obs_t_end"] = fields[13]
(*doc)["obs_rad"] = fields[14]
(*doc)["obs_thr"] = fields[15]
(*doc)["fcst_var"] = fields[16]
(*doc)["fcst_units"] = fields[17]
(*doc)["obs_var"] = fields[19]
(*doc)["obs_units"] = fields[20]
(*doc)["obs_lev"] = fields[21]
}

func (s *TCST_TCDIAG) fill_TCST_TCDIAG_Header(metaData VxMetadata, fields []string, doc *map[string]interface{}) {
// fill the met fields
(*doc)["ID"] = metaData.ID
(*doc)["Subset"] = metaData.Subset
(*doc)["Type"] = metaData.Type
(*doc)["SubType"] = metaData.SubType
(*doc)["version"] = fields[0]
(*doc)["amodel"] = fields[1]
(*doc)["bmodel"] = fields[2]
(*doc)["desc"] = fields[3]
(*doc)["storm_id"] = fields[4]
(*doc)["basin"] = fields[5]
(*doc)["cyclone"] = fields[6]
(*doc)["storm_name"] = fields[7]
(*doc)["init"] = fields[8]
(*doc)["lead"] = fields[9]
(*doc)["valid"] = fields[10]
(*doc)["init_mask"] = fields[11]
(*doc)["valid_mask"] = fields[12]
(*doc)["line_type"] = fields[13]
}


//line data struct definitions
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

type MODE_OBJ struct {
    Obtype                     string  `json:"obtype"`
    Object_id                  string  `json:"object_id"`
    Object_cat                 string  `json:"object_cat"`
    Centroid_x                 string  `json:"centroid_x"`
    Centroid_y                 string  `json:"centroid_y"`
    Centroid_lat               string  `json:"centroid_lat"`
    Centroid_lon               string  `json:"centroid_lon"`
    Axis_ang                   string  `json:"axis_ang"`
    Length                     string  `json:"length"`
    Width                      string  `json:"width"`
    Area                       string  `json:"area"`
    Area_thresh                string  `json:"area_thresh"`
    Curvature                  string  `json:"curvature"`
    Curvature_x                string  `json:"curvature_x"`
    Curvature_y                string  `json:"curvature_y"`
    Complexity                 string  `json:"complexity"`
    Intensity_10               string  `json:"intensity_10"`
    Intensity_25               string  `json:"intensity_25"`
    Intensity_50               string  `json:"intensity_50"`
    Intensity_75               string  `json:"intensity_75"`
    Intensity_90               string  `json:"intensity_90"`
    Intensity_user             string  `json:"intensity_user"`
    Intensity_sum              string  `json:"intensity_sum"`
    Centroid_dist              string  `json:"centroid_dist"`
    Boundary_dist              string  `json:"boundary_dist"`
    Convex_hull_dist           string  `json:"convex_hull_dist"`
    Angle_diff                 string  `json:"angle_diff"`
    Aspect_diff                string  `json:"aspect_diff"`
    Area_ratio                 string  `json:"area_ratio"`
    Intersection_area          string  `json:"intersection_area"`
    Union_area                 string  `json:"union_area"`
    Symmetric_diff             string  `json:"symmetric_diff"`
    Intersection_over_area     string  `json:"intersection_over_area"`
    Curvature_ratio            string  `json:"curvature_ratio"`
    Complexity_ratio           string  `json:"complexity_ratio"`
    Percentile_intensity_ratio string  `json:"percentile_intensity_ratio"`
    Interest                   string  `json:"interest"`
}

type STAT_PRC struct {
    Total         int     `json:"total"`
    N_thresh      int     `json:"n_thresh"`
    Thresh_i      string  `json:"thresh_i"`
    Pody_i        string  `json:"pody_i"`
    Pofd_i        string  `json:"pofd_i"`
}

type STAT_PHIST struct {
    Total      int     `json:"total"`
    Bin_size   float64 `json:"bin_size"`
    N_bin      int     `json:"n_bin"`
    Bin_i      string  `json:"bin_i"`
}

type MODE_CTS struct {
    Obtype string  `json:"obtype"`
    Field  string  `json:"field"`
    Total  int     `json:"total"`
    Fy_oy  float64 `json:"fy_oy"`
    Fy_on  float64 `json:"fy_on"`
    Fn_oy  float64 `json:"fn_oy"`
    Fn_on  float64 `json:"fn_on"`
    Baser  float64 `json:"baser"`
    Fmean  string  `json:"fmean"`
    Acc    string  `json:"acc"`
    Fbias  float64 `json:"fbias"`
    Pody   string  `json:"pody"`
    Podn   string  `json:"podn"`
    Pofd   string  `json:"pofd"`
    Far    string  `json:"far"`
    Csi    string  `json:"csi"`
    Gss    string  `json:"gss"`
    Hk     string  `json:"hk"`
    Hss    string  `json:"hss"`
    Odds   string  `json:"odds"`
    Lodds  string  `json:"lodds"`
    Orss   string  `json:"orss"`
    Eds    string  `json:"eds"`
    Seds   string  `json:"seds"`
    Edi    string  `json:"edi"`
    Sedi   string  `json:"sedi"`
    Bagss  string  `json:"bagss"`
}

type STAT_GENMPR struct {
    Total      int     `json:"total"`
    Index      int     `json:"index"`
    Storm_id   string  `json:"storm_id"`
    Prob_lead  string  `json:"prob_lead"`
    Prob_val   string  `json:"prob_val"`
    Agen_init  string  `json:"agen_init"`
    Agen_fhr   string  `json:"agen_fhr"`
    Agen_lat   string  `json:"agen_lat"`
    Agen_lon   string  `json:"agen_lon"`
    Agen_dland string  `json:"agen_dland"`
    Bgen_lat   string  `json:"bgen_lat"`
    Bgen_lon   string  `json:"bgen_lon"`
    Bgen_dland string  `json:"bgen_dland"`
    Gen_dist   string  `json:"gen_dist"`
    Gen_tdiff  string  `json:"gen_tdiff"`
    Init_tdiff string  `json:"init_tdiff"`
    Dev_cat    string  `json:"dev_cat"`
    Ops_cat    string  `json:"ops_cat"`
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

type STAT_PSTD struct {
    Total         int     `json:"total"`
    N_thresh      int     `json:"n_thresh"`
    Baser         float64 `json:"baser"`
    Baser_ncl     string  `json:"baser_ncl"`
    Baser_ncu     string  `json:"baser_ncu"`
    Reliability   string  `json:"reliability"`
    Resolution    string  `json:"resolution"`
    Uncertainty   string  `json:"uncertainty"`
    Roc_auc       string  `json:"roc_auc"`
    Brier         string  `json:"brier"`
    Brier_ncl     string  `json:"brier_ncl"`
    Brier_ncu     string  `json:"brier_ncu"`
    Briercl       string  `json:"briercl"`
    Briercl_ncl   string  `json:"briercl_ncl"`
    Briercl_ncu   string  `json:"briercl_ncu"`
    Bss           string  `json:"bss"`
    Bss_smpl      string  `json:"bss_smpl"`
    Thresh_i      string  `json:"thresh_i"`
}

type STAT_ECLV struct {
    Total        int     `json:"total"`
    Baser        float64 `json:"baser"`
    Value_baser  string  `json:"value_baser"`
    N_pts        string  `json:"n_pts"`
    Cl_i         string  `json:"cl_i"`
    Value_i      string  `json:"value_i"`
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

type STAT_PJC struct {
    Total              int     `json:"total"`
    N_thresh           int     `json:"n_thresh"`
    Thresh_i           string  `json:"thresh_i"`
    Oy_tp_i            string  `json:"oy_tp_i"`
    On_tp_i            string  `json:"on_tp_i"`
    Calibration_i      string  `json:"calibration_i"`
    Refinement_i       string  `json:"refinement_i"`
    Likelihood_i       string  `json:"likelihood_i"`
    Baser_i            string  `json:"baser_i"`
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

type STAT_SSIDX struct {
    Fcst_model string  `json:"fcst_model"`
    Ref_model  string  `json:"ref_model"`
    N_init     string  `json:"n_init"`
    N_term     string  `json:"n_term"`
    N_vld      string  `json:"n_vld"`
    Ss_index   string  `json:"ss_index"`
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

type MTD_3DPAIR struct {
    Object_id           string  `json:"object_id"`
    Object_cat          string  `json:"object_cat"`
    Space_centroid_dist string  `json:"space_centroid_dist"`
    Time_centroid_delta string  `json:"time_centroid_delta"`
    Axis_diff           string  `json:"axis_diff"`
    Speed_delta         string  `json:"speed_delta"`
    Direction_diff      string  `json:"direction_diff"`
    Volume_ratio        string  `json:"volume_ratio"`
    Start_time_delta    string  `json:"start_time_delta"`
    End_time_delta      string  `json:"end_time_delta"`
    Intersection_volume string  `json:"intersection_volume"`
    Duration_diff       string  `json:"duration_diff"`
    Interest            string  `json:"interest"`
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

type MTD_2DSINGLE struct {
    Object_id      string  `json:"object_id"`
    Object_cat     string  `json:"object_cat"`
    Time_index     string  `json:"time_index"`
    Area           string  `json:"area"`
    Centroid_x     string  `json:"centroid_x"`
    Centroid_y     string  `json:"centroid_y"`
    Centroid_lat   string  `json:"centroid_lat"`
    Centroid_lon   string  `json:"centroid_lon"`
    Axis_ang       string  `json:"axis_ang"`
    Intensity_10   string  `json:"intensity_10"`
    Intensity_25   string  `json:"intensity_25"`
    Intensity_50   string  `json:"intensity_50"`
    Intensity_75   string  `json:"intensity_75"`
    Intensity_90   string  `json:"intensity_90"`
    Intensity_user string  `json:"intensity_user"`
}

type TCST_TCMPR struct {
    Total          int     `json:"total"`
    Index          int     `json:"index"`
    Level          string  `json:"level"`
    Watch_warn     string  `json:"watch_warn"`
    Initials       string  `json:"initials"`
    Alat           string  `json:"alat"`
    Alon           string  `json:"alon"`
    Blat           string  `json:"blat"`
    Blon           string  `json:"blon"`
    Tk_err         string  `json:"tk_err"`
    X_err          string  `json:"x_err"`
    Y_err          string  `json:"y_err"`
    Altk_err       string  `json:"altk_err"`
    Crtk_err       string  `json:"crtk_err"`
    Adland         string  `json:"adland"`
    Bdland         string  `json:"bdland"`
    Amslp          string  `json:"amslp"`
    Bmslp          string  `json:"bmslp"`
    Amax_wind      string  `json:"amax_wind"`
    Bmax_wind      string  `json:"bmax_wind"`
    Aal_wind_34    string  `json:"aal_wind_34"`
    Bal_wind_34    string  `json:"bal_wind_34"`
    Ane_wind_34    string  `json:"ane_wind_34"`
    Bne_wind_34    string  `json:"bne_wind_34"`
    Ase_wind_34    string  `json:"ase_wind_34"`
    Bse_wind_34    string  `json:"bse_wind_34"`
    Asw_wind_34    string  `json:"asw_wind_34"`
    Bsw_wind_34    string  `json:"bsw_wind_34"`
    Anw_wind_34    string  `json:"anw_wind_34"`
    Bnw_wind_34    string  `json:"bnw_wind_34"`
    Aal_wind_50    string  `json:"aal_wind_50"`
    Bal_wind_50    string  `json:"bal_wind_50"`
    Ane_wind_50    string  `json:"ane_wind_50"`
    Bne_wind_50    string  `json:"bne_wind_50"`
    Ase_wind_50    string  `json:"ase_wind_50"`
    Bse_wind_50    string  `json:"bse_wind_50"`
    Asw_wind_50    string  `json:"asw_wind_50"`
    Bsw_wind_50    string  `json:"bsw_wind_50"`
    Anw_wind_50    string  `json:"anw_wind_50"`
    Bnw_wind_50    string  `json:"bnw_wind_50"`
    Aal_wind_64    string  `json:"aal_wind_64"`
    Bal_wind_64    string  `json:"bal_wind_64"`
    Ane_wind_64    string  `json:"ane_wind_64"`
    Bne_wind_64    string  `json:"bne_wind_64"`
    Ase_wind_64    string  `json:"ase_wind_64"`
    Bse_wind_64    string  `json:"bse_wind_64"`
    Asw_wind_64    string  `json:"asw_wind_64"`
    Bsw_wind_64    string  `json:"bsw_wind_64"`
    Anw_wind_64    string  `json:"anw_wind_64"`
    Bnw_wind_64    string  `json:"bnw_wind_64"`
    Aradp          string  `json:"aradp"`
    Bradp          string  `json:"bradp"`
    Arrp           string  `json:"arrp"`
    Brrp           string  `json:"brrp"`
    Amrd           string  `json:"amrd"`
    Bmrd           string  `json:"bmrd"`
    Agusts         string  `json:"agusts"`
    Bgusts         string  `json:"bgusts"`
    Aeye           string  `json:"aeye"`
    Beye           string  `json:"beye"`
    Adir           string  `json:"adir"`
    Bdir           string  `json:"bdir"`
    Aspeed         string  `json:"aspeed"`
    Bspeed         string  `json:"bspeed"`
    Adepth         string  `json:"adepth"`
    Bdepth         string  `json:"bdepth"`
    Num_members    string  `json:"num_members"`
    Track_spread   string  `json:"track_spread"`
    Track_stdev    string  `json:"track_stdev"`
    Mslp_stdev     string  `json:"mslp_stdev"`
    Max_wind_stdev string  `json:"max_wind_stdev"`
}

type TCST_PROBRIRW struct {
    Alat          string  `json:"alat"`
    Alon          string  `json:"alon"`
    Blat          string  `json:"blat"`
    Blon          string  `json:"blon"`
    Initials      string  `json:"initials"`
    Tk_err        string  `json:"tk_err"`
    X_err         string  `json:"x_err"`
    Y_err         string  `json:"y_err"`
    Adland        string  `json:"adland"`
    Bdland        string  `json:"bdland"`
    Rirw_beg      string  `json:"rirw_beg"`
    Rirw_end      string  `json:"rirw_end"`
    Rirw_window   string  `json:"rirw_window"`
    Awind_end     string  `json:"awind_end"`
    Bwind_beg     string  `json:"bwind_beg"`
    Bwind_end     string  `json:"bwind_end"`
    Bdelta        string  `json:"bdelta"`
    Bdelta_max    string  `json:"bdelta_max"`
    Blevel_beg    string  `json:"blevel_beg"`
    Blevel_end    string  `json:"blevel_end"`
    N_thresh      int     `json:"n_thresh"`
    Thresh_i      string  `json:"thresh_i"`
    Prob_i        string  `json:"prob_i"`
}

type MTD_3DSINGLE struct {
    Object_id       string  `json:"object_id"`
    Object_cat      string  `json:"object_cat"`
    Centroid_x      string  `json:"centroid_x"`
    Centroid_y      string  `json:"centroid_y"`
    Centroid_t      string  `json:"centroid_t"`
    Centroid_lat    string  `json:"centroid_lat"`
    Centroid_lon    string  `json:"centroid_lon"`
    X_dot           string  `json:"x_dot"`
    Y_dot           string  `json:"y_dot"`
    Axis_ang        string  `json:"axis_ang"`
    Volume          string  `json:"volume"`
    Start_time      string  `json:"start_time"`
    End_time        string  `json:"end_time"`
    Cdist_travelled string  `json:"cdist_travelled"`
    Intensity_10    string  `json:"intensity_10"`
    Intensity_25    string  `json:"intensity_25"`
    Intensity_50    string  `json:"intensity_50"`
    Intensity_75    string  `json:"intensity_75"`
    Intensity_90    string  `json:"intensity_90"`
    Intensity_user  string  `json:"intensity_user"`
}

type STAT_CTC struct {
    Total    int     `json:"total"`
    Fy_oy    float64 `json:"fy_oy"`
    Fy_on    float64 `json:"fy_on"`
    Fn_oy    float64 `json:"fn_oy"`
    Fn_on    float64 `json:"fn_on"`
    Ec_value float64 `json:"ec_value"`
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

type STAT_RHIST struct {
    Total       int     `json:"total"`
    N_rank      int     `json:"n_rank"`
    Rank_i      string  `json:"rank_i"`
}

type STAT_FHO struct {
    Total  int     `json:"total"`
    F_rate float64 `json:"f_rate"`
    H_rate float64 `json:"h_rate"`
    O_rate float64 `json:"o_rate"`
}

type STAT_RELP struct {
    Total       int     `json:"total"`
    N_ens       float64 `json:"n_ens"`
    Relp_i      string  `json:"relp_i"`
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

type TCST_TCDIAG struct {
    Total        int     `json:"total"`
    Index        int     `json:"index"`
    Diag_source  string  `json:"diag_source"`
    Track_source string  `json:"track_source"`
    Field_source string  `json:"field_source"`
    N_diag       string  `json:"n_diag"`
    Diag_i       string  `json:"diag_i"`
    Value_i      string  `json:"value_i"`
}

type STAT_MCTC struct {
    Total           int     `json:"total"`
    N_cat           int     `json:"n_cat"`
    Fi_oi           string  `json:"fi_oi"`
    Ec_value        float64 `json:"ec_value"`
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

type STAT_PCT struct {
    Total         int     `json:"total"`
    N_thresh      int     `json:"n_thresh"`
    Thresh_i      string  `json:"thresh_i"`
    Oy_i          string  `json:"oy_i"`
    On_i          string  `json:"on_i"`
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


//fillStructure functions
func (s *STAT_SEEPS_MPR) fill_STAT_SEEPS_MPR(fields []string) {
s.Obs_sid = fields[0]
s.Obs_lat, _ = strconv.ParseFloat(fields[1], 64)
s.Obs_lon, _ = strconv.ParseFloat(fields[2], 64)
s.Fcst, _ = strconv.ParseFloat(fields[3], 64)
s.Obs, _ = strconv.ParseFloat(fields[4], 64)
s.Obs_qc = fields[5]
s.Fcst_cat, _ = strconv.Atoi(fields[6])
s.Obs_cat, _ = strconv.Atoi(fields[7])
s.P1, _ = strconv.ParseFloat(fields[8], 64)
s.P2, _ = strconv.ParseFloat(fields[9], 64)
s.T1, _ = strconv.ParseFloat(fields[10], 64)
s.T2, _ = strconv.ParseFloat(fields[11], 64)
s.Seeps, _ = strconv.ParseFloat(fields[12], 64)
}

func (s *STAT_NBRCNT) fill_STAT_NBRCNT(fields []string) {
s.Total, _ = strconv.Atoi(fields[0])
s.Fbs, _ = strconv.ParseFloat(fields[1], 64)
s.Fbs_bcl = fields[2]
s.Fbs_bcu = fields[3]
s.Fss, _ = strconv.ParseFloat(fields[4], 64)
s.Fss_bcl = fields[5]
s.Fss_bcu = fields[6]
s.Afss, _ = strconv.ParseFloat(fields[7], 64)
s.Afss_bcl = fields[8]
s.Afss_bcu = fields[9]
s.Ufss, _ = strconv.ParseFloat(fields[10], 64)
s.Ufss_bcl = fields[11]
s.Ufss_bcu = fields[12]
s.F_rate, _ = strconv.ParseFloat(fields[13], 64)
s.F_rate_bcl = fields[14]
s.F_rate_bcu = fields[15]
s.O_rate, _ = strconv.ParseFloat(fields[16], 64)
s.O_rate_bcl = fields[17]
s.O_rate_bcu = fields[18]
}

func (s *STAT_RHIST) fill_STAT_RHIST(fields []string) {
s.Total, _ = strconv.Atoi(fields[0])
s.N_rank, _ = strconv.Atoi(fields[1])
s.Rank_i = fields[2]
}

func (s *STAT_PHIST) fill_STAT_PHIST(fields []string) {
s.Total, _ = strconv.Atoi(fields[0])
s.Bin_size, _ = strconv.ParseFloat(fields[1], 64)
s.N_bin, _ = strconv.Atoi(fields[2])
s.Bin_i = fields[3]
}

func (s *STAT_RELP) fill_STAT_RELP(fields []string) {
s.Total, _ = strconv.Atoi(fields[0])
s.N_ens, _ = strconv.ParseFloat(fields[1], 64)
s.Relp_i = fields[2]
}

func (s *TCST_TCMPR) fill_TCST_TCMPR(fields []string) {
s.Total, _ = strconv.Atoi(fields[0])
s.Index, _ = strconv.Atoi(fields[1])
s.Level = fields[2]
s.Watch_warn = fields[3]
s.Initials = fields[4]
s.Alat = fields[5]
s.Alon = fields[6]
s.Blat = fields[7]
s.Blon = fields[8]
s.Tk_err = fields[9]
s.X_err = fields[10]
s.Y_err = fields[11]
s.Altk_err = fields[12]
s.Crtk_err = fields[13]
s.Adland = fields[14]
s.Bdland = fields[15]
s.Amslp = fields[16]
s.Bmslp = fields[17]
s.Amax_wind = fields[18]
s.Bmax_wind = fields[19]
s.Aal_wind_34 = fields[20]
s.Bal_wind_34 = fields[21]
s.Ane_wind_34 = fields[22]
s.Bne_wind_34 = fields[23]
s.Ase_wind_34 = fields[24]
s.Bse_wind_34 = fields[25]
s.Asw_wind_34 = fields[26]
s.Bsw_wind_34 = fields[27]
s.Anw_wind_34 = fields[28]
s.Bnw_wind_34 = fields[29]
s.Aal_wind_50 = fields[30]
s.Bal_wind_50 = fields[31]
s.Ane_wind_50 = fields[32]
s.Bne_wind_50 = fields[33]
s.Ase_wind_50 = fields[34]
s.Bse_wind_50 = fields[35]
s.Asw_wind_50 = fields[36]
s.Bsw_wind_50 = fields[37]
s.Anw_wind_50 = fields[38]
s.Bnw_wind_50 = fields[39]
s.Aal_wind_64 = fields[40]
s.Bal_wind_64 = fields[41]
s.Ane_wind_64 = fields[42]
s.Bne_wind_64 = fields[43]
s.Ase_wind_64 = fields[44]
s.Bse_wind_64 = fields[45]
s.Asw_wind_64 = fields[46]
s.Bsw_wind_64 = fields[47]
s.Anw_wind_64 = fields[48]
s.Bnw_wind_64 = fields[49]
s.Aradp = fields[50]
s.Bradp = fields[51]
s.Arrp = fields[52]
s.Brrp = fields[53]
s.Amrd = fields[54]
s.Bmrd = fields[55]
s.Agusts = fields[56]
s.Bgusts = fields[57]
s.Aeye = fields[58]
s.Beye = fields[59]
s.Adir = fields[60]
s.Bdir = fields[61]
s.Aspeed = fields[62]
s.Bspeed = fields[63]
s.Adepth = fields[64]
s.Bdepth = fields[65]
s.Num_members = fields[66]
s.Track_spread = fields[67]
s.Track_stdev = fields[68]
s.Mslp_stdev = fields[69]
s.Max_wind_stdev = fields[70]
}

func (s *TCST_PROBRIRW) fill_TCST_PROBRIRW(fields []string) {
s.Alat = fields[0]
s.Alon = fields[1]
s.Blat = fields[2]
s.Blon = fields[3]
s.Initials = fields[4]
s.Tk_err = fields[5]
s.X_err = fields[6]
s.Y_err = fields[7]
s.Adland = fields[8]
s.Bdland = fields[9]
s.Rirw_beg = fields[10]
s.Rirw_end = fields[11]
s.Rirw_window = fields[12]
s.Awind_end = fields[13]
s.Bwind_beg = fields[14]
s.Bwind_end = fields[15]
s.Bdelta = fields[16]
s.Bdelta_max = fields[17]
s.Blevel_beg = fields[18]
s.Blevel_end = fields[19]
s.N_thresh, _ = strconv.Atoi(fields[20])
s.Thresh_i = fields[21]
s.Prob_i = fields[22]
}

func (s *STAT_MCTS) fill_STAT_MCTS(fields []string) {
s.Total, _ = strconv.Atoi(fields[0])
s.N_cat, _ = strconv.Atoi(fields[1])
s.Acc = fields[2]
s.Acc_ncl = fields[3]
s.Acc_ncu = fields[4]
s.Acc_bcl = fields[5]
s.Acc_bcu = fields[6]
s.Hk = fields[7]
s.Hk_bcl = fields[8]
s.Hk_bcu = fields[9]
s.Hss = fields[10]
s.Hss_bcl = fields[11]
s.Hss_bcu = fields[12]
s.Ger = fields[13]
s.Ger_bcl = fields[14]
s.Ger_bcu = fields[15]
s.Hss_ec = fields[16]
s.Hss_ec_bcl = fields[17]
s.Hss_ec_bcu = fields[18]
s.Ec_value, _ = strconv.ParseFloat(fields[19], 64)
}

func (s *STAT_PCT) fill_STAT_PCT(fields []string) {
s.Total, _ = strconv.Atoi(fields[0])
s.N_thresh, _ = strconv.Atoi(fields[1])
s.Thresh_i = fields[2]
s.Oy_i = fields[3]
s.On_i = fields[4]
}

func (s *STAT_SAL1L2) fill_STAT_SAL1L2(fields []string) {
s.Total, _ = strconv.Atoi(fields[0])
s.Fabar, _ = strconv.ParseFloat(fields[1], 64)
s.Oabar, _ = strconv.ParseFloat(fields[2], 64)
s.Foabar, _ = strconv.ParseFloat(fields[3], 64)
s.Ffabar, _ = strconv.ParseFloat(fields[4], 64)
s.Ooabar, _ = strconv.ParseFloat(fields[5], 64)
s.Mae, _ = strconv.ParseFloat(fields[6], 64)
}

func (s *STAT_VL1L2) fill_STAT_VL1L2(fields []string) {
s.Total, _ = strconv.Atoi(fields[0])
s.Ufbar, _ = strconv.ParseFloat(fields[1], 64)
s.Vfbar, _ = strconv.ParseFloat(fields[2], 64)
s.Uobar, _ = strconv.ParseFloat(fields[3], 64)
s.Vobar, _ = strconv.ParseFloat(fields[4], 64)
s.Uvfobar, _ = strconv.ParseFloat(fields[5], 64)
s.Uvffbar, _ = strconv.ParseFloat(fields[6], 64)
s.Uvoobar, _ = strconv.ParseFloat(fields[7], 64)
s.F_speed_bar, _ = strconv.ParseFloat(fields[8], 64)
s.O_speed_bar, _ = strconv.ParseFloat(fields[9], 64)
s.Total_dir, _ = strconv.Atoi(fields[10])
s.Dir_me, _ = strconv.ParseFloat(fields[11], 64)
s.Dir_mae, _ = strconv.ParseFloat(fields[12], 64)
s.Dir_mse, _ = strconv.ParseFloat(fields[13], 64)
}

func (s *STAT_GENMPR) fill_STAT_GENMPR(fields []string) {
s.Total, _ = strconv.Atoi(fields[0])
s.Index, _ = strconv.Atoi(fields[1])
s.Storm_id = fields[2]
s.Prob_lead = fields[3]
s.Prob_val = fields[4]
s.Agen_init = fields[5]
s.Agen_fhr = fields[6]
s.Agen_lat = fields[7]
s.Agen_lon = fields[8]
s.Agen_dland = fields[9]
s.Bgen_lat = fields[10]
s.Bgen_lon = fields[11]
s.Bgen_dland = fields[12]
s.Gen_dist = fields[13]
s.Gen_tdiff = fields[14]
s.Init_tdiff = fields[15]
s.Dev_cat = fields[16]
s.Ops_cat = fields[17]
}

func (s *MTD_3DSINGLE) fill_MTD_3DSINGLE(fields []string) {
s.Object_id = fields[0]
s.Object_cat = fields[1]
s.Centroid_x = fields[2]
s.Centroid_y = fields[3]
s.Centroid_t = fields[4]
s.Centroid_lat = fields[5]
s.Centroid_lon = fields[6]
s.X_dot = fields[7]
s.Y_dot = fields[8]
s.Axis_ang = fields[9]
s.Volume = fields[10]
s.Start_time = fields[11]
s.End_time = fields[12]
s.Cdist_travelled = fields[13]
s.Intensity_10 = fields[14]
s.Intensity_25 = fields[15]
s.Intensity_50 = fields[16]
s.Intensity_75 = fields[17]
s.Intensity_90 = fields[18]
s.Intensity_user = fields[19]
}

func (s *STAT_CNT) fill_STAT_CNT(fields []string) {
s.Total, _ = strconv.Atoi(fields[0])
s.Fbar, _ = strconv.ParseFloat(fields[1], 64)
s.Fbar_ncl = fields[2]
s.Fbar_ncu = fields[3]
s.Fbar_bcl = fields[4]
s.Fbar_bcu = fields[5]
s.Fstdev = fields[6]
s.Fstdev_ncl = fields[7]
s.Fstdev_ncu = fields[8]
s.Fstdev_bcl = fields[9]
s.Fstdev_bcu = fields[10]
s.Obar, _ = strconv.ParseFloat(fields[11], 64)
s.Obar_ncl = fields[12]
s.Obar_ncu = fields[13]
s.Obar_bcl = fields[14]
s.Obar_bcu = fields[15]
s.Ostdev = fields[16]
s.Ostdev_ncl = fields[17]
s.Ostdev_ncu = fields[18]
s.Ostdev_bcl = fields[19]
s.Ostdev_bcu = fields[20]
s.Pr_corr = fields[21]
s.Pr_corr_ncl = fields[22]
s.Pr_corr_ncu = fields[23]
s.Pr_corr_bcl = fields[24]
s.Pr_corr_bcu = fields[25]
s.Sp_corr = fields[26]
s.Kt_corr = fields[27]
s.Ranks = fields[28]
s.Frank_ties = fields[29]
s.Orank_ties = fields[30]
s.Me, _ = strconv.ParseFloat(fields[31], 64)
s.Me_ncl = fields[32]
s.Me_ncu = fields[33]
s.Me_bcl = fields[34]
s.Me_bcu = fields[35]
s.Estdev = fields[36]
s.Estdev_ncl = fields[37]
s.Estdev_ncu = fields[38]
s.Estdev_bcl = fields[39]
s.Estdev_bcu = fields[40]
s.Mbias = fields[41]
s.Mbias_bcl = fields[42]
s.Mbias_bcu = fields[43]
s.Mae, _ = strconv.ParseFloat(fields[44], 64)
s.Mae_bcl = fields[45]
s.Mae_bcu = fields[46]
s.Mse, _ = strconv.ParseFloat(fields[47], 64)
s.Mse_bcl = fields[48]
s.Mse_bcu = fields[49]
s.Bcmse = fields[50]
s.Bcmse_bcl = fields[51]
s.Bcmse_bcu = fields[52]
s.Rmse, _ = strconv.ParseFloat(fields[53], 64)
s.Rmse_bcl = fields[54]
s.Rmse_bcu = fields[55]
s.E10 = fields[56]
s.E10_bcl = fields[57]
s.E10_bcu = fields[58]
s.E25 = fields[59]
s.E25_bcl = fields[60]
s.E25_bcu = fields[61]
s.E50 = fields[62]
s.E50_bcl = fields[63]
s.E50_bcu = fields[64]
s.E75 = fields[65]
s.E75_bcl = fields[66]
s.E75_bcu = fields[67]
s.E90 = fields[68]
s.E90_bcl = fields[69]
s.E90_bcu = fields[70]
s.Eiqr = fields[71]
s.Eiqr_bcl = fields[72]
s.Eiqr_bcu = fields[73]
s.Mad = fields[74]
s.Mad_bcl = fields[75]
s.Mad_bcu = fields[76]
s.Anom_corr = fields[77]
s.Anom_corr_ncl = fields[78]
s.Anom_corr_ncu = fields[79]
s.Anom_corr_bcl = fields[80]
s.Anom_corr_bcu = fields[81]
s.Me2 = fields[82]
s.Me2_bcl = fields[83]
s.Me2_bcu = fields[84]
s.Msess = fields[85]
s.Msess_bcl = fields[86]
s.Msess_bcu = fields[87]
s.Rmsfa = fields[88]
s.Rmsfa_bcl = fields[89]
s.Rmsfa_bcu = fields[90]
s.Rmsoa = fields[91]
s.Rmsoa_bcl = fields[92]
s.Rmsoa_bcu = fields[93]
s.Anom_corr_uncntr = fields[94]
s.Anom_corr_uncntr_bcl = fields[95]
s.Anom_corr_uncntr_bcu = fields[96]
s.Si, _ = strconv.ParseFloat(fields[97], 64)
s.Si_bcl = fields[98]
s.Si_bcu = fields[99]
}

func (s *STAT_MCTC) fill_STAT_MCTC(fields []string) {
s.Total, _ = strconv.Atoi(fields[0])
s.N_cat, _ = strconv.Atoi(fields[1])
s.Fi_oi = fields[2]
s.Ec_value, _ = strconv.ParseFloat(fields[3], 64)
}

func (s *STAT_NBRCTS) fill_STAT_NBRCTS(fields []string) {
s.Total, _ = strconv.Atoi(fields[0])
s.Baser, _ = strconv.ParseFloat(fields[1], 64)
s.Baser_ncl = fields[2]
s.Baser_ncu = fields[3]
s.Baser_bcl = fields[4]
s.Baser_bcu = fields[5]
s.Fmean = fields[6]
s.Fmean_ncl = fields[7]
s.Fmean_ncu = fields[8]
s.Fmean_bcl = fields[9]
s.Fmean_bcu = fields[10]
s.Acc = fields[11]
s.Acc_ncl = fields[12]
s.Acc_ncu = fields[13]
s.Acc_bcl = fields[14]
s.Acc_bcu = fields[15]
s.Fbias, _ = strconv.ParseFloat(fields[16], 64)
s.Fbias_bcl = fields[17]
s.Fbias_bcu = fields[18]
s.Pody = fields[19]
s.Pody_ncl = fields[20]
s.Pody_ncu = fields[21]
s.Pody_bcl = fields[22]
s.Pody_bcu = fields[23]
s.Podn = fields[24]
s.Podn_ncl = fields[25]
s.Podn_ncu = fields[26]
s.Podn_bcl = fields[27]
s.Podn_bcu = fields[28]
s.Pofd = fields[29]
s.Pofd_ncl = fields[30]
s.Pofd_ncu = fields[31]
s.Pofd_bcl = fields[32]
s.Pofd_bcu = fields[33]
s.Far = fields[34]
s.Far_ncl = fields[35]
s.Far_ncu = fields[36]
s.Far_bcl = fields[37]
s.Far_bcu = fields[38]
s.Csi = fields[39]
s.Csi_ncl = fields[40]
s.Csi_ncu = fields[41]
s.Csi_bcl = fields[42]
s.Csi_bcu = fields[43]
s.Gss = fields[44]
s.Gss_bcl = fields[45]
s.Gss_bcu = fields[46]
s.Hk = fields[47]
s.Hk_ncl = fields[48]
s.Hk_ncu = fields[49]
s.Hk_bcl = fields[50]
s.Hk_bcu = fields[51]
s.Hss = fields[52]
s.Hss_bcl = fields[53]
s.Hss_bcu = fields[54]
s.Odds = fields[55]
s.Odds_ncl = fields[56]
s.Odds_ncu = fields[57]
s.Odds_bcl = fields[58]
s.Odds_bcu = fields[59]
s.Lodds = fields[60]
s.Lodds_ncl = fields[61]
s.Lodds_ncu = fields[62]
s.Lodds_bcl = fields[63]
s.Lodds_bcu = fields[64]
s.Orss = fields[65]
s.Orss_ncl = fields[66]
s.Orss_ncu = fields[67]
s.Orss_bcl = fields[68]
s.Orss_bcu = fields[69]
s.Eds = fields[70]
s.Eds_ncl = fields[71]
s.Eds_ncu = fields[72]
s.Eds_bcl = fields[73]
s.Eds_bcu = fields[74]
s.Seds = fields[75]
s.Seds_ncl = fields[76]
s.Seds_ncu = fields[77]
s.Seds_bcl = fields[78]
s.Seds_bcu = fields[79]
s.Edi = fields[80]
s.Edi_ncl = fields[81]
s.Edi_ncu = fields[82]
s.Edi_bcl = fields[83]
s.Edi_bcu = fields[84]
s.Sedi = fields[85]
s.Sedi_ncl = fields[86]
s.Sedi_ncu = fields[87]
s.Sedi_bcl = fields[88]
s.Sedi_bcu = fields[89]
s.Bagss = fields[90]
s.Bagss_bcl = fields[91]
s.Bagss_bcu = fields[92]
}

func (s *STAT_PRC) fill_STAT_PRC(fields []string) {
s.Total, _ = strconv.Atoi(fields[0])
s.N_thresh, _ = strconv.Atoi(fields[1])
s.Thresh_i = fields[2]
s.Pody_i = fields[3]
s.Pofd_i = fields[4]
}

func (s *STAT_ECNT) fill_STAT_ECNT(fields []string) {
s.Total, _ = strconv.Atoi(fields[0])
s.N_ens, _ = strconv.ParseFloat(fields[1], 64)
s.Crps, _ = strconv.ParseFloat(fields[2], 64)
s.Crpss, _ = strconv.ParseFloat(fields[3], 64)
s.Ign, _ = strconv.ParseFloat(fields[4], 64)
s.Me, _ = strconv.ParseFloat(fields[5], 64)
s.Rmse, _ = strconv.ParseFloat(fields[6], 64)
s.Spread, _ = strconv.ParseFloat(fields[7], 64)
s.Me_oerr, _ = strconv.ParseFloat(fields[8], 64)
s.Rmse_oerr, _ = strconv.ParseFloat(fields[9], 64)
s.Spread_oerr, _ = strconv.ParseFloat(fields[10], 64)
s.Spread_plus_oerr, _ = strconv.ParseFloat(fields[11], 64)
s.Crpscl, _ = strconv.ParseFloat(fields[12], 64)
s.Crps_emp, _ = strconv.ParseFloat(fields[13], 64)
s.Crpscl_emp, _ = strconv.ParseFloat(fields[14], 64)
s.Crpss_emp, _ = strconv.ParseFloat(fields[15], 64)
s.Crps_emp_fair, _ = strconv.ParseFloat(fields[16], 64)
s.Spread_md, _ = strconv.ParseFloat(fields[17], 64)
s.Mae, _ = strconv.ParseFloat(fields[18], 64)
s.Mae_oerr, _ = strconv.ParseFloat(fields[19], 64)
s.Bias_ratio, _ = strconv.ParseFloat(fields[20], 64)
s.N_ge_obs, _ = strconv.Atoi(fields[21])
s.Me_ge_obs, _ = strconv.ParseFloat(fields[22], 64)
s.N_lt_obs, _ = strconv.Atoi(fields[23])
s.Me_lt_obs, _ = strconv.ParseFloat(fields[24], 64)
s.Ign_conv_oerr, _ = strconv.ParseFloat(fields[25], 64)
s.Ign_corr_oerr, _ = strconv.ParseFloat(fields[26], 64)
}

func (s *STAT_SSIDX) fill_STAT_SSIDX(fields []string) {
s.Fcst_model = fields[0]
s.Ref_model = fields[1]
s.N_init = fields[2]
s.N_term = fields[3]
s.N_vld = fields[4]
s.Ss_index = fields[5]
}

func (s *MODE_OBJ) fill_MODE_OBJ(fields []string) {
s.Obtype = fields[0]
s.Object_id = fields[1]
s.Object_cat = fields[2]
s.Centroid_x = fields[3]
s.Centroid_y = fields[4]
s.Centroid_lat = fields[5]
s.Centroid_lon = fields[6]
s.Axis_ang = fields[7]
s.Length = fields[8]
s.Width = fields[9]
s.Area = fields[10]
s.Area_thresh = fields[11]
s.Curvature = fields[12]
s.Curvature_x = fields[13]
s.Curvature_y = fields[14]
s.Complexity = fields[15]
s.Intensity_10 = fields[16]
s.Intensity_25 = fields[17]
s.Intensity_50 = fields[18]
s.Intensity_75 = fields[19]
s.Intensity_90 = fields[20]
s.Intensity_user = fields[21]
s.Intensity_sum = fields[22]
s.Centroid_dist = fields[23]
s.Boundary_dist = fields[24]
s.Convex_hull_dist = fields[25]
s.Angle_diff = fields[26]
s.Aspect_diff = fields[27]
s.Area_ratio = fields[28]
s.Intersection_area = fields[29]
s.Union_area = fields[30]
s.Symmetric_diff = fields[31]
s.Intersection_over_area = fields[32]
s.Curvature_ratio = fields[33]
s.Complexity_ratio = fields[34]
s.Percentile_intensity_ratio = fields[35]
s.Interest = fields[36]
}

func (s *STAT_ECLV) fill_STAT_ECLV(fields []string) {
s.Total, _ = strconv.Atoi(fields[0])
s.Baser, _ = strconv.ParseFloat(fields[1], 64)
s.Value_baser = fields[2]
s.N_pts = fields[3]
s.Cl_i = fields[4]
s.Value_i = fields[5]
}

func (s *STAT_SL1L2) fill_STAT_SL1L2(fields []string) {
s.Total, _ = strconv.Atoi(fields[0])
s.Fbar, _ = strconv.ParseFloat(fields[1], 64)
s.Obar, _ = strconv.ParseFloat(fields[2], 64)
s.Fobar, _ = strconv.ParseFloat(fields[3], 64)
s.Ffbar, _ = strconv.ParseFloat(fields[4], 64)
s.Oobar, _ = strconv.ParseFloat(fields[5], 64)
s.Mae, _ = strconv.ParseFloat(fields[6], 64)
}

func (s *STAT_VAL1L2) fill_STAT_VAL1L2(fields []string) {
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

func (s *STAT_SEEPS) fill_STAT_SEEPS(fields []string) {
s.Total, _ = strconv.Atoi(fields[0])
s.Odfl, _ = strconv.ParseFloat(fields[1], 64)
s.Odfh, _ = strconv.ParseFloat(fields[2], 64)
s.Olfd, _ = strconv.ParseFloat(fields[3], 64)
s.Olfh, _ = strconv.ParseFloat(fields[4], 64)
s.Ohfd, _ = strconv.ParseFloat(fields[5], 64)
s.Ohfl, _ = strconv.ParseFloat(fields[6], 64)
s.Pf1, _ = strconv.ParseFloat(fields[7], 64)
s.Pf2, _ = strconv.ParseFloat(fields[8], 64)
s.Pf3, _ = strconv.ParseFloat(fields[9], 64)
s.Pv1, _ = strconv.ParseFloat(fields[10], 64)
s.Pv2, _ = strconv.ParseFloat(fields[11], 64)
s.Pv3, _ = strconv.ParseFloat(fields[12], 64)
s.Mean_fcst, _ = strconv.ParseFloat(fields[13], 64)
s.Mean_obs, _ = strconv.ParseFloat(fields[14], 64)
s.Seeps, _ = strconv.ParseFloat(fields[15], 64)
}

func (s *STAT_DMAP) fill_STAT_DMAP(fields []string) {
s.Total, _ = strconv.Atoi(fields[0])
s.Fy, _ = strconv.ParseFloat(fields[1], 64)
s.Oy, _ = strconv.ParseFloat(fields[2], 64)
s.Fbias, _ = strconv.ParseFloat(fields[3], 64)
s.Baddeley = fields[4]
s.Hausdorff = fields[5]
s.Med_fo = fields[6]
s.Med_of = fields[7]
s.Med_min = fields[8]
s.Med_max = fields[9]
s.Med_mean = fields[10]
s.Fom_fo = fields[11]
s.Fom_of = fields[12]
s.Fom_min = fields[13]
s.Fom_max = fields[14]
s.Fom_mean = fields[15]
s.Zhu_fo = fields[16]
s.Zhu_of = fields[17]
s.Zhu_min = fields[18]
s.Zhu_max = fields[19]
s.Zhu_mean = fields[20]
s.G, _ = strconv.ParseFloat(fields[21], 64)
s.Gbeta = fields[22]
s.Beta_value = fields[23]
}

func (s *STAT_PSTD) fill_STAT_PSTD(fields []string) {
s.Total, _ = strconv.Atoi(fields[0])
s.N_thresh, _ = strconv.Atoi(fields[1])
s.Baser, _ = strconv.ParseFloat(fields[2], 64)
s.Baser_ncl = fields[3]
s.Baser_ncu = fields[4]
s.Reliability = fields[5]
s.Resolution = fields[6]
s.Uncertainty = fields[7]
s.Roc_auc = fields[8]
s.Brier = fields[9]
s.Brier_ncl = fields[10]
s.Brier_ncu = fields[11]
s.Briercl = fields[12]
s.Briercl_ncl = fields[13]
s.Briercl_ncu = fields[14]
s.Bss = fields[15]
s.Bss_smpl = fields[16]
s.Thresh_i = fields[17]
}

func (s *MTD_2DSINGLE) fill_MTD_2DSINGLE(fields []string) {
s.Object_id = fields[0]
s.Object_cat = fields[1]
s.Time_index = fields[2]
s.Area = fields[3]
s.Centroid_x = fields[4]
s.Centroid_y = fields[5]
s.Centroid_lat = fields[6]
s.Centroid_lon = fields[7]
s.Axis_ang = fields[8]
s.Intensity_10 = fields[9]
s.Intensity_25 = fields[10]
s.Intensity_50 = fields[11]
s.Intensity_75 = fields[12]
s.Intensity_90 = fields[13]
s.Intensity_user = fields[14]
}

func (s *STAT_CTS) fill_STAT_CTS(fields []string) {
s.Total, _ = strconv.Atoi(fields[0])
s.Baser, _ = strconv.ParseFloat(fields[1], 64)
s.Baser_ncl = fields[2]
s.Baser_ncu = fields[3]
s.Baser_bcl = fields[4]
s.Baser_bcu = fields[5]
s.Fmean = fields[6]
s.Fmean_ncl = fields[7]
s.Fmean_ncu = fields[8]
s.Fmean_bcl = fields[9]
s.Fmean_bcu = fields[10]
s.Acc = fields[11]
s.Acc_ncl = fields[12]
s.Acc_ncu = fields[13]
s.Acc_bcl = fields[14]
s.Acc_bcu = fields[15]
s.Fbias, _ = strconv.ParseFloat(fields[16], 64)
s.Fbias_bcl = fields[17]
s.Fbias_bcu = fields[18]
s.Pody = fields[19]
s.Pody_ncl = fields[20]
s.Pody_ncu = fields[21]
s.Pody_bcl = fields[22]
s.Pody_bcu = fields[23]
s.Podn = fields[24]
s.Podn_ncl = fields[25]
s.Podn_ncu = fields[26]
s.Podn_bcl = fields[27]
s.Podn_bcu = fields[28]
s.Pofd = fields[29]
s.Pofd_ncl = fields[30]
s.Pofd_ncu = fields[31]
s.Pofd_bcl = fields[32]
s.Pofd_bcu = fields[33]
s.Far = fields[34]
s.Far_ncl = fields[35]
s.Far_ncu = fields[36]
s.Far_bcl = fields[37]
s.Far_bcu = fields[38]
s.Csi = fields[39]
s.Csi_ncl = fields[40]
s.Csi_ncu = fields[41]
s.Csi_bcl = fields[42]
s.Csi_bcu = fields[43]
s.Gss = fields[44]
s.Gss_bcl = fields[45]
s.Gss_bcu = fields[46]
s.Hk = fields[47]
s.Hk_ncl = fields[48]
s.Hk_ncu = fields[49]
s.Hk_bcl = fields[50]
s.Hk_bcu = fields[51]
s.Hss = fields[52]
s.Hss_bcl = fields[53]
s.Hss_bcu = fields[54]
s.Odds = fields[55]
s.Odds_ncl = fields[56]
s.Odds_ncu = fields[57]
s.Odds_bcl = fields[58]
s.Odds_bcu = fields[59]
s.Lodds = fields[60]
s.Lodds_ncl = fields[61]
s.Lodds_ncu = fields[62]
s.Lodds_bcl = fields[63]
s.Lodds_bcu = fields[64]
s.Orss = fields[65]
s.Orss_ncl = fields[66]
s.Orss_ncu = fields[67]
s.Orss_bcl = fields[68]
s.Orss_bcu = fields[69]
s.Eds = fields[70]
s.Eds_ncl = fields[71]
s.Eds_ncu = fields[72]
s.Eds_bcl = fields[73]
s.Eds_bcu = fields[74]
s.Seds = fields[75]
s.Seds_ncl = fields[76]
s.Seds_ncu = fields[77]
s.Seds_bcl = fields[78]
s.Seds_bcu = fields[79]
s.Edi = fields[80]
s.Edi_ncl = fields[81]
s.Edi_ncu = fields[82]
s.Edi_bcl = fields[83]
s.Edi_bcu = fields[84]
s.Sedi = fields[85]
s.Sedi_ncl = fields[86]
s.Sedi_ncu = fields[87]
s.Sedi_bcl = fields[88]
s.Sedi_bcu = fields[89]
s.Bagss = fields[90]
s.Bagss_bcl = fields[91]
s.Bagss_bcu = fields[92]
s.Hss_ec = fields[93]
s.Hss_ec_bcl = fields[94]
s.Hss_ec_bcu = fields[95]
s.Ec_value, _ = strconv.ParseFloat(fields[96], 64)
}

func (s *STAT_NBRCTC) fill_STAT_NBRCTC(fields []string) {
s.Total, _ = strconv.Atoi(fields[0])
s.Fy_oy, _ = strconv.ParseFloat(fields[1], 64)
s.Fy_on, _ = strconv.ParseFloat(fields[2], 64)
s.Fn_oy, _ = strconv.ParseFloat(fields[3], 64)
s.Fn_on, _ = strconv.ParseFloat(fields[4], 64)
}

func (s *STAT_GRAD) fill_STAT_GRAD(fields []string) {
s.Total, _ = strconv.Atoi(fields[0])
s.Fgbar, _ = strconv.ParseFloat(fields[1], 64)
s.Ogbar, _ = strconv.ParseFloat(fields[2], 64)
s.Mgbar, _ = strconv.ParseFloat(fields[3], 64)
s.Egbar, _ = strconv.ParseFloat(fields[4], 64)
s.S1 = fields[5]
s.S1_og = fields[6]
s.Fgog_ratio = fields[7]
s.Dx, _ = strconv.Atoi(fields[8])
s.Dy, _ = strconv.Atoi(fields[9])
}

func (s *STAT_ORANK) fill_STAT_ORANK(fields []string) {
s.Total, _ = strconv.Atoi(fields[0])
s.Index, _ = strconv.Atoi(fields[1])
s.Obs_sid = fields[2]
s.Obs_lat, _ = strconv.ParseFloat(fields[3], 64)
s.Obs_lon, _ = strconv.ParseFloat(fields[4], 64)
s.Obs_lvl, _ = strconv.ParseFloat(fields[5], 64)
s.Obs_elv, _ = strconv.ParseFloat(fields[6], 64)
s.Obs, _ = strconv.ParseFloat(fields[7], 64)
s.Pit, _ = strconv.ParseFloat(fields[8], 64)
s.Rank, _ = strconv.Atoi(fields[9])
s.N_ens_vld, _ = strconv.Atoi(fields[10])
s.N_ens, _ = strconv.ParseFloat(fields[11], 64)
s.Ens_i = fields[12]
s.Obs_qc = fields[13]
s.Ens_mean, _ = strconv.ParseFloat(fields[14], 64)
s.Obs_climo_mean, _ = strconv.ParseFloat(fields[15], 64)
s.Spread, _ = strconv.ParseFloat(fields[16], 64)
s.Ens_mean_oerr, _ = strconv.ParseFloat(fields[17], 64)
s.Spread_oerr, _ = strconv.ParseFloat(fields[18], 64)
s.Spread_plus_oerr, _ = strconv.ParseFloat(fields[19], 64)
s.Obs_climo_stdev, _ = strconv.ParseFloat(fields[20], 64)
s.Fcst_climo_mean, _ = strconv.ParseFloat(fields[21], 64)
s.Fcst_climo_stdev, _ = strconv.ParseFloat(fields[22], 64)
}

func (s *STAT_PJC) fill_STAT_PJC(fields []string) {
s.Total, _ = strconv.Atoi(fields[0])
s.N_thresh, _ = strconv.Atoi(fields[1])
s.Thresh_i = fields[2]
s.Oy_tp_i = fields[3]
s.On_tp_i = fields[4]
s.Calibration_i = fields[5]
s.Refinement_i = fields[6]
s.Likelihood_i = fields[7]
s.Baser_i = fields[8]
}

func (s *STAT_RPS) fill_STAT_RPS(fields []string) {
s.Total, _ = strconv.Atoi(fields[0])
s.N_prob, _ = strconv.ParseFloat(fields[1], 64)
s.Rps_rel, _ = strconv.ParseFloat(fields[2], 64)
s.Rps_res, _ = strconv.ParseFloat(fields[3], 64)
s.Rps_unc, _ = strconv.ParseFloat(fields[4], 64)
s.Rps, _ = strconv.ParseFloat(fields[5], 64)
s.Rpss, _ = strconv.ParseFloat(fields[6], 64)
s.Rpss_smpl, _ = strconv.ParseFloat(fields[7], 64)
s.Rps_comp = fields[8]
}

func (s *STAT_VCNT) fill_STAT_VCNT(fields []string) {
s.Total, _ = strconv.Atoi(fields[0])
s.Fbar, _ = strconv.ParseFloat(fields[1], 64)
s.Fbar_bcl = fields[2]
s.Fbar_bcu = fields[3]
s.Obar, _ = strconv.ParseFloat(fields[4], 64)
s.Obar_bcl = fields[5]
s.Obar_bcu = fields[6]
s.Fs_rms = fields[7]
s.Fs_rms_bcl = fields[8]
s.Fs_rms_bcu = fields[9]
s.Os_rms = fields[10]
s.Os_rms_bcl = fields[11]
s.Os_rms_bcu = fields[12]
s.Msve = fields[13]
s.Msve_bcl = fields[14]
s.Msve_bcu = fields[15]
s.Rmsve = fields[16]
s.Rmsve_bcl = fields[17]
s.Rmsve_bcu = fields[18]
s.Fstdev = fields[19]
s.Fstdev_bcl = fields[20]
s.Fstdev_bcu = fields[21]
s.Ostdev = fields[22]
s.Ostdev_bcl = fields[23]
s.Ostdev_bcu = fields[24]
s.Fdir = fields[25]
s.Fdir_bcl = fields[26]
s.Fdir_bcu = fields[27]
s.Odir = fields[28]
s.Odir_bcl = fields[29]
s.Odir_bcu = fields[30]
s.Fbar_speed = fields[31]
s.Fbar_speed_bcl = fields[32]
s.Fbar_speed_bcu = fields[33]
s.Obar_speed = fields[34]
s.Obar_speed_bcl = fields[35]
s.Obar_speed_bcu = fields[36]
s.Vdiff_speed = fields[37]
s.Vdiff_speed_bcl = fields[38]
s.Vdiff_speed_bcu = fields[39]
s.Vdiff_dir = fields[40]
s.Vdiff_dir_bcl = fields[41]
s.Vdiff_dir_bcu = fields[42]
s.Speed_err = fields[43]
s.Speed_err_bcl = fields[44]
s.Speed_err_bcu = fields[45]
s.Speed_abserr = fields[46]
s.Speed_abserr_bcl = fields[47]
s.Speed_abserr_bcu = fields[48]
s.Dir_err = fields[49]
s.Dir_err_bcl = fields[50]
s.Dir_err_bcu = fields[51]
s.Dir_abserr = fields[52]
s.Dir_abserr_bcl = fields[53]
s.Dir_abserr_bcu = fields[54]
s.Anom_corr = fields[55]
s.Anom_corr_ncl = fields[56]
s.Anom_corr_ncu = fields[57]
s.Anom_corr_bcl = fields[58]
s.Anom_corr_bcu = fields[59]
s.Anom_corr_uncntr = fields[60]
s.Anom_corr_uncntr_bcl = fields[61]
s.Anom_corr_uncntr_bcu = fields[62]
s.Total_dir, _ = strconv.Atoi(fields[63])
s.Dir_me, _ = strconv.ParseFloat(fields[64], 64)
s.Dir_me_bcl = fields[65]
s.Dir_me_bcu = fields[66]
s.Dir_mae, _ = strconv.ParseFloat(fields[67], 64)
s.Dir_mae_bcl = fields[68]
s.Dir_mae_bcu = fields[69]
s.Dir_mse, _ = strconv.ParseFloat(fields[70], 64)
s.Dir_mse_bcl = fields[71]
s.Dir_mse_bcu = fields[72]
s.Dir_rmse = fields[73]
s.Dir_rmse_bcl = fields[74]
s.Dir_rmse_bcu = fields[75]
}

func (s *TCST_TCDIAG) fill_TCST_TCDIAG(fields []string) {
s.Total, _ = strconv.Atoi(fields[0])
s.Index, _ = strconv.Atoi(fields[1])
s.Diag_source = fields[2]
s.Track_source = fields[3]
s.Field_source = fields[4]
s.N_diag = fields[5]
s.Diag_i = fields[6]
s.Value_i = fields[7]
}

func (s *STAT_FHO) fill_STAT_FHO(fields []string) {
s.Total, _ = strconv.Atoi(fields[0])
s.F_rate, _ = strconv.ParseFloat(fields[1], 64)
s.H_rate, _ = strconv.ParseFloat(fields[2], 64)
s.O_rate, _ = strconv.ParseFloat(fields[3], 64)
}

func (s *STAT_CTC) fill_STAT_CTC(fields []string) {
s.Total, _ = strconv.Atoi(fields[0])
s.Fy_oy, _ = strconv.ParseFloat(fields[1], 64)
s.Fy_on, _ = strconv.ParseFloat(fields[2], 64)
s.Fn_oy, _ = strconv.ParseFloat(fields[3], 64)
s.Fn_on, _ = strconv.ParseFloat(fields[4], 64)
s.Ec_value, _ = strconv.ParseFloat(fields[5], 64)
}

func (s *STAT_ISC) fill_STAT_ISC(fields []string) {
s.Total, _ = strconv.Atoi(fields[0])
s.Tile_dim, _ = strconv.Atoi(fields[1])
s.Tile_xll, _ = strconv.Atoi(fields[2])
s.Tile_yll, _ = strconv.Atoi(fields[3])
s.Nscale, _ = strconv.Atoi(fields[4])
s.Iscale, _ = strconv.Atoi(fields[5])
s.Mse, _ = strconv.ParseFloat(fields[6], 64)
s.Isc, _ = strconv.Atoi(fields[7])
s.Fenergy2, _ = strconv.ParseFloat(fields[8], 64)
s.Oenergy2, _ = strconv.ParseFloat(fields[9], 64)
s.Baser, _ = strconv.ParseFloat(fields[10], 64)
s.Fbias, _ = strconv.ParseFloat(fields[11], 64)
}

func (s *STAT_MPR) fill_STAT_MPR(fields []string) {
s.Total, _ = strconv.Atoi(fields[0])
s.Index, _ = strconv.Atoi(fields[1])
s.Obs_sid = fields[2]
s.Obs_lat, _ = strconv.ParseFloat(fields[3], 64)
s.Obs_lon, _ = strconv.ParseFloat(fields[4], 64)
s.Obs_lvl, _ = strconv.ParseFloat(fields[5], 64)
s.Obs_elv, _ = strconv.ParseFloat(fields[6], 64)
s.Fcst, _ = strconv.ParseFloat(fields[7], 64)
s.Obs, _ = strconv.ParseFloat(fields[8], 64)
s.Obs_qc = fields[9]
s.Obs_climo_mean, _ = strconv.ParseFloat(fields[10], 64)
s.Obs_climo_stdev, _ = strconv.ParseFloat(fields[11], 64)
s.Obs_climo_cdf, _ = strconv.ParseFloat(fields[12], 64)
s.Fcst_climo_mean, _ = strconv.ParseFloat(fields[13], 64)
s.Fcst_climo_stdev, _ = strconv.ParseFloat(fields[14], 64)
}

func (s *STAT_SSVAR) fill_STAT_SSVAR(fields []string) {
s.Total, _ = strconv.Atoi(fields[0])
s.N_bin, _ = strconv.Atoi(fields[1])
s.Bin_i = fields[2]
s.Bin_n, _ = strconv.Atoi(fields[3])
s.Var_min, _ = strconv.ParseFloat(fields[4], 64)
s.Var_max, _ = strconv.ParseFloat(fields[5], 64)
s.Var_mean, _ = strconv.ParseFloat(fields[6], 64)
s.Fbar, _ = strconv.ParseFloat(fields[7], 64)
s.Obar, _ = strconv.ParseFloat(fields[8], 64)
s.Fobar, _ = strconv.ParseFloat(fields[9], 64)
s.Ffbar, _ = strconv.ParseFloat(fields[10], 64)
s.Oobar, _ = strconv.ParseFloat(fields[11], 64)
s.Fbar_ncl = fields[12]
s.Fbar_ncu = fields[13]
s.Fstdev = fields[14]
s.Fstdev_ncl = fields[15]
s.Fstdev_ncu = fields[16]
s.Obar_ncl = fields[17]
s.Obar_ncu = fields[18]
s.Ostdev = fields[19]
s.Ostdev_ncl = fields[20]
s.Ostdev_ncu = fields[21]
s.Pr_corr = fields[22]
s.Pr_corr_ncl = fields[23]
s.Pr_corr_ncu = fields[24]
s.Me, _ = strconv.ParseFloat(fields[25], 64)
s.Me_ncl = fields[26]
s.Me_ncu = fields[27]
s.Estdev = fields[28]
s.Estdev_ncl = fields[29]
s.Estdev_ncu = fields[30]
s.Mbias = fields[31]
s.Mse, _ = strconv.ParseFloat(fields[32], 64)
s.Bcmse = fields[33]
s.Rmse, _ = strconv.ParseFloat(fields[34], 64)
}

func (s *MODE_CTS) fill_MODE_CTS(fields []string) {
s.Obtype = fields[0]
s.Field = fields[1]
s.Total, _ = strconv.Atoi(fields[2])
s.Fy_oy, _ = strconv.ParseFloat(fields[3], 64)
s.Fy_on, _ = strconv.ParseFloat(fields[4], 64)
s.Fn_oy, _ = strconv.ParseFloat(fields[5], 64)
s.Fn_on, _ = strconv.ParseFloat(fields[6], 64)
s.Baser, _ = strconv.ParseFloat(fields[7], 64)
s.Fmean = fields[8]
s.Acc = fields[9]
s.Fbias, _ = strconv.ParseFloat(fields[10], 64)
s.Pody = fields[11]
s.Podn = fields[12]
s.Pofd = fields[13]
s.Far = fields[14]
s.Csi = fields[15]
s.Gss = fields[16]
s.Hk = fields[17]
s.Hss = fields[18]
s.Odds = fields[19]
s.Lodds = fields[20]
s.Orss = fields[21]
s.Eds = fields[22]
s.Seds = fields[23]
s.Edi = fields[24]
s.Sedi = fields[25]
s.Bagss = fields[26]
}

func (s *MTD_3DPAIR) fill_MTD_3DPAIR(fields []string) {
s.Object_id = fields[0]
s.Object_cat = fields[1]
s.Space_centroid_dist = fields[2]
s.Time_centroid_delta = fields[3]
s.Axis_diff = fields[4]
s.Speed_delta = fields[5]
s.Direction_diff = fields[6]
s.Volume_ratio = fields[7]
s.Start_time_delta = fields[8]
s.End_time_delta = fields[9]
s.Intersection_volume = fields[10]
s.Duration_diff = fields[11]
s.Interest = fields[12]
}


//getDocForId functions
func GetDocForId(fileLineType string, metaData VxMetadata, headerData []string, dataData []string, dataKey string) (interface{}, error) {
	switch fileLineType {
	case "STAT_CNT":
	{
		elem := STAT_CNT{}
		elem.fill_STAT_CNT_Header(data)
		if exists := (*doc)["data"]; exists == nil {
			(*doc)["data"] = make(map[string]STAT_CNT)
		}
		if val, ok := (*doc)["data"].(map[string]STAT_CNT); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
}
	case "STAT_CTC":
	{
		elem := STAT_CTC{}
		elem.fill_STAT_CTC_Header(data)
		if exists := (*doc)["data"]; exists == nil {
			(*doc)["data"] = make(map[string]STAT_CTC)
		}
		if val, ok := (*doc)["data"].(map[string]STAT_CTC); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
}
	case "STAT_CTS":
	{
		elem := STAT_CTS{}
		elem.fill_STAT_CTS_Header(data)
		if exists := (*doc)["data"]; exists == nil {
			(*doc)["data"] = make(map[string]STAT_CTS)
		}
		if val, ok := (*doc)["data"].(map[string]STAT_CTS); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
}
	case "STAT_FHO":
	{
		elem := STAT_FHO{}
		elem.fill_STAT_FHO_Header(data)
		if exists := (*doc)["data"]; exists == nil {
			(*doc)["data"] = make(map[string]STAT_FHO)
		}
		if val, ok := (*doc)["data"].(map[string]STAT_FHO); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
}
	case "STAT_ISC":
	{
		elem := STAT_ISC{}
		elem.fill_STAT_ISC_Header(data)
		if exists := (*doc)["data"]; exists == nil {
			(*doc)["data"] = make(map[string]STAT_ISC)
		}
		if val, ok := (*doc)["data"].(map[string]STAT_ISC); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
}
	case "STAT_MCTC":
	{
		elem := STAT_MCTC{}
		elem.fill_STAT_MCTC_Header(data)
		if exists := (*doc)["data"]; exists == nil {
			(*doc)["data"] = make(map[string]STAT_MCTC)
		}
		if val, ok := (*doc)["data"].(map[string]STAT_MCTC); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
}
	case "STAT_MCTS":
	{
		elem := STAT_MCTS{}
		elem.fill_STAT_MCTS_Header(data)
		if exists := (*doc)["data"]; exists == nil {
			(*doc)["data"] = make(map[string]STAT_MCTS)
		}
		if val, ok := (*doc)["data"].(map[string]STAT_MCTS); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
}
	case "STAT_MPR":
	{
		elem := STAT_MPR{}
		elem.fill_STAT_MPR_Header(data)
		if exists := (*doc)["data"]; exists == nil {
			(*doc)["data"] = make(map[string]STAT_MPR)
		}
		if val, ok := (*doc)["data"].(map[string]STAT_MPR); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
}
	case "STAT_SEEPS":
	{
		elem := STAT_SEEPS{}
		elem.fill_STAT_SEEPS_Header(data)
		if exists := (*doc)["data"]; exists == nil {
			(*doc)["data"] = make(map[string]STAT_SEEPS)
		}
		if val, ok := (*doc)["data"].(map[string]STAT_SEEPS); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
}
	case "STAT_SEEPS_MPR":
	{
		elem := STAT_SEEPS_MPR{}
		elem.fill_STAT_SEEPS_MPR_Header(data)
		if exists := (*doc)["data"]; exists == nil {
			(*doc)["data"] = make(map[string]STAT_SEEPS_MPR)
		}
		if val, ok := (*doc)["data"].(map[string]STAT_SEEPS_MPR); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
}
	case "STAT_NBRCNT":
	{
		elem := STAT_NBRCNT{}
		elem.fill_STAT_NBRCNT_Header(data)
		if exists := (*doc)["data"]; exists == nil {
			(*doc)["data"] = make(map[string]STAT_NBRCNT)
		}
		if val, ok := (*doc)["data"].(map[string]STAT_NBRCNT); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
}
	case "STAT_NBRCTC":
	{
		elem := STAT_NBRCTC{}
		elem.fill_STAT_NBRCTC_Header(data)
		if exists := (*doc)["data"]; exists == nil {
			(*doc)["data"] = make(map[string]STAT_NBRCTC)
		}
		if val, ok := (*doc)["data"].(map[string]STAT_NBRCTC); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
}
	case "STAT_NBRCTS":
	{
		elem := STAT_NBRCTS{}
		elem.fill_STAT_NBRCTS_Header(data)
		if exists := (*doc)["data"]; exists == nil {
			(*doc)["data"] = make(map[string]STAT_NBRCTS)
		}
		if val, ok := (*doc)["data"].(map[string]STAT_NBRCTS); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
}
	case "STAT_GRAD":
	{
		elem := STAT_GRAD{}
		elem.fill_STAT_GRAD_Header(data)
		if exists := (*doc)["data"]; exists == nil {
			(*doc)["data"] = make(map[string]STAT_GRAD)
		}
		if val, ok := (*doc)["data"].(map[string]STAT_GRAD); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
}
	case "STAT_DMAP":
	{
		elem := STAT_DMAP{}
		elem.fill_STAT_DMAP_Header(data)
		if exists := (*doc)["data"]; exists == nil {
			(*doc)["data"] = make(map[string]STAT_DMAP)
		}
		if val, ok := (*doc)["data"].(map[string]STAT_DMAP); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
}
	case "STAT_ORANK":
	{
		elem := STAT_ORANK{}
		elem.fill_STAT_ORANK_Header(data)
		if exists := (*doc)["data"]; exists == nil {
			(*doc)["data"] = make(map[string]STAT_ORANK)
		}
		if val, ok := (*doc)["data"].(map[string]STAT_ORANK); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
}
	case "STAT_PCT":
	{
		elem := STAT_PCT{}
		elem.fill_STAT_PCT_Header(data)
		if exists := (*doc)["data"]; exists == nil {
			(*doc)["data"] = make(map[string]STAT_PCT)
		}
		if val, ok := (*doc)["data"].(map[string]STAT_PCT); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
}
	case "STAT_PJC":
	{
		elem := STAT_PJC{}
		elem.fill_STAT_PJC_Header(data)
		if exists := (*doc)["data"]; exists == nil {
			(*doc)["data"] = make(map[string]STAT_PJC)
		}
		if val, ok := (*doc)["data"].(map[string]STAT_PJC); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
}
	case "STAT_PRC":
	{
		elem := STAT_PRC{}
		elem.fill_STAT_PRC_Header(data)
		if exists := (*doc)["data"]; exists == nil {
			(*doc)["data"] = make(map[string]STAT_PRC)
		}
		if val, ok := (*doc)["data"].(map[string]STAT_PRC); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
}
	case "STAT_PSTD":
	{
		elem := STAT_PSTD{}
		elem.fill_STAT_PSTD_Header(data)
		if exists := (*doc)["data"]; exists == nil {
			(*doc)["data"] = make(map[string]STAT_PSTD)
		}
		if val, ok := (*doc)["data"].(map[string]STAT_PSTD); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
}
	case "STAT_ECLV":
	{
		elem := STAT_ECLV{}
		elem.fill_STAT_ECLV_Header(data)
		if exists := (*doc)["data"]; exists == nil {
			(*doc)["data"] = make(map[string]STAT_ECLV)
		}
		if val, ok := (*doc)["data"].(map[string]STAT_ECLV); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
}
	case "STAT_ECNT":
	{
		elem := STAT_ECNT{}
		elem.fill_STAT_ECNT_Header(data)
		if exists := (*doc)["data"]; exists == nil {
			(*doc)["data"] = make(map[string]STAT_ECNT)
		}
		if val, ok := (*doc)["data"].(map[string]STAT_ECNT); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
}
	case "STAT_RPS":
	{
		elem := STAT_RPS{}
		elem.fill_STAT_RPS_Header(data)
		if exists := (*doc)["data"]; exists == nil {
			(*doc)["data"] = make(map[string]STAT_RPS)
		}
		if val, ok := (*doc)["data"].(map[string]STAT_RPS); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
}
	case "STAT_RHIST":
	{
		elem := STAT_RHIST{}
		elem.fill_STAT_RHIST_Header(data)
		if exists := (*doc)["data"]; exists == nil {
			(*doc)["data"] = make(map[string]STAT_RHIST)
		}
		if val, ok := (*doc)["data"].(map[string]STAT_RHIST); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
}
	case "STAT_PHIST":
	{
		elem := STAT_PHIST{}
		elem.fill_STAT_PHIST_Header(data)
		if exists := (*doc)["data"]; exists == nil {
			(*doc)["data"] = make(map[string]STAT_PHIST)
		}
		if val, ok := (*doc)["data"].(map[string]STAT_PHIST); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
}
	case "STAT_RELP":
	{
		elem := STAT_RELP{}
		elem.fill_STAT_RELP_Header(data)
		if exists := (*doc)["data"]; exists == nil {
			(*doc)["data"] = make(map[string]STAT_RELP)
		}
		if val, ok := (*doc)["data"].(map[string]STAT_RELP); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
}
	case "STAT_SAL1L2":
	{
		elem := STAT_SAL1L2{}
		elem.fill_STAT_SAL1L2_Header(data)
		if exists := (*doc)["data"]; exists == nil {
			(*doc)["data"] = make(map[string]STAT_SAL1L2)
		}
		if val, ok := (*doc)["data"].(map[string]STAT_SAL1L2); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
}
	case "STAT_SL1L2":
	{
		elem := STAT_SL1L2{}
		elem.fill_STAT_SL1L2_Header(data)
		if exists := (*doc)["data"]; exists == nil {
			(*doc)["data"] = make(map[string]STAT_SL1L2)
		}
		if val, ok := (*doc)["data"].(map[string]STAT_SL1L2); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
}
	case "STAT_SSVAR":
	{
		elem := STAT_SSVAR{}
		elem.fill_STAT_SSVAR_Header(data)
		if exists := (*doc)["data"]; exists == nil {
			(*doc)["data"] = make(map[string]STAT_SSVAR)
		}
		if val, ok := (*doc)["data"].(map[string]STAT_SSVAR); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
}
	case "STAT_VAL1L2":
	{
		elem := STAT_VAL1L2{}
		elem.fill_STAT_VAL1L2_Header(data)
		if exists := (*doc)["data"]; exists == nil {
			(*doc)["data"] = make(map[string]STAT_VAL1L2)
		}
		if val, ok := (*doc)["data"].(map[string]STAT_VAL1L2); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
}
	case "STAT_VL1L2":
	{
		elem := STAT_VL1L2{}
		elem.fill_STAT_VL1L2_Header(data)
		if exists := (*doc)["data"]; exists == nil {
			(*doc)["data"] = make(map[string]STAT_VL1L2)
		}
		if val, ok := (*doc)["data"].(map[string]STAT_VL1L2); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
}
	case "STAT_VCNT":
	{
		elem := STAT_VCNT{}
		elem.fill_STAT_VCNT_Header(data)
		if exists := (*doc)["data"]; exists == nil {
			(*doc)["data"] = make(map[string]STAT_VCNT)
		}
		if val, ok := (*doc)["data"].(map[string]STAT_VCNT); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
}
	case "STAT_GENMPR":
	{
		elem := STAT_GENMPR{}
		elem.fill_STAT_GENMPR_Header(data)
		if exists := (*doc)["data"]; exists == nil {
			(*doc)["data"] = make(map[string]STAT_GENMPR)
		}
		if val, ok := (*doc)["data"].(map[string]STAT_GENMPR); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
}
	case "STAT_SSIDX":
	{
		elem := STAT_SSIDX{}
		elem.fill_STAT_SSIDX_Header(data)
		if exists := (*doc)["data"]; exists == nil {
			(*doc)["data"] = make(map[string]STAT_SSIDX)
		}
		if val, ok := (*doc)["data"].(map[string]STAT_SSIDX); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
}
	case "MODE_OBJ":
	{
		elem := MODE_OBJ{}
		elem.fill_MODE_OBJ_Header(data)
		if exists := (*doc)["data"]; exists == nil {
			(*doc)["data"] = make(map[string]MODE_OBJ)
		}
		if val, ok := (*doc)["data"].(map[string]MODE_OBJ); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
}
	case "MODE_CTS":
	{
		elem := MODE_CTS{}
		elem.fill_MODE_CTS_Header(data)
		if exists := (*doc)["data"]; exists == nil {
			(*doc)["data"] = make(map[string]MODE_CTS)
		}
		if val, ok := (*doc)["data"].(map[string]MODE_CTS); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
}
	case "MTD_2DSINGLE":
	{
		elem := MTD_2DSINGLE{}
		elem.fill_MTD_2DSINGLE_Header(data)
		if exists := (*doc)["data"]; exists == nil {
			(*doc)["data"] = make(map[string]MTD_2DSINGLE)
		}
		if val, ok := (*doc)["data"].(map[string]MTD_2DSINGLE); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
}
	case "MTD_3DSINGLE":
	{
		elem := MTD_3DSINGLE{}
		elem.fill_MTD_3DSINGLE_Header(data)
		if exists := (*doc)["data"]; exists == nil {
			(*doc)["data"] = make(map[string]MTD_3DSINGLE)
		}
		if val, ok := (*doc)["data"].(map[string]MTD_3DSINGLE); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
}
	case "MTD_3DPAIR":
	{
		elem := MTD_3DPAIR{}
		elem.fill_MTD_3DPAIR_Header(data)
		if exists := (*doc)["data"]; exists == nil {
			(*doc)["data"] = make(map[string]MTD_3DPAIR)
		}
		if val, ok := (*doc)["data"].(map[string]MTD_3DPAIR); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
}
	case "TCST_TCMPR":
	{
		elem := TCST_TCMPR{}
		elem.fill_TCST_TCMPR_Header(data)
		if exists := (*doc)["data"]; exists == nil {
			(*doc)["data"] = make(map[string]TCST_TCMPR)
		}
		if val, ok := (*doc)["data"].(map[string]TCST_TCMPR); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
}
	case "TCST_TCDIAG":
	{
		elem := TCST_TCDIAG{}
		elem.fill_TCST_TCDIAG_Header(data)
		if exists := (*doc)["data"]; exists == nil {
			(*doc)["data"] = make(map[string]TCST_TCDIAG)
		}
		if val, ok := (*doc)["data"].(map[string]TCST_TCDIAG); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
}
	case "TCST_PROBRIRW":
	{
		elem := TCST_PROBRIRW{}
		elem.fill_TCST_PROBRIRW_Header(data)
		if exists := (*doc)["data"]; exists == nil {
			(*doc)["data"] = make(map[string]TCST_PROBRIRW)
		}
		if val, ok := (*doc)["data"].(map[string]TCST_PROBRIRW); ok {
			val[dataKey] = elem
			(*doc)["data"] = val
		}
}
	default:
return errors.New("Unknonwn file_line type:" + fileLineType)
	}
		return nil}


//data key definitions
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
    "MTD_2DSINGLE":  {"FCST_LEAD, FCST_LEV"},
    "MTD_3DSINGLE":  {"FCST_LEAD, FCST_LEV"},
    "MTD_3DPAIR":    {"FCST_LEAD, FCST_LEV"},
    "TCST_TCMPR":     {"FCST_LEAD"},
    "TCST_TCDIAG":    {"FCST_LEAD"},
    "TCST_PROBRIRW":  {"FCST_LEAD"},

}
