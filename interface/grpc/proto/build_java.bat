

@echo on

set batPath=%~dp0
@echo  %batPath%
cd %batPath%
cd ../go_out
set outPath=%cd%
@echo  %outPath%
cd ../../../../../../
set proPath=%cd%
@echo  %proPath%
cd %proPath%

set javagen="/java/github/grpc-java/compiler/build/exe/java_plugin/protoc-gen-grpc-java"

protoc --proto_path=%proPath% --proto_path=%batPath% --grpc-java_out=%batPath% global_def.proto
protoc --proto_path=%proPath% --proto_path=%batPath% --grpc-java_out=%batPath% fs_gw.proto
protoc --plugin=protoc-gen-grpc-java=%javagen% --grpc-java_out="./" global_def.proto
protoc --plugin=protoc-gen-grpc-java=%javagen% --grpc-java_out="./" fs_gw.proto

cd %outPath%


cd %batPath%

