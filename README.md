# Data-Structure
Data structure written in Golang

# How to Interconnect Go Language with C Language
export GOPATH="($TOPDIR)/Data-Structure/cgo"  

($TOPDIR) is your computer's directory location.  
And Data-Sturcture is our github repo.  

After write that instruction, then follow below instructions.  

cd ($TOPDIR)/Data-Structure/cgo/src/tut  
gcc -c test.c  
go install  
cd ../../bin  
./tut  
