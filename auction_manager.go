package main

import (
"net/http"
	"encoding/json"
	"strconv"
	"strings"
)

func startAuction(adPlacements []AdPlacement, bidders []Bidder, w http.ResponseWriter){
	var bidderResponseList []BidResponse
	
	for _,bidder := range bidders {
		
		var bids []Bid
		var bidderResponse BidResponse
		for _,adPlacement := range adPlacements {
			
			bids = append(bids, simulateBid(adPlacement,bidder))
		}
		bidderResponse.Bids = bids
		bidderResponse.BidderName = bidder
		bidderResponseList = append(bidderResponseList, bidderResponse)
	}
	determineWinner(adPlacements, bidderResponseList, w)
}

func determineWinner(adPlacements []AdPlacement, bidderResponseList []BidResponse, w http.ResponseWriter) (AdExResponse){
	
	var adPlacementSizes map[Size]bool
	var maxBid = 0.0
	
	var adExResponse AdExResponse
	var adExPlacements []AdPlacement_AdEx
	
	for index,adPlacement :=range adPlacements {
		
		adPlacementSizes = make(map[Size]bool)
		
		for _,size := range adPlacement.Sizes{			
			adPlacementSizes[size] = true
		}
		var adExPlacement AdPlacement_AdEx
		adExPlacement.PlacementID = adPlacement.PlacementID
		
		for _ , bidderResponse := range bidderResponseList {
			if bidderResponse.Bids[index].Price > maxBid {
				for _,adCreative := range bidderResponse.Bids[index].AdCreatives {
					if adPlacementSizes[adCreative.Size] == true {
						adExPlacement.Price = bidderResponse.Bids[index].Price
						adExPlacement.URLAdEx, _ = strconv.Unquote("\"" + adCreative.URL.String() + "\"" )
						maxBid = bidderResponse.Bids[index].Price
					}
				}	
			}
		}
		maxBid = 0.0
		adExPlacements = append(adExPlacements, adExPlacement)
		
	}
	adExResponse.AdPlacements = adExPlacements
	adExResponseJSON, _ := json.Marshal(adExResponse)
	var modAdExResponseJSON, _ = (strconv.Unquote(strings.Replace(strconv.Quote(string(adExResponseJSON)), `\\u`, `\u`, -1)))
	if w != nil { 
		w.Write([]byte (modAdExResponseJSON))
	}
	return adExResponse
}
