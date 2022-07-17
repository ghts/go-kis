package kis

import (
	"context"
	"log"
	"net/http"

	"moul.io/http2curl"
)

type DomesticStockQuotationsService service

type InquirePriceResponse struct {
	Output struct {
		IscdStatClsCode      string `json:"iscd_stat_cls_code"`
		MargRate             string `json:"marg_rate"`
		RprsMrktKorName      string `json:"rprs_mrkt_kor_name"`
		BstpKorIsnm          string `json:"bstp_kor_isnm"`
		TempStopYn           string `json:"temp_stop_yn"`
		OprcRangContYn       string `json:"oprc_rang_cont_yn"`
		ClprRangContYn       string `json:"clpr_rang_cont_yn"`
		CrdtAbleYn           string `json:"crdt_able_yn"`
		GrmnRateClsCode      string `json:"grmn_rate_cls_code"`
		ElwPblcYn            string `json:"elw_pblc_yn"`
		StckPrpr             string `json:"stck_prpr"`
		PrdyVrss             string `json:"prdy_vrss"`
		PrdyVrssSign         string `json:"prdy_vrss_sign"`
		PrdyCtrt             string `json:"prdy_ctrt"`
		AcmlTrPbmn           string `json:"acml_tr_pbmn"`
		AcmlVol              string `json:"acml_vol"`
		PrdyVrssVolRate      string `json:"prdy_vrss_vol_rate"`
		StckOprc             string `json:"stck_oprc"`
		StckHgpr             string `json:"stck_hgpr"`
		StckLwpr             string `json:"stck_lwpr"`
		StckMxpr             string `json:"stck_mxpr"`
		StckLlam             string `json:"stck_llam"`
		StckSdpr             string `json:"stck_sdpr"`
		WghnAvrgStckPrc      string `json:"wghn_avrg_stck_prc"`
		HtsFrgnEhrt          string `json:"hts_frgn_ehrt"`
		FrgnNtbyQty          string `json:"frgn_ntby_qty"`
		PgtrNtbyQty          string `json:"pgtr_ntby_qty"`
		PvtScndDmrsPrc       string `json:"pvt_scnd_dmrs_prc"`
		PvtFrstDmrsPrc       string `json:"pvt_frst_dmrs_prc"`
		PvtPontVal           string `json:"pvt_pont_val"`
		PvtFrstDmspPrc       string `json:"pvt_frst_dmsp_prc"`
		PvtScndDmspPrc       string `json:"pvt_scnd_dmsp_prc"`
		DmrsVal              string `json:"dmrs_val"`
		DmspVal              string `json:"dmsp_val"`
		Cpfn                 string `json:"cpfn"`
		RstcWdthPrc          string `json:"rstc_wdth_prc"`
		StckFcam             string `json:"stck_fcam"`
		StckSspr             string `json:"stck_sspr"`
		AsprUnit             string `json:"aspr_unit"`
		HtsDealQtyUnitVal    string `json:"hts_deal_qty_unit_val"`
		LstnStcn             string `json:"lstn_stcn"`
		HtsAvls              string `json:"hts_avls"`
		Per                  string `json:"per"`
		Pbr                  string `json:"pbr"`
		StacMonth            string `json:"stac_month"`
		VolTnrt              string `json:"vol_tnrt"`
		Eps                  string `json:"eps"`
		Bps                  string `json:"bps"`
		D250Hgpr             string `json:"d250_hgpr"`
		D250HgprDate         string `json:"d250_hgpr_date"`
		D250HgprVrssPrprRate string `json:"d250_hgpr_vrss_prpr_rate"`
		D250Lwpr             string `json:"d250_lwpr"`
		D250LwprDate         string `json:"d250_lwpr_date"`
		D250LwprVrssPrprRate string `json:"d250_lwpr_vrss_prpr_rate"`
		StckDryyHgpr         string `json:"stck_dryy_hgpr"`
		DryyHgprVrssPrprRate string `json:"dryy_hgpr_vrss_prpr_rate"`
		DryyHgprDate         string `json:"dryy_hgpr_date"`
		StckDryyLwpr         string `json:"stck_dryy_lwpr"`
		DryyLwprVrssPrprRate string `json:"dryy_lwpr_vrss_prpr_rate"`
		DryyLwprDate         string `json:"dryy_lwpr_date"`
		W52Hgpr              string `json:"w52_hgpr"`
		W52HgprVrssPrprCtrt  string `json:"w52_hgpr_vrss_prpr_ctrt"`
		W52HgprDate          string `json:"w52_hgpr_date"`
		W52Lwpr              string `json:"w52_lwpr"`
		W52LwprVrssPrprCtrt  string `json:"w52_lwpr_vrss_prpr_ctrt"`
		W52LwprDate          string `json:"w52_lwpr_date"`
		WholLoanRmndRate     string `json:"whol_loan_rmnd_rate"`
		SstsYn               string `json:"ssts_yn"`
		StckShrnIscd         string `json:"stck_shrn_iscd"`
		FcamCnnm             string `json:"fcam_cnnm"`
		CpfnCnnm             string `json:"cpfn_cnnm"`
		FrgnHldnQty          string `json:"frgn_hldn_qty"`
		ViClsCode            string `json:"vi_cls_code"`
		OvtmViClsCode        string `json:"ovtm_vi_cls_code"`
		LastSstsCntgQty      string `json:"last_ssts_cntg_qty"`
		InvtCafulYn          string `json:"invt_caful_yn"`
		MrktWarnClsCode      string `json:"mrkt_warn_cls_code"`
		ShortOverYn          string `json:"short_over_yn"`
		SltrYn               string `json:"sltr_yn"`
	} `json:"output"`
	RtCd  string `json:"rt_cd"`
	MsgCd string `json:"msg_cd"`
	Msg1  string `json:"msg1"`
}

// https://apiportal.koreainvestment.com/apiservice/apiservice-domestic-stock-quotations#L_07802512-4f49-4486-91b4-1050b6f5dc9d
func (s *DomesticStockQuotationsService) InquirePrice(ctx context.Context, FidCondMrktDivCode string, FidInputIscd string) (*InquirePriceResponse, *Response, error) {
	u := "/uapi/domestic-stock/v1/quotations/inquire-price"

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	q := req.URL.Query()
	q.Add("FID_COND_MRKT_DIV_CODE", FidCondMrktDivCode)
	q.Add("FID_INPUT_ISCD", FidInputIscd)
	req.URL.RawQuery = q.Encode()

	req.Header.Add("tr_id", "FHKST01010100")

	req.Header.Add("appkey", s.client.AppKey)
	req.Header.Add("appsecret", s.client.AppSecret)
	req.Header.Add("authorization", "Bearer "+s.client.AccessToken)
	req.Header.Add("content-type", "application/json; charset=utf-8")

	if s.client.Debug {
		log.Println(http2curl.GetCurlCommand(req))
	}

	respBody := new(InquirePriceResponse)
	resp, err := s.client.Do(ctx, req, respBody)
	if err != nil {
		return nil, resp, err
	}

	return respBody, resp, nil
}

type InquireDailyPriceResponse struct {
	Output []struct {
		StckBsopDate    string `json:"stck_bsop_date"`
		StckOprc        string `json:"stck_oprc"`
		StckHgpr        string `json:"stck_hgpr"`
		StckLwpr        string `json:"stck_lwpr"`
		StckClpr        string `json:"stck_clpr"`
		AcmlVol         string `json:"acml_vol"`
		PrdyVrssVolRate string `json:"prdy_vrss_vol_rate"`
		PrdyVrss        string `json:"prdy_vrss"`
		PrdyVrssSign    string `json:"prdy_vrss_sign"`
		PrdyCtrt        string `json:"prdy_ctrt"`
		HtsFrgnEhrt     string `json:"hts_frgn_ehrt"`
		FrgnNtbyQty     string `json:"frgn_ntby_qty"`
		FlngClsCode     string `json:"flng_cls_code"`
		AcmlPrttRate    string `json:"acml_prtt_rate"`
	} `json:"output"`
	RtCd  string `json:"rt_cd"`
	MsgCd string `json:"msg_cd"`
	Msg1  string `json:"msg1"`
}

// https://apiportal.koreainvestment.com/apiservice/apiservice-domestic-stock-quotations#L_011d4de2-a4a0-47c0-aa47-20c65a26a763
func (s *DomesticStockQuotationsService) InquireDailyPrice(ctx context.Context, FidCondMrktDivCode string, FidInputIscd string, FidPeriodDivCode string, FidOrgAdjPrc string) (*InquireDailyPriceResponse, *Response, error) {
	u := "/uapi/domestic-stock/v1/quotations/inquire-daily-price"

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	q := req.URL.Query()
	q.Add("FID_COND_MRKT_DIV_CODE", FidCondMrktDivCode)
	q.Add("FID_INPUT_ISCD", FidInputIscd)
	q.Add("FID_PERIOD_DIV_CODE", FidPeriodDivCode)
	q.Add("FID_ORG_ADJ_PRC", FidOrgAdjPrc)
	req.URL.RawQuery = q.Encode()

	req.Header.Add("tr_id", "FHKST01010400")

	req.Header.Add("appkey", s.client.AppKey)
	req.Header.Add("appsecret", s.client.AppSecret)
	req.Header.Add("authorization", "Bearer "+s.client.AccessToken)
	req.Header.Add("content-type", "application/json; charset=utf-8")

	if s.client.Debug {
		log.Println(http2curl.GetCurlCommand(req))
	}

	respBody := new(InquireDailyPriceResponse)
	resp, err := s.client.Do(ctx, req, respBody)
	if err != nil {
		return nil, resp, err
	}

	return respBody, resp, nil
}
