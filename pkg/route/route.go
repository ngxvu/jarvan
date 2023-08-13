package route

import (
	"github.com/gin-contrib/pprof"
	"github.com/hibiken/asynq"
	"gitlab.com/merakilab9/j4/conf"
	handlerJ4 "gitlab.com/merakilab9/j4/pkg/handler"
	"gitlab.com/merakilab9/j4/pkg/repo/pg"
	serviceJ4 "gitlab.com/merakilab9/j4/pkg/service"
	"gitlab.com/merakilab9/j4/pkg/utils"
	"gitlab.com/merakilab9/meracore/ginext"
	"gitlab.com/merakilab9/meracore/service"
	"log"
)

type Service struct {
	*service.BaseApp
}

func NewService() *Service {
	s := &Service{
		service.NewApp("jarvanService", "v1.0"),
	}

	db := s.GetDB()

	if !conf.LoadEnv().DbDebugEnable {
		db = db.Debug()
	}

	repoPG := pg.NewPGRepo(db)

	cateService := serviceJ4.NewCateService(repoPG)
	cateHandler := handlerJ4.NewCateHandlers(cateService)
	cateHandlersQueue := handlerJ4.NewSendAPIToQueueHandlers(cateService)

	shopIdService := serviceJ4.NewShopIdService(repoPG)
	shopIdHandler := handlerJ4.NewShopIdHandlers(shopIdService)

	itemService := serviceJ4.NewItemService(repoPG)
	itemHandler := handlerJ4.NewItemHandlers(itemService)

	migrateHandler := handlerJ4.NewMigrationHandler(db)
	pprof.Register(s.Router)

	v1Api := s.Router.Group("/api/v1")
	v1Api.POST("/internal/migrate", migrateHandler.Migrate)

	//v1Api.GET("/shopee/crawl-cate-url", ginext.WrapHandler(cateHandler.CrawlCate))

	v1Api.GET("/shopee/get-cate-url", ginext.WrapHandler(cateHandler.GetUrlCate))

	v1Api.GET("/shopee/get-shopid-url", ginext.WrapHandler(shopIdHandler.GetUrlShopId))

	v1Api.GET("/shopee/get-item-url", ginext.WrapHandler(itemHandler.GetUrlItem))

	client := asynq.NewClient(asynq.RedisClientOpt{Addr: utils.RedisAddr})
	defer client.Close()
	///=================== MssBroker ===================///

	task, err := cateHandlersQueue.NewAPIDeliveryTask()
	if err != nil {
		log.Fatalf("could not create tasks: %v", err)
	}

	info, err := client.Enqueue(task)
	if err != nil {
		log.Fatalf("could not enqueue tasks: %v", err)
	}
	log.Printf("enqueued tasks: id=%s queue=%s", info.ID, info.Queue)

	return s
}
