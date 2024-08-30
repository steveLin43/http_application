package dao

import "http_application/internal/model"

func (d *Dao) GetAuth(appkey, appSecret string) (model.Auth, error) {
	auth := model.Auth{AppKey: appkey, AppSecret: appSecret}
	return auth.Get(d.engine)
}
