package main

import (
	"context"
	"encoding/json"
	"errors"
	"graphql-plugin/internal/logger"
	"graphql-plugin/internal/response"
	"net/http"
)

// func main() {}

type registerer string

var pluginName = "graphql-plugin"

var HandlerRegisterer = registerer(pluginName)

func (r registerer) RegisterLogger(v interface{}) {
	logger.SetLogger(v)
}

func (r registerer) RegisterHandlers(
	f func(
		name string,
		handler func(context.Context, map[string]interface{}, http.Handler) (http.Handler, error),
	)) {
	f(string(r), r.registerHandlers)
}

func (r registerer) registerHandlers(_ context.Context, extra map[string]interface{}, h http.Handler) (http.Handler, error) {

	config, ok := extra[pluginName].(map[string]interface{})
	if !ok {
		return h, errors.New("configuration not found")
	}

	// The plugin will look for this path:
	path, _ := config["path"].(string)
	logger.Debugf("The plugin is now hijacking the path %s", path)

	logger.Debugf("registering pluging [%s]", pluginName)
	logger.Infof("[Plugin: %s] Loaded")

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		if req.URL.Path != path {
			h.ServeHTTP(w, req)
			return
		}

		myResponseWriter := response.CreateResponseWriter(w)
		h.ServeHTTP(myResponseWriter, req)

		logger.Debugf("response from backend: %#v", myResponseWriter)
		logger.Debugf("response content: %s", string(myResponseWriter.Read()))

		var jsonResponse map[string]interface{}

		if err := json.Unmarshal(myResponseWriter.Read(), &jsonResponse); err != nil {
			logger.Error(err)
			json.NewEncoder(w).Encode(response.Response{
				Data: map[string]interface{}{
					"Error": err,
				},
			})
		}

		json.NewEncoder(w).Encode(response.Response{
			Data: map[string]interface{}{
				"Status": jsonResponse["status"],
			},
		})

	}), nil

}
