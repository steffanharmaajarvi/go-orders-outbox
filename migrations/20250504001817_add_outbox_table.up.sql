CREATE TABLE outbox (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    aggregate_type TEXT NOT NULL,         -- тип сущности (например, 'order')
    aggregate_id UUID NOT NULL,           -- id этой сущности (например, order_id)
    type TEXT NOT NULL,                   -- тип события (например, 'OrderCreated')
    payload JSONB NOT NULL,               -- само событие
    occurred_at TIMESTAMP WITH TIME ZONE DEFAULT now(),  -- когда событие создано
    sent_at TIMESTAMP WITH TIME ZONE      -- когда отправлено (NULL = не отправлено)
);