mkdir -p src/sim_pack  
export GOPATH=$(pwd)	// pack dir  
cd src/sim_pack  
go install sim_pack  
cd ../../  
ls						// Check pkg dir  
go run use_pack.go  
