echo -n >  analysis/5cow.log
echo -n >  analysis/583cow.log
echo -n >  analysis/1024cow.log
rm test/*.db


go build cowPerfTest.go
echo "testing 5 bytes workload"
./cowPerfTest workload/5.byte > analysis/5cow.log
php analyser.php analysis/5cow.log
php ms2s.php analysis/push_5cow.log
php ms2s.php analysis/pop_5cow.log

rm test/*.db
echo "testing 583 bytes workload"
./cowPerfTest workload/583.byte > analysis/583cow.log
php analyser.php analysis/583cow.log
php ms2s.php analysis/push_583cow.log
php ms2s.php analysis/pop_583cow.log

rm test/*.db
echo "testing 1024 bytes workload"
./cowPerfTest workload/1024.byte > analysis/1024cow.log
php analyser.php analysis/1024cow.log
php ms2s.php analysis/push_1024cow.log
php ms2s.php analysis/pop_1024cow.log
