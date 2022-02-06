package eth

import (
	"fmt"
	"testing"
)

func TestGetEthCompilers(t *testing.T) {
	ethClient := GetBaseSetupTest()
	baseResponse, err := GetEthCompilers(ethClient, []string{})
	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println(baseResponse)
}
