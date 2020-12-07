package y3_9_202012071427

import (
	"context"

	"configcenter/src/common/blog"
	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/storage/dal"
)

func init() {
	upgrader.RegistUpgrader("y3.9.202012071427", upgrade)
}

func upgrade(ctx context.Context, db dal.RDB, conf *upgrader.Config) (err error) {
	blog.Infof("start execute y3.9.202012071427")

	err = changeBindInfoSubAttr(ctx, db, conf)
	if err != nil {
		blog.Errorf("[upgrade y3.9.202012071427] changeBindInfoSubAttr failed, error  %s", err.Error())
		return err
	}

	return nil
}
