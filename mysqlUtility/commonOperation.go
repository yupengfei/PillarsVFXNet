package mysqlUtility

import (
	"PillarsPhenomVFXWeb/pillarsLog"
	"database/sql"
)

func TransactionOperation(tx *sql.Tx, err error) (bool, error) {
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return false, err
	}
	err = tx.Commit()
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		err = tx.Rollback()
		if err != nil {
			pillarsLog.PillarsLogger.Print(err.Error())
		}
		return false, err
	}
	return true, err
}
