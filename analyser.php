<?php


$file = fopen($argv[1], "r");

$ipush=1;
$ipop=1;

$fnamepush=str_replace("analysis/","analysis/push_",$argv[1]);
$fnamepop=str_replace("analysis/","analysis/pop_",$argv[1]);

$pushfile = fopen($fnamepush, 'w');
$popfile = fopen($fnamepop, 'w');

while(!feof($file)){
    $line = fgets($file);
    
    echo $line;
    if(strpos($line,"push")!==false){
    
    
        $time=explode("took  ",$line)[1];
        $time=str_replace("s","",$time);
        echo $ipush." ".$time;
        fwrite($pushfile,$time);
        $ipush++;
    
    }
    
    if(strpos($line,"pop")!==false){
    
    
        $time=explode("took  ",$line)[1];
        $time=str_replace("s","",$time);
        echo $ipop." ".$time;
        fwrite($popfile,$time);
        $ipop++;
    
    }
    
    
}
fclose($file);
