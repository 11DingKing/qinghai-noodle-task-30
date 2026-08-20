package brand

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestQinghaiBrandTask30(t *testing.T) {
	now := time.Now()
	r := NewRegistry()
	require.NoError(t, r.PublishCatalog(context.Background(), CatalogSnapshot{StoreID: "store-1", GeneratedAt: now, Listings: []ProductListing{{SKU: "yak", StoreID: "store-1", Published: true}}}))
	s := NewService(r, func() time.Time { return now })
	c := CultureCampaign{StoreID: "store-1", Title: "大美青海", ContentVersion: 2, ApprovedVersion: 2, StartsAt: now.Add(-time.Hour), EndsAt: now.Add(time.Hour), FeaturedSKUs: []string{"yak"}, DestinationCodes: []string{"TJ"}, RegionalBrandLogo: true}
	_, err := s.LaunchCampaign(context.Background(), c, compliantStore(now), activeLicense(now))
	require.NoError(t, err)
}
