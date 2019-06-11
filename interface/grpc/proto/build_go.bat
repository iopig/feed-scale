

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
protoc --proto_path=%proPath% --proto_path=%batPath% --go_out=plugins=grpc:./ global_def.proto
protoc --proto_path=%proPath% --proto_path=%batPath% --go_out=plugins=grpc:./ fs_gw.proto


cd %outPath%
cd fsapi

go build

cd %batPath%

