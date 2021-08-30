package main

import (
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/gobuffalo/envy"
)

const (
	AlgoliaAppIDEnv       = "ALGOLIA_APP_ID"
	AlgoliaAdminAPIKeyEnv = "ALGOLIA_API_KEY"
	AlgoliaItemsIndexEnv  = "ALGOLIA_ITEMS_INDEX"
	AlgoliaOffersIndexEnv = "ALGOLIA_OFFERS_INDEX"
)

func main() {
	algAppID := envy.Get(AlgoliaAppIDEnv, "")
	algAPIKey := envy.Get(AlgoliaAdminAPIKeyEnv, "")
	algItemsIndex := envy.Get(AlgoliaItemsIndexEnv, "items_test")
	algOffersIndex := envy.Get(AlgoliaOffersIndexEnv, "offers_test")

	client := search.NewClient(algAppID, algAPIKey)
	itemIndex := client.InitIndex(algItemsIndex)
	offerIndex := client.InitIndex(algOffersIndex)

}
