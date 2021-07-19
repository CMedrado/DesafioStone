package transfer

import "context"

func (a *Storage) SaveTransfer(transfer Transfer) error {
	statement := `INSERT INTO accounts(id, origin_account_id, destination_account_id, amount, created_at)
				  VALUES ($1, $2, $3, $4, $5)`
	_, err := a.pool.Exec(context.Background(), statement, transfer.ID, transfer.OriginAccountID, transfer.DestinationAccountID, transfer.Amount, transfer.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}