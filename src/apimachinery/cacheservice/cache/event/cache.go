/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.
 * Copyright (C) 2017-2018 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package event

import (
	"context"
	"net/http"

	"configcenter/src/common"
	"configcenter/src/common/errors"
	"configcenter/src/common/watch"
)

// WatchEvent TODO
func (e *eventCache) WatchEvent(ctx context.Context, h http.Header, opts *watch.WatchEventOptions) (*string,
	errors.CCErrorCoder) {

	resp, err := e.client.Post().
		WithContext(ctx).
		Body(opts).
		SubResourcef("/watch/cache/event").
		WithHeaders(h).
		Do().
		IntoJsonString()

	if err != nil {
		return nil, errors.New(common.CCErrCommHTTPDoRequestFailed, err.Error())
	}

	if err := resp.CCError(); err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

// InnerWatchEvent watch event for inner api
func (e *eventCache) InnerWatchEvent(ctx context.Context, h http.Header, opts *watch.WatchEventOptions) (
	*watch.WatchResp, errors.CCErrorCoder) {

	resp := new(watch.WatchEventResp)
	err := e.client.Post().
		WithContext(ctx).
		Body(opts).
		SubResourcef("/inner/watch/cache/event").
		WithHeaders(h).
		Do().
		Into(resp)

	if err != nil {
		return nil, errors.CCHttpError
	}

	if err := resp.CCError(); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
