package seeder

import (
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func LoadAlloffCats() []*domain.AlloffCategoryDAO {
	tshirtOID, _ := primitive.ObjectIDFromHex("60feb9f98adeef23689cbff5")
	blouseOID, _ := primitive.ObjectIDFromHex("60feb9f98adeef23689cbff6")
	blouse2ndOID, _ := primitive.ObjectIDFromHex("60feb9f98adeef23689cbff7")
	knitwareOID, _ := primitive.ObjectIDFromHex("60feb9f98adeef23689cbff8")
	knit2ndOID, _ := primitive.ObjectIDFromHex("60feb9f98adeef23689cbff9")
	onepieceOID, _ := primitive.ObjectIDFromHex("60feb9f98adeef23689cbffa")
	outerOID, _ := primitive.ObjectIDFromHex("60feb9f98adeef23689cbffc")
	skirtOID, _ := primitive.ObjectIDFromHex("60feb9f98adeef23689cbffb")
	coat2ndOID, _ := primitive.ObjectIDFromHex("60feb9f98adeef23689cc001")
	jumper2ndOID, _ := primitive.ObjectIDFromHex("60feb9f98adeef23689cc002")
	pantsOID, _ := primitive.ObjectIDFromHex("60feb9f98adeef23689cc003")
	pants2ndOID, _ := primitive.ObjectIDFromHex("60feb9f98adeef23689cc004")
	denim2ndOID, _ := primitive.ObjectIDFromHex("60feb9fa8adeef23689cc008")
	jacket2ndOID, _ := primitive.ObjectIDFromHex("60feb9f98adeef23689cc005")
	shirt2ndOID, _ := primitive.ObjectIDFromHex("60feb9fa8adeef23689cc006")
	vest2ndOID, _ := primitive.ObjectIDFromHex("60feb9fa8adeef23689cc007")
	cardigan2ndOID, _ := primitive.ObjectIDFromHex("60feb9fa8adeef23689cc009")
	accessoryOID, _ := primitive.ObjectIDFromHex("60feb9fa8adeef23689cc00d")
	trench2ndOID, _ := primitive.ObjectIDFromHex("60feb9fa8adeef23689cc010")
	underwearOID, _ := primitive.ObjectIDFromHex("60feb9fa8adeef23689cc011")
	shoes2ndOID, _ := primitive.ObjectIDFromHex("60feb9fa8adeef23689cc013")
	bag2ndOID, _ := primitive.ObjectIDFromHex("60feb9fa8adeef23689cc014")
	padding2ndOID, _ := primitive.ObjectIDFromHex("6103b81dba5c475596777312")
	notshowOID, _ := primitive.ObjectIDFromHex("6108a77e4228e1fff07db971")

	return []*domain.AlloffCategoryDAO{
		{
			ID:           tshirtOID,
			Name:         "티셔츠",
			KeyName:      "1_TSHIRT",
			Level:        1,
			ParentId:     primitive.NilObjectID,
			CategoryType: "NORMAL",
			ImgURL:       "https://alloff.s3.ap-northeast-2.amazonaws.com/categories/1_TSHIRT.png",
		},
		/* 2 */
		{
			ID:           blouseOID,
			Name:         "셔츠/블라우스",
			KeyName:      "1_SHIRT/BLOUSE",
			Level:        1,
			ParentId:     primitive.NilObjectID,
			CategoryType: "NORMAL",
			ImgURL:       "https://alloff.s3.ap-northeast-2.amazonaws.com/categories/BLOUSE.png",
		},
		/* 3 */
		{
			ID:           blouse2ndOID,
			Name:         "블라우스",
			KeyName:      "2_BLOUSE",
			Level:        2,
			ParentId:     blouseOID,
			CategoryType: "NORMAL",
		},
		/* 4 */
		{
			ID:           knitwareOID,
			Name:         "니트웨어",
			KeyName:      "1_KNITWARE",
			Level:        1,
			ParentId:     primitive.NilObjectID,
			CategoryType: "NORMAL",
			ImgURL:       "https://alloff.s3.ap-northeast-2.amazonaws.com/categories/1_KNITWARE.png",
		},
		/* 5 */
		{
			ID:           knit2ndOID,
			Name:         "니트/스웨터",
			KeyName:      "2_KNIT/SWEATER",
			Level:        2,
			ParentId:     knitwareOID,
			CategoryType: "NORMAL",
		},
		/* 6 */
		{
			ID:           onepieceOID,
			Name:         "원피스",
			KeyName:      "1_ONEPIECE",
			Level:        1,
			ParentId:     primitive.NilObjectID,
			CategoryType: "NORMAL",
			ImgURL:       "https://alloff.s3.ap-northeast-2.amazonaws.com/categories/1_ONEPIECE.png",
		},
		/* 7 */
		{
			ID:           skirtOID,
			Name:         "스커트",
			KeyName:      "1_SKIRT",
			Level:        1,
			ParentId:     primitive.NilObjectID,
			CategoryType: "NORMAL",
			ImgURL:       "https://alloff.s3.ap-northeast-2.amazonaws.com/categories/1_SKIRT.png",
		},
		/* 8 */
		{
			ID:           outerOID,
			Name:         "아우터",
			KeyName:      "1_OUTER",
			Level:        1,
			ParentId:     primitive.NilObjectID,
			CategoryType: "NORMAL",
			ImgURL:       "https://alloff.s3.ap-northeast-2.amazonaws.com/categories/1_OUTER.png",
		},
		/* 9 */
		{
			ID:           coat2ndOID,
			Name:         "코트",
			KeyName:      "2_COAT",
			Level:        2,
			ParentId:     outerOID,
			CategoryType: "NORMAL",
		},
		/* 10 */
		{
			ID:           jumper2ndOID,
			Name:         "점퍼",
			KeyName:      "2_JUMPER",
			Level:        2,
			ParentId:     outerOID,
			CategoryType: "NORMAL",
		},
		/* 11 */
		{
			ID:           pantsOID,
			Name:         "팬츠/데님",
			KeyName:      "1_PANTS/DENIM",
			Level:        1,
			ParentId:     primitive.NilObjectID,
			CategoryType: "NORMAL",
			ImgURL:       "https://alloff.s3.ap-northeast-2.amazonaws.com/categories/1_PANTS.DENIM.png",
		},
		/* 12 */
		{
			ID:           pants2ndOID,
			Name:         "팬츠",
			KeyName:      "2_PANTS",
			Level:        2,
			ParentId:     pantsOID,
			CategoryType: "NORMAL",
		},
		/* 13 */
		{
			ID:           jacket2ndOID,
			Name:         "자켓",
			KeyName:      "2_JACKET",
			Level:        2,
			ParentId:     outerOID,
			CategoryType: "NORMAL",
		},
		/* 14 */
		{
			ID:           shirt2ndOID,
			Name:         "셔츠",
			KeyName:      "2_SHIRT",
			Level:        2,
			ParentId:     blouseOID,
			CategoryType: "NORMAL",
		},
		/* 15 */
		{
			ID:           vest2ndOID,
			Name:         "베스트",
			KeyName:      "2_VEST",
			Level:        2,
			ParentId:     outerOID,
			CategoryType: "NORMAL",
		},
		/* 16 */
		{
			ID:           denim2ndOID,
			Name:         "데님",
			KeyName:      "2_DENIM",
			Level:        2,
			ParentId:     pantsOID,
			CategoryType: "NORMAL",
		},
		/* 17 */
		{
			ID:           cardigan2ndOID,
			Name:         "가디건",
			KeyName:      "2_CARDIGAN",
			Level:        2,
			ParentId:     knitwareOID,
			CategoryType: "NORMAL",
		},
		/* 18 */
		{
			ID:           accessoryOID,
			Name:         "패션잡화",
			KeyName:      "1_ACCESSORY",
			Level:        1,
			ParentId:     primitive.NilObjectID,
			CategoryType: "NORMAL",
			ImgURL:       "https://alloff.s3.ap-northeast-2.amazonaws.com/categories/1_ACCESSORY.png",
		},
		/* 19 */
		{
			ID:           trench2ndOID,
			Name:         "트렌치",
			KeyName:      "2_TRENCH",
			Level:        2,
			ParentId:     outerOID,
			CategoryType: "NORMAL",
		},
		/* 20 */
		{
			ID:           underwearOID,
			Name:         "라운지/언더웨어",
			KeyName:      "1_LOUNGE/UNDERWARE",
			Level:        1,
			ParentId:     primitive.NilObjectID,
			CategoryType: "NORMAL",
			ImgURL:       "https://alloff.s3.ap-northeast-2.amazonaws.com/categories/1_LOUNGE.UNDERWARE.png",
		},
		/* 21 */
		{
			ID:           shoes2ndOID,
			Name:         "신발",
			KeyName:      "2_SHOES",
			Level:        2,
			ParentId:     accessoryOID,
			CategoryType: "NORMAL",
		},
		/* 22 */
		{
			ID:           bag2ndOID,
			Name:         "가방/잡화",
			KeyName:      "2_BAG_ETC",
			Level:        2,
			ParentId:     accessoryOID,
			CategoryType: "NORMAL",
		},
		/* 23 */
		{
			ID:           padding2ndOID,
			Name:         "패딩",
			KeyName:      "2_PADDING",
			Level:        2,
			ParentId:     outerOID,
			CategoryType: "NORMAL",
		},
		/* 24 */
		{
			ID:           notshowOID,
			Name:         "❌미노출❌",
			CategoryType: "DO_NOT_SHOW",
			Level:        1,
			ParentId:     primitive.NilObjectID,
			KeyName:      "1_DO_NOT_SHOW",
		},
	}
}
