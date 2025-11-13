package bootstrap

import (
	_ "gozero-sso-service/internal/dataaccess/adapter/repository/user"
	_ "gozero-sso-service/internal/logic/adapter/service/auth"
	_ "gozero-sso-service/internal/logic/adapter/service/user"
)
