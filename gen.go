// Copyright 2018 Hajime Hoshi
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

package sourcehansans

//go:generate sh ./subset.sh
//go:generate mkdir -p scregular
//go:generate mkdir -p tcregular
//go:generate file2byteslice -input scregular.subset.ttf -output scregular/ttf.go -package scregular -var CompressedTTF -compress
//go:generate file2byteslice -input tcregular.subset.ttf -output tcregular/ttf.go -package tcregular -var CompressedTTF -compress
//go:generate gofmt -s -w .
