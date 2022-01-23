package eth

import (
	"fmt"
	"testing"
)

func TestGetNetListing(t *testing.T) {
	ethClient := GetBaseSetupTest()
	netListing, err := GetNetListing(ethClient)
	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println(netListing)
}
