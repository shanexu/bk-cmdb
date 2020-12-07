package y3_9_202012071427

import (
	"context"

	"configcenter/src/common"
	"configcenter/src/common/mapstr"
	"configcenter/src/common/metadata"
	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/storage/dal"
)

func changeBindInfoSubAttr(ctx context.Context, db dal.RDB, conf *upgrader.Config) error {
	cond := mapstr.MapStr{}
	cond.Set(common.BKObjIDField, "process")
	cond.Set(common.BKPropertyIDField, "bind_info")
	procBindIPAttr := &Attribute{}
	err := db.Table(common.BKTableNameObjAttDes).Find(cond).One(ctx, procBindIPAttr)
	if err != nil {
		return err
	}
	subAttributes := append([]metadata.SubAttribute{{
		PropertyID:    "name",
		PropertyName:  "名字",
		Placeholder:   "名字",
		IsEditable:    true,
		IsRequired:    false,
		PropertyType:  common.FieldTypeSingleChar,
		Option: common.FieldTypeStrictCharRegexp,
		PropertyGroup: common.BKProcBindInfo,
	}}, procBindIPAttr.Option...)

	err = db.Table(common.BKTableNameObjAttDes).Update(ctx, cond, mapstr.MapStr{
		common.BKOptionField: subAttributes,
	})

	return err
}

type Attribute struct {
	ObjectID string                  `field:"bk_obj_id" json:"bk_obj_id" bson:"bk_obj_id"`
	Option   []metadata.SubAttribute `field:"option" json:"option" bson:"option"`
}
