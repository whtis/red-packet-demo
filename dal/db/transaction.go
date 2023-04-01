package db

import (
	"context"
	"ginDemo/model"
	"gorm.io/gorm"
)

func UpdateSendAndCreateReceiveRecordTx(ctx context.Context, sRecord *model.RpSendRecord, rRecord *model.RpReceiveRecord) error {
	return Rdb.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		if err := tx.WithContext(ctx).Table("rp_receive_record").Create(rRecord).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}
		if err := tx.WithContext(ctx).Table("rp_send_record").Where("rp_id = ?", rRecord.RpId).Update("receive_amount", sRecord.ReceiveAmount).Error; err != nil {
			return err
		}

		// 返回 nil 提交事务
		return nil
	})

}
