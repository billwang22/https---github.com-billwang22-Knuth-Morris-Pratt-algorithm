<?php
set_time_limit(0);
error_reporting(E_ALL);
/*
 * the Knuth-Morris-Pratt string searching algorithms. 
 */

/*
 * find the max length common element in tow array
 * 
 */
function findBigCom($arrayPrefix,$arraysuffix,$len){
	$com = null;
	$maxLen = 0;
	for($i=0;$i<$len;$i++){
		for($j=0;$j<$len;$j++){
			if($arrayPrefix[$i] == $arraysuffix[$j]){
				$com = $arrayPrefix[$i];
				if(strlen($com) > $maxLen){
					$maxLen = strlen($com);
				}
			}
		}
	}
	return $maxLen;
}
/*
 * generate the part match table for specific word
 */
function generatePartMatch($pattenWord){
	$partMatch = array();
	$leng = strlen($pattenWord);
	for($i=1;$i<=$leng;$i++){
		$prefix = array(); // prefix set 
		$suffix = array(); //suffix set
		$subString = substr($pattenWord,0,$i);
		for($j=1;$j<$i;$j++){
			$prefix[]= substr($subString,0,$j);
			$suffix[]= substr($subString,$j);
		}
		$partMatch[$i] = findBigCom($prefix,$suffix,$i-1);
	}
	return $partMatch;
}

/*
 * BMP search
 */
function BMPSearch($bigString,$smallString){
	$foundFlag = false;
	$partMatchTable = generatePartMatch($smallString);
	$i = 0 ; $j = 0;
	$matchLen = 0;
	$bigStringLen = strlen($bigString);
	while ($i<=$bigStringLen){
		$index = substr($bigString,$i,1);
		$cword = substr($smallString,$j,1);
		if($index == $cword){
			$matchLen +=1;
			$j++;
			$i++;
		}else {
			if($matchLen != 0){ 
				$noMatchPos = $j;
				$j = $j - ($matchLen - $partMatchTable[$noMatchPos]);
				$matchLen = $j; // jump to next match position 
			}else {
				$i++;   
				$j = 0;
				// jump to next match position 
			}
		}
		
		if($matchLen == strlen($smallString)){
			$start = strlen($smallString);
			$end   = $i;
			$start = $i-$start;
			echo "found...from  ...".$start."to   ".$end;
			echo substr($bigString,$start,$end);
			$foundFlag = true;
			
		}
	}
	if(!$foundFlag){
		echo "not found";
	}

}


