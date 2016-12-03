rm test/*.meta
rm test/*.data

go build bfqPerfTest.go
echo "start testing 5 bytes workload"
./bfqPerfTest workload/5.byte 500000 > analysis/500000_5bfq.log
php analyser.php analysis/500000_5bfq.log

rm test/*.meta
rm test/*.data
echo "testing 583 bytes workload"
./bfqPerfTest workload/583.byte 500000 > analysis/500000_583bfq.log
php analyser.php analysis/500000_583bfq.log


rm test/*.meta
rm test/*.data
echo "testing 1024 bytes workload"
./bfqPerfTest workload/1024.byte 500000 > analysis/500000_1024bfq.log
php analyser.php analysis/500000_1024bfq.log

