package structColumnTypes

import (
	"errors"
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

// fillHeader functions
func (s *STAT_RHIST) fill_STAT_RHIST_Header(fields []string, doc *map[string]interface{}) {
	// fill the met fields leaving out "" and NA values
	if fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != "" && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != "" && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != "" && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != "" && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != "" && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != "" && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != "" && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != "" && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != "" && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != "" && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != "" && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != "" && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != "" && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != "" && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != "" && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != "" && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_PHIST) fill_STAT_PHIST_Header(fields []string, doc *map[string]interface{}) {
	// fill the met fields leaving out "" and NA values
	if fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != "" && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != "" && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != "" && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != "" && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != "" && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != "" && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != "" && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != "" && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != "" && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != "" && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != "" && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != "" && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != "" && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != "" && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != "" && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != "" && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *MODE_CTS) fill_MODE_CTS_Header(fields []string, doc *map[string]interface{}) {
	// fill the met fields leaving out "" and NA values
	if fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != "" && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != "" && fields[0] != "NA" {
		(*doc)["N_VALID"] = fields[2]
	}
	if fields[3] != "" && fields[0] != "NA" {
		(*doc)["GRID_RES"] = fields[3]
	}
	if fields[4] != "" && fields[0] != "NA" {
		(*doc)["DESC"] = fields[4]
	}
	if fields[6] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID"] = fields[6]
	}
	if fields[7] != "" && fields[0] != "NA" {
		(*doc)["FCST_ACCUM"] = fields[7]
	}
	if fields[8] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[8]
	}
	if fields[9] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID"] = fields[9]
	}
	if fields[10] != "" && fields[0] != "NA" {
		(*doc)["OBS_ACCUM"] = fields[10]
	}
	if fields[11] != "" && fields[0] != "NA" {
		(*doc)["FCST_RAD"] = fields[11]
	}
	if fields[12] != "" && fields[0] != "NA" {
		(*doc)["FCST_THR"] = fields[12]
	}
	if fields[13] != "" && fields[0] != "NA" {
		(*doc)["OBS_RAD"] = fields[13]
	}
	if fields[14] != "" && fields[0] != "NA" {
		(*doc)["OBS_THR"] = fields[14]
	}
	if fields[15] != "" && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[15]
	}
	if fields[16] != "" && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[16]
	}
	if fields[18] != "" && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[18]
	}
	if fields[19] != "" && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[19]
	}
	if fields[20] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[20]
	}
	if fields[21] != "" && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[21]
	}
	if fields[22] != "" && fields[0] != "NA" {
		(*doc)["FIELD"] = fields[22]
	}
	if fields[23] != "" && fields[0] != "NA" {
		(*doc)["TOTAL"] = fields[23]
	}
	if fields[24] != "" && fields[0] != "NA" {
		(*doc)["FY_OY"] = fields[24]
	}
	if fields[25] != "" && fields[0] != "NA" {
		(*doc)["FY_ON"] = fields[25]
	}
	if fields[26] != "" && fields[0] != "NA" {
		(*doc)["FN_OY"] = fields[26]
	}
	if fields[27] != "" && fields[0] != "NA" {
		(*doc)["FN_ON"] = fields[27]
	}
	if fields[28] != "" && fields[0] != "NA" {
		(*doc)["BASER"] = fields[28]
	}
	if fields[29] != "" && fields[0] != "NA" {
		(*doc)["FMEAN"] = fields[29]
	}
	if fields[30] != "" && fields[0] != "NA" {
		(*doc)["ACC"] = fields[30]
	}
	if fields[31] != "" && fields[0] != "NA" {
		(*doc)["FBIAS"] = fields[31]
	}
	if fields[32] != "" && fields[0] != "NA" {
		(*doc)["PODY"] = fields[32]
	}
	if fields[33] != "" && fields[0] != "NA" {
		(*doc)["PODN"] = fields[33]
	}
	if fields[34] != "" && fields[0] != "NA" {
		(*doc)["POFD"] = fields[34]
	}
	if fields[35] != "" && fields[0] != "NA" {
		(*doc)["FAR"] = fields[35]
	}
	if fields[36] != "" && fields[0] != "NA" {
		(*doc)["CSI"] = fields[36]
	}
	if fields[37] != "" && fields[0] != "NA" {
		(*doc)["GSS"] = fields[37]
	}
	if fields[38] != "" && fields[0] != "NA" {
		(*doc)["HK"] = fields[38]
	}
	if fields[39] != "" && fields[0] != "NA" {
		(*doc)["HSS"] = fields[39]
	}
	if fields[40] != "" && fields[0] != "NA" {
		(*doc)["ODDS"] = fields[40]
	}
	if fields[41] != "" && fields[0] != "NA" {
		(*doc)["LODDS"] = fields[41]
	}
	if fields[42] != "" && fields[0] != "NA" {
		(*doc)["ORSS"] = fields[42]
	}
	if fields[43] != "" && fields[0] != "NA" {
		(*doc)["EDS"] = fields[43]
	}
	if fields[44] != "" && fields[0] != "NA" {
		(*doc)["SEDS"] = fields[44]
	}
	if fields[45] != "" && fields[0] != "NA" {
		(*doc)["EDI"] = fields[45]
	}
	if fields[46] != "" && fields[0] != "NA" {
		(*doc)["SEDI"] = fields[46]
	}
	if fields[47] != "" && fields[0] != "NA" {
		(*doc)["BAGSS"] = fields[47]
	}
}

func (s *STAT_SSIDX) fill_STAT_SSIDX_Header(fields []string, doc *map[string]interface{}) {
	// fill the met fields leaving out "" and NA values
	if fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != "" && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != "" && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != "" && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != "" && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != "" && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != "" && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != "" && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != "" && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != "" && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != "" && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != "" && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != "" && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != "" && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != "" && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != "" && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != "" && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *MTD_3DPAIR) fill_MTD_3DPAIR_Header(fields []string, doc *map[string]interface{}) {
	// fill the met fields leaving out "" and NA values
	if fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != "" && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != "" && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID"] = fields[4]
	}
	if fields[5] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[5]
	}
	if fields[6] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID"] = fields[6]
	}
	if fields[7] != "" && fields[0] != "NA" {
		(*doc)["T_DELTA"] = fields[7]
	}
	if fields[8] != "" && fields[0] != "NA" {
		(*doc)["FCST_T_BEG"] = fields[8]
	}
	if fields[9] != "" && fields[0] != "NA" {
		(*doc)["FCST_T_END"] = fields[9]
	}
	if fields[10] != "" && fields[0] != "NA" {
		(*doc)["FCST_RAD"] = fields[10]
	}
	if fields[11] != "" && fields[0] != "NA" {
		(*doc)["FCST_THR"] = fields[11]
	}
	if fields[12] != "" && fields[0] != "NA" {
		(*doc)["OBS_T_BEG"] = fields[12]
	}
	if fields[13] != "" && fields[0] != "NA" {
		(*doc)["OBS_T_END"] = fields[13]
	}
	if fields[14] != "" && fields[0] != "NA" {
		(*doc)["OBS_RAD"] = fields[14]
	}
	if fields[15] != "" && fields[0] != "NA" {
		(*doc)["OBS_THR"] = fields[15]
	}
	if fields[16] != "" && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[16]
	}
	if fields[17] != "" && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[17]
	}
	if fields[19] != "" && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[19]
	}
	if fields[20] != "" && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[20]
	}
	if fields[21] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[21]
	}
}

func (s *TCST_TCDIAG) fill_TCST_TCDIAG_Header(fields []string, doc *map[string]interface{}) {
	// fill the met fields leaving out "" and NA values
	if fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != "" && fields[0] != "NA" {
		(*doc)["AMODEL"] = fields[1]
	}
	if fields[2] != "" && fields[0] != "NA" {
		(*doc)["BMODEL"] = fields[2]
	}
	if fields[3] != "" && fields[0] != "NA" {
		(*doc)["DESC"] = fields[3]
	}
	if fields[4] != "" && fields[0] != "NA" {
		(*doc)["STORM_ID"] = fields[4]
	}
	if fields[5] != "" && fields[0] != "NA" {
		(*doc)["BASIN"] = fields[5]
	}
	if fields[6] != "" && fields[0] != "NA" {
		(*doc)["CYCLONE"] = fields[6]
	}
	if fields[7] != "" && fields[0] != "NA" {
		(*doc)["STORM_NAME"] = fields[7]
	}
	if fields[8] != "" && fields[0] != "NA" {
		(*doc)["INIT"] = fields[8]
	}
	if fields[9] != "" && fields[0] != "NA" {
		(*doc)["LEAD"] = fields[9]
	}
	if fields[10] != "" && fields[0] != "NA" {
		(*doc)["VALID"] = fields[10]
	}
	if fields[11] != "" && fields[0] != "NA" {
		(*doc)["INIT_MASK"] = fields[11]
	}
	if fields[12] != "" && fields[0] != "NA" {
		(*doc)["VALID_MASK"] = fields[12]
	}
	if fields[13] != "" && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[13]
	}
}

func (s *STAT_MCTS) fill_STAT_MCTS_Header(fields []string, doc *map[string]interface{}) {
	// fill the met fields leaving out "" and NA values
	if fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != "" && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != "" && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != "" && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != "" && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != "" && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != "" && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != "" && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != "" && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != "" && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != "" && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != "" && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != "" && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != "" && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != "" && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != "" && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != "" && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_ECNT) fill_STAT_ECNT_Header(fields []string, doc *map[string]interface{}) {
	// fill the met fields leaving out "" and NA values
	if fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != "" && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != "" && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != "" && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != "" && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != "" && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != "" && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != "" && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != "" && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != "" && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != "" && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != "" && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != "" && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != "" && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != "" && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != "" && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != "" && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_RPS) fill_STAT_RPS_Header(fields []string, doc *map[string]interface{}) {
	// fill the met fields leaving out "" and NA values
	if fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != "" && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != "" && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != "" && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != "" && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != "" && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != "" && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != "" && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != "" && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != "" && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != "" && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != "" && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != "" && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != "" && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != "" && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != "" && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != "" && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_SAL1L2) fill_STAT_SAL1L2_Header(fields []string, doc *map[string]interface{}) {
	// fill the met fields leaving out "" and NA values
	if fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != "" && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != "" && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != "" && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != "" && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != "" && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != "" && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != "" && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != "" && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != "" && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != "" && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != "" && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != "" && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != "" && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != "" && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != "" && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != "" && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_VL1L2) fill_STAT_VL1L2_Header(fields []string, doc *map[string]interface{}) {
	// fill the met fields leaving out "" and NA values
	if fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != "" && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != "" && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != "" && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != "" && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != "" && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != "" && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != "" && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != "" && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != "" && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != "" && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != "" && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != "" && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != "" && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != "" && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != "" && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != "" && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_SEEPS) fill_STAT_SEEPS_Header(fields []string, doc *map[string]interface{}) {
	// fill the met fields leaving out "" and NA values
	if fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != "" && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != "" && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != "" && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != "" && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != "" && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != "" && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != "" && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != "" && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != "" && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != "" && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != "" && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != "" && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != "" && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != "" && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != "" && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != "" && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_NBRCNT) fill_STAT_NBRCNT_Header(fields []string, doc *map[string]interface{}) {
	// fill the met fields leaving out "" and NA values
	if fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != "" && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != "" && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != "" && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != "" && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != "" && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != "" && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != "" && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != "" && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != "" && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != "" && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != "" && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != "" && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != "" && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != "" && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != "" && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != "" && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_PSTD) fill_STAT_PSTD_Header(fields []string, doc *map[string]interface{}) {
	// fill the met fields leaving out "" and NA values
	if fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != "" && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != "" && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != "" && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != "" && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != "" && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != "" && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != "" && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != "" && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != "" && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != "" && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != "" && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != "" && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != "" && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != "" && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != "" && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != "" && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_CNT) fill_STAT_CNT_Header(fields []string, doc *map[string]interface{}) {
	// fill the met fields leaving out "" and NA values
	if fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != "" && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != "" && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != "" && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != "" && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != "" && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != "" && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != "" && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != "" && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != "" && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != "" && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != "" && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != "" && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != "" && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != "" && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != "" && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != "" && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_CTC) fill_STAT_CTC_Header(fields []string, doc *map[string]interface{}) {
	// fill the met fields leaving out "" and NA values
	if fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != "" && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != "" && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != "" && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != "" && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != "" && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != "" && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != "" && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != "" && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != "" && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != "" && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != "" && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != "" && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != "" && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != "" && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != "" && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != "" && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_MCTC) fill_STAT_MCTC_Header(fields []string, doc *map[string]interface{}) {
	// fill the met fields leaving out "" and NA values
	if fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != "" && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != "" && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != "" && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != "" && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != "" && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != "" && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != "" && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != "" && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != "" && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != "" && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != "" && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != "" && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != "" && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != "" && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != "" && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != "" && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_SL1L2) fill_STAT_SL1L2_Header(fields []string, doc *map[string]interface{}) {
	// fill the met fields leaving out "" and NA values
	if fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != "" && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != "" && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != "" && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != "" && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != "" && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != "" && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != "" && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != "" && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != "" && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != "" && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != "" && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != "" && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != "" && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != "" && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != "" && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != "" && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *MTD_3DSINGLE) fill_MTD_3DSINGLE_Header(fields []string, doc *map[string]interface{}) {
	// fill the met fields leaving out "" and NA values
	if fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != "" && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != "" && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID"] = fields[4]
	}
	if fields[5] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[5]
	}
	if fields[6] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID"] = fields[6]
	}
	if fields[7] != "" && fields[0] != "NA" {
		(*doc)["T_DELTA"] = fields[7]
	}
	if fields[8] != "" && fields[0] != "NA" {
		(*doc)["FCST_T_BEG"] = fields[8]
	}
	if fields[9] != "" && fields[0] != "NA" {
		(*doc)["FCST_T_END"] = fields[9]
	}
	if fields[10] != "" && fields[0] != "NA" {
		(*doc)["FCST_RAD"] = fields[10]
	}
	if fields[11] != "" && fields[0] != "NA" {
		(*doc)["FCST_THR"] = fields[11]
	}
	if fields[12] != "" && fields[0] != "NA" {
		(*doc)["OBS_T_BEG"] = fields[12]
	}
	if fields[13] != "" && fields[0] != "NA" {
		(*doc)["OBS_T_END"] = fields[13]
	}
	if fields[14] != "" && fields[0] != "NA" {
		(*doc)["OBS_RAD"] = fields[14]
	}
	if fields[15] != "" && fields[0] != "NA" {
		(*doc)["OBS_THR"] = fields[15]
	}
	if fields[16] != "" && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[16]
	}
	if fields[17] != "" && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[17]
	}
	if fields[19] != "" && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[19]
	}
	if fields[20] != "" && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[20]
	}
	if fields[21] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[21]
	}
}

func (s *MODE_OBJ) fill_MODE_OBJ_Header(fields []string, doc *map[string]interface{}) {
	// fill the met fields leaving out "" and NA values
	if fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != "" && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != "" && fields[0] != "NA" {
		(*doc)["N_VALID"] = fields[2]
	}
	if fields[3] != "" && fields[0] != "NA" {
		(*doc)["GRID_RES"] = fields[3]
	}
	if fields[4] != "" && fields[0] != "NA" {
		(*doc)["DESC"] = fields[4]
	}
	if fields[6] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID"] = fields[6]
	}
	if fields[7] != "" && fields[0] != "NA" {
		(*doc)["FCST_ACCUM"] = fields[7]
	}
	if fields[8] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[8]
	}
	if fields[9] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID"] = fields[9]
	}
	if fields[10] != "" && fields[0] != "NA" {
		(*doc)["OBS_ACCUM"] = fields[10]
	}
	if fields[11] != "" && fields[0] != "NA" {
		(*doc)["FCST_RAD"] = fields[11]
	}
	if fields[12] != "" && fields[0] != "NA" {
		(*doc)["FCST_THR"] = fields[12]
	}
	if fields[13] != "" && fields[0] != "NA" {
		(*doc)["OBS_RAD"] = fields[13]
	}
	if fields[14] != "" && fields[0] != "NA" {
		(*doc)["OBS_THR"] = fields[14]
	}
	if fields[15] != "" && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[15]
	}
	if fields[16] != "" && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[16]
	}
	if fields[18] != "" && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[18]
	}
	if fields[19] != "" && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[19]
	}
	if fields[20] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[20]
	}
	if fields[21] != "" && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[21]
	}
	if fields[22] != "" && fields[0] != "NA" {
		(*doc)["OBJECT_ID"] = fields[22]
	}
	if fields[23] != "" && fields[0] != "NA" {
		(*doc)["OBJECT_CAT"] = fields[23]
	}
	if fields[24] != "" && fields[0] != "NA" {
		(*doc)["CENTROID_X"] = fields[24]
	}
	if fields[25] != "" && fields[0] != "NA" {
		(*doc)["CENTROID_Y"] = fields[25]
	}
	if fields[26] != "" && fields[0] != "NA" {
		(*doc)["CENTROID_LAT"] = fields[26]
	}
	if fields[27] != "" && fields[0] != "NA" {
		(*doc)["CENTROID_LON"] = fields[27]
	}
	if fields[28] != "" && fields[0] != "NA" {
		(*doc)["AXIS_ANG"] = fields[28]
	}
	if fields[29] != "" && fields[0] != "NA" {
		(*doc)["LENGTH"] = fields[29]
	}
	if fields[30] != "" && fields[0] != "NA" {
		(*doc)["WIDTH"] = fields[30]
	}
	if fields[31] != "" && fields[0] != "NA" {
		(*doc)["AREA"] = fields[31]
	}
	if fields[32] != "" && fields[0] != "NA" {
		(*doc)["AREA_THRESH"] = fields[32]
	}
	if fields[33] != "" && fields[0] != "NA" {
		(*doc)["CURVATURE"] = fields[33]
	}
	if fields[34] != "" && fields[0] != "NA" {
		(*doc)["CURVATURE_X"] = fields[34]
	}
	if fields[35] != "" && fields[0] != "NA" {
		(*doc)["CURVATURE_Y"] = fields[35]
	}
	if fields[36] != "" && fields[0] != "NA" {
		(*doc)["COMPLEXITY"] = fields[36]
	}
	if fields[37] != "" && fields[0] != "NA" {
		(*doc)["INTENSITY_10"] = fields[37]
	}
	if fields[38] != "" && fields[0] != "NA" {
		(*doc)["INTENSITY_25"] = fields[38]
	}
	if fields[39] != "" && fields[0] != "NA" {
		(*doc)["INTENSITY_50"] = fields[39]
	}
	if fields[40] != "" && fields[0] != "NA" {
		(*doc)["INTENSITY_75"] = fields[40]
	}
	if fields[41] != "" && fields[0] != "NA" {
		(*doc)["INTENSITY_90"] = fields[41]
	}
	if fields[42] != "" && fields[0] != "NA" {
		(*doc)["INTENSITY_USER"] = fields[42]
	}
	if fields[43] != "" && fields[0] != "NA" {
		(*doc)["INTENSITY_SUM"] = fields[43]
	}
	if fields[44] != "" && fields[0] != "NA" {
		(*doc)["CENTROID_DIST"] = fields[44]
	}
	if fields[45] != "" && fields[0] != "NA" {
		(*doc)["BOUNDARY_DIST"] = fields[45]
	}
	if fields[46] != "" && fields[0] != "NA" {
		(*doc)["CONVEX_HULL_DIST"] = fields[46]
	}
	if fields[47] != "" && fields[0] != "NA" {
		(*doc)["ANGLE_DIFF"] = fields[47]
	}
	if fields[48] != "" && fields[0] != "NA" {
		(*doc)["ASPECT_DIFF"] = fields[48]
	}
	if fields[49] != "" && fields[0] != "NA" {
		(*doc)["AREA_RATIO"] = fields[49]
	}
	if fields[50] != "" && fields[0] != "NA" {
		(*doc)["INTERSECTION_AREA"] = fields[50]
	}
	if fields[51] != "" && fields[0] != "NA" {
		(*doc)["UNION_AREA"] = fields[51]
	}
	if fields[52] != "" && fields[0] != "NA" {
		(*doc)["SYMMETRIC_DIFF"] = fields[52]
	}
	if fields[53] != "" && fields[0] != "NA" {
		(*doc)["INTERSECTION_OVER_AREA"] = fields[53]
	}
	if fields[54] != "" && fields[0] != "NA" {
		(*doc)["CURVATURE_RATIO"] = fields[54]
	}
	if fields[55] != "" && fields[0] != "NA" {
		(*doc)["COMPLEXITY_RATIO"] = fields[55]
	}
	if fields[56] != "" && fields[0] != "NA" {
		(*doc)["PERCENTILE_INTENSITY_RATIO"] = fields[56]
	}
	if fields[57] != "" && fields[0] != "NA" {
		(*doc)["INTEREST"] = fields[57]
	}
}

func (s *TCST_TCMPR) fill_TCST_TCMPR_Header(fields []string, doc *map[string]interface{}) {
	// fill the met fields leaving out "" and NA values
	if fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != "" && fields[0] != "NA" {
		(*doc)["AMODEL"] = fields[1]
	}
	if fields[2] != "" && fields[0] != "NA" {
		(*doc)["BMODEL"] = fields[2]
	}
	if fields[3] != "" && fields[0] != "NA" {
		(*doc)["DESC"] = fields[3]
	}
	if fields[4] != "" && fields[0] != "NA" {
		(*doc)["STORM_ID"] = fields[4]
	}
	if fields[5] != "" && fields[0] != "NA" {
		(*doc)["BASIN"] = fields[5]
	}
	if fields[6] != "" && fields[0] != "NA" {
		(*doc)["CYCLONE"] = fields[6]
	}
	if fields[7] != "" && fields[0] != "NA" {
		(*doc)["STORM_NAME"] = fields[7]
	}
	if fields[8] != "" && fields[0] != "NA" {
		(*doc)["INIT"] = fields[8]
	}
	if fields[9] != "" && fields[0] != "NA" {
		(*doc)["LEAD"] = fields[9]
	}
	if fields[10] != "" && fields[0] != "NA" {
		(*doc)["VALID"] = fields[10]
	}
	if fields[11] != "" && fields[0] != "NA" {
		(*doc)["INIT_MASK"] = fields[11]
	}
	if fields[12] != "" && fields[0] != "NA" {
		(*doc)["VALID_MASK"] = fields[12]
	}
	if fields[13] != "" && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[13]
	}
}

func (s *STAT_ORANK) fill_STAT_ORANK_Header(fields []string, doc *map[string]interface{}) {
	// fill the met fields leaving out "" and NA values
	if fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != "" && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != "" && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != "" && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != "" && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != "" && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != "" && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != "" && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != "" && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != "" && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != "" && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != "" && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != "" && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != "" && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != "" && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != "" && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != "" && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_PJC) fill_STAT_PJC_Header(fields []string, doc *map[string]interface{}) {
	// fill the met fields leaving out "" and NA values
	if fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != "" && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != "" && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != "" && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != "" && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != "" && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != "" && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != "" && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != "" && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != "" && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != "" && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != "" && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != "" && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != "" && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != "" && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != "" && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != "" && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_RELP) fill_STAT_RELP_Header(fields []string, doc *map[string]interface{}) {
	// fill the met fields leaving out "" and NA values
	if fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != "" && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != "" && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != "" && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != "" && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != "" && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != "" && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != "" && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != "" && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != "" && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != "" && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != "" && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != "" && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != "" && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != "" && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != "" && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != "" && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_SSVAR) fill_STAT_SSVAR_Header(fields []string, doc *map[string]interface{}) {
	// fill the met fields leaving out "" and NA values
	if fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != "" && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != "" && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != "" && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != "" && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != "" && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != "" && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != "" && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != "" && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != "" && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != "" && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != "" && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != "" && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != "" && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != "" && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != "" && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != "" && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_VCNT) fill_STAT_VCNT_Header(fields []string, doc *map[string]interface{}) {
	// fill the met fields leaving out "" and NA values
	if fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != "" && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != "" && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != "" && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != "" && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != "" && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != "" && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != "" && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != "" && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != "" && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != "" && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != "" && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != "" && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != "" && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != "" && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != "" && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != "" && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_SEEPS_MPR) fill_STAT_SEEPS_MPR_Header(fields []string, doc *map[string]interface{}) {
	// fill the met fields leaving out "" and NA values
	if fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != "" && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != "" && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != "" && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != "" && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != "" && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != "" && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != "" && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != "" && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != "" && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != "" && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != "" && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != "" && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != "" && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != "" && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != "" && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != "" && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_VAL1L2) fill_STAT_VAL1L2_Header(fields []string, doc *map[string]interface{}) {
	// fill the met fields leaving out "" and NA values
	if fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != "" && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != "" && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != "" && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != "" && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != "" && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != "" && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != "" && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != "" && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != "" && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != "" && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != "" && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != "" && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != "" && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != "" && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != "" && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != "" && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *MTD_2DSINGLE) fill_MTD_2DSINGLE_Header(fields []string, doc *map[string]interface{}) {
	// fill the met fields leaving out "" and NA values
	if fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != "" && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != "" && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID"] = fields[4]
	}
	if fields[5] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[5]
	}
	if fields[6] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID"] = fields[6]
	}
	if fields[7] != "" && fields[0] != "NA" {
		(*doc)["T_DELTA"] = fields[7]
	}
	if fields[8] != "" && fields[0] != "NA" {
		(*doc)["FCST_T_BEG"] = fields[8]
	}
	if fields[9] != "" && fields[0] != "NA" {
		(*doc)["FCST_T_END"] = fields[9]
	}
	if fields[10] != "" && fields[0] != "NA" {
		(*doc)["FCST_RAD"] = fields[10]
	}
	if fields[11] != "" && fields[0] != "NA" {
		(*doc)["FCST_THR"] = fields[11]
	}
	if fields[12] != "" && fields[0] != "NA" {
		(*doc)["OBS_T_BEG"] = fields[12]
	}
	if fields[13] != "" && fields[0] != "NA" {
		(*doc)["OBS_T_END"] = fields[13]
	}
	if fields[14] != "" && fields[0] != "NA" {
		(*doc)["OBS_RAD"] = fields[14]
	}
	if fields[15] != "" && fields[0] != "NA" {
		(*doc)["OBS_THR"] = fields[15]
	}
	if fields[16] != "" && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[16]
	}
	if fields[17] != "" && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[17]
	}
	if fields[19] != "" && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[19]
	}
	if fields[20] != "" && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[20]
	}
	if fields[21] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[21]
	}
}

func (s *STAT_NBRCTS) fill_STAT_NBRCTS_Header(fields []string, doc *map[string]interface{}) {
	// fill the met fields leaving out "" and NA values
	if fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != "" && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != "" && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != "" && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != "" && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != "" && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != "" && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != "" && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != "" && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != "" && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != "" && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != "" && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != "" && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != "" && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != "" && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != "" && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != "" && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_GRAD) fill_STAT_GRAD_Header(fields []string, doc *map[string]interface{}) {
	// fill the met fields leaving out "" and NA values
	if fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != "" && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != "" && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != "" && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != "" && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != "" && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != "" && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != "" && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != "" && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != "" && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != "" && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != "" && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != "" && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != "" && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != "" && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != "" && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != "" && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_DMAP) fill_STAT_DMAP_Header(fields []string, doc *map[string]interface{}) {
	// fill the met fields leaving out "" and NA values
	if fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != "" && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != "" && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != "" && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != "" && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != "" && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != "" && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != "" && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != "" && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != "" && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != "" && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != "" && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != "" && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != "" && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != "" && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != "" && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != "" && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_CTS) fill_STAT_CTS_Header(fields []string, doc *map[string]interface{}) {
	// fill the met fields leaving out "" and NA values
	if fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != "" && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != "" && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != "" && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != "" && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != "" && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != "" && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != "" && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != "" && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != "" && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != "" && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != "" && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != "" && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != "" && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != "" && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != "" && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != "" && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_FHO) fill_STAT_FHO_Header(fields []string, doc *map[string]interface{}) {
	// fill the met fields leaving out "" and NA values
	if fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != "" && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != "" && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != "" && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != "" && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != "" && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != "" && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != "" && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != "" && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != "" && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != "" && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != "" && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != "" && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != "" && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != "" && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != "" && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != "" && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_ISC) fill_STAT_ISC_Header(fields []string, doc *map[string]interface{}) {
	// fill the met fields leaving out "" and NA values
	if fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != "" && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != "" && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != "" && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != "" && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != "" && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != "" && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != "" && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != "" && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != "" && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != "" && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != "" && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != "" && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != "" && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != "" && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != "" && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != "" && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_MPR) fill_STAT_MPR_Header(fields []string, doc *map[string]interface{}) {
	// fill the met fields leaving out "" and NA values
	if fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != "" && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != "" && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != "" && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != "" && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != "" && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != "" && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != "" && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != "" && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != "" && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != "" && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != "" && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != "" && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != "" && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != "" && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != "" && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != "" && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_NBRCTC) fill_STAT_NBRCTC_Header(fields []string, doc *map[string]interface{}) {
	// fill the met fields leaving out "" and NA values
	if fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != "" && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != "" && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != "" && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != "" && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != "" && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != "" && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != "" && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != "" && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != "" && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != "" && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != "" && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != "" && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != "" && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != "" && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != "" && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != "" && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_PCT) fill_STAT_PCT_Header(fields []string, doc *map[string]interface{}) {
	// fill the met fields leaving out "" and NA values
	if fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != "" && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != "" && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != "" && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != "" && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != "" && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != "" && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != "" && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != "" && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != "" && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != "" && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != "" && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != "" && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != "" && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != "" && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != "" && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != "" && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_ECLV) fill_STAT_ECLV_Header(fields []string, doc *map[string]interface{}) {
	// fill the met fields leaving out "" and NA values
	if fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != "" && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != "" && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != "" && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != "" && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != "" && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != "" && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != "" && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != "" && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != "" && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != "" && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != "" && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != "" && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != "" && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != "" && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != "" && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != "" && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *TCST_PROBRIRW) fill_TCST_PROBRIRW_Header(fields []string, doc *map[string]interface{}) {
	// fill the met fields leaving out "" and NA values
	if fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != "" && fields[0] != "NA" {
		(*doc)["AMODEL"] = fields[1]
	}
	if fields[2] != "" && fields[0] != "NA" {
		(*doc)["BMODEL"] = fields[2]
	}
	if fields[3] != "" && fields[0] != "NA" {
		(*doc)["DESC"] = fields[3]
	}
	if fields[4] != "" && fields[0] != "NA" {
		(*doc)["STORM_ID"] = fields[4]
	}
	if fields[5] != "" && fields[0] != "NA" {
		(*doc)["BASIN"] = fields[5]
	}
	if fields[6] != "" && fields[0] != "NA" {
		(*doc)["CYCLONE"] = fields[6]
	}
	if fields[7] != "" && fields[0] != "NA" {
		(*doc)["STORM_NAME"] = fields[7]
	}
	if fields[8] != "" && fields[0] != "NA" {
		(*doc)["INIT"] = fields[8]
	}
	if fields[9] != "" && fields[0] != "NA" {
		(*doc)["LEAD"] = fields[9]
	}
	if fields[10] != "" && fields[0] != "NA" {
		(*doc)["VALID"] = fields[10]
	}
	if fields[11] != "" && fields[0] != "NA" {
		(*doc)["INIT_MASK"] = fields[11]
	}
	if fields[12] != "" && fields[0] != "NA" {
		(*doc)["VALID_MASK"] = fields[12]
	}
	if fields[13] != "" && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[13]
	}
}

func (s *STAT_PRC) fill_STAT_PRC_Header(fields []string, doc *map[string]interface{}) {
	// fill the met fields leaving out "" and NA values
	if fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != "" && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != "" && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != "" && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != "" && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != "" && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != "" && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != "" && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != "" && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != "" && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != "" && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != "" && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != "" && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != "" && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != "" && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != "" && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != "" && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

func (s *STAT_GENMPR) fill_STAT_GENMPR_Header(fields []string, doc *map[string]interface{}) {
	// fill the met fields leaving out "" and NA values
	if fields[0] != "" && fields[0] != "NA" {
		(*doc)["VERSION"] = fields[0]
	}
	if fields[1] != "" && fields[0] != "NA" {
		(*doc)["MODEL"] = fields[1]
	}
	if fields[2] != "" && fields[0] != "NA" {
		(*doc)["DESC"] = fields[2]
	}
	if fields[4] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_BEG"] = fields[4]
	}
	if fields[5] != "" && fields[0] != "NA" {
		(*doc)["FCST_VALID_END"] = fields[5]
	}
	if fields[6] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEAD"] = fields[6]
	}
	if fields[7] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_BEG"] = fields[7]
	}
	if fields[8] != "" && fields[0] != "NA" {
		(*doc)["OBS_VALID_END"] = fields[8]
	}
	if fields[9] != "" && fields[0] != "NA" {
		(*doc)["FCST_VAR"] = fields[9]
	}
	if fields[10] != "" && fields[0] != "NA" {
		(*doc)["FCST_UNITS"] = fields[10]
	}
	if fields[11] != "" && fields[0] != "NA" {
		(*doc)["FCST_LEV"] = fields[11]
	}
	if fields[12] != "" && fields[0] != "NA" {
		(*doc)["OBS_VAR"] = fields[12]
	}
	if fields[13] != "" && fields[0] != "NA" {
		(*doc)["OBS_UNITS"] = fields[13]
	}
	if fields[14] != "" && fields[0] != "NA" {
		(*doc)["OBS_LEV"] = fields[14]
	}
	if fields[15] != "" && fields[0] != "NA" {
		(*doc)["OBTYPE"] = fields[15]
	}
	if fields[16] != "" && fields[0] != "NA" {
		(*doc)["VX_MASK"] = fields[16]
	}
	if fields[17] != "" && fields[0] != "NA" {
		(*doc)["INTERP_MTHD"] = fields[17]
	}
	if fields[18] != "" && fields[0] != "NA" {
		(*doc)["INTERP_PNTS"] = fields[18]
	}
	if fields[19] != "" && fields[0] != "NA" {
		(*doc)["FCST_THRESH"] = fields[19]
	}
	if fields[20] != "" && fields[0] != "NA" {
		(*doc)["OBS_THRESH"] = fields[20]
	}
	if fields[21] != "" && fields[0] != "NA" {
		(*doc)["COV_THRESH"] = fields[21]
	}
	if fields[22] != "" && fields[0] != "NA" {
		(*doc)["ALPHA"] = fields[22]
	}
	if fields[23] != "" && fields[0] != "NA" {
		(*doc)["LINE_TYPE"] = fields[23]
	}
}

//line data struct definitions
type STAT_CTS struct {
	TOTAL      int     `json:"TOTAL"`
	BASER      float64 `json:"BASER"`
	BASER_NCL  float64 `json:"BASER_NCL"`
	BASER_NCU  float64 `json:"BASER_NCU"`
	BASER_BCL  float64 `json:"BASER_BCL"`
	BASER_BCU  float64 `json:"BASER_BCU"`
	FMEAN      float64 `json:"FMEAN"`
	FMEAN_NCL  float64 `json:"FMEAN_NCL"`
	FMEAN_NCU  float64 `json:"FMEAN_NCU"`
	FMEAN_BCL  float64 `json:"FMEAN_BCL"`
	FMEAN_BCU  float64 `json:"FMEAN_BCU"`
	ACC        float64 `json:"ACC"`
	ACC_NCL    float64 `json:"ACC_NCL"`
	ACC_NCU    float64 `json:"ACC_NCU"`
	ACC_BCL    float64 `json:"ACC_BCL"`
	ACC_BCU    float64 `json:"ACC_BCU"`
	FBIAS      float64 `json:"FBIAS"`
	FBIAS_BCL  float64 `json:"FBIAS_BCL"`
	FBIAS_BCU  float64 `json:"FBIAS_BCU"`
	PODY       float64 `json:"PODY"`
	PODY_NCL   float64 `json:"PODY_NCL"`
	PODY_NCU   float64 `json:"PODY_NCU"`
	PODY_BCL   float64 `json:"PODY_BCL"`
	PODY_BCU   float64 `json:"PODY_BCU"`
	PODN       float64 `json:"PODN"`
	PODN_NCL   float64 `json:"PODN_NCL"`
	PODN_NCU   float64 `json:"PODN_NCU"`
	PODN_BCL   float64 `json:"PODN_BCL"`
	PODN_BCU   float64 `json:"PODN_BCU"`
	POFD       float64 `json:"POFD"`
	POFD_NCL   float64 `json:"POFD_NCL"`
	POFD_NCU   float64 `json:"POFD_NCU"`
	POFD_BCL   float64 `json:"POFD_BCL"`
	POFD_BCU   float64 `json:"POFD_BCU"`
	FAR        float64 `json:"FAR"`
	FAR_NCL    float64 `json:"FAR_NCL"`
	FAR_NCU    float64 `json:"FAR_NCU"`
	FAR_BCL    float64 `json:"FAR_BCL"`
	FAR_BCU    float64 `json:"FAR_BCU"`
	CSI        float64 `json:"CSI"`
	CSI_NCL    float64 `json:"CSI_NCL"`
	CSI_NCU    float64 `json:"CSI_NCU"`
	CSI_BCL    float64 `json:"CSI_BCL"`
	CSI_BCU    float64 `json:"CSI_BCU"`
	GSS        float64 `json:"GSS"`
	GSS_BCL    float64 `json:"GSS_BCL"`
	GSS_BCU    float64 `json:"GSS_BCU"`
	HK         float64 `json:"HK"`
	HK_NCL     float64 `json:"HK_NCL"`
	HK_NCU     float64 `json:"HK_NCU"`
	HK_BCL     float64 `json:"HK_BCL"`
	HK_BCU     float64 `json:"HK_BCU"`
	HSS        float64 `json:"HSS"`
	HSS_BCL    float64 `json:"HSS_BCL"`
	HSS_BCU    float64 `json:"HSS_BCU"`
	ODDS       float64 `json:"ODDS"`
	ODDS_NCL   float64 `json:"ODDS_NCL"`
	ODDS_NCU   float64 `json:"ODDS_NCU"`
	ODDS_BCL   float64 `json:"ODDS_BCL"`
	ODDS_BCU   float64 `json:"ODDS_BCU"`
	LODDS      float64 `json:"LODDS"`
	LODDS_NCL  float64 `json:"LODDS_NCL"`
	LODDS_NCU  float64 `json:"LODDS_NCU"`
	LODDS_BCL  float64 `json:"LODDS_BCL"`
	LODDS_BCU  float64 `json:"LODDS_BCU"`
	ORSS       float64 `json:"ORSS"`
	ORSS_NCL   float64 `json:"ORSS_NCL"`
	ORSS_NCU   float64 `json:"ORSS_NCU"`
	ORSS_BCL   float64 `json:"ORSS_BCL"`
	ORSS_BCU   float64 `json:"ORSS_BCU"`
	EDS        float64 `json:"EDS"`
	EDS_NCL    float64 `json:"EDS_NCL"`
	EDS_NCU    float64 `json:"EDS_NCU"`
	EDS_BCL    float64 `json:"EDS_BCL"`
	EDS_BCU    float64 `json:"EDS_BCU"`
	SEDS       float64 `json:"SEDS"`
	SEDS_NCL   float64 `json:"SEDS_NCL"`
	SEDS_NCU   float64 `json:"SEDS_NCU"`
	SEDS_BCL   float64 `json:"SEDS_BCL"`
	SEDS_BCU   float64 `json:"SEDS_BCU"`
	EDI        float64 `json:"EDI"`
	EDI_NCL    float64 `json:"EDI_NCL"`
	EDI_NCU    float64 `json:"EDI_NCU"`
	EDI_BCL    float64 `json:"EDI_BCL"`
	EDI_BCU    float64 `json:"EDI_BCU"`
	SEDI       float64 `json:"SEDI"`
	SEDI_NCL   float64 `json:"SEDI_NCL"`
	SEDI_NCU   float64 `json:"SEDI_NCU"`
	SEDI_BCL   float64 `json:"SEDI_BCL"`
	SEDI_BCU   float64 `json:"SEDI_BCU"`
	BAGSS      float64 `json:"BAGSS"`
	BAGSS_BCL  float64 `json:"BAGSS_BCL"`
	BAGSS_BCU  float64 `json:"BAGSS_BCU"`
	HSS_EC     float64 `json:"HSS_EC"`
	HSS_EC_BCL float64 `json:"HSS_EC_BCL"`
	HSS_EC_BCU float64 `json:"HSS_EC_BCU"`
	EC_VALUE   float64 `json:"EC_VALUE"`
}

type STAT_MPR struct {
	TOTAL            int     `json:"TOTAL"`
	INDEX            int     `json:"INDEX"`
	OBS_SID          string  `json:"OBS_SID"`
	OBS_LAT          float64 `json:"OBS_LAT"`
	OBS_LON          float64 `json:"OBS_LON"`
	OBS_LVL          float64 `json:"OBS_LVL"`
	OBS_ELV          float64 `json:"OBS_ELV"`
	FCST             float64 `json:"FCST"`
	OBS              float64 `json:"OBS"`
	OBS_QC           string  `json:"OBS_QC"`
	OBS_CLIMO_MEAN   float64 `json:"OBS_CLIMO_MEAN"`
	OBS_CLIMO_STDEV  float64 `json:"OBS_CLIMO_STDEV"`
	OBS_CLIMO_CDF    float64 `json:"OBS_CLIMO_CDF"`
	FCST_CLIMO_MEAN  float64 `json:"FCST_CLIMO_MEAN"`
	FCST_CLIMO_STDEV float64 `json:"FCST_CLIMO_STDEV"`
}

type STAT_VAL1L2 struct {
	TOTAL        int     `json:"TOTAL"`
	UFABAR       float64 `json:"UFABAR"`
	VFABAR       float64 `json:"VFABAR"`
	UOABAR       float64 `json:"UOABAR"`
	VOABAR       float64 `json:"VOABAR"`
	UVFOABAR     float64 `json:"UVFOABAR"`
	UVFFABAR     float64 `json:"UVFFABAR"`
	UVOOABAR     float64 `json:"UVOOABAR"`
	FA_SPEED_BAR float64 `json:"FA_SPEED_BAR"`
	OA_SPEED_BAR float64 `json:"OA_SPEED_BAR"`
	TOTAL_DIR    float64 `json:"TOTAL_DIR"`
	DIRA_ME      float64 `json:"DIRA_ME"`
	DIRA_MAE     float64 `json:"DIRA_MAE"`
	DIRA_MSE     float64 `json:"DIRA_MSE"`
}

type STAT_VL1L2 struct {
	TOTAL       int     `json:"TOTAL"`
	UFBAR       float64 `json:"UFBAR"`
	VFBAR       float64 `json:"VFBAR"`
	UOBAR       float64 `json:"UOBAR"`
	VOBAR       float64 `json:"VOBAR"`
	UVFOBAR     float64 `json:"UVFOBAR"`
	UVFFBAR     float64 `json:"UVFFBAR"`
	UVOOBAR     float64 `json:"UVOOBAR"`
	F_SPEED_BAR float64 `json:"F_SPEED_BAR"`
	O_SPEED_BAR float64 `json:"O_SPEED_BAR"`
	TOTAL_DIR   float64 `json:"TOTAL_DIR"`
	DIR_ME      float64 `json:"DIR_ME"`
	DIR_MAE     float64 `json:"DIR_MAE"`
	DIR_MSE     float64 `json:"DIR_MSE"`
}

type MTD_2DSINGLE struct {
	OBJECT_ID      string  `json:"OBJECT_ID"`
	OBJECT_CAT     string  `json:"OBJECT_CAT"`
	TIME_INDEX     int     `json:"TIME_INDEX"`
	AREA           int     `json:"AREA"`
	CENTROID_X     float64 `json:"CENTROID_X"`
	CENTROID_Y     float64 `json:"CENTROID_Y"`
	CENTROID_LAT   float64 `json:"CENTROID_LAT"`
	CENTROID_LON   float64 `json:"CENTROID_LON"`
	AXIS_ANG       float64 `json:"AXIS_ANG"`
	INTENSITY_10   float64 `json:"INTENSITY_10"`
	INTENSITY_25   float64 `json:"INTENSITY_25"`
	INTENSITY_50   float64 `json:"INTENSITY_50"`
	INTENSITY_75   float64 `json:"INTENSITY_75"`
	INTENSITY_90   float64 `json:"INTENSITY_90"`
	INTENSITY_USER float64 `json:"INTENSITY_USER"`
}

type TCST_TCMPR struct {
	TOTAL          int     `json:"TOTAL"`
	INDEX          int     `json:"INDEX"`
	LEVEL          string  `json:"LEVEL"`
	WATCH_WARN     string  `json:"WATCH_WARN"`
	INITIALS       string  `json:"INITIALS"`
	ALAT           float64 `json:"ALAT"`
	ALON           float64 `json:"ALON"`
	BLAT           float64 `json:"BLAT"`
	BLON           float64 `json:"BLON"`
	TK_ERR         float64 `json:"TK_ERR"`
	X_ERR          float64 `json:"X_ERR"`
	Y_ERR          float64 `json:"Y_ERR"`
	ALTK_ERR       float64 `json:"ALTK_ERR"`
	CRTK_ERR       float64 `json:"CRTK_ERR"`
	ADLAND         float64 `json:"ADLAND"`
	BDLAND         float64 `json:"BDLAND"`
	AMSLP          float64 `json:"AMSLP"`
	BMSLP          float64 `json:"BMSLP"`
	AMAX_WIND      float64 `json:"AMAX_WIND"`
	BMAX_WIND      float64 `json:"BMAX_WIND"`
	AAL_WIND_34    float64 `json:"AAL_WIND_34"`
	BAL_WIND_34    float64 `json:"BAL_WIND_34"`
	ANE_WIND_34    float64 `json:"ANE_WIND_34"`
	BNE_WIND_34    float64 `json:"BNE_WIND_34"`
	ASE_WIND_34    float64 `json:"ASE_WIND_34"`
	BSE_WIND_34    float64 `json:"BSE_WIND_34"`
	ASW_WIND_34    float64 `json:"ASW_WIND_34"`
	BSW_WIND_34    float64 `json:"BSW_WIND_34"`
	ANW_WIND_34    float64 `json:"ANW_WIND_34"`
	BNW_WIND_34    float64 `json:"BNW_WIND_34"`
	AAL_WIND_50    float64 `json:"AAL_WIND_50"`
	BAL_WIND_50    float64 `json:"BAL_WIND_50"`
	ANE_WIND_50    float64 `json:"ANE_WIND_50"`
	BNE_WIND_50    float64 `json:"BNE_WIND_50"`
	ASE_WIND_50    float64 `json:"ASE_WIND_50"`
	BSE_WIND_50    float64 `json:"BSE_WIND_50"`
	ASW_WIND_50    float64 `json:"ASW_WIND_50"`
	BSW_WIND_50    float64 `json:"BSW_WIND_50"`
	ANW_WIND_50    float64 `json:"ANW_WIND_50"`
	BNW_WIND_50    float64 `json:"BNW_WIND_50"`
	AAL_WIND_64    float64 `json:"AAL_WIND_64"`
	BAL_WIND_64    float64 `json:"BAL_WIND_64"`
	ANE_WIND_64    float64 `json:"ANE_WIND_64"`
	BNE_WIND_64    float64 `json:"BNE_WIND_64"`
	ASE_WIND_64    float64 `json:"ASE_WIND_64"`
	BSE_WIND_64    float64 `json:"BSE_WIND_64"`
	ASW_WIND_64    float64 `json:"ASW_WIND_64"`
	BSW_WIND_64    float64 `json:"BSW_WIND_64"`
	ANW_WIND_64    float64 `json:"ANW_WIND_64"`
	BNW_WIND_64    float64 `json:"BNW_WIND_64"`
	ARADP          string  `json:"ARADP"`
	BRADP          float64 `json:"BRADP"`
	ARRP           int     `json:"ARRP"`
	BRRP           float64 `json:"BRRP"`
	AMRD           int     `json:"AMRD"`
	BMRD           float64 `json:"BMRD"`
	AGUSTS         int     `json:"AGUSTS"`
	BGUSTS         float64 `json:"BGUSTS"`
	AEYE           int     `json:"AEYE"`
	BEYE           float64 `json:"BEYE"`
	ADIR           int     `json:"ADIR"`
	BDIR           float64 `json:"BDIR"`
	ASPEED         int     `json:"ASPEED"`
	BSPEED         float64 `json:"BSPEED"`
	ADEPTH         int     `json:"ADEPTH"`
	BDEPTH         float64 `json:"BDEPTH"`
	NUM_MEMBERS    float64 `json:"NUM_MEMBERS"`
	TRACK_SPREAD   float64 `json:"TRACK_SPREAD"`
	TRACK_STDEV    float64 `json:"TRACK_STDEV"`
	MSLP_STDEV     float64 `json:"MSLP_STDEV"`
	MAX_WIND_STDEV float64 `json:"MAX_WIND_STDEV"`
}

type STAT_CTC struct {
	TOTAL    int     `json:"TOTAL"`
	FY_OY    float64 `json:"FY_OY"`
	FY_ON    float64 `json:"FY_ON"`
	FN_OY    float64 `json:"FN_OY"`
	FN_ON    float64 `json:"FN_ON"`
	EC_VALUE float64 `json:"EC_VALUE"`
}

type STAT_ECLV struct {
	TOTAL       int                    `json:"TOTAL"`
	BASER       float64                `json:"BASER"`
	VALUE_BASER int                    `json:"VALUE_BASER"`
	PTS         map[string]interface{} `json:"PTS"`
	CL_I        float64                `json:"CL_I"`
	VALUE_I     int                    `json:"VALUE_I"`
}

type STAT_PHIST struct {
	TOTAL    int            `json:"TOTAL"`
	BIN_SIZE int            `json:"BIN_SIZE"`
	BIN      map[string]int `json:"BIN"`
	BIN_I    int            `json:"BIN_I"`
}

type TCST_PROBRIRW struct {
	ALAT        float64                `json:"ALAT"`
	ALON        float64                `json:"ALON"`
	BLAT        float64                `json:"BLAT"`
	BLON        float64                `json:"BLON"`
	INITIALS    string                 `json:"INITIALS"`
	TK_ERR      float64                `json:"TK_ERR"`
	X_ERR       float64                `json:"X_ERR"`
	Y_ERR       float64                `json:"Y_ERR"`
	ADLAND      float64                `json:"ADLAND"`
	BDLAND      float64                `json:"BDLAND"`
	RIRW_BEG    int                    `json:"RIRW_BEG"`
	RIRW_END    int                    `json:"RIRW_END"`
	RIRW_WINDOW int                    `json:"RIRW_WINDOW"`
	AWIND_END   float64                `json:"AWIND_END"`
	BWIND_BEG   float64                `json:"BWIND_BEG"`
	BWIND_END   float64                `json:"BWIND_END"`
	BDELTA      float64                `json:"BDELTA"`
	BDELTA_MAX  float64                `json:"BDELTA_MAX"`
	BLEVEL_BEG  string                 `json:"BLEVEL_BEG"`
	BLEVEL_END  string                 `json:"BLEVEL_END"`
	THRESH      map[string]interface{} `json:"THRESH"`
	THRESH_I    int                    `json:"THRESH_I"`
	PROB_I      float64                `json:"PROB_I"`
}

type STAT_CNT struct {
	TOTAL                int     `json:"TOTAL"`
	FBAR                 float64 `json:"FBAR"`
	FBAR_NCL             float64 `json:"FBAR_NCL"`
	FBAR_NCU             float64 `json:"FBAR_NCU"`
	FBAR_BCL             float64 `json:"FBAR_BCL"`
	FBAR_BCU             float64 `json:"FBAR_BCU"`
	FSTDEV               float64 `json:"FSTDEV"`
	FSTDEV_NCL           float64 `json:"FSTDEV_NCL"`
	FSTDEV_NCU           float64 `json:"FSTDEV_NCU"`
	FSTDEV_BCL           float64 `json:"FSTDEV_BCL"`
	FSTDEV_BCU           float64 `json:"FSTDEV_BCU"`
	OBAR                 float64 `json:"OBAR"`
	OBAR_NCL             float64 `json:"OBAR_NCL"`
	OBAR_NCU             float64 `json:"OBAR_NCU"`
	OBAR_BCL             float64 `json:"OBAR_BCL"`
	OBAR_BCU             float64 `json:"OBAR_BCU"`
	OSTDEV               float64 `json:"OSTDEV"`
	OSTDEV_NCL           float64 `json:"OSTDEV_NCL"`
	OSTDEV_NCU           float64 `json:"OSTDEV_NCU"`
	OSTDEV_BCL           float64 `json:"OSTDEV_BCL"`
	OSTDEV_BCU           float64 `json:"OSTDEV_BCU"`
	PR_CORR              float64 `json:"PR_CORR"`
	PR_CORR_NCL          float64 `json:"PR_CORR_NCL"`
	PR_CORR_NCU          float64 `json:"PR_CORR_NCU"`
	PR_CORR_BCL          float64 `json:"PR_CORR_BCL"`
	PR_CORR_BCU          float64 `json:"PR_CORR_BCU"`
	SP_CORR              float64 `json:"SP_CORR"`
	KT_CORR              float64 `json:"KT_CORR"`
	RANKS                int     `json:"RANKS"`
	FRANK_TIES           int     `json:"FRANK_TIES"`
	ORANK_TIES           int     `json:"ORANK_TIES"`
	ME                   float64 `json:"ME"`
	ME_NCL               float64 `json:"ME_NCL"`
	ME_NCU               float64 `json:"ME_NCU"`
	ME_BCL               float64 `json:"ME_BCL"`
	ME_BCU               float64 `json:"ME_BCU"`
	ESTDEV               float64 `json:"ESTDEV"`
	ESTDEV_NCL           float64 `json:"ESTDEV_NCL"`
	ESTDEV_NCU           float64 `json:"ESTDEV_NCU"`
	ESTDEV_BCL           float64 `json:"ESTDEV_BCL"`
	ESTDEV_BCU           float64 `json:"ESTDEV_BCU"`
	MBIAS                float64 `json:"MBIAS"`
	MBIAS_BCL            float64 `json:"MBIAS_BCL"`
	MBIAS_BCU            float64 `json:"MBIAS_BCU"`
	MAE                  float64 `json:"MAE"`
	MAE_BCL              float64 `json:"MAE_BCL"`
	MAE_BCU              float64 `json:"MAE_BCU"`
	MSE                  float64 `json:"MSE"`
	MSE_BCL              float64 `json:"MSE_BCL"`
	MSE_BCU              float64 `json:"MSE_BCU"`
	BCMSE                float64 `json:"BCMSE"`
	BCMSE_BCL            float64 `json:"BCMSE_BCL"`
	BCMSE_BCU            float64 `json:"BCMSE_BCU"`
	RMSE                 float64 `json:"RMSE"`
	RMSE_BCL             float64 `json:"RMSE_BCL"`
	RMSE_BCU             float64 `json:"RMSE_BCU"`
	E10                  float64 `json:"E10"`
	E10_BCL              float64 `json:"E10_BCL"`
	E10_BCU              float64 `json:"E10_BCU"`
	E25                  float64 `json:"E25"`
	E25_BCL              float64 `json:"E25_BCL"`
	E25_BCU              float64 `json:"E25_BCU"`
	E50                  float64 `json:"E50"`
	E50_BCL              float64 `json:"E50_BCL"`
	E50_BCU              float64 `json:"E50_BCU"`
	E75                  float64 `json:"E75"`
	E75_BCL              float64 `json:"E75_BCL"`
	E75_BCU              float64 `json:"E75_BCU"`
	E90                  float64 `json:"E90"`
	E90_BCL              float64 `json:"E90_BCL"`
	E90_BCU              float64 `json:"E90_BCU"`
	EIQR                 float64 `json:"EIQR"`
	EIQR_BCL             float64 `json:"EIQR_BCL"`
	EIQR_BCU             float64 `json:"EIQR_BCU"`
	MAD                  float64 `json:"MAD"`
	MAD_BCL              float64 `json:"MAD_BCL"`
	MAD_BCU              float64 `json:"MAD_BCU"`
	ANOM_CORR            float64 `json:"ANOM_CORR"`
	ANOM_CORR_NCL        float64 `json:"ANOM_CORR_NCL"`
	ANOM_CORR_NCU        float64 `json:"ANOM_CORR_NCU"`
	ANOM_CORR_BCL        float64 `json:"ANOM_CORR_BCL"`
	ANOM_CORR_BCU        float64 `json:"ANOM_CORR_BCU"`
	ME2                  float64 `json:"ME2"`
	ME2_BCL              float64 `json:"ME2_BCL"`
	ME2_BCU              float64 `json:"ME2_BCU"`
	MSESS                float64 `json:"MSESS"`
	MSESS_BCL            float64 `json:"MSESS_BCL"`
	MSESS_BCU            float64 `json:"MSESS_BCU"`
	RMSFA                float64 `json:"RMSFA"`
	RMSFA_BCL            float64 `json:"RMSFA_BCL"`
	RMSFA_BCU            float64 `json:"RMSFA_BCU"`
	RMSOA                float64 `json:"RMSOA"`
	RMSOA_BCL            float64 `json:"RMSOA_BCL"`
	RMSOA_BCU            float64 `json:"RMSOA_BCU"`
	ANOM_CORR_UNCNTR     float64 `json:"ANOM_CORR_UNCNTR"`
	ANOM_CORR_UNCNTR_BCL float64 `json:"ANOM_CORR_UNCNTR_BCL"`
	ANOM_CORR_UNCNTR_BCU float64 `json:"ANOM_CORR_UNCNTR_BCU"`
	SI                   float64 `json:"SI"`
	SI_BCL               float64 `json:"SI_BCL"`
	SI_BCU               float64 `json:"SI_BCU"`
}

type STAT_MCTC struct {
	TOTAL    int            `json:"TOTAL"`
	CAT      map[string]int `json:"CAT"`
	FI_OI    string         `json:"FI_OI"`
	EC_VALUE float64        `json:"EC_VALUE"`
}

type STAT_SEEPS_MPR struct {
	OBS_SID  string  `json:"OBS_SID"`
	OBS_LAT  float64 `json:"OBS_LAT"`
	OBS_LON  float64 `json:"OBS_LON"`
	FCST     float64 `json:"FCST"`
	OBS      float64 `json:"OBS"`
	OBS_QC   string  `json:"OBS_QC"`
	FCST_CAT int     `json:"FCST_CAT"`
	OBS_CAT  int     `json:"OBS_CAT"`
	P1       float64 `json:"P1"`
	P2       float64 `json:"P2"`
	T1       float64 `json:"T1"`
	T2       float64 `json:"T2"`
	SEEPS    float64 `json:"SEEPS"`
}

type STAT_GRAD struct {
	TOTAL      int     `json:"TOTAL"`
	FGBAR      float64 `json:"FGBAR"`
	OGBAR      float64 `json:"OGBAR"`
	MGBAR      float64 `json:"MGBAR"`
	EGBAR      float64 `json:"EGBAR"`
	S1         float64 `json:"S1"`
	S1_OG      float64 `json:"S1_OG"`
	FGOG_RATIO float64 `json:"FGOG_RATIO"`
	DX         float64 `json:"DX"`
	DY         float64 `json:"DY"`
}

type STAT_PCT struct {
	TOTAL    int                    `json:"TOTAL"`
	THRESH   map[string]interface{} `json:"THRESH"`
	THRESH_I int                    `json:"THRESH_I"`
	OY_I     float64                `json:"OY_I"`
	ON_I     float64                `json:"ON_I"`
}

type STAT_RHIST struct {
	TOTAL  int            `json:"TOTAL"`
	RANK   map[string]int `json:"RANK"`
	RANK_I int            `json:"RANK_I"`
}

type STAT_SL1L2 struct {
	TOTAL int     `json:"TOTAL"`
	FBAR  float64 `json:"FBAR"`
	OBAR  float64 `json:"OBAR"`
	FOBAR float64 `json:"FOBAR"`
	FFBAR float64 `json:"FFBAR"`
	OOBAR float64 `json:"OOBAR"`
	MAE   float64 `json:"MAE"`
}

type STAT_FHO struct {
	TOTAL  int     `json:"TOTAL"`
	F_RATE float64 `json:"F_RATE"`
	H_RATE float64 `json:"H_RATE"`
	O_RATE float64 `json:"O_RATE"`
}

type STAT_NBRCTS struct {
	TOTAL     int     `json:"TOTAL"`
	BASER     float64 `json:"BASER"`
	BASER_NCL float64 `json:"BASER_NCL"`
	BASER_NCU float64 `json:"BASER_NCU"`
	BASER_BCL float64 `json:"BASER_BCL"`
	BASER_BCU float64 `json:"BASER_BCU"`
	FMEAN     float64 `json:"FMEAN"`
	FMEAN_NCL float64 `json:"FMEAN_NCL"`
	FMEAN_NCU float64 `json:"FMEAN_NCU"`
	FMEAN_BCL float64 `json:"FMEAN_BCL"`
	FMEAN_BCU float64 `json:"FMEAN_BCU"`
	ACC       float64 `json:"ACC"`
	ACC_NCL   float64 `json:"ACC_NCL"`
	ACC_NCU   float64 `json:"ACC_NCU"`
	ACC_BCL   float64 `json:"ACC_BCL"`
	ACC_BCU   float64 `json:"ACC_BCU"`
	FBIAS     float64 `json:"FBIAS"`
	FBIAS_BCL float64 `json:"FBIAS_BCL"`
	FBIAS_BCU float64 `json:"FBIAS_BCU"`
	PODY      float64 `json:"PODY"`
	PODY_NCL  float64 `json:"PODY_NCL"`
	PODY_NCU  float64 `json:"PODY_NCU"`
	PODY_BCL  float64 `json:"PODY_BCL"`
	PODY_BCU  float64 `json:"PODY_BCU"`
	PODN      float64 `json:"PODN"`
	PODN_NCL  float64 `json:"PODN_NCL"`
	PODN_NCU  float64 `json:"PODN_NCU"`
	PODN_BCL  float64 `json:"PODN_BCL"`
	PODN_BCU  float64 `json:"PODN_BCU"`
	POFD      float64 `json:"POFD"`
	POFD_NCL  float64 `json:"POFD_NCL"`
	POFD_NCU  float64 `json:"POFD_NCU"`
	POFD_BCL  float64 `json:"POFD_BCL"`
	POFD_BCU  float64 `json:"POFD_BCU"`
	FAR       float64 `json:"FAR"`
	FAR_NCL   float64 `json:"FAR_NCL"`
	FAR_NCU   float64 `json:"FAR_NCU"`
	FAR_BCL   float64 `json:"FAR_BCL"`
	FAR_BCU   float64 `json:"FAR_BCU"`
	CSI       float64 `json:"CSI"`
	CSI_NCL   float64 `json:"CSI_NCL"`
	CSI_NCU   float64 `json:"CSI_NCU"`
	CSI_BCL   float64 `json:"CSI_BCL"`
	CSI_BCU   float64 `json:"CSI_BCU"`
	GSS       float64 `json:"GSS"`
	GSS_BCL   float64 `json:"GSS_BCL"`
	GSS_BCU   float64 `json:"GSS_BCU"`
	HK        float64 `json:"HK"`
	HK_NCL    float64 `json:"HK_NCL"`
	HK_NCU    float64 `json:"HK_NCU"`
	HK_BCL    float64 `json:"HK_BCL"`
	HK_BCU    float64 `json:"HK_BCU"`
	HSS       float64 `json:"HSS"`
	HSS_BCL   float64 `json:"HSS_BCL"`
	HSS_BCU   float64 `json:"HSS_BCU"`
	ODDS      float64 `json:"ODDS"`
	ODDS_NCL  float64 `json:"ODDS_NCL"`
	ODDS_NCU  float64 `json:"ODDS_NCU"`
	ODDS_BCL  float64 `json:"ODDS_BCL"`
	ODDS_BCU  float64 `json:"ODDS_BCU"`
	LODDS     float64 `json:"LODDS"`
	LODDS_NCL float64 `json:"LODDS_NCL"`
	LODDS_NCU float64 `json:"LODDS_NCU"`
	LODDS_BCL float64 `json:"LODDS_BCL"`
	LODDS_BCU float64 `json:"LODDS_BCU"`
	ORSS      float64 `json:"ORSS"`
	ORSS_NCL  float64 `json:"ORSS_NCL"`
	ORSS_NCU  float64 `json:"ORSS_NCU"`
	ORSS_BCL  float64 `json:"ORSS_BCL"`
	ORSS_BCU  float64 `json:"ORSS_BCU"`
	EDS       float64 `json:"EDS"`
	EDS_NCL   float64 `json:"EDS_NCL"`
	EDS_NCU   float64 `json:"EDS_NCU"`
	EDS_BCL   float64 `json:"EDS_BCL"`
	EDS_BCU   float64 `json:"EDS_BCU"`
	SEDS      float64 `json:"SEDS"`
	SEDS_NCL  float64 `json:"SEDS_NCL"`
	SEDS_NCU  float64 `json:"SEDS_NCU"`
	SEDS_BCL  float64 `json:"SEDS_BCL"`
	SEDS_BCU  float64 `json:"SEDS_BCU"`
	EDI       float64 `json:"EDI"`
	EDI_NCL   float64 `json:"EDI_NCL"`
	EDI_NCU   float64 `json:"EDI_NCU"`
	EDI_BCL   float64 `json:"EDI_BCL"`
	EDI_BCU   float64 `json:"EDI_BCU"`
	SEDI      float64 `json:"SEDI"`
	SEDI_NCL  float64 `json:"SEDI_NCL"`
	SEDI_NCU  float64 `json:"SEDI_NCU"`
	SEDI_BCL  float64 `json:"SEDI_BCL"`
	SEDI_BCU  float64 `json:"SEDI_BCU"`
	BAGSS     float64 `json:"BAGSS"`
	BAGSS_BCL float64 `json:"BAGSS_BCL"`
	BAGSS_BCU float64 `json:"BAGSS_BCU"`
}

type STAT_ECNT struct {
	TOTAL            int     `json:"TOTAL"`
	N_ENS            int     `json:"N_ENS"`
	CRPS             float64 `json:"CRPS"`
	CRPSS            float64 `json:"CRPSS"`
	IGN              float64 `json:"IGN"`
	ME               float64 `json:"ME"`
	RMSE             float64 `json:"RMSE"`
	SPREAD           float64 `json:"SPREAD"`
	ME_OERR          float64 `json:"ME_OERR"`
	RMSE_OERR        float64 `json:"RMSE_OERR"`
	SPREAD_OERR      float64 `json:"SPREAD_OERR"`
	SPREAD_PLUS_OERR float64 `json:"SPREAD_PLUS_OERR"`
	CRPSCL           float64 `json:"CRPSCL"`
	CRPS_EMP         float64 `json:"CRPS_EMP"`
	CRPSCL_EMP       float64 `json:"CRPSCL_EMP"`
	CRPSS_EMP        float64 `json:"CRPSS_EMP"`
	CRPS_EMP_FAIR    float64 `json:"CRPS_EMP_FAIR"`
	SPREAD_MD        float64 `json:"SPREAD_MD"`
	MAE              float64 `json:"MAE"`
	MAE_OERR         float64 `json:"MAE_OERR"`
	BIAS_RATIO       float64 `json:"BIAS_RATIO"`
	N_GE_OBS         int     `json:"N_GE_OBS"`
	ME_GE_OBS        float64 `json:"ME_GE_OBS"`
	N_LT_OBS         int     `json:"N_LT_OBS"`
	ME_LT_OBS        float64 `json:"ME_LT_OBS"`
	IGN_CONV_OERR    float64 `json:"IGN_CONV_OERR"`
	IGN_CORR_OERR    float64 `json:"IGN_CORR_OERR"`
}

type STAT_MCTS struct {
	TOTAL      int     `json:"TOTAL"`
	N_CAT      int     `json:"N_CAT"`
	ACC        float64 `json:"ACC"`
	ACC_NCL    float64 `json:"ACC_NCL"`
	ACC_NCU    float64 `json:"ACC_NCU"`
	ACC_BCL    float64 `json:"ACC_BCL"`
	ACC_BCU    float64 `json:"ACC_BCU"`
	HK         float64 `json:"HK"`
	HK_BCL     float64 `json:"HK_BCL"`
	HK_BCU     float64 `json:"HK_BCU"`
	HSS        float64 `json:"HSS"`
	HSS_BCL    float64 `json:"HSS_BCL"`
	HSS_BCU    float64 `json:"HSS_BCU"`
	GER        float64 `json:"GER"`
	GER_BCL    float64 `json:"GER_BCL"`
	GER_BCU    float64 `json:"GER_BCU"`
	HSS_EC     float64 `json:"HSS_EC"`
	HSS_EC_BCL float64 `json:"HSS_EC_BCL"`
	HSS_EC_BCU float64 `json:"HSS_EC_BCU"`
	EC_VALUE   float64 `json:"EC_VALUE"`
}

type STAT_PRC struct {
	TOTAL    int                    `json:"TOTAL"`
	THRESH   map[string]interface{} `json:"THRESH"`
	THRESH_I int                    `json:"THRESH_I"`
	PODY_I   float64                `json:"PODY_I"`
	POFD_I   float64                `json:"POFD_I"`
}

type STAT_VCNT struct {
	TOTAL                int     `json:"TOTAL"`
	FBAR                 float64 `json:"FBAR"`
	FBAR_BCL             float64 `json:"FBAR_BCL"`
	FBAR_BCU             float64 `json:"FBAR_BCU"`
	OBAR                 float64 `json:"OBAR"`
	OBAR_BCL             float64 `json:"OBAR_BCL"`
	OBAR_BCU             float64 `json:"OBAR_BCU"`
	FS_RMS               float64 `json:"FS_RMS"`
	FS_RMS_BCL           float64 `json:"FS_RMS_BCL"`
	FS_RMS_BCU           float64 `json:"FS_RMS_BCU"`
	OS_RMS               float64 `json:"OS_RMS"`
	OS_RMS_BCL           float64 `json:"OS_RMS_BCL"`
	OS_RMS_BCU           float64 `json:"OS_RMS_BCU"`
	MSVE                 float64 `json:"MSVE"`
	MSVE_BCL             float64 `json:"MSVE_BCL"`
	MSVE_BCU             float64 `json:"MSVE_BCU"`
	RMSVE                float64 `json:"RMSVE"`
	RMSVE_BCL            float64 `json:"RMSVE_BCL"`
	RMSVE_BCU            float64 `json:"RMSVE_BCU"`
	FSTDEV               float64 `json:"FSTDEV"`
	FSTDEV_BCL           float64 `json:"FSTDEV_BCL"`
	FSTDEV_BCU           float64 `json:"FSTDEV_BCU"`
	OSTDEV               float64 `json:"OSTDEV"`
	OSTDEV_BCL           float64 `json:"OSTDEV_BCL"`
	OSTDEV_BCU           float64 `json:"OSTDEV_BCU"`
	FDIR                 float64 `json:"FDIR"`
	FDIR_BCL             float64 `json:"FDIR_BCL"`
	FDIR_BCU             float64 `json:"FDIR_BCU"`
	ODIR                 float64 `json:"ODIR"`
	ODIR_BCL             float64 `json:"ODIR_BCL"`
	ODIR_BCU             float64 `json:"ODIR_BCU"`
	FBAR_SPEED           float64 `json:"FBAR_SPEED"`
	FBAR_SPEED_BCL       float64 `json:"FBAR_SPEED_BCL"`
	FBAR_SPEED_BCU       float64 `json:"FBAR_SPEED_BCU"`
	OBAR_SPEED           float64 `json:"OBAR_SPEED"`
	OBAR_SPEED_BCL       float64 `json:"OBAR_SPEED_BCL"`
	OBAR_SPEED_BCU       float64 `json:"OBAR_SPEED_BCU"`
	VDIFF_SPEED          float64 `json:"VDIFF_SPEED"`
	VDIFF_SPEED_BCL      float64 `json:"VDIFF_SPEED_BCL"`
	VDIFF_SPEED_BCU      float64 `json:"VDIFF_SPEED_BCU"`
	VDIFF_DIR            float64 `json:"VDIFF_DIR"`
	VDIFF_DIR_BCL        float64 `json:"VDIFF_DIR_BCL"`
	VDIFF_DIR_BCU        float64 `json:"VDIFF_DIR_BCU"`
	SPEED_ERR            float64 `json:"SPEED_ERR"`
	SPEED_ERR_BCL        float64 `json:"SPEED_ERR_BCL"`
	SPEED_ERR_BCU        float64 `json:"SPEED_ERR_BCU"`
	SPEED_ABSERR         float64 `json:"SPEED_ABSERR"`
	SPEED_ABSERR_BCL     float64 `json:"SPEED_ABSERR_BCL"`
	SPEED_ABSERR_BCU     float64 `json:"SPEED_ABSERR_BCU"`
	DIR_ERR              float64 `json:"DIR_ERR"`
	DIR_ERR_BCL          float64 `json:"DIR_ERR_BCL"`
	DIR_ERR_BCU          float64 `json:"DIR_ERR_BCU"`
	DIR_ABSERR           float64 `json:"DIR_ABSERR"`
	DIR_ABSERR_BCL       float64 `json:"DIR_ABSERR_BCL"`
	DIR_ABSERR_BCU       float64 `json:"DIR_ABSERR_BCU"`
	ANOM_CORR            float64 `json:"ANOM_CORR"`
	ANOM_CORR_NCL        float64 `json:"ANOM_CORR_NCL"`
	ANOM_CORR_NCU        float64 `json:"ANOM_CORR_NCU"`
	ANOM_CORR_BCL        float64 `json:"ANOM_CORR_BCL"`
	ANOM_CORR_BCU        float64 `json:"ANOM_CORR_BCU"`
	ANOM_CORR_UNCNTR     float64 `json:"ANOM_CORR_UNCNTR"`
	ANOM_CORR_UNCNTR_BCL float64 `json:"ANOM_CORR_UNCNTR_BCL"`
	ANOM_CORR_UNCNTR_BCU float64 `json:"ANOM_CORR_UNCNTR_BCU"`
	TOTAL_DIR            float64 `json:"TOTAL_DIR"`
	DIR_ME               float64 `json:"DIR_ME"`
	DIR_ME_BCL           float64 `json:"DIR_ME_BCL"`
	DIR_ME_BCU           float64 `json:"DIR_ME_BCU"`
	DIR_MAE              float64 `json:"DIR_MAE"`
	DIR_MAE_BCL          float64 `json:"DIR_MAE_BCL"`
	DIR_MAE_BCU          float64 `json:"DIR_MAE_BCU"`
	DIR_MSE              float64 `json:"DIR_MSE"`
	DIR_MSE_BCL          float64 `json:"DIR_MSE_BCL"`
	DIR_MSE_BCU          float64 `json:"DIR_MSE_BCU"`
	DIR_RMSE             float64 `json:"DIR_RMSE"`
	DIR_RMSE_BCL         float64 `json:"DIR_RMSE_BCL"`
	DIR_RMSE_BCU         float64 `json:"DIR_RMSE_BCU"`
}

type STAT_GENMPR struct {
	TOTAL      int     `json:"TOTAL"`
	INDEX      int     `json:"INDEX"`
	STORM_ID   string  `json:"STORM_ID"`
	PROB_LEAD  float64 `json:"PROB_LEAD"`
	PROB_VAL   float64 `json:"PROB_VAL"`
	AGEN_INIT  string  `json:"AGEN_INIT"`
	AGEN_FHR   string  `json:"AGEN_FHR"`
	AGEN_LAT   float64 `json:"AGEN_LAT"`
	AGEN_LON   float64 `json:"AGEN_LON"`
	AGEN_DLAND float64 `json:"AGEN_DLAND"`
	BGEN_LAT   float64 `json:"BGEN_LAT"`
	BGEN_LON   float64 `json:"BGEN_LON"`
	BGEN_DLAND float64 `json:"BGEN_DLAND"`
	GEN_DIST   float64 `json:"GEN_DIST"`
	GEN_TDIFF  string  `json:"GEN_TDIFF"`
	INIT_TDIFF string  `json:"INIT_TDIFF"`
	DEV_CAT    string  `json:"DEV_CAT"`
	OPS_CAT    string  `json:"OPS_CAT"`
}

type MODE_OBJ struct {
}

type MODE_CTS struct {
}

type STAT_SEEPS struct {
	TOTAL     int     `json:"TOTAL"`
	ODFL      float64 `json:"ODFL"`
	ODFH      float64 `json:"ODFH"`
	OLFD      float64 `json:"OLFD"`
	OLFH      float64 `json:"OLFH"`
	OHFD      float64 `json:"OHFD"`
	OHFL      float64 `json:"OHFL"`
	PF1       float64 `json:"PF1"`
	PF2       float64 `json:"PF2"`
	PF3       float64 `json:"PF3"`
	PV1       float64 `json:"PV1"`
	PV2       float64 `json:"PV2"`
	PV3       float64 `json:"PV3"`
	MEAN_FCST float64 `json:"MEAN_FCST"`
	MEAN_OBS  float64 `json:"MEAN_OBS"`
	SEEPS     float64 `json:"SEEPS"`
}

type STAT_NBRCTC struct {
	TOTAL int     `json:"TOTAL"`
	FY_OY float64 `json:"FY_OY"`
	FY_ON float64 `json:"FY_ON"`
	FN_OY float64 `json:"FN_OY"`
	FN_ON float64 `json:"FN_ON"`
}

type STAT_DMAP struct {
	TOTAL      int     `json:"TOTAL"`
	FY         int     `json:"FY"`
	OY         int     `json:"OY"`
	FBIAS      float64 `json:"FBIAS"`
	BADDELEY   float64 `json:"BADDELEY"`
	HAUSDORFF  float64 `json:"HAUSDORFF"`
	MED_FO     float64 `json:"MED_FO"`
	MED_OF     float64 `json:"MED_OF"`
	MED_MIN    float64 `json:"MED_MIN"`
	MED_MAX    float64 `json:"MED_MAX"`
	MED_MEAN   float64 `json:"MED_MEAN"`
	FOM_FO     float64 `json:"FOM_FO"`
	FOM_OF     float64 `json:"FOM_OF"`
	FOM_MIN    float64 `json:"FOM_MIN"`
	FOM_MAX    float64 `json:"FOM_MAX"`
	FOM_MEAN   float64 `json:"FOM_MEAN"`
	ZHU_FO     float64 `json:"ZHU_FO"`
	ZHU_OF     float64 `json:"ZHU_OF"`
	ZHU_MIN    float64 `json:"ZHU_MIN"`
	ZHU_MAX    float64 `json:"ZHU_MAX"`
	ZHU_MEAN   float64 `json:"ZHU_MEAN"`
	G          float64 `json:"G"`
	GBETA      float64 `json:"GBETA"`
	BETA_VALUE float64 `json:"BETA_VALUE"`
}

type STAT_RPS struct {
	TOTAL     int     `json:"TOTAL"`
	N_PROB    int     `json:"N_PROB"`
	RPS_REL   float64 `json:"RPS_REL"`
	RPS_RES   float64 `json:"RPS_RES"`
	RPS_UNC   float64 `json:"RPS_UNC"`
	RPS       float64 `json:"RPS"`
	RPSS      float64 `json:"RPSS"`
	RPSS_SMPL float64 `json:"RPSS_SMPL"`
	RPS_COMP  float64 `json:"RPS_COMP"`
}

type STAT_RELP struct {
	TOTAL  int                    `json:"TOTAL"`
	ENS    map[string]interface{} `json:"ENS"`
	RELP_I float64                `json:"RELP_I"`
}

type STAT_SAL1L2 struct {
	TOTAL  int     `json:"TOTAL"`
	FABAR  float64 `json:"FABAR"`
	OABAR  float64 `json:"OABAR"`
	FOABAR float64 `json:"FOABAR"`
	FFABAR float64 `json:"FFABAR"`
	OOABAR float64 `json:"OOABAR"`
	MAE    float64 `json:"MAE"`
}

type MTD_3DPAIR struct {
	OBJECT_ID           string  `json:"OBJECT_ID"`
	OBJECT_CAT          string  `json:"OBJECT_CAT"`
	SPACE_CENTROID_DIST float64 `json:"SPACE_CENTROID_DIST"`
	TIME_CENTROID_DELTA float64 `json:"TIME_CENTROID_DELTA"`
	AXIS_DIFF           float64 `json:"AXIS_DIFF"`
	SPEED_DELTA         float64 `json:"SPEED_DELTA"`
	DIRECTION_DIFF      float64 `json:"DIRECTION_DIFF"`
	VOLUME_RATIO        float64 `json:"VOLUME_RATIO"`
	START_TIME_DELTA    int     `json:"START_TIME_DELTA"`
	END_TIME_DELTA      int     `json:"END_TIME_DELTA"`
	INTERSECTION_VOLUME float64 `json:"INTERSECTION_VOLUME"`
	DURATION_DIFF       float64 `json:"DURATION_DIFF"`
	INTEREST            float64 `json:"INTEREST"`
}

type TCST_TCDIAG struct {
	TOTAL        int                    `json:"TOTAL"`
	INDEX        int                    `json:"INDEX"`
	DIAG_SOURCE  float64                `json:"DIAG_SOURCE"`
	TRACK_SOURCE string                 `json:"TRACK_SOURCE"`
	FIELD_SOURCE string                 `json:"FIELD_SOURCE"`
	DIAG         map[string]interface{} `json:"DIAG"`
	DIAG_I       float64                `json:"DIAG_I"`
	VALUE_I      int                    `json:"VALUE_I"`
}

type STAT_ISC struct {
	TOTAL    int     `json:"TOTAL"`
	TILE_DIM int     `json:"TILE_DIM"`
	TILE_XLL int     `json:"TILE_XLL"`
	TILE_YLL int     `json:"TILE_YLL"`
	NSCALE   int     `json:"NSCALE"`
	ISCALE   int     `json:"ISCALE"`
	MSE      float64 `json:"MSE"`
	ISC      float64 `json:"ISC"`
	FENERGY2 float64 `json:"FENERGY2"`
	OENERGY2 float64 `json:"OENERGY2"`
	BASER    float64 `json:"BASER"`
	FBIAS    float64 `json:"FBIAS"`
}

type STAT_PSTD struct {
	TOTAL       int                    `json:"TOTAL"`
	THRESH      map[string]interface{} `json:"THRESH"`
	BASER       float64                `json:"BASER"`
	BASER_NCL   float64                `json:"BASER_NCL"`
	BASER_NCU   float64                `json:"BASER_NCU"`
	RELIABILITY float64                `json:"RELIABILITY"`
	RESOLUTION  float64                `json:"RESOLUTION"`
	UNCERTAINTY float64                `json:"UNCERTAINTY"`
	ROC_AUC     float64                `json:"ROC_AUC"`
	BRIER       float64                `json:"BRIER"`
	BRIER_NCL   float64                `json:"BRIER_NCL"`
	BRIER_NCU   float64                `json:"BRIER_NCU"`
	BRIERCL     float64                `json:"BRIERCL"`
	BRIERCL_NCL float64                `json:"BRIERCL_NCL"`
	BRIERCL_NCU float64                `json:"BRIERCL_NCU"`
	BSS         float64                `json:"BSS"`
	BSS_SMPL    float64                `json:"BSS_SMPL"`
	THRESH_I    int                    `json:"THRESH_I"`
}

type STAT_SSVAR struct {
	TOTAL       int     `json:"TOTAL"`
	N_BIN       int     `json:"N_BIN"`
	BIN_i       int     `json:"BIN_i"`
	BIN_N       int     `json:"BIN_N"`
	VAR_MIN     float64 `json:"VAR_MIN"`
	VAR_MAX     float64 `json:"VAR_MAX"`
	VAR_MEAN    float64 `json:"VAR_MEAN"`
	FBAR        float64 `json:"FBAR"`
	OBAR        float64 `json:"OBAR"`
	FOBAR       float64 `json:"FOBAR"`
	FFBAR       float64 `json:"FFBAR"`
	OOBAR       float64 `json:"OOBAR"`
	FBAR_NCL    float64 `json:"FBAR_NCL"`
	FBAR_NCU    float64 `json:"FBAR_NCU"`
	FSTDEV      float64 `json:"FSTDEV"`
	FSTDEV_NCL  float64 `json:"FSTDEV_NCL"`
	FSTDEV_NCU  float64 `json:"FSTDEV_NCU"`
	OBAR_NCL    float64 `json:"OBAR_NCL"`
	OBAR_NCU    float64 `json:"OBAR_NCU"`
	OSTDEV      float64 `json:"OSTDEV"`
	OSTDEV_NCL  float64 `json:"OSTDEV_NCL"`
	OSTDEV_NCU  float64 `json:"OSTDEV_NCU"`
	PR_CORR     float64 `json:"PR_CORR"`
	PR_CORR_NCL float64 `json:"PR_CORR_NCL"`
	PR_CORR_NCU float64 `json:"PR_CORR_NCU"`
	ME          float64 `json:"ME"`
	ME_NCL      float64 `json:"ME_NCL"`
	ME_NCU      float64 `json:"ME_NCU"`
	ESTDEV      float64 `json:"ESTDEV"`
	ESTDEV_NCL  float64 `json:"ESTDEV_NCL"`
	ESTDEV_NCU  float64 `json:"ESTDEV_NCU"`
	MBIAS       float64 `json:"MBIAS"`
	MSE         float64 `json:"MSE"`
	BCMSE       float64 `json:"BCMSE"`
	RMSE        float64 `json:"RMSE"`
}

type MTD_3DSINGLE struct {
	OBJECT_ID       string  `json:"OBJECT_ID"`
	OBJECT_CAT      string  `json:"OBJECT_CAT"`
	CENTROID_X      float64 `json:"CENTROID_X"`
	CENTROID_Y      float64 `json:"CENTROID_Y"`
	CENTROID_T      float64 `json:"CENTROID_T"`
	CENTROID_LAT    float64 `json:"CENTROID_LAT"`
	CENTROID_LON    float64 `json:"CENTROID_LON"`
	X_DOT           float64 `json:"X_DOT"`
	Y_DOT           float64 `json:"Y_DOT"`
	AXIS_ANG        float64 `json:"AXIS_ANG"`
	VOLUME          int     `json:"VOLUME"`
	START_TIME      int     `json:"START_TIME"`
	END_TIME        int     `json:"END_TIME"`
	CDIST_TRAVELLED float64 `json:"CDIST_TRAVELLED"`
	INTENSITY_10    float64 `json:"INTENSITY_10"`
	INTENSITY_25    float64 `json:"INTENSITY_25"`
	INTENSITY_50    float64 `json:"INTENSITY_50"`
	INTENSITY_75    float64 `json:"INTENSITY_75"`
	INTENSITY_90    float64 `json:"INTENSITY_90"`
	INTENSITY_USER  float64 `json:"INTENSITY_USER"`
}

type STAT_NBRCNT struct {
	TOTAL      int     `json:"TOTAL"`
	FBS        float64 `json:"FBS"`
	FBS_BCL    float64 `json:"FBS_BCL"`
	FBS_BCU    float64 `json:"FBS_BCU"`
	FSS        float64 `json:"FSS"`
	FSS_BCL    float64 `json:"FSS_BCL"`
	FSS_BCU    float64 `json:"FSS_BCU"`
	AFSS       float64 `json:"AFSS"`
	AFSS_BCL   float64 `json:"AFSS_BCL"`
	AFSS_BCU   float64 `json:"AFSS_BCU"`
	UFSS       float64 `json:"UFSS"`
	UFSS_BCL   float64 `json:"UFSS_BCL"`
	UFSS_BCU   float64 `json:"UFSS_BCU"`
	F_RATE     float64 `json:"F_RATE"`
	F_RATE_BCL float64 `json:"F_RATE_BCL"`
	F_RATE_BCU float64 `json:"F_RATE_BCU"`
	O_RATE     float64 `json:"O_RATE"`
	O_RATE_BCL float64 `json:"O_RATE_BCL"`
	O_RATE_BCU float64 `json:"O_RATE_BCU"`
}

type STAT_ORANK struct {
	TOTAL            int                    `json:"TOTAL"`
	INDEX            int                    `json:"INDEX"`
	OBS_SID          string                 `json:"OBS_SID"`
	OBS_LAT          float64                `json:"OBS_LAT"`
	OBS_LON          float64                `json:"OBS_LON"`
	OBS_LVL          float64                `json:"OBS_LVL"`
	OBS_ELV          float64                `json:"OBS_ELV"`
	OBS              float64                `json:"OBS"`
	PIT              float64                `json:"PIT"`
	RANK             int                    `json:"RANK"`
	N_ENS_VLD        int                    `json:"N_ENS_VLD"`
	ENS              map[string]interface{} `json:"ENS"`
	ENS_I            int                    `json:"ENS_I"`
	OBS_QC           string                 `json:"OBS_QC"`
	ENS_MEAN         int                    `json:"ENS_MEAN"`
	OBS_CLIMO_MEAN   float64                `json:"OBS_CLIMO_MEAN"`
	SPREAD           float64                `json:"SPREAD"`
	ENS_MEAN_OERR    int                    `json:"ENS_MEAN_OERR"`
	SPREAD_OERR      float64                `json:"SPREAD_OERR"`
	SPREAD_PLUS_OERR float64                `json:"SPREAD_PLUS_OERR"`
	OBS_CLIMO_STDEV  float64                `json:"OBS_CLIMO_STDEV"`
	FCST_CLIMO_MEAN  float64                `json:"FCST_CLIMO_MEAN"`
	FCST_CLIMO_STDEV float64                `json:"FCST_CLIMO_STDEV"`
}

type STAT_PJC struct {
	TOTAL         int                    `json:"TOTAL"`
	THRESH        map[string]interface{} `json:"THRESH"`
	THRESH_I      int                    `json:"THRESH_I"`
	OY_TP_I       float64                `json:"OY_TP_I"`
	ON_TP_I       float64                `json:"ON_TP_I"`
	CALIBRATION_I float64                `json:"CALIBRATION_I"`
	REFINEMENT_I  float64                `json:"REFINEMENT_I"`
	LIKELIHOOD_I  float64                `json:"LIKELIHOOD_I"`
	BASER_I       map[string]float64     `json:"BASER_I"`
}

type STAT_SSIDX struct {
	FCST_MODEL string  `json:"FCST_MODEL"`
	REF_MODEL  string  `json:"REF_MODEL"`
	N_INIT     int     `json:"N_INIT"`
	N_TERM     int     `json:"N_TERM"`
	N_VLD      int     `json:"N_VLD"`
	SS_INDEX   float64 `json:"SS_INDEX"`
}

// fillStructure functions
func (s *STAT_RPS) fill_STAT_RPS(fields []string) {
	s.TOTAL, _ = strconv.Atoi(fields[0])
	s.N_PROB, _ = strconv.Atoi(fields[1])
	s.RPS_REL, _ = strconv.ParseFloat(fields[2], 64)
	s.RPS_RES, _ = strconv.ParseFloat(fields[3], 64)
	s.RPS_UNC, _ = strconv.ParseFloat(fields[4], 64)
	s.RPS, _ = strconv.ParseFloat(fields[5], 64)
	s.RPSS, _ = strconv.ParseFloat(fields[6], 64)
	s.RPSS_SMPL, _ = strconv.ParseFloat(fields[7], 64)
	s.RPS_COMP, _ = strconv.ParseFloat(fields[8], 64)
}

func (s *STAT_RHIST) fill_STAT_RHIST(fields []string) {
	s.TOTAL, _ = strconv.Atoi(fields[0])
	s.RANK = make(map[string]int)
	s.RANK_I, _ = strconv.Atoi(fields[2])
}

func (s *STAT_SAL1L2) fill_STAT_SAL1L2(fields []string) {
	s.TOTAL, _ = strconv.Atoi(fields[0])
	s.FABAR, _ = strconv.ParseFloat(fields[1], 64)
	s.OABAR, _ = strconv.ParseFloat(fields[2], 64)
	s.FOABAR, _ = strconv.ParseFloat(fields[3], 64)
	s.FFABAR, _ = strconv.ParseFloat(fields[4], 64)
	s.OOABAR, _ = strconv.ParseFloat(fields[5], 64)
	s.MAE, _ = strconv.ParseFloat(fields[6], 64)
}

func (s *STAT_SL1L2) fill_STAT_SL1L2(fields []string) {
	s.TOTAL, _ = strconv.Atoi(fields[0])
	s.FBAR, _ = strconv.ParseFloat(fields[1], 64)
	s.OBAR, _ = strconv.ParseFloat(fields[2], 64)
	s.FOBAR, _ = strconv.ParseFloat(fields[3], 64)
	s.FFBAR, _ = strconv.ParseFloat(fields[4], 64)
	s.OOBAR, _ = strconv.ParseFloat(fields[5], 64)
	s.MAE, _ = strconv.ParseFloat(fields[6], 64)
}

func (s *STAT_VAL1L2) fill_STAT_VAL1L2(fields []string) {
	s.TOTAL, _ = strconv.Atoi(fields[0])
	s.UFABAR, _ = strconv.ParseFloat(fields[1], 64)
	s.VFABAR, _ = strconv.ParseFloat(fields[2], 64)
	s.UOABAR, _ = strconv.ParseFloat(fields[3], 64)
	s.VOABAR, _ = strconv.ParseFloat(fields[4], 64)
	s.UVFOABAR, _ = strconv.ParseFloat(fields[5], 64)
	s.UVFFABAR, _ = strconv.ParseFloat(fields[6], 64)
	s.UVOOABAR, _ = strconv.ParseFloat(fields[7], 64)
	s.FA_SPEED_BAR, _ = strconv.ParseFloat(fields[8], 64)
	s.OA_SPEED_BAR, _ = strconv.ParseFloat(fields[9], 64)
	s.TOTAL_DIR, _ = strconv.ParseFloat(fields[10], 64)
	s.DIRA_ME, _ = strconv.ParseFloat(fields[11], 64)
	s.DIRA_MAE, _ = strconv.ParseFloat(fields[12], 64)
	s.DIRA_MSE, _ = strconv.ParseFloat(fields[13], 64)
}

func (s *STAT_PCT) fill_STAT_PCT(fields []string) {
	s.TOTAL, _ = strconv.Atoi(fields[0])
	s.THRESH = make(map[string]interface{})
	s.THRESH_I, _ = strconv.Atoi(fields[2])
	s.OY_I, _ = strconv.ParseFloat(fields[3], 64)
	s.ON_I, _ = strconv.ParseFloat(fields[4], 64)
}

func (s *STAT_CTS) fill_STAT_CTS(fields []string) {
	s.TOTAL, _ = strconv.Atoi(fields[0])
	s.BASER, _ = strconv.ParseFloat(fields[1], 64)
	s.BASER_NCL, _ = strconv.ParseFloat(fields[2], 64)
	s.BASER_NCU, _ = strconv.ParseFloat(fields[3], 64)
	s.BASER_BCL, _ = strconv.ParseFloat(fields[4], 64)
	s.BASER_BCU, _ = strconv.ParseFloat(fields[5], 64)
	s.FMEAN, _ = strconv.ParseFloat(fields[6], 64)
	s.FMEAN_NCL, _ = strconv.ParseFloat(fields[7], 64)
	s.FMEAN_NCU, _ = strconv.ParseFloat(fields[8], 64)
	s.FMEAN_BCL, _ = strconv.ParseFloat(fields[9], 64)
	s.FMEAN_BCU, _ = strconv.ParseFloat(fields[10], 64)
	s.ACC, _ = strconv.ParseFloat(fields[11], 64)
	s.ACC_NCL, _ = strconv.ParseFloat(fields[12], 64)
	s.ACC_NCU, _ = strconv.ParseFloat(fields[13], 64)
	s.ACC_BCL, _ = strconv.ParseFloat(fields[14], 64)
	s.ACC_BCU, _ = strconv.ParseFloat(fields[15], 64)
	s.FBIAS, _ = strconv.ParseFloat(fields[16], 64)
	s.FBIAS_BCL, _ = strconv.ParseFloat(fields[17], 64)
	s.FBIAS_BCU, _ = strconv.ParseFloat(fields[18], 64)
	s.PODY, _ = strconv.ParseFloat(fields[19], 64)
	s.PODY_NCL, _ = strconv.ParseFloat(fields[20], 64)
	s.PODY_NCU, _ = strconv.ParseFloat(fields[21], 64)
	s.PODY_BCL, _ = strconv.ParseFloat(fields[22], 64)
	s.PODY_BCU, _ = strconv.ParseFloat(fields[23], 64)
	s.PODN, _ = strconv.ParseFloat(fields[24], 64)
	s.PODN_NCL, _ = strconv.ParseFloat(fields[25], 64)
	s.PODN_NCU, _ = strconv.ParseFloat(fields[26], 64)
	s.PODN_BCL, _ = strconv.ParseFloat(fields[27], 64)
	s.PODN_BCU, _ = strconv.ParseFloat(fields[28], 64)
	s.POFD, _ = strconv.ParseFloat(fields[29], 64)
	s.POFD_NCL, _ = strconv.ParseFloat(fields[30], 64)
	s.POFD_NCU, _ = strconv.ParseFloat(fields[31], 64)
	s.POFD_BCL, _ = strconv.ParseFloat(fields[32], 64)
	s.POFD_BCU, _ = strconv.ParseFloat(fields[33], 64)
	s.FAR, _ = strconv.ParseFloat(fields[34], 64)
	s.FAR_NCL, _ = strconv.ParseFloat(fields[35], 64)
	s.FAR_NCU, _ = strconv.ParseFloat(fields[36], 64)
	s.FAR_BCL, _ = strconv.ParseFloat(fields[37], 64)
	s.FAR_BCU, _ = strconv.ParseFloat(fields[38], 64)
	s.CSI, _ = strconv.ParseFloat(fields[39], 64)
	s.CSI_NCL, _ = strconv.ParseFloat(fields[40], 64)
	s.CSI_NCU, _ = strconv.ParseFloat(fields[41], 64)
	s.CSI_BCL, _ = strconv.ParseFloat(fields[42], 64)
	s.CSI_BCU, _ = strconv.ParseFloat(fields[43], 64)
	s.GSS, _ = strconv.ParseFloat(fields[44], 64)
	s.GSS_BCL, _ = strconv.ParseFloat(fields[45], 64)
	s.GSS_BCU, _ = strconv.ParseFloat(fields[46], 64)
	s.HK, _ = strconv.ParseFloat(fields[47], 64)
	s.HK_NCL, _ = strconv.ParseFloat(fields[48], 64)
	s.HK_NCU, _ = strconv.ParseFloat(fields[49], 64)
	s.HK_BCL, _ = strconv.ParseFloat(fields[50], 64)
	s.HK_BCU, _ = strconv.ParseFloat(fields[51], 64)
	s.HSS, _ = strconv.ParseFloat(fields[52], 64)
	s.HSS_BCL, _ = strconv.ParseFloat(fields[53], 64)
	s.HSS_BCU, _ = strconv.ParseFloat(fields[54], 64)
	s.ODDS, _ = strconv.ParseFloat(fields[55], 64)
	s.ODDS_NCL, _ = strconv.ParseFloat(fields[56], 64)
	s.ODDS_NCU, _ = strconv.ParseFloat(fields[57], 64)
	s.ODDS_BCL, _ = strconv.ParseFloat(fields[58], 64)
	s.ODDS_BCU, _ = strconv.ParseFloat(fields[59], 64)
	s.LODDS, _ = strconv.ParseFloat(fields[60], 64)
	s.LODDS_NCL, _ = strconv.ParseFloat(fields[61], 64)
	s.LODDS_NCU, _ = strconv.ParseFloat(fields[62], 64)
	s.LODDS_BCL, _ = strconv.ParseFloat(fields[63], 64)
	s.LODDS_BCU, _ = strconv.ParseFloat(fields[64], 64)
	s.ORSS, _ = strconv.ParseFloat(fields[65], 64)
	s.ORSS_NCL, _ = strconv.ParseFloat(fields[66], 64)
	s.ORSS_NCU, _ = strconv.ParseFloat(fields[67], 64)
	s.ORSS_BCL, _ = strconv.ParseFloat(fields[68], 64)
	s.ORSS_BCU, _ = strconv.ParseFloat(fields[69], 64)
	s.EDS, _ = strconv.ParseFloat(fields[70], 64)
	s.EDS_NCL, _ = strconv.ParseFloat(fields[71], 64)
	s.EDS_NCU, _ = strconv.ParseFloat(fields[72], 64)
	s.EDS_BCL, _ = strconv.ParseFloat(fields[73], 64)
	s.EDS_BCU, _ = strconv.ParseFloat(fields[74], 64)
	s.SEDS, _ = strconv.ParseFloat(fields[75], 64)
	s.SEDS_NCL, _ = strconv.ParseFloat(fields[76], 64)
	s.SEDS_NCU, _ = strconv.ParseFloat(fields[77], 64)
	s.SEDS_BCL, _ = strconv.ParseFloat(fields[78], 64)
	s.SEDS_BCU, _ = strconv.ParseFloat(fields[79], 64)
	s.EDI, _ = strconv.ParseFloat(fields[80], 64)
	s.EDI_NCL, _ = strconv.ParseFloat(fields[81], 64)
	s.EDI_NCU, _ = strconv.ParseFloat(fields[82], 64)
	s.EDI_BCL, _ = strconv.ParseFloat(fields[83], 64)
	s.EDI_BCU, _ = strconv.ParseFloat(fields[84], 64)
	s.SEDI, _ = strconv.ParseFloat(fields[85], 64)
	s.SEDI_NCL, _ = strconv.ParseFloat(fields[86], 64)
	s.SEDI_NCU, _ = strconv.ParseFloat(fields[87], 64)
	s.SEDI_BCL, _ = strconv.ParseFloat(fields[88], 64)
	s.SEDI_BCU, _ = strconv.ParseFloat(fields[89], 64)
	s.BAGSS, _ = strconv.ParseFloat(fields[90], 64)
	s.BAGSS_BCL, _ = strconv.ParseFloat(fields[91], 64)
	s.BAGSS_BCU, _ = strconv.ParseFloat(fields[92], 64)
	s.HSS_EC, _ = strconv.ParseFloat(fields[93], 64)
	s.HSS_EC_BCL, _ = strconv.ParseFloat(fields[94], 64)
	s.HSS_EC_BCU, _ = strconv.ParseFloat(fields[95], 64)
	s.EC_VALUE, _ = strconv.ParseFloat(fields[96], 64)
}

func (s *STAT_RELP) fill_STAT_RELP(fields []string) {
	s.TOTAL, _ = strconv.Atoi(fields[0])
	s.ENS = make(map[string]interface{})
	s.RELP_I, _ = strconv.ParseFloat(fields[2], 64)
}

func (s *MODE_CTS) fill_MODE_CTS(fields []string) {
}

func (s *TCST_TCMPR) fill_TCST_TCMPR(fields []string) {
	s.TOTAL, _ = strconv.Atoi(fields[0])
	s.INDEX, _ = strconv.Atoi(fields[1])
	s.LEVEL = fields[2]
	s.WATCH_WARN = fields[3]
	s.INITIALS = fields[4]
	s.ALAT, _ = strconv.ParseFloat(fields[5], 64)
	s.ALON, _ = strconv.ParseFloat(fields[6], 64)
	s.BLAT, _ = strconv.ParseFloat(fields[7], 64)
	s.BLON, _ = strconv.ParseFloat(fields[8], 64)
	s.TK_ERR, _ = strconv.ParseFloat(fields[9], 64)
	s.X_ERR, _ = strconv.ParseFloat(fields[10], 64)
	s.Y_ERR, _ = strconv.ParseFloat(fields[11], 64)
	s.ALTK_ERR, _ = strconv.ParseFloat(fields[12], 64)
	s.CRTK_ERR, _ = strconv.ParseFloat(fields[13], 64)
	s.ADLAND, _ = strconv.ParseFloat(fields[14], 64)
	s.BDLAND, _ = strconv.ParseFloat(fields[15], 64)
	s.AMSLP, _ = strconv.ParseFloat(fields[16], 64)
	s.BMSLP, _ = strconv.ParseFloat(fields[17], 64)
	s.AMAX_WIND, _ = strconv.ParseFloat(fields[18], 64)
	s.BMAX_WIND, _ = strconv.ParseFloat(fields[19], 64)
	s.AAL_WIND_34, _ = strconv.ParseFloat(fields[20], 64)
	s.BAL_WIND_34, _ = strconv.ParseFloat(fields[21], 64)
	s.ANE_WIND_34, _ = strconv.ParseFloat(fields[22], 64)
	s.BNE_WIND_34, _ = strconv.ParseFloat(fields[23], 64)
	s.ASE_WIND_34, _ = strconv.ParseFloat(fields[24], 64)
	s.BSE_WIND_34, _ = strconv.ParseFloat(fields[25], 64)
	s.ASW_WIND_34, _ = strconv.ParseFloat(fields[26], 64)
	s.BSW_WIND_34, _ = strconv.ParseFloat(fields[27], 64)
	s.ANW_WIND_34, _ = strconv.ParseFloat(fields[28], 64)
	s.BNW_WIND_34, _ = strconv.ParseFloat(fields[29], 64)
	s.AAL_WIND_50, _ = strconv.ParseFloat(fields[30], 64)
	s.BAL_WIND_50, _ = strconv.ParseFloat(fields[31], 64)
	s.ANE_WIND_50, _ = strconv.ParseFloat(fields[32], 64)
	s.BNE_WIND_50, _ = strconv.ParseFloat(fields[33], 64)
	s.ASE_WIND_50, _ = strconv.ParseFloat(fields[34], 64)
	s.BSE_WIND_50, _ = strconv.ParseFloat(fields[35], 64)
	s.ASW_WIND_50, _ = strconv.ParseFloat(fields[36], 64)
	s.BSW_WIND_50, _ = strconv.ParseFloat(fields[37], 64)
	s.ANW_WIND_50, _ = strconv.ParseFloat(fields[38], 64)
	s.BNW_WIND_50, _ = strconv.ParseFloat(fields[39], 64)
	s.AAL_WIND_64, _ = strconv.ParseFloat(fields[40], 64)
	s.BAL_WIND_64, _ = strconv.ParseFloat(fields[41], 64)
	s.ANE_WIND_64, _ = strconv.ParseFloat(fields[42], 64)
	s.BNE_WIND_64, _ = strconv.ParseFloat(fields[43], 64)
	s.ASE_WIND_64, _ = strconv.ParseFloat(fields[44], 64)
	s.BSE_WIND_64, _ = strconv.ParseFloat(fields[45], 64)
	s.ASW_WIND_64, _ = strconv.ParseFloat(fields[46], 64)
	s.BSW_WIND_64, _ = strconv.ParseFloat(fields[47], 64)
	s.ANW_WIND_64, _ = strconv.ParseFloat(fields[48], 64)
	s.BNW_WIND_64, _ = strconv.ParseFloat(fields[49], 64)
	s.ARADP = fields[50]
	s.BRADP, _ = strconv.ParseFloat(fields[51], 64)
	s.ARRP, _ = strconv.Atoi(fields[52])
	s.BRRP, _ = strconv.ParseFloat(fields[53], 64)
	s.AMRD, _ = strconv.Atoi(fields[54])
	s.BMRD, _ = strconv.ParseFloat(fields[55], 64)
	s.AGUSTS, _ = strconv.Atoi(fields[56])
	s.BGUSTS, _ = strconv.ParseFloat(fields[57], 64)
	s.AEYE, _ = strconv.Atoi(fields[58])
	s.BEYE, _ = strconv.ParseFloat(fields[59], 64)
	s.ADIR, _ = strconv.Atoi(fields[60])
	s.BDIR, _ = strconv.ParseFloat(fields[61], 64)
	s.ASPEED, _ = strconv.Atoi(fields[62])
	s.BSPEED, _ = strconv.ParseFloat(fields[63], 64)
	s.ADEPTH, _ = strconv.Atoi(fields[64])
	s.BDEPTH, _ = strconv.ParseFloat(fields[65], 64)
	s.NUM_MEMBERS, _ = strconv.ParseFloat(fields[66], 64)
	s.TRACK_SPREAD, _ = strconv.ParseFloat(fields[67], 64)
	s.TRACK_STDEV, _ = strconv.ParseFloat(fields[68], 64)
	s.MSLP_STDEV, _ = strconv.ParseFloat(fields[69], 64)
	s.MAX_WIND_STDEV, _ = strconv.ParseFloat(fields[70], 64)
}

func (s *TCST_TCDIAG) fill_TCST_TCDIAG(fields []string) {
	s.TOTAL, _ = strconv.Atoi(fields[0])
	s.INDEX, _ = strconv.Atoi(fields[1])
	s.DIAG_SOURCE, _ = strconv.ParseFloat(fields[2], 64)
	s.TRACK_SOURCE = fields[3]
	s.FIELD_SOURCE = fields[4]
	s.DIAG = make(map[string]interface{})
	s.DIAG_I, _ = strconv.ParseFloat(fields[6], 64)
	s.VALUE_I, _ = strconv.Atoi(fields[7])
}

func (s *STAT_CTC) fill_STAT_CTC(fields []string) {
	s.TOTAL, _ = strconv.Atoi(fields[0])
	s.FY_OY, _ = strconv.ParseFloat(fields[1], 64)
	s.FY_ON, _ = strconv.ParseFloat(fields[2], 64)
	s.FN_OY, _ = strconv.ParseFloat(fields[3], 64)
	s.FN_ON, _ = strconv.ParseFloat(fields[4], 64)
	s.EC_VALUE, _ = strconv.ParseFloat(fields[5], 64)
}

func (s *STAT_SSVAR) fill_STAT_SSVAR(fields []string) {
	s.TOTAL, _ = strconv.Atoi(fields[0])
	s.N_BIN, _ = strconv.Atoi(fields[1])
	s.BIN_i, _ = strconv.Atoi(fields[2])
	s.BIN_N, _ = strconv.Atoi(fields[3])
	s.VAR_MIN, _ = strconv.ParseFloat(fields[4], 64)
	s.VAR_MAX, _ = strconv.ParseFloat(fields[5], 64)
	s.VAR_MEAN, _ = strconv.ParseFloat(fields[6], 64)
	s.FBAR, _ = strconv.ParseFloat(fields[7], 64)
	s.OBAR, _ = strconv.ParseFloat(fields[8], 64)
	s.FOBAR, _ = strconv.ParseFloat(fields[9], 64)
	s.FFBAR, _ = strconv.ParseFloat(fields[10], 64)
	s.OOBAR, _ = strconv.ParseFloat(fields[11], 64)
	s.FBAR_NCL, _ = strconv.ParseFloat(fields[12], 64)
	s.FBAR_NCU, _ = strconv.ParseFloat(fields[13], 64)
	s.FSTDEV, _ = strconv.ParseFloat(fields[14], 64)
	s.FSTDEV_NCL, _ = strconv.ParseFloat(fields[15], 64)
	s.FSTDEV_NCU, _ = strconv.ParseFloat(fields[16], 64)
	s.OBAR_NCL, _ = strconv.ParseFloat(fields[17], 64)
	s.OBAR_NCU, _ = strconv.ParseFloat(fields[18], 64)
	s.OSTDEV, _ = strconv.ParseFloat(fields[19], 64)
	s.OSTDEV_NCL, _ = strconv.ParseFloat(fields[20], 64)
	s.OSTDEV_NCU, _ = strconv.ParseFloat(fields[21], 64)
	s.PR_CORR, _ = strconv.ParseFloat(fields[22], 64)
	s.PR_CORR_NCL, _ = strconv.ParseFloat(fields[23], 64)
	s.PR_CORR_NCU, _ = strconv.ParseFloat(fields[24], 64)
	s.ME, _ = strconv.ParseFloat(fields[25], 64)
	s.ME_NCL, _ = strconv.ParseFloat(fields[26], 64)
	s.ME_NCU, _ = strconv.ParseFloat(fields[27], 64)
	s.ESTDEV, _ = strconv.ParseFloat(fields[28], 64)
	s.ESTDEV_NCL, _ = strconv.ParseFloat(fields[29], 64)
	s.ESTDEV_NCU, _ = strconv.ParseFloat(fields[30], 64)
	s.MBIAS, _ = strconv.ParseFloat(fields[31], 64)
	s.MSE, _ = strconv.ParseFloat(fields[32], 64)
	s.BCMSE, _ = strconv.ParseFloat(fields[33], 64)
	s.RMSE, _ = strconv.ParseFloat(fields[34], 64)
}

func (s *TCST_PROBRIRW) fill_TCST_PROBRIRW(fields []string) {
	s.ALAT, _ = strconv.ParseFloat(fields[0], 64)
	s.ALON, _ = strconv.ParseFloat(fields[1], 64)
	s.BLAT, _ = strconv.ParseFloat(fields[2], 64)
	s.BLON, _ = strconv.ParseFloat(fields[3], 64)
	s.INITIALS = fields[4]
	s.TK_ERR, _ = strconv.ParseFloat(fields[5], 64)
	s.X_ERR, _ = strconv.ParseFloat(fields[6], 64)
	s.Y_ERR, _ = strconv.ParseFloat(fields[7], 64)
	s.ADLAND, _ = strconv.ParseFloat(fields[8], 64)
	s.BDLAND, _ = strconv.ParseFloat(fields[9], 64)
	s.RIRW_BEG, _ = strconv.Atoi(fields[10])
	s.RIRW_END, _ = strconv.Atoi(fields[11])
	s.RIRW_WINDOW, _ = strconv.Atoi(fields[12])
	s.AWIND_END, _ = strconv.ParseFloat(fields[13], 64)
	s.BWIND_BEG, _ = strconv.ParseFloat(fields[14], 64)
	s.BWIND_END, _ = strconv.ParseFloat(fields[15], 64)
	s.BDELTA, _ = strconv.ParseFloat(fields[16], 64)
	s.BDELTA_MAX, _ = strconv.ParseFloat(fields[17], 64)
	s.BLEVEL_BEG = fields[18]
	s.BLEVEL_END = fields[19]
	s.THRESH = make(map[string]interface{})
	s.THRESH_I, _ = strconv.Atoi(fields[21])
	s.PROB_I, _ = strconv.ParseFloat(fields[22], 64)
}

func (s *STAT_FHO) fill_STAT_FHO(fields []string) {
	s.TOTAL, _ = strconv.Atoi(fields[0])
	s.F_RATE, _ = strconv.ParseFloat(fields[1], 64)
	s.H_RATE, _ = strconv.ParseFloat(fields[2], 64)
	s.O_RATE, _ = strconv.ParseFloat(fields[3], 64)
}

func (s *STAT_MCTS) fill_STAT_MCTS(fields []string) {
	s.TOTAL, _ = strconv.Atoi(fields[0])
	s.N_CAT, _ = strconv.Atoi(fields[1])
	s.ACC, _ = strconv.ParseFloat(fields[2], 64)
	s.ACC_NCL, _ = strconv.ParseFloat(fields[3], 64)
	s.ACC_NCU, _ = strconv.ParseFloat(fields[4], 64)
	s.ACC_BCL, _ = strconv.ParseFloat(fields[5], 64)
	s.ACC_BCU, _ = strconv.ParseFloat(fields[6], 64)
	s.HK, _ = strconv.ParseFloat(fields[7], 64)
	s.HK_BCL, _ = strconv.ParseFloat(fields[8], 64)
	s.HK_BCU, _ = strconv.ParseFloat(fields[9], 64)
	s.HSS, _ = strconv.ParseFloat(fields[10], 64)
	s.HSS_BCL, _ = strconv.ParseFloat(fields[11], 64)
	s.HSS_BCU, _ = strconv.ParseFloat(fields[12], 64)
	s.GER, _ = strconv.ParseFloat(fields[13], 64)
	s.GER_BCL, _ = strconv.ParseFloat(fields[14], 64)
	s.GER_BCU, _ = strconv.ParseFloat(fields[15], 64)
	s.HSS_EC, _ = strconv.ParseFloat(fields[16], 64)
	s.HSS_EC_BCL, _ = strconv.ParseFloat(fields[17], 64)
	s.HSS_EC_BCU, _ = strconv.ParseFloat(fields[18], 64)
	s.EC_VALUE, _ = strconv.ParseFloat(fields[19], 64)
}

func (s *STAT_VCNT) fill_STAT_VCNT(fields []string) {
	s.TOTAL, _ = strconv.Atoi(fields[0])
	s.FBAR, _ = strconv.ParseFloat(fields[1], 64)
	s.FBAR_BCL, _ = strconv.ParseFloat(fields[2], 64)
	s.FBAR_BCU, _ = strconv.ParseFloat(fields[3], 64)
	s.OBAR, _ = strconv.ParseFloat(fields[4], 64)
	s.OBAR_BCL, _ = strconv.ParseFloat(fields[5], 64)
	s.OBAR_BCU, _ = strconv.ParseFloat(fields[6], 64)
	s.FS_RMS, _ = strconv.ParseFloat(fields[7], 64)
	s.FS_RMS_BCL, _ = strconv.ParseFloat(fields[8], 64)
	s.FS_RMS_BCU, _ = strconv.ParseFloat(fields[9], 64)
	s.OS_RMS, _ = strconv.ParseFloat(fields[10], 64)
	s.OS_RMS_BCL, _ = strconv.ParseFloat(fields[11], 64)
	s.OS_RMS_BCU, _ = strconv.ParseFloat(fields[12], 64)
	s.MSVE, _ = strconv.ParseFloat(fields[13], 64)
	s.MSVE_BCL, _ = strconv.ParseFloat(fields[14], 64)
	s.MSVE_BCU, _ = strconv.ParseFloat(fields[15], 64)
	s.RMSVE, _ = strconv.ParseFloat(fields[16], 64)
	s.RMSVE_BCL, _ = strconv.ParseFloat(fields[17], 64)
	s.RMSVE_BCU, _ = strconv.ParseFloat(fields[18], 64)
	s.FSTDEV, _ = strconv.ParseFloat(fields[19], 64)
	s.FSTDEV_BCL, _ = strconv.ParseFloat(fields[20], 64)
	s.FSTDEV_BCU, _ = strconv.ParseFloat(fields[21], 64)
	s.OSTDEV, _ = strconv.ParseFloat(fields[22], 64)
	s.OSTDEV_BCL, _ = strconv.ParseFloat(fields[23], 64)
	s.OSTDEV_BCU, _ = strconv.ParseFloat(fields[24], 64)
	s.FDIR, _ = strconv.ParseFloat(fields[25], 64)
	s.FDIR_BCL, _ = strconv.ParseFloat(fields[26], 64)
	s.FDIR_BCU, _ = strconv.ParseFloat(fields[27], 64)
	s.ODIR, _ = strconv.ParseFloat(fields[28], 64)
	s.ODIR_BCL, _ = strconv.ParseFloat(fields[29], 64)
	s.ODIR_BCU, _ = strconv.ParseFloat(fields[30], 64)
	s.FBAR_SPEED, _ = strconv.ParseFloat(fields[31], 64)
	s.FBAR_SPEED_BCL, _ = strconv.ParseFloat(fields[32], 64)
	s.FBAR_SPEED_BCU, _ = strconv.ParseFloat(fields[33], 64)
	s.OBAR_SPEED, _ = strconv.ParseFloat(fields[34], 64)
	s.OBAR_SPEED_BCL, _ = strconv.ParseFloat(fields[35], 64)
	s.OBAR_SPEED_BCU, _ = strconv.ParseFloat(fields[36], 64)
	s.VDIFF_SPEED, _ = strconv.ParseFloat(fields[37], 64)
	s.VDIFF_SPEED_BCL, _ = strconv.ParseFloat(fields[38], 64)
	s.VDIFF_SPEED_BCU, _ = strconv.ParseFloat(fields[39], 64)
	s.VDIFF_DIR, _ = strconv.ParseFloat(fields[40], 64)
	s.VDIFF_DIR_BCL, _ = strconv.ParseFloat(fields[41], 64)
	s.VDIFF_DIR_BCU, _ = strconv.ParseFloat(fields[42], 64)
	s.SPEED_ERR, _ = strconv.ParseFloat(fields[43], 64)
	s.SPEED_ERR_BCL, _ = strconv.ParseFloat(fields[44], 64)
	s.SPEED_ERR_BCU, _ = strconv.ParseFloat(fields[45], 64)
	s.SPEED_ABSERR, _ = strconv.ParseFloat(fields[46], 64)
	s.SPEED_ABSERR_BCL, _ = strconv.ParseFloat(fields[47], 64)
	s.SPEED_ABSERR_BCU, _ = strconv.ParseFloat(fields[48], 64)
	s.DIR_ERR, _ = strconv.ParseFloat(fields[49], 64)
	s.DIR_ERR_BCL, _ = strconv.ParseFloat(fields[50], 64)
	s.DIR_ERR_BCU, _ = strconv.ParseFloat(fields[51], 64)
	s.DIR_ABSERR, _ = strconv.ParseFloat(fields[52], 64)
	s.DIR_ABSERR_BCL, _ = strconv.ParseFloat(fields[53], 64)
	s.DIR_ABSERR_BCU, _ = strconv.ParseFloat(fields[54], 64)
	s.ANOM_CORR, _ = strconv.ParseFloat(fields[55], 64)
	s.ANOM_CORR_NCL, _ = strconv.ParseFloat(fields[56], 64)
	s.ANOM_CORR_NCU, _ = strconv.ParseFloat(fields[57], 64)
	s.ANOM_CORR_BCL, _ = strconv.ParseFloat(fields[58], 64)
	s.ANOM_CORR_BCU, _ = strconv.ParseFloat(fields[59], 64)
	s.ANOM_CORR_UNCNTR, _ = strconv.ParseFloat(fields[60], 64)
	s.ANOM_CORR_UNCNTR_BCL, _ = strconv.ParseFloat(fields[61], 64)
	s.ANOM_CORR_UNCNTR_BCU, _ = strconv.ParseFloat(fields[62], 64)
	s.TOTAL_DIR, _ = strconv.ParseFloat(fields[63], 64)
	s.DIR_ME, _ = strconv.ParseFloat(fields[64], 64)
	s.DIR_ME_BCL, _ = strconv.ParseFloat(fields[65], 64)
	s.DIR_ME_BCU, _ = strconv.ParseFloat(fields[66], 64)
	s.DIR_MAE, _ = strconv.ParseFloat(fields[67], 64)
	s.DIR_MAE_BCL, _ = strconv.ParseFloat(fields[68], 64)
	s.DIR_MAE_BCU, _ = strconv.ParseFloat(fields[69], 64)
	s.DIR_MSE, _ = strconv.ParseFloat(fields[70], 64)
	s.DIR_MSE_BCL, _ = strconv.ParseFloat(fields[71], 64)
	s.DIR_MSE_BCU, _ = strconv.ParseFloat(fields[72], 64)
	s.DIR_RMSE, _ = strconv.ParseFloat(fields[73], 64)
	s.DIR_RMSE_BCL, _ = strconv.ParseFloat(fields[74], 64)
	s.DIR_RMSE_BCU, _ = strconv.ParseFloat(fields[75], 64)
}

func (s *MTD_3DSINGLE) fill_MTD_3DSINGLE(fields []string) {
	s.OBJECT_ID = fields[0]
	s.OBJECT_CAT = fields[1]
	s.CENTROID_X, _ = strconv.ParseFloat(fields[2], 64)
	s.CENTROID_Y, _ = strconv.ParseFloat(fields[3], 64)
	s.CENTROID_T, _ = strconv.ParseFloat(fields[4], 64)
	s.CENTROID_LAT, _ = strconv.ParseFloat(fields[5], 64)
	s.CENTROID_LON, _ = strconv.ParseFloat(fields[6], 64)
	s.X_DOT, _ = strconv.ParseFloat(fields[7], 64)
	s.Y_DOT, _ = strconv.ParseFloat(fields[8], 64)
	s.AXIS_ANG, _ = strconv.ParseFloat(fields[9], 64)
	s.VOLUME, _ = strconv.Atoi(fields[10])
	s.START_TIME, _ = strconv.Atoi(fields[11])
	s.END_TIME, _ = strconv.Atoi(fields[12])
	s.CDIST_TRAVELLED, _ = strconv.ParseFloat(fields[13], 64)
	s.INTENSITY_10, _ = strconv.ParseFloat(fields[14], 64)
	s.INTENSITY_25, _ = strconv.ParseFloat(fields[15], 64)
	s.INTENSITY_50, _ = strconv.ParseFloat(fields[16], 64)
	s.INTENSITY_75, _ = strconv.ParseFloat(fields[17], 64)
	s.INTENSITY_90, _ = strconv.ParseFloat(fields[18], 64)
	s.INTENSITY_USER, _ = strconv.ParseFloat(fields[19], 64)
}

func (s *STAT_ISC) fill_STAT_ISC(fields []string) {
	s.TOTAL, _ = strconv.Atoi(fields[0])
	s.TILE_DIM, _ = strconv.Atoi(fields[1])
	s.TILE_XLL, _ = strconv.Atoi(fields[2])
	s.TILE_YLL, _ = strconv.Atoi(fields[3])
	s.NSCALE, _ = strconv.Atoi(fields[4])
	s.ISCALE, _ = strconv.Atoi(fields[5])
	s.MSE, _ = strconv.ParseFloat(fields[6], 64)
	s.ISC, _ = strconv.ParseFloat(fields[7], 64)
	s.FENERGY2, _ = strconv.ParseFloat(fields[8], 64)
	s.OENERGY2, _ = strconv.ParseFloat(fields[9], 64)
	s.BASER, _ = strconv.ParseFloat(fields[10], 64)
	s.FBIAS, _ = strconv.ParseFloat(fields[11], 64)
}

func (s *STAT_GENMPR) fill_STAT_GENMPR(fields []string) {
	s.TOTAL, _ = strconv.Atoi(fields[0])
	s.INDEX, _ = strconv.Atoi(fields[1])
	s.STORM_ID = fields[2]
	s.PROB_LEAD, _ = strconv.ParseFloat(fields[3], 64)
	s.PROB_VAL, _ = strconv.ParseFloat(fields[4], 64)
	s.AGEN_INIT = fields[5]
	s.AGEN_FHR = fields[6]
	s.AGEN_LAT, _ = strconv.ParseFloat(fields[7], 64)
	s.AGEN_LON, _ = strconv.ParseFloat(fields[8], 64)
	s.AGEN_DLAND, _ = strconv.ParseFloat(fields[9], 64)
	s.BGEN_LAT, _ = strconv.ParseFloat(fields[10], 64)
	s.BGEN_LON, _ = strconv.ParseFloat(fields[11], 64)
	s.BGEN_DLAND, _ = strconv.ParseFloat(fields[12], 64)
	s.GEN_DIST, _ = strconv.ParseFloat(fields[13], 64)
	s.GEN_TDIFF = fields[14]
	s.INIT_TDIFF = fields[15]
	s.DEV_CAT = fields[16]
	s.OPS_CAT = fields[17]
}

func (s *MTD_2DSINGLE) fill_MTD_2DSINGLE(fields []string) {
	s.OBJECT_ID = fields[0]
	s.OBJECT_CAT = fields[1]
	s.TIME_INDEX, _ = strconv.Atoi(fields[2])
	s.AREA, _ = strconv.Atoi(fields[3])
	s.CENTROID_X, _ = strconv.ParseFloat(fields[4], 64)
	s.CENTROID_Y, _ = strconv.ParseFloat(fields[5], 64)
	s.CENTROID_LAT, _ = strconv.ParseFloat(fields[6], 64)
	s.CENTROID_LON, _ = strconv.ParseFloat(fields[7], 64)
	s.AXIS_ANG, _ = strconv.ParseFloat(fields[8], 64)
	s.INTENSITY_10, _ = strconv.ParseFloat(fields[9], 64)
	s.INTENSITY_25, _ = strconv.ParseFloat(fields[10], 64)
	s.INTENSITY_50, _ = strconv.ParseFloat(fields[11], 64)
	s.INTENSITY_75, _ = strconv.ParseFloat(fields[12], 64)
	s.INTENSITY_90, _ = strconv.ParseFloat(fields[13], 64)
	s.INTENSITY_USER, _ = strconv.ParseFloat(fields[14], 64)
}

func (s *STAT_MCTC) fill_STAT_MCTC(fields []string) {
	s.TOTAL, _ = strconv.Atoi(fields[0])
	s.CAT = make(map[string]int)
	s.FI_OI = fields[2]
	s.EC_VALUE, _ = strconv.ParseFloat(fields[3], 64)
}

func (s *STAT_GRAD) fill_STAT_GRAD(fields []string) {
	s.TOTAL, _ = strconv.Atoi(fields[0])
	s.FGBAR, _ = strconv.ParseFloat(fields[1], 64)
	s.OGBAR, _ = strconv.ParseFloat(fields[2], 64)
	s.MGBAR, _ = strconv.ParseFloat(fields[3], 64)
	s.EGBAR, _ = strconv.ParseFloat(fields[4], 64)
	s.S1, _ = strconv.ParseFloat(fields[5], 64)
	s.S1_OG, _ = strconv.ParseFloat(fields[6], 64)
	s.FGOG_RATIO, _ = strconv.ParseFloat(fields[7], 64)
	s.DX, _ = strconv.ParseFloat(fields[8], 64)
	s.DY, _ = strconv.ParseFloat(fields[9], 64)
}

func (s *STAT_DMAP) fill_STAT_DMAP(fields []string) {
	s.TOTAL, _ = strconv.Atoi(fields[0])
	s.FY, _ = strconv.Atoi(fields[1])
	s.OY, _ = strconv.Atoi(fields[2])
	s.FBIAS, _ = strconv.ParseFloat(fields[3], 64)
	s.BADDELEY, _ = strconv.ParseFloat(fields[4], 64)
	s.HAUSDORFF, _ = strconv.ParseFloat(fields[5], 64)
	s.MED_FO, _ = strconv.ParseFloat(fields[6], 64)
	s.MED_OF, _ = strconv.ParseFloat(fields[7], 64)
	s.MED_MIN, _ = strconv.ParseFloat(fields[8], 64)
	s.MED_MAX, _ = strconv.ParseFloat(fields[9], 64)
	s.MED_MEAN, _ = strconv.ParseFloat(fields[10], 64)
	s.FOM_FO, _ = strconv.ParseFloat(fields[11], 64)
	s.FOM_OF, _ = strconv.ParseFloat(fields[12], 64)
	s.FOM_MIN, _ = strconv.ParseFloat(fields[13], 64)
	s.FOM_MAX, _ = strconv.ParseFloat(fields[14], 64)
	s.FOM_MEAN, _ = strconv.ParseFloat(fields[15], 64)
	s.ZHU_FO, _ = strconv.ParseFloat(fields[16], 64)
	s.ZHU_OF, _ = strconv.ParseFloat(fields[17], 64)
	s.ZHU_MIN, _ = strconv.ParseFloat(fields[18], 64)
	s.ZHU_MAX, _ = strconv.ParseFloat(fields[19], 64)
	s.ZHU_MEAN, _ = strconv.ParseFloat(fields[20], 64)
	s.G, _ = strconv.ParseFloat(fields[21], 64)
	s.GBETA, _ = strconv.ParseFloat(fields[22], 64)
	s.BETA_VALUE, _ = strconv.ParseFloat(fields[23], 64)
}

func (s *STAT_ORANK) fill_STAT_ORANK(fields []string) {
	s.TOTAL, _ = strconv.Atoi(fields[0])
	s.INDEX, _ = strconv.Atoi(fields[1])
	s.OBS_SID = fields[2]
	s.OBS_LAT, _ = strconv.ParseFloat(fields[3], 64)
	s.OBS_LON, _ = strconv.ParseFloat(fields[4], 64)
	s.OBS_LVL, _ = strconv.ParseFloat(fields[5], 64)
	s.OBS_ELV, _ = strconv.ParseFloat(fields[6], 64)
	s.OBS, _ = strconv.ParseFloat(fields[7], 64)
	s.PIT, _ = strconv.ParseFloat(fields[8], 64)
	s.RANK, _ = strconv.Atoi(fields[9])
	s.N_ENS_VLD, _ = strconv.Atoi(fields[10])
	s.ENS = make(map[string]interface{})
	s.ENS_I, _ = strconv.Atoi(fields[12])
	s.OBS_QC = fields[13]
	s.ENS_MEAN, _ = strconv.Atoi(fields[14])
	s.OBS_CLIMO_MEAN, _ = strconv.ParseFloat(fields[15], 64)
	s.SPREAD, _ = strconv.ParseFloat(fields[16], 64)
	s.ENS_MEAN_OERR, _ = strconv.Atoi(fields[17])
	s.SPREAD_OERR, _ = strconv.ParseFloat(fields[18], 64)
	s.SPREAD_PLUS_OERR, _ = strconv.ParseFloat(fields[19], 64)
	s.OBS_CLIMO_STDEV, _ = strconv.ParseFloat(fields[20], 64)
	s.FCST_CLIMO_MEAN, _ = strconv.ParseFloat(fields[21], 64)
	s.FCST_CLIMO_STDEV, _ = strconv.ParseFloat(fields[22], 64)
}

func (s *STAT_PRC) fill_STAT_PRC(fields []string) {
	s.TOTAL, _ = strconv.Atoi(fields[0])
	s.THRESH = make(map[string]interface{})
	s.THRESH_I, _ = strconv.Atoi(fields[2])
	s.PODY_I, _ = strconv.ParseFloat(fields[3], 64)
	s.POFD_I, _ = strconv.ParseFloat(fields[4], 64)
}

func (s *STAT_PSTD) fill_STAT_PSTD(fields []string) {
	s.TOTAL, _ = strconv.Atoi(fields[0])
	s.THRESH = make(map[string]interface{})
	s.BASER, _ = strconv.ParseFloat(fields[2], 64)
	s.BASER_NCL, _ = strconv.ParseFloat(fields[3], 64)
	s.BASER_NCU, _ = strconv.ParseFloat(fields[4], 64)
	s.RELIABILITY, _ = strconv.ParseFloat(fields[5], 64)
	s.RESOLUTION, _ = strconv.ParseFloat(fields[6], 64)
	s.UNCERTAINTY, _ = strconv.ParseFloat(fields[7], 64)
	s.ROC_AUC, _ = strconv.ParseFloat(fields[8], 64)
	s.BRIER, _ = strconv.ParseFloat(fields[9], 64)
	s.BRIER_NCL, _ = strconv.ParseFloat(fields[10], 64)
	s.BRIER_NCU, _ = strconv.ParseFloat(fields[11], 64)
	s.BRIERCL, _ = strconv.ParseFloat(fields[12], 64)
	s.BRIERCL_NCL, _ = strconv.ParseFloat(fields[13], 64)
	s.BRIERCL_NCU, _ = strconv.ParseFloat(fields[14], 64)
	s.BSS, _ = strconv.ParseFloat(fields[15], 64)
	s.BSS_SMPL, _ = strconv.ParseFloat(fields[16], 64)
	s.THRESH_I, _ = strconv.Atoi(fields[17])
}

func (s *STAT_ECLV) fill_STAT_ECLV(fields []string) {
	s.TOTAL, _ = strconv.Atoi(fields[0])
	s.BASER, _ = strconv.ParseFloat(fields[1], 64)
	s.VALUE_BASER, _ = strconv.Atoi(fields[2])
	s.PTS = make(map[string]interface{})
	s.CL_I, _ = strconv.ParseFloat(fields[4], 64)
	s.VALUE_I, _ = strconv.Atoi(fields[5])
}

func (s *STAT_ECNT) fill_STAT_ECNT(fields []string) {
	s.TOTAL, _ = strconv.Atoi(fields[0])
	s.N_ENS, _ = strconv.Atoi(fields[1])
	s.CRPS, _ = strconv.ParseFloat(fields[2], 64)
	s.CRPSS, _ = strconv.ParseFloat(fields[3], 64)
	s.IGN, _ = strconv.ParseFloat(fields[4], 64)
	s.ME, _ = strconv.ParseFloat(fields[5], 64)
	s.RMSE, _ = strconv.ParseFloat(fields[6], 64)
	s.SPREAD, _ = strconv.ParseFloat(fields[7], 64)
	s.ME_OERR, _ = strconv.ParseFloat(fields[8], 64)
	s.RMSE_OERR, _ = strconv.ParseFloat(fields[9], 64)
	s.SPREAD_OERR, _ = strconv.ParseFloat(fields[10], 64)
	s.SPREAD_PLUS_OERR, _ = strconv.ParseFloat(fields[11], 64)
	s.CRPSCL, _ = strconv.ParseFloat(fields[12], 64)
	s.CRPS_EMP, _ = strconv.ParseFloat(fields[13], 64)
	s.CRPSCL_EMP, _ = strconv.ParseFloat(fields[14], 64)
	s.CRPSS_EMP, _ = strconv.ParseFloat(fields[15], 64)
	s.CRPS_EMP_FAIR, _ = strconv.ParseFloat(fields[16], 64)
	s.SPREAD_MD, _ = strconv.ParseFloat(fields[17], 64)
	s.MAE, _ = strconv.ParseFloat(fields[18], 64)
	s.MAE_OERR, _ = strconv.ParseFloat(fields[19], 64)
	s.BIAS_RATIO, _ = strconv.ParseFloat(fields[20], 64)
	s.N_GE_OBS, _ = strconv.Atoi(fields[21])
	s.ME_GE_OBS, _ = strconv.ParseFloat(fields[22], 64)
	s.N_LT_OBS, _ = strconv.Atoi(fields[23])
	s.ME_LT_OBS, _ = strconv.ParseFloat(fields[24], 64)
	s.IGN_CONV_OERR, _ = strconv.ParseFloat(fields[25], 64)
	s.IGN_CORR_OERR, _ = strconv.ParseFloat(fields[26], 64)
}

func (s *STAT_NBRCTC) fill_STAT_NBRCTC(fields []string) {
	s.TOTAL, _ = strconv.Atoi(fields[0])
	s.FY_OY, _ = strconv.ParseFloat(fields[1], 64)
	s.FY_ON, _ = strconv.ParseFloat(fields[2], 64)
	s.FN_OY, _ = strconv.ParseFloat(fields[3], 64)
	s.FN_ON, _ = strconv.ParseFloat(fields[4], 64)
}

func (s *STAT_PHIST) fill_STAT_PHIST(fields []string) {
	s.TOTAL, _ = strconv.Atoi(fields[0])
	s.BIN_SIZE, _ = strconv.Atoi(fields[1])
	s.BIN = make(map[string]int)
	s.BIN_I, _ = strconv.Atoi(fields[3])
}

func (s *STAT_SEEPS) fill_STAT_SEEPS(fields []string) {
	s.TOTAL, _ = strconv.Atoi(fields[0])
	s.ODFL, _ = strconv.ParseFloat(fields[1], 64)
	s.ODFH, _ = strconv.ParseFloat(fields[2], 64)
	s.OLFD, _ = strconv.ParseFloat(fields[3], 64)
	s.OLFH, _ = strconv.ParseFloat(fields[4], 64)
	s.OHFD, _ = strconv.ParseFloat(fields[5], 64)
	s.OHFL, _ = strconv.ParseFloat(fields[6], 64)
	s.PF1, _ = strconv.ParseFloat(fields[7], 64)
	s.PF2, _ = strconv.ParseFloat(fields[8], 64)
	s.PF3, _ = strconv.ParseFloat(fields[9], 64)
	s.PV1, _ = strconv.ParseFloat(fields[10], 64)
	s.PV2, _ = strconv.ParseFloat(fields[11], 64)
	s.PV3, _ = strconv.ParseFloat(fields[12], 64)
	s.MEAN_FCST, _ = strconv.ParseFloat(fields[13], 64)
	s.MEAN_OBS, _ = strconv.ParseFloat(fields[14], 64)
	s.SEEPS, _ = strconv.ParseFloat(fields[15], 64)
}

func (s *STAT_SEEPS_MPR) fill_STAT_SEEPS_MPR(fields []string) {
	s.OBS_SID = fields[0]
	s.OBS_LAT, _ = strconv.ParseFloat(fields[1], 64)
	s.OBS_LON, _ = strconv.ParseFloat(fields[2], 64)
	s.FCST, _ = strconv.ParseFloat(fields[3], 64)
	s.OBS, _ = strconv.ParseFloat(fields[4], 64)
	s.OBS_QC = fields[5]
	s.FCST_CAT, _ = strconv.Atoi(fields[6])
	s.OBS_CAT, _ = strconv.Atoi(fields[7])
	s.P1, _ = strconv.ParseFloat(fields[8], 64)
	s.P2, _ = strconv.ParseFloat(fields[9], 64)
	s.T1, _ = strconv.ParseFloat(fields[10], 64)
	s.T2, _ = strconv.ParseFloat(fields[11], 64)
	s.SEEPS, _ = strconv.ParseFloat(fields[12], 64)
}

func (s *STAT_SSIDX) fill_STAT_SSIDX(fields []string) {
	s.FCST_MODEL = fields[0]
	s.REF_MODEL = fields[1]
	s.N_INIT, _ = strconv.Atoi(fields[2])
	s.N_TERM, _ = strconv.Atoi(fields[3])
	s.N_VLD, _ = strconv.Atoi(fields[4])
	s.SS_INDEX, _ = strconv.ParseFloat(fields[5], 64)
}

func (s *MTD_3DPAIR) fill_MTD_3DPAIR(fields []string) {
	s.OBJECT_ID = fields[0]
	s.OBJECT_CAT = fields[1]
	s.SPACE_CENTROID_DIST, _ = strconv.ParseFloat(fields[2], 64)
	s.TIME_CENTROID_DELTA, _ = strconv.ParseFloat(fields[3], 64)
	s.AXIS_DIFF, _ = strconv.ParseFloat(fields[4], 64)
	s.SPEED_DELTA, _ = strconv.ParseFloat(fields[5], 64)
	s.DIRECTION_DIFF, _ = strconv.ParseFloat(fields[6], 64)
	s.VOLUME_RATIO, _ = strconv.ParseFloat(fields[7], 64)
	s.START_TIME_DELTA, _ = strconv.Atoi(fields[8])
	s.END_TIME_DELTA, _ = strconv.Atoi(fields[9])
	s.INTERSECTION_VOLUME, _ = strconv.ParseFloat(fields[10], 64)
	s.DURATION_DIFF, _ = strconv.ParseFloat(fields[11], 64)
	s.INTEREST, _ = strconv.ParseFloat(fields[12], 64)
}

func (s *STAT_CNT) fill_STAT_CNT(fields []string) {
	s.TOTAL, _ = strconv.Atoi(fields[0])
	s.FBAR, _ = strconv.ParseFloat(fields[1], 64)
	s.FBAR_NCL, _ = strconv.ParseFloat(fields[2], 64)
	s.FBAR_NCU, _ = strconv.ParseFloat(fields[3], 64)
	s.FBAR_BCL, _ = strconv.ParseFloat(fields[4], 64)
	s.FBAR_BCU, _ = strconv.ParseFloat(fields[5], 64)
	s.FSTDEV, _ = strconv.ParseFloat(fields[6], 64)
	s.FSTDEV_NCL, _ = strconv.ParseFloat(fields[7], 64)
	s.FSTDEV_NCU, _ = strconv.ParseFloat(fields[8], 64)
	s.FSTDEV_BCL, _ = strconv.ParseFloat(fields[9], 64)
	s.FSTDEV_BCU, _ = strconv.ParseFloat(fields[10], 64)
	s.OBAR, _ = strconv.ParseFloat(fields[11], 64)
	s.OBAR_NCL, _ = strconv.ParseFloat(fields[12], 64)
	s.OBAR_NCU, _ = strconv.ParseFloat(fields[13], 64)
	s.OBAR_BCL, _ = strconv.ParseFloat(fields[14], 64)
	s.OBAR_BCU, _ = strconv.ParseFloat(fields[15], 64)
	s.OSTDEV, _ = strconv.ParseFloat(fields[16], 64)
	s.OSTDEV_NCL, _ = strconv.ParseFloat(fields[17], 64)
	s.OSTDEV_NCU, _ = strconv.ParseFloat(fields[18], 64)
	s.OSTDEV_BCL, _ = strconv.ParseFloat(fields[19], 64)
	s.OSTDEV_BCU, _ = strconv.ParseFloat(fields[20], 64)
	s.PR_CORR, _ = strconv.ParseFloat(fields[21], 64)
	s.PR_CORR_NCL, _ = strconv.ParseFloat(fields[22], 64)
	s.PR_CORR_NCU, _ = strconv.ParseFloat(fields[23], 64)
	s.PR_CORR_BCL, _ = strconv.ParseFloat(fields[24], 64)
	s.PR_CORR_BCU, _ = strconv.ParseFloat(fields[25], 64)
	s.SP_CORR, _ = strconv.ParseFloat(fields[26], 64)
	s.KT_CORR, _ = strconv.ParseFloat(fields[27], 64)
	s.RANKS, _ = strconv.Atoi(fields[28])
	s.FRANK_TIES, _ = strconv.Atoi(fields[29])
	s.ORANK_TIES, _ = strconv.Atoi(fields[30])
	s.ME, _ = strconv.ParseFloat(fields[31], 64)
	s.ME_NCL, _ = strconv.ParseFloat(fields[32], 64)
	s.ME_NCU, _ = strconv.ParseFloat(fields[33], 64)
	s.ME_BCL, _ = strconv.ParseFloat(fields[34], 64)
	s.ME_BCU, _ = strconv.ParseFloat(fields[35], 64)
	s.ESTDEV, _ = strconv.ParseFloat(fields[36], 64)
	s.ESTDEV_NCL, _ = strconv.ParseFloat(fields[37], 64)
	s.ESTDEV_NCU, _ = strconv.ParseFloat(fields[38], 64)
	s.ESTDEV_BCL, _ = strconv.ParseFloat(fields[39], 64)
	s.ESTDEV_BCU, _ = strconv.ParseFloat(fields[40], 64)
	s.MBIAS, _ = strconv.ParseFloat(fields[41], 64)
	s.MBIAS_BCL, _ = strconv.ParseFloat(fields[42], 64)
	s.MBIAS_BCU, _ = strconv.ParseFloat(fields[43], 64)
	s.MAE, _ = strconv.ParseFloat(fields[44], 64)
	s.MAE_BCL, _ = strconv.ParseFloat(fields[45], 64)
	s.MAE_BCU, _ = strconv.ParseFloat(fields[46], 64)
	s.MSE, _ = strconv.ParseFloat(fields[47], 64)
	s.MSE_BCL, _ = strconv.ParseFloat(fields[48], 64)
	s.MSE_BCU, _ = strconv.ParseFloat(fields[49], 64)
	s.BCMSE, _ = strconv.ParseFloat(fields[50], 64)
	s.BCMSE_BCL, _ = strconv.ParseFloat(fields[51], 64)
	s.BCMSE_BCU, _ = strconv.ParseFloat(fields[52], 64)
	s.RMSE, _ = strconv.ParseFloat(fields[53], 64)
	s.RMSE_BCL, _ = strconv.ParseFloat(fields[54], 64)
	s.RMSE_BCU, _ = strconv.ParseFloat(fields[55], 64)
	s.E10, _ = strconv.ParseFloat(fields[56], 64)
	s.E10_BCL, _ = strconv.ParseFloat(fields[57], 64)
	s.E10_BCU, _ = strconv.ParseFloat(fields[58], 64)
	s.E25, _ = strconv.ParseFloat(fields[59], 64)
	s.E25_BCL, _ = strconv.ParseFloat(fields[60], 64)
	s.E25_BCU, _ = strconv.ParseFloat(fields[61], 64)
	s.E50, _ = strconv.ParseFloat(fields[62], 64)
	s.E50_BCL, _ = strconv.ParseFloat(fields[63], 64)
	s.E50_BCU, _ = strconv.ParseFloat(fields[64], 64)
	s.E75, _ = strconv.ParseFloat(fields[65], 64)
	s.E75_BCL, _ = strconv.ParseFloat(fields[66], 64)
	s.E75_BCU, _ = strconv.ParseFloat(fields[67], 64)
	s.E90, _ = strconv.ParseFloat(fields[68], 64)
	s.E90_BCL, _ = strconv.ParseFloat(fields[69], 64)
	s.E90_BCU, _ = strconv.ParseFloat(fields[70], 64)
	s.EIQR, _ = strconv.ParseFloat(fields[71], 64)
	s.EIQR_BCL, _ = strconv.ParseFloat(fields[72], 64)
	s.EIQR_BCU, _ = strconv.ParseFloat(fields[73], 64)
	s.MAD, _ = strconv.ParseFloat(fields[74], 64)
	s.MAD_BCL, _ = strconv.ParseFloat(fields[75], 64)
	s.MAD_BCU, _ = strconv.ParseFloat(fields[76], 64)
	s.ANOM_CORR, _ = strconv.ParseFloat(fields[77], 64)
	s.ANOM_CORR_NCL, _ = strconv.ParseFloat(fields[78], 64)
	s.ANOM_CORR_NCU, _ = strconv.ParseFloat(fields[79], 64)
	s.ANOM_CORR_BCL, _ = strconv.ParseFloat(fields[80], 64)
	s.ANOM_CORR_BCU, _ = strconv.ParseFloat(fields[81], 64)
	s.ME2, _ = strconv.ParseFloat(fields[82], 64)
	s.ME2_BCL, _ = strconv.ParseFloat(fields[83], 64)
	s.ME2_BCU, _ = strconv.ParseFloat(fields[84], 64)
	s.MSESS, _ = strconv.ParseFloat(fields[85], 64)
	s.MSESS_BCL, _ = strconv.ParseFloat(fields[86], 64)
	s.MSESS_BCU, _ = strconv.ParseFloat(fields[87], 64)
	s.RMSFA, _ = strconv.ParseFloat(fields[88], 64)
	s.RMSFA_BCL, _ = strconv.ParseFloat(fields[89], 64)
	s.RMSFA_BCU, _ = strconv.ParseFloat(fields[90], 64)
	s.RMSOA, _ = strconv.ParseFloat(fields[91], 64)
	s.RMSOA_BCL, _ = strconv.ParseFloat(fields[92], 64)
	s.RMSOA_BCU, _ = strconv.ParseFloat(fields[93], 64)
	s.ANOM_CORR_UNCNTR, _ = strconv.ParseFloat(fields[94], 64)
	s.ANOM_CORR_UNCNTR_BCL, _ = strconv.ParseFloat(fields[95], 64)
	s.ANOM_CORR_UNCNTR_BCU, _ = strconv.ParseFloat(fields[96], 64)
	s.SI, _ = strconv.ParseFloat(fields[97], 64)
	s.SI_BCL, _ = strconv.ParseFloat(fields[98], 64)
	s.SI_BCU, _ = strconv.ParseFloat(fields[99], 64)
}

func (s *STAT_NBRCNT) fill_STAT_NBRCNT(fields []string) {
	s.TOTAL, _ = strconv.Atoi(fields[0])
	s.FBS, _ = strconv.ParseFloat(fields[1], 64)
	s.FBS_BCL, _ = strconv.ParseFloat(fields[2], 64)
	s.FBS_BCU, _ = strconv.ParseFloat(fields[3], 64)
	s.FSS, _ = strconv.ParseFloat(fields[4], 64)
	s.FSS_BCL, _ = strconv.ParseFloat(fields[5], 64)
	s.FSS_BCU, _ = strconv.ParseFloat(fields[6], 64)
	s.AFSS, _ = strconv.ParseFloat(fields[7], 64)
	s.AFSS_BCL, _ = strconv.ParseFloat(fields[8], 64)
	s.AFSS_BCU, _ = strconv.ParseFloat(fields[9], 64)
	s.UFSS, _ = strconv.ParseFloat(fields[10], 64)
	s.UFSS_BCL, _ = strconv.ParseFloat(fields[11], 64)
	s.UFSS_BCU, _ = strconv.ParseFloat(fields[12], 64)
	s.F_RATE, _ = strconv.ParseFloat(fields[13], 64)
	s.F_RATE_BCL, _ = strconv.ParseFloat(fields[14], 64)
	s.F_RATE_BCU, _ = strconv.ParseFloat(fields[15], 64)
	s.O_RATE, _ = strconv.ParseFloat(fields[16], 64)
	s.O_RATE_BCL, _ = strconv.ParseFloat(fields[17], 64)
	s.O_RATE_BCU, _ = strconv.ParseFloat(fields[18], 64)
}

func (s *STAT_NBRCTS) fill_STAT_NBRCTS(fields []string) {
	s.TOTAL, _ = strconv.Atoi(fields[0])
	s.BASER, _ = strconv.ParseFloat(fields[1], 64)
	s.BASER_NCL, _ = strconv.ParseFloat(fields[2], 64)
	s.BASER_NCU, _ = strconv.ParseFloat(fields[3], 64)
	s.BASER_BCL, _ = strconv.ParseFloat(fields[4], 64)
	s.BASER_BCU, _ = strconv.ParseFloat(fields[5], 64)
	s.FMEAN, _ = strconv.ParseFloat(fields[6], 64)
	s.FMEAN_NCL, _ = strconv.ParseFloat(fields[7], 64)
	s.FMEAN_NCU, _ = strconv.ParseFloat(fields[8], 64)
	s.FMEAN_BCL, _ = strconv.ParseFloat(fields[9], 64)
	s.FMEAN_BCU, _ = strconv.ParseFloat(fields[10], 64)
	s.ACC, _ = strconv.ParseFloat(fields[11], 64)
	s.ACC_NCL, _ = strconv.ParseFloat(fields[12], 64)
	s.ACC_NCU, _ = strconv.ParseFloat(fields[13], 64)
	s.ACC_BCL, _ = strconv.ParseFloat(fields[14], 64)
	s.ACC_BCU, _ = strconv.ParseFloat(fields[15], 64)
	s.FBIAS, _ = strconv.ParseFloat(fields[16], 64)
	s.FBIAS_BCL, _ = strconv.ParseFloat(fields[17], 64)
	s.FBIAS_BCU, _ = strconv.ParseFloat(fields[18], 64)
	s.PODY, _ = strconv.ParseFloat(fields[19], 64)
	s.PODY_NCL, _ = strconv.ParseFloat(fields[20], 64)
	s.PODY_NCU, _ = strconv.ParseFloat(fields[21], 64)
	s.PODY_BCL, _ = strconv.ParseFloat(fields[22], 64)
	s.PODY_BCU, _ = strconv.ParseFloat(fields[23], 64)
	s.PODN, _ = strconv.ParseFloat(fields[24], 64)
	s.PODN_NCL, _ = strconv.ParseFloat(fields[25], 64)
	s.PODN_NCU, _ = strconv.ParseFloat(fields[26], 64)
	s.PODN_BCL, _ = strconv.ParseFloat(fields[27], 64)
	s.PODN_BCU, _ = strconv.ParseFloat(fields[28], 64)
	s.POFD, _ = strconv.ParseFloat(fields[29], 64)
	s.POFD_NCL, _ = strconv.ParseFloat(fields[30], 64)
	s.POFD_NCU, _ = strconv.ParseFloat(fields[31], 64)
	s.POFD_BCL, _ = strconv.ParseFloat(fields[32], 64)
	s.POFD_BCU, _ = strconv.ParseFloat(fields[33], 64)
	s.FAR, _ = strconv.ParseFloat(fields[34], 64)
	s.FAR_NCL, _ = strconv.ParseFloat(fields[35], 64)
	s.FAR_NCU, _ = strconv.ParseFloat(fields[36], 64)
	s.FAR_BCL, _ = strconv.ParseFloat(fields[37], 64)
	s.FAR_BCU, _ = strconv.ParseFloat(fields[38], 64)
	s.CSI, _ = strconv.ParseFloat(fields[39], 64)
	s.CSI_NCL, _ = strconv.ParseFloat(fields[40], 64)
	s.CSI_NCU, _ = strconv.ParseFloat(fields[41], 64)
	s.CSI_BCL, _ = strconv.ParseFloat(fields[42], 64)
	s.CSI_BCU, _ = strconv.ParseFloat(fields[43], 64)
	s.GSS, _ = strconv.ParseFloat(fields[44], 64)
	s.GSS_BCL, _ = strconv.ParseFloat(fields[45], 64)
	s.GSS_BCU, _ = strconv.ParseFloat(fields[46], 64)
	s.HK, _ = strconv.ParseFloat(fields[47], 64)
	s.HK_NCL, _ = strconv.ParseFloat(fields[48], 64)
	s.HK_NCU, _ = strconv.ParseFloat(fields[49], 64)
	s.HK_BCL, _ = strconv.ParseFloat(fields[50], 64)
	s.HK_BCU, _ = strconv.ParseFloat(fields[51], 64)
	s.HSS, _ = strconv.ParseFloat(fields[52], 64)
	s.HSS_BCL, _ = strconv.ParseFloat(fields[53], 64)
	s.HSS_BCU, _ = strconv.ParseFloat(fields[54], 64)
	s.ODDS, _ = strconv.ParseFloat(fields[55], 64)
	s.ODDS_NCL, _ = strconv.ParseFloat(fields[56], 64)
	s.ODDS_NCU, _ = strconv.ParseFloat(fields[57], 64)
	s.ODDS_BCL, _ = strconv.ParseFloat(fields[58], 64)
	s.ODDS_BCU, _ = strconv.ParseFloat(fields[59], 64)
	s.LODDS, _ = strconv.ParseFloat(fields[60], 64)
	s.LODDS_NCL, _ = strconv.ParseFloat(fields[61], 64)
	s.LODDS_NCU, _ = strconv.ParseFloat(fields[62], 64)
	s.LODDS_BCL, _ = strconv.ParseFloat(fields[63], 64)
	s.LODDS_BCU, _ = strconv.ParseFloat(fields[64], 64)
	s.ORSS, _ = strconv.ParseFloat(fields[65], 64)
	s.ORSS_NCL, _ = strconv.ParseFloat(fields[66], 64)
	s.ORSS_NCU, _ = strconv.ParseFloat(fields[67], 64)
	s.ORSS_BCL, _ = strconv.ParseFloat(fields[68], 64)
	s.ORSS_BCU, _ = strconv.ParseFloat(fields[69], 64)
	s.EDS, _ = strconv.ParseFloat(fields[70], 64)
	s.EDS_NCL, _ = strconv.ParseFloat(fields[71], 64)
	s.EDS_NCU, _ = strconv.ParseFloat(fields[72], 64)
	s.EDS_BCL, _ = strconv.ParseFloat(fields[73], 64)
	s.EDS_BCU, _ = strconv.ParseFloat(fields[74], 64)
	s.SEDS, _ = strconv.ParseFloat(fields[75], 64)
	s.SEDS_NCL, _ = strconv.ParseFloat(fields[76], 64)
	s.SEDS_NCU, _ = strconv.ParseFloat(fields[77], 64)
	s.SEDS_BCL, _ = strconv.ParseFloat(fields[78], 64)
	s.SEDS_BCU, _ = strconv.ParseFloat(fields[79], 64)
	s.EDI, _ = strconv.ParseFloat(fields[80], 64)
	s.EDI_NCL, _ = strconv.ParseFloat(fields[81], 64)
	s.EDI_NCU, _ = strconv.ParseFloat(fields[82], 64)
	s.EDI_BCL, _ = strconv.ParseFloat(fields[83], 64)
	s.EDI_BCU, _ = strconv.ParseFloat(fields[84], 64)
	s.SEDI, _ = strconv.ParseFloat(fields[85], 64)
	s.SEDI_NCL, _ = strconv.ParseFloat(fields[86], 64)
	s.SEDI_NCU, _ = strconv.ParseFloat(fields[87], 64)
	s.SEDI_BCL, _ = strconv.ParseFloat(fields[88], 64)
	s.SEDI_BCU, _ = strconv.ParseFloat(fields[89], 64)
	s.BAGSS, _ = strconv.ParseFloat(fields[90], 64)
	s.BAGSS_BCL, _ = strconv.ParseFloat(fields[91], 64)
	s.BAGSS_BCU, _ = strconv.ParseFloat(fields[92], 64)
}

func (s *STAT_PJC) fill_STAT_PJC(fields []string) {
	s.TOTAL, _ = strconv.Atoi(fields[0])
	s.THRESH = make(map[string]interface{})
	s.THRESH_I, _ = strconv.Atoi(fields[2])
	s.OY_TP_I, _ = strconv.ParseFloat(fields[3], 64)
	s.ON_TP_I, _ = strconv.ParseFloat(fields[4], 64)
	s.CALIBRATION_I, _ = strconv.ParseFloat(fields[5], 64)
	s.REFINEMENT_I, _ = strconv.ParseFloat(fields[6], 64)
	s.LIKELIHOOD_I, _ = strconv.ParseFloat(fields[7], 64)
	s.BASER_I = make(map[string]float64)
}

func (s *STAT_VL1L2) fill_STAT_VL1L2(fields []string) {
	s.TOTAL, _ = strconv.Atoi(fields[0])
	s.UFBAR, _ = strconv.ParseFloat(fields[1], 64)
	s.VFBAR, _ = strconv.ParseFloat(fields[2], 64)
	s.UOBAR, _ = strconv.ParseFloat(fields[3], 64)
	s.VOBAR, _ = strconv.ParseFloat(fields[4], 64)
	s.UVFOBAR, _ = strconv.ParseFloat(fields[5], 64)
	s.UVFFBAR, _ = strconv.ParseFloat(fields[6], 64)
	s.UVOOBAR, _ = strconv.ParseFloat(fields[7], 64)
	s.F_SPEED_BAR, _ = strconv.ParseFloat(fields[8], 64)
	s.O_SPEED_BAR, _ = strconv.ParseFloat(fields[9], 64)
	s.TOTAL_DIR, _ = strconv.ParseFloat(fields[10], 64)
	s.DIR_ME, _ = strconv.ParseFloat(fields[11], 64)
	s.DIR_MAE, _ = strconv.ParseFloat(fields[12], 64)
	s.DIR_MSE, _ = strconv.ParseFloat(fields[13], 64)
}

func (s *MODE_OBJ) fill_MODE_OBJ(fields []string) {
}

func (s *STAT_MPR) fill_STAT_MPR(fields []string) {
	s.TOTAL, _ = strconv.Atoi(fields[0])
	s.INDEX, _ = strconv.Atoi(fields[1])
	s.OBS_SID = fields[2]
	s.OBS_LAT, _ = strconv.ParseFloat(fields[3], 64)
	s.OBS_LON, _ = strconv.ParseFloat(fields[4], 64)
	s.OBS_LVL, _ = strconv.ParseFloat(fields[5], 64)
	s.OBS_ELV, _ = strconv.ParseFloat(fields[6], 64)
	s.FCST, _ = strconv.ParseFloat(fields[7], 64)
	s.OBS, _ = strconv.ParseFloat(fields[8], 64)
	s.OBS_QC = fields[9]
	s.OBS_CLIMO_MEAN, _ = strconv.ParseFloat(fields[10], 64)
	s.OBS_CLIMO_STDEV, _ = strconv.ParseFloat(fields[11], 64)
	s.OBS_CLIMO_CDF, _ = strconv.ParseFloat(fields[12], 64)
	s.FCST_CLIMO_MEAN, _ = strconv.ParseFloat(fields[13], 64)
	s.FCST_CLIMO_STDEV, _ = strconv.ParseFloat(fields[14], 64)
}

// getDocForId functions
func GetDocForId(fileLineType string, metaDataMap map[string]interface{}, headerData []string, dataData []string, dataKey string) (interface{}, error) {
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
		return nil, errors.New("Unknown file_line type:" + fileLineType)
	}
	return doc, nil
}

// addDataElement functions
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
