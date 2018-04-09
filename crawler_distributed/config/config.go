package config

const (
	// Parser names
	ParseCity     = "ParseCity"
	ParseCityList = "ParseCityList"
	ParseProfile  = "ParseProfile"
	ProfileParser = "ProfileParser"
	NilParser     = "NilParser"

	// Service ports
	ItemSaverPort = 1234
	WorkPort0     = 9000
	// ElasticSearch
	ElasticIndex = "dating_profile"
	// RPC Endpoints
	ItemSaverRpc    = "ItemSaverService.Save"
	CrawlServiceRpc = "CrawlService.Process"
)
