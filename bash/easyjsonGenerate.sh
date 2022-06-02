bash easyjsonClear.sh
cd .. || return
find . -regex ".*/microservices/.*[pP]roto.*pb.go" -exec easyjson -all {} \;
find . -regex ".*/structs/.*.go" -exec easyjson -all {} \;
