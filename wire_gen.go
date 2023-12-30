// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/google/wire"
	"peken-be/app"
	"peken-be/controller"
	"peken-be/helper"
	"peken-be/repository"
	"peken-be/service"
)

// Injectors from injector.go:

func InitializedServer() *gin.Engine {
	db := app.ConnectToDb()
	roleRepositoryImpl := repository.NewRoleRepository(db)
	userRepositoryImpl := repository.NewUserRepository(db)
	passwordUtilsImpl := helper.NewPasswordUtils()
	validate := validator.New()
	userServiceImpl := service.NewUserService(roleRepositoryImpl, userRepositoryImpl, passwordUtilsImpl, validate)
	userControllerImpl := controller.NewUserController(userServiceImpl)
	loginServiceImpl := service.NewLoginService(userRepositoryImpl, passwordUtilsImpl, validate)
	loginControllerImpl := controller.NewLoginController(loginServiceImpl)
	memberRepositoryImpl := repository.NewMemberRepository(db)
	memberServiceImpl := service.NewMemberService(memberRepositoryImpl, validate)
	memberControllerImpl := controller.NewMemberController(memberServiceImpl)
	engine := app.InitRoute(userControllerImpl, loginControllerImpl, userRepositoryImpl, memberControllerImpl)
	return engine
}

// injector.go:

var roleSet = wire.NewSet(repository.NewRoleRepository, wire.Bind(new(repository.RoleRepository), new(*repository.RoleRepositoryImpl)))

var userSet = wire.NewSet(repository.NewUserRepository, wire.Bind(new(repository.UserRepository), new(*repository.UserRepositoryImpl)), service.NewUserService, wire.Bind(new(service.UserService), new(*service.UserServiceImpl)), controller.NewUserController, wire.Bind(new(controller.UserController), new(*controller.UserControllerImpl)))

var loginSet = wire.NewSet(service.NewLoginService, wire.Bind(new(service.LoginService), new(*service.LoginServiceImpl)), controller.NewLoginController, wire.Bind(new(controller.LoginController), new(*controller.LoginControllerImpl)))

var passGenSet = wire.NewSet(helper.NewPasswordUtils, wire.Bind(new(helper.PasswordUtils), new(*helper.PasswordUtilsImpl)))

var memberSet = wire.NewSet(repository.NewMemberRepository, wire.Bind(new(repository.MemberRepository), new(*repository.MemberRepositoryImpl)), service.NewMemberService, wire.Bind(new(service.MemberService), new(*service.MemberServiceImpl)), controller.NewMemberController, wire.Bind(new(controller.MemberController), new(*controller.MemberControllerImpl)))
