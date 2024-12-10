package structColumnTypes

type vxMetadata struct {    ID string `json:"ID"`
    subset string `json:"MET"`
    version string `json:"version"`
    typ string `json:"type"`  // like "DD"
	subType string `json:"subtype"`  // like STAT, MODE, etc.
}

type STAT_header_V12_0 struct {
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


type MODE_header_V12_0 struct {
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
}


type TCST_header_V12_0 struct {
    Version string `json:"VERSION"`
    Amodel string `json:"AMODEL"`
    Bmodel string `json:"BMODEL"`
    Desc string `json:"DESC"`
    Storm_id string `json:"STORM_ID"`
    Basin string `json:"BASIN"`
    Cyclone string `json:"CYCLONE"`
    Storm_name string `json:"STORM_NAME"`
    Init string `json:"INIT"`
    Lead string `json:"LEAD"`
    Valid string `json:"VALID"`
    Init_mask string `json:"INIT_MASK"`
    Valid_mask string `json:"VALID_MASK"`
    Line_type string `json:"LINE_TYPE"`
}


type STAT_ECLV_V12_0 struct {
    vxMetadata
    STAT_header_V12_0
    Total int `json:"TOTAL"`
    Baser float64 `json:"BASER"`
    Value_baser string `json:"VALUE_BASER"`
    (n_pts) string `json:"(N_PTS)"`
    Cl_[0-9]* float64 `json:"CL_[0-9]*"`
    Value_[0-9]* string `json:"VALUE_[0-9]*"`
}


type STAT_GENMPR_V12_0 struct {
    vxMetadata
    STAT_header_V12_0
    Total int `json:"TOTAL"`
    Index int `json:"INDEX"`
    Storm_id string `json:"STORM_ID"`
    Prob_lead string `json:"PROB_LEAD"`
    Prob_val string `json:"PROB_VAL"`
    Agen_init string `json:"AGEN_INIT"`
    Agen_fhr string `json:"AGEN_FHR"`
    Agen_lat string `json:"AGEN_LAT"`
    Agen_lon string `json:"AGEN_LON"`
    Agen_dland string `json:"AGEN_DLAND"`
    Bgen_lat string `json:"BGEN_LAT"`
    Bgen_lon string `json:"BGEN_LON"`
    Bgen_dland string `json:"BGEN_DLAND"`
    Gen_dist string `json:"GEN_DIST"`
    Gen_tdiff string `json:"GEN_TDIFF"`
    Init_tdiff string `json:"INIT_TDIFF"`
    Dev_cat string `json:"DEV_CAT"`
    Ops_cat string `json:"OPS_CAT"`
}


type STAT_SEEPS_V12_0 struct {
    vxMetadata
    STAT_header_V12_0
    Total int `json:"TOTAL"`
    Odfl float64 `json:"ODFL"`
    Odfh float64 `json:"ODFH"`
    Olfd float64 `json:"OLFD"`
    Olfh float64 `json:"OLFH"`
    Ohfd float64 `json:"OHFD"`
    Ohfl float64 `json:"OHFL"`
    Pf1 float64 `json:"PF1"`
    Pf2 float64 `json:"PF2"`
    Pf3 float64 `json:"PF3"`
    Pv1 float64 `json:"PV1"`
    Pv2 float64 `json:"PV2"`
    Pv3 float64 `json:"PV3"`
    Mean_fcst float64 `json:"MEAN_FCST"`
    Mean_obs float64 `json:"MEAN_OBS"`
    Seeps float64 `json:"SEEPS"`
}


type STAT_GRAD_V12_0 struct {
    vxMetadata
    STAT_header_V12_0
    Total int `json:"TOTAL"`
    Fgbar float64 `json:"FGBAR"`
    Ogbar float64 `json:"OGBAR"`
    Mgbar float64 `json:"MGBAR"`
    Egbar float64 `json:"EGBAR"`
    S1 string `json:"S1"`
    S1_og string `json:"S1_OG"`
    Fgog_ratio string `json:"FGOG_RATIO"`
    Dx int `json:"DX"`
    Dy int `json:"DY"`
}


type STAT_PCT_V12_0 struct {
    vxMetadata
    STAT_header_V12_0
    Total int `json:"TOTAL"`
    (n_thresh) int `json:"(N_THRESH)"`
    Thresh_[0-9]* string `json:"THRESH_[0-9]*"`
    Oy_[0-9]* string `json:"OY_[0-9]*"`
    On_[0-9]* string `json:"ON_[0-9]*"`
}


type STAT_PSTD_V12_0 struct {
    vxMetadata
    STAT_header_V12_0
    Total int `json:"TOTAL"`
    (n_thresh) int `json:"(N_THRESH)"`
    Baser float64 `json:"BASER"`
    Baser_ncl string `json:"BASER_NCL"`
    Baser_ncu string `json:"BASER_NCU"`
    Reliability string `json:"RELIABILITY"`
    Resolution string `json:"RESOLUTION"`
    Uncertainty string `json:"UNCERTAINTY"`
    Roc_auc string `json:"ROC_AUC"`
    Brier string `json:"BRIER"`
    Brier_ncl string `json:"BRIER_NCL"`
    Brier_ncu string `json:"BRIER_NCU"`
    Briercl string `json:"BRIERCL"`
    Briercl_ncl string `json:"BRIERCL_NCL"`
    Briercl_ncu string `json:"BRIERCL_NCU"`
    Bss string `json:"BSS"`
    Bss_smpl string `json:"BSS_SMPL"`
    Thresh_[0-9]* string `json:"THRESH_[0-9]*"`
}


type TCST_PROBRIRW_V12_0 struct {
    vxMetadata
    TCST_header_V12_0
    Alat string `json:"ALAT"`
    Alon string `json:"ALON"`
    Blat string `json:"BLAT"`
    Blon string `json:"BLON"`
    Initials string `json:"INITIALS"`
    Tk_err string `json:"TK_ERR"`
    X_err string `json:"X_ERR"`
    Y_err string `json:"Y_ERR"`
    Adland string `json:"ADLAND"`
    Bdland string `json:"BDLAND"`
    Rirw_beg string `json:"RIRW_BEG"`
    Rirw_end string `json:"RIRW_END"`
    Rirw_window string `json:"RIRW_WINDOW"`
    Awind_end string `json:"AWIND_END"`
    Bwind_beg string `json:"BWIND_BEG"`
    Bwind_end string `json:"BWIND_END"`
    Bdelta string `json:"BDELTA"`
    Bdelta_max string `json:"BDELTA_MAX"`
    Blevel_beg string `json:"BLEVEL_BEG"`
    Blevel_end string `json:"BLEVEL_END"`
    (n_thresh) int `json:"(N_THRESH)"`
    Thresh_[0-9]* string `json:"THRESH_[0-9]*"`
    Prob_[0-9]* string `json:"PROB_[0-9]*"`
}


type STAT_NBRCTC_V12_0 struct {
    vxMetadata
    STAT_header_V12_0
    Total int `json:"TOTAL"`
    Fy_oy float64 `json:"FY_OY"`
    Fy_on float64 `json:"FY_ON"`
    Fn_oy float64 `json:"FN_OY"`
    Fn_on float64 `json:"FN_ON"`
}


type STAT_SEEPS_MPR_V12_0 struct {
    vxMetadata
    STAT_header_V12_0
    Obs_sid string `json:"OBS_SID"`
    Obs_lat float64 `json:"OBS_LAT"`
    Obs_lon float64 `json:"OBS_LON"`
    Fcst float64 `json:"FCST"`
    Obs float64 `json:"OBS"`
    Obs_qc string `json:"OBS_QC"`
    Fcst_cat int `json:"FCST_CAT"`
    Obs_cat int `json:"OBS_CAT"`
    P1 float64 `json:"P1"`
    P2 float64 `json:"P2"`
    T1 float64 `json:"T1"`
    T2 float64 `json:"T2"`
    Seeps float64 `json:"SEEPS"`
}


type STAT_ECNT_V12_0 struct {
    vxMetadata
    STAT_header_V12_0
    Total int `json:"TOTAL"`
    N_ens float64 `json:"N_ENS"`
    Crps float64 `json:"CRPS"`
    Crpss float64 `json:"CRPSS"`
    Ign float64 `json:"IGN"`
    Me float64 `json:"ME"`
    Rmse float64 `json:"RMSE"`
    Spread float64 `json:"SPREAD"`
    Me_oerr float64 `json:"ME_OERR"`
    Rmse_oerr float64 `json:"RMSE_OERR"`
    Spread_oerr float64 `json:"SPREAD_OERR"`
    Spread_plus_oerr float64 `json:"SPREAD_PLUS_OERR"`
    Crpscl float64 `json:"CRPSCL"`
    Crps_emp float64 `json:"CRPS_EMP"`
    Crpscl_emp float64 `json:"CRPSCL_EMP"`
    Crpss_emp float64 `json:"CRPSS_EMP"`
    Crps_emp_fair float64 `json:"CRPS_EMP_FAIR"`
    Spread_md float64 `json:"SPREAD_MD"`
    Mae float64 `json:"MAE"`
    Mae_oerr float64 `json:"MAE_OERR"`
    Bias_ratio float64 `json:"BIAS_RATIO"`
    N_ge_obs int `json:"N_GE_OBS"`
    Me_ge_obs float64 `json:"ME_GE_OBS"`
    N_lt_obs int `json:"N_LT_OBS"`
    Me_lt_obs float64 `json:"ME_LT_OBS"`
    Ign_conv_oerr float64 `json:"IGN_CONV_OERR"`
    Ign_corr_oerr float64 `json:"IGN_CORR_OERR"`
}


type MODE_CTS_V12_0 struct {
    vxMetadata
    MODE_header_V12_0
    Field string `json:"FIELD"`
    Total int `json:"TOTAL"`
    Fy_oy float64 `json:"FY_OY"`
    Fy_on float64 `json:"FY_ON"`
    Fn_oy float64 `json:"FN_OY"`
    Fn_on float64 `json:"FN_ON"`
    Baser float64 `json:"BASER"`
    Fmean string `json:"FMEAN"`
    Acc string `json:"ACC"`
    Fbias float64 `json:"FBIAS"`
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
}


type STAT_CTS_V12_0 struct {
    vxMetadata
    STAT_header_V12_0
    Total int `json:"TOTAL"`
    Baser float64 `json:"BASER"`
    Baser_ncl string `json:"BASER_NCL"`
    Baser_ncu string `json:"BASER_NCU"`
    Baser_bcl string `json:"BASER_BCL"`
    Baser_bcu string `json:"BASER_BCU"`
    Fmean string `json:"FMEAN"`
    Fmean_ncl string `json:"FMEAN_NCL"`
    Fmean_ncu string `json:"FMEAN_NCU"`
    Fmean_bcl string `json:"FMEAN_BCL"`
    Fmean_bcu string `json:"FMEAN_BCU"`
    Acc string `json:"ACC"`
    Acc_ncl string `json:"ACC_NCL"`
    Acc_ncu string `json:"ACC_NCU"`
    Acc_bcl string `json:"ACC_BCL"`
    Acc_bcu string `json:"ACC_BCU"`
    Fbias float64 `json:"FBIAS"`
    Fbias_bcl string `json:"FBIAS_BCL"`
    Fbias_bcu string `json:"FBIAS_BCU"`
    Pody string `json:"PODY"`
    Pody_ncl string `json:"PODY_NCL"`
    Pody_ncu string `json:"PODY_NCU"`
    Pody_bcl string `json:"PODY_BCL"`
    Pody_bcu string `json:"PODY_BCU"`
    Podn string `json:"PODN"`
    Podn_ncl string `json:"PODN_NCL"`
    Podn_ncu string `json:"PODN_NCU"`
    Podn_bcl string `json:"PODN_BCL"`
    Podn_bcu string `json:"PODN_BCU"`
    Pofd string `json:"POFD"`
    Pofd_ncl string `json:"POFD_NCL"`
    Pofd_ncu string `json:"POFD_NCU"`
    Pofd_bcl string `json:"POFD_BCL"`
    Pofd_bcu string `json:"POFD_BCU"`
    Far string `json:"FAR"`
    Far_ncl string `json:"FAR_NCL"`
    Far_ncu string `json:"FAR_NCU"`
    Far_bcl string `json:"FAR_BCL"`
    Far_bcu string `json:"FAR_BCU"`
    Csi string `json:"CSI"`
    Csi_ncl string `json:"CSI_NCL"`
    Csi_ncu string `json:"CSI_NCU"`
    Csi_bcl string `json:"CSI_BCL"`
    Csi_bcu string `json:"CSI_BCU"`
    Gss string `json:"GSS"`
    Gss_bcl string `json:"GSS_BCL"`
    Gss_bcu string `json:"GSS_BCU"`
    Hk string `json:"HK"`
    Hk_ncl string `json:"HK_NCL"`
    Hk_ncu string `json:"HK_NCU"`
    Hk_bcl string `json:"HK_BCL"`
    Hk_bcu string `json:"HK_BCU"`
    Hss string `json:"HSS"`
    Hss_bcl string `json:"HSS_BCL"`
    Hss_bcu string `json:"HSS_BCU"`
    Odds string `json:"ODDS"`
    Odds_ncl string `json:"ODDS_NCL"`
    Odds_ncu string `json:"ODDS_NCU"`
    Odds_bcl string `json:"ODDS_BCL"`
    Odds_bcu string `json:"ODDS_BCU"`
    Lodds string `json:"LODDS"`
    Lodds_ncl string `json:"LODDS_NCL"`
    Lodds_ncu string `json:"LODDS_NCU"`
    Lodds_bcl string `json:"LODDS_BCL"`
    Lodds_bcu string `json:"LODDS_BCU"`
    Orss string `json:"ORSS"`
    Orss_ncl string `json:"ORSS_NCL"`
    Orss_ncu string `json:"ORSS_NCU"`
    Orss_bcl string `json:"ORSS_BCL"`
    Orss_bcu string `json:"ORSS_BCU"`
    Eds string `json:"EDS"`
    Eds_ncl string `json:"EDS_NCL"`
    Eds_ncu string `json:"EDS_NCU"`
    Eds_bcl string `json:"EDS_BCL"`
    Eds_bcu string `json:"EDS_BCU"`
    Seds string `json:"SEDS"`
    Seds_ncl string `json:"SEDS_NCL"`
    Seds_ncu string `json:"SEDS_NCU"`
    Seds_bcl string `json:"SEDS_BCL"`
    Seds_bcu string `json:"SEDS_BCU"`
    Edi string `json:"EDI"`
    Edi_ncl string `json:"EDI_NCL"`
    Edi_ncu string `json:"EDI_NCU"`
    Edi_bcl string `json:"EDI_BCL"`
    Edi_bcu string `json:"EDI_BCU"`
    Sedi string `json:"SEDI"`
    Sedi_ncl string `json:"SEDI_NCL"`
    Sedi_ncu string `json:"SEDI_NCU"`
    Sedi_bcl string `json:"SEDI_BCL"`
    Sedi_bcu string `json:"SEDI_BCU"`
    Bagss string `json:"BAGSS"`
    Bagss_bcl string `json:"BAGSS_BCL"`
    Bagss_bcu string `json:"BAGSS_BCU"`
    Hss_ec string `json:"HSS_EC"`
    Hss_ec_bcl string `json:"HSS_EC_BCL"`
    Hss_ec_bcu string `json:"HSS_EC_BCU"`
    Ec_value float64 `json:"EC_VALUE"`
}


type STAT_MPR_V12_0 struct {
    vxMetadata
    STAT_header_V12_0
    Total int `json:"TOTAL"`
    Index int `json:"INDEX"`
    Obs_sid string `json:"OBS_SID"`
    Obs_lat float64 `json:"OBS_LAT"`
    Obs_lon float64 `json:"OBS_LON"`
    Obs_lvl float64 `json:"OBS_LVL"`
    Obs_elv float64 `json:"OBS_ELV"`
    Fcst float64 `json:"FCST"`
    Obs float64 `json:"OBS"`
    Obs_qc string `json:"OBS_QC"`
    Obs_climo_mean float64 `json:"OBS_CLIMO_MEAN"`
    Obs_climo_stdev float64 `json:"OBS_CLIMO_STDEV"`
    Obs_climo_cdf float64 `json:"OBS_CLIMO_CDF"`
    Fcst_climo_mean float64 `json:"FCST_CLIMO_MEAN"`
    Fcst_climo_stdev float64 `json:"FCST_CLIMO_STDEV"`
}


type STAT_SL1L2_V12_0 struct {
    vxMetadata
    STAT_header_V12_0
    Total int `json:"TOTAL"`
    Fbar float64 `json:"FBAR"`
    Obar float64 `json:"OBAR"`
    Fobar float64 `json:"FOBAR"`
    Ffbar float64 `json:"FFBAR"`
    Oobar float64 `json:"OOBAR"`
    Mae float64 `json:"MAE"`
}


type STAT_SSIDX_V12_0 struct {
    vxMetadata
    STAT_header_V12_0
    Fcst_model string `json:"FCST_MODEL"`
    Ref_model string `json:"REF_MODEL"`
    N_init string `json:"N_INIT"`
    N_term string `json:"N_TERM"`
    N_vld string `json:"N_VLD"`
    Ss_index string `json:"SS_INDEX"`
}


type STAT_CNT_V12_0 struct {
    vxMetadata
    STAT_header_V12_0
    Total int `json:"TOTAL"`
    Fbar float64 `json:"FBAR"`
    Fbar_ncl string `json:"FBAR_NCL"`
    Fbar_ncu string `json:"FBAR_NCU"`
    Fbar_bcl string `json:"FBAR_BCL"`
    Fbar_bcu string `json:"FBAR_BCU"`
    Fstdev string `json:"FSTDEV"`
    Fstdev_ncl string `json:"FSTDEV_NCL"`
    Fstdev_ncu string `json:"FSTDEV_NCU"`
    Fstdev_bcl string `json:"FSTDEV_BCL"`
    Fstdev_bcu string `json:"FSTDEV_BCU"`
    Obar float64 `json:"OBAR"`
    Obar_ncl string `json:"OBAR_NCL"`
    Obar_ncu string `json:"OBAR_NCU"`
    Obar_bcl string `json:"OBAR_BCL"`
    Obar_bcu string `json:"OBAR_BCU"`
    Ostdev string `json:"OSTDEV"`
    Ostdev_ncl string `json:"OSTDEV_NCL"`
    Ostdev_ncu string `json:"OSTDEV_NCU"`
    Ostdev_bcl string `json:"OSTDEV_BCL"`
    Ostdev_bcu string `json:"OSTDEV_BCU"`
    Pr_corr string `json:"PR_CORR"`
    Pr_corr_ncl string `json:"PR_CORR_NCL"`
    Pr_corr_ncu string `json:"PR_CORR_NCU"`
    Pr_corr_bcl string `json:"PR_CORR_BCL"`
    Pr_corr_bcu string `json:"PR_CORR_BCU"`
    Sp_corr string `json:"SP_CORR"`
    Kt_corr string `json:"KT_CORR"`
    Ranks string `json:"RANKS"`
    Frank_ties string `json:"FRANK_TIES"`
    Orank_ties string `json:"ORANK_TIES"`
    Me float64 `json:"ME"`
    Me_ncl string `json:"ME_NCL"`
    Me_ncu string `json:"ME_NCU"`
    Me_bcl string `json:"ME_BCL"`
    Me_bcu string `json:"ME_BCU"`
    Estdev string `json:"ESTDEV"`
    Estdev_ncl string `json:"ESTDEV_NCL"`
    Estdev_ncu string `json:"ESTDEV_NCU"`
    Estdev_bcl string `json:"ESTDEV_BCL"`
    Estdev_bcu string `json:"ESTDEV_BCU"`
    Mbias string `json:"MBIAS"`
    Mbias_bcl string `json:"MBIAS_BCL"`
    Mbias_bcu string `json:"MBIAS_BCU"`
    Mae float64 `json:"MAE"`
    Mae_bcl string `json:"MAE_BCL"`
    Mae_bcu string `json:"MAE_BCU"`
    Mse float64 `json:"MSE"`
    Mse_bcl string `json:"MSE_BCL"`
    Mse_bcu string `json:"MSE_BCU"`
    Bcmse string `json:"BCMSE"`
    Bcmse_bcl string `json:"BCMSE_BCL"`
    Bcmse_bcu string `json:"BCMSE_BCU"`
    Rmse float64 `json:"RMSE"`
    Rmse_bcl string `json:"RMSE_BCL"`
    Rmse_bcu string `json:"RMSE_BCU"`
    E10 string `json:"E10"`
    E10_bcl string `json:"E10_BCL"`
    E10_bcu string `json:"E10_BCU"`
    E25 string `json:"E25"`
    E25_bcl string `json:"E25_BCL"`
    E25_bcu string `json:"E25_BCU"`
    E50 string `json:"E50"`
    E50_bcl string `json:"E50_BCL"`
    E50_bcu string `json:"E50_BCU"`
    E75 string `json:"E75"`
    E75_bcl string `json:"E75_BCL"`
    E75_bcu string `json:"E75_BCU"`
    E90 string `json:"E90"`
    E90_bcl string `json:"E90_BCL"`
    E90_bcu string `json:"E90_BCU"`
    Eiqr string `json:"EIQR"`
    Eiqr_bcl string `json:"EIQR_BCL"`
    Eiqr_bcu string `json:"EIQR_BCU"`
    Mad string `json:"MAD"`
    Mad_bcl string `json:"MAD_BCL"`
    Mad_bcu string `json:"MAD_BCU"`
    Anom_corr string `json:"ANOM_CORR"`
    Anom_corr_ncl string `json:"ANOM_CORR_NCL"`
    Anom_corr_ncu string `json:"ANOM_CORR_NCU"`
    Anom_corr_bcl string `json:"ANOM_CORR_BCL"`
    Anom_corr_bcu string `json:"ANOM_CORR_BCU"`
    Me2 string `json:"ME2"`
    Me2_bcl string `json:"ME2_BCL"`
    Me2_bcu string `json:"ME2_BCU"`
    Msess string `json:"MSESS"`
    Msess_bcl string `json:"MSESS_BCL"`
    Msess_bcu string `json:"MSESS_BCU"`
    Rmsfa string `json:"RMSFA"`
    Rmsfa_bcl string `json:"RMSFA_BCL"`
    Rmsfa_bcu string `json:"RMSFA_BCU"`
    Rmsoa string `json:"RMSOA"`
    Rmsoa_bcl string `json:"RMSOA_BCL"`
    Rmsoa_bcu string `json:"RMSOA_BCU"`
    Anom_corr_uncntr string `json:"ANOM_CORR_UNCNTR"`
    Anom_corr_uncntr_bcl string `json:"ANOM_CORR_UNCNTR_BCL"`
    Anom_corr_uncntr_bcu string `json:"ANOM_CORR_UNCNTR_BCU"`
    Si float64 `json:"SI"`
    Si_bcl string `json:"SI_BCL"`
    Si_bcu string `json:"SI_BCU"`
}


type STAT_MCTC_V12_0 struct {
    vxMetadata
    STAT_header_V12_0
    Total int `json:"TOTAL"`
    (n_cat) int `json:"(N_CAT)"`
    F[0-9]*_o[0-9]* string `json:"F[0-9]*_O[0-9]*"`
    Ec_value float64 `json:"EC_VALUE"`
}


type STAT_NBRCTS_V12_0 struct {
    vxMetadata
    STAT_header_V12_0
    Total int `json:"TOTAL"`
    Baser float64 `json:"BASER"`
    Baser_ncl string `json:"BASER_NCL"`
    Baser_ncu string `json:"BASER_NCU"`
    Baser_bcl string `json:"BASER_BCL"`
    Baser_bcu string `json:"BASER_BCU"`
    Fmean string `json:"FMEAN"`
    Fmean_ncl string `json:"FMEAN_NCL"`
    Fmean_ncu string `json:"FMEAN_NCU"`
    Fmean_bcl string `json:"FMEAN_BCL"`
    Fmean_bcu string `json:"FMEAN_BCU"`
    Acc string `json:"ACC"`
    Acc_ncl string `json:"ACC_NCL"`
    Acc_ncu string `json:"ACC_NCU"`
    Acc_bcl string `json:"ACC_BCL"`
    Acc_bcu string `json:"ACC_BCU"`
    Fbias float64 `json:"FBIAS"`
    Fbias_bcl string `json:"FBIAS_BCL"`
    Fbias_bcu string `json:"FBIAS_BCU"`
    Pody string `json:"PODY"`
    Pody_ncl string `json:"PODY_NCL"`
    Pody_ncu string `json:"PODY_NCU"`
    Pody_bcl string `json:"PODY_BCL"`
    Pody_bcu string `json:"PODY_BCU"`
    Podn string `json:"PODN"`
    Podn_ncl string `json:"PODN_NCL"`
    Podn_ncu string `json:"PODN_NCU"`
    Podn_bcl string `json:"PODN_BCL"`
    Podn_bcu string `json:"PODN_BCU"`
    Pofd string `json:"POFD"`
    Pofd_ncl string `json:"POFD_NCL"`
    Pofd_ncu string `json:"POFD_NCU"`
    Pofd_bcl string `json:"POFD_BCL"`
    Pofd_bcu string `json:"POFD_BCU"`
    Far string `json:"FAR"`
    Far_ncl string `json:"FAR_NCL"`
    Far_ncu string `json:"FAR_NCU"`
    Far_bcl string `json:"FAR_BCL"`
    Far_bcu string `json:"FAR_BCU"`
    Csi string `json:"CSI"`
    Csi_ncl string `json:"CSI_NCL"`
    Csi_ncu string `json:"CSI_NCU"`
    Csi_bcl string `json:"CSI_BCL"`
    Csi_bcu string `json:"CSI_BCU"`
    Gss string `json:"GSS"`
    Gss_bcl string `json:"GSS_BCL"`
    Gss_bcu string `json:"GSS_BCU"`
    Hk string `json:"HK"`
    Hk_ncl string `json:"HK_NCL"`
    Hk_ncu string `json:"HK_NCU"`
    Hk_bcl string `json:"HK_BCL"`
    Hk_bcu string `json:"HK_BCU"`
    Hss string `json:"HSS"`
    Hss_bcl string `json:"HSS_BCL"`
    Hss_bcu string `json:"HSS_BCU"`
    Odds string `json:"ODDS"`
    Odds_ncl string `json:"ODDS_NCL"`
    Odds_ncu string `json:"ODDS_NCU"`
    Odds_bcl string `json:"ODDS_BCL"`
    Odds_bcu string `json:"ODDS_BCU"`
    Lodds string `json:"LODDS"`
    Lodds_ncl string `json:"LODDS_NCL"`
    Lodds_ncu string `json:"LODDS_NCU"`
    Lodds_bcl string `json:"LODDS_BCL"`
    Lodds_bcu string `json:"LODDS_BCU"`
    Orss string `json:"ORSS"`
    Orss_ncl string `json:"ORSS_NCL"`
    Orss_ncu string `json:"ORSS_NCU"`
    Orss_bcl string `json:"ORSS_BCL"`
    Orss_bcu string `json:"ORSS_BCU"`
    Eds string `json:"EDS"`
    Eds_ncl string `json:"EDS_NCL"`
    Eds_ncu string `json:"EDS_NCU"`
    Eds_bcl string `json:"EDS_BCL"`
    Eds_bcu string `json:"EDS_BCU"`
    Seds string `json:"SEDS"`
    Seds_ncl string `json:"SEDS_NCL"`
    Seds_ncu string `json:"SEDS_NCU"`
    Seds_bcl string `json:"SEDS_BCL"`
    Seds_bcu string `json:"SEDS_BCU"`
    Edi string `json:"EDI"`
    Edi_ncl string `json:"EDI_NCL"`
    Edi_ncu string `json:"EDI_NCU"`
    Edi_bcl string `json:"EDI_BCL"`
    Edi_bcu string `json:"EDI_BCU"`
    Sedi string `json:"SEDI"`
    Sedi_ncl string `json:"SEDI_NCL"`
    Sedi_ncu string `json:"SEDI_NCU"`
    Sedi_bcl string `json:"SEDI_BCL"`
    Sedi_bcu string `json:"SEDI_BCU"`
    Bagss string `json:"BAGSS"`
    Bagss_bcl string `json:"BAGSS_BCL"`
    Bagss_bcu string `json:"BAGSS_BCU"`
}


type STAT_DMAP_V12_0 struct {
    vxMetadata
    STAT_header_V12_0
    Total int `json:"TOTAL"`
    Fy float64 `json:"FY"`
    Oy float64 `json:"OY"`
    Fbias float64 `json:"FBIAS"`
    Baddeley string `json:"BADDELEY"`
    Hausdorff string `json:"HAUSDORFF"`
    Med_fo string `json:"MED_FO"`
    Med_of string `json:"MED_OF"`
    Med_min string `json:"MED_MIN"`
    Med_max string `json:"MED_MAX"`
    Med_mean string `json:"MED_MEAN"`
    Fom_fo string `json:"FOM_FO"`
    Fom_of string `json:"FOM_OF"`
    Fom_min string `json:"FOM_MIN"`
    Fom_max string `json:"FOM_MAX"`
    Fom_mean string `json:"FOM_MEAN"`
    Zhu_fo string `json:"ZHU_FO"`
    Zhu_of string `json:"ZHU_OF"`
    Zhu_min string `json:"ZHU_MIN"`
    Zhu_max string `json:"ZHU_MAX"`
    Zhu_mean string `json:"ZHU_MEAN"`
    G float64 `json:"G"`
    Gbeta string `json:"GBETA"`
    Beta_value string `json:"BETA_VALUE"`
}


type STAT_VL1L2_V12_0 struct {
    vxMetadata
    STAT_header_V12_0
    Total int `json:"TOTAL"`
    Ufbar float64 `json:"UFBAR"`
    Vfbar float64 `json:"VFBAR"`
    Uobar float64 `json:"UOBAR"`
    Vobar float64 `json:"VOBAR"`
    Uvfobar float64 `json:"UVFOBAR"`
    Uvffbar float64 `json:"UVFFBAR"`
    Uvoobar float64 `json:"UVOOBAR"`
    F_speed_bar float64 `json:"F_SPEED_BAR"`
    O_speed_bar float64 `json:"O_SPEED_BAR"`
    Total_dir int `json:"TOTAL_DIR"`
    Dir_me float64 `json:"DIR_ME"`
    Dir_mae float64 `json:"DIR_MAE"`
    Dir_mse float64 `json:"DIR_MSE"`
}


type STAT_FHO_V12_0 struct {
    vxMetadata
    STAT_header_V12_0
    Total int `json:"TOTAL"`
    F_rate float64 `json:"F_RATE"`
    H_rate float64 `json:"H_RATE"`
    O_rate float64 `json:"O_RATE"`
}


type STAT_NBRCNT_V12_0 struct {
    vxMetadata
    STAT_header_V12_0
    Total int `json:"TOTAL"`
    Fbs float64 `json:"FBS"`
    Fbs_bcl string `json:"FBS_BCL"`
    Fbs_bcu string `json:"FBS_BCU"`
    Fss float64 `json:"FSS"`
    Fss_bcl string `json:"FSS_BCL"`
    Fss_bcu string `json:"FSS_BCU"`
    Afss float64 `json:"AFSS"`
    Afss_bcl string `json:"AFSS_BCL"`
    Afss_bcu string `json:"AFSS_BCU"`
    Ufss float64 `json:"UFSS"`
    Ufss_bcl string `json:"UFSS_BCL"`
    Ufss_bcu string `json:"UFSS_BCU"`
    F_rate float64 `json:"F_RATE"`
    F_rate_bcl string `json:"F_RATE_BCL"`
    F_rate_bcu string `json:"F_RATE_BCU"`
    O_rate float64 `json:"O_RATE"`
    O_rate_bcl string `json:"O_RATE_BCL"`
    O_rate_bcu string `json:"O_RATE_BCU"`
}


type STAT_PJC_V12_0 struct {
    vxMetadata
    STAT_header_V12_0
    Total int `json:"TOTAL"`
    (n_thresh) int `json:"(N_THRESH)"`
    Thresh_[0-9]* string `json:"THRESH_[0-9]*"`
    Oy_tp_[0-9]* string `json:"OY_TP_[0-9]*"`
    On_tp_[0-9]* string `json:"ON_TP_[0-9]*"`
    Calibration_[0-9]* string `json:"CALIBRATION_[0-9]*"`
    Refinement_[0-9]* string `json:"REFINEMENT_[0-9]*"`
    Likelihood_[0-9]* string `json:"LIKELIHOOD_[0-9]*"`
    Baser_[0-9]* string `json:"BASER_[0-9]*"`
}


type STAT_RPS_V12_0 struct {
    vxMetadata
    STAT_header_V12_0
    Total int `json:"TOTAL"`
    N_prob float64 `json:"N_PROB"`
    Rps_rel float64 `json:"RPS_REL"`
    Rps_res float64 `json:"RPS_RES"`
    Rps_unc float64 `json:"RPS_UNC"`
    Rps float64 `json:"RPS"`
    Rpss float64 `json:"RPSS"`
    Rpss_smpl float64 `json:"RPSS_SMPL"`
    Rps_comp string `json:"RPS_COMP"`
}


type STAT_SAL1L2_V12_0 struct {
    vxMetadata
    STAT_header_V12_0
    Total int `json:"TOTAL"`
    Fabar float64 `json:"FABAR"`
    Oabar float64 `json:"OABAR"`
    Foabar float64 `json:"FOABAR"`
    Ffabar float64 `json:"FFABAR"`
    Ooabar float64 `json:"OOABAR"`
    Mae float64 `json:"MAE"`
}


type TCST_TCMPR_V12_0 struct {
    vxMetadata
    TCST_header_V12_0
    Total int `json:"TOTAL"`
    Index int `json:"INDEX"`
    Level string `json:"LEVEL"`
    Watch_warn string `json:"WATCH_WARN"`
    Initials string `json:"INITIALS"`
    Alat string `json:"ALAT"`
    Alon string `json:"ALON"`
    Blat string `json:"BLAT"`
    Blon string `json:"BLON"`
    Tk_err string `json:"TK_ERR"`
    X_err string `json:"X_ERR"`
    Y_err string `json:"Y_ERR"`
    Altk_err string `json:"ALTK_ERR"`
    Crtk_err string `json:"CRTK_ERR"`
    Adland string `json:"ADLAND"`
    Bdland string `json:"BDLAND"`
    Amslp string `json:"AMSLP"`
    Bmslp string `json:"BMSLP"`
    Amax_wind string `json:"AMAX_WIND"`
    Bmax_wind string `json:"BMAX_WIND"`
    Aal_wind_34 string `json:"AAL_WIND_34"`
    Bal_wind_34 string `json:"BAL_WIND_34"`
    Ane_wind_34 string `json:"ANE_WIND_34"`
    Bne_wind_34 string `json:"BNE_WIND_34"`
    Ase_wind_34 string `json:"ASE_WIND_34"`
    Bse_wind_34 string `json:"BSE_WIND_34"`
    Asw_wind_34 string `json:"ASW_WIND_34"`
    Bsw_wind_34 string `json:"BSW_WIND_34"`
    Anw_wind_34 string `json:"ANW_WIND_34"`
    Bnw_wind_34 string `json:"BNW_WIND_34"`
    Aal_wind_50 string `json:"AAL_WIND_50"`
    Bal_wind_50 string `json:"BAL_WIND_50"`
    Ane_wind_50 string `json:"ANE_WIND_50"`
    Bne_wind_50 string `json:"BNE_WIND_50"`
    Ase_wind_50 string `json:"ASE_WIND_50"`
    Bse_wind_50 string `json:"BSE_WIND_50"`
    Asw_wind_50 string `json:"ASW_WIND_50"`
    Bsw_wind_50 string `json:"BSW_WIND_50"`
    Anw_wind_50 string `json:"ANW_WIND_50"`
    Bnw_wind_50 string `json:"BNW_WIND_50"`
    Aal_wind_64 string `json:"AAL_WIND_64"`
    Bal_wind_64 string `json:"BAL_WIND_64"`
    Ane_wind_64 string `json:"ANE_WIND_64"`
    Bne_wind_64 string `json:"BNE_WIND_64"`
    Ase_wind_64 string `json:"ASE_WIND_64"`
    Bse_wind_64 string `json:"BSE_WIND_64"`
    Asw_wind_64 string `json:"ASW_WIND_64"`
    Bsw_wind_64 string `json:"BSW_WIND_64"`
    Anw_wind_64 string `json:"ANW_WIND_64"`
    Bnw_wind_64 string `json:"BNW_WIND_64"`
    Aradp string `json:"ARADP"`
    Bradp string `json:"BRADP"`
    Arrp string `json:"ARRP"`
    Brrp string `json:"BRRP"`
    Amrd string `json:"AMRD"`
    Bmrd string `json:"BMRD"`
    Agusts string `json:"AGUSTS"`
    Bgusts string `json:"BGUSTS"`
    Aeye string `json:"AEYE"`
    Beye string `json:"BEYE"`
    Adir string `json:"ADIR"`
    Bdir string `json:"BDIR"`
    Aspeed string `json:"ASPEED"`
    Bspeed string `json:"BSPEED"`
    Adepth string `json:"ADEPTH"`
    Bdepth string `json:"BDEPTH"`
    Num_members string `json:"NUM_MEMBERS"`
    Track_spread string `json:"TRACK_SPREAD"`
    Track_stdev string `json:"TRACK_STDEV"`
    Mslp_stdev string `json:"MSLP_STDEV"`
    Max_wind_stdev string `json:"MAX_WIND_STDEV"`
}


type STAT_ISC_V12_0 struct {
    vxMetadata
    STAT_header_V12_0
    Total int `json:"TOTAL"`
    Tile_dim int `json:"TILE_DIM"`
    Tile_xll int `json:"TILE_XLL"`
    Tile_yll int `json:"TILE_YLL"`
    Nscale int `json:"NSCALE"`
    Iscale int `json:"ISCALE"`
    Mse float64 `json:"MSE"`
    Isc int `json:"ISC"`
    Fenergy2 float64 `json:"FENERGY2"`
    Oenergy2 float64 `json:"OENERGY2"`
    Baser float64 `json:"BASER"`
    Fbias float64 `json:"FBIAS"`
}


type STAT_ORANK_V12_0 struct {
    vxMetadata
    STAT_header_V12_0
    Total int `json:"TOTAL"`
    Index int `json:"INDEX"`
    Obs_sid string `json:"OBS_SID"`
    Obs_lat float64 `json:"OBS_LAT"`
    Obs_lon float64 `json:"OBS_LON"`
    Obs_lvl float64 `json:"OBS_LVL"`
    Obs_elv float64 `json:"OBS_ELV"`
    Obs float64 `json:"OBS"`
    Pit float64 `json:"PIT"`
    Rank int `json:"RANK"`
    N_ens_vld int `json:"N_ENS_VLD"`
    (n_ens) float64 `json:"(N_ENS)"`
    Ens_[0-9]* int `json:"ENS_[0-9]*"`
    Obs_qc string `json:"OBS_QC"`
    Ens_mean float64 `json:"ENS_MEAN"`
    Obs_climo_mean float64 `json:"OBS_CLIMO_MEAN"`
    Spread float64 `json:"SPREAD"`
    Ens_mean_oerr float64 `json:"ENS_MEAN_OERR"`
    Spread_oerr float64 `json:"SPREAD_OERR"`
    Spread_plus_oerr float64 `json:"SPREAD_PLUS_OERR"`
    Obs_climo_stdev float64 `json:"OBS_CLIMO_STDEV"`
    Fcst_climo_mean float64 `json:"FCST_CLIMO_MEAN"`
    Fcst_climo_stdev float64 `json:"FCST_CLIMO_STDEV"`
}


type STAT_PRC_V12_0 struct {
    vxMetadata
    STAT_header_V12_0
    Total int `json:"TOTAL"`
    (n_thresh) int `json:"(N_THRESH)"`
    Thresh_[0-9]* string `json:"THRESH_[0-9]*"`
    Pody_[0-9]* string `json:"PODY_[0-9]*"`
    Pofd_[0-9]* string `json:"POFD_[0-9]*"`
}


type STAT_PHIST_V12_0 struct {
    vxMetadata
    STAT_header_V12_0
    Total int `json:"TOTAL"`
    Bin_size float64 `json:"BIN_SIZE"`
    (n_bin) int `json:"(N_BIN)"`
    Bin_[0-9]* float64 `json:"BIN_[0-9]*"`
}


type STAT_VAL1L2_V12_0 struct {
    vxMetadata
    STAT_header_V12_0
    Total int `json:"TOTAL"`
    Ufabar float64 `json:"UFABAR"`
    Vfabar float64 `json:"VFABAR"`
    Uoabar float64 `json:"UOABAR"`
    Voabar float64 `json:"VOABAR"`
    Uvfoabar float64 `json:"UVFOABAR"`
    Uvffabar float64 `json:"UVFFABAR"`
    Uvooabar float64 `json:"UVOOABAR"`
    Fa_speed_bar float64 `json:"FA_SPEED_BAR"`
    Oa_speed_bar float64 `json:"OA_SPEED_BAR"`
    Total_dir int `json:"TOTAL_DIR"`
    Dira_me float64 `json:"DIRA_ME"`
    Dira_mae float64 `json:"DIRA_MAE"`
    Dira_mse float64 `json:"DIRA_MSE"`
}


type STAT_VCNT_V12_0 struct {
    vxMetadata
    STAT_header_V12_0
    Total int `json:"TOTAL"`
    Fbar float64 `json:"FBAR"`
    Fbar_bcl string `json:"FBAR_BCL"`
    Fbar_bcu string `json:"FBAR_BCU"`
    Obar float64 `json:"OBAR"`
    Obar_bcl string `json:"OBAR_BCL"`
    Obar_bcu string `json:"OBAR_BCU"`
    Fs_rms string `json:"FS_RMS"`
    Fs_rms_bcl string `json:"FS_RMS_BCL"`
    Fs_rms_bcu string `json:"FS_RMS_BCU"`
    Os_rms string `json:"OS_RMS"`
    Os_rms_bcl string `json:"OS_RMS_BCL"`
    Os_rms_bcu string `json:"OS_RMS_BCU"`
    Msve string `json:"MSVE"`
    Msve_bcl string `json:"MSVE_BCL"`
    Msve_bcu string `json:"MSVE_BCU"`
    Rmsve string `json:"RMSVE"`
    Rmsve_bcl string `json:"RMSVE_BCL"`
    Rmsve_bcu string `json:"RMSVE_BCU"`
    Fstdev string `json:"FSTDEV"`
    Fstdev_bcl string `json:"FSTDEV_BCL"`
    Fstdev_bcu string `json:"FSTDEV_BCU"`
    Ostdev string `json:"OSTDEV"`
    Ostdev_bcl string `json:"OSTDEV_BCL"`
    Ostdev_bcu string `json:"OSTDEV_BCU"`
    Fdir string `json:"FDIR"`
    Fdir_bcl string `json:"FDIR_BCL"`
    Fdir_bcu string `json:"FDIR_BCU"`
    Odir string `json:"ODIR"`
    Odir_bcl string `json:"ODIR_BCL"`
    Odir_bcu string `json:"ODIR_BCU"`
    Fbar_speed string `json:"FBAR_SPEED"`
    Fbar_speed_bcl string `json:"FBAR_SPEED_BCL"`
    Fbar_speed_bcu string `json:"FBAR_SPEED_BCU"`
    Obar_speed string `json:"OBAR_SPEED"`
    Obar_speed_bcl string `json:"OBAR_SPEED_BCL"`
    Obar_speed_bcu string `json:"OBAR_SPEED_BCU"`
    Vdiff_speed string `json:"VDIFF_SPEED"`
    Vdiff_speed_bcl string `json:"VDIFF_SPEED_BCL"`
    Vdiff_speed_bcu string `json:"VDIFF_SPEED_BCU"`
    Vdiff_dir string `json:"VDIFF_DIR"`
    Vdiff_dir_bcl string `json:"VDIFF_DIR_BCL"`
    Vdiff_dir_bcu string `json:"VDIFF_DIR_BCU"`
    Speed_err string `json:"SPEED_ERR"`
    Speed_err_bcl string `json:"SPEED_ERR_BCL"`
    Speed_err_bcu string `json:"SPEED_ERR_BCU"`
    Speed_abserr string `json:"SPEED_ABSERR"`
    Speed_abserr_bcl string `json:"SPEED_ABSERR_BCL"`
    Speed_abserr_bcu string `json:"SPEED_ABSERR_BCU"`
    Dir_err string `json:"DIR_ERR"`
    Dir_err_bcl string `json:"DIR_ERR_BCL"`
    Dir_err_bcu string `json:"DIR_ERR_BCU"`
    Dir_abserr string `json:"DIR_ABSERR"`
    Dir_abserr_bcl string `json:"DIR_ABSERR_BCL"`
    Dir_abserr_bcu string `json:"DIR_ABSERR_BCU"`
    Anom_corr string `json:"ANOM_CORR"`
    Anom_corr_ncl string `json:"ANOM_CORR_NCL"`
    Anom_corr_ncu string `json:"ANOM_CORR_NCU"`
    Anom_corr_bcl string `json:"ANOM_CORR_BCL"`
    Anom_corr_bcu string `json:"ANOM_CORR_BCU"`
    Anom_corr_uncntr string `json:"ANOM_CORR_UNCNTR"`
    Anom_corr_uncntr_bcl string `json:"ANOM_CORR_UNCNTR_BCL"`
    Anom_corr_uncntr_bcu string `json:"ANOM_CORR_UNCNTR_BCU"`
    Total_dir int `json:"TOTAL_DIR"`
    Dir_me float64 `json:"DIR_ME"`
    Dir_me_bcl string `json:"DIR_ME_BCL"`
    Dir_me_bcu string `json:"DIR_ME_BCU"`
    Dir_mae float64 `json:"DIR_MAE"`
    Dir_mae_bcl string `json:"DIR_MAE_BCL"`
    Dir_mae_bcu string `json:"DIR_MAE_BCU"`
    Dir_mse float64 `json:"DIR_MSE"`
    Dir_mse_bcl string `json:"DIR_MSE_BCL"`
    Dir_mse_bcu string `json:"DIR_MSE_BCU"`
    Dir_rmse string `json:"DIR_RMSE"`
    Dir_rmse_bcl string `json:"DIR_RMSE_BCL"`
    Dir_rmse_bcu string `json:"DIR_RMSE_BCU"`
}


type TCST_TCDIAG_V12_0 struct {
    vxMetadata
    TCST_header_V12_0
    Total int `json:"TOTAL"`
    Index int `json:"INDEX"`
    Diag_source string `json:"DIAG_SOURCE"`
    Track_source string `json:"TRACK_SOURCE"`
    Field_source string `json:"FIELD_SOURCE"`
    (n_diag) string `json:"(N_DIAG)"`
    Diag_[0-9]* string `json:"DIAG_[0-9]*"`
    Value_[0-9]* string `json:"VALUE_[0-9]*"`
}


type STAT_MCTS_V12_0 struct {
    vxMetadata
    STAT_header_V12_0
    Total int `json:"TOTAL"`
    N_cat int `json:"N_CAT"`
    Acc string `json:"ACC"`
    Acc_ncl string `json:"ACC_NCL"`
    Acc_ncu string `json:"ACC_NCU"`
    Acc_bcl string `json:"ACC_BCL"`
    Acc_bcu string `json:"ACC_BCU"`
    Hk string `json:"HK"`
    Hk_bcl string `json:"HK_BCL"`
    Hk_bcu string `json:"HK_BCU"`
    Hss string `json:"HSS"`
    Hss_bcl string `json:"HSS_BCL"`
    Hss_bcu string `json:"HSS_BCU"`
    Ger string `json:"GER"`
    Ger_bcl string `json:"GER_BCL"`
    Ger_bcu string `json:"GER_BCU"`
    Hss_ec string `json:"HSS_EC"`
    Hss_ec_bcl string `json:"HSS_EC_BCL"`
    Hss_ec_bcu string `json:"HSS_EC_BCU"`
    Ec_value float64 `json:"EC_VALUE"`
}


type STAT_RHIST_V12_0 struct {
    vxMetadata
    STAT_header_V12_0
    Total int `json:"TOTAL"`
    (n_rank) int `json:"(N_RANK)"`
    Rank_[0-9]* string `json:"RANK_[0-9]*"`
}


type STAT_RELP_V12_0 struct {
    vxMetadata
    STAT_header_V12_0
    Total int `json:"TOTAL"`
    (n_ens) float64 `json:"(N_ENS)"`
    Relp_[0-9]* string `json:"RELP_[0-9]*"`
}


type STAT_SSVAR_V12_0 struct {
    vxMetadata
    STAT_header_V12_0
    Total int `json:"TOTAL"`
    N_bin int `json:"N_BIN"`
    Bin_i string `json:"BIN_i"`
    Bin_n int `json:"BIN_N"`
    Var_min float64 `json:"VAR_MIN"`
    Var_max float64 `json:"VAR_MAX"`
    Var_mean float64 `json:"VAR_MEAN"`
    Fbar float64 `json:"FBAR"`
    Obar float64 `json:"OBAR"`
    Fobar float64 `json:"FOBAR"`
    Ffbar float64 `json:"FFBAR"`
    Oobar float64 `json:"OOBAR"`
    Fbar_ncl string `json:"FBAR_NCL"`
    Fbar_ncu string `json:"FBAR_NCU"`
    Fstdev string `json:"FSTDEV"`
    Fstdev_ncl string `json:"FSTDEV_NCL"`
    Fstdev_ncu string `json:"FSTDEV_NCU"`
    Obar_ncl string `json:"OBAR_NCL"`
    Obar_ncu string `json:"OBAR_NCU"`
    Ostdev string `json:"OSTDEV"`
    Ostdev_ncl string `json:"OSTDEV_NCL"`
    Ostdev_ncu string `json:"OSTDEV_NCU"`
    Pr_corr string `json:"PR_CORR"`
    Pr_corr_ncl string `json:"PR_CORR_NCL"`
    Pr_corr_ncu string `json:"PR_CORR_NCU"`
    Me float64 `json:"ME"`
    Me_ncl string `json:"ME_NCL"`
    Me_ncu string `json:"ME_NCU"`
    Estdev string `json:"ESTDEV"`
    Estdev_ncl string `json:"ESTDEV_NCL"`
    Estdev_ncu string `json:"ESTDEV_NCU"`
    Mbias string `json:"MBIAS"`
    Mse float64 `json:"MSE"`
    Bcmse string `json:"BCMSE"`
    Rmse float64 `json:"RMSE"`
}


type MODE_OBJ_V12_0 struct {
    vxMetadata
    MODE_header_V12_0
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
}


type STAT_CTC_V12_0 struct {
    vxMetadata
    STAT_header_V12_0
    Total int `json:"TOTAL"`
    Fy_oy float64 `json:"FY_OY"`
    Fy_on float64 `json:"FY_ON"`
    Fn_oy float64 `json:"FN_OY"`
    Fn_on float64 `json:"FN_ON"`
    Ec_value float64 `json:"EC_VALUE"`
}


var parserMap = map[string]ColumnDef{
    "STAT_SAL1L2_V12_0": {
		Name: "STAT_SAL1L2_V12_0",
		doc:  STAT_SAL1L2_V12_0{},
	},
    "TCST_TCMPR_V12_0": {
		Name: "TCST_TCMPR_V12_0",
		doc:  TCST_TCMPR_V12_0{},
	},
    "STAT_ISC_V12_0": {
		Name: "STAT_ISC_V12_0",
		doc:  STAT_ISC_V12_0{},
	},
    "STAT_NBRCNT_V12_0": {
		Name: "STAT_NBRCNT_V12_0",
		doc:  STAT_NBRCNT_V12_0{},
	},
    "STAT_PJC_V12_0": {
		Name: "STAT_PJC_V12_0",
		doc:  STAT_PJC_V12_0{},
	},
    "STAT_RPS_V12_0": {
		Name: "STAT_RPS_V12_0",
		doc:  STAT_RPS_V12_0{},
	},
    "STAT_VAL1L2_V12_0": {
		Name: "STAT_VAL1L2_V12_0",
		doc:  STAT_VAL1L2_V12_0{},
	},
    "STAT_VCNT_V12_0": {
		Name: "STAT_VCNT_V12_0",
		doc:  STAT_VCNT_V12_0{},
	},
    "TCST_TCDIAG_V12_0": {
		Name: "TCST_TCDIAG_V12_0",
		doc:  TCST_TCDIAG_V12_0{},
	},
    "STAT_MCTS_V12_0": {
		Name: "STAT_MCTS_V12_0",
		doc:  STAT_MCTS_V12_0{},
	},
    "STAT_ORANK_V12_0": {
		Name: "STAT_ORANK_V12_0",
		doc:  STAT_ORANK_V12_0{},
	},
    "STAT_PRC_V12_0": {
		Name: "STAT_PRC_V12_0",
		doc:  STAT_PRC_V12_0{},
	},
    "STAT_PHIST_V12_0": {
		Name: "STAT_PHIST_V12_0",
		doc:  STAT_PHIST_V12_0{},
	},
    "MODE_OBJ_V12_0": {
		Name: "MODE_OBJ_V12_0",
		doc:  MODE_OBJ_V12_0{},
	},
    "STAT_CTC_V12_0": {
		Name: "STAT_CTC_V12_0",
		doc:  STAT_CTC_V12_0{},
	},
    "STAT_RHIST_V12_0": {
		Name: "STAT_RHIST_V12_0",
		doc:  STAT_RHIST_V12_0{},
	},
    "STAT_RELP_V12_0": {
		Name: "STAT_RELP_V12_0",
		doc:  STAT_RELP_V12_0{},
	},
    "STAT_SSVAR_V12_0": {
		Name: "STAT_SSVAR_V12_0",
		doc:  STAT_SSVAR_V12_0{},
	},
    "STAT_SEEPS_V12_0": {
		Name: "STAT_SEEPS_V12_0",
		doc:  STAT_SEEPS_V12_0{},
	},
    "STAT_ECLV_V12_0": {
		Name: "STAT_ECLV_V12_0",
		doc:  STAT_ECLV_V12_0{},
	},
    "STAT_GENMPR_V12_0": {
		Name: "STAT_GENMPR_V12_0",
		doc:  STAT_GENMPR_V12_0{},
	},
    "TCST_PROBRIRW_V12_0": {
		Name: "TCST_PROBRIRW_V12_0",
		doc:  TCST_PROBRIRW_V12_0{},
	},
    "STAT_NBRCTC_V12_0": {
		Name: "STAT_NBRCTC_V12_0",
		doc:  STAT_NBRCTC_V12_0{},
	},
    "STAT_GRAD_V12_0": {
		Name: "STAT_GRAD_V12_0",
		doc:  STAT_GRAD_V12_0{},
	},
    "STAT_PCT_V12_0": {
		Name: "STAT_PCT_V12_0",
		doc:  STAT_PCT_V12_0{},
	},
    "STAT_PSTD_V12_0": {
		Name: "STAT_PSTD_V12_0",
		doc:  STAT_PSTD_V12_0{},
	},
    "STAT_CTS_V12_0": {
		Name: "STAT_CTS_V12_0",
		doc:  STAT_CTS_V12_0{},
	},
    "STAT_SEEPS_MPR_V12_0": {
		Name: "STAT_SEEPS_MPR_V12_0",
		doc:  STAT_SEEPS_MPR_V12_0{},
	},
    "STAT_ECNT_V12_0": {
		Name: "STAT_ECNT_V12_0",
		doc:  STAT_ECNT_V12_0{},
	},
    "MODE_CTS_V12_0": {
		Name: "MODE_CTS_V12_0",
		doc:  MODE_CTS_V12_0{},
	},
    "STAT_CNT_V12_0": {
		Name: "STAT_CNT_V12_0",
		doc:  STAT_CNT_V12_0{},
	},
    "STAT_MPR_V12_0": {
		Name: "STAT_MPR_V12_0",
		doc:  STAT_MPR_V12_0{},
	},
    "STAT_SL1L2_V12_0": {
		Name: "STAT_SL1L2_V12_0",
		doc:  STAT_SL1L2_V12_0{},
	},
    "STAT_SSIDX_V12_0": {
		Name: "STAT_SSIDX_V12_0",
		doc:  STAT_SSIDX_V12_0{},
	},
    "STAT_VL1L2_V12_0": {
		Name: "STAT_VL1L2_V12_0",
		doc:  STAT_VL1L2_V12_0{},
	},
    "STAT_FHO_V12_0": {
		Name: "STAT_FHO_V12_0",
		doc:  STAT_FHO_V12_0{},
	},
    "STAT_MCTC_V12_0": {
		Name: "STAT_MCTC_V12_0",
		doc:  STAT_MCTC_V12_0{},
	},
    "STAT_NBRCTS_V12_0": {
		Name: "STAT_NBRCTS_V12_0",
		doc:  STAT_NBRCTS_V12_0{},
	},
    "STAT_DMAP_V12_0": {
		Name: "STAT_DMAP_V12_0",
		doc:  STAT_DMAP_V12_0{},
	},
}
