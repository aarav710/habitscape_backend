package main

import (
	"context"
	"database/sql"
	"fmt"
	adminRepo "habitscape/backend/admin/repo"
	commentRepo "habitscape/backend/comment/repo"
	"habitscape/backend/ent"
	habitRepo "habitscape/backend/habit/repo"
	invitationRepo "habitscape/backend/invitation/repo"
	postRepo "habitscape/backend/post/repo"
	userRepo "habitscape/backend/user/repo"

	adminService "habitscape/backend/admin/service"
	commentService "habitscape/backend/comment/service"
	habitService "habitscape/backend/habit/service"
	invitationService "habitscape/backend/invitation/service"
	postService "habitscape/backend/post/service"
	userService "habitscape/backend/user/service"

	adminController "habitscape/backend/admin/controller"
	commentController "habitscape/backend/comment/controller"
	habitController "habitscape/backend/habit/controller"
	invitationController "habitscape/backend/invitation/controller"
	postController "habitscape/backend/post/controller"
	userController "habitscape/backend/user/controller"

	"habitscape/backend/cache"
	"habitscape/backend/cache/sessions"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	gorillaSessions "github.com/gorilla/sessions"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/markbates/goth/gothic"
)

func Open(databaseUrl string) *ent.Client {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		fmt.Println(err)
	} 

	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv))
}

func main() {
	client := Open("postgresql://aaravjain:@localhost:5432/habitscape")
  ctx := context.Background()
  if err := client.Schema.Create(ctx); err != nil {
    fmt.Println(err)
  } 
 
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
  })

  cacheService := cache.ProvideCache(redisClient)
	
  sessionsService, err := sessions.NewSessionsService(ctx, redisClient)
	if err != nil {
		fmt.Println(err)
	}

	sessionsService.Store.Options(gorillaSessions.Options{
    MaxAge: 86400 * 60,
		HttpOnly: true,
	})

	gothic.Store = sessionsService.Store
  
	router := gin.Default()

	// repo init 
	commentRepo := commentRepo.ProvideCommentRepo(client, ctx)
  habitRepo := habitRepo.ProvideHabitRepo(client, ctx)
	invitationRepo := invitationRepo.ProvideInvitationRepo(client, ctx)
	postRepo := postRepo.ProvidePostRepo(client, ctx)
	userRepo := userRepo.ProvideUserRepo(client, ctx)
	adminRepo := adminRepo.ProvideAdminRepo(client, ctx)

	//service init
  commentService := commentService.ProvideCommentService(commentRepo)
	habitService := habitService.ProvideHabitService(habitRepo)
	invitationService := invitationService.ProvideInvitationService(invitationRepo)
	postService := postService.ProvidePostService(postRepo)
	userService := userService.ProvideUserService(userRepo)
	adminService := adminService.ProvideAdminService(adminRepo)

	//controller init
  commentController.ProvideCommentController(router, commentService)
	habitController.ProvideHabitController(router, habitService)
	invitationController.ProvideInvitationController(router, invitationService)
	postController.ProvidePostController(router, postService)
	userController.ProvideUserController(router, userService, &cacheService)
	adminController.ProvideAdminController(router, adminService)

	router.Run()
}