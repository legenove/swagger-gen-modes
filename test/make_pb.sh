echo "build *** struct go begin"

CUR_DIR=`pwd`
DIR_NAME=`basename $CUR_DIR`
`cd $CUR_DIR`
protoc --proto_path=./out/ --gofast_out=plugins=grpc:./pb/ testPet.proto
#protoc --proto_path=./out/ --gofast_out=./pb/ testPet.proto

echo "build *** struct go end"