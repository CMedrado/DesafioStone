package transfer

import (
	"context"
	"github.com/CMedrado/DesafioStone/domain"
)

func (a *Storage) SaveTransfer(transfer domain.Transfer) error {
	statement := `INSERT INTO transfers(id, origin_account_id, destination_account_id, amount, created_at)
				  VALUES ($1, $2, $3, $4, $5)`
	comand, err := a.pool.Exec(context.Background(), statement, transfer.ID, transfer.OriginAccountID, transfer.DestinationAccountID, transfer.Amount, transfer.CreatedAt)
	if comand.RowsAffected() > 0 {
		return nil
	}
	if err != nil {
		return err
	}
	return nil
}
