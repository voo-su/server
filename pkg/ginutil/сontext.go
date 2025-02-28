package ginutil

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"net/http"
	"voo.su/pkg/jwtutil"
	"voo.su/pkg/validator"
)

var MarshalOptions = protojson.MarshalOptions{
	UseProtoNames:   true,
	EmitUnpopulated: true,
}

type Context struct {
	Context *gin.Context
}

func New(ctx *gin.Context) *Context {
	return &Context{ctx}
}

func (c *Context) Ctx() context.Context {
	return c.Context.Request.Context()
}

func (c *Context) Unauthorized(message string) error {
	c.Context.AbortWithStatusJSON(http.StatusUnauthorized, &Response{
		Code:    http.StatusUnauthorized,
		Message: message,
	})

	return nil
}

func (c *Context) Forbidden(message string) error {
	c.Context.AbortWithStatusJSON(http.StatusForbidden, &Response{
		Code:    http.StatusForbidden,
		Message: message,
	})

	return nil
}

func (c *Context) InvalidParams(message any) error {
	resp := &Response{
		Code:    305,
		Message: "Invalid parameters",
	}
	switch msg := message.(type) {
	case error:
		resp.Message = validator.Translate(msg)
	case string:
		resp.Message = msg
	default:
		resp.Message = fmt.Sprintf("%v", msg)
	}

	c.Context.AbortWithStatusJSON(http.StatusBadRequest, resp)

	return nil
}

func (c *Context) Error(message any) error {
	resp := &Response{
		Code:    400,
		Message: "Error",
	}
	switch msg := message.(type) {
	case error:
		resp.Message = msg.Error()
	case string:
		resp.Message = msg
	default:
		resp.Message = fmt.Sprintf("%v", msg)
	}

	c.Context.AbortWithStatusJSON(http.StatusBadRequest, resp)

	return nil
}

func (c *Context) Success(data any, message ...string) error {
	resp := &Response{
		Code:    200,
		Message: "success",
		Data:    data,
	}
	if len(message) > 0 {
		resp.Message = message[0]
	}

	if value, ok := data.(proto.Message); ok {
		bt, _ := MarshalOptions.Marshal(value)
		var body map[string]any
		_ = json.Unmarshal(bt, &body)
		resp.Data = body
	}

	c.Context.AbortWithStatusJSON(http.StatusOK, resp)

	return nil
}

func (c *Context) Raw(value string) error {
	c.Context.Abort()
	c.Context.String(http.StatusOK, value)

	return nil
}

func (c *Context) JwtSession() *jwtutil.JSession {
	data, isOk := c.Context.Get(jwtutil.JWTSession)
	if !isOk {
		return nil
	}

	return data.(*jwtutil.JSession)
}

func (c *Context) UserId() int {
	if session := c.JwtSession(); session != nil {
		return session.Uid
	}

	return 0
}

func (c *Context) UserToken() string {
	if session := c.JwtSession(); session != nil {
		return session.Token
	}

	return ""
}
