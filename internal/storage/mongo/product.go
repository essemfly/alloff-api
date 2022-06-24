package mongo

import (
	"context"
	"log"
	"time"

	"github.com/lessbutter/alloff-api/config"
	"go.uber.org/zap"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/core/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type productRepo struct {
	col *mongo.Collection
}

type productMetaInfoRepo struct {
	col *mongo.Collection
}

func (repo *productRepo) Get(ID string) (*domain.ProductDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	productObjectId, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": productObjectId}
	var product *domain.ProductDAO
	if err := repo.col.FindOne(ctx, filter).Decode(&product); err != nil {
		return nil, err
	}
	return product, nil
}

func (repo *productRepo) GetByMetaID(metaID, exhibitionID, productGroupID string) (*domain.ProductDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	productObjectId, _ := primitive.ObjectIDFromHex(metaID)
	filter := bson.M{"productinfo._id": productObjectId, "isremoved": false}
	if exhibitionID != "" {
		filter["exhibitionid"] = exhibitionID
	}
	if productGroupID != "" {
		filter["productgroupid"] = productGroupID
	}
	var product *domain.ProductDAO
	if err := repo.col.FindOne(ctx, filter).Decode(&product); err != nil {
		return nil, err
	}
	return product, nil
}

func (repo *productRepo) ListByMetaID(metaID string) ([]*domain.ProductDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	productObjectId, _ := primitive.ObjectIDFromHex(metaID)
	filter := bson.M{"productinfo._id": productObjectId, "isremoved": false}
	cursor, err := repo.col.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var products []*domain.ProductDAO
	err = cursor.All(ctx, &products)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (repo *productRepo) List(offset, limit int, filter, sortingOptions interface{}) ([]*domain.ProductDAO, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	options := options.Find()
	options.SetSort(sortingOptions)
	options.SetSkip(int64(offset))
	options.SetLimit(int64(limit))

	totalCount, _ := repo.col.CountDocuments(ctx, filter)
	cursor, err := repo.col.Find(ctx, filter, options)
	if err != nil {
		return nil, 0, err
	}

	var products []*domain.ProductDAO
	err = cursor.All(ctx, &products)
	if err != nil {
		return nil, 0, err
	}

	return products, int(totalCount), nil
}

func (repo *productRepo) Aggregate(filter interface{}, pipelines []interface{}) ([]*domain.ProductDAO, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	totalCount, _ := repo.col.CountDocuments(ctx, filter)
	cursor, err := repo.col.Aggregate(ctx, pipelines)
	if err != nil {
		config.Logger.Error("error on aggregating products : ", zap.Error(err))
		return nil, 0, err
	}

	var products []*domain.ProductDAO
	err = cursor.All(ctx, &products)
	if err != nil {
		config.Logger.Error("error on cursor to products : ", zap.Error(err))
		return nil, 0, err
	}

	return products, int(totalCount), nil
}

func (repo *productRepo) Insert(product *domain.ProductDAO) (*domain.ProductDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	oid, err := repo.col.InsertOne(ctx, product)
	if err != nil {
		return nil, err
	}

	var oldProduct domain.ProductDAO
	filter := bson.M{"_id": oid.InsertedID}
	err = repo.col.FindOne(ctx, filter).Decode(&oldProduct)
	if err != nil {
		return nil, err
	}

	return &oldProduct, nil
}

func (repo *productRepo) Upsert(product *domain.ProductDAO) (*domain.ProductDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	opts := options.Update().SetUpsert(true)
	filter := bson.M{"_id": product.ID}
	if _, err := repo.col.UpdateOne(ctx, filter, bson.M{"$set": &product}, opts); err != nil {
		return nil, err
	}

	var updatedProduct *domain.ProductDAO
	if err := repo.col.FindOne(ctx, filter).Decode(&updatedProduct); err != nil {
		return nil, err
	}

	return updatedProduct, nil
}

func (repo *productRepo) ListDistinctBrands(alloffCategoryID string) ([]*domain.BrandDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	filter := bson.M{"removed": false, "alloffcategories.done": true}

	alloffCat, err := ioc.Repo.AlloffCategories.Get(alloffCategoryID)
	if err != nil {
		return nil, err
	}

	if alloffCat.Level == 1 {
		filter["alloffcategories.first._id"] = alloffCat.ID
	} else if alloffCat.Level == 2 {
		filter["alloffcategories.second._id"] = alloffCat.ID
	}

	rows, err := repo.col.Distinct(ctx, "brand", filter)
	if err != nil {
		return nil, err
	}

	brands := []*domain.BrandDAO{}

	for _, row := range rows {
		var brand *domain.BrandDAO
		data, err := bson.Marshal(row)
		if err != nil {
			log.Println("Err in marshaling")
		}

		err = bson.Unmarshal(data, &brand)
		if err != nil {
			log.Println("Err in unmarshaling")
		}
		brands = append(brands, brand)
	}

	return brands, nil
}

func (repo *productRepo) ListDistinctInfos(filter interface{}) (brands []*domain.BrandCountsData, cats []*domain.AlloffCategoryDAO, sizes []*domain.AlloffSizeDAO) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	brandRows, _ := repo.col.Distinct(ctx, "productinfo.brand", filter)
	for _, row := range brandRows {
		var brand *domain.BrandDAO
		data, _ := bson.Marshal(row)
		bson.Unmarshal(data, &brand)
		brands = append(brands, &domain.BrandCountsData{
			Brand:  brand,
			Counts: 1882,
		})
	}

	catRows, _ := repo.col.Distinct(ctx, "productinfo.alloffcategory.first", filter)
	for _, row := range catRows {
		var cat *domain.AlloffCategoryDAO
		data, _ := bson.Marshal(row)
		bson.Unmarshal(data, &cat)
		cats = append(cats, cat)
	}

	sizeRows, _ := repo.col.Distinct(ctx, "productinfo.inventory.alloffsizes", filter)
	for _, row := range sizeRows {
		var size *domain.AlloffSizeDAO
		data, _ := bson.Marshal(row)
		bson.Unmarshal(data, &size)
		sizes = append(sizes, size)
	}

	return
}

func (repo *productRepo) ListInfos(filter interface{}) (brands []*domain.BrandCountsData, cats []*domain.AlloffCategoryDAO, sizes []*domain.AlloffSizeDAO) {
	ctx, cancel := context.WithTimeout(context.Background(), 70*time.Second)
	defer cancel()

	cursor, err := repo.col.Find(ctx, filter)
	if err != nil {
		return nil, nil, nil
	}

	var pds []*domain.ProductDAO
	err = cursor.All(ctx, &pds)
	if err != nil {
		return nil, nil, nil
	}

	for _, pd := range pds {
		// ******* mapping brands *******
		brandExists := false
		existAt := -1
		for idx, bd := range brands {
			// 이미 있는 브랜드면 그게 brands 몇번쨰인지 알려준다.
			if pd.ProductInfo.Brand.ID == bd.Brand.ID {
				brandExists = true
				existAt = idx
			}
		}
		// 아직 입력안된 브랜드면 brands 에 입력한다
		if !brandExists {
			brands = append(brands, &domain.BrandCountsData{
				Brand:  pd.ProductInfo.Brand,
				Counts: 1,
			})
		} else {
			// 이미 입력된 브랜드면 counts 만 +1을 한다.
			brands[existAt].Counts = brands[existAt].Counts + 1
		}

		// ******* mapping cats *******
		catExists := false
		for _, cat := range cats {
			if pd.ProductInfo.AlloffCategory == nil || pd.ProductInfo.AlloffCategory.First == nil {
				// nilcase 인경우 true 로 처리하여 cats에 빈 category가 들어가는걸 막는다.
				catExists = true
				continue
			}
			if pd.ProductInfo.AlloffCategory.First.ID == cat.ID {
				catExists = true
			}
		}
		if !catExists {
			cats = append(cats, pd.ProductInfo.AlloffCategory.First)
		}

		// ******* mapping sizes *******
		alloffSizes := []*domain.AlloffSizeDAO{}
		for _, inv := range pd.ProductInfo.Inventory {
			for _, alloffSize := range inv.AlloffSizes {
				alloffSizes = append(alloffSizes, alloffSize)
			}
		}
		for _, target := range alloffSizes {
			isThere := false
			for _, ns := range sizes {
				if ns.ID == target.ID {
					isThere = true
				}
			}
			if !isThere {
				sizes = append(sizes, target)
			}
		}
	}
	return
}

func MongoProductsRepo(conn *MongoDB) repository.ProductsRepository {
	return &productRepo{
		col: conn.productCol,
	}
}

func (repo *productMetaInfoRepo) Get(ID string) (*domain.ProductMetaInfoDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	productObjectId, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": productObjectId}
	var product *domain.ProductMetaInfoDAO
	if err := repo.col.FindOne(ctx, filter).Decode(&product); err != nil {
		return nil, err
	}
	return product, nil
}

func (repo *productMetaInfoRepo) GetByProductID(brandKeyname string, productID string) (*domain.ProductMetaInfoDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	var oldProduct domain.ProductMetaInfoDAO
	filter := bson.M{"productid": productID, "brand.keyname": brandKeyname}
	err := repo.col.FindOne(ctx, filter).Decode(&oldProduct)
	if err != nil {
		return nil, err
	}

	return &oldProduct, nil
}

func (repo *productMetaInfoRepo) List(offset, limit int, filter, sortingOptions interface{}) ([]*domain.ProductMetaInfoDAO, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	options := options.Find()
	options.SetSort(sortingOptions)
	options.SetSkip(int64(offset))
	options.SetLimit(int64(limit))

	totalCount, _ := repo.col.CountDocuments(ctx, filter)
	cursor, err := repo.col.Find(ctx, filter, options)
	if err != nil {
		return nil, 0, err
	}

	var productInfos []*domain.ProductMetaInfoDAO
	err = cursor.All(ctx, &productInfos)
	if err != nil {
		return nil, 0, err
	}

	return productInfos, int(totalCount), nil
}

func (repo *productMetaInfoRepo) Insert(pd *domain.ProductMetaInfoDAO) (*domain.ProductMetaInfoDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	var oldProduct domain.ProductMetaInfoDAO
	oid, err := repo.col.InsertOne(ctx, pd)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": oid.InsertedID}
	err = repo.col.FindOne(ctx, filter).Decode(&oldProduct)
	if err != nil {
		return nil, err
	}

	return &oldProduct, nil
}

func (repo *productMetaInfoRepo) Upsert(product *domain.ProductMetaInfoDAO) (*domain.ProductMetaInfoDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	opts := options.Update().SetUpsert(true)
	filter := bson.M{"_id": product.ID}
	if _, err := repo.col.UpdateOne(ctx, filter, bson.M{"$set": &product}, opts); err != nil {
		return nil, err
	}

	var updatedProduct *domain.ProductMetaInfoDAO
	if err := repo.col.FindOne(ctx, filter).Decode(&updatedProduct); err != nil {
		return nil, err
	}

	return updatedProduct, nil
}

func (repo *productMetaInfoRepo) CountNewProducts(brandModules []string, lastUpdatedDate time.Time) int {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"isremoved": false, "source.crawlmodulename": bson.M{"$in": brandModules}, "created": bson.M{
		"$gte": primitive.NewDateTimeFromTime(lastUpdatedDate),
	}}
	newProducts, err := repo.col.CountDocuments(ctx, filter)
	if err != nil {
		config.Logger.Error("Find num of new crawled products error", zap.Error(err))
	}

	return int(newProducts)
}

func (repo *productMetaInfoRepo) MakeOutdatedProducts(brandModules []string, lastUpdatedDate time.Time) int {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	outProducts, err := repo.col.UpdateMany(
		ctx,
		bson.M{
			"isremoved":              false,
			"source.crawlmodulename": bson.M{"$in": brandModules},
			"updated": bson.M{
				"$lte": time.Now().Add(-24 * 7 * time.Hour),
			},
		}, bson.M{"$set": bson.M{"isremoved": true}})
	if err != nil {
		log.Println("Find num of outdated products error", err)
	}

	return int(outProducts.ModifiedCount)
}

func MongoProductMetaInfosRepo(conn *MongoDB) repository.ProductMetaInfoRepository {
	return &productMetaInfoRepo{
		col: conn.productMetaInfoCol,
	}
}
