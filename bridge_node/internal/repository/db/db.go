package db

import (
	"context"
	"database/sql"

	"github.com/QwaQ-dev/stacks-base-bridge/internal/models"
	_ "github.com/lib/pq"
)

type DB struct {
	conn *sql.DB
}

func NewFromSQL(conn *sql.DB) (*DB, error) {
	if conn == nil {
		return nil, sql.ErrConnDone
	}
	return &DB{conn: conn}, nil
}

func (db *DB) Close() error {
	return db.conn.Close()
}

// Bridge Events
func (db *DB) SaveBridgeEvent(ctx context.Context, sourceChain, targetChain, txHash string, logIndex int,
	blockNumber int64, blockHash, sender, receiver string, amount, nonce string, status string) error {

	_, err := db.conn.ExecContext(ctx, `
	INSERT INTO bridge_events
	(source_chain,target_chain,tx_hash,log_index,block_number,block_hash,sender,receiver,amount,nonce,status,created_at,updated_at)
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,NOW(),NOW())
	`, sourceChain, targetChain, txHash, logIndex, blockNumber, blockHash, sender, receiver, amount, nonce, status)
	return err
}

func (db *DB) UpdateBridgeEventStatus(ctx context.Context, eventID int64, status string) error {
	_, err := db.conn.ExecContext(ctx, `
		UPDATE bridge_events SET status=$1, updated_at=NOW() WHERE id=$2
	`, status, eventID)
	return err
}

// GetPendingEvents return events with 'detected'status for sourceChain.
func (db *DB) GetPendingEvents(ctx context.Context, sourceChain string) ([]models.BridgeEvent, error) {
	rows, err := db.conn.QueryContext(ctx, `
		SELECT id, source_chain, target_chain, tx_hash, log_index, block_number,
		       sender, receiver, amount, nonce, status
		FROM bridge_events
		WHERE source_chain=$1 AND status='detected'
		ORDER BY block_number ASC
		LIMIT 100
	`, sourceChain)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []models.BridgeEvent
	for rows.Next() {
		var e models.BridgeEvent
		if err := rows.Scan(&e.ID, &e.SourceChain, &e.TargetChain, &e.TxHash, &e.LogIndex,
			&e.BlockNumber, &e.Sender, &e.Receiver, &e.Amount, &e.Nonce, &e.Status); err != nil {
			return nil, err
		}
		events = append(events, e)
	}
	return events, rows.Err()
}

// ─── Relay Transactions ───────────────────────────────────────────────────────

// ClaimRelay атомарно создаёт запись в relay_transactions.
// Если другой нод уже создал запись для этого event_id — возвращает (false, nil).
// Это гарантирует что только один нод выполнит relay.
func (db *DB) ClaimRelay(ctx context.Context, eventID int64, targetChain string) (claimed bool, err error) {
	res, err := db.conn.ExecContext(ctx, `
		INSERT INTO relay_transactions (event_id, target_chain, status, created_at, updated_at)
		VALUES ($1, $2, 'pending', NOW(), NOW())
		ON CONFLICT (event_id) DO NOTHING
	`, eventID, targetChain)
	if err != nil {
		return false, err
	}
	rows, err := res.RowsAffected()
	return rows == 1, err
}

// UpdateRelayStatus обновляет статус и tx_hash relay транзакции.
func (db *DB) UpdateRelayStatus(ctx context.Context, eventID int64, status, txHash string) error {
	_, err := db.conn.ExecContext(ctx, `
		UPDATE relay_transactions
		SET status=$1, tx_hash=$2, updated_at=NOW()
		WHERE event_id=$3
	`, status, txHash, eventID)
	return err
}

// ─── Indexer State ────────────────────────────────────────────────────────────

func (db *DB) GetLastBlock(ctx context.Context, chain string) (int64, error) {
	var block int64
	err := db.conn.QueryRowContext(ctx, `
		SELECT last_block FROM indexer_state WHERE chain=$1
	`, chain).Scan(&block)
	if err == sql.ErrNoRows {
		return 0, nil
	}
	return block, err
}

func (db *DB) SetLastBlock(ctx context.Context, chain string, block int64) error {
	_, err := db.conn.ExecContext(ctx, `
		INSERT INTO indexer_state(chain, last_block, updated_at)
		VALUES ($1, $2, NOW())
		ON CONFLICT (chain) DO UPDATE SET last_block = EXCLUDED.last_block, updated_at=NOW()
	`, chain, block)
	return err
}
