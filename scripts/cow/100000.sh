go build cowPerfTest.go
echo "start testing 5 bytes workload"
./cowPerfTest workload/5.byte 100000 > analysis/100000_5cow.log
php analyser.php analysis/100000_5cow.log


rm test/*.db
echo "testing 583 bytes workload"
./cowPerfTest workload/583.byte 100000 > analysis/100000_583cow.log
php analyser.php analysis/100000_583cow.log


rm test/*.db
echo "testing 1024 bytes workload"
./cowPerfTest workload/1024.byte 100000 > analysis/100000_1024cow.log
php analyser.php analysis/100000_1024cow.log

