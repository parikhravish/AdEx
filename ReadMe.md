## Running the AdEx service

curl -X POST --data @PATHTOJSONREQUEST http://localhost:5000/auction

- PATHTOJSONREQUEST : AdPlacement reqeust (added a few sample requests in the SampleRequests Package)

Notes:

- bid_simulator: responsible for simulating bids given an adplacement and a bidder
- auction_manager: responsible for managing an auction (Starting an auction and at the end determining the winner)
- adex_response: Object representing AdEx Response

- Bid Responses are generated randomly. 
	- Number of AdCreatives items is also generated randomly.
	- Sizes associated with each AdCreatives is a static list of sizes in bid_simulater which we choose from randomly.{{200,400}, {350,450}, {150,300}, {400,800}, {700,900}, {10,20}, {125,140},{250,540},{225,240},{425,540}}
- For simplicity purposes, auction logic is designed serially. In an ideal scenario, it would be a concurrent design where we would simultaneously receive bids. 
