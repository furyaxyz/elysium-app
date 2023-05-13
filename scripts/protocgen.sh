#!/usr/bin/env bash

set -eo pipefail

protoc_gen_gocosmos() {
  if ! grep "github.com/gogo/protobuf => github.com/regen-network/protobuf" go.mod &>/dev/null ; then
    echo -e "\tPlease run this command from somewhere inside the cosmos-sdk folder."
    return 1
  fi
}

protoc_gen_gocosmos

cd proto
proto_dirs=$(find . -path ./cosmos -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do
  for file in $(find "${dir}" -maxdepth 1 -name '*.proto'); do
    if grep "option go_package" $file &> /dev/null ; then
      buf generate --template buf.gen.gogo.yaml $file
    fi
  done
done

cd ..

# temporary import hack to use cosmos-sdk implementation of Any type.
# check https://github.com/furyaxyz/elysium-app/issues/507 for more information.
sed -i 's/types "github.com\/furyaxyz\/elysium-app\/codec\/types"/types "github.com\/cosmos\/cosmos-sdk\/codec\/types"/g' \
 github.com/furyaxyz/elysium-app/x/qgb/types/query.pb.go

# move proto files to the right places
cp -r github.com/furyaxyz/elysium-app/* ./
rm -rf github.com
