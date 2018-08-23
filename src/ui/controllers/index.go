// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package controllers

import (
	"github.com/astaxie/beego"
	"github.com/goharbor/harbor/src/ui/config"
)

// IndexController handles request to /
type IndexController struct {
	beego.Controller
}

//Prepare to check if incoming requests should be served
func (ic *IndexController) Prepare() {
	if config.WithAdmiral() {
		ic.Redirect(config.AdmiralEndpoint(), 302)
	}
}

// Get renders the index page
func (ic *IndexController) Get() {
	ic.TplExt = "html"
	ic.TplName = "index.html"
}
