# Copyright 2022 Marcelo Cantos
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

all: test lint

test:
	go test -cover ./...

coverage:
	go test -covermode count -coverprofile=coverage.out ./... && go tool cover -func=coverage.out \
		| perl -ne 's{^'$$(awk '/^module/{print$$2}' go.mod)'/}{}; print unless m{^total:|100\.0%$$}' \
		| sort -rn -k3

lint:
	golangci-lint run --max-same-issues 50
