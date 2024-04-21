// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/ITK13201/rss-generator/ent/feed"
	"github.com/ITK13201/rss-generator/ent/feeditem"
	"github.com/ITK13201/rss-generator/ent/schema"
	"github.com/ITK13201/rss-generator/ent/scrapingselector"
	"github.com/ITK13201/rss-generator/ent/site"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	feedFields := schema.Feed{}.Fields()
	_ = feedFields
	// feedDescTitle is the schema descriptor for title field.
	feedDescTitle := feedFields[1].Descriptor()
	// feed.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	feed.TitleValidator = func() func(string) error {
		validators := feedDescTitle.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(title string) error {
			for _, fn := range fns {
				if err := fn(title); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// feedDescDescription is the schema descriptor for description field.
	feedDescDescription := feedFields[2].Descriptor()
	// feed.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	feed.DescriptionValidator = func() func(string) error {
		validators := feedDescDescription.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(description string) error {
			for _, fn := range fns {
				if err := fn(description); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// feedDescLink is the schema descriptor for link field.
	feedDescLink := feedFields[3].Descriptor()
	// feed.LinkValidator is a validator for the "link" field. It is called by the builders before save.
	feed.LinkValidator = func() func(string) error {
		validators := feedDescLink.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(link string) error {
			for _, fn := range fns {
				if err := fn(link); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// feedDescIsTest is the schema descriptor for is_test field.
	feedDescIsTest := feedFields[5].Descriptor()
	// feed.DefaultIsTest holds the default value on creation for the is_test field.
	feed.DefaultIsTest = feedDescIsTest.Default.(bool)
	// feedDescCreatedAt is the schema descriptor for created_at field.
	feedDescCreatedAt := feedFields[6].Descriptor()
	// feed.DefaultCreatedAt holds the default value on creation for the created_at field.
	feed.DefaultCreatedAt = feedDescCreatedAt.Default.(func() time.Time)
	// feedDescUpdatedAt is the schema descriptor for updated_at field.
	feedDescUpdatedAt := feedFields[7].Descriptor()
	// feed.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	feed.DefaultUpdatedAt = feedDescUpdatedAt.Default.(func() time.Time)
	// feed.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	feed.UpdateDefaultUpdatedAt = feedDescUpdatedAt.UpdateDefault.(func() time.Time)
	// feedDescID is the schema descriptor for id field.
	feedDescID := feedFields[0].Descriptor()
	// feed.DefaultID holds the default value on creation for the id field.
	feed.DefaultID = feedDescID.Default.(func() uuid.UUID)
	feeditemFields := schema.FeedItem{}.Fields()
	_ = feeditemFields
	// feeditemDescTitle is the schema descriptor for title field.
	feeditemDescTitle := feeditemFields[1].Descriptor()
	// feeditem.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	feeditem.TitleValidator = func() func(string) error {
		validators := feeditemDescTitle.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(title string) error {
			for _, fn := range fns {
				if err := fn(title); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// feeditemDescDescription is the schema descriptor for description field.
	feeditemDescDescription := feeditemFields[2].Descriptor()
	// feeditem.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	feeditem.DescriptionValidator = func() func(string) error {
		validators := feeditemDescDescription.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(description string) error {
			for _, fn := range fns {
				if err := fn(description); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// feeditemDescLink is the schema descriptor for link field.
	feeditemDescLink := feeditemFields[3].Descriptor()
	// feeditem.LinkValidator is a validator for the "link" field. It is called by the builders before save.
	feeditem.LinkValidator = feeditemDescLink.Validators[0].(func(string) error)
	// feeditemDescCreatedAt is the schema descriptor for created_at field.
	feeditemDescCreatedAt := feeditemFields[5].Descriptor()
	// feeditem.DefaultCreatedAt holds the default value on creation for the created_at field.
	feeditem.DefaultCreatedAt = feeditemDescCreatedAt.Default.(func() time.Time)
	// feeditemDescUpdatedAt is the schema descriptor for updated_at field.
	feeditemDescUpdatedAt := feeditemFields[6].Descriptor()
	// feeditem.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	feeditem.DefaultUpdatedAt = feeditemDescUpdatedAt.Default.(func() time.Time)
	// feeditem.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	feeditem.UpdateDefaultUpdatedAt = feeditemDescUpdatedAt.UpdateDefault.(func() time.Time)
	scrapingselectorFields := schema.ScrapingSelector{}.Fields()
	_ = scrapingselectorFields
	// scrapingselectorDescSelector is the schema descriptor for selector field.
	scrapingselectorDescSelector := scrapingselectorFields[1].Descriptor()
	// scrapingselector.SelectorValidator is a validator for the "selector" field. It is called by the builders before save.
	scrapingselector.SelectorValidator = scrapingselectorDescSelector.Validators[0].(func(string) error)
	// scrapingselectorDescInnerSelector is the schema descriptor for inner_selector field.
	scrapingselectorDescInnerSelector := scrapingselectorFields[2].Descriptor()
	// scrapingselector.InnerSelectorValidator is a validator for the "inner_selector" field. It is called by the builders before save.
	scrapingselector.InnerSelectorValidator = scrapingselectorDescInnerSelector.Validators[0].(func(string) error)
	// scrapingselectorDescTitleSelector is the schema descriptor for title_selector field.
	scrapingselectorDescTitleSelector := scrapingselectorFields[3].Descriptor()
	// scrapingselector.TitleSelectorValidator is a validator for the "title_selector" field. It is called by the builders before save.
	scrapingselector.TitleSelectorValidator = scrapingselectorDescTitleSelector.Validators[0].(func(string) error)
	// scrapingselectorDescCreatedAt is the schema descriptor for created_at field.
	scrapingselectorDescCreatedAt := scrapingselectorFields[6].Descriptor()
	// scrapingselector.DefaultCreatedAt holds the default value on creation for the created_at field.
	scrapingselector.DefaultCreatedAt = scrapingselectorDescCreatedAt.Default.(func() time.Time)
	// scrapingselectorDescUpdatedAt is the schema descriptor for updated_at field.
	scrapingselectorDescUpdatedAt := scrapingselectorFields[7].Descriptor()
	// scrapingselector.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	scrapingselector.DefaultUpdatedAt = scrapingselectorDescUpdatedAt.Default.(func() time.Time)
	// scrapingselector.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	scrapingselector.UpdateDefaultUpdatedAt = scrapingselectorDescUpdatedAt.UpdateDefault.(func() time.Time)
	siteFields := schema.Site{}.Fields()
	_ = siteFields
	// siteDescSlug is the schema descriptor for slug field.
	siteDescSlug := siteFields[1].Descriptor()
	// site.SlugValidator is a validator for the "slug" field. It is called by the builders before save.
	site.SlugValidator = siteDescSlug.Validators[0].(func(string) error)
	// siteDescTitle is the schema descriptor for title field.
	siteDescTitle := siteFields[2].Descriptor()
	// site.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	site.TitleValidator = siteDescTitle.Validators[0].(func(string) error)
	// siteDescURL is the schema descriptor for url field.
	siteDescURL := siteFields[4].Descriptor()
	// site.URLValidator is a validator for the "url" field. It is called by the builders before save.
	site.URLValidator = siteDescURL.Validators[0].(func(string) error)
	// siteDescEnableJsRendering is the schema descriptor for enable_js_rendering field.
	siteDescEnableJsRendering := siteFields[5].Descriptor()
	// site.DefaultEnableJsRendering holds the default value on creation for the enable_js_rendering field.
	site.DefaultEnableJsRendering = siteDescEnableJsRendering.Default.(bool)
	// siteDescCreatedAt is the schema descriptor for created_at field.
	siteDescCreatedAt := siteFields[6].Descriptor()
	// site.DefaultCreatedAt holds the default value on creation for the created_at field.
	site.DefaultCreatedAt = siteDescCreatedAt.Default.(func() time.Time)
	// siteDescUpdatedAt is the schema descriptor for updated_at field.
	siteDescUpdatedAt := siteFields[7].Descriptor()
	// site.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	site.DefaultUpdatedAt = siteDescUpdatedAt.Default.(func() time.Time)
	// site.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	site.UpdateDefaultUpdatedAt = siteDescUpdatedAt.UpdateDefault.(func() time.Time)
}
