#!/bin/bash

#Copyright © 2021 Aurelio Calegari
#
#Licensed under the Apache License, Version 2.0 (the "License");
#you may not use this file except in compliance with the License.
#You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
#Unless required by applicable law or agreed to in writing, software
#distributed under the License is distributed on an "AS IS" BASIS,
#WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#See the License for the specific language governing permissions and
#limitations under the License.

CURV=$(git describe --tags --abbrev=0)
IFS='.' read -ra VR <<< "$CURV"
INC=`expr ${VR[2]} + 1`
FV="${VR[0]}.${VR[1]}.$INC"
git tag ${FV} && git push origin ${FV}
GOPROXY=proxy.golang.org go list -m github.com/aurc/plist@${FV}