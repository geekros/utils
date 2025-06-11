// Copyright 2025 GEEKROS, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package response

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func requestLog(c *gin.Context, status string, response string) {

}

func Success(c *gin.Context, data interface{}) {

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data":    data,
	})

	logJson, _ := json.Marshal(gin.H{"code": 0, "message": "success", "data": data})

	requestLog(c, "success", string(logJson))
}

func Warning(c *gin.Context, code int, message string, data interface{}) {

	if message == "" {
		message = "warning"
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": message,
		"data":    data,
	})

	logJson, _ := json.Marshal(gin.H{"code": code, "message": message, "data": data})

	requestLog(c, "warning", string(logJson))
}

func Error(c *gin.Context, data interface{}) {

	c.JSON(http.StatusOK, gin.H{
		"code":    10000,
		"message": "error",
		"data":    data,
	})

	logJson, _ := json.Marshal(gin.H{"code": 10000, "message": "error", "data": data})

	requestLog(c, "error", string(logJson))
}
