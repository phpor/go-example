<?php

$gzip = "1";  // from form
$validator = getenv("validator"); // from header X-Validator

$throwLogs = file_get_contents("php://stdin"); // from form

// printf("throwLogs: %s\n\n", $throwLogs);

if ($validator) {
    $arr = parseServerTime($validator);
    if (!$arr) {
        echo "parse validator fail";
        exit();
    }
//     print_r($arr);
// var_dump($arr);
    $throwLogs = xor_encode($throwLogs, $arr["hour"]);

//     echo $throwLogs;exit;
//     printf("throwLogs2: %s\n\n", $throwLogs);

}

if ($gzip == "1") {
    $throwLogs = gzinflate(substr($throwLogs, 10, -8));
}

printf("throwLogs3: %s\n\n", $throwLogs);
// $throwLogs = gzinflate($throwLogs);


function xor_encode($data, $key) {
    $dataLen = strlen($data);
    $keyLen = strlen($key);

    for ($text = '', $i = 0; $i < $dataLen; $i++) {
        $text .= $data{$i} ^ $key{$i % $keyLen};
    }

    return $text;
}

function parseServerTime($validator) {
    $secret = "obiewelibom_anis";
    echo $validator;
    $text = openssl_decrypt($validator, "aes-128-ecb", $secret);
    if (!$text) {
        return "";
    }
//     printf("text: %s\n", $text);
    return [
        "timestamp" => substr($text, 0, 10),
        "hour" => substr($text, -10),
    ];
}
