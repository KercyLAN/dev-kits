// Copyright 2019 Xusixxxx Author. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hanzi

import (
	"github.com/KercyLAN/dev-kits/components/hanzi/pinyin"
)

// 汉字结构
type hanzi struct {
	char 			string						// 文字
	pinyin 			[]*pinyin.Pinyin			// 拼音信息
	rhyme 			[]string 					// 韵脚（通常韵脚取韵母一位，如 chi : i , li xiang : i ang）
}

func (slf *hanzi) String() string {
	return slf.char
}

func (slf *hanzi) Rhyme() []string {
	return slf.rhyme
}

func (slf *hanzi) Pinyin() []*pinyin.Pinyin {
	return slf.pinyin
}
