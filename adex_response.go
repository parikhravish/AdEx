package main

// AdExResponse represents an adex response to the request for a bid from publisher
type AdExResponse struct {
	AdPlacements       []AdPlacement_AdEx  `json:"ad_placements"`
}

// Bid represents a bid for a single ad placement
type AdPlacement_AdEx struct {
	PlacementID int          `json:"placement_id"`
	Price       float64      `json:"bid_price"`
	URLAdEx  string `json:"url"`
}
