// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ITK13201/rss-generator/ent/site"
)

// Site is the model entity for the Site schema.
type Site struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Slug holds the value of the "slug" field.
	Slug string `json:"slug,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// URL holds the value of the "url" field.
	URL string `json:"url,omitempty"`
	// EnableJsRendering holds the value of the "enable_js_rendering" field.
	EnableJsRendering bool `json:"enable_js_rendering,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SiteQuery when eager-loading is set.
	Edges        SiteEdges `json:"edges"`
	selectValues sql.SelectValues
}

// SiteEdges holds the relations/edges for other nodes in the graph.
type SiteEdges struct {
	// ScrapingSettings holds the value of the scraping_settings edge.
	ScrapingSettings []*ScrapingSetting `json:"scraping_settings,omitempty"`
	// Feeds holds the value of the feeds edge.
	Feeds []*Feed `json:"feeds,omitempty"`
	// TestFeeds holds the value of the test_feeds edge.
	TestFeeds []*TestFeed `json:"test_feeds,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// ScrapingSettingsOrErr returns the ScrapingSettings value or an error if the edge
// was not loaded in eager-loading.
func (e SiteEdges) ScrapingSettingsOrErr() ([]*ScrapingSetting, error) {
	if e.loadedTypes[0] {
		return e.ScrapingSettings, nil
	}
	return nil, &NotLoadedError{edge: "scraping_settings"}
}

// FeedsOrErr returns the Feeds value or an error if the edge
// was not loaded in eager-loading.
func (e SiteEdges) FeedsOrErr() ([]*Feed, error) {
	if e.loadedTypes[1] {
		return e.Feeds, nil
	}
	return nil, &NotLoadedError{edge: "feeds"}
}

// TestFeedsOrErr returns the TestFeeds value or an error if the edge
// was not loaded in eager-loading.
func (e SiteEdges) TestFeedsOrErr() ([]*TestFeed, error) {
	if e.loadedTypes[2] {
		return e.TestFeeds, nil
	}
	return nil, &NotLoadedError{edge: "test_feeds"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Site) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case site.FieldEnableJsRendering:
			values[i] = new(sql.NullBool)
		case site.FieldID:
			values[i] = new(sql.NullInt64)
		case site.FieldSlug, site.FieldTitle, site.FieldDescription, site.FieldURL:
			values[i] = new(sql.NullString)
		case site.FieldCreatedAt, site.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Site fields.
func (s *Site) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case site.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case site.FieldSlug:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field slug", values[i])
			} else if value.Valid {
				s.Slug = value.String
			}
		case site.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				s.Title = value.String
			}
		case site.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				s.Description = value.String
			}
		case site.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				s.URL = value.String
			}
		case site.FieldEnableJsRendering:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field enable_js_rendering", values[i])
			} else if value.Valid {
				s.EnableJsRendering = value.Bool
			}
		case site.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				s.CreatedAt = value.Time
			}
		case site.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				s.UpdatedAt = value.Time
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Site.
// This includes values selected through modifiers, order, etc.
func (s *Site) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// QueryScrapingSettings queries the "scraping_settings" edge of the Site entity.
func (s *Site) QueryScrapingSettings() *ScrapingSettingQuery {
	return NewSiteClient(s.config).QueryScrapingSettings(s)
}

// QueryFeeds queries the "feeds" edge of the Site entity.
func (s *Site) QueryFeeds() *FeedQuery {
	return NewSiteClient(s.config).QueryFeeds(s)
}

// QueryTestFeeds queries the "test_feeds" edge of the Site entity.
func (s *Site) QueryTestFeeds() *TestFeedQuery {
	return NewSiteClient(s.config).QueryTestFeeds(s)
}

// Update returns a builder for updating this Site.
// Note that you need to call Site.Unwrap() before calling this method if this Site
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Site) Update() *SiteUpdateOne {
	return NewSiteClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Site entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Site) Unwrap() *Site {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Site is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Site) String() string {
	var builder strings.Builder
	builder.WriteString("Site(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("slug=")
	builder.WriteString(s.Slug)
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(s.Title)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(s.Description)
	builder.WriteString(", ")
	builder.WriteString("url=")
	builder.WriteString(s.URL)
	builder.WriteString(", ")
	builder.WriteString("enable_js_rendering=")
	builder.WriteString(fmt.Sprintf("%v", s.EnableJsRendering))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(s.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(s.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Sites is a parsable slice of Site.
type Sites []*Site
