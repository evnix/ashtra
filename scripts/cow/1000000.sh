# now for 1000,000 ---------------------------------------------------------
echo "1000,000 testing 5 bytes workload"
./cowPerfTest workload/5.byte 1000000 > analysis/1000000_5cow.log
php analyser.php analysis/1000000_5cow.log


rm test/*.db
echo "testing 583 bytes workload"
./cowPerfTest workload/583.byte 1000000 > analysis/1000000_583cow.log
php analyser.php analysis/1000000_583cow.log


rm test/*.db
echo "testing 1024 bytes workload"
./cowPerfTest workload/1024.byte 1000000 > analysis/1000000_1024cow.log
php analyser.php analysis/1000000_1024cow.log

