package outbox

import (
	"context"
	"database/sql"
	"github.com/rs/zerolog/log"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"orders.go/m/internal/models"
)

type OutboxRepo struct {
	db *sql.DB
}

func NewOutboxRepo(db *sql.DB) *OutboxRepo {
	return &OutboxRepo{db: db}
}

func (repo *OutboxRepo) Save(ctx context.Context, exec boil.ContextExecutor, outbox *models.Outbox) error {
	var executor boil.ContextExecutor
	if exec == nil {
		executor = repo.db
	} else {
		executor = exec
	}

	err := outbox.Upsert(ctx, executor, true, []string{"id"}, boil.Whitelist("sent_at"), boil.Whitelist("id", "aggregate_type", "aggregate_id", "type", "payload", "sent_at"))
	if err != nil {
		log.Err(err).Msg("failed to upsert outbox")

		return err
	}

	return nil
}

func (repo *OutboxRepo) FetchUnsent(ctx context.Context) ([]*models.Outbox, error) {
	return models.Outboxes(
		models.OutboxWhere.SentAt.IsNull(),
	).All(ctx, repo.db)
}
