package main

import (
	"math/rand"
	"net/url"
	"fmt"
	"strconv"
)

//10 Image sizes supported to limit the possibilites of generated sizes.
var sizeList = []Size{{200,400}, {350,450}, {150,300}, {400,800}, {700,900}, {10,20}, {125,140},{250,540},{225,240},{425,540}}

func simulateBid(adplacement AdPlacement, bidder Bidder) Bid{
	
	// List to hold instances of adcreatives.
	var adCreativeList []AdCreative
	
	//Randomize the number of adCreatives generated for each bid.
	var adCreativesSize = 1 + rand.Intn(len(sizeList));
	
	//For each adCreative randomly generate the index of a size from sizelist.
	for i := 0; i < adCreativesSize; i++ {
		var size Size = sizeList[1 + rand.Intn(len(sizeList)-2)]
		var urlStr,err = url.Parse(fmt.Sprintf("http://www.%s.com/img/size?w=%d&h=%d",bidder, size.Width, size.Height))
		if err ==nil{
			adCreativeList = append(adCreativeList, AdCreative{Size:size, URL:*urlStr})
		}
	}
	
	//Bid range is 1 cent to 100 dollars.
	var generatedBidPrice, _ = strconv.ParseFloat((fmt.Sprintf("%.2f", ((0.1 + rand.Float64() * (100 - 0.1))*0.05)/0.05)),64)
	
	var bid Bid
	
	bid.PlacementID = adplacement.PlacementID
	bid.AdCreatives = adCreativeList
	bid.Price = generatedBidPrice
	return bid
}


