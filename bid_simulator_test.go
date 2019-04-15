package main

import (
    "testing"
)

func testSimulateBid(t *testing.T) {
	var adplacement AdPlacement 
	adplacement.PlacementID = 1
	adplacement.Sizes[0].Height = 200
	adplacement.Sizes[0].Width = 100
	var bid = simulateBid(adplacement, GetAllBidders()[0])
	
	if len(bid.AdCreatives) == 0  {
	t.Errorf("Error generating a AdCreatives for PlacementID: %d by Bidder: %s", adplacement.PlacementID, GetAllBidders()[0])
	}
	
	if bid.Price < 0.1 {
	t.Errorf("Error generating a Price for PlacementID: %d by Bidder: %s", adplacement.PlacementID, GetAllBidders()[0])
	}
}