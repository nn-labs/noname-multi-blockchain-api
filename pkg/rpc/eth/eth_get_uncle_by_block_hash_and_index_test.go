package eth

import (
	"fmt"
	"testing"
)

func TestGetEthUncleByBlockHashAndIndex(t *testing.T) {
	ethClient := GetBaseSetupTest()
	params := []string{
		"0xc6ef2fc5426d6ad6fd9e2a26abeab0aa2411b7ab17f30a99d3cb96aed1d1055b",
		"0x0",
	}
	baseResponse, err := GetEthUncleByBlockHashAndIndex(ethClient, params)
	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println(baseResponse)
}
