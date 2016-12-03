
# now for 500,000 ---------------------------------------------------------
echo "500,000 testing 5 bytes workload"
./cowPerfTest workload/5.byte 500000 > analysis/500000_5cow.log
php analyser.php analysis/500000_5cow.log


rm test/*.db
echo "testing 583 bytes workload"
./cowPerfTest workload/583.byte 500000 > analysis/500000_583cow.log
php analyser.php analysis/500000_583cow.log


rm test/*.db
echo "testing 1024 bytes workload"
./cowPerfTest workload/1024.byte 500000 > analysis/500000_1024cow.log
php analyser.php analysis/500000_1024cow.log

