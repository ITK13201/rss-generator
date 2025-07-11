// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// FeedsColumns holds the columns for the "feeds" table.
	FeedsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "title", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "link", Type: field.TypeString},
		{Name: "published_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "site_id", Type: field.TypeInt},
	}
	// FeedsTable holds the schema information for the "feeds" table.
	FeedsTable = &schema.Table{
		Name:       "feeds",
		Columns:    FeedsColumns,
		PrimaryKey: []*schema.Column{FeedsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "feeds_sites_site",
				Columns:    []*schema.Column{FeedsColumns[7]},
				RefColumns: []*schema.Column{SitesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// FeedItemsColumns holds the columns for the "feed_items" table.
	FeedItemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "link", Type: field.TypeString, Nullable: true},
		{Name: "published_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "feed_id", Type: field.TypeUUID},
	}
	// FeedItemsTable holds the schema information for the "feed_items" table.
	FeedItemsTable = &schema.Table{
		Name:       "feed_items",
		Columns:    FeedItemsColumns,
		PrimaryKey: []*schema.Column{FeedItemsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "feed_items_feeds_feed",
				Columns:    []*schema.Column{FeedItemsColumns[7]},
				RefColumns: []*schema.Column{FeedsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// ScrapingSettingsColumns holds the columns for the "scraping_settings" table.
	ScrapingSettingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "selector", Type: field.TypeString},
		{Name: "inner_selector", Type: field.TypeString},
		{Name: "title_selector", Type: field.TypeString},
		{Name: "description_selector", Type: field.TypeString, Nullable: true},
		{Name: "link_selector", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "site_id", Type: field.TypeInt},
	}
	// ScrapingSettingsTable holds the schema information for the "scraping_settings" table.
	ScrapingSettingsTable = &schema.Table{
		Name:       "scraping_settings",
		Columns:    ScrapingSettingsColumns,
		PrimaryKey: []*schema.Column{ScrapingSettingsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "scraping_settings_sites_site",
				Columns:    []*schema.Column{ScrapingSettingsColumns[8]},
				RefColumns: []*schema.Column{SitesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// SitesColumns holds the columns for the "sites" table.
	SitesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "slug", Type: field.TypeString, Unique: true},
		{Name: "title", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "url", Type: field.TypeString},
		{Name: "enable_js_rendering", Type: field.TypeBool, Default: false},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// SitesTable holds the schema information for the "sites" table.
	SitesTable = &schema.Table{
		Name:       "sites",
		Columns:    SitesColumns,
		PrimaryKey: []*schema.Column{SitesColumns[0]},
	}
	// TestFeedsColumns holds the columns for the "test_feeds" table.
	TestFeedsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "title", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "link", Type: field.TypeString},
		{Name: "published_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "site_id", Type: field.TypeInt},
	}
	// TestFeedsTable holds the schema information for the "test_feeds" table.
	TestFeedsTable = &schema.Table{
		Name:       "test_feeds",
		Columns:    TestFeedsColumns,
		PrimaryKey: []*schema.Column{TestFeedsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "test_feeds_sites_site",
				Columns:    []*schema.Column{TestFeedsColumns[7]},
				RefColumns: []*schema.Column{SitesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// TestFeedItemsColumns holds the columns for the "test_feed_items" table.
	TestFeedItemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "link", Type: field.TypeString, Nullable: true},
		{Name: "published_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "test_feed_id", Type: field.TypeUUID},
	}
	// TestFeedItemsTable holds the schema information for the "test_feed_items" table.
	TestFeedItemsTable = &schema.Table{
		Name:       "test_feed_items",
		Columns:    TestFeedItemsColumns,
		PrimaryKey: []*schema.Column{TestFeedItemsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "test_feed_items_test_feeds_test_feed",
				Columns:    []*schema.Column{TestFeedItemsColumns[7]},
				RefColumns: []*schema.Column{TestFeedsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		FeedsTable,
		FeedItemsTable,
		ScrapingSettingsTable,
		SitesTable,
		TestFeedsTable,
		TestFeedItemsTable,
	}
)

func init() {
	FeedsTable.ForeignKeys[0].RefTable = SitesTable
	FeedsTable.Annotation = &entsql.Annotation{
		Table: "feeds",
	}
	FeedItemsTable.ForeignKeys[0].RefTable = FeedsTable
	FeedItemsTable.Annotation = &entsql.Annotation{
		Table: "feed_items",
	}
	ScrapingSettingsTable.ForeignKeys[0].RefTable = SitesTable
	ScrapingSettingsTable.Annotation = &entsql.Annotation{
		Table: "scraping_settings",
	}
	SitesTable.Annotation = &entsql.Annotation{
		Table: "sites",
	}
	TestFeedsTable.ForeignKeys[0].RefTable = SitesTable
	TestFeedsTable.Annotation = &entsql.Annotation{
		Table: "test_feeds",
	}
	TestFeedItemsTable.ForeignKeys[0].RefTable = TestFeedsTable
	TestFeedItemsTable.Annotation = &entsql.Annotation{
		Table: "test_feed_items",
	}
}
