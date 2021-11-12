package cart

import (
	"context"
	"encoding/json"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
	"gorm.io/gorm"
)

// Repository is the interface that describes a cart repository
// object.
type Repository interface {
	SaveCartItem(ctx context.Context, item *CartItem) error
}

// CartRepo is the default implementation for Repository interface.
type CartRepo struct {
	orm    *gorm.DB
	tracer opentracing.Tracer
}

// NewRepository returns a new cart repository object.
func NewRepository(orm *gorm.DB, tracer opentracing.Tracer) *CartRepo {
	return &CartRepo{
		orm:    orm,
		tracer: tracer,
	}
}

func (r *CartRepo) setPostGresComponentTags(span opentracing.Span, tableName string) {
	ext.DBInstance.Set(span, tableName)
	ext.DBType.Set(span, "postgres")
	ext.SpanKindRPCClient.Set(span)
}

// SaveCartItem saves a cart item to the database.
func (r *CartRepo) SaveCartItem(ctx context.Context, item *CartItem) error {
	span := r.tracer.StartSpan("SaveCartItem", opentracing.ChildOf(opentracing.SpanFromContext(ctx).Context()))
	defer span.Finish()
	r.setPostGresComponentTags(span, "cart")
	span.LogFields(log.Object("param.item", toJSON(item)))
	item.TimeAdded = time.Now()
	item.LastUpdated = time.Now()
	tx := r.orm.Save(item)
	if tx.Error != nil {
		ext.Error.Set(span, true)
		span.LogFields(log.Error(tx.Error), log.Event("gorm.db.Save"))
		return tx.Error
	}
	return nil
}

func toJSON(i interface{}) string {
	iJSON, _ := json.Marshal(i)
	return string(iJSON)
}
