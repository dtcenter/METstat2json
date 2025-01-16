package structColumnTypes

import (
	"strconv"
	"errors"
)

/*
THIS CODE IS AUTOMATICALLY GENERATED - DO NOT EDIT THIS CODE
To modify this code - modify the buildHeaderLineTypes.go file and run the buildHeaderLineTypes.go program
cd  <repo path>/metlinetypes/pkg/buildHeaderLineTypes
go run . > /tmp/types.go
cp /tmp/types.go ../structColumnTypes/structColumnTypes.go
*/

const DOC_NOT_FOUND = "Document not found:"

//Header struct definitions
type STAT_MPR_header struct {
    Version        string `json:"VERSION"`
    Model          string `json:"MODEL"`
    Desc           string `json:"DESC"`
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

type STAT_DMAP_header struct {
    Version        string `json:"VERSION"`
    Model          string `json:"MODEL"`
    Desc           string `json:"DESC"`
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

type STAT_RPS_header struct {
    Version        string `json:"VERSION"`
    Model          string `json:"MODEL"`
    Desc           string `json:"DESC"`
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

type STAT_VAL1L2_header struct {
    Version        string `json:"VERSION"`
    Model          string `json:"MODEL"`
    Desc           string `json:"DESC"`
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

type STAT_GENMPR_header struct {
    Version        string `json:"VERSION"`
    Model          string `json:"MODEL"`
    Desc           string `json:"DESC"`
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

type STAT_CTS_header struct {
    Version        string `json:"VERSION"`
    Model          string `json:"MODEL"`
    Desc           string `json:"DESC"`
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

type STAT_SEEPS_MPR_header struct {
    Version        string `json:"VERSION"`
    Model          string `json:"MODEL"`
    Desc           string `json:"DESC"`
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

type STAT_ECNT_header struct {
    Version        string `json:"VERSION"`
    Model          string `json:"MODEL"`
    Desc           string `json:"DESC"`
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

type MTD_3DSINGLE_header struct {
    Version    string `json:"VERSION"`
    Model      string `json:"MODEL"`
    Desc       string `json:"DESC"`
    Fcst_valid string `json:"FCST_VALID"`
    Obs_lead   string `json:"OBS_LEAD"`
    Obs_valid  string `json:"OBS_VALID"`
    T_delta    string `json:"T_DELTA"`
    Fcst_t_beg string `json:"FCST_T_BEG"`
    Fcst_t_end string `json:"FCST_T_END"`
    Fcst_rad   string `json:"FCST_RAD"`
    Fcst_thr   string `json:"FCST_THR"`
    Obs_t_beg  string `json:"OBS_T_BEG"`
    Obs_t_end  string `json:"OBS_T_END"`
    Obs_rad    string `json:"OBS_RAD"`
    Obs_thr    string `json:"OBS_THR"`
    Fcst_var   string `json:"FCST_VAR"`
    Fcst_units string `json:"FCST_UNITS"`
    Obs_var    string `json:"OBS_VAR"`
    Obs_units  string `json:"OBS_UNITS"`
    Obs_lev    string `json:"OBS_LEV"`
}

type STAT_NBRCTS_header struct {
    Version        string `json:"VERSION"`
    Model          string `json:"MODEL"`
    Desc           string `json:"DESC"`
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

type STAT_SSVAR_header struct {
    Version        string `json:"VERSION"`
    Model          string `json:"MODEL"`
    Desc           string `json:"DESC"`
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

type STAT_VCNT_header struct {
    Version        string `json:"VERSION"`
    Model          string `json:"MODEL"`
    Desc           string `json:"DESC"`
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

type STAT_ISC_header struct {
    Version        string `json:"VERSION"`
    Model          string `json:"MODEL"`
    Desc           string `json:"DESC"`
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

type STAT_GRAD_header struct {
    Version        string `json:"VERSION"`
    Model          string `json:"MODEL"`
    Desc           string `json:"DESC"`
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

type STAT_PSTD_header struct {
    Version        string `json:"VERSION"`
    Model          string `json:"MODEL"`
    Desc           string `json:"DESC"`
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

type STAT_PHIST_header struct {
    Version        string `json:"VERSION"`
    Model          string `json:"MODEL"`
    Desc           string `json:"DESC"`
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

type STAT_SL1L2_header struct {
    Version        string `json:"VERSION"`
    Model          string `json:"MODEL"`
    Desc           string `json:"DESC"`
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

type STAT_SSIDX_header struct {
    Version        string `json:"VERSION"`
    Model          string `json:"MODEL"`
    Desc           string `json:"DESC"`
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

type MODE_OBJ_header struct {
    Version                    string `json:"VERSION"`
    Model                      string `json:"MODEL"`
    N_valid                    string `json:"N_VALID"`
    Grid_res                   string `json:"GRID_RES"`
    Desc                       string `json:"DESC"`
    Fcst_valid                 string `json:"FCST_VALID"`
    Fcst_accum                 string `json:"FCST_ACCUM"`
    Obs_lead                   string `json:"OBS_LEAD"`
    Obs_valid                  string `json:"OBS_VALID"`
    Obs_accum                  string `json:"OBS_ACCUM"`
    Fcst_rad                   string `json:"FCST_RAD"`
    Fcst_thr                   string `json:"FCST_THR"`
    Obs_rad                    string `json:"OBS_RAD"`
    Obs_thr                    string `json:"OBS_THR"`
    Fcst_var                   string `json:"FCST_VAR"`
    Fcst_units                 string `json:"FCST_UNITS"`
    Obs_var                    string `json:"OBS_VAR"`
    Obs_units                  string `json:"OBS_UNITS"`
    Obs_lev                    string `json:"OBS_LEV"`
    Obtype                     string `json:"OBTYPE"`
    Object_id                  string `json:"OBJECT_ID"`
    Object_cat                 string `json:"OBJECT_CAT"`
    Centroid_x                 string `json:"CENTROID_X"`
    Centroid_y                 string `json:"CENTROID_Y"`
    Centroid_lat               string `json:"CENTROID_LAT"`
    Centroid_lon               string `json:"CENTROID_LON"`
    Axis_ang                   string `json:"AXIS_ANG"`
    Length                     string `json:"LENGTH"`
    Width                      string `json:"WIDTH"`
    Area                       string `json:"AREA"`
    Area_thresh                string `json:"AREA_THRESH"`
    Curvature                  string `json:"CURVATURE"`
    Curvature_x                string `json:"CURVATURE_X"`
    Curvature_y                string `json:"CURVATURE_Y"`
    Complexity                 string `json:"COMPLEXITY"`
    Intensity_10               string `json:"INTENSITY_10"`
    Intensity_25               string `json:"INTENSITY_25"`
    Intensity_50               string `json:"INTENSITY_50"`
    Intensity_75               string `json:"INTENSITY_75"`
    Intensity_90               string `json:"INTENSITY_90"`
    Intensity_user             string `json:"INTENSITY_USER"`
    Intensity_sum              string `json:"INTENSITY_SUM"`
    Centroid_dist              string `json:"CENTROID_DIST"`
    Boundary_dist              string `json:"BOUNDARY_DIST"`
    Convex_hull_dist           string `json:"CONVEX_HULL_DIST"`
    Angle_diff                 string `json:"ANGLE_DIFF"`
    Aspect_diff                string `json:"ASPECT_DIFF"`
    Area_ratio                 string `json:"AREA_RATIO"`
    Intersection_area          string `json:"INTERSECTION_AREA"`
    Union_area                 string `json:"UNION_AREA"`
    Symmetric_diff             string `json:"SYMMETRIC_DIFF"`
    Intersection_over_area     string `json:"INTERSECTION_OVER_AREA"`
    Curvature_ratio            string `json:"CURVATURE_RATIO"`
    Complexity_ratio           string `json:"COMPLEXITY_RATIO"`
    Percentile_intensity_ratio string `json:"PERCENTILE_INTENSITY_RATIO"`
    Interest                   string `json:"INTEREST"`
}

type MODE_CTS_header struct {
    Version    string `json:"VERSION"`
    Model      string `json:"MODEL"`
    N_valid    string `json:"N_VALID"`
    Grid_res   string `json:"GRID_RES"`
    Desc       string `json:"DESC"`
    Fcst_valid string `json:"FCST_VALID"`
    Fcst_accum string `json:"FCST_ACCUM"`
    Obs_lead   string `json:"OBS_LEAD"`
    Obs_valid  string `json:"OBS_VALID"`
    Obs_accum  string `json:"OBS_ACCUM"`
    Fcst_rad   string `json:"FCST_RAD"`
    Fcst_thr   string `json:"FCST_THR"`
    Obs_rad    string `json:"OBS_RAD"`
    Obs_thr    string `json:"OBS_THR"`
    Fcst_var   string `json:"FCST_VAR"`
    Fcst_units string `json:"FCST_UNITS"`
    Obs_var    string `json:"OBS_VAR"`
    Obs_units  string `json:"OBS_UNITS"`
    Obs_lev    string `json:"OBS_LEV"`
    Obtype     string `json:"OBTYPE"`
    Field      string `json:"FIELD"`
    Total      string `json:"TOTAL"`
    Fy_oy      string `json:"FY_OY"`
    Fy_on      string `json:"FY_ON"`
    Fn_oy      string `json:"FN_OY"`
    Fn_on      string `json:"FN_ON"`
    Baser      string `json:"BASER"`
    Fmean      string `json:"FMEAN"`
    Acc        string `json:"ACC"`
    Fbias      string `json:"FBIAS"`
    Pody       string `json:"PODY"`
    Podn       string `json:"PODN"`
    Pofd       string `json:"POFD"`
    Far        string `json:"FAR"`
    Csi        string `json:"CSI"`
    Gss        string `json:"GSS"`
    Hk         string `json:"HK"`
    Hss        string `json:"HSS"`
    Odds       string `json:"ODDS"`
    Lodds      string `json:"LODDS"`
    Orss       string `json:"ORSS"`
    Eds        string `json:"EDS"`
    Seds       string `json:"SEDS"`
    Edi        string `json:"EDI"`
    Sedi       string `json:"SEDI"`
    Bagss      string `json:"BAGSS"`
}

type MTD_2DSINGLE_header struct {
    Version    string `json:"VERSION"`
    Model      string `json:"MODEL"`
    Desc       string `json:"DESC"`
    Fcst_valid string `json:"FCST_VALID"`
    Obs_lead   string `json:"OBS_LEAD"`
    Obs_valid  string `json:"OBS_VALID"`
    T_delta    string `json:"T_DELTA"`
    Fcst_t_beg string `json:"FCST_T_BEG"`
    Fcst_t_end string `json:"FCST_T_END"`
    Fcst_rad   string `json:"FCST_RAD"`
    Fcst_thr   string `json:"FCST_THR"`
    Obs_t_beg  string `json:"OBS_T_BEG"`
    Obs_t_end  string `json:"OBS_T_END"`
    Obs_rad    string `json:"OBS_RAD"`
    Obs_thr    string `json:"OBS_THR"`
    Fcst_var   string `json:"FCST_VAR"`
    Fcst_units string `json:"FCST_UNITS"`
    Obs_var    string `json:"OBS_VAR"`
    Obs_units  string `json:"OBS_UNITS"`
    Obs_lev    string `json:"OBS_LEV"`
}

type MTD_3DPAIR_header struct {
    Version    string `json:"VERSION"`
    Model      string `json:"MODEL"`
    Desc       string `json:"DESC"`
    Fcst_valid string `json:"FCST_VALID"`
    Obs_lead   string `json:"OBS_LEAD"`
    Obs_valid  string `json:"OBS_VALID"`
    T_delta    string `json:"T_DELTA"`
    Fcst_t_beg string `json:"FCST_T_BEG"`
    Fcst_t_end string `json:"FCST_T_END"`
    Fcst_rad   string `json:"FCST_RAD"`
    Fcst_thr   string `json:"FCST_THR"`
    Obs_t_beg  string `json:"OBS_T_BEG"`
    Obs_t_end  string `json:"OBS_T_END"`
    Obs_rad    string `json:"OBS_RAD"`
    Obs_thr    string `json:"OBS_THR"`
    Fcst_var   string `json:"FCST_VAR"`
    Fcst_units string `json:"FCST_UNITS"`
    Obs_var    string `json:"OBS_VAR"`
    Obs_units  string `json:"OBS_UNITS"`
    Obs_lev    string `json:"OBS_LEV"`
}

type STAT_MCTC_header struct {
    Version        string `json:"VERSION"`
    Model          string `json:"MODEL"`
    Desc           string `json:"DESC"`
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

type STAT_NBRCNT_header struct {
    Version        string `json:"VERSION"`
    Model          string `json:"MODEL"`
    Desc           string `json:"DESC"`
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

type STAT_NBRCTC_header struct {
    Version        string `json:"VERSION"`
    Model          string `json:"MODEL"`
    Desc           string `json:"DESC"`
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

type STAT_ORANK_header struct {
    Version        string `json:"VERSION"`
    Model          string `json:"MODEL"`
    Desc           string `json:"DESC"`
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

type STAT_SAL1L2_header struct {
    Version        string `json:"VERSION"`
    Model          string `json:"MODEL"`
    Desc           string `json:"DESC"`
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

type TCST_PROBRIRW_header struct {
    Version    string `json:"VERSION"`
    Amodel     string `json:"AMODEL"`
    Bmodel     string `json:"BMODEL"`
    Desc       string `json:"DESC"`
    Storm_id   string `json:"STORM_ID"`
    Basin      string `json:"BASIN"`
    Cyclone    string `json:"CYCLONE"`
    Storm_name string `json:"STORM_NAME"`
    Init       string `json:"INIT"`
    Lead       string `json:"LEAD"`
    Valid      string `json:"VALID"`
    Init_mask  string `json:"INIT_MASK"`
    Valid_mask string `json:"VALID_MASK"`
    Line_type  string `json:"LINE_TYPE"`
}

type STAT_MCTS_header struct {
    Version        string `json:"VERSION"`
    Model          string `json:"MODEL"`
    Desc           string `json:"DESC"`
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

type STAT_SEEPS_header struct {
    Version        string `json:"VERSION"`
    Model          string `json:"MODEL"`
    Desc           string `json:"DESC"`
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

type STAT_VL1L2_header struct {
    Version        string `json:"VERSION"`
    Model          string `json:"MODEL"`
    Desc           string `json:"DESC"`
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

type STAT_CNT_header struct {
    Version        string `json:"VERSION"`
    Model          string `json:"MODEL"`
    Desc           string `json:"DESC"`
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

type STAT_PJC_header struct {
    Version        string `json:"VERSION"`
    Model          string `json:"MODEL"`
    Desc           string `json:"DESC"`
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

type STAT_PRC_header struct {
    Version        string `json:"VERSION"`
    Model          string `json:"MODEL"`
    Desc           string `json:"DESC"`
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

type STAT_ECLV_header struct {
    Version        string `json:"VERSION"`
    Model          string `json:"MODEL"`
    Desc           string `json:"DESC"`
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

type STAT_RHIST_header struct {
    Version        string `json:"VERSION"`
    Model          string `json:"MODEL"`
    Desc           string `json:"DESC"`
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

type STAT_RELP_header struct {
    Version        string `json:"VERSION"`
    Model          string `json:"MODEL"`
    Desc           string `json:"DESC"`
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

type TCST_TCDIAG_header struct {
    Version    string `json:"VERSION"`
    Amodel     string `json:"AMODEL"`
    Bmodel     string `json:"BMODEL"`
    Desc       string `json:"DESC"`
    Storm_id   string `json:"STORM_ID"`
    Basin      string `json:"BASIN"`
    Cyclone    string `json:"CYCLONE"`
    Storm_name string `json:"STORM_NAME"`
    Init       string `json:"INIT"`
    Lead       string `json:"LEAD"`
    Valid      string `json:"VALID"`
    Init_mask  string `json:"INIT_MASK"`
    Valid_mask string `json:"VALID_MASK"`
    Line_type  string `json:"LINE_TYPE"`
}

type STAT_CTC_header struct {
    Version        string `json:"VERSION"`
    Model          string `json:"MODEL"`
    Desc           string `json:"DESC"`
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

type STAT_FHO_header struct {
    Version        string `json:"VERSION"`
    Model          string `json:"MODEL"`
    Desc           string `json:"DESC"`
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

type STAT_PCT_header struct {
    Version        string `json:"VERSION"`
    Model          string `json:"MODEL"`
    Desc           string `json:"DESC"`
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

type TCST_TCMPR_header struct {
    Version    string `json:"VERSION"`
    Amodel     string `json:"AMODEL"`
    Bmodel     string `json:"BMODEL"`
    Desc       string `json:"DESC"`
    Storm_id   string `json:"STORM_ID"`
    Basin      string `json:"BASIN"`
    Cyclone    string `json:"CYCLONE"`
    Storm_name string `json:"STORM_NAME"`
    Init       string `json:"INIT"`
    Lead       string `json:"LEAD"`
    Valid      string `json:"VALID"`
    Init_mask  string `json:"INIT_MASK"`
    Valid_mask string `json:"VALID_MASK"`
    Line_type  string `json:"LINE_TYPE"`
}


//fillHeader functions
func (s *STAT_ORANK) fill_STAT_ORANK_Header(fields []string, doc *map[string]interface{}){
	// fill the met fields leaving out "" and NA values
	if fields[0] != ""  && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != ""  && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != ""  && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != ""  && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != ""  && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != ""  && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != ""  && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != ""  && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != ""  && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != ""  && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != ""  && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != ""  && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != ""  && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_ECNT) fill_STAT_ECNT_Header(fields []string, doc *map[string]interface{}){
	// fill the met fields leaving out "" and NA values
	if fields[0] != ""  && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != ""  && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != ""  && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != ""  && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != ""  && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != ""  && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != ""  && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != ""  && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != ""  && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != ""  && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != ""  && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != ""  && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != ""  && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_VL1L2) fill_STAT_VL1L2_Header(fields []string, doc *map[string]interface{}){
	// fill the met fields leaving out "" and NA values
	if fields[0] != ""  && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != ""  && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != ""  && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != ""  && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != ""  && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != ""  && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != ""  && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != ""  && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != ""  && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != ""  && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != ""  && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != ""  && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != ""  && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_PCT) fill_STAT_PCT_Header(fields []string, doc *map[string]interface{}){
	// fill the met fields leaving out "" and NA values
	if fields[0] != ""  && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != ""  && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != ""  && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != ""  && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != ""  && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != ""  && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != ""  && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != ""  && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != ""  && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != ""  && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != ""  && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != ""  && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != ""  && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_PRC) fill_STAT_PRC_Header(fields []string, doc *map[string]interface{}){
	// fill the met fields leaving out "" and NA values
	if fields[0] != ""  && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != ""  && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != ""  && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != ""  && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != ""  && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != ""  && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != ""  && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != ""  && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != ""  && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != ""  && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != ""  && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != ""  && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != ""  && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_SL1L2) fill_STAT_SL1L2_Header(fields []string, doc *map[string]interface{}){
	// fill the met fields leaving out "" and NA values
	if fields[0] != ""  && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != ""  && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != ""  && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != ""  && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != ""  && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != ""  && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != ""  && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != ""  && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != ""  && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != ""  && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != ""  && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != ""  && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != ""  && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_MCTC) fill_STAT_MCTC_Header(fields []string, doc *map[string]interface{}){
	// fill the met fields leaving out "" and NA values
	if fields[0] != ""  && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != ""  && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != ""  && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != ""  && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != ""  && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != ""  && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != ""  && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != ""  && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != ""  && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != ""  && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != ""  && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != ""  && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != ""  && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_SEEPS_MPR) fill_STAT_SEEPS_MPR_Header(fields []string, doc *map[string]interface{}){
	// fill the met fields leaving out "" and NA values
	if fields[0] != ""  && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != ""  && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != ""  && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != ""  && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != ""  && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != ""  && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != ""  && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != ""  && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != ""  && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != ""  && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != ""  && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != ""  && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != ""  && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_NBRCNT) fill_STAT_NBRCNT_Header(fields []string, doc *map[string]interface{}){
	// fill the met fields leaving out "" and NA values
	if fields[0] != ""  && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != ""  && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != ""  && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != ""  && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != ""  && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != ""  && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != ""  && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != ""  && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != ""  && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != ""  && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != ""  && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != ""  && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != ""  && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_RELP) fill_STAT_RELP_Header(fields []string, doc *map[string]interface{}){
	// fill the met fields leaving out "" and NA values
	if fields[0] != ""  && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != ""  && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != ""  && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != ""  && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != ""  && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != ""  && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != ""  && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != ""  && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != ""  && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != ""  && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != ""  && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != ""  && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != ""  && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_GENMPR) fill_STAT_GENMPR_Header(fields []string, doc *map[string]interface{}){
	// fill the met fields leaving out "" and NA values
	if fields[0] != ""  && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != ""  && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != ""  && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != ""  && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != ""  && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != ""  && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != ""  && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != ""  && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != ""  && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != ""  && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != ""  && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != ""  && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != ""  && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *MTD_3DPAIR) fill_MTD_3DPAIR_Header(fields []string, doc *map[string]interface{}){
	// fill the met fields leaving out "" and NA values
	if fields[0] != ""  && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != ""  && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != ""  && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID"] = fields[4]
	}
	if fields[5] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[5]
	}
	if fields[6] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID"] = fields[6]
	}
	if fields[7] != ""  && fields[0] != "NA" {
		(*doc)["T_DELTA"] = fields[7]
	}
	if fields[8] != ""  && fields[0] != "NA" {
		(*doc)["FCST_T_BEG"] = fields[8]
	}
	if fields[9] != ""  && fields[0] != "NA" {
		(*doc)["FCST_T_END"] = fields[9]
	}
	if fields[10] != ""  && fields[0] != "NA" {
		(*doc)["FCST_RAD"] = fields[10]
	}
	if fields[11] != ""  && fields[0] != "NA" {
		(*doc)["FCST_THR"] = fields[11]
	}
	if fields[12] != ""  && fields[0] != "NA" {
		(*doc)["OBS_T_BEG"] = fields[12]
	}
	if fields[13] != ""  && fields[0] != "NA" {
		(*doc)["OBS_T_END"] = fields[13]
	}
	if fields[14] != ""  && fields[0] != "NA" {
		(*doc)["OBS_RAD"] = fields[14]
	}
	if fields[15] != ""  && fields[0] != "NA" {
		(*doc)["OBS_THR"] = fields[15]
	}
	if fields[16] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[16]
	}
	if fields[17] != ""  && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[17]
	}
	if fields[19] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[19]
	}
	if fields[20] != ""  && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[20]
	}
	if fields[21] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[21]
	}
}

func (s *TCST_TCDIAG) fill_TCST_TCDIAG_Header(fields []string, doc *map[string]interface{}){
	// fill the met fields leaving out "" and NA values
	if fields[0] != ""  && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != ""  && fields[0] != "NA" {
		(*doc)["AMODEL"] = fields[1]
	}
	if fields[2] != ""  && fields[0] != "NA" {
		(*doc)["BMODEL"] = fields[2]
	}
	if fields[3] != ""  && fields[0] != "NA" {
		(*doc)["DESC"] = fields[3]
	}
	if fields[4] != ""  && fields[0] != "NA" {
		(*doc)["STORM_ID"] = fields[4]
	}
	if fields[5] != ""  && fields[0] != "NA" {
		(*doc)["BASIN"] = fields[5]
	}
	if fields[6] != ""  && fields[0] != "NA" {
		(*doc)["CYCLONE"] = fields[6]
	}
	if fields[7] != ""  && fields[0] != "NA" {
		(*doc)["STORM_NAME"] = fields[7]
	}
	if fields[8] != ""  && fields[0] != "NA" {
		(*doc)["INIT"] = fields[8]
	}
	if fields[9] != ""  && fields[0] != "NA" {
		(*doc)["LEAD"] = fields[9]
	}
	if fields[10] != ""  && fields[0] != "NA" {
		(*doc)["VALID"] = fields[10]
	}
	if fields[11] != ""  && fields[0] != "NA" {
		(*doc)["INIT_MASK"] = fields[11]
	}
	if fields[12] != ""  && fields[0] != "NA" {
		(*doc)["VALID_MASK"] = fields[12]
	}
	if fields[13] != ""  && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[13]
	}
}

func (s *STAT_CTS) fill_STAT_CTS_Header(fields []string, doc *map[string]interface{}){
	// fill the met fields leaving out "" and NA values
	if fields[0] != ""  && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != ""  && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != ""  && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != ""  && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != ""  && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != ""  && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != ""  && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != ""  && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != ""  && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != ""  && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != ""  && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != ""  && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != ""  && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_FHO) fill_STAT_FHO_Header(fields []string, doc *map[string]interface{}){
	// fill the met fields leaving out "" and NA values
	if fields[0] != ""  && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != ""  && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != ""  && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != ""  && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != ""  && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != ""  && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != ""  && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != ""  && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != ""  && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != ""  && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != ""  && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != ""  && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != ""  && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_ECLV) fill_STAT_ECLV_Header(fields []string, doc *map[string]interface{}){
	// fill the met fields leaving out "" and NA values
	if fields[0] != ""  && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != ""  && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != ""  && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != ""  && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != ""  && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != ""  && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != ""  && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != ""  && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != ""  && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != ""  && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != ""  && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != ""  && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != ""  && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *MTD_3DSINGLE) fill_MTD_3DSINGLE_Header(fields []string, doc *map[string]interface{}){
	// fill the met fields leaving out "" and NA values
	if fields[0] != ""  && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != ""  && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != ""  && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID"] = fields[4]
	}
	if fields[5] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[5]
	}
	if fields[6] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID"] = fields[6]
	}
	if fields[7] != ""  && fields[0] != "NA" {
		(*doc)["T_DELTA"] = fields[7]
	}
	if fields[8] != ""  && fields[0] != "NA" {
		(*doc)["FCST_T_BEG"] = fields[8]
	}
	if fields[9] != ""  && fields[0] != "NA" {
		(*doc)["FCST_T_END"] = fields[9]
	}
	if fields[10] != ""  && fields[0] != "NA" {
		(*doc)["FCST_RAD"] = fields[10]
	}
	if fields[11] != ""  && fields[0] != "NA" {
		(*doc)["FCST_THR"] = fields[11]
	}
	if fields[12] != ""  && fields[0] != "NA" {
		(*doc)["OBS_T_BEG"] = fields[12]
	}
	if fields[13] != ""  && fields[0] != "NA" {
		(*doc)["OBS_T_END"] = fields[13]
	}
	if fields[14] != ""  && fields[0] != "NA" {
		(*doc)["OBS_RAD"] = fields[14]
	}
	if fields[15] != ""  && fields[0] != "NA" {
		(*doc)["OBS_THR"] = fields[15]
	}
	if fields[16] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[16]
	}
	if fields[17] != ""  && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[17]
	}
	if fields[19] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[19]
	}
	if fields[20] != ""  && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[20]
	}
	if fields[21] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[21]
	}
}

func (s *STAT_MCTS) fill_STAT_MCTS_Header(fields []string, doc *map[string]interface{}){
	// fill the met fields leaving out "" and NA values
	if fields[0] != ""  && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != ""  && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != ""  && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != ""  && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != ""  && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != ""  && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != ""  && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != ""  && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != ""  && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != ""  && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != ""  && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != ""  && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != ""  && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_DMAP) fill_STAT_DMAP_Header(fields []string, doc *map[string]interface{}){
	// fill the met fields leaving out "" and NA values
	if fields[0] != ""  && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != ""  && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != ""  && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != ""  && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != ""  && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != ""  && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != ""  && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != ""  && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != ""  && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != ""  && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != ""  && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != ""  && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != ""  && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_PHIST) fill_STAT_PHIST_Header(fields []string, doc *map[string]interface{}){
	// fill the met fields leaving out "" and NA values
	if fields[0] != ""  && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != ""  && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != ""  && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != ""  && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != ""  && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != ""  && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != ""  && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != ""  && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != ""  && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != ""  && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != ""  && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != ""  && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != ""  && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_NBRCTC) fill_STAT_NBRCTC_Header(fields []string, doc *map[string]interface{}){
	// fill the met fields leaving out "" and NA values
	if fields[0] != ""  && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != ""  && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != ""  && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != ""  && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != ""  && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != ""  && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != ""  && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != ""  && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != ""  && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != ""  && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != ""  && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != ""  && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != ""  && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_NBRCTS) fill_STAT_NBRCTS_Header(fields []string, doc *map[string]interface{}){
	// fill the met fields leaving out "" and NA values
	if fields[0] != ""  && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != ""  && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != ""  && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != ""  && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != ""  && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != ""  && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != ""  && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != ""  && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != ""  && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != ""  && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != ""  && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != ""  && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != ""  && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_GRAD) fill_STAT_GRAD_Header(fields []string, doc *map[string]interface{}){
	// fill the met fields leaving out "" and NA values
	if fields[0] != ""  && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != ""  && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != ""  && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != ""  && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != ""  && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != ""  && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != ""  && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != ""  && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != ""  && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != ""  && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != ""  && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != ""  && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != ""  && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_RPS) fill_STAT_RPS_Header(fields []string, doc *map[string]interface{}){
	// fill the met fields leaving out "" and NA values
	if fields[0] != ""  && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != ""  && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != ""  && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != ""  && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != ""  && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != ""  && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != ""  && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != ""  && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != ""  && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != ""  && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != ""  && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != ""  && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != ""  && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_RHIST) fill_STAT_RHIST_Header(fields []string, doc *map[string]interface{}){
	// fill the met fields leaving out "" and NA values
	if fields[0] != ""  && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != ""  && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != ""  && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != ""  && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != ""  && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != ""  && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != ""  && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != ""  && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != ""  && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != ""  && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != ""  && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != ""  && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != ""  && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_ISC) fill_STAT_ISC_Header(fields []string, doc *map[string]interface{}){
	// fill the met fields leaving out "" and NA values
	if fields[0] != ""  && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != ""  && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != ""  && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != ""  && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != ""  && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != ""  && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != ""  && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != ""  && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != ""  && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != ""  && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != ""  && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != ""  && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != ""  && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_MPR) fill_STAT_MPR_Header(fields []string, doc *map[string]interface{}){
	// fill the met fields leaving out "" and NA values
	if fields[0] != ""  && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != ""  && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != ""  && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != ""  && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != ""  && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != ""  && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != ""  && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != ""  && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != ""  && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != ""  && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != ""  && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != ""  && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != ""  && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_SEEPS) fill_STAT_SEEPS_Header(fields []string, doc *map[string]interface{}){
	// fill the met fields leaving out "" and NA values
	if fields[0] != ""  && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != ""  && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != ""  && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != ""  && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != ""  && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != ""  && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != ""  && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != ""  && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != ""  && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != ""  && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != ""  && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != ""  && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != ""  && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_VAL1L2) fill_STAT_VAL1L2_Header(fields []string, doc *map[string]interface{}){
	// fill the met fields leaving out "" and NA values
	if fields[0] != ""  && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != ""  && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != ""  && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != ""  && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != ""  && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != ""  && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != ""  && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != ""  && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != ""  && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != ""  && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != ""  && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != ""  && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != ""  && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *TCST_PROBRIRW) fill_TCST_PROBRIRW_Header(fields []string, doc *map[string]interface{}){
	// fill the met fields leaving out "" and NA values
	if fields[0] != ""  && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != ""  && fields[0] != "NA" {
		(*doc)["AMODEL"] = fields[1]
	}
	if fields[2] != ""  && fields[0] != "NA" {
		(*doc)["BMODEL"] = fields[2]
	}
	if fields[3] != ""  && fields[0] != "NA" {
		(*doc)["DESC"] = fields[3]
	}
	if fields[4] != ""  && fields[0] != "NA" {
		(*doc)["STORM_ID"] = fields[4]
	}
	if fields[5] != ""  && fields[0] != "NA" {
		(*doc)["BASIN"] = fields[5]
	}
	if fields[6] != ""  && fields[0] != "NA" {
		(*doc)["CYCLONE"] = fields[6]
	}
	if fields[7] != ""  && fields[0] != "NA" {
		(*doc)["STORM_NAME"] = fields[7]
	}
	if fields[8] != ""  && fields[0] != "NA" {
		(*doc)["INIT"] = fields[8]
	}
	if fields[9] != ""  && fields[0] != "NA" {
		(*doc)["LEAD"] = fields[9]
	}
	if fields[10] != ""  && fields[0] != "NA" {
		(*doc)["VALID"] = fields[10]
	}
	if fields[11] != ""  && fields[0] != "NA" {
		(*doc)["INIT_MASK"] = fields[11]
	}
	if fields[12] != ""  && fields[0] != "NA" {
		(*doc)["VALID_MASK"] = fields[12]
	}
	if fields[13] != ""  && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[13]
	}
}

func (s *STAT_PJC) fill_STAT_PJC_Header(fields []string, doc *map[string]interface{}){
	// fill the met fields leaving out "" and NA values
	if fields[0] != ""  && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != ""  && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != ""  && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != ""  && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != ""  && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != ""  && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != ""  && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != ""  && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != ""  && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != ""  && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != ""  && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != ""  && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != ""  && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_SAL1L2) fill_STAT_SAL1L2_Header(fields []string, doc *map[string]interface{}){
	// fill the met fields leaving out "" and NA values
	if fields[0] != ""  && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != ""  && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != ""  && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != ""  && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != ""  && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != ""  && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != ""  && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != ""  && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != ""  && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != ""  && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != ""  && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != ""  && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != ""  && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *MODE_CTS) fill_MODE_CTS_Header(fields []string, doc *map[string]interface{}){
	// fill the met fields leaving out "" and NA values
	if fields[0] != ""  && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != ""  && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != ""  && fields[0] != "NA" {
		(*doc)["N_VALID"] = fields[2]
	}
	if fields[3] != ""  && fields[0] != "NA" {
		(*doc)["GRID_RES"] = fields[3]
	}
	if fields[4] != ""  && fields[0] != "NA" {
		(*doc)["DESC"] = fields[4]
	}
	if fields[6] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID"] = fields[6]
	}
	if fields[7] != ""  && fields[0] != "NA" {
		(*doc)["FCST_ACCUM"] = fields[7]
	}
	if fields[8] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[8]
	}
	if fields[9] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID"] = fields[9]
	}
	if fields[10] != ""  && fields[0] != "NA" {
		(*doc)["OBS_ACCUM"] = fields[10]
	}
	if fields[11] != ""  && fields[0] != "NA" {
		(*doc)["FCST_RAD"] = fields[11]
	}
	if fields[12] != ""  && fields[0] != "NA" {
		(*doc)["FCST_THR"] = fields[12]
	}
	if fields[13] != ""  && fields[0] != "NA" {
		(*doc)["OBS_RAD"] = fields[13]
	}
	if fields[14] != ""  && fields[0] != "NA" {
		(*doc)["OBS_THR"] = fields[14]
	}
	if fields[15] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[15]
	}
	if fields[16] != ""  && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[16]
	}
	if fields[18] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[18]
	}
	if fields[19] != ""  && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[19]
	}
	if fields[20] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[20]
	}
	if fields[21] != ""  && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[21]
	}
	if fields[22] != ""  && fields[0] != "NA" {
		(*doc)["FIELD"] = fields[22]
	}
	if fields[23] != ""  && fields[0] != "NA" {
		(*doc)["TOTAL"] = fields[23]
	}
	if fields[24] != ""  && fields[0] != "NA" {
		(*doc)["FY_OY"] = fields[24]
	}
	if fields[25] != ""  && fields[0] != "NA" {
		(*doc)["FY_ON"] = fields[25]
	}
	if fields[26] != ""  && fields[0] != "NA" {
		(*doc)["FN_OY"] = fields[26]
	}
	if fields[27] != ""  && fields[0] != "NA" {
		(*doc)["FN_ON"] = fields[27]
	}
	if fields[28] != ""  && fields[0] != "NA" {
		(*doc)["BASER"] = fields[28]
	}
	if fields[29] != ""  && fields[0] != "NA" {
		(*doc)["FMEAN"] = fields[29]
	}
	if fields[30] != ""  && fields[0] != "NA" {
		(*doc)["ACC"] = fields[30]
	}
	if fields[31] != ""  && fields[0] != "NA" {
		(*doc)["FBIAS"] = fields[31]
	}
	if fields[32] != ""  && fields[0] != "NA" {
		(*doc)["PODY"] = fields[32]
	}
	if fields[33] != ""  && fields[0] != "NA" {
		(*doc)["PODN"] = fields[33]
	}
	if fields[34] != ""  && fields[0] != "NA" {
		(*doc)["POFD"] = fields[34]
	}
	if fields[35] != ""  && fields[0] != "NA" {
		(*doc)["FAR"] = fields[35]
	}
	if fields[36] != ""  && fields[0] != "NA" {
		(*doc)["CSI"] = fields[36]
	}
	if fields[37] != ""  && fields[0] != "NA" {
		(*doc)["GSS"] = fields[37]
	}
	if fields[38] != ""  && fields[0] != "NA" {
		(*doc)["HK"] = fields[38]
	}
	if fields[39] != ""  && fields[0] != "NA" {
		(*doc)["HSS"] = fields[39]
	}
	if fields[40] != ""  && fields[0] != "NA" {
		(*doc)["ODDS"] = fields[40]
	}
	if fields[41] != ""  && fields[0] != "NA" {
		(*doc)["LODDS"] = fields[41]
	}
	if fields[42] != ""  && fields[0] != "NA" {
		(*doc)["ORSS"] = fields[42]
	}
	if fields[43] != ""  && fields[0] != "NA" {
		(*doc)["EDS"] = fields[43]
	}
	if fields[44] != ""  && fields[0] != "NA" {
		(*doc)["SEDS"] = fields[44]
	}
	if fields[45] != ""  && fields[0] != "NA" {
		(*doc)["EDI"] = fields[45]
	}
	if fields[46] != ""  && fields[0] != "NA" {
		(*doc)["SEDI"] = fields[46]
	}
	if fields[47] != ""  && fields[0] != "NA" {
		(*doc)["BAGSS"] = fields[47]
	}
}

func (s *STAT_CTC) fill_STAT_CTC_Header(fields []string, doc *map[string]interface{}){
	// fill the met fields leaving out "" and NA values
	if fields[0] != ""  && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != ""  && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != ""  && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != ""  && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != ""  && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != ""  && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != ""  && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != ""  && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != ""  && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != ""  && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != ""  && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != ""  && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != ""  && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *MODE_OBJ) fill_MODE_OBJ_Header(fields []string, doc *map[string]interface{}){
	// fill the met fields leaving out "" and NA values
	if fields[0] != ""  && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != ""  && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != ""  && fields[0] != "NA" {
		(*doc)["N_VALID"] = fields[2]
	}
	if fields[3] != ""  && fields[0] != "NA" {
		(*doc)["GRID_RES"] = fields[3]
	}
	if fields[4] != ""  && fields[0] != "NA" {
		(*doc)["DESC"] = fields[4]
	}
	if fields[6] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID"] = fields[6]
	}
	if fields[7] != ""  && fields[0] != "NA" {
		(*doc)["FCST_ACCUM"] = fields[7]
	}
	if fields[8] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[8]
	}
	if fields[9] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID"] = fields[9]
	}
	if fields[10] != ""  && fields[0] != "NA" {
		(*doc)["OBS_ACCUM"] = fields[10]
	}
	if fields[11] != ""  && fields[0] != "NA" {
		(*doc)["FCST_RAD"] = fields[11]
	}
	if fields[12] != ""  && fields[0] != "NA" {
		(*doc)["FCST_THR"] = fields[12]
	}
	if fields[13] != ""  && fields[0] != "NA" {
		(*doc)["OBS_RAD"] = fields[13]
	}
	if fields[14] != ""  && fields[0] != "NA" {
		(*doc)["OBS_THR"] = fields[14]
	}
	if fields[15] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[15]
	}
	if fields[16] != ""  && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[16]
	}
	if fields[18] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[18]
	}
	if fields[19] != ""  && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[19]
	}
	if fields[20] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[20]
	}
	if fields[21] != ""  && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[21]
	}
	if fields[22] != ""  && fields[0] != "NA" {
		(*doc)["OBJECT_ID"] = fields[22]
	}
	if fields[23] != ""  && fields[0] != "NA" {
		(*doc)["OBJECT_CAT"] = fields[23]
	}
	if fields[24] != ""  && fields[0] != "NA" {
		(*doc)["CENTROID_X"] = fields[24]
	}
	if fields[25] != ""  && fields[0] != "NA" {
		(*doc)["CENTROID_Y"] = fields[25]
	}
	if fields[26] != ""  && fields[0] != "NA" {
		(*doc)["CENTROID_LAT"] = fields[26]
	}
	if fields[27] != ""  && fields[0] != "NA" {
		(*doc)["CENTROID_LON"] = fields[27]
	}
	if fields[28] != ""  && fields[0] != "NA" {
		(*doc)["AXIS_ANG"] = fields[28]
	}
	if fields[29] != ""  && fields[0] != "NA" {
		(*doc)["LENGTH"] = fields[29]
	}
	if fields[30] != ""  && fields[0] != "NA" {
		(*doc)["WIDTH"] = fields[30]
	}
	if fields[31] != ""  && fields[0] != "NA" {
		(*doc)["AREA"] = fields[31]
	}
	if fields[32] != ""  && fields[0] != "NA" {
		(*doc)["AREA_THRESH"] = fields[32]
	}
	if fields[33] != ""  && fields[0] != "NA" {
		(*doc)["CURVATURE"] = fields[33]
	}
	if fields[34] != ""  && fields[0] != "NA" {
		(*doc)["CURVATURE_X"] = fields[34]
	}
	if fields[35] != ""  && fields[0] != "NA" {
		(*doc)["CURVATURE_Y"] = fields[35]
	}
	if fields[36] != ""  && fields[0] != "NA" {
		(*doc)["COMPLEXITY"] = fields[36]
	}
	if fields[37] != ""  && fields[0] != "NA" {
		(*doc)["INTENSITY_10"] = fields[37]
	}
	if fields[38] != ""  && fields[0] != "NA" {
		(*doc)["INTENSITY_25"] = fields[38]
	}
	if fields[39] != ""  && fields[0] != "NA" {
		(*doc)["INTENSITY_50"] = fields[39]
	}
	if fields[40] != ""  && fields[0] != "NA" {
		(*doc)["INTENSITY_75"] = fields[40]
	}
	if fields[41] != ""  && fields[0] != "NA" {
		(*doc)["INTENSITY_90"] = fields[41]
	}
	if fields[42] != ""  && fields[0] != "NA" {
		(*doc)["INTENSITY_USER"] = fields[42]
	}
	if fields[43] != ""  && fields[0] != "NA" {
		(*doc)["INTENSITY_SUM"] = fields[43]
	}
	if fields[44] != ""  && fields[0] != "NA" {
		(*doc)["CENTROID_DIST"] = fields[44]
	}
	if fields[45] != ""  && fields[0] != "NA" {
		(*doc)["BOUNDARY_DIST"] = fields[45]
	}
	if fields[46] != ""  && fields[0] != "NA" {
		(*doc)["CONVEX_HULL_DIST"] = fields[46]
	}
	if fields[47] != ""  && fields[0] != "NA" {
		(*doc)["ANGLE_DIFF"] = fields[47]
	}
	if fields[48] != ""  && fields[0] != "NA" {
		(*doc)["ASPECT_DIFF"] = fields[48]
	}
	if fields[49] != ""  && fields[0] != "NA" {
		(*doc)["AREA_RATIO"] = fields[49]
	}
	if fields[50] != ""  && fields[0] != "NA" {
		(*doc)["INTERSECTION_AREA"] = fields[50]
	}
	if fields[51] != ""  && fields[0] != "NA" {
		(*doc)["UNION_AREA"] = fields[51]
	}
	if fields[52] != ""  && fields[0] != "NA" {
		(*doc)["SYMMETRIC_DIFF"] = fields[52]
	}
	if fields[53] != ""  && fields[0] != "NA" {
		(*doc)["INTERSECTION_OVER_AREA"] = fields[53]
	}
	if fields[54] != ""  && fields[0] != "NA" {
		(*doc)["CURVATURE_RATIO"] = fields[54]
	}
	if fields[55] != ""  && fields[0] != "NA" {
		(*doc)["COMPLEXITY_RATIO"] = fields[55]
	}
	if fields[56] != ""  && fields[0] != "NA" {
		(*doc)["PERCENTILE_INTENSITY_RATIO"] = fields[56]
	}
	if fields[57] != ""  && fields[0] != "NA" {
		(*doc)["INTEREST"] = fields[57]
	}
}

func (s *MTD_2DSINGLE) fill_MTD_2DSINGLE_Header(fields []string, doc *map[string]interface{}){
	// fill the met fields leaving out "" and NA values
	if fields[0] != ""  && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != ""  && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != ""  && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID"] = fields[4]
	}
	if fields[5] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[5]
	}
	if fields[6] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID"] = fields[6]
	}
	if fields[7] != ""  && fields[0] != "NA" {
		(*doc)["T_DELTA"] = fields[7]
	}
	if fields[8] != ""  && fields[0] != "NA" {
		(*doc)["FCST_T_BEG"] = fields[8]
	}
	if fields[9] != ""  && fields[0] != "NA" {
		(*doc)["FCST_T_END"] = fields[9]
	}
	if fields[10] != ""  && fields[0] != "NA" {
		(*doc)["FCST_RAD"] = fields[10]
	}
	if fields[11] != ""  && fields[0] != "NA" {
		(*doc)["FCST_THR"] = fields[11]
	}
	if fields[12] != ""  && fields[0] != "NA" {
		(*doc)["OBS_T_BEG"] = fields[12]
	}
	if fields[13] != ""  && fields[0] != "NA" {
		(*doc)["OBS_T_END"] = fields[13]
	}
	if fields[14] != ""  && fields[0] != "NA" {
		(*doc)["OBS_RAD"] = fields[14]
	}
	if fields[15] != ""  && fields[0] != "NA" {
		(*doc)["OBS_THR"] = fields[15]
	}
	if fields[16] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[16]
	}
	if fields[17] != ""  && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[17]
	}
	if fields[19] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[19]
	}
	if fields[20] != ""  && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[20]
	}
	if fields[21] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[21]
	}
}

func (s *STAT_VCNT) fill_STAT_VCNT_Header(fields []string, doc *map[string]interface{}){
	// fill the met fields leaving out "" and NA values
	if fields[0] != ""  && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != ""  && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != ""  && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != ""  && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != ""  && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != ""  && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != ""  && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != ""  && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != ""  && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != ""  && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != ""  && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != ""  && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != ""  && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_SSIDX) fill_STAT_SSIDX_Header(fields []string, doc *map[string]interface{}){
	// fill the met fields leaving out "" and NA values
	if fields[0] != ""  && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != ""  && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != ""  && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != ""  && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != ""  && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != ""  && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != ""  && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != ""  && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != ""  && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != ""  && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != ""  && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != ""  && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != ""  && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *TCST_TCMPR) fill_TCST_TCMPR_Header(fields []string, doc *map[string]interface{}){
	// fill the met fields leaving out "" and NA values
	if fields[0] != ""  && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != ""  && fields[0] != "NA" {
		(*doc)["AMODEL"] = fields[1]
	}
	if fields[2] != ""  && fields[0] != "NA" {
		(*doc)["BMODEL"] = fields[2]
	}
	if fields[3] != ""  && fields[0] != "NA" {
		(*doc)["DESC"] = fields[3]
	}
	if fields[4] != ""  && fields[0] != "NA" {
		(*doc)["STORM_ID"] = fields[4]
	}
	if fields[5] != ""  && fields[0] != "NA" {
		(*doc)["BASIN"] = fields[5]
	}
	if fields[6] != ""  && fields[0] != "NA" {
		(*doc)["CYCLONE"] = fields[6]
	}
	if fields[7] != ""  && fields[0] != "NA" {
		(*doc)["STORM_NAME"] = fields[7]
	}
	if fields[8] != ""  && fields[0] != "NA" {
		(*doc)["INIT"] = fields[8]
	}
	if fields[9] != ""  && fields[0] != "NA" {
		(*doc)["LEAD"] = fields[9]
	}
	if fields[10] != ""  && fields[0] != "NA" {
		(*doc)["VALID"] = fields[10]
	}
	if fields[11] != ""  && fields[0] != "NA" {
		(*doc)["INIT_MASK"] = fields[11]
	}
	if fields[12] != ""  && fields[0] != "NA" {
		(*doc)["VALID_MASK"] = fields[12]
	}
	if fields[13] != ""  && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[13]
	}
}

func (s *STAT_CNT) fill_STAT_CNT_Header(fields []string, doc *map[string]interface{}){
	// fill the met fields leaving out "" and NA values
	if fields[0] != ""  && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != ""  && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != ""  && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != ""  && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != ""  && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != ""  && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != ""  && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != ""  && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != ""  && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != ""  && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != ""  && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != ""  && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != ""  && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_PSTD) fill_STAT_PSTD_Header(fields []string, doc *map[string]interface{}){
	// fill the met fields leaving out "" and NA values
	if fields[0] != ""  && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != ""  && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != ""  && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != ""  && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != ""  && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != ""  && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != ""  && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != ""  && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != ""  && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != ""  && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != ""  && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != ""  && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != ""  && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_SSVAR) fill_STAT_SSVAR_Header(fields []string, doc *map[string]interface{}){
	// fill the met fields leaving out "" and NA values
	if fields[0] != ""  && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != ""  && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != ""  && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != ""  && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != ""  && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != ""  && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != ""  && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != ""  && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != ""  && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != ""  && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != ""  && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != ""  && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != ""  && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != ""  && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != ""  && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != ""  && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != ""  && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}


//line data struct definitions
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
    N_ens            int     `json:"n_ens"`
    Ens_i                    `json:"ens_i"`
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

type STAT_RHIST struct {
    Total       int     `json:"total"`
    N_rank      int     `json:"n_rank"`
    Rank_i              `json:"rank_i"`
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
    Total_dir    float64 `json:"total_dir"`
    Dira_me      float64 `json:"dira_me"`
    Dira_mae     float64 `json:"dira_mae"`
    Dira_mse     float64 `json:"dira_mse"`
}

type STAT_CTC struct {
    Total    int     `json:"total"`
    Fy_oy    float64 `json:"fy_oy"`
    Fy_on    float64 `json:"fy_on"`
    Fn_oy    float64 `json:"fn_oy"`
    Fn_on    float64 `json:"fn_on"`
    Ec_value float64 `json:"ec_value"`
}

type STAT_PCT struct {
    Total         int     `json:"total"`
    N_thresh      int     `json:"n_thresh"`
    Thresh_i              `json:"thresh_i"`
    Oy_i                  `json:"oy_i"`
    On_i                  `json:"on_i"`
}

type STAT_GRAD struct {
    Total      int     `json:"total"`
    Fgbar      float64 `json:"fgbar"`
    Ogbar      float64 `json:"ogbar"`
    Mgbar      float64 `json:"mgbar"`
    Egbar      float64 `json:"egbar"`
    S1         float64 `json:"s1"`
    S1_og      float64 `json:"s1_og"`
    Fgog_ratio float64 `json:"fgog_ratio"`
    Dx         float64 `json:"dx"`
    Dy         float64 `json:"dy"`
}

type STAT_PRC struct {
    Total         int     `json:"total"`
    N_thresh      int     `json:"n_thresh"`
    Thresh_i              `json:"thresh_i"`
    Pody_i                `json:"pody_i"`
    Pofd_i                `json:"pofd_i"`
}

type STAT_ECLV struct {
    Total        int     `json:"total"`
    Baser        float64 `json:"baser"`
    Value_baser  float64 `json:"value_baser"`
    N_pts                `json:"n_pts"`
    Cl_i                 `json:"cl_i"`
    Value_i              `json:"value_i"`
}

type MTD_3DSINGLE struct {
    Object_id       string  `json:"object_id"`
    Object_cat      string  `json:"object_cat"`
    Centroid_x      float64 `json:"centroid_x"`
    Centroid_y      float64 `json:"centroid_y"`
    Centroid_t      float64 `json:"centroid_t"`
    Centroid_lat    float64 `json:"centroid_lat"`
    Centroid_lon    float64 `json:"centroid_lon"`
    X_dot           float64 `json:"x_dot"`
    Y_dot           float64 `json:"y_dot"`
    Axis_ang        float64 `json:"axis_ang"`
    Volume          int     `json:"volume"`
    Start_time      int     `json:"start_time"`
    End_time        int     `json:"end_time"`
    Cdist_travelled float64 `json:"cdist_travelled"`
    Intensity_10    float64 `json:"intensity_10"`
    Intensity_25    float64 `json:"intensity_25"`
    Intensity_50    float64 `json:"intensity_50"`
    Intensity_75    float64 `json:"intensity_75"`
    Intensity_90    float64 `json:"intensity_90"`
    Intensity_user  string  `json:"intensity_user"`
}

type TCST_PROBRIRW struct {
    Alat          float64 `json:"alat"`
    Alon          float64 `json:"alon"`
    Blat          float64 `json:"blat"`
    Blon          float64 `json:"blon"`
    Initials      string  `json:"initials"`
    Tk_err        float64 `json:"tk_err"`
    X_err         float64 `json:"x_err"`
    Y_err         float64 `json:"y_err"`
    Adland        float64 `json:"adland"`
    Bdland        float64 `json:"bdland"`
    Rirw_beg      int     `json:"rirw_beg"`
    Rirw_end      int     `json:"rirw_end"`
    Rirw_window   float64 `json:"rirw_window"`
    Awind_end     float64 `json:"awind_end"`
    Bwind_beg     float64 `json:"bwind_beg"`
    Bwind_end     float64 `json:"bwind_end"`
    Bdelta        float64 `json:"bdelta"`
    Bdelta_max    float64 `json:"bdelta_max"`
    Blevel_beg    string  `json:"blevel_beg"`
    Blevel_end    string  `json:"blevel_end"`
    N_thresh      int     `json:"n_thresh"`
    Thresh_i              `json:"thresh_i"`
    Prob_i                `json:"prob_i"`
}

type STAT_CNT struct {
    Total                int     `json:"total"`
    Fbar                 float64 `json:"fbar"`
    Fbar_ncl             float64 `json:"fbar_ncl"`
    Fbar_ncu             float64 `json:"fbar_ncu"`
    Fbar_bcl             float64 `json:"fbar_bcl"`
    Fbar_bcu             float64 `json:"fbar_bcu"`
    Fstdev               float64 `json:"fstdev"`
    Fstdev_ncl           float64 `json:"fstdev_ncl"`
    Fstdev_ncu           float64 `json:"fstdev_ncu"`
    Fstdev_bcl           float64 `json:"fstdev_bcl"`
    Fstdev_bcu           float64 `json:"fstdev_bcu"`
    Obar                 float64 `json:"obar"`
    Obar_ncl             float64 `json:"obar_ncl"`
    Obar_ncu             float64 `json:"obar_ncu"`
    Obar_bcl             float64 `json:"obar_bcl"`
    Obar_bcu             float64 `json:"obar_bcu"`
    Ostdev               float64 `json:"ostdev"`
    Ostdev_ncl           float64 `json:"ostdev_ncl"`
    Ostdev_ncu           float64 `json:"ostdev_ncu"`
    Ostdev_bcl           float64 `json:"ostdev_bcl"`
    Ostdev_bcu           float64 `json:"ostdev_bcu"`
    Pr_corr              float64 `json:"pr_corr"`
    Pr_corr_ncl          float64 `json:"pr_corr_ncl"`
    Pr_corr_ncu          float64 `json:"pr_corr_ncu"`
    Pr_corr_bcl          float64 `json:"pr_corr_bcl"`
    Pr_corr_bcu          float64 `json:"pr_corr_bcu"`
    Sp_corr              float64 `json:"sp_corr"`
    Kt_corr              float64 `json:"kt_corr"`
    Ranks                int     `json:"ranks"`
    Frank_ties           int     `json:"frank_ties"`
    Orank_ties           int     `json:"orank_ties"`
    Me                   float64 `json:"me"`
    Me_ncl               float64 `json:"me_ncl"`
    Me_ncu               float64 `json:"me_ncu"`
    Me_bcl               float64 `json:"me_bcl"`
    Me_bcu               float64 `json:"me_bcu"`
    Estdev               float64 `json:"estdev"`
    Estdev_ncl           float64 `json:"estdev_ncl"`
    Estdev_ncu           float64 `json:"estdev_ncu"`
    Estdev_bcl           float64 `json:"estdev_bcl"`
    Estdev_bcu           float64 `json:"estdev_bcu"`
    Mbias                float64 `json:"mbias"`
    Mbias_bcl            float64 `json:"mbias_bcl"`
    Mbias_bcu            float64 `json:"mbias_bcu"`
    Mae                  float64 `json:"mae"`
    Mae_bcl              float64 `json:"mae_bcl"`
    Mae_bcu              float64 `json:"mae_bcu"`
    Mse                  float64 `json:"mse"`
    Mse_bcl              float64 `json:"mse_bcl"`
    Mse_bcu              float64 `json:"mse_bcu"`
    Bcmse                float64 `json:"bcmse"`
    Bcmse_bcl            float64 `json:"bcmse_bcl"`
    Bcmse_bcu            float64 `json:"bcmse_bcu"`
    Rmse                 float64 `json:"rmse"`
    Rmse_bcl             float64 `json:"rmse_bcl"`
    Rmse_bcu             float64 `json:"rmse_bcu"`
    E10                  float64 `json:"e10"`
    E10_bcl              float64 `json:"e10_bcl"`
    E10_bcu              float64 `json:"e10_bcu"`
    E25                  float64 `json:"e25"`
    E25_bcl              float64 `json:"e25_bcl"`
    E25_bcu              float64 `json:"e25_bcu"`
    E50                  float64 `json:"e50"`
    E50_bcl              float64 `json:"e50_bcl"`
    E50_bcu              float64 `json:"e50_bcu"`
    E75                  float64 `json:"e75"`
    E75_bcl              float64 `json:"e75_bcl"`
    E75_bcu              float64 `json:"e75_bcu"`
    E90                  float64 `json:"e90"`
    E90_bcl              float64 `json:"e90_bcl"`
    E90_bcu              float64 `json:"e90_bcu"`
    Eiqr                 float64 `json:"eiqr"`
    Eiqr_bcl             int     `json:"eiqr_bcl"`
    Eiqr_bcu             int     `json:"eiqr_bcu"`
    Mad                  float64 `json:"mad"`
    Mad_bcl              float64 `json:"mad_bcl"`
    Mad_bcu              float64 `json:"mad_bcu"`
    Anom_corr            float64 `json:"anom_corr"`
    Anom_corr_ncl        float64 `json:"anom_corr_ncl"`
    Anom_corr_ncu        float64 `json:"anom_corr_ncu"`
    Anom_corr_bcl        float64 `json:"anom_corr_bcl"`
    Anom_corr_bcu        float64 `json:"anom_corr_bcu"`
    Me2                  float64 `json:"me2"`
    Me2_bcl              float64 `json:"me2_bcl"`
    Me2_bcu              float64 `json:"me2_bcu"`
    Msess                float64 `json:"msess"`
    Msess_bcl            float64 `json:"msess_bcl"`
    Msess_bcu            float64 `json:"msess_bcu"`
    Rmsfa                float64 `json:"rmsfa"`
    Rmsfa_bcl            float64 `json:"rmsfa_bcl"`
    Rmsfa_bcu            float64 `json:"rmsfa_bcu"`
    Rmsoa                float64 `json:"rmsoa"`
    Rmsoa_bcl            float64 `json:"rmsoa_bcl"`
    Rmsoa_bcu            float64 `json:"rmsoa_bcu"`
    Anom_corr_uncntr     float64 `json:"anom_corr_uncntr"`
    Anom_corr_uncntr_bcl float64 `json:"anom_corr_uncntr_bcl"`
    Anom_corr_uncntr_bcu float64 `json:"anom_corr_uncntr_bcu"`
    Si                   float64 `json:"si"`
    Si_bcl               float64 `json:"si_bcl"`
    Si_bcu               float64 `json:"si_bcu"`
}

type STAT_MCTC struct {
    Total           int     `json:"total"`
    N_cat           int     `json:"n_cat"`
    Fi_oi                   `json:"fi_oi"`
    Ec_value        float64 `json:"ec_value"`
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

type STAT_DMAP struct {
    Total      int     `json:"total"`
    Fy         int     `json:"fy"`
    Oy         int     `json:"oy"`
    Fbias      float64 `json:"fbias"`
    Baddeley   float64 `json:"baddeley"`
    Hausdorff  float64 `json:"hausdorff"`
    Med_fo     float64 `json:"med_fo"`
    Med_of     float64 `json:"med_of"`
    Med_min    float64 `json:"med_min"`
    Med_max    float64 `json:"med_max"`
    Med_mean   float64 `json:"med_mean"`
    Fom_fo     float64 `json:"fom_fo"`
    Fom_of     float64 `json:"fom_of"`
    Fom_min    float64 `json:"fom_min"`
    Fom_max    float64 `json:"fom_max"`
    Fom_mean   float64 `json:"fom_mean"`
    Zhu_fo     float64 `json:"zhu_fo"`
    Zhu_of     float64 `json:"zhu_of"`
    Zhu_min    float64 `json:"zhu_min"`
    Zhu_max    float64 `json:"zhu_max"`
    Zhu_mean   float64 `json:"zhu_mean"`
    G          float64 `json:"g"`
    Gbeta      float64 `json:"gbeta"`
    Beta_value float64 `json:"beta_value"`
}

type STAT_GENMPR struct {
    Total      int     `json:"total"`
    Index      int     `json:"index"`
    Storm_id   string  `json:"storm_id"`
    Prob_lead  string  `json:"prob_lead"`
    Prob_val   float64 `json:"prob_val"`
    Agen_init  string  `json:"agen_init"`
    Agen_fhr   string  `json:"agen_fhr"`
    Agen_lat   float64 `json:"agen_lat"`
    Agen_lon   float64 `json:"agen_lon"`
    Agen_dland float64 `json:"agen_dland"`
    Bgen_lat   float64 `json:"bgen_lat"`
    Bgen_lon   float64 `json:"bgen_lon"`
    Bgen_dland float64 `json:"bgen_dland"`
    Gen_dist   float64 `json:"gen_dist"`
    Gen_tdiff  string  `json:"gen_tdiff"`
    Init_tdiff string  `json:"init_tdiff"`
    Dev_cat    string  `json:"dev_cat"`
    Ops_cat    string  `json:"ops_cat"`
}

type STAT_SSIDX struct {
    Fcst_model string  `json:"fcst_model"`
    Ref_model  string  `json:"ref_model"`
    N_init     int     `json:"n_init"`
    N_term     int     `json:"n_term"`
    N_vld      int     `json:"n_vld"`
    Ss_index   float64 `json:"ss_index"`
}

type MTD_2DSINGLE struct {
    Object_id      string  `json:"object_id"`
    Object_cat     string  `json:"object_cat"`
    Time_index     int     `json:"time_index"`
    Area           int     `json:"area"`
    Centroid_x     float64 `json:"centroid_x"`
    Centroid_y     float64 `json:"centroid_y"`
    Centroid_lat   float64 `json:"centroid_lat"`
    Centroid_lon   float64 `json:"centroid_lon"`
    Axis_ang       float64 `json:"axis_ang"`
    Intensity_10   float64 `json:"intensity_10"`
    Intensity_25   float64 `json:"intensity_25"`
    Intensity_50   float64 `json:"intensity_50"`
    Intensity_75   float64 `json:"intensity_75"`
    Intensity_90   float64 `json:"intensity_90"`
    Intensity_user string  `json:"intensity_user"`
}

type STAT_CTS struct {
    Total      int     `json:"total"`
    Baser      float64 `json:"baser"`
    Baser_ncl  float64 `json:"baser_ncl"`
    Baser_ncu  float64 `json:"baser_ncu"`
    Baser_bcl  float64 `json:"baser_bcl"`
    Baser_bcu  float64 `json:"baser_bcu"`
    Fmean      float64 `json:"fmean"`
    Fmean_ncl  float64 `json:"fmean_ncl"`
    Fmean_ncu  float64 `json:"fmean_ncu"`
    Fmean_bcl  float64 `json:"fmean_bcl"`
    Fmean_bcu  float64 `json:"fmean_bcu"`
    Acc        float64 `json:"acc"`
    Acc_ncl    float64 `json:"acc_ncl"`
    Acc_ncu    float64 `json:"acc_ncu"`
    Acc_bcl    float64 `json:"acc_bcl"`
    Acc_bcu    float64 `json:"acc_bcu"`
    Fbias      float64 `json:"fbias"`
    Fbias_bcl  float64 `json:"fbias_bcl"`
    Fbias_bcu  float64 `json:"fbias_bcu"`
    Pody       float64 `json:"pody"`
    Pody_ncl   float64 `json:"pody_ncl"`
    Pody_ncu   float64 `json:"pody_ncu"`
    Pody_bcl   float64 `json:"pody_bcl"`
    Pody_bcu   float64 `json:"pody_bcu"`
    Podn       float64 `json:"podn"`
    Podn_ncl   float64 `json:"podn_ncl"`
    Podn_ncu   float64 `json:"podn_ncu"`
    Podn_bcl   float64 `json:"podn_bcl"`
    Podn_bcu   float64 `json:"podn_bcu"`
    Pofd       float64 `json:"pofd"`
    Pofd_ncl   float64 `json:"pofd_ncl"`
    Pofd_ncu   float64 `json:"pofd_ncu"`
    Pofd_bcl   float64 `json:"pofd_bcl"`
    Pofd_bcu   float64 `json:"pofd_bcu"`
    Far        float64 `json:"far"`
    Far_ncl    float64 `json:"far_ncl"`
    Far_ncu    float64 `json:"far_ncu"`
    Far_bcl    float64 `json:"far_bcl"`
    Far_bcu    float64 `json:"far_bcu"`
    Csi        float64 `json:"csi"`
    Csi_ncl    float64 `json:"csi_ncl"`
    Csi_ncu    float64 `json:"csi_ncu"`
    Csi_bcl    float64 `json:"csi_bcl"`
    Csi_bcu    float64 `json:"csi_bcu"`
    Gss        float64 `json:"gss"`
    Gss_bcl    float64 `json:"gss_bcl"`
    Gss_bcu    float64 `json:"gss_bcu"`
    Hk         float64 `json:"hk"`
    Hk_ncl     float64 `json:"hk_ncl"`
    Hk_ncu     float64 `json:"hk_ncu"`
    Hk_bcl     float64 `json:"hk_bcl"`
    Hk_bcu     float64 `json:"hk_bcu"`
    Hss        float64 `json:"hss"`
    Hss_bcl    float64 `json:"hss_bcl"`
    Hss_bcu    float64 `json:"hss_bcu"`
    Odds       float64 `json:"odds"`
    Odds_ncl   float64 `json:"odds_ncl"`
    Odds_ncu   float64 `json:"odds_ncu"`
    Odds_bcl   float64 `json:"odds_bcl"`
    Odds_bcu   float64 `json:"odds_bcu"`
    Lodds      float64 `json:"lodds"`
    Lodds_ncl  float64 `json:"lodds_ncl"`
    Lodds_ncu  float64 `json:"lodds_ncu"`
    Lodds_bcl  float64 `json:"lodds_bcl"`
    Lodds_bcu  float64 `json:"lodds_bcu"`
    Orss       float64 `json:"orss"`
    Orss_ncl   float64 `json:"orss_ncl"`
    Orss_ncu   float64 `json:"orss_ncu"`
    Orss_bcl   float64 `json:"orss_bcl"`
    Orss_bcu   float64 `json:"orss_bcu"`
    Eds        float64 `json:"eds"`
    Eds_ncl    float64 `json:"eds_ncl"`
    Eds_ncu    float64 `json:"eds_ncu"`
    Eds_bcl    float64 `json:"eds_bcl"`
    Eds_bcu    float64 `json:"eds_bcu"`
    Seds       float64 `json:"seds"`
    Seds_ncl   float64 `json:"seds_ncl"`
    Seds_ncu   float64 `json:"seds_ncu"`
    Seds_bcl   float64 `json:"seds_bcl"`
    Seds_bcu   float64 `json:"seds_bcu"`
    Edi        float64 `json:"edi"`
    Edi_ncl    float64 `json:"edi_ncl"`
    Edi_ncu    float64 `json:"edi_ncu"`
    Edi_bcl    float64 `json:"edi_bcl"`
    Edi_bcu    float64 `json:"edi_bcu"`
    Sedi       float64 `json:"sedi"`
    Sedi_ncl   float64 `json:"sedi_ncl"`
    Sedi_ncu   float64 `json:"sedi_ncu"`
    Sedi_bcl   float64 `json:"sedi_bcl"`
    Sedi_bcu   float64 `json:"sedi_bcu"`
    Bagss      float64 `json:"bagss"`
    Bagss_bcl  float64 `json:"bagss_bcl"`
    Bagss_bcu  float64 `json:"bagss_bcu"`
    Hss_ec     float64 `json:"hss_ec"`
    Hss_ec_bcl float64 `json:"hss_ec_bcl"`
    Hss_ec_bcu float64 `json:"hss_ec_bcu"`
    Ec_value   float64 `json:"ec_value"`
}

type STAT_MCTS struct {
    Total      int     `json:"total"`
    N_cat      int     `json:"n_cat"`
    Acc        float64 `json:"acc"`
    Acc_ncl    float64 `json:"acc_ncl"`
    Acc_ncu    float64 `json:"acc_ncu"`
    Acc_bcl    float64 `json:"acc_bcl"`
    Acc_bcu    float64 `json:"acc_bcu"`
    Hk         float64 `json:"hk"`
    Hk_bcl     float64 `json:"hk_bcl"`
    Hk_bcu     float64 `json:"hk_bcu"`
    Hss        float64 `json:"hss"`
    Hss_bcl    float64 `json:"hss_bcl"`
    Hss_bcu    float64 `json:"hss_bcu"`
    Ger        float64 `json:"ger"`
    Ger_bcl    float64 `json:"ger_bcl"`
    Ger_bcu    float64 `json:"ger_bcu"`
    Hss_ec     float64 `json:"hss_ec"`
    Hss_ec_bcl float64 `json:"hss_ec_bcl"`
    Hss_ec_bcu float64 `json:"hss_ec_bcu"`
    Ec_value   float64 `json:"ec_value"`
}

type STAT_PHIST struct {
    Total      int     `json:"total"`
    Bin_size   float64 `json:"bin_size"`
    N_bin      int     `json:"n_bin"`
    Bin_i              `json:"bin_i"`
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
    Total_dir   float64 `json:"total_dir"`
    Dir_me      float64 `json:"dir_me"`
    Dir_mae     float64 `json:"dir_mae"`
    Dir_mse     float64 `json:"dir_mse"`
}

type MTD_3DPAIR struct {
    Object_id           string  `json:"object_id"`
    Object_cat          string  `json:"object_cat"`
    Space_centroid_dist float64 `json:"space_centroid_dist"`
    Time_centroid_delta float64 `json:"time_centroid_delta"`
    Axis_diff           float64 `json:"axis_diff"`
    Speed_delta         float64 `json:"speed_delta"`
    Direction_diff      float64 `json:"direction_diff"`
    Volume_ratio        float64 `json:"volume_ratio"`
    Start_time_delta    int     `json:"start_time_delta"`
    End_time_delta      int     `json:"end_time_delta"`
    Intersection_volume int     `json:"intersection_volume"`
    Duration_diff       int     `json:"duration_diff"`
    Interest            float64 `json:"interest"`
}

type STAT_NBRCNT struct {
    Total      int     `json:"total"`
    Fbs        float64 `json:"fbs"`
    Fbs_bcl    float64 `json:"fbs_bcl"`
    Fbs_bcu    float64 `json:"fbs_bcu"`
    Fss        float64 `json:"fss"`
    Fss_bcl    float64 `json:"fss_bcl"`
    Fss_bcu    float64 `json:"fss_bcu"`
    Afss       float64 `json:"afss"`
    Afss_bcl   float64 `json:"afss_bcl"`
    Afss_bcu   float64 `json:"afss_bcu"`
    Ufss       float64 `json:"ufss"`
    Ufss_bcl   float64 `json:"ufss_bcl"`
    Ufss_bcu   float64 `json:"ufss_bcu"`
    F_rate     float64 `json:"f_rate"`
    F_rate_bcl float64 `json:"f_rate_bcl"`
    F_rate_bcu float64 `json:"f_rate_bcu"`
    O_rate     float64 `json:"o_rate"`
    O_rate_bcl float64 `json:"o_rate_bcl"`
    O_rate_bcu float64 `json:"o_rate_bcu"`
}

type STAT_RPS struct {
    Total     int     `json:"total"`
    N_prob    int     `json:"n_prob"`
    Rps_rel   float64 `json:"rps_rel"`
    Rps_res   float64 `json:"rps_res"`
    Rps_unc   float64 `json:"rps_unc"`
    Rps       float64 `json:"rps"`
    Rpss      float64 `json:"rpss"`
    Rpss_smpl float64 `json:"rpss_smpl"`
    Rps_comp  string  `json:"rps_comp"`
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

type STAT_PJC struct {
    Total              int     `json:"total"`
    N_thresh           int     `json:"n_thresh"`
    Thresh_i                   `json:"thresh_i"`
    Oy_tp_i                    `json:"oy_tp_i"`
    On_tp_i                    `json:"on_tp_i"`
    Calibration_i              `json:"calibration_i"`
    Refinement_i               `json:"refinement_i"`
    Likelihood_i       float64 `json:"likelihood_i"`
    Baser_i                    `json:"baser_i"`
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

type MODE_OBJ struct {
}

type MODE_CTS struct {
}

type TCST_TCDIAG struct {
    Total        int     `json:"total"`
    Index        int     `json:"index"`
    Diag_source  string  `json:"diag_source"`
    Track_source string  `json:"track_source"`
    Field_source string  `json:"field_source"`
    N_diag       int     `json:"n_diag"`
    Diag_i               `json:"diag_i"`
    Value_i              `json:"value_i"`
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

type STAT_SSVAR struct {
    Total       int     `json:"total"`
    N_bin       int     `json:"n_bin"`
    Bin_i               `json:"bin_i"`
    Bin_n       int     `json:"bin_n"`
    Var_min     float64 `json:"var_min"`
    Var_max     float64 `json:"var_max"`
    Var_mean    float64 `json:"var_mean"`
    Fbar        float64 `json:"fbar"`
    Obar        float64 `json:"obar"`
    Fobar       float64 `json:"fobar"`
    Ffbar       float64 `json:"ffbar"`
    Oobar       float64 `json:"oobar"`
    Fbar_ncl    float64 `json:"fbar_ncl"`
    Fbar_ncu    float64 `json:"fbar_ncu"`
    Fstdev      float64 `json:"fstdev"`
    Fstdev_ncl  float64 `json:"fstdev_ncl"`
    Fstdev_ncu  float64 `json:"fstdev_ncu"`
    Obar_ncl    float64 `json:"obar_ncl"`
    Obar_ncu    float64 `json:"obar_ncu"`
    Ostdev      float64 `json:"ostdev"`
    Ostdev_ncl  float64 `json:"ostdev_ncl"`
    Ostdev_ncu  float64 `json:"ostdev_ncu"`
    Pr_corr     float64 `json:"pr_corr"`
    Pr_corr_ncl float64 `json:"pr_corr_ncl"`
    Pr_corr_ncu float64 `json:"pr_corr_ncu"`
    Me          float64 `json:"me"`
    Me_ncl      float64 `json:"me_ncl"`
    Me_ncu      float64 `json:"me_ncu"`
    Estdev      float64 `json:"estdev"`
    Estdev_ncl  float64 `json:"estdev_ncl"`
    Estdev_ncu  float64 `json:"estdev_ncu"`
    Mbias       float64 `json:"mbias"`
    Mse         float64 `json:"mse"`
    Bcmse       float64 `json:"bcmse"`
    Rmse        float64 `json:"rmse"`
}

type TCST_TCMPR struct {
    Total          int     `json:"total"`
    Index          int     `json:"index"`
    Level          string  `json:"level"`
    Watch_warn     string  `json:"watch_warn"`
    Initials       string  `json:"initials"`
    Alat           float64 `json:"alat"`
    Alon           float64 `json:"alon"`
    Blat           float64 `json:"blat"`
    Blon           float64 `json:"blon"`
    Tk_err         float64 `json:"tk_err"`
    X_err          float64 `json:"x_err"`
    Y_err          float64 `json:"y_err"`
    Altk_err       float64 `json:"altk_err"`
    Crtk_err       float64 `json:"crtk_err"`
    Adland         float64 `json:"adland"`
    Bdland         float64 `json:"bdland"`
    Amslp          float64 `json:"amslp"`
    Bmslp          float64 `json:"bmslp"`
    Amax_wind      float64 `json:"amax_wind"`
    Bmax_wind      float64 `json:"bmax_wind"`
    Aal_wind_34    float64 `json:"aal_wind_34"`
    Bal_wind_34    float64 `json:"bal_wind_34"`
    Ane_wind_34    float64 `json:"ane_wind_34"`
    Bne_wind_34    float64 `json:"bne_wind_34"`
    Ase_wind_34    float64 `json:"ase_wind_34"`
    Bse_wind_34    float64 `json:"bse_wind_34"`
    Asw_wind_34    float64 `json:"asw_wind_34"`
    Bsw_wind_34    float64 `json:"bsw_wind_34"`
    Anw_wind_34    float64 `json:"anw_wind_34"`
    Bnw_wind_34    float64 `json:"bnw_wind_34"`
    Aal_wind_50    float64 `json:"aal_wind_50"`
    Bal_wind_50    float64 `json:"bal_wind_50"`
    Ane_wind_50    float64 `json:"ane_wind_50"`
    Bne_wind_50    float64 `json:"bne_wind_50"`
    Ase_wind_50    float64 `json:"ase_wind_50"`
    Bse_wind_50    float64 `json:"bse_wind_50"`
    Asw_wind_50    float64 `json:"asw_wind_50"`
    Bsw_wind_50    float64 `json:"bsw_wind_50"`
    Anw_wind_50    float64 `json:"anw_wind_50"`
    Bnw_wind_50    float64 `json:"bnw_wind_50"`
    Aal_wind_64    float64 `json:"aal_wind_64"`
    Bal_wind_64    float64 `json:"bal_wind_64"`
    Ane_wind_64    float64 `json:"ane_wind_64"`
    Bne_wind_64    float64 `json:"bne_wind_64"`
    Ase_wind_64    float64 `json:"ase_wind_64"`
    Bse_wind_64    float64 `json:"bse_wind_64"`
    Asw_wind_64    float64 `json:"asw_wind_64"`
    Bsw_wind_64    float64 `json:"bsw_wind_64"`
    Anw_wind_64    float64 `json:"anw_wind_64"`
    Bnw_wind_64    float64 `json:"bnw_wind_64"`
    Aradp          string  `json:"aradp"`
    Bradp          float64 `json:"bradp"`
    Arrp           string  `json:"arrp"`
    Brrp           float64 `json:"brrp"`
    Amrd           string  `json:"amrd"`
    Bmrd           float64 `json:"bmrd"`
    Agusts         string  `json:"agusts"`
    Bgusts         float64 `json:"bgusts"`
    Aeye           string  `json:"aeye"`
    Beye           float64 `json:"beye"`
    Adir           string  `json:"adir"`
    Bdir           float64 `json:"bdir"`
    Aspeed         float64 `json:"aspeed"`
    Bspeed         float64 `json:"bspeed"`
    Adepth         string  `json:"adepth"`
    Bdepth         float64 `json:"bdepth"`
    Num_members    float64 `json:"num_members"`
    Track_spread   float64 `json:"track_spread"`
    Track_stdev    float64 `json:"track_stdev"`
    Mslp_stdev     float64 `json:"mslp_stdev"`
    Max_wind_stdev float64 `json:"max_wind_stdev"`
}

type STAT_NBRCTC struct {
    Total int     `json:"total"`
    Fy_oy float64 `json:"fy_oy"`
    Fy_on float64 `json:"fy_on"`
    Fn_oy float64 `json:"fn_oy"`
    Fn_on float64 `json:"fn_on"`
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

type STAT_PSTD struct {
    Total         int     `json:"total"`
    N_thresh      int     `json:"n_thresh"`
    Baser         float64 `json:"baser"`
    Baser_ncl     float64 `json:"baser_ncl"`
    Baser_ncu     float64 `json:"baser_ncu"`
    Reliability   float64 `json:"reliability"`
    Resolution    float64 `json:"resolution"`
    Uncertainty   float64 `json:"uncertainty"`
    Roc_auc       float64 `json:"roc_auc"`
    Brier         float64 `json:"brier"`
    Brier_ncl     float64 `json:"brier_ncl"`
    Brier_ncu     float64 `json:"brier_ncu"`
    Briercl       float64 `json:"briercl"`
    Briercl_ncl   float64 `json:"briercl_ncl"`
    Briercl_ncu   float64 `json:"briercl_ncu"`
    Bss           float64 `json:"bss"`
    Bss_smpl      float64 `json:"bss_smpl"`
    Thresh_i              `json:"thresh_i"`
}

type STAT_ECNT struct {
    Total            int     `json:"total"`
    N_ens            int     `json:"n_ens"`
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

type STAT_RELP struct {
    Total       int     `json:"total"`
    N_ens       int     `json:"n_ens"`
    Relp_i              `json:"relp_i"`
}

type STAT_VCNT struct {
    Total                int     `json:"total"`
    Fbar                 float64 `json:"fbar"`
    Fbar_bcl             float64 `json:"fbar_bcl"`
    Fbar_bcu             float64 `json:"fbar_bcu"`
    Obar                 float64 `json:"obar"`
    Obar_bcl             float64 `json:"obar_bcl"`
    Obar_bcu             float64 `json:"obar_bcu"`
    Fs_rms               float64 `json:"fs_rms"`
    Fs_rms_bcl           float64 `json:"fs_rms_bcl"`
    Fs_rms_bcu           float64 `json:"fs_rms_bcu"`
    Os_rms               float64 `json:"os_rms"`
    Os_rms_bcl           float64 `json:"os_rms_bcl"`
    Os_rms_bcu           float64 `json:"os_rms_bcu"`
    Msve                 float64 `json:"msve"`
    Msve_bcl             float64 `json:"msve_bcl"`
    Msve_bcu             float64 `json:"msve_bcu"`
    Rmsve                float64 `json:"rmsve"`
    Rmsve_bcl            float64 `json:"rmsve_bcl"`
    Rmsve_bcu            float64 `json:"rmsve_bcu"`
    Fstdev               float64 `json:"fstdev"`
    Fstdev_bcl           float64 `json:"fstdev_bcl"`
    Fstdev_bcu           float64 `json:"fstdev_bcu"`
    Ostdev               float64 `json:"ostdev"`
    Ostdev_bcl           float64 `json:"ostdev_bcl"`
    Ostdev_bcu           float64 `json:"ostdev_bcu"`
    Fdir                 float64 `json:"fdir"`
    Fdir_bcl             float64 `json:"fdir_bcl"`
    Fdir_bcu             float64 `json:"fdir_bcu"`
    Odir                 float64 `json:"odir"`
    Odir_bcl             float64 `json:"odir_bcl"`
    Odir_bcu             float64 `json:"odir_bcu"`
    Fbar_speed           float64 `json:"fbar_speed"`
    Fbar_speed_bcl       float64 `json:"fbar_speed_bcl"`
    Fbar_speed_bcu       float64 `json:"fbar_speed_bcu"`
    Obar_speed           float64 `json:"obar_speed"`
    Obar_speed_bcl       float64 `json:"obar_speed_bcl"`
    Obar_speed_bcu       float64 `json:"obar_speed_bcu"`
    Vdiff_speed          float64 `json:"vdiff_speed"`
    Vdiff_speed_bcl      float64 `json:"vdiff_speed_bcl"`
    Vdiff_speed_bcu      float64 `json:"vdiff_speed_bcu"`
    Vdiff_dir            float64 `json:"vdiff_dir"`
    Vdiff_dir_bcl        float64 `json:"vdiff_dir_bcl"`
    Vdiff_dir_bcu        float64 `json:"vdiff_dir_bcu"`
    Speed_err            float64 `json:"speed_err"`
    Speed_err_bcl        float64 `json:"speed_err_bcl"`
    Speed_err_bcu        float64 `json:"speed_err_bcu"`
    Speed_abserr         float64 `json:"speed_abserr"`
    Speed_abserr_bcl     float64 `json:"speed_abserr_bcl"`
    Speed_abserr_bcu     float64 `json:"speed_abserr_bcu"`
    Dir_err              float64 `json:"dir_err"`
    Dir_err_bcl          float64 `json:"dir_err_bcl"`
    Dir_err_bcu          float64 `json:"dir_err_bcu"`
    Dir_abserr           float64 `json:"dir_abserr"`
    Dir_abserr_bcl       float64 `json:"dir_abserr_bcl"`
    Dir_abserr_bcu       float64 `json:"dir_abserr_bcu"`
    Anom_corr            float64 `json:"anom_corr"`
    Anom_corr_ncl        float64 `json:"anom_corr_ncl"`
    Anom_corr_ncu        float64 `json:"anom_corr_ncu"`
    Anom_corr_bcl        float64 `json:"anom_corr_bcl"`
    Anom_corr_bcu        float64 `json:"anom_corr_bcu"`
    Anom_corr_uncntr     float64 `json:"anom_corr_uncntr"`
    Anom_corr_uncntr_bcl float64 `json:"anom_corr_uncntr_bcl"`
    Anom_corr_uncntr_bcu float64 `json:"anom_corr_uncntr_bcu"`
    Total_dir            float64 `json:"total_dir"`
    Dir_me               float64 `json:"dir_me"`
    Dir_me_bcl           float64 `json:"dir_me_bcl"`
    Dir_me_bcu           float64 `json:"dir_me_bcu"`
    Dir_mae              float64 `json:"dir_mae"`
    Dir_mae_bcl          float64 `json:"dir_mae_bcl"`
    Dir_mae_bcu          float64 `json:"dir_mae_bcu"`
    Dir_mse              float64 `json:"dir_mse"`
    Dir_mse_bcl          float64 `json:"dir_mse_bcl"`
    Dir_mse_bcu          float64 `json:"dir_mse_bcu"`
    Dir_rmse             float64 `json:"dir_rmse"`
    Dir_rmse_bcl         float64 `json:"dir_rmse_bcl"`
    Dir_rmse_bcu         float64 `json:"dir_rmse_bcu"`
}

type STAT_ISC struct {
    Total    int     `json:"total"`
    Tile_dim int     `json:"tile_dim"`
    Tile_xll int     `json:"tile_xll"`
    Tile_yll int     `json:"tile_yll"`
    Nscale   int     `json:"nscale"`
    Iscale   int     `json:"iscale"`
    Mse      float64 `json:"mse"`
    Isc      float64 `json:"isc"`
    Fenergy2 float64 `json:"fenergy2"`
    Oenergy2 float64 `json:"oenergy2"`
    Baser    float64 `json:"baser"`
    Fbias    float64 `json:"fbias"`
}

type STAT_NBRCTS struct {
    Total     int     `json:"total"`
    Baser     float64 `json:"baser"`
    Baser_ncl float64 `json:"baser_ncl"`
    Baser_ncu float64 `json:"baser_ncu"`
    Baser_bcl float64 `json:"baser_bcl"`
    Baser_bcu float64 `json:"baser_bcu"`
    Fmean     float64 `json:"fmean"`
    Fmean_ncl float64 `json:"fmean_ncl"`
    Fmean_ncu float64 `json:"fmean_ncu"`
    Fmean_bcl float64 `json:"fmean_bcl"`
    Fmean_bcu float64 `json:"fmean_bcu"`
    Acc       float64 `json:"acc"`
    Acc_ncl   float64 `json:"acc_ncl"`
    Acc_ncu   float64 `json:"acc_ncu"`
    Acc_bcl   float64 `json:"acc_bcl"`
    Acc_bcu   float64 `json:"acc_bcu"`
    Fbias     float64 `json:"fbias"`
    Fbias_bcl float64 `json:"fbias_bcl"`
    Fbias_bcu float64 `json:"fbias_bcu"`
    Pody      float64 `json:"pody"`
    Pody_ncl  float64 `json:"pody_ncl"`
    Pody_ncu  float64 `json:"pody_ncu"`
    Pody_bcl  float64 `json:"pody_bcl"`
    Pody_bcu  float64 `json:"pody_bcu"`
    Podn      float64 `json:"podn"`
    Podn_ncl  float64 `json:"podn_ncl"`
    Podn_ncu  float64 `json:"podn_ncu"`
    Podn_bcl  float64 `json:"podn_bcl"`
    Podn_bcu  float64 `json:"podn_bcu"`
    Pofd      float64 `json:"pofd"`
    Pofd_ncl  float64 `json:"pofd_ncl"`
    Pofd_ncu  float64 `json:"pofd_ncu"`
    Pofd_bcl  float64 `json:"pofd_bcl"`
    Pofd_bcu  float64 `json:"pofd_bcu"`
    Far       float64 `json:"far"`
    Far_ncl   float64 `json:"far_ncl"`
    Far_ncu   float64 `json:"far_ncu"`
    Far_bcl   float64 `json:"far_bcl"`
    Far_bcu   float64 `json:"far_bcu"`
    Csi       float64 `json:"csi"`
    Csi_ncl   float64 `json:"csi_ncl"`
    Csi_ncu   float64 `json:"csi_ncu"`
    Csi_bcl   float64 `json:"csi_bcl"`
    Csi_bcu   float64 `json:"csi_bcu"`
    Gss       float64 `json:"gss"`
    Gss_bcl   float64 `json:"gss_bcl"`
    Gss_bcu   float64 `json:"gss_bcu"`
    Hk        float64 `json:"hk"`
    Hk_ncl    float64 `json:"hk_ncl"`
    Hk_ncu    float64 `json:"hk_ncu"`
    Hk_bcl    float64 `json:"hk_bcl"`
    Hk_bcu    float64 `json:"hk_bcu"`
    Hss       float64 `json:"hss"`
    Hss_bcl   float64 `json:"hss_bcl"`
    Hss_bcu   float64 `json:"hss_bcu"`
    Odds      float64 `json:"odds"`
    Odds_ncl  float64 `json:"odds_ncl"`
    Odds_ncu  float64 `json:"odds_ncu"`
    Odds_bcl  float64 `json:"odds_bcl"`
    Odds_bcu  float64 `json:"odds_bcu"`
    Lodds     float64 `json:"lodds"`
    Lodds_ncl float64 `json:"lodds_ncl"`
    Lodds_ncu float64 `json:"lodds_ncu"`
    Lodds_bcl float64 `json:"lodds_bcl"`
    Lodds_bcu float64 `json:"lodds_bcu"`
    Orss      float64 `json:"orss"`
    Orss_ncl  float64 `json:"orss_ncl"`
    Orss_ncu  float64 `json:"orss_ncu"`
    Orss_bcl  float64 `json:"orss_bcl"`
    Orss_bcu  float64 `json:"orss_bcu"`
    Eds       float64 `json:"eds"`
    Eds_ncl   float64 `json:"eds_ncl"`
    Eds_ncu   float64 `json:"eds_ncu"`
    Eds_bcl   float64 `json:"eds_bcl"`
    Eds_bcu   float64 `json:"eds_bcu"`
    Seds      float64 `json:"seds"`
    Seds_ncl  float64 `json:"seds_ncl"`
    Seds_ncu  float64 `json:"seds_ncu"`
    Seds_bcl  float64 `json:"seds_bcl"`
    Seds_bcu  float64 `json:"seds_bcu"`
    Edi       float64 `json:"edi"`
    Edi_ncl   float64 `json:"edi_ncl"`
    Edi_ncu   float64 `json:"edi_ncu"`
    Edi_bcl   float64 `json:"edi_bcl"`
    Edi_bcu   float64 `json:"edi_bcu"`
    Sedi      float64 `json:"sedi"`
    Sedi_ncl  float64 `json:"sedi_ncl"`
    Sedi_ncu  float64 `json:"sedi_ncu"`
    Sedi_bcl  float64 `json:"sedi_bcl"`
    Sedi_bcu  float64 `json:"sedi_bcu"`
    Bagss     float64 `json:"bagss"`
    Bagss_bcl float64 `json:"bagss_bcl"`
    Bagss_bcu float64 `json:"bagss_bcu"`
}


//fillStructure functions
func (s *MODE_OBJ) fill_MODE_OBJ(fields []string) {
}

func (s *TCST_TCMPR) fill_TCST_TCMPR(fields []string) {
	s.Total, _ = strconv.Atoi(fields[0])
	s.Index, _ = strconv.Atoi(fields[1])
	s.Level = fields[2]
	s.Watch_warn = fields[3]
	s.Initials = fields[4]
	s.Alat, _ = strconv.ParseFloat(fields[5], 64)
	s.Alon, _ = strconv.ParseFloat(fields[6], 64)
	s.Blat, _ = strconv.ParseFloat(fields[7], 64)
	s.Blon, _ = strconv.ParseFloat(fields[8], 64)
	s.Tk_err, _ = strconv.ParseFloat(fields[9], 64)
	s.X_err, _ = strconv.ParseFloat(fields[10], 64)
	s.Y_err, _ = strconv.ParseFloat(fields[11], 64)
	s.Altk_err, _ = strconv.ParseFloat(fields[12], 64)
	s.Crtk_err, _ = strconv.ParseFloat(fields[13], 64)
	s.Adland, _ = strconv.ParseFloat(fields[14], 64)
	s.Bdland, _ = strconv.ParseFloat(fields[15], 64)
	s.Amslp, _ = strconv.ParseFloat(fields[16], 64)
	s.Bmslp, _ = strconv.ParseFloat(fields[17], 64)
	s.Amax_wind, _ = strconv.ParseFloat(fields[18], 64)
	s.Bmax_wind, _ = strconv.ParseFloat(fields[19], 64)
	s.Aal_wind_34, _ = strconv.ParseFloat(fields[20], 64)
	s.Bal_wind_34, _ = strconv.ParseFloat(fields[21], 64)
	s.Ane_wind_34, _ = strconv.ParseFloat(fields[22], 64)
	s.Bne_wind_34, _ = strconv.ParseFloat(fields[23], 64)
	s.Ase_wind_34, _ = strconv.ParseFloat(fields[24], 64)
	s.Bse_wind_34, _ = strconv.ParseFloat(fields[25], 64)
	s.Asw_wind_34, _ = strconv.ParseFloat(fields[26], 64)
	s.Bsw_wind_34, _ = strconv.ParseFloat(fields[27], 64)
	s.Anw_wind_34, _ = strconv.ParseFloat(fields[28], 64)
	s.Bnw_wind_34, _ = strconv.ParseFloat(fields[29], 64)
	s.Aal_wind_50, _ = strconv.ParseFloat(fields[30], 64)
	s.Bal_wind_50, _ = strconv.ParseFloat(fields[31], 64)
	s.Ane_wind_50, _ = strconv.ParseFloat(fields[32], 64)
	s.Bne_wind_50, _ = strconv.ParseFloat(fields[33], 64)
	s.Ase_wind_50, _ = strconv.ParseFloat(fields[34], 64)
	s.Bse_wind_50, _ = strconv.ParseFloat(fields[35], 64)
	s.Asw_wind_50, _ = strconv.ParseFloat(fields[36], 64)
	s.Bsw_wind_50, _ = strconv.ParseFloat(fields[37], 64)
	s.Anw_wind_50, _ = strconv.ParseFloat(fields[38], 64)
	s.Bnw_wind_50, _ = strconv.ParseFloat(fields[39], 64)
	s.Aal_wind_64, _ = strconv.ParseFloat(fields[40], 64)
	s.Bal_wind_64, _ = strconv.ParseFloat(fields[41], 64)
	s.Ane_wind_64, _ = strconv.ParseFloat(fields[42], 64)
	s.Bne_wind_64, _ = strconv.ParseFloat(fields[43], 64)
	s.Ase_wind_64, _ = strconv.ParseFloat(fields[44], 64)
	s.Bse_wind_64, _ = strconv.ParseFloat(fields[45], 64)
	s.Asw_wind_64, _ = strconv.ParseFloat(fields[46], 64)
	s.Bsw_wind_64, _ = strconv.ParseFloat(fields[47], 64)
	s.Anw_wind_64, _ = strconv.ParseFloat(fields[48], 64)
	s.Bnw_wind_64, _ = strconv.ParseFloat(fields[49], 64)
	s.Aradp = fields[50]
	s.Bradp, _ = strconv.ParseFloat(fields[51], 64)
	s.Arrp = fields[52]
	s.Brrp, _ = strconv.ParseFloat(fields[53], 64)
	s.Amrd = fields[54]
	s.Bmrd, _ = strconv.ParseFloat(fields[55], 64)
	s.Agusts = fields[56]
	s.Bgusts, _ = strconv.ParseFloat(fields[57], 64)
	s.Aeye = fields[58]
	s.Beye, _ = strconv.ParseFloat(fields[59], 64)
	s.Adir = fields[60]
	s.Bdir, _ = strconv.ParseFloat(fields[61], 64)
	s.Aspeed, _ = strconv.ParseFloat(fields[62], 64)
	s.Bspeed, _ = strconv.ParseFloat(fields[63], 64)
	s.Adepth = fields[64]
	s.Bdepth, _ = strconv.ParseFloat(fields[65], 64)
	s.Num_members, _ = strconv.ParseFloat(fields[66], 64)
	s.Track_spread, _ = strconv.ParseFloat(fields[67], 64)
	s.Track_stdev, _ = strconv.ParseFloat(fields[68], 64)
	s.Mslp_stdev, _ = strconv.ParseFloat(fields[69], 64)
	s.Max_wind_stdev, _ = strconv.ParseFloat(fields[70], 64)
}

func (s *TCST_PROBRIRW) fill_TCST_PROBRIRW(fields []string) {
	s.Alat, _ = strconv.ParseFloat(fields[0], 64)
	s.Alon, _ = strconv.ParseFloat(fields[1], 64)
	s.Blat, _ = strconv.ParseFloat(fields[2], 64)
	s.Blon, _ = strconv.ParseFloat(fields[3], 64)
	s.Initials = fields[4]
	s.Tk_err, _ = strconv.ParseFloat(fields[5], 64)
	s.X_err, _ = strconv.ParseFloat(fields[6], 64)
	s.Y_err, _ = strconv.ParseFloat(fields[7], 64)
	s.Adland, _ = strconv.ParseFloat(fields[8], 64)
	s.Bdland, _ = strconv.ParseFloat(fields[9], 64)
	s.Rirw_beg, _ = strconv.Atoi(fields[10])
	s.Rirw_end, _ = strconv.Atoi(fields[11])
	s.Rirw_window, _ = strconv.ParseFloat(fields[12], 64)
	s.Awind_end, _ = strconv.ParseFloat(fields[13], 64)
	s.Bwind_beg, _ = strconv.ParseFloat(fields[14], 64)
	s.Bwind_end, _ = strconv.ParseFloat(fields[15], 64)
	s.Bdelta, _ = strconv.ParseFloat(fields[16], 64)
	s.Bdelta_max, _ = strconv.ParseFloat(fields[17], 64)
	s.Blevel_beg = fields[18]
	s.Blevel_end = fields[19]
	s.N_thresh, _ = strconv.Atoi(fields[20])
	s.Thresh_i = fields[21]
	s.Prob_i = fields[22]
}

func (s *STAT_CTC) fill_STAT_CTC(fields []string) {
	s.Total, _ = strconv.Atoi(fields[0])
	s.Fy_oy, _ = strconv.ParseFloat(fields[1], 64)
	s.Fy_on, _ = strconv.ParseFloat(fields[2], 64)
	s.Fn_oy, _ = strconv.ParseFloat(fields[3], 64)
	s.Fn_on, _ = strconv.ParseFloat(fields[4], 64)
	s.Ec_value, _ = strconv.ParseFloat(fields[5], 64)
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
	s.Total_dir, _ = strconv.ParseFloat(fields[10], 64)
	s.Dir_me, _ = strconv.ParseFloat(fields[11], 64)
	s.Dir_mae, _ = strconv.ParseFloat(fields[12], 64)
	s.Dir_mse, _ = strconv.ParseFloat(fields[13], 64)
}

func (s *STAT_VCNT) fill_STAT_VCNT(fields []string) {
	s.Total, _ = strconv.Atoi(fields[0])
	s.Fbar, _ = strconv.ParseFloat(fields[1], 64)
	s.Fbar_bcl, _ = strconv.ParseFloat(fields[2], 64)
	s.Fbar_bcu, _ = strconv.ParseFloat(fields[3], 64)
	s.Obar, _ = strconv.ParseFloat(fields[4], 64)
	s.Obar_bcl, _ = strconv.ParseFloat(fields[5], 64)
	s.Obar_bcu, _ = strconv.ParseFloat(fields[6], 64)
	s.Fs_rms, _ = strconv.ParseFloat(fields[7], 64)
	s.Fs_rms_bcl, _ = strconv.ParseFloat(fields[8], 64)
	s.Fs_rms_bcu, _ = strconv.ParseFloat(fields[9], 64)
	s.Os_rms, _ = strconv.ParseFloat(fields[10], 64)
	s.Os_rms_bcl, _ = strconv.ParseFloat(fields[11], 64)
	s.Os_rms_bcu, _ = strconv.ParseFloat(fields[12], 64)
	s.Msve, _ = strconv.ParseFloat(fields[13], 64)
	s.Msve_bcl, _ = strconv.ParseFloat(fields[14], 64)
	s.Msve_bcu, _ = strconv.ParseFloat(fields[15], 64)
	s.Rmsve, _ = strconv.ParseFloat(fields[16], 64)
	s.Rmsve_bcl, _ = strconv.ParseFloat(fields[17], 64)
	s.Rmsve_bcu, _ = strconv.ParseFloat(fields[18], 64)
	s.Fstdev, _ = strconv.ParseFloat(fields[19], 64)
	s.Fstdev_bcl, _ = strconv.ParseFloat(fields[20], 64)
	s.Fstdev_bcu, _ = strconv.ParseFloat(fields[21], 64)
	s.Ostdev, _ = strconv.ParseFloat(fields[22], 64)
	s.Ostdev_bcl, _ = strconv.ParseFloat(fields[23], 64)
	s.Ostdev_bcu, _ = strconv.ParseFloat(fields[24], 64)
	s.Fdir, _ = strconv.ParseFloat(fields[25], 64)
	s.Fdir_bcl, _ = strconv.ParseFloat(fields[26], 64)
	s.Fdir_bcu, _ = strconv.ParseFloat(fields[27], 64)
	s.Odir, _ = strconv.ParseFloat(fields[28], 64)
	s.Odir_bcl, _ = strconv.ParseFloat(fields[29], 64)
	s.Odir_bcu, _ = strconv.ParseFloat(fields[30], 64)
	s.Fbar_speed, _ = strconv.ParseFloat(fields[31], 64)
	s.Fbar_speed_bcl, _ = strconv.ParseFloat(fields[32], 64)
	s.Fbar_speed_bcu, _ = strconv.ParseFloat(fields[33], 64)
	s.Obar_speed, _ = strconv.ParseFloat(fields[34], 64)
	s.Obar_speed_bcl, _ = strconv.ParseFloat(fields[35], 64)
	s.Obar_speed_bcu, _ = strconv.ParseFloat(fields[36], 64)
	s.Vdiff_speed, _ = strconv.ParseFloat(fields[37], 64)
	s.Vdiff_speed_bcl, _ = strconv.ParseFloat(fields[38], 64)
	s.Vdiff_speed_bcu, _ = strconv.ParseFloat(fields[39], 64)
	s.Vdiff_dir, _ = strconv.ParseFloat(fields[40], 64)
	s.Vdiff_dir_bcl, _ = strconv.ParseFloat(fields[41], 64)
	s.Vdiff_dir_bcu, _ = strconv.ParseFloat(fields[42], 64)
	s.Speed_err, _ = strconv.ParseFloat(fields[43], 64)
	s.Speed_err_bcl, _ = strconv.ParseFloat(fields[44], 64)
	s.Speed_err_bcu, _ = strconv.ParseFloat(fields[45], 64)
	s.Speed_abserr, _ = strconv.ParseFloat(fields[46], 64)
	s.Speed_abserr_bcl, _ = strconv.ParseFloat(fields[47], 64)
	s.Speed_abserr_bcu, _ = strconv.ParseFloat(fields[48], 64)
	s.Dir_err, _ = strconv.ParseFloat(fields[49], 64)
	s.Dir_err_bcl, _ = strconv.ParseFloat(fields[50], 64)
	s.Dir_err_bcu, _ = strconv.ParseFloat(fields[51], 64)
	s.Dir_abserr, _ = strconv.ParseFloat(fields[52], 64)
	s.Dir_abserr_bcl, _ = strconv.ParseFloat(fields[53], 64)
	s.Dir_abserr_bcu, _ = strconv.ParseFloat(fields[54], 64)
	s.Anom_corr, _ = strconv.ParseFloat(fields[55], 64)
	s.Anom_corr_ncl, _ = strconv.ParseFloat(fields[56], 64)
	s.Anom_corr_ncu, _ = strconv.ParseFloat(fields[57], 64)
	s.Anom_corr_bcl, _ = strconv.ParseFloat(fields[58], 64)
	s.Anom_corr_bcu, _ = strconv.ParseFloat(fields[59], 64)
	s.Anom_corr_uncntr, _ = strconv.ParseFloat(fields[60], 64)
	s.Anom_corr_uncntr_bcl, _ = strconv.ParseFloat(fields[61], 64)
	s.Anom_corr_uncntr_bcu, _ = strconv.ParseFloat(fields[62], 64)
	s.Total_dir, _ = strconv.ParseFloat(fields[63], 64)
	s.Dir_me, _ = strconv.ParseFloat(fields[64], 64)
	s.Dir_me_bcl, _ = strconv.ParseFloat(fields[65], 64)
	s.Dir_me_bcu, _ = strconv.ParseFloat(fields[66], 64)
	s.Dir_mae, _ = strconv.ParseFloat(fields[67], 64)
	s.Dir_mae_bcl, _ = strconv.ParseFloat(fields[68], 64)
	s.Dir_mae_bcu, _ = strconv.ParseFloat(fields[69], 64)
	s.Dir_mse, _ = strconv.ParseFloat(fields[70], 64)
	s.Dir_mse_bcl, _ = strconv.ParseFloat(fields[71], 64)
	s.Dir_mse_bcu, _ = strconv.ParseFloat(fields[72], 64)
	s.Dir_rmse, _ = strconv.ParseFloat(fields[73], 64)
	s.Dir_rmse_bcl, _ = strconv.ParseFloat(fields[74], 64)
	s.Dir_rmse_bcu, _ = strconv.ParseFloat(fields[75], 64)
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
	s.N_ens, _ = strconv.Atoi(fields[11])
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

func (s *STAT_NBRCNT) fill_STAT_NBRCNT(fields []string) {
	s.Total, _ = strconv.Atoi(fields[0])
	s.Fbs, _ = strconv.ParseFloat(fields[1], 64)
	s.Fbs_bcl, _ = strconv.ParseFloat(fields[2], 64)
	s.Fbs_bcu, _ = strconv.ParseFloat(fields[3], 64)
	s.Fss, _ = strconv.ParseFloat(fields[4], 64)
	s.Fss_bcl, _ = strconv.ParseFloat(fields[5], 64)
	s.Fss_bcu, _ = strconv.ParseFloat(fields[6], 64)
	s.Afss, _ = strconv.ParseFloat(fields[7], 64)
	s.Afss_bcl, _ = strconv.ParseFloat(fields[8], 64)
	s.Afss_bcu, _ = strconv.ParseFloat(fields[9], 64)
	s.Ufss, _ = strconv.ParseFloat(fields[10], 64)
	s.Ufss_bcl, _ = strconv.ParseFloat(fields[11], 64)
	s.Ufss_bcu, _ = strconv.ParseFloat(fields[12], 64)
	s.F_rate, _ = strconv.ParseFloat(fields[13], 64)
	s.F_rate_bcl, _ = strconv.ParseFloat(fields[14], 64)
	s.F_rate_bcu, _ = strconv.ParseFloat(fields[15], 64)
	s.O_rate, _ = strconv.ParseFloat(fields[16], 64)
	s.O_rate_bcl, _ = strconv.ParseFloat(fields[17], 64)
	s.O_rate_bcu, _ = strconv.ParseFloat(fields[18], 64)
}

func (s *STAT_PRC) fill_STAT_PRC(fields []string) {
	s.Total, _ = strconv.Atoi(fields[0])
	s.N_thresh, _ = strconv.Atoi(fields[1])
	s.Thresh_i = fields[2]
	s.Pody_i = fields[3]
	s.Pofd_i = fields[4]
}

func (s *STAT_CTS) fill_STAT_CTS(fields []string) {
	s.Total, _ = strconv.Atoi(fields[0])
	s.Baser, _ = strconv.ParseFloat(fields[1], 64)
	s.Baser_ncl, _ = strconv.ParseFloat(fields[2], 64)
	s.Baser_ncu, _ = strconv.ParseFloat(fields[3], 64)
	s.Baser_bcl, _ = strconv.ParseFloat(fields[4], 64)
	s.Baser_bcu, _ = strconv.ParseFloat(fields[5], 64)
	s.Fmean, _ = strconv.ParseFloat(fields[6], 64)
	s.Fmean_ncl, _ = strconv.ParseFloat(fields[7], 64)
	s.Fmean_ncu, _ = strconv.ParseFloat(fields[8], 64)
	s.Fmean_bcl, _ = strconv.ParseFloat(fields[9], 64)
	s.Fmean_bcu, _ = strconv.ParseFloat(fields[10], 64)
	s.Acc, _ = strconv.ParseFloat(fields[11], 64)
	s.Acc_ncl, _ = strconv.ParseFloat(fields[12], 64)
	s.Acc_ncu, _ = strconv.ParseFloat(fields[13], 64)
	s.Acc_bcl, _ = strconv.ParseFloat(fields[14], 64)
	s.Acc_bcu, _ = strconv.ParseFloat(fields[15], 64)
	s.Fbias, _ = strconv.ParseFloat(fields[16], 64)
	s.Fbias_bcl, _ = strconv.ParseFloat(fields[17], 64)
	s.Fbias_bcu, _ = strconv.ParseFloat(fields[18], 64)
	s.Pody, _ = strconv.ParseFloat(fields[19], 64)
	s.Pody_ncl, _ = strconv.ParseFloat(fields[20], 64)
	s.Pody_ncu, _ = strconv.ParseFloat(fields[21], 64)
	s.Pody_bcl, _ = strconv.ParseFloat(fields[22], 64)
	s.Pody_bcu, _ = strconv.ParseFloat(fields[23], 64)
	s.Podn, _ = strconv.ParseFloat(fields[24], 64)
	s.Podn_ncl, _ = strconv.ParseFloat(fields[25], 64)
	s.Podn_ncu, _ = strconv.ParseFloat(fields[26], 64)
	s.Podn_bcl, _ = strconv.ParseFloat(fields[27], 64)
	s.Podn_bcu, _ = strconv.ParseFloat(fields[28], 64)
	s.Pofd, _ = strconv.ParseFloat(fields[29], 64)
	s.Pofd_ncl, _ = strconv.ParseFloat(fields[30], 64)
	s.Pofd_ncu, _ = strconv.ParseFloat(fields[31], 64)
	s.Pofd_bcl, _ = strconv.ParseFloat(fields[32], 64)
	s.Pofd_bcu, _ = strconv.ParseFloat(fields[33], 64)
	s.Far, _ = strconv.ParseFloat(fields[34], 64)
	s.Far_ncl, _ = strconv.ParseFloat(fields[35], 64)
	s.Far_ncu, _ = strconv.ParseFloat(fields[36], 64)
	s.Far_bcl, _ = strconv.ParseFloat(fields[37], 64)
	s.Far_bcu, _ = strconv.ParseFloat(fields[38], 64)
	s.Csi, _ = strconv.ParseFloat(fields[39], 64)
	s.Csi_ncl, _ = strconv.ParseFloat(fields[40], 64)
	s.Csi_ncu, _ = strconv.ParseFloat(fields[41], 64)
	s.Csi_bcl, _ = strconv.ParseFloat(fields[42], 64)
	s.Csi_bcu, _ = strconv.ParseFloat(fields[43], 64)
	s.Gss, _ = strconv.ParseFloat(fields[44], 64)
	s.Gss_bcl, _ = strconv.ParseFloat(fields[45], 64)
	s.Gss_bcu, _ = strconv.ParseFloat(fields[46], 64)
	s.Hk, _ = strconv.ParseFloat(fields[47], 64)
	s.Hk_ncl, _ = strconv.ParseFloat(fields[48], 64)
	s.Hk_ncu, _ = strconv.ParseFloat(fields[49], 64)
	s.Hk_bcl, _ = strconv.ParseFloat(fields[50], 64)
	s.Hk_bcu, _ = strconv.ParseFloat(fields[51], 64)
	s.Hss, _ = strconv.ParseFloat(fields[52], 64)
	s.Hss_bcl, _ = strconv.ParseFloat(fields[53], 64)
	s.Hss_bcu, _ = strconv.ParseFloat(fields[54], 64)
	s.Odds, _ = strconv.ParseFloat(fields[55], 64)
	s.Odds_ncl, _ = strconv.ParseFloat(fields[56], 64)
	s.Odds_ncu, _ = strconv.ParseFloat(fields[57], 64)
	s.Odds_bcl, _ = strconv.ParseFloat(fields[58], 64)
	s.Odds_bcu, _ = strconv.ParseFloat(fields[59], 64)
	s.Lodds, _ = strconv.ParseFloat(fields[60], 64)
	s.Lodds_ncl, _ = strconv.ParseFloat(fields[61], 64)
	s.Lodds_ncu, _ = strconv.ParseFloat(fields[62], 64)
	s.Lodds_bcl, _ = strconv.ParseFloat(fields[63], 64)
	s.Lodds_bcu, _ = strconv.ParseFloat(fields[64], 64)
	s.Orss, _ = strconv.ParseFloat(fields[65], 64)
	s.Orss_ncl, _ = strconv.ParseFloat(fields[66], 64)
	s.Orss_ncu, _ = strconv.ParseFloat(fields[67], 64)
	s.Orss_bcl, _ = strconv.ParseFloat(fields[68], 64)
	s.Orss_bcu, _ = strconv.ParseFloat(fields[69], 64)
	s.Eds, _ = strconv.ParseFloat(fields[70], 64)
	s.Eds_ncl, _ = strconv.ParseFloat(fields[71], 64)
	s.Eds_ncu, _ = strconv.ParseFloat(fields[72], 64)
	s.Eds_bcl, _ = strconv.ParseFloat(fields[73], 64)
	s.Eds_bcu, _ = strconv.ParseFloat(fields[74], 64)
	s.Seds, _ = strconv.ParseFloat(fields[75], 64)
	s.Seds_ncl, _ = strconv.ParseFloat(fields[76], 64)
	s.Seds_ncu, _ = strconv.ParseFloat(fields[77], 64)
	s.Seds_bcl, _ = strconv.ParseFloat(fields[78], 64)
	s.Seds_bcu, _ = strconv.ParseFloat(fields[79], 64)
	s.Edi, _ = strconv.ParseFloat(fields[80], 64)
	s.Edi_ncl, _ = strconv.ParseFloat(fields[81], 64)
	s.Edi_ncu, _ = strconv.ParseFloat(fields[82], 64)
	s.Edi_bcl, _ = strconv.ParseFloat(fields[83], 64)
	s.Edi_bcu, _ = strconv.ParseFloat(fields[84], 64)
	s.Sedi, _ = strconv.ParseFloat(fields[85], 64)
	s.Sedi_ncl, _ = strconv.ParseFloat(fields[86], 64)
	s.Sedi_ncu, _ = strconv.ParseFloat(fields[87], 64)
	s.Sedi_bcl, _ = strconv.ParseFloat(fields[88], 64)
	s.Sedi_bcu, _ = strconv.ParseFloat(fields[89], 64)
	s.Bagss, _ = strconv.ParseFloat(fields[90], 64)
	s.Bagss_bcl, _ = strconv.ParseFloat(fields[91], 64)
	s.Bagss_bcu, _ = strconv.ParseFloat(fields[92], 64)
	s.Hss_ec, _ = strconv.ParseFloat(fields[93], 64)
	s.Hss_ec_bcl, _ = strconv.ParseFloat(fields[94], 64)
	s.Hss_ec_bcu, _ = strconv.ParseFloat(fields[95], 64)
	s.Ec_value, _ = strconv.ParseFloat(fields[96], 64)
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
	s.Fbar_ncl, _ = strconv.ParseFloat(fields[12], 64)
	s.Fbar_ncu, _ = strconv.ParseFloat(fields[13], 64)
	s.Fstdev, _ = strconv.ParseFloat(fields[14], 64)
	s.Fstdev_ncl, _ = strconv.ParseFloat(fields[15], 64)
	s.Fstdev_ncu, _ = strconv.ParseFloat(fields[16], 64)
	s.Obar_ncl, _ = strconv.ParseFloat(fields[17], 64)
	s.Obar_ncu, _ = strconv.ParseFloat(fields[18], 64)
	s.Ostdev, _ = strconv.ParseFloat(fields[19], 64)
	s.Ostdev_ncl, _ = strconv.ParseFloat(fields[20], 64)
	s.Ostdev_ncu, _ = strconv.ParseFloat(fields[21], 64)
	s.Pr_corr, _ = strconv.ParseFloat(fields[22], 64)
	s.Pr_corr_ncl, _ = strconv.ParseFloat(fields[23], 64)
	s.Pr_corr_ncu, _ = strconv.ParseFloat(fields[24], 64)
	s.Me, _ = strconv.ParseFloat(fields[25], 64)
	s.Me_ncl, _ = strconv.ParseFloat(fields[26], 64)
	s.Me_ncu, _ = strconv.ParseFloat(fields[27], 64)
	s.Estdev, _ = strconv.ParseFloat(fields[28], 64)
	s.Estdev_ncl, _ = strconv.ParseFloat(fields[29], 64)
	s.Estdev_ncu, _ = strconv.ParseFloat(fields[30], 64)
	s.Mbias, _ = strconv.ParseFloat(fields[31], 64)
	s.Mse, _ = strconv.ParseFloat(fields[32], 64)
	s.Bcmse, _ = strconv.ParseFloat(fields[33], 64)
	s.Rmse, _ = strconv.ParseFloat(fields[34], 64)
}

func (s *STAT_NBRCTC) fill_STAT_NBRCTC(fields []string) {
	s.Total, _ = strconv.Atoi(fields[0])
	s.Fy_oy, _ = strconv.ParseFloat(fields[1], 64)
	s.Fy_on, _ = strconv.ParseFloat(fields[2], 64)
	s.Fn_oy, _ = strconv.ParseFloat(fields[3], 64)
	s.Fn_on, _ = strconv.ParseFloat(fields[4], 64)
}

func (s *STAT_NBRCTS) fill_STAT_NBRCTS(fields []string) {
	s.Total, _ = strconv.Atoi(fields[0])
	s.Baser, _ = strconv.ParseFloat(fields[1], 64)
	s.Baser_ncl, _ = strconv.ParseFloat(fields[2], 64)
	s.Baser_ncu, _ = strconv.ParseFloat(fields[3], 64)
	s.Baser_bcl, _ = strconv.ParseFloat(fields[4], 64)
	s.Baser_bcu, _ = strconv.ParseFloat(fields[5], 64)
	s.Fmean, _ = strconv.ParseFloat(fields[6], 64)
	s.Fmean_ncl, _ = strconv.ParseFloat(fields[7], 64)
	s.Fmean_ncu, _ = strconv.ParseFloat(fields[8], 64)
	s.Fmean_bcl, _ = strconv.ParseFloat(fields[9], 64)
	s.Fmean_bcu, _ = strconv.ParseFloat(fields[10], 64)
	s.Acc, _ = strconv.ParseFloat(fields[11], 64)
	s.Acc_ncl, _ = strconv.ParseFloat(fields[12], 64)
	s.Acc_ncu, _ = strconv.ParseFloat(fields[13], 64)
	s.Acc_bcl, _ = strconv.ParseFloat(fields[14], 64)
	s.Acc_bcu, _ = strconv.ParseFloat(fields[15], 64)
	s.Fbias, _ = strconv.ParseFloat(fields[16], 64)
	s.Fbias_bcl, _ = strconv.ParseFloat(fields[17], 64)
	s.Fbias_bcu, _ = strconv.ParseFloat(fields[18], 64)
	s.Pody, _ = strconv.ParseFloat(fields[19], 64)
	s.Pody_ncl, _ = strconv.ParseFloat(fields[20], 64)
	s.Pody_ncu, _ = strconv.ParseFloat(fields[21], 64)
	s.Pody_bcl, _ = strconv.ParseFloat(fields[22], 64)
	s.Pody_bcu, _ = strconv.ParseFloat(fields[23], 64)
	s.Podn, _ = strconv.ParseFloat(fields[24], 64)
	s.Podn_ncl, _ = strconv.ParseFloat(fields[25], 64)
	s.Podn_ncu, _ = strconv.ParseFloat(fields[26], 64)
	s.Podn_bcl, _ = strconv.ParseFloat(fields[27], 64)
	s.Podn_bcu, _ = strconv.ParseFloat(fields[28], 64)
	s.Pofd, _ = strconv.ParseFloat(fields[29], 64)
	s.Pofd_ncl, _ = strconv.ParseFloat(fields[30], 64)
	s.Pofd_ncu, _ = strconv.ParseFloat(fields[31], 64)
	s.Pofd_bcl, _ = strconv.ParseFloat(fields[32], 64)
	s.Pofd_bcu, _ = strconv.ParseFloat(fields[33], 64)
	s.Far, _ = strconv.ParseFloat(fields[34], 64)
	s.Far_ncl, _ = strconv.ParseFloat(fields[35], 64)
	s.Far_ncu, _ = strconv.ParseFloat(fields[36], 64)
	s.Far_bcl, _ = strconv.ParseFloat(fields[37], 64)
	s.Far_bcu, _ = strconv.ParseFloat(fields[38], 64)
	s.Csi, _ = strconv.ParseFloat(fields[39], 64)
	s.Csi_ncl, _ = strconv.ParseFloat(fields[40], 64)
	s.Csi_ncu, _ = strconv.ParseFloat(fields[41], 64)
	s.Csi_bcl, _ = strconv.ParseFloat(fields[42], 64)
	s.Csi_bcu, _ = strconv.ParseFloat(fields[43], 64)
	s.Gss, _ = strconv.ParseFloat(fields[44], 64)
	s.Gss_bcl, _ = strconv.ParseFloat(fields[45], 64)
	s.Gss_bcu, _ = strconv.ParseFloat(fields[46], 64)
	s.Hk, _ = strconv.ParseFloat(fields[47], 64)
	s.Hk_ncl, _ = strconv.ParseFloat(fields[48], 64)
	s.Hk_ncu, _ = strconv.ParseFloat(fields[49], 64)
	s.Hk_bcl, _ = strconv.ParseFloat(fields[50], 64)
	s.Hk_bcu, _ = strconv.ParseFloat(fields[51], 64)
	s.Hss, _ = strconv.ParseFloat(fields[52], 64)
	s.Hss_bcl, _ = strconv.ParseFloat(fields[53], 64)
	s.Hss_bcu, _ = strconv.ParseFloat(fields[54], 64)
	s.Odds, _ = strconv.ParseFloat(fields[55], 64)
	s.Odds_ncl, _ = strconv.ParseFloat(fields[56], 64)
	s.Odds_ncu, _ = strconv.ParseFloat(fields[57], 64)
	s.Odds_bcl, _ = strconv.ParseFloat(fields[58], 64)
	s.Odds_bcu, _ = strconv.ParseFloat(fields[59], 64)
	s.Lodds, _ = strconv.ParseFloat(fields[60], 64)
	s.Lodds_ncl, _ = strconv.ParseFloat(fields[61], 64)
	s.Lodds_ncu, _ = strconv.ParseFloat(fields[62], 64)
	s.Lodds_bcl, _ = strconv.ParseFloat(fields[63], 64)
	s.Lodds_bcu, _ = strconv.ParseFloat(fields[64], 64)
	s.Orss, _ = strconv.ParseFloat(fields[65], 64)
	s.Orss_ncl, _ = strconv.ParseFloat(fields[66], 64)
	s.Orss_ncu, _ = strconv.ParseFloat(fields[67], 64)
	s.Orss_bcl, _ = strconv.ParseFloat(fields[68], 64)
	s.Orss_bcu, _ = strconv.ParseFloat(fields[69], 64)
	s.Eds, _ = strconv.ParseFloat(fields[70], 64)
	s.Eds_ncl, _ = strconv.ParseFloat(fields[71], 64)
	s.Eds_ncu, _ = strconv.ParseFloat(fields[72], 64)
	s.Eds_bcl, _ = strconv.ParseFloat(fields[73], 64)
	s.Eds_bcu, _ = strconv.ParseFloat(fields[74], 64)
	s.Seds, _ = strconv.ParseFloat(fields[75], 64)
	s.Seds_ncl, _ = strconv.ParseFloat(fields[76], 64)
	s.Seds_ncu, _ = strconv.ParseFloat(fields[77], 64)
	s.Seds_bcl, _ = strconv.ParseFloat(fields[78], 64)
	s.Seds_bcu, _ = strconv.ParseFloat(fields[79], 64)
	s.Edi, _ = strconv.ParseFloat(fields[80], 64)
	s.Edi_ncl, _ = strconv.ParseFloat(fields[81], 64)
	s.Edi_ncu, _ = strconv.ParseFloat(fields[82], 64)
	s.Edi_bcl, _ = strconv.ParseFloat(fields[83], 64)
	s.Edi_bcu, _ = strconv.ParseFloat(fields[84], 64)
	s.Sedi, _ = strconv.ParseFloat(fields[85], 64)
	s.Sedi_ncl, _ = strconv.ParseFloat(fields[86], 64)
	s.Sedi_ncu, _ = strconv.ParseFloat(fields[87], 64)
	s.Sedi_bcl, _ = strconv.ParseFloat(fields[88], 64)
	s.Sedi_bcu, _ = strconv.ParseFloat(fields[89], 64)
	s.Bagss, _ = strconv.ParseFloat(fields[90], 64)
	s.Bagss_bcl, _ = strconv.ParseFloat(fields[91], 64)
	s.Bagss_bcu, _ = strconv.ParseFloat(fields[92], 64)
}

func (s *STAT_ECNT) fill_STAT_ECNT(fields []string) {
	s.Total, _ = strconv.Atoi(fields[0])
	s.N_ens, _ = strconv.Atoi(fields[1])
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

func (s *STAT_RPS) fill_STAT_RPS(fields []string) {
	s.Total, _ = strconv.Atoi(fields[0])
	s.N_prob, _ = strconv.Atoi(fields[1])
	s.Rps_rel, _ = strconv.ParseFloat(fields[2], 64)
	s.Rps_res, _ = strconv.ParseFloat(fields[3], 64)
	s.Rps_unc, _ = strconv.ParseFloat(fields[4], 64)
	s.Rps, _ = strconv.ParseFloat(fields[5], 64)
	s.Rpss, _ = strconv.ParseFloat(fields[6], 64)
	s.Rpss_smpl, _ = strconv.ParseFloat(fields[7], 64)
	s.Rps_comp = fields[8]
}

func (s *STAT_RHIST) fill_STAT_RHIST(fields []string) {
	s.Total, _ = strconv.Atoi(fields[0])
	s.N_rank, _ = strconv.Atoi(fields[1])
	s.Rank_i = fields[2]
}

func (s *TCST_TCDIAG) fill_TCST_TCDIAG(fields []string) {
	s.Total, _ = strconv.Atoi(fields[0])
	s.Index, _ = strconv.Atoi(fields[1])
	s.Diag_source = fields[2]
	s.Track_source = fields[3]
	s.Field_source = fields[4]
	s.N_diag, _ = strconv.Atoi(fields[5])
	s.Diag_i = fields[6]
	s.Value_i = fields[7]
}

func (s *STAT_FHO) fill_STAT_FHO(fields []string) {
	s.Total, _ = strconv.Atoi(fields[0])
	s.F_rate, _ = strconv.ParseFloat(fields[1], 64)
	s.H_rate, _ = strconv.ParseFloat(fields[2], 64)
	s.O_rate, _ = strconv.ParseFloat(fields[3], 64)
}

func (s *STAT_MCTS) fill_STAT_MCTS(fields []string) {
	s.Total, _ = strconv.Atoi(fields[0])
	s.N_cat, _ = strconv.Atoi(fields[1])
	s.Acc, _ = strconv.ParseFloat(fields[2], 64)
	s.Acc_ncl, _ = strconv.ParseFloat(fields[3], 64)
	s.Acc_ncu, _ = strconv.ParseFloat(fields[4], 64)
	s.Acc_bcl, _ = strconv.ParseFloat(fields[5], 64)
	s.Acc_bcu, _ = strconv.ParseFloat(fields[6], 64)
	s.Hk, _ = strconv.ParseFloat(fields[7], 64)
	s.Hk_bcl, _ = strconv.ParseFloat(fields[8], 64)
	s.Hk_bcu, _ = strconv.ParseFloat(fields[9], 64)
	s.Hss, _ = strconv.ParseFloat(fields[10], 64)
	s.Hss_bcl, _ = strconv.ParseFloat(fields[11], 64)
	s.Hss_bcu, _ = strconv.ParseFloat(fields[12], 64)
	s.Ger, _ = strconv.ParseFloat(fields[13], 64)
	s.Ger_bcl, _ = strconv.ParseFloat(fields[14], 64)
	s.Ger_bcu, _ = strconv.ParseFloat(fields[15], 64)
	s.Hss_ec, _ = strconv.ParseFloat(fields[16], 64)
	s.Hss_ec_bcl, _ = strconv.ParseFloat(fields[17], 64)
	s.Hss_ec_bcu, _ = strconv.ParseFloat(fields[18], 64)
	s.Ec_value, _ = strconv.ParseFloat(fields[19], 64)
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
	s.Fy, _ = strconv.Atoi(fields[1])
	s.Oy, _ = strconv.Atoi(fields[2])
	s.Fbias, _ = strconv.ParseFloat(fields[3], 64)
	s.Baddeley, _ = strconv.ParseFloat(fields[4], 64)
	s.Hausdorff, _ = strconv.ParseFloat(fields[5], 64)
	s.Med_fo, _ = strconv.ParseFloat(fields[6], 64)
	s.Med_of, _ = strconv.ParseFloat(fields[7], 64)
	s.Med_min, _ = strconv.ParseFloat(fields[8], 64)
	s.Med_max, _ = strconv.ParseFloat(fields[9], 64)
	s.Med_mean, _ = strconv.ParseFloat(fields[10], 64)
	s.Fom_fo, _ = strconv.ParseFloat(fields[11], 64)
	s.Fom_of, _ = strconv.ParseFloat(fields[12], 64)
	s.Fom_min, _ = strconv.ParseFloat(fields[13], 64)
	s.Fom_max, _ = strconv.ParseFloat(fields[14], 64)
	s.Fom_mean, _ = strconv.ParseFloat(fields[15], 64)
	s.Zhu_fo, _ = strconv.ParseFloat(fields[16], 64)
	s.Zhu_of, _ = strconv.ParseFloat(fields[17], 64)
	s.Zhu_min, _ = strconv.ParseFloat(fields[18], 64)
	s.Zhu_max, _ = strconv.ParseFloat(fields[19], 64)
	s.Zhu_mean, _ = strconv.ParseFloat(fields[20], 64)
	s.G, _ = strconv.ParseFloat(fields[21], 64)
	s.Gbeta, _ = strconv.ParseFloat(fields[22], 64)
	s.Beta_value, _ = strconv.ParseFloat(fields[23], 64)
}

func (s *STAT_ECLV) fill_STAT_ECLV(fields []string) {
	s.Total, _ = strconv.Atoi(fields[0])
	s.Baser, _ = strconv.ParseFloat(fields[1], 64)
	s.Value_baser, _ = strconv.ParseFloat(fields[2], 64)
	s.N_pts = fields[3]
	s.Cl_i = fields[4]
	s.Value_i = fields[5]
}

func (s *STAT_PHIST) fill_STAT_PHIST(fields []string) {
	s.Total, _ = strconv.Atoi(fields[0])
	s.Bin_size, _ = strconv.ParseFloat(fields[1], 64)
	s.N_bin, _ = strconv.Atoi(fields[2])
	s.Bin_i = fields[3]
}

func (s *STAT_RELP) fill_STAT_RELP(fields []string) {
	s.Total, _ = strconv.Atoi(fields[0])
	s.N_ens, _ = strconv.Atoi(fields[1])
	s.Relp_i = fields[2]
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

func (s *STAT_ISC) fill_STAT_ISC(fields []string) {
	s.Total, _ = strconv.Atoi(fields[0])
	s.Tile_dim, _ = strconv.Atoi(fields[1])
	s.Tile_xll, _ = strconv.Atoi(fields[2])
	s.Tile_yll, _ = strconv.Atoi(fields[3])
	s.Nscale, _ = strconv.Atoi(fields[4])
	s.Iscale, _ = strconv.Atoi(fields[5])
	s.Mse, _ = strconv.ParseFloat(fields[6], 64)
	s.Isc, _ = strconv.ParseFloat(fields[7], 64)
	s.Fenergy2, _ = strconv.ParseFloat(fields[8], 64)
	s.Oenergy2, _ = strconv.ParseFloat(fields[9], 64)
	s.Baser, _ = strconv.ParseFloat(fields[10], 64)
	s.Fbias, _ = strconv.ParseFloat(fields[11], 64)
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
	s.Total_dir, _ = strconv.ParseFloat(fields[10], 64)
	s.Dira_me, _ = strconv.ParseFloat(fields[11], 64)
	s.Dira_mae, _ = strconv.ParseFloat(fields[12], 64)
	s.Dira_mse, _ = strconv.ParseFloat(fields[13], 64)
}

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

func (s *STAT_GRAD) fill_STAT_GRAD(fields []string) {
	s.Total, _ = strconv.Atoi(fields[0])
	s.Fgbar, _ = strconv.ParseFloat(fields[1], 64)
	s.Ogbar, _ = strconv.ParseFloat(fields[2], 64)
	s.Mgbar, _ = strconv.ParseFloat(fields[3], 64)
	s.Egbar, _ = strconv.ParseFloat(fields[4], 64)
	s.S1, _ = strconv.ParseFloat(fields[5], 64)
	s.S1_og, _ = strconv.ParseFloat(fields[6], 64)
	s.Fgog_ratio, _ = strconv.ParseFloat(fields[7], 64)
	s.Dx, _ = strconv.ParseFloat(fields[8], 64)
	s.Dy, _ = strconv.ParseFloat(fields[9], 64)
}

func (s *STAT_PJC) fill_STAT_PJC(fields []string) {
	s.Total, _ = strconv.Atoi(fields[0])
	s.N_thresh, _ = strconv.Atoi(fields[1])
	s.Thresh_i = fields[2]
	s.Oy_tp_i = fields[3]
	s.On_tp_i = fields[4]
	s.Calibration_i = fields[5]
	s.Refinement_i = fields[6]
	s.Likelihood_i, _ = strconv.ParseFloat(fields[7], 64)
	s.Baser_i = fields[8]
}

func (s *STAT_GENMPR) fill_STAT_GENMPR(fields []string) {
	s.Total, _ = strconv.Atoi(fields[0])
	s.Index, _ = strconv.Atoi(fields[1])
	s.Storm_id = fields[2]
	s.Prob_lead = fields[3]
	s.Prob_val, _ = strconv.ParseFloat(fields[4], 64)
	s.Agen_init = fields[5]
	s.Agen_fhr = fields[6]
	s.Agen_lat, _ = strconv.ParseFloat(fields[7], 64)
	s.Agen_lon, _ = strconv.ParseFloat(fields[8], 64)
	s.Agen_dland, _ = strconv.ParseFloat(fields[9], 64)
	s.Bgen_lat, _ = strconv.ParseFloat(fields[10], 64)
	s.Bgen_lon, _ = strconv.ParseFloat(fields[11], 64)
	s.Bgen_dland, _ = strconv.ParseFloat(fields[12], 64)
	s.Gen_dist, _ = strconv.ParseFloat(fields[13], 64)
	s.Gen_tdiff = fields[14]
	s.Init_tdiff = fields[15]
	s.Dev_cat = fields[16]
	s.Ops_cat = fields[17]
}

func (s *MTD_3DPAIR) fill_MTD_3DPAIR(fields []string) {
	s.Object_id = fields[0]
	s.Object_cat = fields[1]
	s.Space_centroid_dist, _ = strconv.ParseFloat(fields[2], 64)
	s.Time_centroid_delta, _ = strconv.ParseFloat(fields[3], 64)
	s.Axis_diff, _ = strconv.ParseFloat(fields[4], 64)
	s.Speed_delta, _ = strconv.ParseFloat(fields[5], 64)
	s.Direction_diff, _ = strconv.ParseFloat(fields[6], 64)
	s.Volume_ratio, _ = strconv.ParseFloat(fields[7], 64)
	s.Start_time_delta, _ = strconv.Atoi(fields[8])
	s.End_time_delta, _ = strconv.Atoi(fields[9])
	s.Intersection_volume, _ = strconv.Atoi(fields[10])
	s.Duration_diff, _ = strconv.Atoi(fields[11])
	s.Interest, _ = strconv.ParseFloat(fields[12], 64)
}

func (s *STAT_CNT) fill_STAT_CNT(fields []string) {
	s.Total, _ = strconv.Atoi(fields[0])
	s.Fbar, _ = strconv.ParseFloat(fields[1], 64)
	s.Fbar_ncl, _ = strconv.ParseFloat(fields[2], 64)
	s.Fbar_ncu, _ = strconv.ParseFloat(fields[3], 64)
	s.Fbar_bcl, _ = strconv.ParseFloat(fields[4], 64)
	s.Fbar_bcu, _ = strconv.ParseFloat(fields[5], 64)
	s.Fstdev, _ = strconv.ParseFloat(fields[6], 64)
	s.Fstdev_ncl, _ = strconv.ParseFloat(fields[7], 64)
	s.Fstdev_ncu, _ = strconv.ParseFloat(fields[8], 64)
	s.Fstdev_bcl, _ = strconv.ParseFloat(fields[9], 64)
	s.Fstdev_bcu, _ = strconv.ParseFloat(fields[10], 64)
	s.Obar, _ = strconv.ParseFloat(fields[11], 64)
	s.Obar_ncl, _ = strconv.ParseFloat(fields[12], 64)
	s.Obar_ncu, _ = strconv.ParseFloat(fields[13], 64)
	s.Obar_bcl, _ = strconv.ParseFloat(fields[14], 64)
	s.Obar_bcu, _ = strconv.ParseFloat(fields[15], 64)
	s.Ostdev, _ = strconv.ParseFloat(fields[16], 64)
	s.Ostdev_ncl, _ = strconv.ParseFloat(fields[17], 64)
	s.Ostdev_ncu, _ = strconv.ParseFloat(fields[18], 64)
	s.Ostdev_bcl, _ = strconv.ParseFloat(fields[19], 64)
	s.Ostdev_bcu, _ = strconv.ParseFloat(fields[20], 64)
	s.Pr_corr, _ = strconv.ParseFloat(fields[21], 64)
	s.Pr_corr_ncl, _ = strconv.ParseFloat(fields[22], 64)
	s.Pr_corr_ncu, _ = strconv.ParseFloat(fields[23], 64)
	s.Pr_corr_bcl, _ = strconv.ParseFloat(fields[24], 64)
	s.Pr_corr_bcu, _ = strconv.ParseFloat(fields[25], 64)
	s.Sp_corr, _ = strconv.ParseFloat(fields[26], 64)
	s.Kt_corr, _ = strconv.ParseFloat(fields[27], 64)
	s.Ranks, _ = strconv.Atoi(fields[28])
	s.Frank_ties, _ = strconv.Atoi(fields[29])
	s.Orank_ties, _ = strconv.Atoi(fields[30])
	s.Me, _ = strconv.ParseFloat(fields[31], 64)
	s.Me_ncl, _ = strconv.ParseFloat(fields[32], 64)
	s.Me_ncu, _ = strconv.ParseFloat(fields[33], 64)
	s.Me_bcl, _ = strconv.ParseFloat(fields[34], 64)
	s.Me_bcu, _ = strconv.ParseFloat(fields[35], 64)
	s.Estdev, _ = strconv.ParseFloat(fields[36], 64)
	s.Estdev_ncl, _ = strconv.ParseFloat(fields[37], 64)
	s.Estdev_ncu, _ = strconv.ParseFloat(fields[38], 64)
	s.Estdev_bcl, _ = strconv.ParseFloat(fields[39], 64)
	s.Estdev_bcu, _ = strconv.ParseFloat(fields[40], 64)
	s.Mbias, _ = strconv.ParseFloat(fields[41], 64)
	s.Mbias_bcl, _ = strconv.ParseFloat(fields[42], 64)
	s.Mbias_bcu, _ = strconv.ParseFloat(fields[43], 64)
	s.Mae, _ = strconv.ParseFloat(fields[44], 64)
	s.Mae_bcl, _ = strconv.ParseFloat(fields[45], 64)
	s.Mae_bcu, _ = strconv.ParseFloat(fields[46], 64)
	s.Mse, _ = strconv.ParseFloat(fields[47], 64)
	s.Mse_bcl, _ = strconv.ParseFloat(fields[48], 64)
	s.Mse_bcu, _ = strconv.ParseFloat(fields[49], 64)
	s.Bcmse, _ = strconv.ParseFloat(fields[50], 64)
	s.Bcmse_bcl, _ = strconv.ParseFloat(fields[51], 64)
	s.Bcmse_bcu, _ = strconv.ParseFloat(fields[52], 64)
	s.Rmse, _ = strconv.ParseFloat(fields[53], 64)
	s.Rmse_bcl, _ = strconv.ParseFloat(fields[54], 64)
	s.Rmse_bcu, _ = strconv.ParseFloat(fields[55], 64)
	s.E10, _ = strconv.ParseFloat(fields[56], 64)
	s.E10_bcl, _ = strconv.ParseFloat(fields[57], 64)
	s.E10_bcu, _ = strconv.ParseFloat(fields[58], 64)
	s.E25, _ = strconv.ParseFloat(fields[59], 64)
	s.E25_bcl, _ = strconv.ParseFloat(fields[60], 64)
	s.E25_bcu, _ = strconv.ParseFloat(fields[61], 64)
	s.E50, _ = strconv.ParseFloat(fields[62], 64)
	s.E50_bcl, _ = strconv.ParseFloat(fields[63], 64)
	s.E50_bcu, _ = strconv.ParseFloat(fields[64], 64)
	s.E75, _ = strconv.ParseFloat(fields[65], 64)
	s.E75_bcl, _ = strconv.ParseFloat(fields[66], 64)
	s.E75_bcu, _ = strconv.ParseFloat(fields[67], 64)
	s.E90, _ = strconv.ParseFloat(fields[68], 64)
	s.E90_bcl, _ = strconv.ParseFloat(fields[69], 64)
	s.E90_bcu, _ = strconv.ParseFloat(fields[70], 64)
	s.Eiqr, _ = strconv.ParseFloat(fields[71], 64)
	s.Eiqr_bcl, _ = strconv.Atoi(fields[72])
	s.Eiqr_bcu, _ = strconv.Atoi(fields[73])
	s.Mad, _ = strconv.ParseFloat(fields[74], 64)
	s.Mad_bcl, _ = strconv.ParseFloat(fields[75], 64)
	s.Mad_bcu, _ = strconv.ParseFloat(fields[76], 64)
	s.Anom_corr, _ = strconv.ParseFloat(fields[77], 64)
	s.Anom_corr_ncl, _ = strconv.ParseFloat(fields[78], 64)
	s.Anom_corr_ncu, _ = strconv.ParseFloat(fields[79], 64)
	s.Anom_corr_bcl, _ = strconv.ParseFloat(fields[80], 64)
	s.Anom_corr_bcu, _ = strconv.ParseFloat(fields[81], 64)
	s.Me2, _ = strconv.ParseFloat(fields[82], 64)
	s.Me2_bcl, _ = strconv.ParseFloat(fields[83], 64)
	s.Me2_bcu, _ = strconv.ParseFloat(fields[84], 64)
	s.Msess, _ = strconv.ParseFloat(fields[85], 64)
	s.Msess_bcl, _ = strconv.ParseFloat(fields[86], 64)
	s.Msess_bcu, _ = strconv.ParseFloat(fields[87], 64)
	s.Rmsfa, _ = strconv.ParseFloat(fields[88], 64)
	s.Rmsfa_bcl, _ = strconv.ParseFloat(fields[89], 64)
	s.Rmsfa_bcu, _ = strconv.ParseFloat(fields[90], 64)
	s.Rmsoa, _ = strconv.ParseFloat(fields[91], 64)
	s.Rmsoa_bcl, _ = strconv.ParseFloat(fields[92], 64)
	s.Rmsoa_bcu, _ = strconv.ParseFloat(fields[93], 64)
	s.Anom_corr_uncntr, _ = strconv.ParseFloat(fields[94], 64)
	s.Anom_corr_uncntr_bcl, _ = strconv.ParseFloat(fields[95], 64)
	s.Anom_corr_uncntr_bcu, _ = strconv.ParseFloat(fields[96], 64)
	s.Si, _ = strconv.ParseFloat(fields[97], 64)
	s.Si_bcl, _ = strconv.ParseFloat(fields[98], 64)
	s.Si_bcu, _ = strconv.ParseFloat(fields[99], 64)
}

func (s *STAT_PSTD) fill_STAT_PSTD(fields []string) {
	s.Total, _ = strconv.Atoi(fields[0])
	s.N_thresh, _ = strconv.Atoi(fields[1])
	s.Baser, _ = strconv.ParseFloat(fields[2], 64)
	s.Baser_ncl, _ = strconv.ParseFloat(fields[3], 64)
	s.Baser_ncu, _ = strconv.ParseFloat(fields[4], 64)
	s.Reliability, _ = strconv.ParseFloat(fields[5], 64)
	s.Resolution, _ = strconv.ParseFloat(fields[6], 64)
	s.Uncertainty, _ = strconv.ParseFloat(fields[7], 64)
	s.Roc_auc, _ = strconv.ParseFloat(fields[8], 64)
	s.Brier, _ = strconv.ParseFloat(fields[9], 64)
	s.Brier_ncl, _ = strconv.ParseFloat(fields[10], 64)
	s.Brier_ncu, _ = strconv.ParseFloat(fields[11], 64)
	s.Briercl, _ = strconv.ParseFloat(fields[12], 64)
	s.Briercl_ncl, _ = strconv.ParseFloat(fields[13], 64)
	s.Briercl_ncu, _ = strconv.ParseFloat(fields[14], 64)
	s.Bss, _ = strconv.ParseFloat(fields[15], 64)
	s.Bss_smpl, _ = strconv.ParseFloat(fields[16], 64)
	s.Thresh_i = fields[17]
}

func (s *STAT_SSIDX) fill_STAT_SSIDX(fields []string) {
	s.Fcst_model = fields[0]
	s.Ref_model = fields[1]
	s.N_init, _ = strconv.Atoi(fields[2])
	s.N_term, _ = strconv.Atoi(fields[3])
	s.N_vld, _ = strconv.Atoi(fields[4])
	s.Ss_index, _ = strconv.ParseFloat(fields[5], 64)
}

func (s *MODE_CTS) fill_MODE_CTS(fields []string) {
}

func (s *MTD_2DSINGLE) fill_MTD_2DSINGLE(fields []string) {
	s.Object_id = fields[0]
	s.Object_cat = fields[1]
	s.Time_index, _ = strconv.Atoi(fields[2])
	s.Area, _ = strconv.Atoi(fields[3])
	s.Centroid_x, _ = strconv.ParseFloat(fields[4], 64)
	s.Centroid_y, _ = strconv.ParseFloat(fields[5], 64)
	s.Centroid_lat, _ = strconv.ParseFloat(fields[6], 64)
	s.Centroid_lon, _ = strconv.ParseFloat(fields[7], 64)
	s.Axis_ang, _ = strconv.ParseFloat(fields[8], 64)
	s.Intensity_10, _ = strconv.ParseFloat(fields[9], 64)
	s.Intensity_25, _ = strconv.ParseFloat(fields[10], 64)
	s.Intensity_50, _ = strconv.ParseFloat(fields[11], 64)
	s.Intensity_75, _ = strconv.ParseFloat(fields[12], 64)
	s.Intensity_90, _ = strconv.ParseFloat(fields[13], 64)
	s.Intensity_user = fields[14]
}

func (s *MTD_3DSINGLE) fill_MTD_3DSINGLE(fields []string) {
	s.Object_id = fields[0]
	s.Object_cat = fields[1]
	s.Centroid_x, _ = strconv.ParseFloat(fields[2], 64)
	s.Centroid_y, _ = strconv.ParseFloat(fields[3], 64)
	s.Centroid_t, _ = strconv.ParseFloat(fields[4], 64)
	s.Centroid_lat, _ = strconv.ParseFloat(fields[5], 64)
	s.Centroid_lon, _ = strconv.ParseFloat(fields[6], 64)
	s.X_dot, _ = strconv.ParseFloat(fields[7], 64)
	s.Y_dot, _ = strconv.ParseFloat(fields[8], 64)
	s.Axis_ang, _ = strconv.ParseFloat(fields[9], 64)
	s.Volume, _ = strconv.Atoi(fields[10])
	s.Start_time, _ = strconv.Atoi(fields[11])
	s.End_time, _ = strconv.Atoi(fields[12])
	s.Cdist_travelled, _ = strconv.ParseFloat(fields[13], 64)
	s.Intensity_10, _ = strconv.ParseFloat(fields[14], 64)
	s.Intensity_25, _ = strconv.ParseFloat(fields[15], 64)
	s.Intensity_50, _ = strconv.ParseFloat(fields[16], 64)
	s.Intensity_75, _ = strconv.ParseFloat(fields[17], 64)
	s.Intensity_90, _ = strconv.ParseFloat(fields[18], 64)
	s.Intensity_user = fields[19]
}

func (s *STAT_MCTC) fill_STAT_MCTC(fields []string) {
	s.Total, _ = strconv.Atoi(fields[0])
	s.N_cat, _ = strconv.Atoi(fields[1])
	s.Fi_oi = fields[2]
	s.Ec_value, _ = strconv.ParseFloat(fields[3], 64)
}


//getDocForId functions
func GetDocForId(fileLineType string, metaDataMap map[string]interface{}, headerData []string, dataData []string, dataKey string) (interface{}, error) {
	doc := make(map[string] interface{})
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
		return nil, errors.New("Unknown file_line type:" + fileLineType)
	}
	return doc, nil
}


//addDataElement functions
func AddDataElement(dataKey string, fileLineType string, dataData []string, doc *map[string]interface{}) error {
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
		return errors.New("Unknown file_line type:" + fileLineType)
	}
	return nil
}


	// DateFieldNames is a list of the date fields that need to be converted to epochs
var DateFieldNames = []string{"FCST_VALID_BEG", "FCST_VALID_END", "OBS_VALID_BEG", "OBS_VALID_END"}

