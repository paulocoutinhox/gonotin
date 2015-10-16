<?php

// general
ini_set('memory_limit', -1);

$mode = $argv[3];

// functions
function debug($msg)
{
	$msg = trim($msg);
	echo("$msg\n");
}

if ($mode == 1)
{
    // it is generating wrong results
    debug('it is generating wrong results');

    /*
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

    $result = array_diff($dataB, $dataA);

    foreach ($result as $resultItem)
    {
        debug($resultItem);
    }
    */
}
else if ($mode == 2)
{
    // it is slow a lot
    debug('it is slow a lot');

    /*
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

	foreach ($dataA as $dataAItem)
	{
		if (!in_array($dataAItem, $dataB))
	    {
		    debug($dataAItem);
	    }
	}
	*/
}