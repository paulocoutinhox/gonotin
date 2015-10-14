<?php

// general
$dataA  = [];
$dataB  = [];

// get data A
$handle = fopen($argv[1], 'r');

if ($handle)
{
    while (($line = fgets($handle)) !== false)
    {
        $dataA[] = $line;
    }

    fclose($handle);
}

// get data B
$handle = fopen($argv[2], 'r');

if ($handle)
{
    while (($line = fgets($handle)) !== false)
    {
        $dataB[] = $line;
    }

    fclose($handle);
}

// functions
function debug($msg)
{
	$msg = trim($msg);
	echo("$msg\n");
}

// logic (1 or 2 for difference performance)
$mode = 1;

if ($mode == 1)
{
	foreach ($dataA as $dataAItem)
	{
		if (!in_array($dataAItem, $dataB))
	    {
		    debug($dataAItem);
	    }
	}
}
else if ($mode == 2)
{
	foreach ($dataA as $dataAItem)
	{
		$exists = false;
	    
	    foreach ($dataB as $dataBItem)
	    {
		    if ($dataAItem == $dataBItem)
		    {
			    $exists = true;
		    }
	    }
	    
	    if (!$exists)
	    {
		    debug($dataAItem);
	    }
	}
}