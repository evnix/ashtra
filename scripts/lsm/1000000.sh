rm test/*level.db
go build lsmPerfTest.go
echo "start testing 5 bytes workload"
./lsmPerfTest workload/5.byte 1000000 > analysis/1000000_5lsm.log
php analyser.php analysis/1000000_5lsm.log


rm test/*level.db
echo "testing 583 bytes workload"
./lsmPerfTest workload/583.byte 1000000 > analysis/1000000_583lsm.log
php analyser.php analysis/1000000_583lsm.log


rm test/*level.db
echo "testing 1024 bytes workload"
./lsmPerfTest workload/1024.byte 1000000 > analysis/1000000_1024lsm.log
php analyser.php analysis/1000000_1024lsm.log


