echo -n >  analysis/5cow.log
echo -n >  analysis/583cow.log
echo -n >  analysis/1024cow.log
rm test/*.db


go build cowPerfTest.go
echo "testing 5 bytes workload"
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




# now for 500,000 ---------------------------------------------------------
echo "testing 5 bytes workload"
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



# now for 1000,000 ---------------------------------------------------------
echo "testing 5 bytes workload"
./cowPerfTest workload/5.byte 500000 > analysis/1000000_5cow.log
php analyser.php analysis/1000000_5cow.log


rm test/*.db
echo "testing 583 bytes workload"
./cowPerfTest workload/583.byte 500000 > analysis/1000000_583cow.log
php analyser.php analysis/1000000_583cow.log


rm test/*.db
echo "testing 1024 bytes workload"
./cowPerfTest workload/1024.byte 500000 > analysis/1000000_1024cow.log
php analyser.php analysis/1000000_1024cow.log

