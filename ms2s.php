<?php


$c = file_get_contents($argv[1]);

$lines = explode("\n",$c);

$nlines=[];
foreach($lines as $line){

    if(strpos($line,"m")!==false){
    
        $flt = (float)$line;
        
        $line = ($flt/1000.0)+""; 
    }
    
    $nlines[]=$line;

}


$c2=implode("\n",$nlines);

file_put_contents($argv[1],$c2);
