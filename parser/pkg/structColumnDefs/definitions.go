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


type vxMetadata struct {
	ID string `json:"ID"`
	// etc...
}

// header type definitions
// These have been derived from the header fields in the text files in
// the test_data/textfiles directory which are copied from the MET met_regression test data
// using the bin/build_header_types.sh script.

type statHeader_0 struct {
    Version string `json:"VERSION"`
    Model string `json:"MODEL"`
    Desc string `json:"DESC"`
    Fcst_lead string `json:"FCST_LEAD"`
    Fcst_valid string `json:"FCST_VALID"`
    Obs_lead string `json:"OBS_LEAD"`
    Obs_valid string `json:"OBS_VALID"`
    T_delta string `json:"T_DELTA"`
    Fcst_t_beg string `json:"FCST_T_BEG"`
    Fcst_t_end string `json:"FCST_T_END"`
    Fcst_rad string `json:"FCST_RAD"`
    Fcst_thr string `json:"FCST_THR"`
    Obs_t_beg string `json:"OBS_T_BEG"`
    Obs_t_end string `json:"OBS_T_END"`
    Obs_rad string `json:"OBS_RAD"`
    Obs_thr string `json:"OBS_THR"`
    Fcst_var string `json:"FCST_VAR"`
    Fcst_units string `json:"FCST_UNITS"`
    Fcst_lev string `json:"FCST_LEV"`
    Obs_var string `json:"OBS_VAR"`
    Obs_units string `json:"OBS_UNITS"`
    Obs_lev string `json:"OBS_LEV"`
    Object_id string `json:"OBJECT_ID"`
    Object_cat string `json:"OBJECT_CAT"`
    Centroid_x string `json:"CENTROID_X"`
    Centroid_y string `json:"CENTROID_Y"`
    Centroid_t string `json:"CENTROID_T"`
    Centroid_lat string `json:"CENTROID_LAT"`
    Centroid_lon string `json:"CENTROID_LON"`
    X_dot string `json:"X_DOT"`
    Y_dot string `json:"Y_DOT"`
    Axis_ang string `json:"AXIS_ANG"`
    Volume string `json:"VOLUME"`
    Start_time string `json:"START_TIME"`
    End_time string `json:"END_TIME"`
    Cdist_travelled string `json:"CDIST_TRAVELLED"`
    Intensity_10 string `json:"INTENSITY_10"`
    Intensity_25 string `json:"INTENSITY_25"`
    Intensity_50 string `json:"INTENSITY_50"`
    Intensity_75 string `json:"INTENSITY_75"`
    Intensity_90 string `json:"INTENSITY_90"`
    Intensity_99 string `json:"INTENSITY_99"`
    Alpha string `json:"ALPHA"`
    Line_type string `json:"LINE_TYPE"`
}


type statHeader_1 struct {
    Version string `json:"VERSION"`
    Model string `json:"MODEL"`
    Desc string `json:"DESC"`
    Fcst_lead string `json:"FCST_LEAD"`
    Fcst_valid_beg string `json:"FCST_VALID_BEG"`
    Fcst_valid_end string `json:"FCST_VALID_END"`
    Obs_lead string `json:"OBS_LEAD"`
    Obs_valid_beg string `json:"OBS_VALID_BEG"`
    Obs_valid_end string `json:"OBS_VALID_END"`
    Fcst_var string `json:"FCST_VAR"`
    Fcst_units string `json:"FCST_UNITS"`
    Fcst_lev string `json:"FCST_LEV"`
    Obs_var string `json:"OBS_VAR"`
    Obs_units string `json:"OBS_UNITS"`
    Obs_lev string `json:"OBS_LEV"`
    Obtype string `json:"OBTYPE"`
    Vx_mask string `json:"VX_MASK"`
    Interp_mthd string `json:"INTERP_MTHD"`
    Interp_pnts string `json:"INTERP_PNTS"`
    Fcst_thresh string `json:"FCST_THRESH"`
    Obs_thresh string `json:"OBS_THRESH"`
    Cov_thresh string `json:"COV_THRESH"`
    Alpha string `json:"ALPHA"`
    Line_type string `json:"LINE_TYPE"`
}


type statHeader_2 struct {
    Version string `json:"VERSION"`
    Model string `json:"MODEL"`
    N_valid string `json:"N_VALID"`
    Grid_res string `json:"GRID_RES"`
    Desc string `json:"DESC"`
    Fcst_lead string `json:"FCST_LEAD"`
    Fcst_valid string `json:"FCST_VALID"`
    Fcst_accum string `json:"FCST_ACCUM"`
    Obs_lead string `json:"OBS_LEAD"`
    Obs_valid string `json:"OBS_VALID"`
    Obs_accum string `json:"OBS_ACCUM"`
    Fcst_rad string `json:"FCST_RAD"`
    Fcst_thr string `json:"FCST_THR"`
    Obs_rad string `json:"OBS_RAD"`
    Obs_thr string `json:"OBS_THR"`
    Fcst_var string `json:"FCST_VAR"`
    Fcst_units string `json:"FCST_UNITS"`
    Fcst_lev string `json:"FCST_LEV"`
    Obs_var string `json:"OBS_VAR"`
    Obs_units string `json:"OBS_UNITS"`
    Obs_lev string `json:"OBS_LEV"`
    Obtype string `json:"OBTYPE"`
    Object_id string `json:"OBJECT_ID"`
    Object_cat string `json:"OBJECT_CAT"`
    Centroid_x string `json:"CENTROID_X"`
    Centroid_y string `json:"CENTROID_Y"`
    Centroid_lat string `json:"CENTROID_LAT"`
    Centroid_lon string `json:"CENTROID_LON"`
    Axis_ang string `json:"AXIS_ANG"`
    Length string `json:"LENGTH"`
    Width string `json:"WIDTH"`
    Area string `json:"AREA"`
    Area_thresh string `json:"AREA_THRESH"`
    Curvature string `json:"CURVATURE"`
    Curvature_x string `json:"CURVATURE_X"`
    Curvature_y string `json:"CURVATURE_Y"`
    Complexity string `json:"COMPLEXITY"`
    Intensity_10 string `json:"INTENSITY_10"`
    Intensity_25 string `json:"INTENSITY_25"`
    Intensity_50 string `json:"INTENSITY_50"`
    Intensity_75 string `json:"INTENSITY_75"`
    Intensity_90 string `json:"INTENSITY_90"`
    Intensity_user string `json:"INTENSITY_USER"`
    Intensity_sum string `json:"INTENSITY_SUM"`
    Centroid_dist string `json:"CENTROID_DIST"`
    Boundary_dist string `json:"BOUNDARY_DIST"`
    Convex_hull_dist string `json:"CONVEX_HULL_DIST"`
    Angle_diff string `json:"ANGLE_DIFF"`
    Aspect_diff string `json:"ASPECT_DIFF"`
    Area_ratio string `json:"AREA_RATIO"`
    Intersection_area string `json:"INTERSECTION_AREA"`
    Union_area string `json:"UNION_AREA"`
    Symmetric_diff string `json:"SYMMETRIC_DIFF"`
    Intersection_over_area string `json:"INTERSECTION_OVER_AREA"`
    Curvature_ratio string `json:"CURVATURE_RATIO"`
    Complexity_ratio string `json:"COMPLEXITY_RATIO"`
    Percentile_intensity_ratio string `json:"PERCENTILE_INTENSITY_RATIO"`
    Interest string `json:"INTEREST"`
    Alpha string `json:"ALPHA"`
    Line_type string `json:"LINE_TYPE"`
}


type statHeader_3 struct {
    Version string `json:"VERSION"`
    Model string `json:"MODEL"`
    N_valid string `json:"N_VALID"`
    Grid_res string `json:"GRID_RES"`
    Desc string `json:"DESC"`
    Fcst_lead string `json:"FCST_LEAD"`
    Fcst_valid string `json:"FCST_VALID"`
    Fcst_accum string `json:"FCST_ACCUM"`
    Obs_lead string `json:"OBS_LEAD"`
    Obs_valid string `json:"OBS_VALID"`
    Obs_accum string `json:"OBS_ACCUM"`
    Fcst_rad string `json:"FCST_RAD"`
    Fcst_thr string `json:"FCST_THR"`
    Obs_rad string `json:"OBS_RAD"`
    Obs_thr string `json:"OBS_THR"`
    Fcst_var string `json:"FCST_VAR"`
    Fcst_units string `json:"FCST_UNITS"`
    Fcst_lev string `json:"FCST_LEV"`
    Obs_var string `json:"OBS_VAR"`
    Obs_units string `json:"OBS_UNITS"`
    Obs_lev string `json:"OBS_LEV"`
    Obtype string `json:"OBTYPE"`
    Field string `json:"FIELD"`
    Total string `json:"TOTAL"`
    Fy_oy string `json:"FY_OY"`
    Fy_on string `json:"FY_ON"`
    Fn_oy string `json:"FN_OY"`
    Fn_on string `json:"FN_ON"`
    Baser string `json:"BASER"`
    Fmean string `json:"FMEAN"`
    Acc string `json:"ACC"`
    Fbias string `json:"FBIAS"`
    Pody string `json:"PODY"`
    Podn string `json:"PODN"`
    Pofd string `json:"POFD"`
    Far string `json:"FAR"`
    Csi string `json:"CSI"`
    Gss string `json:"GSS"`
    Hk string `json:"HK"`
    Hss string `json:"HSS"`
    Odds string `json:"ODDS"`
    Lodds string `json:"LODDS"`
    Orss string `json:"ORSS"`
    Eds string `json:"EDS"`
    Seds string `json:"SEDS"`
    Edi string `json:"EDI"`
    Sedi string `json:"SEDI"`
    Bagss string `json:"BAGSS"`
    Alpha string `json:"ALPHA"`
    Line_type string `json:"LINE_TYPE"`
}


type statHeader_4 struct {
    Version string `json:"VERSION"`
    Model string `json:"MODEL"`
    N_valid string `json:"N_VALID"`
    Grid_res string `json:"GRID_RES"`
    Desc string `json:"DESC"`
    Fcst_lead string `json:"FCST_LEAD"`
    Fcst_valid string `json:"FCST_VALID"`
    Fcst_accum string `json:"FCST_ACCUM"`
    Obs_lead string `json:"OBS_LEAD"`
    Obs_valid string `json:"OBS_VALID"`
    Obs_accum string `json:"OBS_ACCUM"`
    Fcst_rad string `json:"FCST_RAD"`
    Fcst_thr string `json:"FCST_THR"`
    Obs_rad string `json:"OBS_RAD"`
    Obs_thr string `json:"OBS_THR"`
    Fcst_var string `json:"FCST_VAR"`
    Fcst_units string `json:"FCST_UNITS"`
    Fcst_lev string `json:"FCST_LEV"`
    Obs_var string `json:"OBS_VAR"`
    Obs_units string `json:"OBS_UNITS"`
    Obs_lev string `json:"OBS_LEV"`
    Obtype string `json:"OBTYPE"`
    Object_id string `json:"OBJECT_ID"`
    Object_cat string `json:"OBJECT_CAT"`
    Centroid_x string `json:"CENTROID_X"`
    Centroid_y string `json:"CENTROID_Y"`
    Centroid_lat string `json:"CENTROID_LAT"`
    Centroid_lon string `json:"CENTROID_LON"`
    Axis_ang string `json:"AXIS_ANG"`
    Length string `json:"LENGTH"`
    Width string `json:"WIDTH"`
    Area string `json:"AREA"`
    Area_thresh string `json:"AREA_THRESH"`
    Curvature string `json:"CURVATURE"`
    Curvature_x string `json:"CURVATURE_X"`
    Curvature_y string `json:"CURVATURE_Y"`
    Complexity string `json:"COMPLEXITY"`
    Intensity_10 string `json:"INTENSITY_10"`
    Intensity_25 string `json:"INTENSITY_25"`
    Intensity_50 string `json:"INTENSITY_50"`
    Intensity_75 string `json:"INTENSITY_75"`
    Intensity_90 string `json:"INTENSITY_90"`
    Intensity_sum string `json:"INTENSITY_SUM"`
    Centroid_dist string `json:"CENTROID_DIST"`
    Boundary_dist string `json:"BOUNDARY_DIST"`
    Convex_hull_dist string `json:"CONVEX_HULL_DIST"`
    Angle_diff string `json:"ANGLE_DIFF"`
    Aspect_diff string `json:"ASPECT_DIFF"`
    Area_ratio string `json:"AREA_RATIO"`
    Intersection_area string `json:"INTERSECTION_AREA"`
    Union_area string `json:"UNION_AREA"`
    Symmetric_diff string `json:"SYMMETRIC_DIFF"`
    Intersection_over_area string `json:"INTERSECTION_OVER_AREA"`
    Curvature_ratio string `json:"CURVATURE_RATIO"`
    Complexity_ratio string `json:"COMPLEXITY_RATIO"`
    Percentile_intensity_ratio string `json:"PERCENTILE_INTENSITY_RATIO"`
    Interest string `json:"INTEREST"`
    Alpha string `json:"ALPHA"`
    Line_type string `json:"LINE_TYPE"`
}


type statHeader_5 struct {
    Version string `json:"VERSION"`
    Model string `json:"MODEL"`
    Desc string `json:"DESC"`
    Fcst_lead string `json:"FCST_LEAD"`
    Fcst_valid string `json:"FCST_VALID"`
    Obs_lead string `json:"OBS_LEAD"`
    Obs_valid string `json:"OBS_VALID"`
    T_delta string `json:"T_DELTA"`
    Fcst_t_beg string `json:"FCST_T_BEG"`
    Fcst_t_end string `json:"FCST_T_END"`
    Fcst_rad string `json:"FCST_RAD"`
    Fcst_thr string `json:"FCST_THR"`
    Obs_t_beg string `json:"OBS_T_BEG"`
    Obs_t_end string `json:"OBS_T_END"`
    Obs_rad string `json:"OBS_RAD"`
    Obs_thr string `json:"OBS_THR"`
    Fcst_var string `json:"FCST_VAR"`
    Fcst_units string `json:"FCST_UNITS"`
    Fcst_lev string `json:"FCST_LEV"`
    Obs_var string `json:"OBS_VAR"`
    Obs_units string `json:"OBS_UNITS"`
    Obs_lev string `json:"OBS_LEV"`
    Object_id string `json:"OBJECT_ID"`
    Object_cat string `json:"OBJECT_CAT"`
    Time_index string `json:"TIME_INDEX"`
    Area string `json:"AREA"`
    Centroid_x string `json:"CENTROID_X"`
    Centroid_y string `json:"CENTROID_Y"`
    Centroid_lat string `json:"CENTROID_LAT"`
    Centroid_lon string `json:"CENTROID_LON"`
    Axis_ang string `json:"AXIS_ANG"`
    Intensity_10 string `json:"INTENSITY_10"`
    Intensity_25 string `json:"INTENSITY_25"`
    Intensity_50 string `json:"INTENSITY_50"`
    Intensity_75 string `json:"INTENSITY_75"`
    Intensity_90 string `json:"INTENSITY_90"`
    Intensity_99 string `json:"INTENSITY_99"`
    Alpha string `json:"ALPHA"`
    Line_type string `json:"LINE_TYPE"`
}


type statHeader_6 struct {
    Version string `json:"VERSION"`
    Model string `json:"MODEL"`
    Desc string `json:"DESC"`
    Fcst_lead string `json:"FCST_LEAD"`
    Fcst_valid string `json:"FCST_VALID"`
    Obs_lead string `json:"OBS_LEAD"`
    Obs_valid string `json:"OBS_VALID"`
    T_delta string `json:"T_DELTA"`
    Fcst_t_beg string `json:"FCST_T_BEG"`
    Fcst_t_end string `json:"FCST_T_END"`
    Fcst_rad string `json:"FCST_RAD"`
    Fcst_thr string `json:"FCST_THR"`
    Obs_t_beg string `json:"OBS_T_BEG"`
    Obs_t_end string `json:"OBS_T_END"`
    Obs_rad string `json:"OBS_RAD"`
    Obs_thr string `json:"OBS_THR"`
    Fcst_var string `json:"FCST_VAR"`
    Fcst_units string `json:"FCST_UNITS"`
    Fcst_lev string `json:"FCST_LEV"`
    Obs_var string `json:"OBS_VAR"`
    Obs_units string `json:"OBS_UNITS"`
    Obs_lev string `json:"OBS_LEV"`
    Object_id string `json:"OBJECT_ID"`
    Object_cat string `json:"OBJECT_CAT"`
    Space_centroid_dist string `json:"SPACE_CENTROID_DIST"`
    Time_centroid_delta string `json:"TIME_CENTROID_DELTA"`
    Axis_diff string `json:"AXIS_DIFF"`
    Speed_delta string `json:"SPEED_DELTA"`
    Direction_diff string `json:"DIRECTION_DIFF"`
    Volume_ratio string `json:"VOLUME_RATIO"`
    Start_time_delta string `json:"START_TIME_DELTA"`
    End_time_delta string `json:"END_TIME_DELTA"`
    Intersection_volume string `json:"INTERSECTION_VOLUME"`
    Duration_diff string `json:"DURATION_DIFF"`
    Interest string `json:"INTEREST"`
    Alpha string `json:"ALPHA"`
    Line_type string `json:"LINE_TYPE"`
}

//line data definitions
// These have been derived from the data fields in the text files in
// the test_data/textfiles directory which are copied from the MET met_regression test data.
// using the bin/build_line_types.sh script.
type VL1L2Doc struct {
	vxMetadata
	statHeader_0
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
