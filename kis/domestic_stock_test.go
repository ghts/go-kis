package kis

import (
	"context"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

var testClient *Client

func TestMain(m *testing.M) {
	opts := AuthOptionsFromEnv()
	testClient = NewClient(http.DefaultClient, opts)
	testClient.Debug = true

	accessToken := os.Getenv("KIS_ACCESS_TOKEN")
	if accessToken == "" {
		resp, _, err := testClient.Oauth2.TokenP(context.TODO())
		if err != nil {
			log.Fatal(err)
		}
		_ = resp
		testClient.AccessToken = resp.AccessToken
	} else {
		testClient.AccessToken = accessToken
	}

	os.Exit(m.Run())
}

func TestDomesticStockQuotationsService_InquirePrice(t *testing.T) {
	resp, _, err := testClient.DomesticStockQuotations.InquirePrice(context.TODO(), "J", "000660")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "전기.전자", resp.Output.BstpKorIsnm)
}

func TestDomesticStockQuotationsService_InquireDailyPrice(t *testing.T) {
	InquirePriceResponse, _, err := testClient.DomesticStockQuotations.InquireDailyPrice(context.TODO(), "J", "000660", "D", "0000000001")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "0", InquirePriceResponse.RtCd)

	spew.Dump(InquirePriceResponse.RtCd)
}
