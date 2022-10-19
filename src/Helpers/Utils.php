<?php

namespace Hightemp\WappSnanimageServer\Helpers;

class Utils
{
    public static function fnSanitiseURL($url, $encode = false) {
        $sNewURL = filter_var(urldecode($url), FILTER_SANITIZE_SPECIAL_CHARS);
        if (! filter_var($sNewURL, FILTER_VALIDATE_URL))
            return false;
        return $encode ? urlencode($url) : $url;
    }
}