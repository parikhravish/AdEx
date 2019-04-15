package main

import (
    "testing"
)
//Simple winner with same size bidders
func testDetermineWinnerSimple(t *testing.T) {
	//Single Ad Placement test
	var adplacement AdPlacement 
	adplacement.PlacementID = 1
	
	adplacement.Sizes[0].Height = 200
	adplacement.Sizes[0].Width = 100

	var bidresponseNike BidResponse
	bidresponseNike.BidderName =  Nike
	bidresponseNike.Bids[0].PlacementID = 1
	bidresponseNike.Bids[0].Price = 100
	bidresponseNike.Bids[0].AdCreatives[0].Size.Height = 200
	bidresponseNike.Bids[0].AdCreatives[0].Size.Width = 100
	
	var bidresponseApple BidResponse
	bidresponseApple.BidderName =  Apple
	bidresponseApple.Bids[0].PlacementID = 1
	bidresponseApple.Bids[0].Price = 200
	bidresponseApple.Bids[0].AdCreatives[0].Size.Height = 200
	bidresponseApple.Bids[0].AdCreatives[0].Size.Width = 100
	
	var bidderResponseList []BidResponse
	bidderResponseList = append(bidderResponseList, bidresponseNike)
	bidderResponseList = append(bidderResponseList, bidresponseApple)
	
	var adPlacements []AdPlacement
	
	adPlacements = append(adPlacements, adplacement)
	
	var adExResponse AdExResponse = determineWinner(adPlacements, bidderResponseList, nil)
	
	if adExResponse.AdPlacements[0].Price != 200 {
		t.Errorf("Incorrect bid generated for PlacementID: %d", adplacement.PlacementID)
	}
}
// No Matching Size
func testDetermineWinner_NoWinner(t *testing.T) {
	//Single Ad Placement test
	var adplacement AdPlacement 
	adplacement.PlacementID = 1
	
	adplacement.Sizes[0].Height = 200
	adplacement.Sizes[0].Width = 100

	var bidresponseNike BidResponse
	bidresponseNike.BidderName =  Nike
	bidresponseNike.Bids[0].PlacementID = 1
	bidresponseNike.Bids[0].Price = 100
	bidresponseNike.Bids[0].AdCreatives[0].Size.Height = 250
	bidresponseNike.Bids[0].AdCreatives[0].Size.Width = 100
	
	var bidresponseApple BidResponse
	bidresponseApple.BidderName =  Apple
	bidresponseApple.Bids[0].PlacementID = 1
	bidresponseApple.Bids[0].Price = 200
	bidresponseApple.Bids[0].AdCreatives[0].Size.Height = 300
	bidresponseApple.Bids[0].AdCreatives[0].Size.Width = 100
	
	var bidderResponseList []BidResponse
	bidderResponseList = append(bidderResponseList, bidresponseNike)
	bidderResponseList = append(bidderResponseList, bidresponseApple)
	
	var adPlacements []AdPlacement
	
	adPlacements = append(adPlacements, adplacement)
	
	var adExResponse AdExResponse = determineWinner(adPlacements, bidderResponseList, nil)
	
	if adExResponse.AdPlacements[0].Price != 0 {
		t.Errorf("No winner expected for PlacementID: %d", adplacement.PlacementID)
	}
}