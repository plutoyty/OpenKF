// Copyright © 2023 OpenIM open source community. All rights reserved.
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

package requestparams

// RegisterBotParams register params for bot.
type RegisterBotParams struct {
	BotAddr     string `json:"bot_addr"     binding:"required"`
	BotPort     int    `json:"bot_port"     binding:"required"`
	BotToken    string `json:"bot_token"    binding:"required"`
	Nickname    string `json:"nickname"     binding:"required"`
	Avatar      string `json:"avatar"       binding:"required"`
	CommunityId uint   `json:"community_id" binding:"required"`
}
