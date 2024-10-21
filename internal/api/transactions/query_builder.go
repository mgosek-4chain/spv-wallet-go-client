package transactions

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/bitcoin-sv/spv-wallet/models"
	"github.com/bitcoin-sv/spv-wallet/models/filter"
)

type QueryBuilderOption func(*QueryBuilder)

func WithQueryParamsFilter(q filter.QueryParams) QueryBuilderOption {
	return func(qb *QueryBuilder) {
		qb.builders = append(qb.builders, &QueryParamsFilterQueryBuilder{q})
	}
}

func WithMetadataFilter(m Metadata) QueryBuilderOption {
	return func(qb *QueryBuilder) {
		qb.builders = append(qb.builders, &MetadataFilterQueryBuilder{MaxDepth: DefaultMaxDepth, Metadata: m})
	}
}

func WithTransactionFilter(tf filter.TransactionFilter) QueryBuilderOption {
	return func(qb *QueryBuilder) {
		qb.builders = append(qb.builders, &TransactionFilterQueryBuilder{
			TransactionFilter:       tf,
			ModelFilterQueryBuilder: ModelFilterQueryBuilder{ModelFilter: tf.ModelFilter},
		})
	}
}

func WithFilterQueryBuilder(b FilterQueryBuilder) QueryBuilderOption {
	return func(qb *QueryBuilder) {
		if b != nil {
			qb.builders = append(qb.builders, b)
		}
	}
}

type FilterQueryBuilder interface {
	Build() (url.Values, error)
	String() string
}

type QueryBuilder struct {
	builders []FilterQueryBuilder
}

func (q *QueryBuilder) Build() (url.Values, error) {
	params := NewExtendedURLValues()
	for _, b := range q.builders {
		bparams, err := b.Build()
		if err != nil {
			return nil, NewErrFilterQueryBuilder(b.String()).Wrap(err)
		}
		if len(bparams) > 0 {
			params.Append(bparams)
		}
	}
	return params.Values, nil
}

func NewQueryBuilder(opts ...QueryBuilderOption) *QueryBuilder {
	var qb QueryBuilder
	for _, o := range opts {
		o(&qb)
	}
	return &qb
}

func NewErrFilterQueryBuilder(s string) models.SPVError {
	err := models.SPVError{
		Message:    fmt.Sprintf("failed to build transactions query parameters - filter query builder: %s", s),
		StatusCode: http.StatusInternalServerError,
		Code:       "filter-query-builder-transactions-parameters-build-failure",
	}
	return err
}
