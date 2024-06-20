package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/run-bigpig/skplayer/internal/config"
	"github.com/run-bigpig/skplayer/internal/core"
	"github.com/run-bigpig/skplayer/internal/types"
	"github.com/run-bigpig/skplayer/internal/utils"
	"strconv"
)

func Index(ctx *fiber.Ctx) error {
	return ctx.SendFile("web/index.html")
}

func Class(ctx *fiber.Ctx) error {
	var response []*types.ClassResponse
	s := core.NewSpider(config.Get().DefaultSource.Url)
	class, err := s.Class()
	if err != nil {
		return utils.Fail(ctx, 400, err)
	}
	for _, item := range class.Class {
		response = append(response, &types.ClassResponse{
			Id:   item.TypeId,
			Pid:  item.TypePid,
			Name: item.TypeName,
		})
	}
	return utils.Success(ctx, response)
}

func List(ctx *fiber.Ctx) error {
	var (
		req       types.ListRequest
		listItems []*types.ListItem
		response  types.ListResponse
	)
	if err := ctx.BodyParser(&req); err != nil {
		return utils.Fail(ctx, 400, err)
	}
	s := core.NewSpider(config.Get().DefaultSource.Url)
	list, err := s.DetailList(&types.VodListRequest{
		Page:     req.Page,
		PageSize: req.PageSize,
		TypeId:   req.Class,
	})
	if err != nil {
		return utils.Fail(ctx, 400, err)
	}
	for _, item := range list.List {
		listItems = append(listItems, &types.ListItem{
			Id:       item.VodId,
			Name:     item.VodName,
			Actor:    item.VodActor,
			Director: item.VodDirector,
			Img:      item.VodPic,
		})
	}
	response = types.ListResponse{
		List:  listItems,
		Total: list.Total,
	}
	return utils.Success(ctx, response)
}

func Detail(ctx *fiber.Ctx) error {
	var (
		req      types.DetailRequest
		response types.DetailResponse
	)
	if err := ctx.BodyParser(&req); err != nil {
		return utils.Fail(ctx, 400, err)
	}
	s := core.NewSpider(config.Get().DefaultSource.Url)
	detail, err := s.Detail(&types.VodDetailRequest{
		Ids: strconv.Itoa(req.Id),
	})
	if err != nil {
		return utils.Fail(ctx, 400, err)
	}
	if len(detail.List) == 1 {
		response = types.DetailResponse{
			Id:       detail.List[0].VodId,
			Name:     detail.List[0].VodName,
			Actor:    detail.List[0].VodActor,
			Director: detail.List[0].VodDirector,
			Img:      detail.List[0].VodPic,
			Desc:     utils.RemoveHTMLTags(detail.List[0].VodContent),
			Play:     utils.DealPlayUrl(detail.List[0].VodPlayUrl),
		}
	}
	return utils.Success(ctx, response)
}

func Search(ctx *fiber.Ctx) error {
	var (
		req       types.SearchRequest
		listItems []*types.ListItem
		response  types.ListResponse
	)
	if err := ctx.BodyParser(&req); err != nil {
		return utils.Fail(ctx, 400, err)
	}
	s := core.NewSpider(config.Get().DefaultSource.Url)
	list, err := s.Search(&types.VodSearchRequest{
		Keyword:  req.Keyword,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return utils.Fail(ctx, 400, err)
	}
	for _, item := range list.List {
		listItems = append(listItems, &types.ListItem{
			Id:       item.VodId,
			Name:     item.VodName,
			Actor:    item.VodActor,
			Director: item.VodDirector,
			Img:      item.VodPic,
		})
	}
	response = types.ListResponse{
		List:  listItems,
		Total: list.Total,
	}
	return utils.Success(ctx, response)
}

func GetSourceList(ctx *fiber.Ctx) error {
	return utils.Success(ctx, &types.SourceResponse{
		List:    config.Get().Source,
		Default: config.Get().DefaultSource,
	})
}

func Setting(ctx *fiber.Ctx) error {
	var (
		req types.SettingRequest
	)
	if err := ctx.BodyParser(&req); err != nil {
		return utils.Fail(ctx, 400, err)
	}
	config.SetDefaultSource(req.Source)
	return utils.Success(ctx, "ok")
}
