// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ITK13201/rss-generator/ent/scrapingsetting"
	"github.com/ITK13201/rss-generator/ent/site"
)

// ScrapingSetting is the model entity for the ScrapingSetting schema.
type ScrapingSetting struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Selector holds the value of the "selector" field.
	Selector string `json:"selector,omitempty"`
	// InnerSelector holds the value of the "inner_selector" field.
	InnerSelector string `json:"inner_selector,omitempty"`
	// TitleSelector holds the value of the "title_selector" field.
	TitleSelector string `json:"title_selector,omitempty"`
	// DescriptionSelector holds the value of the "description_selector" field.
	DescriptionSelector string `json:"description_selector,omitempty"`
	// LinkSelector holds the value of the "link_selector" field.
	LinkSelector string `json:"link_selector,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ScrapingSettingQuery when eager-loading is set.
	Edges        ScrapingSettingEdges `json:"edges"`
	site_id      *int
	selectValues sql.SelectValues
}

// ScrapingSettingEdges holds the relations/edges for other nodes in the graph.
type ScrapingSettingEdges struct {
	// Site holds the value of the site edge.
	Site *Site `json:"site,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// SiteOrErr returns the Site value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ScrapingSettingEdges) SiteOrErr() (*Site, error) {
	if e.Site != nil {
		return e.Site, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: site.Label}
	}
	return nil, &NotLoadedError{edge: "site"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ScrapingSetting) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case scrapingsetting.FieldID:
			values[i] = new(sql.NullInt64)
		case scrapingsetting.FieldSelector, scrapingsetting.FieldInnerSelector, scrapingsetting.FieldTitleSelector, scrapingsetting.FieldDescriptionSelector, scrapingsetting.FieldLinkSelector:
			values[i] = new(sql.NullString)
		case scrapingsetting.FieldCreatedAt, scrapingsetting.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case scrapingsetting.ForeignKeys[0]: // site_id
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ScrapingSetting fields.
func (ss *ScrapingSetting) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case scrapingsetting.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ss.ID = int(value.Int64)
		case scrapingsetting.FieldSelector:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field selector", values[i])
			} else if value.Valid {
				ss.Selector = value.String
			}
		case scrapingsetting.FieldInnerSelector:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field inner_selector", values[i])
			} else if value.Valid {
				ss.InnerSelector = value.String
			}
		case scrapingsetting.FieldTitleSelector:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title_selector", values[i])
			} else if value.Valid {
				ss.TitleSelector = value.String
			}
		case scrapingsetting.FieldDescriptionSelector:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description_selector", values[i])
			} else if value.Valid {
				ss.DescriptionSelector = value.String
			}
		case scrapingsetting.FieldLinkSelector:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field link_selector", values[i])
			} else if value.Valid {
				ss.LinkSelector = value.String
			}
		case scrapingsetting.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ss.CreatedAt = value.Time
			}
		case scrapingsetting.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ss.UpdatedAt = value.Time
			}
		case scrapingsetting.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field site_id", value)
			} else if value.Valid {
				ss.site_id = new(int)
				*ss.site_id = int(value.Int64)
			}
		default:
			ss.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ScrapingSetting.
// This includes values selected through modifiers, order, etc.
func (ss *ScrapingSetting) Value(name string) (ent.Value, error) {
	return ss.selectValues.Get(name)
}

// QuerySite queries the "site" edge of the ScrapingSetting entity.
func (ss *ScrapingSetting) QuerySite() *SiteQuery {
	return NewScrapingSettingClient(ss.config).QuerySite(ss)
}

// Update returns a builder for updating this ScrapingSetting.
// Note that you need to call ScrapingSetting.Unwrap() before calling this method if this ScrapingSetting
// was returned from a transaction, and the transaction was committed or rolled back.
func (ss *ScrapingSetting) Update() *ScrapingSettingUpdateOne {
	return NewScrapingSettingClient(ss.config).UpdateOne(ss)
}

// Unwrap unwraps the ScrapingSetting entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ss *ScrapingSetting) Unwrap() *ScrapingSetting {
	_tx, ok := ss.config.driver.(*txDriver)
	if !ok {
		panic("ent: ScrapingSetting is not a transactional entity")
	}
	ss.config.driver = _tx.drv
	return ss
}

// String implements the fmt.Stringer.
func (ss *ScrapingSetting) String() string {
	var builder strings.Builder
	builder.WriteString("ScrapingSetting(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ss.ID))
	builder.WriteString("selector=")
	builder.WriteString(ss.Selector)
	builder.WriteString(", ")
	builder.WriteString("inner_selector=")
	builder.WriteString(ss.InnerSelector)
	builder.WriteString(", ")
	builder.WriteString("title_selector=")
	builder.WriteString(ss.TitleSelector)
	builder.WriteString(", ")
	builder.WriteString("description_selector=")
	builder.WriteString(ss.DescriptionSelector)
	builder.WriteString(", ")
	builder.WriteString("link_selector=")
	builder.WriteString(ss.LinkSelector)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(ss.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ss.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// ScrapingSettings is a parsable slice of ScrapingSetting.
type ScrapingSettings []*ScrapingSetting