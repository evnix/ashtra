rm test/*level.db
go build lsmPerfTest.go
echo "start testing 5 bytes workload"
./lsmPerfTest workload/5.byte 500000 > analysis/500000_5lsm.log
php analyser.php analysis/500000_5lsm.log


rm test/*level.db
echo "testing 583 bytes workload"
./lsmPerfTest workload/583.byte 500000 > analysis/500000_583lsm.log
php analyser.php analysis/500000_583lsm.log


rm test/*level.db
echo "testing 1024 bytes workload"
./lsmPerfTest workload/1024.byte 500000 > analysis/500000_1024lsm.log
php analyser.php analysis/500000_1024lsm.log

