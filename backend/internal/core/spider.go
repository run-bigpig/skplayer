package core

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/run-bigpig/skplayer/internal/constant"
	"github.com/run-bigpig/skplayer/internal/types"
)

type Spider struct {
	api    string
	header map[string]string
}

type Unmarshaler interface {
	Unmarshal([]byte) error
}

func NewSpider(api string) *Spider {
	return &Spider{
		api:    api,
		header: map[string]string{"User-Agent": constant.UserAgent, "Content-Type": fiber.MIMEApplicationJSONCharsetUTF8},
	}
}

// Class 获取分类
func (s *Spider) Class() (*types.VodListResponse, error) {
	var (
		response types.VodListResponse
	)
	err := s.send(fiber.MethodGet, s.api, nil, &response)
	if err != nil {
		return nil, err
	}
	return &response, err
}

// List 获取列表
func (s *Spider) List(params *types.VodListRequest) (*types.VodListResponse, error) {
	var (
		response types.VodListResponse
	)
	err := s.send(fiber.MethodGet, s.api, []byte(fmt.Sprintf("ac=%s&pg=%d&pagesize=%d&t=%d", "list", params.Page, params.PageSize, params.TypeId)), &response)
	if err != nil {
		return nil, err
	}
	return &response, err
}

// DetailList 获取详情列表
func (s *Spider) DetailList(params *types.VodListRequest) (*types.VodDetailResponse, error) {
	var (
		response types.VodDetailResponse
	)
	if params.Page == 0 {
		params.Page = 1
	}
	if params.PageSize == 0 {
		params.PageSize = 24
	}
	err := s.send(fiber.MethodGet, s.api, []byte(fmt.Sprintf("ac=%s&pg=%d&pagesize=%d&t=%d", "videolist", params.Page, params.PageSize, params.TypeId)), &response)
	if err != nil {
		return nil, err
	}
	return &response, err
}

// Detail 获取详情
func (s *Spider) Detail(params *types.VodDetailRequest) (*types.VodDetailResponse, error) {
	var (
		response types.VodDetailResponse
	)
	err := s.send(fiber.MethodGet, s.api, []byte(fmt.Sprintf("ac=%s&ids=%s", "videolist", params.Ids)), &response)
	if err != nil {
		return nil, err
	}
	return &response, err
}

// Search 搜索
func (s *Spider) Search(params *types.VodSearchRequest) (*types.VodDetailResponse, error) {
	var (
		response types.VodDetailResponse
	)
	err := s.send(fiber.MethodGet, s.api, []byte(fmt.Sprintf("ac=%s&wd=%s&pg=%d&pagesize=%d", "videolist", params.Keyword, params.Page, params.PageSize)), &response)
	if err != nil {
		return nil, err
	}
	return &response, err
}

func (s *Spider) send(method string, url string, params []byte, response Unmarshaler) error {
	agent := fiber.AcquireAgent()
	for key, v := range s.header {
		agent.Set(key, v)
	}
	req := agent.Request()
	req.Header.SetMethod(method)
	req.SetRequestURI(url)
	switch method {
	case fiber.MethodGet:
		agent.QueryStringBytes(params)
	case fiber.MethodPost:
		agent.Body(params)
	}
	if err := agent.Parse(); err != nil {
		return err
	}
	code, body, errs := agent.Bytes()
	if errs != nil {
		return fmt.Errorf("request error: %v", errs)
	}
	if code != fiber.StatusOK {
		return fmt.Errorf("request error: %d", code)
	}
	return response.Unmarshal(body)
}
