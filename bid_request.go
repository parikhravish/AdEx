package main

// BidRequest represents a publisher's request for a bid
type BidRequest struct {
	AdPlacements       []AdPlacement  `json:"ad_placements"`
}

// Bid represents a bid for a single ad placement
type AdPlacement struct {
	PlacementID int          `json:"id"`
	Sizes []Size `json:"size"`
}
