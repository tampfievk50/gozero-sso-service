package vtmo

import (
	"fmt"
	"gorm.io/gorm"
	"gozero-sso-service/internal/dataaccess/model"
)

func (r *repository) GetMO8x79(tx *gorm.DB, id uint) ([]model.VT8x79, error) {
	var mo8x97 []model.VT8x79
	sql := "SELECT id, operator, service_id, user_id, DATE_FORMAT(send_time,'%Y%m%d%H%i%s') AS submit_date, " +
		"DATE_FORMAT(send_time,'%Y%m%d%H%i%s') AS done_date, send_time, deliver_time, keyword, info, content_type, " +
		"message_type, number_message, request_id, thirdparty, process_result, money, autotimestamp, loggingTime " +
		fmt.Sprintf("FROM %s ", r.config.Job.VT8x79Table) +
		"WHERE operator='viettel' AND service_id LIKE '8%79' " +
		"AND user_id NOT IN ('84904061617','84936180038','841223388311','841213000005','841262412986','84912816396','841253134221','841236022442','841273070757','84945366105','841696941026','841663153143','841676219770','841656538118','841689468650','84989068604','84979756683','84923000328','841998378533','84988087935','84906071118','84946651563','84904858285','84983801346') " +
		"AND keyword IS NOT NULL " +
		"AND keyword <> '' AND id > ? ORDER BY id ASC LIMIT 1000"
	err := tx.Raw(sql, id).Scan(&mo8x97).Error
	if err != nil {
		return nil, err
	}
	return mo8x97, nil
}
