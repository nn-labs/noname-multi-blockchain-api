package eth

import (
	"fmt"
	"testing"
)

func TestGetEthSendRawTransaction(t *testing.T) {
	ethClient := GetBaseSetupTest()
	params := []string{
		"0xd46e8dd67c5d32be8d46e8dd67c5d32be8058bb8eb970870f072445675058bb8eb970870f072445675",
	}
	ethProtocolVersion, err := GetEthSendRawTransaction(ethClient, params)
	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println(ethProtocolVersion)
}
